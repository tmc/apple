// Code generated from Apple documentation. DO NOT EDIT.

package espresso

import (
	"github.com/tmc/apple/objc"
)

// VoidHandler is the signature for a completion handler block.
//
// Used by:
//   - [ETDataPoint.IterateBuffersByKey]
//   - [ETTask.FitNumberOfBatchesOutputNamesBatchCallback]
//   - [ETTask.FitNumberOfBatchesWithProgress]
//   - [ETTask.FitNumberOfEpochsOutputNamesBatchCallback]
//   - [ETTask.FitNumberOfEpochsWithProgress]
//   - [ETTask.RunBatchesNumberOfBatchesOutputNamesBatchCallback]
//   - [ETTask.RunInferenceOutputNamesBatchCallback]
//   - [EspressoDataFrameStorageExecutor.ExecuteDataFrameStorageWithNetworkBlockBlockPrepareForIndex]
//   - [EspressoDataFrameStorageExecutor.ExecuteDataFrameStorageWithNetworkBlock]
//   - [EspressoDataFrameStorageExecutor.ExecuteDataFrameStorageWithNetworkReferenceNetworkBlockBlockPrepareForIndex]
type VoidHandler = func()

// NewVoidBlock wraps a Go [VoidHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [ETDataPoint.IterateBuffersByKey]
//   - [ETTask.FitNumberOfBatchesOutputNamesBatchCallback]
//   - [ETTask.FitNumberOfBatchesWithProgress]
//   - [ETTask.FitNumberOfEpochsOutputNamesBatchCallback]
//   - [ETTask.FitNumberOfEpochsWithProgress]
//   - [ETTask.RunBatchesNumberOfBatchesOutputNamesBatchCallback]
//   - [ETTask.RunInferenceOutputNamesBatchCallback]
//   - [EspressoDataFrameStorageExecutor.ExecuteDataFrameStorageWithNetworkBlockBlockPrepareForIndex]
//   - [EspressoDataFrameStorageExecutor.ExecuteDataFrameStorageWithNetworkBlock]
//   - [EspressoDataFrameStorageExecutor.ExecuteDataFrameStorageWithNetworkReferenceNetworkBlockBlockPrepareForIndex]
func NewVoidBlock(handler VoidHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}

