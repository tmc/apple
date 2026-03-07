// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLHeapDescriptor] class.
var (
	_MTLHeapDescriptorClass     MTLHeapDescriptorClass
	_MTLHeapDescriptorClassOnce sync.Once
)

func getMTLHeapDescriptorClass() MTLHeapDescriptorClass {
	_MTLHeapDescriptorClassOnce.Do(func() {
		_MTLHeapDescriptorClass = MTLHeapDescriptorClass{class: objc.GetClass("MTLHeapDescriptor")}
	})
	return _MTLHeapDescriptorClass
}

// GetMTLHeapDescriptorClass returns the class object for MTLHeapDescriptor.
func GetMTLHeapDescriptorClass() MTLHeapDescriptorClass {
	return getMTLHeapDescriptorClass()
}

type MTLHeapDescriptorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLHeapDescriptorClass) Alloc() MTLHeapDescriptor {
	rv := objc.Send[MTLHeapDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// A configuration that customizes the behavior for a Metal memory heap.
//
// # Overview
// 
// Create an [MTLHeap] by configuring an [MTLHeapDescriptor] instance’s
// properties and passing it to the [NewHeapWithDescriptor] method of an
// [MTLDevice].
// 
// Each new heap inherits the descriptor’s configuration as you create it,
// which means you can modify and reuse a descriptor to create other heaps.
//
// # Configuring a heap
//
//   - [MTLHeapDescriptor.Type]: The memory placement strategy for any resources you allocate from the heaps you create with this descriptor.
//   - [MTLHeapDescriptor.SetType]
//   - [MTLHeapDescriptor.StorageMode]: The storage mode for the heaps you create with this descriptor.
//   - [MTLHeapDescriptor.SetStorageMode]
//   - [MTLHeapDescriptor.CpuCacheMode]: The CPU cache behavior for any resources you allocate from the heaps you create with this descriptor.
//   - [MTLHeapDescriptor.SetCpuCacheMode]
//   - [MTLHeapDescriptor.HazardTrackingMode]: The hazard tracking behavior for any resources you allocate from the heaps you create with this descriptor.
//   - [MTLHeapDescriptor.SetHazardTrackingMode]
//   - [MTLHeapDescriptor.ResourceOptions]: The combined behavior for any resources you allocate from the heaps you create with this descriptor.
//   - [MTLHeapDescriptor.SetResourceOptions]
//   - [MTLHeapDescriptor.Size]: The total amount of memory, in bytes, for the heaps you create with this descriptor.
//   - [MTLHeapDescriptor.SetSize]
//   - [MTLHeapDescriptor.SparsePageSize]: The page size for any resources you allocate from the heaps you create with this descriptor.
//   - [MTLHeapDescriptor.SetSparsePageSize]
//
// # Instance Properties
//
//   - [MTLHeapDescriptor.MaxCompatiblePlacementSparsePageSize]: Specifies the largest sparse page size that the Metal heap supports.
//   - [MTLHeapDescriptor.SetMaxCompatiblePlacementSparsePageSize]
//
// See: https://developer.apple.com/documentation/Metal/MTLHeapDescriptor
type MTLHeapDescriptor struct {
	objectivec.Object
}

// MTLHeapDescriptorFromID constructs a [MTLHeapDescriptor] from an objc.ID.
//
// A configuration that customizes the behavior for a Metal memory heap.
func MTLHeapDescriptorFromID(id objc.ID) MTLHeapDescriptor {
	return MTLHeapDescriptor{objectivec.Object{id}}
}
// NOTE: MTLHeapDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [MTLHeapDescriptor] class.
//
// # Configuring a heap
//
//   - [IMTLHeapDescriptor.Type]: The memory placement strategy for any resources you allocate from the heaps you create with this descriptor.
//   - [IMTLHeapDescriptor.SetType]
//   - [IMTLHeapDescriptor.StorageMode]: The storage mode for the heaps you create with this descriptor.
//   - [IMTLHeapDescriptor.SetStorageMode]
//   - [IMTLHeapDescriptor.CpuCacheMode]: The CPU cache behavior for any resources you allocate from the heaps you create with this descriptor.
//   - [IMTLHeapDescriptor.SetCpuCacheMode]
//   - [IMTLHeapDescriptor.HazardTrackingMode]: The hazard tracking behavior for any resources you allocate from the heaps you create with this descriptor.
//   - [IMTLHeapDescriptor.SetHazardTrackingMode]
//   - [IMTLHeapDescriptor.ResourceOptions]: The combined behavior for any resources you allocate from the heaps you create with this descriptor.
//   - [IMTLHeapDescriptor.SetResourceOptions]
//   - [IMTLHeapDescriptor.Size]: The total amount of memory, in bytes, for the heaps you create with this descriptor.
//   - [IMTLHeapDescriptor.SetSize]
//   - [IMTLHeapDescriptor.SparsePageSize]: The page size for any resources you allocate from the heaps you create with this descriptor.
//   - [IMTLHeapDescriptor.SetSparsePageSize]
//
// # Instance Properties
//
//   - [IMTLHeapDescriptor.MaxCompatiblePlacementSparsePageSize]: Specifies the largest sparse page size that the Metal heap supports.
//   - [IMTLHeapDescriptor.SetMaxCompatiblePlacementSparsePageSize]
//
// See: https://developer.apple.com/documentation/Metal/MTLHeapDescriptor
type IMTLHeapDescriptor interface {
	objectivec.IObject

	// Topic: Configuring a heap

	// The memory placement strategy for any resources you allocate from the heaps you create with this descriptor.
	Type() MTLHeapType
	SetType(value MTLHeapType)
	// The storage mode for the heaps you create with this descriptor.
	StorageMode() MTLStorageMode
	SetStorageMode(value MTLStorageMode)
	// The CPU cache behavior for any resources you allocate from the heaps you create with this descriptor.
	CpuCacheMode() MTLCPUCacheMode
	SetCpuCacheMode(value MTLCPUCacheMode)
	// The hazard tracking behavior for any resources you allocate from the heaps you create with this descriptor.
	HazardTrackingMode() MTLHazardTrackingMode
	SetHazardTrackingMode(value MTLHazardTrackingMode)
	// The combined behavior for any resources you allocate from the heaps you create with this descriptor.
	ResourceOptions() MTLResourceOptions
	SetResourceOptions(value MTLResourceOptions)
	// The total amount of memory, in bytes, for the heaps you create with this descriptor.
	Size() uint
	SetSize(value uint)
	// The page size for any resources you allocate from the heaps you create with this descriptor.
	SparsePageSize() MTLSparsePageSize
	SetSparsePageSize(value MTLSparsePageSize)

	// Topic: Instance Properties

	// Specifies the largest sparse page size that the Metal heap supports.
	MaxCompatiblePlacementSparsePageSize() MTLSparsePageSize
	SetMaxCompatiblePlacementSparsePageSize(value MTLSparsePageSize)
}





// Init initializes the instance.
func (h MTLHeapDescriptor) Init() MTLHeapDescriptor {
	rv := objc.Send[MTLHeapDescriptor](h.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h MTLHeapDescriptor) Autorelease() MTLHeapDescriptor {
	rv := objc.Send[MTLHeapDescriptor](h.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLHeapDescriptor creates a new MTLHeapDescriptor instance.
func NewMTLHeapDescriptor() MTLHeapDescriptor {
	class := getMTLHeapDescriptorClass()
	rv := objc.Send[MTLHeapDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}





















// The memory placement strategy for any resources you allocate from the heaps
// you create with this descriptor.
//
// # Discussion
// 
// This property’s default value is [HeapTypeAutomatic].
//
// See: https://developer.apple.com/documentation/Metal/MTLHeapDescriptor/type
func (h MTLHeapDescriptor) Type() MTLHeapType {
	rv := objc.Send[MTLHeapType](h.ID, objc.Sel("type"))
	return MTLHeapType(rv)
}
func (h MTLHeapDescriptor) SetType(value MTLHeapType) {
	objc.Send[struct{}](h.ID, objc.Sel("setType:"), value)
}



// The storage mode for the heaps you create with this descriptor.
//
// # Discussion
// 
// For devices with Apple silicon, you can create a heap with either the
// [StorageModePrivate] or the [StorageModeShared] storage mode. However, you
// can only create heaps with private storage on macOS devices without Apple
// silicon.
// 
// The resources you allocate from a heap inherit that heap’s storage mode.
// This property’s default value is [StorageModePrivate].
//
// See: https://developer.apple.com/documentation/Metal/MTLHeapDescriptor/storageMode
func (h MTLHeapDescriptor) StorageMode() MTLStorageMode {
	rv := objc.Send[MTLStorageMode](h.ID, objc.Sel("storageMode"))
	return MTLStorageMode(rv)
}
func (h MTLHeapDescriptor) SetStorageMode(value MTLStorageMode) {
	objc.Send[struct{}](h.ID, objc.Sel("setStorageMode:"), value)
}



// The CPU cache behavior for any resources you allocate from the heaps you
// create with this descriptor.
//
// # Discussion
// 
// This property’s default value is [CPUCacheModeDefaultCache].
// 
// The resources you allocate from a heap inherit that heap’s CPU cache
// mode.
//
// See: https://developer.apple.com/documentation/Metal/MTLHeapDescriptor/cpuCacheMode
func (h MTLHeapDescriptor) CpuCacheMode() MTLCPUCacheMode {
	rv := objc.Send[MTLCPUCacheMode](h.ID, objc.Sel("cpuCacheMode"))
	return MTLCPUCacheMode(rv)
}
func (h MTLHeapDescriptor) SetCpuCacheMode(value MTLCPUCacheMode) {
	objc.Send[struct{}](h.ID, objc.Sel("setCpuCacheMode:"), value)
}



// The hazard tracking behavior for any resources you allocate from the heaps
// you create with this descriptor.
//
// # Discussion
// 
// This property’s default value is [HazardTrackingModeDefault], which is
// equivalent to [HazardTrackingModeUntracked] for a heap.
// 
// The resources you allocate from a heap inherit that heap’s hazard
// tracking mode.
//
// See: https://developer.apple.com/documentation/Metal/MTLHeapDescriptor/hazardTrackingMode
func (h MTLHeapDescriptor) HazardTrackingMode() MTLHazardTrackingMode {
	rv := objc.Send[MTLHazardTrackingMode](h.ID, objc.Sel("hazardTrackingMode"))
	return MTLHazardTrackingMode(rv)
}
func (h MTLHeapDescriptor) SetHazardTrackingMode(value MTLHazardTrackingMode) {
	objc.Send[struct{}](h.ID, objc.Sel("setHazardTrackingMode:"), value)
}



// The combined behavior for any resources you allocate from the heaps you
// create with this descriptor.
//
// # Discussion
// 
// This property aggregates the values of [StorageMode], [CpuCacheMode], and
// [HazardTrackingMode]. Any modifications you make to this property affect
// the other properties, and vice versa.
//
// See: https://developer.apple.com/documentation/Metal/MTLHeapDescriptor/resourceOptions
func (h MTLHeapDescriptor) ResourceOptions() MTLResourceOptions {
	rv := objc.Send[MTLResourceOptions](h.ID, objc.Sel("resourceOptions"))
	return MTLResourceOptions(rv)
}
func (h MTLHeapDescriptor) SetResourceOptions(value MTLResourceOptions) {
	objc.Send[struct{}](h.ID, objc.Sel("setResourceOptions:"), value)
}



// The total amount of memory, in bytes, for the heaps you create with this
// descriptor.
//
// # Discussion
// 
// You can use various [MTLDevice] methods to help you estimate an appropriate
// heap size, including the following:
// 
// - [HeapBufferSizeAndAlignWithLengthOptions] -
// [HeapTextureSizeAndAlignWithDescriptor] -
// [HeapAccelerationStructureSizeAndAlignWithSize] -
// [HeapAccelerationStructureSizeAndAlignWithDescriptor]
// 
// This property’s default value is `0`.
//
// See: https://developer.apple.com/documentation/Metal/MTLHeapDescriptor/size
func (h MTLHeapDescriptor) Size() uint {
	rv := objc.Send[uint](h.ID, objc.Sel("size"))
	return rv
}
func (h MTLHeapDescriptor) SetSize(value uint) {
	objc.Send[struct{}](h.ID, objc.Sel("setSize:"), value)
}



// The page size for any resources you allocate from the heaps you create with
// this descriptor.
//
// # Discussion
// 
// This property’s default value is 16 kilobytes ([SparsePageSize16]), which
// is a smaller page size option that can help reduce your app’s memory
// usage. However, you can reduce operational overhead for sparse textures
// with larger page sizes, such as [SparsePageSize64] and [SparsePageSize256].
// These operations include blit commands and the configuration of sparse
// texture mappings (see [Blit passes] and [MTLResourceStateCommandEncoder],
// respectively).
//
// [Blit passes]: https://developer.apple.com/documentation/Metal/blit-passes
//
// See: https://developer.apple.com/documentation/Metal/MTLHeapDescriptor/sparsePageSize
func (h MTLHeapDescriptor) SparsePageSize() MTLSparsePageSize {
	rv := objc.Send[MTLSparsePageSize](h.ID, objc.Sel("sparsePageSize"))
	return MTLSparsePageSize(rv)
}
func (h MTLHeapDescriptor) SetSparsePageSize(value MTLSparsePageSize) {
	objc.Send[struct{}](h.ID, objc.Sel("setSparsePageSize:"), value)
}



// Specifies the largest sparse page size that the Metal heap supports.
//
// # Discussion
// 
// This parameter only affects the heap if you set the [Type] property of this
// descriptor to [HeapTypePlacement].
// 
// The value you assign to this property determines the compatibility of the
// Metal heap with with placement sparse resources, because placement sparse
// resources require that their sparse page size be less than or equal to the
// placement sparse page of the Metal heap that this property controls.
//
// See: https://developer.apple.com/documentation/Metal/MTLHeapDescriptor/maxCompatiblePlacementSparsePageSize
func (h MTLHeapDescriptor) MaxCompatiblePlacementSparsePageSize() MTLSparsePageSize {
	rv := objc.Send[MTLSparsePageSize](h.ID, objc.Sel("maxCompatiblePlacementSparsePageSize"))
	return MTLSparsePageSize(rv)
}
func (h MTLHeapDescriptor) SetMaxCompatiblePlacementSparsePageSize(value MTLSparsePageSize) {
	objc.Send[struct{}](h.ID, objc.Sel("setMaxCompatiblePlacementSparsePageSize:"), value)
}
























