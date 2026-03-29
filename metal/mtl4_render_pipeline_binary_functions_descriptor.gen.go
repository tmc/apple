// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTL4RenderPipelineBinaryFunctionsDescriptor] class.
var (
	_MTL4RenderPipelineBinaryFunctionsDescriptorClass     MTL4RenderPipelineBinaryFunctionsDescriptorClass
	_MTL4RenderPipelineBinaryFunctionsDescriptorClassOnce sync.Once
)

func getMTL4RenderPipelineBinaryFunctionsDescriptorClass() MTL4RenderPipelineBinaryFunctionsDescriptorClass {
	_MTL4RenderPipelineBinaryFunctionsDescriptorClassOnce.Do(func() {
		_MTL4RenderPipelineBinaryFunctionsDescriptorClass = MTL4RenderPipelineBinaryFunctionsDescriptorClass{class: objc.GetClass("MTL4RenderPipelineBinaryFunctionsDescriptor")}
	})
	return _MTL4RenderPipelineBinaryFunctionsDescriptorClass
}

// GetMTL4RenderPipelineBinaryFunctionsDescriptorClass returns the class object for MTL4RenderPipelineBinaryFunctionsDescriptor.
func GetMTL4RenderPipelineBinaryFunctionsDescriptorClass() MTL4RenderPipelineBinaryFunctionsDescriptorClass {
	return getMTL4RenderPipelineBinaryFunctionsDescriptorClass()
}

type MTL4RenderPipelineBinaryFunctionsDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTL4RenderPipelineBinaryFunctionsDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTL4RenderPipelineBinaryFunctionsDescriptorClass) Alloc() MTL4RenderPipelineBinaryFunctionsDescriptor {
	rv := objc.Send[MTL4RenderPipelineBinaryFunctionsDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// Allows you to specify additional binary functions to link to each stage of
// a render pipeline.
//
// # Instance Properties
//
//   - [MTL4RenderPipelineBinaryFunctionsDescriptor.FragmentAdditionalBinaryFunctions]: Provides an array of binary functions representing additional binary fragment shader functions.
//   - [MTL4RenderPipelineBinaryFunctionsDescriptor.SetFragmentAdditionalBinaryFunctions]
//   - [MTL4RenderPipelineBinaryFunctionsDescriptor.MeshAdditionalBinaryFunctions]: Provides an array of binary functions representing additional binary mesh shader functions.
//   - [MTL4RenderPipelineBinaryFunctionsDescriptor.SetMeshAdditionalBinaryFunctions]
//   - [MTL4RenderPipelineBinaryFunctionsDescriptor.ObjectAdditionalBinaryFunctions]: Provides an array of binary functions representing additional binary object shader functions.
//   - [MTL4RenderPipelineBinaryFunctionsDescriptor.SetObjectAdditionalBinaryFunctions]
//   - [MTL4RenderPipelineBinaryFunctionsDescriptor.TileAdditionalBinaryFunctions]: Provides an array of binary functions representing additional binary tile shader functions.
//   - [MTL4RenderPipelineBinaryFunctionsDescriptor.SetTileAdditionalBinaryFunctions]
//   - [MTL4RenderPipelineBinaryFunctionsDescriptor.VertexAdditionalBinaryFunctions]: Provides an array of binary functions representing additional binary vertex shader functions.
//   - [MTL4RenderPipelineBinaryFunctionsDescriptor.SetVertexAdditionalBinaryFunctions]
//
// # Instance Methods
//
//   - [MTL4RenderPipelineBinaryFunctionsDescriptor.Reset]: Resets this descriptor to its default state.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineBinaryFunctionsDescriptor
type MTL4RenderPipelineBinaryFunctionsDescriptor struct {
	objectivec.Object
}

// MTL4RenderPipelineBinaryFunctionsDescriptorFromID constructs a [MTL4RenderPipelineBinaryFunctionsDescriptor] from an objc.ID.
//
// Allows you to specify additional binary functions to link to each stage of
// a render pipeline.
func MTL4RenderPipelineBinaryFunctionsDescriptorFromID(id objc.ID) MTL4RenderPipelineBinaryFunctionsDescriptor {
	return MTL4RenderPipelineBinaryFunctionsDescriptor{objectivec.Object{ID: id}}
}
// NOTE: MTL4RenderPipelineBinaryFunctionsDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTL4RenderPipelineBinaryFunctionsDescriptor] class.
//
// # Instance Properties
//
//   - [IMTL4RenderPipelineBinaryFunctionsDescriptor.FragmentAdditionalBinaryFunctions]: Provides an array of binary functions representing additional binary fragment shader functions.
//   - [IMTL4RenderPipelineBinaryFunctionsDescriptor.SetFragmentAdditionalBinaryFunctions]
//   - [IMTL4RenderPipelineBinaryFunctionsDescriptor.MeshAdditionalBinaryFunctions]: Provides an array of binary functions representing additional binary mesh shader functions.
//   - [IMTL4RenderPipelineBinaryFunctionsDescriptor.SetMeshAdditionalBinaryFunctions]
//   - [IMTL4RenderPipelineBinaryFunctionsDescriptor.ObjectAdditionalBinaryFunctions]: Provides an array of binary functions representing additional binary object shader functions.
//   - [IMTL4RenderPipelineBinaryFunctionsDescriptor.SetObjectAdditionalBinaryFunctions]
//   - [IMTL4RenderPipelineBinaryFunctionsDescriptor.TileAdditionalBinaryFunctions]: Provides an array of binary functions representing additional binary tile shader functions.
//   - [IMTL4RenderPipelineBinaryFunctionsDescriptor.SetTileAdditionalBinaryFunctions]
//   - [IMTL4RenderPipelineBinaryFunctionsDescriptor.VertexAdditionalBinaryFunctions]: Provides an array of binary functions representing additional binary vertex shader functions.
//   - [IMTL4RenderPipelineBinaryFunctionsDescriptor.SetVertexAdditionalBinaryFunctions]
//
// # Instance Methods
//
//   - [IMTL4RenderPipelineBinaryFunctionsDescriptor.Reset]: Resets this descriptor to its default state.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineBinaryFunctionsDescriptor
type IMTL4RenderPipelineBinaryFunctionsDescriptor interface {
	objectivec.IObject

	// Topic: Instance Properties

	// Provides an array of binary functions representing additional binary fragment shader functions.
	FragmentAdditionalBinaryFunctions() []objectivec.IObject
	SetFragmentAdditionalBinaryFunctions(value []objectivec.IObject)
	// Provides an array of binary functions representing additional binary mesh shader functions.
	MeshAdditionalBinaryFunctions() []objectivec.IObject
	SetMeshAdditionalBinaryFunctions(value []objectivec.IObject)
	// Provides an array of binary functions representing additional binary object shader functions.
	ObjectAdditionalBinaryFunctions() []objectivec.IObject
	SetObjectAdditionalBinaryFunctions(value []objectivec.IObject)
	// Provides an array of binary functions representing additional binary tile shader functions.
	TileAdditionalBinaryFunctions() []objectivec.IObject
	SetTileAdditionalBinaryFunctions(value []objectivec.IObject)
	// Provides an array of binary functions representing additional binary vertex shader functions.
	VertexAdditionalBinaryFunctions() []objectivec.IObject
	SetVertexAdditionalBinaryFunctions(value []objectivec.IObject)

	// Topic: Instance Methods

	// Resets this descriptor to its default state.
	Reset()
}

// Init initializes the instance.
func (m MTL4RenderPipelineBinaryFunctionsDescriptor) Init() MTL4RenderPipelineBinaryFunctionsDescriptor {
	rv := objc.Send[MTL4RenderPipelineBinaryFunctionsDescriptor](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MTL4RenderPipelineBinaryFunctionsDescriptor) Autorelease() MTL4RenderPipelineBinaryFunctionsDescriptor {
	rv := objc.Send[MTL4RenderPipelineBinaryFunctionsDescriptor](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4RenderPipelineBinaryFunctionsDescriptor creates a new MTL4RenderPipelineBinaryFunctionsDescriptor instance.
func NewMTL4RenderPipelineBinaryFunctionsDescriptor() MTL4RenderPipelineBinaryFunctionsDescriptor {
	class := getMTL4RenderPipelineBinaryFunctionsDescriptorClass()
	rv := objc.Send[MTL4RenderPipelineBinaryFunctionsDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Resets this descriptor to its default state.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineBinaryFunctionsDescriptor/reset()
func (m MTL4RenderPipelineBinaryFunctionsDescriptor) Reset() {
	objc.Send[objc.ID](m.ID, objc.Sel("reset"))
}

// Provides an array of binary functions representing additional binary
// fragment shader functions.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineBinaryFunctionsDescriptor/fragmentAdditionalBinaryFunctions
func (m MTL4RenderPipelineBinaryFunctionsDescriptor) FragmentAdditionalBinaryFunctions() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("fragmentAdditionalBinaryFunctions"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}
func (m MTL4RenderPipelineBinaryFunctionsDescriptor) SetFragmentAdditionalBinaryFunctions(value []objectivec.IObject) {
	objc.Send[struct{}](m.ID, objc.Sel("setFragmentAdditionalBinaryFunctions:"), objectivec.IObjectSliceToNSArray(value))
}
// Provides an array of binary functions representing additional binary mesh
// shader functions.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineBinaryFunctionsDescriptor/meshAdditionalBinaryFunctions
func (m MTL4RenderPipelineBinaryFunctionsDescriptor) MeshAdditionalBinaryFunctions() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("meshAdditionalBinaryFunctions"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}
func (m MTL4RenderPipelineBinaryFunctionsDescriptor) SetMeshAdditionalBinaryFunctions(value []objectivec.IObject) {
	objc.Send[struct{}](m.ID, objc.Sel("setMeshAdditionalBinaryFunctions:"), objectivec.IObjectSliceToNSArray(value))
}
// Provides an array of binary functions representing additional binary object
// shader functions.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineBinaryFunctionsDescriptor/objectAdditionalBinaryFunctions
func (m MTL4RenderPipelineBinaryFunctionsDescriptor) ObjectAdditionalBinaryFunctions() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("objectAdditionalBinaryFunctions"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}
func (m MTL4RenderPipelineBinaryFunctionsDescriptor) SetObjectAdditionalBinaryFunctions(value []objectivec.IObject) {
	objc.Send[struct{}](m.ID, objc.Sel("setObjectAdditionalBinaryFunctions:"), objectivec.IObjectSliceToNSArray(value))
}
// Provides an array of binary functions representing additional binary tile
// shader functions.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineBinaryFunctionsDescriptor/tileAdditionalBinaryFunctions
func (m MTL4RenderPipelineBinaryFunctionsDescriptor) TileAdditionalBinaryFunctions() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("tileAdditionalBinaryFunctions"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}
func (m MTL4RenderPipelineBinaryFunctionsDescriptor) SetTileAdditionalBinaryFunctions(value []objectivec.IObject) {
	objc.Send[struct{}](m.ID, objc.Sel("setTileAdditionalBinaryFunctions:"), objectivec.IObjectSliceToNSArray(value))
}
// Provides an array of binary functions representing additional binary vertex
// shader functions.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineBinaryFunctionsDescriptor/vertexAdditionalBinaryFunctions
func (m MTL4RenderPipelineBinaryFunctionsDescriptor) VertexAdditionalBinaryFunctions() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("vertexAdditionalBinaryFunctions"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}
func (m MTL4RenderPipelineBinaryFunctionsDescriptor) SetVertexAdditionalBinaryFunctions(value []objectivec.IObject) {
	objc.Send[struct{}](m.ID, objc.Sel("setVertexAdditionalBinaryFunctions:"), objectivec.IObjectSliceToNSArray(value))
}

