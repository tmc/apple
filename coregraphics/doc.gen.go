// Code generated from Apple documentation for CoreGraphics. DO NOT EDIT.

// Package coregraphics provides Go bindings for the CoreGraphics framework.
//
// Harness the power of Quartz technology to perform lightweight 2D rendering
// with high-fidelity output. Handle path-based drawing, antialiased
// rendering, gradients, images, color management, PDF documents, and more.
//
// The Core Graphics framework is based on the Quartz advanced drawing engine.
// It provides low-level, lightweight 2D rendering with unmatched output
// fidelity. You use this framework to handle path-based drawing,
// transformations, color management, offscreen rendering, patterns, gradients
// and shadings, image data management, image creation, and image masking, as
// well as PDF document creation, display, and parsing.
//
// # Opaque Types
//
//   - [CGContextRef]: A Quartz 2D drawing environment. ([CGBitmapContextReleaseDataCallback], [CGPathDrawingMode], [CGInterpolationQuality], [CGGradientDrawingOptions], [CGTextDrawingMode])
//   - [CGColorRef]: A set of components that define a color, with a color space specifying how to interpret them.
//   - [CGColorConversionInfoRef]: An object that describes how to convert between color spaces for use by other system services. ([CGColorConversionInfoTransformType])
//   - [CGColorSpaceRef]: A profile that specifies how to interpret a color value for display. ([CGColorSpaceModel], [CGColorRenderingIntent])
//   - [CGDataConsumerRef]: An abstraction for data-writing tasks that eliminates the need to manage a raw memory buffer. ([CGDataConsumerCallbacks], [CGDataConsumerPutBytesCallback], [CGDataConsumerReleaseInfoCallback])
//   - [CGDataProviderRef]: An abstraction for data-reading tasks that eliminates the need to manage a raw memory buffer. ([CGDataProviderSequentialCallbacks], [CGDataProviderRewindCallback], [CGDataProviderGetBytesCallback], [CGDataProviderSkipForwardCallback], [CGDataProviderReleaseInfoCallback])
//   - [CGFontRef]: A set of character glyphs and layout information for drawing text. ([CGFontPostScriptFormat], [CGGlyph], [CGFontIndex])
//   - [CGFunctionRef]: A general facility for defining and using callback functions. ([CGFunctionCallbacks], [CGFunctionEvaluateCallback], [CGFunctionReleaseInfoCallback])
//   - [CGGradientRef]: A definition for a smooth transition between colors for drawing radial and axial gradient fills.
//   - [CGImageRef]: A bitmap image or image mask. ([CGImageAlphaInfo], [CGBitmapInfo], [CGImageAlphaInfo], [CGBitmapInfo])
//   - [CGLayerRef]: An offscreen context for reusing content drawn with Core Graphics.
//   - [CGPathRef]: An immutable graphics path: a mathematical description of shapes or lines to be drawn in a graphics context. ([CGPathApplierFunction], [CGPathElement], [CGPathElementType])
//   - [CGPatternRef]: A 2D pattern to be used for drawing graphics paths. ([CGPatternCallbacks], [CGPatternDrawPatternCallback], [CGPatternReleaseInfoCallback], [CGPatternTiling])
//   - CGPDFArray: An array structure within a PDF document. ([CGPDFArrayRef])
//   - CGPDFContentStream: A representation of one or more content data streams in a PDF page. ([CGPDFContentStreamRef])
//   - CGPDFDictionary: A dictionary structure within a PDF document. ([CGPDFDictionaryApplierFunction], [CGPDFDictionaryRef])
//   - [CGPDFDocumentRef]: A document that contains PDF (Portable Document Format) drawing information.
//   - CGPDFObject: An object representing content within a PDF document. ([CGPDFObjectRef], [CGPDFBoolean], [CGPDFInteger], [CGPDFReal], [CGPDFObjectType])
//   - CGPDFOperatorTable: A set of callback functions for operators used when scanning content in a PDF document. ([CGPDFOperatorCallback], [CGPDFOperatorTableRef])
//   - [CGPDFPageRef]: A type that represents a page in a PDF document. ([CGPDFBox])
//   - CGPDFScanner: A parser object for handling content and operators in a PDF content stream. ([CGPDFScannerRef])
//   - CGPDFStream: A stream or sequence of data bytes in a PDF document. ([CGPDFStreamRef], [CGPDFDataFormat])
//   - CGPDFString: A text string in a PDF document. ([CGPDFStringRef])
//   - [CGPSConverterRef]: An opaque data type used to convert PostScript data to PDF data.
//   - [CGShadingRef]: A definition for a smooth transition between colors, controlled by a custom function you provide, for drawing radial and axial gradient fills.
//
// # Services
//
//   - [Quartz Display Services]: Provides direct access to features in the macOS window server for configuring and controlling display hardware. ([CGDisplayReconfigurationCallBack], [CGScreenRefreshCallback], [CGScreenUpdateMoveCallback], [CGDirectDisplayID], [CGDisplayBlendFraction])
//   - [Quartz Event Services]: Provides features for managing —filters for observing and altering the stream of low-level user input events in macOS. ([CGEventTapCallBack], [CGButtonCount], [CGCharCode], [CGEventMask], [CGEventSourceKeyboardType])
//   - [Quartz Window Services]: Provides information about the windows managed by the macOS window server. ([CGWindowID], [CGWindowListOption], [CGWindowImageOption], [CGWindowSharingType], [CGWindowBackingType])
//
// # Variables
//
//   - [KCGAdaptiveMaximumBitDepth]
//   - [KCGContentAverageLightLevel]
//   - [KCGContentAverageLightLevelNits]
//   - [KCGDynamicRangeConstrained]
//   - [KCGDynamicRangeHigh]
//   - [KCGDynamicRangeStandard]
//   - [KCGPreferredDynamicRange]
//
// # Functions
//
//   - [CGBitmapContextCreateAdaptive]
//   - [CGBitmapInfoMake]
//   - [CGColorGetContentHeadroom]
//   - [CGColorCreateWithContentHeadroom]
//   - [CGBitmapContextCreate]
//   - [CGBitmapContextCreateWithData]
//   - [CGContextSynchronizeAttributes]
//   - [CGContextGetContentToneMappingInfo]
//   - [CGContextSetContentToneMappingInfo]
//   - [CGEXRToneMappingGammaGetDefaultOptions]
//   - [CGGradientGetContentHeadroom]
//   - [CGGradientCreateWithContentHeadroom]
//   - [CGRenderingBufferLockBytePtr]
//   - [CGRenderingBufferProviderCreate]
//   - [CGRenderingBufferProviderCreateWithCFData]
//   - [CGRenderingBufferProviderGetSize]
//   - [CGRenderingBufferProviderGetTypeID]
//   - [CGRenderingBufferUnlockBytePtr]
//   - [CGShadingGetContentHeadroom]
//   - [CGShadingCreateAxialWithContentHeadroom]
//   - [CGShadingCreateRadialWithContentHeadroom]
//
// # Type Aliases
//
//   - [CGRenderingBufferProviderRef]
//
// # Enumerations
//
//   - [CGBitmapLayout]
//   - [CGColorModel]
//   - [CGComponent]
//   - [CGImageComponentInfo]//
//
// [Quartz Display Services]: https://developer.apple.com/documentation/coregraphics/quartz-display-services
// [Quartz Event Services]: https://developer.apple.com/documentation/coregraphics/quartz-event-services
// [Quartz Window Services]: https://developer.apple.com/documentation/coregraphics/quartz-window-services
package coregraphics

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the CoreGraphics library,
// in order. Frameworks whose symbols live in a known dylib resolve to that
// dylib alone; the rest try the framework bundle first and then a /usr/lib
// dylib fallback, which covers C-API frameworks that are not in the dyld
// shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/CoreGraphics.framework/CoreGraphics",
	"/usr/lib/libCoreGraphics.dylib",
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
		fmt.Fprintf(os.Stderr, "warning: CoreGraphics: failed to load framework from any known path\n")
	}
}
