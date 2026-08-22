// Code generated from Apple documentation. DO NOT EDIT.

package mediaextension

import (
	"github.com/tmc/apple/avfoundation"
	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/corevideo"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// AVMetadataItemArrayErrorHandler handles The completion block to execute when the load operation finishes.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [MEFormatReader.LoadMetadataWithCompletionHandler]
//   - [METrackReader.LoadMetadataWithCompletionHandler]
type AVMetadataItemArrayErrorHandler = func(*[]avfoundation.AVMetadataItem, error)

// NewAVMetadataItemArrayErrorBlock wraps a Go [AVMetadataItemArrayErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [MEFormatReader.LoadMetadataWithCompletionHandler]
//   - [METrackReader.LoadMetadataWithCompletionHandler]
func NewAVMetadataItemArrayErrorBlock(handler AVMetadataItemArrayErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *[]avfoundation.AVMetadataItem
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			obj := foundation.NSArrayFromID(resultID)
			count := obj.Count()
			res := make([]avfoundation.AVMetadataItem, count)
			for i := uint(0); i < count; i++ {
				item := obj.ObjectAtIndex(i)
				res[i] = avfoundation.AVMetadataItemFromID(item.GetID())
			}
			result = &res
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// CMSampleBufferRefErrorHandler handles The completion block to execute when the load operation finishes.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [MESampleCursor.LoadSampleBufferContainingSamplesToEndCursorCompletionHandler]
type CMSampleBufferRefErrorHandler = func(coremedia.CMSampleBufferRef, error)

// NewCMSampleBufferRefErrorBlock wraps a Go [CMSampleBufferRefErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [MESampleCursor.LoadSampleBufferContainingSamplesToEndCursorCompletionHandler]
func NewCMSampleBufferRefErrorBlock(handler CMSampleBufferRefErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal coremedia.CMSampleBufferRef, errID objc.ID) {
		handler(primitiveVal, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// CMTimeBoolErrorHandler handles The completion block to execute when the move operation finishes.
//
// Used by:
//   - [MESampleCursor.StepByDecodeTimeCompletionHandler]
//   - [MESampleCursor.StepByPresentationTimeCompletionHandler]
type CMTimeBoolErrorHandler = func(bool, error)

// CMTimeErrorHandler handles The completion block to execute when the load operation finishes.
//
// Used by:
//   - [METrackReader.LoadUneditedDurationWithCompletionHandler]
type CMTimeErrorHandler = func(error)

// CVImageBufferRefMEDecodeFrameStatusErrorHandler handles The completion block to execute when the decode operation finishes.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [MEVideoDecoder.DecodeFrameFromSampleBufferOptionsCompletionHandler]
type CVImageBufferRefMEDecodeFrameStatusErrorHandler = func(corevideo.CVImageBufferRef, MEDecodeFrameStatus, error)

// NewCVImageBufferRefMEDecodeFrameStatusErrorBlock wraps a Go [CVImageBufferRefMEDecodeFrameStatusErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [MEVideoDecoder.DecodeFrameFromSampleBufferOptionsCompletionHandler]
func NewCVImageBufferRefMEDecodeFrameStatusErrorBlock(handler CVImageBufferRefMEDecodeFrameStatusErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive corevideo.CVImageBufferRef, extra0 MEDecodeFrameStatus, errID objc.ID) {
		handler(primitive, extra0, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// CVPixelBufferRefErrorHandler handles The handler is invoked when a frame processes and is ready to be sent back to the caller.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [MERAWProcessor.ProcessFrameFromImageBufferCompletionHandler]
type CVPixelBufferRefErrorHandler = func(corevideo.CVImageBufferRef, error)

// NewCVPixelBufferRefErrorBlock wraps a Go [CVPixelBufferRefErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [MERAWProcessor.ProcessFrameFromImageBufferCompletionHandler]
func NewCVPixelBufferRefErrorBlock(handler CVPixelBufferRefErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal corevideo.CVImageBufferRef, errID objc.ID) {
		handler(primitiveVal, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// DataErrorHandler handles The completion block to execute when the read operation finishes.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [MEByteSource.ReadDataOfLengthFromOffsetCompletionHandler]
type DataErrorHandler = func(*foundation.NSData, error)

// NewDataErrorBlock wraps a Go [DataErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [MEByteSource.ReadDataOfLengthFromOffsetCompletionHandler]
func NewDataErrorBlock(handler DataErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *foundation.NSData
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := foundation.NSDataFromID(resultID)
			result = &v
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// ErrorHandler handles The completion block to execute when the load operation finishes.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [MEFormatReader.LoadTrackReadersWithCompletionHandler]
type ErrorHandler = func(error)

// NewErrorBlock wraps a Go [ErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [MEFormatReader.LoadTrackReadersWithCompletionHandler]
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

// Float32ErrorHandler handles The completion block to execute when the load operation finishes.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [METrackReader.LoadEstimatedDataRateWithCompletionHandler]
type Float32ErrorHandler = func(float32, error)

// NewFloat32ErrorBlock wraps a Go [Float32ErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [METrackReader.LoadEstimatedDataRateWithCompletionHandler]
func NewFloat32ErrorBlock(handler Float32ErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal float32, errID objc.ID) {
		handler(primitiveVal, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// MEFileInfoErrorHandler handles The completion block to execute when the load operation finishes.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [MEFormatReader.LoadFileInfoWithCompletionHandler]
type MEFileInfoErrorHandler = func(*MEFileInfo, error)

// NewMEFileInfoErrorBlock wraps a Go [MEFileInfoErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [MEFormatReader.LoadFileInfoWithCompletionHandler]
func NewMEFileInfoErrorBlock(handler MEFileInfoErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *MEFileInfo
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := MEFileInfoFromID(resultID)
			result = &v
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// MEFormatReaderParseAdditionalFragmentsStatusErrorHandler handles The completion block to execute when the parse operation finishes.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [MEFormatReader.ParseAdditionalFragmentsWithCompletionHandler]
type MEFormatReaderParseAdditionalFragmentsStatusErrorHandler = func(MEFormatReaderParseAdditionalFragmentsStatus, error)

// NewMEFormatReaderParseAdditionalFragmentsStatusErrorBlock wraps a Go [MEFormatReaderParseAdditionalFragmentsStatusErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [MEFormatReader.ParseAdditionalFragmentsWithCompletionHandler]
func NewMEFormatReaderParseAdditionalFragmentsStatusErrorBlock(handler MEFormatReaderParseAdditionalFragmentsStatusErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal MEFormatReaderParseAdditionalFragmentsStatus, errID objc.ID) {
		handler(primitiveVal, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// MESampleCursorErrorHandler handles The completion block to execute when the generate operation finishes.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [METrackReader.GenerateSampleCursorAtFirstSampleInDecodeOrderWithCompletionHandler]
//   - [METrackReader.GenerateSampleCursorAtLastSampleInDecodeOrderWithCompletionHandler]
//   - [METrackReader.GenerateSampleCursorAtPresentationTimeStampCompletionHandler]
type MESampleCursorErrorHandler = func(MESampleCursor, error)

// NewMESampleCursorErrorBlock wraps a Go [MESampleCursorErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [METrackReader.GenerateSampleCursorAtFirstSampleInDecodeOrderWithCompletionHandler]
//   - [METrackReader.GenerateSampleCursorAtLastSampleInDecodeOrderWithCompletionHandler]
//   - [METrackReader.GenerateSampleCursorAtPresentationTimeStampCompletionHandler]
func NewMESampleCursorErrorBlock(handler MESampleCursorErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result MESampleCursor
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			result = MESampleCursorObjectFromID(resultID)
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// METrackInfoErrorHandler handles The completion block to execute when the load operation finishes.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [METrackReader.LoadTrackInfoWithCompletionHandler]
type METrackInfoErrorHandler = func(*METrackInfo, error)

// NewMETrackInfoErrorBlock wraps a Go [METrackInfoErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [METrackReader.LoadTrackInfoWithCompletionHandler]
func NewMETrackInfoErrorBlock(handler METrackInfoErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *METrackInfo
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := METrackInfoFromID(resultID)
			result = &v
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// StringidDictionaryErrorHandler handles The handler that will be invoked when the method completes.
//
// Used by:
//   - [MESampleCursor.LoadPostDecodeProcessingMetadataWithCompletionHandler]
type StringidDictionaryErrorHandler = func(*foundation.INSDictionary, error)

// int64_tErrorHandler handles The completion block to execute when the move operation finishes.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [MESampleCursor.StepInDecodeOrderByCountCompletionHandler]
//   - [MESampleCursor.StepInPresentationOrderByCountCompletionHandler]
//   - [METrackReader.LoadTotalSampleDataLengthWithCompletionHandler]
type int64_tErrorHandler = func(int64, error)

// Newint64_tErrorBlock wraps a Go [int64_tErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [MESampleCursor.StepInDecodeOrderByCountCompletionHandler]
//   - [MESampleCursor.StepInPresentationOrderByCountCompletionHandler]
//   - [METrackReader.LoadTotalSampleDataLengthWithCompletionHandler]
func Newint64_tErrorBlock(handler int64_tErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal int64, errID objc.ID) {
		handler(primitiveVal, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}
