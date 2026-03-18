// Code generated from Apple documentation for CoreGraphics. DO NOT EDIT.

package coregraphics
import (
	"github.com/tmc/apple/corefoundation"
)

// C struct types
// CGBitmapParameters
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBitmapParameters-1cm7j
type CGBitmapParameters struct {
	AlignedBytesPerRow uintptr
	ByteOrder int32
	BytesPerPixel uintptr
	ColorSpace CGColorSpaceRef
	Component CGComponent
	EdrTargetHeadroom float32
	Format CGImagePixelFormatInfo
	HasPremultipliedAlpha bool
	Height uintptr
	Layout CGBitmapLayout
	Width uintptr

}

// CGColorBufferFormat
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorBufferFormat
type CGColorBufferFormat struct {
	Version uint32
	BitmapInfo CGBitmapInfo
	BitsPerComponent uintptr
	BitsPerPixel uintptr
	BytesPerRow uintptr

}

// CGColorDataFormat
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorDataFormat
type CGColorDataFormat struct {
	Version uint32
	Colorspace_info corefoundation.CFTypeRef
	Bitmap_info CGBitmapInfo
	Bits_per_component uintptr
	Bytes_per_row uintptr
	Intent CGColorRenderingIntent
	Decode *float64

}

// CGContentInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContentInfo
type CGContentInfo struct {
	DeepestImageComponent CGComponent
	ContentColorModels CGColorModel
	HasWideGamut bool
	HasTransparency bool
	LargestContentHeadroom float32

}

// CGContentToneMappingInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContentToneMappingInfo-c.struct
type CGContentToneMappingInfo struct {
	Method CGToneMapping
	Options corefoundation.CFDictionaryRef

}

// CGDataConsumerCallbacks - A structure that contains pointers to callback functions that manage the copying of data for a data consumer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDataConsumerCallbacks
type CGDataConsumerCallbacks struct {
	PutBytes CGDataConsumerPutBytesCallback // A pointer to a function that copies data to the data consumer. For more information, see [CGDataConsumerPutBytesCallback](<doc://com.apple.coregraphics/documentation/CoreGraphics/CGDataConsumerPutBytesCallback>).
	ReleaseConsumer CGDataConsumerReleaseInfoCallback // A pointer to a function that handles clean-up for the data consumer, or [NULL].

}

// CGDataProviderDirectCallbacks - Defines pointers to client-defined callback functions that manage the sending of data for a direct-access data provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDataProviderDirectCallbacks
type CGDataProviderDirectCallbacks struct {
	Version uint // The version of this structure. It should be set to 0.
	GetBytePointer CGDataProviderGetBytePointerCallback // A pointer to a function that returns a pointer to the provider’s data. For more information, see [CGDataProviderGetBytePointerCallback](<doc://com.apple.coregraphics/documentation/CoreGraphics/CGDataProviderGetBytePointerCallback>).
	ReleaseBytePointer CGDataProviderReleaseBytePointerCallback // A pointer to a function that Core Graphics calls to release a pointer to the provider’s data. For more information, see [CGDataProviderReleaseBytePointerCallback](<doc://com.apple.coregraphics/documentation/CoreGraphics/CGDataProviderReleaseBytePointerCallback>).
	GetBytesAtPosition CGDataProviderGetBytesAtPositionCallback // A pointer to a function that copies data from the provider.
	ReleaseInfo CGDataProviderReleaseInfoCallback // A pointer to a function that handles clean-up for the data provider, or [NULL]. For more information, see [CGDataProviderReleaseInfoCallback](<doc://com.apple.coregraphics/documentation/CoreGraphics/CGDataProviderReleaseInfoCallback>).

}

// CGDataProviderSequentialCallbacks - Defines a structure containing pointers to client-defined callback functions that manage the sending of data for a sequential-access data provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDataProviderSequentialCallbacks
type CGDataProviderSequentialCallbacks struct {
	Version uint // The version of this structure. It should be set to 0.
	GetBytes CGDataProviderGetBytesCallback // A pointer to a function that copies data from the provider. For more information, see [CGDataProviderGetBytesCallback](<doc://com.apple.coregraphics/documentation/CoreGraphics/CGDataProviderGetBytesCallback>).
	SkipForward CGDataProviderSkipForwardCallback // A pointer to a function that Core Graphics calls to advance the stream of data supplied by the provider.
	Rewind CGDataProviderRewindCallback // A pointer to a function Core Graphics calls to return the provider to the beginning of the data stream. For more information, see [CGDataProviderRewindCallback](<doc://com.apple.coregraphics/documentation/CoreGraphics/CGDataProviderRewindCallback>).
	ReleaseInfo CGDataProviderReleaseInfoCallback // A pointer to a function that handles clean-up for the data provider, or [NULL]. For more information, see [CGDataProviderReleaseInfoCallback](<doc://com.apple.coregraphics/documentation/CoreGraphics/CGDataProviderReleaseInfoCallback>).

}

// CGDeviceColor
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDeviceColor
type CGDeviceColor struct {
	Red float32
	Green float32
	Blue float32

}

// CGFunctionCallbacks - A structure that contains callbacks needed by a [CGFunctionRef] object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGFunctionCallbacks
type CGFunctionCallbacks struct {
	Version uint // The structure version number. For this structure,the version should be `0`.
	Evaluate CGFunctionEvaluateCallback // The callback that evaluates the function.
	ReleaseInfo CGFunctionReleaseInfoCallback // If non-[NULL],the callback used to release the `info` parameterpassed to [init(info:domainDimension:domain:rangeDimension:range:callbacks:)](<doc://com.apple.coregraphics/documentation/CoreGraphics/CGFunction/init(info:domainDimension:domain:rangeDimension:range:callbacks:)>).

}

// CGPSConverterCallbacks - A structure for holding the callbacks provided when you create a PostScript converter object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPSConverterCallbacks
type CGPSConverterCallbacks struct {
	Version uint // The version number of the structure passed in as a parameter to the converter creation functions. The structure defined below is version `0`.
	BeginDocument CGPSConverterBeginDocumentCallback // The callback called at the beginning of the conversion of the PostScript document, or [NULL].
	EndDocument CGPSConverterEndDocumentCallback // The callback called at the end of conversion of the PostScript document, or [NULL].
	BeginPage CGPSConverterBeginPageCallback // The callback called at the start of the conversion of each page in the PostScript document, or [NULL].
	EndPage CGPSConverterEndPageCallback // The callback called at the end of the conversion of each page in the PostScript document, or [NULL].
	NoteProgress CGPSConverterProgressCallback // The callback called periodically during the conversion to indicate that conversion is proceeding, or [NULL].
	NoteMessage CGPSConverterMessageCallback // The callback called to pass any messages that might result during the conversion, or [NULL].
	ReleaseInfo CGPSConverterReleaseInfoCallback // The callback called when the converter is deallocated, or [NULL].

}

// CGPathElement - A data structure that provides information about a path element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPathElement
type CGPathElement struct {
	Type CGPathElementType // An element type (or operation).
	Points *corefoundation.CGPoint // An array of one or more points that serve as arguments.

}

// CGPatternCallbacks - A structure that holds a version and two callback functions for drawing a custom pattern.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPatternCallbacks
type CGPatternCallbacks struct {
	Version uint // The version of the structure passed in as a parameterto the [init(info:bounds:matrix:xStep:yStep:tiling:isColored:callbacks:)](<doc://com.apple.coregraphics/documentation/CoreGraphics/CGPattern/init(info:bounds:matrix:xStep:yStep:tiling:isColored:callbacks:)>). Forthis version of the structure, you should set this value to zero.
	DrawPattern CGPatternDrawPatternCallback // A pointer to a custom function that draws thepattern. For information about this callback function, see [CGPatternDrawPatternCallback](<doc://com.apple.coregraphics/documentation/CoreGraphics/CGPatternDrawPatternCallback>).
	ReleaseInfo CGPatternReleaseInfoCallback // An optional pointer to a custom function that’sinvoked when the pattern is released. [CGPatternReleaseInfoCallback](<doc://com.apple.coregraphics/documentation/CoreGraphics/CGPatternReleaseInfoCallback>).

}

// CGScreenUpdateMoveDelta - The distance, in pixel units, that an onscreen region moves.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGScreenUpdateMoveDelta
type CGScreenUpdateMoveDelta struct {
	DX int32
	DY int32

}

