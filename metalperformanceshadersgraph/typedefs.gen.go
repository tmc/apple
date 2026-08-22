// Code generated from Apple documentation. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"github.com/tmc/apple/foundation"
)

// MPSGraphCallableMap is a dictionary of symbol names and the corresponding executables for them.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphCallableMap
type MPSGraphCallableMap = foundation.INSDictionary

// MPSGraphCompilationCompletionHandler is a notification that appears when compilation finishes.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphCompilationCompletionHandler
type MPSGraphCompilationCompletionHandler = func(executable MPSGraphExecutable, error_ foundation.NSError)

// MPSGraphCompletionHandler is a notification that appears when graph execution finishes.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphCompletionHandler
type MPSGraphCompletionHandler = func(resultsDictionary foundation.INSDictionary, error_ foundation.NSError)

// MPSGraphControlFlowDependencyBlock is the scope where all the operations defined in this block get control-dependency operations.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphControlFlowDependencyBlock
type MPSGraphControlFlowDependencyBlock = func() []MPSGraphTensor

// MPSGraphExecutableCompletionHandler is a notification when graph executable execution finishes.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutableCompletionHandler
type MPSGraphExecutableCompletionHandler = func(results []MPSGraphTensorData, error_ foundation.NSError)

// MPSGraphExecutableScheduledHandler is a notification when graph executable execution schedules.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutableScheduledHandler
type MPSGraphExecutableScheduledHandler = func(results []MPSGraphTensorData, error_ foundation.NSError)

// MPSGraphForLoopBodyBlock is a block for the body in the for loop.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphForLoopBodyBlock
type MPSGraphForLoopBodyBlock = func(index MPSGraphTensor, iterationArguments []MPSGraphTensor) []MPSGraphTensor

// MPSGraphIfThenElseBlock is a block of operations executed under either the if or else condition.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphIfThenElseBlock
type MPSGraphIfThenElseBlock = func() []MPSGraphTensor

// MPSGraphScheduledHandler is a notification that appears when graph execution schedules.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphScheduledHandler
type MPSGraphScheduledHandler = func(resultsDictionary foundation.INSDictionary, error_ foundation.NSError)

// MPSGraphTensorDataDictionary is a dictionary of tensors and corresponding tensor data.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorDataDictionary
type MPSGraphTensorDataDictionary = foundation.INSDictionary

// MPSGraphTensorShapedTypeDictionary is a dictionary of tensors and corresponding shapes for them.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorShapedTypeDictionary
type MPSGraphTensorShapedTypeDictionary = foundation.INSDictionary

// MPSGraphWhileAfterBlock is the block that executes after the condition evaluates for each iteration.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphWhileAfterBlock
type MPSGraphWhileAfterBlock = func(bodyBlockArguments []MPSGraphTensor) []MPSGraphTensor

// MPSGraphWhileBeforeBlock is the block that executes before the condition evaluates for each iteration.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphWhileBeforeBlock
type MPSGraphWhileBeforeBlock = func(inputTensors []MPSGraphTensor, resultTensors foundation.INSArray) MPSGraphTensor
