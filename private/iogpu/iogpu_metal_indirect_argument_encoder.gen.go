// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOGPUMetalIndirectArgumentEncoder] class.
var (
	_IOGPUMetalIndirectArgumentEncoderClass     IOGPUMetalIndirectArgumentEncoderClass
	_IOGPUMetalIndirectArgumentEncoderClassOnce sync.Once
)

func getIOGPUMetalIndirectArgumentEncoderClass() IOGPUMetalIndirectArgumentEncoderClass {
	_IOGPUMetalIndirectArgumentEncoderClassOnce.Do(func() {
		_IOGPUMetalIndirectArgumentEncoderClass = IOGPUMetalIndirectArgumentEncoderClass{class: objc.GetClass("IOGPUMetalIndirectArgumentEncoder")}
	})
	return _IOGPUMetalIndirectArgumentEncoderClass
}

// GetIOGPUMetalIndirectArgumentEncoderClass returns the class object for IOGPUMetalIndirectArgumentEncoder.
func GetIOGPUMetalIndirectArgumentEncoderClass() IOGPUMetalIndirectArgumentEncoderClass {
	return getIOGPUMetalIndirectArgumentEncoderClass()
}

type IOGPUMetalIndirectArgumentEncoderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOGPUMetalIndirectArgumentEncoderClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOGPUMetalIndirectArgumentEncoderClass) Alloc() IOGPUMetalIndirectArgumentEncoder {
	rv := objc.SendIfResponds[IOGPUMetalIndirectArgumentEncoder](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [IOGPUMetalIndirectArgumentEncoder.ConstantDataAtIndex]
//   - [IOGPUMetalIndirectArgumentEncoder.EncodedLength]
//   - [IOGPUMetalIndirectArgumentEncoder.SetBufferOffsetAtIndex]
//   - [IOGPUMetalIndirectArgumentEncoder.SetBuffersOffsetsWithRange]
//   - [IOGPUMetalIndirectArgumentEncoder.SetFunctionTableAtIndex]
//   - [IOGPUMetalIndirectArgumentEncoder.SetFunctionTablesWithRange]
//   - [IOGPUMetalIndirectArgumentEncoder.SetIndirectArgumentBufferOffset]
//   - [IOGPUMetalIndirectArgumentEncoder.SetIndirectCommandBufferAtIndex]
//   - [IOGPUMetalIndirectArgumentEncoder.SetIndirectCommandBuffersWithRange]
//   - [IOGPUMetalIndirectArgumentEncoder.SetIntersectionFunctionTableAtBufferIndex]
//   - [IOGPUMetalIndirectArgumentEncoder.SetIntersectionFunctionTableAtIndex]
//   - [IOGPUMetalIndirectArgumentEncoder.SetIntersectionFunctionTablesWithBufferRange]
//   - [IOGPUMetalIndirectArgumentEncoder.SetIntersectionFunctionTablesWithRange]
//   - [IOGPUMetalIndirectArgumentEncoder.SetSamplerStateAtIndex]
//   - [IOGPUMetalIndirectArgumentEncoder.SetSamplerStatesWithRange]
//   - [IOGPUMetalIndirectArgumentEncoder.SetTextureAtIndex]
//   - [IOGPUMetalIndirectArgumentEncoder.SetTexturesWithRange]
//   - [IOGPUMetalIndirectArgumentEncoder.SetVisibleFunctionTableAtBufferIndex]
//   - [IOGPUMetalIndirectArgumentEncoder.SetVisibleFunctionTableAtIndex]
//   - [IOGPUMetalIndirectArgumentEncoder.SetVisibleFunctionTablesWithBufferRange]
//   - [IOGPUMetalIndirectArgumentEncoder.SetVisibleFunctionTablesWithRange]
type IOGPUMetalIndirectArgumentEncoder struct {
	objectivec.Object
}

// IOGPUMetalIndirectArgumentEncoderFromID constructs a [IOGPUMetalIndirectArgumentEncoder] from an objc.ID.
func IOGPUMetalIndirectArgumentEncoderFromID(id objc.ID) IOGPUMetalIndirectArgumentEncoder {
	return IOGPUMetalIndirectArgumentEncoder{objectivec.Object{ID: id}}
}

// NOTE: IOGPUMetalIndirectArgumentEncoder embeds objectivec.Object because the parent type is
// unavailable, but IIOGPUMetalIndirectArgumentEncoder embeds IMTLIndirectArgumentEncoder, which that fallback
// cannot satisfy; skip compile-time assertion.

// An interface definition for the [IOGPUMetalIndirectArgumentEncoder] class.
//
// # Methods
//
//   - [IIOGPUMetalIndirectArgumentEncoder.ConstantDataAtIndex]
//   - [IIOGPUMetalIndirectArgumentEncoder.EncodedLength]
//   - [IIOGPUMetalIndirectArgumentEncoder.SetBufferOffsetAtIndex]
//   - [IIOGPUMetalIndirectArgumentEncoder.SetBuffersOffsetsWithRange]
//   - [IIOGPUMetalIndirectArgumentEncoder.SetFunctionTableAtIndex]
//   - [IIOGPUMetalIndirectArgumentEncoder.SetFunctionTablesWithRange]
//   - [IIOGPUMetalIndirectArgumentEncoder.SetIndirectArgumentBufferOffset]
//   - [IIOGPUMetalIndirectArgumentEncoder.SetIndirectCommandBufferAtIndex]
//   - [IIOGPUMetalIndirectArgumentEncoder.SetIndirectCommandBuffersWithRange]
//   - [IIOGPUMetalIndirectArgumentEncoder.SetIntersectionFunctionTableAtBufferIndex]
//   - [IIOGPUMetalIndirectArgumentEncoder.SetIntersectionFunctionTableAtIndex]
//   - [IIOGPUMetalIndirectArgumentEncoder.SetIntersectionFunctionTablesWithBufferRange]
//   - [IIOGPUMetalIndirectArgumentEncoder.SetIntersectionFunctionTablesWithRange]
//   - [IIOGPUMetalIndirectArgumentEncoder.SetSamplerStateAtIndex]
//   - [IIOGPUMetalIndirectArgumentEncoder.SetSamplerStatesWithRange]
//   - [IIOGPUMetalIndirectArgumentEncoder.SetTextureAtIndex]
//   - [IIOGPUMetalIndirectArgumentEncoder.SetTexturesWithRange]
//   - [IIOGPUMetalIndirectArgumentEncoder.SetVisibleFunctionTableAtBufferIndex]
//   - [IIOGPUMetalIndirectArgumentEncoder.SetVisibleFunctionTableAtIndex]
//   - [IIOGPUMetalIndirectArgumentEncoder.SetVisibleFunctionTablesWithBufferRange]
//   - [IIOGPUMetalIndirectArgumentEncoder.SetVisibleFunctionTablesWithRange]
type IIOGPUMetalIndirectArgumentEncoder interface {
	IMTLIndirectArgumentEncoder

	// Topic: Methods

	ConstantDataAtIndex(index uint64) unsafe.Pointer
	EncodedLength() uint64
	SetBufferOffsetAtIndex(buffer objectivec.IObject, offset uint64, index uint64)
	SetBuffersOffsetsWithRange(buffers []objectivec.IObject, offsets *uint64, range_ foundation.NSRange)
	SetFunctionTableAtIndex(table objectivec.IObject, index uint64)
	SetFunctionTablesWithRange(tables []objectivec.IObject, range_ foundation.NSRange)
	SetIndirectArgumentBufferOffset(buffer objectivec.IObject, offset uint64)
	SetIndirectCommandBufferAtIndex(buffer objectivec.IObject, index uint64)
	SetIndirectCommandBuffersWithRange(buffers []objectivec.IObject, range_ foundation.NSRange)
	SetIntersectionFunctionTableAtBufferIndex(table objectivec.IObject, index uint64)
	SetIntersectionFunctionTableAtIndex(table objectivec.IObject, index uint64)
	SetIntersectionFunctionTablesWithBufferRange(tables []objectivec.IObject, range_ foundation.NSRange)
	SetIntersectionFunctionTablesWithRange(tables []objectivec.IObject, range_ foundation.NSRange)
	SetSamplerStateAtIndex(state objectivec.IObject, index uint64)
	SetSamplerStatesWithRange(states []objectivec.IObject, range_ foundation.NSRange)
	SetTextureAtIndex(texture objectivec.IObject, index uint64)
	SetTexturesWithRange(textures []objectivec.IObject, range_ foundation.NSRange)
	SetVisibleFunctionTableAtBufferIndex(table objectivec.IObject, index uint64)
	SetVisibleFunctionTableAtIndex(table objectivec.IObject, index uint64)
	SetVisibleFunctionTablesWithBufferRange(tables []objectivec.IObject, range_ foundation.NSRange)
	SetVisibleFunctionTablesWithRange(tables []objectivec.IObject, range_ foundation.NSRange)
}

// Init initializes the instance.
func (i IOGPUMetalIndirectArgumentEncoder) Init() IOGPUMetalIndirectArgumentEncoder {
	rv := objc.SendIfResponds[IOGPUMetalIndirectArgumentEncoder](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i IOGPUMetalIndirectArgumentEncoder) Autorelease() IOGPUMetalIndirectArgumentEncoder {
	rv := objc.SendIfResponds[IOGPUMetalIndirectArgumentEncoder](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOGPUMetalIndirectArgumentEncoder creates a new IOGPUMetalIndirectArgumentEncoder instance.
func NewIOGPUMetalIndirectArgumentEncoder() IOGPUMetalIndirectArgumentEncoder {
	class := getIOGPUMetalIndirectArgumentEncoderClass()
	rv := objc.SendIfResponds[IOGPUMetalIndirectArgumentEncoder](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (i IOGPUMetalIndirectArgumentEncoder) ConstantDataAtIndex(index uint64) unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](i.ID, objc.Sel("constantDataAtIndex:"), index)
	return rv
}
func (i IOGPUMetalIndirectArgumentEncoder) SetBufferOffsetAtIndex(buffer objectivec.IObject, offset uint64, index uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setBuffer:offset:atIndex:"), buffer, offset, index)
}
func (i IOGPUMetalIndirectArgumentEncoder) SetBuffersOffsetsWithRange(buffers []objectivec.IObject, offsets *uint64, range_ foundation.NSRange) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setBuffers:offsets:withRange:"), objectivec.IObjectSliceToNSArray(buffers), unsafe.Pointer(offsets), range_)
}
func (i IOGPUMetalIndirectArgumentEncoder) SetFunctionTableAtIndex(table objectivec.IObject, index uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setFunctionTable:atIndex:"), table, index)
}
func (i IOGPUMetalIndirectArgumentEncoder) SetFunctionTablesWithRange(tables []objectivec.IObject, range_ foundation.NSRange) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setFunctionTables:withRange:"), objectivec.IObjectSliceToNSArray(tables), range_)
}
func (i IOGPUMetalIndirectArgumentEncoder) SetIndirectArgumentBufferOffset(buffer objectivec.IObject, offset uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setIndirectArgumentBuffer:offset:"), buffer, offset)
}
func (i IOGPUMetalIndirectArgumentEncoder) SetIndirectCommandBufferAtIndex(buffer objectivec.IObject, index uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setIndirectCommandBuffer:atIndex:"), buffer, index)
}
func (i IOGPUMetalIndirectArgumentEncoder) SetIndirectCommandBuffersWithRange(buffers []objectivec.IObject, range_ foundation.NSRange) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setIndirectCommandBuffers:withRange:"), objectivec.IObjectSliceToNSArray(buffers), range_)
}
func (i IOGPUMetalIndirectArgumentEncoder) SetIntersectionFunctionTableAtBufferIndex(table objectivec.IObject, index uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setIntersectionFunctionTable:atBufferIndex:"), table, index)
}
func (i IOGPUMetalIndirectArgumentEncoder) SetIntersectionFunctionTableAtIndex(table objectivec.IObject, index uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setIntersectionFunctionTable:atIndex:"), table, index)
}
func (i IOGPUMetalIndirectArgumentEncoder) SetIntersectionFunctionTablesWithBufferRange(tables []objectivec.IObject, range_ foundation.NSRange) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setIntersectionFunctionTables:withBufferRange:"), objectivec.IObjectSliceToNSArray(tables), range_)
}
func (i IOGPUMetalIndirectArgumentEncoder) SetIntersectionFunctionTablesWithRange(tables []objectivec.IObject, range_ foundation.NSRange) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setIntersectionFunctionTables:withRange:"), objectivec.IObjectSliceToNSArray(tables), range_)
}
func (i IOGPUMetalIndirectArgumentEncoder) SetSamplerStateAtIndex(state objectivec.IObject, index uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setSamplerState:atIndex:"), state, index)
}
func (i IOGPUMetalIndirectArgumentEncoder) SetSamplerStatesWithRange(states []objectivec.IObject, range_ foundation.NSRange) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setSamplerStates:withRange:"), objectivec.IObjectSliceToNSArray(states), range_)
}
func (i IOGPUMetalIndirectArgumentEncoder) SetTextureAtIndex(texture objectivec.IObject, index uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setTexture:atIndex:"), texture, index)
}
func (i IOGPUMetalIndirectArgumentEncoder) SetTexturesWithRange(textures []objectivec.IObject, range_ foundation.NSRange) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setTextures:withRange:"), objectivec.IObjectSliceToNSArray(textures), range_)
}
func (i IOGPUMetalIndirectArgumentEncoder) SetVisibleFunctionTableAtBufferIndex(table objectivec.IObject, index uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setVisibleFunctionTable:atBufferIndex:"), table, index)
}
func (i IOGPUMetalIndirectArgumentEncoder) SetVisibleFunctionTableAtIndex(table objectivec.IObject, index uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setVisibleFunctionTable:atIndex:"), table, index)
}
func (i IOGPUMetalIndirectArgumentEncoder) SetVisibleFunctionTablesWithBufferRange(tables []objectivec.IObject, range_ foundation.NSRange) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setVisibleFunctionTables:withBufferRange:"), objectivec.IObjectSliceToNSArray(tables), range_)
}
func (i IOGPUMetalIndirectArgumentEncoder) SetVisibleFunctionTablesWithRange(tables []objectivec.IObject, range_ foundation.NSRange) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setVisibleFunctionTables:withRange:"), objectivec.IObjectSliceToNSArray(tables), range_)
}

func (i IOGPUMetalIndirectArgumentEncoder) EncodedLength() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("encodedLength"))
	return rv
}
