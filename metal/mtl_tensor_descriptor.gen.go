// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLTensorDescriptor] class.
var (
	_MTLTensorDescriptorClass     MTLTensorDescriptorClass
	_MTLTensorDescriptorClassOnce sync.Once
)

func getMTLTensorDescriptorClass() MTLTensorDescriptorClass {
	_MTLTensorDescriptorClassOnce.Do(func() {
		_MTLTensorDescriptorClass = MTLTensorDescriptorClass{class: objc.GetClass("MTLTensorDescriptor")}
	})
	return _MTLTensorDescriptorClass
}

// GetMTLTensorDescriptorClass returns the class object for MTLTensorDescriptor.
func GetMTLTensorDescriptorClass() MTLTensorDescriptorClass {
	return getMTLTensorDescriptorClass()
}

type MTLTensorDescriptorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLTensorDescriptorClass) Alloc() MTLTensorDescriptor {
	rv := objc.Send[MTLTensorDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A configuration type for creating new tensor instances.
//
// # Instance Properties
//
//   - [MTLTensorDescriptor.CpuCacheMode]: A value that configures the cache mode of CPU mapping of tensors you create with this descriptor.
//   - [MTLTensorDescriptor.SetCpuCacheMode]
//   - [MTLTensorDescriptor.DataType]: A data format for the tensors you create with this descriptor.
//   - [MTLTensorDescriptor.SetDataType]
//   - [MTLTensorDescriptor.Dimensions]: An array of sizes, in elements, one for each dimension of the tensors you create with this descriptor.
//   - [MTLTensorDescriptor.SetDimensions]
//   - [MTLTensorDescriptor.HazardTrackingMode]: A value that configures the hazard tracking of tensors you create with this descriptor.
//   - [MTLTensorDescriptor.SetHazardTrackingMode]
//   - [MTLTensorDescriptor.ResourceOptions]: A packed set of the `storageMode`, `cpuCacheMode` and `hazardTrackingMode` properties.
//   - [MTLTensorDescriptor.SetResourceOptions]
//   - [MTLTensorDescriptor.StorageMode]: A value that configures the memory location and access permissions of tensors you create with this descriptor.
//   - [MTLTensorDescriptor.SetStorageMode]
//   - [MTLTensorDescriptor.Strides]: An array of strides, in elements, one for each dimension in the tensors you create with this descriptor, if applicable.
//   - [MTLTensorDescriptor.SetStrides]
//   - [MTLTensorDescriptor.Usage]: A set of contexts in which you can use tensors you create with this descriptor.
//   - [MTLTensorDescriptor.SetUsage]
//
// See: https://developer.apple.com/documentation/Metal/MTLTensorDescriptor
type MTLTensorDescriptor struct {
	objectivec.Object
}

// MTLTensorDescriptorFromID constructs a [MTLTensorDescriptor] from an objc.ID.
//
// A configuration type for creating new tensor instances.
func MTLTensorDescriptorFromID(id objc.ID) MTLTensorDescriptor {
	return MTLTensorDescriptor{objectivec.Object{ID: id}}
}
// NOTE: MTLTensorDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLTensorDescriptor] class.
//
// # Instance Properties
//
//   - [IMTLTensorDescriptor.CpuCacheMode]: A value that configures the cache mode of CPU mapping of tensors you create with this descriptor.
//   - [IMTLTensorDescriptor.SetCpuCacheMode]
//   - [IMTLTensorDescriptor.DataType]: A data format for the tensors you create with this descriptor.
//   - [IMTLTensorDescriptor.SetDataType]
//   - [IMTLTensorDescriptor.Dimensions]: An array of sizes, in elements, one for each dimension of the tensors you create with this descriptor.
//   - [IMTLTensorDescriptor.SetDimensions]
//   - [IMTLTensorDescriptor.HazardTrackingMode]: A value that configures the hazard tracking of tensors you create with this descriptor.
//   - [IMTLTensorDescriptor.SetHazardTrackingMode]
//   - [IMTLTensorDescriptor.ResourceOptions]: A packed set of the `storageMode`, `cpuCacheMode` and `hazardTrackingMode` properties.
//   - [IMTLTensorDescriptor.SetResourceOptions]
//   - [IMTLTensorDescriptor.StorageMode]: A value that configures the memory location and access permissions of tensors you create with this descriptor.
//   - [IMTLTensorDescriptor.SetStorageMode]
//   - [IMTLTensorDescriptor.Strides]: An array of strides, in elements, one for each dimension in the tensors you create with this descriptor, if applicable.
//   - [IMTLTensorDescriptor.SetStrides]
//   - [IMTLTensorDescriptor.Usage]: A set of contexts in which you can use tensors you create with this descriptor.
//   - [IMTLTensorDescriptor.SetUsage]
//
// See: https://developer.apple.com/documentation/Metal/MTLTensorDescriptor
type IMTLTensorDescriptor interface {
	objectivec.IObject

	// Topic: Instance Properties

	// A value that configures the cache mode of CPU mapping of tensors you create with this descriptor.
	CpuCacheMode() MTLCPUCacheMode
	SetCpuCacheMode(value MTLCPUCacheMode)
	// A data format for the tensors you create with this descriptor.
	DataType() MTLTensorDataType
	SetDataType(value MTLTensorDataType)
	// An array of sizes, in elements, one for each dimension of the tensors you create with this descriptor.
	Dimensions() IMTLTensorExtents
	SetDimensions(value IMTLTensorExtents)
	// A value that configures the hazard tracking of tensors you create with this descriptor.
	HazardTrackingMode() MTLHazardTrackingMode
	SetHazardTrackingMode(value MTLHazardTrackingMode)
	// A packed set of the `storageMode`, `cpuCacheMode` and `hazardTrackingMode` properties.
	ResourceOptions() MTLResourceOptions
	SetResourceOptions(value MTLResourceOptions)
	// A value that configures the memory location and access permissions of tensors you create with this descriptor.
	StorageMode() MTLStorageMode
	SetStorageMode(value MTLStorageMode)
	// An array of strides, in elements, one for each dimension in the tensors you create with this descriptor, if applicable.
	Strides() IMTLTensorExtents
	SetStrides(value IMTLTensorExtents)
	// A set of contexts in which you can use tensors you create with this descriptor.
	Usage() MTLTensorUsage
	SetUsage(value MTLTensorUsage)

	// An error domain for errors that pertain to creating a tensor.
	MTLTensorDomain() string
	MTL_TENSOR_MAX_RANK() objectivec.IObject
	SetMTL_TENSOR_MAX_RANK(value objectivec.IObject)
}

// Init initializes the instance.
func (t MTLTensorDescriptor) Init() MTLTensorDescriptor {
	rv := objc.Send[MTLTensorDescriptor](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t MTLTensorDescriptor) Autorelease() MTLTensorDescriptor {
	rv := objc.Send[MTLTensorDescriptor](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLTensorDescriptor creates a new MTLTensorDescriptor instance.
func NewMTLTensorDescriptor() MTLTensorDescriptor {
	class := getMTLTensorDescriptorClass()
	rv := objc.Send[MTLTensorDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A value that configures the cache mode of CPU mapping of tensors you create
// with this descriptor.
//
// # Discussion
// 
// The default value of this property is [CPUCacheModeDefaultCache].
//
// See: https://developer.apple.com/documentation/Metal/MTLTensorDescriptor/cpuCacheMode
func (t MTLTensorDescriptor) CpuCacheMode() MTLCPUCacheMode {
	rv := objc.Send[MTLCPUCacheMode](t.ID, objc.Sel("cpuCacheMode"))
	return MTLCPUCacheMode(rv)
}
func (t MTLTensorDescriptor) SetCpuCacheMode(value MTLCPUCacheMode) {
	objc.Send[struct{}](t.ID, objc.Sel("setCpuCacheMode:"), value)
}

// A data format for the tensors you create with this descriptor.
//
// # Discussion
// 
// The default value of this property is [TensorDataTypeFloat32].
//
// See: https://developer.apple.com/documentation/Metal/MTLTensorDescriptor/dataType
func (t MTLTensorDescriptor) DataType() MTLTensorDataType {
	rv := objc.Send[MTLTensorDataType](t.ID, objc.Sel("dataType"))
	return MTLTensorDataType(rv)
}
func (t MTLTensorDescriptor) SetDataType(value MTLTensorDataType) {
	objc.Send[struct{}](t.ID, objc.Sel("setDataType:"), value)
}

// An array of sizes, in elements, one for each dimension of the tensors you
// create with this descriptor.
//
// # Discussion
// 
// The default value of this property is a rank one extents with size one.
//
// See: https://developer.apple.com/documentation/Metal/MTLTensorDescriptor/dimensions
func (t MTLTensorDescriptor) Dimensions() IMTLTensorExtents {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("dimensions"))
	return MTLTensorExtentsFromID(objc.ID(rv))
}
func (t MTLTensorDescriptor) SetDimensions(value IMTLTensorExtents) {
	objc.Send[struct{}](t.ID, objc.Sel("setDimensions:"), value)
}

// A value that configures the hazard tracking of tensors you create with this
// descriptor.
//
// # Discussion
// 
// The default value of this property is [HazardTrackingModeDefault].
//
// See: https://developer.apple.com/documentation/Metal/MTLTensorDescriptor/hazardTrackingMode
func (t MTLTensorDescriptor) HazardTrackingMode() MTLHazardTrackingMode {
	rv := objc.Send[MTLHazardTrackingMode](t.ID, objc.Sel("hazardTrackingMode"))
	return MTLHazardTrackingMode(rv)
}
func (t MTLTensorDescriptor) SetHazardTrackingMode(value MTLHazardTrackingMode) {
	objc.Send[struct{}](t.ID, objc.Sel("setHazardTrackingMode:"), value)
}

// A packed set of the `storageMode`, `cpuCacheMode` and `hazardTrackingMode`
// properties.
//
// See: https://developer.apple.com/documentation/Metal/MTLTensorDescriptor/resourceOptions
func (t MTLTensorDescriptor) ResourceOptions() MTLResourceOptions {
	rv := objc.Send[MTLResourceOptions](t.ID, objc.Sel("resourceOptions"))
	return MTLResourceOptions(rv)
}
func (t MTLTensorDescriptor) SetResourceOptions(value MTLResourceOptions) {
	objc.Send[struct{}](t.ID, objc.Sel("setResourceOptions:"), value)
}

// A value that configures the memory location and access permissions of
// tensors you create with this descriptor.
//
// # Discussion
// 
// The default value of this property defaults to [StorageModeShared].
//
// See: https://developer.apple.com/documentation/Metal/MTLTensorDescriptor/storageMode
func (t MTLTensorDescriptor) StorageMode() MTLStorageMode {
	rv := objc.Send[MTLStorageMode](t.ID, objc.Sel("storageMode"))
	return MTLStorageMode(rv)
}
func (t MTLTensorDescriptor) SetStorageMode(value MTLStorageMode) {
	objc.Send[struct{}](t.ID, objc.Sel("setStorageMode:"), value)
}

// An array of strides, in elements, one for each dimension in the tensors you
// create with this descriptor, if applicable.
//
// # Discussion
// 
// This property only applies to tensors you create from a buffer, otherwise
// it is nil. You are responsible for ensuring `strides` meets the following
// requirements:
// 
// - Elements of `strides`are in monotonically non-decreasing order. - The
// first element of `strides` is one. - For any `i` larger than zero,
// `strides[i]` is greater than or equal to `strides[i-1] * dimensions[i-1]`.
// - If `usage` contains [TensorUsageMachineLearning], the second element of
// `strides` is aligned to 64 bytes, and for any `i` larger than one,
// `strides[i]` is equal to `strides[i-1] * dimensions[i-1]`.
//
// See: https://developer.apple.com/documentation/Metal/MTLTensorDescriptor/strides
func (t MTLTensorDescriptor) Strides() IMTLTensorExtents {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("strides"))
	return MTLTensorExtentsFromID(objc.ID(rv))
}
func (t MTLTensorDescriptor) SetStrides(value IMTLTensorExtents) {
	objc.Send[struct{}](t.ID, objc.Sel("setStrides:"), value)
}

// A set of contexts in which you can use tensors you create with this
// descriptor.
//
// # Discussion
// 
// The default value for this property is a bitwise [OR] of:
// 
// - [TensorUsageRender]
// - [TensorUsageCompute]
//
// See: https://developer.apple.com/documentation/Metal/MTLTensorDescriptor/usage
func (t MTLTensorDescriptor) Usage() MTLTensorUsage {
	rv := objc.Send[MTLTensorUsage](t.ID, objc.Sel("usage"))
	return MTLTensorUsage(rv)
}
func (t MTLTensorDescriptor) SetUsage(value MTLTensorUsage) {
	objc.Send[struct{}](t.ID, objc.Sel("setUsage:"), value)
}

// An error domain for errors that pertain to creating a tensor.
//
// See: https://developer.apple.com/documentation/metal/mtltensordomain
func (t MTLTensorDescriptor) MTLTensorDomain() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("MTLTensorDomain"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/metal/mtl_tensor_max_rank
func (t MTLTensorDescriptor) MTL_TENSOR_MAX_RANK() objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("MTL_TENSOR_MAX_RANK"))
	return objectivec.Object{ID: rv}
}
func (t MTLTensorDescriptor) SetMTL_TENSOR_MAX_RANK(value objectivec.IObject) {
	objc.Send[struct{}](t.ID, objc.Sel("setMTL_TENSOR_MAX_RANK:"), value)
}

