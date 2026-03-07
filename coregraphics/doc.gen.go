
// Code generated from Apple documentation for CoreGraphics. DO NOT EDIT.

// Package coregraphics provides Go bindings for the CoreGraphics framework.
//
// Harness the power of Quartz technology to perform lightweight 2D rendering with high-fidelity output. Handle path-based drawing, antialiased rendering, gradients, images, color management, PDF documents, and more.
//
// The Core Graphics framework is based on the Quartz advanced drawing engine. It provides low-level, lightweight 2D rendering with unmatched output fidelity. You use this framework to handle path-based drawing, transformations, color management, offscreen rendering, patterns, gradients and shadings, image data management, image creation, and image masking, as well as PDF document creation, display, and parsing.
//
// # Geometric Data Types
//
//   - CGFloat: The basic type for floating-point scalar values in Core Graphics and related frameworks.
//   - CGPoint
//   - CGSize: A structure that contains width and height values.
//   - CGRect
//   - CGVector: A structure that contains a two-dimensional vector.
//   - CGAffineTransform
//
// # 2D Drawing
//
//   - CGContext: A Quartz 2D drawing environment. ([CGBitmapContextReleaseDataCallback], [CGPathDrawingMode], [CGInterpolationQuality], [CGGradientDrawingOptions], [CGTextDrawingMode])
//   - CGImage: A bitmap image or image mask. ([CGImageAlphaInfo], [CGBitmapInfo], [CGImageAlphaInfo], [CGBitmapInfo])
//   - CGPath: An immutable graphics path: a mathematical description of shapes or lines to be drawn in a graphics context. ([CGPathApplierFunction], [CGPathElement], [CGPathElementType])
//   - CGMutablePath: A mutable graphics path: a mathematical description of shapes or lines to be drawn in a graphics context.
//   - CGLayer: An offscreen context for reusing content drawn with Core Graphics.
//
// # Colors and Fonts
//
//   - CGColor: A set of components that define a color, with a color space specifying how to interpret them.
//   - CGColorConversionInfo: An object that describes how to convert between color spaces for use by other system services. ([CGColorConversionInfoTransformType])
//   - CGColorSpace: A profile that specifies how to interpret a color value for display. ([CGColorSpaceModel], [CGColorRenderingIntent])
//   - CGFont: A set of character glyphs and layout information for drawing text. ([CGFontPostScriptFormat], [CGGlyph], [CGFontIndex])
//
// # Working with PDF Documents
//
//   - CGPDFDocument: A document that contains PDF (Portable Document Format) drawing information.
//
// # Utility and Support Classes
//
//   - CGDataConsumer: An abstraction for data-writing tasks that eliminates the need to manage a raw memory buffer. ([CGDataConsumerCallbacks], [CGDataConsumerPutBytesCallback], [CGDataConsumerReleaseInfoCallback])
//   - CGDataProvider: An abstraction for data-reading tasks that eliminates the need to manage a raw memory buffer. ([CGDataProviderSequentialCallbacks], [CGDataProviderRewindCallback], [CGDataProviderGetBytesCallback], [CGDataProviderSkipForwardCallback], [CGDataProviderReleaseInfoCallback])
//   - CGShading: A definition for a smooth transition between colors, controlled by a custom function you provide, for drawing radial and axial gradient fills.
//   - CGGradient: A definition for a smooth transition between colors for drawing radial and axial gradient fills.
//   - CGFunction: A general facility for defining and using callback functions. ([CGFunctionCallbacks], [CGFunctionEvaluateCallback], [CGFunctionReleaseInfoCallback])
//   - CGPattern: A 2D pattern to be used for drawing graphics paths. ([CGPatternCallbacks], [CGPatternDrawPatternCallback], [CGPatternReleaseInfoCallback], [CGPatternTiling])
//
// # Services
//
//   - Quartz Display Services: Provides direct access to features in the macOS window server for configuring and controlling display hardware. ([CGDisplayReconfigurationCallBack], [CGScreenRefreshCallback], [CGScreenUpdateMoveCallback], [CGDirectDisplayID], [CGDisplayBlendFraction])
//   - Quartz Event Services: Provides features for managing —filters for observing and altering the stream of low-level user input events in macOS. ([CGEventTapCallBack], [CGButtonCount], [CGCharCode], [CGEventMask], [CGEventSourceKeyboardType])
//   - Quartz Window Services: Provides information about the windows managed by the macOS window server. ([CGWindowID], [CGWindowListOption], [CGWindowImageOption], [CGWindowSharingType], [CGWindowBackingType])
//
// # Classes
//
//   - CGRenderingBufferProvider
//
// # Enumerations
//
//   - CGBitmapLayout
//   - CGComponent
//   - CGContentToneMappingInfo
//   - CGImageComponentInfo
//
// [CoreGraphics Documentation]: https://developer.apple.com/documentation/CoreGraphics
package coregraphics

import (
	"fmt"
	"os"
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/CoreGraphics.framework/CoreGraphics"
// frameworkHandle is the handle to the loaded framework.
var frameworkHandle uintptr

func init() {
	var err error
	frameworkHandle, err = purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: CoreGraphics: failed to load %s: %v\n", frameworkPath, err)
		return 
	}
}

