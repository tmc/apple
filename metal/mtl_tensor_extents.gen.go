// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"

	"github.com/tmc/apple/objc"
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
// # Instance Properties
//
//   - [MTLTensorExtents.Rank]: Obtains the rank of the tensor.
//
// See: https://developer.apple.com/documentation/Metal/MTLTensorExtents
//
// [MTL_TENSOR_MAX_RANK]: https://developer.apple.com/documentation/Metal/MTL_TENSOR_MAX_RANK
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
//   - [IMTLTensorExtents.Rank]: Obtains the rank of the tensor.
//
// See: https://developer.apple.com/documentation/Metal/MTLTensorExtents
type IMTLTensorExtents interface {
	objectivec.IObject

	// Topic: Instance Properties

	// Obtains the rank of the tensor.
	Rank() uint

	// Returns the extent at an index.
	ExtentAtDimensionIndex(dimensionIndex uint) int
	// Creates a new tensor extents with the rank and extent values you provide.
	InitWithRankValues(rank uint, values *int) MTLTensorExtents
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
// # Discussion
//
// Zero rank extents represent scalars. `values` can only be `nil`if `rank` is
// 0.
//
// See: https://developer.apple.com/documentation/Metal/MTLTensorExtents/initWithRank:values:
//
// [MTL_TENSOR_MAX_RANK]: https://developer.apple.com/documentation/Metal/MTL_TENSOR_MAX_RANK
func NewTensorExtentsWithRankValues(rank uint, values *int) MTLTensorExtents {
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
// # Discussion
//
// Zero rank extents represent scalars. `values` can only be `nil`if `rank` is
// 0.
//
// See: https://developer.apple.com/documentation/Metal/MTLTensorExtents/initWithRank:values:
//
// [MTL_TENSOR_MAX_RANK]: https://developer.apple.com/documentation/Metal/MTL_TENSOR_MAX_RANK
func (t MTLTensorExtents) InitWithRankValues(rank uint, values *int) MTLTensorExtents {
	rv := objc.Send[MTLTensorExtents](t.ID, objc.Sel("initWithRank:values:"), rank, values)
	return rv
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
