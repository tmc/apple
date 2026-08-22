// Code generated from Apple documentation. DO NOT EDIT.

package cryptotokenkit

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// BoolErrorHandler handles success: Whether the session could be established successfully.
//   - success: Whether the session could be established successfully.
//   - error: Contains information about the error preventing the transaction from being established.
//
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [TKSmartCard.BeginSessionWithReply]
//   - [TKSmartCardUserInteraction.RunWithReply]
type BoolErrorHandler = func(bool, error)

// NewBoolErrorBlock wraps a Go [BoolErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [TKSmartCard.BeginSessionWithReply]
//   - [TKSmartCardUserInteraction.RunWithReply]
func NewBoolErrorBlock(handler BoolErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal bool, errID objc.ID) {
		handler(primitiveVal, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// DataErrorHandler handles response: The APDU response data, or `nil` if communication with the Smart Card failed.
//   - response: The APDU response data, or `nil` if communication with the Smart Card failed.
//   - error: Contains information about the the error preventing the transaction from being established.
//
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [TKSmartCard.TransmitRequestReply]
type DataErrorHandler = func(*foundation.NSData, error)

// NewDataErrorBlock wraps a Go [DataErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [TKSmartCard.TransmitRequestReply]
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

// IntVoidHandler handles The block providing a stream of data for an ATR.
//
// Used by:
//   - [TKSmartCardATR.InitWithSource]
type IntVoidHandler = func() int

// NewIntVoidBlock wraps a Go [IntVoidHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [TKSmartCardATR.InitWithSource]
func NewIntVoidBlock(handler IntVoidHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block) int {
		return handler()
	})
	return objc.ID(block), func() { block.Release() }
}

// StringHandler handles A block to be called when the specified token is removed.
//   - tokenID: The identifier of the removed token.
//
// Used by:
//   - [TKTokenWatcher.AddRemovalHandlerForTokenID]
//   - [TKTokenWatcher.SetInsertionHandler]
type StringHandler = func(*string)

// NewStringBlock wraps a Go [StringHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [TKTokenWatcher.AddRemovalHandlerForTokenID]
//   - [TKTokenWatcher.SetInsertionHandler]
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

// TKSmartCardSlotHandler handles slot: The Smart Card reader slot corresponding to the specified name.
//   - slot: The Smart Card reader slot corresponding to the specified name. If no slot exists with that name, this argument is `nil`.
//
// Used by:
//   - [TKSmartCardSlotManager.GetSlotWithNameReply]
type TKSmartCardSlotHandler = func(*TKSmartCardSlot)

// NewTKSmartCardSlotBlock wraps a Go [TKSmartCardSlotHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [TKSmartCardSlotManager.GetSlotWithNameReply]
func NewTKSmartCardSlotBlock(handler TKSmartCardSlotHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *TKSmartCardSlot
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := TKSmartCardSlotFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// TKSmartCardSlotNFCSessionErrorHandler handles Completion handler which returns the NFC session of the created slot or an error on failure.
//
// Used by:
//   - [TKSmartCardSlotManager.CreateNFCSlotWithMessageCompletion]
type TKSmartCardSlotNFCSessionErrorHandler = func(*uintptr, error)
