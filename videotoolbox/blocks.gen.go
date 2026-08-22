// Code generated from Apple documentation. DO NOT EDIT.

package videotoolbox

import (
	"unsafe"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/corevideo"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// ErrorHandler is the signature for a completion handler block.
//
// Used by:
//   - [VTSuperResolutionScalerConfiguration.DownloadConfigurationModelWithCompletionHandler]
type ErrorHandler = func(error)

// NewErrorBlock wraps a Go [ErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [VTSuperResolutionScalerConfiguration.DownloadConfigurationModelWithCompletionHandler]
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

// VTCompressionOutputHandler handles A callback for the system to invoke when it’s finished compressing a frame.

// NewVTCompressionOutputHandlerBlock wraps a Go [VTCompressionOutputHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewVTCompressionOutputHandlerBlock(handler VTCompressionOutputHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive int32, extra0 uint, extra1 unsafe.Pointer) {
		handler(primitive, extra0, extra1)
	})
	return objc.ID(block), func() { block.Release() }
}

// VTDecompressionMultiImageCapableOutputHandler handles A type alias for callback that the system invokes when it finishes decompressing a frame.

// VTDecompressionOutputHandler handles The prototype for the block invoked when frame decompression is complete.

// VTFrameProcessorParametersCMTimeBoolErrorHandler handles This frame output handler is called once for each destination frame in the provided parameters if no errors are encountered.
//
// Used by:
//   - [VTFrameProcessor.ProcessWithParametersFrameOutputHandler]
type VTFrameProcessorParametersCMTimeBoolErrorHandler = func(VTFrameProcessorParameters, coremedia.CMTime, bool, error)

// VTFrameProcessorParametersErrorHandler handles This completion handler is called when the frame processing is completed.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [VTFrameProcessor.ProcessWithParametersCompletionHandler]
type VTFrameProcessorParametersErrorHandler = func(VTFrameProcessorParameters, error)

// NewVTFrameProcessorParametersErrorBlock wraps a Go [VTFrameProcessorParametersErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [VTFrameProcessor.ProcessWithParametersCompletionHandler]
func NewVTFrameProcessorParametersErrorBlock(handler VTFrameProcessorParametersErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result VTFrameProcessorParameters
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			result = VTFrameProcessorParametersObjectFromID(resultID)
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// VTMotionEstimationOutputHandler handles A block invoked by motion-estimation session when frame processing is complete.

// NewVTMotionEstimationOutputHandlerBlock wraps a Go [VTMotionEstimationOutputHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewVTMotionEstimationOutputHandlerBlock(handler VTMotionEstimationOutputHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive int32, extra0 uint, extra1 corefoundation.CFDictionaryRef, extra2 corevideo.CVImageBufferRef) {
		handler(primitive, extra0, extra1, extra2)
	})
	return objc.ID(block), func() { block.Release() }
}

// VTRAWProcessingOutputHandler handles A block the system calls when frame processing is complete.

// NewVTRAWProcessingOutputHandlerBlock wraps a Go [VTRAWProcessingOutputHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewVTRAWProcessingOutputHandlerBlock(handler VTRAWProcessingOutputHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive int32, extra0 corevideo.CVImageBufferRef) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// VTRAWProcessingParameterChangeHandler handles A function the system calls when processing parameters change.

// NewVTRAWProcessingParameterChangeHandlerBlock wraps a Go [VTRAWProcessingParameterChangeHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewVTRAWProcessingParameterChangeHandlerBlock(handler VTRAWProcessingParameterChangeHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal corefoundation.CFArrayRef) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}
