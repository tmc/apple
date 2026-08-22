// Code generated from Apple documentation for ParavirtualizedGraphics. DO NOT EDIT.

package paravirtualizedgraphics

import (
	"sync"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [PGDeviceDescriptor] class.
var (
	_PGDeviceDescriptorClass     PGDeviceDescriptorClass
	_PGDeviceDescriptorClassOnce sync.Once
)

func getPGDeviceDescriptorClass() PGDeviceDescriptorClass {
	_PGDeviceDescriptorClassOnce.Do(func() {
		_PGDeviceDescriptorClass = PGDeviceDescriptorClass{class: objc.GetClass("PGDeviceDescriptor")}
	})
	return _PGDeviceDescriptorClass
}

// GetPGDeviceDescriptorClass returns the class object for PGDeviceDescriptor.
func GetPGDeviceDescriptorClass() PGDeviceDescriptorClass {
	return getPGDeviceDescriptorClass()
}

type PGDeviceDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (pc PGDeviceDescriptorClass) Class() objc.Class {
	return pc.class
}

// Alloc allocates memory for a new instance of the class.
func (pc PGDeviceDescriptorClass) Alloc() PGDeviceDescriptor {
	rv := objc.Send[PGDeviceDescriptor](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// A description of the paravirtualized graphics device to create.
//
// # Specifying the GPU
//
//   - [PGDeviceDescriptor.Device]: The Metal device object to use to back the virtual graphics device.
//   - [PGDeviceDescriptor.SetDevice]
//
// # Managing Tasks
//
//   - [PGDeviceDescriptor.CreateTask]: A handler that the framework calls to create a task object.
//   - [PGDeviceDescriptor.SetCreateTask]
//   - [PGDeviceDescriptor.DestroyTask]: A handler that the framework calls to destroy a task object.
//   - [PGDeviceDescriptor.SetDestroyTask]
//
// # Managing Memory Operations
//
//   - [PGDeviceDescriptor.MapMemory]: A handler that the framework calls to map memory into the virtual machine.
//   - [PGDeviceDescriptor.SetMapMemory]
//   - [PGDeviceDescriptor.UnmapMemory]: A handler that the framework calls to unmap memory from the virtual machine.
//   - [PGDeviceDescriptor.SetUnmapMemory]
//   - [PGDeviceDescriptor.ReadMemory]: A handler that the framework calls to read data from the guest’s memory.
//   - [PGDeviceDescriptor.SetReadMemory]
//
// # Specifying Trace Behavior
//
//   - [PGDeviceDescriptor.AddTraceRange]: A handler that the framework calls to add a trace range.
//   - [PGDeviceDescriptor.SetAddTraceRange]
//   - [PGDeviceDescriptor.RemoveTraceRange]: A handler that the framework calls to remove a trace range.
//   - [PGDeviceDescriptor.SetRemoveTraceRange]
//
// # Handling Interrupts
//
//   - [PGDeviceDescriptor.RaiseInterrupt]: A handler that the system calls to raise an interrupt in the guest environment.
//   - [PGDeviceDescriptor.SetRaiseInterrupt]
//
// # Specifying Virtual Device Properties
//
//   - [PGDeviceDescriptor.MmioLength]: The length in bytes of the memory-mapped IO section.
//   - [PGDeviceDescriptor.SetMmioLength]
//
// # Instance Properties
//
//   - [PGDeviceDescriptor.DisplayPortCount]
//   - [PGDeviceDescriptor.SetDisplayPortCount]
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDeviceDescriptor
type PGDeviceDescriptor struct {
	objectivec.Object
}

// PGDeviceDescriptorFromID constructs a [PGDeviceDescriptor] from an objc.ID.
//
// A description of the paravirtualized graphics device to create.
func PGDeviceDescriptorFromID(id objc.ID) PGDeviceDescriptor {
	return PGDeviceDescriptor{objectivec.Object{ID: id}}
}

// NOTE: PGDeviceDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [PGDeviceDescriptor] class.
//
// # Specifying the GPU
//
//   - [IPGDeviceDescriptor.Device]: The Metal device object to use to back the virtual graphics device.
//   - [IPGDeviceDescriptor.SetDevice]
//
// # Managing Tasks
//
//   - [IPGDeviceDescriptor.CreateTask]: A handler that the framework calls to create a task object.
//   - [IPGDeviceDescriptor.SetCreateTask]
//   - [IPGDeviceDescriptor.DestroyTask]: A handler that the framework calls to destroy a task object.
//   - [IPGDeviceDescriptor.SetDestroyTask]
//
// # Managing Memory Operations
//
//   - [IPGDeviceDescriptor.MapMemory]: A handler that the framework calls to map memory into the virtual machine.
//   - [IPGDeviceDescriptor.SetMapMemory]
//   - [IPGDeviceDescriptor.UnmapMemory]: A handler that the framework calls to unmap memory from the virtual machine.
//   - [IPGDeviceDescriptor.SetUnmapMemory]
//   - [IPGDeviceDescriptor.ReadMemory]: A handler that the framework calls to read data from the guest’s memory.
//   - [IPGDeviceDescriptor.SetReadMemory]
//
// # Specifying Trace Behavior
//
//   - [IPGDeviceDescriptor.AddTraceRange]: A handler that the framework calls to add a trace range.
//   - [IPGDeviceDescriptor.SetAddTraceRange]
//   - [IPGDeviceDescriptor.RemoveTraceRange]: A handler that the framework calls to remove a trace range.
//   - [IPGDeviceDescriptor.SetRemoveTraceRange]
//
// # Handling Interrupts
//
//   - [IPGDeviceDescriptor.RaiseInterrupt]: A handler that the system calls to raise an interrupt in the guest environment.
//   - [IPGDeviceDescriptor.SetRaiseInterrupt]
//
// # Specifying Virtual Device Properties
//
//   - [IPGDeviceDescriptor.MmioLength]: The length in bytes of the memory-mapped IO section.
//   - [IPGDeviceDescriptor.SetMmioLength]
//
// # Instance Properties
//
//   - [IPGDeviceDescriptor.DisplayPortCount]
//   - [IPGDeviceDescriptor.SetDisplayPortCount]
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDeviceDescriptor
type IPGDeviceDescriptor interface {
	objectivec.IObject

	// Topic: Specifying the GPU

	// The Metal device object to use to back the virtual graphics device.
	Device() metal.MTLDevice
	SetDevice(value metal.MTLDevice)

	// Topic: Managing Tasks

	// A handler that the framework calls to create a task object.
	CreateTask() UintptrUint64Handler
	SetCreateTask(value UintptrUint64Handler)
	// A handler that the framework calls to destroy a task object.
	DestroyTask() PGDestroyTask
	SetDestroyTask(value objc.ID)

	// Topic: Managing Memory Operations

	// A handler that the framework calls to map memory into the virtual machine.
	MapMemory() PGMapMemory
	SetMapMemory(value objc.ID)
	// A handler that the framework calls to unmap memory from the virtual machine.
	UnmapMemory() PGUnmapMemory
	SetUnmapMemory(value objc.ID)
	// A handler that the framework calls to read data from the guest’s memory.
	ReadMemory() BoolUint64Handler
	SetReadMemory(value BoolUint64Handler)

	// Topic: Specifying Trace Behavior

	// A handler that the framework calls to add a trace range.
	AddTraceRange() UintptrPGPhysicalMemoryRange_sHandler
	SetAddTraceRange(value UintptrPGPhysicalMemoryRange_sHandler)
	// A handler that the framework calls to remove a trace range.
	RemoveTraceRange() PGRemoveTraceRange
	SetRemoveTraceRange(value objc.ID)

	// Topic: Handling Interrupts

	// A handler that the system calls to raise an interrupt in the guest environment.
	RaiseInterrupt() PGRaiseInterrupt
	SetRaiseInterrupt(value PGRaiseInterrupt)

	// Topic: Specifying Virtual Device Properties

	// The length in bytes of the memory-mapped IO section.
	MmioLength() uintptr
	SetMmioLength(value uintptr)

	// Topic: Instance Properties

	DisplayPortCount() uint32
	SetDisplayPortCount(value uint32)
}

// Init initializes the instance.
func (p PGDeviceDescriptor) Init() PGDeviceDescriptor {
	rv := objc.Send[PGDeviceDescriptor](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p PGDeviceDescriptor) Autorelease() PGDeviceDescriptor {
	rv := objc.Send[PGDeviceDescriptor](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewPGDeviceDescriptor creates a new PGDeviceDescriptor instance.
func NewPGDeviceDescriptor() PGDeviceDescriptor {
	class := getPGDeviceDescriptorClass()
	rv := objc.Send[PGDeviceDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The Metal device object to use to back the virtual graphics device.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDeviceDescriptor/device
func (p PGDeviceDescriptor) Device() metal.MTLDevice {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("device"))
	return metal.MTLDeviceObjectFromID(rv)
}
func (p PGDeviceDescriptor) SetDevice(value metal.MTLDevice) {
	objc.Send[struct{}](p.ID, objc.Sel("setDevice:"), value)
}

// A handler that the framework calls to create a task object.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDeviceDescriptor/createTask
func (p PGDeviceDescriptor) CreateTask() UintptrUint64Handler {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("createTask"))
	_ = rv
	return nil
}
func (p PGDeviceDescriptor) SetCreateTask(value UintptrUint64Handler) {
	block, cleanup := NewUintptrUint64Block(value)
	defer cleanup()
	objc.Send[struct{}](p.ID, objc.Sel("setCreateTask:"), block)
}

// A handler that the framework calls to destroy a task object.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDeviceDescriptor/destroyTask
func (p PGDeviceDescriptor) DestroyTask() PGDestroyTask {
	rv := objc.Send[PGDestroyTask](p.ID, objc.Sel("destroyTask"))
	return PGDestroyTask(rv)
}
func (p PGDeviceDescriptor) SetDestroyTask(value objc.ID) {
	objc.Send[struct{}](p.ID, objc.Sel("setDestroyTask:"), value)
}

// A handler that the framework calls to map memory into the virtual machine.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDeviceDescriptor/mapMemory
func (p PGDeviceDescriptor) MapMemory() PGMapMemory {
	rv := objc.Send[PGMapMemory](p.ID, objc.Sel("mapMemory"))
	return PGMapMemory(rv)
}
func (p PGDeviceDescriptor) SetMapMemory(value objc.ID) {
	objc.Send[struct{}](p.ID, objc.Sel("setMapMemory:"), value)
}

// A handler that the framework calls to unmap memory from the virtual
// machine.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDeviceDescriptor/unmapMemory
func (p PGDeviceDescriptor) UnmapMemory() PGUnmapMemory {
	rv := objc.Send[PGUnmapMemory](p.ID, objc.Sel("unmapMemory"))
	return PGUnmapMemory(rv)
}
func (p PGDeviceDescriptor) SetUnmapMemory(value objc.ID) {
	objc.Send[struct{}](p.ID, objc.Sel("setUnmapMemory:"), value)
}

// A handler that the framework calls to read data from the guest’s memory.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDeviceDescriptor/readMemory
func (p PGDeviceDescriptor) ReadMemory() BoolUint64Handler {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("readMemory"))
	_ = rv
	return nil
}
func (p PGDeviceDescriptor) SetReadMemory(value BoolUint64Handler) {
	block, cleanup := NewBoolUint64Block(value)
	defer cleanup()
	objc.Send[struct{}](p.ID, objc.Sel("setReadMemory:"), block)
}

// A handler that the framework calls to add a trace range.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDeviceDescriptor/addTraceRange
func (p PGDeviceDescriptor) AddTraceRange() UintptrPGPhysicalMemoryRange_sHandler {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("addTraceRange"))
	_ = rv
	return nil
}
func (p PGDeviceDescriptor) SetAddTraceRange(value UintptrPGPhysicalMemoryRange_sHandler) {
	block, cleanup := NewUintptrPGPhysicalMemoryRange_sBlock(value)
	defer cleanup()
	objc.Send[struct{}](p.ID, objc.Sel("setAddTraceRange:"), block)
}

// A handler that the framework calls to remove a trace range.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDeviceDescriptor/removeTraceRange
func (p PGDeviceDescriptor) RemoveTraceRange() PGRemoveTraceRange {
	rv := objc.Send[PGRemoveTraceRange](p.ID, objc.Sel("removeTraceRange"))
	return PGRemoveTraceRange(rv)
}
func (p PGDeviceDescriptor) SetRemoveTraceRange(value objc.ID) {
	objc.Send[struct{}](p.ID, objc.Sel("setRemoveTraceRange:"), value)
}

// A handler that the system calls to raise an interrupt in the guest
// environment.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDeviceDescriptor/raiseInterrupt
func (p PGDeviceDescriptor) RaiseInterrupt() PGRaiseInterrupt {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("raiseInterrupt"))
	_ = rv
	return nil
}
func (p PGDeviceDescriptor) SetRaiseInterrupt(value PGRaiseInterrupt) {
	block, cleanup := NewPGRaiseInterruptBlock(value)
	defer cleanup()
	objc.Send[struct{}](p.ID, objc.Sel("setRaiseInterrupt:"), block)
}

// The length in bytes of the memory-mapped IO section.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDeviceDescriptor/mmioLength
func (p PGDeviceDescriptor) MmioLength() uintptr {
	rv := objc.Send[uintptr](p.ID, objc.Sel("mmioLength"))
	return rv
}
func (p PGDeviceDescriptor) SetMmioLength(value uintptr) {
	objc.Send[struct{}](p.ID, objc.Sel("setMmioLength:"), value)
}

// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDeviceDescriptor/displayPortCount
func (p PGDeviceDescriptor) DisplayPortCount() uint32 {
	rv := objc.Send[uint32](p.ID, objc.Sel("displayPortCount"))
	return rv
}
func (p PGDeviceDescriptor) SetDisplayPortCount(value uint32) {
	objc.Send[struct{}](p.ID, objc.Sel("setDisplayPortCount:"), value)
}
