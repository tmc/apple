// Code generated from Apple documentation. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// MPSGraphCompilationCompletionHandler handles A notification that appears when compilation finishes.

// NewMPSGraphCompilationCompletionHandlerBlock wraps a Go [MPSGraphCompilationCompletionHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewMPSGraphCompilationCompletionHandlerBlock(handler MPSGraphCompilationCompletionHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive MPSGraphExecutable, extra0 foundation.NSError) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// MPSGraphCompletionHandler handles A notification that appears when graph execution finishes.

// NewMPSGraphCompletionHandlerBlock wraps a Go [MPSGraphCompletionHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewMPSGraphCompletionHandlerBlock(handler MPSGraphCompletionHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveID objc.ID, extra0 foundation.NSError) {
		var primitive foundation.INSDictionary
		if primitiveID != 0 {
			objc.Send[objc.ID](primitiveID, objc.Sel("retain"))
			primitive = foundation.NSDictionaryFromID(primitiveID)
		}
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// MPSGraphExecutableCompletionHandler handles A notification when graph executable execution finishes.

// MPSGraphExecutableErrorHandler is the signature for a completion handler block.
type MPSGraphExecutableErrorHandler = func(*MPSGraphExecutable, error)

// NewMPSGraphExecutableErrorBlock wraps a Go [MPSGraphExecutableErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewMPSGraphExecutableErrorBlock(handler MPSGraphExecutableErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *MPSGraphExecutable
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := MPSGraphExecutableFromID(resultID)
			result = &v
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// MPSGraphExecutableScheduledHandler handles A notification when graph executable execution schedules.

// MPSGraphScheduledHandler handles A notification that appears when graph execution schedules.

// NewMPSGraphScheduledHandlerBlock wraps a Go [MPSGraphScheduledHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewMPSGraphScheduledHandlerBlock(handler MPSGraphScheduledHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveID objc.ID, extra0 foundation.NSError) {
		var primitive foundation.INSDictionary
		if primitiveID != 0 {
			objc.Send[objc.ID](primitiveID, objc.Sel("retain"))
			primitive = foundation.NSDictionaryFromID(primitiveID)
		}
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// MPSGraphTensorArrayHandler handles `afterBlock`, this will execute after the condition evaluation.
//
// Used by:
//   - [MPSGraph.WhileWithInitialInputsBeforeAfterName]
type MPSGraphTensorArrayHandler = func(*[]MPSGraphTensor) []MPSGraphTensor

// MPSGraphTensorDataArrayErrorHandler is the signature for a completion handler block.
type MPSGraphTensorDataArrayErrorHandler = func(*[]MPSGraphTensorData, error)

// NewMPSGraphTensorDataArrayErrorBlock wraps a Go [MPSGraphTensorDataArrayErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewMPSGraphTensorDataArrayErrorBlock(handler MPSGraphTensorDataArrayErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *[]MPSGraphTensorData
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			obj := foundation.NSArrayFromID(resultID)
			count := obj.Count()
			res := make([]MPSGraphTensorData, count)
			for i := uint(0); i < count; i++ {
				item := obj.ObjectAtIndex(i)
				res[i] = MPSGraphTensorDataFromID(item.GetID())
			}
			result = &res
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// MPSGraphTensorMPSGraphTensorArrayHandler handles This block will execute the body of the for loop.
//
// Used by:
//   - [MPSGraph.ForLoopWithLowerBoundUpperBoundStepInitialBodyArgumentsBodyName]
//   - [MPSGraph.ForLoopWithNumberOfIterationsInitialBodyArgumentsBodyName]
type MPSGraphTensorMPSGraphTensorArrayHandler = func(*MPSGraphTensor, *[]MPSGraphTensor) []MPSGraphTensor

// MPSGraphTensorMPSGraphTensorDataDictionaryErrorHandler is the signature for a completion handler block.
type MPSGraphTensorMPSGraphTensorDataDictionaryErrorHandler = func(*foundation.INSDictionary, error)

// MPSGraphTensorVoidHandler handles MPSGraphControlFlowDependencyBlock which is provided by caller to create dependent ops
//
// Used by:
//   - [MPSGraph.ControlDependencyWithOperationsDependentBlockName]
//   - [MPSGraph.IfWithPredicateTensorThenBlockElseBlockName]
//   - [MPSGraph.WhileWithInitialInputsBeforeAfterName]
type MPSGraphTensorVoidHandler = func() []MPSGraphTensor

// MPSGraphWhileAfterBlock handles The block that executes after the condition evaluates for each iteration.
