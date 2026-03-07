// Code generated from Apple documentation. DO NOT EDIT.

package mlruntime

import (
	"github.com/tmc/apple/objc"
)

// ErrorHandler is the signature for a completion handler block.
//
// Used by:
//   - [MLRDonationManager.EncodeAndUploadToDediscoWithIdentifierMeasurementsWithEncodingSchemasMetadataCompletion]
//   - [MLROnDemandConnectionHandler.PerformTaskCompletionHandler]
//   - [MLROnDemandConnectionHandler.PerformTaskInternalCompletionHandler]
//   - [MLROnDemandPlugin.PerformTaskForPluginIDCompletionHandler]
//   - [MLRServiceClient.DonateJSONResultIdentifierCompletion]
//   - [MLRServiceClient.PerformOnRemoteObjectWithBlockErrorHandler]
//   - [MLRServiceClient.PerformOnRemoteObjectWithBlockIsSynchronousErrorHandler]
//   - [MLRServiceClient.PerformSynchronouslyOnRemoteObjectWithBlockErrorHandler]
type ErrorHandler = func()

// NewErrorBlock wraps a Go [ErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [MLRDonationManager.EncodeAndUploadToDediscoWithIdentifierMeasurementsWithEncodingSchemasMetadataCompletion]
//   - [MLROnDemandConnectionHandler.PerformTaskCompletionHandler]
//   - [MLROnDemandConnectionHandler.PerformTaskInternalCompletionHandler]
//   - [MLROnDemandPlugin.PerformTaskForPluginIDCompletionHandler]
//   - [MLRServiceClient.DonateJSONResultIdentifierCompletion]
//   - [MLRServiceClient.PerformOnRemoteObjectWithBlockErrorHandler]
//   - [MLRServiceClient.PerformOnRemoteObjectWithBlockIsSynchronousErrorHandler]
//   - [MLRServiceClient.PerformSynchronouslyOnRemoteObjectWithBlockErrorHandler]
func NewErrorBlock(handler ErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}




