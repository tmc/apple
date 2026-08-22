// Code generated from Apple documentation. DO NOT EDIT.

package coretext

import (
	"unsafe"

	"github.com/tmc/apple/corefoundation"
)

// See: https://developer.apple.com/documentation/CoreText/ATSFontRef
type ATSFontRef = uint32

// See: https://developer.apple.com/documentation/CoreText/BslnBaselineClass
type BslnBaselineClass = uint32

// See: https://developer.apple.com/documentation/CoreText/BslnBaselineRecord
type BslnBaselineRecord = [32]int32

// See: https://developer.apple.com/documentation/CoreText/BslnTableFormat
type BslnTableFormat = uint16

// See: https://developer.apple.com/documentation/CoreText/BslnTablePtr
type BslnTablePtr = *BslnTable

// CTFontCollectionRef is a font collection.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontCollection
type CTFontCollectionRef uintptr

// CTFontCollectionSortDescriptorsCallback is the collection sorting callback type.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontCollectionSortDescriptorsCallback
type CTFontCollectionSortDescriptorsCallback = func(first CTFontDescriptorRef, second CTFontDescriptorRef, refCon unsafe.Pointer) corefoundation.CFComparisonResult

// CTFontDescriptorProgressHandler is the progress callback type.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontDescriptorProgressHandler
type CTFontDescriptorProgressHandler = func(CTFontDescriptorMatchingState, corefoundation.CFDictionaryRef) bool

// CTFontDescriptorRef is a font descriptor.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontDescriptor
type CTFontDescriptorRef uintptr

// CTFontPriority is the priority of font descriptors when resolving duplicates and sorting match results.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontPriority
type CTFontPriority = uint32

// CTFontRef is a font object.
//
// See: https://developer.apple.com/documentation/CoreText/CTFont
type CTFontRef uintptr

// CTFontTableTag is font table tags provide access to font table data.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontTableTag
type CTFontTableTag = uint32

// CTFrameRef is a frame.
//
// See: https://developer.apple.com/documentation/CoreText/CTFrame
type CTFrameRef uintptr

// CTFramesetterRef is generate text frames.
//
// See: https://developer.apple.com/documentation/CoreText/CTFramesetter
type CTFramesetterRef uintptr

// CTGlyphInfoRef is override a font’s specified mapping from Unicode to the glyph ID.
//
// See: https://developer.apple.com/documentation/CoreText/CTGlyphInfo
type CTGlyphInfoRef uintptr

// CTLineRef is a line of text.
//
// See: https://developer.apple.com/documentation/CoreText/CTLine
type CTLineRef uintptr

// CTMutableFontCollectionRef is a reference to a mutable font collection.
//
// See: https://developer.apple.com/documentation/CoreText/CTMutableFontCollection
type CTMutableFontCollectionRef uintptr

// CTParagraphStyleRef is paragraph or ruler attributes in an attributed string.
//
// See: https://developer.apple.com/documentation/CoreText/CTParagraphStyle
type CTParagraphStyleRef uintptr

// See: https://developer.apple.com/documentation/CoreText/CTRubyAnnotation
type CTRubyAnnotationRef uintptr

// CTRunDelegateDeallocateCallback is defines a pointer to a function that is invoked when a CTRunDelegate object is deallocated.
//
// See: https://developer.apple.com/documentation/CoreText/CTRunDelegateDeallocateCallback
type CTRunDelegateDeallocateCallback = func(refCon unsafe.Pointer)

// CTRunDelegateGetAscentCallback is defines a pointer to a function that determines typographic ascent of glyphs in the run.
//
// See: https://developer.apple.com/documentation/CoreText/CTRunDelegateGetAscentCallback
type CTRunDelegateGetAscentCallback = func(refCon unsafe.Pointer) float64

// CTRunDelegateGetDescentCallback is defines a pointer to a function that determines typographic descent of glyphs in the run.
//
// See: https://developer.apple.com/documentation/CoreText/CTRunDelegateGetDescentCallback
type CTRunDelegateGetDescentCallback = func(refCon unsafe.Pointer) float64

// CTRunDelegateGetWidthCallback is defines a pointer to a function that determines the typographic width of glyphs in the run.
//
// See: https://developer.apple.com/documentation/CoreText/CTRunDelegateGetWidthCallback
type CTRunDelegateGetWidthCallback = func(refCon unsafe.Pointer) float64

// CTRunDelegateRef is a run delegate.
//
// See: https://developer.apple.com/documentation/CoreText/CTRunDelegate
type CTRunDelegateRef uintptr

// CTRunRef is a glyph run.
//
// See: https://developer.apple.com/documentation/CoreText/CTRun
type CTRunRef uintptr

// CTTextTabRef is a tab in a paragraph style, storing an alignment type and location.
//
// See: https://developer.apple.com/documentation/CoreText/CTTextTab
type CTTextTabRef uintptr

// CTTypesetterRef is a typesetter which performs line layout.
//
// See: https://developer.apple.com/documentation/CoreText/CTTypesetter
type CTTypesetterRef uintptr

// See: https://developer.apple.com/documentation/CoreText/FontLanguageCode
type FontLanguageCode = uint32

// See: https://developer.apple.com/documentation/CoreText/FontNameCode
type FontNameCode = uint32

// See: https://developer.apple.com/documentation/CoreText/FontPlatformCode
type FontPlatformCode = uint32

// See: https://developer.apple.com/documentation/CoreText/FontScriptCode
type FontScriptCode = uint32

// See: https://developer.apple.com/documentation/CoreText/JustPCActionType
type JustPCActionType = uint16

// See: https://developer.apple.com/documentation/CoreText/JustPCUnconditionalAddAction
type JustPCUnconditionalAddAction = uint16

// See: https://developer.apple.com/documentation/CoreText/JustificationFlags
type JustificationFlags = uint16

// See: https://developer.apple.com/documentation/CoreText/KernArrayOffset
type KernArrayOffset = uint16

// See: https://developer.apple.com/documentation/CoreText/KernKerningValue
type KernKerningValue = int16

// See: https://developer.apple.com/documentation/CoreText/KernOffsetTablePtr
type KernOffsetTablePtr = *KernOffsetTable

// See: https://developer.apple.com/documentation/CoreText/KernOrderedListEntryPtr
type KernOrderedListEntryPtr = *KernOrderedListEntry

// See: https://developer.apple.com/documentation/CoreText/KernSubtableHeaderPtr
type KernSubtableHeaderPtr = *KernSubtableHeader

// See: https://developer.apple.com/documentation/CoreText/KernSubtableInfo
type KernSubtableInfo = uint16

// See: https://developer.apple.com/documentation/CoreText/KernTableFormat
type KernTableFormat = uint8

// See: https://developer.apple.com/documentation/CoreText/KernTableHeaderHandle
type KernTableHeaderHandle = *KernTableHeaderPtr

// See: https://developer.apple.com/documentation/CoreText/KernTableHeaderPtr
type KernTableHeaderPtr = *KernTableHeader

// See: https://developer.apple.com/documentation/CoreText/KerxArrayOffset
type KerxArrayOffset = uint32

// See: https://developer.apple.com/documentation/CoreText/KerxOrderedListEntryPtr
type KerxOrderedListEntryPtr = *KerxOrderedListEntry

// See: https://developer.apple.com/documentation/CoreText/KerxSubtableCoverage
type KerxSubtableCoverage = uint32

// See: https://developer.apple.com/documentation/CoreText/KerxSubtableHeaderPtr
type KerxSubtableHeaderPtr = *KerxSubtableHeader

// See: https://developer.apple.com/documentation/CoreText/KerxTableHeaderHandle
type KerxTableHeaderHandle = *KerxTableHeaderPtr

// See: https://developer.apple.com/documentation/CoreText/KerxTableHeaderPtr
type KerxTableHeaderPtr = *KerxTableHeader

// See: https://developer.apple.com/documentation/CoreText/LcarCaretTablePtr
type LcarCaretTablePtr = *LcarCaretTable

// See: https://developer.apple.com/documentation/CoreText/MortLigatureActionEntry
type MortLigatureActionEntry = uint32

// See: https://developer.apple.com/documentation/CoreText/MortSubtableMaskFlags
type MortSubtableMaskFlags = uint32

// See: https://developer.apple.com/documentation/CoreText/OpbdTableFormat
type OpbdTableFormat = uint16

// See: https://developer.apple.com/documentation/CoreText/PropCharProperties
type PropCharProperties = uint16

// See: https://developer.apple.com/documentation/CoreText/SFNTLookupKind
type SFNTLookupKind = uint32

// See: https://developer.apple.com/documentation/CoreText/SFNTLookupOffset
type SFNTLookupOffset = uint16

// See: https://developer.apple.com/documentation/CoreText/SFNTLookupTableFormat
type SFNTLookupTableFormat = uint16

// See: https://developer.apple.com/documentation/CoreText/SFNTLookupTableHandle
type SFNTLookupTableHandle = *SFNTLookupTablePtr

// See: https://developer.apple.com/documentation/CoreText/SFNTLookupTablePtr
type SFNTLookupTablePtr = *SFNTLookupTable

// See: https://developer.apple.com/documentation/CoreText/SFNTLookupValue
type SFNTLookupValue = uint16

// See: https://developer.apple.com/documentation/CoreText/STClass
type STClass = uint8

// See: https://developer.apple.com/documentation/CoreText/STEntryIndex
type STEntryIndex = uint8

// See: https://developer.apple.com/documentation/CoreText/STXClass
type STXClass = uint16

// See: https://developer.apple.com/documentation/CoreText/STXClassTable
type STXClassTable = SFNTLookupTable

// See: https://developer.apple.com/documentation/CoreText/STXEntryIndex
type STXEntryIndex = uint16

// See: https://developer.apple.com/documentation/CoreText/STXStateIndex
type STXStateIndex = uint16

// See: https://developer.apple.com/documentation/CoreText/TrakValue
type TrakValue = int16
