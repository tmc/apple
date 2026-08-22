// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MTLIntersectionFunctionTableSPI protocol.
type MTLIntersectionFunctionTableSPI interface {
	objectivec.IObject

	// BufferAddressAtIndex protocol.
	BufferAddressAtIndex(index uint64) uint64

	// GlobalBuffer protocol.
	GlobalBuffer() objectivec.IObject

	// GlobalBufferOffset protocol.
	GlobalBufferOffset() uint64

	// GpuHandle protocol.
	GpuHandle() uint64

	// ResourceIndex protocol.
	ResourceIndex() uint64

	// SetGlobalBuffer protocol.
	SetGlobalBuffer(buffer objectivec.IObject)

	// SetGlobalBufferOffset protocol.
	SetGlobalBufferOffset(offset uint64)

	// UniqueIdentifier protocol.
	UniqueIdentifier() uint64
}

// MTLIntersectionFunctionTableSPIObject wraps an existing Objective-C object that conforms to the MTLIntersectionFunctionTableSPI protocol.
type MTLIntersectionFunctionTableSPIObject struct {
	objectivec.Object
}

func (o MTLIntersectionFunctionTableSPIObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTLIntersectionFunctionTableSPIObjectFromID constructs a [MTLIntersectionFunctionTableSPIObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLIntersectionFunctionTableSPIObjectFromID(id objc.ID) MTLIntersectionFunctionTableSPIObject {
	return MTLIntersectionFunctionTableSPIObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o MTLIntersectionFunctionTableSPIObject) BufferAddressAtIndex(index uint64) uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("bufferAddressAtIndex:"), index)
	return rv
}
func (o MTLIntersectionFunctionTableSPIObject) GlobalBuffer() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("globalBuffer"))
	return objectivec.Object{ID: rv}
}
func (o MTLIntersectionFunctionTableSPIObject) GlobalBufferOffset() uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("globalBufferOffset"))
	return rv
}
func (o MTLIntersectionFunctionTableSPIObject) GpuHandle() uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("gpuHandle"))
	return rv
}
func (o MTLIntersectionFunctionTableSPIObject) ResourceIndex() uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("resourceIndex"))
	return rv
}
func (o MTLIntersectionFunctionTableSPIObject) SetGlobalBuffer(buffer objectivec.IObject) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("setGlobalBuffer:"), buffer)
}
func (o MTLIntersectionFunctionTableSPIObject) SetGlobalBufferOffset(offset uint64) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("setGlobalBufferOffset:"), offset)
}
func (o MTLIntersectionFunctionTableSPIObject) UniqueIdentifier() uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("uniqueIdentifier"))
	return rv
}
