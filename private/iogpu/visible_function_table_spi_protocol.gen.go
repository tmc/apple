// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MTLVisibleFunctionTableSPI protocol.
type MTLVisibleFunctionTableSPI interface {
	objectivec.IObject

	// BufferAddressAtIndex protocol.
	BufferAddressAtIndex(index uint64) uint64

	// GlobalBuffer protocol.
	GlobalBuffer() objectivec.IObject

	// GlobalBufferOffset protocol.
	GlobalBufferOffset() uint64

	// GpuAddress protocol.
	GpuAddress() uint64

	// GpuHandle protocol.
	GpuHandle() uint64

	// ResourceIndex protocol.
	ResourceIndex() uint64

	// SetBufferOffsetAtIndex protocol.
	SetBufferOffsetAtIndex(buffer objectivec.IObject, offset uint64, index uint64)

	// SetBuffersOffsetsWithRange protocol.
	SetBuffersOffsetsWithRange(buffers []objectivec.IObject, offsets *uint64, range_ foundation.NSRange)

	// SetGlobalBuffer protocol.
	SetGlobalBuffer(buffer objectivec.IObject)

	// SetGlobalBufferOffset protocol.
	SetGlobalBufferOffset(offset uint64)

	// SetValueAtIndex protocol.
	SetValueAtIndex(value uint64, index uint64)

	// SetValueWithRange protocol.
	SetValueWithRange(value uint64, range_ foundation.NSRange)

	// SetVisibleFunctionTableAtBufferIndex protocol.
	SetVisibleFunctionTableAtBufferIndex(table objectivec.IObject, index uint64)

	// SetVisibleFunctionTablesWithBufferRange protocol.
	SetVisibleFunctionTablesWithBufferRange(tables []objectivec.IObject, range_ foundation.NSRange)

	// UniqueIdentifier protocol.
	UniqueIdentifier() uint64
}

// MTLVisibleFunctionTableSPIObject wraps an existing Objective-C object that conforms to the MTLVisibleFunctionTableSPI protocol.
type MTLVisibleFunctionTableSPIObject struct {
	objectivec.Object
}

func (o MTLVisibleFunctionTableSPIObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTLVisibleFunctionTableSPIObjectFromID constructs a [MTLVisibleFunctionTableSPIObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLVisibleFunctionTableSPIObjectFromID(id objc.ID) MTLVisibleFunctionTableSPIObject {
	return MTLVisibleFunctionTableSPIObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o MTLVisibleFunctionTableSPIObject) BufferAddressAtIndex(index uint64) uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("bufferAddressAtIndex:"), index)
	return rv
}
func (o MTLVisibleFunctionTableSPIObject) GlobalBuffer() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("globalBuffer"))
	return objectivec.Object{ID: rv}
}
func (o MTLVisibleFunctionTableSPIObject) GlobalBufferOffset() uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("globalBufferOffset"))
	return rv
}
func (o MTLVisibleFunctionTableSPIObject) GpuAddress() uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("gpuAddress"))
	return rv
}
func (o MTLVisibleFunctionTableSPIObject) GpuHandle() uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("gpuHandle"))
	return rv
}
func (o MTLVisibleFunctionTableSPIObject) ResourceIndex() uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("resourceIndex"))
	return rv
}
func (o MTLVisibleFunctionTableSPIObject) SetBufferOffsetAtIndex(buffer objectivec.IObject, offset uint64, index uint64) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("setBuffer:offset:atIndex:"), buffer, offset, index)
}
func (o MTLVisibleFunctionTableSPIObject) SetBuffersOffsetsWithRange(buffers []objectivec.IObject, offsets *uint64, range_ foundation.NSRange) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("setBuffers:offsets:withRange:"), objectivec.IObjectSliceToNSArray(buffers), unsafe.Pointer(offsets), range_)
}
func (o MTLVisibleFunctionTableSPIObject) SetGlobalBuffer(buffer objectivec.IObject) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("setGlobalBuffer:"), buffer)
}
func (o MTLVisibleFunctionTableSPIObject) SetGlobalBufferOffset(offset uint64) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("setGlobalBufferOffset:"), offset)
}
func (o MTLVisibleFunctionTableSPIObject) SetValueAtIndex(value uint64, index uint64) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("setValue:atIndex:"), value, index)
}
func (o MTLVisibleFunctionTableSPIObject) SetValueWithRange(value uint64, range_ foundation.NSRange) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("setValue:withRange:"), value, range_)
}
func (o MTLVisibleFunctionTableSPIObject) SetVisibleFunctionTableAtBufferIndex(table objectivec.IObject, index uint64) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("setVisibleFunctionTable:atBufferIndex:"), table, index)
}
func (o MTLVisibleFunctionTableSPIObject) SetVisibleFunctionTablesWithBufferRange(tables []objectivec.IObject, range_ foundation.NSRange) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("setVisibleFunctionTables:withBufferRange:"), objectivec.IObjectSliceToNSArray(tables), range_)
}
func (o MTLVisibleFunctionTableSPIObject) UniqueIdentifier() uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("uniqueIdentifier"))
	return rv
}
