//go:build darwin

package ane

import (
	"fmt"

	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/iosurface"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// MetalDevice wraps a Metal GPU device for zero-copy interop with ANE.
type MetalDevice struct {
	device metal.MTLDeviceObject
}

// OpenMetal opens the system default Metal device.
func OpenMetal() (*MetalDevice, error) {
	ptr := metal.MTLCreateSystemDefaultDevice()
	if ptr == nil {
		return nil, fmt.Errorf("ane: no Metal device available")
	}
	dev := metal.MTLDeviceObjectFromID(objc.ID(uintptr(ptr)))
	return &MetalDevice{device: dev}, nil
}

// Close releases the Metal device.
func (d *MetalDevice) Close() error {
	return nil
}

// Device returns the underlying MTLDeviceObject.
func (d *MetalDevice) Device() metal.MTLDeviceObject { return d.device }

// MetalInputBuffer returns the i-th input IOSurface as a zero-copy MTLBuffer.
// The returned buffer shares physical pages with the ANE surface.
func (k *Kernel) MetalInputBuffer(d *MetalDevice, i int) (metal.MTLBufferObject, error) {
	if !k.mapped {
		return metal.MTLBufferObject{}, &ANEError{Op: "metal", Err: fmt.Errorf("kernel not mapped")}
	}
	if i < 0 || i >= len(k.inputs) {
		return metal.MTLBufferObject{}, fmt.Errorf("ane: input index %d out of range [0,%d)", i, len(k.inputs))
	}
	return metalBufferFromSurface(d, k.inputs[i])
}

// MetalOutputBuffer returns the i-th output IOSurface as a zero-copy MTLBuffer.
// The returned buffer shares physical pages with the ANE surface.
func (k *Kernel) MetalOutputBuffer(d *MetalDevice, i int) (metal.MTLBufferObject, error) {
	if !k.mapped {
		return metal.MTLBufferObject{}, &ANEError{Op: "metal", Err: fmt.Errorf("kernel not mapped")}
	}
	if i < 0 || i >= len(k.outputs) {
		return metal.MTLBufferObject{}, fmt.Errorf("ane: output index %d out of range [0,%d)", i, len(k.outputs))
	}
	return metalBufferFromSurface(d, k.outputs[i])
}

// metalBufferFromSurface wraps an IOSurface's memory as a zero-copy MTLBuffer.
// The buffer uses MTLResourceStorageModeShared (= 0) with no deallocator,
// since the IOSurface owns the backing memory.
func metalBufferFromSurface(d *MetalDevice, ref coregraphics.IOSurfaceRef) (metal.MTLBufferObject, error) {
	surfRef := iosurface.IOSurfaceRef(ref)
	base := iosurface.IOSurfaceGetBaseAddress(surfRef)
	if base == nil {
		return metal.MTLBufferObject{}, &ANEError{Op: "metal", Err: fmt.Errorf("IOSurface base address is nil")}
	}
	length := iosurface.IOSurfaceGetAllocSize(surfRef)
	if length == 0 {
		return metal.MTLBufferObject{}, &ANEError{Op: "metal", Err: fmt.Errorf("IOSurface alloc size is 0")}
	}

	buf := d.device.NewBufferWithBytesNoCopyLengthOptionsDeallocator(
		base,
		uint(length),
		metal.MTLResourceStorageModeShared,
		0, // no deallocator — IOSurface owns the memory
	)
	if buf == nil || buf.GetID() == 0 {
		return metal.MTLBufferObject{}, &ANEError{Op: "metal", Err: fmt.Errorf("failed to create MTLBuffer")}
	}
	return metal.MTLBufferObjectFromID(buf.GetID()), nil
}

// MetalSharedEvent creates a Metal shared event from an ANE SharedEvent.
// Both events share the same Mach port, enabling hardware-level signaling
// between the ANE and Metal GPU.
func (d *MetalDevice) MetalSharedEvent(ev *SharedEvent) (metal.MTLSharedEventObject, error) {
	port := ev.Port()
	if port == 0 {
		return metal.MTLSharedEventObject{}, fmt.Errorf("ane: shared event has no port")
	}
	mtlEv := d.device.NewSharedEventWithMachPort(port)
	if mtlEv == nil || mtlEv.GetID() == 0 {
		return metal.MTLSharedEventObject{}, fmt.Errorf("ane: failed to create MTLSharedEvent from port %d", port)
	}
	return metal.MTLSharedEventObjectFromID(mtlEv.GetID()), nil
}

// NewMetalSharedEvent creates a new shared event backed by the same Mach port
// on both the Metal and ANE sides, enabling hardware-level synchronization.
func (d *MetalDevice) NewMetalSharedEvent() (metal.MTLSharedEventObject, *SharedEvent, error) {
	// Create the ANE-side event first (it owns the Mach port).
	aneEvent, err := NewSharedEvent()
	if err != nil {
		return metal.MTLSharedEventObject{}, nil, fmt.Errorf("ane: %w", err)
	}

	// Bridge to Metal via the Mach port.
	port := aneEvent.Port()
	mtlEv := d.device.NewSharedEventWithMachPort(port)
	if mtlEv == nil || mtlEv.GetID() == 0 {
		aneEvent.Close()
		return metal.MTLSharedEventObject{}, nil, fmt.Errorf("ane: failed to bridge MTLSharedEvent from port %d", port)
	}

	return metal.MTLSharedEventObjectFromID(mtlEv.GetID()), aneEvent, nil
}
