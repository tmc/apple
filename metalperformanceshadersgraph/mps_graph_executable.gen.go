// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/metalperformanceshaders"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSGraphExecutable] class.
var (
	_MPSGraphExecutableClass     MPSGraphExecutableClass
	_MPSGraphExecutableClassOnce sync.Once
)

func getMPSGraphExecutableClass() MPSGraphExecutableClass {
	_MPSGraphExecutableClassOnce.Do(func() {
		_MPSGraphExecutableClass = MPSGraphExecutableClass{class: objc.GetClass("MPSGraphExecutable")}
	})
	return _MPSGraphExecutableClass
}

// GetMPSGraphExecutableClass returns the class object for MPSGraphExecutable.
func GetMPSGraphExecutableClass() MPSGraphExecutableClass {
	return getMPSGraphExecutableClass()
}

type MPSGraphExecutableClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSGraphExecutableClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSGraphExecutableClass) Alloc() MPSGraphExecutable {
	rv := objc.Send[MPSGraphExecutable](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// The compiled representation of a compute graph executable.
//
// # Overview
//
// An [MPSGraphExecutable] is a compiled graph for specific feeds for specific
// target tensors and target operations.
//
// # Initializers
//
//   - [MPSGraphExecutable.InitWithCoreMLPackageAtURLCompilationDescriptor]: Initialize the executable with the Core ML model package at the provided URL.
//   - [MPSGraphExecutable.InitWithMPSGraphPackageAtURLCompilationDescriptor]: Initialize the executable with the Metal Performance Shaders Graph package at the provided URL.
//
// # Instance Properties
//
//   - [MPSGraphExecutable.FeedTensors]: Tensors fed to the graph, can be used to order the inputs when executable is created with a graph.
//   - [MPSGraphExecutable.Options]: Options for the graph executable.
//   - [MPSGraphExecutable.SetOptions]
//   - [MPSGraphExecutable.TargetTensors]: Tensors targeted by the graph, can be used to order the outputs when executable was created with a graph.
//
// # Instance Methods
//
//   - [MPSGraphExecutable.EncodeToCommandBufferInputsArrayResultsArrayExecutionDescriptor]: Runs the graph for the given feeds and returns the target tensor values, ensuring all target operations also executed. This call is asynchronous and will return immediately after finishing encoding.
//   - [MPSGraphExecutable.GetOutputTypesWithDeviceInputTypesCompilationDescriptor]: Get output shapes for a specialized executable.
//   - [MPSGraphExecutable.RunWithMTLCommandQueueInputsArrayResultsArrayExecutionDescriptor]: Runs the graph for the given feeds and returns the target tensor values, ensuring all target operations also executed.
//   - [MPSGraphExecutable.RunAsyncWithMTLCommandQueueInputsArrayResultsArrayExecutionDescriptor]: Runs the graph for the given feeds and returns the target tensor values, ensuring all target operations also executed. This call is asynchronous and will return immediately.
//   - [MPSGraphExecutable.SerializeToMPSGraphPackageAtURLDescriptor]: Serialize the MPSGraph executable at the provided url.
//   - [MPSGraphExecutable.SpecializeWithDeviceInputTypesCompilationDescriptor]: Specialize the executable and optimize it.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutable
type MPSGraphExecutable struct {
	MPSGraphObject
}

// MPSGraphExecutableFromID constructs a [MPSGraphExecutable] from an objc.ID.
//
// The compiled representation of a compute graph executable.
func MPSGraphExecutableFromID(id objc.ID) MPSGraphExecutable {
	return MPSGraphExecutable{MPSGraphObject: MPSGraphObjectFromID(id)}
}

// NOTE: MPSGraphExecutable adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSGraphExecutable] class.
//
// # Initializers
//
//   - [IMPSGraphExecutable.InitWithCoreMLPackageAtURLCompilationDescriptor]: Initialize the executable with the Core ML model package at the provided URL.
//   - [IMPSGraphExecutable.InitWithMPSGraphPackageAtURLCompilationDescriptor]: Initialize the executable with the Metal Performance Shaders Graph package at the provided URL.
//
// # Instance Properties
//
//   - [IMPSGraphExecutable.FeedTensors]: Tensors fed to the graph, can be used to order the inputs when executable is created with a graph.
//   - [IMPSGraphExecutable.Options]: Options for the graph executable.
//   - [IMPSGraphExecutable.SetOptions]
//   - [IMPSGraphExecutable.TargetTensors]: Tensors targeted by the graph, can be used to order the outputs when executable was created with a graph.
//
// # Instance Methods
//
//   - [IMPSGraphExecutable.EncodeToCommandBufferInputsArrayResultsArrayExecutionDescriptor]: Runs the graph for the given feeds and returns the target tensor values, ensuring all target operations also executed. This call is asynchronous and will return immediately after finishing encoding.
//   - [IMPSGraphExecutable.GetOutputTypesWithDeviceInputTypesCompilationDescriptor]: Get output shapes for a specialized executable.
//   - [IMPSGraphExecutable.RunWithMTLCommandQueueInputsArrayResultsArrayExecutionDescriptor]: Runs the graph for the given feeds and returns the target tensor values, ensuring all target operations also executed.
//   - [IMPSGraphExecutable.RunAsyncWithMTLCommandQueueInputsArrayResultsArrayExecutionDescriptor]: Runs the graph for the given feeds and returns the target tensor values, ensuring all target operations also executed. This call is asynchronous and will return immediately.
//   - [IMPSGraphExecutable.SerializeToMPSGraphPackageAtURLDescriptor]: Serialize the MPSGraph executable at the provided url.
//   - [IMPSGraphExecutable.SpecializeWithDeviceInputTypesCompilationDescriptor]: Specialize the executable and optimize it.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutable
type IMPSGraphExecutable interface {
	IMPSGraphObject

	// Topic: Initializers

	// Initialize the executable with the Core ML model package at the provided URL.
	InitWithCoreMLPackageAtURLCompilationDescriptor(coreMLPackageURL foundation.NSURL, compilationDescriptor IMPSGraphCompilationDescriptor) MPSGraphExecutable
	// Initialize the executable with the Metal Performance Shaders Graph package at the provided URL.
	InitWithMPSGraphPackageAtURLCompilationDescriptor(mpsgraphPackageURL foundation.NSURL, compilationDescriptor IMPSGraphCompilationDescriptor) MPSGraphExecutable

	// Topic: Instance Properties

	// Tensors fed to the graph, can be used to order the inputs when executable is created with a graph.
	FeedTensors() []MPSGraphTensor
	// Options for the graph executable.
	Options() MPSGraphOptions
	SetOptions(value MPSGraphOptions)
	// Tensors targeted by the graph, can be used to order the outputs when executable was created with a graph.
	TargetTensors() []MPSGraphTensor

	// Topic: Instance Methods

	// Runs the graph for the given feeds and returns the target tensor values, ensuring all target operations also executed. This call is asynchronous and will return immediately after finishing encoding.
	EncodeToCommandBufferInputsArrayResultsArrayExecutionDescriptor(commandBuffer *metalperformanceshaders.MPSCommandBuffer, inputsArray []MPSGraphTensorData, resultsArray []MPSGraphTensorData, executionDescriptor IMPSGraphExecutableExecutionDescriptor) []MPSGraphTensorData
	// Get output shapes for a specialized executable.
	GetOutputTypesWithDeviceInputTypesCompilationDescriptor(device IMPSGraphDevice, inputTypes []MPSGraphType, compilationDescriptor IMPSGraphCompilationDescriptor) []MPSGraphShapedType
	// Runs the graph for the given feeds and returns the target tensor values, ensuring all target operations also executed.
	RunWithMTLCommandQueueInputsArrayResultsArrayExecutionDescriptor(commandQueue metal.MTLCommandQueue, inputsArray []MPSGraphTensorData, resultsArray []MPSGraphTensorData, executionDescriptor IMPSGraphExecutableExecutionDescriptor) []MPSGraphTensorData
	// Runs the graph for the given feeds and returns the target tensor values, ensuring all target operations also executed. This call is asynchronous and will return immediately.
	RunAsyncWithMTLCommandQueueInputsArrayResultsArrayExecutionDescriptor(commandQueue metal.MTLCommandQueue, inputsArray []MPSGraphTensorData, resultsArray []MPSGraphTensorData, executionDescriptor IMPSGraphExecutableExecutionDescriptor) []MPSGraphTensorData
	// Serialize the MPSGraph executable at the provided url.
	SerializeToMPSGraphPackageAtURLDescriptor(url foundation.NSURL, descriptor IMPSGraphExecutableSerializationDescriptor)
	// Specialize the executable and optimize it.
	SpecializeWithDeviceInputTypesCompilationDescriptor(device IMPSGraphDevice, inputTypes []MPSGraphType, compilationDescriptor IMPSGraphCompilationDescriptor)
}

// Init initializes the instance.
func (g MPSGraphExecutable) Init() MPSGraphExecutable {
	rv := objc.Send[MPSGraphExecutable](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g MPSGraphExecutable) Autorelease() MPSGraphExecutable {
	rv := objc.Send[MPSGraphExecutable](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSGraphExecutable creates a new MPSGraphExecutable instance.
func NewMPSGraphExecutable() MPSGraphExecutable {
	class := getMPSGraphExecutableClass()
	rv := objc.Send[MPSGraphExecutable](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initialize the executable with the Core ML model package at the provided
// URL.
//
// coreMLPackageURL: The URL where to read the Core ML model package.
//
// compilationDescriptor: Compilation descriptor to be used to specialize, since the executable was
// created with a compilationDescriptor already this one overrides those
// settings to the extent it can.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutable/init(coreMLPackageAtURL:descriptor:)
func NewGraphExecutableWithCoreMLPackageAtURLCompilationDescriptor(coreMLPackageURL foundation.NSURL, compilationDescriptor IMPSGraphCompilationDescriptor) MPSGraphExecutable {
	instance := getMPSGraphExecutableClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoreMLPackageAtURL:compilationDescriptor:"), coreMLPackageURL, compilationDescriptor)
	return MPSGraphExecutableFromID(rv)
}

// Initialize the executable with the Metal Performance Shaders Graph package
// at the provided URL.
//
// mpsgraphPackageURL: The URL where to read the serialized MPSGraphExecutable.
//
// compilationDescriptor: Compilation descriptor to be used to specialize, since the executable was
// created with a compilationDescriptor already this one overrides those
// settings to the extent it can.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutable/init(package:descriptor:)
func NewGraphExecutableWithMPSGraphPackageAtURLCompilationDescriptor(mpsgraphPackageURL foundation.NSURL, compilationDescriptor IMPSGraphCompilationDescriptor) MPSGraphExecutable {
	instance := getMPSGraphExecutableClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMPSGraphPackageAtURL:compilationDescriptor:"), mpsgraphPackageURL, compilationDescriptor)
	return MPSGraphExecutableFromID(rv)
}

// Initialize the executable with the Core ML model package at the provided
// URL.
//
// coreMLPackageURL: The URL where to read the Core ML model package.
//
// compilationDescriptor: Compilation descriptor to be used to specialize, since the executable was
// created with a compilationDescriptor already this one overrides those
// settings to the extent it can.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutable/init(coreMLPackageAtURL:descriptor:)
func (g MPSGraphExecutable) InitWithCoreMLPackageAtURLCompilationDescriptor(coreMLPackageURL foundation.NSURL, compilationDescriptor IMPSGraphCompilationDescriptor) MPSGraphExecutable {
	rv := objc.Send[MPSGraphExecutable](g.ID, objc.Sel("initWithCoreMLPackageAtURL:compilationDescriptor:"), coreMLPackageURL, compilationDescriptor)
	return rv
}

// Initialize the executable with the Metal Performance Shaders Graph package
// at the provided URL.
//
// mpsgraphPackageURL: The URL where to read the serialized MPSGraphExecutable.
//
// compilationDescriptor: Compilation descriptor to be used to specialize, since the executable was
// created with a compilationDescriptor already this one overrides those
// settings to the extent it can.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutable/init(package:descriptor:)
func (g MPSGraphExecutable) InitWithMPSGraphPackageAtURLCompilationDescriptor(mpsgraphPackageURL foundation.NSURL, compilationDescriptor IMPSGraphCompilationDescriptor) MPSGraphExecutable {
	rv := objc.Send[MPSGraphExecutable](g.ID, objc.Sel("initWithMPSGraphPackageAtURL:compilationDescriptor:"), mpsgraphPackageURL, compilationDescriptor)
	return rv
}

// Runs the graph for the given feeds and returns the target tensor values,
// ensuring all target operations also executed. This call is asynchronous and
// will return immediately after finishing encoding.
//
// commandBuffer: CommandBuffer passed to exectute the graph on, commitAndContinue might be
// called, please don’t rely on underlying MTLCommandBuffer to remain
// uncommitted
//
// inputsArray: Feeds tensorData for the placeholder tensors, same order as arguments of
// main function
//
// resultsArray: Tensors for which the caller wishes MPSGraphTensorData to be returned
//
// executionDescriptor: ExecutionDescriptor to be passed in and used,
//
// # Return Value
//
// A valid MPSGraphTensorData array with results synchronized to the CPU
// memory if MPSGraphOptionsSynchronizeResults set.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutable/encode(to:inputs:results:executionDescriptor:)
func (g MPSGraphExecutable) EncodeToCommandBufferInputsArrayResultsArrayExecutionDescriptor(commandBuffer *metalperformanceshaders.MPSCommandBuffer, inputsArray []MPSGraphTensorData, resultsArray []MPSGraphTensorData, executionDescriptor IMPSGraphExecutableExecutionDescriptor) []MPSGraphTensorData {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("encodeToCommandBuffer:inputsArray:resultsArray:executionDescriptor:"), commandBuffer.ID, objectivec.IObjectSliceToNSArray(inputsArray), objectivec.IObjectSliceToNSArray(resultsArray), executionDescriptor)
	return objc.ConvertSlice(rv, func(id objc.ID) MPSGraphTensorData {
		return MPSGraphTensorDataFromID(id)
	})
}

// Get output shapes for a specialized executable.
//
// device: Optional MPSGraph device to compile with
//
// inputTypes: Input types expected to be passed to the executable.
//
// compilationDescriptor: CompilationDescriptor to be used to specialize, since the executable was
// created with a compilationDescriptor already this one overrides those
// settings to the extent it can.
//
// # Discussion
//
// In case specialization has not been done yet then calling this function
// will specialize for the given input shapes.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutable/getOutputTypes(with:inputTypes:compilationDescriptor:)
func (g MPSGraphExecutable) GetOutputTypesWithDeviceInputTypesCompilationDescriptor(device IMPSGraphDevice, inputTypes []MPSGraphType, compilationDescriptor IMPSGraphCompilationDescriptor) []MPSGraphShapedType {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("getOutputTypesWithDevice:inputTypes:compilationDescriptor:"), device, objectivec.IObjectSliceToNSArray(inputTypes), compilationDescriptor)
	return objc.ConvertSlice(rv, func(id objc.ID) MPSGraphShapedType {
		return MPSGraphShapedTypeFromID(id)
	})
}

// Runs the graph for the given feeds and returns the target tensor values,
// ensuring all target operations also executed.
//
// commandQueue: CommandQueue passed to exectute the graph on.
//
// inputsArray: Feeds tensorData for the placeholder tensors, same order as arguments of
// main function.
//
// resultsArray: Results tensorData for which the caller wishes MPSGraphTensorData to be
// returned.
//
// # Return Value
//
// A valid MPSGraphTensorData array with results synchronized to the CPU
// memory if MPSGraphOptionsSynchronizeResults set.
//
// # Discussion
//
// This call is synchronous and will return on completion of execution.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutable/run(with:inputs:results:executionDescriptor:)
func (g MPSGraphExecutable) RunWithMTLCommandQueueInputsArrayResultsArrayExecutionDescriptor(commandQueue metal.MTLCommandQueue, inputsArray []MPSGraphTensorData, resultsArray []MPSGraphTensorData, executionDescriptor IMPSGraphExecutableExecutionDescriptor) []MPSGraphTensorData {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("runWithMTLCommandQueue:inputsArray:resultsArray:executionDescriptor:"), commandQueue, objectivec.IObjectSliceToNSArray(inputsArray), objectivec.IObjectSliceToNSArray(resultsArray), executionDescriptor)
	return objc.ConvertSlice(rv, func(id objc.ID) MPSGraphTensorData {
		return MPSGraphTensorDataFromID(id)
	})
}

// Runs the graph for the given feeds and returns the target tensor values,
// ensuring all target operations also executed. This call is asynchronous and
// will return immediately.
//
// commandQueue: CommandQueue passed to exectute the graph on.
//
// inputsArray: Feeds tensorData for the placeholder tensors, same order as arguments of
// main function.
//
// resultsArray: Tensors for which the caller wishes MPSGraphTensorData to be returned.
//
// executionDescriptor: ExecutionDescriptor to be passed in and used.
//
// # Return Value
//
// A valid MPSGraphTensorData array with results synchronized to the CPU
// memory if MPSGraphOptionsSynchronizeResults set.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutable/runAsync(with:inputs:results:executionDescriptor:)
func (g MPSGraphExecutable) RunAsyncWithMTLCommandQueueInputsArrayResultsArrayExecutionDescriptor(commandQueue metal.MTLCommandQueue, inputsArray []MPSGraphTensorData, resultsArray []MPSGraphTensorData, executionDescriptor IMPSGraphExecutableExecutionDescriptor) []MPSGraphTensorData {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("runAsyncWithMTLCommandQueue:inputsArray:resultsArray:executionDescriptor:"), commandQueue, objectivec.IObjectSliceToNSArray(inputsArray), objectivec.IObjectSliceToNSArray(resultsArray), executionDescriptor)
	return objc.ConvertSlice(rv, func(id objc.ID) MPSGraphTensorData {
		return MPSGraphTensorDataFromID(id)
	})
}

// Serialize the MPSGraph executable at the provided url.
//
// url: The URL where to serialize the MPSGraph executable.
//
// descriptor: The descriptor to be used to serialize the graph.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutable/serialize(package:descriptor:)
func (g MPSGraphExecutable) SerializeToMPSGraphPackageAtURLDescriptor(url foundation.NSURL, descriptor IMPSGraphExecutableSerializationDescriptor) {
	objc.Send[objc.ID](g.ID, objc.Sel("serializeToMPSGraphPackageAtURL:descriptor:"), url, descriptor)
}

// Specialize the executable and optimize it.
//
// device: Ooptional MPSGraph device to compile with.
//
// inputTypes: Input types expected to be passed to the executable.
//
// compilationDescriptor: Compilation descriptor to be used to specialize, since the executable was
// created with a compilationDescriptor already this one overrides those
// settings to the extent it can.
//
// # Discussion
//
// Use this method to choose when specialization happens, else it occurs at
// encode time automatically.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutable/specialize(with:inputTypes:compilationDescriptor:)
func (g MPSGraphExecutable) SpecializeWithDeviceInputTypesCompilationDescriptor(device IMPSGraphDevice, inputTypes []MPSGraphType, compilationDescriptor IMPSGraphCompilationDescriptor) {
	objc.Send[objc.ID](g.ID, objc.Sel("specializeWithDevice:inputTypes:compilationDescriptor:"), device, objectivec.IObjectSliceToNSArray(inputTypes), compilationDescriptor)
}

// Tensors fed to the graph, can be used to order the inputs when executable
// is created with a graph.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutable/feedTensors
func (g MPSGraphExecutable) FeedTensors() []MPSGraphTensor {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("feedTensors"))
	return objc.ConvertSlice(rv, func(id objc.ID) MPSGraphTensor {
		return MPSGraphTensorFromID(id)
	})
}

// Options for the graph executable.
//
// # Discussion
//
// Default value is [MPSGraphOptionsDefault].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutable/options
func (g MPSGraphExecutable) Options() MPSGraphOptions {
	rv := objc.Send[MPSGraphOptions](g.ID, objc.Sel("options"))
	return MPSGraphOptions(rv)
}
func (g MPSGraphExecutable) SetOptions(value MPSGraphOptions) {
	objc.Send[struct{}](g.ID, objc.Sel("setOptions:"), value)
}

// Tensors targeted by the graph, can be used to order the outputs when
// executable was created with a graph.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutable/targetTensors
func (g MPSGraphExecutable) TargetTensors() []MPSGraphTensor {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("targetTensors"))
	return objc.ConvertSlice(rv, func(id objc.ID) MPSGraphTensor {
		return MPSGraphTensorFromID(id)
	})
}
