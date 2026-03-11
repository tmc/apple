// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"fmt"
)

type NS uint

const (
	// NSBackTabCharacter: The back tab character: `0x0019`
	NSBackTabCharacter NS = 25
	// NSBackspaceCharacter: The backspace character: `0x0008`
	NSBackspaceCharacter NS = 8
	// NSBacktabTextMovement: The Backtab (Shift-Tab) key was pressed.
	NSBacktabTextMovement NS = 18
	// NSBeginFunctionKey: The begin key.
	NSBeginFunctionKey NS = 63274
	// NSBreakFunctionKey: The break key.
	NSBreakFunctionKey NS = 63282
	// NSCancelTextMovement: The user cancelled the completion.
	NSCancelTextMovement NS = 23
	// NSCarriageReturnCharacter: The carriage return character: `0x000d`
	NSCarriageReturnCharacter NS = 13
	// NSClearDisplayFunctionKey: The clear display key.
	NSClearDisplayFunctionKey NS = 63290
	// NSClearLineFunctionKey: The clear or num lock key.
	NSClearLineFunctionKey NS = 63289
	// NSControlGlyph: The reserved code for a control glyph.
	NSControlGlyph NS = 16777215
	// NSDeleteCharFunctionKey: The delete character key.
	NSDeleteCharFunctionKey NS = 63294
	// NSDeleteCharacter: The delete character: `0x007f`
	NSDeleteCharacter NS = 127
	// NSDeleteFunctionKey: The forward delete key.
	NSDeleteFunctionKey NS = 63272
	// NSDeleteLineFunctionKey: The delete line key.
	NSDeleteLineFunctionKey NS = 63292
	// NSDownArrowFunctionKey: The down arrow key.
	NSDownArrowFunctionKey NS = 63233
	// NSDownTextMovement: The down arrow key was pressed.
	NSDownTextMovement NS = 22
	// NSEndFunctionKey: The end key.
	NSEndFunctionKey NS = 63275
	// NSEnterCharacter: The enter character: `0x0003`
	NSEnterCharacter NS = 3
	// NSExecuteFunctionKey: The execute key.
	NSExecuteFunctionKey NS = 63298
	// NSF10FunctionKey: The F10 key.
	NSF10FunctionKey NS = 63245
	// NSF11FunctionKey: The F11 key.
	NSF11FunctionKey NS = 63246
	// NSF12FunctionKey: The F12 key.
	NSF12FunctionKey NS = 63247
	// NSF13FunctionKey: The F13 key.
	NSF13FunctionKey NS = 63248
	// NSF14FunctionKey: The F14 key.
	NSF14FunctionKey NS = 63249
	// NSF15FunctionKey: The F15 key.
	NSF15FunctionKey NS = 63250
	// NSF16FunctionKey: The F16 key.
	NSF16FunctionKey NS = 63251
	// NSF17FunctionKey: The F17 key.
	NSF17FunctionKey NS = 63252
	// NSF18FunctionKey: The F18 key.
	NSF18FunctionKey NS = 63253
	// NSF19FunctionKey: The F19 key.
	NSF19FunctionKey NS = 63254
	// NSF1FunctionKey: The F1 key.
	NSF1FunctionKey NS = 63236
	// NSF20FunctionKey: The F20 key.
	NSF20FunctionKey NS = 63255
	// NSF21FunctionKey: The F21 key.
	NSF21FunctionKey NS = 63256
	// NSF22FunctionKey: The F22 key.
	NSF22FunctionKey NS = 63257
	// NSF23FunctionKey: The F23 key.
	NSF23FunctionKey NS = 63258
	// NSF24FunctionKey: The F24 key.
	NSF24FunctionKey NS = 63259
	// NSF25FunctionKey: The F25 key.
	NSF25FunctionKey NS = 63260
	// NSF26FunctionKey: The F26 key.
	NSF26FunctionKey NS = 63261
	// NSF27FunctionKey: The F27 key.
	NSF27FunctionKey NS = 63262
	// NSF28FunctionKey: The F28 key.
	NSF28FunctionKey NS = 63263
	// NSF29FunctionKey: The F29 key.
	NSF29FunctionKey NS = 63264
	// NSF2FunctionKey: The F2 key.
	NSF2FunctionKey NS = 63237
	// NSF30FunctionKey: The F30 key.
	NSF30FunctionKey NS = 63265
	// NSF31FunctionKey: The F31 key.
	NSF31FunctionKey NS = 63266
	// NSF32FunctionKey: The F32 key.
	NSF32FunctionKey NS = 63267
	// NSF33FunctionKey: The F33 key.
	NSF33FunctionKey NS = 63268
	// NSF34FunctionKey: The F34 key.
	NSF34FunctionKey NS = 63269
	// NSF35FunctionKey: The F35 key.
	NSF35FunctionKey NS = 63270
	// NSF3FunctionKey: The F3 key.
	NSF3FunctionKey NS = 63238
	// NSF4FunctionKey: The F4 key.
	NSF4FunctionKey NS = 63239
	// NSF5FunctionKey: The F5 key.
	NSF5FunctionKey NS = 63240
	// NSF6FunctionKey: The F6 key.
	NSF6FunctionKey NS = 63241
	// NSF7FunctionKey: The F7 key.
	NSF7FunctionKey NS = 63242
	// NSF8FunctionKey: The F8 key.
	NSF8FunctionKey NS = 63243
	// NSF9FunctionKey: The F9 key.
	NSF9FunctionKey NS = 63244
	// NSFindFunctionKey: The find key.
	NSFindFunctionKey NS = 63301
	NSFontAssetDownloadError NS = 66304
	NSFontErrorMaximum NS = 66335
	NSFontErrorMinimum NS = 66304
	// NSFormFeedCharacter: The form feed character: `0x000c`
	NSFormFeedCharacter NS = 12
	// NSHelpFunctionKey: The help key.
	NSHelpFunctionKey NS = 63302
	// NSHomeFunctionKey: The home key.
	NSHomeFunctionKey NS = 63273
	// NSIllegalTextMovement: Currently unused.
	NSIllegalTextMovement NS = 0
	// NSInsertCharFunctionKey: The insert character key.
	NSInsertCharFunctionKey NS = 63293
	// NSInsertFunctionKey: The insert key.
	NSInsertFunctionKey NS = 63271
	// NSInsertLineFunctionKey: The insert line key.
	NSInsertLineFunctionKey NS = 63291
	// NSLeftArrowFunctionKey: The left arrow key.
	NSLeftArrowFunctionKey NS = 63234
	// NSLeftTextMovement: The left arrow key was pressed.
	NSLeftTextMovement NS = 19
	// NSLineSeparatorCharacter: The line separator character: `0x2028`
	NSLineSeparatorCharacter NS = 8232
	// NSMenuFunctionKey: The menu key.
	NSMenuFunctionKey NS = 63285
	// NSModeSwitchFunctionKey: The mode switch key.
	NSModeSwitchFunctionKey NS = 63303
	// NSNewlineCharacter: The newline character: `0x000a`
	NSNewlineCharacter NS = 10
	// NSNextFunctionKey: The next key.
	NSNextFunctionKey NS = 63296
	// NSNullGlyph: The reserved code for a null glyph.
	NSNullGlyph NS = 0
	// NSOtherTextMovement: The user performed some undefined action.
	NSOtherTextMovement NS = 0
	// NSPageDownFunctionKey: The page down key.
	NSPageDownFunctionKey NS = 63277
	// NSPageUpFunctionKey: The page up key.
	NSPageUpFunctionKey NS = 63276
	// NSParagraphSeparatorCharacter: The paragraph separator character: `0x2029`
	NSParagraphSeparatorCharacter NS = 8233
	NSPasteboardCommunicationError NS = 67585
	NSPasteboardContentsNotAvailableError NS = 67587
	NSPasteboardErrorMaximum NS = 67839
	NSPasteboardErrorMinimum NS = 67584
	NSPasteboardInvalidArgumentError NS = 67586
	NSPasteboardMiscellaneousError NS = 67584
	// NSPauseFunctionKey: The pause key.
	NSPauseFunctionKey NS = 63280
	// NSPrevFunctionKey: Previous key.
	NSPrevFunctionKey NS = 63295
	// NSPrintFunctionKey: The print key.
	NSPrintFunctionKey NS = 63288
	// NSPrintScreenFunctionKey: The print screen key.
	NSPrintScreenFunctionKey NS = 63278
	// NSRedoFunctionKey: The redo key.
	NSRedoFunctionKey NS = 63300
	// NSResetFunctionKey: The reset key.
	NSResetFunctionKey NS = 63283
	// NSReturnTextMovement: The Return key was pressed.
	NSReturnTextMovement NS = 16
	// NSRightArrowFunctionKey: The right arrow key.
	NSRightArrowFunctionKey NS = 63235
	// NSRightTextMovement: The right arrow key was pressed.
	NSRightTextMovement NS = 20
	// NSScrollLockFunctionKey: The scroll lock key.
	NSScrollLockFunctionKey NS = 63279
	// NSSelectFunctionKey: The select key.
	NSSelectFunctionKey NS = 63297
	// NSServiceApplicationLaunchFailedError: The service providing application could not be launched.
	NSServiceApplicationLaunchFailedError NS = 66561
	// NSServiceApplicationNotFoundError: The service provider could not be found.
	NSServiceApplicationNotFoundError NS = 66560
	// NSServiceErrorMaximum: Inclusive service error range, for checking future error codes.
	NSServiceErrorMaximum NS = 66817
	// NSServiceErrorMinimum: Inclusive service error range, for checking future error codes.
	NSServiceErrorMinimum NS = 66560
	// NSServiceInvalidPasteboardDataError: The service providing app did not return a pasteboard with any of the promised types, or we couldn’t write the data from the pasteboard to the object receiving the returned data.
	NSServiceInvalidPasteboardDataError NS = 66563
	// NSServiceMalformedServiceDictionaryError: The service dictionary did not contain the necessary keys.
	NSServiceMalformedServiceDictionaryError NS = 66564
	// NSServiceMiscellaneousError: Other errors, representing programmatic mistakes in the service consuming application.
	NSServiceMiscellaneousError NS = 66800
	// NSServiceRequestTimedOutError: The service providing application did not open its service listening port in time, or the app didn’t respond to the request in time; see the Console log to figure out which (the errors are typically reported the same way to the user).
	NSServiceRequestTimedOutError NS = 66562
	NSSharingServiceErrorMaximum NS = 67327
	NSSharingServiceErrorMinimum NS = 67072
	NSSharingServiceNotConfiguredError NS = 67072
	// NSShowControlGlyphs: Generates displayable glyphs for control characters.
	NSShowControlGlyphs NS = 1
	// NSShowInvisibleGlyphs: Generates displayable glyphs for invisible characters.
	NSShowInvisibleGlyphs NS = 2
	// NSStopFunctionKey: The stop key.
	NSStopFunctionKey NS = 63284
	// NSSysReqFunctionKey: The system request key.
	NSSysReqFunctionKey NS = 63281
	// NSSystemFunctionKey: The system key.
	NSSystemFunctionKey NS = 63287
	// NSTabCharacter: The tab character: `0x0009`
	NSTabCharacter NS = 9
	// NSTabTextMovement: The Tab key was pressed.
	NSTabTextMovement NS = 17
	// NSTextReadInapplicableDocumentTypeError: Indicates a problem reading data with the specified format.
	NSTextReadInapplicableDocumentTypeError NS = 65806
	// NSTextReadWriteErrorMaximum: The end of a range of error codes reserved for future use.
	NSTextReadWriteErrorMaximum NS = 66303
	// NSTextReadWriteErrorMinimum: The beginning of a range of error codes reserved for future use.
	NSTextReadWriteErrorMinimum NS = 65792
	// NSTextWriteInapplicableDocumentTypeError: Indicates a problem writing data of the specified format.
	NSTextWriteInapplicableDocumentTypeError NS = 66062
	// NSUndoFunctionKey: The undo key.
	NSUndoFunctionKey NS = 63299
	// NSUpArrowFunctionKey: The up arrow key.
	NSUpArrowFunctionKey NS = 63232
	// NSUpTextMovement: The up arrow key was pressed.
	NSUpTextMovement NS = 21
	// NSUserFunctionKey: The user key.
	NSUserFunctionKey NS = 63286
	// NSWantsBidiLevels: Generates directional formatting codes for bidirectional text.
	NSWantsBidiLevels NS = 4
	NSWindowSharingErrorMaximum NS = 67466
	NSWindowSharingErrorMinimum NS = 67456
	NSWindowSharingRequestAlreadyRequested NS = 67456
	NSWindowSharingRequestNoEligibleSession NS = 67457
	NSWindowSharingRequestUnspecifiedError NS = 67458
	// NSWorkspaceAuthorizationInvalidError: The provided workspace authorization credentials expired or are invalid.
	NSWorkspaceAuthorizationInvalidError NS = 67328
	NSWorkspaceErrorMaximum NS = 67455
	NSWorkspaceErrorMinimum NS = 67328
	// Deprecated: use NSNullGlyph.
	NSAnyType NS = 0
	// Deprecated: use NSNullGlyph.
	NSCancelButton NS = 0
	// Deprecated.
	NSDoubleType NS = 6
	// Deprecated: use NSEnterCharacter.
	NSFloatType NS = 3
	// Deprecated: use NSShowControlGlyphs.
	NSIntType NS = 1
	// Deprecated: use NSEnterCharacter.
	NSMacintoshInterfaceStyle NS = 3
	// Deprecated: use NSShowControlGlyphs.
	NSNextStepInterfaceStyle NS = 1
	// Deprecated: use NSNullGlyph.
	NSNoInterfaceStyle NS = 0
	// Deprecated: use NSNullGlyph.
	NSNoUnderlineStyle NS = 0
	// Deprecated: use NSNullGlyph.
	NSOKButton NS = 0
	// Deprecated.
	NSPositiveDoubleType NS = 7
	// Deprecated: use NSWantsBidiLevels.
	NSPositiveFloatType NS = 4
	// Deprecated: use NSShowInvisibleGlyphs.
	NSPositiveIntType NS = 2
	// Deprecated: use NSShowControlGlyphs.
	NSSingleUnderlineStyle NS = 1
	// Deprecated: use NSShowInvisibleGlyphs.
	NSWindows95InterfaceStyle NS = 2
)


func (e NS) String() string {
	switch e {
	case NSBackTabCharacter:
		return "NSBackTabCharacter"
	case NSBackspaceCharacter:
		return "NSBackspaceCharacter"
	case NSBacktabTextMovement:
		return "NSBacktabTextMovement"
	case NSBeginFunctionKey:
		return "NSBeginFunctionKey"
	case NSBreakFunctionKey:
		return "NSBreakFunctionKey"
	case NSCancelTextMovement:
		return "NSCancelTextMovement"
	case NSCarriageReturnCharacter:
		return "NSCarriageReturnCharacter"
	case NSClearDisplayFunctionKey:
		return "NSClearDisplayFunctionKey"
	case NSClearLineFunctionKey:
		return "NSClearLineFunctionKey"
	case NSControlGlyph:
		return "NSControlGlyph"
	case NSDeleteCharFunctionKey:
		return "NSDeleteCharFunctionKey"
	case NSDeleteCharacter:
		return "NSDeleteCharacter"
	case NSDeleteFunctionKey:
		return "NSDeleteFunctionKey"
	case NSDeleteLineFunctionKey:
		return "NSDeleteLineFunctionKey"
	case NSDownArrowFunctionKey:
		return "NSDownArrowFunctionKey"
	case NSDownTextMovement:
		return "NSDownTextMovement"
	case NSEndFunctionKey:
		return "NSEndFunctionKey"
	case NSEnterCharacter:
		return "NSEnterCharacter"
	case NSExecuteFunctionKey:
		return "NSExecuteFunctionKey"
	case NSF10FunctionKey:
		return "NSF10FunctionKey"
	case NSF11FunctionKey:
		return "NSF11FunctionKey"
	case NSF12FunctionKey:
		return "NSF12FunctionKey"
	case NSF13FunctionKey:
		return "NSF13FunctionKey"
	case NSF14FunctionKey:
		return "NSF14FunctionKey"
	case NSF15FunctionKey:
		return "NSF15FunctionKey"
	case NSF16FunctionKey:
		return "NSF16FunctionKey"
	case NSF17FunctionKey:
		return "NSF17FunctionKey"
	case NSF18FunctionKey:
		return "NSF18FunctionKey"
	case NSF19FunctionKey:
		return "NSF19FunctionKey"
	case NSF1FunctionKey:
		return "NSF1FunctionKey"
	case NSF20FunctionKey:
		return "NSF20FunctionKey"
	case NSF21FunctionKey:
		return "NSF21FunctionKey"
	case NSF22FunctionKey:
		return "NSF22FunctionKey"
	case NSF23FunctionKey:
		return "NSF23FunctionKey"
	case NSF24FunctionKey:
		return "NSF24FunctionKey"
	case NSF25FunctionKey:
		return "NSF25FunctionKey"
	case NSF26FunctionKey:
		return "NSF26FunctionKey"
	case NSF27FunctionKey:
		return "NSF27FunctionKey"
	case NSF28FunctionKey:
		return "NSF28FunctionKey"
	case NSF29FunctionKey:
		return "NSF29FunctionKey"
	case NSF2FunctionKey:
		return "NSF2FunctionKey"
	case NSF30FunctionKey:
		return "NSF30FunctionKey"
	case NSF31FunctionKey:
		return "NSF31FunctionKey"
	case NSF32FunctionKey:
		return "NSF32FunctionKey"
	case NSF33FunctionKey:
		return "NSF33FunctionKey"
	case NSF34FunctionKey:
		return "NSF34FunctionKey"
	case NSF35FunctionKey:
		return "NSF35FunctionKey"
	case NSF3FunctionKey:
		return "NSF3FunctionKey"
	case NSF4FunctionKey:
		return "NSF4FunctionKey"
	case NSF5FunctionKey:
		return "NSF5FunctionKey"
	case NSF6FunctionKey:
		return "NSF6FunctionKey"
	case NSF7FunctionKey:
		return "NSF7FunctionKey"
	case NSF8FunctionKey:
		return "NSF8FunctionKey"
	case NSF9FunctionKey:
		return "NSF9FunctionKey"
	case NSFindFunctionKey:
		return "NSFindFunctionKey"
	case NSFontAssetDownloadError:
		return "NSFontAssetDownloadError"
	case NSFontErrorMaximum:
		return "NSFontErrorMaximum"
	case NSFormFeedCharacter:
		return "NSFormFeedCharacter"
	case NSHelpFunctionKey:
		return "NSHelpFunctionKey"
	case NSHomeFunctionKey:
		return "NSHomeFunctionKey"
	case NSIllegalTextMovement:
		return "NSIllegalTextMovement"
	case NSInsertCharFunctionKey:
		return "NSInsertCharFunctionKey"
	case NSInsertFunctionKey:
		return "NSInsertFunctionKey"
	case NSInsertLineFunctionKey:
		return "NSInsertLineFunctionKey"
	case NSLeftArrowFunctionKey:
		return "NSLeftArrowFunctionKey"
	case NSLeftTextMovement:
		return "NSLeftTextMovement"
	case NSLineSeparatorCharacter:
		return "NSLineSeparatorCharacter"
	case NSMenuFunctionKey:
		return "NSMenuFunctionKey"
	case NSModeSwitchFunctionKey:
		return "NSModeSwitchFunctionKey"
	case NSNewlineCharacter:
		return "NSNewlineCharacter"
	case NSNextFunctionKey:
		return "NSNextFunctionKey"
	case NSPageDownFunctionKey:
		return "NSPageDownFunctionKey"
	case NSPageUpFunctionKey:
		return "NSPageUpFunctionKey"
	case NSParagraphSeparatorCharacter:
		return "NSParagraphSeparatorCharacter"
	case NSPasteboardCommunicationError:
		return "NSPasteboardCommunicationError"
	case NSPasteboardContentsNotAvailableError:
		return "NSPasteboardContentsNotAvailableError"
	case NSPasteboardErrorMaximum:
		return "NSPasteboardErrorMaximum"
	case NSPasteboardErrorMinimum:
		return "NSPasteboardErrorMinimum"
	case NSPasteboardInvalidArgumentError:
		return "NSPasteboardInvalidArgumentError"
	case NSPauseFunctionKey:
		return "NSPauseFunctionKey"
	case NSPrevFunctionKey:
		return "NSPrevFunctionKey"
	case NSPrintFunctionKey:
		return "NSPrintFunctionKey"
	case NSPrintScreenFunctionKey:
		return "NSPrintScreenFunctionKey"
	case NSRedoFunctionKey:
		return "NSRedoFunctionKey"
	case NSResetFunctionKey:
		return "NSResetFunctionKey"
	case NSReturnTextMovement:
		return "NSReturnTextMovement"
	case NSRightArrowFunctionKey:
		return "NSRightArrowFunctionKey"
	case NSRightTextMovement:
		return "NSRightTextMovement"
	case NSScrollLockFunctionKey:
		return "NSScrollLockFunctionKey"
	case NSSelectFunctionKey:
		return "NSSelectFunctionKey"
	case NSServiceApplicationLaunchFailedError:
		return "NSServiceApplicationLaunchFailedError"
	case NSServiceApplicationNotFoundError:
		return "NSServiceApplicationNotFoundError"
	case NSServiceErrorMaximum:
		return "NSServiceErrorMaximum"
	case NSServiceInvalidPasteboardDataError:
		return "NSServiceInvalidPasteboardDataError"
	case NSServiceMalformedServiceDictionaryError:
		return "NSServiceMalformedServiceDictionaryError"
	case NSServiceMiscellaneousError:
		return "NSServiceMiscellaneousError"
	case NSServiceRequestTimedOutError:
		return "NSServiceRequestTimedOutError"
	case NSSharingServiceErrorMaximum:
		return "NSSharingServiceErrorMaximum"
	case NSSharingServiceErrorMinimum:
		return "NSSharingServiceErrorMinimum"
	case NSShowControlGlyphs:
		return "NSShowControlGlyphs"
	case NSShowInvisibleGlyphs:
		return "NSShowInvisibleGlyphs"
	case NSStopFunctionKey:
		return "NSStopFunctionKey"
	case NSSysReqFunctionKey:
		return "NSSysReqFunctionKey"
	case NSSystemFunctionKey:
		return "NSSystemFunctionKey"
	case NSTabCharacter:
		return "NSTabCharacter"
	case NSTabTextMovement:
		return "NSTabTextMovement"
	case NSTextReadInapplicableDocumentTypeError:
		return "NSTextReadInapplicableDocumentTypeError"
	case NSTextReadWriteErrorMaximum:
		return "NSTextReadWriteErrorMaximum"
	case NSTextReadWriteErrorMinimum:
		return "NSTextReadWriteErrorMinimum"
	case NSTextWriteInapplicableDocumentTypeError:
		return "NSTextWriteInapplicableDocumentTypeError"
	case NSUndoFunctionKey:
		return "NSUndoFunctionKey"
	case NSUpArrowFunctionKey:
		return "NSUpArrowFunctionKey"
	case NSUpTextMovement:
		return "NSUpTextMovement"
	case NSUserFunctionKey:
		return "NSUserFunctionKey"
	case NSWantsBidiLevels:
		return "NSWantsBidiLevels"
	case NSWindowSharingErrorMaximum:
		return "NSWindowSharingErrorMaximum"
	case NSWindowSharingErrorMinimum:
		return "NSWindowSharingErrorMinimum"
	case NSWindowSharingRequestNoEligibleSession:
		return "NSWindowSharingRequestNoEligibleSession"
	case NSWindowSharingRequestUnspecifiedError:
		return "NSWindowSharingRequestUnspecifiedError"
	case NSWorkspaceAuthorizationInvalidError:
		return "NSWorkspaceAuthorizationInvalidError"
	case NSWorkspaceErrorMaximum:
		return "NSWorkspaceErrorMaximum"
	case NSDoubleType:
		return "NSDoubleType"
	case NSPositiveDoubleType:
		return "NSPositiveDoubleType"
	default:
		return fmt.Sprintf("NS(%d)", e)
	}
}




const (
	NSAccessibilityHourMinuteDateTimeComponentsFlag uint = 12
	NSAccessibilityHourMinuteSecondDateTimeComponentsFlag uint = 14
	NSAccessibilityYearMonthDateTimeComponentsFlag uint = 192
	NSAccessibilityYearMonthDayDateTimeComponentsFlag uint = 224
)


// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityAnnotationPosition
type NSAccessibilityAnnotationPosition int

const (
	NSAccessibilityAnnotationPositionEnd NSAccessibilityAnnotationPosition = 2
	NSAccessibilityAnnotationPositionFullRange NSAccessibilityAnnotationPosition = 0
	NSAccessibilityAnnotationPositionStart NSAccessibilityAnnotationPosition = 1
)


func (e NSAccessibilityAnnotationPosition) String() string {
	switch e {
	case NSAccessibilityAnnotationPositionEnd:
		return "NSAccessibilityAnnotationPositionEnd"
	case NSAccessibilityAnnotationPositionFullRange:
		return "NSAccessibilityAnnotationPositionFullRange"
	case NSAccessibilityAnnotationPositionStart:
		return "NSAccessibilityAnnotationPositionStart"
	default:
		return fmt.Sprintf("NSAccessibilityAnnotationPosition(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/SearchDirection
type NSAccessibilityCustomRotorSearchDirection int

const (
	// NSAccessibilityCustomRotorSearchDirectionNext: The next search item.
	NSAccessibilityCustomRotorSearchDirectionNext NSAccessibilityCustomRotorSearchDirection = 1
	// NSAccessibilityCustomRotorSearchDirectionPrevious: The previous search item.
	NSAccessibilityCustomRotorSearchDirectionPrevious NSAccessibilityCustomRotorSearchDirection = 0
)


func (e NSAccessibilityCustomRotorSearchDirection) String() string {
	switch e {
	case NSAccessibilityCustomRotorSearchDirectionNext:
		return "NSAccessibilityCustomRotorSearchDirectionNext"
	case NSAccessibilityCustomRotorSearchDirectionPrevious:
		return "NSAccessibilityCustomRotorSearchDirectionPrevious"
	default:
		return fmt.Sprintf("NSAccessibilityCustomRotorSearchDirection(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/RotorType
type NSAccessibilityCustomRotorType int

const (
	// NSAccessibilityCustomRotorTypeAnnotation: An annotation.
	NSAccessibilityCustomRotorTypeAnnotation NSAccessibilityCustomRotorType = 2
	// NSAccessibilityCustomRotorTypeAny: Any type of item.
	NSAccessibilityCustomRotorTypeAny NSAccessibilityCustomRotorType = 1
	// NSAccessibilityCustomRotorTypeBoldText: Any bold text.
	NSAccessibilityCustomRotorTypeBoldText NSAccessibilityCustomRotorType = 3
	// NSAccessibilityCustomRotorTypeCustom: A rotor with a custom label.
	NSAccessibilityCustomRotorTypeCustom NSAccessibilityCustomRotorType = 0
	// NSAccessibilityCustomRotorTypeHeading: Any heading-level text.
	NSAccessibilityCustomRotorTypeHeading NSAccessibilityCustomRotorType = 4
	// NSAccessibilityCustomRotorTypeHeadingLevel1: A first-level heading.
	NSAccessibilityCustomRotorTypeHeadingLevel1 NSAccessibilityCustomRotorType = 5
	// NSAccessibilityCustomRotorTypeHeadingLevel2: A second-level heading.
	NSAccessibilityCustomRotorTypeHeadingLevel2 NSAccessibilityCustomRotorType = 6
	// NSAccessibilityCustomRotorTypeHeadingLevel3: A third-level heading.
	NSAccessibilityCustomRotorTypeHeadingLevel3 NSAccessibilityCustomRotorType = 7
	// NSAccessibilityCustomRotorTypeHeadingLevel4: A fourth-level heading.
	NSAccessibilityCustomRotorTypeHeadingLevel4 NSAccessibilityCustomRotorType = 8
	// NSAccessibilityCustomRotorTypeHeadingLevel5: A fifth-level heading.
	NSAccessibilityCustomRotorTypeHeadingLevel5 NSAccessibilityCustomRotorType = 9
	// NSAccessibilityCustomRotorTypeHeadingLevel6: A sixth-level heading.
	NSAccessibilityCustomRotorTypeHeadingLevel6 NSAccessibilityCustomRotorType = 10
	// NSAccessibilityCustomRotorTypeImage: An image.
	NSAccessibilityCustomRotorTypeImage NSAccessibilityCustomRotorType = 11
	// NSAccessibilityCustomRotorTypeItalicText: Any italicized text.
	NSAccessibilityCustomRotorTypeItalicText NSAccessibilityCustomRotorType = 12
	// NSAccessibilityCustomRotorTypeLandmark: A landmark.
	NSAccessibilityCustomRotorTypeLandmark NSAccessibilityCustomRotorType = 13
	// NSAccessibilityCustomRotorTypeLink: A link.
	NSAccessibilityCustomRotorTypeLink NSAccessibilityCustomRotorType = 14
	// NSAccessibilityCustomRotorTypeList: A list of items.
	NSAccessibilityCustomRotorTypeList NSAccessibilityCustomRotorType = 15
	// NSAccessibilityCustomRotorTypeMisspelledWord: A misspelled word.
	NSAccessibilityCustomRotorTypeMisspelledWord NSAccessibilityCustomRotorType = 16
	// NSAccessibilityCustomRotorTypeTable: A table of information.
	NSAccessibilityCustomRotorTypeTable NSAccessibilityCustomRotorType = 17
	// NSAccessibilityCustomRotorTypeTextField: A text field.
	NSAccessibilityCustomRotorTypeTextField NSAccessibilityCustomRotorType = 18
	// NSAccessibilityCustomRotorTypeUnderlinedText: Any underlined text.
	NSAccessibilityCustomRotorTypeUnderlinedText NSAccessibilityCustomRotorType = 19
	// NSAccessibilityCustomRotorTypeVisitedLink: A visited link.
	NSAccessibilityCustomRotorTypeVisitedLink NSAccessibilityCustomRotorType = 20
)


func (e NSAccessibilityCustomRotorType) String() string {
	switch e {
	case NSAccessibilityCustomRotorTypeAnnotation:
		return "NSAccessibilityCustomRotorTypeAnnotation"
	case NSAccessibilityCustomRotorTypeAny:
		return "NSAccessibilityCustomRotorTypeAny"
	case NSAccessibilityCustomRotorTypeBoldText:
		return "NSAccessibilityCustomRotorTypeBoldText"
	case NSAccessibilityCustomRotorTypeCustom:
		return "NSAccessibilityCustomRotorTypeCustom"
	case NSAccessibilityCustomRotorTypeHeading:
		return "NSAccessibilityCustomRotorTypeHeading"
	case NSAccessibilityCustomRotorTypeHeadingLevel1:
		return "NSAccessibilityCustomRotorTypeHeadingLevel1"
	case NSAccessibilityCustomRotorTypeHeadingLevel2:
		return "NSAccessibilityCustomRotorTypeHeadingLevel2"
	case NSAccessibilityCustomRotorTypeHeadingLevel3:
		return "NSAccessibilityCustomRotorTypeHeadingLevel3"
	case NSAccessibilityCustomRotorTypeHeadingLevel4:
		return "NSAccessibilityCustomRotorTypeHeadingLevel4"
	case NSAccessibilityCustomRotorTypeHeadingLevel5:
		return "NSAccessibilityCustomRotorTypeHeadingLevel5"
	case NSAccessibilityCustomRotorTypeHeadingLevel6:
		return "NSAccessibilityCustomRotorTypeHeadingLevel6"
	case NSAccessibilityCustomRotorTypeImage:
		return "NSAccessibilityCustomRotorTypeImage"
	case NSAccessibilityCustomRotorTypeItalicText:
		return "NSAccessibilityCustomRotorTypeItalicText"
	case NSAccessibilityCustomRotorTypeLandmark:
		return "NSAccessibilityCustomRotorTypeLandmark"
	case NSAccessibilityCustomRotorTypeLink:
		return "NSAccessibilityCustomRotorTypeLink"
	case NSAccessibilityCustomRotorTypeList:
		return "NSAccessibilityCustomRotorTypeList"
	case NSAccessibilityCustomRotorTypeMisspelledWord:
		return "NSAccessibilityCustomRotorTypeMisspelledWord"
	case NSAccessibilityCustomRotorTypeTable:
		return "NSAccessibilityCustomRotorTypeTable"
	case NSAccessibilityCustomRotorTypeTextField:
		return "NSAccessibilityCustomRotorTypeTextField"
	case NSAccessibilityCustomRotorTypeUnderlinedText:
		return "NSAccessibilityCustomRotorTypeUnderlinedText"
	case NSAccessibilityCustomRotorTypeVisitedLink:
		return "NSAccessibilityCustomRotorTypeVisitedLink"
	default:
		return fmt.Sprintf("NSAccessibilityCustomRotorType(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityOrientation
type NSAccessibilityOrientation int

const (
	// NSAccessibilityOrientationHorizontal: The element is oriented horizontally.
	NSAccessibilityOrientationHorizontal NSAccessibilityOrientation = 2
	// NSAccessibilityOrientationUnknown: The element has unknown orientation.
	NSAccessibilityOrientationUnknown NSAccessibilityOrientation = 0
	// NSAccessibilityOrientationVertical: The element is oriented vertically.
	NSAccessibilityOrientationVertical NSAccessibilityOrientation = 1
)


func (e NSAccessibilityOrientation) String() string {
	switch e {
	case NSAccessibilityOrientationHorizontal:
		return "NSAccessibilityOrientationHorizontal"
	case NSAccessibilityOrientationUnknown:
		return "NSAccessibilityOrientationUnknown"
	case NSAccessibilityOrientationVertical:
		return "NSAccessibilityOrientationVertical"
	default:
		return fmt.Sprintf("NSAccessibilityOrientation(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityPriorityLevel
type NSAccessibilityPriorityLevel int

const (
	// NSAccessibilityPriorityHigh: The notification is a high priority.
	NSAccessibilityPriorityHigh NSAccessibilityPriorityLevel = 90
	// NSAccessibilityPriorityLow: The notification is a low priority.
	NSAccessibilityPriorityLow NSAccessibilityPriorityLevel = 10
	// NSAccessibilityPriorityMedium: The notification is a medium priority.
	NSAccessibilityPriorityMedium NSAccessibilityPriorityLevel = 50
)


func (e NSAccessibilityPriorityLevel) String() string {
	switch e {
	case NSAccessibilityPriorityHigh:
		return "NSAccessibilityPriorityHigh"
	case NSAccessibilityPriorityLow:
		return "NSAccessibilityPriorityLow"
	case NSAccessibilityPriorityMedium:
		return "NSAccessibilityPriorityMedium"
	default:
		return fmt.Sprintf("NSAccessibilityPriorityLevel(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityRulerMarkerType
type NSAccessibilityRulerMarkerType int

const (
	// NSAccessibilityRulerMarkerTypeIndentFirstLine: First line indent marker.
	NSAccessibilityRulerMarkerTypeIndentFirstLine NSAccessibilityRulerMarkerType = 7
	// NSAccessibilityRulerMarkerTypeIndentHead: Head indent marker.
	NSAccessibilityRulerMarkerTypeIndentHead NSAccessibilityRulerMarkerType = 5
	// NSAccessibilityRulerMarkerTypeIndentTail: Tail indent marker.
	NSAccessibilityRulerMarkerTypeIndentTail NSAccessibilityRulerMarkerType = 6
	// NSAccessibilityRulerMarkerTypeTabStopCenter: Center tab stop.
	NSAccessibilityRulerMarkerTypeTabStopCenter NSAccessibilityRulerMarkerType = 3
	// NSAccessibilityRulerMarkerTypeTabStopDecimal: Decimal tab stop.
	NSAccessibilityRulerMarkerTypeTabStopDecimal NSAccessibilityRulerMarkerType = 4
	// NSAccessibilityRulerMarkerTypeTabStopLeft: Left tab stop.
	NSAccessibilityRulerMarkerTypeTabStopLeft NSAccessibilityRulerMarkerType = 1
	// NSAccessibilityRulerMarkerTypeTabStopRight: Right tab stop.
	NSAccessibilityRulerMarkerTypeTabStopRight NSAccessibilityRulerMarkerType = 2
	// NSAccessibilityRulerMarkerTypeUnknown: Unknown marker type.
	NSAccessibilityRulerMarkerTypeUnknown NSAccessibilityRulerMarkerType = 0
)


func (e NSAccessibilityRulerMarkerType) String() string {
	switch e {
	case NSAccessibilityRulerMarkerTypeIndentFirstLine:
		return "NSAccessibilityRulerMarkerTypeIndentFirstLine"
	case NSAccessibilityRulerMarkerTypeIndentHead:
		return "NSAccessibilityRulerMarkerTypeIndentHead"
	case NSAccessibilityRulerMarkerTypeIndentTail:
		return "NSAccessibilityRulerMarkerTypeIndentTail"
	case NSAccessibilityRulerMarkerTypeTabStopCenter:
		return "NSAccessibilityRulerMarkerTypeTabStopCenter"
	case NSAccessibilityRulerMarkerTypeTabStopDecimal:
		return "NSAccessibilityRulerMarkerTypeTabStopDecimal"
	case NSAccessibilityRulerMarkerTypeTabStopLeft:
		return "NSAccessibilityRulerMarkerTypeTabStopLeft"
	case NSAccessibilityRulerMarkerTypeTabStopRight:
		return "NSAccessibilityRulerMarkerTypeTabStopRight"
	case NSAccessibilityRulerMarkerTypeUnknown:
		return "NSAccessibilityRulerMarkerTypeUnknown"
	default:
		return fmt.Sprintf("NSAccessibilityRulerMarkerType(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySortDirection
type NSAccessibilitySortDirection int

const (
	// NSAccessibilitySortDirectionAscending: The column is sorted in ascending values.
	NSAccessibilitySortDirectionAscending NSAccessibilitySortDirection = 1
	// NSAccessibilitySortDirectionDescending: The column is sorted in descending values.
	NSAccessibilitySortDirectionDescending NSAccessibilitySortDirection = 2
	// NSAccessibilitySortDirectionUnknown: The sort direction is unknown.
	NSAccessibilitySortDirectionUnknown NSAccessibilitySortDirection = 0
)


func (e NSAccessibilitySortDirection) String() string {
	switch e {
	case NSAccessibilitySortDirectionAscending:
		return "NSAccessibilitySortDirectionAscending"
	case NSAccessibilitySortDirectionDescending:
		return "NSAccessibilitySortDirectionDescending"
	case NSAccessibilitySortDirectionUnknown:
		return "NSAccessibilitySortDirectionUnknown"
	default:
		return fmt.Sprintf("NSAccessibilitySortDirection(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityUnits
type NSAccessibilityUnits int

const (
	// NSAccessibilityUnitsCentimeters: The units are centimeters.
	NSAccessibilityUnitsCentimeters NSAccessibilityUnits = 2
	// NSAccessibilityUnitsInches: The units are inches.
	NSAccessibilityUnitsInches NSAccessibilityUnits = 1
	// NSAccessibilityUnitsPicas: The units are picas.
	NSAccessibilityUnitsPicas NSAccessibilityUnits = 4
	// NSAccessibilityUnitsPoints: The units are points.
	NSAccessibilityUnitsPoints NSAccessibilityUnits = 3
	// NSAccessibilityUnitsUnknown: The units are unknown.
	NSAccessibilityUnitsUnknown NSAccessibilityUnits = 0
)


func (e NSAccessibilityUnits) String() string {
	switch e {
	case NSAccessibilityUnitsCentimeters:
		return "NSAccessibilityUnitsCentimeters"
	case NSAccessibilityUnitsInches:
		return "NSAccessibilityUnitsInches"
	case NSAccessibilityUnitsPicas:
		return "NSAccessibilityUnitsPicas"
	case NSAccessibilityUnitsPoints:
		return "NSAccessibilityUnitsPoints"
	case NSAccessibilityUnitsUnknown:
		return "NSAccessibilityUnitsUnknown"
	default:
		return fmt.Sprintf("NSAccessibilityUnits(%d)", e)
	}
}




const (
	// Deprecated.
	NSAlertAlternateReturn int = 0
	// Deprecated.
	NSAlertDefaultReturn int = 1
	// Deprecated.
	NSAlertErrorReturn int = -2
	// Deprecated.
	NSAlertOtherReturn int = -1
)


// See: https://developer.apple.com/documentation/AppKit/NSAlert/Style
type NSAlertStyle int

const (
	// NSAlertStyleCritical: An alert style to inform someone about a critical event.
	NSAlertStyleCritical NSAlertStyle = 2
	// NSAlertStyleInformational: An alert style to inform someone about a current or impending event.
	NSAlertStyleInformational NSAlertStyle = 1
	// NSAlertStyleWarning: An alert style to warn someone about a current or impending event.
	NSAlertStyleWarning NSAlertStyle = 0
)


func (e NSAlertStyle) String() string {
	switch e {
	case NSAlertStyleCritical:
		return "NSAlertStyleCritical"
	case NSAlertStyleInformational:
		return "NSAlertStyleInformational"
	case NSAlertStyleWarning:
		return "NSAlertStyleWarning"
	default:
		return fmt.Sprintf("NSAlertStyle(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSAnimation/BlockingMode
type NSAnimationBlockingMode int

const (
	// NSAnimationBlocking: Requests the animation to run in the main thread in a custom run-loop mode that blocks user input.
	NSAnimationBlocking NSAnimationBlockingMode = 0
	// NSAnimationNonblocking: Requests the animation to run in a standard or specified run-loop mode that allows user input.
	NSAnimationNonblocking NSAnimationBlockingMode = 1
	// NSAnimationNonblockingThreaded: Requests the animation to run in a separate thread that is spawned by the  object.
	NSAnimationNonblockingThreaded NSAnimationBlockingMode = 2
)


func (e NSAnimationBlockingMode) String() string {
	switch e {
	case NSAnimationBlocking:
		return "NSAnimationBlocking"
	case NSAnimationNonblocking:
		return "NSAnimationNonblocking"
	case NSAnimationNonblockingThreaded:
		return "NSAnimationNonblockingThreaded"
	default:
		return fmt.Sprintf("NSAnimationBlockingMode(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSAnimation/Curve
type NSAnimationCurve int

const (
	// NSAnimationEaseIn: Describes an animation that slows down as it reaches the end.
	NSAnimationEaseIn NSAnimationCurve = 1
	// NSAnimationEaseInOut: Describes an S-curve in which the animation slowly speeds up and then slows down near the end of the animation.
	NSAnimationEaseInOut NSAnimationCurve = 0
	// NSAnimationEaseOut: Describes an animation that slowly speeds up from the start.
	NSAnimationEaseOut NSAnimationCurve = 2
	// NSAnimationLinear: Describes an animation in which there is no change in frame rate.
	NSAnimationLinear NSAnimationCurve = 3
)


func (e NSAnimationCurve) String() string {
	switch e {
	case NSAnimationEaseIn:
		return "NSAnimationEaseIn"
	case NSAnimationEaseInOut:
		return "NSAnimationEaseInOut"
	case NSAnimationEaseOut:
		return "NSAnimationEaseOut"
	case NSAnimationLinear:
		return "NSAnimationLinear"
	default:
		return fmt.Sprintf("NSAnimationCurve(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSAnimationEffect
type NSAnimationEffect int

const (
	// NSAnimationEffectDisappearingItemDefault: The default effect.
	NSAnimationEffectDisappearingItemDefault NSAnimationEffect = 0
	// NSAnimationEffectPoof: An effect showing a puff of smoke.
	NSAnimationEffectPoof NSAnimationEffect = 10
)


func (e NSAnimationEffect) String() string {
	switch e {
	case NSAnimationEffectDisappearingItemDefault:
		return "NSAnimationEffectDisappearingItemDefault"
	case NSAnimationEffectPoof:
		return "NSAnimationEffectPoof"
	default:
		return fmt.Sprintf("NSAnimationEffect(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSApplication/ActivationOptions
type NSApplicationActivationOptions int

const (
	// NSApplicationActivateAllWindows: By default, activation brings only the main and key windows forward.
	NSApplicationActivateAllWindows NSApplicationActivationOptions = 1
	// Deprecated.
	NSApplicationActivateIgnoringOtherApps NSApplicationActivationOptions = 2
)


func (e NSApplicationActivationOptions) String() string {
	switch e {
	case NSApplicationActivateAllWindows:
		return "NSApplicationActivateAllWindows"
	case NSApplicationActivateIgnoringOtherApps:
		return "NSApplicationActivateIgnoringOtherApps"
	default:
		return fmt.Sprintf("NSApplicationActivationOptions(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSApplication/ActivationPolicy-swift.enum
type NSApplicationActivationPolicy int

const (
	// NSApplicationActivationPolicyAccessory: The application doesn’t appear in the Dock and doesn’t have a menu bar, but it may be activated programmatically or by clicking on one of its windows.
	NSApplicationActivationPolicyAccessory NSApplicationActivationPolicy = 1
	// NSApplicationActivationPolicyProhibited: The application doesn’t appear in the Dock and may not create windows or be activated.
	NSApplicationActivationPolicyProhibited NSApplicationActivationPolicy = 2
	// NSApplicationActivationPolicyRegular: The application is an ordinary app that appears in the Dock and may have a user interface.
	NSApplicationActivationPolicyRegular NSApplicationActivationPolicy = 0
)


func (e NSApplicationActivationPolicy) String() string {
	switch e {
	case NSApplicationActivationPolicyAccessory:
		return "NSApplicationActivationPolicyAccessory"
	case NSApplicationActivationPolicyProhibited:
		return "NSApplicationActivationPolicyProhibited"
	case NSApplicationActivationPolicyRegular:
		return "NSApplicationActivationPolicyRegular"
	default:
		return fmt.Sprintf("NSApplicationActivationPolicy(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSApplication/DelegateReply
type NSApplicationDelegateReply int

const (
	// NSApplicationDelegateReplyCancel: Indicates the user cancelled the operation.
	NSApplicationDelegateReplyCancel NSApplicationDelegateReply = 1
	// NSApplicationDelegateReplyFailure: Indicates an error occurred processing the operation.
	NSApplicationDelegateReplyFailure NSApplicationDelegateReply = 2
	// NSApplicationDelegateReplySuccess: Indicates the operation succeeded.
	NSApplicationDelegateReplySuccess NSApplicationDelegateReply = 0
)


func (e NSApplicationDelegateReply) String() string {
	switch e {
	case NSApplicationDelegateReplyCancel:
		return "NSApplicationDelegateReplyCancel"
	case NSApplicationDelegateReplyFailure:
		return "NSApplicationDelegateReplyFailure"
	case NSApplicationDelegateReplySuccess:
		return "NSApplicationDelegateReplySuccess"
	default:
		return fmt.Sprintf("NSApplicationDelegateReply(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSApplication/PresentationOptions-swift.struct
type NSApplicationPresentationOptions int

const (
	// NSApplicationPresentationAutoHideDock: The dock is normally hidden, but automatically appears when moused near.
	NSApplicationPresentationAutoHideDock NSApplicationPresentationOptions = 1
	// NSApplicationPresentationAutoHideMenuBar: The menu bar is normally hidden, but automatically appears when moused near.
	NSApplicationPresentationAutoHideMenuBar NSApplicationPresentationOptions = 4
	// NSApplicationPresentationAutoHideToolbar: When in fullscreen mode the window toolbar is detached from window and hides and shows with autoHidden menuBar.
	NSApplicationPresentationAutoHideToolbar NSApplicationPresentationOptions = 2048
	// NSApplicationPresentationDisableAppleMenu: All Apple Menu items are disabled.
	NSApplicationPresentationDisableAppleMenu NSApplicationPresentationOptions = 16
	// NSApplicationPresentationDisableCursorLocationAssistance: The behavior that allows the user to shake the mouse to locate the cursor is disabled.
	NSApplicationPresentationDisableCursorLocationAssistance NSApplicationPresentationOptions = 4096
	// NSApplicationPresentationDisableForceQuit: The force quit panel (displayed by pressing Command + Option + Esc) is disabled
	NSApplicationPresentationDisableForceQuit NSApplicationPresentationOptions = 64
	// NSApplicationPresentationDisableHideApplication: The app’s “Hide” menu item is disabled.
	NSApplicationPresentationDisableHideApplication NSApplicationPresentationOptions = 256
	// NSApplicationPresentationDisableMenuBarTransparency: The menu bar transparency appearance is disabled.
	NSApplicationPresentationDisableMenuBarTransparency NSApplicationPresentationOptions = 512
	// NSApplicationPresentationDisableProcessSwitching: The process switching user interface (Command + Tab to cycle through apps) is disabled.
	NSApplicationPresentationDisableProcessSwitching NSApplicationPresentationOptions = 32
	// NSApplicationPresentationDisableSessionTermination: The panel that shows the Restart, Shut Down, and Log Out options that are displayed as a result of pushing the power key is disabled.
	NSApplicationPresentationDisableSessionTermination NSApplicationPresentationOptions = 128
	// NSApplicationPresentationFullScreen: The app is in fullscreen mode.
	NSApplicationPresentationFullScreen NSApplicationPresentationOptions = 1024
	// NSApplicationPresentationHideDock: The dock is entirely hidden and disabled.
	NSApplicationPresentationHideDock NSApplicationPresentationOptions = 2
	// NSApplicationPresentationHideMenuBar: The menu bar is entirely hidden and disabled.
	NSApplicationPresentationHideMenuBar NSApplicationPresentationOptions = 8
)


func (e NSApplicationPresentationOptions) String() string {
	switch e {
	case NSApplicationPresentationAutoHideDock:
		return "NSApplicationPresentationAutoHideDock"
	case NSApplicationPresentationAutoHideMenuBar:
		return "NSApplicationPresentationAutoHideMenuBar"
	case NSApplicationPresentationAutoHideToolbar:
		return "NSApplicationPresentationAutoHideToolbar"
	case NSApplicationPresentationDisableAppleMenu:
		return "NSApplicationPresentationDisableAppleMenu"
	case NSApplicationPresentationDisableCursorLocationAssistance:
		return "NSApplicationPresentationDisableCursorLocationAssistance"
	case NSApplicationPresentationDisableForceQuit:
		return "NSApplicationPresentationDisableForceQuit"
	case NSApplicationPresentationDisableHideApplication:
		return "NSApplicationPresentationDisableHideApplication"
	case NSApplicationPresentationDisableMenuBarTransparency:
		return "NSApplicationPresentationDisableMenuBarTransparency"
	case NSApplicationPresentationDisableProcessSwitching:
		return "NSApplicationPresentationDisableProcessSwitching"
	case NSApplicationPresentationDisableSessionTermination:
		return "NSApplicationPresentationDisableSessionTermination"
	case NSApplicationPresentationFullScreen:
		return "NSApplicationPresentationFullScreen"
	case NSApplicationPresentationHideDock:
		return "NSApplicationPresentationHideDock"
	case NSApplicationPresentationHideMenuBar:
		return "NSApplicationPresentationHideMenuBar"
	default:
		return fmt.Sprintf("NSApplicationPresentationOptions(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSApplication/PrintReply
type NSApplicationPrintReply int

const (
	// NSPrintingCancelled: Printing was cancelled.
	NSPrintingCancelled NSApplicationPrintReply = 0
	// NSPrintingFailure: Printing failed.
	NSPrintingFailure NSApplicationPrintReply = 3
	// NSPrintingSuccess: Printing was successful.
	NSPrintingSuccess NSApplicationPrintReply = 1
)


func (e NSApplicationPrintReply) String() string {
	switch e {
	case NSPrintingCancelled:
		return "NSPrintingCancelled"
	case NSPrintingFailure:
		return "NSPrintingFailure"
	case NSPrintingSuccess:
		return "NSPrintingSuccess"
	default:
		return fmt.Sprintf("NSApplicationPrintReply(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSApplication/TerminateReply
type NSApplicationTerminateReply int

const (
	// NSTerminateCancel: The app should not be terminated.
	NSTerminateCancel NSApplicationTerminateReply = 0
	// NSTerminateNow: It is OK to proceed with termination.
	NSTerminateNow NSApplicationTerminateReply = 1
)


func (e NSApplicationTerminateReply) String() string {
	switch e {
	case NSTerminateCancel:
		return "NSTerminateCancel"
	case NSTerminateNow:
		return "NSTerminateNow"
	default:
		return fmt.Sprintf("NSApplicationTerminateReply(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSView/BackgroundStyle
type NSBackgroundStyle int

const (
	// NSBackgroundStyleEmphasized: A style that adds emphasis to the background using an alternate color or visual effect.
	NSBackgroundStyleEmphasized NSBackgroundStyle = 1
	// NSBackgroundStyleLowered: A style that makes the background appear lower than the content drawn on it.
	NSBackgroundStyleLowered NSBackgroundStyle = 3
	// NSBackgroundStyleNormal: A style that reflects the predominant color scheme of the view’s appearance.
	NSBackgroundStyleNormal NSBackgroundStyle = 0
	// NSBackgroundStyleRaised: A style that makes the background appear higher than the content drawn on it.
	NSBackgroundStyleRaised NSBackgroundStyle = 2
)


func (e NSBackgroundStyle) String() string {
	switch e {
	case NSBackgroundStyleEmphasized:
		return "NSBackgroundStyleEmphasized"
	case NSBackgroundStyleLowered:
		return "NSBackgroundStyleLowered"
	case NSBackgroundStyleNormal:
		return "NSBackgroundStyleNormal"
	case NSBackgroundStyleRaised:
		return "NSBackgroundStyleRaised"
	default:
		return fmt.Sprintf("NSBackgroundStyle(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSWindow/BackingStoreType
type NSBackingStoreType int

const (
	// NSBackingStoreBuffered: The window renders all drawing into a display buffer and then flushes it to the screen.
	NSBackingStoreBuffered NSBackingStoreType = 2
	// NSBackingStoreNonretained: The window draws directly to the screen without using any buffer.
	NSBackingStoreNonretained NSBackingStoreType = 1
	// NSBackingStoreRetained: The window uses a buffer, but draws directly to the screen where possible and to the buffer for obscured portions.
	NSBackingStoreRetained NSBackingStoreType = 0
)


func (e NSBackingStoreType) String() string {
	switch e {
	case NSBackingStoreBuffered:
		return "NSBackingStoreBuffered"
	case NSBackingStoreNonretained:
		return "NSBackingStoreNonretained"
	case NSBackingStoreRetained:
		return "NSBackingStoreRetained"
	default:
		return fmt.Sprintf("NSBackingStoreType(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSButton/BezelStyle-swift.enum
type NSBezelStyle int

const (
	// NSBezelStyleAccessoryBar: A button style that’s typically used in the context of an accessory toolbar for buttons that narrow the focus of a search or other operation.
	NSBezelStyleAccessoryBar NSBezelStyle = 13
	// NSBezelStyleAccessoryBarAction: A button style that you use for extra actions in an accessory toolbar.
	NSBezelStyleAccessoryBarAction NSBezelStyle = 12
	// NSBezelStyleAutomatic: The default button style based on the button’s contents and position within the window.
	NSBezelStyleAutomatic NSBezelStyle = 0
	// NSBezelStyleBadge: A button style suitable for displaying additional information.
	NSBezelStyleBadge NSBezelStyle = 15
	// NSBezelStyleCircular: A round button that can contain either a single character or an icon.
	NSBezelStyleCircular NSBezelStyle = 7
	// NSBezelStyleDisclosure: A bezel style button for use with a disclosure triangle.
	NSBezelStyleDisclosure NSBezelStyle = 5
	// NSBezelStyleFlexiblePush: A push button with a flexible height to accommodate longer text labels or an image.
	NSBezelStyleFlexiblePush NSBezelStyle = 2
	// NSBezelStyleGlass: A bezel style with a glass effect
	NSBezelStyleGlass NSBezelStyle = 16
	// NSBezelStyleHelpButton: A round button with a question mark, providing the standard help button look.
	NSBezelStyleHelpButton NSBezelStyle = 9
	// NSBezelStylePush: A standard push style button.
	NSBezelStylePush NSBezelStyle = 1
	// NSBezelStylePushDisclosure: A bezel style push button with a disclosure triangle.
	NSBezelStylePushDisclosure NSBezelStyle = 14
	// NSBezelStyleSmallSquare: A simple square bezel style that can scale to any size.
	NSBezelStyleSmallSquare NSBezelStyle = 10
	// NSBezelStyleToolbar: A button style that’s appropriate for a toolbar item.
	NSBezelStyleToolbar NSBezelStyle = 11
)


func (e NSBezelStyle) String() string {
	switch e {
	case NSBezelStyleAccessoryBar:
		return "NSBezelStyleAccessoryBar"
	case NSBezelStyleAccessoryBarAction:
		return "NSBezelStyleAccessoryBarAction"
	case NSBezelStyleAutomatic:
		return "NSBezelStyleAutomatic"
	case NSBezelStyleBadge:
		return "NSBezelStyleBadge"
	case NSBezelStyleCircular:
		return "NSBezelStyleCircular"
	case NSBezelStyleDisclosure:
		return "NSBezelStyleDisclosure"
	case NSBezelStyleFlexiblePush:
		return "NSBezelStyleFlexiblePush"
	case NSBezelStyleGlass:
		return "NSBezelStyleGlass"
	case NSBezelStyleHelpButton:
		return "NSBezelStyleHelpButton"
	case NSBezelStylePush:
		return "NSBezelStylePush"
	case NSBezelStylePushDisclosure:
		return "NSBezelStylePushDisclosure"
	case NSBezelStyleSmallSquare:
		return "NSBezelStyleSmallSquare"
	case NSBezelStyleToolbar:
		return "NSBezelStyleToolbar"
	default:
		return fmt.Sprintf("NSBezelStyle(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/ElementType
type NSBezierPathElement int

const (
	// NSBezierPathElementClosePath: Marks the end of the current subpath at the specified point.
	NSBezierPathElementClosePath NSBezierPathElement = 3
	// NSBezierPathElementLineTo: Creates a straight line from the current drawing point to the specified point.
	NSBezierPathElementLineTo NSBezierPathElement = 1
	// NSBezierPathElementMoveTo: Moves the path object’s current drawing point to the specified point.
	NSBezierPathElementMoveTo NSBezierPathElement = 0
)


func (e NSBezierPathElement) String() string {
	switch e {
	case NSBezierPathElementClosePath:
		return "NSBezierPathElementClosePath"
	case NSBezierPathElementLineTo:
		return "NSBezierPathElementLineTo"
	case NSBezierPathElementMoveTo:
		return "NSBezierPathElementMoveTo"
	default:
		return fmt.Sprintf("NSBezierPathElement(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/Format
type NSBitmapFormat int

const (
	// NSBitmapFormatAlphaFirst: A format where the alpha value comes first.
	NSBitmapFormatAlphaFirst NSBitmapFormat = 1
	// NSBitmapFormatAlphaNonpremultiplied: A format where alpha values are not premultiplied.
	NSBitmapFormatAlphaNonpremultiplied NSBitmapFormat = 2
	// NSBitmapFormatFloatingPointSamples: A format where samples are specified using floating-point numbers.
	NSBitmapFormatFloatingPointSamples NSBitmapFormat = 4
	// NSBitmapFormatSixteenBitBigEndian: A 16-bit, big endian format.
	NSBitmapFormatSixteenBitBigEndian NSBitmapFormat = 1024
	// NSBitmapFormatSixteenBitLittleEndian: A 16-bit, little endian format.
	NSBitmapFormatSixteenBitLittleEndian NSBitmapFormat = 256
	// NSBitmapFormatThirtyTwoBitBigEndian: A 32-bit, big endian format.
	NSBitmapFormatThirtyTwoBitBigEndian NSBitmapFormat = 2048
	// NSBitmapFormatThirtyTwoBitLittleEndian: A 32-bit, little endian format.
	NSBitmapFormatThirtyTwoBitLittleEndian NSBitmapFormat = 512
)


func (e NSBitmapFormat) String() string {
	switch e {
	case NSBitmapFormatAlphaFirst:
		return "NSBitmapFormatAlphaFirst"
	case NSBitmapFormatAlphaNonpremultiplied:
		return "NSBitmapFormatAlphaNonpremultiplied"
	case NSBitmapFormatFloatingPointSamples:
		return "NSBitmapFormatFloatingPointSamples"
	case NSBitmapFormatSixteenBitBigEndian:
		return "NSBitmapFormatSixteenBitBigEndian"
	case NSBitmapFormatSixteenBitLittleEndian:
		return "NSBitmapFormatSixteenBitLittleEndian"
	case NSBitmapFormatThirtyTwoBitBigEndian:
		return "NSBitmapFormatThirtyTwoBitBigEndian"
	case NSBitmapFormatThirtyTwoBitLittleEndian:
		return "NSBitmapFormatThirtyTwoBitLittleEndian"
	default:
		return fmt.Sprintf("NSBitmapFormat(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/FileType
type NSBitmapImageFileType int

const (
	// NSBitmapImageFileTypeBMP: Windows bitmap image (BMP) format.
	NSBitmapImageFileTypeBMP NSBitmapImageFileType = 1
	// NSBitmapImageFileTypeGIF: Graphics Image Format (GIF), originally created by CompuServe for online downloads.
	NSBitmapImageFileTypeGIF NSBitmapImageFileType = 2
	// NSBitmapImageFileTypeJPEG: JPEG format.
	NSBitmapImageFileTypeJPEG NSBitmapImageFileType = 3
	// NSBitmapImageFileTypeJPEG2000: JPEG 2000 file format.
	NSBitmapImageFileTypeJPEG2000 NSBitmapImageFileType = 5
	// NSBitmapImageFileTypePNG: Portable Network Graphics (PNG) format.
	NSBitmapImageFileTypePNG NSBitmapImageFileType = 4
	// NSBitmapImageFileTypeTIFF: Tagged Image File Format (TIFF).
	NSBitmapImageFileTypeTIFF NSBitmapImageFileType = 0
)


func (e NSBitmapImageFileType) String() string {
	switch e {
	case NSBitmapImageFileTypeBMP:
		return "NSBitmapImageFileTypeBMP"
	case NSBitmapImageFileTypeGIF:
		return "NSBitmapImageFileTypeGIF"
	case NSBitmapImageFileTypeJPEG:
		return "NSBitmapImageFileTypeJPEG"
	case NSBitmapImageFileTypeJPEG2000:
		return "NSBitmapImageFileTypeJPEG2000"
	case NSBitmapImageFileTypePNG:
		return "NSBitmapImageFileTypePNG"
	case NSBitmapImageFileTypeTIFF:
		return "NSBitmapImageFileTypeTIFF"
	default:
		return fmt.Sprintf("NSBitmapImageFileType(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSBorderType
type NSBorderType int

const (
	// NSBezelBorder: A concave border that makes the view look sunken.
	NSBezelBorder NSBorderType = 2
	// NSGrooveBorder: A thin border that looks etched around the image.
	NSGrooveBorder NSBorderType = 3
	// NSLineBorder: A black line border around the view.
	NSLineBorder NSBorderType = 1
	// NSNoBorder: No border.
	NSNoBorder NSBorderType = 0
)


func (e NSBorderType) String() string {
	switch e {
	case NSBezelBorder:
		return "NSBezelBorder"
	case NSGrooveBorder:
		return "NSGrooveBorder"
	case NSLineBorder:
		return "NSLineBorder"
	case NSNoBorder:
		return "NSNoBorder"
	default:
		return fmt.Sprintf("NSBorderType(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSBox/BoxType-swift.enum
type NSBoxType int

const (
	NSBoxCustom NSBoxType = 4
	NSBoxPrimary NSBoxType = 0
	NSBoxSeparator NSBoxType = 2
)


func (e NSBoxType) String() string {
	switch e {
	case NSBoxCustom:
		return "NSBoxCustom"
	case NSBoxPrimary:
		return "NSBoxPrimary"
	case NSBoxSeparator:
		return "NSBoxSeparator"
	default:
		return fmt.Sprintf("NSBoxType(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSBrowser/ColumnResizingType-swift.enum
type NSBrowserColumnResizingType int

const (
	NSBrowserAutoColumnResizing NSBrowserColumnResizingType = 1
	NSBrowserNoColumnResizing NSBrowserColumnResizingType = 0
	NSBrowserUserColumnResizing NSBrowserColumnResizingType = 2
)


func (e NSBrowserColumnResizingType) String() string {
	switch e {
	case NSBrowserAutoColumnResizing:
		return "NSBrowserAutoColumnResizing"
	case NSBrowserNoColumnResizing:
		return "NSBrowserNoColumnResizing"
	case NSBrowserUserColumnResizing:
		return "NSBrowserUserColumnResizing"
	default:
		return fmt.Sprintf("NSBrowserColumnResizingType(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSBrowser/DropOperation
type NSBrowserDropOperation int

const (
	NSBrowserDropAbove NSBrowserDropOperation = 1
	NSBrowserDropOn NSBrowserDropOperation = 0
)


func (e NSBrowserDropOperation) String() string {
	switch e {
	case NSBrowserDropAbove:
		return "NSBrowserDropAbove"
	case NSBrowserDropOn:
		return "NSBrowserDropOn"
	default:
		return fmt.Sprintf("NSBrowserDropOperation(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSButton/ButtonType
type NSButtonType int

const (
	// NSButtonTypeAccelerator: A button that sends repeating actions as pressure changes occur.
	NSButtonTypeAccelerator NSButtonType = 8
	// NSButtonTypeMomentaryChange: A button that displays its alternate content when clicked and returns to its normal content when the user releases it.
	NSButtonTypeMomentaryChange NSButtonType = 5
	// NSButtonTypeMomentaryLight: A button that displays a highlight when the user clicks it and returns to its normal state when the user releases it.
	NSButtonTypeMomentaryLight NSButtonType = 0
	// NSButtonTypeMomentaryPushIn: A button that illuminates when the user clicks it.
	NSButtonTypeMomentaryPushIn NSButtonType = 7
	// NSButtonTypeMultiLevelAccelerator: A button that allows for a configurable number of stepped pressure levels and provides tactile feedback as the user reaches each step.
	NSButtonTypeMultiLevelAccelerator NSButtonType = 9
	// NSButtonTypeOnOff: A button that switches between a normal and emphasized bezel on each click.
	NSButtonTypeOnOff NSButtonType = 6
	// NSButtonTypePushOnPushOff: A button that switches between on and off states with each click.
	NSButtonTypePushOnPushOff NSButtonType = 1
	// NSButtonTypeRadio: A button that displays a single selected value from group of possible choices.
	NSButtonTypeRadio NSButtonType = 4
	// NSButtonTypeSwitch: A standard checkbox button.
	NSButtonTypeSwitch NSButtonType = 3
	// NSButtonTypeToggle: A button that switches between its normal and alternate content on each click.
	NSButtonTypeToggle NSButtonType = 2
)


func (e NSButtonType) String() string {
	switch e {
	case NSButtonTypeAccelerator:
		return "NSButtonTypeAccelerator"
	case NSButtonTypeMomentaryChange:
		return "NSButtonTypeMomentaryChange"
	case NSButtonTypeMomentaryLight:
		return "NSButtonTypeMomentaryLight"
	case NSButtonTypeMomentaryPushIn:
		return "NSButtonTypeMomentaryPushIn"
	case NSButtonTypeMultiLevelAccelerator:
		return "NSButtonTypeMultiLevelAccelerator"
	case NSButtonTypeOnOff:
		return "NSButtonTypeOnOff"
	case NSButtonTypePushOnPushOff:
		return "NSButtonTypePushOnPushOff"
	case NSButtonTypeRadio:
		return "NSButtonTypeRadio"
	case NSButtonTypeSwitch:
		return "NSButtonTypeSwitch"
	case NSButtonTypeToggle:
		return "NSButtonTypeToggle"
	default:
		return fmt.Sprintf("NSButtonType(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSCell/Attribute
type NSCellAttribute int

const (
	// NSCellAllowsMixedState: Lets the cell’s state be , as well as  and .
	NSCellAllowsMixedState NSCellAttribute = 16
	// NSCellChangesContents: If the cell’s state is  or , displays the cell’s alternate image.
	NSCellChangesContents NSCellAttribute = 14
	// NSCellDisabled: Does not let the user manipulate the cell.
	NSCellDisabled NSCellAttribute = 0
	// NSCellEditable: Lets the user edit the cell’s contents.
	NSCellEditable NSCellAttribute = 3
	// NSCellHasImageHorizontal: Controls the position of the cell’s image: places the image on the right of any text in the cell.
	NSCellHasImageHorizontal NSCellAttribute = 12
	// NSCellHasImageOnLeftOrBottom: Controls the position of the cell’s image: places the image on the left of or below any text in the cell.
	NSCellHasImageOnLeftOrBottom NSCellAttribute = 13
	// NSCellHasOverlappingImage: Controls the position of the cell’s image: places the image over any text in the cell.
	NSCellHasOverlappingImage NSCellAttribute = 11
	// NSCellHighlighted: Draws the cell with a highlighted appearance.
	NSCellHighlighted NSCellAttribute = 5
	// NSCellIsBordered: Draws a border around the cell.
	NSCellIsBordered NSCellAttribute = 10
	// NSCellIsInsetButton: Insets the cell’s contents from the border.
	NSCellIsInsetButton NSCellAttribute = 15
	// NSCellLightsByBackground: If the cell is pushed in, changes the cell’s background color from gray to white.
	NSCellLightsByBackground NSCellAttribute = 9
	// NSCellLightsByContents: If the cell is pushed in, displays the cell’s alternate image.
	NSCellLightsByContents NSCellAttribute = 6
	// NSCellLightsByGray: If the cell is pushed in, displays the cell’s image as darkened.
	NSCellLightsByGray NSCellAttribute = 7
	// NSCellState: The cell’s state.
	NSCellState NSCellAttribute = 1
	// NSChangeBackgroundCell: If the cell’s state is  or , changes the cell’s background color from gray to white.
	NSChangeBackgroundCell NSCellAttribute = 8
	// NSChangeGrayCell: If the cell’s state is  or , displays the cell’s image as darkened.
	NSChangeGrayCell NSCellAttribute = 4
	// NSPushInCell: Determines whether the cell’s image and text appear to be shifted down and to the right.
	NSPushInCell NSCellAttribute = 2
)


func (e NSCellAttribute) String() string {
	switch e {
	case NSCellAllowsMixedState:
		return "NSCellAllowsMixedState"
	case NSCellChangesContents:
		return "NSCellChangesContents"
	case NSCellDisabled:
		return "NSCellDisabled"
	case NSCellEditable:
		return "NSCellEditable"
	case NSCellHasImageHorizontal:
		return "NSCellHasImageHorizontal"
	case NSCellHasImageOnLeftOrBottom:
		return "NSCellHasImageOnLeftOrBottom"
	case NSCellHasOverlappingImage:
		return "NSCellHasOverlappingImage"
	case NSCellHighlighted:
		return "NSCellHighlighted"
	case NSCellIsBordered:
		return "NSCellIsBordered"
	case NSCellIsInsetButton:
		return "NSCellIsInsetButton"
	case NSCellLightsByBackground:
		return "NSCellLightsByBackground"
	case NSCellLightsByContents:
		return "NSCellLightsByContents"
	case NSCellLightsByGray:
		return "NSCellLightsByGray"
	case NSCellState:
		return "NSCellState"
	case NSChangeBackgroundCell:
		return "NSChangeBackgroundCell"
	case NSChangeGrayCell:
		return "NSChangeGrayCell"
	case NSPushInCell:
		return "NSPushInCell"
	default:
		return fmt.Sprintf("NSCellAttribute(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSCell/HitResult
type NSCellHitResult int

const (
	// NSCellHitContentArea: A content area in the cell.
	NSCellHitContentArea NSCellHitResult = 1
	// NSCellHitEditableTextArea: An editable text area of the cell.
	NSCellHitEditableTextArea NSCellHitResult = 2
	// NSCellHitTrackableArea: A trackable area in the cell.
	NSCellHitTrackableArea NSCellHitResult = 4
)


func (e NSCellHitResult) String() string {
	switch e {
	case NSCellHitContentArea:
		return "NSCellHitContentArea"
	case NSCellHitEditableTextArea:
		return "NSCellHitEditableTextArea"
	case NSCellHitTrackableArea:
		return "NSCellHitTrackableArea"
	default:
		return fmt.Sprintf("NSCellHitResult(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSControl/ImagePosition
type NSCellImagePosition int

const (
	// NSImageAbove: The image is above the title.
	NSImageAbove NSCellImagePosition = 5
	// NSImageBelow: The image is below the title.
	NSImageBelow NSCellImagePosition = 4
	// NSImageLeading: The image is on the title’s leading edge.
	NSImageLeading NSCellImagePosition = 7
	// NSImageLeft: The image is to the left of the title.
	NSImageLeft NSCellImagePosition = 2
	// NSImageOnly: The cell displays an image but not a title.
	NSImageOnly NSCellImagePosition = 1
	// NSImageOverlaps: The image overlaps the title.
	NSImageOverlaps NSCellImagePosition = 6
	// NSImageRight: The image is to the right of the title.
	NSImageRight NSCellImagePosition = 3
	// NSImageTrailing: The image is on the title’s trailing edge.
	NSImageTrailing NSCellImagePosition = 8
	// NSNoImage: The cell doesn’t display an image.
	NSNoImage NSCellImagePosition = 0
)


func (e NSCellImagePosition) String() string {
	switch e {
	case NSImageAbove:
		return "NSImageAbove"
	case NSImageBelow:
		return "NSImageBelow"
	case NSImageLeading:
		return "NSImageLeading"
	case NSImageLeft:
		return "NSImageLeft"
	case NSImageOnly:
		return "NSImageOnly"
	case NSImageOverlaps:
		return "NSImageOverlaps"
	case NSImageRight:
		return "NSImageRight"
	case NSImageTrailing:
		return "NSImageTrailing"
	case NSNoImage:
		return "NSNoImage"
	default:
		return fmt.Sprintf("NSCellImagePosition(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSCell/StyleMask
type NSCellStyleMask int

const (
	// NSChangeBackgroundCellMask: Same as [NSChangeGrayCellMask], but only background pixels are changed.
	NSChangeBackgroundCellMask NSCellStyleMask = 8
	// NSChangeGrayCellMask: The button cell swaps the “control color” (the controlColor method of [NSColor]) and white pixels on its background and icon.
	NSChangeGrayCellMask NSCellStyleMask = 4
	// NSContentsCellMask: The button cell displays its alternate icon and/or title.
	NSContentsCellMask NSCellStyleMask = 1
	// NSPushInCellMask: The button cell “pushes in” if it has a border.
	NSPushInCellMask NSCellStyleMask = 2
)


func (e NSCellStyleMask) String() string {
	switch e {
	case NSChangeBackgroundCellMask:
		return "NSChangeBackgroundCellMask"
	case NSChangeGrayCellMask:
		return "NSChangeGrayCellMask"
	case NSContentsCellMask:
		return "NSContentsCellMask"
	case NSPushInCellMask:
		return "NSPushInCellMask"
	default:
		return fmt.Sprintf("NSCellStyleMask(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSCell/CellType
type NSCellType int

const (
	// NSImageCellType: Cell displays images.
	NSImageCellType NSCellType = 2
	// NSNullCellType: Cell displays nothing.
	NSNullCellType NSCellType = 0
	// NSTextCellType: Cell displays text.
	NSTextCellType NSCellType = 1
)


func (e NSCellType) String() string {
	switch e {
	case NSImageCellType:
		return "NSImageCellType"
	case NSNullCellType:
		return "NSNullCellType"
	case NSTextCellType:
		return "NSTextCellType"
	default:
		return fmt.Sprintf("NSCellType(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSCharacterCollection
type NSCharacterCollection int

const (
	// NSAdobeCNS1CharacterCollection: Indicates the Adobe-CNS1 mapping.
	NSAdobeCNS1CharacterCollection NSCharacterCollection = 1
	// NSAdobeGB1CharacterCollection: Indicates the Adobe-GB1 mapping.
	NSAdobeGB1CharacterCollection NSCharacterCollection = 2
	// NSAdobeJapan1CharacterCollection: Indicates the Adobe-Japan1 mapping.
	NSAdobeJapan1CharacterCollection NSCharacterCollection = 3
	// NSAdobeJapan2CharacterCollection: Indicates the Adobe-Japan2 mapping.
	NSAdobeJapan2CharacterCollection NSCharacterCollection = 4
	// NSAdobeKorea1CharacterCollection: Indicates the Adobe-Korea1 mapping.
	NSAdobeKorea1CharacterCollection NSCharacterCollection = 5
	// NSIdentityMappingCharacterCollection: Indicates that the character identifier is equal to the glyph index.
	NSIdentityMappingCharacterCollection NSCharacterCollection = 0
)


func (e NSCharacterCollection) String() string {
	switch e {
	case NSAdobeCNS1CharacterCollection:
		return "NSAdobeCNS1CharacterCollection"
	case NSAdobeGB1CharacterCollection:
		return "NSAdobeGB1CharacterCollection"
	case NSAdobeJapan1CharacterCollection:
		return "NSAdobeJapan1CharacterCollection"
	case NSAdobeJapan2CharacterCollection:
		return "NSAdobeJapan2CharacterCollection"
	case NSAdobeKorea1CharacterCollection:
		return "NSAdobeKorea1CharacterCollection"
	case NSIdentityMappingCharacterCollection:
		return "NSIdentityMappingCharacterCollection"
	default:
		return fmt.Sprintf("NSCharacterCollection(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSSharingService/CloudKitOptions
type NSCloudKitSharingServiceOptions int

const (
	// NSCloudKitSharingServiceAllowPrivate: An option that allows the participant to privately distribute the share to other iCloud users.
	NSCloudKitSharingServiceAllowPrivate NSCloudKitSharingServiceOptions = 2
	// NSCloudKitSharingServiceAllowPublic: An option that allows the participant to publicly distribute the share to other iCloud users.
	NSCloudKitSharingServiceAllowPublic NSCloudKitSharingServiceOptions = 1
	// NSCloudKitSharingServiceAllowReadOnly: An option that allows the participant to grant other participants read-only permissions.
	NSCloudKitSharingServiceAllowReadOnly NSCloudKitSharingServiceOptions = 16
	// NSCloudKitSharingServiceAllowReadWrite: An option that allows the participant to grant other participants read-write permissions.
	NSCloudKitSharingServiceAllowReadWrite NSCloudKitSharingServiceOptions = 32
	// NSCloudKitSharingServiceStandard: An option that allows the participant to configure the share with a standard set of options.
	NSCloudKitSharingServiceStandard NSCloudKitSharingServiceOptions = 0
)


func (e NSCloudKitSharingServiceOptions) String() string {
	switch e {
	case NSCloudKitSharingServiceAllowPrivate:
		return "NSCloudKitSharingServiceAllowPrivate"
	case NSCloudKitSharingServiceAllowPublic:
		return "NSCloudKitSharingServiceAllowPublic"
	case NSCloudKitSharingServiceAllowReadOnly:
		return "NSCloudKitSharingServiceAllowReadOnly"
	case NSCloudKitSharingServiceAllowReadWrite:
		return "NSCloudKitSharingServiceAllowReadWrite"
	case NSCloudKitSharingServiceStandard:
		return "NSCloudKitSharingServiceStandard"
	default:
		return fmt.Sprintf("NSCloudKitSharingServiceOptions(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSCollectionElementCategory
type NSCollectionElementCategory int

const (
	// NSCollectionElementCategoryDecorationView: The element is a decoration view.
	NSCollectionElementCategoryDecorationView NSCollectionElementCategory = 2
	// NSCollectionElementCategoryInterItemGap: The element is an inter-item gap.
	NSCollectionElementCategoryInterItemGap NSCollectionElementCategory = 3
	// NSCollectionElementCategoryItem: The element is an item.
	NSCollectionElementCategoryItem NSCollectionElementCategory = 0
	// NSCollectionElementCategorySupplementaryView: The element is a supplementary view.
	NSCollectionElementCategorySupplementaryView NSCollectionElementCategory = 1
)


func (e NSCollectionElementCategory) String() string {
	switch e {
	case NSCollectionElementCategoryDecorationView:
		return "NSCollectionElementCategoryDecorationView"
	case NSCollectionElementCategoryInterItemGap:
		return "NSCollectionElementCategoryInterItemGap"
	case NSCollectionElementCategoryItem:
		return "NSCollectionElementCategoryItem"
	case NSCollectionElementCategorySupplementaryView:
		return "NSCollectionElementCategorySupplementaryView"
	default:
		return fmt.Sprintf("NSCollectionElementCategory(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSectionOrthogonalScrollingBehavior
type NSCollectionLayoutSectionOrthogonalScrollingBehavior int

const (
	// NSCollectionLayoutSectionOrthogonalScrollingBehaviorContinuous: The section allows users to scroll its content orthogonally with continuous scrolling.
	NSCollectionLayoutSectionOrthogonalScrollingBehaviorContinuous NSCollectionLayoutSectionOrthogonalScrollingBehavior = 1
	// NSCollectionLayoutSectionOrthogonalScrollingBehaviorContinuousGroupLeadingBoundary: The section allows users to scroll its content orthogonally, coming to a natural stop at the leading boundary of the visible group.
	NSCollectionLayoutSectionOrthogonalScrollingBehaviorContinuousGroupLeadingBoundary NSCollectionLayoutSectionOrthogonalScrollingBehavior = 2
	// NSCollectionLayoutSectionOrthogonalScrollingBehaviorGroupPaging: The section allows users to page its content orthogonally one group at a time.
	NSCollectionLayoutSectionOrthogonalScrollingBehaviorGroupPaging NSCollectionLayoutSectionOrthogonalScrollingBehavior = 4
	// NSCollectionLayoutSectionOrthogonalScrollingBehaviorGroupPagingCentered: The section allows users to page its content orthogonally one group at a time, centering each group.
	NSCollectionLayoutSectionOrthogonalScrollingBehaviorGroupPagingCentered NSCollectionLayoutSectionOrthogonalScrollingBehavior = 5
	// NSCollectionLayoutSectionOrthogonalScrollingBehaviorNone: The section does not allow users to scroll its content orthogonally.
	NSCollectionLayoutSectionOrthogonalScrollingBehaviorNone NSCollectionLayoutSectionOrthogonalScrollingBehavior = 0
	// NSCollectionLayoutSectionOrthogonalScrollingBehaviorPaging: The section allows users to page its content orthogonally.
	NSCollectionLayoutSectionOrthogonalScrollingBehaviorPaging NSCollectionLayoutSectionOrthogonalScrollingBehavior = 3
)


func (e NSCollectionLayoutSectionOrthogonalScrollingBehavior) String() string {
	switch e {
	case NSCollectionLayoutSectionOrthogonalScrollingBehaviorContinuous:
		return "NSCollectionLayoutSectionOrthogonalScrollingBehaviorContinuous"
	case NSCollectionLayoutSectionOrthogonalScrollingBehaviorContinuousGroupLeadingBoundary:
		return "NSCollectionLayoutSectionOrthogonalScrollingBehaviorContinuousGroupLeadingBoundary"
	case NSCollectionLayoutSectionOrthogonalScrollingBehaviorGroupPaging:
		return "NSCollectionLayoutSectionOrthogonalScrollingBehaviorGroupPaging"
	case NSCollectionLayoutSectionOrthogonalScrollingBehaviorGroupPagingCentered:
		return "NSCollectionLayoutSectionOrthogonalScrollingBehaviorGroupPagingCentered"
	case NSCollectionLayoutSectionOrthogonalScrollingBehaviorNone:
		return "NSCollectionLayoutSectionOrthogonalScrollingBehaviorNone"
	case NSCollectionLayoutSectionOrthogonalScrollingBehaviorPaging:
		return "NSCollectionLayoutSectionOrthogonalScrollingBehaviorPaging"
	default:
		return fmt.Sprintf("NSCollectionLayoutSectionOrthogonalScrollingBehavior(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/UpdateAction
type NSCollectionUpdateAction int

const (
	// NSCollectionUpdateActionDelete: Remove the action from the collection view.
	NSCollectionUpdateActionDelete NSCollectionUpdateAction = 1
	// NSCollectionUpdateActionInsert: Insert the item into the collection view.
	NSCollectionUpdateActionInsert NSCollectionUpdateAction = 0
	// NSCollectionUpdateActionMove: Move the item from its current location to a new location.
	NSCollectionUpdateActionMove NSCollectionUpdateAction = 3
	// NSCollectionUpdateActionNone: Take no action on the item.
	NSCollectionUpdateActionNone NSCollectionUpdateAction = 4
	// NSCollectionUpdateActionReload: Reload the item, which consists of deleting and then inserting the item.
	NSCollectionUpdateActionReload NSCollectionUpdateAction = 2
)


func (e NSCollectionUpdateAction) String() string {
	switch e {
	case NSCollectionUpdateActionDelete:
		return "NSCollectionUpdateActionDelete"
	case NSCollectionUpdateActionInsert:
		return "NSCollectionUpdateActionInsert"
	case NSCollectionUpdateActionMove:
		return "NSCollectionUpdateActionMove"
	case NSCollectionUpdateActionNone:
		return "NSCollectionUpdateActionNone"
	case NSCollectionUpdateActionReload:
		return "NSCollectionUpdateActionReload"
	default:
		return fmt.Sprintf("NSCollectionUpdateAction(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/DropOperation
type NSCollectionViewDropOperation int

const (
	// NSCollectionViewDropBefore: The drop occurs before the collection view item to which the item was dragged.
	NSCollectionViewDropBefore NSCollectionViewDropOperation = 1
	// NSCollectionViewDropOn: The drop occurs at the collection view item to which the item was dragged.
	NSCollectionViewDropOn NSCollectionViewDropOperation = 0
)


func (e NSCollectionViewDropOperation) String() string {
	switch e {
	case NSCollectionViewDropBefore:
		return "NSCollectionViewDropBefore"
	case NSCollectionViewDropOn:
		return "NSCollectionViewDropOn"
	default:
		return fmt.Sprintf("NSCollectionViewDropOperation(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewItem/HighlightState-swift.enum
type NSCollectionViewItemHighlightState int

const (
	// NSCollectionViewItemHighlightAsDropTarget: The drop target highlight state.
	NSCollectionViewItemHighlightAsDropTarget NSCollectionViewItemHighlightState = 3
	// NSCollectionViewItemHighlightForDeselection: The deselection highlight state.
	NSCollectionViewItemHighlightForDeselection NSCollectionViewItemHighlightState = 2
	// NSCollectionViewItemHighlightForSelection: The selected highlight state.
	NSCollectionViewItemHighlightForSelection NSCollectionViewItemHighlightState = 1
	// NSCollectionViewItemHighlightNone: No highlight state.
	NSCollectionViewItemHighlightNone NSCollectionViewItemHighlightState = 0
)


func (e NSCollectionViewItemHighlightState) String() string {
	switch e {
	case NSCollectionViewItemHighlightAsDropTarget:
		return "NSCollectionViewItemHighlightAsDropTarget"
	case NSCollectionViewItemHighlightForDeselection:
		return "NSCollectionViewItemHighlightForDeselection"
	case NSCollectionViewItemHighlightForSelection:
		return "NSCollectionViewItemHighlightForSelection"
	case NSCollectionViewItemHighlightNone:
		return "NSCollectionViewItemHighlightNone"
	default:
		return fmt.Sprintf("NSCollectionViewItemHighlightState(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/ScrollDirection
type NSCollectionViewScrollDirection int

const (
	// NSCollectionViewScrollDirectionHorizontal: The layout scrolls content horizontally.
	NSCollectionViewScrollDirectionHorizontal NSCollectionViewScrollDirection = 1
	// NSCollectionViewScrollDirectionVertical: The layout scrolls content vertically.
	NSCollectionViewScrollDirectionVertical NSCollectionViewScrollDirection = 0
)


func (e NSCollectionViewScrollDirection) String() string {
	switch e {
	case NSCollectionViewScrollDirectionHorizontal:
		return "NSCollectionViewScrollDirectionHorizontal"
	case NSCollectionViewScrollDirectionVertical:
		return "NSCollectionViewScrollDirectionVertical"
	default:
		return fmt.Sprintf("NSCollectionViewScrollDirection(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/ScrollPosition
type NSCollectionViewScrollPosition int

const (
	// NSCollectionViewScrollPositionBottom: Scroll so that the bottom edge of the bounding box is adjacent to the bottom of the collection view’s bounds.
	NSCollectionViewScrollPositionBottom NSCollectionViewScrollPosition = 4
	// NSCollectionViewScrollPositionCenteredHorizontally: Scroll so that the selected items’ bounding box is centered horizontally in the collection view’s bounds.
	NSCollectionViewScrollPositionCenteredHorizontally NSCollectionViewScrollPosition = 16
	// NSCollectionViewScrollPositionCenteredVertically: Scroll so that the bounding box of the selected items is centered vertically in the collection view’s bounds.
	NSCollectionViewScrollPositionCenteredVertically NSCollectionViewScrollPosition = 2
	// NSCollectionViewScrollPositionLeadingEdge: Scroll so that the leading edge of the selected items’ bounding box is adjacent to the leading edge of the collection view’s bounds.
	NSCollectionViewScrollPositionLeadingEdge NSCollectionViewScrollPosition = 64
	// NSCollectionViewScrollPositionLeft: Scroll so that the left edge of the selected items’ bounding box is adjacent to the left edge of the collection view’s bounds.
	NSCollectionViewScrollPositionLeft NSCollectionViewScrollPosition = 8
	// NSCollectionViewScrollPositionNearestHorizontalEdge: Scroll so that the bounding box is adjacent to the nearest edge (top or bottom) of the collection view.
	NSCollectionViewScrollPositionNearestHorizontalEdge NSCollectionViewScrollPosition = 512
	// NSCollectionViewScrollPositionNearestVerticalEdge: Scroll so that the bounding box is adjacent to the nearest edge (leading or trailing) of the collection view.
	NSCollectionViewScrollPositionNearestVerticalEdge NSCollectionViewScrollPosition = 256
	// NSCollectionViewScrollPositionRight: Scroll so that the right edge of the selected items’ bounding box is adjacent to the right edge of the collection view’s bounds.
	NSCollectionViewScrollPositionRight NSCollectionViewScrollPosition = 32
	// NSCollectionViewScrollPositionTop: Scroll so that the top edge of the selected items’ bounding box is adjacent to the top edge of the collection view’s bounds.
	NSCollectionViewScrollPositionTop NSCollectionViewScrollPosition = 1
	// NSCollectionViewScrollPositionTrailingEdge: Scroll so that the trailing edge of the selected items’ bounding box is adjacent to the trailing edge of the collection view’s bounds.
	NSCollectionViewScrollPositionTrailingEdge NSCollectionViewScrollPosition = 128
)


func (e NSCollectionViewScrollPosition) String() string {
	switch e {
	case NSCollectionViewScrollPositionBottom:
		return "NSCollectionViewScrollPositionBottom"
	case NSCollectionViewScrollPositionCenteredHorizontally:
		return "NSCollectionViewScrollPositionCenteredHorizontally"
	case NSCollectionViewScrollPositionCenteredVertically:
		return "NSCollectionViewScrollPositionCenteredVertically"
	case NSCollectionViewScrollPositionLeadingEdge:
		return "NSCollectionViewScrollPositionLeadingEdge"
	case NSCollectionViewScrollPositionLeft:
		return "NSCollectionViewScrollPositionLeft"
	case NSCollectionViewScrollPositionNearestHorizontalEdge:
		return "NSCollectionViewScrollPositionNearestHorizontalEdge"
	case NSCollectionViewScrollPositionNearestVerticalEdge:
		return "NSCollectionViewScrollPositionNearestVerticalEdge"
	case NSCollectionViewScrollPositionRight:
		return "NSCollectionViewScrollPositionRight"
	case NSCollectionViewScrollPositionTop:
		return "NSCollectionViewScrollPositionTop"
	case NSCollectionViewScrollPositionTrailingEdge:
		return "NSCollectionViewScrollPositionTrailingEdge"
	default:
		return fmt.Sprintf("NSCollectionViewScrollPosition(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSColorPanel/Mode-swift.enum
type NSColorPanelMode int

const (
	// NSColorPanelModeCMYK: The cyan-magenta-yellow-black color mode.
	NSColorPanelModeCMYK NSColorPanelMode = 2
	// NSColorPanelModeColorList: The custom color list mode.
	NSColorPanelModeColorList NSColorPanelMode = 5
	// NSColorPanelModeCrayon: The crayon picker mode.
	NSColorPanelModeCrayon NSColorPanelMode = 7
	// NSColorPanelModeCustomPalette: The custom palette color mode.
	NSColorPanelModeCustomPalette NSColorPanelMode = 4
	// NSColorPanelModeGray: The grayscale-alpha color mode.
	NSColorPanelModeGray NSColorPanelMode = 0
	// NSColorPanelModeHSB: The hue-saturation-brightness color mode.
	NSColorPanelModeHSB NSColorPanelMode = 3
	// NSColorPanelModeNone: No color panel mode.
	NSColorPanelModeNone NSColorPanelMode = -1
	// NSColorPanelModeRGB: The red-green-blue color mode.
	NSColorPanelModeRGB NSColorPanelMode = 1
	// NSColorPanelModeWheel: The color wheel mode.
	NSColorPanelModeWheel NSColorPanelMode = 6
)


func (e NSColorPanelMode) String() string {
	switch e {
	case NSColorPanelModeCMYK:
		return "NSColorPanelModeCMYK"
	case NSColorPanelModeColorList:
		return "NSColorPanelModeColorList"
	case NSColorPanelModeCrayon:
		return "NSColorPanelModeCrayon"
	case NSColorPanelModeCustomPalette:
		return "NSColorPanelModeCustomPalette"
	case NSColorPanelModeGray:
		return "NSColorPanelModeGray"
	case NSColorPanelModeHSB:
		return "NSColorPanelModeHSB"
	case NSColorPanelModeNone:
		return "NSColorPanelModeNone"
	case NSColorPanelModeRGB:
		return "NSColorPanelModeRGB"
	case NSColorPanelModeWheel:
		return "NSColorPanelModeWheel"
	default:
		return fmt.Sprintf("NSColorPanelMode(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSColorPanel/Options
type NSColorPanelOptions int

const (
	// NSColorPanelAllModesMask: All color modes.
	NSColorPanelAllModesMask NSColorPanelOptions = 65535
	// NSColorPanelCMYKModeMask: The cyan-magenta-yellow-black color mode.
	NSColorPanelCMYKModeMask NSColorPanelOptions = 4
	// NSColorPanelColorListModeMask: The custom color list mode.
	NSColorPanelColorListModeMask NSColorPanelOptions = 32
	// NSColorPanelCrayonModeMask: The crayon color mode.
	NSColorPanelCrayonModeMask NSColorPanelOptions = 128
	// NSColorPanelCustomPaletteModeMask: The custom palette color mode.
	NSColorPanelCustomPaletteModeMask NSColorPanelOptions = 16
	// NSColorPanelGrayModeMask: The grayscale-alpha color mode.
	NSColorPanelGrayModeMask NSColorPanelOptions = 1
	// NSColorPanelHSBModeMask: The hue-saturation-brightness color mode.
	NSColorPanelHSBModeMask NSColorPanelOptions = 8
	// NSColorPanelRGBModeMask: The red-green-blue color mode.
	NSColorPanelRGBModeMask NSColorPanelOptions = 2
	// NSColorPanelWheelModeMask: The color wheel mode.
	NSColorPanelWheelModeMask NSColorPanelOptions = 64
)


func (e NSColorPanelOptions) String() string {
	switch e {
	case NSColorPanelAllModesMask:
		return "NSColorPanelAllModesMask"
	case NSColorPanelCMYKModeMask:
		return "NSColorPanelCMYKModeMask"
	case NSColorPanelColorListModeMask:
		return "NSColorPanelColorListModeMask"
	case NSColorPanelCrayonModeMask:
		return "NSColorPanelCrayonModeMask"
	case NSColorPanelCustomPaletteModeMask:
		return "NSColorPanelCustomPaletteModeMask"
	case NSColorPanelGrayModeMask:
		return "NSColorPanelGrayModeMask"
	case NSColorPanelHSBModeMask:
		return "NSColorPanelHSBModeMask"
	case NSColorPanelRGBModeMask:
		return "NSColorPanelRGBModeMask"
	case NSColorPanelWheelModeMask:
		return "NSColorPanelWheelModeMask"
	default:
		return fmt.Sprintf("NSColorPanelOptions(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSColorRenderingIntent
type NSColorRenderingIntent int

const (
	// NSColorRenderingIntentAbsoluteColorimetric: Map colors outside of the gamut of the output device to the closest possible match inside the gamut of the output device.
	NSColorRenderingIntentAbsoluteColorimetric NSColorRenderingIntent = 1
	// NSColorRenderingIntentDefault: Use the default rendering intent for the graphics context.
	NSColorRenderingIntentDefault NSColorRenderingIntent = 0
	// NSColorRenderingIntentPerceptual: Preserve the visual relationship between colors by compressing the gamut of the graphics context to fit inside the gamut of the output device.
	NSColorRenderingIntentPerceptual NSColorRenderingIntent = 3
	// NSColorRenderingIntentRelativeColorimetric: Map colors outside of the gamut of the output device to the closest possible match inside the gamut of the output device.
	NSColorRenderingIntentRelativeColorimetric NSColorRenderingIntent = 2
	// NSColorRenderingIntentSaturation: Preserve the relative saturation value of the colors when converting into the gamut of the output device.
	NSColorRenderingIntentSaturation NSColorRenderingIntent = 4
)


func (e NSColorRenderingIntent) String() string {
	switch e {
	case NSColorRenderingIntentAbsoluteColorimetric:
		return "NSColorRenderingIntentAbsoluteColorimetric"
	case NSColorRenderingIntentDefault:
		return "NSColorRenderingIntentDefault"
	case NSColorRenderingIntentPerceptual:
		return "NSColorRenderingIntentPerceptual"
	case NSColorRenderingIntentRelativeColorimetric:
		return "NSColorRenderingIntentRelativeColorimetric"
	case NSColorRenderingIntentSaturation:
		return "NSColorRenderingIntentSaturation"
	default:
		return fmt.Sprintf("NSColorRenderingIntent(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSColorSpace/Model
type NSColorSpaceModel int

const (
	// NSColorSpaceModelCMYK: The CMYK (cyan-magenta-yellow-black) color-space model.
	NSColorSpaceModelCMYK NSColorSpaceModel = 2
	// NSColorSpaceModelDeviceN: The DeviceN color-space model from Adobe Systems, Inc.
	NSColorSpaceModelDeviceN NSColorSpaceModel = 4
	// NSColorSpaceModelGray: The grayscale color-space model.
	NSColorSpaceModelGray NSColorSpaceModel = 0
	// NSColorSpaceModelIndexed: An indexed color space, which identifies discrete colors in a color list by index number.
	NSColorSpaceModelIndexed NSColorSpaceModel = 5
	// NSColorSpaceModelLAB: The L*a*b* device-independent color-space model, which represents colors relative to a reference white point.
	NSColorSpaceModelLAB NSColorSpaceModel = 3
	// NSColorSpaceModelPatterned: A pattern color space, which is a repeated image that creates a tiled pattern.
	NSColorSpaceModelPatterned NSColorSpaceModel = 6
	// NSColorSpaceModelRGB: The RGB (red-green-blue) color-space model.
	NSColorSpaceModelRGB NSColorSpaceModel = 1
	// NSColorSpaceModelUnknown: An unknown color-space model.
	NSColorSpaceModelUnknown NSColorSpaceModel = -1
)


func (e NSColorSpaceModel) String() string {
	switch e {
	case NSColorSpaceModelCMYK:
		return "NSColorSpaceModelCMYK"
	case NSColorSpaceModelDeviceN:
		return "NSColorSpaceModelDeviceN"
	case NSColorSpaceModelGray:
		return "NSColorSpaceModelGray"
	case NSColorSpaceModelIndexed:
		return "NSColorSpaceModelIndexed"
	case NSColorSpaceModelLAB:
		return "NSColorSpaceModelLAB"
	case NSColorSpaceModelPatterned:
		return "NSColorSpaceModelPatterned"
	case NSColorSpaceModelRGB:
		return "NSColorSpaceModelRGB"
	case NSColorSpaceModelUnknown:
		return "NSColorSpaceModelUnknown"
	default:
		return fmt.Sprintf("NSColorSpaceModel(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSColor/SystemEffect
type NSColorSystemEffect int

const (
	// NSColorSystemEffectDeepPressed: The color that indicates the item received a deep press.
	NSColorSystemEffectDeepPressed NSColorSystemEffect = 2
	// NSColorSystemEffectDisabled: The color that indicates the item is disabled.
	NSColorSystemEffectDisabled NSColorSystemEffect = 3
	// NSColorSystemEffectNone: No additional effects.
	NSColorSystemEffectNone NSColorSystemEffect = 0
	// NSColorSystemEffectPressed: The color that indicates the item was pressed.
	NSColorSystemEffectPressed NSColorSystemEffect = 1
	// NSColorSystemEffectRollover: The color that indicates the mouse rolled over the item.
	NSColorSystemEffectRollover NSColorSystemEffect = 4
)


func (e NSColorSystemEffect) String() string {
	switch e {
	case NSColorSystemEffectDeepPressed:
		return "NSColorSystemEffectDeepPressed"
	case NSColorSystemEffectDisabled:
		return "NSColorSystemEffectDisabled"
	case NSColorSystemEffectNone:
		return "NSColorSystemEffectNone"
	case NSColorSystemEffectPressed:
		return "NSColorSystemEffectPressed"
	case NSColorSystemEffectRollover:
		return "NSColorSystemEffectRollover"
	default:
		return fmt.Sprintf("NSColorSystemEffect(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSColor/ColorType
type NSColorType int

const (
	// NSColorTypeCatalog: Colors that are retrieved from an asset catalog.
	NSColorTypeCatalog NSColorType = 2
	// NSColorTypeComponentBased: Colors that include floating-point color components and a color space.
	NSColorTypeComponentBased NSColorType = 0
	// NSColorTypePattern: Colors that include an image to be used as a pattern.
	NSColorTypePattern NSColorType = 1
)


func (e NSColorType) String() string {
	switch e {
	case NSColorTypeCatalog:
		return "NSColorTypeCatalog"
	case NSColorTypeComponentBased:
		return "NSColorTypeComponentBased"
	case NSColorTypePattern:
		return "NSColorTypePattern"
	default:
		return fmt.Sprintf("NSColorType(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSColorWell/Style
type NSColorWellStyle int

const (
	// NSColorWellStyleDefault: The default style for color wells.
	NSColorWellStyleDefault NSColorWellStyle = 0
	// NSColorWellStyleExpanded: A style that supports a color picker popover for fast interactions, and adds a dedicated button to display the color panel.
	NSColorWellStyleExpanded NSColorWellStyle = 2
	// NSColorWellStyleMinimal: A style that adds minimal adornments to the color well.
	NSColorWellStyleMinimal NSColorWellStyle = 1
)


func (e NSColorWellStyle) String() string {
	switch e {
	case NSColorWellStyleDefault:
		return "NSColorWellStyleDefault"
	case NSColorWellStyleExpanded:
		return "NSColorWellStyleExpanded"
	case NSColorWellStyleMinimal:
		return "NSColorWellStyleMinimal"
	default:
		return fmt.Sprintf("NSColorWellStyle(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSComboButton/Style-swift.enum
type NSComboButtonStyle int

const (
	// NSComboButtonStyleSplit: A style that separates the button’s title and image from the menu indicator people use to activate the button.
	NSComboButtonStyleSplit NSComboButtonStyle = 0
	// NSComboButtonStyleUnified: A style that unifies the button’s title and image with the menu indicator.
	NSComboButtonStyleUnified NSComboButtonStyle = 1
)


func (e NSComboButtonStyle) String() string {
	switch e {
	case NSComboButtonStyleSplit:
		return "NSComboButtonStyleSplit"
	case NSComboButtonStyleUnified:
		return "NSComboButtonStyleUnified"
	default:
		return fmt.Sprintf("NSComboButtonStyle(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSCompositingOperation
type NSCompositingOperation int

const (
	// NSCompositingOperationClear: Transparency everywhere.
	NSCompositingOperationClear NSCompositingOperation = 0
	// NSCompositingOperationColor: Uses the hue and saturation of the source and the luminosity of the destination.
	NSCompositingOperationColor NSCompositingOperation = 27
	// NSCompositingOperationColorBurn: Darkens the destination color to reflect the source.
	NSCompositingOperationColorBurn NSCompositingOperation = 20
	// NSCompositingOperationColorDodge: Brightens the destination to reflect the source.
	NSCompositingOperationColorDodge NSCompositingOperation = 19
	// NSCompositingOperationCopy: The source image.
	NSCompositingOperationCopy NSCompositingOperation = 1
	// NSCompositingOperationDarken: Use the darker of the source and destination colors.
	NSCompositingOperationDarken NSCompositingOperation = 17
	// NSCompositingOperationDestinationAtop: The destination image wherever both images are opaque, the source image wherever it is opaque and the destination image is transparent, and transparent elsehwere.
	NSCompositingOperationDestinationAtop NSCompositingOperation = 9
	// NSCompositingOperationDestinationIn: The destination image wherever both images are opaque, and transparent elsewhere.
	NSCompositingOperationDestinationIn NSCompositingOperation = 7
	// NSCompositingOperationDestinationOut: The destination image wherever it is opaque and the source image is transparent, and transparent elsewhere.
	NSCompositingOperationDestinationOut NSCompositingOperation = 8
	// NSCompositingOperationDestinationOver: The destination image wherever it is opaque, and the source image elsewhere.
	NSCompositingOperationDestinationOver NSCompositingOperation = 6
	// NSCompositingOperationDifference: Subtracts the darker value from the lighter value.
	NSCompositingOperationDifference NSCompositingOperation = 23
	// NSCompositingOperationExclusion: Subtracts the darker value from the lighter value, except lower in contrast.
	NSCompositingOperationExclusion NSCompositingOperation = 24
	// NSCompositingOperationHardLight: Multiplies or screens colors, with the effect of shining a spotlight on the destination.
	NSCompositingOperationHardLight NSCompositingOperation = 22
	// NSCompositingOperationHue: Uses the hue of the source and the saturation and luminosity of the destination.
	NSCompositingOperationHue NSCompositingOperation = 25
	// NSCompositingOperationLighten: Use the lighter of the source and destination colors.
	NSCompositingOperationLighten NSCompositingOperation = 18
	// NSCompositingOperationLuminosity: Uses the luminosity of the source and the hue and saturation of the destination.
	NSCompositingOperationLuminosity NSCompositingOperation = 28
	// NSCompositingOperationMultiply: The source color is multiplied by the destination color.
	NSCompositingOperationMultiply NSCompositingOperation = 14
	// NSCompositingOperationOverlay: Source colors overlay the destination.
	NSCompositingOperationOverlay NSCompositingOperation = 16
	// NSCompositingOperationPlusDarker: The sum of the source and destination images, with color values approach 0 as a limit.
	NSCompositingOperationPlusDarker NSCompositingOperation = 11
	// NSCompositingOperationPlusLighter: The sum of the source and destination images, with color values approach 1 as a limit.
	NSCompositingOperationPlusLighter NSCompositingOperation = 13
	// NSCompositingOperationSaturation: Uses the saturation value of the source and the hue and luminosity of the destination.
	NSCompositingOperationSaturation NSCompositingOperation = 26
	// NSCompositingOperationScreen: Multiplies the complement of the destination and source color values, and then complements the result.
	NSCompositingOperationScreen NSCompositingOperation = 15
	// NSCompositingOperationSoftLight: Darkens or lightens colors, with the effect of shining a diffused spotlight on the destination.
	NSCompositingOperationSoftLight NSCompositingOperation = 21
	// NSCompositingOperationSourceAtop: The source image wherever both images are opaque, the destination image wherever it is opaque but the source image is transparent, and transparent elsewhere
	NSCompositingOperationSourceAtop NSCompositingOperation = 5
	// NSCompositingOperationSourceIn: The source image wherever both images are opaque, and transparent elsewhere.
	NSCompositingOperationSourceIn NSCompositingOperation = 3
	// NSCompositingOperationSourceOut: The source image wherever it is opaque and the destination image is transparent, and transparent elsewhere.
	NSCompositingOperationSourceOut NSCompositingOperation = 4
	// NSCompositingOperationSourceOver: The source image wherever it is opaque, and the destination image elsewhere.
	NSCompositingOperationSourceOver NSCompositingOperation = 2
	// NSCompositingOperationXOR: Exclusive OR of the source and destination images.
	NSCompositingOperationXOR NSCompositingOperation = 10
)


func (e NSCompositingOperation) String() string {
	switch e {
	case NSCompositingOperationClear:
		return "NSCompositingOperationClear"
	case NSCompositingOperationColor:
		return "NSCompositingOperationColor"
	case NSCompositingOperationColorBurn:
		return "NSCompositingOperationColorBurn"
	case NSCompositingOperationColorDodge:
		return "NSCompositingOperationColorDodge"
	case NSCompositingOperationCopy:
		return "NSCompositingOperationCopy"
	case NSCompositingOperationDarken:
		return "NSCompositingOperationDarken"
	case NSCompositingOperationDestinationAtop:
		return "NSCompositingOperationDestinationAtop"
	case NSCompositingOperationDestinationIn:
		return "NSCompositingOperationDestinationIn"
	case NSCompositingOperationDestinationOut:
		return "NSCompositingOperationDestinationOut"
	case NSCompositingOperationDestinationOver:
		return "NSCompositingOperationDestinationOver"
	case NSCompositingOperationDifference:
		return "NSCompositingOperationDifference"
	case NSCompositingOperationExclusion:
		return "NSCompositingOperationExclusion"
	case NSCompositingOperationHardLight:
		return "NSCompositingOperationHardLight"
	case NSCompositingOperationHue:
		return "NSCompositingOperationHue"
	case NSCompositingOperationLighten:
		return "NSCompositingOperationLighten"
	case NSCompositingOperationLuminosity:
		return "NSCompositingOperationLuminosity"
	case NSCompositingOperationMultiply:
		return "NSCompositingOperationMultiply"
	case NSCompositingOperationOverlay:
		return "NSCompositingOperationOverlay"
	case NSCompositingOperationPlusDarker:
		return "NSCompositingOperationPlusDarker"
	case NSCompositingOperationPlusLighter:
		return "NSCompositingOperationPlusLighter"
	case NSCompositingOperationSaturation:
		return "NSCompositingOperationSaturation"
	case NSCompositingOperationScreen:
		return "NSCompositingOperationScreen"
	case NSCompositingOperationSoftLight:
		return "NSCompositingOperationSoftLight"
	case NSCompositingOperationSourceAtop:
		return "NSCompositingOperationSourceAtop"
	case NSCompositingOperationSourceIn:
		return "NSCompositingOperationSourceIn"
	case NSCompositingOperationSourceOut:
		return "NSCompositingOperationSourceOut"
	case NSCompositingOperationSourceOver:
		return "NSCompositingOperationSourceOver"
	case NSCompositingOperationXOR:
		return "NSCompositingOperationXOR"
	default:
		return fmt.Sprintf("NSCompositingOperation(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSControl/BorderShape
type NSControlBorderShape int

const (
	// NSControlBorderShapeAutomatic: The control will resolve this to an appropriate shape for the given control size and context
	NSControlBorderShapeAutomatic NSControlBorderShape = 0
	// NSControlBorderShapeCapsule: The control will resolve this to an appropriate shape for the given control size and context
	NSControlBorderShapeCapsule NSControlBorderShape = 1
	// NSControlBorderShapeCircle: The control will resolve this to an appropriate shape for the given control size and context
	NSControlBorderShapeCircle NSControlBorderShape = 3
	// NSControlBorderShapeRoundedRectangle: The control will resolve this to an appropriate shape for the given control size and context
	NSControlBorderShapeRoundedRectangle NSControlBorderShape = 2
)


func (e NSControlBorderShape) String() string {
	switch e {
	case NSControlBorderShapeAutomatic:
		return "NSControlBorderShapeAutomatic"
	case NSControlBorderShapeCapsule:
		return "NSControlBorderShapeCapsule"
	case NSControlBorderShapeCircle:
		return "NSControlBorderShapeCircle"
	case NSControlBorderShapeRoundedRectangle:
		return "NSControlBorderShapeRoundedRectangle"
	default:
		return fmt.Sprintf("NSControlBorderShape(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/ControlCharacterAction
type NSControlCharacterAction int

const (
	// NSControlCharacterActionContainerBreak: An action that triggers a break in layout for the current container.
	NSControlCharacterActionContainerBreak NSControlCharacterAction = 32
	// NSControlCharacterActionHorizontalTab: An action that inserts a horizontal tab.
	NSControlCharacterActionHorizontalTab NSControlCharacterAction = 4
	// NSControlCharacterActionLineBreak: An action that causes a line break.
	NSControlCharacterActionLineBreak NSControlCharacterAction = 8
	// NSControlCharacterActionParagraphBreak: An action that causes a paragraph break.
	NSControlCharacterActionParagraphBreak NSControlCharacterAction = 16
	// NSControlCharacterActionWhitespace: An action that adds whitespace.
	NSControlCharacterActionWhitespace NSControlCharacterAction = 2
	// NSControlCharacterActionZeroAdvancement: An action that removes the glyph from layout.
	NSControlCharacterActionZeroAdvancement NSControlCharacterAction = 1
)


func (e NSControlCharacterAction) String() string {
	switch e {
	case NSControlCharacterActionContainerBreak:
		return "NSControlCharacterActionContainerBreak"
	case NSControlCharacterActionHorizontalTab:
		return "NSControlCharacterActionHorizontalTab"
	case NSControlCharacterActionLineBreak:
		return "NSControlCharacterActionLineBreak"
	case NSControlCharacterActionParagraphBreak:
		return "NSControlCharacterActionParagraphBreak"
	case NSControlCharacterActionWhitespace:
		return "NSControlCharacterActionWhitespace"
	case NSControlCharacterActionZeroAdvancement:
		return "NSControlCharacterActionZeroAdvancement"
	default:
		return fmt.Sprintf("NSControlCharacterAction(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSControl/ControlSize-swift.enum
type NSControlSize int

const (
	// NSControlSizeLarge: A size larger than the default control size.
	NSControlSizeLarge NSControlSize = 3
	// NSControlSizeMini: The smallest control size.
	NSControlSizeMini NSControlSize = 2
	// NSControlSizeRegular: The default control size.
	NSControlSizeRegular NSControlSize = 0
	// NSControlSizeSmall: A size smaller than the default control size.
	NSControlSizeSmall NSControlSize = 1
)


func (e NSControlSize) String() string {
	switch e {
	case NSControlSizeLarge:
		return "NSControlSizeLarge"
	case NSControlSizeMini:
		return "NSControlSizeMini"
	case NSControlSizeRegular:
		return "NSControlSizeRegular"
	case NSControlSizeSmall:
		return "NSControlSizeSmall"
	default:
		return fmt.Sprintf("NSControlSize(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSControlTint
type NSControlTint int

const (
	// NSBlueControlTint: Aqua control tint.
	NSBlueControlTint NSControlTint = 1
	// NSClearControlTint: Clear control tint.
	NSClearControlTint NSControlTint = 7
	// NSDefaultControlTint: The current default tint setting.
	NSDefaultControlTint NSControlTint = 0
	// NSGraphiteControlTint: Graphite control tint.
	NSGraphiteControlTint NSControlTint = 6
)


func (e NSControlTint) String() string {
	switch e {
	case NSBlueControlTint:
		return "NSBlueControlTint"
	case NSClearControlTint:
		return "NSClearControlTint"
	case NSDefaultControlTint:
		return "NSDefaultControlTint"
	case NSGraphiteControlTint:
		return "NSGraphiteControlTint"
	default:
		return fmt.Sprintf("NSControlTint(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/CorrectionIndicatorType
type NSCorrectionIndicatorType int

const (
	// NSCorrectionIndicatorTypeDefault: The default indicator that shows a proposed correction.
	NSCorrectionIndicatorTypeDefault NSCorrectionIndicatorType = 0
	// NSCorrectionIndicatorTypeGuesses: Shows multiple alternatives from which the user may choose the appropriate spelling.
	NSCorrectionIndicatorTypeGuesses NSCorrectionIndicatorType = 2
	// NSCorrectionIndicatorTypeReversion: Provides the option to revert to the original form after a correction has been made.
	NSCorrectionIndicatorTypeReversion NSCorrectionIndicatorType = 1
)


func (e NSCorrectionIndicatorType) String() string {
	switch e {
	case NSCorrectionIndicatorTypeDefault:
		return "NSCorrectionIndicatorTypeDefault"
	case NSCorrectionIndicatorTypeGuesses:
		return "NSCorrectionIndicatorTypeGuesses"
	case NSCorrectionIndicatorTypeReversion:
		return "NSCorrectionIndicatorTypeReversion"
	default:
		return fmt.Sprintf("NSCorrectionIndicatorType(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/CorrectionResponse
type NSCorrectionResponse int

const (
	// NSCorrectionResponseAccepted: The user accepted the correction.
	NSCorrectionResponseAccepted NSCorrectionResponse = 1
	// NSCorrectionResponseEdited: After the correction was accepted, the user edited the corrected word (to something other than its original form.
	NSCorrectionResponseEdited NSCorrectionResponse = 4
	// NSCorrectionResponseIgnored: The user continued in such a way as to ignore the correction.
	NSCorrectionResponseIgnored NSCorrectionResponse = 3
	// NSCorrectionResponseNone: No response was received from the user.
	NSCorrectionResponseNone NSCorrectionResponse = 0
	// NSCorrectionResponseRejected: The user rejected the correction by dismissing the correction indicator.
	NSCorrectionResponseRejected NSCorrectionResponse = 2
	// NSCorrectionResponseReverted: After the correction was accepted, the user reverted the correction back to the original word.
	NSCorrectionResponseReverted NSCorrectionResponse = 5
)


func (e NSCorrectionResponse) String() string {
	switch e {
	case NSCorrectionResponseAccepted:
		return "NSCorrectionResponseAccepted"
	case NSCorrectionResponseEdited:
		return "NSCorrectionResponseEdited"
	case NSCorrectionResponseIgnored:
		return "NSCorrectionResponseIgnored"
	case NSCorrectionResponseNone:
		return "NSCorrectionResponseNone"
	case NSCorrectionResponseRejected:
		return "NSCorrectionResponseRejected"
	case NSCorrectionResponseReverted:
		return "NSCorrectionResponseReverted"
	default:
		return fmt.Sprintf("NSCorrectionResponse(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSCursorFrameResizeDirections
type NSCursorFrameResizeDirections uint

const (
	NSCursorFrameResizeDirectionsInward NSCursorFrameResizeDirections = 1
	NSCursorFrameResizeDirectionsOutward NSCursorFrameResizeDirections = 2
)


func (e NSCursorFrameResizeDirections) String() string {
	switch e {
	case NSCursorFrameResizeDirectionsInward:
		return "NSCursorFrameResizeDirectionsInward"
	case NSCursorFrameResizeDirectionsOutward:
		return "NSCursorFrameResizeDirectionsOutward"
	default:
		return fmt.Sprintf("NSCursorFrameResizeDirections(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSDatePicker/ElementFlags
type NSDatePickerElementFlags int

const (
	NSDatePickerElementFlagEra NSDatePickerElementFlags = 256
	NSDatePickerElementFlagHourMinute NSDatePickerElementFlags = 12
	NSDatePickerElementFlagHourMinuteSecond NSDatePickerElementFlags = 14
	NSDatePickerElementFlagTimeZone NSDatePickerElementFlags = 16
	NSDatePickerElementFlagYearMonth NSDatePickerElementFlags = 192
	NSDatePickerElementFlagYearMonthDay NSDatePickerElementFlags = 224
)


func (e NSDatePickerElementFlags) String() string {
	switch e {
	case NSDatePickerElementFlagEra:
		return "NSDatePickerElementFlagEra"
	case NSDatePickerElementFlagHourMinute:
		return "NSDatePickerElementFlagHourMinute"
	case NSDatePickerElementFlagHourMinuteSecond:
		return "NSDatePickerElementFlagHourMinuteSecond"
	case NSDatePickerElementFlagTimeZone:
		return "NSDatePickerElementFlagTimeZone"
	case NSDatePickerElementFlagYearMonth:
		return "NSDatePickerElementFlagYearMonth"
	case NSDatePickerElementFlagYearMonthDay:
		return "NSDatePickerElementFlagYearMonthDay"
	default:
		return fmt.Sprintf("NSDatePickerElementFlags(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSDatePicker/Mode
type NSDatePickerMode int

const (
	NSDatePickerModeRange NSDatePickerMode = 1
	NSDatePickerModeSingle NSDatePickerMode = 0
)


func (e NSDatePickerMode) String() string {
	switch e {
	case NSDatePickerModeRange:
		return "NSDatePickerModeRange"
	case NSDatePickerModeSingle:
		return "NSDatePickerModeSingle"
	default:
		return fmt.Sprintf("NSDatePickerMode(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSDatePicker/Style
type NSDatePickerStyle int

const (
	NSDatePickerStyleClockAndCalendar NSDatePickerStyle = 1
	NSDatePickerStyleTextField NSDatePickerStyle = 2
	NSDatePickerStyleTextFieldAndStepper NSDatePickerStyle = 0
)


func (e NSDatePickerStyle) String() string {
	switch e {
	case NSDatePickerStyleClockAndCalendar:
		return "NSDatePickerStyleClockAndCalendar"
	case NSDatePickerStyleTextField:
		return "NSDatePickerStyleTextField"
	case NSDatePickerStyleTextFieldAndStepper:
		return "NSDatePickerStyleTextFieldAndStepper"
	default:
		return fmt.Sprintf("NSDatePickerStyle(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSDirectionalRectEdge
type NSDirectionalRectEdge int

const (
	NSDirectionalRectEdgeAll NSDirectionalRectEdge = 1
	NSDirectionalRectEdgeBottom NSDirectionalRectEdge = 4
	NSDirectionalRectEdgeLeading NSDirectionalRectEdge = 2
	NSDirectionalRectEdgeTop NSDirectionalRectEdge = 1
	NSDirectionalRectEdgeTrailing NSDirectionalRectEdge = 8
)


func (e NSDirectionalRectEdge) String() string {
	switch e {
	case NSDirectionalRectEdgeAll:
		return "NSDirectionalRectEdgeAll"
	case NSDirectionalRectEdgeBottom:
		return "NSDirectionalRectEdgeBottom"
	case NSDirectionalRectEdgeLeading:
		return "NSDirectionalRectEdgeLeading"
	case NSDirectionalRectEdgeTrailing:
		return "NSDirectionalRectEdgeTrailing"
	default:
		return fmt.Sprintf("NSDirectionalRectEdge(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSDisplayGamut
type NSDisplayGamut int

const (
	NSDisplayGamutP3 NSDisplayGamut = 2
	NSDisplayGamutSRGB NSDisplayGamut = 1
)


func (e NSDisplayGamut) String() string {
	switch e {
	case NSDisplayGamutP3:
		return "NSDisplayGamutP3"
	case NSDisplayGamutSRGB:
		return "NSDisplayGamutSRGB"
	default:
		return fmt.Sprintf("NSDisplayGamut(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSDocument/ChangeType
type NSDocumentChangeType int

const (
	// NSChangeAutosaved: The document’s contents have been autosaved.
	NSChangeAutosaved NSDocumentChangeType = 4
	// NSChangeCleared: Set change count to 0.
	NSChangeCleared NSDocumentChangeType = 2
	// NSChangeDiscardable: A discardable change has been done.
	NSChangeDiscardable NSDocumentChangeType = 256
	// NSChangeDone: Increment change count.
	NSChangeDone NSDocumentChangeType = 0
	// NSChangeReadOtherContents: The document has been initialized with the contents of a file or file package other than the one whose location is in the  property, and therefore can’t possibly be synchronized with its persistent representation.
	NSChangeReadOtherContents NSDocumentChangeType = 3
	// NSChangeRedone: A single change has been redone.
	NSChangeRedone NSDocumentChangeType = 5
	// NSChangeUndone: Decrement change count.
	NSChangeUndone NSDocumentChangeType = 1
)


func (e NSDocumentChangeType) String() string {
	switch e {
	case NSChangeAutosaved:
		return "NSChangeAutosaved"
	case NSChangeCleared:
		return "NSChangeCleared"
	case NSChangeDiscardable:
		return "NSChangeDiscardable"
	case NSChangeDone:
		return "NSChangeDone"
	case NSChangeReadOtherContents:
		return "NSChangeReadOtherContents"
	case NSChangeRedone:
		return "NSChangeRedone"
	case NSChangeUndone:
		return "NSChangeUndone"
	default:
		return fmt.Sprintf("NSDocumentChangeType(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSDragOperation
type NSDragOperation int

const (
	// NSDragOperationCopy: A constant that indicates the drag can copy the data that the image represents.
	NSDragOperationCopy NSDragOperation = 1
	// NSDragOperationDelete: A constant that indicates the drag can delete the data.
	NSDragOperationDelete NSDragOperation = 32
	// NSDragOperationEvery: A constant that indicates that drag can perform all of the drag operations.
	NSDragOperationEvery NSDragOperation = 0
	// NSDragOperationGeneric: A constant that indicates the destination can define the drag operation.
	NSDragOperationGeneric NSDragOperation = 4
	// NSDragOperationLink: A constant that indicates the drag can share the data.
	NSDragOperationLink NSDragOperation = 2
	// NSDragOperationMove: A constant that indicates the drag can move the data.
	NSDragOperationMove NSDragOperation = 16
	// NSDragOperationPrivate: A constant that indicates the source and destination negotiate the drag operation privately.
	NSDragOperationPrivate NSDragOperation = 8
	// Deprecated.
	NSDragOperationAll NSDragOperation = 15
	// Deprecated.
	NSDragOperationAll_Obsolete NSDragOperation = 15
)


func (e NSDragOperation) String() string {
	switch e {
	case NSDragOperationCopy:
		return "NSDragOperationCopy"
	case NSDragOperationDelete:
		return "NSDragOperationDelete"
	case NSDragOperationEvery:
		return "NSDragOperationEvery"
	case NSDragOperationGeneric:
		return "NSDragOperationGeneric"
	case NSDragOperationLink:
		return "NSDragOperationLink"
	case NSDragOperationMove:
		return "NSDragOperationMove"
	case NSDragOperationPrivate:
		return "NSDragOperationPrivate"
	case NSDragOperationAll:
		return "NSDragOperationAll"
	default:
		return fmt.Sprintf("NSDragOperation(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSDraggingContext
type NSDraggingContext int

const (
	// NSDraggingContextOutsideApplication: A constant that indicates dragging terminates outside the application.
	NSDraggingContextOutsideApplication NSDraggingContext = 0
	// NSDraggingContextWithinApplication: A constant that indicates dragging terminates within the application.
	NSDraggingContextWithinApplication NSDraggingContext = 1
)


func (e NSDraggingContext) String() string {
	switch e {
	case NSDraggingContextOutsideApplication:
		return "NSDraggingContextOutsideApplication"
	case NSDraggingContextWithinApplication:
		return "NSDraggingContextWithinApplication"
	default:
		return fmt.Sprintf("NSDraggingContext(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSDraggingFormation
type NSDraggingFormation int

const (
	// NSDraggingFormationDefault: A constant that represents the system determined formation.
	NSDraggingFormationDefault NSDraggingFormation = 0
	// NSDraggingFormationList: A constant that represents a list formation, so drag images display vertically, non-overlapping with the left edges aligned.
	NSDraggingFormationList NSDraggingFormation = 3
	// NSDraggingFormationNone: A constant that represents no custom formation, so drag images maintain their set positions relative to each other.
	NSDraggingFormationNone NSDraggingFormation = 1
	// NSDraggingFormationPile: A constant that represents a pile formation, so drag images display on top of each other with random rotations.
	NSDraggingFormationPile NSDraggingFormation = 2
	// NSDraggingFormationStack: A constant that represents a stack formation, so drag images display overlapping diagonally.
	NSDraggingFormationStack NSDraggingFormation = 4
)


func (e NSDraggingFormation) String() string {
	switch e {
	case NSDraggingFormationDefault:
		return "NSDraggingFormationDefault"
	case NSDraggingFormationList:
		return "NSDraggingFormationList"
	case NSDraggingFormationNone:
		return "NSDraggingFormationNone"
	case NSDraggingFormationPile:
		return "NSDraggingFormationPile"
	case NSDraggingFormationStack:
		return "NSDraggingFormationStack"
	default:
		return fmt.Sprintf("NSDraggingFormation(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSDraggingItemEnumerationOptions
type NSDraggingItemEnumerationOptions int

const (
	// NSDraggingItemEnumerationClearNonenumeratedImages: A constant that indicates the enumeration clears the image components provider for all dragging items that don’t meet the classes and search options criteria.
	NSDraggingItemEnumerationClearNonenumeratedImages NSDraggingItemEnumerationOptions = 65536
	// NSDraggingItemEnumerationConcurrent: A constant that indicates the enumeration processes dragging items concurrently.
	NSDraggingItemEnumerationConcurrent NSDraggingItemEnumerationOptions = 1
)


func (e NSDraggingItemEnumerationOptions) String() string {
	switch e {
	case NSDraggingItemEnumerationClearNonenumeratedImages:
		return "NSDraggingItemEnumerationClearNonenumeratedImages"
	case NSDraggingItemEnumerationConcurrent:
		return "NSDraggingItemEnumerationConcurrent"
	default:
		return fmt.Sprintf("NSDraggingItemEnumerationOptions(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSDrawer/State-swift.enum
type NSDrawerState int

const (
	NSDrawerClosedState NSDrawerState = 0
	NSDrawerClosingState NSDrawerState = 3
	NSDrawerOpenState NSDrawerState = 2
	NSDrawerOpeningState NSDrawerState = 1
)


func (e NSDrawerState) String() string {
	switch e {
	case NSDrawerClosedState:
		return "NSDrawerClosedState"
	case NSDrawerClosingState:
		return "NSDrawerClosingState"
	case NSDrawerOpenState:
		return "NSDrawerOpenState"
	case NSDrawerOpeningState:
		return "NSDrawerOpeningState"
	default:
		return fmt.Sprintf("NSDrawerState(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSEvent/ButtonMask-swift.struct
type NSEventButtonMask int

const (
	// NSEventButtonMaskPenLowerSide: A mask that matches the button on the lower side of the device.
	NSEventButtonMaskPenLowerSide NSEventButtonMask = 2
	// NSEventButtonMaskPenTip: A mask that matches the pen tip.
	NSEventButtonMaskPenTip NSEventButtonMask = 1
	// NSEventButtonMaskPenUpperSide: A mask that matches the button on the upper side of the device.
	NSEventButtonMaskPenUpperSide NSEventButtonMask = 4
)


func (e NSEventButtonMask) String() string {
	switch e {
	case NSEventButtonMaskPenLowerSide:
		return "NSEventButtonMaskPenLowerSide"
	case NSEventButtonMaskPenTip:
		return "NSEventButtonMaskPenTip"
	case NSEventButtonMaskPenUpperSide:
		return "NSEventButtonMaskPenUpperSide"
	default:
		return fmt.Sprintf("NSEventButtonMask(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSEvent/GestureAxis
type NSEventGestureAxis int

const (
	// NSEventGestureAxisHorizontal: The horizontal axis.
	NSEventGestureAxisHorizontal NSEventGestureAxis = 1
	// NSEventGestureAxisNone: No specific axis.
	NSEventGestureAxisNone NSEventGestureAxis = 0
	// NSEventGestureAxisVertical: The vertical axis.
	NSEventGestureAxisVertical NSEventGestureAxis = 2
)


func (e NSEventGestureAxis) String() string {
	switch e {
	case NSEventGestureAxisHorizontal:
		return "NSEventGestureAxisHorizontal"
	case NSEventGestureAxisNone:
		return "NSEventGestureAxisNone"
	case NSEventGestureAxisVertical:
		return "NSEventGestureAxisVertical"
	default:
		return fmt.Sprintf("NSEventGestureAxis(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSEvent/EventTypeMask
type NSEventMask int

const (
	// NSEventMaskAny: A mask that matches any type of event.
	NSEventMaskAny NSEventMask = 0
	// NSEventMaskAppKitDefined: A mask for AppKit–defined events.
	NSEventMaskAppKitDefined NSEventMask = 8192
	// NSEventMaskApplicationDefined: A mask for app-defined events.
	NSEventMaskApplicationDefined NSEventMask = 32768
	// NSEventMaskBeginGesture: A mask for begin-gesture events.
	NSEventMaskBeginGesture NSEventMask = 524288
	// NSEventMaskChangeMode: A mask for change-mode events.
	NSEventMaskChangeMode NSEventMask = 274877906944
	// NSEventMaskCursorUpdate: A mask for cursor-update events.
	NSEventMaskCursorUpdate NSEventMask = 131072
	// NSEventMaskDirectTouch: A mask for touch events.
	NSEventMaskDirectTouch NSEventMask = 137438953472
	// NSEventMaskEndGesture: A mask for end-gesture events.
	NSEventMaskEndGesture NSEventMask = 1048576
	// NSEventMaskFlagsChanged: A mask for flags-changed events.
	NSEventMaskFlagsChanged NSEventMask = 4096
	// NSEventMaskGesture: A mask for generic gesture events.
	NSEventMaskGesture NSEventMask = 536870912
	// NSEventMaskKeyDown: A mask for key-down events.
	NSEventMaskKeyDown NSEventMask = 1024
	// NSEventMaskKeyUp: A mask for key-up events.
	NSEventMaskKeyUp NSEventMask = 2048
	// NSEventMaskLeftMouseDown: A mask for left mouse-down events.
	NSEventMaskLeftMouseDown NSEventMask = 2
	// NSEventMaskLeftMouseDragged: A mask for left mouse-dragged events.
	NSEventMaskLeftMouseDragged NSEventMask = 64
	// NSEventMaskLeftMouseUp: A mask for left mouse-up events.
	NSEventMaskLeftMouseUp NSEventMask = 4
	// NSEventMaskMagnify: A mask for magnify-gesture events.
	NSEventMaskMagnify NSEventMask = 1073741824
	NSEventMaskMouseCancelled NSEventMask = 1099511627776
	// NSEventMaskMouseEntered: A mask for mouse-entered events.
	NSEventMaskMouseEntered NSEventMask = 256
	// NSEventMaskMouseExited: A mask for mouse-exited events.
	NSEventMaskMouseExited NSEventMask = 512
	// NSEventMaskMouseMoved: A mask for mouse-moved events.
	NSEventMaskMouseMoved NSEventMask = 32
	// NSEventMaskOtherMouseDown: A mask for tertiary mouse-down events.
	NSEventMaskOtherMouseDown NSEventMask = 33554432
	// NSEventMaskOtherMouseDragged: A mask for tertiary mouse-dragged events.
	NSEventMaskOtherMouseDragged NSEventMask = 134217728
	// NSEventMaskOtherMouseUp: A mask for tertiary mouse-up events.
	NSEventMaskOtherMouseUp NSEventMask = 67108864
	// NSEventMaskPeriodic: A mask for periodic events.
	NSEventMaskPeriodic NSEventMask = 65536
	// NSEventMaskPressure: A mask for pressure-change events.
	NSEventMaskPressure NSEventMask = 17179869184
	// NSEventMaskRightMouseDown: A mask for right mouse-down events.
	NSEventMaskRightMouseDown NSEventMask = 8
	// NSEventMaskRightMouseDragged: A mask for right mouse-dragged events.
	NSEventMaskRightMouseDragged NSEventMask = 128
	// NSEventMaskRightMouseUp: A mask for right mouse-up events.
	NSEventMaskRightMouseUp NSEventMask = 16
	// NSEventMaskRotate: A mask for rotate-gesture events.
	NSEventMaskRotate NSEventMask = 262144
	// NSEventMaskScrollWheel: A mask for scroll-wheel events.
	NSEventMaskScrollWheel NSEventMask = 4194304
	// NSEventMaskSmartMagnify: A mask for smart-zoom gesture events.
	NSEventMaskSmartMagnify NSEventMask = 4294967296
	// NSEventMaskSwipe: A mask for swipe-gesture events.
	NSEventMaskSwipe NSEventMask = 2147483648
	// NSEventMaskSystemDefined: A mask for system-defined events.
	NSEventMaskSystemDefined NSEventMask = 16384
	// NSEventMaskTabletPoint: A mask for tablet-point events.
	NSEventMaskTabletPoint NSEventMask = 8388608
	// NSEventMaskTabletProximity: A mask for tablet-proximity events.
	NSEventMaskTabletProximity NSEventMask = 16777216
)


func (e NSEventMask) String() string {
	switch e {
	case NSEventMaskAny:
		return "NSEventMaskAny"
	case NSEventMaskAppKitDefined:
		return "NSEventMaskAppKitDefined"
	case NSEventMaskApplicationDefined:
		return "NSEventMaskApplicationDefined"
	case NSEventMaskBeginGesture:
		return "NSEventMaskBeginGesture"
	case NSEventMaskChangeMode:
		return "NSEventMaskChangeMode"
	case NSEventMaskCursorUpdate:
		return "NSEventMaskCursorUpdate"
	case NSEventMaskDirectTouch:
		return "NSEventMaskDirectTouch"
	case NSEventMaskEndGesture:
		return "NSEventMaskEndGesture"
	case NSEventMaskFlagsChanged:
		return "NSEventMaskFlagsChanged"
	case NSEventMaskGesture:
		return "NSEventMaskGesture"
	case NSEventMaskKeyDown:
		return "NSEventMaskKeyDown"
	case NSEventMaskKeyUp:
		return "NSEventMaskKeyUp"
	case NSEventMaskLeftMouseDown:
		return "NSEventMaskLeftMouseDown"
	case NSEventMaskLeftMouseDragged:
		return "NSEventMaskLeftMouseDragged"
	case NSEventMaskLeftMouseUp:
		return "NSEventMaskLeftMouseUp"
	case NSEventMaskMagnify:
		return "NSEventMaskMagnify"
	case NSEventMaskMouseCancelled:
		return "NSEventMaskMouseCancelled"
	case NSEventMaskMouseEntered:
		return "NSEventMaskMouseEntered"
	case NSEventMaskMouseExited:
		return "NSEventMaskMouseExited"
	case NSEventMaskMouseMoved:
		return "NSEventMaskMouseMoved"
	case NSEventMaskOtherMouseDown:
		return "NSEventMaskOtherMouseDown"
	case NSEventMaskOtherMouseDragged:
		return "NSEventMaskOtherMouseDragged"
	case NSEventMaskOtherMouseUp:
		return "NSEventMaskOtherMouseUp"
	case NSEventMaskPeriodic:
		return "NSEventMaskPeriodic"
	case NSEventMaskPressure:
		return "NSEventMaskPressure"
	case NSEventMaskRightMouseDown:
		return "NSEventMaskRightMouseDown"
	case NSEventMaskRightMouseDragged:
		return "NSEventMaskRightMouseDragged"
	case NSEventMaskRightMouseUp:
		return "NSEventMaskRightMouseUp"
	case NSEventMaskRotate:
		return "NSEventMaskRotate"
	case NSEventMaskScrollWheel:
		return "NSEventMaskScrollWheel"
	case NSEventMaskSmartMagnify:
		return "NSEventMaskSmartMagnify"
	case NSEventMaskSwipe:
		return "NSEventMaskSwipe"
	case NSEventMaskSystemDefined:
		return "NSEventMaskSystemDefined"
	case NSEventMaskTabletPoint:
		return "NSEventMaskTabletPoint"
	case NSEventMaskTabletProximity:
		return "NSEventMaskTabletProximity"
	default:
		return fmt.Sprintf("NSEventMask(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSEvent/ModifierFlags-swift.struct
type NSEventModifierFlags int

const (
	// NSEventModifierFlagCapsLock: The Caps Lock key has been pressed.
	NSEventModifierFlagCapsLock NSEventModifierFlags = 65536
	// NSEventModifierFlagCommand: The Command key has been pressed.
	NSEventModifierFlagCommand NSEventModifierFlags = 1048576
	// NSEventModifierFlagControl: The Control key has been pressed.
	NSEventModifierFlagControl NSEventModifierFlags = 262144
	// NSEventModifierFlagDeviceIndependentFlagsMask: Device-independent modifier flags are masked.
	NSEventModifierFlagDeviceIndependentFlagsMask NSEventModifierFlags = 4294901760
	// NSEventModifierFlagFunction: A function key has been pressed.
	NSEventModifierFlagFunction NSEventModifierFlags = 8388608
	// NSEventModifierFlagHelp: The Help key has been pressed.
	NSEventModifierFlagHelp NSEventModifierFlags = 4194304
	// NSEventModifierFlagNumericPad: A key in the numeric keypad or an arrow key has been pressed.
	NSEventModifierFlagNumericPad NSEventModifierFlags = 2097152
	// NSEventModifierFlagOption: The Option or Alt key has been pressed.
	NSEventModifierFlagOption NSEventModifierFlags = 524288
	// NSEventModifierFlagShift: The Shift key has been pressed.
	NSEventModifierFlagShift NSEventModifierFlags = 131072
)


func (e NSEventModifierFlags) String() string {
	switch e {
	case NSEventModifierFlagCapsLock:
		return "NSEventModifierFlagCapsLock"
	case NSEventModifierFlagCommand:
		return "NSEventModifierFlagCommand"
	case NSEventModifierFlagControl:
		return "NSEventModifierFlagControl"
	case NSEventModifierFlagDeviceIndependentFlagsMask:
		return "NSEventModifierFlagDeviceIndependentFlagsMask"
	case NSEventModifierFlagFunction:
		return "NSEventModifierFlagFunction"
	case NSEventModifierFlagHelp:
		return "NSEventModifierFlagHelp"
	case NSEventModifierFlagNumericPad:
		return "NSEventModifierFlagNumericPad"
	case NSEventModifierFlagOption:
		return "NSEventModifierFlagOption"
	case NSEventModifierFlagShift:
		return "NSEventModifierFlagShift"
	default:
		return fmt.Sprintf("NSEventModifierFlags(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSEvent/Phase-swift.struct
type NSEventPhase int

const (
	// NSEventPhaseBegan: An event phase has begun.
	NSEventPhaseBegan NSEventPhase = 1
	// NSEventPhaseCancelled: The system canceled the event phase.
	NSEventPhaseCancelled NSEventPhase = 1
	// NSEventPhaseChanged: An event phase has changed.
	NSEventPhaseChanged NSEventPhase = 1
	// NSEventPhaseEnded: The event phase ended.
	NSEventPhaseEnded NSEventPhase = 1
	// NSEventPhaseMayBegin: The system event phase may begin.
	NSEventPhaseMayBegin NSEventPhase = 1
	// NSEventPhaseStationary: An event phase is in progress but hasn’t moved since the previous event.
	NSEventPhaseStationary NSEventPhase = 1
)


func (e NSEventPhase) String() string {
	switch e {
	case NSEventPhaseBegan:
		return "NSEventPhaseBegan"
	default:
		return fmt.Sprintf("NSEventPhase(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSEvent/EventSubtype
type NSEventSubtype int

const (
	// NSEventSubtypeApplicationActivated: An app-activation event occurred.
	NSEventSubtypeApplicationActivated NSEventSubtype = 1
	// NSEventSubtypeApplicationDeactivated: An app-deactivation event occurred.
	NSEventSubtypeApplicationDeactivated NSEventSubtype = 2
	// NSEventSubtypeScreenChanged: An event that indicates a window changed screens.
	NSEventSubtypeScreenChanged NSEventSubtype = 8
	// NSEventSubtypeTouch: A touch event occurred.
	NSEventSubtypeTouch NSEventSubtype = 3
	// NSEventSubtypeWindowExposed: An event that indicates a window’s contents are visible again.
	NSEventSubtypeWindowExposed NSEventSubtype = 0
	// NSEventSubtypeWindowMoved: An event that indicates a window moved.
	NSEventSubtypeWindowMoved NSEventSubtype = 4
)


func (e NSEventSubtype) String() string {
	switch e {
	case NSEventSubtypeApplicationActivated:
		return "NSEventSubtypeApplicationActivated"
	case NSEventSubtypeApplicationDeactivated:
		return "NSEventSubtypeApplicationDeactivated"
	case NSEventSubtypeScreenChanged:
		return "NSEventSubtypeScreenChanged"
	case NSEventSubtypeTouch:
		return "NSEventSubtypeTouch"
	case NSEventSubtypeWindowExposed:
		return "NSEventSubtypeWindowExposed"
	case NSEventSubtypeWindowMoved:
		return "NSEventSubtypeWindowMoved"
	default:
		return fmt.Sprintf("NSEventSubtype(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSEvent/SwipeTrackingOptions
type NSEventSwipeTrackingOptions int

const (
	// NSEventSwipeTrackingClampGestureAmount: Don’t allow gestureAmount to go beyond +/-1.0
	NSEventSwipeTrackingClampGestureAmount NSEventSwipeTrackingOptions = 1
	// NSEventSwipeTrackingLockDirection: Clamp gestureAmount to 0 if the user starts to swipe in the opposite direction than they started.
	NSEventSwipeTrackingLockDirection NSEventSwipeTrackingOptions = 1
)


func (e NSEventSwipeTrackingOptions) String() string {
	switch e {
	case NSEventSwipeTrackingClampGestureAmount:
		return "NSEventSwipeTrackingClampGestureAmount"
	default:
		return fmt.Sprintf("NSEventSwipeTrackingOptions(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSEvent/EventType
type NSEventType int

const (
	// NSEventTypeAppKitDefined: An AppKit-related event occurred.
	NSEventTypeAppKitDefined NSEventType = 13
	// NSEventTypeApplicationDefined: An app-defined event occurred.
	NSEventTypeApplicationDefined NSEventType = 15
	// NSEventTypeBeginGesture: An event marking the beginning of a gesture.
	NSEventTypeBeginGesture NSEventType = 19
	// NSEventTypeChangeMode: The user changed the mode of a connected device.
	NSEventTypeChangeMode NSEventType = 38
	// NSEventTypeCursorUpdate: An event that updates the cursor.
	NSEventTypeCursorUpdate NSEventType = 17
	// NSEventTypeDirectTouch: The user touched a portion of the touch bar.
	NSEventTypeDirectTouch NSEventType = 37
	// NSEventTypeEndGesture: An event that marks the end of a gesture.
	NSEventTypeEndGesture NSEventType = 20
	// NSEventTypeFlagsChanged: The event flags changed.
	NSEventTypeFlagsChanged NSEventType = 12
	// NSEventTypeGesture: The user performed a nonspecific type of gesture.
	NSEventTypeGesture NSEventType = 29
	// NSEventTypeKeyDown: The user pressed a key on the keyboard.
	NSEventTypeKeyDown NSEventType = 10
	// NSEventTypeKeyUp: The user released a key on the keyboard.
	NSEventTypeKeyUp NSEventType = 11
	// NSEventTypeLeftMouseDown: The user pressed the left mouse button.
	NSEventTypeLeftMouseDown NSEventType = 1
	// NSEventTypeLeftMouseDragged: The user moved the mouse while holding down the left mouse button.
	NSEventTypeLeftMouseDragged NSEventType = 6
	// NSEventTypeLeftMouseUp: The user released the left mouse button.
	NSEventTypeLeftMouseUp NSEventType = 2
	// NSEventTypeMagnify: The user performed a pinch-open or pinch-close gesture.
	NSEventTypeMagnify NSEventType = 30
	// NSEventTypeMouseEntered: The cursor entered a well-defined area, such as a view.
	NSEventTypeMouseEntered NSEventType = 8
	// NSEventTypeMouseExited: The cursor exited a well-defined area, such as a view.
	NSEventTypeMouseExited NSEventType = 9
	// NSEventTypeMouseMoved: The user moved the mouse in a way that caused the cursor to move onscreen.
	NSEventTypeMouseMoved NSEventType = 5
	// NSEventTypeOtherMouseDown: The user pressed a tertiary mouse button.
	NSEventTypeOtherMouseDown NSEventType = 25
	// NSEventTypeOtherMouseDragged: The user moved the mouse while holding down a tertiary mouse button.
	NSEventTypeOtherMouseDragged NSEventType = 27
	// NSEventTypeOtherMouseUp: The user released a tertiary mouse button.
	NSEventTypeOtherMouseUp NSEventType = 26
	// NSEventTypePeriodic: An event that provides execution time to periodic tasks.
	NSEventTypePeriodic NSEventType = 16
	// NSEventTypePressure: An event that reports a change in pressure on a pressure-sensitive device.
	NSEventTypePressure NSEventType = 34
	// NSEventTypeQuickLook: An event that initiates a Quick Look request.
	NSEventTypeQuickLook NSEventType = 33
	// NSEventTypeRightMouseDown: The user pressed the right mouse button.
	NSEventTypeRightMouseDown NSEventType = 3
	// NSEventTypeRightMouseDragged: The user moved the mouse while holding down the right mouse button.
	NSEventTypeRightMouseDragged NSEventType = 7
	// NSEventTypeRightMouseUp: The user released the right mouse button.
	NSEventTypeRightMouseUp NSEventType = 4
	// NSEventTypeRotate: The user performed a rotate gesture.
	NSEventTypeRotate NSEventType = 18
	// NSEventTypeScrollWheel: The scroll wheel position changed.
	NSEventTypeScrollWheel NSEventType = 22
	// NSEventTypeSmartMagnify: The user performed a smart-zoom gesture.
	NSEventTypeSmartMagnify NSEventType = 32
	// NSEventTypeSwipe: The user performed a swipe gesture.
	NSEventTypeSwipe NSEventType = 31
	// NSEventTypeSystemDefined: A system-related event occurred.
	NSEventTypeSystemDefined NSEventType = 14
	// NSEventTypeTabletPoint: The user touched a point on a tablet.
	NSEventTypeTabletPoint NSEventType = 23
	// NSEventTypeTabletProximity: A pointing device is near, but not touching, the associated tablet.
	NSEventTypeTabletProximity NSEventType = 24
)


func (e NSEventType) String() string {
	switch e {
	case NSEventTypeAppKitDefined:
		return "NSEventTypeAppKitDefined"
	case NSEventTypeApplicationDefined:
		return "NSEventTypeApplicationDefined"
	case NSEventTypeBeginGesture:
		return "NSEventTypeBeginGesture"
	case NSEventTypeChangeMode:
		return "NSEventTypeChangeMode"
	case NSEventTypeCursorUpdate:
		return "NSEventTypeCursorUpdate"
	case NSEventTypeDirectTouch:
		return "NSEventTypeDirectTouch"
	case NSEventTypeEndGesture:
		return "NSEventTypeEndGesture"
	case NSEventTypeFlagsChanged:
		return "NSEventTypeFlagsChanged"
	case NSEventTypeGesture:
		return "NSEventTypeGesture"
	case NSEventTypeKeyDown:
		return "NSEventTypeKeyDown"
	case NSEventTypeKeyUp:
		return "NSEventTypeKeyUp"
	case NSEventTypeLeftMouseDown:
		return "NSEventTypeLeftMouseDown"
	case NSEventTypeLeftMouseDragged:
		return "NSEventTypeLeftMouseDragged"
	case NSEventTypeLeftMouseUp:
		return "NSEventTypeLeftMouseUp"
	case NSEventTypeMagnify:
		return "NSEventTypeMagnify"
	case NSEventTypeMouseEntered:
		return "NSEventTypeMouseEntered"
	case NSEventTypeMouseExited:
		return "NSEventTypeMouseExited"
	case NSEventTypeMouseMoved:
		return "NSEventTypeMouseMoved"
	case NSEventTypeOtherMouseDown:
		return "NSEventTypeOtherMouseDown"
	case NSEventTypeOtherMouseDragged:
		return "NSEventTypeOtherMouseDragged"
	case NSEventTypeOtherMouseUp:
		return "NSEventTypeOtherMouseUp"
	case NSEventTypePeriodic:
		return "NSEventTypePeriodic"
	case NSEventTypePressure:
		return "NSEventTypePressure"
	case NSEventTypeQuickLook:
		return "NSEventTypeQuickLook"
	case NSEventTypeRightMouseDown:
		return "NSEventTypeRightMouseDown"
	case NSEventTypeRightMouseDragged:
		return "NSEventTypeRightMouseDragged"
	case NSEventTypeRightMouseUp:
		return "NSEventTypeRightMouseUp"
	case NSEventTypeRotate:
		return "NSEventTypeRotate"
	case NSEventTypeScrollWheel:
		return "NSEventTypeScrollWheel"
	case NSEventTypeSmartMagnify:
		return "NSEventTypeSmartMagnify"
	case NSEventTypeSwipe:
		return "NSEventTypeSwipe"
	case NSEventTypeSystemDefined:
		return "NSEventTypeSystemDefined"
	case NSEventTypeTabletPoint:
		return "NSEventTypeTabletPoint"
	case NSEventTypeTabletProximity:
		return "NSEventTypeTabletProximity"
	default:
		return fmt.Sprintf("NSEventType(%d)", e)
	}
}




type NSFileHandlingPanel uint

const (
	// Deprecated.
	NSFileHandlingPanelCancelButton NSFileHandlingPanel = 0
	// Deprecated.
	NSFileHandlingPanelOKButton NSFileHandlingPanel = 0
)


func (e NSFileHandlingPanel) String() string {
	switch e {
	case NSFileHandlingPanelCancelButton:
		return "NSFileHandlingPanelCancelButton"
	default:
		return fmt.Sprintf("NSFileHandlingPanel(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSFindPanelAction
type NSFindPanelAction int

const (
	// NSFindPanelActionNext: Finds the next instance of the queried text.
	NSFindPanelActionNext NSFindPanelAction = 2
	// NSFindPanelActionPrevious: Finds the previous instance of the queried text.
	NSFindPanelActionPrevious NSFindPanelAction = 3
	// NSFindPanelActionReplace: Replaces a single query instance within the text view.
	NSFindPanelActionReplace NSFindPanelAction = 5
	// NSFindPanelActionReplaceAll: Replaces all query instances within the text view.
	NSFindPanelActionReplaceAll NSFindPanelAction = 4
	// NSFindPanelActionReplaceAllInSelection: Replaces all query instances within the selection.
	NSFindPanelActionReplaceAllInSelection NSFindPanelAction = 8
	// NSFindPanelActionReplaceAndFind: Replaces a single query instance and finds the next.
	NSFindPanelActionReplaceAndFind NSFindPanelAction = 6
	// NSFindPanelActionSelectAll: Selects all query instances in the text view.
	NSFindPanelActionSelectAll NSFindPanelAction = 9
	// NSFindPanelActionSelectAllInSelection: Selects all query instances within the selection.
	NSFindPanelActionSelectAllInSelection NSFindPanelAction = 10
	// NSFindPanelActionSetFindString: Sets the query string to the current selection.
	NSFindPanelActionSetFindString NSFindPanelAction = 7
	// NSFindPanelActionShowFindPanel: Displays the find panel.
	NSFindPanelActionShowFindPanel NSFindPanelAction = 1
)


func (e NSFindPanelAction) String() string {
	switch e {
	case NSFindPanelActionNext:
		return "NSFindPanelActionNext"
	case NSFindPanelActionPrevious:
		return "NSFindPanelActionPrevious"
	case NSFindPanelActionReplace:
		return "NSFindPanelActionReplace"
	case NSFindPanelActionReplaceAll:
		return "NSFindPanelActionReplaceAll"
	case NSFindPanelActionReplaceAllInSelection:
		return "NSFindPanelActionReplaceAllInSelection"
	case NSFindPanelActionReplaceAndFind:
		return "NSFindPanelActionReplaceAndFind"
	case NSFindPanelActionSelectAll:
		return "NSFindPanelActionSelectAll"
	case NSFindPanelActionSelectAllInSelection:
		return "NSFindPanelActionSelectAllInSelection"
	case NSFindPanelActionSetFindString:
		return "NSFindPanelActionSetFindString"
	case NSFindPanelActionShowFindPanel:
		return "NSFindPanelActionShowFindPanel"
	default:
		return fmt.Sprintf("NSFindPanelAction(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSFindPanelSubstringMatchType
type NSFindPanelSubstringMatchType int

const (
	NSFindPanelSubstringMatchTypeContains NSFindPanelSubstringMatchType = 0
	NSFindPanelSubstringMatchTypeEndsWith NSFindPanelSubstringMatchType = 3
	NSFindPanelSubstringMatchTypeFullWord NSFindPanelSubstringMatchType = 2
	NSFindPanelSubstringMatchTypeStartsWith NSFindPanelSubstringMatchType = 1
)


func (e NSFindPanelSubstringMatchType) String() string {
	switch e {
	case NSFindPanelSubstringMatchTypeContains:
		return "NSFindPanelSubstringMatchTypeContains"
	case NSFindPanelSubstringMatchTypeEndsWith:
		return "NSFindPanelSubstringMatchTypeEndsWith"
	case NSFindPanelSubstringMatchTypeFullWord:
		return "NSFindPanelSubstringMatchTypeFullWord"
	case NSFindPanelSubstringMatchTypeStartsWith:
		return "NSFindPanelSubstringMatchTypeStartsWith"
	default:
		return fmt.Sprintf("NSFindPanelSubstringMatchType(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSFocusRingPlacement
type NSFocusRingPlacement int

const (
	// NSFocusRingAbove: Draw the focus ring over an image.
	NSFocusRingAbove NSFocusRingPlacement = 2
	// NSFocusRingBelow: Draw the focus ring under text.
	NSFocusRingBelow NSFocusRingPlacement = 1
	// NSFocusRingOnly: Draw the focus ring if you don’t have an image or text.
	NSFocusRingOnly NSFocusRingPlacement = 0
)


func (e NSFocusRingPlacement) String() string {
	switch e {
	case NSFocusRingAbove:
		return "NSFocusRingAbove"
	case NSFocusRingBelow:
		return "NSFocusRingBelow"
	case NSFocusRingOnly:
		return "NSFocusRingOnly"
	default:
		return fmt.Sprintf("NSFocusRingPlacement(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSFocusRingType
type NSFocusRingType int

const (
	// NSFocusRingTypeDefault: The default focus ring type for a view or cell.
	NSFocusRingTypeDefault NSFocusRingType = 0
	// NSFocusRingTypeExterior: The standard Aqua focus ring.
	NSFocusRingTypeExterior NSFocusRingType = 2
	// NSFocusRingTypeNone: No focus ring.
	NSFocusRingTypeNone NSFocusRingType = 1
)


func (e NSFocusRingType) String() string {
	switch e {
	case NSFocusRingTypeDefault:
		return "NSFocusRingTypeDefault"
	case NSFocusRingTypeExterior:
		return "NSFocusRingTypeExterior"
	case NSFocusRingTypeNone:
		return "NSFocusRingTypeNone"
	default:
		return fmt.Sprintf("NSFocusRingType(%d)", e)
	}
}




const (
	// NSFontBoldTrait: The font’s typestyle is boldface.
	NSFontBoldTrait uint = 2
	// NSFontClarendonSerifsClass: A font where the style is a variation of the Oldstyle Serifs and the Transitional Serifs.
	NSFontClarendonSerifsClass uint = 1073741824
	// NSFontCondensedTrait: The font’s typestyle is condensed.
	NSFontCondensedTrait uint = 64
	// NSFontExpandedTrait: The font’s typestyle is expanded.
	NSFontExpandedTrait uint = 32
	// NSFontFreeformSerifsClass: A font where the style includes serifs, but it expresses a design freedom that does not generally fit within the other serif design classifications.
	NSFontFreeformSerifsClass uint = 1879048192
	// NSFontItalicTrait: The font’s typestyle is italic.
	NSFontItalicTrait uint = 1
	// NSFontModernSerifsClass: A font where the style is based on the Latin printing style of the 20th century.
	NSFontModernSerifsClass uint = 805306368
	// NSFontMonoSpaceTrait: The font uses fixed-pitch glyphs if available.
	NSFontMonoSpaceTrait uint = 1024
	// NSFontOldStyleSerifsClass: A font where the style is based on the Latin printing style of the 15th to 17th century.
	NSFontOldStyleSerifsClass uint = 268435456
	// NSFontOrnamentalsClass: A font where the style includes highly decorated or stylized character shapes such as those typically used in headlines.
	NSFontOrnamentalsClass uint = 2415919104
	// NSFontSansSerifClass: A font where the style includes most basic letter forms (excluding Scripts and Ornamentals) that do not have serifs on the strokes.
	NSFontSansSerifClass uint = 2147483648
	// NSFontScriptsClass: A font where the style is among those typefaces designed to simulate handwriting.
	NSFontScriptsClass uint = 2684354560
	// NSFontSlabSerifsClass: A font where the style is characterized by serifs with a square transition between the strokes and the serifs (no brackets).
	NSFontSlabSerifsClass uint = 1342177280
	// NSFontSymbolicClass: A font where the style is generally design independent, making it suitable for special characters (icons, dingbats, technical symbols, and so on) that may be used equally well with any font.
	NSFontSymbolicClass uint = 3221225472
	// NSFontTransitionalSerifsClass: A font where the style is based on the Latin printing style of the 18th to 19th century.
	NSFontTransitionalSerifsClass uint = 536870912
	// NSFontUIOptimizedTrait: The font synthesizes appropriate attributes for user interface rendering, such as control titles, if necessary.
	NSFontUIOptimizedTrait uint = 4096
	// NSFontUnknownClass: A font with no design classification.
	NSFontUnknownClass uint = 0
	// NSFontVerticalTrait: The font uses vertical glyph variants and metrics.
	NSFontVerticalTrait uint = 2048
)


// See: https://developer.apple.com/documentation/AppKit/NSFontAction
type NSFontAction int

const (
	// NSAddTraitFontAction: Converts the font to have an additional trait using .
	NSAddTraitFontAction NSFontAction = 2
	// NSHeavierFontAction: Converts the font to a heavier weight using .
	NSHeavierFontAction NSFontAction = 5
	// NSLighterFontAction: Converts the font to a lighter weight using .
	NSLighterFontAction NSFontAction = 6
	// NSNoFontChangeAction: No action; the font is returned unchanged.
	NSNoFontChangeAction NSFontAction = 0
	// NSRemoveTraitFontAction: Converts the font to remove a trait using .
	NSRemoveTraitFontAction NSFontAction = 7
	// NSSizeDownFontAction: Converts the font to a smaller size using .
	NSSizeDownFontAction NSFontAction = 4
	// NSSizeUpFontAction: Converts the font to a larger size using .
	NSSizeUpFontAction NSFontAction = 3
	// NSViaPanelFontAction: Converts the font according to the  method .
	NSViaPanelFontAction NSFontAction = 1
)


func (e NSFontAction) String() string {
	switch e {
	case NSAddTraitFontAction:
		return "NSAddTraitFontAction"
	case NSHeavierFontAction:
		return "NSHeavierFontAction"
	case NSLighterFontAction:
		return "NSLighterFontAction"
	case NSNoFontChangeAction:
		return "NSNoFontChangeAction"
	case NSRemoveTraitFontAction:
		return "NSRemoveTraitFontAction"
	case NSSizeDownFontAction:
		return "NSSizeDownFontAction"
	case NSSizeUpFontAction:
		return "NSSizeUpFontAction"
	case NSViaPanelFontAction:
		return "NSViaPanelFontAction"
	default:
		return fmt.Sprintf("NSFontAction(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSFontAssetRequest/Options
type NSFontAssetRequestOptions int

const (
	NSFontAssetRequestOptionUsesStandardUI NSFontAssetRequestOptions = 1
)


func (e NSFontAssetRequestOptions) String() string {
	switch e {
	case NSFontAssetRequestOptionUsesStandardUI:
		return "NSFontAssetRequestOptionUsesStandardUI"
	default:
		return fmt.Sprintf("NSFontAssetRequestOptions(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSFontCollectionOptions
type NSFontCollectionOptions int

const (
	// NSFontCollectionApplicationOnlyMask: Makes the collection available only to the application.
	NSFontCollectionApplicationOnlyMask NSFontCollectionOptions = 1
)


func (e NSFontCollectionOptions) String() string {
	switch e {
	case NSFontCollectionApplicationOnlyMask:
		return "NSFontCollectionApplicationOnlyMask"
	default:
		return fmt.Sprintf("NSFontCollectionOptions(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSFontCollection/Visibility
type NSFontCollectionVisibility int

const (
	// NSFontCollectionVisibilityComputer: The font collection is visible to all users and is stored persistently.
	NSFontCollectionVisibilityComputer NSFontCollectionVisibility = 4
	// NSFontCollectionVisibilityProcess: The font collection is visible within this process and is not persistent.
	NSFontCollectionVisibilityProcess NSFontCollectionVisibility = 1
	// NSFontCollectionVisibilityUser: The font collection is visible to all processes and is stored persistently.
	NSFontCollectionVisibilityUser NSFontCollectionVisibility = 2
)


func (e NSFontCollectionVisibility) String() string {
	switch e {
	case NSFontCollectionVisibilityComputer:
		return "NSFontCollectionVisibilityComputer"
	case NSFontCollectionVisibilityProcess:
		return "NSFontCollectionVisibilityProcess"
	case NSFontCollectionVisibilityUser:
		return "NSFontCollectionVisibilityUser"
	default:
		return fmt.Sprintf("NSFontCollectionVisibility(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/SymbolicTraits-swift.struct
type NSFontDescriptorSymbolicTraits int

const (
	// NSFontDescriptorClassClarendonSerifs: The font’s characters include variations of old style and transitional serifs.
	NSFontDescriptorClassClarendonSerifs NSFontDescriptorSymbolicTraits = 1073741824
	// NSFontDescriptorClassFreeformSerifs: The font’s characters include serifs, and don’t generally fit within other serif design classifications.
	NSFontDescriptorClassFreeformSerifs NSFontDescriptorSymbolicTraits = 1879048192
	// NSFontDescriptorClassMask: The font family class mask that you use to access font descriptor values.
	NSFontDescriptorClassMask NSFontDescriptorSymbolicTraits = 4026531840
	// NSFontDescriptorClassModernSerifs: The font’s characters include serifs, and reflect the Latin printing style of the 20th century.
	NSFontDescriptorClassModernSerifs NSFontDescriptorSymbolicTraits = 805306368
	// NSFontDescriptorClassOldStyleSerifs: The font’s characters include serifs, and reflect the Latin printing style of the 15th to 17th centuries.
	NSFontDescriptorClassOldStyleSerifs NSFontDescriptorSymbolicTraits = 268435456
	// NSFontDescriptorClassOrnamentals: The font’s characters use highly decorated or stylized character shapes.
	NSFontDescriptorClassOrnamentals NSFontDescriptorSymbolicTraits = 2415919104
	// NSFontDescriptorClassSansSerif: The font’s characters don’t have serifs.
	NSFontDescriptorClassSansSerif NSFontDescriptorSymbolicTraits = 2147483648
	// NSFontDescriptorClassScripts: The font’s characters simulate handwriting.
	NSFontDescriptorClassScripts NSFontDescriptorSymbolicTraits = 2684354560
	// NSFontDescriptorClassSlabSerifs: The font’s characters use square transitions, without brackets, between strokes and serifs.
	NSFontDescriptorClassSlabSerifs NSFontDescriptorSymbolicTraits = 1342177280
	// NSFontDescriptorClassSymbolic: The font’s characters consist mainly of symbols rather than letters and numbers.
	NSFontDescriptorClassSymbolic NSFontDescriptorSymbolicTraits = 3221225472
	// NSFontDescriptorClassTransitionalSerifs: The font’s characters include serifs, and reflect the Latin printing style of the 18th to 19th centuries.
	NSFontDescriptorClassTransitionalSerifs NSFontDescriptorSymbolicTraits = 536870912
	// NSFontDescriptorTraitBold: The font’s style is boldface.
	NSFontDescriptorTraitBold NSFontDescriptorSymbolicTraits = 2
	// NSFontDescriptorTraitCondensed: The font’s characters have a condensed width.
	NSFontDescriptorTraitCondensed NSFontDescriptorSymbolicTraits = 64
	// NSFontDescriptorTraitExpanded: The font’s characters have an expanded width.
	NSFontDescriptorTraitExpanded NSFontDescriptorSymbolicTraits = 32
	// NSFontDescriptorTraitItalic: The font’s style is italic.
	NSFontDescriptorTraitItalic NSFontDescriptorSymbolicTraits = 1
	// NSFontDescriptorTraitLooseLeading: The font uses a leading value that’s greater than the default.
	NSFontDescriptorTraitLooseLeading NSFontDescriptorSymbolicTraits = 65536
	// NSFontDescriptorTraitMonoSpace: The font’s characters all have the same width.
	NSFontDescriptorTraitMonoSpace NSFontDescriptorSymbolicTraits = 1024
	// NSFontDescriptorTraitTightLeading: The font uses a leading value that’s less than the default.
	NSFontDescriptorTraitTightLeading NSFontDescriptorSymbolicTraits = 32768
	// NSFontDescriptorTraitUIOptimized: The font synthesizes appropriate attributes for user interface rendering, such as in control titles, if necessary.
	NSFontDescriptorTraitUIOptimized NSFontDescriptorSymbolicTraits = 4096
	// NSFontDescriptorTraitVertical: The font uses vertical glyph variants and metrics.
	NSFontDescriptorTraitVertical NSFontDescriptorSymbolicTraits = 2048
)


func (e NSFontDescriptorSymbolicTraits) String() string {
	switch e {
	case NSFontDescriptorClassClarendonSerifs:
		return "NSFontDescriptorClassClarendonSerifs"
	case NSFontDescriptorClassFreeformSerifs:
		return "NSFontDescriptorClassFreeformSerifs"
	case NSFontDescriptorClassMask:
		return "NSFontDescriptorClassMask"
	case NSFontDescriptorClassModernSerifs:
		return "NSFontDescriptorClassModernSerifs"
	case NSFontDescriptorClassOldStyleSerifs:
		return "NSFontDescriptorClassOldStyleSerifs"
	case NSFontDescriptorClassOrnamentals:
		return "NSFontDescriptorClassOrnamentals"
	case NSFontDescriptorClassSansSerif:
		return "NSFontDescriptorClassSansSerif"
	case NSFontDescriptorClassScripts:
		return "NSFontDescriptorClassScripts"
	case NSFontDescriptorClassSlabSerifs:
		return "NSFontDescriptorClassSlabSerifs"
	case NSFontDescriptorClassSymbolic:
		return "NSFontDescriptorClassSymbolic"
	case NSFontDescriptorClassTransitionalSerifs:
		return "NSFontDescriptorClassTransitionalSerifs"
	case NSFontDescriptorTraitBold:
		return "NSFontDescriptorTraitBold"
	case NSFontDescriptorTraitCondensed:
		return "NSFontDescriptorTraitCondensed"
	case NSFontDescriptorTraitExpanded:
		return "NSFontDescriptorTraitExpanded"
	case NSFontDescriptorTraitItalic:
		return "NSFontDescriptorTraitItalic"
	case NSFontDescriptorTraitLooseLeading:
		return "NSFontDescriptorTraitLooseLeading"
	case NSFontDescriptorTraitMonoSpace:
		return "NSFontDescriptorTraitMonoSpace"
	case NSFontDescriptorTraitTightLeading:
		return "NSFontDescriptorTraitTightLeading"
	case NSFontDescriptorTraitUIOptimized:
		return "NSFontDescriptorTraitUIOptimized"
	case NSFontDescriptorTraitVertical:
		return "NSFontDescriptorTraitVertical"
	default:
		return fmt.Sprintf("NSFontDescriptorSymbolicTraits(%d)", e)
	}
}




type NSFontFamilyClassMaskConstants uint

const (
	// NSFontFamilyClassMask: Constant you use to access [NSFontFamilyClass] values in the upper four bits of [NSFontSymbolicTraits].
	NSFontFamilyClassMask NSFontFamilyClassMaskConstants = 4026531840
)


func (e NSFontFamilyClassMaskConstants) String() string {
	switch e {
	case NSFontFamilyClassMask:
		return "NSFontFamilyClassMask"
	default:
		return fmt.Sprintf("NSFontFamilyClassMaskConstants(%d)", e)
	}
}




const (
	// NSFontPanelAllEffectsModeMask: Display all the effects user interface items.
	NSFontPanelAllEffectsModeMask uint = 1048320
	// NSFontPanelAllModesMask: Display all the available adornments.
	NSFontPanelAllModesMask uint = 4294967295
	// NSFontPanelCollectionModeMask: Display the font collections column.
	NSFontPanelCollectionModeMask uint = 4
	// NSFontPanelDocumentColorEffectModeMask: Display the document color button.
	NSFontPanelDocumentColorEffectModeMask uint = 2048
	// NSFontPanelFaceModeMask: Display the typeface column.
	NSFontPanelFaceModeMask uint = 1
	// NSFontPanelShadowEffectModeMask: Display the shadow effects button.
	NSFontPanelShadowEffectModeMask uint = 4096
	// NSFontPanelSizeModeMask: Display the font size column.
	NSFontPanelSizeModeMask uint = 2
	// NSFontPanelStandardModesMask: Display the standard default font panel—that is, including the collections, typeface, and size columns.
	NSFontPanelStandardModesMask uint = 65535
	// NSFontPanelStrikethroughEffectModeMask: Display the strike-through popup menu.
	NSFontPanelStrikethroughEffectModeMask uint = 512
	// NSFontPanelTextColorEffectModeMask: Display the text color button.
	NSFontPanelTextColorEffectModeMask uint = 1024
	// NSFontPanelUnderlineEffectModeMask: Display the underline popup menu.
	NSFontPanelUnderlineEffectModeMask uint = 256
)


// See: https://developer.apple.com/documentation/AppKit/NSFontPanel/ModeMask
type NSFontPanelModeMask int

const (
	NSFontPanelModeMaskAllEffects NSFontPanelModeMask = 1048320
	NSFontPanelModeMaskCollection NSFontPanelModeMask = 4
	NSFontPanelModeMaskDocumentColorEffect NSFontPanelModeMask = 2048
	NSFontPanelModeMaskFace NSFontPanelModeMask = 1
	NSFontPanelModeMaskShadowEffect NSFontPanelModeMask = 4096
	NSFontPanelModeMaskSize NSFontPanelModeMask = 2
	NSFontPanelModeMaskStrikethroughEffect NSFontPanelModeMask = 512
	NSFontPanelModeMaskTextColorEffect NSFontPanelModeMask = 1024
	NSFontPanelModeMaskUnderlineEffect NSFontPanelModeMask = 256
	NSFontPanelModesMaskAllModes NSFontPanelModeMask = 4294967295
	NSFontPanelModesMaskStandardModes NSFontPanelModeMask = 65535
)


func (e NSFontPanelModeMask) String() string {
	switch e {
	case NSFontPanelModeMaskAllEffects:
		return "NSFontPanelModeMaskAllEffects"
	case NSFontPanelModeMaskCollection:
		return "NSFontPanelModeMaskCollection"
	case NSFontPanelModeMaskDocumentColorEffect:
		return "NSFontPanelModeMaskDocumentColorEffect"
	case NSFontPanelModeMaskFace:
		return "NSFontPanelModeMaskFace"
	case NSFontPanelModeMaskShadowEffect:
		return "NSFontPanelModeMaskShadowEffect"
	case NSFontPanelModeMaskSize:
		return "NSFontPanelModeMaskSize"
	case NSFontPanelModeMaskStrikethroughEffect:
		return "NSFontPanelModeMaskStrikethroughEffect"
	case NSFontPanelModeMaskTextColorEffect:
		return "NSFontPanelModeMaskTextColorEffect"
	case NSFontPanelModeMaskUnderlineEffect:
		return "NSFontPanelModeMaskUnderlineEffect"
	case NSFontPanelModesMaskAllModes:
		return "NSFontPanelModesMaskAllModes"
	case NSFontPanelModesMaskStandardModes:
		return "NSFontPanelModesMaskStandardModes"
	default:
		return fmt.Sprintf("NSFontPanelModeMask(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSFontRenderingMode
type NSFontRenderingMode int

const (
	// NSFontAntialiasedIntegerAdvancementsRenderingMode: Specifies antialiased, integer advancements rendering mode.
	NSFontAntialiasedIntegerAdvancementsRenderingMode NSFontRenderingMode = 3
	// NSFontAntialiasedRenderingMode: Specifies antialiased, floating-point advancements rendering mode (synonymous with printerFont).
	NSFontAntialiasedRenderingMode NSFontRenderingMode = 1
	// NSFontDefaultRenderingMode: Determines the actual mode based on the user preference settings.
	NSFontDefaultRenderingMode NSFontRenderingMode = 0
	// NSFontIntegerAdvancementsRenderingMode: Specifies integer advancements rendering mode.
	NSFontIntegerAdvancementsRenderingMode NSFontRenderingMode = 2
)


func (e NSFontRenderingMode) String() string {
	switch e {
	case NSFontAntialiasedIntegerAdvancementsRenderingMode:
		return "NSFontAntialiasedIntegerAdvancementsRenderingMode"
	case NSFontAntialiasedRenderingMode:
		return "NSFontAntialiasedRenderingMode"
	case NSFontDefaultRenderingMode:
		return "NSFontDefaultRenderingMode"
	case NSFontIntegerAdvancementsRenderingMode:
		return "NSFontIntegerAdvancementsRenderingMode"
	default:
		return fmt.Sprintf("NSFontRenderingMode(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSFontTraitMask
type NSFontTraitMask int

const (
	// NSBoldFontMask: A mask that specifies a bold font.
	NSBoldFontMask NSFontTraitMask = 2
	// NSCompressedFontMask: A mask that specifies a compressed font.
	NSCompressedFontMask NSFontTraitMask = 512
	// NSCondensedFontMask: A mask that specifies a condensed font.
	NSCondensedFontMask NSFontTraitMask = 64
	// NSExpandedFontMask: A mask that specifies an expanded font.
	NSExpandedFontMask NSFontTraitMask = 32
	// NSFixedPitchFontMask: A mask that specifies a fixed pitch font.
	NSFixedPitchFontMask NSFontTraitMask = 1024
	// NSItalicFontMask: A mask that specifies an italic font.
	NSItalicFontMask NSFontTraitMask = 1
	// NSNarrowFontMask: A mask that specifies a narrow font.
	NSNarrowFontMask NSFontTraitMask = 16
	// NSNonStandardCharacterSetFontMask: A mask that specifies a font containing a non-standard character set.
	NSNonStandardCharacterSetFontMask NSFontTraitMask = 8
	// NSPosterFontMask: A mask that specifies a poster-style font.
	NSPosterFontMask NSFontTraitMask = 256
	// NSSmallCapsFontMask: A mask that specifies a small-caps font.
	NSSmallCapsFontMask NSFontTraitMask = 128
	// NSUnboldFontMask: A mask that specifies a font that is not bold.
	NSUnboldFontMask NSFontTraitMask = 4
	// NSUnitalicFontMask: A mask that specifies a font that is not italic.
	NSUnitalicFontMask NSFontTraitMask = 16777216
)


func (e NSFontTraitMask) String() string {
	switch e {
	case NSBoldFontMask:
		return "NSBoldFontMask"
	case NSCompressedFontMask:
		return "NSCompressedFontMask"
	case NSCondensedFontMask:
		return "NSCondensedFontMask"
	case NSExpandedFontMask:
		return "NSExpandedFontMask"
	case NSFixedPitchFontMask:
		return "NSFixedPitchFontMask"
	case NSItalicFontMask:
		return "NSItalicFontMask"
	case NSNarrowFontMask:
		return "NSNarrowFontMask"
	case NSNonStandardCharacterSetFontMask:
		return "NSNonStandardCharacterSetFontMask"
	case NSPosterFontMask:
		return "NSPosterFontMask"
	case NSSmallCapsFontMask:
		return "NSSmallCapsFontMask"
	case NSUnboldFontMask:
		return "NSUnboldFontMask"
	case NSUnitalicFontMask:
		return "NSUnitalicFontMask"
	default:
		return fmt.Sprintf("NSFontTraitMask(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/State-swift.enum
type NSGestureRecognizerState int

const (
	// NSGestureRecognizerStateBegan: The gesture recognizer has recognized a sequence of events as a continuous gesture.
	NSGestureRecognizerStateBegan NSGestureRecognizerState = 1
	// NSGestureRecognizerStateCancelled: The gesture recognizer received events that resulted in the cancellation of a continuous gesture.
	NSGestureRecognizerStateCancelled NSGestureRecognizerState = 4
	// NSGestureRecognizerStateChanged: The gesture recognizer has detected a change to a continuous gesture.
	NSGestureRecognizerStateChanged NSGestureRecognizerState = 2
	// NSGestureRecognizerStateEnded: The gesture recognizer has detected the end of a continuous gesture.
	NSGestureRecognizerStateEnded NSGestureRecognizerState = 3
	// NSGestureRecognizerStateFailed: The gesture recognizer failed to recognize its gesture and will not call its action method.
	NSGestureRecognizerStateFailed NSGestureRecognizerState = 5
	// NSGestureRecognizerStatePossible: The gesture recognizer has not yet recognized its gesture but may be evaluating events.
	NSGestureRecognizerStatePossible NSGestureRecognizerState = 0
)


func (e NSGestureRecognizerState) String() string {
	switch e {
	case NSGestureRecognizerStateBegan:
		return "NSGestureRecognizerStateBegan"
	case NSGestureRecognizerStateCancelled:
		return "NSGestureRecognizerStateCancelled"
	case NSGestureRecognizerStateChanged:
		return "NSGestureRecognizerStateChanged"
	case NSGestureRecognizerStateEnded:
		return "NSGestureRecognizerStateEnded"
	case NSGestureRecognizerStateFailed:
		return "NSGestureRecognizerStateFailed"
	case NSGestureRecognizerStatePossible:
		return "NSGestureRecognizerStatePossible"
	default:
		return fmt.Sprintf("NSGestureRecognizerState(%d)", e)
	}
}




type NSGlyphAttribute uint

const (
	// Deprecated.
	NSGlyphAttributeBidiLevel NSGlyphAttribute = 2
	// Deprecated.
	NSGlyphAttributeElastic NSGlyphAttribute = 1
	// Deprecated.
	NSGlyphAttributeInscribe NSGlyphAttribute = 5
	// Deprecated.
	NSGlyphAttributeSoft NSGlyphAttribute = 0
)


func (e NSGlyphAttribute) String() string {
	switch e {
	case NSGlyphAttributeBidiLevel:
		return "NSGlyphAttributeBidiLevel"
	case NSGlyphAttributeElastic:
		return "NSGlyphAttributeElastic"
	case NSGlyphAttributeInscribe:
		return "NSGlyphAttributeInscribe"
	case NSGlyphAttributeSoft:
		return "NSGlyphAttributeSoft"
	default:
		return fmt.Sprintf("NSGlyphAttribute(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSGlyphInscription
type NSGlyphInscription uint

const (
	NSGlyphInscribeAbove NSGlyphInscription = 2
	NSGlyphInscribeBase NSGlyphInscription = 0
	NSGlyphInscribeBelow NSGlyphInscription = 1
	NSGlyphInscribeOverBelow NSGlyphInscription = 4
	NSGlyphInscribeOverstrike NSGlyphInscription = 3
)


func (e NSGlyphInscription) String() string {
	switch e {
	case NSGlyphInscribeAbove:
		return "NSGlyphInscribeAbove"
	case NSGlyphInscribeBase:
		return "NSGlyphInscribeBase"
	case NSGlyphInscribeBelow:
		return "NSGlyphInscribeBelow"
	case NSGlyphInscribeOverBelow:
		return "NSGlyphInscribeOverBelow"
	case NSGlyphInscribeOverstrike:
		return "NSGlyphInscribeOverstrike"
	default:
		return fmt.Sprintf("NSGlyphInscription(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/GlyphProperty
type NSGlyphProperty int

const (
	// NSGlyphPropertyControlCharacter: A glyph representing a control character.
	NSGlyphPropertyControlCharacter NSGlyphProperty = 2
	// NSGlyphPropertyElastic: A glyph with a changeable width, such as a white space character.
	NSGlyphPropertyElastic NSGlyphProperty = 4
	// NSGlyphPropertyNonBaseCharacter: A glyph that combines several properties.
	NSGlyphPropertyNonBaseCharacter NSGlyphProperty = 8
	// NSGlyphPropertyNull: The null glyph, which the layout manager ignores.
	NSGlyphPropertyNull NSGlyphProperty = 1
)


func (e NSGlyphProperty) String() string {
	switch e {
	case NSGlyphPropertyControlCharacter:
		return "NSGlyphPropertyControlCharacter"
	case NSGlyphPropertyElastic:
		return "NSGlyphPropertyElastic"
	case NSGlyphPropertyNonBaseCharacter:
		return "NSGlyphPropertyNonBaseCharacter"
	case NSGlyphPropertyNull:
		return "NSGlyphPropertyNull"
	default:
		return fmt.Sprintf("NSGlyphProperty(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSGradient/DrawingOptions
type NSGradientDrawingOptions int

const (
	// NSGradientDrawsAfterEndingLocation: Drawing extends beyond the gradient end point.
	NSGradientDrawsAfterEndingLocation NSGradientDrawingOptions = 2
	// NSGradientDrawsBeforeStartingLocation: Drawing extends before the gradient starting point.
	NSGradientDrawsBeforeStartingLocation NSGradientDrawingOptions = 1
)


func (e NSGradientDrawingOptions) String() string {
	switch e {
	case NSGradientDrawsAfterEndingLocation:
		return "NSGradientDrawsAfterEndingLocation"
	case NSGradientDrawsBeforeStartingLocation:
		return "NSGradientDrawsBeforeStartingLocation"
	default:
		return fmt.Sprintf("NSGradientDrawingOptions(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSButton/GradientType
type NSGradientType int

const (
	// NSGradientConcaveStrong: As with , the top-left corner is light gray, and the bottom-right corner is dark gray, but the difference between the grays is greater, so the appearance of being pushed in is stronger.
	NSGradientConcaveStrong NSGradientType = 2
	// NSGradientConcaveWeak: The top-left corner is light gray, and the bottom-right corner is dark gray, so the button appears to be pushed in.
	NSGradientConcaveWeak NSGradientType = 1
	// NSGradientConvexStrong: As with , the top-left corner is dark gray, and the bottom-right corner is light gray, but the difference between the grays is greater, so the appearance of sticking out is stronger.
	NSGradientConvexStrong NSGradientType = 4
	// NSGradientConvexWeak: The top-left corner is dark gray, and the bottom-right corner is light gray, so the button appears to be sticking out.
	NSGradientConvexWeak NSGradientType = 3
	// NSGradientNone: There is no gradient, so the button looks flat.
	NSGradientNone NSGradientType = 0
)


func (e NSGradientType) String() string {
	switch e {
	case NSGradientConcaveStrong:
		return "NSGradientConcaveStrong"
	case NSGradientConcaveWeak:
		return "NSGradientConcaveWeak"
	case NSGradientConvexStrong:
		return "NSGradientConvexStrong"
	case NSGradientConvexWeak:
		return "NSGradientConvexWeak"
	case NSGradientNone:
		return "NSGradientNone"
	default:
		return fmt.Sprintf("NSGradientType(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSGridCell/Placement
type NSGridCellPlacement int

const (
	NSGridCellPlacementBottom NSGridCellPlacement = 3
	NSGridCellPlacementCenter NSGridCellPlacement = 4
	NSGridCellPlacementFill NSGridCellPlacement = 5
	NSGridCellPlacementInherited NSGridCellPlacement = 0
	NSGridCellPlacementLeading NSGridCellPlacement = 2
	NSGridCellPlacementNone NSGridCellPlacement = 1
	NSGridCellPlacementTop NSGridCellPlacement = 2
	NSGridCellPlacementTrailing NSGridCellPlacement = 3
)


func (e NSGridCellPlacement) String() string {
	switch e {
	case NSGridCellPlacementBottom:
		return "NSGridCellPlacementBottom"
	case NSGridCellPlacementCenter:
		return "NSGridCellPlacementCenter"
	case NSGridCellPlacementFill:
		return "NSGridCellPlacementFill"
	case NSGridCellPlacementInherited:
		return "NSGridCellPlacementInherited"
	case NSGridCellPlacementLeading:
		return "NSGridCellPlacementLeading"
	case NSGridCellPlacementNone:
		return "NSGridCellPlacementNone"
	default:
		return fmt.Sprintf("NSGridCellPlacement(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSGridRow/Alignment
type NSGridRowAlignment int

const (
	NSGridRowAlignmentFirstBaseline NSGridRowAlignment = 2
	NSGridRowAlignmentInherited NSGridRowAlignment = 0
	NSGridRowAlignmentLastBaseline NSGridRowAlignment = 3
	NSGridRowAlignmentNone NSGridRowAlignment = 1
)


func (e NSGridRowAlignment) String() string {
	switch e {
	case NSGridRowAlignmentFirstBaseline:
		return "NSGridRowAlignmentFirstBaseline"
	case NSGridRowAlignmentInherited:
		return "NSGridRowAlignmentInherited"
	case NSGridRowAlignmentLastBaseline:
		return "NSGridRowAlignmentLastBaseline"
	case NSGridRowAlignmentNone:
		return "NSGridRowAlignmentNone"
	default:
		return fmt.Sprintf("NSGridRowAlignment(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSHapticFeedbackManager/FeedbackPattern
type NSHapticFeedbackPattern int

const (
	// NSHapticFeedbackPatternAlignment: A haptic feedback pattern to be used in response to the alignment of an object the user is dragging around.
	NSHapticFeedbackPatternAlignment NSHapticFeedbackPattern = 1
	// NSHapticFeedbackPatternGeneric: A general haptic feedback pattern.
	NSHapticFeedbackPatternGeneric NSHapticFeedbackPattern = 0
	// NSHapticFeedbackPatternLevelChange: A haptic feedback pattern to be used as the user moves between discrete levels of pressure.
	NSHapticFeedbackPatternLevelChange NSHapticFeedbackPattern = 2
)


func (e NSHapticFeedbackPattern) String() string {
	switch e {
	case NSHapticFeedbackPatternAlignment:
		return "NSHapticFeedbackPatternAlignment"
	case NSHapticFeedbackPatternGeneric:
		return "NSHapticFeedbackPatternGeneric"
	case NSHapticFeedbackPatternLevelChange:
		return "NSHapticFeedbackPatternLevelChange"
	default:
		return fmt.Sprintf("NSHapticFeedbackPattern(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSHapticFeedbackManager/PerformanceTime
type NSHapticFeedbackPerformanceTime int

const (
	// NSHapticFeedbackPerformanceTimeDefault: Allows the system to choose the most appropriate time for feedback to be provided.
	NSHapticFeedbackPerformanceTimeDefault NSHapticFeedbackPerformanceTime = 0
	// NSHapticFeedbackPerformanceTimeDrawCompleted: Instructs the system to provide haptic feedback to the user the next time the screen is updated.
	NSHapticFeedbackPerformanceTimeDrawCompleted NSHapticFeedbackPerformanceTime = 2
	// NSHapticFeedbackPerformanceTimeNow: Instructs the system to provide immediate haptic feedback to the user, rather than waiting for synchronization to occur with something visual occurring on screen.
	NSHapticFeedbackPerformanceTimeNow NSHapticFeedbackPerformanceTime = 1
)


func (e NSHapticFeedbackPerformanceTime) String() string {
	switch e {
	case NSHapticFeedbackPerformanceTimeDefault:
		return "NSHapticFeedbackPerformanceTimeDefault"
	case NSHapticFeedbackPerformanceTimeDrawCompleted:
		return "NSHapticFeedbackPerformanceTimeDrawCompleted"
	case NSHapticFeedbackPerformanceTimeNow:
		return "NSHapticFeedbackPerformanceTimeNow"
	default:
		return fmt.Sprintf("NSHapticFeedbackPerformanceTime(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSHorizontalDirections
type NSHorizontalDirections uint

const (
	NSHorizontalDirectionsLeft NSHorizontalDirections = 1
	NSHorizontalDirectionsRight NSHorizontalDirections = 2
)


func (e NSHorizontalDirections) String() string {
	switch e {
	case NSHorizontalDirectionsLeft:
		return "NSHorizontalDirectionsLeft"
	case NSHorizontalDirectionsRight:
		return "NSHorizontalDirectionsRight"
	default:
		return fmt.Sprintf("NSHorizontalDirections(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSImageAlignment
type NSImageAlignment int

const (
	NSImageAlignBottom NSImageAlignment = 5
	NSImageAlignBottomLeft NSImageAlignment = 6
	NSImageAlignBottomRight NSImageAlignment = 7
	NSImageAlignCenter NSImageAlignment = 0
	NSImageAlignLeft NSImageAlignment = 4
	NSImageAlignRight NSImageAlignment = 8
	NSImageAlignTop NSImageAlignment = 1
	NSImageAlignTopLeft NSImageAlignment = 2
	NSImageAlignTopRight NSImageAlignment = 3
)


func (e NSImageAlignment) String() string {
	switch e {
	case NSImageAlignBottom:
		return "NSImageAlignBottom"
	case NSImageAlignBottomLeft:
		return "NSImageAlignBottomLeft"
	case NSImageAlignBottomRight:
		return "NSImageAlignBottomRight"
	case NSImageAlignCenter:
		return "NSImageAlignCenter"
	case NSImageAlignLeft:
		return "NSImageAlignLeft"
	case NSImageAlignRight:
		return "NSImageAlignRight"
	case NSImageAlignTop:
		return "NSImageAlignTop"
	case NSImageAlignTopLeft:
		return "NSImageAlignTopLeft"
	case NSImageAlignTopRight:
		return "NSImageAlignTopRight"
	default:
		return fmt.Sprintf("NSImageAlignment(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSImage/CacheMode-swift.enum
type NSImageCacheMode int

const (
	// NSImageCacheAlways: Always generate a cache when drawing.
	NSImageCacheAlways NSImageCacheMode = 1
	// NSImageCacheBySize: Cache if the cache size is smaller than the original data.
	NSImageCacheBySize NSImageCacheMode = 2
	// NSImageCacheDefault: Caching is unspecified.
	NSImageCacheDefault NSImageCacheMode = 0
	// NSImageCacheNever: Never cache; always draw direct.
	NSImageCacheNever NSImageCacheMode = 3
)


func (e NSImageCacheMode) String() string {
	switch e {
	case NSImageCacheAlways:
		return "NSImageCacheAlways"
	case NSImageCacheBySize:
		return "NSImageCacheBySize"
	case NSImageCacheDefault:
		return "NSImageCacheDefault"
	case NSImageCacheNever:
		return "NSImageCacheNever"
	default:
		return fmt.Sprintf("NSImageCacheMode(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSImage/DynamicRange
type NSImageDynamicRange int

const (
	// NSImageDynamicRangeConstrainedHigh: Allows for constrained High Dynamic Range (HDR) image content which is useful for mixing HDR and Standard Dynamic Range (SDR) content.
	NSImageDynamicRangeConstrainedHigh NSImageDynamicRange = 1
	// NSImageDynamicRangeHigh: Allows image content to use extended dynamic range if it has dynamic range content.
	NSImageDynamicRangeHigh NSImageDynamicRange = 2
	// NSImageDynamicRangeStandard: Restricts the image content dynamic range to the standard range regardless of the actual range of the image content.
	NSImageDynamicRangeStandard NSImageDynamicRange = 0
	// NSImageDynamicRangeUnspecified: Indicates that the dynamic range treatment of the image is unknown or otherwise unspecified.
	NSImageDynamicRangeUnspecified NSImageDynamicRange = -1
)


func (e NSImageDynamicRange) String() string {
	switch e {
	case NSImageDynamicRangeConstrainedHigh:
		return "NSImageDynamicRangeConstrainedHigh"
	case NSImageDynamicRangeHigh:
		return "NSImageDynamicRangeHigh"
	case NSImageDynamicRangeStandard:
		return "NSImageDynamicRangeStandard"
	case NSImageDynamicRangeUnspecified:
		return "NSImageDynamicRangeUnspecified"
	default:
		return fmt.Sprintf("NSImageDynamicRange(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSImageView/FrameStyle
type NSImageFrameStyle int

const (
	// NSImageFrameButton: A convex bezel that makes the image stand out in relief, like a button.
	NSImageFrameButton NSImageFrameStyle = 4
	// NSImageFrameGrayBezel: A gray, concave bezel that makes the image look sunken.
	NSImageFrameGrayBezel NSImageFrameStyle = 2
	// NSImageFrameGroove: A thin groove that looks etched around the image.
	NSImageFrameGroove NSImageFrameStyle = 3
	// NSImageFrameNone: An invisible frame
	NSImageFrameNone NSImageFrameStyle = 0
	// NSImageFramePhoto: A thin black outline and a dropped shadow.
	NSImageFramePhoto NSImageFrameStyle = 1
)


func (e NSImageFrameStyle) String() string {
	switch e {
	case NSImageFrameButton:
		return "NSImageFrameButton"
	case NSImageFrameGrayBezel:
		return "NSImageFrameGrayBezel"
	case NSImageFrameGroove:
		return "NSImageFrameGroove"
	case NSImageFrameNone:
		return "NSImageFrameNone"
	case NSImageFramePhoto:
		return "NSImageFramePhoto"
	default:
		return fmt.Sprintf("NSImageFrameStyle(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSImageInterpolation
type NSImageInterpolation int

const (
	// NSImageInterpolationDefault: Use the context’s default interpolation.
	NSImageInterpolationDefault NSImageInterpolation = 0
	// NSImageInterpolationHigh: Highest quality, slower than the medium interpolation option.
	NSImageInterpolationHigh NSImageInterpolation = 3
	// NSImageInterpolationLow: Fast, low-quality interpolation.
	NSImageInterpolationLow NSImageInterpolation = 2
	// NSImageInterpolationMedium: Medium quality, slower than the low interpolation option.
	NSImageInterpolationMedium NSImageInterpolation = 4
	// NSImageInterpolationNone: No interpolation.
	NSImageInterpolationNone NSImageInterpolation = 1
)


func (e NSImageInterpolation) String() string {
	switch e {
	case NSImageInterpolationDefault:
		return "NSImageInterpolationDefault"
	case NSImageInterpolationHigh:
		return "NSImageInterpolationHigh"
	case NSImageInterpolationLow:
		return "NSImageInterpolationLow"
	case NSImageInterpolationMedium:
		return "NSImageInterpolationMedium"
	case NSImageInterpolationNone:
		return "NSImageInterpolationNone"
	default:
		return fmt.Sprintf("NSImageInterpolation(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSImage/LayoutDirection
type NSImageLayoutDirection int

const (
	// NSImageLayoutDirectionLeftToRight: A left-to-right layout direction.
	NSImageLayoutDirectionLeftToRight NSImageLayoutDirection = 2
	// NSImageLayoutDirectionRightToLeft: A right-to-left layout direction.
	NSImageLayoutDirectionRightToLeft NSImageLayoutDirection = 3
	// NSImageLayoutDirectionUnspecified: An unspecified layout direction.
	NSImageLayoutDirectionUnspecified NSImageLayoutDirection = -1
)


func (e NSImageLayoutDirection) String() string {
	switch e {
	case NSImageLayoutDirectionLeftToRight:
		return "NSImageLayoutDirectionLeftToRight"
	case NSImageLayoutDirectionRightToLeft:
		return "NSImageLayoutDirectionRightToLeft"
	case NSImageLayoutDirectionUnspecified:
		return "NSImageLayoutDirectionUnspecified"
	default:
		return fmt.Sprintf("NSImageLayoutDirection(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSImage/LoadStatus
type NSImageLoadStatus int

const (
	// NSImageLoadStatusCancelled: Image loading was canceled.
	NSImageLoadStatusCancelled NSImageLoadStatus = 1
	// NSImageLoadStatusCompleted: Enough data is available to completely decompress the image.
	NSImageLoadStatusCompleted NSImageLoadStatus = 0
	// NSImageLoadStatusInvalidData: An error occurred during image decompression.
	NSImageLoadStatusInvalidData NSImageLoadStatus = 2
	// NSImageLoadStatusReadError: Not enough data was available for full decompression of the image.
	NSImageLoadStatusReadError NSImageLoadStatus = 4
	// NSImageLoadStatusUnexpectedEOF: Not enough data was available to fully decompress the image.
	NSImageLoadStatusUnexpectedEOF NSImageLoadStatus = 3
)


func (e NSImageLoadStatus) String() string {
	switch e {
	case NSImageLoadStatusCancelled:
		return "NSImageLoadStatusCancelled"
	case NSImageLoadStatusCompleted:
		return "NSImageLoadStatusCompleted"
	case NSImageLoadStatusInvalidData:
		return "NSImageLoadStatusInvalidData"
	case NSImageLoadStatusReadError:
		return "NSImageLoadStatusReadError"
	case NSImageLoadStatusUnexpectedEOF:
		return "NSImageLoadStatusUnexpectedEOF"
	default:
		return fmt.Sprintf("NSImageLoadStatus(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/LoadStatus
type NSImageRepLoadStatus int

const (
	// NSImageRepLoadStatusCompleted: Enough data has been provided to successfully decompress the image (regardless of the complete: flag).
	NSImageRepLoadStatusCompleted NSImageRepLoadStatus = -6
	// NSImageRepLoadStatusInvalidData: An error occurred during image decompression.
	NSImageRepLoadStatusInvalidData NSImageRepLoadStatus = -4
	// NSImageRepLoadStatusReadingHeader: The image format is known, but not enough data has been read to determine the size, depth, etc., of the image.
	NSImageRepLoadStatusReadingHeader NSImageRepLoadStatus = -2
	// NSImageRepLoadStatusUnexpectedEOF: Not enough data was available to fully decompress the image.
	NSImageRepLoadStatusUnexpectedEOF NSImageRepLoadStatus = -5
	// NSImageRepLoadStatusUnknownType: Not enough data to determine image format.
	NSImageRepLoadStatusUnknownType NSImageRepLoadStatus = -1
	// NSImageRepLoadStatusWillNeedAllData: Incremental loading cannot be supported.
	NSImageRepLoadStatusWillNeedAllData NSImageRepLoadStatus = -3
)


func (e NSImageRepLoadStatus) String() string {
	switch e {
	case NSImageRepLoadStatusCompleted:
		return "NSImageRepLoadStatusCompleted"
	case NSImageRepLoadStatusInvalidData:
		return "NSImageRepLoadStatusInvalidData"
	case NSImageRepLoadStatusReadingHeader:
		return "NSImageRepLoadStatusReadingHeader"
	case NSImageRepLoadStatusUnexpectedEOF:
		return "NSImageRepLoadStatusUnexpectedEOF"
	case NSImageRepLoadStatusUnknownType:
		return "NSImageRepLoadStatusUnknownType"
	case NSImageRepLoadStatusWillNeedAllData:
		return "NSImageRepLoadStatusWillNeedAllData"
	default:
		return fmt.Sprintf("NSImageRepLoadStatus(%d)", e)
	}
}




type NSImageRepMatchesDeviceConstants uint

const (
	// NSImageRepMatchesDevice: A constant indicating that the value of certain attributes, such as the number of colors or bits per sample, will change to match the display device.
	NSImageRepMatchesDevice NSImageRepMatchesDeviceConstants = 0
)


func (e NSImageRepMatchesDeviceConstants) String() string {
	switch e {
	case NSImageRepMatchesDevice:
		return "NSImageRepMatchesDevice"
	default:
		return fmt.Sprintf("NSImageRepMatchesDeviceConstants(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSImage/ResizingMode-swift.enum
type NSImageResizingMode int

const (
	// NSImageResizingModeStretch: The image stretches when it resizes.
	NSImageResizingModeStretch NSImageResizingMode = 1
	// NSImageResizingModeTile: The image tiles when it resizes.
	NSImageResizingModeTile NSImageResizingMode = 0
)


func (e NSImageResizingMode) String() string {
	switch e {
	case NSImageResizingModeStretch:
		return "NSImageResizingModeStretch"
	case NSImageResizingModeTile:
		return "NSImageResizingModeTile"
	default:
		return fmt.Sprintf("NSImageResizingMode(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSImageScaling
type NSImageScaling int

const (
	// NSImageScaleAxesIndependently: Scale each dimension to exactly fit destination.
	NSImageScaleAxesIndependently NSImageScaling = 1
	// NSImageScaleNone: Do not scale the image.
	NSImageScaleNone NSImageScaling = 2
	// NSImageScaleProportionallyDown: If it is too large for the destination, scale the image down while preserving the aspect ratio.
	NSImageScaleProportionallyDown NSImageScaling = 0
	// NSImageScaleProportionallyUpOrDown: Scale the image to its maximum possible dimensions while both staying within the destination area and preserving its aspect ratio.
	NSImageScaleProportionallyUpOrDown NSImageScaling = 3
)


func (e NSImageScaling) String() string {
	switch e {
	case NSImageScaleAxesIndependently:
		return "NSImageScaleAxesIndependently"
	case NSImageScaleNone:
		return "NSImageScaleNone"
	case NSImageScaleProportionallyDown:
		return "NSImageScaleProportionallyDown"
	case NSImageScaleProportionallyUpOrDown:
		return "NSImageScaleProportionallyUpOrDown"
	default:
		return fmt.Sprintf("NSImageScaling(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSImage/SymbolColorRenderingMode
type NSImageSymbolColorRenderingMode int

const (
	// NSImageSymbolColorRenderingModeAutomatic: Automatically uses an appropriate color rendering mode for the symbol’s color layers.
	NSImageSymbolColorRenderingModeAutomatic NSImageSymbolColorRenderingMode = 0
	// NSImageSymbolColorRenderingModeFlat: Renders the symbol’s color layers using flat colors.
	NSImageSymbolColorRenderingModeFlat NSImageSymbolColorRenderingMode = 1
	// NSImageSymbolColorRenderingModeGradient: Renders the symbol’s color layers using gradients.
	NSImageSymbolColorRenderingModeGradient NSImageSymbolColorRenderingMode = 2
)


func (e NSImageSymbolColorRenderingMode) String() string {
	switch e {
	case NSImageSymbolColorRenderingModeAutomatic:
		return "NSImageSymbolColorRenderingModeAutomatic"
	case NSImageSymbolColorRenderingModeFlat:
		return "NSImageSymbolColorRenderingModeFlat"
	case NSImageSymbolColorRenderingModeGradient:
		return "NSImageSymbolColorRenderingModeGradient"
	default:
		return fmt.Sprintf("NSImageSymbolColorRenderingMode(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSImage/SymbolScale
type NSImageSymbolScale int

const (
	// NSImageSymbolScaleLarge: The symbol uses the large scale variant.
	NSImageSymbolScaleLarge NSImageSymbolScale = 3
	// NSImageSymbolScaleMedium: The symbol uses the default medium scale variant.
	NSImageSymbolScaleMedium NSImageSymbolScale = 2
	// NSImageSymbolScaleSmall: The symbol uses the small scale variant.
	NSImageSymbolScaleSmall NSImageSymbolScale = 1
)


func (e NSImageSymbolScale) String() string {
	switch e {
	case NSImageSymbolScaleLarge:
		return "NSImageSymbolScaleLarge"
	case NSImageSymbolScaleMedium:
		return "NSImageSymbolScaleMedium"
	case NSImageSymbolScaleSmall:
		return "NSImageSymbolScaleSmall"
	default:
		return fmt.Sprintf("NSImageSymbolScale(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSImage/SymbolVariableValueMode
type NSImageSymbolVariableValueMode int

const (
	// NSImageSymbolVariableValueModeAutomatic: Automatically selects an appropriate variable value mode for the symbol.
	NSImageSymbolVariableValueModeAutomatic NSImageSymbolVariableValueMode = 0
	// NSImageSymbolVariableValueModeColor: The “color” variable value mode.
	NSImageSymbolVariableValueModeColor NSImageSymbolVariableValueMode = 1
	// NSImageSymbolVariableValueModeDraw: The “draw” variable value mode.
	NSImageSymbolVariableValueModeDraw NSImageSymbolVariableValueMode = 2
)


func (e NSImageSymbolVariableValueMode) String() string {
	switch e {
	case NSImageSymbolVariableValueModeAutomatic:
		return "NSImageSymbolVariableValueModeAutomatic"
	case NSImageSymbolVariableValueModeColor:
		return "NSImageSymbolVariableValueModeColor"
	case NSImageSymbolVariableValueModeDraw:
		return "NSImageSymbolVariableValueModeDraw"
	default:
		return fmt.Sprintf("NSImageSymbolVariableValueMode(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Attribute
type NSLayoutAttribute int

const (
	// NSLayoutAttributeBottom: The bottom of the object’s alignment rectangle.
	NSLayoutAttributeBottom NSLayoutAttribute = 4
	// NSLayoutAttributeCenterX: The center along the x-axis of the object’s alignment rectangle.
	NSLayoutAttributeCenterX NSLayoutAttribute = 9
	// NSLayoutAttributeCenterY: The center along the y-axis of the object’s alignment rectangle.
	NSLayoutAttributeCenterY NSLayoutAttribute = 10
	// NSLayoutAttributeFirstBaseline: The object’s baseline.
	NSLayoutAttributeFirstBaseline NSLayoutAttribute = 12
	// NSLayoutAttributeHeight: The height of the object’s alignment rectangle.
	NSLayoutAttributeHeight NSLayoutAttribute = 8
	// NSLayoutAttributeLastBaseline: The object’s baseline.
	NSLayoutAttributeLastBaseline NSLayoutAttribute = 11
	// NSLayoutAttributeLeading: The leading edge of the object’s alignment rectangle.
	NSLayoutAttributeLeading NSLayoutAttribute = 5
	// NSLayoutAttributeLeft: The left side of the object’s alignment rectangle.
	NSLayoutAttributeLeft NSLayoutAttribute = 1
	// NSLayoutAttributeNotAnAttribute: A placeholder value for indicating that the constraint’s second item and second attribute aren’t used in any calculations.
	NSLayoutAttributeNotAnAttribute NSLayoutAttribute = 0
	// NSLayoutAttributeRight: The right side of the object’s alignment rectangle.
	NSLayoutAttributeRight NSLayoutAttribute = 2
	// NSLayoutAttributeTop: The top of the object’s alignment rectangle.
	NSLayoutAttributeTop NSLayoutAttribute = 3
	// NSLayoutAttributeTrailing: The trailing edge of the object’s alignment rectangle.
	NSLayoutAttributeTrailing NSLayoutAttribute = 6
	// NSLayoutAttributeWidth: The width of the object’s alignment rectangle.
	NSLayoutAttributeWidth NSLayoutAttribute = 7
)


func (e NSLayoutAttribute) String() string {
	switch e {
	case NSLayoutAttributeBottom:
		return "NSLayoutAttributeBottom"
	case NSLayoutAttributeCenterX:
		return "NSLayoutAttributeCenterX"
	case NSLayoutAttributeCenterY:
		return "NSLayoutAttributeCenterY"
	case NSLayoutAttributeFirstBaseline:
		return "NSLayoutAttributeFirstBaseline"
	case NSLayoutAttributeHeight:
		return "NSLayoutAttributeHeight"
	case NSLayoutAttributeLastBaseline:
		return "NSLayoutAttributeLastBaseline"
	case NSLayoutAttributeLeading:
		return "NSLayoutAttributeLeading"
	case NSLayoutAttributeLeft:
		return "NSLayoutAttributeLeft"
	case NSLayoutAttributeNotAnAttribute:
		return "NSLayoutAttributeNotAnAttribute"
	case NSLayoutAttributeRight:
		return "NSLayoutAttributeRight"
	case NSLayoutAttributeTop:
		return "NSLayoutAttributeTop"
	case NSLayoutAttributeTrailing:
		return "NSLayoutAttributeTrailing"
	case NSLayoutAttributeWidth:
		return "NSLayoutAttributeWidth"
	default:
		return fmt.Sprintf("NSLayoutAttribute(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Orientation
type NSLayoutConstraintOrientation int

const (
	// NSLayoutConstraintOrientationHorizontal: The constraint orientation applied to laying out the horizontal relationship between objects.
	NSLayoutConstraintOrientationHorizontal NSLayoutConstraintOrientation = 0
	// NSLayoutConstraintOrientationVertical: The constraint orientation applied to laying out the vertical relationship between objects.
	NSLayoutConstraintOrientationVertical NSLayoutConstraintOrientation = 1
)


func (e NSLayoutConstraintOrientation) String() string {
	switch e {
	case NSLayoutConstraintOrientationHorizontal:
		return "NSLayoutConstraintOrientationHorizontal"
	case NSLayoutConstraintOrientationVertical:
		return "NSLayoutConstraintOrientationVertical"
	default:
		return fmt.Sprintf("NSLayoutConstraintOrientation(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/FormatOptions
type NSLayoutFormatOptions int

const (
	// NSLayoutFormatAlignAllBottom: Align all specified interface elements using NSLayoutConstraint.Attribute.bottom on each.
	NSLayoutFormatAlignAllBottom NSLayoutFormatOptions = 16
	// NSLayoutFormatAlignAllCenterX: Align all specified interface elements using NSLayoutConstraint.Attribute.centerX on each.
	NSLayoutFormatAlignAllCenterX NSLayoutFormatOptions = 512
	// NSLayoutFormatAlignAllCenterY: Align all specified interface elements using NSLayoutConstraint.Attribute.centerY on each.
	NSLayoutFormatAlignAllCenterY NSLayoutFormatOptions = 1024
	// NSLayoutFormatAlignAllFirstBaseline: Align all specified interface elements using the first baseline of each one.
	NSLayoutFormatAlignAllFirstBaseline NSLayoutFormatOptions = 4096
	// NSLayoutFormatAlignAllLastBaseline: Align all specified interface elements using the last baseline of each one.
	NSLayoutFormatAlignAllLastBaseline NSLayoutFormatOptions = 2048
	// NSLayoutFormatAlignAllLeading: Align all specified interface elements using NSLayoutConstraint.Attribute.leading on each.
	NSLayoutFormatAlignAllLeading NSLayoutFormatOptions = 32
	// NSLayoutFormatAlignAllLeft: Align all specified interface elements using NSLayoutConstraint.Attribute.left on each.
	NSLayoutFormatAlignAllLeft NSLayoutFormatOptions = 2
	// NSLayoutFormatAlignAllRight: Align all specified interface elements using NSLayoutConstraint.Attribute.right on each.
	NSLayoutFormatAlignAllRight NSLayoutFormatOptions = 4
	// NSLayoutFormatAlignAllTop: Align all specified interface elements using NSLayoutConstraint.Attribute.top on each.
	NSLayoutFormatAlignAllTop NSLayoutFormatOptions = 8
	// NSLayoutFormatAlignAllTrailing: Align all specified interface elements using NSLayoutConstraint.Attribute.trailing on each.
	NSLayoutFormatAlignAllTrailing NSLayoutFormatOptions = 64
	// NSLayoutFormatAlignmentMask: Bit mask that can be combined with an NSLayoutConstraint.FormatOptions variable to yield only the alignment portion of the format options.
	NSLayoutFormatAlignmentMask NSLayoutFormatOptions = 65535
	// NSLayoutFormatDirectionLeadingToTrailing: Arrange objects in order based on the normal text flow for the current user interface language.
	NSLayoutFormatDirectionLeadingToTrailing NSLayoutFormatOptions = 0
	// NSLayoutFormatDirectionLeftToRight: Arrange objects in order from left to right.
	NSLayoutFormatDirectionLeftToRight NSLayoutFormatOptions = 65536
	// NSLayoutFormatDirectionMask: A bit mask that can be combined with an NSLayoutConstraint.FormatOptions variable to yield only the direction portion of the format options.
	NSLayoutFormatDirectionMask NSLayoutFormatOptions = 3
	// NSLayoutFormatDirectionRightToLeft: Arrange objects in order from right to left.
	NSLayoutFormatDirectionRightToLeft NSLayoutFormatOptions = 131072
)


func (e NSLayoutFormatOptions) String() string {
	switch e {
	case NSLayoutFormatAlignAllBottom:
		return "NSLayoutFormatAlignAllBottom"
	case NSLayoutFormatAlignAllCenterX:
		return "NSLayoutFormatAlignAllCenterX"
	case NSLayoutFormatAlignAllCenterY:
		return "NSLayoutFormatAlignAllCenterY"
	case NSLayoutFormatAlignAllFirstBaseline:
		return "NSLayoutFormatAlignAllFirstBaseline"
	case NSLayoutFormatAlignAllLastBaseline:
		return "NSLayoutFormatAlignAllLastBaseline"
	case NSLayoutFormatAlignAllLeading:
		return "NSLayoutFormatAlignAllLeading"
	case NSLayoutFormatAlignAllLeft:
		return "NSLayoutFormatAlignAllLeft"
	case NSLayoutFormatAlignAllRight:
		return "NSLayoutFormatAlignAllRight"
	case NSLayoutFormatAlignAllTop:
		return "NSLayoutFormatAlignAllTop"
	case NSLayoutFormatAlignAllTrailing:
		return "NSLayoutFormatAlignAllTrailing"
	case NSLayoutFormatAlignmentMask:
		return "NSLayoutFormatAlignmentMask"
	case NSLayoutFormatDirectionLeadingToTrailing:
		return "NSLayoutFormatDirectionLeadingToTrailing"
	case NSLayoutFormatDirectionLeftToRight:
		return "NSLayoutFormatDirectionLeftToRight"
	case NSLayoutFormatDirectionMask:
		return "NSLayoutFormatDirectionMask"
	case NSLayoutFormatDirectionRightToLeft:
		return "NSLayoutFormatDirectionRightToLeft"
	default:
		return fmt.Sprintf("NSLayoutFormatOptions(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Relation-swift.enum
type NSLayoutRelation int

const (
	// NSLayoutRelationEqual: The constraint requires the first attribute to be exactly equal to the modified second attribute.
	NSLayoutRelationEqual NSLayoutRelation = 0
	// NSLayoutRelationGreaterThanOrEqual: The constraint requires the first attribute to be greater than or equal to the modified second attribute.
	NSLayoutRelationGreaterThanOrEqual NSLayoutRelation = 1
	// NSLayoutRelationLessThanOrEqual: The constraint requires the first attribute to be less than or equal to the modified second attribute.
	NSLayoutRelationLessThanOrEqual NSLayoutRelation = -1
)


func (e NSLayoutRelation) String() string {
	switch e {
	case NSLayoutRelationEqual:
		return "NSLayoutRelationEqual"
	case NSLayoutRelationGreaterThanOrEqual:
		return "NSLayoutRelationGreaterThanOrEqual"
	case NSLayoutRelationLessThanOrEqual:
		return "NSLayoutRelationLessThanOrEqual"
	default:
		return fmt.Sprintf("NSLayoutRelation(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/PlaceholderVisibility-swift.enum
type NSLevelIndicatorPlaceholderVisibility int

const (
	NSLevelIndicatorPlaceholderVisibilityAlways NSLevelIndicatorPlaceholderVisibility = 1
	NSLevelIndicatorPlaceholderVisibilityAutomatic NSLevelIndicatorPlaceholderVisibility = 0
	NSLevelIndicatorPlaceholderVisibilityWhileEditing NSLevelIndicatorPlaceholderVisibility = 2
)


func (e NSLevelIndicatorPlaceholderVisibility) String() string {
	switch e {
	case NSLevelIndicatorPlaceholderVisibilityAlways:
		return "NSLevelIndicatorPlaceholderVisibilityAlways"
	case NSLevelIndicatorPlaceholderVisibilityAutomatic:
		return "NSLevelIndicatorPlaceholderVisibilityAutomatic"
	case NSLevelIndicatorPlaceholderVisibilityWhileEditing:
		return "NSLevelIndicatorPlaceholderVisibilityWhileEditing"
	default:
		return fmt.Sprintf("NSLevelIndicatorPlaceholderVisibility(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/Style
type NSLevelIndicatorStyle int

const (
	// NSLevelIndicatorStyleContinuousCapacity: A style that indicates the capacity of something, such as how much data is on a hard disk.
	NSLevelIndicatorStyleContinuousCapacity NSLevelIndicatorStyle = 1
	// NSLevelIndicatorStyleDiscreteCapacity: A style that displays discrete segments that indicate the capacity of something, such as an audio level.
	NSLevelIndicatorStyleDiscreteCapacity NSLevelIndicatorStyle = 2
	// NSLevelIndicatorStyleRating: A style that indicates a rank, such as a star ranking display.
	NSLevelIndicatorStyleRating NSLevelIndicatorStyle = 3
	// NSLevelIndicatorStyleRelevancy: A style that indicates the relevancy of an item, such as a search result.
	NSLevelIndicatorStyleRelevancy NSLevelIndicatorStyle = 0
)


func (e NSLevelIndicatorStyle) String() string {
	switch e {
	case NSLevelIndicatorStyleContinuousCapacity:
		return "NSLevelIndicatorStyleContinuousCapacity"
	case NSLevelIndicatorStyleDiscreteCapacity:
		return "NSLevelIndicatorStyleDiscreteCapacity"
	case NSLevelIndicatorStyleRating:
		return "NSLevelIndicatorStyleRating"
	case NSLevelIndicatorStyleRelevancy:
		return "NSLevelIndicatorStyleRelevancy"
	default:
		return fmt.Sprintf("NSLevelIndicatorStyle(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSLineBreakMode
type NSLineBreakMode int

const (
	// NSLineBreakByCharWrapping: The value that indicates wrapping occurs before the first character that doesn’t fit.
	NSLineBreakByCharWrapping NSLineBreakMode = 1
	// NSLineBreakByClipping: The value that indicates lines don’t extend past the edge of the text container.
	NSLineBreakByClipping NSLineBreakMode = 2
	// NSLineBreakByTruncatingHead: The value that indicates that a line displays so that the end fits in the container and an ellipsis glyph indicates the missing text at the beginning of the line.
	NSLineBreakByTruncatingHead NSLineBreakMode = 3
	// NSLineBreakByTruncatingMiddle: The value that indicates that a line displays so that the beginning and end fit in the container and an ellipsis glyph indicates the missing text in the middle.
	NSLineBreakByTruncatingMiddle NSLineBreakMode = 5
	// NSLineBreakByTruncatingTail: The value that indicates a line displays so that the beginning fits in the container and an ellipsis glyph indicates the missing text at the end of the line.
	NSLineBreakByTruncatingTail NSLineBreakMode = 4
	// NSLineBreakByWordWrapping: The value that indicates wrapping occurs at word boundaries, unless the word doesn’t fit on a single line.
	NSLineBreakByWordWrapping NSLineBreakMode = 0
)


func (e NSLineBreakMode) String() string {
	switch e {
	case NSLineBreakByCharWrapping:
		return "NSLineBreakByCharWrapping"
	case NSLineBreakByClipping:
		return "NSLineBreakByClipping"
	case NSLineBreakByTruncatingHead:
		return "NSLineBreakByTruncatingHead"
	case NSLineBreakByTruncatingMiddle:
		return "NSLineBreakByTruncatingMiddle"
	case NSLineBreakByTruncatingTail:
		return "NSLineBreakByTruncatingTail"
	case NSLineBreakByWordWrapping:
		return "NSLineBreakByWordWrapping"
	default:
		return fmt.Sprintf("NSLineBreakMode(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/LineBreakStrategy-swift.struct
type NSLineBreakStrategy int

const (
	// NSLineBreakStrategyHangulWordPriority: The text system prohibits breaking between Hangul characters.
	NSLineBreakStrategyHangulWordPriority NSLineBreakStrategy = 2
	// NSLineBreakStrategyPushOut: The text system pushes out individual lines to avoid an orphan word on the last line of the paragraph.
	NSLineBreakStrategyPushOut NSLineBreakStrategy = 1
	// NSLineBreakStrategyStandard: The text system uses the same configuration of line-break strategies that it uses for standard UI labels.
	NSLineBreakStrategyStandard NSLineBreakStrategy = 65535
)


func (e NSLineBreakStrategy) String() string {
	switch e {
	case NSLineBreakStrategyHangulWordPriority:
		return "NSLineBreakStrategyHangulWordPriority"
	case NSLineBreakStrategyPushOut:
		return "NSLineBreakStrategyPushOut"
	case NSLineBreakStrategyStandard:
		return "NSLineBreakStrategyStandard"
	default:
		return fmt.Sprintf("NSLineBreakStrategy(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/LineCapStyle-swift.enum
type NSLineCapStyle int

const (
	// NSLineCapStyleButt: Specifies a butt line cap style for endpoints for an open path when stroked.
	NSLineCapStyleButt NSLineCapStyle = 0
	// NSLineCapStyleRound: Specifies a round line cap style for endpoints for an open path when stroked.
	NSLineCapStyleRound NSLineCapStyle = 1
	// NSLineCapStyleSquare: Specifies a square line cap style for endpoints for an open path when stroked.
	NSLineCapStyleSquare NSLineCapStyle = 2
)


func (e NSLineCapStyle) String() string {
	switch e {
	case NSLineCapStyleButt:
		return "NSLineCapStyleButt"
	case NSLineCapStyleRound:
		return "NSLineCapStyleRound"
	case NSLineCapStyleSquare:
		return "NSLineCapStyleSquare"
	default:
		return fmt.Sprintf("NSLineCapStyle(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/LineJoinStyle-swift.enum
type NSLineJoinStyle int

const (
	// NSLineJoinStyleBevel: Specifies a bevel line shape of the joints between connected segments of a stroked path.
	NSLineJoinStyleBevel NSLineJoinStyle = 2
	// NSLineJoinStyleMiter: Specifies a miter line shape of the joints between connected segments of a stroked path.
	NSLineJoinStyleMiter NSLineJoinStyle = 0
	// NSLineJoinStyleRound: Specifies a round line shape of the joints between connected segments of a stroked path.
	NSLineJoinStyleRound NSLineJoinStyle = 1
)


func (e NSLineJoinStyle) String() string {
	switch e {
	case NSLineJoinStyleBevel:
		return "NSLineJoinStyleBevel"
	case NSLineJoinStyleMiter:
		return "NSLineJoinStyleMiter"
	case NSLineJoinStyleRound:
		return "NSLineJoinStyleRound"
	default:
		return fmt.Sprintf("NSLineJoinStyle(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSLineMovementDirection
type NSLineMovementDirection uint

const (
	NSLineDoesntMove NSLineMovementDirection = 0
	NSLineMovesDown NSLineMovementDirection = 3
	NSLineMovesLeft NSLineMovementDirection = 1
	NSLineMovesRight NSLineMovementDirection = 2
	NSLineMovesUp NSLineMovementDirection = 4
)


func (e NSLineMovementDirection) String() string {
	switch e {
	case NSLineDoesntMove:
		return "NSLineDoesntMove"
	case NSLineMovesDown:
		return "NSLineMovesDown"
	case NSLineMovesLeft:
		return "NSLineMovesLeft"
	case NSLineMovesRight:
		return "NSLineMovesRight"
	case NSLineMovesUp:
		return "NSLineMovesUp"
	default:
		return fmt.Sprintf("NSLineMovementDirection(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSLineSweepDirection
type NSLineSweepDirection uint

const (
	NSLineSweepDown NSLineSweepDirection = 2
	NSLineSweepLeft NSLineSweepDirection = 0
	NSLineSweepRight NSLineSweepDirection = 1
	NSLineSweepUp NSLineSweepDirection = 3
)


func (e NSLineSweepDirection) String() string {
	switch e {
	case NSLineSweepDown:
		return "NSLineSweepDown"
	case NSLineSweepLeft:
		return "NSLineSweepLeft"
	case NSLineSweepRight:
		return "NSLineSweepRight"
	case NSLineSweepUp:
		return "NSLineSweepUp"
	default:
		return fmt.Sprintf("NSLineSweepDirection(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSMatrix/Mode-swift.enum
type NSMatrixMode int

const (
	// NSHighlightModeMatrix: An  is highlighted before it’s asked to track the mouse, then unhighlighted when it’s done tracking.
	NSHighlightModeMatrix NSMatrixMode = 1
	// NSListModeMatrix: objects are highlighted, but don’t track the mouse.
	NSListModeMatrix NSMatrixMode = 2
	// NSRadioModeMatrix: Selects no more than one  at a time.
	NSRadioModeMatrix NSMatrixMode = 0
	// NSTrackModeMatrix: The  objects are asked to track the mouse with  whenever the cursor is inside their bounds.
	NSTrackModeMatrix NSMatrixMode = 3
)


func (e NSMatrixMode) String() string {
	switch e {
	case NSHighlightModeMatrix:
		return "NSHighlightModeMatrix"
	case NSListModeMatrix:
		return "NSListModeMatrix"
	case NSRadioModeMatrix:
		return "NSRadioModeMatrix"
	case NSTrackModeMatrix:
		return "NSTrackModeMatrix"
	default:
		return fmt.Sprintf("NSMatrixMode(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSMediaLibraryBrowserController/Library
type NSMediaLibrary int

const (
	// NSMediaLibraryAudio: Display audio media.
	NSMediaLibraryAudio NSMediaLibrary = 1
	// NSMediaLibraryImage: Display image media.
	NSMediaLibraryImage NSMediaLibrary = 2
	// NSMediaLibraryMovie: Display movie media.
	NSMediaLibraryMovie NSMediaLibrary = 4
)


func (e NSMediaLibrary) String() string {
	switch e {
	case NSMediaLibraryAudio:
		return "NSMediaLibraryAudio"
	case NSMediaLibraryImage:
		return "NSMediaLibraryImage"
	case NSMediaLibraryMovie:
		return "NSMediaLibraryMovie"
	default:
		return fmt.Sprintf("NSMediaLibrary(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSMenuItemBadge/BadgeType
type NSMenuItemBadgeType int

const (
	// NSMenuItemBadgeTypeAlerts: A badge representing the number of alerts.
	NSMenuItemBadgeTypeAlerts NSMenuItemBadgeType = 3
	// NSMenuItemBadgeTypeNewItems: A badge representing the number of new items.
	NSMenuItemBadgeTypeNewItems NSMenuItemBadgeType = 2
	// NSMenuItemBadgeTypeNone: A badge with no string portion.
	NSMenuItemBadgeTypeNone NSMenuItemBadgeType = 0
	// NSMenuItemBadgeTypeUpdates: A badge representing the number of available updates.
	NSMenuItemBadgeTypeUpdates NSMenuItemBadgeType = 1
)


func (e NSMenuItemBadgeType) String() string {
	switch e {
	case NSMenuItemBadgeTypeAlerts:
		return "NSMenuItemBadgeTypeAlerts"
	case NSMenuItemBadgeTypeNewItems:
		return "NSMenuItemBadgeTypeNewItems"
	case NSMenuItemBadgeTypeNone:
		return "NSMenuItemBadgeTypeNone"
	case NSMenuItemBadgeTypeUpdates:
		return "NSMenuItemBadgeTypeUpdates"
	default:
		return fmt.Sprintf("NSMenuItemBadgeType(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSMenu/PresentationStyle-swift.enum
type NSMenuPresentationStyle int

const (
	// NSMenuPresentationStylePalette: A menu presentation style where items to display align horizontally.
	NSMenuPresentationStylePalette NSMenuPresentationStyle = 1
	// NSMenuPresentationStyleRegular: The default presentation style for a menu.
	NSMenuPresentationStyleRegular NSMenuPresentationStyle = 0
)


func (e NSMenuPresentationStyle) String() string {
	switch e {
	case NSMenuPresentationStylePalette:
		return "NSMenuPresentationStylePalette"
	case NSMenuPresentationStyleRegular:
		return "NSMenuPresentationStyleRegular"
	default:
		return fmt.Sprintf("NSMenuPresentationStyle(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSMenu/Properties
type NSMenuProperties int

const (
	// NSMenuPropertyItemAccessibilityDescription: The menu item’s accessibility description.
	NSMenuPropertyItemAccessibilityDescription NSMenuProperties = 32
	// NSMenuPropertyItemAttributedTitle: The menu item’s attributed string title.
	NSMenuPropertyItemAttributedTitle NSMenuProperties = 2
	// NSMenuPropertyItemEnabled: Whether the menu item is enabled or disabled.
	NSMenuPropertyItemEnabled NSMenuProperties = 16
	// NSMenuPropertyItemImage: The menu image.
	NSMenuPropertyItemImage NSMenuProperties = 8
	// NSMenuPropertyItemKeyEquivalent: The menu item’s key equivalent.
	NSMenuPropertyItemKeyEquivalent NSMenuProperties = 4
	// NSMenuPropertyItemTitle: The menu item’s title.
	NSMenuPropertyItemTitle NSMenuProperties = 1
)


func (e NSMenuProperties) String() string {
	switch e {
	case NSMenuPropertyItemAccessibilityDescription:
		return "NSMenuPropertyItemAccessibilityDescription"
	case NSMenuPropertyItemAttributedTitle:
		return "NSMenuPropertyItemAttributedTitle"
	case NSMenuPropertyItemEnabled:
		return "NSMenuPropertyItemEnabled"
	case NSMenuPropertyItemImage:
		return "NSMenuPropertyItemImage"
	case NSMenuPropertyItemKeyEquivalent:
		return "NSMenuPropertyItemKeyEquivalent"
	case NSMenuPropertyItemTitle:
		return "NSMenuPropertyItemTitle"
	default:
		return fmt.Sprintf("NSMenuProperties(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSMenu/SelectionMode-swift.enum
type NSMenuSelectionMode int

const (
	// NSMenuSelectionModeAutomatic: A selection mode where the menu determines the appropriate selection mode based on the context and its constants.
	NSMenuSelectionModeAutomatic NSMenuSelectionMode = 0
	// NSMenuSelectionModeSelectAny: A selection mode where someone can select multiple items in the menu.
	NSMenuSelectionModeSelectAny NSMenuSelectionMode = 2
	// NSMenuSelectionModeSelectOne: A selection mode where someone can select at most one menu item in the same selection group at the same time.
	NSMenuSelectionModeSelectOne NSMenuSelectionMode = 1
)


func (e NSMenuSelectionMode) String() string {
	switch e {
	case NSMenuSelectionModeAutomatic:
		return "NSMenuSelectionModeAutomatic"
	case NSMenuSelectionModeSelectAny:
		return "NSMenuSelectionModeSelectAny"
	case NSMenuSelectionModeSelectOne:
		return "NSMenuSelectionModeSelectOne"
	default:
		return fmt.Sprintf("NSMenuSelectionMode(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSMultibyteGlyphPacking
type NSMultibyteGlyphPacking int

const (
	// NSNativeShortGlyphPacking: The native format for macOS.
	NSNativeShortGlyphPacking NSMultibyteGlyphPacking = 5
)


func (e NSMultibyteGlyphPacking) String() string {
	switch e {
	case NSNativeShortGlyphPacking:
		return "NSNativeShortGlyphPacking"
	default:
		return fmt.Sprintf("NSMultibyteGlyphPacking(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSOpenGLContext/Parameter
type NSOpenGLContextParameter int

const (
	// NSOpenGLContextParameterCurrentRendererID: Get the current renderer ID.
	NSOpenGLContextParameterCurrentRendererID NSOpenGLContextParameter = 5
	// NSOpenGLContextParameterGPUFragmentProcessing: Get whether the CPU is currently processing fragments with the GPU.
	NSOpenGLContextParameterGPUFragmentProcessing NSOpenGLContextParameter = 7
	// NSOpenGLContextParameterGPUVertexProcessing: Get whether the CPU is currently processing vertices with the GPU.
	NSOpenGLContextParameterGPUVertexProcessing NSOpenGLContextParameter = 6
	// NSOpenGLContextParameterHasDrawable: Returns a Boolean that indicates whether a drawable is attached to the context.
	NSOpenGLContextParameterHasDrawable NSOpenGLContextParameter = 8
	// NSOpenGLContextParameterMPSwapsInFlight: The number of frames that the multithreaded OpenGL engine can process before stalling.
	NSOpenGLContextParameterMPSwapsInFlight NSOpenGLContextParameter = 9
	// NSOpenGLContextParameterRasterizationEnable: If disabled, all rasterization of 2D and 3D primitives is disabled.
	NSOpenGLContextParameterRasterizationEnable NSOpenGLContextParameter = 221
	// NSOpenGLContextParameterReclaimResources: Enable or disable reclaiming resources.
	NSOpenGLContextParameterReclaimResources NSOpenGLContextParameter = 4
	// NSOpenGLContextParameterStateValidation: If enabled, OpenGL inspects the context state each time the update method is called to ensure that it is in an appropriate state for switching between renderers.
	NSOpenGLContextParameterStateValidation NSOpenGLContextParameter = 301
	// NSOpenGLContextParameterSurfaceBackingSize: Set or get the height and width of the back buffer.
	NSOpenGLContextParameterSurfaceBackingSize NSOpenGLContextParameter = 3
	// NSOpenGLContextParameterSurfaceOpacity: Set or get the surface opacity.
	NSOpenGLContextParameterSurfaceOpacity NSOpenGLContextParameter = 2
	// NSOpenGLContextParameterSurfaceOrder: Set or get the surface order.
	NSOpenGLContextParameterSurfaceOrder NSOpenGLContextParameter = 1
	// NSOpenGLContextParameterSwapInterval: Set or get the swap interval.
	NSOpenGLContextParameterSwapInterval NSOpenGLContextParameter = 0
	// NSOpenGLContextParameterSwapRectangle: Sets or gets the swap rectangle.
	NSOpenGLContextParameterSwapRectangle NSOpenGLContextParameter = 200
	// NSOpenGLContextParameterSwapRectangleEnable: Enables or disables the swap rectangle in the context’s drawable object.
	NSOpenGLContextParameterSwapRectangleEnable NSOpenGLContextParameter = 201
)


func (e NSOpenGLContextParameter) String() string {
	switch e {
	case NSOpenGLContextParameterCurrentRendererID:
		return "NSOpenGLContextParameterCurrentRendererID"
	case NSOpenGLContextParameterGPUFragmentProcessing:
		return "NSOpenGLContextParameterGPUFragmentProcessing"
	case NSOpenGLContextParameterGPUVertexProcessing:
		return "NSOpenGLContextParameterGPUVertexProcessing"
	case NSOpenGLContextParameterHasDrawable:
		return "NSOpenGLContextParameterHasDrawable"
	case NSOpenGLContextParameterMPSwapsInFlight:
		return "NSOpenGLContextParameterMPSwapsInFlight"
	case NSOpenGLContextParameterRasterizationEnable:
		return "NSOpenGLContextParameterRasterizationEnable"
	case NSOpenGLContextParameterReclaimResources:
		return "NSOpenGLContextParameterReclaimResources"
	case NSOpenGLContextParameterStateValidation:
		return "NSOpenGLContextParameterStateValidation"
	case NSOpenGLContextParameterSurfaceBackingSize:
		return "NSOpenGLContextParameterSurfaceBackingSize"
	case NSOpenGLContextParameterSurfaceOpacity:
		return "NSOpenGLContextParameterSurfaceOpacity"
	case NSOpenGLContextParameterSurfaceOrder:
		return "NSOpenGLContextParameterSurfaceOrder"
	case NSOpenGLContextParameterSwapInterval:
		return "NSOpenGLContextParameterSwapInterval"
	case NSOpenGLContextParameterSwapRectangle:
		return "NSOpenGLContextParameterSwapRectangle"
	case NSOpenGLContextParameterSwapRectangleEnable:
		return "NSOpenGLContextParameterSwapRectangleEnable"
	default:
		return fmt.Sprintf("NSOpenGLContextParameter(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSOpenGLGlobalOption
type NSOpenGLGlobalOption int

const (
	// NSOpenGLGOClearFormatCache: Resets the pixel format cache if true.
	NSOpenGLGOClearFormatCache NSOpenGLGlobalOption = 1
	// NSOpenGLGOFormatCacheSize: Sets the size of the pixel format cache.
	NSOpenGLGOFormatCacheSize NSOpenGLGlobalOption = 0
	// NSOpenGLGORetainRenderers: Whether to retain loaded renderers in memory.
	NSOpenGLGORetainRenderers NSOpenGLGlobalOption = 2
	// NSOpenGLGOUseBuildCache: Whether to enable the function compilation block cache.
	NSOpenGLGOUseBuildCache NSOpenGLGlobalOption = 3
)


func (e NSOpenGLGlobalOption) String() string {
	switch e {
	case NSOpenGLGOClearFormatCache:
		return "NSOpenGLGOClearFormatCache"
	case NSOpenGLGOFormatCacheSize:
		return "NSOpenGLGOFormatCacheSize"
	case NSOpenGLGORetainRenderers:
		return "NSOpenGLGORetainRenderers"
	case NSOpenGLGOUseBuildCache:
		return "NSOpenGLGOUseBuildCache"
	default:
		return fmt.Sprintf("NSOpenGLGlobalOption(%d)", e)
	}
}




type NSOpenGLPFA uint

const (
	// Deprecated.
	NSOpenGLPFAAccelerated NSOpenGLPFA = 20
	// Deprecated.
	NSOpenGLPFAAcceleratedCompute NSOpenGLPFA = 25
	// Deprecated.
	NSOpenGLPFAAccumSize NSOpenGLPFA = 8
	// Deprecated.
	NSOpenGLPFAAllRenderers NSOpenGLPFA = 0
	// Deprecated.
	NSOpenGLPFAAllowOfflineRenderers NSOpenGLPFA = 24
	// Deprecated.
	NSOpenGLPFAAlphaSize NSOpenGLPFA = 5
	// Deprecated.
	NSOpenGLPFAAuxBuffers NSOpenGLPFA = 3
	// Deprecated.
	NSOpenGLPFAAuxDepthStencil NSOpenGLPFA = 13
	// Deprecated.
	NSOpenGLPFABackingStore NSOpenGLPFA = 22
	// Deprecated.
	NSOpenGLPFAClosestPolicy NSOpenGLPFA = 21
	// Deprecated.
	NSOpenGLPFAColorFloat NSOpenGLPFA = 14
	// Deprecated.
	NSOpenGLPFAColorSize NSOpenGLPFA = 4
	// Deprecated.
	NSOpenGLPFACompliant NSOpenGLPFA = 83
	// Deprecated.
	NSOpenGLPFADepthSize NSOpenGLPFA = 6
	// Deprecated.
	NSOpenGLPFADoubleBuffer NSOpenGLPFA = 2
	// Deprecated.
	NSOpenGLPFAFullScreen NSOpenGLPFA = 54
	// Deprecated.
	NSOpenGLPFAMPSafe NSOpenGLPFA = 78
	// Deprecated.
	NSOpenGLPFAMaximumPolicy NSOpenGLPFA = 10
	// Deprecated.
	NSOpenGLPFAMinimumPolicy NSOpenGLPFA = 9
	// Deprecated.
	NSOpenGLPFAMultiScreen NSOpenGLPFA = 81
	// Deprecated.
	NSOpenGLPFAMultisample NSOpenGLPFA = 15
	// Deprecated.
	NSOpenGLPFANoRecovery NSOpenGLPFA = 19
	// Deprecated.
	NSOpenGLPFAOffScreen NSOpenGLPFA = 53
	// Deprecated.
	NSOpenGLPFAOpenGLProfile NSOpenGLPFA = 26
	// Deprecated.
	NSOpenGLPFAPixelBuffer NSOpenGLPFA = 90
	// Deprecated.
	NSOpenGLPFARemotePixelBuffer NSOpenGLPFA = 91
	// Deprecated.
	NSOpenGLPFARendererID NSOpenGLPFA = 18
	// Deprecated.
	NSOpenGLPFARobust NSOpenGLPFA = 75
	// Deprecated.
	NSOpenGLPFASampleAlpha NSOpenGLPFA = 17
	// Deprecated.
	NSOpenGLPFASampleBuffers NSOpenGLPFA = 11
	// Deprecated.
	NSOpenGLPFASamples NSOpenGLPFA = 12
	// Deprecated.
	NSOpenGLPFAScreenMask NSOpenGLPFA = 23
	// Deprecated.
	NSOpenGLPFASingleRenderer NSOpenGLPFA = 71
	// Deprecated.
	NSOpenGLPFAStencilSize NSOpenGLPFA = 7
	// Deprecated.
	NSOpenGLPFAStereo NSOpenGLPFA = 6
	// Deprecated.
	NSOpenGLPFASupersample NSOpenGLPFA = 16
	// Deprecated.
	NSOpenGLPFATripleBuffer NSOpenGLPFA = 1
	// Deprecated.
	NSOpenGLPFAVirtualScreenCount NSOpenGLPFA = 27
	// Deprecated.
	NSOpenGLPFAWindow NSOpenGLPFA = 80
)


func (e NSOpenGLPFA) String() string {
	switch e {
	case NSOpenGLPFAAccelerated:
		return "NSOpenGLPFAAccelerated"
	case NSOpenGLPFAAcceleratedCompute:
		return "NSOpenGLPFAAcceleratedCompute"
	case NSOpenGLPFAAccumSize:
		return "NSOpenGLPFAAccumSize"
	case NSOpenGLPFAAllRenderers:
		return "NSOpenGLPFAAllRenderers"
	case NSOpenGLPFAAllowOfflineRenderers:
		return "NSOpenGLPFAAllowOfflineRenderers"
	case NSOpenGLPFAAlphaSize:
		return "NSOpenGLPFAAlphaSize"
	case NSOpenGLPFAAuxBuffers:
		return "NSOpenGLPFAAuxBuffers"
	case NSOpenGLPFAAuxDepthStencil:
		return "NSOpenGLPFAAuxDepthStencil"
	case NSOpenGLPFABackingStore:
		return "NSOpenGLPFABackingStore"
	case NSOpenGLPFAClosestPolicy:
		return "NSOpenGLPFAClosestPolicy"
	case NSOpenGLPFAColorFloat:
		return "NSOpenGLPFAColorFloat"
	case NSOpenGLPFAColorSize:
		return "NSOpenGLPFAColorSize"
	case NSOpenGLPFACompliant:
		return "NSOpenGLPFACompliant"
	case NSOpenGLPFADepthSize:
		return "NSOpenGLPFADepthSize"
	case NSOpenGLPFADoubleBuffer:
		return "NSOpenGLPFADoubleBuffer"
	case NSOpenGLPFAFullScreen:
		return "NSOpenGLPFAFullScreen"
	case NSOpenGLPFAMPSafe:
		return "NSOpenGLPFAMPSafe"
	case NSOpenGLPFAMaximumPolicy:
		return "NSOpenGLPFAMaximumPolicy"
	case NSOpenGLPFAMinimumPolicy:
		return "NSOpenGLPFAMinimumPolicy"
	case NSOpenGLPFAMultiScreen:
		return "NSOpenGLPFAMultiScreen"
	case NSOpenGLPFAMultisample:
		return "NSOpenGLPFAMultisample"
	case NSOpenGLPFANoRecovery:
		return "NSOpenGLPFANoRecovery"
	case NSOpenGLPFAOffScreen:
		return "NSOpenGLPFAOffScreen"
	case NSOpenGLPFAOpenGLProfile:
		return "NSOpenGLPFAOpenGLProfile"
	case NSOpenGLPFAPixelBuffer:
		return "NSOpenGLPFAPixelBuffer"
	case NSOpenGLPFARemotePixelBuffer:
		return "NSOpenGLPFARemotePixelBuffer"
	case NSOpenGLPFARendererID:
		return "NSOpenGLPFARendererID"
	case NSOpenGLPFARobust:
		return "NSOpenGLPFARobust"
	case NSOpenGLPFASampleAlpha:
		return "NSOpenGLPFASampleAlpha"
	case NSOpenGLPFASampleBuffers:
		return "NSOpenGLPFASampleBuffers"
	case NSOpenGLPFASamples:
		return "NSOpenGLPFASamples"
	case NSOpenGLPFAScreenMask:
		return "NSOpenGLPFAScreenMask"
	case NSOpenGLPFASingleRenderer:
		return "NSOpenGLPFASingleRenderer"
	case NSOpenGLPFAStencilSize:
		return "NSOpenGLPFAStencilSize"
	case NSOpenGLPFASupersample:
		return "NSOpenGLPFASupersample"
	case NSOpenGLPFATripleBuffer:
		return "NSOpenGLPFATripleBuffer"
	case NSOpenGLPFAVirtualScreenCount:
		return "NSOpenGLPFAVirtualScreenCount"
	case NSOpenGLPFAWindow:
		return "NSOpenGLPFAWindow"
	default:
		return fmt.Sprintf("NSOpenGLPFA(%d)", e)
	}
}




type NSOpenGLProfile uint

const (
	// Deprecated.
	NSOpenGLProfileVersion3_2Core NSOpenGLProfile = 1
	// Deprecated.
	NSOpenGLProfileVersion4_1Core NSOpenGLProfile = 2
	// Deprecated.
	NSOpenGLProfileVersionLegacy NSOpenGLProfile = 0
)


func (e NSOpenGLProfile) String() string {
	switch e {
	case NSOpenGLProfileVersion3_2Core:
		return "NSOpenGLProfileVersion3_2Core"
	case NSOpenGLProfileVersion4_1Core:
		return "NSOpenGLProfileVersion4_1Core"
	case NSOpenGLProfileVersionLegacy:
		return "NSOpenGLProfileVersionLegacy"
	default:
		return fmt.Sprintf("NSOpenGLProfile(%d)", e)
	}
}




type NSOutlineViewDropOnItemIndexConstants int

const (
	// NSOutlineViewDropOnItemIndex: May be used as a valid child index of a drop target item.
	NSOutlineViewDropOnItemIndex NSOutlineViewDropOnItemIndexConstants = -1
)


func (e NSOutlineViewDropOnItemIndexConstants) String() string {
	switch e {
	case NSOutlineViewDropOnItemIndex:
		return "NSOutlineViewDropOnItemIndex"
	default:
		return fmt.Sprintf("NSOutlineViewDropOnItemIndexConstants(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSPDFPanel/Options-swift.struct
type NSPDFPanelOptions int

const (
	// NSPDFPanelRequestsParentDirectory: The PDF panel doesn’t show a name field; instead, it allows the user to identify a directory in which to save multiple PDF files.
	NSPDFPanelRequestsParentDirectory NSPDFPanelOptions = 16777216
	// NSPDFPanelShowsOrientation: The PDF panel shows the current orientation of the PDF contents, such as landscape or portrait.
	NSPDFPanelShowsOrientation NSPDFPanelOptions = 8
	// NSPDFPanelShowsPaperSize: The PDF panel shows a menu of paper sizes.
	NSPDFPanelShowsPaperSize NSPDFPanelOptions = 4
)


func (e NSPDFPanelOptions) String() string {
	switch e {
	case NSPDFPanelRequestsParentDirectory:
		return "NSPDFPanelRequestsParentDirectory"
	case NSPDFPanelShowsOrientation:
		return "NSPDFPanelShowsOrientation"
	case NSPDFPanelShowsPaperSize:
		return "NSPDFPanelShowsPaperSize"
	default:
		return fmt.Sprintf("NSPDFPanelOptions(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSPageController/TransitionStyle-swift.enum
type NSPageControllerTransitionStyle int

const (
	// NSPageControllerTransitionStyleHorizontalStrip: Each page is laid out next to each other in one long horizontal strip
	NSPageControllerTransitionStyleHorizontalStrip NSPageControllerTransitionStyle = 2
	// NSPageControllerTransitionStyleStackBook: Pages are stacked on top of each other.
	NSPageControllerTransitionStyleStackBook NSPageControllerTransitionStyle = 1
	// NSPageControllerTransitionStyleStackHistory: Pages are stacked on top of each other.
	NSPageControllerTransitionStyleStackHistory NSPageControllerTransitionStyle = 0
)


func (e NSPageControllerTransitionStyle) String() string {
	switch e {
	case NSPageControllerTransitionStyleHorizontalStrip:
		return "NSPageControllerTransitionStyleHorizontalStrip"
	case NSPageControllerTransitionStyleStackBook:
		return "NSPageControllerTransitionStyleStackBook"
	case NSPageControllerTransitionStyleStackHistory:
		return "NSPageControllerTransitionStyleStackHistory"
	default:
		return fmt.Sprintf("NSPageControllerTransitionStyle(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSPageLayout/Result
type NSPageLayoutResult int

const (
	NSPageLayoutResultCancelled NSPageLayoutResult = 0
	NSPageLayoutResultChanged NSPageLayoutResult = 1
)


func (e NSPageLayoutResult) String() string {
	switch e {
	case NSPageLayoutResultCancelled:
		return "NSPageLayoutResultCancelled"
	case NSPageLayoutResultChanged:
		return "NSPageLayoutResultChanged"
	default:
		return fmt.Sprintf("NSPageLayoutResult(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/PaperOrientation
type NSPaperOrientation int

const (
	// NSPaperOrientationLandscape: Pages are printed in landscape orientation.
	NSPaperOrientationLandscape NSPaperOrientation = 1
	// NSPaperOrientationPortrait: Pages are printed in portrait orientation.
	NSPaperOrientationPortrait NSPaperOrientation = 0
)


func (e NSPaperOrientation) String() string {
	switch e {
	case NSPaperOrientationLandscape:
		return "NSPaperOrientationLandscape"
	case NSPaperOrientationPortrait:
		return "NSPaperOrientationPortrait"
	default:
		return fmt.Sprintf("NSPaperOrientation(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSPasteboard/AccessBehavior-swift.enum
type NSPasteboardAccessBehavior int

const (
	// NSPasteboardAccessBehaviorAlwaysAllow: The system will automatically allow all pasteboard access, without notifying the user.
	NSPasteboardAccessBehaviorAlwaysAllow NSPasteboardAccessBehavior = 2
	// NSPasteboardAccessBehaviorAlwaysDeny: The system will automatically deny all pasteboard access, without notifying the user.
	NSPasteboardAccessBehaviorAlwaysDeny NSPasteboardAccessBehavior = 3
	// NSPasteboardAccessBehaviorAsk: The system will notify the user and ask for permission before granting pasteboard access.
	NSPasteboardAccessBehaviorAsk NSPasteboardAccessBehavior = 1
	// NSPasteboardAccessBehaviorDefault: The default behavior for the General pasteboard is to ask upon programmatic access.
	NSPasteboardAccessBehaviorDefault NSPasteboardAccessBehavior = 0
)


func (e NSPasteboardAccessBehavior) String() string {
	switch e {
	case NSPasteboardAccessBehaviorAlwaysAllow:
		return "NSPasteboardAccessBehaviorAlwaysAllow"
	case NSPasteboardAccessBehaviorAlwaysDeny:
		return "NSPasteboardAccessBehaviorAlwaysDeny"
	case NSPasteboardAccessBehaviorAsk:
		return "NSPasteboardAccessBehaviorAsk"
	case NSPasteboardAccessBehaviorDefault:
		return "NSPasteboardAccessBehaviorDefault"
	default:
		return fmt.Sprintf("NSPasteboardAccessBehavior(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSPasteboard/ContentsOptions
type NSPasteboardContentsOptions int

const (
	// NSPasteboardContentsCurrentHostOnly: The pasteboard contents are available only on the current device, and not on any other devices.
	NSPasteboardContentsCurrentHostOnly NSPasteboardContentsOptions = 1
)


func (e NSPasteboardContentsOptions) String() string {
	switch e {
	case NSPasteboardContentsCurrentHostOnly:
		return "NSPasteboardContentsCurrentHostOnly"
	default:
		return fmt.Sprintf("NSPasteboardContentsOptions(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSPasteboard/ReadingOptions
type NSPasteboardReadingOptions int

const (
	// NSPasteboardReadingAsData: An option to read data from the pasteboard as-is and return it as a data object.
	NSPasteboardReadingAsData NSPasteboardReadingOptions = 0
	// NSPasteboardReadingAsKeyedArchive: An option to read data from the pasteboard and use it to initialize the object.
	NSPasteboardReadingAsKeyedArchive NSPasteboardReadingOptions = 4
	// NSPasteboardReadingAsPropertyList: An option to read data from the pasteboard and unserialize it as a property list.
	NSPasteboardReadingAsPropertyList NSPasteboardReadingOptions = 2
	// NSPasteboardReadingAsString: An option to read data from the pasteboard and convert it to a string object.
	NSPasteboardReadingAsString NSPasteboardReadingOptions = 1
)


func (e NSPasteboardReadingOptions) String() string {
	switch e {
	case NSPasteboardReadingAsData:
		return "NSPasteboardReadingAsData"
	case NSPasteboardReadingAsKeyedArchive:
		return "NSPasteboardReadingAsKeyedArchive"
	case NSPasteboardReadingAsPropertyList:
		return "NSPasteboardReadingAsPropertyList"
	case NSPasteboardReadingAsString:
		return "NSPasteboardReadingAsString"
	default:
		return fmt.Sprintf("NSPasteboardReadingOptions(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSPasteboard/WritingOptions
type NSPasteboardWritingOptions int

const (
	// NSPasteboardWritingPromised: Data for a type with this option is promised, not immediately written.
	NSPasteboardWritingPromised NSPasteboardWritingOptions = 512
)


func (e NSPasteboardWritingOptions) String() string {
	switch e {
	case NSPasteboardWritingPromised:
		return "NSPasteboardWritingPromised"
	default:
		return fmt.Sprintf("NSPasteboardWritingOptions(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSPathControl/Style
type NSPathStyle int

const (
	NSPathStyleNavigationBar NSPathStyle = 1
	NSPathStylePopUp NSPathStyle = 2
	NSPathStyleStandard NSPathStyle = 0
)


func (e NSPathStyle) String() string {
	switch e {
	case NSPathStyleNavigationBar:
		return "NSPathStyleNavigationBar"
	case NSPathStylePopUp:
		return "NSPathStylePopUp"
	case NSPathStyleStandard:
		return "NSPathStyleStandard"
	default:
		return fmt.Sprintf("NSPathStyle(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/ControlRepresentation-swift.enum
type NSPickerTouchBarItemControlRepresentation int

const (
	// NSPickerTouchBarItemControlRepresentationAutomatic: The system dynamically changes the display mode based on the available space.
	NSPickerTouchBarItemControlRepresentationAutomatic NSPickerTouchBarItemControlRepresentation = 0
	// NSPickerTouchBarItemControlRepresentationCollapsed: The system displays the control’s options through a popover.
	NSPickerTouchBarItemControlRepresentationCollapsed NSPickerTouchBarItemControlRepresentation = 2
	// NSPickerTouchBarItemControlRepresentationExpanded: The system displays the control and all of its options directly.
	NSPickerTouchBarItemControlRepresentationExpanded NSPickerTouchBarItemControlRepresentation = 1
)


func (e NSPickerTouchBarItemControlRepresentation) String() string {
	switch e {
	case NSPickerTouchBarItemControlRepresentationAutomatic:
		return "NSPickerTouchBarItemControlRepresentationAutomatic"
	case NSPickerTouchBarItemControlRepresentationCollapsed:
		return "NSPickerTouchBarItemControlRepresentationCollapsed"
	case NSPickerTouchBarItemControlRepresentationExpanded:
		return "NSPickerTouchBarItemControlRepresentationExpanded"
	default:
		return fmt.Sprintf("NSPickerTouchBarItemControlRepresentation(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/SelectionMode-swift.enum
type NSPickerTouchBarItemSelectionMode int

const (
	// NSPickerTouchBarItemSelectionModeMomentary: A mode in which an option is only selected while a person is interacting within the bounds of that option.
	NSPickerTouchBarItemSelectionModeMomentary NSPickerTouchBarItemSelectionMode = 2
	// NSPickerTouchBarItemSelectionModeSelectAny: A mode in which a person can select one or more options in the control at a time.
	NSPickerTouchBarItemSelectionModeSelectAny NSPickerTouchBarItemSelectionMode = 1
	// NSPickerTouchBarItemSelectionModeSelectOne: A mode in which a person can only select one option in the control at a time.
	NSPickerTouchBarItemSelectionModeSelectOne NSPickerTouchBarItemSelectionMode = 0
)


func (e NSPickerTouchBarItemSelectionMode) String() string {
	switch e {
	case NSPickerTouchBarItemSelectionModeMomentary:
		return "NSPickerTouchBarItemSelectionModeMomentary"
	case NSPickerTouchBarItemSelectionModeSelectAny:
		return "NSPickerTouchBarItemSelectionModeSelectAny"
	case NSPickerTouchBarItemSelectionModeSelectOne:
		return "NSPickerTouchBarItemSelectionModeSelectOne"
	default:
		return fmt.Sprintf("NSPickerTouchBarItemSelectionMode(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSEvent/PointingDeviceType-swift.enum
type NSPointingDeviceType int

const (
	// NSPointingDeviceTypeCursor: Represents a cursor pointing device.
	NSPointingDeviceTypeCursor NSPointingDeviceType = 2
	// NSPointingDeviceTypeEraser: Represents the eraser end of a stylus-like pointing device.
	NSPointingDeviceTypeEraser NSPointingDeviceType = 3
	// NSPointingDeviceTypePen: Represents the tip end of a stylus-like pointing device.
	NSPointingDeviceTypePen NSPointingDeviceType = 1
	// NSPointingDeviceTypeUnknown: Represents an unknown type of pointing device.
	NSPointingDeviceTypeUnknown NSPointingDeviceType = 0
)


func (e NSPointingDeviceType) String() string {
	switch e {
	case NSPointingDeviceTypeCursor:
		return "NSPointingDeviceTypeCursor"
	case NSPointingDeviceTypeEraser:
		return "NSPointingDeviceTypeEraser"
	case NSPointingDeviceTypePen:
		return "NSPointingDeviceTypePen"
	case NSPointingDeviceTypeUnknown:
		return "NSPointingDeviceTypeUnknown"
	default:
		return fmt.Sprintf("NSPointingDeviceType(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSPopUpButton/ArrowPosition
type NSPopUpArrowPosition int

const (
	// NSPopUpArrowAtBottom: Arrow is drawn at the edge of the button, pointing toward the .
	NSPopUpArrowAtBottom NSPopUpArrowPosition = 2
	// NSPopUpArrowAtCenter: Arrow is centered vertically, pointing toward the .
	NSPopUpArrowAtCenter NSPopUpArrowPosition = 1
	// NSPopUpNoArrow: Does not display any arrow in the control.
	NSPopUpNoArrow NSPopUpArrowPosition = 0
)


func (e NSPopUpArrowPosition) String() string {
	switch e {
	case NSPopUpArrowAtBottom:
		return "NSPopUpArrowAtBottom"
	case NSPopUpArrowAtCenter:
		return "NSPopUpArrowAtCenter"
	case NSPopUpNoArrow:
		return "NSPopUpNoArrow"
	default:
		return fmt.Sprintf("NSPopUpArrowPosition(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSPopover/Appearance-swift.enum
type NSPopoverAppearance int

const (
	// NSPopoverAppearanceHUD: The popover draws with a HUD appearance.
	NSPopoverAppearanceHUD NSPopoverAppearance = 1
	// NSPopoverAppearanceMinimal: The popover draws with a minimal appearance.
	NSPopoverAppearanceMinimal NSPopoverAppearance = 0
)


func (e NSPopoverAppearance) String() string {
	switch e {
	case NSPopoverAppearanceHUD:
		return "NSPopoverAppearanceHUD"
	case NSPopoverAppearanceMinimal:
		return "NSPopoverAppearanceMinimal"
	default:
		return fmt.Sprintf("NSPopoverAppearance(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSPopover/Behavior-swift.enum
type NSPopoverBehavior int

const (
	// NSPopoverBehaviorApplicationDefined: Your application assumes responsibility for closing the popover.
	NSPopoverBehaviorApplicationDefined NSPopoverBehavior = 0
	// NSPopoverBehaviorSemitransient: The system will close the popover when the user interacts with user interface elements in the window containing the popover’s positioning view.
	NSPopoverBehaviorSemitransient NSPopoverBehavior = 2
	// NSPopoverBehaviorTransient: The system will close the popover when the user interacts with a user interface element outside the popover.
	NSPopoverBehaviorTransient NSPopoverBehavior = 1
)


func (e NSPopoverBehavior) String() string {
	switch e {
	case NSPopoverBehaviorApplicationDefined:
		return "NSPopoverBehaviorApplicationDefined"
	case NSPopoverBehaviorSemitransient:
		return "NSPopoverBehaviorSemitransient"
	case NSPopoverBehaviorTransient:
		return "NSPopoverBehaviorTransient"
	default:
		return fmt.Sprintf("NSPopoverBehavior(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSEvent/PressureBehavior-swift.enum
type NSPressureBehavior int

const (
	// NSPressureBehaviorPrimaryAccelerator: A pressure gesture’s behavior begins on left mouse-down events.
	NSPressureBehaviorPrimaryAccelerator NSPressureBehavior = 3
	// NSPressureBehaviorPrimaryClick: A pressure gesture’s behavior begins on left mouse-down events.
	NSPressureBehaviorPrimaryClick NSPressureBehavior = 1
	// NSPressureBehaviorPrimaryDeepClick: A pressure gesture’s behavior begins on left mouse-down events.
	NSPressureBehaviorPrimaryDeepClick NSPressureBehavior = 5
	// NSPressureBehaviorPrimaryDeepDrag: A pressure gesture’s behavior begins on left mouse-down events.
	NSPressureBehaviorPrimaryDeepDrag NSPressureBehavior = 6
	// NSPressureBehaviorPrimaryDefault: This is the default behavior when a pressure gesture’s behavior has not been explicitly configured.
	NSPressureBehaviorPrimaryDefault NSPressureBehavior = 0
	// NSPressureBehaviorPrimaryGeneric: A pressure gesture’s behavior begins on left mouse-down events.
	NSPressureBehaviorPrimaryGeneric NSPressureBehavior = 2
	// NSPressureBehaviorUnknown: A pressure gesture’s behavior is not known, perhaps because the input device does not support pressure gestures.
	NSPressureBehaviorUnknown NSPressureBehavior = -1
)


func (e NSPressureBehavior) String() string {
	switch e {
	case NSPressureBehaviorPrimaryAccelerator:
		return "NSPressureBehaviorPrimaryAccelerator"
	case NSPressureBehaviorPrimaryClick:
		return "NSPressureBehaviorPrimaryClick"
	case NSPressureBehaviorPrimaryDeepClick:
		return "NSPressureBehaviorPrimaryDeepClick"
	case NSPressureBehaviorPrimaryDeepDrag:
		return "NSPressureBehaviorPrimaryDeepDrag"
	case NSPressureBehaviorPrimaryDefault:
		return "NSPressureBehaviorPrimaryDefault"
	case NSPressureBehaviorPrimaryGeneric:
		return "NSPressureBehaviorPrimaryGeneric"
	case NSPressureBehaviorUnknown:
		return "NSPressureBehaviorUnknown"
	default:
		return fmt.Sprintf("NSPressureBehavior(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSPrintPanel/Options-swift.struct
type NSPrintPanelOptions int

const (
	// NSPrintPanelShowsCopies: The Print panel includes a field for manipulating the number of copies being printed.
	NSPrintPanelShowsCopies NSPrintPanelOptions = 1
	// NSPrintPanelShowsOrientation: The Print panel includes a control for manipulating the page orientation.
	NSPrintPanelShowsOrientation NSPrintPanelOptions = 8
	// NSPrintPanelShowsPageRange: The Print panel includes a set of fields for manipulating the range of pages being printed.
	NSPrintPanelShowsPageRange NSPrintPanelOptions = 2
	// NSPrintPanelShowsPageSetupAccessory: The Print panel includes a separate accessory view for manipulating the paper size, orientation, and scaling attributes.
	NSPrintPanelShowsPageSetupAccessory NSPrintPanelOptions = 256
	// NSPrintPanelShowsPaperSize: The Print panel includes a control for manipulating the paper size of the printer.
	NSPrintPanelShowsPaperSize NSPrintPanelOptions = 4
	// NSPrintPanelShowsPreview: The Print panel displays a built-in preview of the document contents.
	NSPrintPanelShowsPreview NSPrintPanelOptions = 131072
	// NSPrintPanelShowsPrintSelection: The Print panel includes an additional selection option for paper range.
	NSPrintPanelShowsPrintSelection NSPrintPanelOptions = 32
	// NSPrintPanelShowsScaling: The Print panel includes a control for scaling the printed output.
	NSPrintPanelShowsScaling NSPrintPanelOptions = 16
)


func (e NSPrintPanelOptions) String() string {
	switch e {
	case NSPrintPanelShowsCopies:
		return "NSPrintPanelShowsCopies"
	case NSPrintPanelShowsOrientation:
		return "NSPrintPanelShowsOrientation"
	case NSPrintPanelShowsPageRange:
		return "NSPrintPanelShowsPageRange"
	case NSPrintPanelShowsPageSetupAccessory:
		return "NSPrintPanelShowsPageSetupAccessory"
	case NSPrintPanelShowsPaperSize:
		return "NSPrintPanelShowsPaperSize"
	case NSPrintPanelShowsPreview:
		return "NSPrintPanelShowsPreview"
	case NSPrintPanelShowsPrintSelection:
		return "NSPrintPanelShowsPrintSelection"
	case NSPrintPanelShowsScaling:
		return "NSPrintPanelShowsScaling"
	default:
		return fmt.Sprintf("NSPrintPanelOptions(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSPrintPanel/Result
type NSPrintPanelResult int

const (
	NSPrintPanelResultCancelled NSPrintPanelResult = 0
	NSPrintPanelResultPrinted NSPrintPanelResult = 1
)


func (e NSPrintPanelResult) String() string {
	switch e {
	case NSPrintPanelResultCancelled:
		return "NSPrintPanelResultCancelled"
	case NSPrintPanelResultPrinted:
		return "NSPrintPanelResultPrinted"
	default:
		return fmt.Sprintf("NSPrintPanelResult(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSPrintOperation/RenderingQuality
type NSPrintRenderingQuality int

const (
	// NSPrintRenderingQualityBest: Renders the printing at the best possible quality, regardless of speed.
	NSPrintRenderingQualityBest NSPrintRenderingQuality = 0
	// NSPrintRenderingQualityResponsive: Sacrifices the least possible amount of rendering quality for speed to maintain a responsive user interface.
	NSPrintRenderingQualityResponsive NSPrintRenderingQuality = 1
)


func (e NSPrintRenderingQuality) String() string {
	switch e {
	case NSPrintRenderingQualityBest:
		return "NSPrintRenderingQualityBest"
	case NSPrintRenderingQualityResponsive:
		return "NSPrintRenderingQualityResponsive"
	default:
		return fmt.Sprintf("NSPrintRenderingQuality(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSPrinter/TableStatus
type NSPrinterTableStatus int

const (
	// NSPrinterTableError: Printer table is not valid.
	NSPrinterTableError NSPrinterTableStatus = 2
	// NSPrinterTableNotFound: Printer table was not found.
	NSPrinterTableNotFound NSPrinterTableStatus = 1
	// NSPrinterTableOK: Printer table was found and is valid.
	NSPrinterTableOK NSPrinterTableStatus = 0
)


func (e NSPrinterTableStatus) String() string {
	switch e {
	case NSPrinterTableError:
		return "NSPrinterTableError"
	case NSPrinterTableNotFound:
		return "NSPrinterTableNotFound"
	case NSPrinterTableOK:
		return "NSPrinterTableOK"
	default:
		return fmt.Sprintf("NSPrinterTableStatus(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/Orientation-swift.enum
type NSPrintingOrientation int

const (
	// NSLandscapeOrientation: Orientation is landscape (page is wider than it is tall).
	NSLandscapeOrientation NSPrintingOrientation = 1
	// NSPortraitOrientation: Orientation is portrait (page is taller than it is wide).
	NSPortraitOrientation NSPrintingOrientation = 0
)


func (e NSPrintingOrientation) String() string {
	switch e {
	case NSLandscapeOrientation:
		return "NSLandscapeOrientation"
	case NSPortraitOrientation:
		return "NSPortraitOrientation"
	default:
		return fmt.Sprintf("NSPrintingOrientation(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSPrintOperation/PageOrder-swift.enum
type NSPrintingPageOrder int

const (
	// NSAscendingPageOrder: Ascending (back to front) page order.
	NSAscendingPageOrder NSPrintingPageOrder = 1
	// NSDescendingPageOrder: Descending (front to back) page order.
	NSDescendingPageOrder NSPrintingPageOrder = -1
	// NSSpecialPageOrder: The spooler does not rearrange pages—they are printed in the order received by the spooler.
	NSSpecialPageOrder NSPrintingPageOrder = 0
	// NSUnknownPageOrder: No page order specified.
	NSUnknownPageOrder NSPrintingPageOrder = 2
)


func (e NSPrintingPageOrder) String() string {
	switch e {
	case NSAscendingPageOrder:
		return "NSAscendingPageOrder"
	case NSDescendingPageOrder:
		return "NSDescendingPageOrder"
	case NSSpecialPageOrder:
		return "NSSpecialPageOrder"
	case NSUnknownPageOrder:
		return "NSUnknownPageOrder"
	default:
		return fmt.Sprintf("NSPrintingPageOrder(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/PaginationMode
type NSPrintingPaginationMode int

const (
	NSPrintingPaginationModeAutomatic NSPrintingPaginationMode = 0
	NSPrintingPaginationModeClip NSPrintingPaginationMode = 2
	NSPrintingPaginationModeFit NSPrintingPaginationMode = 1
)


func (e NSPrintingPaginationMode) String() string {
	switch e {
	case NSPrintingPaginationModeAutomatic:
		return "NSPrintingPaginationModeAutomatic"
	case NSPrintingPaginationModeClip:
		return "NSPrintingPaginationModeClip"
	case NSPrintingPaginationModeFit:
		return "NSPrintingPaginationModeFit"
	default:
		return fmt.Sprintf("NSPrintingPaginationMode(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/Style-swift.enum
type NSProgressIndicatorStyle int

const (
	// NSProgressIndicatorStyleBar: A rectangular indicator that can be determinate or indeterminate.
	NSProgressIndicatorStyleBar NSProgressIndicatorStyle = 0
	// NSProgressIndicatorStyleSpinning: A small circular indicator that can be determinate or indeterminate.
	NSProgressIndicatorStyleSpinning NSProgressIndicatorStyle = 1
)


func (e NSProgressIndicatorStyle) String() string {
	switch e {
	case NSProgressIndicatorStyleBar:
		return "NSProgressIndicatorStyleBar"
	case NSProgressIndicatorStyleSpinning:
		return "NSProgressIndicatorStyleSpinning"
	default:
		return fmt.Sprintf("NSProgressIndicatorStyle(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSProgressIndicatorThickness
type NSProgressIndicatorThickness uint

const (
	NSProgressIndicatorPreferredAquaThickness NSProgressIndicatorThickness = 12
	NSProgressIndicatorPreferredLargeThickness NSProgressIndicatorThickness = 18
	NSProgressIndicatorPreferredSmallThickness NSProgressIndicatorThickness = 10
	NSProgressIndicatorPreferredThickness NSProgressIndicatorThickness = 14
)


func (e NSProgressIndicatorThickness) String() string {
	switch e {
	case NSProgressIndicatorPreferredAquaThickness:
		return "NSProgressIndicatorPreferredAquaThickness"
	case NSProgressIndicatorPreferredLargeThickness:
		return "NSProgressIndicatorPreferredLargeThickness"
	case NSProgressIndicatorPreferredSmallThickness:
		return "NSProgressIndicatorPreferredSmallThickness"
	case NSProgressIndicatorPreferredThickness:
		return "NSProgressIndicatorPreferredThickness"
	default:
		return fmt.Sprintf("NSProgressIndicatorThickness(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSRectAlignment
type NSRectAlignment int

const (
	// NSRectAlignmentBottom: Aligns to the bottom edge.
	NSRectAlignmentBottom NSRectAlignment = 5
	// NSRectAlignmentBottomLeading: Aligns to the bottom and leading edges.
	NSRectAlignmentBottomLeading NSRectAlignment = 4
	// NSRectAlignmentBottomTrailing: Aligns to the bottom and trailing edges.
	NSRectAlignmentBottomTrailing NSRectAlignment = 6
	// NSRectAlignmentLeading: Aligns to the leading edge.
	NSRectAlignmentLeading NSRectAlignment = 3
	// NSRectAlignmentNone: Has no specified alignment.
	NSRectAlignmentNone NSRectAlignment = 0
	// NSRectAlignmentTop: Aligns to the top edge.
	NSRectAlignmentTop NSRectAlignment = 1
	// NSRectAlignmentTopLeading: Aligns to the top and leading edges.
	NSRectAlignmentTopLeading NSRectAlignment = 2
	// NSRectAlignmentTopTrailing: Aligns to the top and trailing edges.
	NSRectAlignmentTopTrailing NSRectAlignment = 8
	// NSRectAlignmentTrailing: Aligns to the trailing edge.
	NSRectAlignmentTrailing NSRectAlignment = 7
)


func (e NSRectAlignment) String() string {
	switch e {
	case NSRectAlignmentBottom:
		return "NSRectAlignmentBottom"
	case NSRectAlignmentBottomLeading:
		return "NSRectAlignmentBottomLeading"
	case NSRectAlignmentBottomTrailing:
		return "NSRectAlignmentBottomTrailing"
	case NSRectAlignmentLeading:
		return "NSRectAlignmentLeading"
	case NSRectAlignmentNone:
		return "NSRectAlignmentNone"
	case NSRectAlignmentTop:
		return "NSRectAlignmentTop"
	case NSRectAlignmentTopLeading:
		return "NSRectAlignmentTopLeading"
	case NSRectAlignmentTopTrailing:
		return "NSRectAlignmentTopTrailing"
	case NSRectAlignmentTrailing:
		return "NSRectAlignmentTrailing"
	default:
		return fmt.Sprintf("NSRectAlignment(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSApplication/RemoteNotificationType
type NSRemoteNotificationType int

const (
	// NSRemoteNotificationTypeAlert: The app should display an alert.
	NSRemoteNotificationTypeAlert NSRemoteNotificationType = 4
	// NSRemoteNotificationTypeBadge: The app should display a badge.
	NSRemoteNotificationTypeBadge NSRemoteNotificationType = 1
	// NSRemoteNotificationTypeSound: The app should play a sound.
	NSRemoteNotificationTypeSound NSRemoteNotificationType = 2
)


func (e NSRemoteNotificationType) String() string {
	switch e {
	case NSRemoteNotificationTypeAlert:
		return "NSRemoteNotificationTypeAlert"
	case NSRemoteNotificationTypeBadge:
		return "NSRemoteNotificationTypeBadge"
	case NSRemoteNotificationTypeSound:
		return "NSRemoteNotificationTypeSound"
	default:
		return fmt.Sprintf("NSRemoteNotificationType(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSApplication/RequestUserAttentionType
type NSRequestUserAttentionType int

const (
	// NSCriticalRequest: The user attention request is a critical request.
	NSCriticalRequest NSRequestUserAttentionType = 0
	// NSInformationalRequest: The user attention request is an informational request.
	NSInformationalRequest NSRequestUserAttentionType = 10
)


func (e NSRequestUserAttentionType) String() string {
	switch e {
	case NSCriticalRequest:
		return "NSCriticalRequest"
	case NSInformationalRequest:
		return "NSInformationalRequest"
	default:
		return fmt.Sprintf("NSRequestUserAttentionType(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSRuleEditor/NestingMode-swift.enum
type NSRuleEditorNestingMode int

const (
	// NSRuleEditorNestingModeCompound: Unlimited nesting and compound rows.
	NSRuleEditorNestingModeCompound NSRuleEditorNestingMode = 2
	// NSRuleEditorNestingModeList: Allows a single list, with no nesting and no compound rows.
	NSRuleEditorNestingModeList NSRuleEditorNestingMode = 1
	// NSRuleEditorNestingModeSimple: One compound row at the top with subrows beneath it, and no further nesting allowed.
	NSRuleEditorNestingModeSimple NSRuleEditorNestingMode = 3
	// NSRuleEditorNestingModeSingle: Only a single row is allowed.
	NSRuleEditorNestingModeSingle NSRuleEditorNestingMode = 0
)


func (e NSRuleEditorNestingMode) String() string {
	switch e {
	case NSRuleEditorNestingModeCompound:
		return "NSRuleEditorNestingModeCompound"
	case NSRuleEditorNestingModeList:
		return "NSRuleEditorNestingModeList"
	case NSRuleEditorNestingModeSimple:
		return "NSRuleEditorNestingModeSimple"
	case NSRuleEditorNestingModeSingle:
		return "NSRuleEditorNestingModeSingle"
	default:
		return fmt.Sprintf("NSRuleEditorNestingMode(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSRuleEditor/RowType
type NSRuleEditorRowType int

const (
	// NSRuleEditorRowTypeCompound: Specifies a compound row.
	NSRuleEditorRowTypeCompound NSRuleEditorRowType = 1
	// NSRuleEditorRowTypeSimple: Specifies a simple row.
	NSRuleEditorRowTypeSimple NSRuleEditorRowType = 0
)


func (e NSRuleEditorRowType) String() string {
	switch e {
	case NSRuleEditorRowTypeCompound:
		return "NSRuleEditorRowTypeCompound"
	case NSRuleEditorRowTypeSimple:
		return "NSRuleEditorRowTypeSimple"
	default:
		return fmt.Sprintf("NSRuleEditorRowType(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSRulerView/Orientation-swift.enum
type NSRulerOrientation int

const (
	// NSHorizontalRuler: Ruler is oriented horizontally.
	NSHorizontalRuler NSRulerOrientation = 0
	// NSVerticalRuler: Ruler is oriented vertically.
	NSVerticalRuler NSRulerOrientation = 1
)


func (e NSRulerOrientation) String() string {
	switch e {
	case NSHorizontalRuler:
		return "NSHorizontalRuler"
	case NSVerticalRuler:
		return "NSVerticalRuler"
	default:
		return fmt.Sprintf("NSRulerOrientation(%d)", e)
	}
}




type NSRun uint

const (
	// Deprecated.
	NSRunAbortedResponse NSRun = 1
	// Deprecated.
	NSRunContinuesResponse NSRun = 2
	// Deprecated.
	NSRunStoppedResponse NSRun = 0
)


func (e NSRun) String() string {
	switch e {
	case NSRunAbortedResponse:
		return "NSRunAbortedResponse"
	case NSRunContinuesResponse:
		return "NSRunContinuesResponse"
	case NSRunStoppedResponse:
		return "NSRunStoppedResponse"
	default:
		return fmt.Sprintf("NSRun(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSDocument/SaveOperationType
type NSSaveOperationType int

const (
	// NSAutosaveAsOperation: An operation that writes a document’s contents to a new file or file package even though the user has not explicitly requested it, then changes the document’s current location to point to the just-written file or file package.
	NSAutosaveAsOperation NSSaveOperationType = 5
	// NSAutosaveElsewhereOperation: An operation that writes an autosave version of the file to a different location.
	NSAutosaveElsewhereOperation NSSaveOperationType = 3
	// NSAutosaveInPlaceOperation: An operation that overwrites the document’s current contents with autosave data.
	NSAutosaveInPlaceOperation NSSaveOperationType = 4
	// NSSaveAsOperation: An operation that writes the document’s contents to a new location and updates the document to point to that location
	NSSaveAsOperation NSSaveOperationType = 1
	// NSSaveOperation: An operation that overwrites a document’s file or file package with the document’s contents.
	NSSaveOperation NSSaveOperationType = 0
	// NSSaveToOperation: An operation that writes a copy of the document’s contents to the specified location, without changing the original document’s location.
	NSSaveToOperation NSSaveOperationType = 2
)


func (e NSSaveOperationType) String() string {
	switch e {
	case NSAutosaveAsOperation:
		return "NSAutosaveAsOperation"
	case NSAutosaveElsewhereOperation:
		return "NSAutosaveElsewhereOperation"
	case NSAutosaveInPlaceOperation:
		return "NSAutosaveInPlaceOperation"
	case NSSaveAsOperation:
		return "NSSaveAsOperation"
	case NSSaveOperation:
		return "NSSaveOperation"
	case NSSaveToOperation:
		return "NSSaveToOperation"
	default:
		return fmt.Sprintf("NSSaveOperationType(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSScroller/ArrowPosition
type NSScrollArrowPosition int

const (
	// NSScrollerArrowsMaxEnd: Buttons at bottom or right.
	NSScrollerArrowsMaxEnd NSScrollArrowPosition = 0
	// NSScrollerArrowsMinEnd: Buttons at top or left.
	NSScrollerArrowsMinEnd NSScrollArrowPosition = 1
	// NSScrollerArrowsNone: No buttons.
	NSScrollerArrowsNone NSScrollArrowPosition = 2
)


func (e NSScrollArrowPosition) String() string {
	switch e {
	case NSScrollerArrowsMaxEnd:
		return "NSScrollerArrowsMaxEnd"
	case NSScrollerArrowsMinEnd:
		return "NSScrollerArrowsMinEnd"
	case NSScrollerArrowsNone:
		return "NSScrollerArrowsNone"
	default:
		return fmt.Sprintf("NSScrollArrowPosition(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSScrollView/Elasticity
type NSScrollElasticity int

const (
	// NSScrollElasticityAllowed: Allow content to be scrolled past its bounds on this axis in an elastic fashion.
	NSScrollElasticityAllowed NSScrollElasticity = 2
	// NSScrollElasticityAutomatic: Automatically determine whether to allow elasticity on this axis.
	NSScrollElasticityAutomatic NSScrollElasticity = 0
	// NSScrollElasticityNone: Disallow scrolling beyond document bounds on this axis.
	NSScrollElasticityNone NSScrollElasticity = 1
)


func (e NSScrollElasticity) String() string {
	switch e {
	case NSScrollElasticityAllowed:
		return "NSScrollElasticityAllowed"
	case NSScrollElasticityAutomatic:
		return "NSScrollElasticityAutomatic"
	case NSScrollElasticityNone:
		return "NSScrollElasticityNone"
	default:
		return fmt.Sprintf("NSScrollElasticity(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSScrollView/FindBarPosition-swift.enum
type NSScrollViewFindBarPosition int

const (
	// NSScrollViewFindBarPositionAboveContent: The find bar is displayed above the scroll view content.
	NSScrollViewFindBarPositionAboveContent NSScrollViewFindBarPosition = 1
	// NSScrollViewFindBarPositionAboveHorizontalRuler: The find bar is displayed above the horizontal ruler, if visible.
	NSScrollViewFindBarPositionAboveHorizontalRuler NSScrollViewFindBarPosition = 0
	// NSScrollViewFindBarPositionBelowContent: The find bar is displayed below the scroll view content.
	NSScrollViewFindBarPositionBelowContent NSScrollViewFindBarPosition = 2
)


func (e NSScrollViewFindBarPosition) String() string {
	switch e {
	case NSScrollViewFindBarPositionAboveContent:
		return "NSScrollViewFindBarPositionAboveContent"
	case NSScrollViewFindBarPositionAboveHorizontalRuler:
		return "NSScrollViewFindBarPositionAboveHorizontalRuler"
	case NSScrollViewFindBarPositionBelowContent:
		return "NSScrollViewFindBarPositionBelowContent"
	default:
		return fmt.Sprintf("NSScrollViewFindBarPosition(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSScroller/Arrow
type NSScrollerArrow int

const (
	// NSScrollerDecrementArrow: The up or left scroll button.
	NSScrollerDecrementArrow NSScrollerArrow = 1
	// NSScrollerIncrementArrow: The down or right scroll button.
	NSScrollerIncrementArrow NSScrollerArrow = 0
)


func (e NSScrollerArrow) String() string {
	switch e {
	case NSScrollerDecrementArrow:
		return "NSScrollerDecrementArrow"
	case NSScrollerIncrementArrow:
		return "NSScrollerIncrementArrow"
	default:
		return fmt.Sprintf("NSScrollerArrow(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSScroller/KnobStyle-swift.enum
type NSScrollerKnobStyle int

const (
	// NSScrollerKnobStyleDark: Specifies a dark knob.
	NSScrollerKnobStyleDark NSScrollerKnobStyle = 1
	// NSScrollerKnobStyleDefault: Specifies a dark knob with a light border.
	NSScrollerKnobStyleDefault NSScrollerKnobStyle = 0
	// NSScrollerKnobStyleLight: Specifies a light knob.
	NSScrollerKnobStyleLight NSScrollerKnobStyle = 2
)


func (e NSScrollerKnobStyle) String() string {
	switch e {
	case NSScrollerKnobStyleDark:
		return "NSScrollerKnobStyleDark"
	case NSScrollerKnobStyleDefault:
		return "NSScrollerKnobStyleDefault"
	case NSScrollerKnobStyleLight:
		return "NSScrollerKnobStyleLight"
	default:
		return fmt.Sprintf("NSScrollerKnobStyle(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSScroller/Part
type NSScrollerPart int

const (
	// NSScrollerDecrementLine: Up or left by a small amount.
	NSScrollerDecrementLine NSScrollerPart = 4
	// NSScrollerDecrementPage: Up or left by a large amount.
	NSScrollerDecrementPage NSScrollerPart = 1
	// NSScrollerIncrementLine: Down or right by a small amount.
	NSScrollerIncrementLine NSScrollerPart = 5
	// NSScrollerIncrementPage: Down or right by a large amount.
	NSScrollerIncrementPage NSScrollerPart = 3
	// NSScrollerKnob: Directly to the scroller’s value, as given by .
	NSScrollerKnob NSScrollerPart = 2
	// NSScrollerKnobSlot: Directly to the scroller’s value, as given by .
	NSScrollerKnobSlot NSScrollerPart = 6
	// NSScrollerNoPart: Don’t scroll at all.
	NSScrollerNoPart NSScrollerPart = 0
)


func (e NSScrollerPart) String() string {
	switch e {
	case NSScrollerDecrementLine:
		return "NSScrollerDecrementLine"
	case NSScrollerDecrementPage:
		return "NSScrollerDecrementPage"
	case NSScrollerIncrementLine:
		return "NSScrollerIncrementLine"
	case NSScrollerIncrementPage:
		return "NSScrollerIncrementPage"
	case NSScrollerKnob:
		return "NSScrollerKnob"
	case NSScrollerKnobSlot:
		return "NSScrollerKnobSlot"
	case NSScrollerNoPart:
		return "NSScrollerNoPart"
	default:
		return fmt.Sprintf("NSScrollerPart(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSScroller/Style
type NSScrollerStyle int

const (
	// NSScrollerStyleLegacy: Specifies legacy-style scrollers as prior to macOS 10.7.
	NSScrollerStyleLegacy NSScrollerStyle = 0
	// NSScrollerStyleOverlay: Specifies overlay-style scrollers in macOS 10.7 and later.
	NSScrollerStyleOverlay NSScrollerStyle = 1
)


func (e NSScrollerStyle) String() string {
	switch e {
	case NSScrollerStyleLegacy:
		return "NSScrollerStyleLegacy"
	case NSScrollerStyleOverlay:
		return "NSScrollerStyleOverlay"
	default:
		return fmt.Sprintf("NSScrollerStyle(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSScrubber/Alignment
type NSScrubberAlignment int

const (
	// NSScrubberAlignmentCenter: Center alignment of items within the scrubber.
	NSScrubberAlignmentCenter NSScrubberAlignment = 3
	// NSScrubberAlignmentLeading: Leading alignment of items within the scrubber.
	NSScrubberAlignmentLeading NSScrubberAlignment = 1
	// NSScrubberAlignmentNone: No preference for item alignment.
	NSScrubberAlignmentNone NSScrubberAlignment = 0
	// NSScrubberAlignmentTrailing: Trailing alignment of items within the scrubber.
	NSScrubberAlignmentTrailing NSScrubberAlignment = 2
)


func (e NSScrubberAlignment) String() string {
	switch e {
	case NSScrubberAlignmentCenter:
		return "NSScrubberAlignmentCenter"
	case NSScrubberAlignmentLeading:
		return "NSScrubberAlignmentLeading"
	case NSScrubberAlignmentNone:
		return "NSScrubberAlignmentNone"
	case NSScrubberAlignmentTrailing:
		return "NSScrubberAlignmentTrailing"
	default:
		return fmt.Sprintf("NSScrubberAlignment(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSScrubber/Mode-swift.enum
type NSScrubberMode int

const (
	// NSScrubberModeFixed: A scrolling mode in which scrubber items remain fixed in place, and the item under the user’s finger is highlighted.
	NSScrubberModeFixed NSScrubberMode = 0
	// NSScrubberModeFree: A scrolling mode in which the scrubber scrolls as the user swipes horizontally across the scrubber.
	NSScrubberModeFree NSScrubberMode = 1
)


func (e NSScrubberMode) String() string {
	switch e {
	case NSScrubberModeFixed:
		return "NSScrubberModeFixed"
	case NSScrubberModeFree:
		return "NSScrubberModeFree"
	default:
		return fmt.Sprintf("NSScrubberMode(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/Distribution
type NSSegmentDistribution int

const (
	NSSegmentDistributionFill NSSegmentDistribution = 1
	NSSegmentDistributionFillEqually NSSegmentDistribution = 2
	NSSegmentDistributionFillProportionally NSSegmentDistribution = 3
	NSSegmentDistributionFit NSSegmentDistribution = 0
)


func (e NSSegmentDistribution) String() string {
	switch e {
	case NSSegmentDistributionFill:
		return "NSSegmentDistributionFill"
	case NSSegmentDistributionFillEqually:
		return "NSSegmentDistributionFillEqually"
	case NSSegmentDistributionFillProportionally:
		return "NSSegmentDistributionFillProportionally"
	case NSSegmentDistributionFit:
		return "NSSegmentDistributionFit"
	default:
		return fmt.Sprintf("NSSegmentDistribution(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/Style
type NSSegmentStyle int

const (
	// NSSegmentStyleAutomatic: The appearance of the segmented control is automatically determined based on the type of window in which the control is displayed and the position within the window.
	NSSegmentStyleAutomatic NSSegmentStyle = 0
	// NSSegmentStyleCapsule: The control is displayed using the capsule style.
	NSSegmentStyleCapsule NSSegmentStyle = 5
	// NSSegmentStyleRoundRect: The control is displayed using the round rect style.
	NSSegmentStyleRoundRect NSSegmentStyle = 3
	// NSSegmentStyleRounded: The control is displayed using the rounded style.
	NSSegmentStyleRounded NSSegmentStyle = 1
	// NSSegmentStyleSeparated: The segments in the control are displayed very close to each other but not touching.
	NSSegmentStyleSeparated NSSegmentStyle = 8
	// NSSegmentStyleSmallSquare: The control is displayed using the small square style.
	NSSegmentStyleSmallSquare NSSegmentStyle = 6
	// NSSegmentStyleTexturedRounded: The control is displayed using the textured rounded style.
	NSSegmentStyleTexturedRounded NSSegmentStyle = 2
	// NSSegmentStyleTexturedSquare: The control is displayed using the textured square style.
	NSSegmentStyleTexturedSquare NSSegmentStyle = 4
)


func (e NSSegmentStyle) String() string {
	switch e {
	case NSSegmentStyleAutomatic:
		return "NSSegmentStyleAutomatic"
	case NSSegmentStyleCapsule:
		return "NSSegmentStyleCapsule"
	case NSSegmentStyleRoundRect:
		return "NSSegmentStyleRoundRect"
	case NSSegmentStyleRounded:
		return "NSSegmentStyleRounded"
	case NSSegmentStyleSeparated:
		return "NSSegmentStyleSeparated"
	case NSSegmentStyleSmallSquare:
		return "NSSegmentStyleSmallSquare"
	case NSSegmentStyleTexturedRounded:
		return "NSSegmentStyleTexturedRounded"
	case NSSegmentStyleTexturedSquare:
		return "NSSegmentStyleTexturedSquare"
	default:
		return fmt.Sprintf("NSSegmentStyle(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/SwitchTracking
type NSSegmentSwitchTracking int

const (
	// NSSegmentSwitchTrackingMomentary: A segment is selected only when the user is pressing the mouse down within the bounds of the segment.
	NSSegmentSwitchTrackingMomentary NSSegmentSwitchTracking = 2
	// NSSegmentSwitchTrackingMomentaryAccelerator: On pressure-sensitive systems, when the user force clicks a segment, a momentary accelerator segmented control sends repeating actions as pressure changes occur.
	NSSegmentSwitchTrackingMomentaryAccelerator NSSegmentSwitchTracking = 3
	// NSSegmentSwitchTrackingSelectAny: One or more segment cells in the control can be selected at a time.
	NSSegmentSwitchTrackingSelectAny NSSegmentSwitchTracking = 1
	// NSSegmentSwitchTrackingSelectOne: Only one segment in the control can be selected at a time.
	NSSegmentSwitchTrackingSelectOne NSSegmentSwitchTracking = 0
)


func (e NSSegmentSwitchTracking) String() string {
	switch e {
	case NSSegmentSwitchTrackingMomentary:
		return "NSSegmentSwitchTrackingMomentary"
	case NSSegmentSwitchTrackingMomentaryAccelerator:
		return "NSSegmentSwitchTrackingMomentaryAccelerator"
	case NSSegmentSwitchTrackingSelectAny:
		return "NSSegmentSwitchTrackingSelectAny"
	case NSSegmentSwitchTrackingSelectOne:
		return "NSSegmentSwitchTrackingSelectOne"
	default:
		return fmt.Sprintf("NSSegmentSwitchTracking(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSSelectionAffinity
type NSSelectionAffinity int

const (
	// NSSelectionAffinityDownstream: The selection is moving toward the bottom of the document.
	NSSelectionAffinityDownstream NSSelectionAffinity = 1
	// NSSelectionAffinityUpstream: The selection is moving toward the top of the document.
	NSSelectionAffinityUpstream NSSelectionAffinity = 0
)


func (e NSSelectionAffinity) String() string {
	switch e {
	case NSSelectionAffinityDownstream:
		return "NSSelectionAffinityDownstream"
	case NSSelectionAffinityUpstream:
		return "NSSelectionAffinityUpstream"
	default:
		return fmt.Sprintf("NSSelectionAffinity(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSWindow/SelectionDirection
type NSSelectionDirection int

const (
	// NSDirectSelection: The window isn’t traversing the key view loop.
	NSDirectSelection NSSelectionDirection = 0
	// NSSelectingNext: The window is proceeding to the next valid key view.
	NSSelectingNext NSSelectionDirection = 1
	// NSSelectingPrevious: The window is proceeding to the previous valid key view.
	NSSelectingPrevious NSSelectionDirection = 2
)


func (e NSSelectionDirection) String() string {
	switch e {
	case NSDirectSelection:
		return "NSDirectSelection"
	case NSSelectingNext:
		return "NSSelectingNext"
	case NSSelectingPrevious:
		return "NSSelectingPrevious"
	default:
		return fmt.Sprintf("NSSelectionDirection(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSSelectionGranularity
type NSSelectionGranularity int

const (
	// NSSelectByCharacter: Extends the selection character by character.
	NSSelectByCharacter NSSelectionGranularity = 0
	// NSSelectByParagraph: Extends the selection paragraph by paragraph.
	NSSelectByParagraph NSSelectionGranularity = 2
	// NSSelectByWord: Extends the selection word by word.
	NSSelectByWord NSSelectionGranularity = 1
)


func (e NSSelectionGranularity) String() string {
	switch e {
	case NSSelectByCharacter:
		return "NSSelectByCharacter"
	case NSSelectByParagraph:
		return "NSSelectByParagraph"
	case NSSelectByWord:
		return "NSSelectByWord"
	default:
		return fmt.Sprintf("NSSelectionGranularity(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSSharingCollaborationMode
type NSSharingCollaborationMode int

const (
	NSSharingCollaborationModeCollaborate NSSharingCollaborationMode = 1
	NSSharingCollaborationModeSendCopy NSSharingCollaborationMode = 0
)


func (e NSSharingCollaborationMode) String() string {
	switch e {
	case NSSharingCollaborationModeCollaborate:
		return "NSSharingCollaborationModeCollaborate"
	case NSSharingCollaborationModeSendCopy:
		return "NSSharingCollaborationModeSendCopy"
	default:
		return fmt.Sprintf("NSSharingCollaborationMode(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSSharingService/SharingContentScope
type NSSharingContentScope int

const (
	// NSSharingContentScopeFull: Used when sharing the whole content of the current document, for example, the URL of the webpage.
	NSSharingContentScopeFull NSSharingContentScope = 2
	// NSSharingContentScopeItem: Used when sharing a clearly identified item, for example, a file represented by its icon.
	NSSharingContentScopeItem NSSharingContentScope = 0
	// NSSharingContentScopePartial: Used when sharing a portion of a more global content, for example, part of a webpage.
	NSSharingContentScopePartial NSSharingContentScope = 1
)


func (e NSSharingContentScope) String() string {
	switch e {
	case NSSharingContentScopeFull:
		return "NSSharingContentScopeFull"
	case NSSharingContentScopeItem:
		return "NSSharingContentScopeItem"
	case NSSharingContentScopePartial:
		return "NSSharingContentScopePartial"
	default:
		return fmt.Sprintf("NSSharingContentScope(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSSlider/SliderType-swift.enum
type NSSliderType int

const (
	// NSSliderTypeCircular: A dial representing an angular range.
	NSSliderTypeCircular NSSliderType = 1
	// NSSliderTypeLinear: A bar representing a range, and a knob indicating the currently selected value.
	NSSliderTypeLinear NSSliderType = 0
)


func (e NSSliderType) String() string {
	switch e {
	case NSSliderTypeCircular:
		return "NSSliderTypeCircular"
	case NSSliderTypeLinear:
		return "NSSliderTypeLinear"
	default:
		return fmt.Sprintf("NSSliderType(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/Boundary
type NSSpeechBoundary int

const (
	// NSSpeechImmediateBoundary: Speech should be paused or stopped immediately.
	NSSpeechImmediateBoundary NSSpeechBoundary = 0
	// NSSpeechSentenceBoundary: Speech should be paused or stopped at the end of the sentence.
	NSSpeechSentenceBoundary NSSpeechBoundary = 2
	// NSSpeechWordBoundary: Speech should be paused or stopped at the end of the word.
	NSSpeechWordBoundary NSSpeechBoundary = 1
)


func (e NSSpeechBoundary) String() string {
	switch e {
	case NSSpeechImmediateBoundary:
		return "NSSpeechImmediateBoundary"
	case NSSpeechSentenceBoundary:
		return "NSSpeechSentenceBoundary"
	case NSSpeechWordBoundary:
		return "NSSpeechWordBoundary"
	default:
		return fmt.Sprintf("NSSpeechBoundary(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSSpellingState
type NSSpellingState int

const (
	NSSpellingStateGrammarFlag NSSpellingState = 2
	NSSpellingStateSpellingFlag NSSpellingState = 1
)


func (e NSSpellingState) String() string {
	switch e {
	case NSSpellingStateGrammarFlag:
		return "NSSpellingStateGrammarFlag"
	case NSSpellingStateSpellingFlag:
		return "NSSpellingStateSpellingFlag"
	default:
		return fmt.Sprintf("NSSpellingState(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSSplitView/DividerStyle-swift.enum
type NSSplitViewDividerStyle int

const (
	// NSSplitViewDividerStylePaneSplitter: A thick style divider with a 3D appearance displays between subviews.
	NSSplitViewDividerStylePaneSplitter NSSplitViewDividerStyle = 3
	// NSSplitViewDividerStyleThick: A thick style divider displays between subviews.
	NSSplitViewDividerStyleThick NSSplitViewDividerStyle = 1
	// NSSplitViewDividerStyleThin: A thin style divider displays between subviews.
	NSSplitViewDividerStyleThin NSSplitViewDividerStyle = 2
)


func (e NSSplitViewDividerStyle) String() string {
	switch e {
	case NSSplitViewDividerStylePaneSplitter:
		return "NSSplitViewDividerStylePaneSplitter"
	case NSSplitViewDividerStyleThick:
		return "NSSplitViewDividerStyleThick"
	case NSSplitViewDividerStyleThin:
		return "NSSplitViewDividerStyleThin"
	default:
		return fmt.Sprintf("NSSplitViewDividerStyle(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/Behavior-swift.enum
type NSSplitViewItemBehavior int

const (
	// NSSplitViewItemBehaviorContentList: The content list behavior.
	NSSplitViewItemBehaviorContentList NSSplitViewItemBehavior = 2
	// NSSplitViewItemBehaviorDefault: The default split view item behavior.
	NSSplitViewItemBehaviorDefault NSSplitViewItemBehavior = 0
	// NSSplitViewItemBehaviorInspector: The inspector behavior.
	NSSplitViewItemBehaviorInspector NSSplitViewItemBehavior = 3
	// NSSplitViewItemBehaviorSidebar: The sidebar behavior.
	NSSplitViewItemBehaviorSidebar NSSplitViewItemBehavior = 1
)


func (e NSSplitViewItemBehavior) String() string {
	switch e {
	case NSSplitViewItemBehaviorContentList:
		return "NSSplitViewItemBehaviorContentList"
	case NSSplitViewItemBehaviorDefault:
		return "NSSplitViewItemBehaviorDefault"
	case NSSplitViewItemBehaviorInspector:
		return "NSSplitViewItemBehaviorInspector"
	case NSSplitViewItemBehaviorSidebar:
		return "NSSplitViewItemBehaviorSidebar"
	default:
		return fmt.Sprintf("NSSplitViewItemBehavior(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/CollapseBehavior-swift.enum
type NSSplitViewItemCollapseBehavior int

const (
	// NSSplitViewItemCollapseBehaviorDefault: The item uses the default collapsing behavior.
	NSSplitViewItemCollapseBehaviorDefault NSSplitViewItemCollapseBehavior = 0
	// NSSplitViewItemCollapseBehaviorPreferResizingSiblingsWithFixedSplitView: The item’s preference is to resize the other split panes.
	NSSplitViewItemCollapseBehaviorPreferResizingSiblingsWithFixedSplitView NSSplitViewItemCollapseBehavior = 2
	// NSSplitViewItemCollapseBehaviorPreferResizingSplitViewWithFixedSiblings: The item’s preference is to keep the other panes at their current size and position onscreen, potentially growing or shrinking the window in the direction to best preserve that.
	NSSplitViewItemCollapseBehaviorPreferResizingSplitViewWithFixedSiblings NSSplitViewItemCollapseBehavior = 1
	// NSSplitViewItemCollapseBehaviorUseConstraints: The item collapses and expands using a constraint animation, with a constraint priority of the item’s holding priority.
	NSSplitViewItemCollapseBehaviorUseConstraints NSSplitViewItemCollapseBehavior = 3
)


func (e NSSplitViewItemCollapseBehavior) String() string {
	switch e {
	case NSSplitViewItemCollapseBehaviorDefault:
		return "NSSplitViewItemCollapseBehaviorDefault"
	case NSSplitViewItemCollapseBehaviorPreferResizingSiblingsWithFixedSplitView:
		return "NSSplitViewItemCollapseBehaviorPreferResizingSiblingsWithFixedSplitView"
	case NSSplitViewItemCollapseBehaviorPreferResizingSplitViewWithFixedSiblings:
		return "NSSplitViewItemCollapseBehaviorPreferResizingSplitViewWithFixedSiblings"
	case NSSplitViewItemCollapseBehaviorUseConstraints:
		return "NSSplitViewItemCollapseBehaviorUseConstraints"
	default:
		return fmt.Sprintf("NSSplitViewItemCollapseBehavior(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSSpringLoadingHighlight
type NSSpringLoadingHighlight int

const (
	// NSSpringLoadingHighlightEmphasized: A constant that indicates emphasized highlighting to show active spring-loading on the destination.
	NSSpringLoadingHighlightEmphasized NSSpringLoadingHighlight = 2
	// NSSpringLoadingHighlightNone: A constant that indicates no highlighting.
	NSSpringLoadingHighlightNone NSSpringLoadingHighlight = 0
	// NSSpringLoadingHighlightStandard: A constant that indicates standard highlighting to show the destination supports spring-loading.
	NSSpringLoadingHighlightStandard NSSpringLoadingHighlight = 1
)


func (e NSSpringLoadingHighlight) String() string {
	switch e {
	case NSSpringLoadingHighlightEmphasized:
		return "NSSpringLoadingHighlightEmphasized"
	case NSSpringLoadingHighlightNone:
		return "NSSpringLoadingHighlightNone"
	case NSSpringLoadingHighlightStandard:
		return "NSSpringLoadingHighlightStandard"
	default:
		return fmt.Sprintf("NSSpringLoadingHighlight(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSSpringLoadingOptions
type NSSpringLoadingOptions int

const (
	// NSSpringLoadingContinuousActivation: Spring-loading on the destination object is enabled.
	NSSpringLoadingContinuousActivation NSSpringLoadingOptions = 2
	// NSSpringLoadingDisabled: Spring-loading on the destination object is disabled.
	NSSpringLoadingDisabled NSSpringLoadingOptions = 0
	// NSSpringLoadingEnabled: Spring-loading on the destination object is enabled.
	NSSpringLoadingEnabled NSSpringLoadingOptions = 1
	// NSSpringLoadingNoHover: Spring-loading on the destination object is enabled, but cannot be invoked by hovering.
	NSSpringLoadingNoHover NSSpringLoadingOptions = 8
)


func (e NSSpringLoadingOptions) String() string {
	switch e {
	case NSSpringLoadingContinuousActivation:
		return "NSSpringLoadingContinuousActivation"
	case NSSpringLoadingDisabled:
		return "NSSpringLoadingDisabled"
	case NSSpringLoadingEnabled:
		return "NSSpringLoadingEnabled"
	case NSSpringLoadingNoHover:
		return "NSSpringLoadingNoHover"
	default:
		return fmt.Sprintf("NSSpringLoadingOptions(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSStackView/Distribution-swift.enum
type NSStackViewDistribution int

const (
	NSStackViewDistributionEqualCentering NSStackViewDistribution = 4
	NSStackViewDistributionEqualSpacing NSStackViewDistribution = 3
	NSStackViewDistributionFill NSStackViewDistribution = 0
	NSStackViewDistributionFillEqually NSStackViewDistribution = 1
	NSStackViewDistributionFillProportionally NSStackViewDistribution = 2
	NSStackViewDistributionGravityAreas NSStackViewDistribution = -1
)


func (e NSStackViewDistribution) String() string {
	switch e {
	case NSStackViewDistributionEqualCentering:
		return "NSStackViewDistributionEqualCentering"
	case NSStackViewDistributionEqualSpacing:
		return "NSStackViewDistributionEqualSpacing"
	case NSStackViewDistributionFill:
		return "NSStackViewDistributionFill"
	case NSStackViewDistributionFillEqually:
		return "NSStackViewDistributionFillEqually"
	case NSStackViewDistributionFillProportionally:
		return "NSStackViewDistributionFillProportionally"
	case NSStackViewDistributionGravityAreas:
		return "NSStackViewDistributionGravityAreas"
	default:
		return fmt.Sprintf("NSStackViewDistribution(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSStackView/Gravity
type NSStackViewGravity int

const (
	// NSStackViewGravityBottom: The bottommost gravity area in a vertically oriented stack view.
	NSStackViewGravityBottom NSStackViewGravity = 3
	// NSStackViewGravityCenter: The center gravity area, regardless of stack view layout direction or user interface language.
	NSStackViewGravityCenter NSStackViewGravity = 2
	// NSStackViewGravityTop: The topmost gravity area in a vertically oriented stack view.
	NSStackViewGravityTop NSStackViewGravity = 1
)


func (e NSStackViewGravity) String() string {
	switch e {
	case NSStackViewGravityBottom:
		return "NSStackViewGravityBottom"
	case NSStackViewGravityCenter:
		return "NSStackViewGravityCenter"
	case NSStackViewGravityTop:
		return "NSStackViewGravityTop"
	default:
		return fmt.Sprintf("NSStackViewGravity(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSStatusItem/Behavior-swift.struct
type NSStatusItemBehavior int

const (
	// NSStatusItemBehaviorRemovalAllowed: A status item that allows interactive removal.
	NSStatusItemBehaviorRemovalAllowed NSStatusItemBehavior = 2
	// NSStatusItemBehaviorTerminationOnRemoval: A status item that quits the application upon removal.
	NSStatusItemBehaviorTerminationOnRemoval NSStatusItemBehavior = 4
)


func (e NSStatusItemBehavior) String() string {
	switch e {
	case NSStatusItemBehaviorRemovalAllowed:
		return "NSStatusItemBehaviorRemovalAllowed"
	case NSStatusItemBehaviorTerminationOnRemoval:
		return "NSStatusItemBehaviorTerminationOnRemoval"
	default:
		return fmt.Sprintf("NSStatusItemBehavior(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSStringDrawingOptions
type NSStringDrawingOptions int

const (
	NSStringDrawingDisableScreenFontSubstitution NSStringDrawingOptions = 4
	NSStringDrawingOneShot NSStringDrawingOptions = 16
	NSStringDrawingOptionsResolvesNaturalAlignmentWithBaseWritingDirection NSStringDrawingOptions = 512
	NSStringDrawingTruncatesLastVisibleLine NSStringDrawingOptions = 32
	NSStringDrawingUsesDeviceMetrics NSStringDrawingOptions = 8
	NSStringDrawingUsesFontLeading NSStringDrawingOptions = 2
	NSStringDrawingUsesLineFragmentOrigin NSStringDrawingOptions = 1
)


func (e NSStringDrawingOptions) String() string {
	switch e {
	case NSStringDrawingDisableScreenFontSubstitution:
		return "NSStringDrawingDisableScreenFontSubstitution"
	case NSStringDrawingOneShot:
		return "NSStringDrawingOneShot"
	case NSStringDrawingOptionsResolvesNaturalAlignmentWithBaseWritingDirection:
		return "NSStringDrawingOptionsResolvesNaturalAlignmentWithBaseWritingDirection"
	case NSStringDrawingTruncatesLastVisibleLine:
		return "NSStringDrawingTruncatesLastVisibleLine"
	case NSStringDrawingUsesDeviceMetrics:
		return "NSStringDrawingUsesDeviceMetrics"
	case NSStringDrawingUsesFontLeading:
		return "NSStringDrawingUsesFontLeading"
	case NSStringDrawingUsesLineFragmentOrigin:
		return "NSStringDrawingUsesLineFragmentOrigin"
	default:
		return fmt.Sprintf("NSStringDrawingOptions(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/TIFFCompression
type NSTIFFCompression int

const (
	// NSTIFFCompressionCCITTFAX3: CCITT Fax Group 3 compression.
	NSTIFFCompressionCCITTFAX3 NSTIFFCompression = 3
	// NSTIFFCompressionCCITTFAX4: CCITT Fax Group 4 compression.
	NSTIFFCompressionCCITTFAX4 NSTIFFCompression = 4
	// NSTIFFCompressionJPEG: JPEG compression.
	NSTIFFCompressionJPEG NSTIFFCompression = 6
	// NSTIFFCompressionLZW: LZW compression.
	NSTIFFCompressionLZW NSTIFFCompression = 5
	// NSTIFFCompressionNEXT: NeXT compressed.
	NSTIFFCompressionNEXT NSTIFFCompression = 32766
	// NSTIFFCompressionNone: No compression.
	NSTIFFCompressionNone NSTIFFCompression = 1
	// NSTIFFCompressionOldJPEG: Old JPEG compression.
	NSTIFFCompressionOldJPEG NSTIFFCompression = 32865
	// NSTIFFCompressionPackBits: PackBits compression.
	NSTIFFCompressionPackBits NSTIFFCompression = 32773
)


func (e NSTIFFCompression) String() string {
	switch e {
	case NSTIFFCompressionCCITTFAX3:
		return "NSTIFFCompressionCCITTFAX3"
	case NSTIFFCompressionCCITTFAX4:
		return "NSTIFFCompressionCCITTFAX4"
	case NSTIFFCompressionJPEG:
		return "NSTIFFCompressionJPEG"
	case NSTIFFCompressionLZW:
		return "NSTIFFCompressionLZW"
	case NSTIFFCompressionNEXT:
		return "NSTIFFCompressionNEXT"
	case NSTIFFCompressionNone:
		return "NSTIFFCompressionNone"
	case NSTIFFCompressionOldJPEG:
		return "NSTIFFCompressionOldJPEG"
	case NSTIFFCompressionPackBits:
		return "NSTIFFCompressionPackBits"
	default:
		return fmt.Sprintf("NSTIFFCompression(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSTabView/TabPosition-swift.enum
type NSTabPosition int

const (
	NSTabPositionBottom NSTabPosition = 3
	NSTabPositionLeft NSTabPosition = 2
	NSTabPositionNone NSTabPosition = 0
	NSTabPositionRight NSTabPosition = 4
	NSTabPositionTop NSTabPosition = 1
)


func (e NSTabPosition) String() string {
	switch e {
	case NSTabPositionBottom:
		return "NSTabPositionBottom"
	case NSTabPositionLeft:
		return "NSTabPositionLeft"
	case NSTabPositionNone:
		return "NSTabPositionNone"
	case NSTabPositionRight:
		return "NSTabPositionRight"
	case NSTabPositionTop:
		return "NSTabPositionTop"
	default:
		return fmt.Sprintf("NSTabPosition(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSTabViewItem/State
type NSTabState int

const (
	// NSBackgroundTab: A tab that’s not being displayed.
	NSBackgroundTab NSTabState = 1
	// NSPressedTab: A tab that the user is in the process of clicking.
	NSPressedTab NSTabState = 2
	// NSSelectedTab: The tab that’s being displayed.
	NSSelectedTab NSTabState = 0
)


func (e NSTabState) String() string {
	switch e {
	case NSBackgroundTab:
		return "NSBackgroundTab"
	case NSPressedTab:
		return "NSPressedTab"
	case NSSelectedTab:
		return "NSSelectedTab"
	default:
		return fmt.Sprintf("NSTabState(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSTabView/TabViewBorderType-swift.enum
type NSTabViewBorderType int

const (
	NSTabViewBorderTypeBezel NSTabViewBorderType = 2
	NSTabViewBorderTypeLine NSTabViewBorderType = 1
	NSTabViewBorderTypeNone NSTabViewBorderType = 0
)


func (e NSTabViewBorderType) String() string {
	switch e {
	case NSTabViewBorderTypeBezel:
		return "NSTabViewBorderTypeBezel"
	case NSTabViewBorderTypeLine:
		return "NSTabViewBorderTypeLine"
	case NSTabViewBorderTypeNone:
		return "NSTabViewBorderTypeNone"
	default:
		return fmt.Sprintf("NSTabViewBorderType(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSTabViewController/TabStyle-swift.enum
type NSTabViewControllerTabStyle int

const (
	// NSTabViewControllerTabStyleSegmentedControlOnBottom: A style that displays a segmented control along the bottom edge of the tab view interface.
	NSTabViewControllerTabStyleSegmentedControlOnBottom NSTabViewControllerTabStyle = 1
	// NSTabViewControllerTabStyleSegmentedControlOnTop: A style that displays a segmented control along the top edge of the tab view interface.
	NSTabViewControllerTabStyleSegmentedControlOnTop NSTabViewControllerTabStyle = 0
	// NSTabViewControllerTabStyleToolbar: A style that automatically adds any tabs to the window’s toolbar.
	NSTabViewControllerTabStyleToolbar NSTabViewControllerTabStyle = 2
	// NSTabViewControllerTabStyleUnspecified: A style that indicates the tab view controller does not provide the tab selection UI.
	NSTabViewControllerTabStyleUnspecified NSTabViewControllerTabStyle = -1
)


func (e NSTabViewControllerTabStyle) String() string {
	switch e {
	case NSTabViewControllerTabStyleSegmentedControlOnBottom:
		return "NSTabViewControllerTabStyleSegmentedControlOnBottom"
	case NSTabViewControllerTabStyleSegmentedControlOnTop:
		return "NSTabViewControllerTabStyleSegmentedControlOnTop"
	case NSTabViewControllerTabStyleToolbar:
		return "NSTabViewControllerTabStyleToolbar"
	case NSTabViewControllerTabStyleUnspecified:
		return "NSTabViewControllerTabStyleUnspecified"
	default:
		return fmt.Sprintf("NSTabViewControllerTabStyle(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSTabView/TabType
type NSTabViewType int

const (
	// NSBottomTabsBezelBorder: Tabs are on the bottom of the view with a bezeled border.
	NSBottomTabsBezelBorder NSTabViewType = 2
	// NSLeftTabsBezelBorder: Tabs are on the left of the view with a bezeled border.
	NSLeftTabsBezelBorder NSTabViewType = 1
	// NSNoTabsBezelBorder: The view does not include tabs and has a bezeled border.
	NSNoTabsBezelBorder NSTabViewType = 4
	// NSNoTabsLineBorder: The view does not include tabs and has a lined border.
	NSNoTabsLineBorder NSTabViewType = 5
	// NSNoTabsNoBorder: The view does not include tabs and has no border.
	NSNoTabsNoBorder NSTabViewType = 6
	// NSRightTabsBezelBorder: Tabs are on the right of the view with a bezeled border.
	NSRightTabsBezelBorder NSTabViewType = 3
	// NSTopTabsBezelBorder: The view includes tabs on the top of the view and has a bezeled border (the default).
	NSTopTabsBezelBorder NSTabViewType = 0
)


func (e NSTabViewType) String() string {
	switch e {
	case NSBottomTabsBezelBorder:
		return "NSBottomTabsBezelBorder"
	case NSLeftTabsBezelBorder:
		return "NSLeftTabsBezelBorder"
	case NSNoTabsBezelBorder:
		return "NSNoTabsBezelBorder"
	case NSNoTabsLineBorder:
		return "NSNoTabsLineBorder"
	case NSNoTabsNoBorder:
		return "NSNoTabsNoBorder"
	case NSRightTabsBezelBorder:
		return "NSRightTabsBezelBorder"
	case NSTopTabsBezelBorder:
		return "NSTopTabsBezelBorder"
	default:
		return fmt.Sprintf("NSTabViewType(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSTableColumn/ResizingOptions
type NSTableColumnResizingOptions int

const (
	// NSTableColumnAutoresizingMask: Allows the table column to resize automatically in response to resizing the table view.
	NSTableColumnAutoresizingMask NSTableColumnResizingOptions = 1
	// NSTableColumnUserResizingMask: Allows the table column to be resized by the user.
	NSTableColumnUserResizingMask NSTableColumnResizingOptions = 2
)


func (e NSTableColumnResizingOptions) String() string {
	switch e {
	case NSTableColumnAutoresizingMask:
		return "NSTableColumnAutoresizingMask"
	case NSTableColumnUserResizingMask:
		return "NSTableColumnUserResizingMask"
	default:
		return fmt.Sprintf("NSTableColumnResizingOptions(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSTableView/RowActionEdge
type NSTableRowActionEdge int

const (
	// NSTableRowActionEdgeLeading: Denotes the leading, or left, edge of a table row view.
	NSTableRowActionEdgeLeading NSTableRowActionEdge = 0
	// NSTableRowActionEdgeTrailing: Denotes the trailing, or right, edge of a table row view.
	NSTableRowActionEdgeTrailing NSTableRowActionEdge = 1
)


func (e NSTableRowActionEdge) String() string {
	switch e {
	case NSTableRowActionEdgeLeading:
		return "NSTableRowActionEdgeLeading"
	case NSTableRowActionEdgeTrailing:
		return "NSTableRowActionEdgeTrailing"
	default:
		return fmt.Sprintf("NSTableRowActionEdge(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSTableView/AnimationOptions
type NSTableViewAnimationOptions int

const (
	// NSTableViewAnimationEffectFade: Use a fade for row or column removal.
	NSTableViewAnimationEffectFade NSTableViewAnimationOptions = 1
	// NSTableViewAnimationEffectGap: Creates a gap for newly inserted rows.
	NSTableViewAnimationEffectGap NSTableViewAnimationOptions = 2
	// NSTableViewAnimationSlideDown: Animates a row insertion or removal by sliding downward.
	NSTableViewAnimationSlideDown NSTableViewAnimationOptions = 32
	// NSTableViewAnimationSlideLeft: Animates a row insertion by sliding from the left.
	NSTableViewAnimationSlideLeft NSTableViewAnimationOptions = 48
	// NSTableViewAnimationSlideRight: Animates a row insertion by sliding from the right.
	NSTableViewAnimationSlideRight NSTableViewAnimationOptions = 64
	// NSTableViewAnimationSlideUp: Animates a row insertion or removal by sliding upward.
	NSTableViewAnimationSlideUp NSTableViewAnimationOptions = 16
)


func (e NSTableViewAnimationOptions) String() string {
	switch e {
	case NSTableViewAnimationEffectFade:
		return "NSTableViewAnimationEffectFade"
	case NSTableViewAnimationEffectGap:
		return "NSTableViewAnimationEffectGap"
	case NSTableViewAnimationSlideDown:
		return "NSTableViewAnimationSlideDown"
	case NSTableViewAnimationSlideLeft:
		return "NSTableViewAnimationSlideLeft"
	case NSTableViewAnimationSlideRight:
		return "NSTableViewAnimationSlideRight"
	case NSTableViewAnimationSlideUp:
		return "NSTableViewAnimationSlideUp"
	default:
		return fmt.Sprintf("NSTableViewAnimationOptions(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSTableView/ColumnAutoresizingStyle-swift.enum
type NSTableViewColumnAutoresizingStyle int

const (
	// NSTableViewFirstColumnOnlyAutoresizingStyle: Autoresize only the first table column.
	NSTableViewFirstColumnOnlyAutoresizingStyle NSTableViewColumnAutoresizingStyle = 5
	// NSTableViewLastColumnOnlyAutoresizingStyle: Autoresize only the last table column.
	NSTableViewLastColumnOnlyAutoresizingStyle NSTableViewColumnAutoresizingStyle = 4
	// NSTableViewNoColumnAutoresizing: Disable table column autoresizing.
	NSTableViewNoColumnAutoresizing NSTableViewColumnAutoresizingStyle = 0
	// NSTableViewReverseSequentialColumnAutoresizingStyle: Autoresize each table column sequentially, from the first auto-resizable column to the last auto-resizable column; proceed to the next column when the current column has reached its minimum or maximum size.
	NSTableViewReverseSequentialColumnAutoresizingStyle NSTableViewColumnAutoresizingStyle = 3
	// NSTableViewSequentialColumnAutoresizingStyle: Autoresize each table column sequentially, from the last auto-resizable column to the first auto-resizable column; proceed to the next column when the current column has reached its minimum or maximum size.
	NSTableViewSequentialColumnAutoresizingStyle NSTableViewColumnAutoresizingStyle = 2
	// NSTableViewUniformColumnAutoresizingStyle: Autoresize all columns by distributing space equally, simultaneously.
	NSTableViewUniformColumnAutoresizingStyle NSTableViewColumnAutoresizingStyle = 1
)


func (e NSTableViewColumnAutoresizingStyle) String() string {
	switch e {
	case NSTableViewFirstColumnOnlyAutoresizingStyle:
		return "NSTableViewFirstColumnOnlyAutoresizingStyle"
	case NSTableViewLastColumnOnlyAutoresizingStyle:
		return "NSTableViewLastColumnOnlyAutoresizingStyle"
	case NSTableViewNoColumnAutoresizing:
		return "NSTableViewNoColumnAutoresizing"
	case NSTableViewReverseSequentialColumnAutoresizingStyle:
		return "NSTableViewReverseSequentialColumnAutoresizingStyle"
	case NSTableViewSequentialColumnAutoresizingStyle:
		return "NSTableViewSequentialColumnAutoresizingStyle"
	case NSTableViewUniformColumnAutoresizingStyle:
		return "NSTableViewUniformColumnAutoresizingStyle"
	default:
		return fmt.Sprintf("NSTableViewColumnAutoresizingStyle(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSTableView/DraggingDestinationFeedbackStyle-swift.enum
type NSTableViewDraggingDestinationFeedbackStyle int

const (
	// NSTableViewDraggingDestinationFeedbackStyleGap: Provides a gap insertion when dragging over the table.
	NSTableViewDraggingDestinationFeedbackStyleGap NSTableViewDraggingDestinationFeedbackStyle = 2
	// NSTableViewDraggingDestinationFeedbackStyleNone: Provides no feedback when the user drags over the table view.
	NSTableViewDraggingDestinationFeedbackStyleNone NSTableViewDraggingDestinationFeedbackStyle = -1
	// NSTableViewDraggingDestinationFeedbackStyleRegular: Draws a solid round-rect background on drop target rows, and an insertion marker between rows.
	NSTableViewDraggingDestinationFeedbackStyleRegular NSTableViewDraggingDestinationFeedbackStyle = 0
	// NSTableViewDraggingDestinationFeedbackStyleSourceList: Draws an outline on drop target rows, and an insertion marker between rows.
	NSTableViewDraggingDestinationFeedbackStyleSourceList NSTableViewDraggingDestinationFeedbackStyle = 1
)


func (e NSTableViewDraggingDestinationFeedbackStyle) String() string {
	switch e {
	case NSTableViewDraggingDestinationFeedbackStyleGap:
		return "NSTableViewDraggingDestinationFeedbackStyleGap"
	case NSTableViewDraggingDestinationFeedbackStyleNone:
		return "NSTableViewDraggingDestinationFeedbackStyleNone"
	case NSTableViewDraggingDestinationFeedbackStyleRegular:
		return "NSTableViewDraggingDestinationFeedbackStyleRegular"
	case NSTableViewDraggingDestinationFeedbackStyleSourceList:
		return "NSTableViewDraggingDestinationFeedbackStyleSourceList"
	default:
		return fmt.Sprintf("NSTableViewDraggingDestinationFeedbackStyle(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSTableView/DropOperation
type NSTableViewDropOperation int

const (
	// NSTableViewDropAbove: Specifies that the drop should occur above the specified row.
	NSTableViewDropAbove NSTableViewDropOperation = 1
	// NSTableViewDropOn: Specifies that the drop should occur on the specified row.
	NSTableViewDropOn NSTableViewDropOperation = 0
)


func (e NSTableViewDropOperation) String() string {
	switch e {
	case NSTableViewDropAbove:
		return "NSTableViewDropAbove"
	case NSTableViewDropOn:
		return "NSTableViewDropOn"
	default:
		return fmt.Sprintf("NSTableViewDropOperation(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSTableView/GridLineStyle
type NSTableViewGridLineStyle int

const (
	// NSTableViewDashedHorizontalGridLineMask: Specifies that the horizontal grid lines should be drawn dashed.
	NSTableViewDashedHorizontalGridLineMask NSTableViewGridLineStyle = 8
	// NSTableViewSolidHorizontalGridLineMask: Specifies that horizontal grid lines should be displayed.
	NSTableViewSolidHorizontalGridLineMask NSTableViewGridLineStyle = 2
	// NSTableViewSolidVerticalGridLineMask: Specifies that vertical grid lines should be displayed.
	NSTableViewSolidVerticalGridLineMask NSTableViewGridLineStyle = 1
)


func (e NSTableViewGridLineStyle) String() string {
	switch e {
	case NSTableViewDashedHorizontalGridLineMask:
		return "NSTableViewDashedHorizontalGridLineMask"
	case NSTableViewSolidHorizontalGridLineMask:
		return "NSTableViewSolidHorizontalGridLineMask"
	case NSTableViewSolidVerticalGridLineMask:
		return "NSTableViewSolidVerticalGridLineMask"
	default:
		return fmt.Sprintf("NSTableViewGridLineStyle(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSTableViewRowAction/Style-swift.enum
type NSTableViewRowActionStyle int

const (
	// NSTableViewRowActionStyleDestructive: Apply a style that indicates that the action might change or delete data.
	NSTableViewRowActionStyleDestructive NSTableViewRowActionStyle = 1
	// NSTableViewRowActionStyleRegular: Apply the default style to the button.
	NSTableViewRowActionStyleRegular NSTableViewRowActionStyle = 0
)


func (e NSTableViewRowActionStyle) String() string {
	switch e {
	case NSTableViewRowActionStyleDestructive:
		return "NSTableViewRowActionStyleDestructive"
	case NSTableViewRowActionStyleRegular:
		return "NSTableViewRowActionStyleRegular"
	default:
		return fmt.Sprintf("NSTableViewRowActionStyle(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSTableView/RowSizeStyle-swift.enum
type NSTableViewRowSizeStyle int

const (
	// NSTableViewRowSizeStyleCustom: The table will use the  or invoke the delegate method , if implemented.
	NSTableViewRowSizeStyleCustom NSTableViewRowSizeStyle = 0
	// NSTableViewRowSizeStyleDefault: The table will use the system default layout size: small, medium or large.
	NSTableViewRowSizeStyleDefault NSTableViewRowSizeStyle = -1
	// NSTableViewRowSizeStyleLarge: The table will use a row height specified for a large table.
	NSTableViewRowSizeStyleLarge NSTableViewRowSizeStyle = 3
	// NSTableViewRowSizeStyleMedium: The table will use a row height specified for a medium table.
	NSTableViewRowSizeStyleMedium NSTableViewRowSizeStyle = 2
	// NSTableViewRowSizeStyleSmall: The table will use a row height specified for a small table.
	NSTableViewRowSizeStyleSmall NSTableViewRowSizeStyle = 1
)


func (e NSTableViewRowSizeStyle) String() string {
	switch e {
	case NSTableViewRowSizeStyleCustom:
		return "NSTableViewRowSizeStyleCustom"
	case NSTableViewRowSizeStyleDefault:
		return "NSTableViewRowSizeStyleDefault"
	case NSTableViewRowSizeStyleLarge:
		return "NSTableViewRowSizeStyleLarge"
	case NSTableViewRowSizeStyleMedium:
		return "NSTableViewRowSizeStyleMedium"
	case NSTableViewRowSizeStyleSmall:
		return "NSTableViewRowSizeStyleSmall"
	default:
		return fmt.Sprintf("NSTableViewRowSizeStyle(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSTableView/SelectionHighlightStyle-swift.enum
type NSTableViewSelectionHighlightStyle int

const (
	NSTableViewSelectionHighlightStyleNone NSTableViewSelectionHighlightStyle = -1
	NSTableViewSelectionHighlightStyleRegular NSTableViewSelectionHighlightStyle = 0
	NSTableViewSelectionHighlightStyleSourceList NSTableViewSelectionHighlightStyle = 1
)


func (e NSTableViewSelectionHighlightStyle) String() string {
	switch e {
	case NSTableViewSelectionHighlightStyleNone:
		return "NSTableViewSelectionHighlightStyleNone"
	case NSTableViewSelectionHighlightStyleRegular:
		return "NSTableViewSelectionHighlightStyleRegular"
	case NSTableViewSelectionHighlightStyleSourceList:
		return "NSTableViewSelectionHighlightStyleSourceList"
	default:
		return fmt.Sprintf("NSTableViewSelectionHighlightStyle(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSTableView/Style-swift.enum
type NSTableViewStyle int

const (
	// NSTableViewStyleAutomatic: The system resolves the table view style based on the table view hierarchy.
	NSTableViewStyleAutomatic NSTableViewStyle = 0
	// NSTableViewStyleFullWidth: The table view style resolves to a full-width style.
	NSTableViewStyleFullWidth NSTableViewStyle = 1
	// NSTableViewStyleInset: The table view style resolves to an inset style.
	NSTableViewStyleInset NSTableViewStyle = 2
	// NSTableViewStylePlain: The table view style resolves to a plain style.
	NSTableViewStylePlain NSTableViewStyle = 4
	// NSTableViewStyleSourceList: The table view style resolves to a source-list style.
	NSTableViewStyleSourceList NSTableViewStyle = 3
)


func (e NSTableViewStyle) String() string {
	switch e {
	case NSTableViewStyleAutomatic:
		return "NSTableViewStyleAutomatic"
	case NSTableViewStyleFullWidth:
		return "NSTableViewStyleFullWidth"
	case NSTableViewStyleInset:
		return "NSTableViewStyleInset"
	case NSTableViewStylePlain:
		return "NSTableViewStylePlain"
	case NSTableViewStyleSourceList:
		return "NSTableViewStyleSourceList"
	default:
		return fmt.Sprintf("NSTableViewStyle(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSTextAlignment
type NSTextAlignment int

const (
	// NSTextAlignmentCenter: Text is center-aligned.
	NSTextAlignmentCenter NSTextAlignment = 1
	// NSTextAlignmentJustified: Text is justified.
	NSTextAlignmentJustified NSTextAlignment = 3
	// NSTextAlignmentLeft: Text is left-aligned.
	NSTextAlignmentLeft NSTextAlignment = 0
	// NSTextAlignmentNatural: Text uses the default alignment for the current localization of the app.
	NSTextAlignmentNatural NSTextAlignment = 4
	// NSTextAlignmentRight: Text is right-aligned.
	NSTextAlignmentRight NSTextAlignment = 2
)


func (e NSTextAlignment) String() string {
	switch e {
	case NSTextAlignmentCenter:
		return "NSTextAlignmentCenter"
	case NSTextAlignmentJustified:
		return "NSTextAlignmentJustified"
	case NSTextAlignmentLeft:
		return "NSTextAlignmentLeft"
	case NSTextAlignmentNatural:
		return "NSTextAlignmentNatural"
	case NSTextAlignmentRight:
		return "NSTextAlignmentRight"
	default:
		return fmt.Sprintf("NSTextAlignment(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSTextBlock/Dimension
type NSTextBlockDimension int

const (
	// NSTextBlockHeight: Height of the text block.
	NSTextBlockHeight NSTextBlockDimension = 4
	// NSTextBlockMaximumHeight: Maximum height of the text block.
	NSTextBlockMaximumHeight NSTextBlockDimension = 6
	// NSTextBlockMaximumWidth: Maximum width of the text block.
	NSTextBlockMaximumWidth NSTextBlockDimension = 2
	// NSTextBlockMinimumHeight: Minimum height of the text block.
	NSTextBlockMinimumHeight NSTextBlockDimension = 5
	// NSTextBlockMinimumWidth: Minimum width of the text block.
	NSTextBlockMinimumWidth NSTextBlockDimension = 1
	// NSTextBlockWidth: Width of the text block.
	NSTextBlockWidth NSTextBlockDimension = 0
)


func (e NSTextBlockDimension) String() string {
	switch e {
	case NSTextBlockHeight:
		return "NSTextBlockHeight"
	case NSTextBlockMaximumHeight:
		return "NSTextBlockMaximumHeight"
	case NSTextBlockMaximumWidth:
		return "NSTextBlockMaximumWidth"
	case NSTextBlockMinimumHeight:
		return "NSTextBlockMinimumHeight"
	case NSTextBlockMinimumWidth:
		return "NSTextBlockMinimumWidth"
	case NSTextBlockWidth:
		return "NSTextBlockWidth"
	default:
		return fmt.Sprintf("NSTextBlockDimension(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSTextBlock/Layer
type NSTextBlockLayer int

const (
	// NSTextBlockBorder: The border of the text block.
	NSTextBlockBorder NSTextBlockLayer = 0
	// NSTextBlockMargin: Margin of the text block: space surrounding the border.
	NSTextBlockMargin NSTextBlockLayer = 1
	// NSTextBlockPadding: Padding of the text block: space surrounding the content area extending to the border.
	NSTextBlockPadding NSTextBlockLayer = -1
)


func (e NSTextBlockLayer) String() string {
	switch e {
	case NSTextBlockBorder:
		return "NSTextBlockBorder"
	case NSTextBlockMargin:
		return "NSTextBlockMargin"
	case NSTextBlockPadding:
		return "NSTextBlockPadding"
	default:
		return fmt.Sprintf("NSTextBlockLayer(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSTextBlock/ValueType
type NSTextBlockValueType int

const (
	// NSTextBlockAbsoluteValueType: Absolute value in points.
	NSTextBlockAbsoluteValueType NSTextBlockValueType = 0
	// NSTextBlockPercentageValueType: Percentage value (out of 100).
	NSTextBlockPercentageValueType NSTextBlockValueType = 1
)


func (e NSTextBlockValueType) String() string {
	switch e {
	case NSTextBlockAbsoluteValueType:
		return "NSTextBlockAbsoluteValueType"
	case NSTextBlockPercentageValueType:
		return "NSTextBlockPercentageValueType"
	default:
		return fmt.Sprintf("NSTextBlockValueType(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSTextBlock/VerticalAlignment-swift.enum
type NSTextBlockVerticalAlignment int

const (
	// NSTextBlockBaselineAlignment: Aligns adjacent blocks at the baseline of the first line of text in the block.
	NSTextBlockBaselineAlignment NSTextBlockVerticalAlignment = 3
	// NSTextBlockBottomAlignment: Aligns adjacent blocks at their bottom.
	NSTextBlockBottomAlignment NSTextBlockVerticalAlignment = 2
	// NSTextBlockMiddleAlignment: Aligns adjacent blocks at their middle.
	NSTextBlockMiddleAlignment NSTextBlockVerticalAlignment = 1
	// NSTextBlockTopAlignment: Aligns adjacent blocks at their top.
	NSTextBlockTopAlignment NSTextBlockVerticalAlignment = 0
)


func (e NSTextBlockVerticalAlignment) String() string {
	switch e {
	case NSTextBlockBaselineAlignment:
		return "NSTextBlockBaselineAlignment"
	case NSTextBlockBottomAlignment:
		return "NSTextBlockBottomAlignment"
	case NSTextBlockMiddleAlignment:
		return "NSTextBlockMiddleAlignment"
	case NSTextBlockTopAlignment:
		return "NSTextBlockTopAlignment"
	default:
		return fmt.Sprintf("NSTextBlockVerticalAlignment(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSTextContentManager/EnumerationOptions
type NSTextContentManagerEnumerationOptions int

const (
	// NSTextContentManagerEnumerationOptionsReverse: Returns whether enumerations start from the end of the text element.
	NSTextContentManagerEnumerationOptionsReverse NSTextContentManagerEnumerationOptions = 1
)


func (e NSTextContentManagerEnumerationOptions) String() string {
	switch e {
	case NSTextContentManagerEnumerationOptionsReverse:
		return "NSTextContentManagerEnumerationOptionsReverse"
	default:
		return fmt.Sprintf("NSTextContentManagerEnumerationOptions(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSTextCursorAccessoryPlacement
type NSTextCursorAccessoryPlacement int

const (
	NSTextCursorAccessoryPlacementBackward NSTextCursorAccessoryPlacement = 1
	NSTextCursorAccessoryPlacementCenter NSTextCursorAccessoryPlacement = 4
	NSTextCursorAccessoryPlacementForward NSTextCursorAccessoryPlacement = 2
	NSTextCursorAccessoryPlacementInvisible NSTextCursorAccessoryPlacement = 3
	NSTextCursorAccessoryPlacementOffscreenBottom NSTextCursorAccessoryPlacement = 8
	NSTextCursorAccessoryPlacementOffscreenLeft NSTextCursorAccessoryPlacement = 5
	NSTextCursorAccessoryPlacementOffscreenRight NSTextCursorAccessoryPlacement = 7
	NSTextCursorAccessoryPlacementOffscreenTop NSTextCursorAccessoryPlacement = 6
	NSTextCursorAccessoryPlacementUnspecified NSTextCursorAccessoryPlacement = 0
)


func (e NSTextCursorAccessoryPlacement) String() string {
	switch e {
	case NSTextCursorAccessoryPlacementBackward:
		return "NSTextCursorAccessoryPlacementBackward"
	case NSTextCursorAccessoryPlacementCenter:
		return "NSTextCursorAccessoryPlacementCenter"
	case NSTextCursorAccessoryPlacementForward:
		return "NSTextCursorAccessoryPlacementForward"
	case NSTextCursorAccessoryPlacementInvisible:
		return "NSTextCursorAccessoryPlacementInvisible"
	case NSTextCursorAccessoryPlacementOffscreenBottom:
		return "NSTextCursorAccessoryPlacementOffscreenBottom"
	case NSTextCursorAccessoryPlacementOffscreenLeft:
		return "NSTextCursorAccessoryPlacementOffscreenLeft"
	case NSTextCursorAccessoryPlacementOffscreenRight:
		return "NSTextCursorAccessoryPlacementOffscreenRight"
	case NSTextCursorAccessoryPlacementOffscreenTop:
		return "NSTextCursorAccessoryPlacementOffscreenTop"
	case NSTextCursorAccessoryPlacementUnspecified:
		return "NSTextCursorAccessoryPlacementUnspecified"
	default:
		return fmt.Sprintf("NSTextCursorAccessoryPlacement(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSTextField/BezelStyle-swift.enum
type NSTextFieldBezelStyle int

const (
	// NSTextFieldRoundedBezel: A style that draws a bezel with rounded corners around a single-line text field.
	NSTextFieldRoundedBezel NSTextFieldBezelStyle = 1
	// NSTextFieldSquareBezel: A style that draws a bezel with square corners around a text field.
	NSTextFieldSquareBezel NSTextFieldBezelStyle = 0
)


func (e NSTextFieldBezelStyle) String() string {
	switch e {
	case NSTextFieldRoundedBezel:
		return "NSTextFieldRoundedBezel"
	case NSTextFieldSquareBezel:
		return "NSTextFieldSquareBezel"
	default:
		return fmt.Sprintf("NSTextFieldBezelStyle(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSTextFinder/Action
type NSTextFinderAction int

const (
	// NSTextFinderActionHideFindInterface: Hides the find bar interface.
	NSTextFinderActionHideFindInterface NSTextFinderAction = 11
	// NSTextFinderActionHideReplaceInterface: Displays the find bar interface including the replace functionality.
	NSTextFinderActionHideReplaceInterface NSTextFinderAction = 13
	// NSTextFinderActionNextMatch: The next match, if any, is displayed.
	NSTextFinderActionNextMatch NSTextFinderAction = 2
	// NSTextFinderActionPreviousMatch: The previous match, if any, is displayed.
	NSTextFinderActionPreviousMatch NSTextFinderAction = 3
	// NSTextFinderActionReplace: Replaces a single instance of the string.
	NSTextFinderActionReplace NSTextFinderAction = 5
	// NSTextFinderActionReplaceAll: All occurrences of the string are replaced.
	NSTextFinderActionReplaceAll NSTextFinderAction = 4
	// NSTextFinderActionReplaceAllInSelection: Replaces all occurrences of the string within the current selection.
	NSTextFinderActionReplaceAllInSelection NSTextFinderAction = 8
	// NSTextFinderActionReplaceAndFind: Replaces a single instance of the string and searches for the next match.
	NSTextFinderActionReplaceAndFind NSTextFinderAction = 6
	// NSTextFinderActionSelectAll: Selects all matching search strings.
	NSTextFinderActionSelectAll NSTextFinderAction = 9
	// NSTextFinderActionSelectAllInSelection: Selects all matching search strings within the current selection.
	NSTextFinderActionSelectAllInSelection NSTextFinderAction = 10
	// NSTextFinderActionSetSearchString: Sets the search string.
	NSTextFinderActionSetSearchString NSTextFinderAction = 7
	// NSTextFinderActionShowFindInterface: The find bar interface is displayed.
	NSTextFinderActionShowFindInterface NSTextFinderAction = 1
	// NSTextFinderActionShowReplaceInterface: Displays the find bar interface including the replace functionality.
	NSTextFinderActionShowReplaceInterface NSTextFinderAction = 12
)


func (e NSTextFinderAction) String() string {
	switch e {
	case NSTextFinderActionHideFindInterface:
		return "NSTextFinderActionHideFindInterface"
	case NSTextFinderActionHideReplaceInterface:
		return "NSTextFinderActionHideReplaceInterface"
	case NSTextFinderActionNextMatch:
		return "NSTextFinderActionNextMatch"
	case NSTextFinderActionPreviousMatch:
		return "NSTextFinderActionPreviousMatch"
	case NSTextFinderActionReplace:
		return "NSTextFinderActionReplace"
	case NSTextFinderActionReplaceAll:
		return "NSTextFinderActionReplaceAll"
	case NSTextFinderActionReplaceAllInSelection:
		return "NSTextFinderActionReplaceAllInSelection"
	case NSTextFinderActionReplaceAndFind:
		return "NSTextFinderActionReplaceAndFind"
	case NSTextFinderActionSelectAll:
		return "NSTextFinderActionSelectAll"
	case NSTextFinderActionSelectAllInSelection:
		return "NSTextFinderActionSelectAllInSelection"
	case NSTextFinderActionSetSearchString:
		return "NSTextFinderActionSetSearchString"
	case NSTextFinderActionShowFindInterface:
		return "NSTextFinderActionShowFindInterface"
	case NSTextFinderActionShowReplaceInterface:
		return "NSTextFinderActionShowReplaceInterface"
	default:
		return fmt.Sprintf("NSTextFinderAction(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSTextFinder/MatchingType
type NSTextFinderMatchingType int

const (
	// NSTextFinderMatchingTypeContains: The match contains the string.
	NSTextFinderMatchingTypeContains NSTextFinderMatchingType = 0
	// NSTextFinderMatchingTypeEndsWith: The match ends with the string.
	NSTextFinderMatchingTypeEndsWith NSTextFinderMatchingType = 3
	// NSTextFinderMatchingTypeFullWord: The match exactly matches the string.
	NSTextFinderMatchingTypeFullWord NSTextFinderMatchingType = 2
	// NSTextFinderMatchingTypeStartsWith: The match begins with the string.
	NSTextFinderMatchingTypeStartsWith NSTextFinderMatchingType = 1
)


func (e NSTextFinderMatchingType) String() string {
	switch e {
	case NSTextFinderMatchingTypeContains:
		return "NSTextFinderMatchingTypeContains"
	case NSTextFinderMatchingTypeEndsWith:
		return "NSTextFinderMatchingTypeEndsWith"
	case NSTextFinderMatchingTypeFullWord:
		return "NSTextFinderMatchingTypeFullWord"
	case NSTextFinderMatchingTypeStartsWith:
		return "NSTextFinderMatchingTypeStartsWith"
	default:
		return fmt.Sprintf("NSTextFinderMatchingType(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraitType
type NSTextInputTraitType int

const (
	NSTextInputTraitTypeDefault NSTextInputTraitType = 0
	NSTextInputTraitTypeNo NSTextInputTraitType = 1
	NSTextInputTraitTypeYes NSTextInputTraitType = 2
)


func (e NSTextInputTraitType) String() string {
	switch e {
	case NSTextInputTraitTypeDefault:
		return "NSTextInputTraitTypeDefault"
	case NSTextInputTraitTypeNo:
		return "NSTextInputTraitTypeNo"
	case NSTextInputTraitTypeYes:
		return "NSTextInputTraitTypeYes"
	default:
		return fmt.Sprintf("NSTextInputTraitType(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSTextInsertionIndicator/AutomaticModeOptions-swift.struct
type NSTextInsertionIndicatorAutomaticModeOptions int

const (
	// NSTextInsertionIndicatorAutomaticModeOptionsShowEffectsView: Specifies whether a trailing glow displays during dictation.
	NSTextInsertionIndicatorAutomaticModeOptionsShowEffectsView NSTextInsertionIndicatorAutomaticModeOptions = 1
	// NSTextInsertionIndicatorAutomaticModeOptionsShowWhileTracking: Specifies whether the insertion indicator shows during a tracking loop.
	NSTextInsertionIndicatorAutomaticModeOptionsShowWhileTracking NSTextInsertionIndicatorAutomaticModeOptions = 2
)


func (e NSTextInsertionIndicatorAutomaticModeOptions) String() string {
	switch e {
	case NSTextInsertionIndicatorAutomaticModeOptionsShowEffectsView:
		return "NSTextInsertionIndicatorAutomaticModeOptionsShowEffectsView"
	case NSTextInsertionIndicatorAutomaticModeOptionsShowWhileTracking:
		return "NSTextInsertionIndicatorAutomaticModeOptionsShowWhileTracking"
	default:
		return fmt.Sprintf("NSTextInsertionIndicatorAutomaticModeOptions(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSTextInsertionIndicator/DisplayMode-swift.enum
type NSTextInsertionIndicatorDisplayMode int

const (
	NSTextInsertionIndicatorDisplayModeAutomatic NSTextInsertionIndicatorDisplayMode = 0
	NSTextInsertionIndicatorDisplayModeHidden NSTextInsertionIndicatorDisplayMode = 1
	NSTextInsertionIndicatorDisplayModeVisible NSTextInsertionIndicatorDisplayMode = 2
)


func (e NSTextInsertionIndicatorDisplayMode) String() string {
	switch e {
	case NSTextInsertionIndicatorDisplayModeAutomatic:
		return "NSTextInsertionIndicatorDisplayModeAutomatic"
	case NSTextInsertionIndicatorDisplayModeHidden:
		return "NSTextInsertionIndicatorDisplayModeHidden"
	case NSTextInsertionIndicatorDisplayModeVisible:
		return "NSTextInsertionIndicatorDisplayModeVisible"
	default:
		return fmt.Sprintf("NSTextInsertionIndicatorDisplayMode(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/EnumerationOptions
type NSTextLayoutFragmentEnumerationOptions int

const (
	// NSTextLayoutFragmentEnumerationOptionsEnsuresExtraLineFragment: Synthesize the extra line fragment when necessary.
	NSTextLayoutFragmentEnumerationOptionsEnsuresExtraLineFragment NSTextLayoutFragmentEnumerationOptions = 8
	// NSTextLayoutFragmentEnumerationOptionsEnsuresLayout: When enumerating, tell the layout fragments to layout their contents.
	NSTextLayoutFragmentEnumerationOptionsEnsuresLayout NSTextLayoutFragmentEnumerationOptions = 4
	// NSTextLayoutFragmentEnumerationOptionsEstimatesSize: When enumerating, tell the layout fragments to estimate their size.
	NSTextLayoutFragmentEnumerationOptionsEstimatesSize NSTextLayoutFragmentEnumerationOptions = 2
	// NSTextLayoutFragmentEnumerationOptionsReverse: Causes the enumeration to start from the last element.
	NSTextLayoutFragmentEnumerationOptionsReverse NSTextLayoutFragmentEnumerationOptions = 1
)


func (e NSTextLayoutFragmentEnumerationOptions) String() string {
	switch e {
	case NSTextLayoutFragmentEnumerationOptionsEnsuresExtraLineFragment:
		return "NSTextLayoutFragmentEnumerationOptionsEnsuresExtraLineFragment"
	case NSTextLayoutFragmentEnumerationOptionsEnsuresLayout:
		return "NSTextLayoutFragmentEnumerationOptionsEnsuresLayout"
	case NSTextLayoutFragmentEnumerationOptionsEstimatesSize:
		return "NSTextLayoutFragmentEnumerationOptionsEstimatesSize"
	case NSTextLayoutFragmentEnumerationOptionsReverse:
		return "NSTextLayoutFragmentEnumerationOptionsReverse"
	default:
		return fmt.Sprintf("NSTextLayoutFragmentEnumerationOptions(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/State-swift.enum
type NSTextLayoutFragmentState int

const (
	// NSTextLayoutFragmentStateCalculatedUsageBounds: The layout fragment measurements are available without text line fragments.
	NSTextLayoutFragmentStateCalculatedUsageBounds NSTextLayoutFragmentState = 2
	// NSTextLayoutFragmentStateEstimatedUsageBounds: The text layout manager hasn’t performed a full layout yet for the region covered by this layout fragment and is returning an estimated bounds.
	NSTextLayoutFragmentStateEstimatedUsageBounds NSTextLayoutFragmentState = 1
	// NSTextLayoutFragmentStateLayoutAvailable: Measurements for the text line fragments and layout fragment are available.
	NSTextLayoutFragmentStateLayoutAvailable NSTextLayoutFragmentState = 3
	// NSTextLayoutFragmentStateNone: No layout information is available.
	NSTextLayoutFragmentStateNone NSTextLayoutFragmentState = 0
)


func (e NSTextLayoutFragmentState) String() string {
	switch e {
	case NSTextLayoutFragmentStateCalculatedUsageBounds:
		return "NSTextLayoutFragmentStateCalculatedUsageBounds"
	case NSTextLayoutFragmentStateEstimatedUsageBounds:
		return "NSTextLayoutFragmentStateEstimatedUsageBounds"
	case NSTextLayoutFragmentStateLayoutAvailable:
		return "NSTextLayoutFragmentStateLayoutAvailable"
	case NSTextLayoutFragmentStateNone:
		return "NSTextLayoutFragmentStateNone"
	default:
		return fmt.Sprintf("NSTextLayoutFragmentState(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/SegmentOptions
type NSTextLayoutManagerSegmentOptions int

const (
	// NSTextLayoutManagerSegmentOptionsHeadSegmentExtended: Returns the value that causes the framework to extend the segment to the tail edge.
	NSTextLayoutManagerSegmentOptionsHeadSegmentExtended NSTextLayoutManagerSegmentOptions = 4
	// NSTextLayoutManagerSegmentOptionsMiddleFragmentsExcluded: Returns the value that causes the framework to enumerate segments in only the first and last line fragments.
	NSTextLayoutManagerSegmentOptionsMiddleFragmentsExcluded NSTextLayoutManagerSegmentOptions = 2
	// NSTextLayoutManagerSegmentOptionsRangeNotRequired: Returns the value that causes the framework enumerate text segment rectangles, but avoids preparing a range object.
	NSTextLayoutManagerSegmentOptionsRangeNotRequired NSTextLayoutManagerSegmentOptions = 1
	// NSTextLayoutManagerSegmentOptionsTailSegmentExtended: Returns the value that causes the framework to extend the segment to the tail edge.
	NSTextLayoutManagerSegmentOptionsTailSegmentExtended NSTextLayoutManagerSegmentOptions = 8
	// NSTextLayoutManagerSegmentOptionsUpstreamAffinity: Returns the value that causes the framework to the place the segment based on the upstream affinity for an empty range.
	NSTextLayoutManagerSegmentOptionsUpstreamAffinity NSTextLayoutManagerSegmentOptions = 16
)


func (e NSTextLayoutManagerSegmentOptions) String() string {
	switch e {
	case NSTextLayoutManagerSegmentOptionsHeadSegmentExtended:
		return "NSTextLayoutManagerSegmentOptionsHeadSegmentExtended"
	case NSTextLayoutManagerSegmentOptionsMiddleFragmentsExcluded:
		return "NSTextLayoutManagerSegmentOptionsMiddleFragmentsExcluded"
	case NSTextLayoutManagerSegmentOptionsRangeNotRequired:
		return "NSTextLayoutManagerSegmentOptionsRangeNotRequired"
	case NSTextLayoutManagerSegmentOptionsTailSegmentExtended:
		return "NSTextLayoutManagerSegmentOptionsTailSegmentExtended"
	case NSTextLayoutManagerSegmentOptionsUpstreamAffinity:
		return "NSTextLayoutManagerSegmentOptionsUpstreamAffinity"
	default:
		return fmt.Sprintf("NSTextLayoutManagerSegmentOptions(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/SegmentType
type NSTextLayoutManagerSegmentType int

const (
	// NSTextLayoutManagerSegmentTypeHighlight: The segment behavior suitable for highlighting.
	NSTextLayoutManagerSegmentTypeHighlight NSTextLayoutManagerSegmentType = 2
	// NSTextLayoutManagerSegmentTypeSelection: The segment behavior suitable for selection rendering.
	NSTextLayoutManagerSegmentTypeSelection NSTextLayoutManagerSegmentType = 1
	// NSTextLayoutManagerSegmentTypeStandard: The standard segment, matching the typographic bounds of the range.
	NSTextLayoutManagerSegmentTypeStandard NSTextLayoutManagerSegmentType = 0
)


func (e NSTextLayoutManagerSegmentType) String() string {
	switch e {
	case NSTextLayoutManagerSegmentTypeHighlight:
		return "NSTextLayoutManagerSegmentTypeHighlight"
	case NSTextLayoutManagerSegmentTypeSelection:
		return "NSTextLayoutManagerSegmentTypeSelection"
	case NSTextLayoutManagerSegmentTypeStandard:
		return "NSTextLayoutManagerSegmentTypeStandard"
	default:
		return fmt.Sprintf("NSTextLayoutManagerSegmentType(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/TextLayoutOrientation
type NSTextLayoutOrientation int

const (
	// NSTextLayoutOrientationHorizontal: Lines render horizontally, each line following the previous from top to bottom.
	NSTextLayoutOrientationHorizontal NSTextLayoutOrientation = 0
	// NSTextLayoutOrientationVertical: Lines render vertically, each line following the previous from right to left.
	NSTextLayoutOrientationVertical NSTextLayoutOrientation = 1
)


func (e NSTextLayoutOrientation) String() string {
	switch e {
	case NSTextLayoutOrientationHorizontal:
		return "NSTextLayoutOrientationHorizontal"
	case NSTextLayoutOrientationVertical:
		return "NSTextLayoutOrientationVertical"
	default:
		return fmt.Sprintf("NSTextLayoutOrientation(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSTextList/Options
type NSTextListOptions int

const (
	// NSTextListPrependEnclosingMarker: Specifies that a nested list should include the marker for its enclosing superlist before its own marker.
	NSTextListPrependEnclosingMarker NSTextListOptions = 1
)


func (e NSTextListOptions) String() string {
	switch e {
	case NSTextListPrependEnclosingMarker:
		return "NSTextListPrependEnclosingMarker"
	default:
		return fmt.Sprintf("NSTextListOptions(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSTextMovement
type NSTextMovement int

const (
	NSTextMovementBacktab NSTextMovement = 18
	NSTextMovementCancel NSTextMovement = 23
	NSTextMovementDown NSTextMovement = 22
	NSTextMovementLeft NSTextMovement = 19
	NSTextMovementOther NSTextMovement = 0
	NSTextMovementReturn NSTextMovement = 16
	NSTextMovementRight NSTextMovement = 20
	NSTextMovementTab NSTextMovement = 17
	NSTextMovementUp NSTextMovement = 21
)


func (e NSTextMovement) String() string {
	switch e {
	case NSTextMovementBacktab:
		return "NSTextMovementBacktab"
	case NSTextMovementCancel:
		return "NSTextMovementCancel"
	case NSTextMovementDown:
		return "NSTextMovementDown"
	case NSTextMovementLeft:
		return "NSTextMovementLeft"
	case NSTextMovementOther:
		return "NSTextMovementOther"
	case NSTextMovementReturn:
		return "NSTextMovementReturn"
	case NSTextMovementRight:
		return "NSTextMovementRight"
	case NSTextMovementTab:
		return "NSTextMovementTab"
	case NSTextMovementUp:
		return "NSTextMovementUp"
	default:
		return fmt.Sprintf("NSTextMovement(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSTextScalingType
type NSTextScalingType int

const (
	// NSTextScalingStandard: Font sizes throughout the document appear visually similar to how they would render in macOS and non-Apple platforms.
	NSTextScalingStandard NSTextScalingType = 0
	// NSTextScalingiOS: Font sizes throughout the document appear visually similar to how they would render in iOS.
	NSTextScalingiOS NSTextScalingType = 1
)


func (e NSTextScalingType) String() string {
	switch e {
	case NSTextScalingStandard:
		return "NSTextScalingStandard"
	case NSTextScalingiOS:
		return "NSTextScalingiOS"
	default:
		return fmt.Sprintf("NSTextScalingType(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSTextSelection/Affinity-swift.enum
type NSTextSelectionAffinity int

const (
	// NSTextSelectionAffinityDownstream: The value that defines the visual location of the text cursor between the head of line that contains the selection location.
	NSTextSelectionAffinityDownstream NSTextSelectionAffinity = 1
	// NSTextSelectionAffinityUpstream: The value that defines the visual location of the text cursor to the tail of the previous line.
	NSTextSelectionAffinityUpstream NSTextSelectionAffinity = 0
)


func (e NSTextSelectionAffinity) String() string {
	switch e {
	case NSTextSelectionAffinityDownstream:
		return "NSTextSelectionAffinityDownstream"
	case NSTextSelectionAffinityUpstream:
		return "NSTextSelectionAffinityUpstream"
	default:
		return fmt.Sprintf("NSTextSelectionAffinity(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSTextSelection/Granularity-swift.enum
type NSTextSelectionGranularity int

const (
	// NSTextSelectionGranularityCharacter: A value that represents selection by character.
	NSTextSelectionGranularityCharacter NSTextSelectionGranularity = 0
	// NSTextSelectionGranularityLine: A value that represents selection by line.
	NSTextSelectionGranularityLine NSTextSelectionGranularity = 3
	// NSTextSelectionGranularityParagraph: A value that represents selection by paragraph.
	NSTextSelectionGranularityParagraph NSTextSelectionGranularity = 2
	// NSTextSelectionGranularitySentence: A value that represents selection by sentence.
	NSTextSelectionGranularitySentence NSTextSelectionGranularity = 4
	// NSTextSelectionGranularityWord: A value that represents selection by word.
	NSTextSelectionGranularityWord NSTextSelectionGranularity = 1
)


func (e NSTextSelectionGranularity) String() string {
	switch e {
	case NSTextSelectionGranularityCharacter:
		return "NSTextSelectionGranularityCharacter"
	case NSTextSelectionGranularityLine:
		return "NSTextSelectionGranularityLine"
	case NSTextSelectionGranularityParagraph:
		return "NSTextSelectionGranularityParagraph"
	case NSTextSelectionGranularitySentence:
		return "NSTextSelectionGranularitySentence"
	case NSTextSelectionGranularityWord:
		return "NSTextSelectionGranularityWord"
	default:
		return fmt.Sprintf("NSTextSelectionGranularity(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/Destination
type NSTextSelectionNavigationDestination int

const (
	// NSTextSelectionNavigationDestinationCharacter: The selection moves to the next extended grapheme cluster boundary.
	NSTextSelectionNavigationDestinationCharacter NSTextSelectionNavigationDestination = 0
	// NSTextSelectionNavigationDestinationContainer: The selection moves to the next container or page boundary after boundary of the current container, ignoring the end of line elastic characters.
	NSTextSelectionNavigationDestinationContainer NSTextSelectionNavigationDestination = 5
	// NSTextSelectionNavigationDestinationDocument: The selection moves to the document boundary.
	NSTextSelectionNavigationDestinationDocument NSTextSelectionNavigationDestination = 6
	// NSTextSelectionNavigationDestinationLine: The selection moves to the next line boundary.
	NSTextSelectionNavigationDestinationLine NSTextSelectionNavigationDestination = 2
	// NSTextSelectionNavigationDestinationParagraph: The selection moves to the next paragraph boundary, ignoring the end of line elastic characters and paragraph separators.
	NSTextSelectionNavigationDestinationParagraph NSTextSelectionNavigationDestination = 4
	// NSTextSelectionNavigationDestinationSentence: The selection moves to the next sentence boundary, ignoring punctuation, whitespace, and format characters preceding the next sentence.
	NSTextSelectionNavigationDestinationSentence NSTextSelectionNavigationDestination = 3
	// NSTextSelectionNavigationDestinationWord: The selection moves to the next word boundary ignoring punctuation, whitespace, and format characters preceding the next word.
	NSTextSelectionNavigationDestinationWord NSTextSelectionNavigationDestination = 1
)


func (e NSTextSelectionNavigationDestination) String() string {
	switch e {
	case NSTextSelectionNavigationDestinationCharacter:
		return "NSTextSelectionNavigationDestinationCharacter"
	case NSTextSelectionNavigationDestinationContainer:
		return "NSTextSelectionNavigationDestinationContainer"
	case NSTextSelectionNavigationDestinationDocument:
		return "NSTextSelectionNavigationDestinationDocument"
	case NSTextSelectionNavigationDestinationLine:
		return "NSTextSelectionNavigationDestinationLine"
	case NSTextSelectionNavigationDestinationParagraph:
		return "NSTextSelectionNavigationDestinationParagraph"
	case NSTextSelectionNavigationDestinationSentence:
		return "NSTextSelectionNavigationDestinationSentence"
	case NSTextSelectionNavigationDestinationWord:
		return "NSTextSelectionNavigationDestinationWord"
	default:
		return fmt.Sprintf("NSTextSelectionNavigationDestination(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/Direction
type NSTextSelectionNavigationDirection int

const (
	// NSTextSelectionNavigationDirectionBackward: The value that represents a backward selection based on the flow of text stored in the document.
	NSTextSelectionNavigationDirectionBackward NSTextSelectionNavigationDirection = 1
	// NSTextSelectionNavigationDirectionDown: The value that represents a selection in the down direction, below the current line.
	NSTextSelectionNavigationDirectionDown NSTextSelectionNavigationDirection = 5
	// NSTextSelectionNavigationDirectionForward: The value that represents a logical forward selection based on the flow of text stored in the document.
	NSTextSelectionNavigationDirectionForward NSTextSelectionNavigationDirection = 0
	// NSTextSelectionNavigationDirectionLeft: The value that represents a selection in the left direction along the current line.
	NSTextSelectionNavigationDirectionLeft NSTextSelectionNavigationDirection = 3
	// NSTextSelectionNavigationDirectionRight: The value that represents a selection in the right direction along the current line.
	NSTextSelectionNavigationDirectionRight NSTextSelectionNavigationDirection = 2
	// NSTextSelectionNavigationDirectionUp: The value that represents a selection in the up direction, above the current line.
	NSTextSelectionNavigationDirectionUp NSTextSelectionNavigationDirection = 4
)


func (e NSTextSelectionNavigationDirection) String() string {
	switch e {
	case NSTextSelectionNavigationDirectionBackward:
		return "NSTextSelectionNavigationDirectionBackward"
	case NSTextSelectionNavigationDirectionDown:
		return "NSTextSelectionNavigationDirectionDown"
	case NSTextSelectionNavigationDirectionForward:
		return "NSTextSelectionNavigationDirectionForward"
	case NSTextSelectionNavigationDirectionLeft:
		return "NSTextSelectionNavigationDirectionLeft"
	case NSTextSelectionNavigationDirectionRight:
		return "NSTextSelectionNavigationDirectionRight"
	case NSTextSelectionNavigationDirectionUp:
		return "NSTextSelectionNavigationDirectionUp"
	default:
		return fmt.Sprintf("NSTextSelectionNavigationDirection(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/LayoutOrientation
type NSTextSelectionNavigationLayoutOrientation int

const (
	// NSTextSelectionNavigationLayoutOrientationHorizontal: The value that defines horizontal layout orientation.
	NSTextSelectionNavigationLayoutOrientationHorizontal NSTextSelectionNavigationLayoutOrientation = 0
	// NSTextSelectionNavigationLayoutOrientationVertical: The value that defines vertical layout orientation.
	NSTextSelectionNavigationLayoutOrientationVertical NSTextSelectionNavigationLayoutOrientation = 1
)


func (e NSTextSelectionNavigationLayoutOrientation) String() string {
	switch e {
	case NSTextSelectionNavigationLayoutOrientationHorizontal:
		return "NSTextSelectionNavigationLayoutOrientationHorizontal"
	case NSTextSelectionNavigationLayoutOrientationVertical:
		return "NSTextSelectionNavigationLayoutOrientationVertical"
	default:
		return fmt.Sprintf("NSTextSelectionNavigationLayoutOrientation(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/Modifier
type NSTextSelectionNavigationModifier int

const (
	// NSTextSelectionNavigationModifierExtend: The value that indicates the framework extends the selection by not moving the initial location while in a drag selection.
	NSTextSelectionNavigationModifierExtend NSTextSelectionNavigationModifier = 1
	// NSTextSelectionNavigationModifierMultiple: The value that indicates the framework extends the selection visually inside the rectangular area defined by the anchor and dragged positions.
	NSTextSelectionNavigationModifierMultiple NSTextSelectionNavigationModifier = 4
	// NSTextSelectionNavigationModifierVisual: The value that indicates the framework extends the selection visually inside the rectangular area defined by the anchor and drag positions.
	NSTextSelectionNavigationModifierVisual NSTextSelectionNavigationModifier = 2
)


func (e NSTextSelectionNavigationModifier) String() string {
	switch e {
	case NSTextSelectionNavigationModifierExtend:
		return "NSTextSelectionNavigationModifierExtend"
	case NSTextSelectionNavigationModifierMultiple:
		return "NSTextSelectionNavigationModifierMultiple"
	case NSTextSelectionNavigationModifierVisual:
		return "NSTextSelectionNavigationModifierVisual"
	default:
		return fmt.Sprintf("NSTextSelectionNavigationModifier(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/WritingDirection
type NSTextSelectionNavigationWritingDirection int

const (
	// NSTextSelectionNavigationWritingDirectionLeftToRight: The value that defines the left to right writing direction.
	NSTextSelectionNavigationWritingDirectionLeftToRight NSTextSelectionNavigationWritingDirection = 0
	// NSTextSelectionNavigationWritingDirectionRightToLeft: The value that defines the right to left writing direction.
	NSTextSelectionNavigationWritingDirectionRightToLeft NSTextSelectionNavigationWritingDirection = 1
)


func (e NSTextSelectionNavigationWritingDirection) String() string {
	switch e {
	case NSTextSelectionNavigationWritingDirectionLeftToRight:
		return "NSTextSelectionNavigationWritingDirectionLeftToRight"
	case NSTextSelectionNavigationWritingDirectionRightToLeft:
		return "NSTextSelectionNavigationWritingDirectionRightToLeft"
	default:
		return fmt.Sprintf("NSTextSelectionNavigationWritingDirection(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSTextStorageEditActions
type NSTextStorageEditActions int

const (
	// NSTextStorageEditedAttributes: Attributes were added, removed, or changed.
	NSTextStorageEditedAttributes NSTextStorageEditActions = 1
	// NSTextStorageEditedCharacters: Characters were added, removed, or replaced.
	NSTextStorageEditedCharacters NSTextStorageEditActions = 2
)


func (e NSTextStorageEditActions) String() string {
	switch e {
	case NSTextStorageEditedAttributes:
		return "NSTextStorageEditedAttributes"
	case NSTextStorageEditedCharacters:
		return "NSTextStorageEditedCharacters"
	default:
		return fmt.Sprintf("NSTextStorageEditActions(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/TextTabType
type NSTextTabType int

const (
	// NSCenterTabStopType: A center-aligned tab stop.
	NSCenterTabStopType NSTextTabType = 2
	// NSDecimalTabStopType: A tab stop that aligns columns of numbers to each number’s decimal point.
	NSDecimalTabStopType NSTextTabType = 3
	// NSLeftTabStopType: A left-aligned tab stop.
	NSLeftTabStopType NSTextTabType = 0
	// NSRightTabStopType: A right-aligned tab stop.
	NSRightTabStopType NSTextTabType = 1
)


func (e NSTextTabType) String() string {
	switch e {
	case NSCenterTabStopType:
		return "NSCenterTabStopType"
	case NSDecimalTabStopType:
		return "NSDecimalTabStopType"
	case NSLeftTabStopType:
		return "NSLeftTabStopType"
	case NSRightTabStopType:
		return "NSRightTabStopType"
	default:
		return fmt.Sprintf("NSTextTabType(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSTextTable/LayoutAlgorithm-swift.enum
type NSTextTableLayoutAlgorithm int

const (
	// NSTextTableAutomaticLayoutAlgorithm: Specifies automatic layout algorithm
	NSTextTableAutomaticLayoutAlgorithm NSTextTableLayoutAlgorithm = 0
	// NSTextTableFixedLayoutAlgorithm: Specifies fixed layout algorithm
	NSTextTableFixedLayoutAlgorithm NSTextTableLayoutAlgorithm = 1
)


func (e NSTextTableLayoutAlgorithm) String() string {
	switch e {
	case NSTextTableAutomaticLayoutAlgorithm:
		return "NSTextTableAutomaticLayoutAlgorithm"
	case NSTextTableFixedLayoutAlgorithm:
		return "NSTextTableFixedLayoutAlgorithm"
	default:
		return fmt.Sprintf("NSTextTableLayoutAlgorithm(%d)", e)
	}
}




type NSTextWritingDirection uint

const (
	// Deprecated.
	NSTextWritingDirectionEmbedding NSTextWritingDirection = 0
	// Deprecated.
	NSTextWritingDirectionOverride NSTextWritingDirection = 2
)


func (e NSTextWritingDirection) String() string {
	switch e {
	case NSTextWritingDirectionEmbedding:
		return "NSTextWritingDirectionEmbedding"
	case NSTextWritingDirectionOverride:
		return "NSTextWritingDirectionOverride"
	default:
		return fmt.Sprintf("NSTextWritingDirection(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSSlider/TickMarkPosition-swift.enum
type NSTickMarkPosition int

const (
	// NSTickMarkPositionAbove: A constant indicating that tick marks are displayed above the slider.
	NSTickMarkPositionAbove NSTickMarkPosition = 1
	// NSTickMarkPositionBelow: A constant indicating that tick marks are displayed below the slider.
	NSTickMarkPositionBelow NSTickMarkPosition = 0
)


func (e NSTickMarkPosition) String() string {
	switch e {
	case NSTickMarkPositionAbove:
		return "NSTickMarkPositionAbove"
	case NSTickMarkPositionBelow:
		return "NSTickMarkPositionBelow"
	default:
		return fmt.Sprintf("NSTickMarkPosition(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSTintProminence
type NSTintProminence int

const (
	NSTintProminenceAutomatic NSTintProminence = 0
	NSTintProminenceNone NSTintProminence = 1
	NSTintProminencePrimary NSTintProminence = 2
	NSTintProminenceSecondary NSTintProminence = 3
)


func (e NSTintProminence) String() string {
	switch e {
	case NSTintProminenceAutomatic:
		return "NSTintProminenceAutomatic"
	case NSTintProminenceNone:
		return "NSTintProminenceNone"
	case NSTintProminencePrimary:
		return "NSTintProminencePrimary"
	case NSTintProminenceSecondary:
		return "NSTintProminenceSecondary"
	default:
		return fmt.Sprintf("NSTintProminence(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSBox/TitlePosition-swift.enum
type NSTitlePosition int

const (
	NSAboveBottom NSTitlePosition = 4
	NSAboveTop NSTitlePosition = 1
	NSAtBottom NSTitlePosition = 5
	NSAtTop NSTitlePosition = 2
	NSBelowBottom NSTitlePosition = 6
	NSBelowTop NSTitlePosition = 3
	NSNoTitle NSTitlePosition = 0
)


func (e NSTitlePosition) String() string {
	switch e {
	case NSAboveBottom:
		return "NSAboveBottom"
	case NSAboveTop:
		return "NSAboveTop"
	case NSAtBottom:
		return "NSAtBottom"
	case NSAtTop:
		return "NSAtTop"
	case NSBelowBottom:
		return "NSBelowBottom"
	case NSBelowTop:
		return "NSBelowTop"
	case NSNoTitle:
		return "NSNoTitle"
	default:
		return fmt.Sprintf("NSTitlePosition(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSTitlebarSeparatorStyle
type NSTitlebarSeparatorStyle int

const (
	// NSTitlebarSeparatorStyleAutomatic: A style indicating that the system determines the type of separator.
	NSTitlebarSeparatorStyleAutomatic NSTitlebarSeparatorStyle = 0
	// NSTitlebarSeparatorStyleLine: A style indicating that the title bar separator is a line.
	NSTitlebarSeparatorStyleLine NSTitlebarSeparatorStyle = 2
	// NSTitlebarSeparatorStyleNone: A style indicating that there’s no title bar separator.
	NSTitlebarSeparatorStyleNone NSTitlebarSeparatorStyle = 1
	// NSTitlebarSeparatorStyleShadow: A style indicating that the title bar separator is a shadow.
	NSTitlebarSeparatorStyleShadow NSTitlebarSeparatorStyle = 3
)


func (e NSTitlebarSeparatorStyle) String() string {
	switch e {
	case NSTitlebarSeparatorStyleAutomatic:
		return "NSTitlebarSeparatorStyleAutomatic"
	case NSTitlebarSeparatorStyleLine:
		return "NSTitlebarSeparatorStyleLine"
	case NSTitlebarSeparatorStyleNone:
		return "NSTitlebarSeparatorStyleNone"
	case NSTitlebarSeparatorStyleShadow:
		return "NSTitlebarSeparatorStyleShadow"
	default:
		return fmt.Sprintf("NSTitlebarSeparatorStyle(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSTokenField/TokenStyle-swift.enum
type NSTokenStyle int

const (
	NSTokenStyleDefault NSTokenStyle = 0
	NSTokenStyleNone NSTokenStyle = 1
	NSTokenStylePlainSquared NSTokenStyle = 4
	NSTokenStyleRounded NSTokenStyle = 2
	NSTokenStyleSquared NSTokenStyle = 3
)


func (e NSTokenStyle) String() string {
	switch e {
	case NSTokenStyleDefault:
		return "NSTokenStyleDefault"
	case NSTokenStyleNone:
		return "NSTokenStyleNone"
	case NSTokenStylePlainSquared:
		return "NSTokenStylePlainSquared"
	case NSTokenStyleRounded:
		return "NSTokenStyleRounded"
	case NSTokenStyleSquared:
		return "NSTokenStyleSquared"
	default:
		return fmt.Sprintf("NSTokenStyle(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSToolbar/DisplayMode-swift.enum
type NSToolbarDisplayMode int

const (
	// NSToolbarDisplayModeDefault: The default display mode.
	NSToolbarDisplayModeDefault NSToolbarDisplayMode = 0
	// NSToolbarDisplayModeIconAndLabel: The toolbar displays an icon and label for each item.
	NSToolbarDisplayModeIconAndLabel NSToolbarDisplayMode = 1
	// NSToolbarDisplayModeIconOnly: The toolbar displays only an icon for each item.
	NSToolbarDisplayModeIconOnly NSToolbarDisplayMode = 2
	// NSToolbarDisplayModeLabelOnly: The toolbar displays only a label for each item.
	NSToolbarDisplayModeLabelOnly NSToolbarDisplayMode = 3
)


func (e NSToolbarDisplayMode) String() string {
	switch e {
	case NSToolbarDisplayModeDefault:
		return "NSToolbarDisplayModeDefault"
	case NSToolbarDisplayModeIconAndLabel:
		return "NSToolbarDisplayModeIconAndLabel"
	case NSToolbarDisplayModeIconOnly:
		return "NSToolbarDisplayModeIconOnly"
	case NSToolbarDisplayModeLabelOnly:
		return "NSToolbarDisplayModeLabelOnly"
	default:
		return fmt.Sprintf("NSToolbarDisplayMode(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/ControlRepresentation-swift.enum
type NSToolbarItemGroupControlRepresentation int

const (
	NSToolbarItemGroupControlRepresentationAutomatic NSToolbarItemGroupControlRepresentation = 0
	NSToolbarItemGroupControlRepresentationCollapsed NSToolbarItemGroupControlRepresentation = 2
	NSToolbarItemGroupControlRepresentationExpanded NSToolbarItemGroupControlRepresentation = 1
)


func (e NSToolbarItemGroupControlRepresentation) String() string {
	switch e {
	case NSToolbarItemGroupControlRepresentationAutomatic:
		return "NSToolbarItemGroupControlRepresentationAutomatic"
	case NSToolbarItemGroupControlRepresentationCollapsed:
		return "NSToolbarItemGroupControlRepresentationCollapsed"
	case NSToolbarItemGroupControlRepresentationExpanded:
		return "NSToolbarItemGroupControlRepresentationExpanded"
	default:
		return fmt.Sprintf("NSToolbarItemGroupControlRepresentation(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/SelectionMode-swift.enum
type NSToolbarItemGroupSelectionMode int

const (
	// NSToolbarItemGroupSelectionModeMomentary: The system temporarily highlights the select group item when the user selects the item.
	NSToolbarItemGroupSelectionModeMomentary NSToolbarItemGroupSelectionMode = 2
	// NSToolbarItemGroupSelectionModeSelectAny: The system toggles a highlight on any item selected.
	NSToolbarItemGroupSelectionModeSelectAny NSToolbarItemGroupSelectionMode = 1
	// NSToolbarItemGroupSelectionModeSelectOne: The system displays a highlighted mode on the most recent item selected.
	NSToolbarItemGroupSelectionModeSelectOne NSToolbarItemGroupSelectionMode = 0
)


func (e NSToolbarItemGroupSelectionMode) String() string {
	switch e {
	case NSToolbarItemGroupSelectionModeMomentary:
		return "NSToolbarItemGroupSelectionModeMomentary"
	case NSToolbarItemGroupSelectionModeSelectAny:
		return "NSToolbarItemGroupSelectionModeSelectAny"
	case NSToolbarItemGroupSelectionModeSelectOne:
		return "NSToolbarItemGroupSelectionModeSelectOne"
	default:
		return fmt.Sprintf("NSToolbarItemGroupSelectionMode(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSToolbarItem/Style-swift.enum
type NSToolbarItemStyle int

const (
	NSToolbarItemStylePlain NSToolbarItemStyle = 0
	NSToolbarItemStyleProminent NSToolbarItemStyle = 1
)


func (e NSToolbarItemStyle) String() string {
	switch e {
	case NSToolbarItemStylePlain:
		return "NSToolbarItemStylePlain"
	case NSToolbarItemStyleProminent:
		return "NSToolbarItemStyleProminent"
	default:
		return fmt.Sprintf("NSToolbarItemStyle(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSToolbar/SizeMode-swift.enum
type NSToolbarSizeMode int

const (
	// NSToolbarSizeModeDefault: The toolbar uses the system-defined default size, which is .
	NSToolbarSizeModeDefault NSToolbarSizeMode = 0
	// NSToolbarSizeModeRegular: The toolbar uses regular-sized controls and 32 by 32 pixel icons.
	NSToolbarSizeModeRegular NSToolbarSizeMode = 1
	// NSToolbarSizeModeSmall: The toolbar uses small-sized controls and 24 by 24 pixel icons.
	NSToolbarSizeModeSmall NSToolbarSizeMode = 2
)


func (e NSToolbarSizeMode) String() string {
	switch e {
	case NSToolbarSizeModeDefault:
		return "NSToolbarSizeModeDefault"
	case NSToolbarSizeModeRegular:
		return "NSToolbarSizeModeRegular"
	case NSToolbarSizeModeSmall:
		return "NSToolbarSizeModeSmall"
	default:
		return fmt.Sprintf("NSToolbarSizeMode(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSTouch/Phase-swift.struct
type NSTouchPhase int

const (
	// NSTouchPhaseAny: Matches any phase of a touch.
	NSTouchPhaseAny NSTouchPhase = 0
	// NSTouchPhaseBegan: A finger touched the device.
	NSTouchPhaseBegan NSTouchPhase = 1
	// NSTouchPhaseCancelled: The system cancelled tracking for the touch, as when (for example) the window associated with the touch resigns key or is deactivated.
	NSTouchPhaseCancelled NSTouchPhase = 16
	// NSTouchPhaseEnded: A finger was lifted from the screen.
	NSTouchPhaseEnded NSTouchPhase = 8
	// NSTouchPhaseMoved: A finger moved on the device.
	NSTouchPhaseMoved NSTouchPhase = 2
	// NSTouchPhaseStationary: A finger is touching the device, but hasn’t moved since the previous event.
	NSTouchPhaseStationary NSTouchPhase = 4
	// NSTouchPhaseTouching: Matches the NSTouchPhaseBegan, NSTouchPhaseMoved, or NSTouchPhaseStationary phases of a touch.
	NSTouchPhaseTouching NSTouchPhase = 1
)


func (e NSTouchPhase) String() string {
	switch e {
	case NSTouchPhaseAny:
		return "NSTouchPhaseAny"
	case NSTouchPhaseBegan:
		return "NSTouchPhaseBegan"
	case NSTouchPhaseCancelled:
		return "NSTouchPhaseCancelled"
	case NSTouchPhaseEnded:
		return "NSTouchPhaseEnded"
	case NSTouchPhaseMoved:
		return "NSTouchPhaseMoved"
	case NSTouchPhaseStationary:
		return "NSTouchPhaseStationary"
	default:
		return fmt.Sprintf("NSTouchPhase(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSTouch/TouchType
type NSTouchType int

const (
	// NSTouchTypeDirect: A direct touch from a user’s finger on a screen.
	NSTouchTypeDirect NSTouchType = 0
	// NSTouchTypeIndirect: An indirect touch that is not on a screen, like a digitizer touch.
	NSTouchTypeIndirect NSTouchType = 1
)


func (e NSTouchType) String() string {
	switch e {
	case NSTouchTypeDirect:
		return "NSTouchTypeDirect"
	case NSTouchTypeIndirect:
		return "NSTouchTypeIndirect"
	default:
		return fmt.Sprintf("NSTouchType(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSTouch/TouchTypeMask
type NSTouchTypeMask int

const (
	// NSTouchTypeMaskDirect: A direct touch from a user’s finger on a screen.
	NSTouchTypeMaskDirect NSTouchTypeMask = 1
	// NSTouchTypeMaskIndirect: An indirect touch that is not on a screen, like a digitizer touch.
	NSTouchTypeMaskIndirect NSTouchTypeMask = 2
)


func (e NSTouchTypeMask) String() string {
	switch e {
	case NSTouchTypeMaskDirect:
		return "NSTouchTypeMaskDirect"
	case NSTouchTypeMaskIndirect:
		return "NSTouchTypeMaskIndirect"
	default:
		return fmt.Sprintf("NSTouchTypeMask(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSTrackingArea/Options-swift.struct
type NSTrackingAreaOptions int

const (
	// NSTrackingActiveAlways: The owner receives messages regardless of first-responder status, window status, or application status.
	NSTrackingActiveAlways NSTrackingAreaOptions = 128
	// NSTrackingActiveInActiveApp: The owner receives messages when the application is active.
	NSTrackingActiveInActiveApp NSTrackingAreaOptions = 64
	// NSTrackingActiveInKeyWindow: The owner receives messages when the view is in the key window.
	NSTrackingActiveInKeyWindow NSTrackingAreaOptions = 32
	// NSTrackingActiveWhenFirstResponder: The owner receives messages when the view is the first responder.
	NSTrackingActiveWhenFirstResponder NSTrackingAreaOptions = 16
	// NSTrackingAssumeInside: The first event is generated when the cursor leaves the tracking area, regardless if the cursor is inside the area when the [NSTrackingArea] is added to a view.
	NSTrackingAssumeInside NSTrackingAreaOptions = 256
	// NSTrackingCursorUpdate: A tracking option that receives events when the mouse cursor enters and exits the tracking area.
	NSTrackingCursorUpdate NSTrackingAreaOptions = 4
	// NSTrackingEnabledDuringMouseDrag: The owner receives NSMouseEntered events when the mouse cursor is dragged into the tracking area.
	NSTrackingEnabledDuringMouseDrag NSTrackingAreaOptions = 1024
	// NSTrackingInVisibleRect: Mouse tracking occurs only in the visible rectangle of the view—in other words, that region of the tracking rectangle that is unobscured.
	NSTrackingInVisibleRect NSTrackingAreaOptions = 512
	// NSTrackingMouseEnteredAndExited: The owner of the tracking area receives mouseEntered(with:) when the mouse cursor enters the area and mouseExited(with:) events when the mouse leaves the area.
	NSTrackingMouseEnteredAndExited NSTrackingAreaOptions = 1
	// NSTrackingMouseMoved: The owner of the tracking area receives mouseMoved(with:) messages while the mouse cursor is within the area.
	NSTrackingMouseMoved NSTrackingAreaOptions = 2
)


func (e NSTrackingAreaOptions) String() string {
	switch e {
	case NSTrackingActiveAlways:
		return "NSTrackingActiveAlways"
	case NSTrackingActiveInActiveApp:
		return "NSTrackingActiveInActiveApp"
	case NSTrackingActiveInKeyWindow:
		return "NSTrackingActiveInKeyWindow"
	case NSTrackingActiveWhenFirstResponder:
		return "NSTrackingActiveWhenFirstResponder"
	case NSTrackingAssumeInside:
		return "NSTrackingAssumeInside"
	case NSTrackingCursorUpdate:
		return "NSTrackingCursorUpdate"
	case NSTrackingEnabledDuringMouseDrag:
		return "NSTrackingEnabledDuringMouseDrag"
	case NSTrackingInVisibleRect:
		return "NSTrackingInVisibleRect"
	case NSTrackingMouseEnteredAndExited:
		return "NSTrackingMouseEnteredAndExited"
	case NSTrackingMouseMoved:
		return "NSTrackingMouseMoved"
	default:
		return fmt.Sprintf("NSTrackingAreaOptions(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/TypesetterBehavior-swift.enum
type NSTypesetterBehavior int

const (
	// NSTypesetterBehavior_10_2: The typesetter behavior introduced in macOS 10.2.
	NSTypesetterBehavior_10_2 NSTypesetterBehavior = 2
	// NSTypesetterBehavior_10_2_WithCompatibility: The macOS 10.2 typesetting behavior that is still compatible with the original typesetter behavior.
	NSTypesetterBehavior_10_2_WithCompatibility NSTypesetterBehavior = 1
	// NSTypesetterBehavior_10_3: The typesetter behavior introduced in macOS 10.3.
	NSTypesetterBehavior_10_3 NSTypesetterBehavior = 3
	// NSTypesetterBehavior_10_4: The typesetter behavior introduced in macOS 10.4.
	NSTypesetterBehavior_10_4 NSTypesetterBehavior = 4
	// NSTypesetterLatestBehavior: The current typesetter behavior in the current operating system.
	NSTypesetterLatestBehavior NSTypesetterBehavior = -1
	// NSTypesetterOriginalBehavior: The original typesetter behavior, as shipped with macOS 10.1 and earlier.
	NSTypesetterOriginalBehavior NSTypesetterBehavior = 0
)


func (e NSTypesetterBehavior) String() string {
	switch e {
	case NSTypesetterBehavior_10_2:
		return "NSTypesetterBehavior_10_2"
	case NSTypesetterBehavior_10_2_WithCompatibility:
		return "NSTypesetterBehavior_10_2_WithCompatibility"
	case NSTypesetterBehavior_10_3:
		return "NSTypesetterBehavior_10_3"
	case NSTypesetterBehavior_10_4:
		return "NSTypesetterBehavior_10_4"
	case NSTypesetterLatestBehavior:
		return "NSTypesetterLatestBehavior"
	case NSTypesetterOriginalBehavior:
		return "NSTypesetterOriginalBehavior"
	default:
		return fmt.Sprintf("NSTypesetterBehavior(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSTypesetterControlCharacterAction
type NSTypesetterControlCharacterAction int

const (
	// NSTypesetterContainerBreakAction: Causes container break.
	NSTypesetterContainerBreakAction NSTypesetterControlCharacterAction = 32
	// NSTypesetterHorizontalTabAction: Treated as tab character.
	NSTypesetterHorizontalTabAction NSTypesetterControlCharacterAction = 4
	// NSTypesetterLineBreakAction: Causes line break.
	NSTypesetterLineBreakAction NSTypesetterControlCharacterAction = 8
	// NSTypesetterParagraphBreakAction: Causes paragraph break; the value returned by firstLineHeadIndent is the advancement used for the following glyph.
	NSTypesetterParagraphBreakAction NSTypesetterControlCharacterAction = 16
	// NSTypesetterWhitespaceAction: The width for glyphs with this action are determined by boundingBox(forControlGlyphAt:for:proposedLineFragment:glyphPosition:characterIndex:), if the method is implemented; otherwise, same as [NSTypesetterZeroAdvancementAction].
	NSTypesetterWhitespaceAction NSTypesetterControlCharacterAction = 2
	// NSTypesetterZeroAdvancementAction: Glyphs with this action are filtered out from layout `(notShownAttribute == YES)`.
	NSTypesetterZeroAdvancementAction NSTypesetterControlCharacterAction = 1
)


func (e NSTypesetterControlCharacterAction) String() string {
	switch e {
	case NSTypesetterContainerBreakAction:
		return "NSTypesetterContainerBreakAction"
	case NSTypesetterHorizontalTabAction:
		return "NSTypesetterHorizontalTabAction"
	case NSTypesetterLineBreakAction:
		return "NSTypesetterLineBreakAction"
	case NSTypesetterParagraphBreakAction:
		return "NSTypesetterParagraphBreakAction"
	case NSTypesetterWhitespaceAction:
		return "NSTypesetterWhitespaceAction"
	case NSTypesetterZeroAdvancementAction:
		return "NSTypesetterZeroAdvancementAction"
	default:
		return fmt.Sprintf("NSTypesetterControlCharacterAction(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSUnderlineStyle
type NSUnderlineStyle int

const (
	// NSUnderlineStyleByWord: Draw the line only beneath or through words, not whitespace.
	NSUnderlineStyleByWord NSUnderlineStyle = 32768
	// NSUnderlineStyleDouble: Draw a double line.
	NSUnderlineStyleDouble NSUnderlineStyle = 9
	// NSUnderlineStylePatternDash: Draw a line of dashes.
	NSUnderlineStylePatternDash NSUnderlineStyle = 512
	// NSUnderlineStylePatternDashDot: Draw a line of alternating dashes and dots.
	NSUnderlineStylePatternDashDot NSUnderlineStyle = 768
	// NSUnderlineStylePatternDashDotDot: Draw a line of alternating dashes and two dots.
	NSUnderlineStylePatternDashDotDot NSUnderlineStyle = 1024
	// NSUnderlineStylePatternDot: Draw a line of dots.
	NSUnderlineStylePatternDot NSUnderlineStyle = 256
	// NSUnderlineStyleSingle: Draw a single line.
	NSUnderlineStyleSingle NSUnderlineStyle = 1
	// NSUnderlineStyleThick: Draw a thick line.
	NSUnderlineStyleThick NSUnderlineStyle = 2
)


func (e NSUnderlineStyle) String() string {
	switch e {
	case NSUnderlineStyleByWord:
		return "NSUnderlineStyleByWord"
	case NSUnderlineStyleDouble:
		return "NSUnderlineStyleDouble"
	case NSUnderlineStylePatternDash:
		return "NSUnderlineStylePatternDash"
	case NSUnderlineStylePatternDashDot:
		return "NSUnderlineStylePatternDashDot"
	case NSUnderlineStylePatternDashDotDot:
		return "NSUnderlineStylePatternDashDotDot"
	case NSUnderlineStylePatternDot:
		return "NSUnderlineStylePatternDot"
	case NSUnderlineStyleSingle:
		return "NSUnderlineStyleSingle"
	case NSUnderlineStyleThick:
		return "NSUnderlineStyleThick"
	default:
		return fmt.Sprintf("NSUnderlineStyle(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSScroller/UsableParts-swift.enum
type NSUsableScrollerParts int

const (
	// NSAllScrollerParts: Specifies that the scroller has at least a knob, possibly also scroll buttons.
	NSAllScrollerParts NSUsableScrollerParts = 2
	// NSNoScrollerParts: Specifies that the scroller has neither a knob nor scroll buttons, only the knob slot.
	NSNoScrollerParts NSUsableScrollerParts = 0
	// NSOnlyScrollerArrows: Specifies that the scroller has only scroll buttons, no knob.
	NSOnlyScrollerArrows NSUsableScrollerParts = 1
)


func (e NSUsableScrollerParts) String() string {
	switch e {
	case NSAllScrollerParts:
		return "NSAllScrollerParts"
	case NSNoScrollerParts:
		return "NSNoScrollerParts"
	case NSOnlyScrollerArrows:
		return "NSOnlyScrollerArrows"
	default:
		return fmt.Sprintf("NSUsableScrollerParts(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSUserInterfaceLayoutDirection
type NSUserInterfaceLayoutDirection int

const (
	// NSUserInterfaceLayoutDirectionLeftToRight: Layout direction is left to right.
	NSUserInterfaceLayoutDirectionLeftToRight NSUserInterfaceLayoutDirection = 0
	// NSUserInterfaceLayoutDirectionRightToLeft: Layout direction is right to left.
	NSUserInterfaceLayoutDirectionRightToLeft NSUserInterfaceLayoutDirection = 1
)


func (e NSUserInterfaceLayoutDirection) String() string {
	switch e {
	case NSUserInterfaceLayoutDirectionLeftToRight:
		return "NSUserInterfaceLayoutDirectionLeftToRight"
	case NSUserInterfaceLayoutDirectionRightToLeft:
		return "NSUserInterfaceLayoutDirectionRightToLeft"
	default:
		return fmt.Sprintf("NSUserInterfaceLayoutDirection(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSUserInterfaceLayoutOrientation
type NSUserInterfaceLayoutOrientation int

const (
	// NSUserInterfaceLayoutOrientationHorizontal: The horizontal orientation.
	NSUserInterfaceLayoutOrientationHorizontal NSUserInterfaceLayoutOrientation = 0
	// NSUserInterfaceLayoutOrientationVertical: The vertical orientation.
	NSUserInterfaceLayoutOrientationVertical NSUserInterfaceLayoutOrientation = 1
)


func (e NSUserInterfaceLayoutOrientation) String() string {
	switch e {
	case NSUserInterfaceLayoutOrientationHorizontal:
		return "NSUserInterfaceLayoutOrientationHorizontal"
	case NSUserInterfaceLayoutOrientationVertical:
		return "NSUserInterfaceLayoutOrientationVertical"
	default:
		return fmt.Sprintf("NSUserInterfaceLayoutOrientation(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSVerticalDirections
type NSVerticalDirections uint

const (
	NSVerticalDirectionsDown NSVerticalDirections = 2
	NSVerticalDirectionsUp NSVerticalDirections = 1
)


func (e NSVerticalDirections) String() string {
	switch e {
	case NSVerticalDirectionsDown:
		return "NSVerticalDirectionsDown"
	case NSVerticalDirectionsUp:
		return "NSVerticalDirectionsUp"
	default:
		return fmt.Sprintf("NSVerticalDirections(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSViewController/TransitionOptions
type NSViewControllerTransitionOptions int

const (
	// NSViewControllerTransitionAllowUserInteraction: A transition animation that allows user interaction during the transition.
	NSViewControllerTransitionAllowUserInteraction NSViewControllerTransitionOptions = 4096
	// NSViewControllerTransitionCrossfade: A transition animation that fades the new view in and simultaneously fades the old view out.
	NSViewControllerTransitionCrossfade NSViewControllerTransitionOptions = 1
	// NSViewControllerTransitionSlideBackward: A transition animation that reflects the user interface layout direction (userInterfaceLayoutDirection) in a “backward” manner, as follows
	NSViewControllerTransitionSlideBackward NSViewControllerTransitionOptions = 384
	// NSViewControllerTransitionSlideDown: A transition animation that slides the old view down while the new view slides into view from the top.
	NSViewControllerTransitionSlideDown NSViewControllerTransitionOptions = 32
	// NSViewControllerTransitionSlideForward: A transition animation that reflects the user interface layout direction (userInterfaceLayoutDirection) in a “forward” manner, as follows:
	NSViewControllerTransitionSlideForward NSViewControllerTransitionOptions = 320
	// NSViewControllerTransitionSlideLeft: A transition animation that slides the old view to the left while the new view slides into view from the right.
	NSViewControllerTransitionSlideLeft NSViewControllerTransitionOptions = 64
	// NSViewControllerTransitionSlideRight: A transition animation that slides the old view to the right while the new view slides into view from the left.
	NSViewControllerTransitionSlideRight NSViewControllerTransitionOptions = 128
	// NSViewControllerTransitionSlideUp: A transition animation that slides the old view up while the new view comes into view from the bottom.
	NSViewControllerTransitionSlideUp NSViewControllerTransitionOptions = 16
)


func (e NSViewControllerTransitionOptions) String() string {
	switch e {
	case NSViewControllerTransitionAllowUserInteraction:
		return "NSViewControllerTransitionAllowUserInteraction"
	case NSViewControllerTransitionCrossfade:
		return "NSViewControllerTransitionCrossfade"
	case NSViewControllerTransitionSlideBackward:
		return "NSViewControllerTransitionSlideBackward"
	case NSViewControllerTransitionSlideDown:
		return "NSViewControllerTransitionSlideDown"
	case NSViewControllerTransitionSlideForward:
		return "NSViewControllerTransitionSlideForward"
	case NSViewControllerTransitionSlideLeft:
		return "NSViewControllerTransitionSlideLeft"
	case NSViewControllerTransitionSlideRight:
		return "NSViewControllerTransitionSlideRight"
	case NSViewControllerTransitionSlideUp:
		return "NSViewControllerTransitionSlideUp"
	default:
		return fmt.Sprintf("NSViewControllerTransitionOptions(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSViewLayoutRegionAdaptivityAxis
type NSViewLayoutRegionAdaptivityAxis int

const (
	NSViewLayoutRegionAdaptivityAxisHorizontal NSViewLayoutRegionAdaptivityAxis = 1
	NSViewLayoutRegionAdaptivityAxisNone NSViewLayoutRegionAdaptivityAxis = 0
	NSViewLayoutRegionAdaptivityAxisVertical NSViewLayoutRegionAdaptivityAxis = 2
)


func (e NSViewLayoutRegionAdaptivityAxis) String() string {
	switch e {
	case NSViewLayoutRegionAdaptivityAxisHorizontal:
		return "NSViewLayoutRegionAdaptivityAxisHorizontal"
	case NSViewLayoutRegionAdaptivityAxisNone:
		return "NSViewLayoutRegionAdaptivityAxisNone"
	case NSViewLayoutRegionAdaptivityAxisVertical:
		return "NSViewLayoutRegionAdaptivityAxisVertical"
	default:
		return fmt.Sprintf("NSViewLayoutRegionAdaptivityAxis(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/BlendingMode-swift.enum
type NSVisualEffectBlendingMode int

const (
	// NSVisualEffectBlendingModeBehindWindow: A mode that blends and blurs the visual effect view with the contents behind the window, such as the desktop or other windows.
	NSVisualEffectBlendingModeBehindWindow NSVisualEffectBlendingMode = 0
	// NSVisualEffectBlendingModeWithinWindow: A mode that blends and blurs the visual effect view with contents behind the view in the current window only.
	NSVisualEffectBlendingModeWithinWindow NSVisualEffectBlendingMode = 1
)


func (e NSVisualEffectBlendingMode) String() string {
	switch e {
	case NSVisualEffectBlendingModeBehindWindow:
		return "NSVisualEffectBlendingModeBehindWindow"
	case NSVisualEffectBlendingModeWithinWindow:
		return "NSVisualEffectBlendingModeWithinWindow"
	default:
		return fmt.Sprintf("NSVisualEffectBlendingMode(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/Material-swift.enum
type NSVisualEffectMaterial int

const (
	// NSVisualEffectMaterialAppearanceBased: A default material for the view’s effective appearance.
	NSVisualEffectMaterialAppearanceBased NSVisualEffectMaterial = 0
	// NSVisualEffectMaterialContentBackground: The material for the background of opaque content.
	NSVisualEffectMaterialContentBackground NSVisualEffectMaterial = 18
	// NSVisualEffectMaterialDark: A material with a dark effect.
	NSVisualEffectMaterialDark NSVisualEffectMaterial = 2
	// NSVisualEffectMaterialFullScreenUI: The material for the background of a full-screen modal interface.
	NSVisualEffectMaterialFullScreenUI NSVisualEffectMaterial = 15
	// NSVisualEffectMaterialHUDWindow: The material for the background of heads-up display (HUD) windows.
	NSVisualEffectMaterialHUDWindow NSVisualEffectMaterial = 13
	// NSVisualEffectMaterialHeaderView: The material for in-line header or footer views.
	NSVisualEffectMaterialHeaderView NSVisualEffectMaterial = 10
	// NSVisualEffectMaterialLight: A material with a light effect.
	NSVisualEffectMaterialLight NSVisualEffectMaterial = 1
	// NSVisualEffectMaterialMediumLight: A material with a medium-light effect.
	NSVisualEffectMaterialMediumLight NSVisualEffectMaterial = 8
	// NSVisualEffectMaterialMenu: The material for menus.
	NSVisualEffectMaterialMenu NSVisualEffectMaterial = 5
	// NSVisualEffectMaterialPopover: The material for the background of popover windows.
	NSVisualEffectMaterialPopover NSVisualEffectMaterial = 6
	// NSVisualEffectMaterialSelection: The material used to indicate a selection.
	NSVisualEffectMaterialSelection NSVisualEffectMaterial = 4
	// NSVisualEffectMaterialSheet: The material for the background of sheet windows.
	NSVisualEffectMaterialSheet NSVisualEffectMaterial = 11
	// NSVisualEffectMaterialSidebar: The material for the background of window sidebars.
	NSVisualEffectMaterialSidebar NSVisualEffectMaterial = 7
	// NSVisualEffectMaterialTitlebar: The material for a window’s titlebar.
	NSVisualEffectMaterialTitlebar NSVisualEffectMaterial = 3
	// NSVisualEffectMaterialToolTip: The material for the background of a tool tip.
	NSVisualEffectMaterialToolTip NSVisualEffectMaterial = 17
	// NSVisualEffectMaterialUltraDark: A material with an ultra-dark effect.
	NSVisualEffectMaterialUltraDark NSVisualEffectMaterial = 9
	// NSVisualEffectMaterialUnderPageBackground: The material for the area behind the pages of a document.
	NSVisualEffectMaterialUnderPageBackground NSVisualEffectMaterial = 22
	// NSVisualEffectMaterialUnderWindowBackground: The material to show under a window’s background.
	NSVisualEffectMaterialUnderWindowBackground NSVisualEffectMaterial = 21
	// NSVisualEffectMaterialWindowBackground: The material for the background of opaque windows.
	NSVisualEffectMaterialWindowBackground NSVisualEffectMaterial = 12
)


func (e NSVisualEffectMaterial) String() string {
	switch e {
	case NSVisualEffectMaterialAppearanceBased:
		return "NSVisualEffectMaterialAppearanceBased"
	case NSVisualEffectMaterialContentBackground:
		return "NSVisualEffectMaterialContentBackground"
	case NSVisualEffectMaterialDark:
		return "NSVisualEffectMaterialDark"
	case NSVisualEffectMaterialFullScreenUI:
		return "NSVisualEffectMaterialFullScreenUI"
	case NSVisualEffectMaterialHUDWindow:
		return "NSVisualEffectMaterialHUDWindow"
	case NSVisualEffectMaterialHeaderView:
		return "NSVisualEffectMaterialHeaderView"
	case NSVisualEffectMaterialLight:
		return "NSVisualEffectMaterialLight"
	case NSVisualEffectMaterialMediumLight:
		return "NSVisualEffectMaterialMediumLight"
	case NSVisualEffectMaterialMenu:
		return "NSVisualEffectMaterialMenu"
	case NSVisualEffectMaterialPopover:
		return "NSVisualEffectMaterialPopover"
	case NSVisualEffectMaterialSelection:
		return "NSVisualEffectMaterialSelection"
	case NSVisualEffectMaterialSheet:
		return "NSVisualEffectMaterialSheet"
	case NSVisualEffectMaterialSidebar:
		return "NSVisualEffectMaterialSidebar"
	case NSVisualEffectMaterialTitlebar:
		return "NSVisualEffectMaterialTitlebar"
	case NSVisualEffectMaterialToolTip:
		return "NSVisualEffectMaterialToolTip"
	case NSVisualEffectMaterialUltraDark:
		return "NSVisualEffectMaterialUltraDark"
	case NSVisualEffectMaterialUnderPageBackground:
		return "NSVisualEffectMaterialUnderPageBackground"
	case NSVisualEffectMaterialUnderWindowBackground:
		return "NSVisualEffectMaterialUnderWindowBackground"
	case NSVisualEffectMaterialWindowBackground:
		return "NSVisualEffectMaterialWindowBackground"
	default:
		return fmt.Sprintf("NSVisualEffectMaterial(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/State-swift.enum
type NSVisualEffectState int

const (
	// NSVisualEffectStateActive: The backdrop should always appear active.
	NSVisualEffectStateActive NSVisualEffectState = 1
	// NSVisualEffectStateFollowsWindowActiveState: The backdrop should automatically appear active when the window is active, and inactive when it is not.
	NSVisualEffectStateFollowsWindowActiveState NSVisualEffectState = 0
	// NSVisualEffectStateInactive: The backdrop should always appear inactive.
	NSVisualEffectStateInactive NSVisualEffectState = 2
)


func (e NSVisualEffectState) String() string {
	switch e {
	case NSVisualEffectStateActive:
		return "NSVisualEffectStateActive"
	case NSVisualEffectStateFollowsWindowActiveState:
		return "NSVisualEffectStateFollowsWindowActiveState"
	case NSVisualEffectStateInactive:
		return "NSVisualEffectStateInactive"
	default:
		return fmt.Sprintf("NSVisualEffectState(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/WindingRule-swift.enum
type NSWindingRule int

const (
	// NSWindingRuleEvenOdd: Specifies the even-odd winding rule.
	NSWindingRuleEvenOdd NSWindingRule = 1
	// NSWindingRuleNonZero: Specifies the non-zero winding rule.
	NSWindingRuleNonZero NSWindingRule = 0
)


func (e NSWindingRule) String() string {
	switch e {
	case NSWindingRuleEvenOdd:
		return "NSWindingRuleEvenOdd"
	case NSWindingRuleNonZero:
		return "NSWindingRuleNonZero"
	default:
		return fmt.Sprintf("NSWindingRule(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSWindow/AnimationBehavior-swift.enum
type NSWindowAnimationBehavior int

const (
	// NSWindowAnimationBehaviorAlertPanel: The animation behavior that’s appropriate to the alert window style.
	NSWindowAnimationBehaviorAlertPanel NSWindowAnimationBehavior = 5
	// NSWindowAnimationBehaviorDefault: The automatic animation that’s appropriate to the window type.
	NSWindowAnimationBehaviorDefault NSWindowAnimationBehavior = 0
	// NSWindowAnimationBehaviorDocumentWindow: The animation behavior that’s appropriate to the document window style.
	NSWindowAnimationBehaviorDocumentWindow NSWindowAnimationBehavior = 3
	// NSWindowAnimationBehaviorNone: No automatic animation used.
	NSWindowAnimationBehaviorNone NSWindowAnimationBehavior = 2
	// NSWindowAnimationBehaviorUtilityWindow: The animation behavior that’s appropriate to the utility window style.
	NSWindowAnimationBehaviorUtilityWindow NSWindowAnimationBehavior = 4
)


func (e NSWindowAnimationBehavior) String() string {
	switch e {
	case NSWindowAnimationBehaviorAlertPanel:
		return "NSWindowAnimationBehaviorAlertPanel"
	case NSWindowAnimationBehaviorDefault:
		return "NSWindowAnimationBehaviorDefault"
	case NSWindowAnimationBehaviorDocumentWindow:
		return "NSWindowAnimationBehaviorDocumentWindow"
	case NSWindowAnimationBehaviorNone:
		return "NSWindowAnimationBehaviorNone"
	case NSWindowAnimationBehaviorUtilityWindow:
		return "NSWindowAnimationBehaviorUtilityWindow"
	default:
		return fmt.Sprintf("NSWindowAnimationBehavior(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSWindow/ButtonType
type NSWindowButton int

const (
	// NSWindowCloseButton: The close button.
	NSWindowCloseButton NSWindowButton = 0
	// NSWindowDocumentIconButton: The document icon button.
	NSWindowDocumentIconButton NSWindowButton = 4
	// NSWindowDocumentVersionsButton: The document versions button.
	NSWindowDocumentVersionsButton NSWindowButton = 6
	// NSWindowMiniaturizeButton: The minimize button.
	NSWindowMiniaturizeButton NSWindowButton = 1
	// NSWindowToolbarButton: The toolbar button.
	NSWindowToolbarButton NSWindowButton = 3
	// NSWindowZoomButton: The zoom button.
	NSWindowZoomButton NSWindowButton = 2
)


func (e NSWindowButton) String() string {
	switch e {
	case NSWindowCloseButton:
		return "NSWindowCloseButton"
	case NSWindowDocumentIconButton:
		return "NSWindowDocumentIconButton"
	case NSWindowDocumentVersionsButton:
		return "NSWindowDocumentVersionsButton"
	case NSWindowMiniaturizeButton:
		return "NSWindowMiniaturizeButton"
	case NSWindowToolbarButton:
		return "NSWindowToolbarButton"
	case NSWindowZoomButton:
		return "NSWindowZoomButton"
	default:
		return fmt.Sprintf("NSWindowButton(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct
type NSWindowCollectionBehavior int

const (
	// NSWindowCollectionBehaviorAuxiliary: The behavior marking this window as auxiliary for both Stage Manager and full screen.
	NSWindowCollectionBehaviorAuxiliary NSWindowCollectionBehavior = 131072
	// NSWindowCollectionBehaviorCanJoinAllApplications: The behavior marking this window as one that can join all apps for both Stage Manager and full screen.
	NSWindowCollectionBehaviorCanJoinAllApplications NSWindowCollectionBehavior = 262144
	// NSWindowCollectionBehaviorCanJoinAllSpaces: The window can appear in all spaces.
	NSWindowCollectionBehaviorCanJoinAllSpaces NSWindowCollectionBehavior = 1
	// NSWindowCollectionBehaviorFullScreenAllowsTiling: The window can be a secondary full screen tile even if it can’t be a full screen window itself.
	NSWindowCollectionBehaviorFullScreenAllowsTiling NSWindowCollectionBehavior = 2048
	// NSWindowCollectionBehaviorFullScreenAuxiliary: The window displays on the same space as the full screen window.
	NSWindowCollectionBehaviorFullScreenAuxiliary NSWindowCollectionBehavior = 256
	// NSWindowCollectionBehaviorFullScreenDisallowsTiling: The window doesn’t support being a full-screen tile window, but may support being a full-screen window.
	NSWindowCollectionBehaviorFullScreenDisallowsTiling NSWindowCollectionBehavior = 4096
	// NSWindowCollectionBehaviorFullScreenNone: The window doesn’t support full-screen mode.
	NSWindowCollectionBehaviorFullScreenNone NSWindowCollectionBehavior = 512
	// NSWindowCollectionBehaviorFullScreenPrimary: The window can enter full-screen mode.
	NSWindowCollectionBehaviorFullScreenPrimary NSWindowCollectionBehavior = 128
	// NSWindowCollectionBehaviorIgnoresCycle: The window isn’t part of the window cycle for use with the Cycle Through Windows menu item.
	NSWindowCollectionBehaviorIgnoresCycle NSWindowCollectionBehavior = 64
	// NSWindowCollectionBehaviorManaged: The window participates in Mission Control and Spaces.
	NSWindowCollectionBehaviorManaged NSWindowCollectionBehavior = 4
	// NSWindowCollectionBehaviorMoveToActiveSpace: When the window becomes active, move it to the active space instead of switching spaces.
	NSWindowCollectionBehaviorMoveToActiveSpace NSWindowCollectionBehavior = 2
	// NSWindowCollectionBehaviorParticipatesInCycle: The window participates in the window cycle for use with the Cycle Through Windows menu item.
	NSWindowCollectionBehaviorParticipatesInCycle NSWindowCollectionBehavior = 32
	// NSWindowCollectionBehaviorPrimary: The behavior marking this window as primary for both Stage Manager and full screen.
	NSWindowCollectionBehaviorPrimary NSWindowCollectionBehavior = 65536
	// NSWindowCollectionBehaviorStationary: Mission Control doesn’t affect the window, so it stays visible and stationary, like the desktop window.
	NSWindowCollectionBehaviorStationary NSWindowCollectionBehavior = 16
	// NSWindowCollectionBehaviorTransient: The window floats in Spaces and hides in Mission Control.
	NSWindowCollectionBehaviorTransient NSWindowCollectionBehavior = 8
)


func (e NSWindowCollectionBehavior) String() string {
	switch e {
	case NSWindowCollectionBehaviorAuxiliary:
		return "NSWindowCollectionBehaviorAuxiliary"
	case NSWindowCollectionBehaviorCanJoinAllApplications:
		return "NSWindowCollectionBehaviorCanJoinAllApplications"
	case NSWindowCollectionBehaviorCanJoinAllSpaces:
		return "NSWindowCollectionBehaviorCanJoinAllSpaces"
	case NSWindowCollectionBehaviorFullScreenAllowsTiling:
		return "NSWindowCollectionBehaviorFullScreenAllowsTiling"
	case NSWindowCollectionBehaviorFullScreenAuxiliary:
		return "NSWindowCollectionBehaviorFullScreenAuxiliary"
	case NSWindowCollectionBehaviorFullScreenDisallowsTiling:
		return "NSWindowCollectionBehaviorFullScreenDisallowsTiling"
	case NSWindowCollectionBehaviorFullScreenNone:
		return "NSWindowCollectionBehaviorFullScreenNone"
	case NSWindowCollectionBehaviorFullScreenPrimary:
		return "NSWindowCollectionBehaviorFullScreenPrimary"
	case NSWindowCollectionBehaviorIgnoresCycle:
		return "NSWindowCollectionBehaviorIgnoresCycle"
	case NSWindowCollectionBehaviorManaged:
		return "NSWindowCollectionBehaviorManaged"
	case NSWindowCollectionBehaviorMoveToActiveSpace:
		return "NSWindowCollectionBehaviorMoveToActiveSpace"
	case NSWindowCollectionBehaviorParticipatesInCycle:
		return "NSWindowCollectionBehaviorParticipatesInCycle"
	case NSWindowCollectionBehaviorPrimary:
		return "NSWindowCollectionBehaviorPrimary"
	case NSWindowCollectionBehaviorStationary:
		return "NSWindowCollectionBehaviorStationary"
	case NSWindowCollectionBehaviorTransient:
		return "NSWindowCollectionBehaviorTransient"
	default:
		return fmt.Sprintf("NSWindowCollectionBehavior(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSWindow/Depth
type NSWindowDepth int

const (
	// NSWindowDepthOnehundredtwentyeightBitRGB: One hundred and twenty eight bit RGB depth limit.
	NSWindowDepthOnehundredtwentyeightBitRGB NSWindowDepth = 544
	// NSWindowDepthSixtyfourBitRGB: Sixty four bit RGB depth limit.
	NSWindowDepthSixtyfourBitRGB NSWindowDepth = 528
	// NSWindowDepthTwentyfourBitRGB: Twenty four bit RGB depth limit.
	NSWindowDepthTwentyfourBitRGB NSWindowDepth = 520
)


func (e NSWindowDepth) String() string {
	switch e {
	case NSWindowDepthOnehundredtwentyeightBitRGB:
		return "NSWindowDepthOnehundredtwentyeightBitRGB"
	case NSWindowDepthSixtyfourBitRGB:
		return "NSWindowDepthSixtyfourBitRGB"
	case NSWindowDepthTwentyfourBitRGB:
		return "NSWindowDepthTwentyfourBitRGB"
	default:
		return fmt.Sprintf("NSWindowDepth(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSWindow/NumberListOptions
type NSWindowNumberListOptions int

const (
	// NSWindowNumberListAllApplications: # Discussion
	NSWindowNumberListAllApplications NSWindowNumberListOptions = 1
	// NSWindowNumberListAllSpaces: # Discussion
	NSWindowNumberListAllSpaces NSWindowNumberListOptions = 16
)


func (e NSWindowNumberListOptions) String() string {
	switch e {
	case NSWindowNumberListAllApplications:
		return "NSWindowNumberListAllApplications"
	case NSWindowNumberListAllSpaces:
		return "NSWindowNumberListAllSpaces"
	default:
		return fmt.Sprintf("NSWindowNumberListOptions(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSWindow/OcclusionState-swift.struct
type NSWindowOcclusionState int

const (
	// NSWindowOcclusionStateVisible: If set, at least part of the window is visible; if not set, the entire window is occluded.
	NSWindowOcclusionStateVisible NSWindowOcclusionState = 2
)


func (e NSWindowOcclusionState) String() string {
	switch e {
	case NSWindowOcclusionStateVisible:
		return "NSWindowOcclusionStateVisible"
	default:
		return fmt.Sprintf("NSWindowOcclusionState(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSWindow/OrderingMode
type NSWindowOrderingMode int

const (
	// NSWindowAbove: Moves the window above the indicated window.
	NSWindowAbove NSWindowOrderingMode = 1
	// NSWindowBelow: Moves the window below the indicated window.
	NSWindowBelow NSWindowOrderingMode = -1
	// NSWindowOut: Moves the window off the screen.
	NSWindowOut NSWindowOrderingMode = 0
)


func (e NSWindowOrderingMode) String() string {
	switch e {
	case NSWindowAbove:
		return "NSWindowAbove"
	case NSWindowBelow:
		return "NSWindowBelow"
	case NSWindowOut:
		return "NSWindowOut"
	default:
		return fmt.Sprintf("NSWindowOrderingMode(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSWindow/SharingType-swift.enum
type NSWindowSharingType int

const (
	// NSWindowSharingNone: A legacy constant that macOS no longer uses.
	NSWindowSharingNone NSWindowSharingType = 0
)


func (e NSWindowSharingType) String() string {
	switch e {
	case NSWindowSharingNone:
		return "NSWindowSharingNone"
	default:
		return fmt.Sprintf("NSWindowSharingType(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct
type NSWindowStyleMask int

const (
	// NSWindowStyleMaskBorderless: The window displays none of the usual peripheral elements.
	NSWindowStyleMaskBorderless NSWindowStyleMask = 0
	// NSWindowStyleMaskClosable: The window displays a close button.
	NSWindowStyleMaskClosable NSWindowStyleMask = 2
	// NSWindowStyleMaskDocModalWindow: The window is a document-modal panel (or  a subclass of NSPanel).
	NSWindowStyleMaskDocModalWindow NSWindowStyleMask = 64
	// NSWindowStyleMaskFullScreen: The window can appear full screen.
	NSWindowStyleMaskFullScreen NSWindowStyleMask = 16384
	// NSWindowStyleMaskFullSizeContentView: When set, the window’s contentView consumes the full size of the window.
	NSWindowStyleMaskFullSizeContentView NSWindowStyleMask = 32768
	// NSWindowStyleMaskHUDWindow: The window is a HUD panel.
	NSWindowStyleMaskHUDWindow NSWindowStyleMask = 8192
	// NSWindowStyleMaskMiniaturizable: The window displays a minimize button.
	NSWindowStyleMaskMiniaturizable NSWindowStyleMask = 4
	// NSWindowStyleMaskNonactivatingPanel: The window is a panel or a subclass of NSPanel that does not activate the owning app.
	NSWindowStyleMaskNonactivatingPanel NSWindowStyleMask = 128
	// NSWindowStyleMaskResizable: The window can be resized by the user.
	NSWindowStyleMaskResizable NSWindowStyleMask = 8
	// NSWindowStyleMaskTitled: The window displays a title bar.
	NSWindowStyleMaskTitled NSWindowStyleMask = 1
	// NSWindowStyleMaskUnifiedTitleAndToolbar: This constant has no effect, because all windows that include a toolbar use the unified style.
	NSWindowStyleMaskUnifiedTitleAndToolbar NSWindowStyleMask = 4096
	// NSWindowStyleMaskUtilityWindow: The window is a panel or a subclass of NSPanel.
	NSWindowStyleMaskUtilityWindow NSWindowStyleMask = 16
	// Deprecated.
	NSWindowStyleMaskTexturedBackground NSWindowStyleMask = 256
)


func (e NSWindowStyleMask) String() string {
	switch e {
	case NSWindowStyleMaskBorderless:
		return "NSWindowStyleMaskBorderless"
	case NSWindowStyleMaskClosable:
		return "NSWindowStyleMaskClosable"
	case NSWindowStyleMaskDocModalWindow:
		return "NSWindowStyleMaskDocModalWindow"
	case NSWindowStyleMaskFullScreen:
		return "NSWindowStyleMaskFullScreen"
	case NSWindowStyleMaskFullSizeContentView:
		return "NSWindowStyleMaskFullSizeContentView"
	case NSWindowStyleMaskHUDWindow:
		return "NSWindowStyleMaskHUDWindow"
	case NSWindowStyleMaskMiniaturizable:
		return "NSWindowStyleMaskMiniaturizable"
	case NSWindowStyleMaskNonactivatingPanel:
		return "NSWindowStyleMaskNonactivatingPanel"
	case NSWindowStyleMaskResizable:
		return "NSWindowStyleMaskResizable"
	case NSWindowStyleMaskTitled:
		return "NSWindowStyleMaskTitled"
	case NSWindowStyleMaskUnifiedTitleAndToolbar:
		return "NSWindowStyleMaskUnifiedTitleAndToolbar"
	case NSWindowStyleMaskUtilityWindow:
		return "NSWindowStyleMaskUtilityWindow"
	case NSWindowStyleMaskTexturedBackground:
		return "NSWindowStyleMaskTexturedBackground"
	default:
		return fmt.Sprintf("NSWindowStyleMask(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSWindow/TabbingMode-swift.enum
type NSWindowTabbingMode int

const (
	// NSWindowTabbingModeAutomatic: A window that automatically tabs together based on the user’s tabbing preference.
	NSWindowTabbingModeAutomatic NSWindowTabbingMode = 0
	// NSWindowTabbingModeDisallowed: A window that explicitly does not prefer to tab together with other windows.
	NSWindowTabbingModeDisallowed NSWindowTabbingMode = 2
	// NSWindowTabbingModePreferred: A window that explicitly prefers to tab together with other windows with the same tabbing identifier.
	NSWindowTabbingModePreferred NSWindowTabbingMode = 1
)


func (e NSWindowTabbingMode) String() string {
	switch e {
	case NSWindowTabbingModeAutomatic:
		return "NSWindowTabbingModeAutomatic"
	case NSWindowTabbingModeDisallowed:
		return "NSWindowTabbingModeDisallowed"
	case NSWindowTabbingModePreferred:
		return "NSWindowTabbingModePreferred"
	default:
		return fmt.Sprintf("NSWindowTabbingMode(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSWindow/TitleVisibility-swift.enum
type NSWindowTitleVisibility int

const (
	// NSWindowTitleVisible: The window has the regular window title and title bar buttons.
	NSWindowTitleVisible NSWindowTitleVisibility = 0
)


func (e NSWindowTitleVisibility) String() string {
	switch e {
	case NSWindowTitleVisible:
		return "NSWindowTitleVisible"
	default:
		return fmt.Sprintf("NSWindowTitleVisibility(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSWindow/ToolbarStyle-swift.enum
type NSWindowToolbarStyle int

const (
	// NSWindowToolbarStyleAutomatic: A style indicating that the system determines the toolbar’s appearance and location.
	NSWindowToolbarStyleAutomatic NSWindowToolbarStyle = 0
	// NSWindowToolbarStyleExpanded: A style indicating that the toolbar appears below the window title.
	NSWindowToolbarStyleExpanded NSWindowToolbarStyle = 1
	// NSWindowToolbarStylePreference: A style indicating that the toolbar appears below the window title with toolbar items centered in the toolbar.
	NSWindowToolbarStylePreference NSWindowToolbarStyle = 2
	// NSWindowToolbarStyleUnified: A style indicating that the toolbar appears next to the window title.
	NSWindowToolbarStyleUnified NSWindowToolbarStyle = 3
	// NSWindowToolbarStyleUnifiedCompact: A style indicating that the toolbar appears next to the window title and with reduced margins to allow more focus on the window’s contents.
	NSWindowToolbarStyleUnifiedCompact NSWindowToolbarStyle = 4
)


func (e NSWindowToolbarStyle) String() string {
	switch e {
	case NSWindowToolbarStyleAutomatic:
		return "NSWindowToolbarStyleAutomatic"
	case NSWindowToolbarStyleExpanded:
		return "NSWindowToolbarStyleExpanded"
	case NSWindowToolbarStylePreference:
		return "NSWindowToolbarStylePreference"
	case NSWindowToolbarStyleUnified:
		return "NSWindowToolbarStyleUnified"
	case NSWindowToolbarStyleUnifiedCompact:
		return "NSWindowToolbarStyleUnifiedCompact"
	default:
		return fmt.Sprintf("NSWindowToolbarStyle(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSWindow/UserTabbingPreference-swift.enum
type NSWindowUserTabbingPreference int

const (
	// NSWindowUserTabbingPreferenceAlways: A value that indicates a window should always display as tabs.
	NSWindowUserTabbingPreferenceAlways NSWindowUserTabbingPreference = 1
	// NSWindowUserTabbingPreferenceInFullScreen: A value that indicates a window should only display as tabs when in full-screen mode.
	NSWindowUserTabbingPreferenceInFullScreen NSWindowUserTabbingPreference = 2
	// NSWindowUserTabbingPreferenceManual: A value that indicates a window should display as tabs according to the window’s tabbing mode.
	NSWindowUserTabbingPreferenceManual NSWindowUserTabbingPreference = 0
)


func (e NSWindowUserTabbingPreference) String() string {
	switch e {
	case NSWindowUserTabbingPreferenceAlways:
		return "NSWindowUserTabbingPreferenceAlways"
	case NSWindowUserTabbingPreferenceInFullScreen:
		return "NSWindowUserTabbingPreferenceInFullScreen"
	case NSWindowUserTabbingPreferenceManual:
		return "NSWindowUserTabbingPreferenceManual"
	default:
		return fmt.Sprintf("NSWindowUserTabbingPreference(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/AuthorizationType
type NSWorkspaceAuthorizationType int

const (
	// NSWorkspaceAuthorizationTypeCreateSymbolicLink: Authorization for the app to create a symbolic link.
	NSWorkspaceAuthorizationTypeCreateSymbolicLink NSWorkspaceAuthorizationType = 0
	// NSWorkspaceAuthorizationTypeReplaceFile: Authorization for the app to perform an atomic file write without changing the target file’s permissions.
	NSWorkspaceAuthorizationTypeReplaceFile NSWorkspaceAuthorizationType = 2
	// NSWorkspaceAuthorizationTypeSetAttributes: Authorization for the app to change specific file attributes.
	NSWorkspaceAuthorizationTypeSetAttributes NSWorkspaceAuthorizationType = 1
)


func (e NSWorkspaceAuthorizationType) String() string {
	switch e {
	case NSWorkspaceAuthorizationTypeCreateSymbolicLink:
		return "NSWorkspaceAuthorizationTypeCreateSymbolicLink"
	case NSWorkspaceAuthorizationTypeReplaceFile:
		return "NSWorkspaceAuthorizationTypeReplaceFile"
	case NSWorkspaceAuthorizationTypeSetAttributes:
		return "NSWorkspaceAuthorizationTypeSetAttributes"
	default:
		return fmt.Sprintf("NSWorkspaceAuthorizationType(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/IconCreationOptions
type NSWorkspaceIconCreationOptions int

const (
	// NSExclude10_4ElementsIconCreationOption: An option to suppress generation of the new higher resolution icon representations that are supported in macOS 10.4.
	NSExclude10_4ElementsIconCreationOption NSWorkspaceIconCreationOptions = 4
	// NSExcludeQuickDrawElementsIconCreationOption: An option to suppress generation of the QuickDraw format icon representations that are used in macOS 10.0 through macOS 10.4.
	NSExcludeQuickDrawElementsIconCreationOption NSWorkspaceIconCreationOptions = 2
)


func (e NSWorkspaceIconCreationOptions) String() string {
	switch e {
	case NSExclude10_4ElementsIconCreationOption:
		return "NSExclude10_4ElementsIconCreationOption"
	case NSExcludeQuickDrawElementsIconCreationOption:
		return "NSExcludeQuickDrawElementsIconCreationOption"
	default:
		return fmt.Sprintf("NSWorkspaceIconCreationOptions(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSWritingDirection
type NSWritingDirection int

const (
	// NSWritingDirectionLeftToRight: The writing direction is left to right.
	NSWritingDirectionLeftToRight NSWritingDirection = 0
	// NSWritingDirectionRightToLeft: The writing direction is right to left.
	NSWritingDirectionRightToLeft NSWritingDirection = 1
)


func (e NSWritingDirection) String() string {
	switch e {
	case NSWritingDirectionLeftToRight:
		return "NSWritingDirectionLeftToRight"
	case NSWritingDirectionRightToLeft:
		return "NSWritingDirectionRightToLeft"
	default:
		return fmt.Sprintf("NSWritingDirection(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSWritingDirectionFormatType
type NSWritingDirectionFormatType int

const (
	// NSWritingDirectionEmbedding: Text is embedded in text with another writing direction.
	NSWritingDirectionEmbedding NSWritingDirectionFormatType = 0
	// NSWritingDirectionOverride: Enables character types with inherent directionality to be overridden when required for special cases, such as for part numbers made of mixed English, digits, and Hebrew letters to be written from right to left.
	NSWritingDirectionOverride NSWritingDirectionFormatType = 2
)


func (e NSWritingDirectionFormatType) String() string {
	switch e {
	case NSWritingDirectionEmbedding:
		return "NSWritingDirectionEmbedding"
	case NSWritingDirectionOverride:
		return "NSWritingDirectionOverride"
	default:
		return fmt.Sprintf("NSWritingDirectionFormatType(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsBehavior
type NSWritingToolsBehavior int

const (
	// NSWritingToolsBehaviorComplete: An option to provide the complete Writing Tools experience for the text view.
	NSWritingToolsBehaviorComplete NSWritingToolsBehavior = 1
	// NSWritingToolsBehaviorDefault: An option to let the system determine the best way to enable Writing Tools for the view.
	NSWritingToolsBehaviorDefault NSWritingToolsBehavior = 0
	// NSWritingToolsBehaviorLimited: An option to provide a limited, overlay-panel experience for the text view.
	NSWritingToolsBehaviorLimited NSWritingToolsBehavior = 2
	// NSWritingToolsBehaviorNone: An option to prevent Writing Tools from modifying the text in the view.
	NSWritingToolsBehaviorNone NSWritingToolsBehavior = -1
)


func (e NSWritingToolsBehavior) String() string {
	switch e {
	case NSWritingToolsBehaviorComplete:
		return "NSWritingToolsBehaviorComplete"
	case NSWritingToolsBehaviorDefault:
		return "NSWritingToolsBehaviorDefault"
	case NSWritingToolsBehaviorLimited:
		return "NSWritingToolsBehaviorLimited"
	case NSWritingToolsBehaviorNone:
		return "NSWritingToolsBehaviorNone"
	default:
		return fmt.Sprintf("NSWritingToolsBehavior(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/ContextScope
type NSWritingToolsCoordinatorContextScope int

const (
	// NSWritingToolsCoordinatorContextScopeFullDocument: An option to provide all of your view’s text.
	NSWritingToolsCoordinatorContextScopeFullDocument NSWritingToolsCoordinatorContextScope = 1
	// NSWritingToolsCoordinatorContextScopeUserSelection: An option to provide only the view’s currently selected text.
	NSWritingToolsCoordinatorContextScopeUserSelection NSWritingToolsCoordinatorContextScope = 0
	// NSWritingToolsCoordinatorContextScopeVisibleArea: An option to provide only the text in the currently visible portion of your view.
	NSWritingToolsCoordinatorContextScopeVisibleArea NSWritingToolsCoordinatorContextScope = 2
)


func (e NSWritingToolsCoordinatorContextScope) String() string {
	switch e {
	case NSWritingToolsCoordinatorContextScopeFullDocument:
		return "NSWritingToolsCoordinatorContextScopeFullDocument"
	case NSWritingToolsCoordinatorContextScopeUserSelection:
		return "NSWritingToolsCoordinatorContextScopeUserSelection"
	case NSWritingToolsCoordinatorContextScopeVisibleArea:
		return "NSWritingToolsCoordinatorContextScopeVisibleArea"
	default:
		return fmt.Sprintf("NSWritingToolsCoordinatorContextScope(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/State-swift.enum
type NSWritingToolsCoordinatorState int

const (
	// NSWritingToolsCoordinatorStateInactive: A state that indicates Writing Tools isn’t currently performing any work on your view’s content.
	NSWritingToolsCoordinatorStateInactive NSWritingToolsCoordinatorState = 0
	// NSWritingToolsCoordinatorStateInteractiveResting: A state that indicates Writing Tools is in the resting state for an inline editing experience.
	NSWritingToolsCoordinatorStateInteractiveResting NSWritingToolsCoordinatorState = 2
	// NSWritingToolsCoordinatorStateInteractiveStreaming: A state that indicates Writing Tools is processing a request and incorporating changes interactively into your view.
	NSWritingToolsCoordinatorStateInteractiveStreaming NSWritingToolsCoordinatorState = 3
	// NSWritingToolsCoordinatorStateNoninteractive: A state that indicates Writing Tools is handling interactions in the system UI, instead of in your view.
	NSWritingToolsCoordinatorStateNoninteractive NSWritingToolsCoordinatorState = 1
)


func (e NSWritingToolsCoordinatorState) String() string {
	switch e {
	case NSWritingToolsCoordinatorStateInactive:
		return "NSWritingToolsCoordinatorStateInactive"
	case NSWritingToolsCoordinatorStateInteractiveResting:
		return "NSWritingToolsCoordinatorStateInteractiveResting"
	case NSWritingToolsCoordinatorStateInteractiveStreaming:
		return "NSWritingToolsCoordinatorStateInteractiveStreaming"
	case NSWritingToolsCoordinatorStateNoninteractive:
		return "NSWritingToolsCoordinatorStateNoninteractive"
	default:
		return fmt.Sprintf("NSWritingToolsCoordinatorState(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/TextAnimation
type NSWritingToolsCoordinatorTextAnimation int

const (
	// NSWritingToolsCoordinatorTextAnimationAnticipate: The animation that Writing Tools performs when waiting to receive results from the large language model.
	NSWritingToolsCoordinatorTextAnimationAnticipate NSWritingToolsCoordinatorTextAnimation = 0
	// NSWritingToolsCoordinatorTextAnimationAnticipateInactive: The animation effect that Writing Tools performs when the view is waiting for results, but the system isn’t actively evaluating the text.
	NSWritingToolsCoordinatorTextAnimationAnticipateInactive NSWritingToolsCoordinatorTextAnimation = 8
	// NSWritingToolsCoordinatorTextAnimationInsert: The animation that Writing Tools performs when inserting text into your view.
	NSWritingToolsCoordinatorTextAnimationInsert NSWritingToolsCoordinatorTextAnimation = 2
	// NSWritingToolsCoordinatorTextAnimationRemove: The animation that Writing Tools performs when removing text from your view.
	NSWritingToolsCoordinatorTextAnimationRemove NSWritingToolsCoordinatorTextAnimation = 1
	// NSWritingToolsCoordinatorTextAnimationTranslate: The animation effect that Writing Tools performs on text situated after the insertion point.
	NSWritingToolsCoordinatorTextAnimationTranslate NSWritingToolsCoordinatorTextAnimation = 9
)


func (e NSWritingToolsCoordinatorTextAnimation) String() string {
	switch e {
	case NSWritingToolsCoordinatorTextAnimationAnticipate:
		return "NSWritingToolsCoordinatorTextAnimationAnticipate"
	case NSWritingToolsCoordinatorTextAnimationAnticipateInactive:
		return "NSWritingToolsCoordinatorTextAnimationAnticipateInactive"
	case NSWritingToolsCoordinatorTextAnimationInsert:
		return "NSWritingToolsCoordinatorTextAnimationInsert"
	case NSWritingToolsCoordinatorTextAnimationRemove:
		return "NSWritingToolsCoordinatorTextAnimationRemove"
	case NSWritingToolsCoordinatorTextAnimationTranslate:
		return "NSWritingToolsCoordinatorTextAnimationTranslate"
	default:
		return fmt.Sprintf("NSWritingToolsCoordinatorTextAnimation(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/TextReplacementReason
type NSWritingToolsCoordinatorTextReplacementReason int

const (
	// NSWritingToolsCoordinatorTextReplacementReasonInteractive: An option to animate the replacement of text in your view.
	NSWritingToolsCoordinatorTextReplacementReasonInteractive NSWritingToolsCoordinatorTextReplacementReason = 0
	// NSWritingToolsCoordinatorTextReplacementReasonNoninteractive: An option to replace the text in your view without animating the change.
	NSWritingToolsCoordinatorTextReplacementReasonNoninteractive NSWritingToolsCoordinatorTextReplacementReason = 1
)


func (e NSWritingToolsCoordinatorTextReplacementReason) String() string {
	switch e {
	case NSWritingToolsCoordinatorTextReplacementReasonInteractive:
		return "NSWritingToolsCoordinatorTextReplacementReasonInteractive"
	case NSWritingToolsCoordinatorTextReplacementReasonNoninteractive:
		return "NSWritingToolsCoordinatorTextReplacementReasonNoninteractive"
	default:
		return fmt.Sprintf("NSWritingToolsCoordinatorTextReplacementReason(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/TextUpdateReason
type NSWritingToolsCoordinatorTextUpdateReason int

const (
	// NSWritingToolsCoordinatorTextUpdateReasonTyping: An operation that involved a person editing the text in your view.
	NSWritingToolsCoordinatorTextUpdateReasonTyping NSWritingToolsCoordinatorTextUpdateReason = 0
	// NSWritingToolsCoordinatorTextUpdateReasonUndoRedo: An operation that changed the view’s text as part of an undo or redo command.
	NSWritingToolsCoordinatorTextUpdateReasonUndoRedo NSWritingToolsCoordinatorTextUpdateReason = 1
)


func (e NSWritingToolsCoordinatorTextUpdateReason) String() string {
	switch e {
	case NSWritingToolsCoordinatorTextUpdateReasonTyping:
		return "NSWritingToolsCoordinatorTextUpdateReasonTyping"
	case NSWritingToolsCoordinatorTextUpdateReasonUndoRedo:
		return "NSWritingToolsCoordinatorTextUpdateReasonUndoRedo"
	default:
		return fmt.Sprintf("NSWritingToolsCoordinatorTextUpdateReason(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsResultOptions
type NSWritingToolsResultOptions int

const (
	// NSWritingToolsResultList: An option to allow list-style formatting in the returned text.
	NSWritingToolsResultList NSWritingToolsResultOptions = 4
	// NSWritingToolsResultPlainText: An option to allow only plain text without any attributes in the returned text.
	NSWritingToolsResultPlainText NSWritingToolsResultOptions = 1
	NSWritingToolsResultPresentationIntent NSWritingToolsResultOptions = 16
	// NSWritingToolsResultRichText: An option to include style attributes consistent with the RTF format in the returned text.
	NSWritingToolsResultRichText NSWritingToolsResultOptions = 2
	// NSWritingToolsResultTable: An option to allow tabular layout attributes in the returned text.
	NSWritingToolsResultTable NSWritingToolsResultOptions = 8
)


func (e NSWritingToolsResultOptions) String() string {
	switch e {
	case NSWritingToolsResultList:
		return "NSWritingToolsResultList"
	case NSWritingToolsResultPlainText:
		return "NSWritingToolsResultPlainText"
	case NSWritingToolsResultPresentationIntent:
		return "NSWritingToolsResultPresentationIntent"
	case NSWritingToolsResultRichText:
		return "NSWritingToolsResultRichText"
	case NSWritingToolsResultTable:
		return "NSWritingToolsResultTable"
	default:
		return fmt.Sprintf("NSWritingToolsResultOptions(%d)", e)
	}
}




type Nsfp uint

const (
	// Deprecated.
	NSFPCurrentField Nsfp = 134
	// Deprecated.
	NSFPPreviewButton Nsfp = 131
	// Deprecated.
	NSFPPreviewField Nsfp = 128
	// Deprecated.
	NSFPRevertButton Nsfp = 130
	// Deprecated.
	NSFPSetButton Nsfp = 132
	// Deprecated.
	NSFPSizeField Nsfp = 129
	// Deprecated.
	NSFPSizeTitle Nsfp = 133
)


func (e Nsfp) String() string {
	switch e {
	case NSFPCurrentField:
		return "NSFPCurrentField"
	case NSFPPreviewButton:
		return "NSFPPreviewButton"
	case NSFPPreviewField:
		return "NSFPPreviewField"
	case NSFPRevertButton:
		return "NSFPRevertButton"
	case NSFPSetButton:
		return "NSFPSetButton"
	case NSFPSizeField:
		return "NSFPSizeField"
	case NSFPSizeTitle:
		return "NSFPSizeTitle"
	default:
		return fmt.Sprintf("Nsfp(%d)", e)
	}
}





