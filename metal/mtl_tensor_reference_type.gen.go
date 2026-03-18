// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLTensorReferenceType] class.
var (
	_MTLTensorReferenceTypeClass     MTLTensorReferenceTypeClass
	_MTLTensorReferenceTypeClassOnce sync.Once
)

func getMTLTensorReferenceTypeClass() MTLTensorReferenceTypeClass {
	_MTLTensorReferenceTypeClassOnce.Do(func() {
		_MTLTensorReferenceTypeClass = MTLTensorReferenceTypeClass{class: objc.GetClass("MTLTensorReferenceType")}
	})
	return _MTLTensorReferenceTypeClass
}

// GetMTLTensorReferenceTypeClass returns the class object for MTLTensorReferenceType.
func GetMTLTensorReferenceTypeClass() MTLTensorReferenceTypeClass {
	return getMTLTensorReferenceTypeClass()
}

type MTLTensorReferenceTypeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLTensorReferenceTypeClass) Alloc() MTLTensorReferenceType {
	rv := objc.Send[MTLTensorReferenceType](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a tensor in the shading language in a struct or
// array.
//
// # Instance Properties
//
//   - [MTLTensorReferenceType.Access]: A value that represents the read/write permissions of the tensor.
//   - [MTLTensorReferenceType.Dimensions]: The array of sizes, in elements, one for each dimension of this tensor.
//   - [MTLTensorReferenceType.IndexType]: The data format you use for indexing into the tensor.
//   - [MTLTensorReferenceType.TensorDataType]: The underlying data format of the tensor.
//
// See: https://developer.apple.com/documentation/Metal/MTLTensorReferenceType
type MTLTensorReferenceType struct {
	MTLType
}

// MTLTensorReferenceTypeFromID constructs a [MTLTensorReferenceType] from an objc.ID.
//
// An object that represents a tensor in the shading language in a struct or
// array.
func MTLTensorReferenceTypeFromID(id objc.ID) MTLTensorReferenceType {
	return MTLTensorReferenceType{MTLType: MTLTypeFromID(id)}
}
// NOTE: MTLTensorReferenceType adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLTensorReferenceType] class.
//
// # Instance Properties
//
//   - [IMTLTensorReferenceType.Access]: A value that represents the read/write permissions of the tensor.
//   - [IMTLTensorReferenceType.Dimensions]: The array of sizes, in elements, one for each dimension of this tensor.
//   - [IMTLTensorReferenceType.IndexType]: The data format you use for indexing into the tensor.
//   - [IMTLTensorReferenceType.TensorDataType]: The underlying data format of the tensor.
//
// See: https://developer.apple.com/documentation/Metal/MTLTensorReferenceType
type IMTLTensorReferenceType interface {
	IMTLType

	// Topic: Instance Properties

	// A value that represents the read/write permissions of the tensor.
	Access() MTLBindingAccess
	// The array of sizes, in elements, one for each dimension of this tensor.
	Dimensions() IMTLTensorExtents
	// The data format you use for indexing into the tensor.
	IndexType() MTLDataType
	// The underlying data format of the tensor.
	TensorDataType() MTLTensorDataType

	// An error domain for errors that pertain to creating a tensor.
	MTLTensorDomain() string
	MTL_TENSOR_MAX_RANK() objectivec.IObject
	SetMTL_TENSOR_MAX_RANK(value objectivec.IObject)
}

// Init initializes the instance.
func (t MTLTensorReferenceType) Init() MTLTensorReferenceType {
	rv := objc.Send[MTLTensorReferenceType](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t MTLTensorReferenceType) Autorelease() MTLTensorReferenceType {
	rv := objc.Send[MTLTensorReferenceType](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLTensorReferenceType creates a new MTLTensorReferenceType instance.
func NewMTLTensorReferenceType() MTLTensorReferenceType {
	class := getMTLTensorReferenceTypeClass()
	rv := objc.Send[MTLTensorReferenceType](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A value that represents the read/write permissions of the tensor.
//
// See: https://developer.apple.com/documentation/Metal/MTLTensorReferenceType/access
func (t MTLTensorReferenceType) Access() MTLBindingAccess {
	rv := objc.Send[MTLBindingAccess](t.ID, objc.Sel("access"))
	return MTLBindingAccess(rv)
}

// The array of sizes, in elements, one for each dimension of this tensor.
//
// # Discussion
// 
// Because shader-bound tensors have dynamic extents, the [Rank] of
// `dimensions` corresponds to the rank the shader function specifies, and
// `MTLTensorExtents/` always returns a value of -1.
//
// See: https://developer.apple.com/documentation/Metal/MTLTensorReferenceType/dimensions
func (t MTLTensorReferenceType) Dimensions() IMTLTensorExtents {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("dimensions"))
	return MTLTensorExtentsFromID(objc.ID(rv))
}

// The data format you use for indexing into the tensor.
//
// See: https://developer.apple.com/documentation/Metal/MTLTensorReferenceType/indexType
func (t MTLTensorReferenceType) IndexType() MTLDataType {
	rv := objc.Send[MTLDataType](t.ID, objc.Sel("indexType"))
	return MTLDataType(rv)
}

// The underlying data format of the tensor.
//
// See: https://developer.apple.com/documentation/Metal/MTLTensorReferenceType/tensorDataType
func (t MTLTensorReferenceType) TensorDataType() MTLTensorDataType {
	rv := objc.Send[MTLTensorDataType](t.ID, objc.Sel("tensorDataType"))
	return MTLTensorDataType(rv)
}

// An error domain for errors that pertain to creating a tensor.
//
// See: https://developer.apple.com/documentation/metal/mtltensordomain
func (t MTLTensorReferenceType) MTLTensorDomain() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("MTLTensorDomain"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/metal/mtl_tensor_max_rank
func (t MTLTensorReferenceType) MTL_TENSOR_MAX_RANK() objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("MTL_TENSOR_MAX_RANK"))
	return objectivec.Object{ID: rv}
}
func (t MTLTensorReferenceType) SetMTL_TENSOR_MAX_RANK(value objectivec.IObject) {
	objc.Send[struct{}](t.ID, objc.Sel("setMTL_TENSOR_MAX_RANK:"), value)
}

