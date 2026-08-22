// Code generated from Apple documentation for IOUSBHost. DO NOT EDIT.

package iousbhost

import (
	"encoding/binary"
	"unsafe"
)

// C struct types

// IOUSBHostCIMessage
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIMessage
type IOUSBHostCIMessage struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [16]byte
}

// Control returns the Control field from the record's packed storage.
func (s *IOUSBHostCIMessage) Control() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetControl updates the Control field in the record's packed storage.
func (s *IOUSBHostCIMessage) SetControl(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// Data0 returns the Data0 field from the record's packed storage.
func (s *IOUSBHostCIMessage) Data0() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetData0 updates the Data0 field in the record's packed storage.
func (s *IOUSBHostCIMessage) SetData0(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// Data1 returns the Data1 field from the record's packed storage.
func (s *IOUSBHostCIMessage) Data1() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[8:16]))
}

// SetData1 updates the Data1 field in the record's packed storage.
func (s *IOUSBHostCIMessage) SetData1(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[8:16], uint64(v))
}

// IOUSBHostIOSourceDescriptors - The descriptors for a single endpoint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostIOSourceDescriptors
type IOUSBHostIOSourceDescriptors struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.

	// A storage array alone has alignment 1. This zero-length array carries
	// the alignment C measured without contributing any size, so an
	// embedding record places this one where C does.
	_       [0]uint16
	storage [24]byte
}

// BcdUSB returns the BcdUSB field from the record's packed storage.
func (s *IOUSBHostIOSourceDescriptors) BcdUSB() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[0:2]))
}

// SetBcdUSB updates the BcdUSB field in the record's packed storage.
func (s *IOUSBHostIOSourceDescriptors) SetBcdUSB(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[0:2], uint16(v))
}

// Descriptor returns the Descriptor field from the record's packed storage.
func (s *IOUSBHostIOSourceDescriptors) Descriptor() [7]byte {
	return *(*[7]byte)(unsafe.Pointer(&s.storage[2]))
}

// SetDescriptor updates the Descriptor field in the record's packed storage.
func (s *IOUSBHostIOSourceDescriptors) SetDescriptor(v [7]byte) {
	*(*[7]byte)(unsafe.Pointer(&s.storage[2])) = v
}

// SsCompanionDescriptor returns the SsCompanionDescriptor field from the record's packed storage.
func (s *IOUSBHostIOSourceDescriptors) SsCompanionDescriptor() [6]byte {
	return *(*[6]byte)(unsafe.Pointer(&s.storage[9]))
}

// SetSsCompanionDescriptor updates the SsCompanionDescriptor field in the record's packed storage.
func (s *IOUSBHostIOSourceDescriptors) SetSsCompanionDescriptor(v [6]byte) {
	*(*[6]byte)(unsafe.Pointer(&s.storage[9])) = v
}

// SspCompanionDescriptor returns the SspCompanionDescriptor field from the record's packed storage.
func (s *IOUSBHostIOSourceDescriptors) SspCompanionDescriptor() [8]byte {
	return *(*[8]byte)(unsafe.Pointer(&s.storage[15]))
}

// SetSspCompanionDescriptor updates the SspCompanionDescriptor field in the record's packed storage.
func (s *IOUSBHostIOSourceDescriptors) SetSspCompanionDescriptor(v [8]byte) {
	*(*[8]byte)(unsafe.Pointer(&s.storage[15])) = v
}

// IOUSBHostIsochronousFrame - A structure that represents a single frame in an isochronous transfer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostIsochronousFrame
type IOUSBHostIsochronousFrame struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [24]byte
}

// Status returns the Status field from the record's packed storage.
func (s *IOUSBHostIsochronousFrame) Status() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetStatus updates the Status field in the record's packed storage.
func (s *IOUSBHostIsochronousFrame) SetStatus(v int32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// RequestCount returns the RequestCount field from the record's packed storage.
func (s *IOUSBHostIsochronousFrame) RequestCount() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetRequestCount updates the RequestCount field in the record's packed storage.
func (s *IOUSBHostIsochronousFrame) SetRequestCount(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// CompleteCount returns the CompleteCount field from the record's packed storage.
func (s *IOUSBHostIsochronousFrame) CompleteCount() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetCompleteCount updates the CompleteCount field in the record's packed storage.
func (s *IOUSBHostIsochronousFrame) SetCompleteCount(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[8:12], uint32(v))
}

// Reserved returns the Reserved field from the record's packed storage.
func (s *IOUSBHostIsochronousFrame) Reserved() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[12:16]))
}

// SetReserved updates the Reserved field in the record's packed storage.
func (s *IOUSBHostIsochronousFrame) SetReserved(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[12:16], uint32(v))
}

// TimeStamp returns the TimeStamp field from the record's packed storage.
func (s *IOUSBHostIsochronousFrame) TimeStamp() IOUSBHostTime {
	return IOUSBHostTime(binary.NativeEndian.Uint64(s.storage[16:24]))
}

// SetTimeStamp updates the TimeStamp field in the record's packed storage.
func (s *IOUSBHostIsochronousFrame) SetTimeStamp(v IOUSBHostTime) {
	binary.NativeEndian.PutUint64(s.storage[16:24], uint64(v))
}

// IOUSBHostIsochronousTransaction
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostIsochronousTransaction
type IOUSBHostIsochronousTransaction struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [28]byte
}

// Status returns the Status field from the record's packed storage.
func (s *IOUSBHostIsochronousTransaction) Status() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetStatus updates the Status field in the record's packed storage.
func (s *IOUSBHostIsochronousTransaction) SetStatus(v int32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// RequestCount returns the RequestCount field from the record's packed storage.
func (s *IOUSBHostIsochronousTransaction) RequestCount() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetRequestCount updates the RequestCount field in the record's packed storage.
func (s *IOUSBHostIsochronousTransaction) SetRequestCount(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// Offset returns the Offset field from the record's packed storage.
func (s *IOUSBHostIsochronousTransaction) Offset() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetOffset updates the Offset field in the record's packed storage.
func (s *IOUSBHostIsochronousTransaction) SetOffset(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[8:12], uint32(v))
}

// CompleteCount returns the CompleteCount field from the record's packed storage.
func (s *IOUSBHostIsochronousTransaction) CompleteCount() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[12:16]))
}

// SetCompleteCount updates the CompleteCount field in the record's packed storage.
func (s *IOUSBHostIsochronousTransaction) SetCompleteCount(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[12:16], uint32(v))
}

// TimeStamp returns the TimeStamp field from the record's packed storage.
func (s *IOUSBHostIsochronousTransaction) TimeStamp() IOUSBHostTime {
	return IOUSBHostTime(binary.NativeEndian.Uint64(s.storage[16:24]))
}

// SetTimeStamp updates the TimeStamp field in the record's packed storage.
func (s *IOUSBHostIsochronousTransaction) SetTimeStamp(v IOUSBHostTime) {
	binary.NativeEndian.PutUint64(s.storage[16:24], uint64(v))
}

// Options returns the Options field from the record's packed storage.
func (s *IOUSBHostIsochronousTransaction) Options() IOUSBHostIsochronousTransactionOptions {
	return *(*IOUSBHostIsochronousTransactionOptions)(unsafe.Pointer(&s.storage[24]))
}

// SetOptions updates the Options field in the record's packed storage.
func (s *IOUSBHostIsochronousTransaction) SetOptions(v IOUSBHostIsochronousTransactionOptions) {
	*(*IOUSBHostIsochronousTransactionOptions)(unsafe.Pointer(&s.storage[24])) = v
}
