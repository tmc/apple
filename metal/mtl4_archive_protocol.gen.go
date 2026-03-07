// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// A read-only container that stores pipeline states from a shader compiler.
//
// See: https://developer.apple.com/documentation/Metal/MTL4Archive
type MTL4Archive interface {
	objectivec.IObject

	// A label that you can associate with this archive.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4Archive/label
	Label() string

	// Synchronously creates a binary version of a GPU visible function or GPU intersection function.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4Archive/makeBinaryFunction(descriptor:)
	NewBinaryFunctionWithDescriptorError(descriptor IMTL4BinaryFunctionDescriptor) (MTL4BinaryFunction, error)

	// Creates a compute pipeline state from the archive with a descriptor.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4Archive/newComputePipelineStateWithDescriptor:error:
	NewComputePipelineStateWithDescriptorError(descriptor IMTL4ComputePipelineDescriptor) (MTLComputePipelineState, error)

	// Creates a compute pipeline state from the archive with a compute descriptor and a dynamic linking descriptor.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4Archive/newComputePipelineStateWithDescriptor:dynamicLinkingDescriptor:error:
	NewComputePipelineStateWithDescriptorDynamicLinkingDescriptorError(descriptor IMTL4ComputePipelineDescriptor, dynamicLinkingDescriptor IMTL4PipelineStageDynamicLinkingDescriptor) (MTLComputePipelineState, error)

	// Creates a render pipeline state from the archive with a descriptor.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4Archive/newRenderPipelineStateWithDescriptor:error:
	NewRenderPipelineStateWithDescriptorError(descriptor IMTL4PipelineDescriptor) (MTLRenderPipelineState, error)

	// Creates a render pipeline state from the archive with a render descriptor and a dynamic linking descriptor.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4Archive/newRenderPipelineStateWithDescriptor:dynamicLinkingDescriptor:error:
	NewRenderPipelineStateWithDescriptorDynamicLinkingDescriptorError(descriptor IMTL4PipelineDescriptor, dynamicLinkingDescriptor IMTL4RenderPipelineDynamicLinkingDescriptor) (MTLRenderPipelineState, error)

	// A label that you can associate with this archive.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4Archive/label
	SetLabel(value string)
}



// MTL4ArchiveObject wraps an existing Objective-C object that conforms to the MTL4Archive protocol.
type MTL4ArchiveObject struct {
	objectivec.Object
}
func (o MTL4ArchiveObject) BaseObject() objectivec.Object {
	return o.Object
}



// MTL4ArchiveObjectFromID constructs a [MTL4ArchiveObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTL4ArchiveObjectFromID(id objc.ID) MTL4ArchiveObject {
	return MTL4ArchiveObject{
		Object: objectivec.ObjectFromID(id),
	}
}




// A label that you can associate with this archive.
//
// See: https://developer.apple.com/documentation/Metal/MTL4Archive/label

func (o MTL4ArchiveObject) Label() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
	}

// Synchronously creates a binary version of a GPU visible function or GPU
// intersection function.
//
// descriptor: A configuration that tells the method which GPU function to make into a
// binary function and which options to apply when compiling it.
//
// # Return Value
// 
// A new GPU binary function instance if the method succeeds; otherwise `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTL4Archive/makeBinaryFunction(descriptor:)

func (o MTL4ArchiveObject) NewBinaryFunctionWithDescriptorError(descriptor IMTL4BinaryFunctionDescriptor) (MTL4BinaryFunction, error) {
	
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("newBinaryFunctionWithDescriptor:error:"), descriptor)
	if err != nil {
		return nil, err
	}
	return MTL4BinaryFunctionObjectFromID(rv), nil
	}

// Creates a compute pipeline state from the archive with a descriptor.
//
// descriptor: A compute pipeline descriptor.
//
// error: On return, if the method fails, a pointer to an error information instance;
// otherwise `nil`.
//
// # Return Value
// 
// A compute pipeline state if the method succeeds, otherwise `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTL4Archive/newComputePipelineStateWithDescriptor:error:

func (o MTL4ArchiveObject) NewComputePipelineStateWithDescriptorError(descriptor IMTL4ComputePipelineDescriptor) (MTLComputePipelineState, error) {
	
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("newComputePipelineStateWithDescriptor:error:"), descriptor)
	if err != nil {
		return nil, err
	}
	return MTLComputePipelineStateObjectFromID(rv), nil
	}

// Creates a compute pipeline state from the archive with a compute descriptor
// and a dynamic linking descriptor.
//
// descriptor: A compute pipeline descriptor.
//
// dynamicLinkingDescriptor: A descriptor that provides additional properties to link other functions
// with the pipeline.
//
// error: On return, if the method fails, a pointer to an error information instance;
// otherwise `nil`.
//
// # Return Value
// 
// A compute pipeline state if the method succeeds, otherwise `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTL4Archive/newComputePipelineStateWithDescriptor:dynamicLinkingDescriptor:error:

func (o MTL4ArchiveObject) NewComputePipelineStateWithDescriptorDynamicLinkingDescriptorError(descriptor IMTL4ComputePipelineDescriptor, dynamicLinkingDescriptor IMTL4PipelineStageDynamicLinkingDescriptor) (MTLComputePipelineState, error) {
	
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("newComputePipelineStateWithDescriptor:dynamicLinkingDescriptor:error:"), descriptor, dynamicLinkingDescriptor)
	if err != nil {
		return nil, err
	}
	return MTLComputePipelineStateObjectFromID(rv), nil
	}

// Creates a render pipeline state from the archive with a descriptor.
//
// descriptor: A render pipeline descriptor.
//
// error: On return, if the method fails, a pointer to an error information instance;
// otherwise `nil`.
//
// # Return Value
// 
// A render pipeline state if the method succeeds, otherwise `nil`.
//
// # Discussion
// 
// You create any kind of render pipeline states with this method, including:
// 
// - Traditional render pipelines - Mesh render pipelines - Tile render
// pipelines
//
// See: https://developer.apple.com/documentation/Metal/MTL4Archive/newRenderPipelineStateWithDescriptor:error:

func (o MTL4ArchiveObject) NewRenderPipelineStateWithDescriptorError(descriptor IMTL4PipelineDescriptor) (MTLRenderPipelineState, error) {
	
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("newRenderPipelineStateWithDescriptor:error:"), descriptor)
	if err != nil {
		return nil, err
	}
	return MTLRenderPipelineStateObjectFromID(rv), nil
	}

// Creates a render pipeline state from the archive with a render descriptor
// and a dynamic linking descriptor.
//
// descriptor: A render pipeline descriptor.
//
// dynamicLinkingDescriptor: A descriptor that provides additional properties to link other functions
// with the pipeline.
//
// error: On return, if the method fails, a pointer to an error information instance;
// otherwise `nil`.
//
// # Return Value
// 
// A render pipeline state if the method succeeds, otherwise `nil`.
//
// # Discussion
// 
// You create any kind of render pipeline states with this method, including:
// 
// - Traditional render pipelines - Mesh render pipelines - Tile render
// pipelines
//
// See: https://developer.apple.com/documentation/Metal/MTL4Archive/newRenderPipelineStateWithDescriptor:dynamicLinkingDescriptor:error:

func (o MTL4ArchiveObject) NewRenderPipelineStateWithDescriptorDynamicLinkingDescriptorError(descriptor IMTL4PipelineDescriptor, dynamicLinkingDescriptor IMTL4RenderPipelineDynamicLinkingDescriptor) (MTLRenderPipelineState, error) {
	
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("newRenderPipelineStateWithDescriptor:dynamicLinkingDescriptor:error:"), descriptor, dynamicLinkingDescriptor)
	if err != nil {
		return nil, err
	}
	return MTLRenderPipelineStateObjectFromID(rv), nil
	}




func (o MTL4ArchiveObject) SetLabel(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setLabel:"), objc.String(value))
}





