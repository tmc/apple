// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// A container for pipeline state descriptors and their associated compiled shader code.
//
// See: https://developer.apple.com/documentation/Metal/MTLBinaryArchive
type MTLBinaryArchive interface {
	objectivec.IObject

	// The Metal device object that created the binary archive.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLBinaryArchive/device
	Device() MTLDevice

	// A string that identifies the library.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLBinaryArchive/label
	Label() string

	// Adds a description of a compute pipeline to the archive.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLBinaryArchive/addComputePipelineFunctions(descriptor:)
	AddComputePipelineFunctionsWithDescriptorError(descriptor IMTLComputePipelineDescriptor) (bool, error)

	// Adds a description of a render pipeline to the archive.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLBinaryArchive/addRenderPipelineFunctions(descriptor:)
	AddRenderPipelineFunctionsWithDescriptorError(descriptor IMTLRenderPipelineDescriptor) (bool, error)

	// Adds a description of a tile renderer pipeline to the archive.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLBinaryArchive/addTileRenderPipelineFunctions(descriptor:)
	AddTileRenderPipelineFunctionsWithDescriptorError(descriptor IMTLTileRenderPipelineDescriptor) (bool, error)

	// Adds a description of a function to the archive.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLBinaryArchive/addFunction(descriptor:library:)
	AddFunctionWithDescriptorLibraryError(descriptor IMTLFunctionDescriptor, library MTLLibrary) (bool, error)

	// Writes the contents of the archive to a file.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLBinaryArchive/serialize(to:)
	SerializeToURLError(url foundation.INSURL) (bool, error)

	// AddLibraryWithDescriptorError protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLBinaryArchive/addLibrary(descriptor:)
	AddLibraryWithDescriptorError(descriptor IMTLStitchedLibraryDescriptor) (bool, error)

	// AddMeshRenderPipelineFunctionsWithDescriptorError protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLBinaryArchive/addMeshRenderPipelineFunctions(descriptor:)
	AddMeshRenderPipelineFunctionsWithDescriptorError(descriptor IMTLMeshRenderPipelineDescriptor) (bool, error)

	// A string that identifies the library.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLBinaryArchive/label
	SetLabel(value string)
}

// MTLBinaryArchiveObject wraps an existing Objective-C object that conforms to the MTLBinaryArchive protocol.
type MTLBinaryArchiveObject struct {
	objectivec.Object
}
func (o MTLBinaryArchiveObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTLBinaryArchiveObjectFromID constructs a [MTLBinaryArchiveObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLBinaryArchiveObjectFromID(id objc.ID) MTLBinaryArchiveObject {
	return MTLBinaryArchiveObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The Metal device object that created the binary archive.
//
// See: https://developer.apple.com/documentation/Metal/MTLBinaryArchive/device
func (o MTLBinaryArchiveObject) Device() MTLDevice {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("device"))
	return MTLDeviceObjectFromID(rv)
	}
// A string that identifies the library.
//
// See: https://developer.apple.com/documentation/Metal/MTLBinaryArchive/label
func (o MTLBinaryArchiveObject) Label() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
	}
// Adds a description of a compute pipeline to the archive.
//
// descriptor: A description of the compute pipeline to archive.
//
// See: https://developer.apple.com/documentation/Metal/MTLBinaryArchive/addComputePipelineFunctions(descriptor:)
func (o MTLBinaryArchiveObject) AddComputePipelineFunctionsWithDescriptorError(descriptor IMTLComputePipelineDescriptor) (bool, error) {
	
	rv, err := objc.SendWithError[bool](o.ID, objc.Sel("addComputePipelineFunctionsWithDescriptor:error:"), descriptor)
	if err != nil {
		return false, err
	}
	return rv, nil
	}
// Adds a description of a render pipeline to the archive.
//
// descriptor: A description of the render pipeline to archive.
//
// See: https://developer.apple.com/documentation/Metal/MTLBinaryArchive/addRenderPipelineFunctions(descriptor:)
func (o MTLBinaryArchiveObject) AddRenderPipelineFunctionsWithDescriptorError(descriptor IMTLRenderPipelineDescriptor) (bool, error) {
	
	rv, err := objc.SendWithError[bool](o.ID, objc.Sel("addRenderPipelineFunctionsWithDescriptor:error:"), descriptor)
	if err != nil {
		return false, err
	}
	return rv, nil
	}
// Adds a description of a tile renderer pipeline to the archive.
//
// descriptor: A description of the tile renderer pipeline to archive.
//
// See: https://developer.apple.com/documentation/Metal/MTLBinaryArchive/addTileRenderPipelineFunctions(descriptor:)
func (o MTLBinaryArchiveObject) AddTileRenderPipelineFunctionsWithDescriptorError(descriptor IMTLTileRenderPipelineDescriptor) (bool, error) {
	
	rv, err := objc.SendWithError[bool](o.ID, objc.Sel("addTileRenderPipelineFunctionsWithDescriptor:error:"), descriptor)
	if err != nil {
		return false, err
	}
	return rv, nil
	}
// Adds a description of a function to the archive.
//
// See: https://developer.apple.com/documentation/Metal/MTLBinaryArchive/addFunction(descriptor:library:)
func (o MTLBinaryArchiveObject) AddFunctionWithDescriptorLibraryError(descriptor IMTLFunctionDescriptor, library MTLLibrary) (bool, error) {
	
	rv, err := objc.SendWithError[bool](o.ID, objc.Sel("addFunctionWithDescriptor:library:error:"), descriptor, library)
	if err != nil {
		return false, err
	}
	return rv, nil
	}
// Writes the contents of the archive to a file.
//
// url: The URL for the destination file.
//
// # Discussion
// 
// The destination folder needs to exist when you call this method.
//
// See: https://developer.apple.com/documentation/Metal/MTLBinaryArchive/serialize(to:)
func (o MTLBinaryArchiveObject) SerializeToURLError(url foundation.INSURL) (bool, error) {
	
	rv, err := objc.SendWithError[bool](o.ID, objc.Sel("serializeToURL:error:"), url)
	if err != nil {
		return false, err
	}
	return rv, nil
	}
//
// See: https://developer.apple.com/documentation/Metal/MTLBinaryArchive/addLibrary(descriptor:)
func (o MTLBinaryArchiveObject) AddLibraryWithDescriptorError(descriptor IMTLStitchedLibraryDescriptor) (bool, error) {
	
	rv, err := objc.SendWithError[bool](o.ID, objc.Sel("addLibraryWithDescriptor:error:"), descriptor)
	if err != nil {
		return false, err
	}
	return rv, nil
	}
//
// See: https://developer.apple.com/documentation/Metal/MTLBinaryArchive/addMeshRenderPipelineFunctions(descriptor:)
func (o MTLBinaryArchiveObject) AddMeshRenderPipelineFunctionsWithDescriptorError(descriptor IMTLMeshRenderPipelineDescriptor) (bool, error) {
	
	rv, err := objc.SendWithError[bool](o.ID, objc.Sel("addMeshRenderPipelineFunctionsWithDescriptor:error:"), descriptor)
	if err != nil {
		return false, err
	}
	return rv, nil
	}

func (o MTLBinaryArchiveObject) SetLabel(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setLabel:"), objc.String(value))
}

