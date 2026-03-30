// Code generated from Apple documentation. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// NSModalResponse values.
const (

	// AlertFirstButtonReturn is the user clicked the first (rightmost) button on the dialog or sheet.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSApplication/ModalResponse/alertFirstButtonReturn
	AlertFirstButtonReturn NSModalResponse = 1000

	// AlertSecondButtonReturn is the user clicked the second button from the right edge of the dialog or sheet.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSApplication/ModalResponse/alertSecondButtonReturn
	AlertSecondButtonReturn NSModalResponse = 1001

	// AlertThirdButtonReturn is the user clicked the third button from the right edge of the dialog or sheet.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSApplication/ModalResponse/alertThirdButtonReturn
	AlertThirdButtonReturn NSModalResponse = 1002
)

// float64 values.
const (

	// EventDurationForever is the longest time duration possible.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSEvent/foreverDuration
	EventDurationForever float64 = 1.7976931348623157

	// SquareStatusItemLength is a status item length that is equal to the status bar’s thickness.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSStatusItem/squareLength
	SquareStatusItemLength float64 = -2.0

	//
	// See: https://developer.apple.com/documentation/AppKit/NSStackView/useDefaultSpacing
	StackViewSpacingUseDefault float64 = 3.40282347

	// VariableStatusItemLength is a status item length that dynamically adjusts to the width of its contents.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSStatusItem/variableLength
	VariableStatusItemLength float64 = -1.0
)

// NSWindowLevel values.
const (

	// FloatingWindowLevel is useful for floating palettes.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWindow/Level-swift.struct/floating
	FloatingWindowLevel NSWindowLevel = 3

	// MainMenuWindowLevel is reserved for the application’s main menu.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWindow/Level-swift.struct/mainMenu
	MainMenuWindowLevel NSWindowLevel = 24

	// ModalPanelWindowLevel is the level for a modal panel.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWindow/Level-swift.struct/modalPanel
	ModalPanelWindowLevel NSWindowLevel = 8

	// NormalWindowLevel is the default level for [NSWindow] objects.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWindow/Level-swift.struct/normal
	NormalWindowLevel NSWindowLevel = 0

	// PopUpMenuWindowLevel is the level for a pop-up menu.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWindow/Level-swift.struct/popUpMenu
	PopUpMenuWindowLevel NSWindowLevel = 101

	// ScreenSaverWindowLevel is the level for a screen saver.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWindow/Level-swift.struct/screenSaver
	ScreenSaverWindowLevel NSWindowLevel = 1000

	// StatusWindowLevel is the level for a status window.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWindow/Level-swift.struct/statusBar
	StatusWindowLevel NSWindowLevel = 25

	// SubmenuWindowLevel is reserved for submenus. Synonymous with [NSTornOffMenuWindowLevel], which is preferred.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWindow/Level-swift.struct/submenu
	SubmenuWindowLevel NSWindowLevel = 3

	// TornOffMenuWindowLevel is the level for a torn-off menu. Synonymous with [NSSubmenuWindowLevel].
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWindow/Level-swift.struct/tornOffMenu
	TornOffMenuWindowLevel NSWindowLevel = 3
)

// int values.
const (

	// SearchFieldClearRecentsMenuItemTag is the menu item for clearing the current set of recent string searches in the menu.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSearchField/clearRecentsMenuItemTag
	SearchFieldClearRecentsMenuItemTag int = 1002

	// SearchFieldNoRecentsMenuItemTag is the menu item that describes a lack of recent search strings.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSearchField/noRecentsMenuItemTag
	SearchFieldNoRecentsMenuItemTag int = 1003

	// SearchFieldRecentsMenuItemTag is the location of recent search strings in the “recents” menu group.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSearchField/recentsMenuItemTag
	SearchFieldRecentsMenuItemTag int = 1001

	// SearchFieldRecentsTitleMenuItemTag is the menu item that provides the title of the menu group for recent search strings.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSearchField/recentsTitleMenuItemTag
	SearchFieldRecentsTitleMenuItemTag int = 1000
)

var (
	// See: https://developer.apple.com/documentation/AppKit/NSAbortModalException
	AbortModalException foundation.NSExceptionName
	// See: https://developer.apple.com/documentation/AppKit/NSAbortPrintingException
	AbortPrintingException foundation.NSExceptionName
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityAnnotationTextAttribute
	AccessibilityAnnotationTextAttribute foundation.NSAttributedStringKey
	// AccessibilityAutocorrectedTextAttribute is autocorrected text ([NSNumber] as a Boolean value).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityAutocorrectedTextAttribute
	AccessibilityAutocorrectedTextAttribute foundation.NSAttributedStringKey
	// AccessibilityBackgroundColorTextAttribute is text background color ([CGColorRef]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityBackgroundColorTextAttribute
	AccessibilityBackgroundColorTextAttribute foundation.NSAttributedStringKey
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomTextAttribute
	AccessibilityCustomTextAttribute foundation.NSAttributedStringKey
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityFontBoldAttribute
	AccessibilityFontBoldAttribute foundation.NSAttributedStringKey
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityFontItalicAttribute
	AccessibilityFontItalicAttribute foundation.NSAttributedStringKey
	// AccessibilityFontTextAttribute is font keys ([NSDictionary]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityFontTextAttribute
	AccessibilityFontTextAttribute foundation.NSAttributedStringKey
	// AccessibilityForegroundColorTextAttribute is text foreground color ([CGColorRef]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityForegroundColorTextAttribute
	AccessibilityForegroundColorTextAttribute foundation.NSAttributedStringKey
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityLanguageTextAttribute
	AccessibilityLanguageTextAttribute foundation.NSAttributedStringKey
	// AccessibilityLinkTextAttribute is text link (`id`).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityLinkTextAttribute
	AccessibilityLinkTextAttribute foundation.NSAttributedStringKey
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityListItemIndexTextAttribute
	AccessibilityListItemIndexTextAttribute foundation.NSAttributedStringKey
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityListItemLevelTextAttribute
	AccessibilityListItemLevelTextAttribute foundation.NSAttributedStringKey
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityListItemPrefixTextAttribute
	AccessibilityListItemPrefixTextAttribute foundation.NSAttributedStringKey
	// AccessibilityMarkedMisspelledTextAttribute is misspelled text that is visibly marked as misspelled ([NSNumber] as a Boolean value). If you’re implementing a custom text-editing app, use [NSAccessibilityMarkedMisspelledTextAttribute] to ensure that VoiceOver properly identifies misspelled text to users.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityMarkedMisspelledTextAttribute
	AccessibilityMarkedMisspelledTextAttribute foundation.NSAttributedStringKey
	// AccessibilityMisspelledTextAttribute is misspelled text that isn’t necessarily visibly marked as misspelled ([NSNumber] as a Boolean value).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityMisspelledTextAttribute
	AccessibilityMisspelledTextAttribute foundation.NSAttributedStringKey
	// AccessibilityShadowTextAttribute is text shadow ([NSNumber] as a Boolean value).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityShadowTextAttribute
	AccessibilityShadowTextAttribute foundation.NSAttributedStringKey
	// AccessibilityStrikethroughColorTextAttribute is text strikethrough color ([CGColorRef]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityStrikethroughColorTextAttribute
	AccessibilityStrikethroughColorTextAttribute foundation.NSAttributedStringKey
	// AccessibilityStrikethroughTextAttribute is text strikethrough ([NSNumber] as a Boolean value).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityStrikethroughTextAttribute
	AccessibilityStrikethroughTextAttribute foundation.NSAttributedStringKey
	// AccessibilitySuperscriptTextAttribute is text superscript style ([NSNumber]). Values > 0 are superscript; values < 0 are subscript.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySuperscriptTextAttribute
	AccessibilitySuperscriptTextAttribute foundation.NSAttributedStringKey
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityTextAlignmentAttribute
	AccessibilityTextAlignmentAttribute foundation.NSAttributedStringKey
	// AccessibilityUnderlineColorTextAttribute is text underline color ([CGColorRef]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityUnderlineColorTextAttribute
	AccessibilityUnderlineColorTextAttribute foundation.NSAttributedStringKey
	// AccessibilityUnderlineTextAttribute is text underline style ([NSNumber]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityUnderlineTextAttribute
	AccessibilityUnderlineTextAttribute foundation.NSAttributedStringKey
	// AdaptiveImageGlyphAttributeName is the adaptive image glyph for the text.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAdaptiveImageGlyphAttributeName
	AdaptiveImageGlyphAttributeName foundation.NSAttributedStringKey
	// AnimationProgressMarkNotification is posted when the current progress of a running animation reaches one of its progress marks.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAnimation/progressMarkNotification
	AnimationProgressMarkNotification foundation.NSNotificationName
	// AntialiasThresholdChangedNotification is posted after the threshold for antialiasing changes.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSFont/antialiasThresholdChangedNotification
	AntialiasThresholdChangedNotification foundation.NSNotificationName
	// See: https://developer.apple.com/documentation/AppKit/NSAppKitIgnoredException
	AppKitIgnoredException foundation.NSExceptionName
	// See: https://developer.apple.com/documentation/AppKit/NSAppKitVirtualMemoryException
	AppKitVirtualMemoryException foundation.NSExceptionName
	// ApplicationDidBecomeActiveNotification is posted immediately after the app becomes active.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSApplication/didBecomeActiveNotification
	ApplicationDidBecomeActiveNotification foundation.NSNotificationName
	// ApplicationDidChangeOcclusionStateNotification is posted when the app’s occlusion state changes.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSApplication/didChangeOcclusionStateNotification
	ApplicationDidChangeOcclusionStateNotification foundation.NSNotificationName
	// ApplicationDidChangeScreenParametersNotification is posted when the configuration of the displays attached to the computer is changed.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSApplication/didChangeScreenParametersNotification
	ApplicationDidChangeScreenParametersNotification foundation.NSNotificationName
	// ApplicationDidFinishLaunchingNotification is posted at the end of the [finishLaunching()] method to indicate that the app has completed launching and is ready to run.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSApplication/didFinishLaunchingNotification
	ApplicationDidFinishLaunchingNotification foundation.NSNotificationName
	// ApplicationDidFinishRestoringWindowsNotification is posted when the app has finished restoring windows.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSApplication/didFinishRestoringWindowsNotification
	ApplicationDidFinishRestoringWindowsNotification foundation.NSNotificationName
	// ApplicationDidHideNotification is posted at the end of the [hide(_:)] method to indicate that the app is now hidden.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSApplication/didHideNotification
	ApplicationDidHideNotification foundation.NSNotificationName
	// ApplicationDidResignActiveNotification is posted immediately after the app gives up its active status to another app.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSApplication/didResignActiveNotification
	ApplicationDidResignActiveNotification foundation.NSNotificationName
	// ApplicationDidUnhideNotification is posted at the end of the [unhideWithoutActivation()] method to indicate that the app is now visible.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSApplication/didUnhideNotification
	ApplicationDidUnhideNotification foundation.NSNotificationName
	// ApplicationDidUpdateNotification is posted at the end of the [updateWindows()] method to indicate that the app has finished updating its windows.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSApplication/didUpdateNotification
	ApplicationDidUpdateNotification foundation.NSNotificationName
	// See: https://developer.apple.com/documentation/AppKit/NSApplicationProtectedDataDidBecomeAvailableNotification
	ApplicationProtectedDataDidBecomeAvailableNotification foundation.NSNotificationName
	// See: https://developer.apple.com/documentation/AppKit/NSApplicationProtectedDataWillBecomeUnavailableNotification
	ApplicationProtectedDataWillBecomeUnavailableNotification foundation.NSNotificationName
	// See: https://developer.apple.com/documentation/AppKit/NSApplicationShouldBeginSuppressingHighDynamicRangeContentNotification
	ApplicationShouldBeginSuppressingHighDynamicRangeContentNotification foundation.NSNotificationName
	// See: https://developer.apple.com/documentation/AppKit/NSApplicationShouldEndSuppressingHighDynamicRangeContentNotification
	ApplicationShouldEndSuppressingHighDynamicRangeContentNotification foundation.NSNotificationName
	// ApplicationWillBecomeActiveNotification is posted immediately before the app becomes active.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSApplication/willBecomeActiveNotification
	ApplicationWillBecomeActiveNotification foundation.NSNotificationName
	// ApplicationWillFinishLaunchingNotification is posted at the start of the [finishLaunching()] method to indicate that the app has completed its initialization process and is about to finish launching.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSApplication/willFinishLaunchingNotification
	ApplicationWillFinishLaunchingNotification foundation.NSNotificationName
	// ApplicationWillHideNotification is posted at the start of the [hide(_:)] method to indicate that the app is about to be hidden.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSApplication/willHideNotification
	ApplicationWillHideNotification foundation.NSNotificationName
	// ApplicationWillResignActiveNotification is posted immediately before the app gives up its active status to another app.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSApplication/willResignActiveNotification
	ApplicationWillResignActiveNotification foundation.NSNotificationName
	// ApplicationWillTerminateNotification is sends a notification to terminate the app.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSApplication/willTerminateNotification
	ApplicationWillTerminateNotification foundation.NSNotificationName
	// ApplicationWillUnhideNotification is posted at the start of the [unhideWithoutActivation()] method to indicate that the app is about to become visible.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSApplication/willUnhideNotification
	ApplicationWillUnhideNotification foundation.NSNotificationName
	// ApplicationWillUpdateNotification is posted at the start of the [updateWindows()] method to indicate that the app is about to update its windows.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSApplication/willUpdateNotification
	ApplicationWillUpdateNotification foundation.NSNotificationName
	// AttachmentAttributeName is the attachment for the text.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAttachmentAttributeName
	AttachmentAttributeName foundation.NSAttributedStringKey
	// BackgroundColorAttributeName is the color of the background behind the text.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBackgroundColorAttributeName
	BackgroundColorAttributeName foundation.NSAttributedStringKey
	// See: https://developer.apple.com/documentation/AppKit/NSBadBitmapParametersException
	BadBitmapParametersException foundation.NSExceptionName
	// See: https://developer.apple.com/documentation/AppKit/NSBadComparisonException
	BadComparisonException foundation.NSExceptionName
	// See: https://developer.apple.com/documentation/AppKit/NSBadRTFColorTableException
	BadRTFColorTableException foundation.NSExceptionName
	// See: https://developer.apple.com/documentation/AppKit/NSBadRTFDirectiveException
	BadRTFDirectiveException foundation.NSExceptionName
	// See: https://developer.apple.com/documentation/AppKit/NSBadRTFFontTableException
	BadRTFFontTableException foundation.NSExceptionName
	// See: https://developer.apple.com/documentation/AppKit/NSBadRTFStyleSheetException
	BadRTFStyleSheetException foundation.NSExceptionName
	// BaselineOffsetAttributeName is the vertical offset for the position of the text.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBaselineOffsetAttributeName
	BaselineOffsetAttributeName foundation.NSAttributedStringKey
	// BrowserColumnConfigurationDidChangeNotification is notifies the delegate when the width of a browser column has changed.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBrowser/columnConfigurationDidChangeNotification
	BrowserColumnConfigurationDidChangeNotification foundation.NSNotificationName
	// See: https://developer.apple.com/documentation/AppKit/NSBrowserIllegalDelegateException
	BrowserIllegalDelegateException foundation.NSExceptionName
	// ColorListDidChangeNotification is posted whenever a color list changes.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSColorList/didChangeNotification
	ColorListDidChangeNotification foundation.NSNotificationName
	// See: https://developer.apple.com/documentation/AppKit/NSColorListIOException
	ColorListIOException foundation.NSExceptionName
	// See: https://developer.apple.com/documentation/AppKit/NSColorListNotEditableException
	ColorListNotEditableException foundation.NSExceptionName
	// ColorPanelColorDidChangeNotification is posted when the color of the [NSColorPanel] is set, as when [NSColorPanel] is invoked.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSColorPanel/colorDidChangeNotification
	ColorPanelColorDidChangeNotification foundation.NSNotificationName
	// ComboBoxSelectionDidChangeNotification is posted after the pop-up list selection of the [NSComboBox] changes.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSComboBox/selectionDidChangeNotification
	ComboBoxSelectionDidChangeNotification foundation.NSNotificationName
	// ComboBoxSelectionIsChangingNotification is posted whenever the pop-up list selection of the [NSComboBox] is changing.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSComboBox/selectionIsChangingNotification
	ComboBoxSelectionIsChangingNotification foundation.NSNotificationName
	// ComboBoxWillDismissNotification is posted whenever the pop-up list of the [NSComboBox] is about to be dismissed.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSComboBox/willDismissNotification
	ComboBoxWillDismissNotification foundation.NSNotificationName
	// ComboBoxWillPopUpNotification is posted whenever the pop-up list of the [NSComboBox] is going to be displayed.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSComboBox/willPopUpNotification
	ComboBoxWillPopUpNotification foundation.NSNotificationName
	// ContextHelpModeDidActivateNotification is posted when the application enters context-sensitive help mode. This typically happens when the user holds down the Help key.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSHelpManager/contextHelpModeDidActivateNotification
	ContextHelpModeDidActivateNotification foundation.NSNotificationName
	// ContextHelpModeDidDeactivateNotification is posted when the application exits context-sensitive help mode. This happens when the user clicks the mouse button while the cursor is anywhere on the screen after displaying a context-sensitive help topic.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSHelpManager/contextHelpModeDidDeactivateNotification
	ContextHelpModeDidDeactivateNotification foundation.NSNotificationName
	// ControlTextDidBeginEditingNotification is sent when a control with editable cells begins an edit session.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSControl/textDidBeginEditingNotification
	ControlTextDidBeginEditingNotification foundation.NSNotificationName
	// ControlTextDidChangeNotification is sent when the text in the receiving control changes.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSControl/textDidChangeNotification
	ControlTextDidChangeNotification foundation.NSNotificationName
	// ControlTextDidEndEditingNotification is sent when a control with editable cells ends an editing session.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSControl/textDidEndEditingNotification
	ControlTextDidEndEditingNotification foundation.NSNotificationName
	// CursorAttributeName is the cursor object.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSCursorAttributeName
	CursorAttributeName foundation.NSAttributedStringKey
	// See: https://developer.apple.com/documentation/AppKit/NSDraggingException
	DraggingException foundation.NSExceptionName
	// EventTrackingRunLoopMode is the mode set when tracking events modally, such as a mouse-dragging loop.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSEventTrackingRunLoopMode
	EventTrackingRunLoopMode foundation.NSRunLoopMode
	// ExpansionAttributeName is the expansion factor of the text.
	//
	// Deprecated: Deprecated since macOS 26.2.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSExpansionAttributeName
	ExpansionAttributeName foundation.NSAttributedStringKey
	// FontAttributeName is the font of the text.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSFontAttributeName
	FontAttributeName foundation.NSAttributedStringKey
	// FontCollectionDidChangeNotification is posted whenever a font collection is changed.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSFontCollection/didChangeNotification
	FontCollectionDidChangeNotification foundation.NSNotificationName
	// FontSetChangedNotification is posted after the currently-set font changes.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSFont/fontSetChangedNotification
	FontSetChangedNotification foundation.NSNotificationName
	// See: https://developer.apple.com/documentation/AppKit/NSFontUnavailableException
	FontUnavailableException foundation.NSExceptionName
	// ForegroundColorAttributeName is the color of the text.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSForegroundColorAttributeName
	ForegroundColorAttributeName foundation.NSAttributedStringKey
	// GlyphInfoAttributeName is the name of a glyph info object.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSGlyphInfoAttributeName
	GlyphInfoAttributeName foundation.NSAttributedStringKey
	// See: https://developer.apple.com/documentation/AppKit/NSIllegalSelectorException
	IllegalSelectorException foundation.NSExceptionName
	// See: https://developer.apple.com/documentation/AppKit/NSImageCacheException
	ImageCacheException foundation.NSExceptionName
	// ImageRepRegistryDidChangeNotification is posted whenever the image representation class registry changes.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSImageRep/registryDidChangeNotification
	ImageRepRegistryDidChangeNotification foundation.NSNotificationName
	// KernAttributeName is the kerning of the text.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSKernAttributeName
	KernAttributeName foundation.NSAttributedStringKey
	// LigatureAttributeName is the ligature of the text.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSLigatureAttributeName
	LigatureAttributeName foundation.NSAttributedStringKey
	// LinkAttributeName is the link for the text.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSLinkAttributeName
	LinkAttributeName foundation.NSAttributedStringKey
	// MarkedClauseSegmentAttributeName is the index of the marked clause segment.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSMarkedClauseSegmentAttributeName
	MarkedClauseSegmentAttributeName foundation.NSAttributedStringKey
	// MenuDidAddItemNotification is posted after a menu item is added to the menu.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSMenu/didAddItemNotification
	MenuDidAddItemNotification foundation.NSNotificationName
	// MenuDidBeginTrackingNotification is posted when menu tracking begins.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSMenu/didBeginTrackingNotification
	MenuDidBeginTrackingNotification foundation.NSNotificationName
	// MenuDidChangeItemNotification is posted after a menu item in the menu changes appearance.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSMenu/didChangeItemNotification
	MenuDidChangeItemNotification foundation.NSNotificationName
	// MenuDidEndTrackingNotification is posted when menu tracking ends, even if no action is sent.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSMenu/didEndTrackingNotification
	MenuDidEndTrackingNotification foundation.NSNotificationName
	// MenuDidRemoveItemNotification is posted after a menu item is removed from the menu.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSMenu/didRemoveItemNotification
	MenuDidRemoveItemNotification foundation.NSNotificationName
	// MenuDidSendActionNotification is posted just after the application dispatches a menu item’s action method to the menu item’s target.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSMenu/didSendActionNotification
	MenuDidSendActionNotification foundation.NSNotificationName
	// MenuWillSendActionNotification is posted just before the application dispatches a menu item’s action method to the menu item’s target.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSMenu/willSendActionNotification
	MenuWillSendActionNotification foundation.NSNotificationName
	// ModalPanelRunLoopMode is the mode set when waiting for input from a modal panel, such as a save or open panel.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSModalPanelRunLoopMode
	ModalPanelRunLoopMode foundation.NSRunLoopMode
	// See: https://developer.apple.com/documentation/AppKit/NSNibLoadingException
	NibLoadingException foundation.NSExceptionName
	// ObliquenessAttributeName is the obliqueness of the text.
	//
	// Deprecated: Deprecated since macOS 26.2.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSObliquenessAttributeName
	ObliquenessAttributeName foundation.NSAttributedStringKey
	// OutlineViewColumnDidMoveNotification is posted whenever a column is moved by user action in an [NSOutlineView] object.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSOutlineView/columnDidMoveNotification
	OutlineViewColumnDidMoveNotification foundation.NSNotificationName
	// OutlineViewColumnDidResizeNotification is posted whenever a column is resized in an [NSOutlineView] object.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSOutlineView/columnDidResizeNotification
	OutlineViewColumnDidResizeNotification foundation.NSNotificationName
	// OutlineViewItemDidCollapseNotification is posted whenever an item is collapsed in an [NSOutlineView] object.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSOutlineView/itemDidCollapseNotification
	OutlineViewItemDidCollapseNotification foundation.NSNotificationName
	// OutlineViewItemDidExpandNotification is posted whenever an item is expanded in an [NSOutlineView] object.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSOutlineView/itemDidExpandNotification
	OutlineViewItemDidExpandNotification foundation.NSNotificationName
	// OutlineViewItemWillCollapseNotification is posted before an item is collapsed (after the user clicks the arrow but before the item is collapsed).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSOutlineView/itemWillCollapseNotification
	OutlineViewItemWillCollapseNotification foundation.NSNotificationName
	// OutlineViewItemWillExpandNotification is posted before an item is expanded (after the user clicks the arrow but before the item is collapsed).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSOutlineView/itemWillExpandNotification
	OutlineViewItemWillExpandNotification foundation.NSNotificationName
	// OutlineViewSelectionDidChangeNotification is posted after the outline view’s selection changes.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSOutlineView/selectionDidChangeNotification
	OutlineViewSelectionDidChangeNotification foundation.NSNotificationName
	// OutlineViewSelectionIsChangingNotification is posted as the outline view’s selection changes (while the mouse button is still down).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSOutlineView/selectionIsChangingNotification
	OutlineViewSelectionIsChangingNotification foundation.NSNotificationName
	// See: https://developer.apple.com/documentation/AppKit/NSPPDIncludeNotFoundException
	PPDIncludeNotFoundException foundation.NSExceptionName
	// See: https://developer.apple.com/documentation/AppKit/NSPPDIncludeStackOverflowException
	PPDIncludeStackOverflowException foundation.NSExceptionName
	// See: https://developer.apple.com/documentation/AppKit/NSPPDIncludeStackUnderflowException
	PPDIncludeStackUnderflowException foundation.NSExceptionName
	// See: https://developer.apple.com/documentation/AppKit/NSPPDParseException
	PPDParseException foundation.NSExceptionName
	// ParagraphStyleAttributeName is the paragraph style of the text.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSParagraphStyleAttributeName
	ParagraphStyleAttributeName foundation.NSAttributedStringKey
	// See: https://developer.apple.com/documentation/AppKit/NSPasteboardCommunicationException
	PasteboardCommunicationException foundation.NSExceptionName
	// PopUpButtonCellWillPopUpNotification is this notification is posted just before a pop-up menu is attached to its window frame.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/willPopUpNotification
	PopUpButtonCellWillPopUpNotification foundation.NSNotificationName
	// PopUpButtonWillPopUpNotification is posted when an [NSPopUpButton] object receives a mouse-down event—that is, when the user is about to select an item from the menu.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPopUpButton/willPopUpNotification
	PopUpButtonWillPopUpNotification foundation.NSNotificationName
	// PopoverDidCloseNotification is sent after the popover has finished animating offscreen.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPopover/didCloseNotification
	PopoverDidCloseNotification foundation.NSNotificationName
	// PopoverDidShowNotification is sent after the popover has finished animating onscreen.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPopover/didShowNotification
	PopoverDidShowNotification foundation.NSNotificationName
	// PopoverWillCloseNotification is sent before the popover is closed.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPopover/willCloseNotification
	PopoverWillCloseNotification foundation.NSNotificationName
	// PopoverWillShowNotification is sent before the popover is shown.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPopover/willShowNotification
	PopoverWillShowNotification foundation.NSNotificationName
	// PreferredScrollerStyleDidChangeNotification is posted if the preferred scroller style changes.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSScroller/preferredScrollerStyleDidChangeNotification
	PreferredScrollerStyleDidChangeNotification foundation.NSNotificationName
	// PrintOperationExistsException is the name of an exception raised when there is already a print operation in process.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPrintOperationExistsException
	PrintOperationExistsException foundation.NSExceptionName
	// See: https://developer.apple.com/documentation/AppKit/NSPrintPackageException
	PrintPackageException foundation.NSExceptionName
	// See: https://developer.apple.com/documentation/AppKit/NSPrintingCommunicationException
	PrintingCommunicationException foundation.NSExceptionName
	// See: https://developer.apple.com/documentation/AppKit/NSRTFPropertyStackOverflowException
	RTFPropertyStackOverflowException foundation.NSExceptionName
	// RuleEditorRowsDidChangeNotification is this notification is posted to the default notification center whenever the view’s rows change.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSRuleEditor/rowsDidChangeNotification
	RuleEditorRowsDidChangeNotification foundation.NSNotificationName
	// ScreenColorSpaceDidChangeNotification is posted when the color space of the screen has changed.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSScreen/colorSpaceDidChangeNotification
	ScreenColorSpaceDidChangeNotification foundation.NSNotificationName
	// ScrollViewDidEndLiveMagnifyNotification is posted at the end of a magnify gesture.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSScrollView/didEndLiveMagnifyNotification
	ScrollViewDidEndLiveMagnifyNotification foundation.NSNotificationName
	// ScrollViewDidEndLiveScrollNotification is posted on the main thread at the end of live scroll tracking.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSScrollView/didEndLiveScrollNotification
	ScrollViewDidEndLiveScrollNotification foundation.NSNotificationName
	// ScrollViewDidLiveScrollNotification is posted on the main thread after changing the clipview bounds origin due to a user-initiated event.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSScrollView/didLiveScrollNotification
	ScrollViewDidLiveScrollNotification foundation.NSNotificationName
	// ScrollViewWillStartLiveMagnifyNotification is posted at the beginning of a magnify gesture.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSScrollView/willStartLiveMagnifyNotification
	ScrollViewWillStartLiveMagnifyNotification foundation.NSNotificationName
	// ScrollViewWillStartLiveScrollNotification is posted on the main thread at the beginning of user-initiated live scroll tracking (gesture scroll or scroller tracking, for example, thumb dragging).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSScrollView/willStartLiveScrollNotification
	ScrollViewWillStartLiveScrollNotification foundation.NSNotificationName
	// ShadowAttributeName is the shadow of the text.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSShadowAttributeName
	ShadowAttributeName foundation.NSAttributedStringKey
	// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/didChangeAutomaticCapitalizationNotification
	SpellCheckerDidChangeAutomaticCapitalizationNotification foundation.NSNotificationName
	// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/didChangeAutomaticDashSubstitutionNotification
	SpellCheckerDidChangeAutomaticDashSubstitutionNotification foundation.NSNotificationName
	// See: https://developer.apple.com/documentation/AppKit/NSSpellCheckerDidChangeAutomaticInlinePredictionNotification
	SpellCheckerDidChangeAutomaticInlinePredictionNotification foundation.NSNotificationName
	// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/didChangeAutomaticPeriodSubstitutionNotification
	SpellCheckerDidChangeAutomaticPeriodSubstitutionNotification foundation.NSNotificationName
	// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/didChangeAutomaticQuoteSubstitutionNotification
	SpellCheckerDidChangeAutomaticQuoteSubstitutionNotification foundation.NSNotificationName
	// SpellCheckerDidChangeAutomaticSpellingCorrectionNotification is this notification is posted when the spell checker did change text using automatic spell checking correction. The are posted to the application’s default notification center.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/didChangeAutomaticSpellingCorrectionNotification
	SpellCheckerDidChangeAutomaticSpellingCorrectionNotification foundation.NSNotificationName
	// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/didChangeAutomaticTextCompletionNotification
	SpellCheckerDidChangeAutomaticTextCompletionNotification foundation.NSNotificationName
	// SpellCheckerDidChangeAutomaticTextReplacementNotification is posted when the spell checker changed text using automatic text replacement. This notification is posted to the app’s default notification center.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/didChangeAutomaticTextReplacementNotification
	SpellCheckerDidChangeAutomaticTextReplacementNotification foundation.NSNotificationName
	// SpellingStateAttributeName is the spelling state of the text.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpellingStateAttributeName
	SpellingStateAttributeName foundation.NSAttributedStringKey
	// SplitViewDidResizeSubviewsNotification is a notification that posts after a change to the size of some or all subviews of a split view.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSplitView/didResizeSubviewsNotification
	SplitViewDidResizeSubviewsNotification foundation.NSNotificationName
	// SplitViewWillResizeSubviewsNotification is a notification that posts before a change to the size of some or all subviews of a split view.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSplitView/willResizeSubviewsNotification
	SplitViewWillResizeSubviewsNotification foundation.NSNotificationName
	// StrikethroughColorAttributeName is the color of the strikethrough.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSStrikethroughColorAttributeName
	StrikethroughColorAttributeName foundation.NSAttributedStringKey
	// StrikethroughStyleAttributeName is the strikethrough style of the text.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSStrikethroughStyleAttributeName
	StrikethroughStyleAttributeName foundation.NSAttributedStringKey
	// StrokeColorAttributeName is the color of the stroke.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSStrokeColorAttributeName
	StrokeColorAttributeName foundation.NSAttributedStringKey
	// StrokeWidthAttributeName is the width of the stroke.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSStrokeWidthAttributeName
	StrokeWidthAttributeName foundation.NSAttributedStringKey
	// SuperscriptAttributeName is the superscript of the text.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSuperscriptAttributeName
	SuperscriptAttributeName foundation.NSAttributedStringKey
	// SystemColorsDidChangeNotification is sent when the system colors have changed, such as through a system control panel interface.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSColor/systemColorsDidChangeNotification
	SystemColorsDidChangeNotification foundation.NSNotificationName
	// See: https://developer.apple.com/documentation/AppKit/NSTIFFException
	TIFFException foundation.NSExceptionName
	// TableViewColumnDidMoveNotification is posted whenever a column is moved by user action in an [NSTableView] object.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTableView/columnDidMoveNotification
	TableViewColumnDidMoveNotification foundation.NSNotificationName
	// TableViewColumnDidResizeNotification is posted whenever a column is resized in an [NSTableView] object.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTableView/columnDidResizeNotification
	TableViewColumnDidResizeNotification foundation.NSNotificationName
	// TableViewSelectionDidChangeNotification is posted after an [NSTableView] object’s selection changes.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTableView/selectionDidChangeNotification
	TableViewSelectionDidChangeNotification foundation.NSNotificationName
	// TableViewSelectionIsChangingNotification is posted as an [NSTableView] object’s selection changes (while the mouse button is still down).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTableView/selectionIsChangingNotification
	TableViewSelectionIsChangingNotification foundation.NSNotificationName
	// TextAlternativesAttributeName is the alternatives for the text.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextAlternativesAttributeName
	TextAlternativesAttributeName foundation.NSAttributedStringKey
	// TextAlternativesSelectedAlternativeStringNotification is posted when the user selects an alternate string.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextAlternatives/selectedAlternativeStringNotification
	TextAlternativesSelectedAlternativeStringNotification foundation.NSNotificationName
	// See: https://developer.apple.com/documentation/AppKit/NSTextContentStorageUnsupportedAttributeAddedNotification
	TextContentStorageUnsupportedAttributeAddedNotification foundation.NSNotificationName
	// TextDidBeginEditingNotification is posted when an [NSText] object begins any operation that changes characters or formatting attributes.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSText/didBeginEditingNotification
	TextDidBeginEditingNotification foundation.NSNotificationName
	// TextDidChangeNotification is posted after an [NSText] object performs any operation that changes characters or formatting attributes.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSText/didChangeNotification
	TextDidChangeNotification foundation.NSNotificationName
	// TextDidEndEditingNotification is posted when focus leaves an [NSText] object, whether or not any operation has changed characters or formatting attributes.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSText/didEndEditingNotification
	TextDidEndEditingNotification foundation.NSNotificationName
	// TextEffectAttributeName is an attribute that applies a text effect to the text.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextEffectAttributeName
	TextEffectAttributeName foundation.NSAttributedStringKey
	// TextHighlightColorSchemeAttributeName is the custom highlight color to apply to the text.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextHighlightColorSchemeAttributeName
	TextHighlightColorSchemeAttributeName foundation.NSAttributedStringKey
	// TextHighlightStyleAttributeName is an attribute that adds a highlight color to the text to emphasize it.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextHighlightStyleAttributeName
	TextHighlightStyleAttributeName foundation.NSAttributedStringKey
	// TextInputContextKeyboardSelectionDidChangeNotification is posted after the selected text input source changes.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextInputContext/keyboardSelectionDidChangeNotification
	TextInputContextKeyboardSelectionDidChangeNotification foundation.NSNotificationName
	// TextLineTooLongException is exception generated if a line is too long in an [NSText] object.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextLineTooLongException
	TextLineTooLongException foundation.NSExceptionName
	// See: https://developer.apple.com/documentation/AppKit/NSTextNoSelectionException
	TextNoSelectionException foundation.NSExceptionName
	// See: https://developer.apple.com/documentation/AppKit/NSTextReadException
	TextReadException foundation.NSExceptionName
	// TextStorageDidProcessEditingNotification is a notification that posts after a text storage finishes processing edits.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextStorage/didProcessEditingNotification
	TextStorageDidProcessEditingNotification foundation.NSNotificationName
	// TextStorageWillProcessEditingNotification is a notification that posts before a text storage begins processing edits.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextStorage/willProcessEditingNotification
	TextStorageWillProcessEditingNotification foundation.NSNotificationName
	// TextViewDidChangeSelectionNotification is posted when the selected range of characters changes.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextView/didChangeSelectionNotification
	TextViewDidChangeSelectionNotification foundation.NSNotificationName
	// TextViewDidChangeTypingAttributesNotification is posted when there is a change in the typing attributes within a text view.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextView/didChangeTypingAttributesNotification
	TextViewDidChangeTypingAttributesNotification foundation.NSNotificationName
	// TextViewDidSwitchToNSLayoutManagerNotification is posted by the framework after switching to using the compatibility mode layout manager.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextView/didSwitchToNSLayoutManagerNotification
	TextViewDidSwitchToNSLayoutManagerNotification foundation.NSNotificationName
	// TextViewWillChangeNotifyingTextViewNotification is posted when a new text view is established as the text view that sends notifications.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextView/willChangeNotifyingTextViewNotification
	TextViewWillChangeNotifyingTextViewNotification foundation.NSNotificationName
	// TextViewWillSwitchToNSLayoutManagerNotification is posted by the framework before switching to the compatibility mode layout manager.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextView/willSwitchToNSLayoutManagerNotification
	TextViewWillSwitchToNSLayoutManagerNotification foundation.NSNotificationName
	// See: https://developer.apple.com/documentation/AppKit/NSTextWriteException
	TextWriteException foundation.NSExceptionName
	// ToolTipAttributeName is the tooltip text.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSToolTipAttributeName
	ToolTipAttributeName foundation.NSAttributedStringKey
	// ToolbarDidRemoveItemNotification is posted after an item is removed from a toolbar.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSToolbar/didRemoveItemNotification
	ToolbarDidRemoveItemNotification foundation.NSNotificationName
	// ToolbarWillAddItemNotification is posts before the toolbar adds a new item.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSToolbar/willAddItemNotification
	ToolbarWillAddItemNotification foundation.NSNotificationName
	// TrackingAttributeName is the amount to modify the default tracking.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTrackingAttributeName
	TrackingAttributeName foundation.NSAttributedStringKey
	// See: https://developer.apple.com/documentation/AppKit/NSTypedStreamVersionException
	TypedStreamVersionException foundation.NSExceptionName
	// UnderlineColorAttributeName is the color of the underline.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSUnderlineColorAttributeName
	UnderlineColorAttributeName foundation.NSAttributedStringKey
	// UnderlineStyleAttributeName is the underline style of the text.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSUnderlineStyleAttributeName
	UnderlineStyleAttributeName foundation.NSAttributedStringKey
	// VerticalGlyphFormAttributeName is the vertical glyph form of the text.
	//
	// Deprecated: Deprecated since macOS 26.2.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSVerticalGlyphFormAttributeName
	VerticalGlyphFormAttributeName foundation.NSAttributedStringKey
	// ViewBoundsDidChangeNotification is a notification that posts when the view’s bounds rectangle changes to a new value independently of the frame rectangle.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSView/boundsDidChangeNotification
	ViewBoundsDidChangeNotification foundation.NSNotificationName
	// ViewDidUpdateTrackingAreasNotification is posted whenever a view recalculates its tracking areas.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSView/didUpdateTrackingAreasNotification
	ViewDidUpdateTrackingAreasNotification foundation.NSNotificationName
	// ViewFrameDidChangeNotification is a notification that posts when the view’s frame rectangle changes to a new value.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSView/frameDidChangeNotification
	ViewFrameDidChangeNotification foundation.NSNotificationName
	// WindowDidBecomeKeyNotification is a notification that the window object became the key window.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWindow/didBecomeKeyNotification
	WindowDidBecomeKeyNotification foundation.NSNotificationName
	// WindowDidBecomeMainNotification is a notification that the window object became the main window.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWindow/didBecomeMainNotification
	WindowDidBecomeMainNotification foundation.NSNotificationName
	// WindowDidChangeBackingPropertiesNotification is a notification that the window object backing properties changed.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWindow/didChangeBackingPropertiesNotification
	WindowDidChangeBackingPropertiesNotification foundation.NSNotificationName
	// WindowDidChangeOcclusionStateNotification is a notification that the window object’s occlusion state changed.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWindow/didChangeOcclusionStateNotification
	WindowDidChangeOcclusionStateNotification foundation.NSNotificationName
	// WindowDidChangeScreenNotification is a notification that a portion of the window object’s frame moved onto or off of a screen.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWindow/didChangeScreenNotification
	WindowDidChangeScreenNotification foundation.NSNotificationName
	// WindowDidChangeScreenProfileNotification is a notification that the screen containing the window changed.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWindow/didChangeScreenProfileNotification
	WindowDidChangeScreenProfileNotification foundation.NSNotificationName
	// WindowDidDeminiaturizeNotification is a notification that the window is no longer minimized.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWindow/didDeminiaturizeNotification
	WindowDidDeminiaturizeNotification foundation.NSNotificationName
	// WindowDidEndLiveResizeNotification is a notification that the user resized the window object.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWindow/didEndLiveResizeNotification
	WindowDidEndLiveResizeNotification foundation.NSNotificationName
	// WindowDidEndSheetNotification is a notification that the window object closed an attached sheet.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWindow/didEndSheetNotification
	WindowDidEndSheetNotification foundation.NSNotificationName
	// WindowDidEnterFullScreenNotification is a notification that the window entered full-screen mode.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWindow/didEnterFullScreenNotification
	WindowDidEnterFullScreenNotification foundation.NSNotificationName
	// WindowDidEnterVersionBrowserNotification is a notification that the window object entered version browser mode.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWindow/didEnterVersionBrowserNotification
	WindowDidEnterVersionBrowserNotification foundation.NSNotificationName
	// WindowDidExitFullScreenNotification is a notification that the window object exited full-screen mode.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWindow/didExitFullScreenNotification
	WindowDidExitFullScreenNotification foundation.NSNotificationName
	// WindowDidExitVersionBrowserNotification is a notification that the window object exited version browser mode.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWindow/didExitVersionBrowserNotification
	WindowDidExitVersionBrowserNotification foundation.NSNotificationName
	// WindowDidExposeNotification is a notification that a window exposed a portion of its nonretained content.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWindow/didExposeNotification
	WindowDidExposeNotification foundation.NSNotificationName
	// WindowDidMiniaturizeNotification is a notification that the window object minimized.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWindow/didMiniaturizeNotification
	WindowDidMiniaturizeNotification foundation.NSNotificationName
	// WindowDidMoveNotification is a notification that the window object moved.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWindow/didMoveNotification
	WindowDidMoveNotification foundation.NSNotificationName
	// WindowDidResignKeyNotification is a notification that the window object resigned its status as key window.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWindow/didResignKeyNotification
	WindowDidResignKeyNotification foundation.NSNotificationName
	// WindowDidResignMainNotification is a notification that the window object resigned its status as main window.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWindow/didResignMainNotification
	WindowDidResignMainNotification foundation.NSNotificationName
	// WindowDidResizeNotification is a notification that the window object size changed.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWindow/didResizeNotification
	WindowDidResizeNotification foundation.NSNotificationName
	// WindowDidUpdateNotification is a notification that the window object received an update message.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWindow/didUpdateNotification
	WindowDidUpdateNotification foundation.NSNotificationName
	// See: https://developer.apple.com/documentation/AppKit/NSWindowServerCommunicationException
	WindowServerCommunicationException foundation.NSExceptionName
	// WindowWillBeginSheetNotification is a notification that the window object is about to open a sheet.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWindow/willBeginSheetNotification
	WindowWillBeginSheetNotification foundation.NSNotificationName
	// WindowWillCloseNotification is a notification that the window object is about to close.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWindow/willCloseNotification
	WindowWillCloseNotification foundation.NSNotificationName
	// WindowWillEnterFullScreenNotification is a notification that the window will enter full-screen mode.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWindow/willEnterFullScreenNotification
	WindowWillEnterFullScreenNotification foundation.NSNotificationName
	// WindowWillEnterVersionBrowserNotification is a notification that the window object will enter version browser mode.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWindow/willEnterVersionBrowserNotification
	WindowWillEnterVersionBrowserNotification foundation.NSNotificationName
	// WindowWillExitFullScreenNotification is a notification that the window object will exit full-screen mode.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWindow/willExitFullScreenNotification
	WindowWillExitFullScreenNotification foundation.NSNotificationName
	// WindowWillExitVersionBrowserNotification is a notification that the window object will exit version browser mode.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWindow/willExitVersionBrowserNotification
	WindowWillExitVersionBrowserNotification foundation.NSNotificationName
	// WindowWillMiniaturizeNotification is a notification that the window object is about to minimize.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWindow/willMiniaturizeNotification
	WindowWillMiniaturizeNotification foundation.NSNotificationName
	// WindowWillMoveNotification is a notification that the window object is about to move.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWindow/willMoveNotification
	WindowWillMoveNotification foundation.NSNotificationName
	// WindowWillStartLiveResizeNotification is a notification that the user is about to resize the window.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWindow/willStartLiveResizeNotification
	WindowWillStartLiveResizeNotification foundation.NSNotificationName
	// See: https://developer.apple.com/documentation/AppKit/NSWordTablesReadException
	WordTablesReadException foundation.NSExceptionName
	// See: https://developer.apple.com/documentation/AppKit/NSWordTablesWriteException
	WordTablesWriteException foundation.NSExceptionName
	// WorkspaceAccessibilityDisplayOptionsDidChangeNotification is a notification that the workspace posts when any of the accessibility display options change.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/accessibilityDisplayOptionsDidChangeNotification
	WorkspaceAccessibilityDisplayOptionsDidChangeNotification foundation.NSNotificationName
	// WorkspaceActiveSpaceDidChangeNotification is a notification that the workspace posts when a Spaces change occurs.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/activeSpaceDidChangeNotification
	WorkspaceActiveSpaceDidChangeNotification foundation.NSNotificationName
	// WorkspaceDidActivateApplicationNotification is a notification that the workspace posts when the Finder is about to activate an app.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/didActivateApplicationNotification
	WorkspaceDidActivateApplicationNotification foundation.NSNotificationName
	// WorkspaceDidChangeFileLabelsNotification is a notification that the workspace posts when the Finder file labels or colors change.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/didChangeFileLabelsNotification
	WorkspaceDidChangeFileLabelsNotification foundation.NSNotificationName
	// WorkspaceDidDeactivateApplicationNotification is a notification that the workspace posts when the Finder deactivates an app.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/didDeactivateApplicationNotification
	WorkspaceDidDeactivateApplicationNotification foundation.NSNotificationName
	// WorkspaceDidHideApplicationNotification is a notification that the workspace posts when the Finder hides an app.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/didHideApplicationNotification
	WorkspaceDidHideApplicationNotification foundation.NSNotificationName
	// WorkspaceDidLaunchApplicationNotification is a notification that the workspace posts when a new app starts up.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/didLaunchApplicationNotification
	WorkspaceDidLaunchApplicationNotification foundation.NSNotificationName
	// WorkspaceDidMountNotification is a notification that the workspace posts when a new device mounts.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/didMountNotification
	WorkspaceDidMountNotification foundation.NSNotificationName
	// WorkspaceDidRenameVolumeNotification is a notification that the workspace posts when a volume changes its name or mount path.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/didRenameVolumeNotification
	WorkspaceDidRenameVolumeNotification foundation.NSNotificationName
	// WorkspaceDidTerminateApplicationNotification is a notification that the workspace posts when an app finishes executing.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/didTerminateApplicationNotification
	WorkspaceDidTerminateApplicationNotification foundation.NSNotificationName
	// WorkspaceDidUnhideApplicationNotification is a notification that the workspace posts when the Finder unhides an app.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/didUnhideApplicationNotification
	WorkspaceDidUnhideApplicationNotification foundation.NSNotificationName
	// WorkspaceDidUnmountNotification is a notification that the workspace posts when the Finder unmounts a device.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/didUnmountNotification
	WorkspaceDidUnmountNotification foundation.NSNotificationName
	// WorkspaceDidWakeNotification is a notification that the workspace posts when the device wakes from sleep.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/didWakeNotification
	WorkspaceDidWakeNotification foundation.NSNotificationName
	// WorkspaceScreensDidSleepNotification is a notification that the workspace posts when the device’s screen goes to sleep.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/screensDidSleepNotification
	WorkspaceScreensDidSleepNotification foundation.NSNotificationName
	// WorkspaceScreensDidWakeNotification is a notification that the workspace posts when the device’s screens wake.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/screensDidWakeNotification
	WorkspaceScreensDidWakeNotification foundation.NSNotificationName
	// WorkspaceSessionDidBecomeActiveNotification is a notification that the workspace posts after a user session switches in.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/sessionDidBecomeActiveNotification
	WorkspaceSessionDidBecomeActiveNotification foundation.NSNotificationName
	// WorkspaceSessionDidResignActiveNotification is a notification that the workspace posts before a user session switches out.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/sessionDidResignActiveNotification
	WorkspaceSessionDidResignActiveNotification foundation.NSNotificationName
	// WorkspaceWillLaunchApplicationNotification is a notification that the workspace posts when the Finder is about to launch an app.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/willLaunchApplicationNotification
	WorkspaceWillLaunchApplicationNotification foundation.NSNotificationName
	// WorkspaceWillPowerOffNotification is a notification that the workspace posts when the user requests a logout or powers off the device.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/willPowerOffNotification
	WorkspaceWillPowerOffNotification foundation.NSNotificationName
	// WorkspaceWillSleepNotification is a notification that the workspace posts before the device goes to sleep.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/willSleepNotification
	WorkspaceWillSleepNotification foundation.NSNotificationName
	// WorkspaceWillUnmountNotification is a notification that the workspace posts when the Finder is about to unmount a device.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/willUnmountNotification
	WorkspaceWillUnmountNotification foundation.NSNotificationName
	// WritingDirectionAttributeName is the writing direction of the text.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWritingDirectionAttributeName
	WritingDirectionAttributeName foundation.NSAttributedStringKey
	// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsExclusionAttributeName
	WritingToolsExclusionAttributeName foundation.NSAttributedStringKey
)

var (
	// AboutPanelOptionApplicationIcon is the icon to display for the app in the About panel.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSApplication/AboutPanelOptionKey/applicationIcon
	AboutPanelOptionApplicationIcon NSAboutPanelOptionKey
	// AboutPanelOptionApplicationName is the name of the application to display in the About panel.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSApplication/AboutPanelOptionKey/applicationName
	AboutPanelOptionApplicationName NSAboutPanelOptionKey
	// AboutPanelOptionApplicationVersion is the version information to display in the About panel.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSApplication/AboutPanelOptionKey/applicationVersion
	AboutPanelOptionApplicationVersion NSAboutPanelOptionKey
	// AboutPanelOptionCredits is the credits string to display in the About panel.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSApplication/AboutPanelOptionKey/credits
	AboutPanelOptionCredits NSAboutPanelOptionKey
	// AboutPanelOptionVersion is the version number to display in the About panel.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSApplication/AboutPanelOptionKey/version
	AboutPanelOptionVersion NSAboutPanelOptionKey
)

var (
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/activationPoint
	AccessibilityActivationPointAttribute string
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/alternateUIVisible
	AccessibilityAlternateUIVisibleAttribute string
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/autoInteractableAttribute
	AccessibilityAutoInteractableAttribute string
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/blockQuoteLevelAttribute
	AccessibilityBlockQuoteLevelAttribute string
	// AccessibilityCancelButtonAttribute is the element that represents the cancel button (`id`).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/cancelButton
	AccessibilityCancelButtonAttribute string
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/childrenInNavigationOrderAttribute
	AccessibilityChildrenInNavigationOrderAttribute string
	// AccessibilityColumnHeaderUIElementsAttribute is the table’s column headers ([NSArray]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/columnHeaderUIElements
	AccessibilityColumnHeaderUIElementsAttribute string
	// AccessibilityColumnIndexRangeAttribute is the column index range of the cell (an [NSValue] instance that contains the row’s starting index and index span in the table).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/columnIndexRange
	AccessibilityColumnIndexRangeAttribute string
	// AccessibilityColumnTitlesAttribute is the elements that represent the column titles ([NSArray]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/columnTitles
	AccessibilityColumnTitlesAttribute string
	// AccessibilityColumnsAttribute is the table’s columns ([NSArray]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/columns
	AccessibilityColumnsAttribute string
	// AccessibilityContainsProtectedContentAttribute is a flag that indicates whether the object contains protected content.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/containsProtectedContent
	AccessibilityContainsProtectedContentAttribute string
	// AccessibilityContentsAttribute is elements that represent the contents in the current element, such as the document view of a scroll view ([NSArray]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/contents
	AccessibilityContentsAttribute string
	// AccessibilityCriticalValueAttribute is the critical value in a level indicator (typically, [NSNumber]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/criticalValue
	AccessibilityCriticalValueAttribute string
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/dateTimeComponentsAttribute
	AccessibilityDateTimeComponentsAttribute string
	// AccessibilityDecrementButtonAttribute is the element that represents a stepper’s decrement button (`id`).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/decrementButton
	AccessibilityDecrementButtonAttribute string
	// AccessibilityDefaultButtonAttribute is the element that represents the default button (`id`).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/defaultButton
	AccessibilityDefaultButtonAttribute string
	// AccessibilityDescriptionAttribute is the purpose of the element, not including the role ([NSString]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/description
	AccessibilityDescriptionAttribute string
	// AccessibilityDisclosedRowsAttribute is the rows disclosed by this row ([NSArray]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/disclosedRows
	AccessibilityDisclosedRowsAttribute string
	// AccessibilityDisclosingAttribute is a flag that indicates whether a row is disclosing other rows ([NSNumber]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/disclosing
	AccessibilityDisclosingAttribute string
	// AccessibilityDisclosureLevelAttribute is the indentation level of this row ([NSNumber]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/disclosureLevel
	AccessibilityDisclosureLevelAttribute string
	// AccessibilityDocumentAttribute is the URL for the file represented by the element ([NSString]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/document
	AccessibilityDocumentAttribute string
	// AccessibilityEditedAttribute is a flag that indicates whether the element has been modified ([NSNumber]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/edited
	AccessibilityEditedAttribute string
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/embeddedImageDescriptionAttribute
	AccessibilityEmbeddedImageDescriptionAttribute string
	// AccessibilityEnabledAttribute is a flag that indicates the enabled state of the element ([NSNumber]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/enabled
	AccessibilityEnabledAttribute string
	// AccessibilityErrorCodeExceptionInfo is an integer error code for debugging.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/ErrorCodeExceptionInfo
	AccessibilityErrorCodeExceptionInfo string
	// AccessibilityExpandedAttribute is a flag that indicates whether the element is expanded ([NSNumber]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/expanded
	AccessibilityExpandedAttribute string
	// AccessibilityExtrasMenuBarAttribute is the app extras menu bar (`id`).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/extrasMenuBar
	AccessibilityExtrasMenuBarAttribute string
	// AccessibilityFilenameAttribute is the filename associated with the element ([NSString]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/filename
	AccessibilityFilenameAttribute string
	// AccessibilityFocusedAttribute is a flag that indicates the presence of keyboard focus ([NSNumber]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/focused
	AccessibilityFocusedAttribute string
	// AccessibilityFocusedWindowAttribute is the app’s window that has current focus (`id`).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/focusedWindow
	AccessibilityFocusedWindowAttribute string
	// AccessibilityFrontmostAttribute is a flag that indicates whether the app is frontmost ([NSNumber]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/frontmost
	AccessibilityFrontmostAttribute string
	// AccessibilityFullScreenButtonAttribute is the element that represents the full-screen button (`id`).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/fullScreenButton
	AccessibilityFullScreenButtonAttribute string
	// AccessibilityGrowAreaAttribute is the element representing the grow area (`id`).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/growArea
	AccessibilityGrowAreaAttribute string
	// AccessibilityHeaderAttribute is the element that represents a table view’s header (`id`).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/header
	AccessibilityHeaderAttribute string
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/headingLevelAttribute
	AccessibilityHeadingLevelAttribute string
	// AccessibilityHelpAttribute is the help text for the element ([NSString]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/help
	AccessibilityHelpAttribute string
	// AccessibilityHiddenAttribute is a flag that indicates whether the app is hidden ([NSNumber]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/hidden
	AccessibilityHiddenAttribute string
	// AccessibilityHorizontalScrollBarAttribute is the element that represents a scroll view’s horizontal scroll bar (`id`).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/horizontalScrollBar
	AccessibilityHorizontalScrollBarAttribute string
	// AccessibilityHorizontalUnitDescriptionAttribute is the description of the layout view’s horizontal units ([NSString]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/horizontalUnitDescription
	AccessibilityHorizontalUnitDescriptionAttribute string
	// AccessibilityIdentifierAttribute is the identity of the element ([NSString]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/identifier
	AccessibilityIdentifierAttribute string
	// AccessibilityIncrementButtonAttribute is the element that represents a stepper’s increment button (`id`).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/incrementButton
	AccessibilityIncrementButtonAttribute string
	// AccessibilityIndexAttribute is the index of the row or column represented by the element ([NSValue]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/index
	AccessibilityIndexAttribute string
	// AccessibilityLabelUIElementsAttribute is the elements that represent the slider’s labels ([NSArray]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/labelUIElements
	AccessibilityLabelUIElementsAttribute string
	// AccessibilityLabelValueAttribute is the value of the label represented by this element ([NSNumber]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/labelValue
	AccessibilityLabelValueAttribute string
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/languageAttribute
	AccessibilityLanguageAttribute string
	// AccessibilityMainAttribute is a flag that indicates whether the window is the main window ([NSNumber]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/main
	AccessibilityMainAttribute string
	// AccessibilityMainWindowAttribute is the app’s main window (`id`).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/mainWindow
	AccessibilityMainWindowAttribute string
	// AccessibilityMarkerTypeAttribute is the type of the marker ([NSString]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/markerType
	AccessibilityMarkerTypeAttribute string
	// AccessibilityMarkerTypeDescriptionAttribute is the description of the marker type ([NSString]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/markerTypeDescription
	AccessibilityMarkerTypeDescriptionAttribute string
	// AccessibilityMarkerUIElementsAttribute is an array of marker user interface elements ([NSArray]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/markerUIElements
	AccessibilityMarkerUIElementsAttribute string
	// AccessibilityMarkerValuesAttribute is the marker values ([NSArray] of [NSNumber]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/markerValues
	AccessibilityMarkerValuesAttribute string
	// AccessibilityMaxValueAttribute is the element’s maximum value (`id`).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/maxValue
	AccessibilityMaxValueAttribute string
	// AccessibilityMenuBarAttribute is the app’s menu bar (`id`).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/menuBar
	AccessibilityMenuBarAttribute string
	// AccessibilityMinValueAttribute is the element’s minimum value (`id`).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/minValue
	AccessibilityMinValueAttribute string
	// AccessibilityMinimizeButtonAttribute is the element that represents the minimize button (`id`).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/minimizeButton
	AccessibilityMinimizeButtonAttribute string
	// AccessibilityMinimizedAttribute is a flag that indicates whether the window is minimized ([NSNumber]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/minimized
	AccessibilityMinimizedAttribute string
	// AccessibilityModalAttribute is a flag that indicates whether the window represented by this element is modal ([NSNumber]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/modal
	AccessibilityModalAttribute string
	// AccessibilityNextContentsAttribute is the elements representing the contents that follow the current divider element, such as a subview adjacent to a split view’s splitter element ([NSArray]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/nextContents
	AccessibilityNextContentsAttribute string
	// AccessibilityNumberOfCharactersAttribute is the number of characters in the text ([NSNumber]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/numberOfCharacters
	AccessibilityNumberOfCharactersAttribute string
	// AccessibilityOrderedByRowAttribute is a flag that indicates whether the grid is in row major or column major order.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/orderedByRow
	AccessibilityOrderedByRowAttribute string
	// AccessibilityOrientationAttribute is the element’s orientation.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/orientation
	AccessibilityOrientationAttribute string
	// AccessibilityOverflowButtonAttribute is the element that represents a toolbar’s overflow button (`id`).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/overflowButton
	AccessibilityOverflowButtonAttribute string
	// AccessibilityParentAttribute is the element’s parent element in the accessibility hierarchy (`id`).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/parent
	AccessibilityParentAttribute string
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/pathAttribute
	AccessibilityPathAttribute string
	// AccessibilityPlaceholderValueAttribute is the placeholder value for a control, such as a text field ([NSString]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/placeholderValue
	AccessibilityPlaceholderValueAttribute string
	// AccessibilityPositionAttribute is the position in points of the element’s lower-left corner in screen-relative coordinates ([NSValue]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/position
	AccessibilityPositionAttribute string
	// AccessibilityPreviousContentsAttribute is the elements representing the contents that precede the current divider element, such as a subview adjacent to a split view’s splitter bar element ([NSArray]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/previousContents
	AccessibilityPreviousContentsAttribute string
	// AccessibilityProxyAttribute is the element that represents the window’s proxy icon (`id`).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/proxy
	AccessibilityProxyAttribute string
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/required
	AccessibilityRequiredAttribute string
	// AccessibilityRoleAttribute is the element’s type, such as [NSAccessibilityRadioButtonRole] ([NSString]). See Roles for a list of available roles.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/role
	AccessibilityRoleAttribute string
	// AccessibilityRoleDescriptionAttribute is a localized, human-intelligible description of the element’s role, such as `radio button` ([NSString]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/roleDescription
	AccessibilityRoleDescriptionAttribute string
	// AccessibilityRowCountAttribute is the number of rows in the grid ([NSNumber] as `intValue`).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/rowCount
	AccessibilityRowCountAttribute string
	// AccessibilityRowHeaderUIElementsAttribute is the table’s row headers ([NSArray]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/rowHeaderUIElements
	AccessibilityRowHeaderUIElementsAttribute string
	// AccessibilityRowsAttribute is the table’s rows ([NSArray]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/rows
	AccessibilityRowsAttribute string
	// AccessibilitySearchButtonAttribute is the element that represents the search button in a search field (`id`).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/searchButton
	AccessibilitySearchButtonAttribute string
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySearchCurrentElementKey
	AccessibilitySearchCurrentElementKey string
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySearchCurrentRangeKey
	AccessibilitySearchCurrentRangeKey string
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySearchDirectionKey
	AccessibilitySearchDirectionKey string
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySearchDirectionNext
	AccessibilitySearchDirectionNext string
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySearchDirectionPrevious
	AccessibilitySearchDirectionPrevious string
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySearchIdentifiersKey
	AccessibilitySearchIdentifiersKey string
	// AccessibilitySearchMenuAttribute is the element that represents the menu in a search field (`id`).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/searchMenu
	AccessibilitySearchMenuAttribute string
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySearchResultDescriptionOverrideKey
	AccessibilitySearchResultDescriptionOverrideKey string
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySearchResultElementKey
	AccessibilitySearchResultElementKey string
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySearchResultLoaderKey
	AccessibilitySearchResultLoaderKey string
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySearchResultRangeKey
	AccessibilitySearchResultRangeKey string
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySearchResultsLimitKey
	AccessibilitySearchResultsLimitKey string
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySearchTextKey
	AccessibilitySearchTextKey string
	// AccessibilitySelectedAttribute is a flag that indicates whether the element is selected ([NSNumber]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/selected
	AccessibilitySelectedAttribute string
	// AccessibilitySelectedChildrenAttribute is the currently selected children of the element ([NSArray]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/selectedChildren
	AccessibilitySelectedChildrenAttribute string
	// AccessibilitySelectedColumnsAttribute is the table’s selected columns ([NSArray]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/selectedColumns
	AccessibilitySelectedColumnsAttribute string
	// AccessibilitySelectedRowsAttribute is the table’s selected rows ([NSArray]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/selectedRows
	AccessibilitySelectedRowsAttribute string
	// AccessibilitySelectedTextAttribute is the currently selected text ([NSString]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/selectedText
	AccessibilitySelectedTextAttribute string
	// AccessibilitySelectedTextRangeAttribute is the range of selected text ([NSValue]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/selectedTextRange
	AccessibilitySelectedTextRangeAttribute string
	// AccessibilitySelectedTextRangesAttribute is an array of [NSValue] (`rangeValue`) ranges of selected text ([NSArray]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/selectedTextRanges
	AccessibilitySelectedTextRangesAttribute string
	// AccessibilityServesAsTitleForUIElementsAttribute is the elements for which this element serves as the title ([NSArray]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/servesAsTitleForUIElements
	AccessibilityServesAsTitleForUIElementsAttribute string
	// AccessibilitySharedCharacterRangeAttribute is the (`rangeValue`) part of shared text in this view ([NSValue]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/sharedCharacterRange
	AccessibilitySharedCharacterRangeAttribute string
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/sharedFocusElements
	AccessibilitySharedFocusElementsAttribute string
	// AccessibilitySharedTextUIElementsAttribute is the elements with which the text of this element is shared ([NSArray]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/sharedTextUIElements
	AccessibilitySharedTextUIElementsAttribute string
	// AccessibilityShownMenuAttribute is the menu currently being displayed (`id`).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/shownMenu
	AccessibilityShownMenuAttribute string
	// AccessibilitySizeAttribute is the element’s size in points ([NSValue]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/size
	AccessibilitySizeAttribute string
	// AccessibilitySortDirectionAttribute is the column’s sort direction ([NSString]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/sortDirection
	AccessibilitySortDirectionAttribute string
	// AccessibilitySplittersAttribute is the views and splitter bar in a split view ([NSArray]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/splitters
	AccessibilitySplittersAttribute string
	// AccessibilitySubroleAttribute is the element’s subrole, such as [NSAccessibilityTableRowSubrole] ([NSString]). See Subroles for a list of available subroles.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/subrole
	AccessibilitySubroleAttribute string
	// AccessibilityTabsAttribute is the tab elements in a tab view ([NSArray]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/tabs
	AccessibilityTabsAttribute string
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityTextCompletionAttribute
	AccessibilityTextCompletionAttribute string
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/textInputMarkedRangeAttribute
	AccessibilityTextInputMarkedRangeAttribute string
	// AccessibilityTitleAttribute is the title of the element, such as a button’s visible text ([NSString]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/title
	AccessibilityTitleAttribute string
	// AccessibilityTitleUIElementAttribute is an element that represents another element’s static text title (`id`).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/titleUIElement
	AccessibilityTitleUIElementAttribute string
	// AccessibilityToolbarButtonAttribute is the element that represents the toolbar button (`id`).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/toolbarButton
	AccessibilityToolbarButtonAttribute string
	// AccessibilityTopLevelUIElementAttribute is the top-level element that contains this element (`id`).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/topLevelUIElement
	AccessibilityTopLevelUIElementAttribute string
	// AccessibilityURLAttribute is the URL associated with the element ([NSURL]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/url
	AccessibilityURLAttribute string
	// AccessibilityUnitDescriptionAttribute is the description of ruler units ([NSString]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/unitDescription
	AccessibilityUnitDescriptionAttribute string
	// AccessibilityUnitsAttribute is the ruler units ([NSString]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/units
	AccessibilityUnitsAttribute string
	// AccessibilityValueAttribute is the element’s value (`id`).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/value
	AccessibilityValueAttribute string
	// AccessibilityValueDescriptionAttribute is the description of the element’s value ([NSString]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/valueDescription
	AccessibilityValueDescriptionAttribute string
	// AccessibilityVerticalScrollBarAttribute is the element that represents the vertical scroll bar in a scroll view (`id`).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/verticalScrollBar
	AccessibilityVerticalScrollBarAttribute string
	// AccessibilityVerticalUnitDescriptionAttribute is the description of the layout view’s vertical units ([NSString]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/verticalUnitDescription
	AccessibilityVerticalUnitDescriptionAttribute string
	// AccessibilityVerticalUnitsAttribute is the units that the layout view uses for vertical values ([NSString]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/verticalUnits
	AccessibilityVerticalUnitsAttribute string
	// AccessibilityVisibleCellsAttribute is the table’s visible cells ([NSArray]). This attribute is required for cell-based tables.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/visibleCells
	AccessibilityVisibleCellsAttribute string
	// AccessibilityVisibleCharacterRangeAttribute is the range of visible text ([NSValue]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/visibleCharacterRange
	AccessibilityVisibleCharacterRangeAttribute string
	// AccessibilityVisibleChildrenAttribute is the element’s child elements that are visible ([NSArray]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/visibleChildren
	AccessibilityVisibleChildrenAttribute string
	// AccessibilityVisibleColumnsAttribute is the table’s visible columns ([NSArray]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/visibleColumns
	AccessibilityVisibleColumnsAttribute string
	// AccessibilityVisibleRowsAttribute is the table’s visible rows ([NSArray]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/visibleRows
	AccessibilityVisibleRowsAttribute string
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/visitedAttribute
	AccessibilityVisitedAttribute string
	// AccessibilityWarningValueAttribute is the warning value in a level indicator (typically, [NSNumber]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/warningValue
	AccessibilityWarningValueAttribute string
	// AccessibilityWindowAttribute is the window containing the current element (`id`).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/window
	AccessibilityWindowAttribute string
	// AccessibilityWindowsAttribute is the app’s windows ([NSArray]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/windows
	AccessibilityWindowsAttribute string
	// AccessibilityZoomButtonAttribute is the element that represents the zoom button (`id`).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/zoomButton
	AccessibilityZoomButtonAttribute string
	// AlignmentBinding is a constant that identifies an alignment binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/alignment
	AlignmentBinding string
	// AllRomanInputSourcesLocaleIdentifier is a meta-locale identifier representing the set of Roman input sources available. You can pass `[NSArray NSAllRomanInputSourcesLocaleIdentifier]` to the [allowedInputSourceLocales] property to restrict allowed input sources to Roman only.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAllRomanInputSourcesLocaleIdentifier
	AllRomanInputSourcesLocaleIdentifier string
	// AlternateImageBinding is a constant that identifies an alternate image binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/alternateImage
	AlternateImageBinding string
	// AlternateTitleBinding is a constant that identifies an alternate title binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/alternateTitle
	AlternateTitleBinding string
	// AnimateBinding is a constant that identifies an animate binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/animate
	AnimateBinding string
	// AnimationDelayBinding is a constant that identifies an animation delay binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/animationDelay
	AnimationDelayBinding string
	// See: https://developer.apple.com/documentation/AppKit/NSAnimation/progressMarkUserInfoKey
	AnimationProgressMark string
	// ApplicationLaunchIsDefaultLaunchKey is a Boolean value that indicates if the app launch is a default launch.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSApplication/launchIsDefaultUserInfoKey
	ApplicationLaunchIsDefaultLaunchKey string
	// ApplicationLaunchUserNotificationKey is a key that indicates your app was launched because a user activated a notification in the Notification Center.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSApplication/launchUserNotificationUserInfoKey
	ApplicationLaunchUserNotificationKey string
	// ArgumentBinding is a constant that identifies an argument binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/argument
	ArgumentBinding string
	// AttributedStringBinding is a constant that identifies an attributed string binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/attributedString
	AttributedStringBinding string
	// BackingPropertyOldColorSpaceKey is an [NSColorSpace] instance containing the old colorspace.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWindow/oldColorSpaceUserInfoKey
	BackingPropertyOldColorSpaceKey string
	// BackingPropertyOldScaleFactorKey is an NSNumber containing the old scale factor.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWindow/oldScaleFactorUserInfoKey
	BackingPropertyOldScaleFactorKey string
	// ContentArrayBinding is a constant that identifies a content array binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/contentArray
	ContentArrayBinding string
	// ContentArrayForMultipleSelectionBinding is a constant that identifies a content array for multiple selection binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/contentArrayForMultipleSelection
	ContentArrayForMultipleSelectionBinding string
	// ContentBinding is a constant that identifies a content binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/content
	ContentBinding string
	// ContentDictionaryBinding is a constant that identifies a content dictionary binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/contentDictionary
	ContentDictionaryBinding string
	// ContentHeightBinding is a constant that identifies a content height binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/contentHeight
	ContentHeightBinding string
	// ContentObjectBinding is a constant that identifies a content object binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/contentObject
	ContentObjectBinding string
	// ContentObjectsBinding is a constant that identifies a content objects binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/contentObjects
	ContentObjectsBinding string
	// ContentSetBinding is a constant that identifies a content set binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/contentSet
	ContentSetBinding string
	// ContentValuesBinding is a constant that identifies a content values binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/contentValues
	ContentValuesBinding string
	// ContentWidthBinding is a constant that identifies a content width binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/contentWidth
	ContentWidthBinding string
	// CriticalValueBinding is a constant that identifies a critical value binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/criticalValue
	CriticalValueBinding string
	// DataBinding is a constant that identifies a data binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/data
	DataBinding string
	// DisplayPatternTitleBinding is a constant that identifies a display pattern title binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/displayPatternTitle
	DisplayPatternTitleBinding string
	// DisplayPatternValueBinding is a constant that identifies a display pattern value binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/displayPatternValue
	DisplayPatternValueBinding string
	// DocumentEditedBinding is a constant that identifies a document edited binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/documentEdited
	DocumentEditedBinding string
	// DoubleClickArgumentBinding is a constant that identifies a double-click argument binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/doubleClickArgument
	DoubleClickArgumentBinding string
	// DoubleClickTargetBinding is a constant that identifies a double-click target binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/doubleClickTarget
	DoubleClickTargetBinding string
	// EditableBinding is a constant that identifies an editable binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/editable
	EditableBinding string
	// EnabledBinding is a constant that identifies an enabled binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/enabled
	EnabledBinding string
	// ExcludedKeysBinding is a constant that identifies an excluded keys binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/excludedKeys
	ExcludedKeysBinding string
	// FilterPredicateBinding is a constant that identifies a filter predicate binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/filterPredicate
	FilterPredicateBinding string
	// FontBinding is a constant that identifies a font binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/font
	FontBinding string
	// FontBoldBinding is a constant that identifies a font bold binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/fontBold
	FontBoldBinding string
	// FontFamilyNameBinding is a constant that identifies a font family name binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/fontFamilyName
	FontFamilyNameBinding string
	// FontItalicBinding is a constant that identifies a font italic binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/fontItalic
	FontItalicBinding string
	// FontNameBinding is a constant that identifies a font name binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/fontName
	FontNameBinding string
	// FontSizeBinding is a constant that identifies a font size binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/fontSize
	FontSizeBinding string
	// GraphicsContextDestinationAttributeName is specifies the destination.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/AttributeKey/destination
	GraphicsContextDestinationAttributeName string
	// GraphicsContextRepresentationFormatAttributeName is specifies the destination file format.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/AttributeKey/representationFormat
	GraphicsContextRepresentationFormatAttributeName string
	// HeaderTitleBinding is a constant that identifies a header title binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/headerTitle
	HeaderTitleBinding string
	// HiddenBinding is a constant that identifies a hidden binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/hidden
	HiddenBinding string
	// ImageBinding is a constant that identifies an image binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/image
	ImageBinding string
	// IncludedKeysBinding is a constant that identifies an included keys binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/includedKeys
	IncludedKeysBinding string
	// InitialKeyBinding is a constant that identifies an initial key binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/initialKey
	InitialKeyBinding string
	// InitialValueBinding is a constant that identifies an initial value binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/initialValue
	InitialValueBinding string
	// IsIndeterminateBinding is a constant that identifies an is indeterminate binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/isIndeterminate
	IsIndeterminateBinding string
	// LabelBinding is a constant that identifies a label binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/label
	LabelBinding string
	// LocalizedKeyDictionaryBinding is a constant that identifies a localized key dictionary binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/localizedKeyDictionary
	LocalizedKeyDictionaryBinding string
	// ManagedObjectContextBinding is a constant that identifies a managed object context binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/managedObjectContext
	ManagedObjectContextBinding string
	// MaxValueBinding is a constant that identifies a maximum value binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/maxValue
	MaxValueBinding string
	// MaxWidthBinding is a constant that identifies a maximum width binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/maxWidth
	MaxWidthBinding string
	// MaximumRecentsBinding is a constant that identifies a maximum recents binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/maximumRecents
	MaximumRecentsBinding string
	// MinValueBinding is a constant that identifies a minimum value binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/minValue
	MinValueBinding string
	// MinWidthBinding is a constant that identifies a minimum width binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/minWidth
	MinWidthBinding string
	// MixedStateImageBinding is a constant that identifies a mixed state image binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/mixedStateImage
	MixedStateImageBinding string
	// OffStateImageBinding is a constant that identifies an off state image binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/offStateImage
	OffStateImageBinding string
	// OnStateImageBinding is a constant that identifies an on state image binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/onStateImage
	OnStateImageBinding string
	// PopoverCloseReasonKey is the `userInfo` key containing the reason for the [willCloseNotification].
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPopover/closeReasonUserInfoKey
	PopoverCloseReasonKey string
	// PositioningRectBinding is a constant that identifies a positioning rectangle binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/positioningRect
	PositioningRectBinding string
	// PredicateBinding is a constant that identifies a predicate binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/predicate
	PredicateBinding string
	// RecentSearchesBinding is a constant that identifies a recent searches binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/recentSearches
	RecentSearchesBinding string
	// RepresentedFilenameBinding is a constant that identifies a represented filename binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/representedFilename
	RepresentedFilenameBinding string
	// RowHeightBinding is a constant that identifies a row height binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/rowHeight
	RowHeightBinding string
	// SelectedIdentifierBinding is a constant that identifies a selected identifier binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/selectedIdentifier
	SelectedIdentifierBinding string
	// SelectedIndexBinding is a constant that identifies a selected index binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/selectedIndex
	SelectedIndexBinding string
	// SelectedLabelBinding is a constant that identifies a selected label binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/selectedLabel
	SelectedLabelBinding string
	// SelectedObjectBinding is a constant that identifies a selected object binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/selectedObject
	SelectedObjectBinding string
	// SelectedObjectsBinding is a constant that identifies a selected objects binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/selectedObjects
	SelectedObjectsBinding string
	// SelectedTagBinding is a constant that identifies a selected tag binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/selectedTag
	SelectedTagBinding string
	// SelectedValueBinding is a constant that identifies a selected value binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/selectedValue
	SelectedValueBinding string
	// SelectedValuesBinding is a constant that identifies a selected values binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/selectedValues
	SelectedValuesBinding string
	// SelectionIndexPathsBinding is a constant that identifies a selection index paths binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/selectionIndexPaths
	SelectionIndexPathsBinding string
	// SelectionIndexesBinding is a constant that identifies a selection indexes binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/selectionIndexes
	SelectionIndexesBinding string
	// SortDescriptorsBinding is a constant that identifies a sort descriptors binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/sortDescriptors
	SortDescriptorsBinding string
	// TargetBinding is a constant that identifies a target binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/target
	TargetBinding string
	// TextColorBinding is a constant that identifies a text color binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/textColor
	TextColorBinding string
	// TextMovementUserInfoKey is the `userInfo` dictionary key for the [didEndEditingNotification] notification.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSText/movementUserInfoKey
	TextMovementUserInfoKey string
	// TitleBinding is a constant that identifies a title binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/title
	TitleBinding string
	// ToolTipBinding is a constant that identifies a tool tip binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/toolTip
	ToolTipBinding string
	// TransparentBinding is a constant that identifies a transparent binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/transparent
	TransparentBinding string
	// See: https://developer.apple.com/documentation/AppKit/NSTypeIdentifierAddressText
	TypeIdentifierAddressText string
	// See: https://developer.apple.com/documentation/AppKit/NSTypeIdentifierDateText
	TypeIdentifierDateText string
	// See: https://developer.apple.com/documentation/AppKit/NSTypeIdentifierPhoneNumberText
	TypeIdentifierPhoneNumberText string
	// See: https://developer.apple.com/documentation/AppKit/NSTypeIdentifierTransitInformationText
	TypeIdentifierTransitInformationText string
	// UserActivityDocumentURLKey is the key that identifies the document associated with a user activity.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSUserActivityDocumentURLKey
	UserActivityDocumentURLKey string
	// ValueBinding is a constant that identifies a value binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/value
	ValueBinding string
	// ValuePathBinding is a constant that identifies a value path binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/valuePath
	ValuePathBinding string
	// ValueURLBinding is a constant that identifies a value URL binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/valueURL
	ValueURLBinding string
	// VisibleBinding is a constant that identifies a visible binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/visible
	VisibleBinding string
	// WarningValueBinding is a constant that identifies a warning value binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/warningValue
	WarningValueBinding string
	// WidthBinding is a constant that identifies a width binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingName/width
	WidthBinding string
	// WorkspaceApplicationKey is the value corresponding to this key is an instance of [NSRunningApplication] that reflects the affected app.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/applicationUserInfoKey
	WorkspaceApplicationKey string
	// WorkspaceVolumeLocalizedNameKey is a string containing the user-visible name of the volume.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/localizedVolumeNameUserInfoKey
	WorkspaceVolumeLocalizedNameKey string
	// WorkspaceVolumeOldLocalizedNameKey is a string containing the old user-visible name of the volume.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/oldLocalizedVolumeNameUserInfoKey
	WorkspaceVolumeOldLocalizedNameKey string
	// WorkspaceVolumeOldURLKey is a URL containing the old mount path of the volume.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/oldVolumeURLUserInfoKey
	WorkspaceVolumeOldURLKey string
	// WorkspaceVolumeURLKey is a URL containing the mount path of the volume.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/volumeURLUserInfoKey
	WorkspaceVolumeURLKey string
)

var (
	// AccessibilityAnnotationElement is the user interface element for the annotation.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/AnnotationAttributeKey/element
	AccessibilityAnnotationElement NSAccessibilityAnnotationAttributeKey
	// AccessibilityAnnotationLabel is a description of the annotation.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/AnnotationAttributeKey/label
	AccessibilityAnnotationLabel NSAccessibilityAnnotationAttributeKey
	// AccessibilityAnnotationLocation is the position where the annotation applies.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/AnnotationAttributeKey/location
	AccessibilityAnnotationLocation NSAccessibilityAnnotationAttributeKey
)

var (
	// AccessibilityAnnouncementKey is the announcement as a localized string.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/NotificationUserInfoKey/announcement
	AccessibilityAnnouncementKey NSAccessibilityNotificationUserInfoKey
	// AccessibilityPriorityKey is a priority level that can help an assistive app determine how to handle the corresponding notification.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/NotificationUserInfoKey/priority
	AccessibilityPriorityKey NSAccessibilityNotificationUserInfoKey
	// AccessibilityUIElementsKey is an array of elements for the notification.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/NotificationUserInfoKey/uiElements
	AccessibilityUIElementsKey NSAccessibilityNotificationUserInfoKey
)

var (
	// AccessibilityAnnouncementRequestedNotification is this notification posts when an app needs to make an announcement to the user. If VoiceOver is enabled, it’s presented via speech and/or braille. Otherwise, it does nothing.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/announcementRequested
	AccessibilityAnnouncementRequestedNotification NSAccessibilityNotificationName
	// AccessibilityApplicationActivatedNotification is this notification is posted after the app has been activated. Post this notification using the [post(element:notification:)] function instead of an [NSNotificationCenter] instance.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/applicationActivated
	AccessibilityApplicationActivatedNotification NSAccessibilityNotificationName
	// AccessibilityApplicationDeactivatedNotification is this notification is posted after the app has been deactivated. Post this notification using the [post(element:notification:)] function instead of an [NSNotificationCenter] instance.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/applicationDeactivated
	AccessibilityApplicationDeactivatedNotification NSAccessibilityNotificationName
	// AccessibilityApplicationHiddenNotification is this notification is posted after the app is hidden. Post this notification using the [post(element:notification:)] function instead of an [NSNotificationCenter] instance.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/applicationHidden
	AccessibilityApplicationHiddenNotification NSAccessibilityNotificationName
	// AccessibilityApplicationShownNotification is this notification is posted after the app is shown. Post this notification using the [post(element:notification:)] function instead of an [NSNotificationCenter] instance.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/applicationShown
	AccessibilityApplicationShownNotification NSAccessibilityNotificationName
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/NSAccessibilityAutocorrectionOccurred
	AccessibilityAutocorrectionOccurredNotification NSAccessibilityNotificationName
	// AccessibilityCreatedNotification is this notification is posted after an accessibility element is created. Post this notification using the [post(element:notification:)] function instead of an [NSNotificationCenter] instance.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/created
	AccessibilityCreatedNotification NSAccessibilityNotificationName
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/NSAccessibilityDraggingDestinationDragAccepted
	AccessibilityDraggingDestinationDragAcceptedNotification NSAccessibilityNotificationName
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/NSAccessibilityDraggingDestinationDragNotAccepted
	AccessibilityDraggingDestinationDragNotAcceptedNotification NSAccessibilityNotificationName
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/NSAccessibilityDraggingDestinationDropAllowed
	AccessibilityDraggingDestinationDropAllowedNotification NSAccessibilityNotificationName
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/NSAccessibilityDraggingDestinationDropNotAllowed
	AccessibilityDraggingDestinationDropNotAllowedNotification NSAccessibilityNotificationName
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/NSAccessibilityDraggingSourceDragBegan
	AccessibilityDraggingSourceDragBeganNotification NSAccessibilityNotificationName
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/NSAccessibilityDraggingSourceDragEnded
	AccessibilityDraggingSourceDragEndedNotification NSAccessibilityNotificationName
	// AccessibilityDrawerCreatedNotification is this notification is posted after a drawer appears. Post this notification using the [post(element:notification:)] function instead of an [NSNotificationCenter] instance.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/drawerCreated
	AccessibilityDrawerCreatedNotification NSAccessibilityNotificationName
	// AccessibilityFocusedUIElementChangedNotification is this notification is posted after an accessibility element gains focus. Post this notification using the [post(element:notification:)] function instead of an [NSNotificationCenter] instance.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/focusedUIElementChanged
	AccessibilityFocusedUIElementChangedNotification NSAccessibilityNotificationName
	// AccessibilityFocusedWindowChangedNotification is this notification is posted after the key window changes. Post this notification using the [post(element:notification:)] function instead of an [NSNotificationCenter] instance.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/focusedWindowChanged
	AccessibilityFocusedWindowChangedNotification NSAccessibilityNotificationName
	// AccessibilityHelpTagCreatedNotification is this notification is posted after a help tag appears. Post this notification using the [post(element:notification:)] function instead of an [NSNotificationCenter] instance.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/helpTagCreated
	AccessibilityHelpTagCreatedNotification NSAccessibilityNotificationName
	// AccessibilityLayoutChangedNotification is this notification is posted after the UI changes in a way that requires the attention of an accessibility client. This notification should be accompanied by a `userInfo` dictionary with the key [uiElements] and an array containing the UI elements that have been added or changed. Post this notification using the [post(element:notification:)] function instead of an [NSNotificationCenter] instance.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/layoutChanged
	AccessibilityLayoutChangedNotification NSAccessibilityNotificationName
	// AccessibilityMainWindowChangedNotification is this notification is posted after the main window changes. Post this notification using the [post(element:notification:)] function instead of an [NSNotificationCenter] instance.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/mainWindowChanged
	AccessibilityMainWindowChangedNotification NSAccessibilityNotificationName
	// AccessibilityMovedNotification is this notification is posted after an accessibility element moves. Post this notification using the [post(element:notification:)] function instead of an [NSNotificationCenter] instance.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/moved
	AccessibilityMovedNotification NSAccessibilityNotificationName
	// AccessibilityResizedNotification is this notification is posted after an accessibility element’s size changes. Post this notification using the [post(element:notification:)] function instead of an [NSNotificationCenter] instance.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/resized
	AccessibilityResizedNotification NSAccessibilityNotificationName
	// AccessibilityRowCollapsedNotification is this notification is posted after a row collapses. Post this notification using the [post(element:notification:)] function instead of an [NSNotificationCenter] instance.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/rowCollapsed
	AccessibilityRowCollapsedNotification NSAccessibilityNotificationName
	// AccessibilityRowCountChangedNotification is this notification is posted after a row is added or deleted. Post this notification using the [post(element:notification:)] function instead of an [NSNotificationCenter] instance.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/rowCountChanged
	AccessibilityRowCountChangedNotification NSAccessibilityNotificationName
	// AccessibilityRowExpandedNotification is this notification is posted after a row expands. Post this notification using the [post(element:notification:)] function instead of an [NSNotificationCenter] instance.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/rowExpanded
	AccessibilityRowExpandedNotification NSAccessibilityNotificationName
	// AccessibilitySelectedCellsChangedNotification is this notification is posted after one or more cells in a cell-based table are selected or deselected. Post this notification using the [post(element:notification:)] function instead of an [NSNotificationCenter] instance.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/selectedCellsChanged
	AccessibilitySelectedCellsChangedNotification NSAccessibilityNotificationName
	// AccessibilitySelectedChildrenChangedNotification is this notification is posted after one or more child elements are selected or deselected. Post this notification using the [post(element:notification:)] function instead of an [NSNotificationCenter] instance.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/selectedChildrenChanged
	AccessibilitySelectedChildrenChangedNotification NSAccessibilityNotificationName
	// AccessibilitySelectedChildrenMovedNotification is this notification is posted after the selected items in a layout area move. Post this notification using the [post(element:notification:)] function instead of an [NSNotificationCenter] instance.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/selectedChildrenMoved
	AccessibilitySelectedChildrenMovedNotification NSAccessibilityNotificationName
	// AccessibilitySelectedColumnsChangedNotification is this notification is posted after one or more columns are selected or deselected. Post this notification using the [post(element:notification:)] function instead of an [NSNotificationCenter] instance.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/selectedColumnsChanged
	AccessibilitySelectedColumnsChangedNotification NSAccessibilityNotificationName
	// AccessibilitySelectedRowsChangedNotification is this notification is posted after one or more rows are selected or deselected. Post this notification using the [post(element:notification:)] function instead of an [NSNotificationCenter] instance.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/selectedRowsChanged
	AccessibilitySelectedRowsChangedNotification NSAccessibilityNotificationName
	// AccessibilitySelectedTextChangedNotification is this notification is posted after text is selected or deselected. Post this notification using the [post(element:notification:)] function instead of an [NSNotificationCenter] instance.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/selectedTextChanged
	AccessibilitySelectedTextChangedNotification NSAccessibilityNotificationName
	// AccessibilitySheetCreatedNotification is this notification is posted after a sheet appears. Post this notification using the [post(element:notification:)] function instead of an [NSNotificationCenter] instance.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/sheetCreated
	AccessibilitySheetCreatedNotification NSAccessibilityNotificationName
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/NSAccessibilityTextInputMarkingSessionBegan
	AccessibilityTextInputMarkingSessionBeganNotification NSAccessibilityNotificationName
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/NSAccessibilityTextInputMarkingSessionEnded
	AccessibilityTextInputMarkingSessionEndedNotification NSAccessibilityNotificationName
	// AccessibilityTitleChangedNotification is this notification is posted after an accessibility element’s title changes. Post this notification using the [post(element:notification:)] function instead of an [NSNotificationCenter] instance.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/titleChanged
	AccessibilityTitleChangedNotification NSAccessibilityNotificationName
	// AccessibilityUIElementDestroyedNotification is this notification is posted after an accessibility element is destroyed. Post this notification using the [post(element:notification:)] function instead of an [NSNotificationCenter] instance.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/uiElementDestroyed
	AccessibilityUIElementDestroyedNotification NSAccessibilityNotificationName
	// AccessibilityUnitsChangedNotification is this notification is posted after the units in a layout area change. Post this notification using the [post(element:notification:)] function instead of an [NSNotificationCenter] instance.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/unitsChanged
	AccessibilityUnitsChangedNotification NSAccessibilityNotificationName
	// AccessibilityValueChangedNotification is this notification is posted after an accessibility element’s value changes. Post this notification using the [post(element:notification:)] function instead of an [NSNotificationCenter] instance.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/valueChanged
	AccessibilityValueChangedNotification NSAccessibilityNotificationName
	// AccessibilityWindowCreatedNotification is this notification is posted after a new window appears. Post this notification using the [post(element:notification:)] function instead of an [NSNotificationCenter] instance.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/windowCreated
	AccessibilityWindowCreatedNotification NSAccessibilityNotificationName
	// AccessibilityWindowDeminiaturizedNotification is this notification is posted after a window is restored to full size from the Dock. Post this notification using the [post(element:notification:)] function instead of an [NSNotificationCenter] instance.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/windowDeminiaturized
	AccessibilityWindowDeminiaturizedNotification NSAccessibilityNotificationName
	// AccessibilityWindowMiniaturizedNotification is this notification is posted after a window is put in the Dock. Post this notification using the [post(element:notification:)] function instead of an [NSNotificationCenter] instance.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/windowMiniaturized
	AccessibilityWindowMiniaturizedNotification NSAccessibilityNotificationName
	// AccessibilityWindowMovedNotification is this notification is posted after a window moves. Post this notification using the [post(element:notification:)] function instead of an [NSNotificationCenter] instance.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/windowMoved
	AccessibilityWindowMovedNotification NSAccessibilityNotificationName
	// AccessibilityWindowResizedNotification is this notification is posted after a window’s size changes. Post this notification using the [post(element:notification:)] function instead of an [NSNotificationCenter] instance.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/windowResized
	AccessibilityWindowResizedNotification NSAccessibilityNotificationName
)

var (
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySearchKey/anyTypeSearchKey
	AccessibilityAnyTypeSearchKey NSAccessibilitySearchKey
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySearchKey/articleSearchKey
	AccessibilityArticleSearchKey NSAccessibilitySearchKey
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySearchKey/blockquoteSameLevelSearchKey
	AccessibilityBlockquoteSameLevelSearchKey NSAccessibilitySearchKey
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySearchKey/blockquoteSearchKey
	AccessibilityBlockquoteSearchKey NSAccessibilitySearchKey
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySearchKey/boldFontSearchKey
	AccessibilityBoldFontSearchKey NSAccessibilitySearchKey
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySearchKey/buttonSearchKey
	AccessibilityButtonSearchKey NSAccessibilitySearchKey
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySearchKey/checkBoxSearchKey
	AccessibilityCheckBoxSearchKey NSAccessibilitySearchKey
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySearchKey/controlSearchKey
	AccessibilityControlSearchKey NSAccessibilitySearchKey
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySearchKey/differentTypeSearchKey
	AccessibilityDifferentTypeSearchKey NSAccessibilitySearchKey
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySearchKey/fontChangeSearchKey
	AccessibilityFontChangeSearchKey NSAccessibilitySearchKey
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySearchKey/fontColorChangeSearchKey
	AccessibilityFontColorChangeSearchKey NSAccessibilitySearchKey
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySearchKey/frameSearchKey
	AccessibilityFrameSearchKey NSAccessibilitySearchKey
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySearchKey/graphicSearchKey
	AccessibilityGraphicSearchKey NSAccessibilitySearchKey
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySearchKey/headingLevel1SearchKey
	AccessibilityHeadingLevel1SearchKey NSAccessibilitySearchKey
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySearchKey/headingLevel2SearchKey
	AccessibilityHeadingLevel2SearchKey NSAccessibilitySearchKey
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySearchKey/headingLevel3SearchKey
	AccessibilityHeadingLevel3SearchKey NSAccessibilitySearchKey
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySearchKey/headingLevel4SearchKey
	AccessibilityHeadingLevel4SearchKey NSAccessibilitySearchKey
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySearchKey/headingLevel5SearchKey
	AccessibilityHeadingLevel5SearchKey NSAccessibilitySearchKey
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySearchKey/headingLevel6SearchKey
	AccessibilityHeadingLevel6SearchKey NSAccessibilitySearchKey
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySearchKey/headingSameLevelSearchKey
	AccessibilityHeadingSameLevelSearchKey NSAccessibilitySearchKey
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySearchKey/headingSearchKey
	AccessibilityHeadingSearchKey NSAccessibilitySearchKey
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySearchKey/italicFontSearchKey
	AccessibilityItalicFontSearchKey NSAccessibilitySearchKey
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySearchKey/keyboardFocusableSearchKey
	AccessibilityKeyboardFocusableSearchKey NSAccessibilitySearchKey
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySearchKey/landmarkSearchKey
	AccessibilityLandmarkSearchKey NSAccessibilitySearchKey
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySearchKey/linkSearchKey
	AccessibilityLinkSearchKey NSAccessibilitySearchKey
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySearchKey/listSearchKey
	AccessibilityListSearchKey NSAccessibilitySearchKey
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySearchKey/liveRegionSearchKey
	AccessibilityLiveRegionSearchKey NSAccessibilitySearchKey
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySearchKey/misspelledWordSearchKey
	AccessibilityMisspelledWordSearchKey NSAccessibilitySearchKey
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySearchKey/outlineSearchKey
	AccessibilityOutlineSearchKey NSAccessibilitySearchKey
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySearchKey/plainTextSearchKey
	AccessibilityPlainTextSearchKey NSAccessibilitySearchKey
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySearchKey/radioGroupSearchKey
	AccessibilityRadioGroupSearchKey NSAccessibilitySearchKey
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySearchKey/sameTypeSearchKey
	AccessibilitySameTypeSearchKey NSAccessibilitySearchKey
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySearchKey/staticTextSearchKey
	AccessibilityStaticTextSearchKey NSAccessibilitySearchKey
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySearchKey/styleChangeSearchKey
	AccessibilityStyleChangeSearchKey NSAccessibilitySearchKey
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySearchKey/tableSameLevelSearchKey
	AccessibilityTableSameLevelSearchKey NSAccessibilitySearchKey
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySearchKey/tableSearchKey
	AccessibilityTableSearchKey NSAccessibilitySearchKey
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySearchKey/textFieldSearchKey
	AccessibilityTextFieldSearchKey NSAccessibilitySearchKey
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySearchKey/textStateChangeTypeKey
	AccessibilityTextStateChangeTypeKey NSAccessibilitySearchKey
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySearchKey/textStateSyncKey
	AccessibilityTextStateSyncKey NSAccessibilitySearchKey
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySearchKey/underlineSearchKey
	AccessibilityUnderlineSearchKey NSAccessibilitySearchKey
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySearchKey/unvisitedLinkSearchKey
	AccessibilityUnvisitedLinkSearchKey NSAccessibilitySearchKey
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySearchKey/visitedLinkSearchKey
	AccessibilityVisitedLinkSearchKey NSAccessibilitySearchKey
)

var (
	// AccessibilityApplicationRole is the app role.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/application
	AccessibilityApplicationRole NSAccessibilityRole
	// AccessibilityBrowserRole is the browser role.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/browser
	AccessibilityBrowserRole NSAccessibilityRole
	// AccessibilityBusyIndicatorRole is the busy indicator role.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/busyIndicator
	AccessibilityBusyIndicatorRole NSAccessibilityRole
	// AccessibilityButtonRole is the button role.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/button
	AccessibilityButtonRole NSAccessibilityRole
	// AccessibilityCellRole is the cell role in a table or matrix.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/cell
	AccessibilityCellRole NSAccessibilityRole
	// AccessibilityCheckBoxRole is the checkbox role.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/checkBox
	AccessibilityCheckBoxRole NSAccessibilityRole
	// AccessibilityColorWellRole is the color well role.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/colorWell
	AccessibilityColorWellRole NSAccessibilityRole
	// AccessibilityColumnRole is the column role.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/column
	AccessibilityColumnRole NSAccessibilityRole
	// AccessibilityComboBoxRole is the combo box role.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/comboBox
	AccessibilityComboBoxRole NSAccessibilityRole
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/dateTimeAreaRole
	AccessibilityDateTimeAreaRole NSAccessibilityRole
	// AccessibilityDisclosureTriangleRole is the disclosure triangle role.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/disclosureTriangle
	AccessibilityDisclosureTriangleRole NSAccessibilityRole
	// AccessibilityDrawerRole is the drawer role.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/drawer
	AccessibilityDrawerRole NSAccessibilityRole
	// AccessibilityGridRole is the grid role.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/grid
	AccessibilityGridRole NSAccessibilityRole
	// AccessibilityGroupRole is the group role.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/group
	AccessibilityGroupRole NSAccessibilityRole
	// AccessibilityGrowAreaRole is the grow (resize) area role in a window.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/growArea
	AccessibilityGrowAreaRole NSAccessibilityRole
	// AccessibilityHandleRole is the drag handle role.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/handle
	AccessibilityHandleRole NSAccessibilityRole
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/headingRole
	AccessibilityHeadingRole NSAccessibilityRole
	// AccessibilityHelpTagRole is the help tag role.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/helpTag
	AccessibilityHelpTagRole NSAccessibilityRole
	// AccessibilityImageRole is the image role.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/image
	AccessibilityImageRole NSAccessibilityRole
	// AccessibilityIncrementorRole is the stepper role.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/incrementor
	AccessibilityIncrementorRole NSAccessibilityRole
	// AccessibilityLayoutAreaRole is the layout area role (a view, such as a graphic view, that contains visual elements that may not have any accessibility representation).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/layoutArea
	AccessibilityLayoutAreaRole NSAccessibilityRole
	// AccessibilityLayoutItemRole is the role for an item in a layout area.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/layoutItem
	AccessibilityLayoutItemRole NSAccessibilityRole
	// AccessibilityLevelIndicatorRole is the level indicator role.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/levelIndicator
	AccessibilityLevelIndicatorRole NSAccessibilityRole
	// AccessibilityLinkRole is the link role.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/link
	AccessibilityLinkRole NSAccessibilityRole
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/listMarkerRole
	AccessibilityListMarkerRole NSAccessibilityRole
	// AccessibilityListRole is the list role.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/list
	AccessibilityListRole NSAccessibilityRole
	// AccessibilityMatteRole is the matte role.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/matte
	AccessibilityMatteRole NSAccessibilityRole
	// AccessibilityMenuBarItemRole is the menu bar item role.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/menuBarItem
	AccessibilityMenuBarItemRole NSAccessibilityRole
	// AccessibilityMenuBarRole is the menu bar role.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/menuBar
	AccessibilityMenuBarRole NSAccessibilityRole
	// AccessibilityMenuButtonRole is the menu button role.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/menuButton
	AccessibilityMenuButtonRole NSAccessibilityRole
	// AccessibilityMenuItemRole is the menu item role.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/menuItem
	AccessibilityMenuItemRole NSAccessibilityRole
	// AccessibilityMenuRole is the menu role.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/menu
	AccessibilityMenuRole NSAccessibilityRole
	// AccessibilityOutlineRole is the outline role.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/outline
	AccessibilityOutlineRole NSAccessibilityRole
	// AccessibilityPageRole is the page role.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/pageRole
	AccessibilityPageRole NSAccessibilityRole
	// AccessibilityPopUpButtonRole is the pop-up button role.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/popUpButton
	AccessibilityPopUpButtonRole NSAccessibilityRole
	// AccessibilityPopoverRole is the popover role.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/popover
	AccessibilityPopoverRole NSAccessibilityRole
	// AccessibilityProgressIndicatorRole is the progress indicator role.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/progressIndicator
	AccessibilityProgressIndicatorRole NSAccessibilityRole
	// AccessibilityRadioButtonRole is the radio button role.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/radioButton
	AccessibilityRadioButtonRole NSAccessibilityRole
	// AccessibilityRadioGroupRole is the radio button group role.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/radioGroup
	AccessibilityRadioGroupRole NSAccessibilityRole
	// AccessibilityRelevanceIndicatorRole is the relevance indicator role.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/relevanceIndicator
	AccessibilityRelevanceIndicatorRole NSAccessibilityRole
	// AccessibilityRowRole is the row role.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/row
	AccessibilityRowRole NSAccessibilityRole
	// AccessibilityRulerMarkerRole is the ruler marker role.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/rulerMarker
	AccessibilityRulerMarkerRole NSAccessibilityRole
	// AccessibilityRulerRole is the ruler role.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/ruler
	AccessibilityRulerRole NSAccessibilityRole
	// AccessibilityScrollAreaRole is the scroll view role.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/scrollArea
	AccessibilityScrollAreaRole NSAccessibilityRole
	// AccessibilityScrollBarRole is the scroll bar role.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/scrollBar
	AccessibilityScrollBarRole NSAccessibilityRole
	// AccessibilitySheetRole is the sheet role.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/sheet
	AccessibilitySheetRole NSAccessibilityRole
	// AccessibilitySliderRole is the slider role.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/slider
	AccessibilitySliderRole NSAccessibilityRole
	// AccessibilitySplitGroupRole is the split view role.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/splitGroup
	AccessibilitySplitGroupRole NSAccessibilityRole
	// AccessibilitySplitterRole is the splitter bar role for a split view.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/splitter
	AccessibilitySplitterRole NSAccessibilityRole
	// AccessibilityStaticTextRole is the uneditable text role.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/staticText
	AccessibilityStaticTextRole NSAccessibilityRole
	// AccessibilitySystemWideRole is the systemwide accessibility object role.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/systemWide
	AccessibilitySystemWideRole NSAccessibilityRole
	// AccessibilityTabGroupRole is the tab group role.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/tabGroup
	AccessibilityTabGroupRole NSAccessibilityRole
	// AccessibilityTableRole is the table role.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/table
	AccessibilityTableRole NSAccessibilityRole
	// AccessibilityTextAreaRole is the text view role.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/textArea
	AccessibilityTextAreaRole NSAccessibilityRole
	// AccessibilityTextFieldRole is the text field role.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/textField
	AccessibilityTextFieldRole NSAccessibilityRole
	// AccessibilityToolbarRole is the toolbar role.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/toolbar
	AccessibilityToolbarRole NSAccessibilityRole
	// AccessibilityUnknownRole is an object with an unknown role.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/unknown
	AccessibilityUnknownRole NSAccessibilityRole
	// AccessibilityValueIndicatorRole is the value indicator role.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/valueIndicator
	AccessibilityValueIndicatorRole NSAccessibilityRole
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/webAreaRole
	AccessibilityWebAreaRole NSAccessibilityRole
	// AccessibilityWindowRole is the window role.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/window
	AccessibilityWindowRole NSAccessibilityRole
)

var (
	// AccessibilityAttributedStringForRangeParameterizedAttribute is does not use attributes from Appkit/AttributedString.h ([NSAttributedString]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/ParameterizedAttribute/attributedStringForRange
	AccessibilityAttributedStringForRangeParameterizedAttribute NSAccessibilityParameterizedAttributeName
	// AccessibilityBoundsForRangeParameterizedAttribute is the rectangle ([NSValue] containing an [NSRect] value) enclosing the specified range of characters ([NSValue] containing an [NSRange] value). If the range crosses a line boundary, the returned rectangle will fully enclose all the lines of characters.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/ParameterizedAttribute/boundsForRange
	AccessibilityBoundsForRangeParameterizedAttribute NSAccessibilityParameterizedAttributeName
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/ParameterizedAttribute/indexForChildUIElementAttribute
	AccessibilityIndexForChildUIElementAttribute NSAccessibilityParameterizedAttributeName
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/ParameterizedAttribute/indexForChildUIElementInNavigationOrderAttribute
	AccessibilityIndexForChildUIElementInNavigationOrderAttribute NSAccessibilityParameterizedAttributeName
	// AccessibilityLayoutSizeForScreenSizeParameterizedAttribute is the size of the layout area in points ([NSValue]) corresponding to the specified screen size ([NSValue]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/ParameterizedAttribute/layoutSizeForScreenSize
	AccessibilityLayoutSizeForScreenSizeParameterizedAttribute NSAccessibilityParameterizedAttributeName
	// AccessibilityRTFForRangeParameterizedAttribute is the RTF data ([NSData]) describing the specified range of characters ([NSValue] containing an [NSRange] value).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/ParameterizedAttribute/rtfForRange
	AccessibilityRTFForRangeParameterizedAttribute NSAccessibilityParameterizedAttributeName
	// AccessibilityRangeForIndexParameterizedAttribute is the full range of characters ([NSValue] containing an [NSRange] value), including the specified character, which compose a single glyph ([NSNumber]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/ParameterizedAttribute/rangeForIndex
	AccessibilityRangeForIndexParameterizedAttribute NSAccessibilityParameterizedAttributeName
	// AccessibilityRangeForLineParameterizedAttribute is the range of characters ([NSValue] containing an [NSRange] value) corresponding to the specified line number ([NSNumber]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/ParameterizedAttribute/rangeForLine
	AccessibilityRangeForLineParameterizedAttribute NSAccessibilityParameterizedAttributeName
	// AccessibilityRangeForPositionParameterizedAttribute is the range of characters ([NSValue] containing an [NSRange] value) composing the glyph at the specified point ([NSValue] containing an [NSPoint] value).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/ParameterizedAttribute/rangeForPosition
	AccessibilityRangeForPositionParameterizedAttribute NSAccessibilityParameterizedAttributeName
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/ParameterizedAttribute/resultsForSearchPredicateParameterizedAttribute
	AccessibilityResultsForSearchPredicateParameterizedAttribute NSAccessibilityParameterizedAttributeName
	// AccessibilityScreenPointForLayoutPointParameterizedAttribute is the screen point ([NSValue]) corresponding to the specified point in the layout area ([NSValue]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/ParameterizedAttribute/screenPointForLayoutPoint
	AccessibilityScreenPointForLayoutPointParameterizedAttribute NSAccessibilityParameterizedAttributeName
	// AccessibilityScreenSizeForLayoutSizeParameterizedAttribute is the size of the screen in points ([NSValue]) corresponding to the specified size of the layout area ([NSValue]).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/ParameterizedAttribute/screenSizeForLayoutSize
	AccessibilityScreenSizeForLayoutSizeParameterizedAttribute NSAccessibilityParameterizedAttributeName
	// AccessibilityStringForRangeParameterizedAttribute is the substring ([NSString]) specified by the range ([NSValue] containing an [NSRange] value).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/ParameterizedAttribute/stringForRange
	AccessibilityStringForRangeParameterizedAttribute NSAccessibilityParameterizedAttributeName
	// AccessibilityStyleRangeForIndexParameterizedAttribute is the full range of characters ([NSValue] containing an [NSRange] value), including the specified character ([NSNumber]), which have the same style.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/ParameterizedAttribute/styleRangeForIndex
	AccessibilityStyleRangeForIndexParameterizedAttribute NSAccessibilityParameterizedAttributeName
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/ParameterizedAttribute/uiElementsForSearchPredicateParameterizedAttribute
	AccessibilityUIElementsForSearchPredicateParameterizedAttribute NSAccessibilityParameterizedAttributeName
)

var (
	// AccessibilityCancelAction is an action that cancels the operation.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Action/cancel
	AccessibilityCancelAction NSAccessibilityActionName
	// AccessibilityDecrementAction is an action that decrements the value of the object.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Action/decrement
	AccessibilityDecrementAction NSAccessibilityActionName
	// AccessibilityDeleteAction is an action that deletes the value of the object.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Action/delete
	AccessibilityDeleteAction NSAccessibilityActionName
	// AccessibilityIncrementAction is an action that increments the value of the object.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Action/increment
	AccessibilityIncrementAction NSAccessibilityActionName
	// AccessibilityPickAction is an action that selects the object, such as a menu item.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Action/pick
	AccessibilityPickAction NSAccessibilityActionName
	// AccessibilityPressAction is an action that simulates clicking an object, such as a button.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Action/press
	AccessibilityPressAction NSAccessibilityActionName
	// AccessibilityRaiseAction is an action that simulates bringing a window forward by clicking on its title bar.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Action/raise
	AccessibilityRaiseAction NSAccessibilityActionName
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Action/scrollToVisibleAction
	AccessibilityScrollToVisibleAction NSAccessibilityActionName
	// AccessibilityShowAlternateUIAction is an action that shows an alternate UI, for example, during a mouse-hover event.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Action/showAlternateUI
	AccessibilityShowAlternateUIAction NSAccessibilityActionName
	// AccessibilityShowDefaultUIAction is an action that shows the original or default UI; for example, during a mouse-hover event.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Action/showDefaultUI
	AccessibilityShowDefaultUIAction NSAccessibilityActionName
	// AccessibilityShowMenuAction is an action that simulates showing a menu by clicking on it.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Action/showMenu
	AccessibilityShowMenuAction NSAccessibilityActionName
)

var (
	// AccessibilityCenterTabStopMarkerTypeValue is center tab stop.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/RulerMarkerTypeValue/centerTabStop
	AccessibilityCenterTabStopMarkerTypeValue NSAccessibilityRulerMarkerTypeValue
	// AccessibilityDecimalTabStopMarkerTypeValue is decimal tab stop.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/RulerMarkerTypeValue/decimalTabStop
	AccessibilityDecimalTabStopMarkerTypeValue NSAccessibilityRulerMarkerTypeValue
	// AccessibilityFirstLineIndentMarkerTypeValue is first line indent marker.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/RulerMarkerTypeValue/firstLineIndent
	AccessibilityFirstLineIndentMarkerTypeValue NSAccessibilityRulerMarkerTypeValue
	// AccessibilityHeadIndentMarkerTypeValue is head indent marker.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/RulerMarkerTypeValue/headIndent
	AccessibilityHeadIndentMarkerTypeValue NSAccessibilityRulerMarkerTypeValue
	// AccessibilityRightTabStopMarkerTypeValue is right tab stop.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/RulerMarkerTypeValue/rightTabStop
	AccessibilityRightTabStopMarkerTypeValue NSAccessibilityRulerMarkerTypeValue
	// AccessibilityTailIndentMarkerTypeValue is tail indent marker.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/RulerMarkerTypeValue/tailIndent
	AccessibilityTailIndentMarkerTypeValue NSAccessibilityRulerMarkerTypeValue
	// AccessibilityUnknownMarkerTypeValue is unknown marker type.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/RulerMarkerTypeValue/unknown
	AccessibilityUnknownMarkerTypeValue NSAccessibilityRulerMarkerTypeValue
)

var (
	// AccessibilityCentimetersUnitValue is the units are centimeters.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/RulerUnitValue/centimeters
	AccessibilityCentimetersUnitValue NSAccessibilityRulerUnitValue
	// AccessibilityPicasUnitValue is the units are picas.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/RulerUnitValue/picas
	AccessibilityPicasUnitValue NSAccessibilityRulerUnitValue
	// AccessibilityPointsUnitValue is the units are points.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/RulerUnitValue/points
	AccessibilityPointsUnitValue NSAccessibilityRulerUnitValue
	// AccessibilityUnknownUnitValue is the units are unknown.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/RulerUnitValue/unknown
	AccessibilityUnknownUnitValue NSAccessibilityRulerUnitValue
)

var (
	// AccessibilityCloseButtonSubrole is a window’s close button subrole.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Subrole/closeButton
	AccessibilityCloseButtonSubrole NSAccessibilitySubrole
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Subrole/collectionListSubrole
	AccessibilityCollectionListSubrole NSAccessibilitySubrole
	// AccessibilityContentListSubrole is a subrole for content that is organized in a list, but is not in a list control or table view.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Subrole/contentList
	AccessibilityContentListSubrole NSAccessibilitySubrole
	// AccessibilityDecrementArrowSubrole is a decrement arrow subrole (the down arrow in a scroll bar).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Subrole/decrementArrow
	AccessibilityDecrementArrowSubrole NSAccessibilitySubrole
	// AccessibilityDecrementPageSubrole is a decrement page subrole (the decrement area in the scroll track of a scroll bar).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Subrole/decrementPage
	AccessibilityDecrementPageSubrole NSAccessibilitySubrole
	// AccessibilityDefinitionListSubrole is a subrole for a content list in a webpage.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Subrole/definitionList
	AccessibilityDefinitionListSubrole NSAccessibilitySubrole
	// AccessibilityDescriptionListSubrole is a description list subrole.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Subrole/descriptionList
	AccessibilityDescriptionListSubrole NSAccessibilitySubrole
	// AccessibilityDialogSubrole is a dialog subrole.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Subrole/dialog
	AccessibilityDialogSubrole NSAccessibilitySubrole
	// AccessibilityFloatingWindowSubrole is a floating window subrole.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Subrole/floatingWindow
	AccessibilityFloatingWindowSubrole NSAccessibilitySubrole
	// AccessibilityFullScreenButtonSubrole is a window’s full-screen button subrole.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Subrole/fullScreenButton
	AccessibilityFullScreenButtonSubrole NSAccessibilitySubrole
	// AccessibilityIncrementArrowSubrole is an increment arrow subrole (the up arrow in a scroll bar).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Subrole/incrementArrow
	AccessibilityIncrementArrowSubrole NSAccessibilitySubrole
	// AccessibilityIncrementPageSubrole is an increment page subrole (the increment area in the scroll track of a scroll bar).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Subrole/incrementPage
	AccessibilityIncrementPageSubrole NSAccessibilitySubrole
	// AccessibilityMinimizeButtonSubrole is a window’s minimize button subrole.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Subrole/minimizeButton
	AccessibilityMinimizeButtonSubrole NSAccessibilitySubrole
	// AccessibilityOutlineRowSubrole is an outline row subrole.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Subrole/outlineRow
	AccessibilityOutlineRowSubrole NSAccessibilitySubrole
	// AccessibilityRatingIndicatorSubrole is a rating indicator subrole.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Subrole/ratingIndicator
	AccessibilityRatingIndicatorSubrole NSAccessibilitySubrole
	// AccessibilitySearchFieldSubrole is a search field subrole.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Subrole/searchField
	AccessibilitySearchFieldSubrole NSAccessibilitySubrole
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Subrole/sectionListSubrole
	AccessibilitySectionListSubrole NSAccessibilitySubrole
	// AccessibilitySecureTextFieldSubrole is a secure text field subrole.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Subrole/secureTextField
	AccessibilitySecureTextFieldSubrole NSAccessibilitySubrole
	// AccessibilitySortButtonSubrole is a sort button subrole for a table or outline view.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Subrole/sortButton
	AccessibilitySortButtonSubrole NSAccessibilitySubrole
	// AccessibilityStandardWindowSubrole is a standard window subrole.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Subrole/standardWindow
	AccessibilityStandardWindowSubrole NSAccessibilitySubrole
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Subrole/suggestionSubrole
	AccessibilitySuggestionSubrole NSAccessibilitySubrole
	// AccessibilitySwitchSubrole is a switch subrole.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Subrole/switch
	AccessibilitySwitchSubrole NSAccessibilitySubrole
	// AccessibilitySystemDialogSubrole is a system dialog subrole (a system-generated dialog that floats on the top layer, regardless of which app is frontmost).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Subrole/systemDialog
	AccessibilitySystemDialogSubrole NSAccessibilitySubrole
	// AccessibilitySystemFloatingWindowSubrole is a system floating window subrole (a system-generated panel).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Subrole/systemFloatingWindow
	AccessibilitySystemFloatingWindowSubrole NSAccessibilitySubrole
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Subrole/tabButtonSubrole
	AccessibilityTabButtonSubrole NSAccessibilitySubrole
	// AccessibilityTableRowSubrole is a table row subrole.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Subrole/tableRow
	AccessibilityTableRowSubrole NSAccessibilitySubrole
	// AccessibilityTextAttachmentSubrole is a text attachment subrole.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Subrole/textAttachment
	AccessibilityTextAttachmentSubrole NSAccessibilitySubrole
	// AccessibilityTextLinkSubrole is a text link subrole.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Subrole/textLink
	AccessibilityTextLinkSubrole NSAccessibilitySubrole
	// AccessibilityTimelineSubrole is a timeline subrole.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Subrole/timeline
	AccessibilityTimelineSubrole NSAccessibilitySubrole
	// AccessibilityToggleSubrole is a toggle subrole.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Subrole/toggle
	AccessibilityToggleSubrole NSAccessibilitySubrole
	// AccessibilityToolbarButtonSubrole is a window’s toolbar button subrole.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Subrole/toolbarButton
	AccessibilityToolbarButtonSubrole NSAccessibilitySubrole
	// AccessibilityUnknownSubrole is an unknown subrole.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Subrole/unknown
	AccessibilityUnknownSubrole NSAccessibilitySubrole
	// AccessibilityZoomButtonSubrole is a window’s zoom button subrole.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Subrole/zoomButton
	AccessibilityZoomButtonSubrole NSAccessibilitySubrole
)

var (
	// AccessibilityDescendingSortDirectionValue is the column is sorted in descending values.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/SortDirectionValue/descending
	AccessibilityDescendingSortDirectionValue NSAccessibilitySortDirectionValue
	// AccessibilityUnknownSortDirectionValue is the sort direction is unknown.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/SortDirectionValue/unknown
	AccessibilityUnknownSortDirectionValue NSAccessibilitySortDirectionValue
)

var (
	// AccessibilityFontFamilyKey is an optional key for a font family.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/FontAttributeKey/fontFamily
	AccessibilityFontFamilyKey NSAccessibilityFontAttributeKey
	// AccessibilityFontNameKey is a required key for a font name.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/FontAttributeKey/fontName
	AccessibilityFontNameKey NSAccessibilityFontAttributeKey
	// AccessibilityFontSizeKey is a required key for a font size.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/FontAttributeKey/fontSize
	AccessibilityFontSizeKey NSAccessibilityFontAttributeKey
	// AccessibilityVisibleNameKey is an optional key for font visibility.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/FontAttributeKey/visibleName
	AccessibilityVisibleNameKey NSAccessibilityFontAttributeKey
)

var (
	// AccessibilityUnknownOrientationValue is the element has unknown orientation.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/OrientationValue/unknown
	AccessibilityUnknownOrientationValue NSAccessibilityOrientationValue
	// AccessibilityVerticalOrientationValue is the element is oriented vertically.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/OrientationValue/vertical
	AccessibilityVerticalOrientationValue NSAccessibilityOrientationValue
)

var (
	// AllowsEditingMultipleValuesSelectionBindingOption is an [NSNumber] object containing a Boolean value that determines if the binding allows editing when the value represents a multiple selection.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingOption/allowsEditingMultipleValuesSelection
	AllowsEditingMultipleValuesSelectionBindingOption NSBindingOption
	// AllowsNullArgumentBindingOption is an [NSNumber] object containing a Boolean value that determines if the argument bindings allows passing argument values of `nil`.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingOption/allowsNullArgument
	AllowsNullArgumentBindingOption NSBindingOption
	// AlwaysPresentsApplicationModalAlertsBindingOption is a number containing a Boolean value that determines if validation and error alert panels displayed as a result of this binding are displayed as application modal alerts.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingOption/alwaysPresentsApplicationModalAlerts
	AlwaysPresentsApplicationModalAlertsBindingOption NSBindingOption
	// ConditionallySetsEditableBindingOption is an [NSNumber] object containing a Boolean value that determines if the editable state of the user interface item is automatically configured based on the controller’s selection.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingOption/conditionallySetsEditable
	ConditionallySetsEditableBindingOption NSBindingOption
	// ConditionallySetsEnabledBindingOption is an [NSNumber] object containing a Boolean value that determines if the enabled state of the user interface item is automatically configured based on the controller’s selection.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingOption/conditionallySetsEnabled
	ConditionallySetsEnabledBindingOption NSBindingOption
	// ConditionallySetsHiddenBindingOption is an [NSNumber] object containing a Boolean value that determines if the hidden state of the user interface item is automatically configured based on the controller’s selection.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingOption/conditionallySetsHidden
	ConditionallySetsHiddenBindingOption NSBindingOption
	// ContentPlacementTagBindingOption is a number that specifies the tag id of the popup menu item to replace with the content of the array.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingOption/contentPlacementTag
	ContentPlacementTagBindingOption NSBindingOption
	// ContinuouslyUpdatesValueBindingOption is an [NSNumber] object containing a Boolean value that determines whether the value of the binding is updated as edits are made to the user interface item or is updated only when the user interface item resigns as the responder.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingOption/continuouslyUpdatesValue
	ContinuouslyUpdatesValueBindingOption NSBindingOption
	// CreatesSortDescriptorBindingOption is an [NSNumber] object containing a Boolean value that determines if a sort descriptor is created for a table column.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingOption/createsSortDescriptor
	CreatesSortDescriptorBindingOption NSBindingOption
	// DeletesObjectsOnRemoveBindingsOption is an [NSNumber] object containing a Boolean value that determines if an object is deleted from the managed context immediately upon being removed from a relationship.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingOption/deletesObjectsOnRemove
	DeletesObjectsOnRemoveBindingsOption NSBindingOption
	// DisplayNameBindingOption is an [NSString] object containing a human readable string to be displayed for a predicate.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingOption/displayName
	DisplayNameBindingOption NSBindingOption
	// DisplayPatternBindingOption is an [NSString] object that specifies a format string used to construct the final value of a string.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingOption/displayPattern
	DisplayPatternBindingOption NSBindingOption
	// HandlesContentAsCompoundValueBindingOption is an [NSNumber] object containing a Boolean value that determines if the content is treated as a compound value.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingOption/handlesContentAsCompoundValue
	HandlesContentAsCompoundValueBindingOption NSBindingOption
	// InsertsNullPlaceholderBindingOption is an [NSNumber] object containing a Boolean value that determines if an additional item which represents `nil` is inserted into a matrix or pop-up menu before the items in the content array.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingOption/insertsNullPlaceholder
	InsertsNullPlaceholderBindingOption NSBindingOption
	// InvokesSeparatelyWithArrayObjectsBindingOption is an [NSNumber] object containing a Boolean value that determines whether the specified selector is invoked with the array as the argument or is invoked repeatedly with each array item as an argument.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingOption/invokesSeparatelyWithArrayObjects
	InvokesSeparatelyWithArrayObjectsBindingOption NSBindingOption
	// MultipleValuesPlaceholderBindingOption is an object that is used as a placeholder when the key path of the bound controller returns the [NSMultipleValuesMarker] marker for a binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingOption/multipleValuesPlaceholder
	MultipleValuesPlaceholderBindingOption NSBindingOption
	// NoSelectionPlaceholderBindingOption is an object that is used as a placeholder when the key path of the bound controller returns the [NSNoSelectionMarker] marker for a binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingOption/noSelectionPlaceholder
	NoSelectionPlaceholderBindingOption NSBindingOption
	// NotApplicablePlaceholderBindingOption is an object that is used as a placeholder when the key path of the bound controller returns the [NSNotApplicableMarker] marker for a binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingOption/notApplicablePlaceholder
	NotApplicablePlaceholderBindingOption NSBindingOption
	// NullPlaceholderBindingOption is an object that is used as a placeholder when the key path of the bound controller returns `nil` for a binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingOption/nullPlaceholder
	NullPlaceholderBindingOption NSBindingOption
	// PredicateFormatBindingOption is an [NSString] object containing the predicate pattern string for the predicate bindings. Use `$value` to refer to the value in the search field.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingOption/predicateFormat
	PredicateFormatBindingOption NSBindingOption
	// RaisesForNotApplicableKeysBindingOption is an [NSNumber] object containing a Boolean value that specifies if an exception is raised when the binding is bound to a key that is not applicable—for example when an object is not key-value coding compliant for a key.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingOption/raisesForNotApplicableKeys
	RaisesForNotApplicableKeysBindingOption NSBindingOption
	// SelectorNameBindingOption is an [NSString] object that specifies the method selector invoked by the target binding when the user interface item is clicked.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingOption/selectorName
	SelectorNameBindingOption NSBindingOption
	// SelectsAllWhenSettingContentBindingOption is an [NSNumber] object containing a Boolean value that specifies if all the items in the array controller are selected when the content is set.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingOption/selectsAllWhenSettingContent
	SelectsAllWhenSettingContentBindingOption NSBindingOption
	// ValidatesImmediatelyBindingOption is an [NSNumber] object containing a Boolean value that determines if the contents of the binding are validated immediately.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingOption/validatesImmediately
	ValidatesImmediatelyBindingOption NSBindingOption
	// ValueTransformerBindingOption is an [NSValueTransformer] instance that is applied to the bound value.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingOption/valueTransformer
	ValueTransformerBindingOption NSBindingOption
	// ValueTransformerNameBindingOption is the value for this key is an identifier of a registered [NSValueTransformer] instance that is applied to the bound value.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingOption/valueTransformerName
	ValueTransformerNameBindingOption NSBindingOption
)

var (
	// See: https://developer.apple.com/documentation/AppKit/NSAnimationTriggerOrderIn
	AnimationTriggerOrderIn NSAnimatablePropertyKey
	// See: https://developer.apple.com/documentation/AppKit/NSAnimationTriggerOrderOut
	AnimationTriggerOrderOut NSAnimatablePropertyKey
)

var (
	// App is the global variable for the shared app instance.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSApp
	App NSApplication
)

var ()

var (
	// AppearanceDocumentAttribute is the appearance of the document.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAppearanceDocumentAttribute
	AppearanceDocumentAttribute NSAttributedStringDocumentAttributeKey
	// AuthorDocumentAttribute is the author of the document.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAuthorDocumentAttribute
	AuthorDocumentAttribute NSAttributedStringDocumentAttributeKey
	// BackgroundColorDocumentAttribute is the background color of the document.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBackgroundColorDocumentAttribute
	BackgroundColorDocumentAttribute NSAttributedStringDocumentAttributeKey
	// BottomMarginDocumentAttribute is the bottom margin of the document.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBottomMarginDocumentAttribute
	BottomMarginDocumentAttribute NSAttributedStringDocumentAttributeKey
	// CategoryDocumentAttribute is the document’s category.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSCategoryDocumentAttribute
	CategoryDocumentAttribute NSAttributedStringDocumentAttributeKey
	// CharacterEncodingDocumentAttribute is the string encoding for the document.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSCharacterEncodingDocumentAttribute
	CharacterEncodingDocumentAttribute NSAttributedStringDocumentAttributeKey
	// CocoaVersionDocumentAttribute is the version of Cocoa that created the file.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSCocoaVersionDocumentAttribute
	CocoaVersionDocumentAttribute NSAttributedStringDocumentAttributeKey
	// CommentDocumentAttribute is the document comments.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSCommentDocumentAttribute
	CommentDocumentAttribute NSAttributedStringDocumentAttributeKey
	// CompanyDocumentAttribute is the company or organization name associated with the document.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSCompanyDocumentAttribute
	CompanyDocumentAttribute NSAttributedStringDocumentAttributeKey
	// ConvertedDocumentAttribute is a value that indicates whether a filter service converted the file.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSConvertedDocumentAttribute
	ConvertedDocumentAttribute NSAttributedStringDocumentAttributeKey
	// CopyrightDocumentAttribute is the document’s copyright information.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSCopyrightDocumentAttribute
	CopyrightDocumentAttribute NSAttributedStringDocumentAttributeKey
	// CreationTimeDocumentAttribute is the creation date of the document.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSCreationTimeDocumentAttribute
	CreationTimeDocumentAttribute NSAttributedStringDocumentAttributeKey
	// DefaultAttributesDocumentAttribute is the default document attributes.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSDefaultAttributesDocumentAttribute
	DefaultAttributesDocumentAttribute NSAttributedStringDocumentAttributeKey
	// See: https://developer.apple.com/documentation/AppKit/NSDefaultFontExcludedDocumentAttribute
	DefaultFontExcludedDocumentAttribute NSAttributedStringDocumentAttributeKey
	// DefaultTabIntervalDocumentAttribute is the default tab stop interval for the document.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSDefaultTabIntervalDocumentAttribute
	DefaultTabIntervalDocumentAttribute NSAttributedStringDocumentAttributeKey
	// DocumentTypeDocumentAttribute is the document type.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSDocumentTypeDocumentAttribute
	DocumentTypeDocumentAttribute NSAttributedStringDocumentAttributeKey
	// EditorDocumentAttribute is the name of person who last edited the document.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSEditorDocumentAttribute
	EditorDocumentAttribute NSAttributedStringDocumentAttributeKey
	// ExcludedElementsDocumentAttribute is the HTML elements to exclude in generated HTML.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSExcludedElementsDocumentAttribute
	ExcludedElementsDocumentAttribute NSAttributedStringDocumentAttributeKey
	// FileTypeDocumentAttribute is the document type for interpreting the document.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSFileTypeDocumentAttribute
	FileTypeDocumentAttribute NSAttributedStringDocumentAttributeKey
	// HyphenationFactorDocumentAttribute is the hyphenation factor of the document.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSHyphenationFactorDocumentAttribute
	HyphenationFactorDocumentAttribute NSAttributedStringDocumentAttributeKey
	// KeywordsDocumentAttribute is the document keywords.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSKeywordsDocumentAttribute
	KeywordsDocumentAttribute NSAttributedStringDocumentAttributeKey
	// LeftMarginDocumentAttribute is the left margin of the document.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSLeftMarginDocumentAttribute
	LeftMarginDocumentAttribute NSAttributedStringDocumentAttributeKey
	// ManagerDocumentAttribute is the name of the author’s manager.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSManagerDocumentAttribute
	ManagerDocumentAttribute NSAttributedStringDocumentAttributeKey
	// ModificationTimeDocumentAttribute is the modification date of the document.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSModificationTimeDocumentAttribute
	ModificationTimeDocumentAttribute NSAttributedStringDocumentAttributeKey
	// PaperSizeDocumentAttribute is the paper size for the document.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPaperSizeDocumentAttribute
	PaperSizeDocumentAttribute NSAttributedStringDocumentAttributeKey
	// PrefixSpacesDocumentAttribute is the number of spaces for indenting nested HTML elements.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPrefixSpacesDocumentAttribute
	PrefixSpacesDocumentAttribute NSAttributedStringDocumentAttributeKey
	// ReadOnlyDocumentAttribute is an indication of whether the document is read-only.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSReadOnlyDocumentAttribute
	ReadOnlyDocumentAttribute NSAttributedStringDocumentAttributeKey
	// RightMarginDocumentAttribute is the right margin of the document.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSRightMarginDocumentAttribute
	RightMarginDocumentAttribute NSAttributedStringDocumentAttributeKey
	// SourceTextScalingDocumentAttribute is the text-scaling mode you used when creating the text.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSourceTextScalingDocumentAttribute
	SourceTextScalingDocumentAttribute NSAttributedStringDocumentAttributeKey
	// SubjectDocumentAttribute is the subject of the document.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSubjectDocumentAttribute
	SubjectDocumentAttribute NSAttributedStringDocumentAttributeKey
	// TextEncodingNameDocumentAttribute is the name of the text encoding to use.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextEncodingNameDocumentAttribute
	TextEncodingNameDocumentAttribute NSAttributedStringDocumentAttributeKey
	// TextLayoutSectionsAttribute is the layout orientations for each section.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutSectionsAttribute
	TextLayoutSectionsAttribute NSAttributedStringDocumentAttributeKey
	// TextScalingDocumentAttribute is the text-scaling mode to use when displaying the text.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextScalingDocumentAttribute
	TextScalingDocumentAttribute NSAttributedStringDocumentAttributeKey
	// TitleDocumentAttribute is the document title.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTitleDocumentAttribute
	TitleDocumentAttribute NSAttributedStringDocumentAttributeKey
	// TopMarginDocumentAttribute is the top margin of the document.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTopMarginDocumentAttribute
	TopMarginDocumentAttribute NSAttributedStringDocumentAttributeKey
	// ViewModeDocumentAttribute is the view mode.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSViewModeDocumentAttribute
	ViewModeDocumentAttribute NSAttributedStringDocumentAttributeKey
	// ViewSizeDocumentAttribute is the view size.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSViewSizeDocumentAttribute
	ViewSizeDocumentAttribute NSAttributedStringDocumentAttributeKey
	// ViewZoomDocumentAttribute is the view zoom.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSViewZoomDocumentAttribute
	ViewZoomDocumentAttribute NSAttributedStringDocumentAttributeKey
)

var ()

var (
	// BaseURLDocumentOption is the base URL for HTML documents.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBaseURLDocumentOption
	BaseURLDocumentOption NSAttributedStringDocumentReadingOptionKey
	// CharacterEncodingDocumentOption is the string encoding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSCharacterEncodingDocumentOption
	CharacterEncodingDocumentOption NSAttributedStringDocumentReadingOptionKey
	// DefaultAttributesDocumentOption is the default attributes to apply to plain files.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSDefaultAttributesDocumentOption
	DefaultAttributesDocumentOption NSAttributedStringDocumentReadingOptionKey
	// DocumentTypeDocumentOption is the document type.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSDocumentTypeDocumentOption
	DocumentTypeDocumentOption NSAttributedStringDocumentReadingOptionKey
	// FileTypeDocumentOption is the file type.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSFileTypeDocumentOption
	FileTypeDocumentOption NSAttributedStringDocumentReadingOptionKey
	// SourceTextScalingDocumentOption is the text-scaling mode to associate with the document’s content.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSourceTextScalingDocumentOption
	SourceTextScalingDocumentOption NSAttributedStringDocumentReadingOptionKey
	// TargetTextScalingDocumentOption is the text scaling mode to use after reading the text from disk.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTargetTextScalingDocumentOption
	TargetTextScalingDocumentOption NSAttributedStringDocumentReadingOptionKey
	// TextEncodingNameDocumentOption is the text encoding to use.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextEncodingNameDocumentOption
	TextEncodingNameDocumentOption NSAttributedStringDocumentReadingOptionKey
	// See: https://developer.apple.com/documentation/AppKit/NSTextKit1ListMarkerFormatDocumentOption
	TextKit1ListMarkerFormatDocumentOption NSAttributedStringDocumentReadingOptionKey
	// TextSizeMultiplierDocumentOption is the scale factor for font sizes.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextSizeMultiplierDocumentOption
	TextSizeMultiplierDocumentOption NSAttributedStringDocumentReadingOptionKey
	// TimeoutDocumentOption is the time, in seconds, to wait for a document to finish loading.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTimeoutDocumentOption
	TimeoutDocumentOption NSAttributedStringDocumentReadingOptionKey
	// WebPreferencesDocumentOption is a WebPreferences object.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWebPreferencesDocumentOption
	WebPreferencesDocumentOption NSAttributedStringDocumentReadingOptionKey
	// WebResourceLoadDelegateDocumentOption is an object to serve as the web resource loading delegate.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWebResourceLoadDelegateDocumentOption
	WebResourceLoadDelegateDocumentOption NSAttributedStringDocumentReadingOptionKey
)

var (
	// Black is a constant that specifies the black shade in the 2-bit deep grayscale color space.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBlack
	Black float64
	// DarkGray is a constant that specifies the dark gray shade in the 2-bit deep grayscale color space.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSDarkGray
	DarkGray float64
	// FontIdentityMatrix is the identify matrix for the font.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSFont/identityMatrix
	FontIdentityMatrix float64
	// GridViewSizeForContent is the default value for row and column sizes.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSGridView/sizedForContent
	GridViewSizeForContent float64
	// LightGray is a constant that specifies the light gray shade in the 2-bit deep grayscale color space.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSLightGray
	LightGray float64
	// SplitViewControllerAutomaticDimension is the default value to apply to a dimension.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSplitViewController/automaticDimension
	SplitViewControllerAutomaticDimension float64
	// SplitViewItemUnspecifiedDimension is a constant that resets a dimension’s value.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/unspecifiedDimension
	SplitViewItemUnspecifiedDimension float64
	// ViewNoIntrinsicMetric is a value that tells the layout system to ignore the intrinsic size value for a given dimension.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSView/noIntrinsicMetric
	ViewNoIntrinsicMetric float64
	// White is a constant that specifies the white shade in the 2-bit deep grayscale color space.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWhite
	White float64
)

var (
	// CalibratedRGBColorSpace is calibrated color space with red, green, blue, and alpha components.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/calibratedRGB
	CalibratedRGBColorSpace NSColorSpaceName
	// CalibratedWhiteColorSpace is calibrated color space with white and alpha components (pure white is 1.0).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/calibratedWhite
	CalibratedWhiteColorSpace NSColorSpaceName
	// CustomColorSpace is custom [NSColorSpace] object and floating-point components describing a color in that space.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/custom
	CustomColorSpace NSColorSpaceName
	// DeviceCMYKColorSpace is device-dependent color space with cyan, magenta, yellow, black, and alpha components.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/deviceCMYK
	DeviceCMYKColorSpace NSColorSpaceName
	// DeviceRGBColorSpace is device-dependent color space with red, green, blue, and alpha components.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/deviceRGB
	DeviceRGBColorSpace NSColorSpaceName
	// DeviceWhiteColorSpace is device-dependent color space with white and alpha components (pure white is 1.0).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/deviceWhite
	DeviceWhiteColorSpace NSColorSpaceName
	// NamedColorSpace is catalog name and color name components.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/named
	NamedColorSpace NSColorSpaceName
	// PatternColorSpace is pattern image (tiled).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/pattern
	PatternColorSpace NSColorSpaceName
)

var (
	// CollectionElementKindInterItemGapIndicator is the element kind string assigned to the attributes object when it represents an inter-item gap.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/elementKindInterItemGapIndicator
	CollectionElementKindInterItemGapIndicator NSCollectionViewSupplementaryElementKind
	// CollectionElementKindSectionFooter is a supplementary view that acts as a footer for a given section.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/elementKindSectionFooter
	CollectionElementKindSectionFooter NSCollectionViewSupplementaryElementKind
	// CollectionElementKindSectionHeader is a supplementary view that acts as a header for a given section.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/elementKindSectionHeader
	CollectionElementKindSectionHeader NSCollectionViewSupplementaryElementKind
)

var ()

var (
	// DefinitionPresentationTypeKey is an optional key in the options dictionary that specifies the presentation type of the definition display.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSView/DefinitionOptionKey/presentationType
	DefinitionPresentationTypeKey NSDefinitionOptionKey
)

var (
	// DeviceBitsPerSample is the corresponding value is an [NSNumber] object containing an integer that gives the bit depth of the window’s raster image (2-bit, 8-bit, and so forth).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSDeviceDescriptionKey/bitsPerSample
	DeviceBitsPerSample NSDeviceDescriptionKey
	// DeviceColorSpaceName is the corresponding value is an [NSString] object giving the name of the window’s color space.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSDeviceDescriptionKey/colorSpaceName
	DeviceColorSpaceName NSDeviceDescriptionKey
	// DeviceIsPrinter is if there is a corresponding value, this indicates that the display device is a printer.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSDeviceDescriptionKey/isPrinter
	DeviceIsPrinter NSDeviceDescriptionKey
	// DeviceIsScreen is if there is a corresponding value, this indicates that the display device is a screen.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSDeviceDescriptionKey/isScreen
	DeviceIsScreen NSDeviceDescriptionKey
	// DeviceResolution is the corresponding value is an [NSValue] object containing a value of type [NSSize] that describes the window’s raster resolution in dots per inch (dpi).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSDeviceDescriptionKey/resolution
	DeviceResolution NSDeviceDescriptionKey
	// DeviceSize is the corresponding value is an [NSValue] object containing a value of type [NSSize] that gives the size of the window’s frame rectangle.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSDeviceDescriptionKey/size
	DeviceSize NSDeviceDescriptionKey
)

var (
	// DirectionalEdgeInsetsZero is a directional edge insets structure whose top, leading, bottom, and trailing fields all have a value of `0`.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSDirectionalEdgeInsetsZero
	DirectionalEdgeInsetsZero NSDirectionalEdgeInsets
)

var (
	// DocFormatTextDocumentType is microsoft Word document.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSDocFormatTextDocumentType
	DocFormatTextDocumentType NSAttributedStringDocumentType
	// HTMLTextDocumentType is hypertext markup language (HTML) document.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSHTMLTextDocumentType
	HTMLTextDocumentType NSAttributedStringDocumentType
	// MacSimpleTextDocumentType is macintosh SimpleText document.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSMacSimpleTextDocumentType
	MacSimpleTextDocumentType NSAttributedStringDocumentType
	// OfficeOpenXMLTextDocumentType is eCMA Office Open XML text document format.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSOfficeOpenXMLTextDocumentType
	OfficeOpenXMLTextDocumentType NSAttributedStringDocumentType
	// OpenDocumentTextDocumentType is oASIS Open Document text document format.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSOpenDocumentTextDocumentType
	OpenDocumentTextDocumentType NSAttributedStringDocumentType
	// PlainTextDocumentType is plain text document.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPlainTextDocumentType
	PlainTextDocumentType NSAttributedStringDocumentType
	// RTFDTextDocumentType is rich text format with attachments document.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSRTFDTextDocumentType
	RTFDTextDocumentType NSAttributedStringDocumentType
	// RTFTextDocumentType is rich text format document.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSRTFTextDocumentType
	RTFTextDocumentType NSAttributedStringDocumentType
	// WebArchiveTextDocumentType is webKit WebArchive document.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWebArchiveTextDocumentType
	WebArchiveTextDocumentType NSAttributedStringDocumentType
	// WordMLTextDocumentType is microsoft Word XML (WordML schema) document.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWordMLTextDocumentType
	WordMLTextDocumentType NSAttributedStringDocumentType
)

var (
	// DraggingImageComponentIconKey is a key for a corresponding value that is a dragging item’s image.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSDraggingItem/ImageComponentKey/icon
	DraggingImageComponentIconKey NSDraggingImageComponentKey
	// DraggingImageComponentLabelKey is a key for a corresponding value that represents a textual label for a dragging item, for example, a file name.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSDraggingItem/ImageComponentKey/label
	DraggingImageComponentLabelKey NSDraggingImageComponentKey
)

var (
	// FileContentsPboardType is a representation of a file’s contents.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPasteboard/PasteboardType/fileContents
	FileContentsPboardType NSPasteboardType
	// FindPanelSearchOptionsPboardType is type for the find panel metadata property list.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPasteboard/PasteboardType/findPanelSearchOptions
	FindPanelSearchOptionsPboardType NSPasteboardType
	// SoundPboardType is [NSSound] data.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSoundPboardType
	SoundPboardType NSPasteboardType
)

var (
	// FindPanelCaseInsensitiveSearch is a Boolean value indicating whether the search is case-insensitive.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPasteboard/PasteboardType/FindPanelSearchOptionKey/findPanelCaseInsensitiveSearch
	FindPanelCaseInsensitiveSearch NSPasteboardTypeFindPanelSearchOptionKey
	// FindPanelSubstringMatch is a number object containing the match type to use in the find panel.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPasteboard/PasteboardType/FindPanelSearchOptionKey/findPanelSubstringMatch
	FindPanelSubstringMatch NSPasteboardTypeFindPanelSearchOptionKey
)

var (
	// FontCascadeListAttribute is an array, each member of which is a sub-descriptor.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/AttributeName/cascadeList
	FontCascadeListAttribute NSFontDescriptorAttributeName
	// FontCharacterSetAttribute is the set of Unicode characters covered by the font.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/AttributeName/characterSet
	FontCharacterSetAttribute NSFontDescriptorAttributeName
	// FontFaceAttribute is an optional string object that specifies the font face.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/AttributeName/face
	FontFaceAttribute NSFontDescriptorAttributeName
	// FontFamilyAttribute is an optional string object that specifies the font family.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/AttributeName/family
	FontFamilyAttribute NSFontDescriptorAttributeName
	// FontFeatureSettingsAttribute is an array of dictionaries representing non-default font feature settings.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/AttributeName/featureSettings
	FontFeatureSettingsAttribute NSFontDescriptorAttributeName
	// FontFixedAdvanceAttribute is a floating-point value that overrides the glyph advancement specified by the font.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/AttributeName/fixedAdvance
	FontFixedAdvanceAttribute NSFontDescriptorAttributeName
	// FontMatrixAttribute is an affine transform that specifies the font’s transformation matrix.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/AttributeName/matrix
	FontMatrixAttribute NSFontDescriptorAttributeName
	// FontNameAttribute is an optional string object that specifies the font name.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/AttributeName/name
	FontNameAttribute NSFontDescriptorAttributeName
	// FontSizeAttribute is an optional floating-point number that specifies the font size.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/AttributeName/size
	FontSizeAttribute NSFontDescriptorAttributeName
	// FontTraitsAttribute is a dictionary that fully describes the font traits.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/AttributeName/traits
	FontTraitsAttribute NSFontDescriptorAttributeName
	// FontVariationAttribute is a dictionary that describes the font’s variation axis.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/AttributeName/variation
	FontVariationAttribute NSFontDescriptorAttributeName
	// FontVisibleNameAttribute is an optional string object that specifies the font’s visible name.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/AttributeName/visibleName
	FontVisibleNameAttribute NSFontDescriptorAttributeName
)

var (
	// FontCollectionActionKey is an action was taken. See `NSFontCollectionAction Key Values` for the possible values. An [NSString].
	//
	// See: https://developer.apple.com/documentation/AppKit/NSFontCollection/actionUserInfoKey
	FontCollectionActionKey NSFontCollectionUserInfoKey
	// FontCollectionNameKey is the font collection’s name. If renamed, this is the new name. An [NSString].
	//
	// See: https://developer.apple.com/documentation/AppKit/NSFontCollection/nameUserInfoKey
	FontCollectionNameKey NSFontCollectionUserInfoKey
	// FontCollectionOldNameKey is included as a value for the [oldNameUserInfoKey] key, if present. This is the previous name. An [NSString].
	//
	// See: https://developer.apple.com/documentation/AppKit/NSFontCollection/oldNameUserInfoKey
	FontCollectionOldNameKey NSFontCollectionUserInfoKey
	// FontCollectionVisibilityKey is the visibly of the font collection. An NSNumber containing a value from the [NSFontCollection.Visibility] enum.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSFontCollection/visibilityUserInfoKey
	FontCollectionVisibilityKey NSFontCollectionUserInfoKey
)

var (
	// FontCollectionAllFonts is all fonts in the system.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSFontCollection/Name/allFonts
	FontCollectionAllFonts NSFontCollectionName
	// FontCollectionFavorites is font collection of the user’s preferred font descriptors.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSFontCollection/Name/favorites
	FontCollectionFavorites NSFontCollectionName
	// FontCollectionRecentlyUsed is font collection automatically maintained by NSFontManager.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSFontCollection/Name/recentlyUsed
	FontCollectionRecentlyUsed NSFontCollectionName
	// FontCollectionUser is per-user unmodifiable collection.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSFontCollection/Name/user
	FontCollectionUser NSFontCollectionName
)

var (
	// FontCollectionDisallowAutoActivationOption is an NSNumber object containing a Boolean value specifying that auto-activation should not be used to find missing fonts.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSFontCollectionMatchingOptionKey/disallowAutoActivationOption
	FontCollectionDisallowAutoActivationOption NSFontCollectionMatchingOptionKey
	// FontCollectionIncludeDisabledFontsOption is an NSNumber object containing a Boolean value specifying whether disabled fonts should be included in the list of matching descriptors.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSFontCollectionMatchingOptionKey/includeDisabledFontsOption
	FontCollectionIncludeDisabledFontsOption NSFontCollectionMatchingOptionKey
	// FontCollectionRemoveDuplicatesOption is an NSNumber object containing a Boolean value controlling whether more than one copy of a font with the same PostScript name should be included in the list of matching descriptors.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSFontCollectionMatchingOptionKey/removeDuplicatesOption
	FontCollectionRemoveDuplicatesOption NSFontCollectionMatchingOptionKey
)

var (
	// FontCollectionWasHidden is the font collection was hidden.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSFontCollection/ActionTypeKey/hidden
	FontCollectionWasHidden NSFontCollectionActionTypeKey
	// FontCollectionWasRenamed is the font collection was renamed.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSFontCollection/ActionTypeKey/renamed
	FontCollectionWasRenamed NSFontCollectionActionTypeKey
	// FontCollectionWasShown is the font collection was shown.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSFontCollection/ActionTypeKey/shown
	FontCollectionWasShown NSFontCollectionActionTypeKey
)

var ()

var (
	// FontFeatureSelectorIdentifierKey is a key that indicates the selector of the font feature.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/FeatureKey/selectorIdentifier
	FontFeatureSelectorIdentifierKey NSFontDescriptorFeatureKey
	// FontFeatureTypeIdentifierKey is a key that indicates the type of the font feature.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/FeatureKey/typeIdentifier
	FontFeatureTypeIdentifierKey NSFontDescriptorFeatureKey
)

var (
	// FontSlantTrait is the relative slant angle value as a number object.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/TraitKey/slant
	FontSlantTrait NSFontDescriptorTraitKey
	// FontSymbolicTrait is the symbolic traits value as a number object.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/TraitKey/symbolic
	FontSymbolicTrait NSFontDescriptorTraitKey
	// FontWeightTrait is the normalized weight value as a number object.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/TraitKey/weight
	FontWeightTrait NSFontDescriptorTraitKey
	// FontWidthTrait is the relative inter-glyph spacing value as a number object.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/TraitKey/width
	FontWidthTrait NSFontDescriptorTraitKey
)

var ()

var (
	// FontVariationAxisDefaultValueKey is the default axis value as a number object.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/VariationKey/defaultValue
	FontVariationAxisDefaultValueKey NSFontDescriptorVariationKey
	// FontVariationAxisIdentifierKey is the axis identifier value as a number object.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/VariationKey/identifier
	FontVariationAxisIdentifierKey NSFontDescriptorVariationKey
	// FontVariationAxisMaximumValueKey is the maximum axis value as a number object.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/VariationKey/maximumValue
	FontVariationAxisMaximumValueKey NSFontDescriptorVariationKey
	// FontVariationAxisMinimumValueKey is the minimum axis value as a number object.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/VariationKey/minimumValue
	FontVariationAxisMinimumValueKey NSFontDescriptorVariationKey
	// FontVariationAxisNameKey is the localized variation axis name.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/VariationKey/name
	FontVariationAxisNameKey NSFontDescriptorVariationKey
)

var ()

var ()

var (
	// FullScreenModeAllScreens is key whose corresponding value specifies whether the view should take over all screens.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSView/FullScreenModeOptionKey/fullScreenModeAllScreens
	FullScreenModeAllScreens NSViewFullScreenModeOptionKey
	// FullScreenModeApplicationPresentationOptions is key whose corresponding value specifies the application presentation options.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSView/FullScreenModeOptionKey/fullScreenModeApplicationPresentationOptions
	FullScreenModeApplicationPresentationOptions NSViewFullScreenModeOptionKey
	// FullScreenModeSetting is key whose corresponding value specifies the full screen mode setting.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSView/FullScreenModeOptionKey/fullScreenModeSetting
	FullScreenModeSetting NSViewFullScreenModeOptionKey
	// FullScreenModeWindowLevel is key whose corresponding value specifies the screen mode window level.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSView/FullScreenModeOptionKey/fullScreenModeWindowLevel
	FullScreenModeWindowLevel NSViewFullScreenModeOptionKey
)

var (
	// GraphicsContextPDFFormat is destination file format is PDF.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/RepresentationFormatName/pdf
	GraphicsContextPDFFormat NSGraphicsContextRepresentationFormatName
	// GraphicsContextPSFormat is destination file format is PostScript.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/RepresentationFormatName/postScript
	GraphicsContextPSFormat NSGraphicsContextRepresentationFormatName
)

var (
	// ImageColorSyncProfileData is identifies an [NSData] object containing the ColorSync profile data.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/PropertyKey/colorSyncProfileData
	ImageColorSyncProfileData NSBitmapImageRepPropertyKey
	// ImageCompressionFactor is identifies an [NSNumber] object containing the compression factor of the image.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/PropertyKey/compressionFactor
	ImageCompressionFactor NSBitmapImageRepPropertyKey
	// ImageCompressionMethod is identifies an [NSNumber] object identifying the compression method of the image.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/PropertyKey/compressionMethod
	ImageCompressionMethod NSBitmapImageRepPropertyKey
	// ImageCurrentFrame is identifies an [NSNumber] object containing the current frame for an animated GIF file.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/PropertyKey/currentFrame
	ImageCurrentFrame NSBitmapImageRepPropertyKey
	// ImageCurrentFrameDuration is identifies an [NSNumber] object containing the duration (in seconds) of the current frame for an animated GIF image.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/PropertyKey/currentFrameDuration
	ImageCurrentFrameDuration NSBitmapImageRepPropertyKey
	// ImageDitherTransparency is identifies an [NSNumber] object containing a Boolean that indicates whether the image is dithered.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/PropertyKey/ditherTransparency
	ImageDitherTransparency NSBitmapImageRepPropertyKey
	// ImageEXIFData is identifies an [NSDictionary] object containing the EXIF data for the image.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/PropertyKey/exifData
	ImageEXIFData NSBitmapImageRepPropertyKey
	// ImageFallbackBackgroundColor is specifies the background color to use when writing to an image format (such as JPEG) that doesn’t support alpha.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/PropertyKey/fallbackBackgroundColor
	ImageFallbackBackgroundColor NSBitmapImageRepPropertyKey
	// ImageFrameCount is identifies an [NSNumber] object containing the number of frames in an animated GIF file.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/PropertyKey/frameCount
	ImageFrameCount NSBitmapImageRepPropertyKey
	// ImageGamma is identifies an [NSNumber] object containing the gamma value for the image.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/PropertyKey/gamma
	ImageGamma NSBitmapImageRepPropertyKey
	// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/PropertyKey/imageIPTCData
	ImageIPTCData NSBitmapImageRepPropertyKey
	// ImageInterlaced is identifies an [NSNumber] object containing a Boolean value that indicates whether the image is interlaced.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/PropertyKey/interlaced
	ImageInterlaced NSBitmapImageRepPropertyKey
	// ImageLoopCount is identifies an [NSNumber] object containing the number of loops to make when animating a GIF image.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/PropertyKey/loopCount
	ImageLoopCount NSBitmapImageRepPropertyKey
	// ImageProgressive is identifies an [NSNumber] object containing a Boolean that indicates whether the image uses progressive encoding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/PropertyKey/progressive
	ImageProgressive NSBitmapImageRepPropertyKey
	// ImageRGBColorTable is identifies an [NSData] object containing the RGB color table.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/PropertyKey/rgbColorTable
	ImageRGBColorTable NSBitmapImageRepPropertyKey
)

var (
	// ImageHintCTM is a context transform hint.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSImageRep/HintKey/ctm
	ImageHintCTM NSImageHintKey
	// ImageHintInterpolation is an interpolation hint.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSImageRep/HintKey/interpolation
	ImageHintInterpolation NSImageHintKey
	// ImageHintUserInterfaceLayoutDirection is a layout direction hint.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSImageRep/HintKey/userInterfaceLayoutDirection
	ImageHintUserInterfaceLayoutDirection NSImageHintKey
)

var ()

var (
	// MenuItemImportFromDeviceIdentifier is the identifier for a Continuity Camera menu item, which takes pictures or scans documents using an iOS device.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSMenuItem/importFromDeviceIdentifier
	MenuItemImportFromDeviceIdentifier NSUserInterfaceItemIdentifier
	// OutlineViewDisclosureButtonKey is the normal triangle disclosure button.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSOutlineView/disclosureButtonIdentifier
	OutlineViewDisclosureButtonKey NSUserInterfaceItemIdentifier
	// OutlineViewShowHideButtonKey is the Show/Hide button.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSOutlineView/showHideButtonIdentifier
	OutlineViewShowHideButtonKey NSUserInterfaceItemIdentifier
	// TableViewRowViewKey is the key associated with the identifier in the nib file containing the template row view.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTableView/rowViewIdentifier
	TableViewRowViewKey NSUserInterfaceItemIdentifier
)

var (
	// ObservedKeyPathKey is an [NSString] object containing the key path of the binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingInfoKey/observedKeyPath
	ObservedKeyPathKey NSBindingInfoKey
	// ObservedObjectKey is the object that is the observable controller of the binding.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingInfoKey/observedObject
	ObservedObjectKey NSBindingInfoKey
	// OptionsKey is an [NSDictionary] object containing key value pairs as specified in the options dictionary when the binding was created.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSBindingInfoKey/options
	OptionsKey NSBindingInfoKey
)

var ()

var (
	// PasteboardMetadataTypeContentType is a metadata type that returns the content type if the pasteboard detects a reference to a file.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPasteboardMetadataTypeContentType
	PasteboardMetadataTypeContentType NSPasteboardMetadataType
)

var ()

var (
	// PasteboardURLReadingContentsConformToTypesKey is option for reading URLs to restrict the results to URLs with contents that conform to any of the provided UTI types.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPasteboard/ReadingOptionKey/urlReadingContentsConformToTypes
	PasteboardURLReadingContentsConformToTypesKey NSPasteboardReadingOptionKey
	// PasteboardURLReadingFileURLsOnlyKey is option for reading URLs to restrict the results to file URLs only.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPasteboard/ReadingOptionKey/urlReadingFileURLsOnly
	PasteboardURLReadingFileURLsOnlyKey NSPasteboardReadingOptionKey
)

var (
	// PopoverCloseReasonDetachToWindow is specifies that the popover has been closed because it is being detached to a window.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPopover/CloseReason/detachToWindow
	PopoverCloseReasonDetachToWindow NSPopoverCloseReasonValue
	// PopoverCloseReasonStandard is specifies that the popover has been closed in a standard way.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPopover/CloseReason/standard
	PopoverCloseReasonStandard NSPopoverCloseReasonValue
)

var (
	// PrintAllPages is an [NSNumber] object containing a Boolean value that specifies whether to include all pages.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/AttributeKey/allPages
	PrintAllPages NSPrintInfoAttributeKey
	// PrintBottomMargin is [NSNumber], containing a floating-point value that specifies the bottom margin, in points.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/AttributeKey/bottomMargin
	PrintBottomMargin NSPrintInfoAttributeKey
	// PrintCopies is an [NSNumber] object containing an integer—the number of copies to spool.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/AttributeKey/copies
	PrintCopies NSPrintInfoAttributeKey
	// PrintDetailedErrorReporting is an [NSNumber] object containing a Boolean value that specifies whether to produce detailed error reports.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/AttributeKey/detailedErrorReporting
	PrintDetailedErrorReporting NSPrintInfoAttributeKey
	// PrintFaxNumber is an [NSString] object that specifies a fax number.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/AttributeKey/faxNumber
	PrintFaxNumber NSPrintInfoAttributeKey
	// PrintFirstPage is an [NSNumber] object containing an integer value that specifies the first page in the print job.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/AttributeKey/firstPage
	PrintFirstPage NSPrintInfoAttributeKey
	// PrintHeaderAndFooter is an [NSNumber] object containing a Boolean value that specifies whether to include a header and footer.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/AttributeKey/headerAndFooter
	PrintHeaderAndFooter NSPrintInfoAttributeKey
	// PrintHorizontalPagination is [NSNumber], containing a [NSPrintingPaginationMode] value.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/AttributeKey/horizontalPagination
	PrintHorizontalPagination NSPrintInfoAttributeKey
	// PrintHorizontallyCentered is an [NSNumber] object containing a Boolean value that specifies whether to horizontally center pages.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/AttributeKey/horizontallyCentered
	PrintHorizontallyCentered NSPrintInfoAttributeKey
	// PrintJobDisposition is an [NSString] object that specifies the job disposition.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/AttributeKey/jobDisposition
	PrintJobDisposition NSPrintInfoAttributeKey
	// PrintJobSavingFileNameExtensionHidden is an [NSNumber] object containing a Boolean value that specifies whether to hide the job’s file name extension.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/AttributeKey/jobSavingFileNameExtensionHidden
	PrintJobSavingFileNameExtensionHidden NSPrintInfoAttributeKey
	// PrintJobSavingURL is an [NSURL] containing the location to which the job file will be saved when the [jobDisposition] is [save].
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/AttributeKey/jobSavingURL
	PrintJobSavingURL NSPrintInfoAttributeKey
	// PrintLastPage is an [NSNumber] object containing an integer value that specifies the last page in the print job.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/AttributeKey/lastPage
	PrintLastPage NSPrintInfoAttributeKey
	// PrintLeftMargin is [NSNumber], containing a floating-point value that specifies the left margin, in points.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/AttributeKey/leftMargin
	PrintLeftMargin NSPrintInfoAttributeKey
	// PrintMustCollate is an [NSNumber] object containing a Boolean value that specifies whether to collate output.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/AttributeKey/mustCollate
	PrintMustCollate NSPrintInfoAttributeKey
	// PrintOrientation is an [NSNumber] object containing an [NSPrintingOrientation].
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/AttributeKey/orientation
	PrintOrientation NSPrintInfoAttributeKey
	// PrintPagesAcross is an [NSNumber] object that specifies the number of logical pages to be tiled horizontally on a physical sheet of paper.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/AttributeKey/pagesAcross
	PrintPagesAcross NSPrintInfoAttributeKey
	// PrintPagesDown is an [NSNumber] object that specifies the number of logical pages to be tiled vertically on a physical sheet of paper.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/AttributeKey/pagesDown
	PrintPagesDown NSPrintInfoAttributeKey
	// PrintPaperName is an [NSString] object containing the paper name.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/AttributeKey/paperName
	PrintPaperName NSPrintInfoAttributeKey
	// PrintPaperSize is an [NSSize] value specifying the height and width of paper in points.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/AttributeKey/paperSize
	PrintPaperSize NSPrintInfoAttributeKey
	// PrintPrinter is an [NSPrinter] object—the printer to use.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/AttributeKey/printer
	PrintPrinter NSPrintInfoAttributeKey
	// PrintPrinterName is an [NSString] object that specifies the name of a printer.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/AttributeKey/printerName
	PrintPrinterName NSPrintInfoAttributeKey
	// PrintReversePageOrder is an [NSNumber] object containing a Boolean value that specifies whether to print pages in reverse order.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/AttributeKey/reversePageOrder
	PrintReversePageOrder NSPrintInfoAttributeKey
	// PrintRightMargin is [NSNumber], containing a floating-point value that specifies the right margin, in points.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/AttributeKey/rightMargin
	PrintRightMargin NSPrintInfoAttributeKey
	// PrintScalingFactor is scale factor percentage before pagination.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/AttributeKey/scalingFactor
	PrintScalingFactor NSPrintInfoAttributeKey
	// PrintSelectionOnly is an [NSNumber] object containing a Boolean value that specifies whether to print the current selection.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/AttributeKey/selectionOnly
	PrintSelectionOnly NSPrintInfoAttributeKey
	// PrintTime is an [NSDate] object that specifies the time at which printing should begin.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/AttributeKey/time
	PrintTime NSPrintInfoAttributeKey
	// PrintTopMargin is [NSNumber], containing a floating-point value that specifies the top margin, in points.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/AttributeKey/topMargin
	PrintTopMargin NSPrintInfoAttributeKey
	// PrintVerticalPagination is [NSNumber], containing a [NSPrintingPaginationMode] value.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/AttributeKey/verticalPagination
	PrintVerticalPagination NSPrintInfoAttributeKey
	// PrintVerticallyCentered is an [NSNumber] object containing a Boolean value that specifies whether to vertically center pages.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/AttributeKey/verticallyCentered
	PrintVerticallyCentered NSPrintInfoAttributeKey
)

var (
	// PrintAllPresetsJobStyleHint is output appropriate to all graphics types.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPrintPanel/JobStyleHint-swift.struct/allPresets
	PrintAllPresetsJobStyleHint NSPrintPanelJobStyleHint
	// PrintNoPresetsJobStyleHint is output excludes all graphics printing.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPrintPanel/JobStyleHint-swift.struct/noPresets
	PrintNoPresetsJobStyleHint NSPrintPanelJobStyleHint
	// PrintPhotoJobStyleHint is output contains photographic data.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPrintPanel/JobStyleHint-swift.struct/photo
	PrintPhotoJobStyleHint NSPrintPanelJobStyleHint
)

var (
	// PrintCancelJob is cancel print job.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/JobDisposition-swift.struct/cancel
	PrintCancelJob NSPrintJobDispositionValue
	// PrintPreviewJob is send to Preview application.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/JobDisposition-swift.struct/preview
	PrintPreviewJob NSPrintJobDispositionValue
	// PrintSaveJob is save to a file.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/JobDisposition-swift.struct/save
	PrintSaveJob NSPrintJobDispositionValue
	// PrintSpoolJob is normal print job.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/JobDisposition-swift.struct/spool
	PrintSpoolJob NSPrintJobDispositionValue
)

var (
	// PrintPanelAccessorySummaryItemDescriptionKey is a key that identfies the current value of the accessory panel setting.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPrintPanel/AccessorySummaryKey/itemDescription
	PrintPanelAccessorySummaryItemDescriptionKey NSPrintPanelAccessorySummaryKey
	// PrintPanelAccessorySummaryItemNameKey is a key that specifies the name of the accessory panel setting.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPrintPanel/AccessorySummaryKey/itemName
	PrintPanelAccessorySummaryItemNameKey NSPrintPanelAccessorySummaryKey
)

var (
	// RuleEditorPredicateComparisonModifier is the corresponding value is an [NSNumber] object representing a [NSComparisonPredicateModifier] constant the of the predicate.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSRuleEditor/PredicatePartKey/comparisonModifier
	RuleEditorPredicateComparisonModifier NSRuleEditorPredicatePartKey
	// RuleEditorPredicateCompoundType is the corresponding value is an [NSNumber] object representing a [NSCompoundPredicateType] constant.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSRuleEditor/PredicatePartKey/compoundType
	RuleEditorPredicateCompoundType NSRuleEditorPredicatePartKey
	// RuleEditorPredicateCustomSelector is the corresponding value is an [NSString] object representing a custom selector.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSRuleEditor/PredicatePartKey/customSelector
	RuleEditorPredicateCustomSelector NSRuleEditorPredicatePartKey
	// RuleEditorPredicateLeftExpression is the corresponding value is an [NSExpression] object representing the left expression in the predicate.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSRuleEditor/PredicatePartKey/leftExpression
	RuleEditorPredicateLeftExpression NSRuleEditorPredicatePartKey
	// RuleEditorPredicateOperatorType is the corresponding value is an [NSNumber] object representing a [NSPredicateOperatorType] constant.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSRuleEditor/PredicatePartKey/operatorType
	RuleEditorPredicateOperatorType NSRuleEditorPredicatePartKey
	// RuleEditorPredicateOptions is the corresponding value is an [NSNumber] object representing an [NSComparisonPredicateOptions] bitfield.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSRuleEditor/PredicatePartKey/options
	RuleEditorPredicateOptions NSRuleEditorPredicatePartKey
	// RuleEditorPredicateRightExpression is the corresponding value is an [NSExpression] object representing the right expression in the predicate.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSRuleEditor/PredicatePartKey/rightExpression
	RuleEditorPredicateRightExpression NSRuleEditorPredicatePartKey
)

var (
	// See: https://developer.apple.com/documentation/AppKit/NSRulerView/UnitName/centimeters
	RulerViewUnitCentimeters NSRulerViewUnitName
	// See: https://developer.apple.com/documentation/AppKit/NSRulerView/UnitName/inches
	RulerViewUnitInches NSRulerViewUnitName
	// See: https://developer.apple.com/documentation/AppKit/NSRulerView/UnitName/picas
	RulerViewUnitPicas NSRulerViewUnitName
	// See: https://developer.apple.com/documentation/AppKit/NSRulerView/UnitName/points
	RulerViewUnitPoints NSRulerViewUnitName
)

var ()

var ()

var (
	// SpeechCharacterModeProperty is get or set the synthesizer’s current text-processing mode.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/SpeechPropertyKey/characterMode
	SpeechCharacterModeProperty NSSpeechPropertyKey
	// SpeechCommandDelimiterProperty is set the embedded speech command delimiter characters to be used for the synthesizer.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/SpeechPropertyKey/commandDelimiter
	SpeechCommandDelimiterProperty NSSpeechPropertyKey
	// SpeechCurrentVoiceProperty is set the current voice on the synthesizer to the specified voice.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/SpeechPropertyKey/currentVoice
	SpeechCurrentVoiceProperty NSSpeechPropertyKey
	// SpeechErrorsProperty is get speech-error information for the synthesizer.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/SpeechPropertyKey/errors
	SpeechErrorsProperty NSSpeechPropertyKey
	// SpeechInputModeProperty is get or set the synthesizer’s current text-processing mode.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/SpeechPropertyKey/inputMode
	SpeechInputModeProperty NSSpeechPropertyKey
	// SpeechNumberModeProperty is get or set the synthesizer’s current number-processing mode.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/SpeechPropertyKey/numberMode
	SpeechNumberModeProperty NSSpeechPropertyKey
	// SpeechOutputToFileURLProperty is set the speech output destination to a file or to the computer’s speakers.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/SpeechPropertyKey/outputToFileURL
	SpeechOutputToFileURLProperty NSSpeechPropertyKey
	// SpeechPhonemeSymbolsProperty is get a list of phoneme symbols and example words defined for the synthesizer.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/SpeechPropertyKey/phonemeSymbols
	SpeechPhonemeSymbolsProperty NSSpeechPropertyKey
	// SpeechPitchBaseProperty is get or set a synthesizer’s baseline speech pitch.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/SpeechPropertyKey/pitchBase
	SpeechPitchBaseProperty NSSpeechPropertyKey
	// SpeechPitchModProperty is get or set a synthesizer’s pitch modulation.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/SpeechPropertyKey/pitchMod
	SpeechPitchModProperty NSSpeechPropertyKey
	// SpeechRateProperty is get or set a synthesizer’s speech rate.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/SpeechPropertyKey/rate
	SpeechRateProperty NSSpeechPropertyKey
	// SpeechRecentSyncProperty is get the message code for the most recently encountered synchronization command.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/SpeechPropertyKey/recentSync
	SpeechRecentSyncProperty NSSpeechPropertyKey
	// SpeechResetProperty is set a synthesizer back to its default state.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/SpeechPropertyKey/reset
	SpeechResetProperty NSSpeechPropertyKey
	// SpeechStatusProperty is get speech-status information for the synthesizer.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/SpeechPropertyKey/status
	SpeechStatusProperty NSSpeechPropertyKey
	// SpeechSynthesizerInfoProperty is get information about the speech synthesizer being used on the specified synthesizer.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/SpeechPropertyKey/synthesizerInfo
	SpeechSynthesizerInfoProperty NSSpeechPropertyKey
	// SpeechVolumeProperty is get or set the speech volume for a synthesizer.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/SpeechPropertyKey/volume
	SpeechVolumeProperty NSSpeechPropertyKey
)

var (
	// SpeechCommandPrefix is the command delimiter string that prefixes a command, by default, this is `[[`.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/SpeechPropertyKey/CommandDelimiterKey/prefix
	SpeechCommandPrefix NSSpeechCommandDelimiterKey
	// SpeechCommandSuffix is the command delimiter string that suffixes a command,by default, this is `]]`.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/SpeechPropertyKey/CommandDelimiterKey/suffix
	SpeechCommandSuffix NSSpeechCommandDelimiterKey
)

var (
	// SpeechDictionaryAbbreviations is an array of dictionary objects containing the keys [NSSpeechDictionaryEntrySpelling] and [NSSpeechDictionaryEntryPhonemes].
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/DictionaryKey/abbreviations
	SpeechDictionaryAbbreviations NSSpeechDictionaryKey
	// SpeechDictionaryEntryPhonemes is the phonemic representation of an entry. An [NSString].
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/DictionaryKey/entryPhonemes
	SpeechDictionaryEntryPhonemes NSSpeechDictionaryKey
	// SpeechDictionaryEntrySpelling is the spelling of an entry. An [NSString].
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/DictionaryKey/entrySpelling
	SpeechDictionaryEntrySpelling NSSpeechDictionaryKey
	// SpeechDictionaryLocaleIdentifier is the canonical locale identifier string describing the dictionary’s locale.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/DictionaryKey/localeIdentifier
	SpeechDictionaryLocaleIdentifier NSSpeechDictionaryKey
	// SpeechDictionaryModificationDate is a string representation of the dictionary’s last modification date in the international format (YYYY-MM-DD HH:MM:SS ±HHMM). If the same word appears across multiple dictionaries, the one from the dictionary with the most recent date will be used.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/DictionaryKey/modificationDate
	SpeechDictionaryModificationDate NSSpeechDictionaryKey
	// SpeechDictionaryPronunciations is an array of dictionary objects containing the keys [NSSpeechDictionaryEntrySpelling] and [NSSpeechDictionaryEntryPhonemes].
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/DictionaryKey/pronunciations
	SpeechDictionaryPronunciations NSSpeechDictionaryKey
)

var (
	// SpeechErrorCount is the number of errors that have occurred in processing the current text string, since the last call to [object(forProperty:)] with the [errors] property. An [NSNumber].
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/SpeechPropertyKey/ErrorKey/count
	SpeechErrorCount NSSpeechErrorKey
	// SpeechErrorNewestCharacterOffset is the position in the text string of the most recent error that occurred since the last call to [object(forProperty:)] with the [errors] property. An [NSNumber].
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/SpeechPropertyKey/ErrorKey/newestCharacterOffset
	SpeechErrorNewestCharacterOffset NSSpeechErrorKey
	// SpeechErrorNewestCode is the error code of the most recent error that occurred since the last call to [object(forProperty:)] with the [errors] property. An [NSNumber].
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/SpeechPropertyKey/ErrorKey/newestCode
	SpeechErrorNewestCode NSSpeechErrorKey
	// SpeechErrorOldestCharacterOffset is the position in the text string of the first error that occurred since the last call to [object(forProperty:)] with the [errors] property. An [NSNumber].
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/SpeechPropertyKey/ErrorKey/oldestCharacterOffset
	SpeechErrorOldestCharacterOffset NSSpeechErrorKey
	// SpeechErrorOldestCode is the error code of the first error that occurred since the last call to [object(forProperty:)] with the [errors] property. An [NSNumber].
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/SpeechPropertyKey/ErrorKey/oldestCode
	SpeechErrorOldestCode NSSpeechErrorKey
)

var ()

var (
	// SpeechPhonemeInfoExample is an example word that illustrates the use of the phoneme.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/SpeechPropertyKey/PhonemeInfoKey/example
	SpeechPhonemeInfoExample NSSpeechPhonemeInfoKey
	// SpeechPhonemeInfoHiliteEnd is the character offset into the example word that identifies the location of the end of the phoneme.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/SpeechPropertyKey/PhonemeInfoKey/hiliteEnd
	SpeechPhonemeInfoHiliteEnd NSSpeechPhonemeInfoKey
	// SpeechPhonemeInfoHiliteStart is the character offset into the example word that identifies the location of the beginning of the phoneme.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/SpeechPropertyKey/PhonemeInfoKey/hiliteStart
	SpeechPhonemeInfoHiliteStart NSSpeechPhonemeInfoKey
	// SpeechPhonemeInfoOpcode is nSNumber.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/SpeechPropertyKey/PhonemeInfoKey/opcode
	SpeechPhonemeInfoOpcode NSSpeechPhonemeInfoKey
	// SpeechPhonemeInfoSymbol is the symbol used to represent the phoneme.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/SpeechPropertyKey/PhonemeInfoKey/symbol
	SpeechPhonemeInfoSymbol NSSpeechPhonemeInfoKey
)

var (
	// SpeechStatusNumberOfCharactersLeft is the number of characters left in the input string of text.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/SpeechPropertyKey/StatusKey/numberOfCharactersLeft
	SpeechStatusNumberOfCharactersLeft NSSpeechStatusKey
	// SpeechStatusOutputBusy is indicates whether the synthesizer is currently producing speech.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/SpeechPropertyKey/StatusKey/outputBusy
	SpeechStatusOutputBusy NSSpeechStatusKey
	// SpeechStatusOutputPaused is indicates whether speech output in the synthesizer has been paused by sending the message [pauseSpeaking(at:)].
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/SpeechPropertyKey/StatusKey/outputPaused
	SpeechStatusOutputPaused NSSpeechStatusKey
	// SpeechStatusPhonemeCode is indicates that the synthesizer is in phoneme-processing mode. When in phoneme-processing mode, a text buffer is interpreted to be a series of characters representing various phonemes and prosodic controls.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/SpeechPropertyKey/StatusKey/phonemeCode
	SpeechStatusPhonemeCode NSSpeechStatusKey
)

var (
	// SpeechSynthesizerInfoIdentifier is the identifier of the speech synthesizer.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/SpeechPropertyKey/SynthesizerInfoKey/identifier
	SpeechSynthesizerInfoIdentifier NSSpeechSynthesizerInfoKey
	// SpeechSynthesizerInfoVersion is the version of the speech synthesizer.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/SpeechPropertyKey/SynthesizerInfoKey/version
	SpeechSynthesizerInfoVersion NSSpeechSynthesizerInfoKey
)

var (
	// TabColumnTerminatorsAttributeName is the value is an [NSCharacterSet] object.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextTab/OptionKey/columnTerminators
	TabColumnTerminatorsAttributeName NSTextTabOptionKey
)

var (
	// TextCheckingDocumentAuthorKey is an NSString containing the name of an author to be associated with the document.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/OptionKey/documentAuthor
	TextCheckingDocumentAuthorKey NSTextCheckingOptionKey
	// TextCheckingDocumentTitleKey is an NSString containing the title to be associated with the document.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/OptionKey/documentTitle
	TextCheckingDocumentTitleKey NSTextCheckingOptionKey
	// TextCheckingDocumentURLKey is an NSURL to be associated with the document.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/OptionKey/documentURL
	TextCheckingDocumentURLKey NSTextCheckingOptionKey
	// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/OptionKey/generateInlinePredictionsKey
	TextCheckingGenerateInlinePredictionsKey NSTextCheckingOptionKey
	// TextCheckingOrthographyKey is an [NSOrthography] instance indicating an orthography to be used as a starting point for orthography checking, or as the orthography if orthography checking is not enabled.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/OptionKey/orthography
	TextCheckingOrthographyKey NSTextCheckingOptionKey
	// TextCheckingQuotesKey is an [NSArray] containing four strings to be used with `quoteCheckingResult()` (opening double quote, closing double quote, opening single quote, and closing single quote in that order).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/OptionKey/quotes
	TextCheckingQuotesKey NSTextCheckingOptionKey
	// TextCheckingReferenceDateKey is an NSDate to be associated with the document, used as a referent for relative dates; if not specified, the current date will be used.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/OptionKey/referenceDate
	TextCheckingReferenceDateKey NSTextCheckingOptionKey
	// TextCheckingReferenceTimeZoneKey is an NSTimeZone to be associated with the document, used as a reference for dates without time zones; if not specified, the current time zone will be used.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/OptionKey/referenceTimeZone
	TextCheckingReferenceTimeZoneKey NSTextCheckingOptionKey
	// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/OptionKey/regularExpressions
	TextCheckingRegularExpressionsKey NSTextCheckingOptionKey
	// TextCheckingReplacementsKey is an NSDictionary containing replacements to be used with NSTextCheckingTypeReplacement; if not specified, values will be taken from user’s preferences.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/OptionKey/replacements
	TextCheckingReplacementsKey NSTextCheckingOptionKey
	// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/OptionKey/selectedRange
	TextCheckingSelectedRangeKey NSTextCheckingOptionKey
)

var ()

var (
	// TextEffectLetterpressStyle is a graphical text effect that gives glyphs the appearance of letterpress printing, which involves pressing the type into the paper.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextEffectLetterpressStyle
	TextEffectLetterpressStyle NSTextEffectStyle
)

var (
	// TextFinderCaseInsensitiveKey is a Boolean value indicating whether the search is case insensitive.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPasteboard/PasteboardType/TextFinderOptionKey/textFinderCaseInsensitiveKey
	TextFinderCaseInsensitiveKey NSPasteboardTypeTextFinderOptionKey
	// TextFinderMatchingTypeKey is a number object containing the match type to use.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPasteboard/PasteboardType/TextFinderOptionKey/textFinderMatchingTypeKey
	TextFinderMatchingTypeKey NSPasteboardTypeTextFinderOptionKey
)

var ()

var (
	// TextHighlightStyleDefault is the default highlight style to apply to text.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextHighlightStyleDefault
	TextHighlightStyleDefault NSTextHighlightStyle
)

var (
	// TextLayoutSectionOrientation is the orientation of the text.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutSectionOrientation
	TextLayoutSectionOrientation NSTextLayoutSectionKey
	// TextLayoutSectionRange is the character range.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutSectionRange
	TextLayoutSectionRange NSTextLayoutSectionKey
)

var (
	// TextListMarkerBox is the value that represents a square-shaped marker that you can apply to a text list item.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextList/MarkerFormat-swift.struct/box
	TextListMarkerBox NSTextListMarkerFormat
	// TextListMarkerCheck is the value that represents a checkmark-shaped marker that you can apply to a text list item.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextList/MarkerFormat-swift.struct/check
	TextListMarkerCheck NSTextListMarkerFormat
	// TextListMarkerCircle is the value that represents a circle-shaped marker that you can apply to a text list item.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextList/MarkerFormat-swift.struct/circle
	TextListMarkerCircle NSTextListMarkerFormat
	// TextListMarkerDecimal is the value that represents a decimal annotation marker that you can apply to a text list item.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextList/MarkerFormat-swift.struct/decimal
	TextListMarkerDecimal NSTextListMarkerFormat
	// TextListMarkerDiamond is the value that represents a diamond-shaped marker that you can apply to a text list item.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextList/MarkerFormat-swift.struct/diamond
	TextListMarkerDiamond NSTextListMarkerFormat
	// TextListMarkerDisc is the value that represents a disc-shaped marker that you can apply to a text list item.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextList/MarkerFormat-swift.struct/disc
	TextListMarkerDisc NSTextListMarkerFormat
	// TextListMarkerHyphen is the value that represents a hyphen-shaped marker that you can apply to a text list item.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextList/MarkerFormat-swift.struct/hyphen
	TextListMarkerHyphen NSTextListMarkerFormat
	// TextListMarkerLowercaseAlpha is the value that represents a lowercase localized alphabetical marker you that can apply to a text list item.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextList/MarkerFormat-swift.struct/lowercaseAlpha
	TextListMarkerLowercaseAlpha NSTextListMarkerFormat
	// TextListMarkerLowercaseHexadecimal is the value that represents a lowercase hexadecimal (base 16) numerical marker that you can apply to a text list item.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextList/MarkerFormat-swift.struct/lowercaseHexadecimal
	TextListMarkerLowercaseHexadecimal NSTextListMarkerFormat
	// TextListMarkerLowercaseLatin is the value that represents a lowercase Latin alphabetical marker that you can apply to a text list item.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextList/MarkerFormat-swift.struct/lowercaseLatin
	TextListMarkerLowercaseLatin NSTextListMarkerFormat
	// TextListMarkerLowercaseRoman is the value that represents a lowercase Roman numeral marker that you can apply to a text list item.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextList/MarkerFormat-swift.struct/lowercaseRoman
	TextListMarkerLowercaseRoman NSTextListMarkerFormat
	// TextListMarkerOctal is the value that represents an octal (base 8) numerical marker that you can apply to a text list item.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextList/MarkerFormat-swift.struct/octal
	TextListMarkerOctal NSTextListMarkerFormat
	// TextListMarkerSquare is the value that represents a square-shaped marker that you can apply to a text list item.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextList/MarkerFormat-swift.struct/square
	TextListMarkerSquare NSTextListMarkerFormat
	// TextListMarkerUppercaseAlpha is the value that represents an uppercase localized alphabetical marker that you can apply to a text list item.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextList/MarkerFormat-swift.struct/uppercaseAlpha
	TextListMarkerUppercaseAlpha NSTextListMarkerFormat
	// TextListMarkerUppercaseHexadecimal is the value that represents an uppercase hexadecimal (base 16) numerical marker that you can apply to a text list item.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextList/MarkerFormat-swift.struct/uppercaseHexadecimal
	TextListMarkerUppercaseHexadecimal NSTextListMarkerFormat
	// TextListMarkerUppercaseLatin is the value that represents an uppercase Latin alphabetical marker that you can apply to a text list item.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextList/MarkerFormat-swift.struct/uppercaseLatin
	TextListMarkerUppercaseLatin NSTextListMarkerFormat
	// TextListMarkerUppercaseRoman is the value that represents an uppercase Roman numeral marker that you can apply to a text list item.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextList/MarkerFormat-swift.struct/uppercaseRoman
	TextListMarkerUppercaseRoman NSTextListMarkerFormat
)

var (
	// ToolbarCloudSharingItemIdentifier is the identifier for a toolbar item that tells your app to display the iCloud sharing interface.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSToolbarItem/Identifier/cloudSharing
	ToolbarCloudSharingItemIdentifier NSToolbarItemIdentifier
	// ToolbarFlexibleSpaceItemIdentifier is the identifier for a toolbar item that displays an empty space with a flexible width.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSToolbarItem/Identifier/flexibleSpace
	ToolbarFlexibleSpaceItemIdentifier NSToolbarItemIdentifier
	// See: https://developer.apple.com/documentation/AppKit/NSToolbarItem/Identifier/inspectorTrackingSeparator
	ToolbarInspectorTrackingSeparatorItemIdentifier NSToolbarItemIdentifier
	// ToolbarPrintItemIdentifier is the identifier for a toolbar item that tells your app to print the current document.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSToolbarItem/Identifier/print
	ToolbarPrintItemIdentifier NSToolbarItemIdentifier
	// ToolbarShowColorsItemIdentifier is the identifier for a toolbar item that shows the standard color panel.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSToolbarItem/Identifier/showColors
	ToolbarShowColorsItemIdentifier NSToolbarItemIdentifier
	// ToolbarShowFontsItemIdentifier is the identifier for a toolbar item that shows the standard font panel.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSToolbarItem/Identifier/showFonts
	ToolbarShowFontsItemIdentifier NSToolbarItemIdentifier
	// ToolbarSidebarTrackingSeparatorItemIdentifier is the identifier for a toolbar item that displays a tracking separator aligned with the sidebar divider in a split view.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSToolbarItem/Identifier/sidebarTrackingSeparator
	ToolbarSidebarTrackingSeparatorItemIdentifier NSToolbarItemIdentifier
	// ToolbarSpaceItemIdentifier is the identifier for a toolbar item that displays an empty space with a standard fixed size.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSToolbarItem/Identifier/space
	ToolbarSpaceItemIdentifier NSToolbarItemIdentifier
	// See: https://developer.apple.com/documentation/AppKit/NSToolbarItem/Identifier/toggleInspector
	ToolbarToggleInspectorItemIdentifier NSToolbarItemIdentifier
	// ToolbarToggleSidebarItemIdentifier is the identifier for a toolbar item that displays a sidebar.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSToolbarItem/Identifier/toggleSidebar
	ToolbarToggleSidebarItemIdentifier NSToolbarItemIdentifier
	// ToolbarWritingToolsItemIdentifier is a standard item that is configured to send -showWritingTools: to the firstResponder when invoked.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSToolbarItem/Identifier/writingToolsItemIdentifier
	ToolbarWritingToolsItemIdentifier NSToolbarItemIdentifier
)

var (
	// ToolbarItemKey is a key that specifies the toolbar item associated with the notification.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSToolbarUserInfoKey/itemKey
	ToolbarItemKey NSToolbarUserInfoKey
	// See: https://developer.apple.com/documentation/AppKit/NSToolbarUserInfoKey/newIndexKey
	ToolbarNewIndexKey NSToolbarUserInfoKey
)

var ()

var (
	// ViewAnimationEffectKey is an effect to apply to the animation.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSViewAnimation/Key/effect
	ViewAnimationEffectKey NSViewAnimationKey
	// ViewAnimationEndFrameKey is the size and location of the window or view at the end of the animation.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSViewAnimation/Key/endFrame
	ViewAnimationEndFrameKey NSViewAnimationKey
	// ViewAnimationStartFrameKey is the size and location of the window or view at the start of the animation.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSViewAnimation/Key/startFrame
	ViewAnimationStartFrameKey NSViewAnimationKey
	// ViewAnimationTargetKey is the target of the animation.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSViewAnimation/Key/target
	ViewAnimationTargetKey NSViewAnimationKey
)

var (
	// ViewAnimationFadeInEffect is specifies a fade-in type of effect.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSViewAnimation/EffectName/fadeIn
	ViewAnimationFadeInEffect NSViewAnimationEffectName
	// ViewAnimationFadeOutEffect is specifies a fade-out type of effect.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSViewAnimation/EffectName/fadeOut
	ViewAnimationFadeOutEffect NSViewAnimationEffectName
)

var (
	// VoiceAge is the perceived age (in years) of the voice. An [NSString].
	//
	// Deprecated: Deprecated since macOS 14.0. Use AVSpeechSynthesizer in AVFoundation instead
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/VoiceAttributeKey/age
	VoiceAge NSVoiceAttributeKey
	// VoiceDemoText is a demonstration string to speak. An [NSString].
	//
	// Deprecated: Deprecated since macOS 14.0. Use AVSpeechSynthesizer in AVFoundation instead
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/VoiceAttributeKey/demoText
	VoiceDemoText NSVoiceAttributeKey
	// VoiceGender is the perceived gender of the voice. The supported values are listed in `Voice Gender Keys`. An [NSString].
	//
	// Deprecated: Deprecated since macOS 14.0. Use AVSpeechSynthesizer in AVFoundation instead
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/VoiceAttributeKey/gender
	VoiceGender NSVoiceAttributeKey
	// VoiceIdentifier is a unique string identifying the voice. The identifiers of the system voices are listed in `Listing 1`.
	//
	// Deprecated: Deprecated since macOS 14.0. Use AVSpeechSynthesizer in AVFoundation instead
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/VoiceAttributeKey/identifier
	VoiceIdentifier NSVoiceAttributeKey
	// VoiceIndividuallySpokenCharacters is a list of Unicode character id ranges that define the Unicode characters that can be spoken in character-by-character mode by this voice.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/VoiceAttributeKey/individuallySpokenCharacters
	VoiceIndividuallySpokenCharacters NSVoiceAttributeKey
	// VoiceLocaleIdentifier is the language of the voice. An [NSString].
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/VoiceAttributeKey/localeIdentifier
	VoiceLocaleIdentifier NSVoiceAttributeKey
	// VoiceName is the name of the voice suitable for display. An [NSString].
	//
	// Deprecated: Deprecated since macOS 14.0. Use AVSpeechSynthesizer in AVFoundation instead
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/VoiceAttributeKey/name
	VoiceName NSVoiceAttributeKey
	// VoiceSupportedCharacters is a list of Unicode character id ranges that define the Unicode characters supported by this voice.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/VoiceAttributeKey/supportedCharacters
	VoiceSupportedCharacters NSVoiceAttributeKey
)

var (
	// VoiceGenderFemale is a female voice.
	//
	// Deprecated: Deprecated since macOS 14.0. Use AVSpeechSynthesizer in AVFoundation instead
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/VoiceGender/female
	VoiceGenderFemale NSVoiceGenderName
	// VoiceGenderMale is a male voice.
	//
	// Deprecated: Deprecated since macOS 14.0. Use AVSpeechSynthesizer in AVFoundation instead
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/VoiceGender/male
	VoiceGenderMale NSVoiceGenderName
	// VoiceGenderNeuter is a neutral voice (or a novelty voice with a humorous or whimsical quality).
	//
	// Deprecated: Deprecated since macOS 14.0. Use AVSpeechSynthesizer in AVFoundation instead
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/VoiceGender/neuter
	VoiceGenderNeuter NSVoiceGenderName
	// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/VoiceGender/neutral
	VoiceGenderNeutral NSVoiceGenderName
)

var (
	// WorkspaceDesktopImageAllowClippingKey is a key that contains the behavior to use when clipping the image.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/DesktopImageOptionKey/allowClipping
	WorkspaceDesktopImageAllowClippingKey NSWorkspaceDesktopImageOptionKey
	// WorkspaceDesktopImageFillColorKey is a key that contains the behavior to use when filling the empty space around the image.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/DesktopImageOptionKey/fillColor
	WorkspaceDesktopImageFillColorKey NSWorkspaceDesktopImageOptionKey
	// WorkspaceDesktopImageScalingKey is a key that contains the behavior to use when scaling the image.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/DesktopImageOptionKey/imageScaling
	WorkspaceDesktopImageScalingKey NSWorkspaceDesktopImageOptionKey
)

func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAbortModalException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AbortModalException = foundation.NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAbortPrintingException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AbortPrintingException = foundation.NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAboutPanelOptionApplicationIcon"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AboutPanelOptionApplicationIcon = NSAboutPanelOptionKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAboutPanelOptionApplicationName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AboutPanelOptionApplicationName = NSAboutPanelOptionKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAboutPanelOptionApplicationVersion"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AboutPanelOptionApplicationVersion = NSAboutPanelOptionKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAboutPanelOptionCredits"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AboutPanelOptionCredits = NSAboutPanelOptionKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAboutPanelOptionVersion"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AboutPanelOptionVersion = NSAboutPanelOptionKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityActivationPointAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityActivationPointAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityAlternateUIVisibleAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityAlternateUIVisibleAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityAnnotationElement"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityAnnotationElement = NSAccessibilityAnnotationAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityAnnotationLabel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityAnnotationLabel = NSAccessibilityAnnotationAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityAnnotationLocation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityAnnotationLocation = NSAccessibilityAnnotationAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityAnnotationTextAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityAnnotationTextAttribute = foundation.NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityAnnouncementKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityAnnouncementKey = NSAccessibilityNotificationUserInfoKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityAnnouncementRequestedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityAnnouncementRequestedNotification = NSAccessibilityNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityAnyTypeSearchKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityAnyTypeSearchKey = NSAccessibilitySearchKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityApplicationActivatedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityApplicationActivatedNotification = NSAccessibilityNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityApplicationDeactivatedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityApplicationDeactivatedNotification = NSAccessibilityNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityApplicationHiddenNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityApplicationHiddenNotification = NSAccessibilityNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityApplicationRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityApplicationRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityApplicationShownNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityApplicationShownNotification = NSAccessibilityNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityArticleSearchKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityArticleSearchKey = NSAccessibilitySearchKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityAttributedStringForRangeParameterizedAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityAttributedStringForRangeParameterizedAttribute = NSAccessibilityParameterizedAttributeName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityAutoInteractableAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityAutoInteractableAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityAutocorrectedTextAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityAutocorrectedTextAttribute = foundation.NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityAutocorrectionOccurredNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityAutocorrectionOccurredNotification = NSAccessibilityNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityBackgroundColorTextAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityBackgroundColorTextAttribute = foundation.NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityBlockQuoteLevelAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityBlockQuoteLevelAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityBlockquoteSameLevelSearchKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityBlockquoteSameLevelSearchKey = NSAccessibilitySearchKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityBlockquoteSearchKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityBlockquoteSearchKey = NSAccessibilitySearchKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityBoldFontSearchKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityBoldFontSearchKey = NSAccessibilitySearchKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityBoundsForRangeParameterizedAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityBoundsForRangeParameterizedAttribute = NSAccessibilityParameterizedAttributeName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityBrowserRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityBrowserRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityBusyIndicatorRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityBusyIndicatorRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityButtonRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityButtonRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityButtonSearchKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityButtonSearchKey = NSAccessibilitySearchKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityCancelAction"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityCancelAction = NSAccessibilityActionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityCancelButtonAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityCancelButtonAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityCellRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityCellRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityCenterTabStopMarkerTypeValue"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityCenterTabStopMarkerTypeValue = NSAccessibilityRulerMarkerTypeValue(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityCentimetersUnitValue"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityCentimetersUnitValue = NSAccessibilityRulerUnitValue(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityCheckBoxRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityCheckBoxRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityCheckBoxSearchKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityCheckBoxSearchKey = NSAccessibilitySearchKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityChildrenInNavigationOrderAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityChildrenInNavigationOrderAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityCloseButtonSubrole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityCloseButtonSubrole = NSAccessibilitySubrole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityCollectionListSubrole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityCollectionListSubrole = NSAccessibilitySubrole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityColorWellRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityColorWellRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityColumnHeaderUIElementsAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityColumnHeaderUIElementsAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityColumnIndexRangeAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityColumnIndexRangeAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityColumnRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityColumnRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityColumnTitlesAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityColumnTitlesAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityColumnsAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityColumnsAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityComboBoxRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityComboBoxRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityContainsProtectedContentAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityContainsProtectedContentAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityContentListSubrole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityContentListSubrole = NSAccessibilitySubrole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityContentsAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityContentsAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityControlSearchKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityControlSearchKey = NSAccessibilitySearchKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityCreatedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityCreatedNotification = NSAccessibilityNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityCriticalValueAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityCriticalValueAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityCustomTextAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityCustomTextAttribute = foundation.NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityDateTimeAreaRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityDateTimeAreaRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityDateTimeComponentsAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityDateTimeComponentsAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityDecimalTabStopMarkerTypeValue"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityDecimalTabStopMarkerTypeValue = NSAccessibilityRulerMarkerTypeValue(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityDecrementAction"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityDecrementAction = NSAccessibilityActionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityDecrementArrowSubrole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityDecrementArrowSubrole = NSAccessibilitySubrole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityDecrementButtonAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityDecrementButtonAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityDecrementPageSubrole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityDecrementPageSubrole = NSAccessibilitySubrole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityDefaultButtonAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityDefaultButtonAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityDefinitionListSubrole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityDefinitionListSubrole = NSAccessibilitySubrole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityDeleteAction"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityDeleteAction = NSAccessibilityActionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityDescendingSortDirectionValue"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityDescendingSortDirectionValue = NSAccessibilitySortDirectionValue(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityDescriptionAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityDescriptionAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityDescriptionListSubrole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityDescriptionListSubrole = NSAccessibilitySubrole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityDialogSubrole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityDialogSubrole = NSAccessibilitySubrole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityDifferentTypeSearchKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityDifferentTypeSearchKey = NSAccessibilitySearchKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityDisclosedRowsAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityDisclosedRowsAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityDisclosingAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityDisclosingAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityDisclosureLevelAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityDisclosureLevelAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityDisclosureTriangleRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityDisclosureTriangleRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityDocumentAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityDocumentAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityDraggingDestinationDragAcceptedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityDraggingDestinationDragAcceptedNotification = NSAccessibilityNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityDraggingDestinationDragNotAcceptedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityDraggingDestinationDragNotAcceptedNotification = NSAccessibilityNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityDraggingDestinationDropAllowedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityDraggingDestinationDropAllowedNotification = NSAccessibilityNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityDraggingDestinationDropNotAllowedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityDraggingDestinationDropNotAllowedNotification = NSAccessibilityNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityDraggingSourceDragBeganNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityDraggingSourceDragBeganNotification = NSAccessibilityNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityDraggingSourceDragEndedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityDraggingSourceDragEndedNotification = NSAccessibilityNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityDrawerCreatedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityDrawerCreatedNotification = NSAccessibilityNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityDrawerRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityDrawerRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityEditedAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityEditedAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityEmbeddedImageDescriptionAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityEmbeddedImageDescriptionAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityEnabledAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityEnabledAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityErrorCodeExceptionInfo"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityErrorCodeExceptionInfo = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityExpandedAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityExpandedAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityExtrasMenuBarAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityExtrasMenuBarAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityFilenameAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityFilenameAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityFirstLineIndentMarkerTypeValue"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityFirstLineIndentMarkerTypeValue = NSAccessibilityRulerMarkerTypeValue(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityFloatingWindowSubrole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityFloatingWindowSubrole = NSAccessibilitySubrole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityFocusedAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityFocusedAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityFocusedUIElementChangedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityFocusedUIElementChangedNotification = NSAccessibilityNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityFocusedWindowAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityFocusedWindowAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityFocusedWindowChangedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityFocusedWindowChangedNotification = NSAccessibilityNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityFontBoldAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityFontBoldAttribute = foundation.NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityFontChangeSearchKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityFontChangeSearchKey = NSAccessibilitySearchKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityFontColorChangeSearchKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityFontColorChangeSearchKey = NSAccessibilitySearchKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityFontFamilyKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityFontFamilyKey = NSAccessibilityFontAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityFontItalicAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityFontItalicAttribute = foundation.NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityFontNameKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityFontNameKey = NSAccessibilityFontAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityFontSizeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityFontSizeKey = NSAccessibilityFontAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityFontTextAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityFontTextAttribute = foundation.NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityForegroundColorTextAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityForegroundColorTextAttribute = foundation.NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityFrameSearchKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityFrameSearchKey = NSAccessibilitySearchKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityFrontmostAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityFrontmostAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityFullScreenButtonAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityFullScreenButtonAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityFullScreenButtonSubrole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityFullScreenButtonSubrole = NSAccessibilitySubrole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityGraphicSearchKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityGraphicSearchKey = NSAccessibilitySearchKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityGridRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityGridRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityGroupRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityGroupRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityGrowAreaAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityGrowAreaAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityGrowAreaRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityGrowAreaRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityHandleRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityHandleRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityHeadIndentMarkerTypeValue"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityHeadIndentMarkerTypeValue = NSAccessibilityRulerMarkerTypeValue(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityHeaderAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityHeaderAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityHeadingLevel1SearchKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityHeadingLevel1SearchKey = NSAccessibilitySearchKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityHeadingLevel2SearchKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityHeadingLevel2SearchKey = NSAccessibilitySearchKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityHeadingLevel3SearchKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityHeadingLevel3SearchKey = NSAccessibilitySearchKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityHeadingLevel4SearchKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityHeadingLevel4SearchKey = NSAccessibilitySearchKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityHeadingLevel5SearchKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityHeadingLevel5SearchKey = NSAccessibilitySearchKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityHeadingLevel6SearchKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityHeadingLevel6SearchKey = NSAccessibilitySearchKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityHeadingLevelAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityHeadingLevelAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityHeadingRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityHeadingRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityHeadingSameLevelSearchKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityHeadingSameLevelSearchKey = NSAccessibilitySearchKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityHeadingSearchKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityHeadingSearchKey = NSAccessibilitySearchKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityHelpAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityHelpAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityHelpTagCreatedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityHelpTagCreatedNotification = NSAccessibilityNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityHelpTagRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityHelpTagRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityHiddenAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityHiddenAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityHorizontalScrollBarAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityHorizontalScrollBarAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityHorizontalUnitDescriptionAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityHorizontalUnitDescriptionAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityIdentifierAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityIdentifierAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityImageRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityImageRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityIncrementAction"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityIncrementAction = NSAccessibilityActionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityIncrementArrowSubrole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityIncrementArrowSubrole = NSAccessibilitySubrole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityIncrementButtonAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityIncrementButtonAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityIncrementPageSubrole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityIncrementPageSubrole = NSAccessibilitySubrole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityIncrementorRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityIncrementorRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityIndexAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityIndexAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityIndexForChildUIElementAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityIndexForChildUIElementAttribute = NSAccessibilityParameterizedAttributeName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityIndexForChildUIElementInNavigationOrderAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityIndexForChildUIElementInNavigationOrderAttribute = NSAccessibilityParameterizedAttributeName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityItalicFontSearchKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityItalicFontSearchKey = NSAccessibilitySearchKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityKeyboardFocusableSearchKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityKeyboardFocusableSearchKey = NSAccessibilitySearchKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityLabelUIElementsAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityLabelUIElementsAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityLabelValueAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityLabelValueAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityLandmarkSearchKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityLandmarkSearchKey = NSAccessibilitySearchKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityLanguageAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityLanguageAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityLanguageTextAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityLanguageTextAttribute = foundation.NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityLayoutAreaRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityLayoutAreaRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityLayoutChangedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityLayoutChangedNotification = NSAccessibilityNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityLayoutItemRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityLayoutItemRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityLayoutSizeForScreenSizeParameterizedAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityLayoutSizeForScreenSizeParameterizedAttribute = NSAccessibilityParameterizedAttributeName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityLevelIndicatorRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityLevelIndicatorRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityLinkRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityLinkRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityLinkSearchKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityLinkSearchKey = NSAccessibilitySearchKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityLinkTextAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityLinkTextAttribute = foundation.NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityListItemIndexTextAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityListItemIndexTextAttribute = foundation.NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityListItemLevelTextAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityListItemLevelTextAttribute = foundation.NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityListItemPrefixTextAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityListItemPrefixTextAttribute = foundation.NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityListMarkerRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityListMarkerRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityListRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityListRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityListSearchKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityListSearchKey = NSAccessibilitySearchKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityLiveRegionSearchKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityLiveRegionSearchKey = NSAccessibilitySearchKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityMainAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityMainAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityMainWindowAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityMainWindowAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityMainWindowChangedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityMainWindowChangedNotification = NSAccessibilityNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityMarkedMisspelledTextAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityMarkedMisspelledTextAttribute = foundation.NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityMarkerTypeAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityMarkerTypeAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityMarkerTypeDescriptionAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityMarkerTypeDescriptionAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityMarkerUIElementsAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityMarkerUIElementsAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityMarkerValuesAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityMarkerValuesAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityMatteRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityMatteRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityMaxValueAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityMaxValueAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityMenuBarAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityMenuBarAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityMenuBarItemRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityMenuBarItemRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityMenuBarRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityMenuBarRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityMenuButtonRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityMenuButtonRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityMenuItemRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityMenuItemRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityMenuRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityMenuRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityMinValueAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityMinValueAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityMinimizeButtonAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityMinimizeButtonAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityMinimizeButtonSubrole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityMinimizeButtonSubrole = NSAccessibilitySubrole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityMinimizedAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityMinimizedAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityMisspelledTextAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityMisspelledTextAttribute = foundation.NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityMisspelledWordSearchKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityMisspelledWordSearchKey = NSAccessibilitySearchKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityModalAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityModalAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityMovedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityMovedNotification = NSAccessibilityNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityNextContentsAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityNextContentsAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityNumberOfCharactersAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityNumberOfCharactersAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityOrderedByRowAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityOrderedByRowAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityOrientationAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityOrientationAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityOutlineRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityOutlineRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityOutlineRowSubrole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityOutlineRowSubrole = NSAccessibilitySubrole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityOutlineSearchKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityOutlineSearchKey = NSAccessibilitySearchKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityOverflowButtonAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityOverflowButtonAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityPageRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityPageRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityParentAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityParentAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityPathAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityPathAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityPicasUnitValue"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityPicasUnitValue = NSAccessibilityRulerUnitValue(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityPickAction"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityPickAction = NSAccessibilityActionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityPlaceholderValueAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityPlaceholderValueAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityPlainTextSearchKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityPlainTextSearchKey = NSAccessibilitySearchKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityPointsUnitValue"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityPointsUnitValue = NSAccessibilityRulerUnitValue(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityPopUpButtonRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityPopUpButtonRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityPopoverRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityPopoverRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityPositionAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityPositionAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityPressAction"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityPressAction = NSAccessibilityActionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityPreviousContentsAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityPreviousContentsAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityPriorityKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityPriorityKey = NSAccessibilityNotificationUserInfoKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityProgressIndicatorRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityProgressIndicatorRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityProxyAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityProxyAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityRTFForRangeParameterizedAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityRTFForRangeParameterizedAttribute = NSAccessibilityParameterizedAttributeName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityRadioButtonRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityRadioButtonRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityRadioGroupRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityRadioGroupRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityRadioGroupSearchKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityRadioGroupSearchKey = NSAccessibilitySearchKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityRaiseAction"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityRaiseAction = NSAccessibilityActionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityRangeForIndexParameterizedAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityRangeForIndexParameterizedAttribute = NSAccessibilityParameterizedAttributeName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityRangeForLineParameterizedAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityRangeForLineParameterizedAttribute = NSAccessibilityParameterizedAttributeName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityRangeForPositionParameterizedAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityRangeForPositionParameterizedAttribute = NSAccessibilityParameterizedAttributeName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityRatingIndicatorSubrole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityRatingIndicatorSubrole = NSAccessibilitySubrole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityRelevanceIndicatorRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityRelevanceIndicatorRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityRequiredAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityRequiredAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityResizedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityResizedNotification = NSAccessibilityNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityResultsForSearchPredicateParameterizedAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityResultsForSearchPredicateParameterizedAttribute = NSAccessibilityParameterizedAttributeName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityRightTabStopMarkerTypeValue"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityRightTabStopMarkerTypeValue = NSAccessibilityRulerMarkerTypeValue(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityRoleAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityRoleAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityRoleDescriptionAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityRoleDescriptionAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityRowCollapsedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityRowCollapsedNotification = NSAccessibilityNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityRowCountAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityRowCountAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityRowCountChangedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityRowCountChangedNotification = NSAccessibilityNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityRowExpandedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityRowExpandedNotification = NSAccessibilityNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityRowHeaderUIElementsAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityRowHeaderUIElementsAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityRowRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityRowRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityRowsAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityRowsAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityRulerMarkerRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityRulerMarkerRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityRulerRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityRulerRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilitySameTypeSearchKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilitySameTypeSearchKey = NSAccessibilitySearchKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityScreenPointForLayoutPointParameterizedAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityScreenPointForLayoutPointParameterizedAttribute = NSAccessibilityParameterizedAttributeName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityScreenSizeForLayoutSizeParameterizedAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityScreenSizeForLayoutSizeParameterizedAttribute = NSAccessibilityParameterizedAttributeName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityScrollAreaRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityScrollAreaRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityScrollBarRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityScrollBarRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityScrollToVisibleAction"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityScrollToVisibleAction = NSAccessibilityActionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilitySearchButtonAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilitySearchButtonAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilitySearchCurrentElementKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilitySearchCurrentElementKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilitySearchCurrentRangeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilitySearchCurrentRangeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilitySearchDirectionKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilitySearchDirectionKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilitySearchDirectionNext"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilitySearchDirectionNext = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilitySearchDirectionPrevious"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilitySearchDirectionPrevious = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilitySearchFieldSubrole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilitySearchFieldSubrole = NSAccessibilitySubrole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilitySearchIdentifiersKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilitySearchIdentifiersKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilitySearchMenuAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilitySearchMenuAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilitySearchResultDescriptionOverrideKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilitySearchResultDescriptionOverrideKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilitySearchResultElementKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilitySearchResultElementKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilitySearchResultLoaderKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilitySearchResultLoaderKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilitySearchResultRangeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilitySearchResultRangeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilitySearchResultsLimitKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilitySearchResultsLimitKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilitySearchTextKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilitySearchTextKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilitySectionListSubrole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilitySectionListSubrole = NSAccessibilitySubrole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilitySecureTextFieldSubrole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilitySecureTextFieldSubrole = NSAccessibilitySubrole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilitySelectedAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilitySelectedAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilitySelectedCellsChangedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilitySelectedCellsChangedNotification = NSAccessibilityNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilitySelectedChildrenAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilitySelectedChildrenAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilitySelectedChildrenChangedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilitySelectedChildrenChangedNotification = NSAccessibilityNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilitySelectedChildrenMovedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilitySelectedChildrenMovedNotification = NSAccessibilityNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilitySelectedColumnsAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilitySelectedColumnsAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilitySelectedColumnsChangedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilitySelectedColumnsChangedNotification = NSAccessibilityNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilitySelectedRowsAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilitySelectedRowsAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilitySelectedRowsChangedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilitySelectedRowsChangedNotification = NSAccessibilityNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilitySelectedTextAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilitySelectedTextAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilitySelectedTextChangedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilitySelectedTextChangedNotification = NSAccessibilityNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilitySelectedTextRangeAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilitySelectedTextRangeAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilitySelectedTextRangesAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilitySelectedTextRangesAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityServesAsTitleForUIElementsAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityServesAsTitleForUIElementsAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityShadowTextAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityShadowTextAttribute = foundation.NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilitySharedCharacterRangeAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilitySharedCharacterRangeAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilitySharedFocusElementsAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilitySharedFocusElementsAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilitySharedTextUIElementsAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilitySharedTextUIElementsAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilitySheetCreatedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilitySheetCreatedNotification = NSAccessibilityNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilitySheetRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilitySheetRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityShowAlternateUIAction"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityShowAlternateUIAction = NSAccessibilityActionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityShowDefaultUIAction"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityShowDefaultUIAction = NSAccessibilityActionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityShowMenuAction"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityShowMenuAction = NSAccessibilityActionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityShownMenuAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityShownMenuAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilitySizeAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilitySizeAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilitySliderRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilitySliderRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilitySortButtonSubrole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilitySortButtonSubrole = NSAccessibilitySubrole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilitySortDirectionAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilitySortDirectionAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilitySplitGroupRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilitySplitGroupRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilitySplitterRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilitySplitterRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilitySplittersAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilitySplittersAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityStandardWindowSubrole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityStandardWindowSubrole = NSAccessibilitySubrole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityStaticTextRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityStaticTextRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityStaticTextSearchKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityStaticTextSearchKey = NSAccessibilitySearchKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityStrikethroughColorTextAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityStrikethroughColorTextAttribute = foundation.NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityStrikethroughTextAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityStrikethroughTextAttribute = foundation.NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityStringForRangeParameterizedAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityStringForRangeParameterizedAttribute = NSAccessibilityParameterizedAttributeName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityStyleChangeSearchKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityStyleChangeSearchKey = NSAccessibilitySearchKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityStyleRangeForIndexParameterizedAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityStyleRangeForIndexParameterizedAttribute = NSAccessibilityParameterizedAttributeName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilitySubroleAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilitySubroleAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilitySuggestionSubrole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilitySuggestionSubrole = NSAccessibilitySubrole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilitySuperscriptTextAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilitySuperscriptTextAttribute = foundation.NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilitySwitchSubrole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilitySwitchSubrole = NSAccessibilitySubrole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilitySystemDialogSubrole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilitySystemDialogSubrole = NSAccessibilitySubrole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilitySystemFloatingWindowSubrole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilitySystemFloatingWindowSubrole = NSAccessibilitySubrole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilitySystemWideRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilitySystemWideRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityTabButtonSubrole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityTabButtonSubrole = NSAccessibilitySubrole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityTabGroupRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityTabGroupRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityTableRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityTableRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityTableRowSubrole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityTableRowSubrole = NSAccessibilitySubrole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityTableSameLevelSearchKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityTableSameLevelSearchKey = NSAccessibilitySearchKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityTableSearchKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityTableSearchKey = NSAccessibilitySearchKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityTabsAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityTabsAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityTailIndentMarkerTypeValue"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityTailIndentMarkerTypeValue = NSAccessibilityRulerMarkerTypeValue(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityTextAlignmentAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityTextAlignmentAttribute = foundation.NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityTextAreaRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityTextAreaRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityTextAttachmentSubrole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityTextAttachmentSubrole = NSAccessibilitySubrole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityTextCompletionAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityTextCompletionAttribute = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityTextFieldRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityTextFieldRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityTextFieldSearchKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityTextFieldSearchKey = NSAccessibilitySearchKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityTextInputMarkedRangeAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityTextInputMarkedRangeAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityTextInputMarkingSessionBeganNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityTextInputMarkingSessionBeganNotification = NSAccessibilityNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityTextInputMarkingSessionEndedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityTextInputMarkingSessionEndedNotification = NSAccessibilityNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityTextLinkSubrole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityTextLinkSubrole = NSAccessibilitySubrole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityTextStateChangeTypeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityTextStateChangeTypeKey = NSAccessibilitySearchKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityTextStateSyncKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityTextStateSyncKey = NSAccessibilitySearchKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityTimelineSubrole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityTimelineSubrole = NSAccessibilitySubrole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityTitleAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityTitleAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityTitleChangedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityTitleChangedNotification = NSAccessibilityNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityTitleUIElementAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityTitleUIElementAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityToggleSubrole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityToggleSubrole = NSAccessibilitySubrole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityToolbarButtonAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityToolbarButtonAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityToolbarButtonSubrole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityToolbarButtonSubrole = NSAccessibilitySubrole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityToolbarRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityToolbarRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityTopLevelUIElementAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityTopLevelUIElementAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityUIElementDestroyedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityUIElementDestroyedNotification = NSAccessibilityNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityUIElementsForSearchPredicateParameterizedAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityUIElementsForSearchPredicateParameterizedAttribute = NSAccessibilityParameterizedAttributeName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityUIElementsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityUIElementsKey = NSAccessibilityNotificationUserInfoKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityURLAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityURLAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityUnderlineColorTextAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityUnderlineColorTextAttribute = foundation.NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityUnderlineSearchKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityUnderlineSearchKey = NSAccessibilitySearchKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityUnderlineTextAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityUnderlineTextAttribute = foundation.NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityUnitDescriptionAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityUnitDescriptionAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityUnitsAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityUnitsAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityUnitsChangedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityUnitsChangedNotification = NSAccessibilityNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityUnknownMarkerTypeValue"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityUnknownMarkerTypeValue = NSAccessibilityRulerMarkerTypeValue(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityUnknownOrientationValue"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityUnknownOrientationValue = NSAccessibilityOrientationValue(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityUnknownRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityUnknownRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityUnknownSortDirectionValue"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityUnknownSortDirectionValue = NSAccessibilitySortDirectionValue(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityUnknownSubrole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityUnknownSubrole = NSAccessibilitySubrole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityUnknownUnitValue"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityUnknownUnitValue = NSAccessibilityRulerUnitValue(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityUnvisitedLinkSearchKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityUnvisitedLinkSearchKey = NSAccessibilitySearchKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityValueAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityValueAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityValueChangedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityValueChangedNotification = NSAccessibilityNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityValueDescriptionAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityValueDescriptionAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityValueIndicatorRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityValueIndicatorRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityVerticalOrientationValue"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityVerticalOrientationValue = NSAccessibilityOrientationValue(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityVerticalScrollBarAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityVerticalScrollBarAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityVerticalUnitDescriptionAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityVerticalUnitDescriptionAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityVerticalUnitsAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityVerticalUnitsAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityVisibleCellsAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityVisibleCellsAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityVisibleCharacterRangeAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityVisibleCharacterRangeAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityVisibleChildrenAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityVisibleChildrenAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityVisibleColumnsAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityVisibleColumnsAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityVisibleNameKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityVisibleNameKey = NSAccessibilityFontAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityVisibleRowsAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityVisibleRowsAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityVisitedAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityVisitedAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityVisitedLinkSearchKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityVisitedLinkSearchKey = NSAccessibilitySearchKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityWarningValueAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityWarningValueAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityWebAreaRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityWebAreaRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityWindowAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityWindowAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityWindowCreatedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityWindowCreatedNotification = NSAccessibilityNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityWindowDeminiaturizedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityWindowDeminiaturizedNotification = NSAccessibilityNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityWindowMiniaturizedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityWindowMiniaturizedNotification = NSAccessibilityNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityWindowMovedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityWindowMovedNotification = NSAccessibilityNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityWindowResizedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityWindowResizedNotification = NSAccessibilityNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityWindowRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityWindowRole = NSAccessibilityRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityWindowsAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityWindowsAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityZoomButtonAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityZoomButtonAttribute = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAccessibilityZoomButtonSubrole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AccessibilityZoomButtonSubrole = NSAccessibilitySubrole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAdaptiveImageGlyphAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AdaptiveImageGlyphAttributeName = foundation.NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAlignmentBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AlignmentBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAllRomanInputSourcesLocaleIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AllRomanInputSourcesLocaleIdentifier = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAllowsEditingMultipleValuesSelectionBindingOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AllowsEditingMultipleValuesSelectionBindingOption = NSBindingOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAllowsNullArgumentBindingOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AllowsNullArgumentBindingOption = NSBindingOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAlternateImageBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AlternateImageBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAlternateTitleBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AlternateTitleBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAlwaysPresentsApplicationModalAlertsBindingOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AlwaysPresentsApplicationModalAlertsBindingOption = NSBindingOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAnimateBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AnimateBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAnimationDelayBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AnimationDelayBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAnimationProgressMark"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AnimationProgressMark = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAnimationProgressMarkNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AnimationProgressMarkNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAnimationTriggerOrderIn"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AnimationTriggerOrderIn = NSAnimatablePropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAnimationTriggerOrderOut"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AnimationTriggerOrderOut = NSAnimatablePropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAntialiasThresholdChangedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AntialiasThresholdChangedNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSApp"); err == nil && ptr != 0 {
		App = *(*NSApplication)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAppKitIgnoredException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AppKitIgnoredException = foundation.NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAppKitVersionNumber"); err == nil && ptr != 0 {
		NSAppKitVersions.Number = *(*NSAppKitVersion)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAppKitVirtualMemoryException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AppKitVirtualMemoryException = foundation.NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAppearanceDocumentAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AppearanceDocumentAttribute = NSAttributedStringDocumentAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAppearanceNameAccessibilityHighContrastAqua"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSAppearanceNames.AccessibilityHighContrastAqua = NSAppearanceName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAppearanceNameAccessibilityHighContrastDarkAqua"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSAppearanceNames.AccessibilityHighContrastDarkAqua = NSAppearanceName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAppearanceNameAccessibilityHighContrastVibrantDark"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSAppearanceNames.AccessibilityHighContrastVibrantDark = NSAppearanceName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAppearanceNameAccessibilityHighContrastVibrantLight"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSAppearanceNames.AccessibilityHighContrastVibrantLight = NSAppearanceName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAppearanceNameAqua"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSAppearanceNames.Aqua = NSAppearanceName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAppearanceNameDarkAqua"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSAppearanceNames.DarkAqua = NSAppearanceName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAppearanceNameVibrantDark"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSAppearanceNames.VibrantDark = NSAppearanceName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAppearanceNameVibrantLight"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSAppearanceNames.VibrantLight = NSAppearanceName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSApplicationDidBecomeActiveNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ApplicationDidBecomeActiveNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSApplicationDidChangeOcclusionStateNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ApplicationDidChangeOcclusionStateNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSApplicationDidChangeScreenParametersNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ApplicationDidChangeScreenParametersNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSApplicationDidFinishLaunchingNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ApplicationDidFinishLaunchingNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSApplicationDidFinishRestoringWindowsNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ApplicationDidFinishRestoringWindowsNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSApplicationDidHideNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ApplicationDidHideNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSApplicationDidResignActiveNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ApplicationDidResignActiveNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSApplicationDidUnhideNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ApplicationDidUnhideNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSApplicationDidUpdateNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ApplicationDidUpdateNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSApplicationLaunchIsDefaultLaunchKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ApplicationLaunchIsDefaultLaunchKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSApplicationLaunchUserNotificationKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ApplicationLaunchUserNotificationKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSApplicationProtectedDataDidBecomeAvailableNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ApplicationProtectedDataDidBecomeAvailableNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSApplicationProtectedDataWillBecomeUnavailableNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ApplicationProtectedDataWillBecomeUnavailableNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSApplicationShouldBeginSuppressingHighDynamicRangeContentNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ApplicationShouldBeginSuppressingHighDynamicRangeContentNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSApplicationShouldEndSuppressingHighDynamicRangeContentNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ApplicationShouldEndSuppressingHighDynamicRangeContentNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSApplicationWillBecomeActiveNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ApplicationWillBecomeActiveNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSApplicationWillFinishLaunchingNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ApplicationWillFinishLaunchingNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSApplicationWillHideNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ApplicationWillHideNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSApplicationWillResignActiveNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ApplicationWillResignActiveNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSApplicationWillTerminateNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ApplicationWillTerminateNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSApplicationWillUnhideNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ApplicationWillUnhideNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSApplicationWillUpdateNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ApplicationWillUpdateNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSArgumentBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ArgumentBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAttachmentAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AttachmentAttributeName = foundation.NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAttributedStringBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AttributedStringBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAuthorDocumentAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AuthorDocumentAttribute = NSAttributedStringDocumentAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSBackgroundColorAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				BackgroundColorAttributeName = foundation.NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSBackgroundColorDocumentAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				BackgroundColorDocumentAttribute = NSAttributedStringDocumentAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSBackingPropertyOldColorSpaceKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				BackingPropertyOldColorSpaceKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSBackingPropertyOldScaleFactorKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				BackingPropertyOldScaleFactorKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSBadBitmapParametersException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				BadBitmapParametersException = foundation.NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSBadComparisonException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				BadComparisonException = foundation.NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSBadRTFColorTableException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				BadRTFColorTableException = foundation.NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSBadRTFDirectiveException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				BadRTFDirectiveException = foundation.NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSBadRTFFontTableException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				BadRTFFontTableException = foundation.NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSBadRTFStyleSheetException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				BadRTFStyleSheetException = foundation.NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSBaseURLDocumentOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				BaseURLDocumentOption = NSAttributedStringDocumentReadingOptionKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSBaselineOffsetAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				BaselineOffsetAttributeName = foundation.NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSBlack"); err == nil && ptr != 0 {
		Black = *(*float64)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSBottomMarginDocumentAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				BottomMarginDocumentAttribute = NSAttributedStringDocumentAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSBrowserColumnConfigurationDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				BrowserColumnConfigurationDidChangeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSBrowserIllegalDelegateException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				BrowserIllegalDelegateException = foundation.NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSCalibratedRGBColorSpace"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CalibratedRGBColorSpace = NSColorSpaceName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSCalibratedWhiteColorSpace"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CalibratedWhiteColorSpace = NSColorSpaceName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSCategoryDocumentAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CategoryDocumentAttribute = NSAttributedStringDocumentAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSCharacterEncodingDocumentAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CharacterEncodingDocumentAttribute = NSAttributedStringDocumentAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSCharacterEncodingDocumentOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CharacterEncodingDocumentOption = NSAttributedStringDocumentReadingOptionKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSCocoaVersionDocumentAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CocoaVersionDocumentAttribute = NSAttributedStringDocumentAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSCollectionElementKindInterItemGapIndicator"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CollectionElementKindInterItemGapIndicator = NSCollectionViewSupplementaryElementKind(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSCollectionElementKindSectionFooter"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CollectionElementKindSectionFooter = NSCollectionViewSupplementaryElementKind(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSCollectionElementKindSectionHeader"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CollectionElementKindSectionHeader = NSCollectionViewSupplementaryElementKind(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSColorListDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ColorListDidChangeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSColorListIOException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ColorListIOException = foundation.NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSColorListNotEditableException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ColorListNotEditableException = foundation.NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSColorPanelColorDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ColorPanelColorDidChangeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSComboBoxSelectionDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ComboBoxSelectionDidChangeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSComboBoxSelectionIsChangingNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ComboBoxSelectionIsChangingNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSComboBoxWillDismissNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ComboBoxWillDismissNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSComboBoxWillPopUpNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ComboBoxWillPopUpNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSCommentDocumentAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CommentDocumentAttribute = NSAttributedStringDocumentAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSCompanyDocumentAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CompanyDocumentAttribute = NSAttributedStringDocumentAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSConditionallySetsEditableBindingOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ConditionallySetsEditableBindingOption = NSBindingOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSConditionallySetsEnabledBindingOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ConditionallySetsEnabledBindingOption = NSBindingOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSConditionallySetsHiddenBindingOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ConditionallySetsHiddenBindingOption = NSBindingOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSContentArrayBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ContentArrayBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSContentArrayForMultipleSelectionBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ContentArrayForMultipleSelectionBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSContentBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ContentBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSContentDictionaryBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ContentDictionaryBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSContentHeightBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ContentHeightBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSContentObjectBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ContentObjectBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSContentObjectsBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ContentObjectsBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSContentPlacementTagBindingOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ContentPlacementTagBindingOption = NSBindingOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSContentSetBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ContentSetBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSContentValuesBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ContentValuesBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSContentWidthBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ContentWidthBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSContextHelpModeDidActivateNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ContextHelpModeDidActivateNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSContextHelpModeDidDeactivateNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ContextHelpModeDidDeactivateNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSContinuouslyUpdatesValueBindingOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ContinuouslyUpdatesValueBindingOption = NSBindingOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSControlTextDidBeginEditingNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ControlTextDidBeginEditingNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSControlTextDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ControlTextDidChangeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSControlTextDidEndEditingNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ControlTextDidEndEditingNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSConvertedDocumentAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ConvertedDocumentAttribute = NSAttributedStringDocumentAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSCopyrightDocumentAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CopyrightDocumentAttribute = NSAttributedStringDocumentAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSCreatesSortDescriptorBindingOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CreatesSortDescriptorBindingOption = NSBindingOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSCreationTimeDocumentAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CreationTimeDocumentAttribute = NSAttributedStringDocumentAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSCriticalValueBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CriticalValueBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSCursorAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CursorAttributeName = foundation.NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSCustomColorSpace"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CustomColorSpace = NSColorSpaceName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSDarkGray"); err == nil && ptr != 0 {
		DarkGray = *(*float64)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSDataBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				DataBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSDefaultAttributesDocumentAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				DefaultAttributesDocumentAttribute = NSAttributedStringDocumentAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSDefaultAttributesDocumentOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				DefaultAttributesDocumentOption = NSAttributedStringDocumentReadingOptionKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSDefaultFontExcludedDocumentAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				DefaultFontExcludedDocumentAttribute = NSAttributedStringDocumentAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSDefaultTabIntervalDocumentAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				DefaultTabIntervalDocumentAttribute = NSAttributedStringDocumentAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSDefinitionPresentationTypeDictionaryApplication"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSDefinitionPresentationTypes.DictionaryApplication = NSDefinitionPresentationType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSDefinitionPresentationTypeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				DefinitionPresentationTypeKey = NSDefinitionOptionKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSDefinitionPresentationTypeOverlay"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSDefinitionPresentationTypes.Overlay = NSDefinitionPresentationType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSDeletesObjectsOnRemoveBindingsOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				DeletesObjectsOnRemoveBindingsOption = NSBindingOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSDeviceBitsPerSample"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				DeviceBitsPerSample = NSDeviceDescriptionKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSDeviceCMYKColorSpace"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				DeviceCMYKColorSpace = NSColorSpaceName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSDeviceColorSpaceName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				DeviceColorSpaceName = NSDeviceDescriptionKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSDeviceIsPrinter"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				DeviceIsPrinter = NSDeviceDescriptionKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSDeviceIsScreen"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				DeviceIsScreen = NSDeviceDescriptionKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSDeviceRGBColorSpace"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				DeviceRGBColorSpace = NSColorSpaceName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSDeviceResolution"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				DeviceResolution = NSDeviceDescriptionKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSDeviceSize"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				DeviceSize = NSDeviceDescriptionKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSDeviceWhiteColorSpace"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				DeviceWhiteColorSpace = NSColorSpaceName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSDirectionalEdgeInsetsZero"); err == nil && ptr != 0 {
		DirectionalEdgeInsetsZero = *(*NSDirectionalEdgeInsets)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSDisplayNameBindingOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				DisplayNameBindingOption = NSBindingOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSDisplayPatternBindingOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				DisplayPatternBindingOption = NSBindingOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSDisplayPatternTitleBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				DisplayPatternTitleBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSDisplayPatternValueBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				DisplayPatternValueBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSDocFormatTextDocumentType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				DocFormatTextDocumentType = NSAttributedStringDocumentType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSDocumentEditedBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				DocumentEditedBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSDocumentTypeDocumentAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				DocumentTypeDocumentAttribute = NSAttributedStringDocumentAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSDocumentTypeDocumentOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				DocumentTypeDocumentOption = NSAttributedStringDocumentReadingOptionKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSDoubleClickArgumentBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				DoubleClickArgumentBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSDoubleClickTargetBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				DoubleClickTargetBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSDraggingException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				DraggingException = foundation.NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSDraggingImageComponentIconKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				DraggingImageComponentIconKey = NSDraggingImageComponentKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSDraggingImageComponentLabelKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				DraggingImageComponentLabelKey = NSDraggingImageComponentKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSEditableBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				EditableBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSEditorDocumentAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				EditorDocumentAttribute = NSAttributedStringDocumentAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSEnabledBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				EnabledBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSEventTrackingRunLoopMode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				EventTrackingRunLoopMode = foundation.NSRunLoopMode(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSExcludedElementsDocumentAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ExcludedElementsDocumentAttribute = NSAttributedStringDocumentAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSExcludedKeysBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ExcludedKeysBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSExpansionAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ExpansionAttributeName = foundation.NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFileContentsPboardType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FileContentsPboardType = NSPasteboardType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFileTypeDocumentAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FileTypeDocumentAttribute = NSAttributedStringDocumentAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFileTypeDocumentOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FileTypeDocumentOption = NSAttributedStringDocumentReadingOptionKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFilterPredicateBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FilterPredicateBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFindPanelCaseInsensitiveSearch"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FindPanelCaseInsensitiveSearch = NSPasteboardTypeFindPanelSearchOptionKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFindPanelSearchOptionsPboardType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FindPanelSearchOptionsPboardType = NSPasteboardType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFindPanelSubstringMatch"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FindPanelSubstringMatch = NSPasteboardTypeFindPanelSearchOptionKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FontAttributeName = foundation.NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FontBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontBoldBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FontBoldBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontCascadeListAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FontCascadeListAttribute = NSFontDescriptorAttributeName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontCharacterSetAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FontCharacterSetAttribute = NSFontDescriptorAttributeName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontCollectionActionKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FontCollectionActionKey = NSFontCollectionUserInfoKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontCollectionAllFonts"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FontCollectionAllFonts = NSFontCollectionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontCollectionDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FontCollectionDidChangeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontCollectionDisallowAutoActivationOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FontCollectionDisallowAutoActivationOption = NSFontCollectionMatchingOptionKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontCollectionFavorites"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FontCollectionFavorites = NSFontCollectionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontCollectionIncludeDisabledFontsOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FontCollectionIncludeDisabledFontsOption = NSFontCollectionMatchingOptionKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontCollectionNameKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FontCollectionNameKey = NSFontCollectionUserInfoKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontCollectionOldNameKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FontCollectionOldNameKey = NSFontCollectionUserInfoKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontCollectionRecentlyUsed"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FontCollectionRecentlyUsed = NSFontCollectionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontCollectionRemoveDuplicatesOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FontCollectionRemoveDuplicatesOption = NSFontCollectionMatchingOptionKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontCollectionUser"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FontCollectionUser = NSFontCollectionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontCollectionVisibilityKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FontCollectionVisibilityKey = NSFontCollectionUserInfoKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontCollectionWasHidden"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FontCollectionWasHidden = NSFontCollectionActionTypeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontCollectionWasRenamed"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FontCollectionWasRenamed = NSFontCollectionActionTypeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontCollectionWasShown"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FontCollectionWasShown = NSFontCollectionActionTypeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontDescriptorSystemDesignDefault"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSFontDescriptorSystemDesigns.Default = NSFontDescriptorSystemDesign(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontDescriptorSystemDesignMonospaced"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSFontDescriptorSystemDesigns.Monospaced = NSFontDescriptorSystemDesign(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontDescriptorSystemDesignRounded"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSFontDescriptorSystemDesigns.Rounded = NSFontDescriptorSystemDesign(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontDescriptorSystemDesignSerif"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSFontDescriptorSystemDesigns.Serif = NSFontDescriptorSystemDesign(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontFaceAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FontFaceAttribute = NSFontDescriptorAttributeName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontFamilyAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FontFamilyAttribute = NSFontDescriptorAttributeName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontFamilyNameBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FontFamilyNameBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontFeatureSelectorIdentifierKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FontFeatureSelectorIdentifierKey = NSFontDescriptorFeatureKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontFeatureSettingsAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FontFeatureSettingsAttribute = NSFontDescriptorAttributeName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontFeatureTypeIdentifierKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FontFeatureTypeIdentifierKey = NSFontDescriptorFeatureKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontFixedAdvanceAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FontFixedAdvanceAttribute = NSFontDescriptorAttributeName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontIdentityMatrix"); err == nil && ptr != 0 {
		FontIdentityMatrix = *(*float64)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontItalicBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FontItalicBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontMatrixAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FontMatrixAttribute = NSFontDescriptorAttributeName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontNameAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FontNameAttribute = NSFontDescriptorAttributeName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontNameBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FontNameBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontSetChangedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FontSetChangedNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontSizeAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FontSizeAttribute = NSFontDescriptorAttributeName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontSizeBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FontSizeBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontSlantTrait"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FontSlantTrait = NSFontDescriptorTraitKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontSymbolicTrait"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FontSymbolicTrait = NSFontDescriptorTraitKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontTextStyleBody"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSFontTextStyles.Body = NSFontTextStyle(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontTextStyleCallout"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSFontTextStyles.Callout = NSFontTextStyle(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontTextStyleCaption1"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSFontTextStyles.Caption1 = NSFontTextStyle(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontTextStyleCaption2"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSFontTextStyles.Caption2 = NSFontTextStyle(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontTextStyleFootnote"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSFontTextStyles.Footnote = NSFontTextStyle(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontTextStyleHeadline"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSFontTextStyles.Headline = NSFontTextStyle(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontTextStyleLargeTitle"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSFontTextStyles.LargeTitle = NSFontTextStyle(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontTextStyleSubheadline"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSFontTextStyles.Subheadline = NSFontTextStyle(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontTextStyleTitle1"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSFontTextStyles.Title1 = NSFontTextStyle(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontTextStyleTitle2"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSFontTextStyles.Title2 = NSFontTextStyle(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontTextStyleTitle3"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSFontTextStyles.Title3 = NSFontTextStyle(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontTraitsAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FontTraitsAttribute = NSFontDescriptorAttributeName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontUnavailableException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FontUnavailableException = foundation.NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontVariationAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FontVariationAttribute = NSFontDescriptorAttributeName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontVariationAxisDefaultValueKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FontVariationAxisDefaultValueKey = NSFontDescriptorVariationKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontVariationAxisIdentifierKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FontVariationAxisIdentifierKey = NSFontDescriptorVariationKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontVariationAxisMaximumValueKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FontVariationAxisMaximumValueKey = NSFontDescriptorVariationKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontVariationAxisMinimumValueKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FontVariationAxisMinimumValueKey = NSFontDescriptorVariationKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontVariationAxisNameKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FontVariationAxisNameKey = NSFontDescriptorVariationKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontVisibleNameAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FontVisibleNameAttribute = NSFontDescriptorAttributeName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontWeightBlack"); err == nil && ptr != 0 {
		NSFontWeights.Black = *(*NSFontWeight)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontWeightBold"); err == nil && ptr != 0 {
		NSFontWeights.Bold = *(*NSFontWeight)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontWeightHeavy"); err == nil && ptr != 0 {
		NSFontWeights.Heavy = *(*NSFontWeight)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontWeightLight"); err == nil && ptr != 0 {
		NSFontWeights.Light = *(*NSFontWeight)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontWeightMedium"); err == nil && ptr != 0 {
		NSFontWeights.Medium = *(*NSFontWeight)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontWeightRegular"); err == nil && ptr != 0 {
		NSFontWeights.Regular = *(*NSFontWeight)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontWeightSemibold"); err == nil && ptr != 0 {
		NSFontWeights.Semibold = *(*NSFontWeight)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontWeightThin"); err == nil && ptr != 0 {
		NSFontWeights.Thin = *(*NSFontWeight)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontWeightTrait"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FontWeightTrait = NSFontDescriptorTraitKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontWeightUltraLight"); err == nil && ptr != 0 {
		NSFontWeights.UltraLight = *(*NSFontWeight)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontWidthCompressed"); err == nil && ptr != 0 {
		NSFontWidths.Compressed = *(*NSFontWidth)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontWidthCondensed"); err == nil && ptr != 0 {
		NSFontWidths.Condensed = *(*NSFontWidth)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontWidthExpanded"); err == nil && ptr != 0 {
		NSFontWidths.Expanded = *(*NSFontWidth)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontWidthStandard"); err == nil && ptr != 0 {
		NSFontWidths.Standard = *(*NSFontWidth)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFontWidthTrait"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FontWidthTrait = NSFontDescriptorTraitKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSForegroundColorAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ForegroundColorAttributeName = foundation.NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFullScreenModeAllScreens"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FullScreenModeAllScreens = NSViewFullScreenModeOptionKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFullScreenModeApplicationPresentationOptions"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FullScreenModeApplicationPresentationOptions = NSViewFullScreenModeOptionKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFullScreenModeSetting"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FullScreenModeSetting = NSViewFullScreenModeOptionKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFullScreenModeWindowLevel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FullScreenModeWindowLevel = NSViewFullScreenModeOptionKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSGlyphInfoAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GlyphInfoAttributeName = foundation.NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSGraphicsContextDestinationAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GraphicsContextDestinationAttributeName = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSGraphicsContextPDFFormat"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GraphicsContextPDFFormat = NSGraphicsContextRepresentationFormatName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSGraphicsContextPSFormat"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GraphicsContextPSFormat = NSGraphicsContextRepresentationFormatName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSGraphicsContextRepresentationFormatAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GraphicsContextRepresentationFormatAttributeName = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSGridViewSizeForContent"); err == nil && ptr != 0 {
		GridViewSizeForContent = *(*float64)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSHTMLTextDocumentType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				HTMLTextDocumentType = NSAttributedStringDocumentType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSHandlesContentAsCompoundValueBindingOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				HandlesContentAsCompoundValueBindingOption = NSBindingOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSHeaderTitleBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				HeaderTitleBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSHiddenBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				HiddenBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSHyphenationFactorDocumentAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				HyphenationFactorDocumentAttribute = NSAttributedStringDocumentAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSIllegalSelectorException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IllegalSelectorException = foundation.NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ImageBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageCacheException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ImageCacheException = foundation.NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageColorSyncProfileData"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ImageColorSyncProfileData = NSBitmapImageRepPropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageCompressionFactor"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ImageCompressionFactor = NSBitmapImageRepPropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageCompressionMethod"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ImageCompressionMethod = NSBitmapImageRepPropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageCurrentFrame"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ImageCurrentFrame = NSBitmapImageRepPropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageCurrentFrameDuration"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ImageCurrentFrameDuration = NSBitmapImageRepPropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageDitherTransparency"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ImageDitherTransparency = NSBitmapImageRepPropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageEXIFData"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ImageEXIFData = NSBitmapImageRepPropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageFallbackBackgroundColor"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ImageFallbackBackgroundColor = NSBitmapImageRepPropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageFrameCount"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ImageFrameCount = NSBitmapImageRepPropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageGamma"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ImageGamma = NSBitmapImageRepPropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageHintCTM"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ImageHintCTM = NSImageHintKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageHintInterpolation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ImageHintInterpolation = NSImageHintKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageHintUserInterfaceLayoutDirection"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ImageHintUserInterfaceLayoutDirection = NSImageHintKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageIPTCData"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ImageIPTCData = NSBitmapImageRepPropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageInterlaced"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ImageInterlaced = NSBitmapImageRepPropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageLoopCount"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ImageLoopCount = NSBitmapImageRepPropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameActionTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.ActionTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameAddTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.AddTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameAdvanced"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.Advanced = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameApplicationIcon"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.ApplicationIcon = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameBluetoothTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.BluetoothTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameBonjour"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.Bonjour = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameBookmarksTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.BookmarksTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameCaution"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.Caution = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameColorPanel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.ColorPanel = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameColumnViewTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.ColumnViewTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameComputer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.Computer = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameEnterFullScreenTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.EnterFullScreenTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameEveryone"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.Everyone = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameExitFullScreenTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.ExitFullScreenTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameFlowViewTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.FlowViewTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameFolder"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.Folder = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameFolderBurnable"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.FolderBurnable = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameFolderSmart"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.FolderSmart = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameFollowLinkFreestandingTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.FollowLinkFreestandingTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameFontPanel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.FontPanel = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameGoBackTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.GoBackTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameGoForwardTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.GoForwardTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameGoLeftTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.GoLeftTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameGoRightTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.GoRightTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameHomeTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.HomeTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameIChatTheaterTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.IChatTheaterTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameIconViewTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.IconViewTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameInfo"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.Info = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameInvalidDataFreestandingTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.InvalidDataFreestandingTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameLeftFacingTriangleTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.LeftFacingTriangleTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameListViewTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.ListViewTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameLockLockedTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.LockLockedTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameLockUnlockedTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.LockUnlockedTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameMenuMixedStateTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.MenuMixedStateTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameMenuOnStateTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.MenuOnStateTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameMobileMe"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.MobileMe = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameMultipleDocuments"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.MultipleDocuments = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameNetwork"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.Network = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNamePathTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.PathTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNamePreferencesGeneral"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.PreferencesGeneral = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameQuickLookTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.QuickLookTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameRefreshFreestandingTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.RefreshFreestandingTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameRefreshTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.RefreshTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameRemoveTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.RemoveTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameRevealFreestandingTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.RevealFreestandingTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameRightFacingTriangleTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.RightFacingTriangleTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameShareTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.ShareTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameSlideshowTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.SlideshowTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameSmartBadgeTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.SmartBadgeTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameStatusAvailable"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.StatusAvailable = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameStatusNone"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.StatusNone = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameStatusPartiallyAvailable"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.StatusPartiallyAvailable = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameStatusUnavailable"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.StatusUnavailable = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameStopProgressFreestandingTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.StopProgressFreestandingTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameStopProgressTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.StopProgressTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarAddDetailTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarAddDetailTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarAddTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarAddTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarAlarmTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarAlarmTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarAudioInputMuteTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarAudioInputMuteTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarAudioInputTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarAudioInputTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarAudioOutputMuteTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarAudioOutputMuteTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarAudioOutputVolumeHighTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarAudioOutputVolumeHighTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarAudioOutputVolumeLowTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarAudioOutputVolumeLowTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarAudioOutputVolumeMediumTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarAudioOutputVolumeMediumTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarAudioOutputVolumeOffTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarAudioOutputVolumeOffTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarBookmarksTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarBookmarksTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarColorPickerFill"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarColorPickerFill = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarColorPickerFont"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarColorPickerFont = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarColorPickerStroke"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarColorPickerStroke = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarCommunicationAudioTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarCommunicationAudioTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarCommunicationVideoTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarCommunicationVideoTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarComposeTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarComposeTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarDeleteTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarDeleteTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarDownloadTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarDownloadTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarEnterFullScreenTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarEnterFullScreenTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarExitFullScreenTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarExitFullScreenTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarFastForwardTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarFastForwardTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarFolderCopyToTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarFolderCopyToTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarFolderMoveToTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarFolderMoveToTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarFolderTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarFolderTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarGetInfoTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarGetInfoTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarGoBackTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarGoBackTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarGoDownTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarGoDownTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarGoForwardTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarGoForwardTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarGoUpTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarGoUpTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarHistoryTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarHistoryTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarIconViewTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarIconViewTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarListViewTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarListViewTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarMailTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarMailTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarNewFolderTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarNewFolderTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarNewMessageTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarNewMessageTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarOpenInBrowserTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarOpenInBrowserTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarPauseTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarPauseTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarPlayPauseTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarPlayPauseTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarPlayTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarPlayTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarPlayheadTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarPlayheadTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarQuickLookTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarQuickLookTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarRecordStartTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarRecordStartTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarRecordStopTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarRecordStopTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarRefreshTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarRefreshTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarRemoveTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarRemoveTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarRewindTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarRewindTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarRotateLeftTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarRotateLeftTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarRotateRightTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarRotateRightTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarSearchTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarSearchTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarShareTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarShareTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarSidebarTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarSidebarTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarSkipAhead15SecondsTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarSkipAhead15SecondsTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarSkipAhead30SecondsTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarSkipAhead30SecondsTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarSkipAheadTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarSkipAheadTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarSkipBack15SecondsTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarSkipBack15SecondsTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarSkipBack30SecondsTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarSkipBack30SecondsTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarSkipBackTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarSkipBackTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarSkipToEndTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarSkipToEndTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarSkipToStartTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarSkipToStartTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarSlideshowTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarSlideshowTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarTagIconTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarTagIconTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarTextBoldTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarTextBoldTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarTextBoxTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarTextBoxTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarTextCenterAlignTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarTextCenterAlignTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarTextItalicTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarTextItalicTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarTextJustifiedAlignTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarTextJustifiedAlignTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarTextLeftAlignTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarTextLeftAlignTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarTextListTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarTextListTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarTextRightAlignTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarTextRightAlignTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarTextStrikethroughTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarTextStrikethroughTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarTextUnderlineTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarTextUnderlineTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarUserAddTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarUserAddTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarUserGroupTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarUserGroupTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarUserTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarUserTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarVolumeDownTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarVolumeDownTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTouchBarVolumeUpTemplate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TouchBarVolumeUpTemplate = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTrashEmpty"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TrashEmpty = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameTrashFull"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.TrashFull = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameUser"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.User = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameUserAccounts"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.UserAccounts = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameUserGroup"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.UserGroup = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageNameUserGuest"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSImageNames.UserGuest = NSImageName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageProgressive"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ImageProgressive = NSBitmapImageRepPropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageRGBColorTable"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ImageRGBColorTable = NSBitmapImageRepPropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageRepRegistryDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ImageRepRegistryDidChangeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSIncludedKeysBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IncludedKeysBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSInitialKeyBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				InitialKeyBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSInitialValueBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				InitialValueBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSInsertsNullPlaceholderBindingOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				InsertsNullPlaceholderBindingOption = NSBindingOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSInvokesSeparatelyWithArrayObjectsBindingOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				InvokesSeparatelyWithArrayObjectsBindingOption = NSBindingOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSIsIndeterminateBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IsIndeterminateBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSKernAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KernAttributeName = foundation.NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSKeywordsDocumentAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KeywordsDocumentAttribute = NSAttributedStringDocumentAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLabelBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				LabelBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLeftMarginDocumentAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				LeftMarginDocumentAttribute = NSAttributedStringDocumentAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLigatureAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				LigatureAttributeName = foundation.NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLightGray"); err == nil && ptr != 0 {
		LightGray = *(*float64)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLinkAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				LinkAttributeName = foundation.NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLocalizedKeyDictionaryBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				LocalizedKeyDictionaryBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMacSimpleTextDocumentType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MacSimpleTextDocumentType = NSAttributedStringDocumentType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSManagedObjectContextBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ManagedObjectContextBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSManagerDocumentAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ManagerDocumentAttribute = NSAttributedStringDocumentAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMarkedClauseSegmentAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MarkedClauseSegmentAttributeName = foundation.NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMaxValueBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MaxValueBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMaxWidthBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MaxWidthBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMaximumRecentsBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MaximumRecentsBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMenuDidAddItemNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MenuDidAddItemNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMenuDidBeginTrackingNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MenuDidBeginTrackingNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMenuDidChangeItemNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MenuDidChangeItemNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMenuDidEndTrackingNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MenuDidEndTrackingNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMenuDidRemoveItemNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MenuDidRemoveItemNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMenuDidSendActionNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MenuDidSendActionNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMenuItemImportFromDeviceIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MenuItemImportFromDeviceIdentifier = NSUserInterfaceItemIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMenuWillSendActionNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MenuWillSendActionNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMinValueBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MinValueBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMinWidthBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MinWidthBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMixedStateImageBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MixedStateImageBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSModalPanelRunLoopMode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ModalPanelRunLoopMode = foundation.NSRunLoopMode(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSModificationTimeDocumentAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ModificationTimeDocumentAttribute = NSAttributedStringDocumentAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMultipleValuesPlaceholderBindingOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MultipleValuesPlaceholderBindingOption = NSBindingOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSNamedColorSpace"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NamedColorSpace = NSColorSpaceName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSNibLoadingException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NibLoadingException = foundation.NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSNoSelectionPlaceholderBindingOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NoSelectionPlaceholderBindingOption = NSBindingOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSNotApplicablePlaceholderBindingOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NotApplicablePlaceholderBindingOption = NSBindingOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSNullPlaceholderBindingOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NullPlaceholderBindingOption = NSBindingOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSObliquenessAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ObliquenessAttributeName = foundation.NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSObservedKeyPathKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ObservedKeyPathKey = NSBindingInfoKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSObservedObjectKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ObservedObjectKey = NSBindingInfoKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSOffStateImageBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				OffStateImageBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSOfficeOpenXMLTextDocumentType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				OfficeOpenXMLTextDocumentType = NSAttributedStringDocumentType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSOnStateImageBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				OnStateImageBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSOpenDocumentTextDocumentType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				OpenDocumentTextDocumentType = NSAttributedStringDocumentType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSOptionsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				OptionsKey = NSBindingInfoKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSOutlineViewColumnDidMoveNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				OutlineViewColumnDidMoveNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSOutlineViewColumnDidResizeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				OutlineViewColumnDidResizeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSOutlineViewDisclosureButtonKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				OutlineViewDisclosureButtonKey = NSUserInterfaceItemIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSOutlineViewItemDidCollapseNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				OutlineViewItemDidCollapseNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSOutlineViewItemDidExpandNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				OutlineViewItemDidExpandNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSOutlineViewItemWillCollapseNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				OutlineViewItemWillCollapseNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSOutlineViewItemWillExpandNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				OutlineViewItemWillExpandNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSOutlineViewSelectionDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				OutlineViewSelectionDidChangeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSOutlineViewSelectionIsChangingNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				OutlineViewSelectionIsChangingNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSOutlineViewShowHideButtonKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				OutlineViewShowHideButtonKey = NSUserInterfaceItemIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPPDIncludeNotFoundException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PPDIncludeNotFoundException = foundation.NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPPDIncludeStackOverflowException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PPDIncludeStackOverflowException = foundation.NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPPDIncludeStackUnderflowException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PPDIncludeStackUnderflowException = foundation.NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPPDParseException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PPDParseException = foundation.NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPaperSizeDocumentAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PaperSizeDocumentAttribute = NSAttributedStringDocumentAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSParagraphStyleAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ParagraphStyleAttributeName = foundation.NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPasteboardCommunicationException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PasteboardCommunicationException = foundation.NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPasteboardDetectionPatternCalendarEvent"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSPasteboardDetectionPatterns.CalendarEvent = NSPasteboardDetectionPattern(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPasteboardDetectionPatternEmailAddress"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSPasteboardDetectionPatterns.EmailAddress = NSPasteboardDetectionPattern(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPasteboardDetectionPatternFlightNumber"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSPasteboardDetectionPatterns.FlightNumber = NSPasteboardDetectionPattern(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPasteboardDetectionPatternLink"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSPasteboardDetectionPatterns.Link = NSPasteboardDetectionPattern(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPasteboardDetectionPatternMoneyAmount"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSPasteboardDetectionPatterns.MoneyAmount = NSPasteboardDetectionPattern(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPasteboardDetectionPatternNumber"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSPasteboardDetectionPatterns.Number = NSPasteboardDetectionPattern(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPasteboardDetectionPatternPhoneNumber"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSPasteboardDetectionPatterns.PhoneNumber = NSPasteboardDetectionPattern(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPasteboardDetectionPatternPostalAddress"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSPasteboardDetectionPatterns.PostalAddress = NSPasteboardDetectionPattern(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPasteboardDetectionPatternProbableWebSearch"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSPasteboardDetectionPatterns.ProbableWebSearch = NSPasteboardDetectionPattern(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPasteboardDetectionPatternProbableWebURL"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSPasteboardDetectionPatterns.ProbableWebURL = NSPasteboardDetectionPattern(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPasteboardDetectionPatternShipmentTrackingNumber"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSPasteboardDetectionPatterns.ShipmentTrackingNumber = NSPasteboardDetectionPattern(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPasteboardMetadataTypeContentType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PasteboardMetadataTypeContentType = NSPasteboardMetadataType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPasteboardNameDrag"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSPasteboardNames.Drag = NSPasteboardName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPasteboardNameFind"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSPasteboardNames.Find = NSPasteboardName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPasteboardNameFont"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSPasteboardNames.Font = NSPasteboardName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPasteboardNameGeneral"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSPasteboardNames.General = NSPasteboardName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPasteboardNameRuler"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSPasteboardNames.Ruler = NSPasteboardName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPasteboardTypeColor"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSPasteboardTypes.Color = NSPasteboardType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPasteboardTypeFileURL"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSPasteboardTypes.FileURL = NSPasteboardType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPasteboardTypeFont"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSPasteboardTypes.Font = NSPasteboardType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPasteboardTypeHTML"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSPasteboardTypes.HTML = NSPasteboardType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPasteboardTypeMultipleTextSelection"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSPasteboardTypes.MultipleTextSelection = NSPasteboardType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPasteboardTypePDF"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSPasteboardTypes.PDF = NSPasteboardType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPasteboardTypePNG"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSPasteboardTypes.PNG = NSPasteboardType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPasteboardTypeRTF"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSPasteboardTypes.RTF = NSPasteboardType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPasteboardTypeRTFD"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSPasteboardTypes.RTFD = NSPasteboardType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPasteboardTypeRuler"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSPasteboardTypes.Ruler = NSPasteboardType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPasteboardTypeSound"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSPasteboardTypes.Sound = NSPasteboardType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPasteboardTypeString"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSPasteboardTypes.String = NSPasteboardType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPasteboardTypeTIFF"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSPasteboardTypes.TIFF = NSPasteboardType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPasteboardTypeTabularText"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSPasteboardTypes.TabularText = NSPasteboardType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPasteboardTypeTextFinderOptions"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSPasteboardTypes.TextFinderOptions = NSPasteboardType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPasteboardTypeURL"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSPasteboardTypes.URL = NSPasteboardType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPasteboardURLReadingContentsConformToTypesKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PasteboardURLReadingContentsConformToTypesKey = NSPasteboardReadingOptionKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPasteboardURLReadingFileURLsOnlyKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PasteboardURLReadingFileURLsOnlyKey = NSPasteboardReadingOptionKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPatternColorSpace"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PatternColorSpace = NSColorSpaceName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPlainTextDocumentType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PlainTextDocumentType = NSAttributedStringDocumentType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPopUpButtonCellWillPopUpNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PopUpButtonCellWillPopUpNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPopUpButtonWillPopUpNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PopUpButtonWillPopUpNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPopoverCloseReasonDetachToWindow"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PopoverCloseReasonDetachToWindow = NSPopoverCloseReasonValue(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPopoverCloseReasonKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PopoverCloseReasonKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPopoverCloseReasonStandard"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PopoverCloseReasonStandard = NSPopoverCloseReasonValue(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPopoverDidCloseNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PopoverDidCloseNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPopoverDidShowNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PopoverDidShowNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPopoverWillCloseNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PopoverWillCloseNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPopoverWillShowNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PopoverWillShowNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPositioningRectBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PositioningRectBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPredicateBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PredicateBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPredicateFormatBindingOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PredicateFormatBindingOption = NSBindingOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPreferredScrollerStyleDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PreferredScrollerStyleDidChangeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPrefixSpacesDocumentAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PrefixSpacesDocumentAttribute = NSAttributedStringDocumentAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPrintAllPages"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PrintAllPages = NSPrintInfoAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPrintAllPresetsJobStyleHint"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PrintAllPresetsJobStyleHint = NSPrintPanelJobStyleHint(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPrintBottomMargin"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PrintBottomMargin = NSPrintInfoAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPrintCancelJob"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PrintCancelJob = NSPrintJobDispositionValue(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPrintCopies"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PrintCopies = NSPrintInfoAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPrintDetailedErrorReporting"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PrintDetailedErrorReporting = NSPrintInfoAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPrintFaxNumber"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PrintFaxNumber = NSPrintInfoAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPrintFirstPage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PrintFirstPage = NSPrintInfoAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPrintHeaderAndFooter"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PrintHeaderAndFooter = NSPrintInfoAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPrintHorizontalPagination"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PrintHorizontalPagination = NSPrintInfoAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPrintHorizontallyCentered"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PrintHorizontallyCentered = NSPrintInfoAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPrintJobDisposition"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PrintJobDisposition = NSPrintInfoAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPrintJobSavingFileNameExtensionHidden"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PrintJobSavingFileNameExtensionHidden = NSPrintInfoAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPrintJobSavingURL"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PrintJobSavingURL = NSPrintInfoAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPrintLastPage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PrintLastPage = NSPrintInfoAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPrintLeftMargin"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PrintLeftMargin = NSPrintInfoAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPrintMustCollate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PrintMustCollate = NSPrintInfoAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPrintNoPresetsJobStyleHint"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PrintNoPresetsJobStyleHint = NSPrintPanelJobStyleHint(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPrintOperationExistsException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PrintOperationExistsException = foundation.NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPrintOrientation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PrintOrientation = NSPrintInfoAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPrintPackageException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PrintPackageException = foundation.NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPrintPagesAcross"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PrintPagesAcross = NSPrintInfoAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPrintPagesDown"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PrintPagesDown = NSPrintInfoAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPrintPanelAccessorySummaryItemDescriptionKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PrintPanelAccessorySummaryItemDescriptionKey = NSPrintPanelAccessorySummaryKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPrintPanelAccessorySummaryItemNameKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PrintPanelAccessorySummaryItemNameKey = NSPrintPanelAccessorySummaryKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPrintPaperName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PrintPaperName = NSPrintInfoAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPrintPaperSize"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PrintPaperSize = NSPrintInfoAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPrintPhotoJobStyleHint"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PrintPhotoJobStyleHint = NSPrintPanelJobStyleHint(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPrintPreviewJob"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PrintPreviewJob = NSPrintJobDispositionValue(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPrintPrinter"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PrintPrinter = NSPrintInfoAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPrintPrinterName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PrintPrinterName = NSPrintInfoAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPrintReversePageOrder"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PrintReversePageOrder = NSPrintInfoAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPrintRightMargin"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PrintRightMargin = NSPrintInfoAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPrintSaveJob"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PrintSaveJob = NSPrintJobDispositionValue(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPrintScalingFactor"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PrintScalingFactor = NSPrintInfoAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPrintSelectionOnly"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PrintSelectionOnly = NSPrintInfoAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPrintSpoolJob"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PrintSpoolJob = NSPrintJobDispositionValue(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPrintTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PrintTime = NSPrintInfoAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPrintTopMargin"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PrintTopMargin = NSPrintInfoAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPrintVerticalPagination"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PrintVerticalPagination = NSPrintInfoAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPrintVerticallyCentered"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PrintVerticallyCentered = NSPrintInfoAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPrintingCommunicationException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PrintingCommunicationException = foundation.NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSRTFDTextDocumentType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				RTFDTextDocumentType = NSAttributedStringDocumentType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSRTFPropertyStackOverflowException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				RTFPropertyStackOverflowException = foundation.NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSRTFTextDocumentType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				RTFTextDocumentType = NSAttributedStringDocumentType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSRaisesForNotApplicableKeysBindingOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				RaisesForNotApplicableKeysBindingOption = NSBindingOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSReadOnlyDocumentAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ReadOnlyDocumentAttribute = NSAttributedStringDocumentAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSRecentSearchesBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				RecentSearchesBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSRepresentedFilenameBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				RepresentedFilenameBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSRightMarginDocumentAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				RightMarginDocumentAttribute = NSAttributedStringDocumentAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSRowHeightBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				RowHeightBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSRuleEditorPredicateComparisonModifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				RuleEditorPredicateComparisonModifier = NSRuleEditorPredicatePartKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSRuleEditorPredicateCompoundType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				RuleEditorPredicateCompoundType = NSRuleEditorPredicatePartKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSRuleEditorPredicateCustomSelector"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				RuleEditorPredicateCustomSelector = NSRuleEditorPredicatePartKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSRuleEditorPredicateLeftExpression"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				RuleEditorPredicateLeftExpression = NSRuleEditorPredicatePartKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSRuleEditorPredicateOperatorType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				RuleEditorPredicateOperatorType = NSRuleEditorPredicatePartKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSRuleEditorPredicateOptions"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				RuleEditorPredicateOptions = NSRuleEditorPredicatePartKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSRuleEditorPredicateRightExpression"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				RuleEditorPredicateRightExpression = NSRuleEditorPredicatePartKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSRuleEditorRowsDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				RuleEditorRowsDidChangeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSRulerViewUnitCentimeters"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				RulerViewUnitCentimeters = NSRulerViewUnitName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSRulerViewUnitInches"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				RulerViewUnitInches = NSRulerViewUnitName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSRulerViewUnitPicas"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				RulerViewUnitPicas = NSRulerViewUnitName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSRulerViewUnitPoints"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				RulerViewUnitPoints = NSRulerViewUnitName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSScreenColorSpaceDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ScreenColorSpaceDidChangeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSScrollViewDidEndLiveMagnifyNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ScrollViewDidEndLiveMagnifyNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSScrollViewDidEndLiveScrollNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ScrollViewDidEndLiveScrollNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSScrollViewDidLiveScrollNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ScrollViewDidLiveScrollNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSScrollViewWillStartLiveMagnifyNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ScrollViewWillStartLiveMagnifyNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSScrollViewWillStartLiveScrollNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ScrollViewWillStartLiveScrollNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSelectedIdentifierBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SelectedIdentifierBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSelectedIndexBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SelectedIndexBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSelectedLabelBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SelectedLabelBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSelectedObjectBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SelectedObjectBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSelectedObjectsBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SelectedObjectsBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSelectedTagBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SelectedTagBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSelectedValueBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SelectedValueBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSelectedValuesBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SelectedValuesBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSelectionIndexPathsBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SelectionIndexPathsBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSelectionIndexesBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SelectionIndexesBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSelectorNameBindingOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SelectorNameBindingOption = NSBindingOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSelectsAllWhenSettingContentBindingOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SelectsAllWhenSettingContentBindingOption = NSBindingOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSShadowAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ShadowAttributeName = foundation.NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSharingServiceNameAddToAperture"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSSharingServiceNames.AddToAperture = NSSharingServiceName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSharingServiceNameAddToIPhoto"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSSharingServiceNames.AddToIPhoto = NSSharingServiceName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSharingServiceNameAddToSafariReadingList"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSSharingServiceNames.AddToSafariReadingList = NSSharingServiceName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSharingServiceNameCloudSharing"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSSharingServiceNames.CloudSharing = NSSharingServiceName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSharingServiceNameComposeEmail"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSSharingServiceNames.ComposeEmail = NSSharingServiceName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSharingServiceNameComposeMessage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSSharingServiceNames.ComposeMessage = NSSharingServiceName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSharingServiceNameSendViaAirDrop"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSSharingServiceNames.SendViaAirDrop = NSSharingServiceName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSharingServiceNameUseAsDesktopPicture"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSSharingServiceNames.UseAsDesktopPicture = NSSharingServiceName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSliderAccessoryWidthDefault"); err == nil && ptr != 0 {
		NSSliderAccessoryWidths.Default = *(*NSSliderAccessoryWidth)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSliderAccessoryWidthWide"); err == nil && ptr != 0 {
		NSSliderAccessoryWidths.Wide = *(*NSSliderAccessoryWidth)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSortDescriptorsBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SortDescriptorsBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSoundPboardType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SoundPboardType = NSPasteboardType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSourceTextScalingDocumentAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SourceTextScalingDocumentAttribute = NSAttributedStringDocumentAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSourceTextScalingDocumentOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SourceTextScalingDocumentOption = NSAttributedStringDocumentReadingOptionKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSpeechCharacterModeProperty"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SpeechCharacterModeProperty = NSSpeechPropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSpeechCommandDelimiterProperty"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SpeechCommandDelimiterProperty = NSSpeechPropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSpeechCommandPrefix"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SpeechCommandPrefix = NSSpeechCommandDelimiterKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSpeechCommandSuffix"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SpeechCommandSuffix = NSSpeechCommandDelimiterKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSpeechCurrentVoiceProperty"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SpeechCurrentVoiceProperty = NSSpeechPropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSpeechDictionaryAbbreviations"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SpeechDictionaryAbbreviations = NSSpeechDictionaryKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSpeechDictionaryEntryPhonemes"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SpeechDictionaryEntryPhonemes = NSSpeechDictionaryKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSpeechDictionaryEntrySpelling"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SpeechDictionaryEntrySpelling = NSSpeechDictionaryKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSpeechDictionaryLocaleIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SpeechDictionaryLocaleIdentifier = NSSpeechDictionaryKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSpeechDictionaryModificationDate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SpeechDictionaryModificationDate = NSSpeechDictionaryKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSpeechDictionaryPronunciations"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SpeechDictionaryPronunciations = NSSpeechDictionaryKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSpeechErrorCount"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SpeechErrorCount = NSSpeechErrorKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSpeechErrorNewestCharacterOffset"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SpeechErrorNewestCharacterOffset = NSSpeechErrorKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSpeechErrorNewestCode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SpeechErrorNewestCode = NSSpeechErrorKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSpeechErrorOldestCharacterOffset"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SpeechErrorOldestCharacterOffset = NSSpeechErrorKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSpeechErrorOldestCode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SpeechErrorOldestCode = NSSpeechErrorKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSpeechErrorsProperty"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SpeechErrorsProperty = NSSpeechPropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSpeechInputModeProperty"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SpeechInputModeProperty = NSSpeechPropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSpeechModeLiteral"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSSpeechModes.Literal = NSSpeechMode(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSpeechModeNormal"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSSpeechModes.Normal = NSSpeechMode(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSpeechModePhoneme"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSSpeechModes.Phoneme = NSSpeechMode(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSpeechModeText"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSSpeechModes.Text = NSSpeechMode(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSpeechNumberModeProperty"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SpeechNumberModeProperty = NSSpeechPropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSpeechOutputToFileURLProperty"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SpeechOutputToFileURLProperty = NSSpeechPropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSpeechPhonemeInfoExample"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SpeechPhonemeInfoExample = NSSpeechPhonemeInfoKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSpeechPhonemeInfoHiliteEnd"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SpeechPhonemeInfoHiliteEnd = NSSpeechPhonemeInfoKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSpeechPhonemeInfoHiliteStart"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SpeechPhonemeInfoHiliteStart = NSSpeechPhonemeInfoKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSpeechPhonemeInfoOpcode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SpeechPhonemeInfoOpcode = NSSpeechPhonemeInfoKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSpeechPhonemeInfoSymbol"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SpeechPhonemeInfoSymbol = NSSpeechPhonemeInfoKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSpeechPhonemeSymbolsProperty"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SpeechPhonemeSymbolsProperty = NSSpeechPropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSpeechPitchBaseProperty"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SpeechPitchBaseProperty = NSSpeechPropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSpeechPitchModProperty"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SpeechPitchModProperty = NSSpeechPropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSpeechRateProperty"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SpeechRateProperty = NSSpeechPropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSpeechRecentSyncProperty"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SpeechRecentSyncProperty = NSSpeechPropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSpeechResetProperty"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SpeechResetProperty = NSSpeechPropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSpeechStatusNumberOfCharactersLeft"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SpeechStatusNumberOfCharactersLeft = NSSpeechStatusKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSpeechStatusOutputBusy"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SpeechStatusOutputBusy = NSSpeechStatusKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSpeechStatusOutputPaused"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SpeechStatusOutputPaused = NSSpeechStatusKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSpeechStatusPhonemeCode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SpeechStatusPhonemeCode = NSSpeechStatusKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSpeechStatusProperty"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SpeechStatusProperty = NSSpeechPropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSpeechSynthesizerInfoIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SpeechSynthesizerInfoIdentifier = NSSpeechSynthesizerInfoKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSpeechSynthesizerInfoProperty"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SpeechSynthesizerInfoProperty = NSSpeechPropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSpeechSynthesizerInfoVersion"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SpeechSynthesizerInfoVersion = NSSpeechSynthesizerInfoKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSpeechVolumeProperty"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SpeechVolumeProperty = NSSpeechPropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSpellCheckerDidChangeAutomaticCapitalizationNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SpellCheckerDidChangeAutomaticCapitalizationNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSpellCheckerDidChangeAutomaticDashSubstitutionNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SpellCheckerDidChangeAutomaticDashSubstitutionNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSpellCheckerDidChangeAutomaticInlinePredictionNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SpellCheckerDidChangeAutomaticInlinePredictionNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSpellCheckerDidChangeAutomaticPeriodSubstitutionNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SpellCheckerDidChangeAutomaticPeriodSubstitutionNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSpellCheckerDidChangeAutomaticQuoteSubstitutionNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SpellCheckerDidChangeAutomaticQuoteSubstitutionNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSpellCheckerDidChangeAutomaticSpellingCorrectionNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SpellCheckerDidChangeAutomaticSpellingCorrectionNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSpellCheckerDidChangeAutomaticTextCompletionNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SpellCheckerDidChangeAutomaticTextCompletionNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSpellCheckerDidChangeAutomaticTextReplacementNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SpellCheckerDidChangeAutomaticTextReplacementNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSpellingStateAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SpellingStateAttributeName = foundation.NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSplitViewControllerAutomaticDimension"); err == nil && ptr != 0 {
		SplitViewControllerAutomaticDimension = *(*float64)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSplitViewDidResizeSubviewsNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SplitViewDidResizeSubviewsNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSplitViewItemUnspecifiedDimension"); err == nil && ptr != 0 {
		SplitViewItemUnspecifiedDimension = *(*float64)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSplitViewWillResizeSubviewsNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SplitViewWillResizeSubviewsNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSStrikethroughColorAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				StrikethroughColorAttributeName = foundation.NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSStrikethroughStyleAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				StrikethroughStyleAttributeName = foundation.NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSStrokeColorAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				StrokeColorAttributeName = foundation.NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSStrokeWidthAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				StrokeWidthAttributeName = foundation.NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSubjectDocumentAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SubjectDocumentAttribute = NSAttributedStringDocumentAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSuperscriptAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SuperscriptAttributeName = foundation.NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSystemColorsDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SystemColorsDidChangeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTIFFException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TIFFException = foundation.NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTabColumnTerminatorsAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TabColumnTerminatorsAttributeName = NSTextTabOptionKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTableViewColumnDidMoveNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TableViewColumnDidMoveNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTableViewColumnDidResizeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TableViewColumnDidResizeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTableViewRowViewKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TableViewRowViewKey = NSUserInterfaceItemIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTableViewSelectionDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TableViewSelectionDidChangeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTableViewSelectionIsChangingNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TableViewSelectionIsChangingNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTargetBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TargetBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTargetTextScalingDocumentOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TargetTextScalingDocumentOption = NSAttributedStringDocumentReadingOptionKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextAlternativesAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextAlternativesAttributeName = foundation.NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextAlternativesSelectedAlternativeStringNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextAlternativesSelectedAlternativeStringNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextCheckingDocumentAuthorKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextCheckingDocumentAuthorKey = NSTextCheckingOptionKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextCheckingDocumentTitleKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextCheckingDocumentTitleKey = NSTextCheckingOptionKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextCheckingDocumentURLKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextCheckingDocumentURLKey = NSTextCheckingOptionKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextCheckingGenerateInlinePredictionsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextCheckingGenerateInlinePredictionsKey = NSTextCheckingOptionKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextCheckingOrthographyKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextCheckingOrthographyKey = NSTextCheckingOptionKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextCheckingQuotesKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextCheckingQuotesKey = NSTextCheckingOptionKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextCheckingReferenceDateKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextCheckingReferenceDateKey = NSTextCheckingOptionKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextCheckingReferenceTimeZoneKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextCheckingReferenceTimeZoneKey = NSTextCheckingOptionKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextCheckingRegularExpressionsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextCheckingRegularExpressionsKey = NSTextCheckingOptionKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextCheckingReplacementsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextCheckingReplacementsKey = NSTextCheckingOptionKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextCheckingSelectedRangeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextCheckingSelectedRangeKey = NSTextCheckingOptionKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextColorBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextColorBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextContentStorageUnsupportedAttributeAddedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextContentStorageUnsupportedAttributeAddedNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextContentTypeAddressCity"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTextContentTypes.AddressCity = NSTextContentType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextContentTypeAddressCityAndState"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTextContentTypes.AddressCityAndState = NSTextContentType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextContentTypeAddressState"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTextContentTypes.AddressState = NSTextContentType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextContentTypeBirthdate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTextContentTypes.Birthdate = NSTextContentType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextContentTypeBirthdateDay"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTextContentTypes.BirthdateDay = NSTextContentType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextContentTypeBirthdateMonth"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTextContentTypes.BirthdateMonth = NSTextContentType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextContentTypeBirthdateYear"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTextContentTypes.BirthdateYear = NSTextContentType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextContentTypeCountryName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTextContentTypes.CountryName = NSTextContentType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextContentTypeCreditCardExpiration"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTextContentTypes.CreditCardExpiration = NSTextContentType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextContentTypeCreditCardExpirationMonth"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTextContentTypes.CreditCardExpirationMonth = NSTextContentType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextContentTypeCreditCardExpirationYear"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTextContentTypes.CreditCardExpirationYear = NSTextContentType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextContentTypeCreditCardFamilyName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTextContentTypes.CreditCardFamilyName = NSTextContentType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextContentTypeCreditCardGivenName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTextContentTypes.CreditCardGivenName = NSTextContentType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextContentTypeCreditCardMiddleName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTextContentTypes.CreditCardMiddleName = NSTextContentType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextContentTypeCreditCardName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTextContentTypes.CreditCardName = NSTextContentType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextContentTypeCreditCardNumber"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTextContentTypes.CreditCardNumber = NSTextContentType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextContentTypeCreditCardSecurityCode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTextContentTypes.CreditCardSecurityCode = NSTextContentType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextContentTypeCreditCardType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTextContentTypes.CreditCardType = NSTextContentType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextContentTypeDateTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTextContentTypes.DateTime = NSTextContentType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextContentTypeEmailAddress"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTextContentTypes.EmailAddress = NSTextContentType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextContentTypeFamilyName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTextContentTypes.FamilyName = NSTextContentType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextContentTypeFlightNumber"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTextContentTypes.FlightNumber = NSTextContentType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextContentTypeFullStreetAddress"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTextContentTypes.FullStreetAddress = NSTextContentType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextContentTypeGivenName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTextContentTypes.GivenName = NSTextContentType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextContentTypeJobTitle"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTextContentTypes.JobTitle = NSTextContentType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextContentTypeLocation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTextContentTypes.Location = NSTextContentType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextContentTypeMiddleName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTextContentTypes.MiddleName = NSTextContentType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextContentTypeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTextContentTypes.Name = NSTextContentType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextContentTypeNamePrefix"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTextContentTypes.NamePrefix = NSTextContentType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextContentTypeNameSuffix"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTextContentTypes.NameSuffix = NSTextContentType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextContentTypeNewPassword"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTextContentTypes.NewPassword = NSTextContentType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextContentTypeNickname"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTextContentTypes.Nickname = NSTextContentType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextContentTypeOneTimeCode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTextContentTypes.OneTimeCode = NSTextContentType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextContentTypeOrganizationName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTextContentTypes.OrganizationName = NSTextContentType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextContentTypePassword"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTextContentTypes.Password = NSTextContentType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextContentTypePostalCode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTextContentTypes.PostalCode = NSTextContentType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextContentTypeShipmentTrackingNumber"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTextContentTypes.ShipmentTrackingNumber = NSTextContentType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextContentTypeStreetAddressLine1"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTextContentTypes.StreetAddressLine1 = NSTextContentType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextContentTypeStreetAddressLine2"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTextContentTypes.StreetAddressLine2 = NSTextContentType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextContentTypeSublocality"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTextContentTypes.Sublocality = NSTextContentType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextContentTypeTelephoneNumber"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTextContentTypes.TelephoneNumber = NSTextContentType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextContentTypeURL"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTextContentTypes.URL = NSTextContentType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextContentTypeUsername"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTextContentTypes.Username = NSTextContentType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextDidBeginEditingNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextDidBeginEditingNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextDidChangeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextDidEndEditingNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextDidEndEditingNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextEffectAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextEffectAttributeName = foundation.NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextEffectLetterpressStyle"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextEffectLetterpressStyle = NSTextEffectStyle(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextEncodingNameDocumentAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextEncodingNameDocumentAttribute = NSAttributedStringDocumentAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextEncodingNameDocumentOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextEncodingNameDocumentOption = NSAttributedStringDocumentReadingOptionKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextFinderCaseInsensitiveKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextFinderCaseInsensitiveKey = NSPasteboardTypeTextFinderOptionKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextFinderMatchingTypeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextFinderMatchingTypeKey = NSPasteboardTypeTextFinderOptionKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextHighlightColorSchemeAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextHighlightColorSchemeAttributeName = foundation.NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextHighlightColorSchemeBlue"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTextHighlightColorSchemes.Blue = NSTextHighlightColorScheme(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextHighlightColorSchemeDefault"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTextHighlightColorSchemes.Default = NSTextHighlightColorScheme(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextHighlightColorSchemeMint"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTextHighlightColorSchemes.Mint = NSTextHighlightColorScheme(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextHighlightColorSchemeOrange"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTextHighlightColorSchemes.Orange = NSTextHighlightColorScheme(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextHighlightColorSchemePink"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTextHighlightColorSchemes.Pink = NSTextHighlightColorScheme(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextHighlightColorSchemePurple"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTextHighlightColorSchemes.Purple = NSTextHighlightColorScheme(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextHighlightStyleAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextHighlightStyleAttributeName = foundation.NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextHighlightStyleDefault"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextHighlightStyleDefault = NSTextHighlightStyle(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextInputContextKeyboardSelectionDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextInputContextKeyboardSelectionDidChangeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextKit1ListMarkerFormatDocumentOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextKit1ListMarkerFormatDocumentOption = NSAttributedStringDocumentReadingOptionKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextLayoutSectionOrientation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextLayoutSectionOrientation = NSTextLayoutSectionKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextLayoutSectionRange"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextLayoutSectionRange = NSTextLayoutSectionKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextLayoutSectionsAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextLayoutSectionsAttribute = NSAttributedStringDocumentAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextLineTooLongException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextLineTooLongException = foundation.NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextListMarkerBox"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextListMarkerBox = NSTextListMarkerFormat(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextListMarkerCheck"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextListMarkerCheck = NSTextListMarkerFormat(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextListMarkerCircle"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextListMarkerCircle = NSTextListMarkerFormat(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextListMarkerDecimal"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextListMarkerDecimal = NSTextListMarkerFormat(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextListMarkerDiamond"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextListMarkerDiamond = NSTextListMarkerFormat(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextListMarkerDisc"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextListMarkerDisc = NSTextListMarkerFormat(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextListMarkerHyphen"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextListMarkerHyphen = NSTextListMarkerFormat(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextListMarkerLowercaseAlpha"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextListMarkerLowercaseAlpha = NSTextListMarkerFormat(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextListMarkerLowercaseHexadecimal"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextListMarkerLowercaseHexadecimal = NSTextListMarkerFormat(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextListMarkerLowercaseLatin"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextListMarkerLowercaseLatin = NSTextListMarkerFormat(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextListMarkerLowercaseRoman"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextListMarkerLowercaseRoman = NSTextListMarkerFormat(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextListMarkerOctal"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextListMarkerOctal = NSTextListMarkerFormat(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextListMarkerSquare"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextListMarkerSquare = NSTextListMarkerFormat(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextListMarkerUppercaseAlpha"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextListMarkerUppercaseAlpha = NSTextListMarkerFormat(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextListMarkerUppercaseHexadecimal"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextListMarkerUppercaseHexadecimal = NSTextListMarkerFormat(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextListMarkerUppercaseLatin"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextListMarkerUppercaseLatin = NSTextListMarkerFormat(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextListMarkerUppercaseRoman"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextListMarkerUppercaseRoman = NSTextListMarkerFormat(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextMovementUserInfoKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextMovementUserInfoKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextNoSelectionException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextNoSelectionException = foundation.NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextReadException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextReadException = foundation.NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextScalingDocumentAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextScalingDocumentAttribute = NSAttributedStringDocumentAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextSizeMultiplierDocumentOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextSizeMultiplierDocumentOption = NSAttributedStringDocumentReadingOptionKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextStorageDidProcessEditingNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextStorageDidProcessEditingNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextStorageWillProcessEditingNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextStorageWillProcessEditingNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextViewDidChangeSelectionNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextViewDidChangeSelectionNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextViewDidChangeTypingAttributesNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextViewDidChangeTypingAttributesNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextViewDidSwitchToNSLayoutManagerNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextViewDidSwitchToNSLayoutManagerNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextViewWillChangeNotifyingTextViewNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextViewWillChangeNotifyingTextViewNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextViewWillSwitchToNSLayoutManagerNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextViewWillSwitchToNSLayoutManagerNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextWriteException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextWriteException = foundation.NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTimeoutDocumentOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TimeoutDocumentOption = NSAttributedStringDocumentReadingOptionKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTitleBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TitleBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTitleDocumentAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TitleDocumentAttribute = NSAttributedStringDocumentAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSToolTipAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ToolTipAttributeName = foundation.NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSToolTipBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ToolTipBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSToolbarCloudSharingItemIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ToolbarCloudSharingItemIdentifier = NSToolbarItemIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSToolbarDidRemoveItemNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ToolbarDidRemoveItemNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSToolbarFlexibleSpaceItemIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ToolbarFlexibleSpaceItemIdentifier = NSToolbarItemIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSToolbarInspectorTrackingSeparatorItemIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ToolbarInspectorTrackingSeparatorItemIdentifier = NSToolbarItemIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSToolbarItemKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ToolbarItemKey = NSToolbarUserInfoKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSToolbarNewIndexKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ToolbarNewIndexKey = NSToolbarUserInfoKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSToolbarPrintItemIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ToolbarPrintItemIdentifier = NSToolbarItemIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSToolbarShowColorsItemIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ToolbarShowColorsItemIdentifier = NSToolbarItemIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSToolbarShowFontsItemIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ToolbarShowFontsItemIdentifier = NSToolbarItemIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSToolbarSidebarTrackingSeparatorItemIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ToolbarSidebarTrackingSeparatorItemIdentifier = NSToolbarItemIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSToolbarSpaceItemIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ToolbarSpaceItemIdentifier = NSToolbarItemIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSToolbarToggleInspectorItemIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ToolbarToggleInspectorItemIdentifier = NSToolbarItemIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSToolbarToggleSidebarItemIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ToolbarToggleSidebarItemIdentifier = NSToolbarItemIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSToolbarWillAddItemNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ToolbarWillAddItemNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSToolbarWritingToolsItemIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ToolbarWritingToolsItemIdentifier = NSToolbarItemIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTopMarginDocumentAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TopMarginDocumentAttribute = NSAttributedStringDocumentAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTouchBarItemIdentifierCandidateList"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTouchBarItemIdentifiers.CandidateList = NSTouchBarItemIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTouchBarItemIdentifierCharacterPicker"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTouchBarItemIdentifiers.CharacterPicker = NSTouchBarItemIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTouchBarItemIdentifierFixedSpaceLarge"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTouchBarItemIdentifiers.FixedSpaceLarge = NSTouchBarItemIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTouchBarItemIdentifierFixedSpaceSmall"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTouchBarItemIdentifiers.FixedSpaceSmall = NSTouchBarItemIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTouchBarItemIdentifierFlexibleSpace"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTouchBarItemIdentifiers.FlexibleSpace = NSTouchBarItemIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTouchBarItemIdentifierOtherItemsProxy"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTouchBarItemIdentifiers.OtherItemsProxy = NSTouchBarItemIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTouchBarItemIdentifierTextAlignment"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTouchBarItemIdentifiers.TextAlignment = NSTouchBarItemIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTouchBarItemIdentifierTextColorPicker"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTouchBarItemIdentifiers.TextColorPicker = NSTouchBarItemIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTouchBarItemIdentifierTextFormat"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTouchBarItemIdentifiers.TextFormat = NSTouchBarItemIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTouchBarItemIdentifierTextList"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTouchBarItemIdentifiers.TextList = NSTouchBarItemIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTouchBarItemIdentifierTextStyle"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSTouchBarItemIdentifiers.TextStyle = NSTouchBarItemIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTrackingAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TrackingAttributeName = foundation.NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTransparentBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TransparentBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTypeIdentifierAddressText"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TypeIdentifierAddressText = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTypeIdentifierDateText"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TypeIdentifierDateText = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTypeIdentifierPhoneNumberText"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TypeIdentifierPhoneNumberText = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTypeIdentifierTransitInformationText"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TypeIdentifierTransitInformationText = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTypedStreamVersionException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TypedStreamVersionException = foundation.NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSUnderlineColorAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				UnderlineColorAttributeName = foundation.NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSUnderlineStyleAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				UnderlineStyleAttributeName = foundation.NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSUserActivityDocumentURLKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				UserActivityDocumentURLKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSValidatesImmediatelyBindingOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ValidatesImmediatelyBindingOption = NSBindingOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSValueBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ValueBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSValuePathBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ValuePathBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSValueTransformerBindingOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ValueTransformerBindingOption = NSBindingOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSValueTransformerNameBindingOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ValueTransformerNameBindingOption = NSBindingOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSValueURLBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ValueURLBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSVerticalGlyphFormAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				VerticalGlyphFormAttributeName = foundation.NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSViewAnimationEffectKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ViewAnimationEffectKey = NSViewAnimationKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSViewAnimationEndFrameKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ViewAnimationEndFrameKey = NSViewAnimationKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSViewAnimationFadeInEffect"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ViewAnimationFadeInEffect = NSViewAnimationEffectName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSViewAnimationFadeOutEffect"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ViewAnimationFadeOutEffect = NSViewAnimationEffectName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSViewAnimationStartFrameKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ViewAnimationStartFrameKey = NSViewAnimationKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSViewAnimationTargetKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ViewAnimationTargetKey = NSViewAnimationKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSViewBoundsDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ViewBoundsDidChangeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSViewDidUpdateTrackingAreasNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ViewDidUpdateTrackingAreasNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSViewFrameDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ViewFrameDidChangeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSViewModeDocumentAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ViewModeDocumentAttribute = NSAttributedStringDocumentAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSViewNoIntrinsicMetric"); err == nil && ptr != 0 {
		ViewNoIntrinsicMetric = *(*float64)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSViewSizeDocumentAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ViewSizeDocumentAttribute = NSAttributedStringDocumentAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSViewZoomDocumentAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ViewZoomDocumentAttribute = NSAttributedStringDocumentAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSVisibleBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				VisibleBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSVoiceAge"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				VoiceAge = NSVoiceAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSVoiceDemoText"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				VoiceDemoText = NSVoiceAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSVoiceGender"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				VoiceGender = NSVoiceAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSVoiceGenderFemale"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				VoiceGenderFemale = NSVoiceGenderName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSVoiceGenderMale"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				VoiceGenderMale = NSVoiceGenderName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSVoiceGenderNeuter"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				VoiceGenderNeuter = NSVoiceGenderName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSVoiceGenderNeutral"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				VoiceGenderNeutral = NSVoiceGenderName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSVoiceIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				VoiceIdentifier = NSVoiceAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSVoiceIndividuallySpokenCharacters"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				VoiceIndividuallySpokenCharacters = NSVoiceAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSVoiceLocaleIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				VoiceLocaleIdentifier = NSVoiceAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSVoiceName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				VoiceName = NSVoiceAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSVoiceSupportedCharacters"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				VoiceSupportedCharacters = NSVoiceAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWarningValueBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WarningValueBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWebArchiveTextDocumentType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WebArchiveTextDocumentType = NSAttributedStringDocumentType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWebPreferencesDocumentOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WebPreferencesDocumentOption = NSAttributedStringDocumentReadingOptionKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWebResourceLoadDelegateDocumentOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WebResourceLoadDelegateDocumentOption = NSAttributedStringDocumentReadingOptionKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWhite"); err == nil && ptr != 0 {
		White = *(*float64)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWidthBinding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WidthBinding = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWindowDidBecomeKeyNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WindowDidBecomeKeyNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWindowDidBecomeMainNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WindowDidBecomeMainNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWindowDidChangeBackingPropertiesNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WindowDidChangeBackingPropertiesNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWindowDidChangeOcclusionStateNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WindowDidChangeOcclusionStateNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWindowDidChangeScreenNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WindowDidChangeScreenNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWindowDidChangeScreenProfileNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WindowDidChangeScreenProfileNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWindowDidDeminiaturizeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WindowDidDeminiaturizeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWindowDidEndLiveResizeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WindowDidEndLiveResizeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWindowDidEndSheetNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WindowDidEndSheetNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWindowDidEnterFullScreenNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WindowDidEnterFullScreenNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWindowDidEnterVersionBrowserNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WindowDidEnterVersionBrowserNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWindowDidExitFullScreenNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WindowDidExitFullScreenNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWindowDidExitVersionBrowserNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WindowDidExitVersionBrowserNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWindowDidExposeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WindowDidExposeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWindowDidMiniaturizeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WindowDidMiniaturizeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWindowDidMoveNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WindowDidMoveNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWindowDidResignKeyNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WindowDidResignKeyNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWindowDidResignMainNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WindowDidResignMainNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWindowDidResizeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WindowDidResizeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWindowDidUpdateNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WindowDidUpdateNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWindowServerCommunicationException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WindowServerCommunicationException = foundation.NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWindowWillBeginSheetNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WindowWillBeginSheetNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWindowWillCloseNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WindowWillCloseNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWindowWillEnterFullScreenNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WindowWillEnterFullScreenNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWindowWillEnterVersionBrowserNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WindowWillEnterVersionBrowserNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWindowWillExitFullScreenNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WindowWillExitFullScreenNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWindowWillExitVersionBrowserNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WindowWillExitVersionBrowserNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWindowWillMiniaturizeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WindowWillMiniaturizeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWindowWillMoveNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WindowWillMoveNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWindowWillStartLiveResizeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WindowWillStartLiveResizeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWordMLTextDocumentType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WordMLTextDocumentType = NSAttributedStringDocumentType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWordTablesReadException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WordTablesReadException = foundation.NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWordTablesWriteException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WordTablesWriteException = foundation.NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWorkspaceAccessibilityDisplayOptionsDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WorkspaceAccessibilityDisplayOptionsDidChangeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWorkspaceActiveSpaceDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WorkspaceActiveSpaceDidChangeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWorkspaceApplicationKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WorkspaceApplicationKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWorkspaceDesktopImageAllowClippingKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WorkspaceDesktopImageAllowClippingKey = NSWorkspaceDesktopImageOptionKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWorkspaceDesktopImageFillColorKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WorkspaceDesktopImageFillColorKey = NSWorkspaceDesktopImageOptionKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWorkspaceDesktopImageScalingKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WorkspaceDesktopImageScalingKey = NSWorkspaceDesktopImageOptionKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWorkspaceDidActivateApplicationNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WorkspaceDidActivateApplicationNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWorkspaceDidChangeFileLabelsNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WorkspaceDidChangeFileLabelsNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWorkspaceDidDeactivateApplicationNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WorkspaceDidDeactivateApplicationNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWorkspaceDidHideApplicationNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WorkspaceDidHideApplicationNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWorkspaceDidLaunchApplicationNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WorkspaceDidLaunchApplicationNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWorkspaceDidMountNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WorkspaceDidMountNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWorkspaceDidRenameVolumeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WorkspaceDidRenameVolumeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWorkspaceDidTerminateApplicationNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WorkspaceDidTerminateApplicationNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWorkspaceDidUnhideApplicationNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WorkspaceDidUnhideApplicationNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWorkspaceDidUnmountNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WorkspaceDidUnmountNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWorkspaceDidWakeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WorkspaceDidWakeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWorkspaceScreensDidSleepNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WorkspaceScreensDidSleepNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWorkspaceScreensDidWakeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WorkspaceScreensDidWakeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWorkspaceSessionDidBecomeActiveNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WorkspaceSessionDidBecomeActiveNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWorkspaceSessionDidResignActiveNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WorkspaceSessionDidResignActiveNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWorkspaceVolumeLocalizedNameKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WorkspaceVolumeLocalizedNameKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWorkspaceVolumeOldLocalizedNameKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WorkspaceVolumeOldLocalizedNameKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWorkspaceVolumeOldURLKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WorkspaceVolumeOldURLKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWorkspaceVolumeURLKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WorkspaceVolumeURLKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWorkspaceWillLaunchApplicationNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WorkspaceWillLaunchApplicationNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWorkspaceWillPowerOffNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WorkspaceWillPowerOffNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWorkspaceWillSleepNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WorkspaceWillSleepNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWorkspaceWillUnmountNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WorkspaceWillUnmountNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWritingDirectionAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WritingDirectionAttributeName = foundation.NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWritingToolsExclusionAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WritingToolsExclusionAttributeName = foundation.NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

}

// NSAppKitVersions provides typed accessors for [NSAppKitVersion] constants.
var NSAppKitVersions struct {
	// Number: The most recent version of AppKit.
	Number NSAppKitVersion
	// Number10_0: The AppKit framework included in OS X v10.0.
	Number10_0 NSAppKitVersion
	// Number10_1: The AppKit framework included in OS X v10.1.
	Number10_1      NSAppKitVersion
	Number10_10     NSAppKitVersion
	Number10_10_2   NSAppKitVersion
	Number10_10_3   NSAppKitVersion
	Number10_10_4   NSAppKitVersion
	Number10_10_5   NSAppKitVersion
	Number10_10_Max NSAppKitVersion
	Number10_11     NSAppKitVersion
	Number10_11_1   NSAppKitVersion
	Number10_11_2   NSAppKitVersion
	Number10_11_3   NSAppKitVersion
	Number10_12     NSAppKitVersion
	Number10_12_1   NSAppKitVersion
	Number10_12_2   NSAppKitVersion
	Number10_13     NSAppKitVersion
	Number10_13_1   NSAppKitVersion
	Number10_13_2   NSAppKitVersion
	Number10_13_4   NSAppKitVersion
	Number10_14     NSAppKitVersion
	Number10_14_1   NSAppKitVersion
	Number10_14_2   NSAppKitVersion
	Number10_14_3   NSAppKitVersion
	Number10_14_4   NSAppKitVersion
	Number10_14_5   NSAppKitVersion
	Number10_15     NSAppKitVersion
	Number10_15_1   NSAppKitVersion
	Number10_15_2   NSAppKitVersion
	Number10_15_3   NSAppKitVersion
	Number10_15_4   NSAppKitVersion
	Number10_15_5   NSAppKitVersion
	Number10_15_6   NSAppKitVersion
	// Number10_2: The AppKit framework included in OS X v10.2.
	Number10_2 NSAppKitVersion
	// Number10_2_3: The AppKit framework included in OS X v10.2.3.
	Number10_2_3 NSAppKitVersion
	// Number10_3: The AppKit framework included in OS X v10.3.
	Number10_3 NSAppKitVersion
	// Number10_3_2: The AppKit framework included in OS X v10.3.2.
	Number10_3_2 NSAppKitVersion
	// Number10_3_3: The AppKit framework included in OS X v10.3.3.
	Number10_3_3 NSAppKitVersion
	// Number10_3_5: The AppKit framework included in OS X v10.3.5.
	Number10_3_5 NSAppKitVersion
	// Number10_3_7: The AppKit framework included in OS X v10.3.7.
	Number10_3_7 NSAppKitVersion
	// Number10_3_9: The AppKit framework included in OS X v10.3.9.
	Number10_3_9 NSAppKitVersion
	// Number10_4: The AppKit framework included in OS X v10.4.
	Number10_4 NSAppKitVersion
	// Number10_4_1: The AppKit framework included in OS X v10.4.1.
	Number10_4_1 NSAppKitVersion
	// Number10_4_3: The AppKit framework included in OS X v10.4.3.
	Number10_4_3 NSAppKitVersion
	// Number10_4_4: The AppKit framework included in OS X v10.4.4.
	Number10_4_4 NSAppKitVersion
	// Number10_4_7: The AppKit framework included in OS X v10.4.7.
	Number10_4_7 NSAppKitVersion
	// Number10_5: The AppKit framework included in OS X v10.5.
	Number10_5 NSAppKitVersion
	// Number10_5_2: The AppKit framework included in OS X v10.5.2.
	Number10_5_2 NSAppKitVersion
	// Number10_5_3: The AppKit framework included in OS X v10.5.3.
	Number10_5_3 NSAppKitVersion
	// Number10_6: The AppKit framework included in OS X v10.6.
	Number10_6 NSAppKitVersion
	// Number10_7: The AppKit framework included in OS X v10.7.
	Number10_7 NSAppKitVersion
	// Number10_7_2: The AppKit framework included in OS X v10.7.2.
	Number10_7_2 NSAppKitVersion
	// Number10_7_3: The AppKit framework included in OS X v10.7.3.
	Number10_7_3 NSAppKitVersion
	// Number10_7_4: The AppKit framework included in OS X v10.7.4.
	Number10_7_4 NSAppKitVersion
	// Number10_8: The AppKit framework included in OS X v10.8.
	Number10_8 NSAppKitVersion
	Number10_9 NSAppKitVersion
	Number11_0 NSAppKitVersion
	Number11_1 NSAppKitVersion
	Number11_2 NSAppKitVersion
	Number11_3 NSAppKitVersion
	Number11_4 NSAppKitVersion
	Number11_5 NSAppKitVersion
	Number12_0 NSAppKitVersion
	Number12_1 NSAppKitVersion
	Number12_2 NSAppKitVersion
	Number12_3 NSAppKitVersion
	Number12_4 NSAppKitVersion
	Number12_5 NSAppKitVersion
	Number13_0 NSAppKitVersion
	Number13_1 NSAppKitVersion
	Number13_2 NSAppKitVersion
	Number13_3 NSAppKitVersion
	Number13_4 NSAppKitVersion
	Number13_5 NSAppKitVersion
	Number13_6 NSAppKitVersion
	Number14_0 NSAppKitVersion
	Number14_1 NSAppKitVersion
	// NumberWithColumnResizingBrowser: The specific version of the AppKit framework that introduced support for resizing individual browser columns.
	NumberWithColumnResizingBrowser NSAppKitVersion
	// NumberWithContinuousScrollingBrowser: The specific version of the AppKit framework that introduced support the continuous scrolling in a browser view.
	NumberWithContinuousScrollingBrowser NSAppKitVersion
	// NumberWithCursorSizeSupport: The specific version of the AppKit framework that introduced support for cursors larger than 16 x 16 pixels in size.
	NumberWithCursorSizeSupport NSAppKitVersion
	// NumberWithCustomSheetPosition: The specific version of the AppKit framework that introduced custom sheet positioning.
	NumberWithCustomSheetPosition NSAppKitVersion
	// NumberWithDeferredWindowDisplaySupport: The specific version of the AppKit framework that introduced support for deferred window display.
	NumberWithDeferredWindowDisplaySupport NSAppKitVersion
	// NumberWithDirectionalTabs: The specific version of the AppKit framework that introduced support for directional tab items.
	NumberWithDirectionalTabs NSAppKitVersion
	// NumberWithDockTilePlugInSupport: The specific version of the AppKit framework that introduced support for dock tile plug-ins.
	NumberWithDockTilePlugInSupport NSAppKitVersion
	// NumberWithPatternColorLeakFix: The specific version of the AppKit framework from OS X 10.1 that correctly autoreleases color objects.
	NumberWithPatternColorLeakFix NSAppKitVersion
}

// NSAppearanceNames provides typed accessors for [NSAppearanceName] constants.
var NSAppearanceNames struct {
	// AccessibilityHighContrastAqua: A high-contrast version of the standard light system appearance.
	AccessibilityHighContrastAqua NSAppearanceName
	// AccessibilityHighContrastDarkAqua: A high-contrast version of the standard dark system appearance.
	AccessibilityHighContrastDarkAqua NSAppearanceName
	// AccessibilityHighContrastVibrantDark: A high-contrast version of the dark vibrant appearance.
	AccessibilityHighContrastVibrantDark NSAppearanceName
	// AccessibilityHighContrastVibrantLight: A high-contrast version of the light vibrant appearance.
	AccessibilityHighContrastVibrantLight NSAppearanceName
	// Aqua: The standard light system appearance.
	Aqua NSAppearanceName
	// DarkAqua: The standard dark system appearance.
	DarkAqua NSAppearanceName
	// VibrantDark: A dark vibrant appearance, available only in visual effect views.
	VibrantDark NSAppearanceName
	// VibrantLight: The light vibrant appearance, available only in visual effect views.
	VibrantLight NSAppearanceName
}

// NSControlStateValues provides typed accessors for [NSControlStateValue] constants.
var NSControlStateValues struct {
	// Mixed: A constant value that indicates a control is in a mixed state, neither on nor off.
	Mixed NSControlStateValue
	// Off: A constant value that indicates a control is off or unselected.
	Off NSControlStateValue
	// On: A constant value that indicates a control is on or selected.
	On NSControlStateValue
}

// NSDefinitionPresentationTypes provides typed accessors for [NSDefinitionPresentationType] constants.
var NSDefinitionPresentationTypes struct {
	// DictionaryApplication: A possible value of the [presentationType](<doc://com.apple.appkit/documentation/AppKit/NSView/DefinitionOptionKey/presentationType>) dictionary key that invokes Dictionary application to display the definition.
	DictionaryApplication NSDefinitionPresentationType
	// Overlay: A possible value of the [presentationType](<doc://com.apple.appkit/documentation/AppKit/NSView/DefinitionOptionKey/presentationType>) dictionary key that produces a small overlay window at the string location,
	Overlay NSDefinitionPresentationType
}

// NSFontDescriptorSystemDesigns provides typed accessors for [NSFontDescriptorSystemDesign] constants.
var NSFontDescriptorSystemDesigns struct {
	// Default: The default font design.
	Default NSFontDescriptorSystemDesign
	// Monospaced: A font with a monospace appearance.
	Monospaced NSFontDescriptorSystemDesign
	// Rounded: A font with a rounded appearance.
	Rounded NSFontDescriptorSystemDesign
	// Serif: A font with a serif design.
	Serif NSFontDescriptorSystemDesign
}

// NSFontTextStyles provides typed accessors for [NSFontTextStyle] constants.
var NSFontTextStyles struct {
	// Body: The font you use for body text.
	Body NSFontTextStyle
	// Callout: The font you use for callouts.
	Callout NSFontTextStyle
	// Caption1: The font you use for standard captions.
	Caption1 NSFontTextStyle
	// Caption2: The font you use for alternate captions.
	Caption2 NSFontTextStyle
	// Footnote: The font you use in footnotes.
	Footnote NSFontTextStyle
	// Headline: The font you use for headings.
	Headline NSFontTextStyle
	// LargeTitle: The font you use for large titles.
	LargeTitle NSFontTextStyle
	// Subheadline: The font you use for subheadings.
	Subheadline NSFontTextStyle
	// Title1: The font you use for first-level hierarchical headings.
	Title1 NSFontTextStyle
	// Title2: The font you use for second-level hierarchical headings.
	Title2 NSFontTextStyle
	// Title3: The font you use for third-level hierarchical headings.
	Title3 NSFontTextStyle
}

// NSFontWeights provides typed accessors for [NSFontWeight] constants.
var NSFontWeights struct {
	// Black: The font weight for system black font.
	Black NSFontWeight
	// Bold: The font weight for system bold font.
	Bold NSFontWeight
	// Heavy: The font weight for system heavy font.
	Heavy NSFontWeight
	// Light: The font weight for system light font.
	Light NSFontWeight
	// Medium: The font weight for system medium font.
	Medium NSFontWeight
	// Regular: The font weight for system regular font.
	Regular NSFontWeight
	// Semibold: The font weight for system semibold font.
	Semibold NSFontWeight
	// Thin: The font weight for system thin font.
	Thin NSFontWeight
	// UltraLight: The font weight for system ultra light font.
	UltraLight NSFontWeight
}

// NSFontWidths provides typed accessors for [NSFontWidth] constants.
var NSFontWidths struct {
	Compressed NSFontWidth
	Condensed  NSFontWidth
	Expanded   NSFontWidth
	Standard   NSFontWidth
}

// NSImageNames provides typed accessors for [NSImageName] constants.
var NSImageNames struct {
	// ActionTemplate: An action menu template image.
	ActionTemplate NSImageName
	// AddTemplate: An add item template image.
	AddTemplate NSImageName
	// Advanced: Advanced preferences toolbar icon for the preferences window.
	Advanced NSImageName
	// ApplicationIcon: The app’s icon.
	ApplicationIcon NSImageName
	// BluetoothTemplate: A Bluetooth template image.
	BluetoothTemplate NSImageName
	// Bonjour: A Bonjour icon.
	Bonjour NSImageName
	// BookmarksTemplate: Bookmarks image suitable for a template.
	BookmarksTemplate NSImageName
	// Caution: A caution image.
	Caution NSImageName
	// ColorPanel: A color panel toolbar icon.
	ColorPanel NSImageName
	// ColumnViewTemplate: A column view mode template image.
	ColumnViewTemplate NSImageName
	// Computer: A computer icon.
	Computer NSImageName
	// EnterFullScreenTemplate: An enter full-screen mode template image.
	EnterFullScreenTemplate NSImageName
	// Everyone: Permissions for all users.
	Everyone NSImageName
	// ExitFullScreenTemplate: An exit full-screen mode template image.
	ExitFullScreenTemplate NSImageName
	// FlowViewTemplate: A cover flow view mode template image.
	FlowViewTemplate NSImageName
	// Folder: A folder image.
	Folder NSImageName
	// FolderBurnable: A burnable folder icon.
	FolderBurnable NSImageName
	// FolderSmart: A smart folder icon.
	FolderSmart NSImageName
	// FollowLinkFreestandingTemplate: A link template image.
	FollowLinkFreestandingTemplate NSImageName
	// FontPanel: A font panel toolbar icon.
	FontPanel NSImageName
	// GoBackTemplate: A “go back” template image.
	GoBackTemplate NSImageName
	// GoForwardTemplate: A “go forward” template image.
	GoForwardTemplate NSImageName
	// GoLeftTemplate: A “go back” template image.
	GoLeftTemplate NSImageName
	// GoRightTemplate: A “go forward” template image.
	GoRightTemplate NSImageName
	// HomeTemplate: Home image suitable for a template.
	HomeTemplate NSImageName
	// IChatTheaterTemplate: An iChat Theater template image.
	IChatTheaterTemplate NSImageName
	// IconViewTemplate: An icon view mode template image.
	IconViewTemplate NSImageName
	// Info: An information toolbar icon.
	Info NSImageName
	// InvalidDataFreestandingTemplate: A template image used to denote invalid data.
	InvalidDataFreestandingTemplate NSImageName
	// LeftFacingTriangleTemplate: A generic left-facing triangle template image.
	LeftFacingTriangleTemplate NSImageName
	// ListViewTemplate: A list view mode template image.
	ListViewTemplate NSImageName
	// LockLockedTemplate: A locked padlock template image.
	LockLockedTemplate NSImageName
	// LockUnlockedTemplate: An unlocked padlock template image.
	LockUnlockedTemplate NSImageName
	// MenuMixedStateTemplate: A horizontal dash, for use in menus.
	MenuMixedStateTemplate NSImageName
	// MenuOnStateTemplate: A check mark template image, for use in menus.
	MenuOnStateTemplate NSImageName
	// MobileMe: A MobileMe icon.
	MobileMe NSImageName
	// MultipleDocuments: A drag image for multiple items.
	MultipleDocuments NSImageName
	// Network: A network icon.
	Network NSImageName
	// PathTemplate: A path button template image.
	PathTemplate NSImageName
	// PreferencesGeneral: General preferences toolbar icon for the preferences window.
	PreferencesGeneral NSImageName
	// QuickLookTemplate: A Quick Look template image.
	QuickLookTemplate NSImageName
	// RefreshFreestandingTemplate: A refresh template image.
	RefreshFreestandingTemplate NSImageName
	// RefreshTemplate: A refresh template image.
	RefreshTemplate NSImageName
	// RemoveTemplate: A remove item template image.
	RemoveTemplate NSImageName
	// RevealFreestandingTemplate: A reveal contents template image.
	RevealFreestandingTemplate NSImageName
	// RightFacingTriangleTemplate: A generic right-facing triangle template image.
	RightFacingTriangleTemplate NSImageName
	// ShareTemplate: A share view template image.
	ShareTemplate NSImageName
	// SlideshowTemplate: A slideshow template image.
	SlideshowTemplate NSImageName
	// SmartBadgeTemplate: A badge for a “smart” item.
	SmartBadgeTemplate NSImageName
	// StatusAvailable: Small green indicator, similar to iChat’s available image.
	StatusAvailable NSImageName
	// StatusNone: Small clear indicator.
	StatusNone NSImageName
	// StatusPartiallyAvailable: Small yellow indicator, similar to iChat’s idle image.
	StatusPartiallyAvailable NSImageName
	// StatusUnavailable: Small red indicator, similar to iChat’s unavailable image.
	StatusUnavailable NSImageName
	// StopProgressFreestandingTemplate: A stop progress template image.
	StopProgressFreestandingTemplate NSImageName
	// StopProgressTemplate: A stop progress button template image.
	StopProgressTemplate NSImageName
	// TouchBarAddDetailTemplate: A template image for showing additional detail for an item.
	TouchBarAddDetailTemplate NSImageName
	// TouchBarAddTemplate: A template image for creating a new item.
	TouchBarAddTemplate NSImageName
	// TouchBarAlarmTemplate: A template image for setting or showing an alarm.
	TouchBarAlarmTemplate NSImageName
	// TouchBarAudioInputMuteTemplate: A template image for muting audio input or denoting that audio input is muted.
	TouchBarAudioInputMuteTemplate NSImageName
	// TouchBarAudioInputTemplate: A template image for denoting audio input.
	TouchBarAudioInputTemplate NSImageName
	// TouchBarAudioOutputMuteTemplate: A template image for muting audio output or for denoting that audio output is muted.
	TouchBarAudioOutputMuteTemplate NSImageName
	// TouchBarAudioOutputVolumeHighTemplate: A template image for setting the audio output volume to a high level, or for denoting that the audio is at the peak volume.
	TouchBarAudioOutputVolumeHighTemplate NSImageName
	// TouchBarAudioOutputVolumeLowTemplate: A template image for setting the audio output volume to a low level, or for denoting that it is set to a low level.
	TouchBarAudioOutputVolumeLowTemplate NSImageName
	// TouchBarAudioOutputVolumeMediumTemplate: A template image for setting the audio output volume to a medium level, or for denoting that it is set to a medium level.
	TouchBarAudioOutputVolumeMediumTemplate NSImageName
	// TouchBarAudioOutputVolumeOffTemplate: A template image for setting the audio output volume to silent, or for denoting that it is set to silent.
	TouchBarAudioOutputVolumeOffTemplate NSImageName
	// TouchBarBookmarksTemplate: A template image for showing app-specific bookmarks.
	TouchBarBookmarksTemplate NSImageName
	// TouchBarColorPickerFill: A template image for showing a color picker so the user can select a fill color.
	TouchBarColorPickerFill NSImageName
	// TouchBarColorPickerFont: A template image for showing a color picker so the user can select a text color.
	TouchBarColorPickerFont NSImageName
	// TouchBarColorPickerStroke: A template image for showing a color picker so the user can select a stroke color.
	TouchBarColorPickerStroke NSImageName
	// TouchBarCommunicationAudioTemplate: A template image for initiating or denoting audio communication.
	TouchBarCommunicationAudioTemplate NSImageName
	// TouchBarCommunicationVideoTemplate: A template image for initiating or denoting video communication.
	TouchBarCommunicationVideoTemplate NSImageName
	// TouchBarComposeTemplate: A template image for opening a new document or view in edit mode.
	TouchBarComposeTemplate NSImageName
	// TouchBarDeleteTemplate: A template image for deleting the current or selected item.
	TouchBarDeleteTemplate NSImageName
	// TouchBarDownloadTemplate: A template image for downloading an item.
	TouchBarDownloadTemplate NSImageName
	// TouchBarEnterFullScreenTemplate: A template image for entering full screen mode.
	TouchBarEnterFullScreenTemplate NSImageName
	// TouchBarExitFullScreenTemplate: A template image for exiting full screen mode.
	TouchBarExitFullScreenTemplate NSImageName
	// TouchBarFastForwardTemplate: A template image for moving forward through media playback or slides.
	TouchBarFastForwardTemplate NSImageName
	// TouchBarFolderCopyToTemplate: A template image for copying an item to a destination.
	TouchBarFolderCopyToTemplate NSImageName
	// TouchBarFolderMoveToTemplate: A template image for moving an item to a destination.
	TouchBarFolderMoveToTemplate NSImageName
	// TouchBarFolderTemplate: A template image for opening or representing a folder.
	TouchBarFolderTemplate NSImageName
	// TouchBarGetInfoTemplate: A template image for showing information about an item.
	TouchBarGetInfoTemplate NSImageName
	// TouchBarGoBackTemplate: A template image for returning to the previous screen or location.
	TouchBarGoBackTemplate NSImageName
	// TouchBarGoDownTemplate: A template image for moving to the next item in a list.
	TouchBarGoDownTemplate NSImageName
	// TouchBarGoForwardTemplate: A template image for moving to the next screen or location.
	TouchBarGoForwardTemplate NSImageName
	// TouchBarGoUpTemplate: A template image for moving to the previous item in a list.
	TouchBarGoUpTemplate NSImageName
	// TouchBarHistoryTemplate: A template image for showing history information, such as recent downloads.
	TouchBarHistoryTemplate NSImageName
	// TouchBarIconViewTemplate: A template image for showing items in an icon view.
	TouchBarIconViewTemplate NSImageName
	// TouchBarListViewTemplate: A template image for showing items in a list view.
	TouchBarListViewTemplate NSImageName
	// TouchBarMailTemplate: A template image for creating an email message.
	TouchBarMailTemplate NSImageName
	// TouchBarNewFolderTemplate: A template image for creating a new folder.
	TouchBarNewFolderTemplate NSImageName
	// TouchBarNewMessageTemplate: A template image for creating a new message, or for denoting the use of messaging.
	TouchBarNewMessageTemplate NSImageName
	// TouchBarOpenInBrowserTemplate: A template image for opening an item in the user’s browser.
	TouchBarOpenInBrowserTemplate NSImageName
	// TouchBarPauseTemplate: A template image for pausing media playback or slides.
	TouchBarPauseTemplate NSImageName
	// TouchBarPlayPauseTemplate: A template image for toggling between playing and pausing media or slides.
	TouchBarPlayPauseTemplate NSImageName
	// TouchBarPlayTemplate: A template image for starting or resuming playback of media or slides.
	TouchBarPlayTemplate NSImageName
	// TouchBarPlayheadTemplate: A template image for denoting the current playback position within a timeline track.
	TouchBarPlayheadTemplate NSImageName
	// TouchBarQuickLookTemplate: A template image for opening an item in Quick Look.
	TouchBarQuickLookTemplate NSImageName
	// TouchBarRecordStartTemplate: A template image for starting recording.
	TouchBarRecordStartTemplate NSImageName
	// TouchBarRecordStopTemplate: A template image for stopping recording or stopping playback of media or slides.
	TouchBarRecordStopTemplate NSImageName
	// TouchBarRefreshTemplate: A template image for refreshing displayed data.
	TouchBarRefreshTemplate NSImageName
	// TouchBarRemoveTemplate: A template image for removing an item.
	TouchBarRemoveTemplate NSImageName
	// TouchBarRewindTemplate: A template image for moving backwards through media or slides.
	TouchBarRewindTemplate NSImageName
	// TouchBarRotateLeftTemplate: A template image for rotating an item counterclockwise.
	TouchBarRotateLeftTemplate NSImageName
	// TouchBarRotateRightTemplate: A template image for rotating an item clockwise.
	TouchBarRotateRightTemplate NSImageName
	// TouchBarSearchTemplate: A template image for showing a search field or for initiating a search.
	TouchBarSearchTemplate NSImageName
	// TouchBarShareTemplate: A template image for sharing content with others directly or via social media.
	TouchBarShareTemplate NSImageName
	// TouchBarSidebarTemplate: A template image for showing a sidebar in the current view.
	TouchBarSidebarTemplate NSImageName
	// TouchBarSkipAhead15SecondsTemplate: A template image for skipping ahead 15 seconds during media playback.
	TouchBarSkipAhead15SecondsTemplate NSImageName
	// TouchBarSkipAhead30SecondsTemplate: A template image for skipping ahead 30 seconds during media playback.
	TouchBarSkipAhead30SecondsTemplate NSImageName
	// TouchBarSkipAheadTemplate: A template image for skipping to the next chapter or location during media playback.
	TouchBarSkipAheadTemplate NSImageName
	// TouchBarSkipBack15SecondsTemplate: A template image for skipping back 15 seconds during media playback.
	TouchBarSkipBack15SecondsTemplate NSImageName
	// TouchBarSkipBack30SecondsTemplate: A template image for skipping back 30 seconds during media playback.
	TouchBarSkipBack30SecondsTemplate NSImageName
	// TouchBarSkipBackTemplate: A template image for skipping to the previous chapter or location during media playback.
	TouchBarSkipBackTemplate NSImageName
	// TouchBarSkipToEndTemplate: A template image for skipping to the end of media playback.
	TouchBarSkipToEndTemplate NSImageName
	// TouchBarSkipToStartTemplate: A template image for skipping to the start of media playback.
	TouchBarSkipToStartTemplate NSImageName
	// TouchBarSlideshowTemplate: A template image for starting a slideshow.
	TouchBarSlideshowTemplate NSImageName
	// TouchBarTagIconTemplate: A template image for applying a tag to an item.
	TouchBarTagIconTemplate NSImageName
	// TouchBarTextBoldTemplate: A template image for making selected text bold.
	TouchBarTextBoldTemplate NSImageName
	// TouchBarTextBoxTemplate: A template image for inserting a text box.
	TouchBarTextBoxTemplate NSImageName
	// TouchBarTextCenterAlignTemplate: A template image for centering text.
	TouchBarTextCenterAlignTemplate NSImageName
	// TouchBarTextItalicTemplate: A template image for italicizing the selected text.
	TouchBarTextItalicTemplate NSImageName
	// TouchBarTextJustifiedAlignTemplate: A template image for fully justifying text.
	TouchBarTextJustifiedAlignTemplate NSImageName
	// TouchBarTextLeftAlignTemplate: A template image for aligning text to the left.
	TouchBarTextLeftAlignTemplate NSImageName
	// TouchBarTextListTemplate: A template image for inserting a list or converting text to list form.
	TouchBarTextListTemplate NSImageName
	// TouchBarTextRightAlignTemplate: A template image for aligning text to the right.
	TouchBarTextRightAlignTemplate NSImageName
	// TouchBarTextStrikethroughTemplate: A template image for striking through text.
	TouchBarTextStrikethroughTemplate NSImageName
	// TouchBarTextUnderlineTemplate: A template image for underlining text.
	TouchBarTextUnderlineTemplate NSImageName
	// TouchBarUserAddTemplate: A template image for creating a new user account.
	TouchBarUserAddTemplate NSImageName
	// TouchBarUserGroupTemplate: A template image for showing or representing a group of users.
	TouchBarUserGroupTemplate NSImageName
	// TouchBarUserTemplate: A template image for showing or representing user information.
	TouchBarUserTemplate NSImageName
	// TouchBarVolumeDownTemplate: A template image for reducing the audio output volume.
	TouchBarVolumeDownTemplate NSImageName
	// TouchBarVolumeUpTemplate: A template image for increasing the audio output volume.
	TouchBarVolumeUpTemplate NSImageName
	// TrashEmpty: An image of the empty trash can.
	TrashEmpty NSImageName
	// TrashFull: An image of the full trash can.
	TrashFull NSImageName
	// User: Permissions for a single user.
	User NSImageName
	// UserAccounts: User account toolbar icon for the preferences window.
	UserAccounts NSImageName
	// UserGroup: Permissions for a group of users.
	UserGroup NSImageName
	// UserGuest: Permissions for guests.
	UserGuest NSImageName
}

// NSLayoutPrioritys provides typed accessors for [NSLayoutPriority] constants.
var NSLayoutPrioritys struct {
	// DefaultHigh: Priority level with which a button resists compressing its content.
	DefaultHigh NSLayoutPriority
	// DefaultLow: Priority level at which a button hugs its contents horizontally.
	DefaultLow NSLayoutPriority
	// DragThatCanResizeWindow: Appropriate priority level for a drag that may end up resizing the window.
	DragThatCanResizeWindow NSLayoutPriority
	// DragThatCannotResizeWindow: Priority level at which a split view divider, say, is dragged.
	DragThatCannotResizeWindow NSLayoutPriority
	// FittingSizeCompression: When you send a [fittingSize](<doc://com.apple.appkit/documentation/AppKit/NSView/fittingSize>) message to a view, the smallest size that is large enough for the view’s contents is computed.
	FittingSizeCompression NSLayoutPriority
	// Required: A required constraint.
	Required NSLayoutPriority
	// WindowSizeStayPut: Priority level for the window’s current size.
	WindowSizeStayPut NSLayoutPriority
}

// NSModalResponses provides typed accessors for [NSModalResponse] constants.
var NSModalResponses struct {
	// Abort: Modal session was broken with [abortModal()](<doc://com.apple.appkit/documentation/AppKit/NSApplication/abortModal()>).
	Abort NSModalResponse
	// Cancel: The presentation or dismissal of the sheet has been canceled.
	Cancel NSModalResponse
	// Continue: Modal session is continuing (returned by [runModalSession(_:)](<doc://com.apple.appkit/documentation/AppKit/NSApplication/runModalSession(_:)>) only).
	Continue NSModalResponse
	// OK: The presentation or dismissal of the sheet has finished.
	OK NSModalResponse
	// Stop: Modal session was broken with [stopModal()](<doc://com.apple.appkit/documentation/AppKit/NSApplication/stopModal()>).
	Stop NSModalResponse
}

// NSPasteboardDetectionPatterns provides typed accessors for [NSPasteboardDetectionPattern] constants.
var NSPasteboardDetectionPatterns struct {
	// CalendarEvent: A pattern that indicates the pasteboard detects a string that contains a calendar event.
	CalendarEvent NSPasteboardDetectionPattern
	// EmailAddress: A pattern that indicates the pasteboard detects a string that contains an email address.
	EmailAddress NSPasteboardDetectionPattern
	// FlightNumber: A pattern that indicates the pasteboard detects a string that contains a flight number.
	FlightNumber NSPasteboardDetectionPattern
	// Link: A pattern that indicates the pasteboard detects a string that contains a URL.
	Link NSPasteboardDetectionPattern
	// MoneyAmount: A pattern that indicates the pasteboard detects a string that contains an amount of money.
	MoneyAmount NSPasteboardDetectionPattern
	// Number: A pattern that indicates the pasteboard detects a string that consists of a numeric value.
	Number NSPasteboardDetectionPattern
	// PhoneNumber: A pattern that indicates the pasteboard detects a string that contains a phone number.
	PhoneNumber NSPasteboardDetectionPattern
	// PostalAddress: A pattern that indicates the pasteboard detects a string that contains a postal address.
	PostalAddress NSPasteboardDetectionPattern
	// ProbableWebSearch: A pattern that indicates the pasteboard detects a string suitable for use as a web search term.
	ProbableWebSearch NSPasteboardDetectionPattern
	// ProbableWebURL: A pattern that indicates the pasteboard detects a string that consists of a web URL.
	ProbableWebURL NSPasteboardDetectionPattern
	// ShipmentTrackingNumber: A pattern that indicates the pasteboard detects a string that contains a parcel tracking number and carrier.
	ShipmentTrackingNumber NSPasteboardDetectionPattern
}

// NSPasteboardNames provides typed accessors for [NSPasteboardName] constants.
var NSPasteboardNames struct {
	// Drag: The pasteboard that stores data to move as the result of a drag operation.
	Drag NSPasteboardName
	// Find: The pasteboard that holds information about the current state of the active application’s find panel.
	Find NSPasteboardName
	// Font: The pasteboard that holds font and character information and supports Copy Font and Paste Font commands that the text editor may implement.
	Font NSPasteboardName
	// General: The pasteboard you use to perform ordinary cut, copy, and paste operations.
	General NSPasteboardName
	// Ruler: The pasteboard that holds information about paragraph formats and supports the Copy Ruler and Paste Ruler commands that the text editor may implement.
	Ruler NSPasteboardName
}

// NSPasteboardTypes provides typed accessors for [NSPasteboardType] constants.
var NSPasteboardTypes struct {
	// Color: Color data.
	Color NSPasteboardType
	// FileURL: A file URL.
	FileURL NSPasteboardType
	// Font: Font and character information.
	Font NSPasteboardType
	// HTML: Type for HTML content.
	HTML NSPasteboardType
	// MultipleTextSelection: Multiple text selection.
	MultipleTextSelection NSPasteboardType
	// PDF: PDF data.
	PDF NSPasteboardType
	// PNG: PNG image data.
	PNG NSPasteboardType
	// RTF: Rich Text Format (RTF) data.
	RTF NSPasteboardType
	// RTFD: RTFD formatted file contents.
	RTFD NSPasteboardType
	// Ruler: Paragraph formatting information.
	Ruler NSPasteboardType
	// Sound: Sound data.
	Sound NSPasteboardType
	// String: String data.
	String NSPasteboardType
	// TIFF: Tag Image File Format (TIFF) data.
	TIFF NSPasteboardType
	// TabularText: Tab-separated fields of text.
	TabularText NSPasteboardType
	// TextFinderOptions: Type for the Find panel metadata property list.
	TextFinderOptions NSPasteboardType
	// URL: URL data for one file or resource.
	URL NSPasteboardType
}

// NSSharingServiceNames provides typed accessors for [NSSharingServiceName] constants.
var NSSharingServiceNames struct {
	// AddToAperture: A service that shares an item provider’s contents with Aperture.
	AddToAperture NSSharingServiceName
	// AddToIPhoto: A service that shares an item provider’s contents with iPhoto.
	AddToIPhoto NSSharingServiceName
	// AddToSafariReadingList: A service that shares an item provider’s contents with Safari’s Reading List.
	AddToSafariReadingList NSSharingServiceName
	// CloudSharing: A service that shares an item provider’s contents with other iCloud users.
	CloudSharing NSSharingServiceName
	// ComposeEmail: A service that uses an item provider’s contents to compose an email.
	ComposeEmail NSSharingServiceName
	// ComposeMessage: A service that uses an item provider’s contents to compose a message.
	ComposeMessage NSSharingServiceName
	// SendViaAirDrop: A service that sends an item provider’s contents to another device using AirDrop.
	SendViaAirDrop NSSharingServiceName
	// UseAsDesktopPicture: A service that sets the item provider’s contents as the current user’s desktop picture.
	UseAsDesktopPicture NSSharingServiceName
}

// NSSliderAccessoryWidths provides typed accessors for [NSSliderAccessoryWidth] constants.
var NSSliderAccessoryWidths struct {
	// Default: The default width for slider accessories.
	Default NSSliderAccessoryWidth
	Wide    NSSliderAccessoryWidth
}

// NSSpeechModes provides typed accessors for [NSSpeechMode] constants.
var NSSpeechModes struct {
	// Literal: Indicates that each digit or character is spoken literally (so that 12 is spoken as “one, two”, or the word “cat” is spoken as “C A T”).
	Literal NSSpeechMode
	// Normal: Indicates that the synthesizer assembles digits into numbers (so that 12 is spoken as “twelve”) and text into words.
	Normal NSSpeechMode
	// Phoneme: Indicates that the synthesizer is in phoneme-processing mode. When in phoneme-processing mode, a text buffer is interpreted to be a series of characters representing various phonemes and prosodic controls.
	Phoneme NSSpeechMode
	// Text: Indicates that the synthesizer is in text-processing mode.
	Text NSSpeechMode
}

// NSStackViewVisibilityPrioritys provides typed accessors for [NSStackViewVisibilityPriority] constants.
var NSStackViewVisibilityPrioritys struct {
	// DetachOnlyIfNecessary: The Auto Layout priority that results in detachment of a view when there is insufficient space in the stack view to display it fully.
	DetachOnlyIfNecessary NSStackViewVisibilityPriority
	// MustHold: The default value, and maximum Auto Layout priority, that results in a view never detaching from the stack view.
	MustHold NSStackViewVisibilityPriority
	// NotVisible: The minimum Auto Layout priority that forces a view to detach from the stack view.
	NotVisible NSStackViewVisibilityPriority
}

// NSTextContentTypes provides typed accessors for [NSTextContentType] constants.
var NSTextContentTypes struct {
	// AddressCity: A property that defines the content in a text input area as a city name.
	AddressCity NSTextContentType
	// AddressCityAndState: A property that defines the content in a text input area as a city name with a state name.
	AddressCityAndState NSTextContentType
	// AddressState: A property that defines the content in a text input area as a state name.
	AddressState NSTextContentType
	// Birthdate: A property that defines the content in a text input area as a date of birth.
	Birthdate NSTextContentType
	// BirthdateDay: A property that defines the content in a text input area as the day component of a birthdate.
	BirthdateDay NSTextContentType
	// BirthdateMonth: A property that defines the content in a text input area as the month component of a birthdate.
	BirthdateMonth NSTextContentType
	// BirthdateYear: A property that defines the content in a text input area as the year component of a birthdate.
	BirthdateYear NSTextContentType
	// CountryName: A property that defines the content in a text input area as a country or region name.
	CountryName NSTextContentType
	// CreditCardExpiration: A property that defines the content in a text input area as an expiration date on a credit card.
	CreditCardExpiration NSTextContentType
	// CreditCardExpirationMonth: A property that defines the content in a text input area as the month component of an expiration date on a credit card.
	CreditCardExpirationMonth NSTextContentType
	// CreditCardExpirationYear: A property that defines the content in a text input area as the year component of an expiration date on a credit card.
	CreditCardExpirationYear NSTextContentType
	// CreditCardFamilyName: A property that defines the content in a text input area as a family name, or last name, on a credit card.
	CreditCardFamilyName NSTextContentType
	// CreditCardGivenName: A property that defines the content in a text input area as a first name on a credit card.
	CreditCardGivenName NSTextContentType
	// CreditCardMiddleName: A property that defines the content in a text input area as a middle name on a credit card.
	CreditCardMiddleName NSTextContentType
	// CreditCardName: A property that defines the content in a text input area as a name on a credit card.
	CreditCardName NSTextContentType
	// CreditCardNumber: A property that defines the content in a text input area as a credit card number.
	CreditCardNumber NSTextContentType
	// CreditCardSecurityCode: A property that defines the content in a text input area as a credit card security code.
	CreditCardSecurityCode NSTextContentType
	// CreditCardType: A property that defines the content in a text input area as a credit card type.
	CreditCardType NSTextContentType
	// DateTime: A property that defines the content in a text input area as a date, time, or duration.
	DateTime NSTextContentType
	// EmailAddress: A property that defines the content in a text input area as an email address.
	EmailAddress NSTextContentType
	// FamilyName: A property that defines the content in a text input area as a family name, or last name.
	FamilyName NSTextContentType
	// FlightNumber: A property that defines the content in a text input area as an airline flight number.
	FlightNumber NSTextContentType
	// FullStreetAddress: A property that defines the content in a text input area as a street address that fully identifies a location.
	FullStreetAddress NSTextContentType
	// GivenName: A property that defines the content in a text input area as a first name.
	GivenName NSTextContentType
	// JobTitle: A property that defines the content in a text input area as a job title.
	JobTitle NSTextContentType
	// Location: A property that defines the content in a text input area as a location, such as a point of interest, an address, or another identifier for a location.
	Location NSTextContentType
	// MiddleName: A property that defines the content in a text input area as a middle name.
	MiddleName NSTextContentType
	// Name: A property that defines the content in a text input area as a name.
	Name NSTextContentType
	// NamePrefix: A property that defines the content in a text input area as a prefix or title, such as .
	NamePrefix NSTextContentType
	// NameSuffix: A property that defines the content in a text input area as a suffix, such as .
	NameSuffix NSTextContentType
	// NewPassword: A property that defines the content in a text input area as a new password.
	NewPassword NSTextContentType
	// Nickname: A property that defines the content in a text input area as a nickname.
	Nickname NSTextContentType
	// OneTimeCode: A property that defines the content in a text input area as a one-time code.
	OneTimeCode NSTextContentType
	// OrganizationName: A property that defines the content in a text input area as an organization name.
	OrganizationName NSTextContentType
	// Password: A property that defines the content in a text input area as a password.
	Password NSTextContentType
	// PostalCode: A property that defines the content in a text input area as a postal code.
	PostalCode NSTextContentType
	// ShipmentTrackingNumber: A property that defines the content in a text input area as a parcel tracking number.
	ShipmentTrackingNumber NSTextContentType
	// StreetAddressLine1: A property that defines the content in a text input area as the first line of a street address.
	StreetAddressLine1 NSTextContentType
	// StreetAddressLine2: A property that defines the content in a text input area as the second line of a street address.
	StreetAddressLine2 NSTextContentType
	// Sublocality: A property that defines the content in a text input area as a sublocality.
	Sublocality NSTextContentType
	// TelephoneNumber: A property that defines the content in a text input area as a telephone number.
	TelephoneNumber NSTextContentType
	// URL: A property that defines the content in a text input area as a URL.
	URL NSTextContentType
	// Username: A property that defines the content in a text input area as an account or login name.
	Username NSTextContentType
}

// NSTextHighlightColorSchemes provides typed accessors for [NSTextHighlightColorScheme] constants.
var NSTextHighlightColorSchemes struct {
	// Blue: A blue highlight color.
	Blue NSTextHighlightColorScheme
	// Default: The default system highlight color.
	Default NSTextHighlightColorScheme
	// Mint: A mint green highlight color.
	Mint NSTextHighlightColorScheme
	// Orange: An orange highlight color.
	Orange NSTextHighlightColorScheme
	// Pink: A pink highlight color.
	Pink NSTextHighlightColorScheme
	// Purple: A purple highlight color.
	Purple NSTextHighlightColorScheme
}

// NSToolbarItemVisibilityPrioritys provides typed accessors for [NSToolbarItemVisibilityPriority] constants.
var NSToolbarItemVisibilityPrioritys struct {
	// High: A high priority that makes it less likely for the toolbar item to move to the overflow item.
	High NSToolbarItemVisibilityPriority
	// Low: The lowest-priority for a toolbar item.
	Low NSToolbarItemVisibilityPriority
	// Standard: The default visibility priority.
	Standard NSToolbarItemVisibilityPriority
	// User: The highest priority for items in the toolbar.
	User NSToolbarItemVisibilityPriority
}

// NSTouchBarItemIdentifiers provides typed accessors for [NSTouchBarItemIdentifier] constants.
var NSTouchBarItemIdentifiers struct {
	// CandidateList: The standard identifier for a candidate list bar item.
	CandidateList NSTouchBarItemIdentifier
	// CharacterPicker: The standard identifier for selecting special characters such as Emoji.
	CharacterPicker NSTouchBarItemIdentifier
	// FixedSpaceLarge: The identifier of an item appropriate for use as a large space in a Touch Bar.
	FixedSpaceLarge NSTouchBarItemIdentifier
	// FixedSpaceSmall: The identifier of an item appropriate for use as a small space in a Touch Bar.
	FixedSpaceSmall NSTouchBarItemIdentifier
	// FlexibleSpace: The identifier of an item appropriate for use as a flexible space in a Touch Bar.
	FlexibleSpace NSTouchBarItemIdentifier
	// OtherItemsProxy: The identifier of the special “other items proxy”, which is used to nest bars up the responder chain.
	OtherItemsProxy NSTouchBarItemIdentifier
	// TextAlignment: The identifier for a Touch Bar item used to select the text alignment.
	TextAlignment NSTouchBarItemIdentifier
	// TextColorPicker: The identifier for a Touch Bar item used to select the text color.
	TextColorPicker NSTouchBarItemIdentifier
	// TextFormat: The identifier for a group of text format controls.
	TextFormat NSTouchBarItemIdentifier
	// TextList: The identifier for a Touch Bar item used to control the text list style.
	TextList NSTouchBarItemIdentifier
	// TextStyle: The identifier for a Touch Bar item used to control the text style.
	TextStyle NSTouchBarItemIdentifier
}

// NSTouchBarItemPrioritys provides typed accessors for [NSTouchBarItemPriority] constants.
var NSTouchBarItemPrioritys struct {
	// High: A constant indicating a high visibility priority.
	High NSTouchBarItemPriority
	// Low: A constant indicating a low visibility priority.
	Low NSTouchBarItemPriority
	// Normal: A constant indicating a normal visibility priority.
	Normal NSTouchBarItemPriority
}
