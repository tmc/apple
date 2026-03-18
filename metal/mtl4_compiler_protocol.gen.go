// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// A abstraction for a pipeline state and shader function compiler.
//
// See: https://developer.apple.com/documentation/Metal/MTL4Compiler
type MTL4Compiler interface {
	objectivec.IObject

	// Returns the device that this compiler belongs to.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4Compiler/device
	Device() MTLDevice

	// Returns the optional label you specify at creation time.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4Compiler/label
	Label() string

	// Returns the pipeline data set serializer into which this compiler stores data for all pipelines it creates.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4Compiler/pipelineDataSetSerializer
	PipelineDataSetSerializer() MTL4PipelineDataSetSerializer

	// Creates a new dynamic library from a library containing Metal IR code synchronously.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4Compiler/makeDynamicLibrary(library:)
	NewDynamicLibraryError(library MTLLibrary) (MTLDynamicLibrary, error)

	// Creates a new dynamic library from the contents of a file at an URL location synchronously.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4Compiler/makeDynamicLibrary(url:)
	NewDynamicLibraryWithURLError(url foundation.INSURL) (MTLDynamicLibrary, error)

	// Creates a new Metal library synchronously.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4Compiler/makeLibrary(descriptor:)
	NewLibraryWithDescriptorError(descriptor IMTL4LibraryDescriptor) (MTLLibrary, error)

	// Returns a new compiler task that asyncrhonously creates a binary version of a GPU visible function or GPU intersection function.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4Compiler/newBinaryFunctionWithDescriptor:compilerTaskOptions:completionHandler:
	NewBinaryFunctionWithDescriptorCompilerTaskOptionsCompletionHandler(descriptor IMTL4BinaryFunctionDescriptor, compilerTaskOptions IMTL4CompilerTaskOptions, completionHandler ErrorHandler) MTL4CompilerTask

	// Creates a new binary visible or intersection function synchronously.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4Compiler/newBinaryFunctionWithDescriptor:compilerTaskOptions:error:
	NewBinaryFunctionWithDescriptorCompilerTaskOptionsError(descriptor IMTL4BinaryFunctionDescriptor, compilerTaskOptions IMTL4CompilerTaskOptions) (MTL4BinaryFunction, error)

	// Creates a new compute pipeline state asynchronously.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4Compiler/newComputePipelineStateWithDescriptor:compilerTaskOptions:completionHandler:
	NewComputePipelineStateWithDescriptorCompilerTaskOptionsCompletionHandler(descriptor IMTL4ComputePipelineDescriptor, compilerTaskOptions IMTL4CompilerTaskOptions, completionHandler ErrorHandler) MTL4CompilerTask

	// Creates a new compute pipeline state object synchronously.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4Compiler/newComputePipelineStateWithDescriptor:compilerTaskOptions:error:
	NewComputePipelineStateWithDescriptorCompilerTaskOptionsError(descriptor IMTL4ComputePipelineDescriptor, compilerTaskOptions IMTL4CompilerTaskOptions) (MTLComputePipelineState, error)

	// Creates a new compute pipeline state asynchronously.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4Compiler/newComputePipelineStateWithDescriptor:dynamicLinkingDescriptor:compilerTaskOptions:completionHandler:
	NewComputePipelineStateWithDescriptorDynamicLinkingDescriptorCompilerTaskOptionsCompletionHandler(descriptor IMTL4ComputePipelineDescriptor, dynamicLinkingDescriptor IMTL4PipelineStageDynamicLinkingDescriptor, compilerTaskOptions IMTL4CompilerTaskOptions, completionHandler ErrorHandler) MTL4CompilerTask

	// Creates a new compute pipeline state synchronously.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4Compiler/newComputePipelineStateWithDescriptor:dynamicLinkingDescriptor:compilerTaskOptions:error:
	NewComputePipelineStateWithDescriptorDynamicLinkingDescriptorCompilerTaskOptionsError(descriptor IMTL4ComputePipelineDescriptor, dynamicLinkingDescriptor IMTL4PipelineStageDynamicLinkingDescriptor, compilerTaskOptions IMTL4CompilerTaskOptions) (MTLComputePipelineState, error)

	// Creates a new dynamic Metal library instance asynchronously.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4Compiler/newDynamicLibrary:completionHandler:
	NewDynamicLibraryCompletionHandler(library MTLLibrary, completionHandler ErrorHandler) MTL4CompilerTask

	// Creates a new dynamic library from the contents of a file at an URL location synchronously.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4Compiler/newDynamicLibraryWithURL:completionHandler:
	NewDynamicLibraryWithURLCompletionHandler(url foundation.INSURL, completionHandler ErrorHandler) MTL4CompilerTask

	// Creates a new Metal library instance asynchronously.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4Compiler/newLibraryWithDescriptor:completionHandler:
	NewLibraryWithDescriptorCompletionHandler(descriptor IMTL4LibraryDescriptor, completionHandler ErrorHandler) MTL4CompilerTask

	// Creates a new machine learning pipeline state asynchronously.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4Compiler/newMachineLearningPipelineStateWithDescriptor:completionHandler:
	NewMachineLearningPipelineStateWithDescriptorCompletionHandler(descriptor IMTL4MachineLearningPipelineDescriptor, completionHandler ErrorHandler) MTL4CompilerTask

	// Creates a new ML pipeline state with descriptor.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4Compiler/newMachineLearningPipelineStateWithDescriptor:error:
	NewMachineLearningPipelineStateWithDescriptorError(descriptor IMTL4MachineLearningPipelineDescriptor) (MTL4MachineLearningPipelineState, error)

	// Creates a new render pipeline state from another, previously unspecialized, pipeline state
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4Compiler/newRenderPipelineStateBySpecializationWithDescriptor:pipeline:completionHandler:
	NewRenderPipelineStateBySpecializationWithDescriptorPipelineCompletionHandler(descriptor IMTL4PipelineDescriptor, pipeline MTLRenderPipelineState, completionHandler ErrorHandler) MTL4CompilerTask

	// Creates a new render pipeline state from another, previously unspecialized, pipeline state.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4Compiler/newRenderPipelineStateBySpecializationWithDescriptor:pipeline:error:
	NewRenderPipelineStateBySpecializationWithDescriptorPipelineError(descriptor IMTL4PipelineDescriptor, pipeline MTLRenderPipelineState) (MTLRenderPipelineState, error)

	// Creates a new render pipeline state asynchronously.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4Compiler/newRenderPipelineStateWithDescriptor:compilerTaskOptions:completionHandler:
	NewRenderPipelineStateWithDescriptorCompilerTaskOptionsCompletionHandler(descriptor IMTL4PipelineDescriptor, compilerTaskOptions IMTL4CompilerTaskOptions, completionHandler ErrorHandler) MTL4CompilerTask

	// Creates a new render pipeline state synchronously.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4Compiler/newRenderPipelineStateWithDescriptor:compilerTaskOptions:error:
	NewRenderPipelineStateWithDescriptorCompilerTaskOptionsError(descriptor IMTL4PipelineDescriptor, compilerTaskOptions IMTL4CompilerTaskOptions) (MTLRenderPipelineState, error)

	// Creates a new render pipeline state asynchronously.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4Compiler/newRenderPipelineStateWithDescriptor:dynamicLinkingDescriptor:compilerTaskOptions:completionHandler:
	NewRenderPipelineStateWithDescriptorDynamicLinkingDescriptorCompilerTaskOptionsCompletionHandler(descriptor IMTL4PipelineDescriptor, dynamicLinkingDescriptor IMTL4RenderPipelineDynamicLinkingDescriptor, compilerTaskOptions IMTL4CompilerTaskOptions, completionHandler ErrorHandler) MTL4CompilerTask

	// Creates a new render pipeline state synchronously.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4Compiler/newRenderPipelineStateWithDescriptor:dynamicLinkingDescriptor:compilerTaskOptions:error:
	NewRenderPipelineStateWithDescriptorDynamicLinkingDescriptorCompilerTaskOptionsError(descriptor IMTL4PipelineDescriptor, dynamicLinkingDescriptor IMTL4RenderPipelineDynamicLinkingDescriptor, compilerTaskOptions IMTL4CompilerTaskOptions) (MTLRenderPipelineState, error)
}

// MTL4CompilerObject wraps an existing Objective-C object that conforms to the MTL4Compiler protocol.
type MTL4CompilerObject struct {
	objectivec.Object
}
func (o MTL4CompilerObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTL4CompilerObjectFromID constructs a [MTL4CompilerObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTL4CompilerObjectFromID(id objc.ID) MTL4CompilerObject {
	return MTL4CompilerObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Returns the device that this compiler belongs to.
//
// See: https://developer.apple.com/documentation/Metal/MTL4Compiler/device

func (o MTL4CompilerObject) Device() MTLDevice {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("device"))
	return MTLDeviceObjectFromID(rv)
	}

// Returns the optional label you specify at creation time.
//
// See: https://developer.apple.com/documentation/Metal/MTL4Compiler/label

func (o MTL4CompilerObject) Label() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
	}

// Returns the pipeline data set serializer into which this compiler stores
// data for all pipelines it creates.
//
// See: https://developer.apple.com/documentation/Metal/MTL4Compiler/pipelineDataSetSerializer

func (o MTL4CompilerObject) PipelineDataSetSerializer() MTL4PipelineDataSetSerializer {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("pipelineDataSetSerializer"))
	return MTL4PipelineDataSetSerializerObjectFromID(rv)
	}

// Creates a new dynamic library from a library containing Metal IR code
// synchronously.
//
// library: A library from which this compiler creates the new a dynamic library
//
// # Return Value
// 
// A new dynamic Metal library upon success, `nil` otherwise.
//
// See: https://developer.apple.com/documentation/Metal/MTL4Compiler/makeDynamicLibrary(library:)

func (o MTL4CompilerObject) NewDynamicLibraryError(library MTLLibrary) (MTLDynamicLibrary, error) {
	
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("newDynamicLibrary:error:"), library)
	if err != nil {
		return nil, err
	}
	return MTLDynamicLibraryObjectFromID(rv), nil
	}

// Creates a new dynamic library from the contents of a file at an URL
// location synchronously.
//
// url: An URL referencing a file whose contents this compiler uses to build a
// dynamic library.
//
// # Return Value
// 
// A new dynamic Metal library upon success, `nil` otherwise.
//
// See: https://developer.apple.com/documentation/Metal/MTL4Compiler/makeDynamicLibrary(url:)

func (o MTL4CompilerObject) NewDynamicLibraryWithURLError(url foundation.INSURL) (MTLDynamicLibrary, error) {
	
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("newDynamicLibraryWithURL:error:"), url)
	if err != nil {
		return nil, err
	}
	return MTLDynamicLibraryObjectFromID(rv), nil
	}

// Creates a new Metal library synchronously.
//
// descriptor: A description of the library to create.
//
// # Return Value
// 
// A Metal library instance upon success, `nil` otherwise.
//
// See: https://developer.apple.com/documentation/Metal/MTL4Compiler/makeLibrary(descriptor:)

func (o MTL4CompilerObject) NewLibraryWithDescriptorError(descriptor IMTL4LibraryDescriptor) (MTLLibrary, error) {
	
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("newLibraryWithDescriptor:error:"), descriptor)
	if err != nil {
		return nil, err
	}
	return MTLLibraryObjectFromID(rv), nil
	}

// Returns a new compiler task that asyncrhonously creates a binary version of
// a GPU visible function or GPU intersection function.
//
// descriptor: A configuration that tells the method which GPU function to make into a
// binary function and which options to apply when compiling it.
//
// compilerTaskOptions: A configuration for the compiler task.
//
// completionHandler: A completetion handler that you provide, which the task calls when it
// finishes compiling the binary function.
//
// See: https://developer.apple.com/documentation/Metal/MTL4Compiler/newBinaryFunctionWithDescriptor:compilerTaskOptions:completionHandler:

func (o MTL4CompilerObject) NewBinaryFunctionWithDescriptorCompilerTaskOptionsCompletionHandler(descriptor IMTL4BinaryFunctionDescriptor, compilerTaskOptions IMTL4CompilerTaskOptions, completionHandler ErrorHandler) MTL4CompilerTask {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newBinaryFunctionWithDescriptor:compilerTaskOptions:completionHandler:"), descriptor, compilerTaskOptions, completionHandler)
	return MTL4CompilerTaskObjectFromID(rv)
	}

// Creates a new binary visible or intersection function synchronously.
//
// descriptor: A binary function descriptor to use for creating the binary function.
//
// compilerTaskOptions: A descriptor of the compilation itself, providing parameters that influence
// execution of the compilation process.
//
// error: An optional parameter into which Metal stores information in case of an
// error.
//
// # Return Value
// 
// A new binary function upon success, `nil` otherwise.
//
// See: https://developer.apple.com/documentation/Metal/MTL4Compiler/newBinaryFunctionWithDescriptor:compilerTaskOptions:error:

func (o MTL4CompilerObject) NewBinaryFunctionWithDescriptorCompilerTaskOptionsError(descriptor IMTL4BinaryFunctionDescriptor, compilerTaskOptions IMTL4CompilerTaskOptions) (MTL4BinaryFunction, error) {
	
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("newBinaryFunctionWithDescriptor:compilerTaskOptions:error:"), descriptor, compilerTaskOptions)
	if err != nil {
		return nil, err
	}
	return MTL4BinaryFunctionObjectFromID(rv), nil
	}

// Creates a new compute pipeline state asynchronously.
//
// descriptor: A compute pipeline state descriptor, describing the compute pipeline to
// create.
//
// compilerTaskOptions: A descriptor of the compilation itself, providing parameters that influence
// execution of the compilation process.
//
// completionHandler: A block Metal calls when it finishes the build task.
//
// # Return Value
// 
// A compiler task representing the asynchronous compilation task.
//
// See: https://developer.apple.com/documentation/Metal/MTL4Compiler/newComputePipelineStateWithDescriptor:compilerTaskOptions:completionHandler:

func (o MTL4CompilerObject) NewComputePipelineStateWithDescriptorCompilerTaskOptionsCompletionHandler(descriptor IMTL4ComputePipelineDescriptor, compilerTaskOptions IMTL4CompilerTaskOptions, completionHandler ErrorHandler) MTL4CompilerTask {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newComputePipelineStateWithDescriptor:compilerTaskOptions:completionHandler:"), descriptor, compilerTaskOptions, completionHandler)
	return MTL4CompilerTaskObjectFromID(rv)
	}

// Creates a new compute pipeline state object synchronously.
//
// descriptor: A compute pipeline state descriptor describing the pipeline this compiler
// creates.
//
// compilerTaskOptions: A description of the compilation process itself, providing parameters that
// influence execution of the compilation process.
//
// error: An optional parameter into which Metal stores information in case of an
// error.
//
// # Return Value
// 
// A new compute pipeline state object upon success, `nil` otherwise.
//
// See: https://developer.apple.com/documentation/Metal/MTL4Compiler/newComputePipelineStateWithDescriptor:compilerTaskOptions:error:

func (o MTL4CompilerObject) NewComputePipelineStateWithDescriptorCompilerTaskOptionsError(descriptor IMTL4ComputePipelineDescriptor, compilerTaskOptions IMTL4CompilerTaskOptions) (MTLComputePipelineState, error) {
	
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("newComputePipelineStateWithDescriptor:compilerTaskOptions:error:"), descriptor, compilerTaskOptions)
	if err != nil {
		return nil, err
	}
	return MTLComputePipelineStateObjectFromID(rv), nil
	}

// Creates a new compute pipeline state asynchronously.
//
// descriptor: A compute pipeline state descriptor, describing the compute pipeline to
// create.
//
// dynamicLinkingDescriptor: An optional parameter that provides additional configuration for linking
// the pipeline state object.
//
// compilerTaskOptions: A description of the compilation process itself, providing parameters that
// influence execution of the compilation process.
//
// completionHandler: A block Metal calls when it finishes the build task.
//
// # Return Value
// 
// A compiler task representing the asynchronous compilation task.
//
// See: https://developer.apple.com/documentation/Metal/MTL4Compiler/newComputePipelineStateWithDescriptor:dynamicLinkingDescriptor:compilerTaskOptions:completionHandler:

func (o MTL4CompilerObject) NewComputePipelineStateWithDescriptorDynamicLinkingDescriptorCompilerTaskOptionsCompletionHandler(descriptor IMTL4ComputePipelineDescriptor, dynamicLinkingDescriptor IMTL4PipelineStageDynamicLinkingDescriptor, compilerTaskOptions IMTL4CompilerTaskOptions, completionHandler ErrorHandler) MTL4CompilerTask {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newComputePipelineStateWithDescriptor:dynamicLinkingDescriptor:compilerTaskOptions:completionHandler:"), descriptor, dynamicLinkingDescriptor, compilerTaskOptions, completionHandler)
	return MTL4CompilerTaskObjectFromID(rv)
	}

// Creates a new compute pipeline state synchronously.
//
// descriptor: A compute pipeline state descriptor describing the pipeline this compiler
// creates.
//
// dynamicLinkingDescriptor: An optional parameter that provides additional configuration for linking
// the pipeline state object.
//
// compilerTaskOptions: A description of the compilation process itself, providing parameters that
// influence execution of the compilation process.
//
// error: An optional parameter into which Metal stores information in case of an
// error.
//
// # Return Value
// 
// A new compute pipeline state object upon success, `nil` otherwise.
//
// See: https://developer.apple.com/documentation/Metal/MTL4Compiler/newComputePipelineStateWithDescriptor:dynamicLinkingDescriptor:compilerTaskOptions:error:

func (o MTL4CompilerObject) NewComputePipelineStateWithDescriptorDynamicLinkingDescriptorCompilerTaskOptionsError(descriptor IMTL4ComputePipelineDescriptor, dynamicLinkingDescriptor IMTL4PipelineStageDynamicLinkingDescriptor, compilerTaskOptions IMTL4CompilerTaskOptions) (MTLComputePipelineState, error) {
	
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("newComputePipelineStateWithDescriptor:dynamicLinkingDescriptor:compilerTaskOptions:error:"), descriptor, dynamicLinkingDescriptor, compilerTaskOptions)
	if err != nil {
		return nil, err
	}
	return MTLComputePipelineStateObjectFromID(rv), nil
	}

// Creates a new dynamic Metal library instance asynchronously.
//
// library: A library from which this compiler creates the new a dynamic library
//
// completionHandler: A block Metal calls when it finishes the build task.
//
// # Return Value
// 
// A compiler task representing the asynchronous compilation task.
//
// See: https://developer.apple.com/documentation/Metal/MTL4Compiler/newDynamicLibrary:completionHandler:

func (o MTL4CompilerObject) NewDynamicLibraryCompletionHandler(library MTLLibrary, completionHandler ErrorHandler) MTL4CompilerTask {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newDynamicLibrary:completionHandler:"), library, completionHandler)
	return MTL4CompilerTaskObjectFromID(rv)
	}

// Creates a new dynamic library from the contents of a file at an URL
// location synchronously.
//
// url: An URL referencing a file whose contents this compiler uses to build a
// dynamic library.
//
// completionHandler: A block Metal calls when it finishes the build task.
//
// # Return Value
// 
// A compiler task representing the asynchronous compilation task.
//
// See: https://developer.apple.com/documentation/Metal/MTL4Compiler/newDynamicLibraryWithURL:completionHandler:

func (o MTL4CompilerObject) NewDynamicLibraryWithURLCompletionHandler(url foundation.INSURL, completionHandler ErrorHandler) MTL4CompilerTask {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newDynamicLibraryWithURL:completionHandler:"), url, completionHandler)
	return MTL4CompilerTaskObjectFromID(rv)
	}

// Creates a new Metal library instance asynchronously.
//
// descriptor: A description of the library to create.
//
// completionHandler: A block Metal calls when it finishes the build task.
//
// # Return Value
// 
// A compiler task representing the asynchronous compilation task.
//
// See: https://developer.apple.com/documentation/Metal/MTL4Compiler/newLibraryWithDescriptor:completionHandler:

func (o MTL4CompilerObject) NewLibraryWithDescriptorCompletionHandler(descriptor IMTL4LibraryDescriptor, completionHandler ErrorHandler) MTL4CompilerTask {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newLibraryWithDescriptor:completionHandler:"), descriptor, completionHandler)
	return MTL4CompilerTaskObjectFromID(rv)
	}

// Creates a new machine learning pipeline state asynchronously.
//
// descriptor: A machine learning pipeline state descriptor to use for creating the new
// pipeline state.
//
// completionHandler: A block Metal calls when it finishes the build task.
//
// # Return Value
// 
// A compiler task representing the asynchronous compilation task.
//
// See: https://developer.apple.com/documentation/Metal/MTL4Compiler/newMachineLearningPipelineStateWithDescriptor:completionHandler:

func (o MTL4CompilerObject) NewMachineLearningPipelineStateWithDescriptorCompletionHandler(descriptor IMTL4MachineLearningPipelineDescriptor, completionHandler ErrorHandler) MTL4CompilerTask {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newMachineLearningPipelineStateWithDescriptor:completionHandler:"), descriptor, completionHandler)
	return MTL4CompilerTaskObjectFromID(rv)
	}

// Creates a new ML pipeline state with descriptor.
//
// descriptor: A machine learning pipeline state descriptor to use for creating the new
// pipeline state.
//
// error: An optional parameter into which Metal stores information in case of an
// error.
//
// # Return Value
// 
// A machine learning pipeline state if operation is successful, otherwise
// `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTL4Compiler/newMachineLearningPipelineStateWithDescriptor:error:

func (o MTL4CompilerObject) NewMachineLearningPipelineStateWithDescriptorError(descriptor IMTL4MachineLearningPipelineDescriptor) (MTL4MachineLearningPipelineState, error) {
	
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("newMachineLearningPipelineStateWithDescriptor:error:"), descriptor)
	if err != nil {
		return nil, err
	}
	return MTL4MachineLearningPipelineStateObjectFromID(rv), nil
	}

// Creates a new render pipeline state from another, previously unspecialized,
// pipeline state
//
// descriptor: A render pipeline state descriptor or any type: default, tile, or mesh
// render pipeline descriptor.
//
// pipeline: A render pipeline state containing unspecialized substate.
//
// completionHandler: A block Metal calls when it finishes the build task.
//
// # Return Value
// 
// A compiler task representing the asynchronous compilation task.
//
// # Discussion
// 
// Metal specializes the pipeline state with new state values the descriptor
// provides, observing the following rules:
// 
// - The compiler only updates properties that were originally specified as .
// It doesn’t modify other already-specialized properties - The compiler
// sets to their default behavior any unspecialized properties that your
// passed-in descriptor doesn’t specialize
// 
// Additionally, there are some cases where the Metal can’t specialize a
// pipeline:
// 
// - If the original pipeline state object doesn’t have any unspecialized
// properties - You can’t re-specialize a previously specialized pipeline
// state object
//
// See: https://developer.apple.com/documentation/Metal/MTL4Compiler/newRenderPipelineStateBySpecializationWithDescriptor:pipeline:completionHandler:

func (o MTL4CompilerObject) NewRenderPipelineStateBySpecializationWithDescriptorPipelineCompletionHandler(descriptor IMTL4PipelineDescriptor, pipeline MTLRenderPipelineState, completionHandler ErrorHandler) MTL4CompilerTask {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newRenderPipelineStateBySpecializationWithDescriptor:pipeline:completionHandler:"), descriptor, pipeline, completionHandler)
	return MTL4CompilerTaskObjectFromID(rv)
	}

// Creates a new render pipeline state from another, previously unspecialized,
// pipeline state.
//
// descriptor: A render pipeline state descriptor or any type: default, tile, or mesh
// render pipeline descriptor.
//
// pipeline: A render pipeline state containing unspecialized substate.
//
// error: An optional parameter into which Metal stores information in case of an
// error.
//
// # Return Value
// 
// A fully-specialized pipeline state object.
//
// # Discussion
// 
// Metal specializes the pipeline state with new state values the descriptor
// provides, observing the following rules:
// 
// - The compiler only updates properties that were originally specified as .
// It doesn’t modify other already-specialized properties - The compiler
// sets to their default behavior any unspecialized properties that your
// passed-in descriptor doesn’t specialize
// 
// Additionally, there are some cases where the Metal can’t specialize a
// pipeline:
// 
// - If the original pipeline state object doesn’t have any unspecialized
// properties - You can’t re-specialize a previously specialized pipeline
// state object
//
// See: https://developer.apple.com/documentation/Metal/MTL4Compiler/newRenderPipelineStateBySpecializationWithDescriptor:pipeline:error:

func (o MTL4CompilerObject) NewRenderPipelineStateBySpecializationWithDescriptorPipelineError(descriptor IMTL4PipelineDescriptor, pipeline MTLRenderPipelineState) (MTLRenderPipelineState, error) {
	
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("newRenderPipelineStateBySpecializationWithDescriptor:pipeline:error:"), descriptor, pipeline)
	if err != nil {
		return nil, err
	}
	return MTLRenderPipelineStateObjectFromID(rv), nil
	}

// Creates a new render pipeline state asynchronously.
//
// descriptor: A render, tile, or mesh pipeline state descriptor that describes the
// pipeline to create.
//
// compilerTaskOptions: A description of the compilation process itself, providing parameters that
// influence execution of the compilation process.
//
// completionHandler: A block Metal calls when it finishes the build task.
//
// # Return Value
// 
// A compiler task representing the asynchronous compilation task.
//
// # Discussion
// 
// Use this method to build any render pipeline type, including render, tile,
// and mesh render pipeline states. The type of the descriptor you pass
// indicates the pipeline type this method builds.
// 
// Passing in a compute pipeline descriptor to the `descriptor` parameter
// produces an error.
//
// See: https://developer.apple.com/documentation/Metal/MTL4Compiler/newRenderPipelineStateWithDescriptor:compilerTaskOptions:completionHandler:

func (o MTL4CompilerObject) NewRenderPipelineStateWithDescriptorCompilerTaskOptionsCompletionHandler(descriptor IMTL4PipelineDescriptor, compilerTaskOptions IMTL4CompilerTaskOptions, completionHandler ErrorHandler) MTL4CompilerTask {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newRenderPipelineStateWithDescriptor:compilerTaskOptions:completionHandler:"), descriptor, compilerTaskOptions, completionHandler)
	return MTL4CompilerTaskObjectFromID(rv)
	}

// Creates a new render pipeline state synchronously.
//
// descriptor: A render, tile, or mesh pipeline state descriptor that describes the
// pipeline to create.
//
// compilerTaskOptions: A description of the compilation process itself, providing parameters that
// influence execution of the compilation process.
//
// error: An optional parameter into which Metal stores information in case of an
// error.
//
// # Return Value
// 
// A new render pipeline state object upon success, `nil` otherwise.
//
// # Discussion
// 
// Use this method to build any render pipeline type, including render, tile,
// and mesh render pipeline states. The type of the descriptor you pass
// indicates the pipeline type this method builds.
// 
// Passing in a compute pipeline descriptor to the `descriptor` parameter
// produces an error.
//
// See: https://developer.apple.com/documentation/Metal/MTL4Compiler/newRenderPipelineStateWithDescriptor:compilerTaskOptions:error:

func (o MTL4CompilerObject) NewRenderPipelineStateWithDescriptorCompilerTaskOptionsError(descriptor IMTL4PipelineDescriptor, compilerTaskOptions IMTL4CompilerTaskOptions) (MTLRenderPipelineState, error) {
	
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("newRenderPipelineStateWithDescriptor:compilerTaskOptions:error:"), descriptor, compilerTaskOptions)
	if err != nil {
		return nil, err
	}
	return MTLRenderPipelineStateObjectFromID(rv), nil
	}

// Creates a new render pipeline state asynchronously.
//
// descriptor: A render, tile, or mesh pipeline state descriptor that describes the
// pipeline to create.
//
// dynamicLinkingDescriptor: An optional parameter that provides additional configuration for linking
// the pipeline state object.
//
// compilerTaskOptions: A description of the compilation process itself, providing parameters that
// influence execution of the compilation process.
//
// completionHandler: A block Metal calls when it finishes the build task.
//
// # Return Value
// 
// A compiler task representing the asynchronous compilation task.
//
// # Discussion
// 
// Use this method to build any render pipeline type, including render, tile,
// and mesh render pipeline states. The type of the descriptor you pass
// indicates the pipeline type this method builds.
// 
// Passing in a compute pipeline descriptor to the `descriptor` parameter
// produces an error.
//
// See: https://developer.apple.com/documentation/Metal/MTL4Compiler/newRenderPipelineStateWithDescriptor:dynamicLinkingDescriptor:compilerTaskOptions:completionHandler:

func (o MTL4CompilerObject) NewRenderPipelineStateWithDescriptorDynamicLinkingDescriptorCompilerTaskOptionsCompletionHandler(descriptor IMTL4PipelineDescriptor, dynamicLinkingDescriptor IMTL4RenderPipelineDynamicLinkingDescriptor, compilerTaskOptions IMTL4CompilerTaskOptions, completionHandler ErrorHandler) MTL4CompilerTask {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newRenderPipelineStateWithDescriptor:dynamicLinkingDescriptor:compilerTaskOptions:completionHandler:"), descriptor, dynamicLinkingDescriptor, compilerTaskOptions, completionHandler)
	return MTL4CompilerTaskObjectFromID(rv)
	}

// Creates a new render pipeline state synchronously.
//
// descriptor: A render, tile, or mesh pipeline state descriptor that describes the
// pipeline to create.
//
// dynamicLinkingDescriptor: An optional parameter that provides additional configuration for linking
// the pipeline state object.
//
// compilerTaskOptions: A description of the compilation process itself, providing parameters that
// influence execution of the compilation process.
//
// error: An optional parameter into which Metal stores information in case of an
// error.
//
// # Return Value
// 
// A new render pipeline state object upon success, `nil` otherwise.
//
// # Discussion
// 
// Use this method to build any render pipeline type, including render, tile,
// and mesh render pipeline states. The type of the descriptor you pass
// indicates the pipeline type this method builds.
// 
// Passing in a compute pipeline descriptor to the `descriptor` parameter
// produces an error.
//
// See: https://developer.apple.com/documentation/Metal/MTL4Compiler/newRenderPipelineStateWithDescriptor:dynamicLinkingDescriptor:compilerTaskOptions:error:

func (o MTL4CompilerObject) NewRenderPipelineStateWithDescriptorDynamicLinkingDescriptorCompilerTaskOptionsError(descriptor IMTL4PipelineDescriptor, dynamicLinkingDescriptor IMTL4RenderPipelineDynamicLinkingDescriptor, compilerTaskOptions IMTL4CompilerTaskOptions) (MTLRenderPipelineState, error) {
	
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("newRenderPipelineStateWithDescriptor:dynamicLinkingDescriptor:compilerTaskOptions:error:"), descriptor, dynamicLinkingDescriptor, compilerTaskOptions)
	if err != nil {
		return nil, err
	}
	return MTLRenderPipelineStateObjectFromID(rv), nil
	}

