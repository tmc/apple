// Code generated from Apple documentation. DO NOT EDIT.

package devicecheck

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// DataErrorHandler handles A closure that the method calls upon completion with the following parameters:
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [DCAppAttestService.AttestKeyClientDataHashCompletionHandler]
//   - [DCAppAttestService.GenerateAssertionClientDataHashCompletionHandler]
//   - [DCDevice.GenerateTokenWithCompletionHandler]
type DataErrorHandler = func(*foundation.NSData, error)

// NewDataErrorBlock wraps a Go [DataErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [DCAppAttestService.AttestKeyClientDataHashCompletionHandler]
//   - [DCAppAttestService.GenerateAssertionClientDataHashCompletionHandler]
//   - [DCDevice.GenerateTokenWithCompletionHandler]
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

// StringErrorHandler handles A closure that the method calls upon completion with the following parameters:
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [DCAppAttestService.GenerateKeyWithCompletionHandler]
type StringErrorHandler = func(*string, error)

// NewStringErrorBlock wraps a Go [StringErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [DCAppAttestService.GenerateKeyWithCompletionHandler]
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
