// Code generated from Apple documentation for MetalKit. DO NOT EDIT.

package metalkit

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTKMesh] class.
var (
	_MTKMeshClass     MTKMeshClass
	_MTKMeshClassOnce sync.Once
)

func getMTKMeshClass() MTKMeshClass {
	_MTKMeshClassOnce.Do(func() {
		_MTKMeshClass = MTKMeshClass{class: objc.GetClass("MTKMesh")}
	})
	return _MTKMeshClass
}

// GetMTKMeshClass returns the class object for MTKMesh.
func GetMTKMeshClass() MTKMeshClass {
	return getMTKMeshClass()
}

type MTKMeshClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTKMeshClass) Alloc() MTKMesh {
	rv := objc.Send[MTKMesh](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A container for the vertex data of a Model I/O mesh, suitable for use in a
// Metal app.
//
// # Initialization
//
//   - [MTKMesh.InitWithMeshDeviceError]: Initializes a MetalKit mesh and its submeshes from a Model I/O mesh.
//
// # Submeshes
//
//   - [MTKMesh.Submeshes]: An array of submeshes containing index buffers referencing the mesh vertices.
//
// # Vertex Properties
//
//   - [MTKMesh.VertexBuffers]: An array of buffers in which mesh vertex data resides.
//   - [MTKMesh.VertexCount]: The number of vertices in the vertex buffers.
//   - [MTKMesh.VertexDescriptor]: A Model I/O vertex descriptor specifying the data layout in the vertex buffers.
//
// # Identifying Properties
//
//   - [MTKMesh.Name]: The name of the mesh.
//   - [MTKMesh.SetName]
//
// See: https://developer.apple.com/documentation/MetalKit/MTKMesh
type MTKMesh struct {
	objectivec.Object
}

// MTKMeshFromID constructs a [MTKMesh] from an objc.ID.
//
// A container for the vertex data of a Model I/O mesh, suitable for use in a
// Metal app.
func MTKMeshFromID(id objc.ID) MTKMesh {
	return MTKMesh{objectivec.Object{ID: id}}
}
// NOTE: MTKMesh adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTKMesh] class.
//
// # Initialization
//
//   - [IMTKMesh.InitWithMeshDeviceError]: Initializes a MetalKit mesh and its submeshes from a Model I/O mesh.
//
// # Submeshes
//
//   - [IMTKMesh.Submeshes]: An array of submeshes containing index buffers referencing the mesh vertices.
//
// # Vertex Properties
//
//   - [IMTKMesh.VertexBuffers]: An array of buffers in which mesh vertex data resides.
//   - [IMTKMesh.VertexCount]: The number of vertices in the vertex buffers.
//   - [IMTKMesh.VertexDescriptor]: A Model I/O vertex descriptor specifying the data layout in the vertex buffers.
//
// # Identifying Properties
//
//   - [IMTKMesh.Name]: The name of the mesh.
//   - [IMTKMesh.SetName]
//
// See: https://developer.apple.com/documentation/MetalKit/MTKMesh
type IMTKMesh interface {
	objectivec.IObject

	// Topic: Initialization

	// Initializes a MetalKit mesh and its submeshes from a Model I/O mesh.
	InitWithMeshDeviceError(mesh objectivec.IObject, device metal.MTLDevice) (MTKMesh, error)

	// Topic: Submeshes

	// An array of submeshes containing index buffers referencing the mesh vertices.
	Submeshes() []MTKSubmesh

	// Topic: Vertex Properties

	// An array of buffers in which mesh vertex data resides.
	VertexBuffers() []MTKMeshBuffer
	// The number of vertices in the vertex buffers.
	VertexCount() uint
	// A Model I/O vertex descriptor specifying the data layout in the vertex buffers.
	VertexDescriptor() objc.ID

	// Topic: Identifying Properties

	// The name of the mesh.
	Name() string
	SetName(value string)
}

// Init initializes the instance.
func (m MTKMesh) Init() MTKMesh {
	rv := objc.Send[MTKMesh](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MTKMesh) Autorelease() MTKMesh {
	rv := objc.Send[MTKMesh](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTKMesh creates a new MTKMesh instance.
func NewMTKMesh() MTKMesh {
	class := getMTKMeshClass()
	rv := objc.Send[MTKMesh](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a MetalKit mesh and its submeshes from a Model I/O mesh.
//
// mesh: The source Model I/O mesh from which to create this MetalKit mesh.
//
// device: The Metal device on which to create MetalKit mesh resources.
//
// # Return Value
// 
// A new MetalKit mesh object, or `nil` if an error occured.
//
// # Discussion
// 
// This initializer does initialize any children meshes of the Model I/O mesh.
// 
// All vertex buffers in the source Model I/O mesh and the index buffer of
// each of its submeshes must have been created with a
// [MTKMeshBufferAllocator] object.
// 
// A Model I/O submesh may have its index type and/or geometric primitive type
// converted to a corresponding Metal type as listed in the tables below. If
// the geometric primitive type cannot be converted, an error is returned
// through `error`.
// 
// [Table data omitted]
// 
// [Table data omitted]
//
// See: https://developer.apple.com/documentation/MetalKit/MTKMesh/init(mesh:device:)
func NewMeshWithMeshDeviceError(mesh objectivec.IObject, device metal.MTLDevice) (MTKMesh, error) {
	var errorPtr objc.ID
	instance := getMTKMeshClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMesh:device:error:"), mesh, device, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MTKMeshFromID(rv), foundation.NSErrorFrom(errorPtr)
	}
	return MTKMeshFromID(rv), nil
}

// Initializes a MetalKit mesh and its submeshes from a Model I/O mesh.
//
// mesh: The source Model I/O mesh from which to create this MetalKit mesh.
//
// device: The Metal device on which to create MetalKit mesh resources.
//
// # Return Value
// 
// A new MetalKit mesh object, or `nil` if an error occured.
//
// # Discussion
// 
// This initializer does initialize any children meshes of the Model I/O mesh.
// 
// All vertex buffers in the source Model I/O mesh and the index buffer of
// each of its submeshes must have been created with a
// [MTKMeshBufferAllocator] object.
// 
// A Model I/O submesh may have its index type and/or geometric primitive type
// converted to a corresponding Metal type as listed in the tables below. If
// the geometric primitive type cannot be converted, an error is returned
// through `error`.
// 
// [Table data omitted]
// 
// [Table data omitted]
//
// See: https://developer.apple.com/documentation/MetalKit/MTKMesh/init(mesh:device:)
func (m MTKMesh) InitWithMeshDeviceError(mesh objectivec.IObject, device metal.MTLDevice) (MTKMesh, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithMesh:device:error:"), mesh, device, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MTKMesh{}, foundation.NSErrorFrom(errorPtr)
	}
	return MTKMeshFromID(rv), nil

}

// An array of submeshes containing index buffers referencing the mesh
// vertices.
//
// # Discussion
// 
// Submeshes may also contain texture materials to apply when rendering the
// mesh object.
//
// See: https://developer.apple.com/documentation/MetalKit/MTKMesh/submeshes
func (m MTKMesh) Submeshes() []MTKSubmesh {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("submeshes"))
	return objc.ConvertSlice(rv, func(id objc.ID) MTKSubmesh {
		return MTKSubmeshFromID(id)
	})
}
// An array of buffers in which mesh vertex data resides.
//
// See: https://developer.apple.com/documentation/MetalKit/MTKMesh/vertexBuffers
func (m MTKMesh) VertexBuffers() []MTKMeshBuffer {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("vertexBuffers"))
	return objc.ConvertSlice(rv, func(id objc.ID) MTKMeshBuffer {
		return MTKMeshBufferFromID(id)
	})
}
// The number of vertices in the vertex buffers.
//
// See: https://developer.apple.com/documentation/MetalKit/MTKMesh/vertexCount
func (m MTKMesh) VertexCount() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("vertexCount"))
	return rv
}
// A Model I/O vertex descriptor specifying the data layout in the vertex
// buffers.
//
// # Discussion
// 
// This is a convenience property. The [MTKMesh] class does not use this
// descriptor, but your application may use this object to determine rendering
// state or create a [MTLVertexDescriptor] object to build a
// [MTLRenderPipelineState] object capable of interpreting the data in
// [VertexBuffers].
//
// [MTLRenderPipelineState]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineState
// [MTLVertexDescriptor]: https://developer.apple.com/documentation/Metal/MTLVertexDescriptor
//
// See: https://developer.apple.com/documentation/MetalKit/MTKMesh/vertexDescriptor
func (m MTKMesh) VertexDescriptor() objc.ID {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("vertexDescriptor"))
	return rv
}
// The name of the mesh.
//
// # Discussion
// 
// Your application can use this value to identify the mesh in a scene.
//
// See: https://developer.apple.com/documentation/MetalKit/MTKMesh/name
func (m MTKMesh) Name() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
func (m MTKMesh) SetName(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setName:"), objc.String(value))
}

