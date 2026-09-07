// Code generated from Apple documentation for virtualization. DO NOT EDIT.

package virtualization

import (
	"context"
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZVirtioSharedMemoryRegion] class.
var (
	_VZVirtioSharedMemoryRegionClass     VZVirtioSharedMemoryRegionClass
	_VZVirtioSharedMemoryRegionClassOnce sync.Once
)

func getVZVirtioSharedMemoryRegionClass() VZVirtioSharedMemoryRegionClass {
	_VZVirtioSharedMemoryRegionClassOnce.Do(func() {
		_VZVirtioSharedMemoryRegionClass = VZVirtioSharedMemoryRegionClass{class: objc.GetClass("VZVirtioSharedMemoryRegion")}
	})
	return _VZVirtioSharedMemoryRegionClass
}

// GetVZVirtioSharedMemoryRegionClass returns the class object for VZVirtioSharedMemoryRegion.
func GetVZVirtioSharedMemoryRegionClass() VZVirtioSharedMemoryRegionClass {
	return getVZVirtioSharedMemoryRegionClass()
}

type VZVirtioSharedMemoryRegionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZVirtioSharedMemoryRegionClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtioSharedMemoryRegionClass) Alloc() VZVirtioSharedMemoryRegion {
	rv := objc.SendIfResponds[VZVirtioSharedMemoryRegion](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZVirtioSharedMemoryRegion.MapMemoryAtOffsetSizeCompletionHandler]
//   - [VZVirtioSharedMemoryRegion.RegionID]
//   - [VZVirtioSharedMemoryRegion.Size]
//   - [VZVirtioSharedMemoryRegion.UnmapMemoryAtOffsetSizeCompletionHandler]
type VZVirtioSharedMemoryRegion struct {
	objectivec.Object
}

// VZVirtioSharedMemoryRegionFromID constructs a [VZVirtioSharedMemoryRegion] from an objc.ID.
func VZVirtioSharedMemoryRegionFromID(id objc.ID) VZVirtioSharedMemoryRegion {
	return VZVirtioSharedMemoryRegion{objectivec.Object{ID: id}}
}

// Ensure VZVirtioSharedMemoryRegion implements IVZVirtioSharedMemoryRegion.
var _ IVZVirtioSharedMemoryRegion = VZVirtioSharedMemoryRegion{}

// An interface definition for the [VZVirtioSharedMemoryRegion] class.
//
// # Methods
//
//   - [IVZVirtioSharedMemoryRegion.MapMemoryAtOffsetSizeCompletionHandler]
//   - [IVZVirtioSharedMemoryRegion.RegionID]
//   - [IVZVirtioSharedMemoryRegion.Size]
//   - [IVZVirtioSharedMemoryRegion.UnmapMemoryAtOffsetSizeCompletionHandler]
type IVZVirtioSharedMemoryRegion interface {
	objectivec.IObject

	// Topic: Methods

	MapMemoryAtOffsetSizeCompletionHandler(memory unsafe.Pointer, offset uint64, size uint64, handler ErrorHandler)
	RegionID() byte
	Size() uint64
	UnmapMemoryAtOffsetSizeCompletionHandler(offset uint64, size uint64, handler ErrorHandler)
}

// Init initializes the instance.
func (v VZVirtioSharedMemoryRegion) Init() VZVirtioSharedMemoryRegion {
	rv := objc.SendIfResponds[VZVirtioSharedMemoryRegion](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtioSharedMemoryRegion) Autorelease() VZVirtioSharedMemoryRegion {
	rv := objc.SendIfResponds[VZVirtioSharedMemoryRegion](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioSharedMemoryRegion creates a new VZVirtioSharedMemoryRegion instance.
func NewVZVirtioSharedMemoryRegion() VZVirtioSharedMemoryRegion {
	class := getVZVirtioSharedMemoryRegionClass()
	rv := objc.SendIfResponds[VZVirtioSharedMemoryRegion](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VZVirtioSharedMemoryRegion) MapMemoryAtOffsetSizeCompletionHandler(memory unsafe.Pointer, offset uint64, size uint64, handler ErrorHandler) {
	_block3, _ := NewErrorBlock(handler)
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("mapMemory:atOffset:size:completionHandler:"), memory, offset, size, _block3)
}
func (v VZVirtioSharedMemoryRegion) UnmapMemoryAtOffsetSizeCompletionHandler(offset uint64, size uint64, handler ErrorHandler) {
	_block2, _ := NewErrorBlock(handler)
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("unmapMemoryAtOffset:size:completionHandler:"), offset, size, _block2)
}

func (v VZVirtioSharedMemoryRegion) RegionID() byte {
	rv := objc.SendIfResponds[byte](v.ID, objc.Sel("regionID"))
	return rv
}
func (v VZVirtioSharedMemoryRegion) Size() uint64 {
	rv := objc.SendIfResponds[uint64](v.ID, objc.Sel("size"))
	return rv
}

// MapMemoryAtOffsetSize is a synchronous wrapper around [VZVirtioSharedMemoryRegion.MapMemoryAtOffsetSizeCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (v VZVirtioSharedMemoryRegion) MapMemoryAtOffsetSize(ctx context.Context, memory unsafe.Pointer, offset uint64, size uint64) error {
	done := make(chan error, 1)
	v.MapMemoryAtOffsetSizeCompletionHandler(memory, offset, size, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// UnmapMemoryAtOffsetSize is a synchronous wrapper around [VZVirtioSharedMemoryRegion.UnmapMemoryAtOffsetSizeCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (v VZVirtioSharedMemoryRegion) UnmapMemoryAtOffsetSize(ctx context.Context, offset uint64, size uint64) error {
	done := make(chan error, 1)
	v.UnmapMemoryAtOffsetSizeCompletionHandler(offset, size, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
