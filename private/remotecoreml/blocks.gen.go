// Code generated from Apple documentation. DO NOT EDIT.

package remotecoreml

import (
	"github.com/tmc/apple/objc"
)

// ErrorHandler is the signature for a completion handler block.
//
// Used by:
//   - [MLNetworkUtilities.SetupListenerStateChangeHandlerUseUDP]
type ErrorHandler = func()

// NewErrorBlock wraps a Go [ErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [MLNetworkUtilities.SetupListenerStateChangeHandlerUseUDP]
func NewErrorBlock(handler ErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}

// VoidHandler is the signature for a completion handler block.
//
// Used by:
//   - [MLNetworkUtilities.CreateSecureConnectionParameterUseUDP]
//   - [MLNetworking.ReceiveLoop]
//   - [MLNetworking.SetListenerReceiveDataCallBack]
//   - [MLNetworking.SetReceiveDataCallBack]
//   - [MLServer.SetLoadCommand]
//   - [MLServer.SetLoadFunction]
//   - [MLServer.SetPredictCommand]
//   - [MLServer.SetPredictFunction]
//   - [MLServer.SetTextCommand]
//   - [MLServer.SetTextFunction]
//   - [MLServer.SetUnLoadCommand]
//   - [MLServer.SetUnLoadFunction]
type VoidHandler = func()

// NewVoidBlock wraps a Go [VoidHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [MLNetworkUtilities.CreateSecureConnectionParameterUseUDP]
//   - [MLNetworking.ReceiveLoop]
//   - [MLNetworking.SetListenerReceiveDataCallBack]
//   - [MLNetworking.SetReceiveDataCallBack]
//   - [MLServer.SetLoadCommand]
//   - [MLServer.SetLoadFunction]
//   - [MLServer.SetPredictCommand]
//   - [MLServer.SetPredictFunction]
//   - [MLServer.SetTextCommand]
//   - [MLServer.SetTextFunction]
//   - [MLServer.SetUnLoadCommand]
//   - [MLServer.SetUnLoadFunction]
func NewVoidBlock(handler VoidHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}

