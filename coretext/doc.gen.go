// Code generated from Apple documentation for CoreText. DO NOT EDIT.

// Package coretext provides Go bindings for the CoreText framework.
//
// Create text layouts, optimize font handling, and access font metrics and
// glyph data.
//
// Core Text provides a low-level programming interface for laying out text
// and handling fonts. The Core Text layout engine is designed for high
// performance, ease of use, and close integration with Core Foundation. The
// text layout API provides high-quality typesetting, including
// character-to-glyph conversion, with ligatures, kerning, and so on. The
// complementary Core Text font technology provides automatic font
// substitution (cascading), font descriptors and collections, easy access to
// font metrics and glyph data, and many other features.
//
// # Opaque Types
//
//   - [CTFontRef]: A font object. ([CTFontUIFontType], [CTFontTableTag], [CTFontTableOptions], [CTFontOptions])
//   - [CTFontCollectionRef]: A font collection. ([CTFontCollectionSortDescriptorsCallback], [CTFontCollectionCopyOptions])
//   - [CTFontDescriptorRef]: A font descriptor. ([CTFontOrientation], [CTFontFormat], [CTFontPriority], [CTFontSymbolicTraits], [CTFontStylisticClass])
//   - [CTFrameRef]: A frame. ([CTFramePathFillRule], [CTFrameProgression])
//   - [CTFramesetterRef]: Generate text frames.
//   - [CTGlyphInfoRef]: Override a font’s specified mapping from Unicode to the glyph ID. ([CTCharacterCollection])
//   - [CTLineRef]: A line of text. ([CTLineTruncationType])
//   - [CTParagraphStyleRef]: Paragraph or ruler attributes in an attributed string. ([CTParagraphStyleSetting], [CTTextAlignment], [CTLineBreakMode], [CTWritingDirection], [CTParagraphStyleSpecifier])
//   - [CTRunRef]: A glyph run. ([CTRunStatus])
//   - [CTRunDelegateRef]: A run delegate. ([CTRunDelegateGetAscentCallback], [CTRunDelegateGetDescentCallback], [CTRunDelegateGetWidthCallback], [CTRunDelegateDeallocateCallback], [CTRunDelegateCallbacks])
//   - [CTTextTabRef]: A tab in a paragraph style, storing an alignment type and location.
//   - [CTTypesetterRef]: A typesetter which performs line layout.
//
// # Protocols
//
//   - [CTAdaptiveImageProviding]
//
// # Variables
//
//   - [KCTFontDescriptorLanguageAttribute]
//
// # Functions
//
//   - [CTFontGetUIFontType]
//
// # Type Aliases
//
//   - [CTRubyAnnotationRef]
package coretext

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the CoreText library,
// in order. Frameworks whose symbols live in a known dylib resolve to that
// dylib alone; the rest try the framework bundle first and then a /usr/lib
// dylib fallback, which covers C-API frameworks that are not in the dyld
// shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/CoreText.framework/CoreText",
	"/usr/lib/libCoreText.dylib",
}

// frameworkHandle is the handle to the loaded framework.
var frameworkHandle uintptr

func init() {
	for _, path := range frameworkPaths {
		h, err := purego.Dlopen(path, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
		if err == nil {
			frameworkHandle = h
			return
		}
	}
	// Loading is best-effort: the warning is silent by default because a missing
	// framework is harmless unless one of its symbols is actually called. Set
	// APPLE_FRAMEWORK_LOAD_DEBUG to surface load failures while diagnosing.
	if os.Getenv("APPLE_FRAMEWORK_LOAD_DEBUG") != "" {
		fmt.Fprintf(os.Stderr, "warning: CoreText: failed to load framework from any known path\n")
	}
}
