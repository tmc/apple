// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLTensorExtents] class.
var (
	_MTLTensorExtentsClass     MTLTensorExtentsClass
	_MTLTensorExtentsClassOnce sync.Once
)

func getMTLTensorExtentsClass() MTLTensorExtentsClass {
	_MTLTensorExtentsClassOnce.Do(func() {
		_MTLTensorExtentsClass = MTLTensorExtentsClass{class: objc.GetClass("MTLTensorExtents")}
	})
	return _MTLTensorExtentsClass
}

// GetMTLTensorExtentsClass returns the class object for MTLTensorExtents.
func GetMTLTensorExtentsClass() MTLTensorExtentsClass {
	return getMTLTensorExtentsClass()
}

type MTLTensorExtentsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTLTensorExtentsClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLTensorExtentsClass) Alloc() MTLTensorExtents {
	rv := objc.Send[MTLTensorExtents](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An array of length matching the rank, holding the dimensions of a tensor.
//
// # Overview
// 
// Supports rank up to [MTL_TENSOR_MAX_RANK].
//
// [MTL_TENSOR_MAX_RANK]: https://developer.apple.com/documentation/Metal/MTL_TENSOR_MAX_RANK
//
// # Instance Properties
//
//   - [MTLTensorExtents.Extents]: Retrieves the extents for this object.
//   - [MTLTensorExtents.SetExtents]
//   - [MTLTensorExtents.Rank]: Obtains the rank of the tensor.
//
// See: https://developer.apple.com/documentation/Metal/MTLTensorExtents
type MTLTensorExtents struct {
	objectivec.Object
}

// MTLTensorExtentsFromID constructs a [MTLTensorExtents] from an objc.ID.
//
// An array of length matching the rank, holding the dimensions of a tensor.
func MTLTensorExtentsFromID(id objc.ID) MTLTensorExtents {
	return MTLTensorExtents{objectivec.Object{ID: id}}
}
// NOTE: MTLTensorExtents adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLTensorExtents] class.
//
// # Instance Properties
//
//   - [IMTLTensorExtents.Extents]: Retrieves the extents for this object.
//   - [IMTLTensorExtents.SetExtents]
//   - [IMTLTensorExtents.Rank]: Obtains the rank of the tensor.
//
// See: https://developer.apple.com/documentation/Metal/MTLTensorExtents
type IMTLTensorExtents interface {
	objectivec.IObject

	// Topic: Instance Properties

	// Retrieves the extents for this object.
	Extents() int
	SetExtents(value int)
	// Obtains the rank of the tensor.
	Rank() uint

	// An error domain for errors that pertain to creating a tensor.
	MTLTensorDomain() string
	MTL_TENSOR_MAX_RANK() objectivec.IObject
	SetMTL_TENSOR_MAX_RANK(value objectivec.IObject)
	// Returns the extent at an index.
	ExtentAtDimensionIndex(dimensionIndex uint) int
	// Creates a new tensor extents with the rank and extent values you provide.
	InitWithRankValues(rank uint, values unsafe.Pointer) MTLTensorExtents
}

// Init initializes the instance.
func (t MTLTensorExtents) Init() MTLTensorExtents {
	rv := objc.Send[MTLTensorExtents](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t MTLTensorExtents) Autorelease() MTLTensorExtents {
	rv := objc.Send[MTLTensorExtents](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLTensorExtents creates a new MTLTensorExtents instance.
func NewMTLTensorExtents() MTLTensorExtents {
	class := getMTLTensorExtentsClass()
	rv := objc.Send[MTLTensorExtents](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new tensor extents with the rank and extent values you provide.
//
// rank: The number of dimensions.
//
// values: An array of length `rank` that specifies the size of each dimension. The
// first dimension is the innermost dimension.
//
// # Return Value
// 
// Tensor extents with the rank and extent values you provide. Returns `nil`
// if `rank` exceeds 0 and `values` is nil or if `rank` exceeds
// [MTL_TENSOR_MAX_RANK].
//
// [MTL_TENSOR_MAX_RANK]: https://developer.apple.com/documentation/Metal/MTL_TENSOR_MAX_RANK
//
// # Discussion
// 
// Zero rank extents represent scalars. `values` can only be `nil`if `rank` is
// 0.
//
// See: https://developer.apple.com/documentation/Metal/MTLTensorExtents/initWithRank:values:
func NewTensorExtentsWithRankValues(rank uint, values unsafe.Pointer) MTLTensorExtents {
	instance := getMTLTensorExtentsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRank:values:"), rank, values)
	return MTLTensorExtentsFromID(rv)
}

// Returns the extent at an index.
//
// dimensionIndex: The index of the dimension. The first dimension is the innermost dimension.
//
// # Return Value
// 
// The extent at `dimensionIndex`. This method returns -1 if `dimensionIndex`
// is greater than or equal to `rank`.
//
// See: https://developer.apple.com/documentation/Metal/MTLTensorExtents/extentAtDimensionIndex:
func (t MTLTensorExtents) ExtentAtDimensionIndex(dimensionIndex uint) int {
	rv := objc.Send[int](t.ID, objc.Sel("extentAtDimensionIndex:"), dimensionIndex)
	return rv
}
// Creates a new tensor extents with the rank and extent values you provide.
//
// rank: The number of dimensions.
//
// values: An array of length `rank` that specifies the size of each dimension. The
// first dimension is the innermost dimension.
//
// # Return Value
// 
// Tensor extents with the rank and extent values you provide. Returns `nil`
// if `rank` exceeds 0 and `values` is nil or if `rank` exceeds
// [MTL_TENSOR_MAX_RANK].
//
// [MTL_TENSOR_MAX_RANK]: https://developer.apple.com/documentation/Metal/MTL_TENSOR_MAX_RANK
//
// # Discussion
// 
// Zero rank extents represent scalars. `values` can only be `nil`if `rank` is
// 0.
//
// See: https://developer.apple.com/documentation/Metal/MTLTensorExtents/initWithRank:values:
func (t MTLTensorExtents) InitWithRankValues(rank uint, values unsafe.Pointer) MTLTensorExtents {
	rv := objc.Send[MTLTensorExtents](t.ID, objc.Sel("initWithRank:values:"), rank, values)
	return rv
}

// Retrieves the extents for this object.
//
// See: https://developer.apple.com/documentation/metal/mtltensorextents/extents
func (t MTLTensorExtents) Extents() int {
	rv := objc.Send[int](t.ID, objc.Sel("extents"))
	return rv
}
func (t MTLTensorExtents) SetExtents(value int) {
	objc.Send[struct{}](t.ID, objc.Sel("setExtents:"), value)
}
// Obtains the rank of the tensor.
//
// # Discussion
// 
// The rank represents the number of dimensions.
//
// See: https://developer.apple.com/documentation/Metal/MTLTensorExtents/rank
func (t MTLTensorExtents) Rank() uint {
	rv := objc.Send[uint](t.ID, objc.Sel("rank"))
	return rv
}
// An error domain for errors that pertain to creating a tensor.
//
// See: https://developer.apple.com/documentation/metal/mtltensordomain
func (t MTLTensorExtents) MTLTensorDomain() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("MTLTensorDomain"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/metal/mtl_tensor_max_rank
func (t MTLTensorExtents) MTL_TENSOR_MAX_RANK() objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("MTL_TENSOR_MAX_RANK"))
	return objectivec.Object{ID: rv}
}
func (t MTLTensorExtents) SetMTL_TENSOR_MAX_RANK(value objectivec.IObject) {
	objc.Send[struct{}](t.ID, objc.Sel("setMTL_TENSOR_MAX_RANK:"), value)
}

