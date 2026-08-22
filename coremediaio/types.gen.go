// Code generated from Apple documentation for CoreMediaIO. DO NOT EDIT.

package coremediaio

import (
	"encoding/binary"
	"unsafe"

	"github.com/tmc/apple/corefoundation"
)

// C struct types

// CMIODeviceAVCCommand
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIODeviceAVCCommand
type CMIODeviceAVCCommand struct {
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
	_       [0]uint32
	storage [28]byte
}

// MCommand returns the MCommand field from the record's packed storage.
func (s *CMIODeviceAVCCommand) MCommand() *uint8 {
	return *(**uint8)(unsafe.Pointer(&s.storage[0]))
}

// SetMCommand updates the MCommand field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *CMIODeviceAVCCommand) SetMCommand(v *uint8) {
	*(**uint8)(unsafe.Pointer(&s.storage[0])) = v
}

// MCommandLength returns the MCommandLength field from the record's packed storage.
func (s *CMIODeviceAVCCommand) MCommandLength() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetMCommandLength updates the MCommandLength field in the record's packed storage.
func (s *CMIODeviceAVCCommand) SetMCommandLength(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[8:12], uint32(v))
}

// MResponse returns the MResponse field from the record's packed storage.
func (s *CMIODeviceAVCCommand) MResponse() *uint8 {
	return *(**uint8)(unsafe.Pointer(&s.storage[12]))
}

// SetMResponse updates the MResponse field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *CMIODeviceAVCCommand) SetMResponse(v *uint8) {
	*(**uint8)(unsafe.Pointer(&s.storage[12])) = v
}

// MResponseLength returns the MResponseLength field from the record's packed storage.
func (s *CMIODeviceAVCCommand) MResponseLength() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[20:24]))
}

// SetMResponseLength updates the MResponseLength field in the record's packed storage.
func (s *CMIODeviceAVCCommand) SetMResponseLength(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[20:24], uint32(v))
}

// MResponseUsed returns the MResponseUsed field from the record's packed storage.
func (s *CMIODeviceAVCCommand) MResponseUsed() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[24:28]))
}

// SetMResponseUsed updates the MResponseUsed field in the record's packed storage.
func (s *CMIODeviceAVCCommand) SetMResponseUsed(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[24:28], uint32(v))
}

// CMIODeviceRS422Command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIODeviceRS422Command
type CMIODeviceRS422Command struct {
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
	_       [0]uint32
	storage [28]byte
}

// MCommand returns the MCommand field from the record's packed storage.
func (s *CMIODeviceRS422Command) MCommand() *uint8 {
	return *(**uint8)(unsafe.Pointer(&s.storage[0]))
}

// SetMCommand updates the MCommand field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *CMIODeviceRS422Command) SetMCommand(v *uint8) {
	*(**uint8)(unsafe.Pointer(&s.storage[0])) = v
}

// MCommandLength returns the MCommandLength field from the record's packed storage.
func (s *CMIODeviceRS422Command) MCommandLength() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetMCommandLength updates the MCommandLength field in the record's packed storage.
func (s *CMIODeviceRS422Command) SetMCommandLength(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[8:12], uint32(v))
}

// MResponse returns the MResponse field from the record's packed storage.
func (s *CMIODeviceRS422Command) MResponse() *uint8 {
	return *(**uint8)(unsafe.Pointer(&s.storage[12]))
}

// SetMResponse updates the MResponse field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *CMIODeviceRS422Command) SetMResponse(v *uint8) {
	*(**uint8)(unsafe.Pointer(&s.storage[12])) = v
}

// MResponseLength returns the MResponseLength field from the record's packed storage.
func (s *CMIODeviceRS422Command) MResponseLength() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[20:24]))
}

// SetMResponseLength updates the MResponseLength field in the record's packed storage.
func (s *CMIODeviceRS422Command) SetMResponseLength(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[20:24], uint32(v))
}

// MResponseUsed returns the MResponseUsed field from the record's packed storage.
func (s *CMIODeviceRS422Command) MResponseUsed() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[24:28]))
}

// SetMResponseUsed updates the MResponseUsed field in the record's packed storage.
func (s *CMIODeviceRS422Command) SetMResponseUsed(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[24:28], uint32(v))
}

// CMIODeviceSMPTETimeCallback
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIODeviceSMPTETimeCallback
type CMIODeviceSMPTETimeCallback struct {
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
	_       [0]uint32
	storage [16]byte
}

// MGetSMPTETimeProc returns a raw C code address from the record's packed storage.
// The address is not callable from Go. To call C through it, use purego to
// bind a Go function value to the address.
func (s *CMIODeviceSMPTETimeCallback) MGetSMPTETimeProc() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[0:8]))
}

// SetMGetSMPTETimeProc stores a raw C code address in the record's packed storage.
// Use purego.NewCallback to obtain the address for a Go callback. The callback
// and its address must outlive every use C can make through this record.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *CMIODeviceSMPTETimeCallback) SetMGetSMPTETimeProc(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[0:8], uint64(v))
}

// MRefCon returns the MRefCon field from the record's packed storage.
func (s *CMIODeviceSMPTETimeCallback) MRefCon() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[8:16]))
}

// SetMRefCon updates the MRefCon field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *CMIODeviceSMPTETimeCallback) SetMRefCon(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[8:16], uint64(v))
}

// CMIODeviceStreamConfiguration
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIODeviceStreamConfiguration
type CMIODeviceStreamConfiguration struct {
	MNumberStreams uint32
}

// CMIOHardwarePlugInInterface
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOHardwarePlugInInterface
type CMIOHardwarePlugInInterface struct {
	_reserved                 unsafe.Pointer
	QueryInterface            func(unsafe.Pointer, corefoundation.CFUUIDBytes, unsafe.Pointer) int32
	AddRef                    func(unsafe.Pointer) uint32
	Release                   func(unsafe.Pointer) uint32
	Initialize                func(uintptr) int32
	InitializeWithObjectID    func(uintptr, uint32) int32
	Teardown                  func(uintptr) int32
	ObjectShow                func(uintptr, uint32)
	ObjectHasProperty         func(uintptr, uint32, uintptr) uint8
	ObjectIsPropertySettable  func(uintptr, uint32, uintptr, *byte) int32
	ObjectGetPropertyDataSize func(uintptr, uint32, uintptr, uint32, unsafe.Pointer, *uint32) int32
	ObjectGetPropertyData     func(uintptr, uint32, uintptr, uint32, unsafe.Pointer, uint32, *uint32, unsafe.Pointer) int32
	ObjectSetPropertyData     func(uintptr, uint32, uintptr, uint32, unsafe.Pointer, uint32, unsafe.Pointer) int32
	DeviceSuspend             func(uintptr, uint32) int32
	DeviceResume              func(uintptr, uint32) int32
	DeviceStartStream         func(uintptr, uint32, uint32) int32
	DeviceStopStream          func(uintptr, uint32, uint32) int32
	DeviceProcessAVCCommand   func(uintptr, uint32, uintptr) int32
	DeviceProcessRS422Command func(uintptr, uint32, uintptr) int32
	StreamCopyBufferQueue     func(uintptr, uint32, CMIODeviceStreamQueueAlteredProc, unsafe.Pointer, uintptr) int32
	StreamDeckPlay            func(uintptr, uint32) int32
	StreamDeckStop            func(uintptr, uint32) int32
	StreamDeckJog             func(uintptr, uint32, int32) int32
	StreamDeckCueTo           func(uintptr, uint32, float64, uint8) int32
}

// CMIOObjectPropertyAddress
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOObjectPropertyAddress
type CMIOObjectPropertyAddress struct {
	MSelector CMIOObjectPropertySelector
	MScope    CMIOObjectPropertyScope
	MElement  CMIOObjectPropertyElement
}

// CMIOStreamDeck
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOStreamDeck
type CMIOStreamDeck struct {
	MStatus uint32
	MState  uint32
	MState2 uint32
}

// CMIOStreamScheduledOutputNotificationProcAndRefCon
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOStreamScheduledOutputNotificationProcAndRefCon
type CMIOStreamScheduledOutputNotificationProcAndRefCon struct {
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
	_       [0]uint32
	storage [16]byte
}

// ScheduledOutputNotificationProc returns a raw C code address from the record's packed storage.
// The address is not callable from Go. To call C through it, use purego to
// bind a Go function value to the address.
func (s *CMIOStreamScheduledOutputNotificationProcAndRefCon) ScheduledOutputNotificationProc() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[0:8]))
}

// SetScheduledOutputNotificationProc stores a raw C code address in the record's packed storage.
// Use purego.NewCallback to obtain the address for a Go callback. The callback
// and its address must outlive every use C can make through this record.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *CMIOStreamScheduledOutputNotificationProcAndRefCon) SetScheduledOutputNotificationProc(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[0:8], uint64(v))
}

// ScheduledOutputNotificationRefCon returns the ScheduledOutputNotificationRefCon field from the record's packed storage.
func (s *CMIOStreamScheduledOutputNotificationProcAndRefCon) ScheduledOutputNotificationRefCon() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[8:16]))
}

// SetScheduledOutputNotificationRefCon updates the ScheduledOutputNotificationRefCon field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *CMIOStreamScheduledOutputNotificationProcAndRefCon) SetScheduledOutputNotificationRefCon(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[8:16], uint64(v))
}
