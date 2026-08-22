// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MTLAccelerationStructureSPI protocol.
type MTLAccelerationStructureSPI interface {
	objectivec.IObject

	// AccelerationStructureUniqueIdentifier protocol.
	AccelerationStructureUniqueIdentifier() uint64

	// Buffer protocol.
	Buffer() objectivec.IObject

	// BufferOffset protocol.
	BufferOffset() uint64

	// Descriptor protocol.
	Descriptor() objectivec.IObject

	// GpuHandle protocol.
	GpuHandle() uint64

	// ResourceIndex protocol.
	ResourceIndex() uint64

	// SetDescriptor protocol.
	SetDescriptor(descriptor objectivec.IObject)

	// UniqueIdentifier protocol.
	UniqueIdentifier() uint64
}

// MTLAccelerationStructureSPIObject wraps an existing Objective-C object that conforms to the MTLAccelerationStructureSPI protocol.
type MTLAccelerationStructureSPIObject struct {
	objectivec.Object
}

func (o MTLAccelerationStructureSPIObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTLAccelerationStructureSPIObjectFromID constructs a [MTLAccelerationStructureSPIObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLAccelerationStructureSPIObjectFromID(id objc.ID) MTLAccelerationStructureSPIObject {
	return MTLAccelerationStructureSPIObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o MTLAccelerationStructureSPIObject) AccelerationStructureUniqueIdentifier() uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("accelerationStructureUniqueIdentifier"))
	return rv
}
func (o MTLAccelerationStructureSPIObject) Buffer() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("buffer"))
	return objectivec.Object{ID: rv}
}
func (o MTLAccelerationStructureSPIObject) BufferOffset() uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("bufferOffset"))
	return rv
}
func (o MTLAccelerationStructureSPIObject) Descriptor() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("descriptor"))
	return objectivec.Object{ID: rv}
}
func (o MTLAccelerationStructureSPIObject) GpuHandle() uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("gpuHandle"))
	return rv
}
func (o MTLAccelerationStructureSPIObject) ResourceIndex() uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("resourceIndex"))
	return rv
}
func (o MTLAccelerationStructureSPIObject) SetDescriptor(descriptor objectivec.IObject) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("setDescriptor:"), descriptor)
}
func (o MTLAccelerationStructureSPIObject) UniqueIdentifier() uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("uniqueIdentifier"))
	return rv
}
