// Code generated from Apple documentation. DO NOT EDIT.

package imagecapturecore

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// DataDataErrorHandler is the signature for a completion handler block.
//
// Used by:
//   - [ICCameraDevice.RequestSendPTPCommandOutDataCompletion]
type DataDataErrorHandler = func(*foundation.NSData, *foundation.NSData, error)

// NewDataDataErrorBlock wraps a Go [DataDataErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [ICCameraDevice.RequestSendPTPCommandOutDataCompletion]
func NewDataDataErrorBlock(handler DataDataErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0ID objc.ID, errID objc.ID) {
		var result *foundation.NSData
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := foundation.NSDataFromID(resultID)
			result = &v
		}
		var extra0 *foundation.NSData
		if extra0ID != 0 {
			objc.Send[objc.ID](extra0ID, objc.Sel("retain"))
			v := foundation.NSDataFromID(extra0ID)
			extra0 = &v
		}
		handler(result, extra0, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// DataErrorHandler is the signature for a completion handler block.
//
// Used by:
//   - [ICCameraFile.RequestReadDataAtOffsetLengthCompletion]
//   - [ICCameraFile.RequestThumbnailDataWithOptionsCompletion]
type DataErrorHandler = func(*foundation.NSData, error)

// NewDataErrorBlock wraps a Go [DataErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [ICCameraFile.RequestReadDataAtOffsetLengthCompletion]
//   - [ICCameraFile.RequestThumbnailDataWithOptionsCompletion]
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

// DataHandler is the signature for a completion handler block.
type DataHandler = func(*foundation.NSData)

// NewDataBlock wraps a Go [DataHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewDataBlock(handler DataHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *foundation.NSData
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := foundation.NSDataFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// DictionaryErrorHandler is the signature for a completion handler block.
//
// Used by:
//   - [ICCameraFile.RequestMetadataDictionaryWithOptionsCompletion]
type DictionaryErrorHandler = func(*foundation.INSDictionary, error)

// ErrorHandler is the signature for a completion handler block.
//
// Used by:
//   - [ICDevice.RequestCloseSessionWithOptionsCompletion]
//   - [ICDevice.RequestEjectWithCompletion]
//   - [ICDevice.RequestOpenSessionWithOptionsCompletion]
type ErrorHandler = func(error)

// NewErrorBlock wraps a Go [ErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [ICDevice.RequestCloseSessionWithOptionsCompletion]
//   - [ICDevice.RequestEjectWithCompletion]
//   - [ICDevice.RequestOpenSessionWithOptionsCompletion]
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

// StringErrorHandler is the signature for a completion handler block.
//
// Used by:
//   - [ICCameraFile.RequestDownloadWithOptionsCompletion]
//   - [ICCameraFile.RequestFingerprintWithCompletion]
type StringErrorHandler = func(*string, error)

// NewStringErrorBlock wraps a Go [StringErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [ICCameraFile.RequestDownloadWithOptionsCompletion]
//   - [ICCameraFile.RequestFingerprintWithCompletion]
func NewStringErrorBlock(handler StringErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *string
		if resultID != 0 {
			v := objc.IDToString(resultID)
			result = &v
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// StringHandler is the signature for a completion handler block.
//
// Used by:
//   - [ICDeviceBrowser.RequestContentsAuthorizationWithCompletion]
//   - [ICDeviceBrowser.RequestControlAuthorizationWithCompletion]
//   - [ICDeviceBrowser.ResetContentsAuthorizationWithCompletion]
//   - [ICDeviceBrowser.ResetControlAuthorizationWithCompletion]
type StringHandler = func(*string)

// NewStringBlock wraps a Go [StringHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [ICDeviceBrowser.RequestContentsAuthorizationWithCompletion]
//   - [ICDeviceBrowser.RequestControlAuthorizationWithCompletion]
//   - [ICDeviceBrowser.ResetContentsAuthorizationWithCompletion]
//   - [ICDeviceBrowser.ResetControlAuthorizationWithCompletion]
func NewStringBlock(handler StringHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *string
		if resultID != 0 {
			v := objc.IDToString(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// StringICCameraItemDictionaryHandler is the signature for a completion handler block.
//
// Used by:
//   - [ICCameraDevice.RequestDeleteFilesDeleteFailedCompletion]
type StringICCameraItemDictionaryHandler = func(*foundation.INSDictionary)

// StringNSArrayICCameraItemDictionaryErrorHandler is the signature for a completion handler block.
//
// Used by:
//   - [ICCameraDevice.RequestDeleteFilesDeleteFailedCompletion]
type StringNSArrayICCameraItemDictionaryErrorHandler = func(*foundation.INSDictionary, error)

// URLErrorHandler is the signature for a completion handler block.
//
// Used by:
//   - [ICCameraFile.RequestSecurityScopedURLWithCompletion]
type URLErrorHandler = func(*foundation.NSURL, error)

// NewURLErrorBlock wraps a Go [URLErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [ICCameraFile.RequestSecurityScopedURLWithCompletion]
func NewURLErrorBlock(handler URLErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *foundation.NSURL
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := foundation.NSURLFromID(resultID)
			result = &v
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}
