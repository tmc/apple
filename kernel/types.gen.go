// Code generated from Apple documentation for kernel. DO NOT EDIT.

package kernel

import (
	"encoding/binary"
	"unsafe"

	"github.com/tmc/apple/objc"
)

// C struct types

// AVCCommandHandlerInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/avccommandhandlerinfo
type AVCCommandHandlerInfo struct {
}

// AVCSubunitInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/avcsubunitinfo
type AVCSubunitInfo struct {
	Create unsafe.Pointer
	Free   unsafe.Pointer
	Init   unsafe.Pointer
}

// FndrExtendedDirInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/fndrextendeddirinfo
type FndrExtendedDirInfo struct {
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
	storage [16]byte
}

// Document_id returns the Document_id field from the record's packed storage.
func (s *FndrExtendedDirInfo) Document_id() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetDocument_id updates the Document_id field in the record's packed storage.
func (s *FndrExtendedDirInfo) SetDocument_id(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// Date_added returns the Date_added field from the record's packed storage.
func (s *FndrExtendedDirInfo) Date_added() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetDate_added updates the Date_added field in the record's packed storage.
func (s *FndrExtendedDirInfo) SetDate_added(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// Extended_flags returns the Extended_flags field from the record's packed storage.
func (s *FndrExtendedDirInfo) Extended_flags() U_int16_t {
	return U_int16_t(binary.NativeEndian.Uint16(s.storage[8:10]))
}

// SetExtended_flags updates the Extended_flags field in the record's packed storage.
func (s *FndrExtendedDirInfo) SetExtended_flags(v U_int16_t) {
	binary.NativeEndian.PutUint16(s.storage[8:10], uint16(v))
}

// Reserved3 returns the Reserved3 field from the record's packed storage.
func (s *FndrExtendedDirInfo) Reserved3() U_int16_t {
	return U_int16_t(binary.NativeEndian.Uint16(s.storage[10:12]))
}

// SetReserved3 updates the Reserved3 field in the record's packed storage.
func (s *FndrExtendedDirInfo) SetReserved3(v U_int16_t) {
	binary.NativeEndian.PutUint16(s.storage[10:12], uint16(v))
}

// Write_gen_counter returns the Write_gen_counter field from the record's packed storage.
func (s *FndrExtendedDirInfo) Write_gen_counter() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[12:16]))
}

// SetWrite_gen_counter updates the Write_gen_counter field in the record's packed storage.
func (s *FndrExtendedDirInfo) SetWrite_gen_counter(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[12:16], uint32(v))
}

// FndrExtendedFileInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/fndrextendedfileinfo
type FndrExtendedFileInfo struct {
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
	storage [16]byte
}

// Document_id returns the Document_id field from the record's packed storage.
func (s *FndrExtendedFileInfo) Document_id() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetDocument_id updates the Document_id field in the record's packed storage.
func (s *FndrExtendedFileInfo) SetDocument_id(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// Date_added returns the Date_added field from the record's packed storage.
func (s *FndrExtendedFileInfo) Date_added() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetDate_added updates the Date_added field in the record's packed storage.
func (s *FndrExtendedFileInfo) SetDate_added(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// Extended_flags returns the Extended_flags field from the record's packed storage.
func (s *FndrExtendedFileInfo) Extended_flags() U_int16_t {
	return U_int16_t(binary.NativeEndian.Uint16(s.storage[8:10]))
}

// SetExtended_flags updates the Extended_flags field in the record's packed storage.
func (s *FndrExtendedFileInfo) SetExtended_flags(v U_int16_t) {
	binary.NativeEndian.PutUint16(s.storage[8:10], uint16(v))
}

// Reserved2 returns the Reserved2 field from the record's packed storage.
func (s *FndrExtendedFileInfo) Reserved2() U_int16_t {
	return U_int16_t(binary.NativeEndian.Uint16(s.storage[10:12]))
}

// SetReserved2 updates the Reserved2 field in the record's packed storage.
func (s *FndrExtendedFileInfo) SetReserved2(v U_int16_t) {
	binary.NativeEndian.PutUint16(s.storage[10:12], uint16(v))
}

// Write_gen_counter returns the Write_gen_counter field from the record's packed storage.
func (s *FndrExtendedFileInfo) Write_gen_counter() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[12:16]))
}

// SetWrite_gen_counter updates the Write_gen_counter field in the record's packed storage.
func (s *FndrExtendedFileInfo) SetWrite_gen_counter(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[12:16], uint32(v))
}

// IOACPIAddressSpaceDescriptor
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioacpiaddressspacedescriptor
type IOACPIAddressSpaceDescriptor struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [80]byte
}

// ResourceType returns the ResourceType field from the record's packed storage.
func (s *IOACPIAddressSpaceDescriptor) ResourceType() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetResourceType updates the ResourceType field in the record's packed storage.
func (s *IOACPIAddressSpaceDescriptor) SetResourceType(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// GeneralFlags returns the GeneralFlags field from the record's packed storage.
func (s *IOACPIAddressSpaceDescriptor) GeneralFlags() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetGeneralFlags updates the GeneralFlags field in the record's packed storage.
func (s *IOACPIAddressSpaceDescriptor) SetGeneralFlags(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// TypeSpecificFlags returns the TypeSpecificFlags field from the record's packed storage.
func (s *IOACPIAddressSpaceDescriptor) TypeSpecificFlags() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetTypeSpecificFlags updates the TypeSpecificFlags field in the record's packed storage.
func (s *IOACPIAddressSpaceDescriptor) SetTypeSpecificFlags(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[8:12], uint32(v))
}

// Reserved1 returns the Reserved1 field from the record's packed storage.
func (s *IOACPIAddressSpaceDescriptor) Reserved1() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[12:16]))
}

// SetReserved1 updates the Reserved1 field in the record's packed storage.
func (s *IOACPIAddressSpaceDescriptor) SetReserved1(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[12:16], uint32(v))
}

// Granularity returns the Granularity field from the record's packed storage.
func (s *IOACPIAddressSpaceDescriptor) Granularity() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[16:24]))
}

// SetGranularity updates the Granularity field in the record's packed storage.
func (s *IOACPIAddressSpaceDescriptor) SetGranularity(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[16:24], uint64(v))
}

// MinAddressRange returns the MinAddressRange field from the record's packed storage.
func (s *IOACPIAddressSpaceDescriptor) MinAddressRange() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[24:32]))
}

// SetMinAddressRange updates the MinAddressRange field in the record's packed storage.
func (s *IOACPIAddressSpaceDescriptor) SetMinAddressRange(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[24:32], uint64(v))
}

// MaxAddressRange returns the MaxAddressRange field from the record's packed storage.
func (s *IOACPIAddressSpaceDescriptor) MaxAddressRange() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[32:40]))
}

// SetMaxAddressRange updates the MaxAddressRange field in the record's packed storage.
func (s *IOACPIAddressSpaceDescriptor) SetMaxAddressRange(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[32:40], uint64(v))
}

// TranslationOffset returns the TranslationOffset field from the record's packed storage.
func (s *IOACPIAddressSpaceDescriptor) TranslationOffset() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[40:48]))
}

// SetTranslationOffset updates the TranslationOffset field in the record's packed storage.
func (s *IOACPIAddressSpaceDescriptor) SetTranslationOffset(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[40:48], uint64(v))
}

// AddressLength returns the AddressLength field from the record's packed storage.
func (s *IOACPIAddressSpaceDescriptor) AddressLength() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[48:56]))
}

// SetAddressLength updates the AddressLength field in the record's packed storage.
func (s *IOACPIAddressSpaceDescriptor) SetAddressLength(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[48:56], uint64(v))
}

// Reserved2 returns the Reserved2 field from the record's packed storage.
func (s *IOACPIAddressSpaceDescriptor) Reserved2() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[56:64]))
}

// SetReserved2 updates the Reserved2 field in the record's packed storage.
func (s *IOACPIAddressSpaceDescriptor) SetReserved2(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[56:64], uint64(v))
}

// Reserved3 returns the Reserved3 field from the record's packed storage.
func (s *IOACPIAddressSpaceDescriptor) Reserved3() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[64:72]))
}

// SetReserved3 updates the Reserved3 field in the record's packed storage.
func (s *IOACPIAddressSpaceDescriptor) SetReserved3(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[64:72], uint64(v))
}

// Reserved4 returns the Reserved4 field from the record's packed storage.
func (s *IOACPIAddressSpaceDescriptor) Reserved4() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[72:80]))
}

// SetReserved4 updates the Reserved4 field in the record's packed storage.
func (s *IOACPIAddressSpaceDescriptor) SetReserved4(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[72:80], uint64(v))
}

// IOATACommand
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioatacommand
type IOATACommand struct {
	Free unsafe.Pointer
	Init unsafe.Pointer
}

// IOATAReg16
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioatareg16
type IOATAReg16 struct {
}

// IOATAReg32
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioatareg32
type IOATAReg32 struct {
}

// IOATAReg8
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioatareg8
type IOATAReg8 struct {
}

// IOAddressSegment
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioaddresssegment
type IOAddressSegment struct {
	Address uint64
	Length  uint64
}

// IODMACommandSpecification
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iodmacommandspecification
type IODMACommandSpecification struct {
	Options        uint64
	MaxAddressBits uint64
	_resv          [16]uint64
}

// IOExternalMethodArguments
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioexternalmethodarguments
type IOExternalMethodArguments struct {
	ScalarInput                   *uint64
	Version                       uint32
	ScalarOutput                  *uint64
	StructureInput                unsafe.Pointer
	StructureOutputDescriptorSize uint32
	StructureInputSize            uint32
	AsyncReference                *Io_user_reference_t
	Selector                      uint32
	StructureInputDescriptor      *IOMemoryDescriptor
	StructureOutput               unsafe.Pointer
	StructureOutputDescriptor     *IOMemoryDescriptor
	ScalarInputCount              uint32
	AsyncWakePort                 uint32
	StructureVariableOutputData   **OSObject
	StructureOutputSize           uint32
	AsyncReferenceCount           uint32
	ScalarOutputCount             uint32
}

// IOFWAsyncPHYCommand - Send an async PHY packet
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofwasyncphycommand
type IOFWAsyncPHYCommand struct {
	Complete unsafe.Pointer
	Execute  unsafe.Pointer
	Free     unsafe.Pointer
	Reinit   unsafe.Pointer
}

// IOFWIsochChannel
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofwisochchannel
type IOFWIsochChannel struct {
	Free  unsafe.Pointer
	Init  unsafe.Pointer
	Start unsafe.Pointer
	Stop  unsafe.Pointer
}

// IOFireWireAVCAsynchronousCommand
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofirewireavcasynchronouscommand
type IOFireWireAVCAsynchronousCommand struct {
	Cancel unsafe.Pointer
	Free   unsafe.Pointer
	Init   unsafe.Pointer
	Reinit unsafe.Pointer
	Submit unsafe.Pointer
}

// IOFireWireBus - IOFireWireBus is a public class the provides access to general FireWire functionality...
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofirewirebus
type IOFireWireBus struct {
}

// IOFireWireMultiIsochReceivePacket
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofirewiremultiisochreceivepacket
type IOFireWireMultiIsochReceivePacket struct {
	Create unsafe.Pointer
	Free   unsafe.Pointer
	Init   unsafe.Pointer
}

// IOFireWireSBP2ORB - Represents an SBP2 normal command ORB. Supplies the APIs for configuring normal command ORBs. This includes setting the command block and writing the page tables for I/O. The ORBs are executed using the submitORB method in IOFireWireSBP2Login.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofirewiresbp2orb
type IOFireWireSBP2ORB struct {
	Free    unsafe.Pointer
	Release unsafe.Pointer // Primary implementation of the release mechanism.

}

// IOFramebuffer - The base class for graphics devices to be made available as part of the desktop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioframebuffer
type IOFramebuffer struct {
	Reserved   unsafe.Pointer
	Attach     unsafe.Pointer
	Close      unsafe.Pointer
	Free       unsafe.Pointer
	Initialize unsafe.Pointer
	Message    unsafe.Pointer
	Open       unsafe.Pointer
	Start      unsafe.Pointer
	Stop       unsafe.Pointer
	Terminate  unsafe.Pointer
}

// IOHIDEventService - IOService represents an device or OS service in IOKit and DriverKit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iohideventservice
type IOHIDEventService struct {
	Free    unsafe.Pointer
	Close   unsafe.Pointer
	Init    unsafe.Pointer
	Message unsafe.Pointer
	Open    unsafe.Pointer
	Start   unsafe.Pointer
	Stop    unsafe.Pointer
}

// IOInterruptDispatchSourcePayload
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iointerruptdispatchsourcepayload
type IOInterruptDispatchSourcePayload struct {
	Time  uint64
	Count uint64
}

// IOMemoryDescriptor - An abstract base class that defines common methods for describing physical or virtual memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iomemorydescriptor
type IOMemoryDescriptor struct {
	Free       unsafe.Pointer // Performs any final cleanup for the memory descriptor object.
	Reserved   unsafe.Pointer
	Map        unsafe.Pointer // Maps an IOMemoryDescriptor into the kernel map.
	Complete   unsafe.Pointer // Complete processing of the memory after an I/O transfer finishes.
	Initialize unsafe.Pointer
	Prepare    unsafe.Pointer // Prepare the memory for an I/O transfer.
	Redirect   unsafe.Pointer
}

// IONDRVControlParameters
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iondrvcontrolparameters
type IONDRVControlParameters struct {
	__reservedA [26]uint8
	Code        uint16
	Params      unsafe.Pointer
	__reservedB [18]uint8
}

// IONVRAMDescriptor
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ionvramdescriptor
type IONVRAMDescriptor struct {
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
	storage [6]byte
}

// Format returns the Format bitfield from the record's packed storage.
func (s *IONVRAMDescriptor) Format() uint64 {
	return uint64((s.storage[0] >> 0) & 0xf)
}

// SetFormat updates the Format bitfield in the record's packed storage.
func (s *IONVRAMDescriptor) SetFormat(v uint64) {
	s.storage[0] = (s.storage[0] &^ uint8(0xf<<0)) | uint8((uint8(v)&0xf)<<0)
}

// Marker returns the Marker bitfield from the record's packed storage.
func (s *IONVRAMDescriptor) Marker() uint64 {
	return uint64((s.storage[0] >> 4) & 0x1)
}

// SetMarker updates the Marker bitfield in the record's packed storage.
func (s *IONVRAMDescriptor) SetMarker(v uint64) {
	s.storage[0] = (s.storage[0] &^ uint8(0x1<<4)) | uint8((uint8(v)&0x1)<<4)
}

// BridgeCount returns the BridgeCount bitfield from the record's packed storage.
func (s *IONVRAMDescriptor) BridgeCount() uint64 {
	return uint64((s.storage[0] >> 5) & 0x7)
}

// SetBridgeCount updates the BridgeCount bitfield in the record's packed storage.
func (s *IONVRAMDescriptor) SetBridgeCount(v uint64) {
	s.storage[0] = (s.storage[0] &^ uint8(0x7<<5)) | uint8((uint8(v)&0x7)<<5)
}

// BusNum returns the BusNum bitfield from the record's packed storage.
func (s *IONVRAMDescriptor) BusNum() uint64 {
	return uint64((s.storage[1] >> 0) & 0x3)
}

// SetBusNum updates the BusNum bitfield in the record's packed storage.
func (s *IONVRAMDescriptor) SetBusNum(v uint64) {
	s.storage[1] = (s.storage[1] &^ uint8(0x3<<0)) | uint8((uint8(v)&0x3)<<0)
}

// BridgeDevices returns the BridgeDevices bitfield from the record's packed storage.
func (s *IONVRAMDescriptor) BridgeDevices() uint64 {
	return uint64((binary.NativeEndian.Uint32(s.storage[1:5]) >> 2) & 0x3fffffff)
}

// SetBridgeDevices updates the BridgeDevices bitfield in the record's packed storage.
func (s *IONVRAMDescriptor) SetBridgeDevices(v uint64) {
	binary.NativeEndian.PutUint32(s.storage[1:5], (binary.NativeEndian.Uint32(s.storage[1:5])&^uint32(0x3fffffff<<2))|uint32((uint32(v)&0x3fffffff)<<2))
}

// FunctionNum returns the FunctionNum bitfield from the record's packed storage.
func (s *IONVRAMDescriptor) FunctionNum() uint64 {
	return uint64((s.storage[5] >> 0) & 0x7)
}

// SetFunctionNum updates the FunctionNum bitfield in the record's packed storage.
func (s *IONVRAMDescriptor) SetFunctionNum(v uint64) {
	s.storage[5] = (s.storage[5] &^ uint8(0x7<<0)) | uint8((uint8(v)&0x7)<<0)
}

// DeviceNum returns the DeviceNum bitfield from the record's packed storage.
func (s *IONVRAMDescriptor) DeviceNum() uint64 {
	return uint64((s.storage[5] >> 3) & 0x1f)
}

// SetDeviceNum updates the DeviceNum bitfield in the record's packed storage.
func (s *IONVRAMDescriptor) SetDeviceNum(v uint64) {
	s.storage[5] = (s.storage[5] &^ uint8(0x1f<<3)) | uint8((uint8(v)&0x1f)<<3)
}

// IONotifier - An abstract base class defining common methods for controlling a notification request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ionotifier
type IONotifier struct {
	Disable unsafe.Pointer // Disables the notification request.
	Enable  unsafe.Pointer // Sets the enable state of the notification request.
	Remove  unsafe.Pointer // Removes the notification request and releases it.

}

// IOPCIPhysicalAddress
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iopciphysicaladdress
type IOPCIPhysicalAddress struct {
	PhysHi   [1]uint32
	PhysLo   uint32
	PhysMid  uint32
	LengthLo uint32
	LengthHi uint32
}

// IORPCMessageErrorReturn
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iorpcmessageerrorreturn
type IORPCMessageErrorReturn struct {
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
	storage [60]byte
}

// Mach returns the Mach field from the record's packed storage.
func (s *IORPCMessageErrorReturn) Mach() IORPCMessageMach {
	return *(*IORPCMessageMach)(unsafe.Pointer(&s.storage[0]))
}

// SetMach updates the Mach field in the record's packed storage.
func (s *IORPCMessageErrorReturn) SetMach(v IORPCMessageMach) {
	*(*IORPCMessageMach)(unsafe.Pointer(&s.storage[0])) = v
}

// Content returns the Content field from the record's packed storage.
func (s *IORPCMessageErrorReturn) Content() IORPCMessageErrorReturnContent {
	return *(*IORPCMessageErrorReturnContent)(unsafe.Pointer(&s.storage[28]))
}

// SetContent updates the Content field in the record's packed storage.
func (s *IORPCMessageErrorReturn) SetContent(v IORPCMessageErrorReturnContent) {
	*(*IORPCMessageErrorReturnContent)(unsafe.Pointer(&s.storage[28])) = v
}

// IORegistryEntry - The base class for all objects in the registry.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioregistryentry
type IORegistryEntry struct {
	Reserved   unsafe.Pointer
	Free       unsafe.Pointer // Standard free method for all IORegistryEntry subclasses.
	Init       unsafe.Pointer
	Initialize unsafe.Pointer
}

// IORegistryIterator - An iterator over the registry.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioregistryiterator
type IORegistryIterator struct {
	Free  unsafe.Pointer
	Reset unsafe.Pointer // Exits all levels of recursion, restoring the iterator to its state at creation.

}

// IOService - The base class for most I/O Kit families, devices, and drivers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioservice
type IOService struct {
	Reserved         unsafe.Pointer
	Start            unsafe.Pointer // During an IOService object's instantiation, starts the IOService object that has been selected to run on the provider.
	Stop             unsafe.Pointer // During an IOService termination, the stop method is called in its clients before they are detached & it is destroyed.
	Free             unsafe.Pointer // Frees data structures that were allocated when power management was initialized on this service.
	Init             unsafe.Pointer
	Attach           unsafe.Pointer // Attaches an IOService client to a provider in the I/O Registry.
	Close            unsafe.Pointer // Releases active access to a provider.
	Command_received unsafe.Pointer
	Detach           unsafe.Pointer // Detaches an IOService client from a provider in the I/O Registry.
	Finalize         unsafe.Pointer // Finalizes the destruction of an IOService object.
	Message          unsafe.Pointer // Receives a generic message delivered from an attached provider.
	Open             unsafe.Pointer // Requests active access to a provider.
	Probe            unsafe.Pointer // During an IOService object's instantiation, probes a matched service to see if it can be used.
	Terminate        unsafe.Pointer // Makes an IOService object inactive and begins its destruction.

}

// IOServiceInterestContent64
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioserviceinterestcontent64
type IOServiceInterestContent64 struct {
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
	storage [12]byte
}

// MessageType returns the MessageType field from the record's packed storage.
func (s *IOServiceInterestContent64) MessageType() Natural_t {
	return Natural_t(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetMessageType updates the MessageType field in the record's packed storage.
func (s *IOServiceInterestContent64) SetMessageType(v Natural_t) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// MessageArgument returns the MessageArgument field from the record's packed storage.
func (s *IOServiceInterestContent64) MessageArgument() [1]Io_user_reference_t {
	return *(*[1]Io_user_reference_t)(unsafe.Pointer(&s.storage[4]))
}

// SetMessageArgument updates the MessageArgument field in the record's packed storage.
func (s *IOServiceInterestContent64) SetMessageArgument(v [1]Io_user_reference_t) {
	*(*[1]Io_user_reference_t)(unsafe.Pointer(&s.storage[4])) = v
}

// IOTrackingCallSiteInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iotrackingcallsiteinfo
type IOTrackingCallSiteInfo struct {
	Count      uint32
	AddressPID int32
	Address    Mach_vm_address_t
	Size       [2]Mach_vm_size_t
	BtPID      int32
	Bt         [2][16]Mach_vm_address_t
}

// IOUSB30HubPortStatusExt - A structure that represents an extension to the USB 3.0 hub port status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iousb30hubportstatusext
type IOUSB30HubPortStatusExt struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [8]byte
}

// WPortStatus returns the WPortStatus field from the record's packed storage.
func (s *IOUSB30HubPortStatusExt) WPortStatus() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[0:2]))
}

// SetWPortStatus updates the WPortStatus field in the record's packed storage.
func (s *IOUSB30HubPortStatusExt) SetWPortStatus(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[0:2], uint16(v))
}

// WPortChange returns the WPortChange field from the record's packed storage.
func (s *IOUSB30HubPortStatusExt) WPortChange() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[2:4]))
}

// SetWPortChange updates the WPortChange field in the record's packed storage.
func (s *IOUSB30HubPortStatusExt) SetWPortChange(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[2:4], uint16(v))
}

// DwExtPortStatus returns the DwExtPortStatus field from the record's packed storage.
func (s *IOUSB30HubPortStatusExt) DwExtPortStatus() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetDwExtPortStatus updates the DwExtPortStatus field in the record's packed storage.
func (s *IOUSB30HubPortStatusExt) SetDwExtPortStatus(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// IOUSBStandardEndpointDescriptors - A container for descriptors for a single endpoint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iousbstandardendpointdescriptors
type IOUSBStandardEndpointDescriptors struct {
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
func (s *IOUSBStandardEndpointDescriptors) BcdUSB() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[0:2]))
}

// SetBcdUSB updates the BcdUSB field in the record's packed storage.
func (s *IOUSBStandardEndpointDescriptors) SetBcdUSB(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[0:2], uint16(v))
}

// Descriptor returns the Descriptor field from the record's packed storage.
func (s *IOUSBStandardEndpointDescriptors) Descriptor() IOUSBEndpointDescriptor {
	return *(*IOUSBEndpointDescriptor)(unsafe.Pointer(&s.storage[2]))
}

// SetDescriptor updates the Descriptor field in the record's packed storage.
func (s *IOUSBStandardEndpointDescriptors) SetDescriptor(v IOUSBEndpointDescriptor) {
	*(*IOUSBEndpointDescriptor)(unsafe.Pointer(&s.storage[2])) = v
}

// SsCompanionDescriptor returns the SsCompanionDescriptor field from the record's packed storage.
func (s *IOUSBStandardEndpointDescriptors) SsCompanionDescriptor() IOUSBSuperSpeedEndpointCompanionDescriptor {
	return *(*IOUSBSuperSpeedEndpointCompanionDescriptor)(unsafe.Pointer(&s.storage[9]))
}

// SetSsCompanionDescriptor updates the SsCompanionDescriptor field in the record's packed storage.
func (s *IOUSBStandardEndpointDescriptors) SetSsCompanionDescriptor(v IOUSBSuperSpeedEndpointCompanionDescriptor) {
	*(*IOUSBSuperSpeedEndpointCompanionDescriptor)(unsafe.Pointer(&s.storage[9])) = v
}

// SspCompanionDescriptor returns the SspCompanionDescriptor field from the record's packed storage.
func (s *IOUSBStandardEndpointDescriptors) SspCompanionDescriptor() IOUSBSuperSpeedPlusIsochronousEndpointCompanionDescriptor {
	return *(*IOUSBSuperSpeedPlusIsochronousEndpointCompanionDescriptor)(unsafe.Pointer(&s.storage[15]))
}

// SetSspCompanionDescriptor updates the SspCompanionDescriptor field in the record's packed storage.
func (s *IOUSBStandardEndpointDescriptors) SetSspCompanionDescriptor(v IOUSBSuperSpeedPlusIsochronousEndpointCompanionDescriptor) {
	*(*IOUSBSuperSpeedPlusIsochronousEndpointCompanionDescriptor)(unsafe.Pointer(&s.storage[15])) = v
}

// IOUserClientMethodArguments
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iouserclientmethodarguments
type IOUserClientMethodArguments struct {
	Version                    uint64
	Selector                   uint64
	Completion                 *OSAction
	ScalarInput                *uint64
	ScalarInputCount           uint32
	StructureInput             *OSData
	StructureInputDescriptor   *IOMemoryDescriptor
	ScalarOutput               *uint64
	ScalarOutputCount          uint32
	StructureOutput            *OSData
	StructureOutputDescriptor  *IOMemoryDescriptor
	StructureOutputMaximumSize uint64
	__reserved                 [30]uint64
}

// IOUserClientMethodDispatch
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iouserclientmethoddispatch
type IOUserClientMethodDispatch struct {
	Function                 unsafe.Pointer
	CheckCompletionExists    uint32
	CheckScalarInputCount    uint32
	CheckStructureInputSize  uint32
	CheckScalarOutputCount   uint32
	CheckStructureOutputSize uint32
}

// Key
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/key
type Key struct {
	_value    uint64
	_modified bool
}

// KeyAttribute
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/keyattribute
type KeyAttribute struct {
	_flags uint32
}

// OSAction
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/osaction
type OSAction struct {
	Free unsafe.Pointer // Performs any final cleanup for the action object.

}

// OSArray - OSArray provides an indexed store of objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/osarray
type OSArray struct {
	Free      unsafe.Pointer // Deallocates or releases any resources used by the OSArray instance.
	Merge     unsafe.Pointer // Appends the contents of an array onto the receiving array.
	Serialize unsafe.Pointer // Archives the receiver into the provided OSSerialize object.

}

// OSBoolean - OSBoolean wraps a boolean value in a C++ object for use in Libkern collections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/osboolean
type OSBoolean struct {
	Free       unsafe.Pointer // Overridden to prevent deallocation of the shared global instances.
	Initialize unsafe.Pointer
	Serialize  unsafe.Pointer // Archives the receiver into the provided OSSerialize object.

}

// OSClassDescription
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/osclassdescription
type OSClassDescription struct {
	DescriptionSize         uint32
	Name                    [96]int8
	SuperName               [96]int8
	MethodOptionsSize       uint32
	MethodOptionsOffset     uint32
	MetaMethodOptionsSize   uint32
	MetaMethodOptionsOffset uint32
	QueueNamesSize          uint32
	QueueNamesOffset        uint32
	MethodNamesSize         uint32
	MethodNamesOffset       uint32
	MetaMethodNamesSize     uint32
	MetaMethodNamesOffset   uint32
	Flags                   uint64
	Resv1                   [8]uint64
}

// OSCollection - The abstract superclass for Libkern collections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/oscollection
type OSCollection struct {
	Init unsafe.Pointer // Initializes the OSCollection object.

}

// OSCollectionIterator
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/oscollectioniterator
type OSCollectionIterator struct {
	Free  unsafe.Pointer // Releases or deallocates any resources used by the OSCollectionIterator object.
	Reset unsafe.Pointer // Resets the iterator to the beginning of the collection, as if it had just been created.

}

// OSData - OSData wraps an array of bytes in a C++ object for use in Libkern collections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/osdata
type OSData struct {
	Free      unsafe.Pointer // Deallocates or releases any resources used by the OSDictionary instance.
	Serialize unsafe.Pointer // Archives the receiver into the provided OSSerialize object.

}

// OSDictionary - OSDictionary provides an associative store using strings for keys.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/osdictionary
type OSDictionary struct {
	Free      unsafe.Pointer // Deallocates or releases any resources used by the OSDictionary instance.
	Merge     unsafe.Pointer // Merges the contents of a dictionary into the receiver.
	Serialize unsafe.Pointer // Archives the receiver into the provided OSSerialize object.

}

// OSIterator - The abstract superclass for Libkern iterators.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ositerator
type OSIterator struct {
	Reset unsafe.Pointer // Resets the iterator to the beginning of the collection, as if it had just been created.

}

// OSMetaClassBase - OSMetaClassBase is the abstract bootstrap class for the Libkern and I/O Kit run-time type information system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/osmetaclassbase
type OSMetaClassBase struct {
	Initialize unsafe.Pointer
	Release    unsafe.Pointer
	Retain     unsafe.Pointer // Abstract declaration of retain().
	Serialize  unsafe.Pointer // Abstract declaration of serialize.

}

// OSNotificationHeader64
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/osnotificationheader64
type OSNotificationHeader64 struct {
	Size      Mach_msg_size_t
	Type      Natural_t
	Reference [8]Io_user_reference_t
}

// OSNumber - OSNumber wraps an integer value in a C++ object for use in Libkern collections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/osnumber
type OSNumber struct {
	Free      unsafe.Pointer // Deallocates or releases any resources used by the OSNumber instance.
	Init      unsafe.Pointer
	Serialize unsafe.Pointer // Archives the receiver into the provided OSSerialize object.

}

// OSObject - OSObject is the concrete root class of the Libkern and I/O Kit C++ class hierarchy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/osobject
type OSObject struct {
	Free      unsafe.Pointer // Deallocates/releases resources held by the object.
	Init      unsafe.Pointer // Initializes a newly-allocated object.
	Release   unsafe.Pointer
	Retain    unsafe.Pointer // Retains a reference to the object.
	Serialize unsafe.Pointer // Overridden by subclasses to archive the receiver into the provided OSSerialize object.

}

// OSOrderedSet - OSOrderedSet provides an ordered set store of objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/osorderedset
type OSOrderedSet struct {
	Free   unsafe.Pointer // Deallocatesand releases any resources used by the OSOrderedSet instance.
	Init   unsafe.Pointer
	Member unsafe.Pointer // Checks the ordered set for the presence of an object.

}

// OSSerialize - OSSerialize coordinates serialization of Libkern C++ objects into an XML stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/osserialize
type OSSerialize struct {
	Free unsafe.Pointer
	Text unsafe.Pointer // Returns the XML text serialized so far.

}

// OSSerializer
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/osserializer
type OSSerializer struct {
	Free      unsafe.Pointer
	Serialize unsafe.Pointer
}

// OSSet - OSSet provides an unordered set store of objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/osset
type OSSet struct {
	Free      unsafe.Pointer // Deallocates or releases any resources used by the OSSet instance.
	Init      unsafe.Pointer
	Member    unsafe.Pointer // Checks the set for the presence of an object.
	Merge     unsafe.Pointer
	Serialize unsafe.Pointer // Archives the receiver into the provided OSSerialize object.

}

// OSString - OSString wraps a C string in a C++ object for use in Libkern collections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/osstring
type OSString struct {
	Free      unsafe.Pointer // Deallocates or releases any resources used by the OSString instance.
	Serialize unsafe.Pointer // Archives the receiver into the provided OSSerialize object.

}

// OSSymbol - OSSymbol wraps a C string in a unique C++ object for use as keys in Libkern collections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ossymbol
type OSSymbol struct {
	Free       unsafe.Pointer // Overrides OSObject::free to synchronize with the symbol pool.
	Initialize unsafe.Pointer
}

// StdFBShmem_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/stdfbshmem_t
type StdFBShmem_t struct {
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
	_       [0]uint64
	storage [5360]byte
}

// CursorSema returns the CursorSema field from the record's packed storage.
func (s *StdFBShmem_t) CursorSema() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetCursorSema updates the CursorSema field in the record's packed storage.
func (s *StdFBShmem_t) SetCursorSema(v int32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// Frame returns the Frame field from the record's packed storage.
func (s *StdFBShmem_t) Frame() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetFrame updates the Frame field in the record's packed storage.
func (s *StdFBShmem_t) SetFrame(v int32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// CursorShow returns the CursorShow field from the record's packed storage.
func (s *StdFBShmem_t) CursorShow() int8 {
	return int8(s.storage[8])
}

// SetCursorShow updates the CursorShow field in the record's packed storage.
func (s *StdFBShmem_t) SetCursorShow(v int8) {
	s.storage[8] = uint8(v)
}

// CursorObscured returns the CursorObscured field from the record's packed storage.
func (s *StdFBShmem_t) CursorObscured() int8 {
	return int8(s.storage[9])
}

// SetCursorObscured updates the CursorObscured field in the record's packed storage.
func (s *StdFBShmem_t) SetCursorObscured(v int8) {
	s.storage[9] = uint8(v)
}

// ShieldFlag returns the ShieldFlag field from the record's packed storage.
func (s *StdFBShmem_t) ShieldFlag() int8 {
	return int8(s.storage[10])
}

// SetShieldFlag updates the ShieldFlag field in the record's packed storage.
func (s *StdFBShmem_t) SetShieldFlag(v int8) {
	s.storage[10] = uint8(v)
}

// Shielded returns the Shielded field from the record's packed storage.
func (s *StdFBShmem_t) Shielded() int8 {
	return int8(s.storage[11])
}

// SetShielded updates the Shielded field in the record's packed storage.
func (s *StdFBShmem_t) SetShielded(v int8) {
	s.storage[11] = uint8(v)
}

// SaveRect returns the SaveRect field from the record's packed storage.
func (s *StdFBShmem_t) SaveRect() IOGBounds {
	return *(*IOGBounds)(unsafe.Pointer(&s.storage[12]))
}

// SetSaveRect updates the SaveRect field in the record's packed storage.
func (s *StdFBShmem_t) SetSaveRect(v IOGBounds) {
	*(*IOGBounds)(unsafe.Pointer(&s.storage[12])) = v
}

// ShieldRect returns the ShieldRect field from the record's packed storage.
func (s *StdFBShmem_t) ShieldRect() IOGBounds {
	return *(*IOGBounds)(unsafe.Pointer(&s.storage[20]))
}

// SetShieldRect updates the ShieldRect field in the record's packed storage.
func (s *StdFBShmem_t) SetShieldRect(v IOGBounds) {
	*(*IOGBounds)(unsafe.Pointer(&s.storage[20])) = v
}

// CursorLoc returns the CursorLoc field from the record's packed storage.
func (s *StdFBShmem_t) CursorLoc() IOGPoint {
	return *(*IOGPoint)(unsafe.Pointer(&s.storage[28]))
}

// SetCursorLoc updates the CursorLoc field in the record's packed storage.
func (s *StdFBShmem_t) SetCursorLoc(v IOGPoint) {
	*(*IOGPoint)(unsafe.Pointer(&s.storage[28])) = v
}

// CursorRect returns the CursorRect field from the record's packed storage.
func (s *StdFBShmem_t) CursorRect() IOGBounds {
	return *(*IOGBounds)(unsafe.Pointer(&s.storage[32]))
}

// SetCursorRect updates the CursorRect field in the record's packed storage.
func (s *StdFBShmem_t) SetCursorRect(v IOGBounds) {
	*(*IOGBounds)(unsafe.Pointer(&s.storage[32])) = v
}

// OldCursorRect returns the OldCursorRect field from the record's packed storage.
func (s *StdFBShmem_t) OldCursorRect() IOGBounds {
	return *(*IOGBounds)(unsafe.Pointer(&s.storage[40]))
}

// SetOldCursorRect updates the OldCursorRect field in the record's packed storage.
func (s *StdFBShmem_t) SetOldCursorRect(v IOGBounds) {
	*(*IOGBounds)(unsafe.Pointer(&s.storage[40])) = v
}

// ScreenBounds returns the ScreenBounds field from the record's packed storage.
func (s *StdFBShmem_t) ScreenBounds() IOGBounds {
	return *(*IOGBounds)(unsafe.Pointer(&s.storage[48]))
}

// SetScreenBounds updates the ScreenBounds field in the record's packed storage.
func (s *StdFBShmem_t) SetScreenBounds(v IOGBounds) {
	*(*IOGBounds)(unsafe.Pointer(&s.storage[48])) = v
}

// Version returns the Version field from the record's packed storage.
func (s *StdFBShmem_t) Version() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[56:60]))
}

// SetVersion updates the Version field in the record's packed storage.
func (s *StdFBShmem_t) SetVersion(v int32) {
	binary.NativeEndian.PutUint32(s.storage[56:60], uint32(v))
}

// StructSize returns the StructSize field from the record's packed storage.
func (s *StdFBShmem_t) StructSize() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[60:64]))
}

// SetStructSize updates the StructSize field in the record's packed storage.
func (s *StdFBShmem_t) SetStructSize(v int32) {
	binary.NativeEndian.PutUint32(s.storage[60:64], uint32(v))
}

// VblTime returns the VblTime field from the record's packed storage.
func (s *StdFBShmem_t) VblTime() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[64:72]))
}

// SetVblTime updates the VblTime field in the record's packed storage.
func (s *StdFBShmem_t) SetVblTime(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[64:72], uint64(v))
}

// VblDelta returns the VblDelta field from the record's packed storage.
func (s *StdFBShmem_t) VblDelta() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[72:80]))
}

// SetVblDelta updates the VblDelta field in the record's packed storage.
func (s *StdFBShmem_t) SetVblDelta(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[72:80], uint64(v))
}

// VblCount returns the VblCount field from the record's packed storage.
func (s *StdFBShmem_t) VblCount() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[80:88]))
}

// SetVblCount updates the VblCount field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *StdFBShmem_t) SetVblCount(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[80:88], uint64(v))
}

// ReservedC returns the ReservedC field from the record's packed storage.
func (s *StdFBShmem_t) ReservedC() [27]uint32 {
	return *(*[27]uint32)(unsafe.Pointer(&s.storage[88]))
}

// SetReservedC updates the ReservedC field in the record's packed storage.
func (s *StdFBShmem_t) SetReservedC(v [27]uint32) {
	*(*[27]uint32)(unsafe.Pointer(&s.storage[88])) = v
}

// HardwareCursorFlags returns the HardwareCursorFlags field from the record's packed storage.
func (s *StdFBShmem_t) HardwareCursorFlags() [4]uint8 {
	return *(*[4]uint8)(unsafe.Pointer(&s.storage[196]))
}

// SetHardwareCursorFlags updates the HardwareCursorFlags field in the record's packed storage.
func (s *StdFBShmem_t) SetHardwareCursorFlags(v [4]uint8) {
	*(*[4]uint8)(unsafe.Pointer(&s.storage[196])) = v
}

// HardwareCursorCapable returns the HardwareCursorCapable field from the record's packed storage.
func (s *StdFBShmem_t) HardwareCursorCapable() uint8 {
	return uint8(s.storage[200])
}

// SetHardwareCursorCapable updates the HardwareCursorCapable field in the record's packed storage.
func (s *StdFBShmem_t) SetHardwareCursorCapable(v uint8) {
	s.storage[200] = uint8(v)
}

// HardwareCursorActive returns the HardwareCursorActive field from the record's packed storage.
func (s *StdFBShmem_t) HardwareCursorActive() uint8 {
	return uint8(s.storage[201])
}

// SetHardwareCursorActive updates the HardwareCursorActive field in the record's packed storage.
func (s *StdFBShmem_t) SetHardwareCursorActive(v uint8) {
	s.storage[201] = uint8(v)
}

// HardwareCursorShields returns the HardwareCursorShields field from the record's packed storage.
func (s *StdFBShmem_t) HardwareCursorShields() uint8 {
	return uint8(s.storage[202])
}

// SetHardwareCursorShields updates the HardwareCursorShields field in the record's packed storage.
func (s *StdFBShmem_t) SetHardwareCursorShields(v uint8) {
	s.storage[202] = uint8(v)
}

// ReservedB returns the ReservedB field from the record's packed storage.
func (s *StdFBShmem_t) ReservedB() [1]uint8 {
	return *(*[1]uint8)(unsafe.Pointer(&s.storage[203]))
}

// SetReservedB updates the ReservedB field in the record's packed storage.
func (s *StdFBShmem_t) SetReservedB(v [1]uint8) {
	*(*[1]uint8)(unsafe.Pointer(&s.storage[203])) = v
}

// CursorSize returns the CursorSize field from the record's packed storage.
func (s *StdFBShmem_t) CursorSize() [4]IOGSize {
	return *(*[4]IOGSize)(unsafe.Pointer(&s.storage[204]))
}

// SetCursorSize updates the CursorSize field in the record's packed storage.
func (s *StdFBShmem_t) SetCursorSize(v [4]IOGSize) {
	*(*[4]IOGSize)(unsafe.Pointer(&s.storage[204])) = v
}

// HotSpot returns the HotSpot field from the record's packed storage.
func (s *StdFBShmem_t) HotSpot() [4]IOGPoint {
	return *(*[4]IOGPoint)(unsafe.Pointer(&s.storage[220]))
}

// SetHotSpot updates the HotSpot field in the record's packed storage.
func (s *StdFBShmem_t) SetHotSpot(v [4]IOGPoint) {
	*(*[4]IOGPoint)(unsafe.Pointer(&s.storage[220])) = v
}

// Cursor returns the Cursor field from the record's packed storage.
func (s *StdFBShmem_t) Cursor() [5124]byte {
	return *(*[5124]byte)(unsafe.Pointer(&s.storage[236]))
}

// SetCursor updates the Cursor field in the record's packed storage.
func (s *StdFBShmem_t) SetCursor(v [5124]byte) {
	*(*[5124]byte)(unsafe.Pointer(&s.storage[236])) = v
}

// Accessx_descriptor
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/accessx_descriptor
type Accessx_descriptor struct {
	Ad_name_offset uint32
	Ad_flags       int32
	Ad_pad         [2]int32
}

// Ah
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ah
type Ah struct {
	Ah_nxt     U_int8_t
	Ah_len     U_int8_t
	Ah_reserve U_int16_t
	Ah_spi     U_int32_t
}

// Applelabel
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/applelabel
type Applelabel struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [512]byte
}

// Al_boot0 returns the Al_boot0 field from the record's packed storage.
func (s *Applelabel) Al_boot0() [416]uint8 {
	return *(*[416]uint8)(unsafe.Pointer(&s.storage[0]))
}

// SetAl_boot0 updates the Al_boot0 field in the record's packed storage.
func (s *Applelabel) SetAl_boot0(v [416]uint8) {
	*(*[416]uint8)(unsafe.Pointer(&s.storage[0])) = v
}

// Al_magic returns the Al_magic field from the record's packed storage.
func (s *Applelabel) Al_magic() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[416:418]))
}

// SetAl_magic updates the Al_magic field in the record's packed storage.
func (s *Applelabel) SetAl_magic(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[416:418], uint16(v))
}

// Al_type returns the Al_type field from the record's packed storage.
func (s *Applelabel) Al_type() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[418:420]))
}

// SetAl_type updates the Al_type field in the record's packed storage.
func (s *Applelabel) SetAl_type(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[418:420], uint16(v))
}

// Al_flags returns the Al_flags field from the record's packed storage.
func (s *Applelabel) Al_flags() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[420:424]))
}

// SetAl_flags updates the Al_flags field in the record's packed storage.
func (s *Applelabel) SetAl_flags(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[420:424], uint32(v))
}

// Al_offset returns the Al_offset field from the record's packed storage.
func (s *Applelabel) Al_offset() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[424:432]))
}

// SetAl_offset updates the Al_offset field in the record's packed storage.
func (s *Applelabel) SetAl_offset(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[424:432], uint64(v))
}

// Al_size returns the Al_size field from the record's packed storage.
func (s *Applelabel) Al_size() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[432:436]))
}

// SetAl_size updates the Al_size field in the record's packed storage.
func (s *Applelabel) SetAl_size(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[432:436], uint32(v))
}

// Al_checksum returns the Al_checksum field from the record's packed storage.
func (s *Applelabel) Al_checksum() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[436:440]))
}

// SetAl_checksum updates the Al_checksum field in the record's packed storage.
func (s *Applelabel) SetAl_checksum(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[436:440], uint32(v))
}

// Al_boot1 returns the Al_boot1 field from the record's packed storage.
func (s *Applelabel) Al_boot1() [72]uint8 {
	return *(*[72]uint8)(unsafe.Pointer(&s.storage[440]))
}

// SetAl_boot1 updates the Al_boot1 field in the record's packed storage.
func (s *Applelabel) SetAl_boot1(v [72]uint8) {
	*(*[72]uint8)(unsafe.Pointer(&s.storage[440])) = v
}

// Arcade_upcall_subsystem
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/arcade_upcall_subsystem-4t8
type Arcade_upcall_subsystem struct {
	Server   unsafe.Pointer
	Start    Mach_msg_id_t
	End      Mach_msg_id_t
	Maxsize  uint32
	Reserved Vm_address_t
	Routine  unsafe.Pointer
}

// Arphdr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/arphdr
type Arphdr struct {
	Ar_hrd U_short
	Ar_pro U_short
	Ar_hln U_char
	Ar_pln U_char
	Ar_op  U_short
}

// Arpreq
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/arpreq
type Arpreq struct {
	Arp_pa    [16]byte
	Arp_ha    [16]byte
	Arp_flags int32
}

// Arpstat
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/arpstat
type Arpstat struct {
	Txrequests  uint32
	Txreplies   uint32
	Txannounces uint32
	Rxrequests  uint32
	Rxreplies   uint32
	Received    uint32
	Txconflicts uint32
	Invalidreqs uint32
	Reqnobufs   uint32
	Dropped     uint32
	Purged      uint32
	Timeouts    uint32
	Dupips      uint32
	Inuse       uint32
	Txurequests uint32
	Held        uint32
}

// Audit_triggers_subsystem
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/audit_triggers_subsystem-162
type Audit_triggers_subsystem struct {
	Server   unsafe.Pointer
	Start    Mach_msg_id_t
	End      Mach_msg_id_t
	Maxsize  uint32
	Reserved Vm_address_t
	Routine  unsafe.Pointer
}

// Backtrace_control
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/backtrace_control
type Backtrace_control struct {
	Btc_flags             Backtrace_flags_t
	Btc_frame_addr        uintptr
	Btc_user_thread       unsafe.Pointer
	Btc_user_copy         unsafe.Pointer
	Btc_user_copy_context unsafe.Pointer
	Btc_addr_offset       int64
}

// Backtrace_user_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/backtrace_user_info
type Backtrace_user_info struct {
	Btui_info              Backtrace_info_t
	Btui_error             Errno_t
	Btui_async_start_index uint32
	Btui_async_frame_addr  uintptr
	Btui_next_frame_addr   uintptr
}

// Bootp
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/bootp
type Bootp struct {
	Bp_op     U_char
	Bp_htype  U_char
	Bp_hlen   U_char
	Bp_hops   U_char
	Bp_xid    U_int32_t
	Bp_secs   U_short
	Bp_unused U_short
	Bp_ciaddr In_addr
	Bp_yiaddr In_addr
	Bp_siaddr In_addr
	Bp_giaddr In_addr
	Bp_chaddr [16]U_char
	Bp_sname  [64]U_char
	Bp_file   [128]U_char
	Bp_vend   [64]U_char
}

// Bootp_packet
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/bootp_packet
type Bootp_packet struct {
	Bp_ip    Ip
	Bp_udp   Udphdr
	Bp_bootp Bootp
}

// Bpf_dltlist
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/bpf_dltlist
type Bpf_dltlist struct {
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
	storage [12]byte
}

// Bfl_len returns the Bfl_len field from the record's packed storage.
func (s *Bpf_dltlist) Bfl_len() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetBfl_len updates the Bfl_len field in the record's packed storage.
func (s *Bpf_dltlist) SetBfl_len(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// Bfl_u returns the Bfl_u field from the record's packed storage.
func (s *Bpf_dltlist) Bfl_u() [2]uint32 {
	return *(*[2]uint32)(unsafe.Pointer(&s.storage[4]))
}

// SetBfl_u updates the Bfl_u field in the record's packed storage.
func (s *Bpf_dltlist) SetBfl_u(v [2]uint32) {
	*(*[2]uint32)(unsafe.Pointer(&s.storage[4])) = v
}

// Bpf_hdr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/bpf_hdr
type Bpf_hdr struct {
	Bh_tstamp  Timeval32
	Bh_caplen  Bpf_u_int32
	Bh_datalen Bpf_u_int32
	Bh_hdrlen  U_short
}

// Bpf_insn
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/bpf_insn
type Bpf_insn struct {
	Code U_short
	Jt   U_char
	Jf   U_char
	K    Bpf_u_int32
}

// Bpf_program
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/bpf_program
type Bpf_program struct {
	Bf_len   U_int
	Bf_insns *Bpf_insn
}

// Bpf_stat
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/bpf_stat
type Bpf_stat struct {
	Bs_recv U_int
	Bs_drop U_int
}

// Bpf_version
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/bpf_version
type Bpf_version struct {
	Bv_major U_short
	Bv_minor U_short
}

// Bt_params
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/bt_params
type Bt_params struct {
	Rate           float64
	Base_local_ts  uint64
	Base_remote_ts uint64
}

// Btinfo_sc_load_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/btinfo_sc_load_info
type Btinfo_sc_load_info struct {
	SharedCacheSlide       uint32
	SharedCacheUUID        [16]uint8
	SharedCacheBaseAddress uint32
}

// Btinfo_sc_load_info64
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/btinfo_sc_load_info64
type Btinfo_sc_load_info64 struct {
	SharedCacheSlide       uint64
	SharedCacheUUID        [16]uint8
	SharedCacheBaseAddress uint64
}

// Btinfo_thread_state_data_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/btinfo_thread_state_data_t
type Btinfo_thread_state_data_t struct {
	Flavor uint32
	Count  uint32
}

// Build_tool_version
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/build_tool_version
type Build_tool_version struct {
	Tool    uint32
	Version uint32
}

// Build_version_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/build_version_command
type Build_version_command struct {
	Cmd      uint32
	Cmdsize  uint32
	Platform uint32
	Minos    uint32
	Sdk      uint32
	Ntools   uint32
}

// Catch_exc_subsystem
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/catch_exc_subsystem-t6n
type Catch_exc_subsystem struct {
	Server   unsafe.Pointer
	Start    Mach_msg_id_t
	End      Mach_msg_id_t
	Maxsize  uint32
	Reserved Vm_address_t
	Routine  unsafe.Pointer
}

// Catch_mach_exc_subsystem
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/catch_mach_exc_subsystem-j9m
type Catch_mach_exc_subsystem struct {
	Server   unsafe.Pointer
	Start    Mach_msg_id_t
	End      Mach_msg_id_t
	Maxsize  uint32
	Reserved Vm_address_t
	Routine  unsafe.Pointer
}

// Chain_len_stats
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/chain_len_stats
type Chain_len_stats struct {
	Cls_one          uint64
	Cls_two          uint64
	Cls_three        uint64
	Cls_four         uint64
	Cls_five_or_more uint64
}

// Clock_reply_subsystem
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/clock_reply_subsystem
type Clock_reply_subsystem struct {
	Server   unsafe.Pointer
	Start    Mach_msg_id_t
	End      Mach_msg_id_t
	Maxsize  uint32
	Reserved Vm_address_t
	Routine  unsafe.Pointer
}

// Clockinfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/clockinfo
type Clockinfo struct {
	Hz      int32
	Tick    int32
	Tickadj int32
	Stathz  int32
	Profhz  int32
}

// Cmsghdr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/cmsghdr
type Cmsghdr struct {
	Cmsg_len   uint32
	Cmsg_level int32
	Cmsg_type  int32
}

// Coalition_notification_subsystem
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/coalition_notification_subsystem-36b
type Coalition_notification_subsystem struct {
	Server   unsafe.Pointer
	Start    Mach_msg_id_t
	End      Mach_msg_id_t
	Maxsize  uint32
	Reserved Vm_address_t
	Routine  unsafe.Pointer
}

// Codesigning_exit_reason_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/codesigning_exit_reason_info
type Codesigning_exit_reason_info struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [2108]byte
}

// Ceri_virt_addr returns the Ceri_virt_addr field from the record's packed storage.
func (s *Codesigning_exit_reason_info) Ceri_virt_addr() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[0:8]))
}

// SetCeri_virt_addr updates the Ceri_virt_addr field in the record's packed storage.
func (s *Codesigning_exit_reason_info) SetCeri_virt_addr(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[0:8], uint64(v))
}

// Ceri_file_offset returns the Ceri_file_offset field from the record's packed storage.
func (s *Codesigning_exit_reason_info) Ceri_file_offset() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[8:16]))
}

// SetCeri_file_offset updates the Ceri_file_offset field in the record's packed storage.
func (s *Codesigning_exit_reason_info) SetCeri_file_offset(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[8:16], uint64(v))
}

// Ceri_pathname returns the Ceri_pathname field from the record's packed storage.
func (s *Codesigning_exit_reason_info) Ceri_pathname() [1024]int8 {
	return *(*[1024]int8)(unsafe.Pointer(&s.storage[16]))
}

// SetCeri_pathname updates the Ceri_pathname field in the record's packed storage.
func (s *Codesigning_exit_reason_info) SetCeri_pathname(v [1024]int8) {
	*(*[1024]int8)(unsafe.Pointer(&s.storage[16])) = v
}

// Ceri_filename returns the Ceri_filename field from the record's packed storage.
func (s *Codesigning_exit_reason_info) Ceri_filename() [1024]int8 {
	return *(*[1024]int8)(unsafe.Pointer(&s.storage[1040]))
}

// SetCeri_filename updates the Ceri_filename field in the record's packed storage.
func (s *Codesigning_exit_reason_info) SetCeri_filename(v [1024]int8) {
	*(*[1024]int8)(unsafe.Pointer(&s.storage[1040])) = v
}

// Ceri_codesig_modtime_secs returns the Ceri_codesig_modtime_secs field from the record's packed storage.
func (s *Codesigning_exit_reason_info) Ceri_codesig_modtime_secs() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[2064:2072]))
}

// SetCeri_codesig_modtime_secs updates the Ceri_codesig_modtime_secs field in the record's packed storage.
func (s *Codesigning_exit_reason_info) SetCeri_codesig_modtime_secs(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[2064:2072], uint64(v))
}

// Ceri_codesig_modtime_nsecs returns the Ceri_codesig_modtime_nsecs field from the record's packed storage.
func (s *Codesigning_exit_reason_info) Ceri_codesig_modtime_nsecs() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[2072:2080]))
}

// SetCeri_codesig_modtime_nsecs updates the Ceri_codesig_modtime_nsecs field in the record's packed storage.
func (s *Codesigning_exit_reason_info) SetCeri_codesig_modtime_nsecs(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[2072:2080], uint64(v))
}

// Ceri_page_modtime_secs returns the Ceri_page_modtime_secs field from the record's packed storage.
func (s *Codesigning_exit_reason_info) Ceri_page_modtime_secs() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[2080:2088]))
}

// SetCeri_page_modtime_secs updates the Ceri_page_modtime_secs field in the record's packed storage.
func (s *Codesigning_exit_reason_info) SetCeri_page_modtime_secs(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[2080:2088], uint64(v))
}

// Ceri_page_modtime_nsecs returns the Ceri_page_modtime_nsecs field from the record's packed storage.
func (s *Codesigning_exit_reason_info) Ceri_page_modtime_nsecs() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[2088:2096]))
}

// SetCeri_page_modtime_nsecs updates the Ceri_page_modtime_nsecs field in the record's packed storage.
func (s *Codesigning_exit_reason_info) SetCeri_page_modtime_nsecs(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[2088:2096], uint64(v))
}

// Ceri_path_truncated returns the Ceri_path_truncated field from the record's packed storage.
func (s *Codesigning_exit_reason_info) Ceri_path_truncated() uint8 {
	return uint8(s.storage[2096])
}

// SetCeri_path_truncated updates the Ceri_path_truncated field in the record's packed storage.
func (s *Codesigning_exit_reason_info) SetCeri_path_truncated(v uint8) {
	s.storage[2096] = uint8(v)
}

// Ceri_object_codesigned returns the Ceri_object_codesigned field from the record's packed storage.
func (s *Codesigning_exit_reason_info) Ceri_object_codesigned() uint8 {
	return uint8(s.storage[2097])
}

// SetCeri_object_codesigned updates the Ceri_object_codesigned field in the record's packed storage.
func (s *Codesigning_exit_reason_info) SetCeri_object_codesigned(v uint8) {
	s.storage[2097] = uint8(v)
}

// Ceri_page_codesig_validated returns the Ceri_page_codesig_validated field from the record's packed storage.
func (s *Codesigning_exit_reason_info) Ceri_page_codesig_validated() uint8 {
	return uint8(s.storage[2098])
}

// SetCeri_page_codesig_validated updates the Ceri_page_codesig_validated field in the record's packed storage.
func (s *Codesigning_exit_reason_info) SetCeri_page_codesig_validated(v uint8) {
	s.storage[2098] = uint8(v)
}

// Ceri_page_codesig_tainted returns the Ceri_page_codesig_tainted field from the record's packed storage.
func (s *Codesigning_exit_reason_info) Ceri_page_codesig_tainted() uint8 {
	return uint8(s.storage[2099])
}

// SetCeri_page_codesig_tainted updates the Ceri_page_codesig_tainted field in the record's packed storage.
func (s *Codesigning_exit_reason_info) SetCeri_page_codesig_tainted(v uint8) {
	s.storage[2099] = uint8(v)
}

// Ceri_page_codesig_nx returns the Ceri_page_codesig_nx field from the record's packed storage.
func (s *Codesigning_exit_reason_info) Ceri_page_codesig_nx() uint8 {
	return uint8(s.storage[2100])
}

// SetCeri_page_codesig_nx updates the Ceri_page_codesig_nx field in the record's packed storage.
func (s *Codesigning_exit_reason_info) SetCeri_page_codesig_nx(v uint8) {
	s.storage[2100] = uint8(v)
}

// Ceri_page_wpmapped returns the Ceri_page_wpmapped field from the record's packed storage.
func (s *Codesigning_exit_reason_info) Ceri_page_wpmapped() uint8 {
	return uint8(s.storage[2101])
}

// SetCeri_page_wpmapped updates the Ceri_page_wpmapped field in the record's packed storage.
func (s *Codesigning_exit_reason_info) SetCeri_page_wpmapped(v uint8) {
	s.storage[2101] = uint8(v)
}

// Ceri_page_slid returns the Ceri_page_slid field from the record's packed storage.
func (s *Codesigning_exit_reason_info) Ceri_page_slid() uint8 {
	return uint8(s.storage[2102])
}

// SetCeri_page_slid updates the Ceri_page_slid field in the record's packed storage.
func (s *Codesigning_exit_reason_info) SetCeri_page_slid(v uint8) {
	s.storage[2102] = uint8(v)
}

// Ceri_page_dirty returns the Ceri_page_dirty field from the record's packed storage.
func (s *Codesigning_exit_reason_info) Ceri_page_dirty() uint8 {
	return uint8(s.storage[2103])
}

// SetCeri_page_dirty updates the Ceri_page_dirty field in the record's packed storage.
func (s *Codesigning_exit_reason_info) SetCeri_page_dirty(v uint8) {
	s.storage[2103] = uint8(v)
}

// Ceri_page_shadow_depth returns the Ceri_page_shadow_depth field from the record's packed storage.
func (s *Codesigning_exit_reason_info) Ceri_page_shadow_depth() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[2104:2108]))
}

// SetCeri_page_shadow_depth updates the Ceri_page_shadow_depth field in the record's packed storage.
func (s *Codesigning_exit_reason_info) SetCeri_page_shadow_depth(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[2104:2108], uint32(v))
}

// Crashinfo_jit_address_range
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/crashinfo_jit_address_range
type Crashinfo_jit_address_range struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [16]byte
}

// Start_address returns the Start_address field from the record's packed storage.
func (s *Crashinfo_jit_address_range) Start_address() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[0:8]))
}

// SetStart_address updates the Start_address field in the record's packed storage.
func (s *Crashinfo_jit_address_range) SetStart_address(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[0:8], uint64(v))
}

// End_address returns the End_address field from the record's packed storage.
func (s *Crashinfo_jit_address_range) End_address() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[8:16]))
}

// SetEnd_address updates the End_address field in the record's packed storage.
func (s *Crashinfo_jit_address_range) SetEnd_address(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[8:16], uint64(v))
}

// Crashinfo_mb
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/crashinfo_mb
type Crashinfo_mb struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [520]byte
}

// Start_address returns the Start_address field from the record's packed storage.
func (s *Crashinfo_mb) Start_address() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[0:8]))
}

// SetStart_address updates the Start_address field in the record's packed storage.
func (s *Crashinfo_mb) SetStart_address(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[0:8], uint64(v))
}

// Data returns the Data field from the record's packed storage.
func (s *Crashinfo_mb) Data() [64]uint64 {
	return *(*[64]uint64)(unsafe.Pointer(&s.storage[8]))
}

// SetData updates the Data field in the record's packed storage.
func (s *Crashinfo_mb) SetData(v [64]uint64) {
	*(*[64]uint64)(unsafe.Pointer(&s.storage[8])) = v
}

// Crashinfo_proc_uniqidentifierinfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/crashinfo_proc_uniqidentifierinfo
type Crashinfo_proc_uniqidentifierinfo struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [56]byte
}

// P_uuid returns the P_uuid field from the record's packed storage.
func (s *Crashinfo_proc_uniqidentifierinfo) P_uuid() [16]uint8 {
	return *(*[16]uint8)(unsafe.Pointer(&s.storage[0]))
}

// SetP_uuid updates the P_uuid field in the record's packed storage.
func (s *Crashinfo_proc_uniqidentifierinfo) SetP_uuid(v [16]uint8) {
	*(*[16]uint8)(unsafe.Pointer(&s.storage[0])) = v
}

// P_uniqueid returns the P_uniqueid field from the record's packed storage.
func (s *Crashinfo_proc_uniqidentifierinfo) P_uniqueid() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[16:24]))
}

// SetP_uniqueid updates the P_uniqueid field in the record's packed storage.
func (s *Crashinfo_proc_uniqidentifierinfo) SetP_uniqueid(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[16:24], uint64(v))
}

// P_puniqueid returns the P_puniqueid field from the record's packed storage.
func (s *Crashinfo_proc_uniqidentifierinfo) P_puniqueid() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[24:32]))
}

// SetP_puniqueid updates the P_puniqueid field in the record's packed storage.
func (s *Crashinfo_proc_uniqidentifierinfo) SetP_puniqueid(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[24:32], uint64(v))
}

// P_reserve2 returns the P_reserve2 field from the record's packed storage.
func (s *Crashinfo_proc_uniqidentifierinfo) P_reserve2() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[32:40]))
}

// SetP_reserve2 updates the P_reserve2 field in the record's packed storage.
func (s *Crashinfo_proc_uniqidentifierinfo) SetP_reserve2(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[32:40], uint64(v))
}

// P_reserve3 returns the P_reserve3 field from the record's packed storage.
func (s *Crashinfo_proc_uniqidentifierinfo) P_reserve3() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[40:48]))
}

// SetP_reserve3 updates the P_reserve3 field in the record's packed storage.
func (s *Crashinfo_proc_uniqidentifierinfo) SetP_reserve3(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[40:48], uint64(v))
}

// P_reserve4 returns the P_reserve4 field from the record's packed storage.
func (s *Crashinfo_proc_uniqidentifierinfo) P_reserve4() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[48:56]))
}

// SetP_reserve4 updates the P_reserve4 field in the record's packed storage.
func (s *Crashinfo_proc_uniqidentifierinfo) SetP_reserve4(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[48:56], uint64(v))
}

// Ctl_event_data
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ctl_event_data
type Ctl_event_data struct {
	Ctl_id   U_int32_t // The kernel control id.
	Ctl_unit U_int32_t // The kernel control unit.

}

// Ctl_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ctl_info
type Ctl_info struct {
	Ctl_id   U_int32_t // The kernel control id, filled out upon return.
	Ctl_name [96]int8  // The kernel control name to find.

}

// Ctlname
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ctlname
type Ctlname struct {
	Ctl_name *byte
	Ctl_type int32
}

// Data_in_code_entry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/data_in_code_entry
type Data_in_code_entry struct {
	Offset uint32
	Length uint16
	Kind   uint16
}

// Dirent
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dirent
type Dirent struct {
	D_ino     uint64
	D_seekoff uint64
	D_reclen  uint16
	D_namlen  uint16
	D_type    uint8
	D_name    [1024]int8
}

// Disk_blk0
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/disk_blk0
type Disk_blk0 struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [512]byte
}

// Bootcode returns the Bootcode field from the record's packed storage.
func (s *Disk_blk0) Bootcode() [446]uint8 {
	return *(*[446]uint8)(unsafe.Pointer(&s.storage[0]))
}

// SetBootcode updates the Bootcode field in the record's packed storage.
func (s *Disk_blk0) SetBootcode(v [446]uint8) {
	*(*[446]uint8)(unsafe.Pointer(&s.storage[0])) = v
}

// Parts returns the Parts field from the record's packed storage.
func (s *Disk_blk0) Parts() [4]Fdisk_part {
	return *(*[4]Fdisk_part)(unsafe.Pointer(&s.storage[446]))
}

// SetParts updates the Parts field in the record's packed storage.
func (s *Disk_blk0) SetParts(v [4]Fdisk_part) {
	*(*[4]Fdisk_part)(unsafe.Pointer(&s.storage[446])) = v
}

// Signature returns the Signature field from the record's packed storage.
func (s *Disk_blk0) Signature() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[510:512]))
}

// SetSignature updates the Signature field in the record's packed storage.
func (s *Disk_blk0) SetSignature(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[510:512], uint16(v))
}

// Do_notify_subsystem
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/do_notify_subsystem-q2j
type Do_notify_subsystem struct {
	Server   unsafe.Pointer
	Start    Mach_msg_id_t
	End      Mach_msg_id_t
	Maxsize  uint32
	Reserved Vm_address_t
	Routine  unsafe.Pointer
}

// Doubleagent_subsystem
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/doubleagent_subsystem-1pl
type Doubleagent_subsystem struct {
	Server   unsafe.Pointer
	Start    Mach_msg_id_t
	End      Mach_msg_id_t
	Maxsize  uint32
	Reserved Vm_address_t
	Routine  unsafe.Pointer
}

// Dyld_aot_cache_uuid_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dyld_aot_cache_uuid_info
type Dyld_aot_cache_uuid_info struct {
	X86SlidBaseAddress uint64
	X86UUID            [16]uint8
	AotSlidBaseAddress uint64
	AotUUID            [16]uint8
}

// Dyld_chained_fixups_header
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dyld_chained_fixups_header
type Dyld_chained_fixups_header struct {
	Fixups_version uint32
	Starts_offset  uint32
	Imports_offset uint32
	Symbols_offset uint32
	Imports_count  uint32
	Imports_format uint32
	Symbols_format uint32
}

// Dyld_chained_import
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dyld_chained_import
type Dyld_chained_import struct {
	bitfield0 uint32
}

// Lib_ordinal returns the Lib_ordinal bitfield.
func (s *Dyld_chained_import) Lib_ordinal() uint32 {
	return (s.bitfield0 >> 0) & ((1 << 8) - 1)
}

// SetLib_ordinal updates the Lib_ordinal bitfield.
func (s *Dyld_chained_import) SetLib_ordinal(v uint32) {
	const mask uint32 = (1 << 8) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 0)) | ((v & mask) << 0)
}

// Weak_import returns the Weak_import bitfield.
func (s *Dyld_chained_import) Weak_import() uint32 {
	return (s.bitfield0 >> 8) & ((1 << 1) - 1)
}

// SetWeak_import updates the Weak_import bitfield.
func (s *Dyld_chained_import) SetWeak_import(v uint32) {
	const mask uint32 = (1 << 1) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 8)) | ((v & mask) << 8)
}

// Name_offset returns the Name_offset bitfield.
func (s *Dyld_chained_import) Name_offset() uint32 {
	return (s.bitfield0 >> 9) & ((1 << 23) - 1)
}

// SetName_offset updates the Name_offset bitfield.
func (s *Dyld_chained_import) SetName_offset(v uint32) {
	const mask uint32 = (1 << 23) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 9)) | ((v & mask) << 9)
}

// Dyld_chained_import_addend
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dyld_chained_import_addend
type Dyld_chained_import_addend struct {
	bitfield0 uint32
	Addend    int32
}

// Lib_ordinal returns the Lib_ordinal bitfield.
func (s *Dyld_chained_import_addend) Lib_ordinal() uint32 {
	return (s.bitfield0 >> 0) & ((1 << 8) - 1)
}

// SetLib_ordinal updates the Lib_ordinal bitfield.
func (s *Dyld_chained_import_addend) SetLib_ordinal(v uint32) {
	const mask uint32 = (1 << 8) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 0)) | ((v & mask) << 0)
}

// Weak_import returns the Weak_import bitfield.
func (s *Dyld_chained_import_addend) Weak_import() uint32 {
	return (s.bitfield0 >> 8) & ((1 << 1) - 1)
}

// SetWeak_import updates the Weak_import bitfield.
func (s *Dyld_chained_import_addend) SetWeak_import(v uint32) {
	const mask uint32 = (1 << 1) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 8)) | ((v & mask) << 8)
}

// Name_offset returns the Name_offset bitfield.
func (s *Dyld_chained_import_addend) Name_offset() uint32 {
	return (s.bitfield0 >> 9) & ((1 << 23) - 1)
}

// SetName_offset updates the Name_offset bitfield.
func (s *Dyld_chained_import_addend) SetName_offset(v uint32) {
	const mask uint32 = (1 << 23) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 9)) | ((v & mask) << 9)
}

// Dyld_chained_import_addend64
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dyld_chained_import_addend64
type Dyld_chained_import_addend64 struct {
	bitfield0 uint64
	Addend    uint64
}

// Lib_ordinal returns the Lib_ordinal bitfield.
func (s *Dyld_chained_import_addend64) Lib_ordinal() uint64 {
	return (s.bitfield0 >> 0) & ((1 << 16) - 1)
}

// SetLib_ordinal updates the Lib_ordinal bitfield.
func (s *Dyld_chained_import_addend64) SetLib_ordinal(v uint64) {
	const mask uint64 = (1 << 16) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 0)) | ((v & mask) << 0)
}

// Weak_import returns the Weak_import bitfield.
func (s *Dyld_chained_import_addend64) Weak_import() uint64 {
	return (s.bitfield0 >> 16) & ((1 << 1) - 1)
}

// SetWeak_import updates the Weak_import bitfield.
func (s *Dyld_chained_import_addend64) SetWeak_import(v uint64) {
	const mask uint64 = (1 << 1) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 16)) | ((v & mask) << 16)
}

// Reserved returns the Reserved bitfield.
func (s *Dyld_chained_import_addend64) Reserved() uint64 {
	return (s.bitfield0 >> 17) & ((1 << 15) - 1)
}

// SetReserved updates the Reserved bitfield.
func (s *Dyld_chained_import_addend64) SetReserved(v uint64) {
	const mask uint64 = (1 << 15) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 17)) | ((v & mask) << 17)
}

// Name_offset returns the Name_offset bitfield.
func (s *Dyld_chained_import_addend64) Name_offset() uint64 {
	return (s.bitfield0 >> 32) & ((1 << 32) - 1)
}

// SetName_offset updates the Name_offset bitfield.
func (s *Dyld_chained_import_addend64) SetName_offset(v uint64) {
	const mask uint64 = (1 << 32) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 32)) | ((v & mask) << 32)
}

// Dyld_chained_ptr_32_bind
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dyld_chained_ptr_32_bind
type Dyld_chained_ptr_32_bind struct {
	bitfield0 uint32
}

// Ordinal returns the Ordinal bitfield.
func (s *Dyld_chained_ptr_32_bind) Ordinal() uint32 {
	return (s.bitfield0 >> 0) & ((1 << 20) - 1)
}

// SetOrdinal updates the Ordinal bitfield.
func (s *Dyld_chained_ptr_32_bind) SetOrdinal(v uint32) {
	const mask uint32 = (1 << 20) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 0)) | ((v & mask) << 0)
}

// Addend returns the Addend bitfield.
func (s *Dyld_chained_ptr_32_bind) Addend() uint32 {
	return (s.bitfield0 >> 20) & ((1 << 6) - 1)
}

// SetAddend updates the Addend bitfield.
func (s *Dyld_chained_ptr_32_bind) SetAddend(v uint32) {
	const mask uint32 = (1 << 6) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 20)) | ((v & mask) << 20)
}

// Next returns the Next bitfield.
func (s *Dyld_chained_ptr_32_bind) Next() uint32 {
	return (s.bitfield0 >> 26) & ((1 << 5) - 1)
}

// SetNext updates the Next bitfield.
func (s *Dyld_chained_ptr_32_bind) SetNext(v uint32) {
	const mask uint32 = (1 << 5) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 26)) | ((v & mask) << 26)
}

// Bind returns the Bind bitfield.
func (s *Dyld_chained_ptr_32_bind) Bind() uint32 {
	return (s.bitfield0 >> 31) & ((1 << 1) - 1)
}

// SetBind updates the Bind bitfield.
func (s *Dyld_chained_ptr_32_bind) SetBind(v uint32) {
	const mask uint32 = (1 << 1) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 31)) | ((v & mask) << 31)
}

// Dyld_chained_ptr_32_cache_rebase
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dyld_chained_ptr_32_cache_rebase
type Dyld_chained_ptr_32_cache_rebase struct {
	bitfield0 uint32
}

// Target returns the Target bitfield.
func (s *Dyld_chained_ptr_32_cache_rebase) Target() uint32 {
	return (s.bitfield0 >> 0) & ((1 << 30) - 1)
}

// SetTarget updates the Target bitfield.
func (s *Dyld_chained_ptr_32_cache_rebase) SetTarget(v uint32) {
	const mask uint32 = (1 << 30) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 0)) | ((v & mask) << 0)
}

// Next returns the Next bitfield.
func (s *Dyld_chained_ptr_32_cache_rebase) Next() uint32 {
	return (s.bitfield0 >> 30) & ((1 << 2) - 1)
}

// SetNext updates the Next bitfield.
func (s *Dyld_chained_ptr_32_cache_rebase) SetNext(v uint32) {
	const mask uint32 = (1 << 2) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 30)) | ((v & mask) << 30)
}

// Dyld_chained_ptr_32_firmware_rebase
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dyld_chained_ptr_32_firmware_rebase
type Dyld_chained_ptr_32_firmware_rebase struct {
	bitfield0 uint32
}

// Target returns the Target bitfield.
func (s *Dyld_chained_ptr_32_firmware_rebase) Target() uint32 {
	return (s.bitfield0 >> 0) & ((1 << 26) - 1)
}

// SetTarget updates the Target bitfield.
func (s *Dyld_chained_ptr_32_firmware_rebase) SetTarget(v uint32) {
	const mask uint32 = (1 << 26) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 0)) | ((v & mask) << 0)
}

// Next returns the Next bitfield.
func (s *Dyld_chained_ptr_32_firmware_rebase) Next() uint32 {
	return (s.bitfield0 >> 26) & ((1 << 6) - 1)
}

// SetNext updates the Next bitfield.
func (s *Dyld_chained_ptr_32_firmware_rebase) SetNext(v uint32) {
	const mask uint32 = (1 << 6) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 26)) | ((v & mask) << 26)
}

// Dyld_chained_ptr_32_rebase
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dyld_chained_ptr_32_rebase
type Dyld_chained_ptr_32_rebase struct {
	bitfield0 uint32
}

// Target returns the Target bitfield.
func (s *Dyld_chained_ptr_32_rebase) Target() uint32 {
	return (s.bitfield0 >> 0) & ((1 << 26) - 1)
}

// SetTarget updates the Target bitfield.
func (s *Dyld_chained_ptr_32_rebase) SetTarget(v uint32) {
	const mask uint32 = (1 << 26) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 0)) | ((v & mask) << 0)
}

// Next returns the Next bitfield.
func (s *Dyld_chained_ptr_32_rebase) Next() uint32 {
	return (s.bitfield0 >> 26) & ((1 << 5) - 1)
}

// SetNext updates the Next bitfield.
func (s *Dyld_chained_ptr_32_rebase) SetNext(v uint32) {
	const mask uint32 = (1 << 5) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 26)) | ((v & mask) << 26)
}

// Bind returns the Bind bitfield.
func (s *Dyld_chained_ptr_32_rebase) Bind() uint32 {
	return (s.bitfield0 >> 31) & ((1 << 1) - 1)
}

// SetBind updates the Bind bitfield.
func (s *Dyld_chained_ptr_32_rebase) SetBind(v uint32) {
	const mask uint32 = (1 << 1) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 31)) | ((v & mask) << 31)
}

// Dyld_chained_ptr_64_bind
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dyld_chained_ptr_64_bind
type Dyld_chained_ptr_64_bind struct {
	bitfield0 uint64
}

// Ordinal returns the Ordinal bitfield.
func (s *Dyld_chained_ptr_64_bind) Ordinal() uint64 {
	return (s.bitfield0 >> 0) & ((1 << 24) - 1)
}

// SetOrdinal updates the Ordinal bitfield.
func (s *Dyld_chained_ptr_64_bind) SetOrdinal(v uint64) {
	const mask uint64 = (1 << 24) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 0)) | ((v & mask) << 0)
}

// Addend returns the Addend bitfield.
func (s *Dyld_chained_ptr_64_bind) Addend() uint64 {
	return (s.bitfield0 >> 24) & ((1 << 8) - 1)
}

// SetAddend updates the Addend bitfield.
func (s *Dyld_chained_ptr_64_bind) SetAddend(v uint64) {
	const mask uint64 = (1 << 8) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 24)) | ((v & mask) << 24)
}

// Reserved returns the Reserved bitfield.
func (s *Dyld_chained_ptr_64_bind) Reserved() uint64 {
	return (s.bitfield0 >> 32) & ((1 << 19) - 1)
}

// SetReserved updates the Reserved bitfield.
func (s *Dyld_chained_ptr_64_bind) SetReserved(v uint64) {
	const mask uint64 = (1 << 19) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 32)) | ((v & mask) << 32)
}

// Next returns the Next bitfield.
func (s *Dyld_chained_ptr_64_bind) Next() uint64 {
	return (s.bitfield0 >> 51) & ((1 << 12) - 1)
}

// SetNext updates the Next bitfield.
func (s *Dyld_chained_ptr_64_bind) SetNext(v uint64) {
	const mask uint64 = (1 << 12) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 51)) | ((v & mask) << 51)
}

// Bind returns the Bind bitfield.
func (s *Dyld_chained_ptr_64_bind) Bind() uint64 {
	return (s.bitfield0 >> 63) & ((1 << 1) - 1)
}

// SetBind updates the Bind bitfield.
func (s *Dyld_chained_ptr_64_bind) SetBind(v uint64) {
	const mask uint64 = (1 << 1) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 63)) | ((v & mask) << 63)
}

// Dyld_chained_ptr_64_kernel_cache_rebase
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dyld_chained_ptr_64_kernel_cache_rebase
type Dyld_chained_ptr_64_kernel_cache_rebase struct {
	bitfield0 uint64
}

// Target returns the Target bitfield.
func (s *Dyld_chained_ptr_64_kernel_cache_rebase) Target() uint64 {
	return (s.bitfield0 >> 0) & ((1 << 30) - 1)
}

// SetTarget updates the Target bitfield.
func (s *Dyld_chained_ptr_64_kernel_cache_rebase) SetTarget(v uint64) {
	const mask uint64 = (1 << 30) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 0)) | ((v & mask) << 0)
}

// CacheLevel returns the CacheLevel bitfield.
func (s *Dyld_chained_ptr_64_kernel_cache_rebase) CacheLevel() uint64 {
	return (s.bitfield0 >> 30) & ((1 << 2) - 1)
}

// SetCacheLevel updates the CacheLevel bitfield.
func (s *Dyld_chained_ptr_64_kernel_cache_rebase) SetCacheLevel(v uint64) {
	const mask uint64 = (1 << 2) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 30)) | ((v & mask) << 30)
}

// Diversity returns the Diversity bitfield.
func (s *Dyld_chained_ptr_64_kernel_cache_rebase) Diversity() uint64 {
	return (s.bitfield0 >> 32) & ((1 << 16) - 1)
}

// SetDiversity updates the Diversity bitfield.
func (s *Dyld_chained_ptr_64_kernel_cache_rebase) SetDiversity(v uint64) {
	const mask uint64 = (1 << 16) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 32)) | ((v & mask) << 32)
}

// AddrDiv returns the AddrDiv bitfield.
func (s *Dyld_chained_ptr_64_kernel_cache_rebase) AddrDiv() uint64 {
	return (s.bitfield0 >> 48) & ((1 << 1) - 1)
}

// SetAddrDiv updates the AddrDiv bitfield.
func (s *Dyld_chained_ptr_64_kernel_cache_rebase) SetAddrDiv(v uint64) {
	const mask uint64 = (1 << 1) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 48)) | ((v & mask) << 48)
}

// Key returns the Key bitfield.
func (s *Dyld_chained_ptr_64_kernel_cache_rebase) Key() uint64 {
	return (s.bitfield0 >> 49) & ((1 << 2) - 1)
}

// SetKey updates the Key bitfield.
func (s *Dyld_chained_ptr_64_kernel_cache_rebase) SetKey(v uint64) {
	const mask uint64 = (1 << 2) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 49)) | ((v & mask) << 49)
}

// Next returns the Next bitfield.
func (s *Dyld_chained_ptr_64_kernel_cache_rebase) Next() uint64 {
	return (s.bitfield0 >> 51) & ((1 << 12) - 1)
}

// SetNext updates the Next bitfield.
func (s *Dyld_chained_ptr_64_kernel_cache_rebase) SetNext(v uint64) {
	const mask uint64 = (1 << 12) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 51)) | ((v & mask) << 51)
}

// IsAuth returns the IsAuth bitfield.
func (s *Dyld_chained_ptr_64_kernel_cache_rebase) IsAuth() uint64 {
	return (s.bitfield0 >> 63) & ((1 << 1) - 1)
}

// SetIsAuth updates the IsAuth bitfield.
func (s *Dyld_chained_ptr_64_kernel_cache_rebase) SetIsAuth(v uint64) {
	const mask uint64 = (1 << 1) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 63)) | ((v & mask) << 63)
}

// Dyld_chained_ptr_64_rebase
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dyld_chained_ptr_64_rebase
type Dyld_chained_ptr_64_rebase struct {
	bitfield0 uint64
}

// Target returns the Target bitfield.
func (s *Dyld_chained_ptr_64_rebase) Target() uint64 {
	return (s.bitfield0 >> 0) & ((1 << 36) - 1)
}

// SetTarget updates the Target bitfield.
func (s *Dyld_chained_ptr_64_rebase) SetTarget(v uint64) {
	const mask uint64 = (1 << 36) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 0)) | ((v & mask) << 0)
}

// High8 returns the High8 bitfield.
func (s *Dyld_chained_ptr_64_rebase) High8() uint64 {
	return (s.bitfield0 >> 36) & ((1 << 8) - 1)
}

// SetHigh8 updates the High8 bitfield.
func (s *Dyld_chained_ptr_64_rebase) SetHigh8(v uint64) {
	const mask uint64 = (1 << 8) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 36)) | ((v & mask) << 36)
}

// Reserved returns the Reserved bitfield.
func (s *Dyld_chained_ptr_64_rebase) Reserved() uint64 {
	return (s.bitfield0 >> 44) & ((1 << 7) - 1)
}

// SetReserved updates the Reserved bitfield.
func (s *Dyld_chained_ptr_64_rebase) SetReserved(v uint64) {
	const mask uint64 = (1 << 7) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 44)) | ((v & mask) << 44)
}

// Next returns the Next bitfield.
func (s *Dyld_chained_ptr_64_rebase) Next() uint64 {
	return (s.bitfield0 >> 51) & ((1 << 12) - 1)
}

// SetNext updates the Next bitfield.
func (s *Dyld_chained_ptr_64_rebase) SetNext(v uint64) {
	const mask uint64 = (1 << 12) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 51)) | ((v & mask) << 51)
}

// Bind returns the Bind bitfield.
func (s *Dyld_chained_ptr_64_rebase) Bind() uint64 {
	return (s.bitfield0 >> 63) & ((1 << 1) - 1)
}

// SetBind updates the Bind bitfield.
func (s *Dyld_chained_ptr_64_rebase) SetBind(v uint64) {
	const mask uint64 = (1 << 1) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 63)) | ((v & mask) << 63)
}

// Dyld_chained_ptr_arm64e_auth_bind
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dyld_chained_ptr_arm64e_auth_bind
type Dyld_chained_ptr_arm64e_auth_bind struct {
	bitfield0 uint64
}

// Ordinal returns the Ordinal bitfield.
func (s *Dyld_chained_ptr_arm64e_auth_bind) Ordinal() uint64 {
	return (s.bitfield0 >> 0) & ((1 << 16) - 1)
}

// SetOrdinal updates the Ordinal bitfield.
func (s *Dyld_chained_ptr_arm64e_auth_bind) SetOrdinal(v uint64) {
	const mask uint64 = (1 << 16) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 0)) | ((v & mask) << 0)
}

// Zero returns the Zero bitfield.
func (s *Dyld_chained_ptr_arm64e_auth_bind) Zero() uint64 {
	return (s.bitfield0 >> 16) & ((1 << 16) - 1)
}

// SetZero updates the Zero bitfield.
func (s *Dyld_chained_ptr_arm64e_auth_bind) SetZero(v uint64) {
	const mask uint64 = (1 << 16) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 16)) | ((v & mask) << 16)
}

// Diversity returns the Diversity bitfield.
func (s *Dyld_chained_ptr_arm64e_auth_bind) Diversity() uint64 {
	return (s.bitfield0 >> 32) & ((1 << 16) - 1)
}

// SetDiversity updates the Diversity bitfield.
func (s *Dyld_chained_ptr_arm64e_auth_bind) SetDiversity(v uint64) {
	const mask uint64 = (1 << 16) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 32)) | ((v & mask) << 32)
}

// AddrDiv returns the AddrDiv bitfield.
func (s *Dyld_chained_ptr_arm64e_auth_bind) AddrDiv() uint64 {
	return (s.bitfield0 >> 48) & ((1 << 1) - 1)
}

// SetAddrDiv updates the AddrDiv bitfield.
func (s *Dyld_chained_ptr_arm64e_auth_bind) SetAddrDiv(v uint64) {
	const mask uint64 = (1 << 1) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 48)) | ((v & mask) << 48)
}

// Key returns the Key bitfield.
func (s *Dyld_chained_ptr_arm64e_auth_bind) Key() uint64 {
	return (s.bitfield0 >> 49) & ((1 << 2) - 1)
}

// SetKey updates the Key bitfield.
func (s *Dyld_chained_ptr_arm64e_auth_bind) SetKey(v uint64) {
	const mask uint64 = (1 << 2) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 49)) | ((v & mask) << 49)
}

// Next returns the Next bitfield.
func (s *Dyld_chained_ptr_arm64e_auth_bind) Next() uint64 {
	return (s.bitfield0 >> 51) & ((1 << 11) - 1)
}

// SetNext updates the Next bitfield.
func (s *Dyld_chained_ptr_arm64e_auth_bind) SetNext(v uint64) {
	const mask uint64 = (1 << 11) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 51)) | ((v & mask) << 51)
}

// Bind returns the Bind bitfield.
func (s *Dyld_chained_ptr_arm64e_auth_bind) Bind() uint64 {
	return (s.bitfield0 >> 62) & ((1 << 1) - 1)
}

// SetBind updates the Bind bitfield.
func (s *Dyld_chained_ptr_arm64e_auth_bind) SetBind(v uint64) {
	const mask uint64 = (1 << 1) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 62)) | ((v & mask) << 62)
}

// Auth returns the Auth bitfield.
func (s *Dyld_chained_ptr_arm64e_auth_bind) Auth() uint64 {
	return (s.bitfield0 >> 63) & ((1 << 1) - 1)
}

// SetAuth updates the Auth bitfield.
func (s *Dyld_chained_ptr_arm64e_auth_bind) SetAuth(v uint64) {
	const mask uint64 = (1 << 1) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 63)) | ((v & mask) << 63)
}

// Dyld_chained_ptr_arm64e_auth_bind24
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dyld_chained_ptr_arm64e_auth_bind24
type Dyld_chained_ptr_arm64e_auth_bind24 struct {
	bitfield0 uint64
}

// Ordinal returns the Ordinal bitfield.
func (s *Dyld_chained_ptr_arm64e_auth_bind24) Ordinal() uint64 {
	return (s.bitfield0 >> 0) & ((1 << 24) - 1)
}

// SetOrdinal updates the Ordinal bitfield.
func (s *Dyld_chained_ptr_arm64e_auth_bind24) SetOrdinal(v uint64) {
	const mask uint64 = (1 << 24) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 0)) | ((v & mask) << 0)
}

// Zero returns the Zero bitfield.
func (s *Dyld_chained_ptr_arm64e_auth_bind24) Zero() uint64 {
	return (s.bitfield0 >> 24) & ((1 << 8) - 1)
}

// SetZero updates the Zero bitfield.
func (s *Dyld_chained_ptr_arm64e_auth_bind24) SetZero(v uint64) {
	const mask uint64 = (1 << 8) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 24)) | ((v & mask) << 24)
}

// Diversity returns the Diversity bitfield.
func (s *Dyld_chained_ptr_arm64e_auth_bind24) Diversity() uint64 {
	return (s.bitfield0 >> 32) & ((1 << 16) - 1)
}

// SetDiversity updates the Diversity bitfield.
func (s *Dyld_chained_ptr_arm64e_auth_bind24) SetDiversity(v uint64) {
	const mask uint64 = (1 << 16) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 32)) | ((v & mask) << 32)
}

// AddrDiv returns the AddrDiv bitfield.
func (s *Dyld_chained_ptr_arm64e_auth_bind24) AddrDiv() uint64 {
	return (s.bitfield0 >> 48) & ((1 << 1) - 1)
}

// SetAddrDiv updates the AddrDiv bitfield.
func (s *Dyld_chained_ptr_arm64e_auth_bind24) SetAddrDiv(v uint64) {
	const mask uint64 = (1 << 1) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 48)) | ((v & mask) << 48)
}

// Key returns the Key bitfield.
func (s *Dyld_chained_ptr_arm64e_auth_bind24) Key() uint64 {
	return (s.bitfield0 >> 49) & ((1 << 2) - 1)
}

// SetKey updates the Key bitfield.
func (s *Dyld_chained_ptr_arm64e_auth_bind24) SetKey(v uint64) {
	const mask uint64 = (1 << 2) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 49)) | ((v & mask) << 49)
}

// Next returns the Next bitfield.
func (s *Dyld_chained_ptr_arm64e_auth_bind24) Next() uint64 {
	return (s.bitfield0 >> 51) & ((1 << 11) - 1)
}

// SetNext updates the Next bitfield.
func (s *Dyld_chained_ptr_arm64e_auth_bind24) SetNext(v uint64) {
	const mask uint64 = (1 << 11) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 51)) | ((v & mask) << 51)
}

// Bind returns the Bind bitfield.
func (s *Dyld_chained_ptr_arm64e_auth_bind24) Bind() uint64 {
	return (s.bitfield0 >> 62) & ((1 << 1) - 1)
}

// SetBind updates the Bind bitfield.
func (s *Dyld_chained_ptr_arm64e_auth_bind24) SetBind(v uint64) {
	const mask uint64 = (1 << 1) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 62)) | ((v & mask) << 62)
}

// Auth returns the Auth bitfield.
func (s *Dyld_chained_ptr_arm64e_auth_bind24) Auth() uint64 {
	return (s.bitfield0 >> 63) & ((1 << 1) - 1)
}

// SetAuth updates the Auth bitfield.
func (s *Dyld_chained_ptr_arm64e_auth_bind24) SetAuth(v uint64) {
	const mask uint64 = (1 << 1) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 63)) | ((v & mask) << 63)
}

// Dyld_chained_ptr_arm64e_auth_rebase
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dyld_chained_ptr_arm64e_auth_rebase
type Dyld_chained_ptr_arm64e_auth_rebase struct {
	bitfield0 uint64
}

// Target returns the Target bitfield.
func (s *Dyld_chained_ptr_arm64e_auth_rebase) Target() uint64 {
	return (s.bitfield0 >> 0) & ((1 << 32) - 1)
}

// SetTarget updates the Target bitfield.
func (s *Dyld_chained_ptr_arm64e_auth_rebase) SetTarget(v uint64) {
	const mask uint64 = (1 << 32) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 0)) | ((v & mask) << 0)
}

// Diversity returns the Diversity bitfield.
func (s *Dyld_chained_ptr_arm64e_auth_rebase) Diversity() uint64 {
	return (s.bitfield0 >> 32) & ((1 << 16) - 1)
}

// SetDiversity updates the Diversity bitfield.
func (s *Dyld_chained_ptr_arm64e_auth_rebase) SetDiversity(v uint64) {
	const mask uint64 = (1 << 16) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 32)) | ((v & mask) << 32)
}

// AddrDiv returns the AddrDiv bitfield.
func (s *Dyld_chained_ptr_arm64e_auth_rebase) AddrDiv() uint64 {
	return (s.bitfield0 >> 48) & ((1 << 1) - 1)
}

// SetAddrDiv updates the AddrDiv bitfield.
func (s *Dyld_chained_ptr_arm64e_auth_rebase) SetAddrDiv(v uint64) {
	const mask uint64 = (1 << 1) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 48)) | ((v & mask) << 48)
}

// Key returns the Key bitfield.
func (s *Dyld_chained_ptr_arm64e_auth_rebase) Key() uint64 {
	return (s.bitfield0 >> 49) & ((1 << 2) - 1)
}

// SetKey updates the Key bitfield.
func (s *Dyld_chained_ptr_arm64e_auth_rebase) SetKey(v uint64) {
	const mask uint64 = (1 << 2) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 49)) | ((v & mask) << 49)
}

// Next returns the Next bitfield.
func (s *Dyld_chained_ptr_arm64e_auth_rebase) Next() uint64 {
	return (s.bitfield0 >> 51) & ((1 << 11) - 1)
}

// SetNext updates the Next bitfield.
func (s *Dyld_chained_ptr_arm64e_auth_rebase) SetNext(v uint64) {
	const mask uint64 = (1 << 11) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 51)) | ((v & mask) << 51)
}

// Bind returns the Bind bitfield.
func (s *Dyld_chained_ptr_arm64e_auth_rebase) Bind() uint64 {
	return (s.bitfield0 >> 62) & ((1 << 1) - 1)
}

// SetBind updates the Bind bitfield.
func (s *Dyld_chained_ptr_arm64e_auth_rebase) SetBind(v uint64) {
	const mask uint64 = (1 << 1) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 62)) | ((v & mask) << 62)
}

// Auth returns the Auth bitfield.
func (s *Dyld_chained_ptr_arm64e_auth_rebase) Auth() uint64 {
	return (s.bitfield0 >> 63) & ((1 << 1) - 1)
}

// SetAuth updates the Auth bitfield.
func (s *Dyld_chained_ptr_arm64e_auth_rebase) SetAuth(v uint64) {
	const mask uint64 = (1 << 1) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 63)) | ((v & mask) << 63)
}

// Dyld_chained_ptr_arm64e_bind
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dyld_chained_ptr_arm64e_bind
type Dyld_chained_ptr_arm64e_bind struct {
	bitfield0 uint64
}

// Ordinal returns the Ordinal bitfield.
func (s *Dyld_chained_ptr_arm64e_bind) Ordinal() uint64 {
	return (s.bitfield0 >> 0) & ((1 << 16) - 1)
}

// SetOrdinal updates the Ordinal bitfield.
func (s *Dyld_chained_ptr_arm64e_bind) SetOrdinal(v uint64) {
	const mask uint64 = (1 << 16) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 0)) | ((v & mask) << 0)
}

// Zero returns the Zero bitfield.
func (s *Dyld_chained_ptr_arm64e_bind) Zero() uint64 {
	return (s.bitfield0 >> 16) & ((1 << 16) - 1)
}

// SetZero updates the Zero bitfield.
func (s *Dyld_chained_ptr_arm64e_bind) SetZero(v uint64) {
	const mask uint64 = (1 << 16) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 16)) | ((v & mask) << 16)
}

// Addend returns the Addend bitfield.
func (s *Dyld_chained_ptr_arm64e_bind) Addend() uint64 {
	return (s.bitfield0 >> 32) & ((1 << 19) - 1)
}

// SetAddend updates the Addend bitfield.
func (s *Dyld_chained_ptr_arm64e_bind) SetAddend(v uint64) {
	const mask uint64 = (1 << 19) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 32)) | ((v & mask) << 32)
}

// Next returns the Next bitfield.
func (s *Dyld_chained_ptr_arm64e_bind) Next() uint64 {
	return (s.bitfield0 >> 51) & ((1 << 11) - 1)
}

// SetNext updates the Next bitfield.
func (s *Dyld_chained_ptr_arm64e_bind) SetNext(v uint64) {
	const mask uint64 = (1 << 11) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 51)) | ((v & mask) << 51)
}

// Bind returns the Bind bitfield.
func (s *Dyld_chained_ptr_arm64e_bind) Bind() uint64 {
	return (s.bitfield0 >> 62) & ((1 << 1) - 1)
}

// SetBind updates the Bind bitfield.
func (s *Dyld_chained_ptr_arm64e_bind) SetBind(v uint64) {
	const mask uint64 = (1 << 1) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 62)) | ((v & mask) << 62)
}

// Auth returns the Auth bitfield.
func (s *Dyld_chained_ptr_arm64e_bind) Auth() uint64 {
	return (s.bitfield0 >> 63) & ((1 << 1) - 1)
}

// SetAuth updates the Auth bitfield.
func (s *Dyld_chained_ptr_arm64e_bind) SetAuth(v uint64) {
	const mask uint64 = (1 << 1) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 63)) | ((v & mask) << 63)
}

// Dyld_chained_ptr_arm64e_bind24
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dyld_chained_ptr_arm64e_bind24
type Dyld_chained_ptr_arm64e_bind24 struct {
	bitfield0 uint64
}

// Ordinal returns the Ordinal bitfield.
func (s *Dyld_chained_ptr_arm64e_bind24) Ordinal() uint64 {
	return (s.bitfield0 >> 0) & ((1 << 24) - 1)
}

// SetOrdinal updates the Ordinal bitfield.
func (s *Dyld_chained_ptr_arm64e_bind24) SetOrdinal(v uint64) {
	const mask uint64 = (1 << 24) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 0)) | ((v & mask) << 0)
}

// Zero returns the Zero bitfield.
func (s *Dyld_chained_ptr_arm64e_bind24) Zero() uint64 {
	return (s.bitfield0 >> 24) & ((1 << 8) - 1)
}

// SetZero updates the Zero bitfield.
func (s *Dyld_chained_ptr_arm64e_bind24) SetZero(v uint64) {
	const mask uint64 = (1 << 8) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 24)) | ((v & mask) << 24)
}

// Addend returns the Addend bitfield.
func (s *Dyld_chained_ptr_arm64e_bind24) Addend() uint64 {
	return (s.bitfield0 >> 32) & ((1 << 19) - 1)
}

// SetAddend updates the Addend bitfield.
func (s *Dyld_chained_ptr_arm64e_bind24) SetAddend(v uint64) {
	const mask uint64 = (1 << 19) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 32)) | ((v & mask) << 32)
}

// Next returns the Next bitfield.
func (s *Dyld_chained_ptr_arm64e_bind24) Next() uint64 {
	return (s.bitfield0 >> 51) & ((1 << 11) - 1)
}

// SetNext updates the Next bitfield.
func (s *Dyld_chained_ptr_arm64e_bind24) SetNext(v uint64) {
	const mask uint64 = (1 << 11) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 51)) | ((v & mask) << 51)
}

// Bind returns the Bind bitfield.
func (s *Dyld_chained_ptr_arm64e_bind24) Bind() uint64 {
	return (s.bitfield0 >> 62) & ((1 << 1) - 1)
}

// SetBind updates the Bind bitfield.
func (s *Dyld_chained_ptr_arm64e_bind24) SetBind(v uint64) {
	const mask uint64 = (1 << 1) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 62)) | ((v & mask) << 62)
}

// Auth returns the Auth bitfield.
func (s *Dyld_chained_ptr_arm64e_bind24) Auth() uint64 {
	return (s.bitfield0 >> 63) & ((1 << 1) - 1)
}

// SetAuth updates the Auth bitfield.
func (s *Dyld_chained_ptr_arm64e_bind24) SetAuth(v uint64) {
	const mask uint64 = (1 << 1) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 63)) | ((v & mask) << 63)
}

// Dyld_chained_ptr_arm64e_rebase
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dyld_chained_ptr_arm64e_rebase
type Dyld_chained_ptr_arm64e_rebase struct {
	bitfield0 uint64
}

// Target returns the Target bitfield.
func (s *Dyld_chained_ptr_arm64e_rebase) Target() uint64 {
	return (s.bitfield0 >> 0) & ((1 << 43) - 1)
}

// SetTarget updates the Target bitfield.
func (s *Dyld_chained_ptr_arm64e_rebase) SetTarget(v uint64) {
	const mask uint64 = (1 << 43) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 0)) | ((v & mask) << 0)
}

// High8 returns the High8 bitfield.
func (s *Dyld_chained_ptr_arm64e_rebase) High8() uint64 {
	return (s.bitfield0 >> 43) & ((1 << 8) - 1)
}

// SetHigh8 updates the High8 bitfield.
func (s *Dyld_chained_ptr_arm64e_rebase) SetHigh8(v uint64) {
	const mask uint64 = (1 << 8) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 43)) | ((v & mask) << 43)
}

// Next returns the Next bitfield.
func (s *Dyld_chained_ptr_arm64e_rebase) Next() uint64 {
	return (s.bitfield0 >> 51) & ((1 << 11) - 1)
}

// SetNext updates the Next bitfield.
func (s *Dyld_chained_ptr_arm64e_rebase) SetNext(v uint64) {
	const mask uint64 = (1 << 11) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 51)) | ((v & mask) << 51)
}

// Bind returns the Bind bitfield.
func (s *Dyld_chained_ptr_arm64e_rebase) Bind() uint64 {
	return (s.bitfield0 >> 62) & ((1 << 1) - 1)
}

// SetBind updates the Bind bitfield.
func (s *Dyld_chained_ptr_arm64e_rebase) SetBind(v uint64) {
	const mask uint64 = (1 << 1) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 62)) | ((v & mask) << 62)
}

// Auth returns the Auth bitfield.
func (s *Dyld_chained_ptr_arm64e_rebase) Auth() uint64 {
	return (s.bitfield0 >> 63) & ((1 << 1) - 1)
}

// SetAuth updates the Auth bitfield.
func (s *Dyld_chained_ptr_arm64e_rebase) SetAuth(v uint64) {
	const mask uint64 = (1 << 1) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 63)) | ((v & mask) << 63)
}

// Dyld_chained_ptr_arm64e_shared_cache_auth_rebase
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dyld_chained_ptr_arm64e_shared_cache_auth_rebase
type Dyld_chained_ptr_arm64e_shared_cache_auth_rebase struct {
	bitfield0 uint64
}

// RuntimeOffset returns the RuntimeOffset bitfield.
func (s *Dyld_chained_ptr_arm64e_shared_cache_auth_rebase) RuntimeOffset() uint64 {
	return (s.bitfield0 >> 0) & ((1 << 34) - 1)
}

// SetRuntimeOffset updates the RuntimeOffset bitfield.
func (s *Dyld_chained_ptr_arm64e_shared_cache_auth_rebase) SetRuntimeOffset(v uint64) {
	const mask uint64 = (1 << 34) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 0)) | ((v & mask) << 0)
}

// Diversity returns the Diversity bitfield.
func (s *Dyld_chained_ptr_arm64e_shared_cache_auth_rebase) Diversity() uint64 {
	return (s.bitfield0 >> 34) & ((1 << 16) - 1)
}

// SetDiversity updates the Diversity bitfield.
func (s *Dyld_chained_ptr_arm64e_shared_cache_auth_rebase) SetDiversity(v uint64) {
	const mask uint64 = (1 << 16) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 34)) | ((v & mask) << 34)
}

// AddrDiv returns the AddrDiv bitfield.
func (s *Dyld_chained_ptr_arm64e_shared_cache_auth_rebase) AddrDiv() uint64 {
	return (s.bitfield0 >> 50) & ((1 << 1) - 1)
}

// SetAddrDiv updates the AddrDiv bitfield.
func (s *Dyld_chained_ptr_arm64e_shared_cache_auth_rebase) SetAddrDiv(v uint64) {
	const mask uint64 = (1 << 1) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 50)) | ((v & mask) << 50)
}

// KeyIsData returns the KeyIsData bitfield.
func (s *Dyld_chained_ptr_arm64e_shared_cache_auth_rebase) KeyIsData() uint64 {
	return (s.bitfield0 >> 51) & ((1 << 1) - 1)
}

// SetKeyIsData updates the KeyIsData bitfield.
func (s *Dyld_chained_ptr_arm64e_shared_cache_auth_rebase) SetKeyIsData(v uint64) {
	const mask uint64 = (1 << 1) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 51)) | ((v & mask) << 51)
}

// Next returns the Next bitfield.
func (s *Dyld_chained_ptr_arm64e_shared_cache_auth_rebase) Next() uint64 {
	return (s.bitfield0 >> 52) & ((1 << 11) - 1)
}

// SetNext updates the Next bitfield.
func (s *Dyld_chained_ptr_arm64e_shared_cache_auth_rebase) SetNext(v uint64) {
	const mask uint64 = (1 << 11) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 52)) | ((v & mask) << 52)
}

// Auth returns the Auth bitfield.
func (s *Dyld_chained_ptr_arm64e_shared_cache_auth_rebase) Auth() uint64 {
	return (s.bitfield0 >> 63) & ((1 << 1) - 1)
}

// SetAuth updates the Auth bitfield.
func (s *Dyld_chained_ptr_arm64e_shared_cache_auth_rebase) SetAuth(v uint64) {
	const mask uint64 = (1 << 1) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 63)) | ((v & mask) << 63)
}

// Dyld_chained_ptr_arm64e_shared_cache_rebase
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dyld_chained_ptr_arm64e_shared_cache_rebase
type Dyld_chained_ptr_arm64e_shared_cache_rebase struct {
	bitfield0 uint64
}

// RuntimeOffset returns the RuntimeOffset bitfield.
func (s *Dyld_chained_ptr_arm64e_shared_cache_rebase) RuntimeOffset() uint64 {
	return (s.bitfield0 >> 0) & ((1 << 34) - 1)
}

// SetRuntimeOffset updates the RuntimeOffset bitfield.
func (s *Dyld_chained_ptr_arm64e_shared_cache_rebase) SetRuntimeOffset(v uint64) {
	const mask uint64 = (1 << 34) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 0)) | ((v & mask) << 0)
}

// High8 returns the High8 bitfield.
func (s *Dyld_chained_ptr_arm64e_shared_cache_rebase) High8() uint64 {
	return (s.bitfield0 >> 34) & ((1 << 8) - 1)
}

// SetHigh8 updates the High8 bitfield.
func (s *Dyld_chained_ptr_arm64e_shared_cache_rebase) SetHigh8(v uint64) {
	const mask uint64 = (1 << 8) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 34)) | ((v & mask) << 34)
}

// Unused returns the Unused bitfield.
func (s *Dyld_chained_ptr_arm64e_shared_cache_rebase) Unused() uint64 {
	return (s.bitfield0 >> 42) & ((1 << 10) - 1)
}

// SetUnused updates the Unused bitfield.
func (s *Dyld_chained_ptr_arm64e_shared_cache_rebase) SetUnused(v uint64) {
	const mask uint64 = (1 << 10) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 42)) | ((v & mask) << 42)
}

// Next returns the Next bitfield.
func (s *Dyld_chained_ptr_arm64e_shared_cache_rebase) Next() uint64 {
	return (s.bitfield0 >> 52) & ((1 << 11) - 1)
}

// SetNext updates the Next bitfield.
func (s *Dyld_chained_ptr_arm64e_shared_cache_rebase) SetNext(v uint64) {
	const mask uint64 = (1 << 11) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 52)) | ((v & mask) << 52)
}

// Auth returns the Auth bitfield.
func (s *Dyld_chained_ptr_arm64e_shared_cache_rebase) Auth() uint64 {
	return (s.bitfield0 >> 63) & ((1 << 1) - 1)
}

// SetAuth updates the Auth bitfield.
func (s *Dyld_chained_ptr_arm64e_shared_cache_rebase) SetAuth(v uint64) {
	const mask uint64 = (1 << 1) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 63)) | ((v & mask) << 63)
}

// Dyld_chained_starts_in_image
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dyld_chained_starts_in_image
type Dyld_chained_starts_in_image struct {
	Seg_count       uint32
	Seg_info_offset [1]uint32
}

// Dyld_chained_starts_in_segment
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dyld_chained_starts_in_segment
type Dyld_chained_starts_in_segment struct {
	Size              uint32
	Page_size         uint16
	Pointer_format    uint16
	Segment_offset    uint64
	Max_valid_pointer uint32
	Page_count        uint16
	Page_start        [1]uint16
}

// Dyld_chained_starts_offsets
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dyld_chained_starts_offsets
type Dyld_chained_starts_offsets struct {
	Pointer_format uint32
	Starts_count   uint32
	Chain_starts   [1]uint32
}

// Dyld_info_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dyld_info_command
type Dyld_info_command struct {
	Cmd            uint32
	Cmdsize        uint32
	Rebase_off     uint32
	Rebase_size    uint32
	Bind_off       uint32
	Bind_size      uint32
	Weak_bind_off  uint32
	Weak_bind_size uint32
	Lazy_bind_off  uint32
	Lazy_bind_size uint32
	Export_off     uint32
	Export_size    uint32
}

// Dyld_shared_cache_loadinfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dyld_shared_cache_loadinfo
type Dyld_shared_cache_loadinfo struct {
	SharedCacheSlide                     uint64
	SharedCacheUUID                      [16]uint8
	SharedCacheUnreliableSlidBaseAddress uint64
	SharedCacheSlidFirstMapping          uint64
}

// Dyld_shared_cache_loadinfo_v2
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dyld_shared_cache_loadinfo_v2
type Dyld_shared_cache_loadinfo_v2 struct {
	SharedCacheSlide                     uint64
	SharedCacheUUID                      [16]uint8
	SharedCacheUnreliableSlidBaseAddress uint64
	SharedCacheSlidFirstMapping          uint64
	SharedCacheID                        uint32
	SharedCacheFlags                     uint32
}

// Dyld_uuid_info_32
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dyld_uuid_info_32
type Dyld_uuid_info_32 struct {
	ImageLoadAddress uint32
	ImageUUID        [16]uint8
}

// Dyld_uuid_info_64
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dyld_uuid_info_64
type Dyld_uuid_info_64 struct {
	ImageLoadAddress uint64
	ImageUUID        [16]uint8
}

// Dyld_uuid_info_64_v2
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dyld_uuid_info_64_v2
type Dyld_uuid_info_64_v2 struct {
	ImageLoadAddress     uint64
	ImageUUID            [16]uint8
	ImageSlidBaseAddress uint64
}

// Dylib
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dylib
type Dylib struct {
	Name                  [1]uint32
	Timestamp             uint32
	Current_version       uint32
	Compatibility_version uint32
}

// Dylib_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dylib_command
type Dylib_command struct {
	Cmd     uint32 // Common to all load command structures. For this structure, set to either `LC_LOAD_DYLIB`, `LC_LOAD_WEAK_DYLIB`, or `LC_ID_DYLIB`.
	Cmdsize uint32
	Dylib   Dylib
}

// Dylib_module
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dylib_module
type Dylib_module struct {
	Module_name           uint32
	Iextdefsym            uint32
	Nextdefsym            uint32
	Irefsym               uint32
	Nrefsym               uint32 // The number of external reference entries provided by this module.
	Ilocalsym             uint32
	Nlocalsym             uint32
	Iextrel               uint32 // The index into the external relocation table of the first entry provided by this module.
	Nextrel               uint32
	Iinit_iterm           uint32
	Ninit_nterm           uint32 // Contains both the number of pointers in the module initialization (the low 16 bits) and the number of pointers in the module termination section (the high 16 bits) for this module.
	Objc_module_info_addr uint32
	Objc_module_info_size uint32
}

// Dylib_module_64
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dylib_module_64
type Dylib_module_64 struct {
	Module_name           uint32
	Iextdefsym            uint32
	Nextdefsym            uint32
	Irefsym               uint32
	Nrefsym               uint32
	Ilocalsym             uint32
	Nlocalsym             uint32
	Iextrel               uint32
	Nextrel               uint32
	Iinit_iterm           uint32
	Ninit_nterm           uint32 // Contains both the number of pointers in the module initialization (the low 16 bits) and the number of pointers in the module termination section (the high 16 bits) for this module.
	Objc_module_info_size uint32
	Objc_module_info_addr uint64
}

// Dylib_reference
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dylib_reference
type Dylib_reference struct {
	bitfield0 uint32
}

// Isym returns the Isym bitfield.
func (s *Dylib_reference) Isym() uint32 {
	return (s.bitfield0 >> 0) & ((1 << 24) - 1)
}

// SetIsym updates the Isym bitfield.
func (s *Dylib_reference) SetIsym(v uint32) {
	const mask uint32 = (1 << 24) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 0)) | ((v & mask) << 0)
}

// Flags returns the Flags bitfield.
func (s *Dylib_reference) Flags() uint32 {
	return (s.bitfield0 >> 24) & ((1 << 8) - 1)
}

// SetFlags updates the Flags bitfield.
func (s *Dylib_reference) SetFlags(v uint32) {
	const mask uint32 = (1 << 8) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 24)) | ((v & mask) << 24)
}

// Dylib_table_of_contents
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dylib_table_of_contents
type Dylib_table_of_contents struct {
	Symbol_index uint32 // An index into the symbol table indicating the defined external symbol to which this entry refers.
	Module_index uint32
}

// Dylinker_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dylinker_command
type Dylinker_command struct {
	Cmd     uint32
	Cmdsize uint32
	Name    [1]uint32
}

// Dysymtab_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dysymtab_command
type Dysymtab_command struct {
	Cmd            uint32 // Common to all load command structures. For this structure, set to `LC_DYSYMTAB`.
	Cmdsize        uint32
	Ilocalsym      uint32
	Nlocalsym      uint32
	Iextdefsym     uint32
	Nextdefsym     uint32
	Iundefsym      uint32
	Nundefsym      uint32
	Tocoff         uint32
	Ntoc           uint32
	Modtaboff      uint32
	Nmodtab        uint32
	Extrefsymoff   uint32 // An integer indicating the byte offset from the start of the file to the external reference table data.
	Nextrefsyms    uint32
	Indirectsymoff uint32
	Nindirectsyms  uint32
	Extreloff      uint32 // An integer indicating the byte offset from the start of the file to the external relocation table data.
	Nextrel        uint32
	Locreloff      uint32
	Nlocrel        uint32
}

// Ecc_event
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ecc_event
type Ecc_event struct {
	Id    uint8
	Count uint8
	Data  [8]uint64
}

// Efi_aurr_extended_panic_log
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/efi_aurr_extended_panic_log
type Efi_aurr_extended_panic_log struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [1310464]byte
}

// Efi_aurr_extended_log_buf returns the Efi_aurr_extended_log_buf field from the record's packed storage.
func (s *Efi_aurr_extended_panic_log) Efi_aurr_extended_log_buf() [1310456]int8 {
	return *(*[1310456]int8)(unsafe.Pointer(&s.storage[0]))
}

// SetEfi_aurr_extended_log_buf updates the Efi_aurr_extended_log_buf field in the record's packed storage.
func (s *Efi_aurr_extended_panic_log) SetEfi_aurr_extended_log_buf(v [1310456]int8) {
	*(*[1310456]int8)(unsafe.Pointer(&s.storage[0])) = v
}

// Efi_aurr_log_tail returns the Efi_aurr_log_tail field from the record's packed storage.
func (s *Efi_aurr_extended_panic_log) Efi_aurr_log_tail() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[1310456:1310460]))
}

// SetEfi_aurr_log_tail updates the Efi_aurr_log_tail field in the record's packed storage.
func (s *Efi_aurr_extended_panic_log) SetEfi_aurr_log_tail(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[1310456:1310460], uint32(v))
}

// Efi_aurr_log_head returns the Efi_aurr_log_head field from the record's packed storage.
func (s *Efi_aurr_extended_panic_log) Efi_aurr_log_head() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[1310460:1310464]))
}

// SetEfi_aurr_log_head updates the Efi_aurr_log_head field in the record's packed storage.
func (s *Efi_aurr_extended_panic_log) SetEfi_aurr_log_head(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[1310460:1310464], uint32(v))
}

// Efi_aurr_panic_header
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/efi_aurr_panic_header
type Efi_aurr_panic_header struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [24]byte
}

// Efi_aurr_magic returns the Efi_aurr_magic field from the record's packed storage.
func (s *Efi_aurr_panic_header) Efi_aurr_magic() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetEfi_aurr_magic updates the Efi_aurr_magic field in the record's packed storage.
func (s *Efi_aurr_panic_header) SetEfi_aurr_magic(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// Efi_aurr_crc returns the Efi_aurr_crc field from the record's packed storage.
func (s *Efi_aurr_panic_header) Efi_aurr_crc() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetEfi_aurr_crc updates the Efi_aurr_crc field in the record's packed storage.
func (s *Efi_aurr_panic_header) SetEfi_aurr_crc(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// Efi_aurr_version returns the Efi_aurr_version field from the record's packed storage.
func (s *Efi_aurr_panic_header) Efi_aurr_version() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetEfi_aurr_version updates the Efi_aurr_version field in the record's packed storage.
func (s *Efi_aurr_panic_header) SetEfi_aurr_version(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[8:12], uint32(v))
}

// Efi_aurr_reset_cause returns the Efi_aurr_reset_cause field from the record's packed storage.
func (s *Efi_aurr_panic_header) Efi_aurr_reset_cause() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[12:16]))
}

// SetEfi_aurr_reset_cause updates the Efi_aurr_reset_cause field in the record's packed storage.
func (s *Efi_aurr_panic_header) SetEfi_aurr_reset_cause(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[12:16], uint32(v))
}

// Efi_aurr_reset_log_offset returns the Efi_aurr_reset_log_offset field from the record's packed storage.
func (s *Efi_aurr_panic_header) Efi_aurr_reset_log_offset() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[16:20]))
}

// SetEfi_aurr_reset_log_offset updates the Efi_aurr_reset_log_offset field in the record's packed storage.
func (s *Efi_aurr_panic_header) SetEfi_aurr_reset_log_offset(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[16:20], uint32(v))
}

// Efi_aurr_reset_log_len returns the Efi_aurr_reset_log_len field from the record's packed storage.
func (s *Efi_aurr_panic_header) Efi_aurr_reset_log_len() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[20:24]))
}

// SetEfi_aurr_reset_log_len updates the Efi_aurr_reset_log_len field in the record's packed storage.
func (s *Efi_aurr_panic_header) SetEfi_aurr_reset_log_len(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[20:24], uint32(v))
}

// Embedded_panic_header
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/embedded_panic_header
type Embedded_panic_header struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [225]byte
}

// Eph_magic returns the Eph_magic field from the record's packed storage.
func (s *Embedded_panic_header) Eph_magic() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetEph_magic updates the Eph_magic field in the record's packed storage.
func (s *Embedded_panic_header) SetEph_magic(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// Eph_crc returns the Eph_crc field from the record's packed storage.
func (s *Embedded_panic_header) Eph_crc() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetEph_crc updates the Eph_crc field in the record's packed storage.
func (s *Embedded_panic_header) SetEph_crc(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// Eph_version returns the Eph_version field from the record's packed storage.
func (s *Embedded_panic_header) Eph_version() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetEph_version updates the Eph_version field in the record's packed storage.
func (s *Embedded_panic_header) SetEph_version(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[8:12], uint32(v))
}

// Eph_panic_flags returns the Eph_panic_flags field from the record's packed storage.
func (s *Embedded_panic_header) Eph_panic_flags() Eph_panic_flags_t {
	return Eph_panic_flags_t(binary.NativeEndian.Uint64(s.storage[12:20]))
}

// SetEph_panic_flags updates the Eph_panic_flags field in the record's packed storage.
func (s *Embedded_panic_header) SetEph_panic_flags(v Eph_panic_flags_t) {
	binary.NativeEndian.PutUint64(s.storage[12:20], uint64(v))
}

// Eph_panic_log_offset returns the Eph_panic_log_offset field from the record's packed storage.
func (s *Embedded_panic_header) Eph_panic_log_offset() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[20:24]))
}

// SetEph_panic_log_offset updates the Eph_panic_log_offset field in the record's packed storage.
func (s *Embedded_panic_header) SetEph_panic_log_offset(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[20:24], uint32(v))
}

// Eph_panic_log_len returns the Eph_panic_log_len field from the record's packed storage.
func (s *Embedded_panic_header) Eph_panic_log_len() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[24:28]))
}

// SetEph_panic_log_len updates the Eph_panic_log_len field in the record's packed storage.
func (s *Embedded_panic_header) SetEph_panic_log_len(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[24:28], uint32(v))
}

// Eph_stackshot_offset returns the Eph_stackshot_offset field from the record's packed storage.
func (s *Embedded_panic_header) Eph_stackshot_offset() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[28:32]))
}

// SetEph_stackshot_offset updates the Eph_stackshot_offset field in the record's packed storage.
func (s *Embedded_panic_header) SetEph_stackshot_offset(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[28:32], uint32(v))
}

// Eph_stackshot_len returns the Eph_stackshot_len field from the record's packed storage.
func (s *Embedded_panic_header) Eph_stackshot_len() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[32:36]))
}

// SetEph_stackshot_len updates the Eph_stackshot_len field in the record's packed storage.
func (s *Embedded_panic_header) SetEph_stackshot_len(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[32:36], uint32(v))
}

// Eph_other_log_offset returns the Eph_other_log_offset field from the record's packed storage.
func (s *Embedded_panic_header) Eph_other_log_offset() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[36:40]))
}

// SetEph_other_log_offset updates the Eph_other_log_offset field in the record's packed storage.
func (s *Embedded_panic_header) SetEph_other_log_offset(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[36:40], uint32(v))
}

// Eph_other_log_len returns the Eph_other_log_len field from the record's packed storage.
func (s *Embedded_panic_header) Eph_other_log_len() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[40:44]))
}

// SetEph_other_log_len updates the Eph_other_log_len field in the record's packed storage.
func (s *Embedded_panic_header) SetEph_other_log_len(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[40:44], uint32(v))
}

// Eph_x86_power_state returns the Eph_x86_power_state bitfield from the record's packed storage.
func (s *Embedded_panic_header) Eph_x86_power_state() uint64 {
	return uint64((s.storage[44] >> 0) & 0xff)
}

// SetEph_x86_power_state updates the Eph_x86_power_state bitfield in the record's packed storage.
func (s *Embedded_panic_header) SetEph_x86_power_state(v uint64) {
	s.storage[44] = (s.storage[44] &^ uint8(0xff<<0)) | uint8((uint8(v)&0xff)<<0)
}

// Eph_x86_efi_boot_state returns the Eph_x86_efi_boot_state bitfield from the record's packed storage.
func (s *Embedded_panic_header) Eph_x86_efi_boot_state() uint64 {
	return uint64((s.storage[45] >> 0) & 0xff)
}

// SetEph_x86_efi_boot_state updates the Eph_x86_efi_boot_state bitfield in the record's packed storage.
func (s *Embedded_panic_header) SetEph_x86_efi_boot_state(v uint64) {
	s.storage[45] = (s.storage[45] &^ uint8(0xff<<0)) | uint8((uint8(v)&0xff)<<0)
}

// Eph_x86_system_state returns the Eph_x86_system_state bitfield from the record's packed storage.
func (s *Embedded_panic_header) Eph_x86_system_state() uint64 {
	return uint64((s.storage[46] >> 0) & 0xff)
}

// SetEph_x86_system_state updates the Eph_x86_system_state bitfield in the record's packed storage.
func (s *Embedded_panic_header) SetEph_x86_system_state(v uint64) {
	s.storage[46] = (s.storage[46] &^ uint8(0xff<<0)) | uint8((uint8(v)&0xff)<<0)
}

// Eph_x86_unused_bits returns the Eph_x86_unused_bits bitfield from the record's packed storage.
func (s *Embedded_panic_header) Eph_x86_unused_bits() uint64 {
	return uint64((binary.NativeEndian.Uint64(s.storage[47:55]) >> 0) & 0xffffffffff)
}

// SetEph_x86_unused_bits updates the Eph_x86_unused_bits bitfield in the record's packed storage.
func (s *Embedded_panic_header) SetEph_x86_unused_bits(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[47:55], (binary.NativeEndian.Uint64(s.storage[47:55])&^(0xffffffffff<<0))|((uint64(v)&0xffffffffff)<<0))
}

// Eph_os_version returns the Eph_os_version field from the record's packed storage.
func (s *Embedded_panic_header) Eph_os_version() [32]int8 {
	return *(*[32]int8)(unsafe.Pointer(&s.storage[52]))
}

// SetEph_os_version updates the Eph_os_version field in the record's packed storage.
func (s *Embedded_panic_header) SetEph_os_version(v [32]int8) {
	*(*[32]int8)(unsafe.Pointer(&s.storage[52])) = v
}

// Eph_macos_version returns the Eph_macos_version field from the record's packed storage.
func (s *Embedded_panic_header) Eph_macos_version() [32]int8 {
	return *(*[32]int8)(unsafe.Pointer(&s.storage[84]))
}

// SetEph_macos_version updates the Eph_macos_version field in the record's packed storage.
func (s *Embedded_panic_header) SetEph_macos_version(v [32]int8) {
	*(*[32]int8)(unsafe.Pointer(&s.storage[84])) = v
}

// Eph_bootsessionuuid_string returns the Eph_bootsessionuuid_string field from the record's packed storage.
func (s *Embedded_panic_header) Eph_bootsessionuuid_string() [37]int8 {
	return *(*[37]int8)(unsafe.Pointer(&s.storage[116]))
}

// SetEph_bootsessionuuid_string updates the Eph_bootsessionuuid_string field in the record's packed storage.
func (s *Embedded_panic_header) SetEph_bootsessionuuid_string(v [37]int8) {
	*(*[37]int8)(unsafe.Pointer(&s.storage[116])) = v
}

// Eph_roots_installed returns the Eph_roots_installed field from the record's packed storage.
func (s *Embedded_panic_header) Eph_roots_installed() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[153:161]))
}

// SetEph_roots_installed updates the Eph_roots_installed field in the record's packed storage.
func (s *Embedded_panic_header) SetEph_roots_installed(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[153:161], uint64(v))
}

// Eph_ext_paniclog_offset returns the Eph_ext_paniclog_offset field from the record's packed storage.
func (s *Embedded_panic_header) Eph_ext_paniclog_offset() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[161:165]))
}

// SetEph_ext_paniclog_offset updates the Eph_ext_paniclog_offset field in the record's packed storage.
func (s *Embedded_panic_header) SetEph_ext_paniclog_offset(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[161:165], uint32(v))
}

// Eph_ext_paniclog_len returns the Eph_ext_paniclog_len field from the record's packed storage.
func (s *Embedded_panic_header) Eph_ext_paniclog_len() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[165:169]))
}

// SetEph_ext_paniclog_len updates the Eph_ext_paniclog_len field in the record's packed storage.
func (s *Embedded_panic_header) SetEph_ext_paniclog_len(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[165:169], uint32(v))
}

// Eph_panic_initiator_offset returns the Eph_panic_initiator_offset field from the record's packed storage.
func (s *Embedded_panic_header) Eph_panic_initiator_offset() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[169:173]))
}

// SetEph_panic_initiator_offset updates the Eph_panic_initiator_offset field in the record's packed storage.
func (s *Embedded_panic_header) SetEph_panic_initiator_offset(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[169:173], uint32(v))
}

// Eph_panic_initiator_len returns the Eph_panic_initiator_len field from the record's packed storage.
func (s *Embedded_panic_header) Eph_panic_initiator_len() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[173:177]))
}

// SetEph_panic_initiator_len updates the Eph_panic_initiator_len field in the record's packed storage.
func (s *Embedded_panic_header) SetEph_panic_initiator_len(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[173:177], uint32(v))
}

// Eph_device_target_type returns the Eph_device_target_type field from the record's packed storage.
func (s *Embedded_panic_header) Eph_device_target_type() [16]int8 {
	return *(*[16]int8)(unsafe.Pointer(&s.storage[177]))
}

// SetEph_device_target_type updates the Eph_device_target_type field in the record's packed storage.
func (s *Embedded_panic_header) SetEph_device_target_type(v [16]int8) {
	*(*[16]int8)(unsafe.Pointer(&s.storage[177])) = v
}

// Eph_device_model_type returns the Eph_device_model_type field from the record's packed storage.
func (s *Embedded_panic_header) Eph_device_model_type() [32]int8 {
	return *(*[32]int8)(unsafe.Pointer(&s.storage[193]))
}

// SetEph_device_model_type updates the Eph_device_model_type field in the record's packed storage.
func (s *Embedded_panic_header) SetEph_device_model_type(v [32]int8) {
	*(*[32]int8)(unsafe.Pointer(&s.storage[193])) = v
}

// Encryption_info_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/encryption_info_command
type Encryption_info_command struct {
	Cmd       uint32
	Cmdsize   uint32
	Cryptoff  uint32
	Cryptsize uint32
	Cryptid   uint32
}

// Encryption_info_command_64
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/encryption_info_command_64
type Encryption_info_command_64 struct {
	Cmd       uint32
	Cmdsize   uint32
	Cryptoff  uint32
	Cryptsize uint32
	Cryptid   uint32
	Pad       uint32
}

// Entry_point_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/entry_point_command
type Entry_point_command struct {
	Cmd       uint32
	Cmdsize   uint32
	Entryoff  uint64
	Stacksize uint64
}

// Esp
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/esp
type Esp struct {
	Esp_spi U_int32_t
}

// Esptail
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/esptail
type Esptail struct {
	Esp_padlen U_int8_t
	Esp_nxt    U_int8_t
}

// Ether_arp
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ether_arp
type Ether_arp struct {
	Ea_hdr  Arphdr
	Arp_sha [6]U_char
	Arp_spa [4]U_char
	Arp_tha [6]U_char
	Arp_tpa [4]U_char
}

// Ether_vlan_header
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ether_vlan_header
type Ether_vlan_header struct {
	Evl_dhost       [6]U_char
	Evl_shost       [6]U_char
	Evl_encap_proto U_int16_t
	Evl_tag         U_int16_t
	Evl_proto       U_int16_t
}

// Exclave_addressspace_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/exclave_addressspace_info
type Exclave_addressspace_info struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [40]byte
}

// Eas_id returns the Eas_id field from the record's packed storage.
func (s *Exclave_addressspace_info) Eas_id() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[0:8]))
}

// SetEas_id updates the Eas_id field in the record's packed storage.
func (s *Exclave_addressspace_info) SetEas_id(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[0:8], uint64(v))
}

// Eas_flags returns the Eas_flags field from the record's packed storage.
func (s *Exclave_addressspace_info) Eas_flags() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[8:16]))
}

// SetEas_flags updates the Eas_flags field in the record's packed storage.
func (s *Exclave_addressspace_info) SetEas_flags(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[8:16], uint64(v))
}

// Eas_layoutid returns the Eas_layoutid field from the record's packed storage.
func (s *Exclave_addressspace_info) Eas_layoutid() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[16:24]))
}

// SetEas_layoutid updates the Eas_layoutid field in the record's packed storage.
func (s *Exclave_addressspace_info) SetEas_layoutid(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[16:24], uint64(v))
}

// Eas_slide returns the Eas_slide field from the record's packed storage.
func (s *Exclave_addressspace_info) Eas_slide() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[24:32]))
}

// SetEas_slide updates the Eas_slide field in the record's packed storage.
func (s *Exclave_addressspace_info) SetEas_slide(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[24:32], uint64(v))
}

// Eas_asroot returns the Eas_asroot field from the record's packed storage.
func (s *Exclave_addressspace_info) Eas_asroot() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[32:40]))
}

// SetEas_asroot updates the Eas_asroot field in the record's packed storage.
func (s *Exclave_addressspace_info) SetEas_asroot(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[32:40], uint64(v))
}

// Exclave_ipcstackentry_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/exclave_ipcstackentry_info
type Exclave_ipcstackentry_info struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [32]byte
}

// Eise_asid returns the Eise_asid field from the record's packed storage.
func (s *Exclave_ipcstackentry_info) Eise_asid() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[0:8]))
}

// SetEise_asid updates the Eise_asid field in the record's packed storage.
func (s *Exclave_ipcstackentry_info) SetEise_asid(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[0:8], uint64(v))
}

// Eise_tnid returns the Eise_tnid field from the record's packed storage.
func (s *Exclave_ipcstackentry_info) Eise_tnid() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[8:16]))
}

// SetEise_tnid updates the Eise_tnid field in the record's packed storage.
func (s *Exclave_ipcstackentry_info) SetEise_tnid(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[8:16], uint64(v))
}

// Eise_invocationid returns the Eise_invocationid field from the record's packed storage.
func (s *Exclave_ipcstackentry_info) Eise_invocationid() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[16:24]))
}

// SetEise_invocationid updates the Eise_invocationid field in the record's packed storage.
func (s *Exclave_ipcstackentry_info) SetEise_invocationid(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[16:24], uint64(v))
}

// Eise_flags returns the Eise_flags field from the record's packed storage.
func (s *Exclave_ipcstackentry_info) Eise_flags() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[24:32]))
}

// SetEise_flags updates the Eise_flags field in the record's packed storage.
func (s *Exclave_ipcstackentry_info) SetEise_flags(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[24:32], uint64(v))
}

// Exclave_scresult_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/exclave_scresult_info
type Exclave_scresult_info struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [16]byte
}

// Esc_id returns the Esc_id field from the record's packed storage.
func (s *Exclave_scresult_info) Esc_id() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[0:8]))
}

// SetEsc_id updates the Esc_id field in the record's packed storage.
func (s *Exclave_scresult_info) SetEsc_id(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[0:8], uint64(v))
}

// Esc_flags returns the Esc_flags field from the record's packed storage.
func (s *Exclave_scresult_info) Esc_flags() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[8:16]))
}

// SetEsc_flags updates the Esc_flags field in the record's packed storage.
func (s *Exclave_scresult_info) SetEsc_flags(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[8:16], uint64(v))
}

// Exclave_textlayout_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/exclave_textlayout_info
type Exclave_textlayout_info struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [20]byte
}

// Layout_id returns the Layout_id field from the record's packed storage.
func (s *Exclave_textlayout_info) Layout_id() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[0:8]))
}

// SetLayout_id updates the Layout_id field in the record's packed storage.
func (s *Exclave_textlayout_info) SetLayout_id(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[0:8], uint64(v))
}

// Etl_flags returns the Etl_flags field from the record's packed storage.
func (s *Exclave_textlayout_info) Etl_flags() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[8:16]))
}

// SetEtl_flags updates the Etl_flags field in the record's packed storage.
func (s *Exclave_textlayout_info) SetEtl_flags(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[8:16], uint64(v))
}

// Sharedcache_index returns the Sharedcache_index field from the record's packed storage.
func (s *Exclave_textlayout_info) Sharedcache_index() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[16:20]))
}

// SetSharedcache_index updates the Sharedcache_index field in the record's packed storage.
func (s *Exclave_textlayout_info) SetSharedcache_index(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[16:20], uint32(v))
}

// Exclave_textlayout_info_v1
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/exclave_textlayout_info_v1
type Exclave_textlayout_info_v1 struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [16]byte
}

// Layout_id returns the Layout_id field from the record's packed storage.
func (s *Exclave_textlayout_info_v1) Layout_id() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[0:8]))
}

// SetLayout_id updates the Layout_id field in the record's packed storage.
func (s *Exclave_textlayout_info_v1) SetLayout_id(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[0:8], uint64(v))
}

// Etl_flags returns the Etl_flags field from the record's packed storage.
func (s *Exclave_textlayout_info_v1) Etl_flags() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[8:16]))
}

// SetEtl_flags updates the Etl_flags field in the record's packed storage.
func (s *Exclave_textlayout_info_v1) SetEtl_flags(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[8:16], uint64(v))
}

// Exclave_textlayout_segment
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/exclave_textlayout_segment
type Exclave_textlayout_segment struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [24]byte
}

// LayoutSegment_uuid returns the LayoutSegment_uuid field from the record's packed storage.
func (s *Exclave_textlayout_segment) LayoutSegment_uuid() [16]uint8 {
	return *(*[16]uint8)(unsafe.Pointer(&s.storage[0]))
}

// SetLayoutSegment_uuid updates the LayoutSegment_uuid field in the record's packed storage.
func (s *Exclave_textlayout_segment) SetLayoutSegment_uuid(v [16]uint8) {
	*(*[16]uint8)(unsafe.Pointer(&s.storage[0])) = v
}

// LayoutSegment_loadAddress returns the LayoutSegment_loadAddress field from the record's packed storage.
func (s *Exclave_textlayout_segment) LayoutSegment_loadAddress() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[16:24]))
}

// SetLayoutSegment_loadAddress updates the LayoutSegment_loadAddress field in the record's packed storage.
func (s *Exclave_textlayout_segment) SetLayoutSegment_loadAddress(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[16:24], uint64(v))
}

// Exclave_textlayout_segment_v2
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/exclave_textlayout_segment_v2
type Exclave_textlayout_segment_v2 struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [32]byte
}

// LayoutSegment_uuid returns the LayoutSegment_uuid field from the record's packed storage.
func (s *Exclave_textlayout_segment_v2) LayoutSegment_uuid() [16]uint8 {
	return *(*[16]uint8)(unsafe.Pointer(&s.storage[0]))
}

// SetLayoutSegment_uuid updates the LayoutSegment_uuid field in the record's packed storage.
func (s *Exclave_textlayout_segment_v2) SetLayoutSegment_uuid(v [16]uint8) {
	*(*[16]uint8)(unsafe.Pointer(&s.storage[0])) = v
}

// LayoutSegment_loadAddress returns the LayoutSegment_loadAddress field from the record's packed storage.
func (s *Exclave_textlayout_segment_v2) LayoutSegment_loadAddress() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[16:24]))
}

// SetLayoutSegment_loadAddress updates the LayoutSegment_loadAddress field in the record's packed storage.
func (s *Exclave_textlayout_segment_v2) SetLayoutSegment_loadAddress(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[16:24], uint64(v))
}

// LayoutSegment_rawLoadAddress returns the LayoutSegment_rawLoadAddress field from the record's packed storage.
func (s *Exclave_textlayout_segment_v2) LayoutSegment_rawLoadAddress() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[24:32]))
}

// SetLayoutSegment_rawLoadAddress updates the LayoutSegment_rawLoadAddress field in the record's packed storage.
func (s *Exclave_textlayout_segment_v2) SetLayoutSegment_rawLoadAddress(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[24:32], uint64(v))
}

// Exit_reason_snapshot
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/exit_reason_snapshot
type Exit_reason_snapshot struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [20]byte
}

// Ers_namespace returns the Ers_namespace field from the record's packed storage.
func (s *Exit_reason_snapshot) Ers_namespace() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetErs_namespace updates the Ers_namespace field in the record's packed storage.
func (s *Exit_reason_snapshot) SetErs_namespace(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// Ers_code returns the Ers_code field from the record's packed storage.
func (s *Exit_reason_snapshot) Ers_code() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[4:12]))
}

// SetErs_code updates the Ers_code field in the record's packed storage.
func (s *Exit_reason_snapshot) SetErs_code(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[4:12], uint64(v))
}

// Ers_flags returns the Ers_flags field from the record's packed storage.
func (s *Exit_reason_snapshot) Ers_flags() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[12:20]))
}

// SetErs_flags updates the Ers_flags field in the record's packed storage.
func (s *Exit_reason_snapshot) SetErs_flags(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[12:20], uint64(v))
}

// Fairplay_subsystem
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/fairplay_subsystem-4tk
type Fairplay_subsystem struct {
	Server   unsafe.Pointer
	Start    Mach_msg_id_t
	End      Mach_msg_id_t
	Maxsize  uint32
	Reserved Vm_address_t
	Routine  unsafe.Pointer
}

// Fat_arch - Describes the location within the binary of an object file targeted at a single architecture. Declared in `/usr/include/mach-o/fat.h`.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/fat_arch
type Fat_arch struct {
	Cputype    int32
	Cpusubtype int32  // An enumeration value of type `cpu_subtype_t`. Specifies the specific member of the CPU family on which this entry may be used or a constant specifying all members.
	Offset     uint32 // Offset to the beginning of the data for this CPU.
	Size       uint32 // Size of the data for this CPU.
	Align      uint32 // The power of 2 alignment for the offset of the object file for the architecture specified in `cputype` within the binary. This is required to ensure that, if this binary is changed, the contents it retains are correctly aligned for virtual memory paging and other uses.

}

// Fat_header
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/fat_header
type Fat_header struct {
	Magic     uint32 // An integer containing the value `0xCAFEBABE` in big-endian byte order format. On a big-endian host CPU, this can be validated using the constant `FAT_MAGIC`; on a little-endian host CPU, it can be validated using the constant `FAT_CIGAM`.
	Nfat_arch uint32 // An integer specifying the number of [fat_arch](<https://developer.apple.com/documentation/kernel/fat_arch>) data structures that follow. This is the number of architectures contained in this binary.

}

// Fdisk_part
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/fdisk_part
type Fdisk_part struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [16]byte
}

// Bootid returns the Bootid field from the record's packed storage.
func (s *Fdisk_part) Bootid() uint8 {
	return uint8(s.storage[0])
}

// SetBootid updates the Bootid field in the record's packed storage.
func (s *Fdisk_part) SetBootid(v uint8) {
	s.storage[0] = uint8(v)
}

// Beghead returns the Beghead field from the record's packed storage.
func (s *Fdisk_part) Beghead() uint8 {
	return uint8(s.storage[1])
}

// SetBeghead updates the Beghead field in the record's packed storage.
func (s *Fdisk_part) SetBeghead(v uint8) {
	s.storage[1] = uint8(v)
}

// Begsect returns the Begsect field from the record's packed storage.
func (s *Fdisk_part) Begsect() uint8 {
	return uint8(s.storage[2])
}

// SetBegsect updates the Begsect field in the record's packed storage.
func (s *Fdisk_part) SetBegsect(v uint8) {
	s.storage[2] = uint8(v)
}

// Begcyl returns the Begcyl field from the record's packed storage.
func (s *Fdisk_part) Begcyl() uint8 {
	return uint8(s.storage[3])
}

// SetBegcyl updates the Begcyl field in the record's packed storage.
func (s *Fdisk_part) SetBegcyl(v uint8) {
	s.storage[3] = uint8(v)
}

// Systid returns the Systid field from the record's packed storage.
func (s *Fdisk_part) Systid() uint8 {
	return uint8(s.storage[4])
}

// SetSystid updates the Systid field in the record's packed storage.
func (s *Fdisk_part) SetSystid(v uint8) {
	s.storage[4] = uint8(v)
}

// Endhead returns the Endhead field from the record's packed storage.
func (s *Fdisk_part) Endhead() uint8 {
	return uint8(s.storage[5])
}

// SetEndhead updates the Endhead field in the record's packed storage.
func (s *Fdisk_part) SetEndhead(v uint8) {
	s.storage[5] = uint8(v)
}

// Endsect returns the Endsect field from the record's packed storage.
func (s *Fdisk_part) Endsect() uint8 {
	return uint8(s.storage[6])
}

// SetEndsect updates the Endsect field in the record's packed storage.
func (s *Fdisk_part) SetEndsect(v uint8) {
	s.storage[6] = uint8(v)
}

// Endcyl returns the Endcyl field from the record's packed storage.
func (s *Fdisk_part) Endcyl() uint8 {
	return uint8(s.storage[7])
}

// SetEndcyl updates the Endcyl field in the record's packed storage.
func (s *Fdisk_part) SetEndcyl(v uint8) {
	s.storage[7] = uint8(v)
}

// Relsect returns the Relsect field from the record's packed storage.
func (s *Fdisk_part) Relsect() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetRelsect updates the Relsect field in the record's packed storage.
func (s *Fdisk_part) SetRelsect(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[8:12], uint32(v))
}

// Numsect returns the Numsect field from the record's packed storage.
func (s *Fdisk_part) Numsect() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[12:16]))
}

// SetNumsect updates the Numsect field in the record's packed storage.
func (s *Fdisk_part) SetNumsect(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[12:16], uint32(v))
}

// Fileset_entry_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/fileset_entry_command
type Fileset_entry_command struct {
	Cmd      uint32
	Cmdsize  uint32
	Vmaddr   uint64
	Fileoff  uint64
	Entry_id [1]uint32
	Reserved uint32
}

// Flock
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/flock
type Flock struct {
	L_start  int64
	L_len    int64
	L_pid    int32
	L_type   int16
	L_whence int16
}

// Flocktimeout
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/flocktimeout
type Flocktimeout struct {
	Fl      Flock
	Timeout Timespec
}

// Frmrinfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/frmrinfo
type Frmrinfo struct {
	Frmr_rej_pdu0    U_int8_t
	Frmr_rej_pdu1    U_int8_t
	Frmr_control     U_int8_t
	Frmr_control_ext U_int8_t
	Frmr_cause       U_int8_t
}

// Fssearchblock
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/fssearchblock
type Fssearchblock struct {
	Returnattrs         *Attrlist
	Returnbuffer        unsafe.Pointer
	Returnbuffersize    uintptr
	Maxmatches          U_long
	Timelimit           Timeval
	Searchparams1       unsafe.Pointer
	Sizeofsearchparams1 uintptr
	Searchparams2       unsafe.Pointer
	Sizeofsearchparams2 uintptr
	Searchattrs         Attrlist
}

// Fvmfile_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/fvmfile_command
type Fvmfile_command struct {
	Cmd         uint32
	Cmdsize     uint32
	Name        [1]uint32
	Header_addr uint32
}

// Fvmlib
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/fvmlib
type Fvmlib struct {
	Name          [1]uint32
	Minor_version uint32
	Header_addr   uint32
}

// Fvmlib_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/fvmlib_command
type Fvmlib_command struct {
	Cmd     uint32
	Cmdsize uint32
	Fvmlib  Fvmlib
}

// Gpt_ent
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/gpt_ent
type Gpt_ent struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [128]byte
}

// Ent_type returns the Ent_type field from the record's packed storage.
func (s *Gpt_ent) Ent_type() [16]uint8 {
	return *(*[16]uint8)(unsafe.Pointer(&s.storage[0]))
}

// SetEnt_type updates the Ent_type field in the record's packed storage.
func (s *Gpt_ent) SetEnt_type(v [16]uint8) {
	*(*[16]uint8)(unsafe.Pointer(&s.storage[0])) = v
}

// Ent_uuid returns the Ent_uuid field from the record's packed storage.
func (s *Gpt_ent) Ent_uuid() [16]uint8 {
	return *(*[16]uint8)(unsafe.Pointer(&s.storage[16]))
}

// SetEnt_uuid updates the Ent_uuid field in the record's packed storage.
func (s *Gpt_ent) SetEnt_uuid(v [16]uint8) {
	*(*[16]uint8)(unsafe.Pointer(&s.storage[16])) = v
}

// Ent_lba_start returns the Ent_lba_start field from the record's packed storage.
func (s *Gpt_ent) Ent_lba_start() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[32:40]))
}

// SetEnt_lba_start updates the Ent_lba_start field in the record's packed storage.
func (s *Gpt_ent) SetEnt_lba_start(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[32:40], uint64(v))
}

// Ent_lba_end returns the Ent_lba_end field from the record's packed storage.
func (s *Gpt_ent) Ent_lba_end() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[40:48]))
}

// SetEnt_lba_end updates the Ent_lba_end field in the record's packed storage.
func (s *Gpt_ent) SetEnt_lba_end(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[40:48], uint64(v))
}

// Ent_attr returns the Ent_attr field from the record's packed storage.
func (s *Gpt_ent) Ent_attr() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[48:56]))
}

// SetEnt_attr updates the Ent_attr field in the record's packed storage.
func (s *Gpt_ent) SetEnt_attr(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[48:56], uint64(v))
}

// Ent_name returns the Ent_name field from the record's packed storage.
func (s *Gpt_ent) Ent_name() [36]uint16 {
	return *(*[36]uint16)(unsafe.Pointer(&s.storage[56]))
}

// SetEnt_name updates the Ent_name field in the record's packed storage.
func (s *Gpt_ent) SetEnt_name(v [36]uint16) {
	*(*[36]uint16)(unsafe.Pointer(&s.storage[56])) = v
}

// Gpt_hdr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/gpt_hdr
type Gpt_hdr struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [96]byte
}

// Hdr_sig returns the Hdr_sig field from the record's packed storage.
func (s *Gpt_hdr) Hdr_sig() [8]uint8 {
	return *(*[8]uint8)(unsafe.Pointer(&s.storage[0]))
}

// SetHdr_sig updates the Hdr_sig field in the record's packed storage.
func (s *Gpt_hdr) SetHdr_sig(v [8]uint8) {
	*(*[8]uint8)(unsafe.Pointer(&s.storage[0])) = v
}

// Hdr_revision returns the Hdr_revision field from the record's packed storage.
func (s *Gpt_hdr) Hdr_revision() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetHdr_revision updates the Hdr_revision field in the record's packed storage.
func (s *Gpt_hdr) SetHdr_revision(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[8:12], uint32(v))
}

// Hdr_size returns the Hdr_size field from the record's packed storage.
func (s *Gpt_hdr) Hdr_size() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[12:16]))
}

// SetHdr_size updates the Hdr_size field in the record's packed storage.
func (s *Gpt_hdr) SetHdr_size(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[12:16], uint32(v))
}

// Hdr_crc_self returns the Hdr_crc_self field from the record's packed storage.
func (s *Gpt_hdr) Hdr_crc_self() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[16:20]))
}

// SetHdr_crc_self updates the Hdr_crc_self field in the record's packed storage.
func (s *Gpt_hdr) SetHdr_crc_self(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[16:20], uint32(v))
}

// __reserved returns the __reserved field from the record's packed storage.
func (s *Gpt_hdr) __reserved() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[20:24]))
}

// Set__reserved updates the __reserved field in the record's packed storage.
func (s *Gpt_hdr) Set__reserved(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[20:24], uint32(v))
}

// Hdr_lba_self returns the Hdr_lba_self field from the record's packed storage.
func (s *Gpt_hdr) Hdr_lba_self() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[24:32]))
}

// SetHdr_lba_self updates the Hdr_lba_self field in the record's packed storage.
func (s *Gpt_hdr) SetHdr_lba_self(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[24:32], uint64(v))
}

// Hdr_lba_alt returns the Hdr_lba_alt field from the record's packed storage.
func (s *Gpt_hdr) Hdr_lba_alt() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[32:40]))
}

// SetHdr_lba_alt updates the Hdr_lba_alt field in the record's packed storage.
func (s *Gpt_hdr) SetHdr_lba_alt(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[32:40], uint64(v))
}

// Hdr_lba_start returns the Hdr_lba_start field from the record's packed storage.
func (s *Gpt_hdr) Hdr_lba_start() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[40:48]))
}

// SetHdr_lba_start updates the Hdr_lba_start field in the record's packed storage.
func (s *Gpt_hdr) SetHdr_lba_start(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[40:48], uint64(v))
}

// Hdr_lba_end returns the Hdr_lba_end field from the record's packed storage.
func (s *Gpt_hdr) Hdr_lba_end() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[48:56]))
}

// SetHdr_lba_end updates the Hdr_lba_end field in the record's packed storage.
func (s *Gpt_hdr) SetHdr_lba_end(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[48:56], uint64(v))
}

// Hdr_uuid returns the Hdr_uuid field from the record's packed storage.
func (s *Gpt_hdr) Hdr_uuid() [16]uint8 {
	return *(*[16]uint8)(unsafe.Pointer(&s.storage[56]))
}

// SetHdr_uuid updates the Hdr_uuid field in the record's packed storage.
func (s *Gpt_hdr) SetHdr_uuid(v [16]uint8) {
	*(*[16]uint8)(unsafe.Pointer(&s.storage[56])) = v
}

// Hdr_lba_table returns the Hdr_lba_table field from the record's packed storage.
func (s *Gpt_hdr) Hdr_lba_table() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[72:80]))
}

// SetHdr_lba_table updates the Hdr_lba_table field in the record's packed storage.
func (s *Gpt_hdr) SetHdr_lba_table(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[72:80], uint64(v))
}

// Hdr_entries returns the Hdr_entries field from the record's packed storage.
func (s *Gpt_hdr) Hdr_entries() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[80:84]))
}

// SetHdr_entries updates the Hdr_entries field in the record's packed storage.
func (s *Gpt_hdr) SetHdr_entries(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[80:84], uint32(v))
}

// Hdr_entsz returns the Hdr_entsz field from the record's packed storage.
func (s *Gpt_hdr) Hdr_entsz() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[84:88]))
}

// SetHdr_entsz updates the Hdr_entsz field in the record's packed storage.
func (s *Gpt_hdr) SetHdr_entsz(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[84:88], uint32(v))
}

// Hdr_crc_table returns the Hdr_crc_table field from the record's packed storage.
func (s *Gpt_hdr) Hdr_crc_table() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[88:92]))
}

// SetHdr_crc_table updates the Hdr_crc_table field in the record's packed storage.
func (s *Gpt_hdr) SetHdr_crc_table(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[88:92], uint32(v))
}

// Padding returns the Padding field from the record's packed storage.
func (s *Gpt_hdr) Padding() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[92:96]))
}

// SetPadding updates the Padding field in the record's packed storage.
func (s *Gpt_hdr) SetPadding(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[92:96], uint32(v))
}

// Group_req
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/group_req
type Group_req struct {
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
	storage [132]byte
}

// Gr_interface returns the Gr_interface field from the record's packed storage.
func (s *Group_req) Gr_interface() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetGr_interface updates the Gr_interface field in the record's packed storage.
func (s *Group_req) SetGr_interface(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// Gr_group returns the Gr_group field from the record's packed storage.
func (s *Group_req) Gr_group() [16]uint64 {
	return *(*[16]uint64)(unsafe.Pointer(&s.storage[4]))
}

// SetGr_group updates the Gr_group field in the record's packed storage.
func (s *Group_req) SetGr_group(v [16]uint64) {
	*(*[16]uint64)(unsafe.Pointer(&s.storage[4])) = v
}

// Group_source_req
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/group_source_req
type Group_source_req struct {
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
	storage [260]byte
}

// Gsr_interface returns the Gsr_interface field from the record's packed storage.
func (s *Group_source_req) Gsr_interface() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetGsr_interface updates the Gsr_interface field in the record's packed storage.
func (s *Group_source_req) SetGsr_interface(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// Gsr_group returns the Gsr_group field from the record's packed storage.
func (s *Group_source_req) Gsr_group() [16]uint64 {
	return *(*[16]uint64)(unsafe.Pointer(&s.storage[4]))
}

// SetGsr_group updates the Gsr_group field in the record's packed storage.
func (s *Group_source_req) SetGsr_group(v [16]uint64) {
	*(*[16]uint64)(unsafe.Pointer(&s.storage[4])) = v
}

// Gsr_source returns the Gsr_source field from the record's packed storage.
func (s *Group_source_req) Gsr_source() [16]uint64 {
	return *(*[16]uint64)(unsafe.Pointer(&s.storage[132]))
}

// SetGsr_source updates the Gsr_source field in the record's packed storage.
func (s *Group_source_req) SetGsr_source(v [16]uint64) {
	*(*[16]uint64)(unsafe.Pointer(&s.storage[132])) = v
}

// Hfs_mount_args
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/hfs_mount_args
type Hfs_mount_args struct {
	Hfs_uid              uint32
	Hfs_gid              uint32
	Hfs_mask             uint16
	Hfs_encoding         U_int32_t
	Hfs_timezone         Timezone
	Flags                int32
	Journal_tbuffer_size int32
	Journal_flags        int32
	Journal_disable      int32
}

// Icmp
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/icmp
type Icmp struct {
	Icmp_type  U_char
	Icmp_code  U_char
	Icmp_cksum U_short
	Icmp_hun   [1]uint32
	Icmp_dun   [5]uint32
}

// Icmp6_filter
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/icmp6_filter
type Icmp6_filter struct {
	Icmp6_filt [8]U_int32_t
}

// Icmp6_hdr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/icmp6_hdr
type Icmp6_hdr struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [8]byte
}

// Icmp6_type returns the Icmp6_type field from the record's packed storage.
func (s *Icmp6_hdr) Icmp6_type() U_int8_t {
	return *(*U_int8_t)(unsafe.Pointer(&s.storage[0]))
}

// SetIcmp6_type updates the Icmp6_type field in the record's packed storage.
func (s *Icmp6_hdr) SetIcmp6_type(v U_int8_t) {
	*(*U_int8_t)(unsafe.Pointer(&s.storage[0])) = v
}

// Icmp6_code returns the Icmp6_code field from the record's packed storage.
func (s *Icmp6_hdr) Icmp6_code() U_int8_t {
	return *(*U_int8_t)(unsafe.Pointer(&s.storage[1]))
}

// SetIcmp6_code updates the Icmp6_code field in the record's packed storage.
func (s *Icmp6_hdr) SetIcmp6_code(v U_int8_t) {
	*(*U_int8_t)(unsafe.Pointer(&s.storage[1])) = v
}

// Icmp6_cksum returns the Icmp6_cksum field from the record's packed storage.
func (s *Icmp6_hdr) Icmp6_cksum() U_int16_t {
	return U_int16_t(binary.NativeEndian.Uint16(s.storage[2:4]))
}

// SetIcmp6_cksum updates the Icmp6_cksum field in the record's packed storage.
func (s *Icmp6_hdr) SetIcmp6_cksum(v U_int16_t) {
	binary.NativeEndian.PutUint16(s.storage[2:4], uint16(v))
}

// Icmp6_dataun returns the Icmp6_dataun field from the record's packed storage.
func (s *Icmp6_hdr) Icmp6_dataun() [4]byte {
	return *(*[4]byte)(unsafe.Pointer(&s.storage[4]))
}

// SetIcmp6_dataun updates the Icmp6_dataun field in the record's packed storage.
func (s *Icmp6_hdr) SetIcmp6_dataun(v [4]byte) {
	*(*[4]byte)(unsafe.Pointer(&s.storage[4])) = v
}

// Icmp6_ifstat
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/icmp6_ifstat
type Icmp6_ifstat struct {
	Ifs6_in_msg              U_quad_t
	Ifs6_in_error            U_quad_t
	Ifs6_in_dstunreach       U_quad_t
	Ifs6_in_adminprohib      U_quad_t
	Ifs6_in_timeexceed       U_quad_t
	Ifs6_in_paramprob        U_quad_t
	Ifs6_in_pkttoobig        U_quad_t
	Ifs6_in_echo             U_quad_t
	Ifs6_in_echoreply        U_quad_t
	Ifs6_in_routersolicit    U_quad_t
	Ifs6_in_routeradvert     U_quad_t
	Ifs6_in_neighborsolicit  U_quad_t
	Ifs6_in_neighboradvert   U_quad_t
	Ifs6_in_redirect         U_quad_t
	Ifs6_in_mldquery         U_quad_t
	Ifs6_in_mldreport        U_quad_t
	Ifs6_in_mlddone          U_quad_t
	Ifs6_out_msg             U_quad_t
	Ifs6_out_error           U_quad_t
	Ifs6_out_dstunreach      U_quad_t
	Ifs6_out_adminprohib     U_quad_t
	Ifs6_out_timeexceed      U_quad_t
	Ifs6_out_paramprob       U_quad_t
	Ifs6_out_pkttoobig       U_quad_t
	Ifs6_out_echo            U_quad_t
	Ifs6_out_echoreply       U_quad_t
	Ifs6_out_routersolicit   U_quad_t
	Ifs6_out_routeradvert    U_quad_t
	Ifs6_out_neighborsolicit U_quad_t
	Ifs6_out_neighboradvert  U_quad_t
	Ifs6_out_redirect        U_quad_t
	Ifs6_out_mldquery        U_quad_t
	Ifs6_out_mldreport       U_quad_t
	Ifs6_out_mlddone         U_quad_t
}

// Icmp6_namelookup
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/icmp6_namelookup
type Icmp6_namelookup struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [20]byte
}

// Icmp6_nl_hdr returns the Icmp6_nl_hdr field from the record's packed storage.
func (s *Icmp6_namelookup) Icmp6_nl_hdr() Icmp6_hdr {
	return *(*Icmp6_hdr)(unsafe.Pointer(&s.storage[0]))
}

// SetIcmp6_nl_hdr updates the Icmp6_nl_hdr field in the record's packed storage.
func (s *Icmp6_namelookup) SetIcmp6_nl_hdr(v Icmp6_hdr) {
	*(*Icmp6_hdr)(unsafe.Pointer(&s.storage[0])) = v
}

// Icmp6_nl_nonce returns the Icmp6_nl_nonce field from the record's packed storage.
func (s *Icmp6_namelookup) Icmp6_nl_nonce() [8]U_int8_t {
	return *(*[8]U_int8_t)(unsafe.Pointer(&s.storage[8]))
}

// SetIcmp6_nl_nonce updates the Icmp6_nl_nonce field in the record's packed storage.
func (s *Icmp6_namelookup) SetIcmp6_nl_nonce(v [8]U_int8_t) {
	*(*[8]U_int8_t)(unsafe.Pointer(&s.storage[8])) = v
}

// Icmp6_nl_ttl returns the Icmp6_nl_ttl field from the record's packed storage.
func (s *Icmp6_namelookup) Icmp6_nl_ttl() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[16:20]))
}

// SetIcmp6_nl_ttl updates the Icmp6_nl_ttl field in the record's packed storage.
func (s *Icmp6_namelookup) SetIcmp6_nl_ttl(v int32) {
	binary.NativeEndian.PutUint32(s.storage[16:20], uint32(v))
}

// Icmp6_nodeinfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/icmp6_nodeinfo
type Icmp6_nodeinfo struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [16]byte
}

// Icmp6_ni_hdr returns the Icmp6_ni_hdr field from the record's packed storage.
func (s *Icmp6_nodeinfo) Icmp6_ni_hdr() Icmp6_hdr {
	return *(*Icmp6_hdr)(unsafe.Pointer(&s.storage[0]))
}

// SetIcmp6_ni_hdr updates the Icmp6_ni_hdr field in the record's packed storage.
func (s *Icmp6_nodeinfo) SetIcmp6_ni_hdr(v Icmp6_hdr) {
	*(*Icmp6_hdr)(unsafe.Pointer(&s.storage[0])) = v
}

// Icmp6_ni_nonce returns the Icmp6_ni_nonce field from the record's packed storage.
func (s *Icmp6_nodeinfo) Icmp6_ni_nonce() [8]U_int8_t {
	return *(*[8]U_int8_t)(unsafe.Pointer(&s.storage[8]))
}

// SetIcmp6_ni_nonce updates the Icmp6_ni_nonce field in the record's packed storage.
func (s *Icmp6_nodeinfo) SetIcmp6_ni_nonce(v [8]U_int8_t) {
	*(*[8]U_int8_t)(unsafe.Pointer(&s.storage[8])) = v
}

// Icmp6_router_renum
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/icmp6_router_renum
type Icmp6_router_renum struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [16]byte
}

// Rr_hdr returns the Rr_hdr field from the record's packed storage.
func (s *Icmp6_router_renum) Rr_hdr() Icmp6_hdr {
	return *(*Icmp6_hdr)(unsafe.Pointer(&s.storage[0]))
}

// SetRr_hdr updates the Rr_hdr field in the record's packed storage.
func (s *Icmp6_router_renum) SetRr_hdr(v Icmp6_hdr) {
	*(*Icmp6_hdr)(unsafe.Pointer(&s.storage[0])) = v
}

// Rr_segnum returns the Rr_segnum field from the record's packed storage.
func (s *Icmp6_router_renum) Rr_segnum() U_int8_t {
	return *(*U_int8_t)(unsafe.Pointer(&s.storage[8]))
}

// SetRr_segnum updates the Rr_segnum field in the record's packed storage.
func (s *Icmp6_router_renum) SetRr_segnum(v U_int8_t) {
	*(*U_int8_t)(unsafe.Pointer(&s.storage[8])) = v
}

// Rr_flags returns the Rr_flags field from the record's packed storage.
func (s *Icmp6_router_renum) Rr_flags() U_int8_t {
	return *(*U_int8_t)(unsafe.Pointer(&s.storage[9]))
}

// SetRr_flags updates the Rr_flags field in the record's packed storage.
func (s *Icmp6_router_renum) SetRr_flags(v U_int8_t) {
	*(*U_int8_t)(unsafe.Pointer(&s.storage[9])) = v
}

// Rr_maxdelay returns the Rr_maxdelay field from the record's packed storage.
func (s *Icmp6_router_renum) Rr_maxdelay() U_int16_t {
	return U_int16_t(binary.NativeEndian.Uint16(s.storage[10:12]))
}

// SetRr_maxdelay updates the Rr_maxdelay field in the record's packed storage.
func (s *Icmp6_router_renum) SetRr_maxdelay(v U_int16_t) {
	binary.NativeEndian.PutUint16(s.storage[10:12], uint16(v))
}

// Rr_reserved returns the Rr_reserved field from the record's packed storage.
func (s *Icmp6_router_renum) Rr_reserved() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[12:16]))
}

// SetRr_reserved updates the Rr_reserved field in the record's packed storage.
func (s *Icmp6_router_renum) SetRr_reserved(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[12:16], uint32(v))
}

// Icmp6errstat
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/icmp6errstat
type Icmp6errstat struct {
	Icp6errs_dst_unreach_noroute     U_quad_t
	Icp6errs_dst_unreach_admin       U_quad_t
	Icp6errs_dst_unreach_beyondscope U_quad_t
	Icp6errs_dst_unreach_addr        U_quad_t
	Icp6errs_dst_unreach_noport      U_quad_t
	Icp6errs_packet_too_big          U_quad_t
	Icp6errs_time_exceed_transit     U_quad_t
	Icp6errs_time_exceed_reassembly  U_quad_t
	Icp6errs_paramprob_header        U_quad_t
	Icp6errs_paramprob_nextheader    U_quad_t
	Icp6errs_paramprob_option        U_quad_t
	Icp6errs_redirect                U_quad_t
	Icp6errs_unknown                 U_quad_t
}

// Icmp6stat
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/icmp6stat
type Icmp6stat struct {
	Icp6s_error         U_quad_t
	Icp6s_canterror     U_quad_t
	Icp6s_toofreq       U_quad_t
	Icp6s_outhist       [256]U_quad_t
	Icp6s_badcode       U_quad_t
	Icp6s_tooshort      U_quad_t
	Icp6s_checksum      U_quad_t
	Icp6s_badlen        U_quad_t
	Icp6s_reflect       U_quad_t
	Icp6s_inhist        [256]U_quad_t
	Icp6s_nd_toomanyopt U_quad_t
	Icp6s_outerrhist    Icmp6errstat
	Icp6s_pmtuchg       U_quad_t
	Icp6s_nd_badopt     U_quad_t
	Icp6s_badns         U_quad_t
	Icp6s_badna         U_quad_t
	Icp6s_badrs         U_quad_t
	Icp6s_badra         U_quad_t
	Icp6s_badredirect   U_quad_t
	Icp6s_rfc6980_drop  U_quad_t
	Icp6s_badpkttoobig  U_quad_t
}

// Icmp_ra_addr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/icmp_ra_addr
type Icmp_ra_addr struct {
	Ira_addr       U_int32_t
	Ira_preference U_int32_t
}

// Icmpstat
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/icmpstat
type Icmpstat struct {
	Icps_error        U_int32_t
	Icps_oldshort     U_int32_t
	Icps_oldicmp      U_int32_t
	Icps_outhist      [41]U_int32_t
	Icps_badcode      U_int32_t
	Icps_tooshort     U_int32_t
	Icps_checksum     U_int32_t
	Icps_badlen       U_int32_t
	Icps_reflect      U_int32_t
	Icps_inhist       [41]U_int32_t
	Icps_bmcastecho   U_int32_t
	Icps_bmcasttstamp U_int32_t
}

// Ident_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ident_command
type Ident_command struct {
	Cmd     uint32
	Cmdsize uint32
}

// If_agentidreq
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_agentidreq
type If_agentidreq struct {
	Ifar_name [16]int8
	Ifar_uuid [16]uint8
}

// If_agentidsreq
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_agentidsreq
type If_agentidsreq struct {
	Ifar_name  [16]int8
	Ifar_count U_int32_t
	Ifar_uuids *[16]byte
}

// If_bandwidths
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_bandwidths
type If_bandwidths struct {
	Eff_bw uint64
	Max_bw uint64
}

// If_cellular_status
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_cellular_status
type If_cellular_status struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [104]byte
}

// If_cell_u returns the If_cell_u field from the record's packed storage.
func (s *If_cellular_status) If_cell_u() [104]byte {
	return *(*[104]byte)(unsafe.Pointer(&s.storage[0]))
}

// SetIf_cell_u updates the If_cell_u field in the record's packed storage.
func (s *If_cellular_status) SetIf_cell_u(v [104]byte) {
	*(*[104]byte)(unsafe.Pointer(&s.storage[0])) = v
}

// If_cellular_status_v1
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_cellular_status_v1
type If_cellular_status_v1 struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [104]byte
}

// Valid_bitmask returns the Valid_bitmask field from the record's packed storage.
func (s *If_cellular_status_v1) Valid_bitmask() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetValid_bitmask updates the Valid_bitmask field in the record's packed storage.
func (s *If_cellular_status_v1) SetValid_bitmask(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// Link_quality_metric returns the Link_quality_metric field from the record's packed storage.
func (s *If_cellular_status_v1) Link_quality_metric() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetLink_quality_metric updates the Link_quality_metric field in the record's packed storage.
func (s *If_cellular_status_v1) SetLink_quality_metric(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// Ul_effective_bandwidth returns the Ul_effective_bandwidth field from the record's packed storage.
func (s *If_cellular_status_v1) Ul_effective_bandwidth() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetUl_effective_bandwidth updates the Ul_effective_bandwidth field in the record's packed storage.
func (s *If_cellular_status_v1) SetUl_effective_bandwidth(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[8:12], uint32(v))
}

// Ul_max_bandwidth returns the Ul_max_bandwidth field from the record's packed storage.
func (s *If_cellular_status_v1) Ul_max_bandwidth() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[12:16]))
}

// SetUl_max_bandwidth updates the Ul_max_bandwidth field in the record's packed storage.
func (s *If_cellular_status_v1) SetUl_max_bandwidth(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[12:16], uint32(v))
}

// Ul_min_latency returns the Ul_min_latency field from the record's packed storage.
func (s *If_cellular_status_v1) Ul_min_latency() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[16:20]))
}

// SetUl_min_latency updates the Ul_min_latency field in the record's packed storage.
func (s *If_cellular_status_v1) SetUl_min_latency(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[16:20], uint32(v))
}

// Ul_effective_latency returns the Ul_effective_latency field from the record's packed storage.
func (s *If_cellular_status_v1) Ul_effective_latency() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[20:24]))
}

// SetUl_effective_latency updates the Ul_effective_latency field in the record's packed storage.
func (s *If_cellular_status_v1) SetUl_effective_latency(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[20:24], uint32(v))
}

// Ul_max_latency returns the Ul_max_latency field from the record's packed storage.
func (s *If_cellular_status_v1) Ul_max_latency() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[24:28]))
}

// SetUl_max_latency updates the Ul_max_latency field in the record's packed storage.
func (s *If_cellular_status_v1) SetUl_max_latency(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[24:28], uint32(v))
}

// Ul_retxt_level returns the Ul_retxt_level field from the record's packed storage.
func (s *If_cellular_status_v1) Ul_retxt_level() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[28:32]))
}

// SetUl_retxt_level updates the Ul_retxt_level field in the record's packed storage.
func (s *If_cellular_status_v1) SetUl_retxt_level(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[28:32], uint32(v))
}

// Ul_bytes_lost returns the Ul_bytes_lost field from the record's packed storage.
func (s *If_cellular_status_v1) Ul_bytes_lost() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[32:36]))
}

// SetUl_bytes_lost updates the Ul_bytes_lost field in the record's packed storage.
func (s *If_cellular_status_v1) SetUl_bytes_lost(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[32:36], uint32(v))
}

// Ul_min_queue_size returns the Ul_min_queue_size field from the record's packed storage.
func (s *If_cellular_status_v1) Ul_min_queue_size() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[36:40]))
}

// SetUl_min_queue_size updates the Ul_min_queue_size field in the record's packed storage.
func (s *If_cellular_status_v1) SetUl_min_queue_size(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[36:40], uint32(v))
}

// Ul_avg_queue_size returns the Ul_avg_queue_size field from the record's packed storage.
func (s *If_cellular_status_v1) Ul_avg_queue_size() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[40:44]))
}

// SetUl_avg_queue_size updates the Ul_avg_queue_size field in the record's packed storage.
func (s *If_cellular_status_v1) SetUl_avg_queue_size(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[40:44], uint32(v))
}

// Ul_max_queue_size returns the Ul_max_queue_size field from the record's packed storage.
func (s *If_cellular_status_v1) Ul_max_queue_size() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[44:48]))
}

// SetUl_max_queue_size updates the Ul_max_queue_size field in the record's packed storage.
func (s *If_cellular_status_v1) SetUl_max_queue_size(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[44:48], uint32(v))
}

// Dl_effective_bandwidth returns the Dl_effective_bandwidth field from the record's packed storage.
func (s *If_cellular_status_v1) Dl_effective_bandwidth() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[48:52]))
}

// SetDl_effective_bandwidth updates the Dl_effective_bandwidth field in the record's packed storage.
func (s *If_cellular_status_v1) SetDl_effective_bandwidth(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[48:52], uint32(v))
}

// Dl_max_bandwidth returns the Dl_max_bandwidth field from the record's packed storage.
func (s *If_cellular_status_v1) Dl_max_bandwidth() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[52:56]))
}

// SetDl_max_bandwidth updates the Dl_max_bandwidth field in the record's packed storage.
func (s *If_cellular_status_v1) SetDl_max_bandwidth(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[52:56], uint32(v))
}

// Config_inactivity_time returns the Config_inactivity_time field from the record's packed storage.
func (s *If_cellular_status_v1) Config_inactivity_time() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[56:60]))
}

// SetConfig_inactivity_time updates the Config_inactivity_time field in the record's packed storage.
func (s *If_cellular_status_v1) SetConfig_inactivity_time(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[56:60], uint32(v))
}

// Config_backoff_time returns the Config_backoff_time field from the record's packed storage.
func (s *If_cellular_status_v1) Config_backoff_time() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[60:64]))
}

// SetConfig_backoff_time updates the Config_backoff_time field in the record's packed storage.
func (s *If_cellular_status_v1) SetConfig_backoff_time(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[60:64], uint32(v))
}

// Mss_recommended returns the Mss_recommended field from the record's packed storage.
func (s *If_cellular_status_v1) Mss_recommended() U_int16_t {
	return U_int16_t(binary.NativeEndian.Uint16(s.storage[64:66]))
}

// SetMss_recommended updates the Mss_recommended field in the record's packed storage.
func (s *If_cellular_status_v1) SetMss_recommended(v U_int16_t) {
	binary.NativeEndian.PutUint16(s.storage[64:66], uint16(v))
}

// Reserved_1 returns the Reserved_1 field from the record's packed storage.
func (s *If_cellular_status_v1) Reserved_1() U_int16_t {
	return U_int16_t(binary.NativeEndian.Uint16(s.storage[66:68]))
}

// SetReserved_1 updates the Reserved_1 field in the record's packed storage.
func (s *If_cellular_status_v1) SetReserved_1(v U_int16_t) {
	binary.NativeEndian.PutUint16(s.storage[66:68], uint16(v))
}

// Reserved_2 returns the Reserved_2 field from the record's packed storage.
func (s *If_cellular_status_v1) Reserved_2() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[68:72]))
}

// SetReserved_2 updates the Reserved_2 field in the record's packed storage.
func (s *If_cellular_status_v1) SetReserved_2(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[68:72], uint32(v))
}

// Reserved_3 returns the Reserved_3 field from the record's packed storage.
func (s *If_cellular_status_v1) Reserved_3() U_int64_t {
	return U_int64_t(binary.NativeEndian.Uint64(s.storage[72:80]))
}

// SetReserved_3 updates the Reserved_3 field in the record's packed storage.
func (s *If_cellular_status_v1) SetReserved_3(v U_int64_t) {
	binary.NativeEndian.PutUint64(s.storage[72:80], uint64(v))
}

// Reserved_4 returns the Reserved_4 field from the record's packed storage.
func (s *If_cellular_status_v1) Reserved_4() U_int64_t {
	return U_int64_t(binary.NativeEndian.Uint64(s.storage[80:88]))
}

// SetReserved_4 updates the Reserved_4 field in the record's packed storage.
func (s *If_cellular_status_v1) SetReserved_4(v U_int64_t) {
	binary.NativeEndian.PutUint64(s.storage[80:88], uint64(v))
}

// Reserved_5 returns the Reserved_5 field from the record's packed storage.
func (s *If_cellular_status_v1) Reserved_5() U_int64_t {
	return U_int64_t(binary.NativeEndian.Uint64(s.storage[88:96]))
}

// SetReserved_5 updates the Reserved_5 field in the record's packed storage.
func (s *If_cellular_status_v1) SetReserved_5(v U_int64_t) {
	binary.NativeEndian.PutUint64(s.storage[88:96], uint64(v))
}

// Reserved_6 returns the Reserved_6 field from the record's packed storage.
func (s *If_cellular_status_v1) Reserved_6() U_int64_t {
	return U_int64_t(binary.NativeEndian.Uint64(s.storage[96:104]))
}

// SetReserved_6 updates the Reserved_6 field in the record's packed storage.
func (s *If_cellular_status_v1) SetReserved_6(v U_int64_t) {
	binary.NativeEndian.PutUint64(s.storage[96:104], uint64(v))
}

// If_clat46req
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_clat46req
type If_clat46req struct {
	Ifclat46_name [16]int8
	Ifclat46_addr If_ipv6_address
}

// If_clonereq
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_clonereq
type If_clonereq struct {
	Ifcr_total  int32
	Ifcr_count  int32
	Ifcr_buffer *byte
}

// If_data
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_data
type If_data struct {
	Ifi_type       U_char
	Ifi_typelen    U_char
	Ifi_physical   U_char
	Ifi_addrlen    U_char
	Ifi_hdrlen     U_char
	Ifi_recvquota  U_char
	Ifi_xmitquota  U_char
	Ifi_unused1    U_char
	Ifi_mtu        U_int32_t
	Ifi_metric     U_int32_t
	Ifi_baudrate   U_int32_t
	Ifi_ipackets   U_int32_t
	Ifi_ierrors    U_int32_t
	Ifi_opackets   U_int32_t
	Ifi_oerrors    U_int32_t
	Ifi_collisions U_int32_t
	Ifi_ibytes     U_int32_t
	Ifi_obytes     U_int32_t
	Ifi_imcasts    U_int32_t
	Ifi_omcasts    U_int32_t
	Ifi_iqdrops    U_int32_t
	Ifi_noproto    U_int32_t
	Ifi_recvtiming U_int32_t
	Ifi_xmittiming U_int32_t
	Ifi_lastchange Timeval32
	Ifi_unused2    U_int32_t
	Ifi_hwassist   U_int32_t
	Ifi_reserved1  U_int32_t
	Ifi_reserved2  U_int32_t
}

// If_data64
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_data64
type If_data64 struct {
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
	storage [128]byte
}

// Ifi_type returns the Ifi_type field from the record's packed storage.
func (s *If_data64) Ifi_type() U_char {
	return *(*U_char)(unsafe.Pointer(&s.storage[0]))
}

// SetIfi_type updates the Ifi_type field in the record's packed storage.
func (s *If_data64) SetIfi_type(v U_char) {
	*(*U_char)(unsafe.Pointer(&s.storage[0])) = v
}

// Ifi_typelen returns the Ifi_typelen field from the record's packed storage.
func (s *If_data64) Ifi_typelen() U_char {
	return *(*U_char)(unsafe.Pointer(&s.storage[1]))
}

// SetIfi_typelen updates the Ifi_typelen field in the record's packed storage.
func (s *If_data64) SetIfi_typelen(v U_char) {
	*(*U_char)(unsafe.Pointer(&s.storage[1])) = v
}

// Ifi_physical returns the Ifi_physical field from the record's packed storage.
func (s *If_data64) Ifi_physical() U_char {
	return *(*U_char)(unsafe.Pointer(&s.storage[2]))
}

// SetIfi_physical updates the Ifi_physical field in the record's packed storage.
func (s *If_data64) SetIfi_physical(v U_char) {
	*(*U_char)(unsafe.Pointer(&s.storage[2])) = v
}

// Ifi_addrlen returns the Ifi_addrlen field from the record's packed storage.
func (s *If_data64) Ifi_addrlen() U_char {
	return *(*U_char)(unsafe.Pointer(&s.storage[3]))
}

// SetIfi_addrlen updates the Ifi_addrlen field in the record's packed storage.
func (s *If_data64) SetIfi_addrlen(v U_char) {
	*(*U_char)(unsafe.Pointer(&s.storage[3])) = v
}

// Ifi_hdrlen returns the Ifi_hdrlen field from the record's packed storage.
func (s *If_data64) Ifi_hdrlen() U_char {
	return *(*U_char)(unsafe.Pointer(&s.storage[4]))
}

// SetIfi_hdrlen updates the Ifi_hdrlen field in the record's packed storage.
func (s *If_data64) SetIfi_hdrlen(v U_char) {
	*(*U_char)(unsafe.Pointer(&s.storage[4])) = v
}

// Ifi_recvquota returns the Ifi_recvquota field from the record's packed storage.
func (s *If_data64) Ifi_recvquota() U_char {
	return *(*U_char)(unsafe.Pointer(&s.storage[5]))
}

// SetIfi_recvquota updates the Ifi_recvquota field in the record's packed storage.
func (s *If_data64) SetIfi_recvquota(v U_char) {
	*(*U_char)(unsafe.Pointer(&s.storage[5])) = v
}

// Ifi_xmitquota returns the Ifi_xmitquota field from the record's packed storage.
func (s *If_data64) Ifi_xmitquota() U_char {
	return *(*U_char)(unsafe.Pointer(&s.storage[6]))
}

// SetIfi_xmitquota updates the Ifi_xmitquota field in the record's packed storage.
func (s *If_data64) SetIfi_xmitquota(v U_char) {
	*(*U_char)(unsafe.Pointer(&s.storage[6])) = v
}

// Ifi_unused1 returns the Ifi_unused1 field from the record's packed storage.
func (s *If_data64) Ifi_unused1() U_char {
	return *(*U_char)(unsafe.Pointer(&s.storage[7]))
}

// SetIfi_unused1 updates the Ifi_unused1 field in the record's packed storage.
func (s *If_data64) SetIfi_unused1(v U_char) {
	*(*U_char)(unsafe.Pointer(&s.storage[7])) = v
}

// Ifi_mtu returns the Ifi_mtu field from the record's packed storage.
func (s *If_data64) Ifi_mtu() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetIfi_mtu updates the Ifi_mtu field in the record's packed storage.
func (s *If_data64) SetIfi_mtu(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[8:12], uint32(v))
}

// Ifi_metric returns the Ifi_metric field from the record's packed storage.
func (s *If_data64) Ifi_metric() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[12:16]))
}

// SetIfi_metric updates the Ifi_metric field in the record's packed storage.
func (s *If_data64) SetIfi_metric(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[12:16], uint32(v))
}

// Ifi_baudrate returns the Ifi_baudrate field from the record's packed storage.
func (s *If_data64) Ifi_baudrate() U_int64_t {
	return U_int64_t(binary.NativeEndian.Uint64(s.storage[16:24]))
}

// SetIfi_baudrate updates the Ifi_baudrate field in the record's packed storage.
func (s *If_data64) SetIfi_baudrate(v U_int64_t) {
	binary.NativeEndian.PutUint64(s.storage[16:24], uint64(v))
}

// Ifi_ipackets returns the Ifi_ipackets field from the record's packed storage.
func (s *If_data64) Ifi_ipackets() U_int64_t {
	return U_int64_t(binary.NativeEndian.Uint64(s.storage[24:32]))
}

// SetIfi_ipackets updates the Ifi_ipackets field in the record's packed storage.
func (s *If_data64) SetIfi_ipackets(v U_int64_t) {
	binary.NativeEndian.PutUint64(s.storage[24:32], uint64(v))
}

// Ifi_ierrors returns the Ifi_ierrors field from the record's packed storage.
func (s *If_data64) Ifi_ierrors() U_int64_t {
	return U_int64_t(binary.NativeEndian.Uint64(s.storage[32:40]))
}

// SetIfi_ierrors updates the Ifi_ierrors field in the record's packed storage.
func (s *If_data64) SetIfi_ierrors(v U_int64_t) {
	binary.NativeEndian.PutUint64(s.storage[32:40], uint64(v))
}

// Ifi_opackets returns the Ifi_opackets field from the record's packed storage.
func (s *If_data64) Ifi_opackets() U_int64_t {
	return U_int64_t(binary.NativeEndian.Uint64(s.storage[40:48]))
}

// SetIfi_opackets updates the Ifi_opackets field in the record's packed storage.
func (s *If_data64) SetIfi_opackets(v U_int64_t) {
	binary.NativeEndian.PutUint64(s.storage[40:48], uint64(v))
}

// Ifi_oerrors returns the Ifi_oerrors field from the record's packed storage.
func (s *If_data64) Ifi_oerrors() U_int64_t {
	return U_int64_t(binary.NativeEndian.Uint64(s.storage[48:56]))
}

// SetIfi_oerrors updates the Ifi_oerrors field in the record's packed storage.
func (s *If_data64) SetIfi_oerrors(v U_int64_t) {
	binary.NativeEndian.PutUint64(s.storage[48:56], uint64(v))
}

// Ifi_collisions returns the Ifi_collisions field from the record's packed storage.
func (s *If_data64) Ifi_collisions() U_int64_t {
	return U_int64_t(binary.NativeEndian.Uint64(s.storage[56:64]))
}

// SetIfi_collisions updates the Ifi_collisions field in the record's packed storage.
func (s *If_data64) SetIfi_collisions(v U_int64_t) {
	binary.NativeEndian.PutUint64(s.storage[56:64], uint64(v))
}

// Ifi_ibytes returns the Ifi_ibytes field from the record's packed storage.
func (s *If_data64) Ifi_ibytes() U_int64_t {
	return U_int64_t(binary.NativeEndian.Uint64(s.storage[64:72]))
}

// SetIfi_ibytes updates the Ifi_ibytes field in the record's packed storage.
func (s *If_data64) SetIfi_ibytes(v U_int64_t) {
	binary.NativeEndian.PutUint64(s.storage[64:72], uint64(v))
}

// Ifi_obytes returns the Ifi_obytes field from the record's packed storage.
func (s *If_data64) Ifi_obytes() U_int64_t {
	return U_int64_t(binary.NativeEndian.Uint64(s.storage[72:80]))
}

// SetIfi_obytes updates the Ifi_obytes field in the record's packed storage.
func (s *If_data64) SetIfi_obytes(v U_int64_t) {
	binary.NativeEndian.PutUint64(s.storage[72:80], uint64(v))
}

// Ifi_imcasts returns the Ifi_imcasts field from the record's packed storage.
func (s *If_data64) Ifi_imcasts() U_int64_t {
	return U_int64_t(binary.NativeEndian.Uint64(s.storage[80:88]))
}

// SetIfi_imcasts updates the Ifi_imcasts field in the record's packed storage.
func (s *If_data64) SetIfi_imcasts(v U_int64_t) {
	binary.NativeEndian.PutUint64(s.storage[80:88], uint64(v))
}

// Ifi_omcasts returns the Ifi_omcasts field from the record's packed storage.
func (s *If_data64) Ifi_omcasts() U_int64_t {
	return U_int64_t(binary.NativeEndian.Uint64(s.storage[88:96]))
}

// SetIfi_omcasts updates the Ifi_omcasts field in the record's packed storage.
func (s *If_data64) SetIfi_omcasts(v U_int64_t) {
	binary.NativeEndian.PutUint64(s.storage[88:96], uint64(v))
}

// Ifi_iqdrops returns the Ifi_iqdrops field from the record's packed storage.
func (s *If_data64) Ifi_iqdrops() U_int64_t {
	return U_int64_t(binary.NativeEndian.Uint64(s.storage[96:104]))
}

// SetIfi_iqdrops updates the Ifi_iqdrops field in the record's packed storage.
func (s *If_data64) SetIfi_iqdrops(v U_int64_t) {
	binary.NativeEndian.PutUint64(s.storage[96:104], uint64(v))
}

// Ifi_noproto returns the Ifi_noproto field from the record's packed storage.
func (s *If_data64) Ifi_noproto() U_int64_t {
	return U_int64_t(binary.NativeEndian.Uint64(s.storage[104:112]))
}

// SetIfi_noproto updates the Ifi_noproto field in the record's packed storage.
func (s *If_data64) SetIfi_noproto(v U_int64_t) {
	binary.NativeEndian.PutUint64(s.storage[104:112], uint64(v))
}

// Ifi_recvtiming returns the Ifi_recvtiming field from the record's packed storage.
func (s *If_data64) Ifi_recvtiming() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[112:116]))
}

// SetIfi_recvtiming updates the Ifi_recvtiming field in the record's packed storage.
func (s *If_data64) SetIfi_recvtiming(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[112:116], uint32(v))
}

// Ifi_xmittiming returns the Ifi_xmittiming field from the record's packed storage.
func (s *If_data64) Ifi_xmittiming() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[116:120]))
}

// SetIfi_xmittiming updates the Ifi_xmittiming field in the record's packed storage.
func (s *If_data64) SetIfi_xmittiming(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[116:120], uint32(v))
}

// Ifi_lastchange returns the Ifi_lastchange field from the record's packed storage.
func (s *If_data64) Ifi_lastchange() Timeval32 {
	return *(*Timeval32)(unsafe.Pointer(&s.storage[120]))
}

// SetIfi_lastchange updates the Ifi_lastchange field in the record's packed storage.
func (s *If_data64) SetIfi_lastchange(v Timeval32) {
	*(*Timeval32)(unsafe.Pointer(&s.storage[120])) = v
}

// If_data_extended
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_data_extended
type If_data_extended struct {
	Ifi_alignerrs U_int64_t
	Ifi_dt_bytes  U_int64_t
	Ifi_fpackets  U_int64_t
	Ifi_fbytes    U_int64_t
	Reserved      [12]U_int64_t
}

// If_descreq
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_descreq
type If_descreq struct {
	Ifdr_name [16]int8
	Ifdr_len  U_int32_t
	Ifdr_desc [128]U_int8_t
}

// If_description
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_description
type If_description struct {
	Ifd_maxlen U_int32_t
	Ifd_len    U_int32_t
	Ifd_desc   *U_int8_t
}

// If_interface_state
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_interface_state
type If_interface_state struct {
	Valid_bitmask          U_int8_t
	Rrc_state              U_int8_t
	Lqm_state              int8
	Interface_availability U_int8_t
}

// If_ipv6_address
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_ipv6_address
type If_ipv6_address struct {
	V6_address   [4]uint32
	V6_prefixlen uint32
}

// If_latencies
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_latencies
type If_latencies struct {
	Eff_lt U_int64_t
	Max_lt U_int64_t
}

// If_lim_perf_stat
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_lim_perf_stat
type If_lim_perf_stat struct {
	Lim_dl_max_bandwidth     U_int64_t
	Lim_ul_max_bandwidth     U_int64_t
	Lim_total_txpkts         U_int64_t
	Lim_total_rxpkts         U_int64_t
	Lim_total_retxpkts       U_int64_t
	Lim_packet_loss_percent  U_int64_t
	Lim_total_oopkts         U_int64_t
	Lim_packet_ooo_percent   U_int64_t
	Lim_rtt_variance         U_int64_t
	Lim_rtt_average          U_int64_t
	Lim_rtt_min              U_int64_t
	Lim_conn_timeouts        U_int64_t
	Lim_conn_attempts        U_int64_t
	Lim_conn_timeout_percent U_int64_t
	Lim_bk_txpkts            U_int64_t
	bitfield15               uint8
}

// Lim_dl_detected returns the Lim_dl_detected bitfield.
func (s *If_lim_perf_stat) Lim_dl_detected() uint8 {
	return (s.bitfield15 >> 0) & ((1 << 1) - 1)
}

// SetLim_dl_detected updates the Lim_dl_detected bitfield.
func (s *If_lim_perf_stat) SetLim_dl_detected(v uint8) {
	const mask uint8 = (1 << 1) - 1
	s.bitfield15 = (s.bitfield15 &^ (mask << 0)) | ((v & mask) << 0)
}

// Lim_ul_detected returns the Lim_ul_detected bitfield.
func (s *If_lim_perf_stat) Lim_ul_detected() uint8 {
	return (s.bitfield15 >> 1) & ((1 << 1) - 1)
}

// SetLim_ul_detected updates the Lim_ul_detected bitfield.
func (s *If_lim_perf_stat) SetLim_ul_detected(v uint8) {
	const mask uint8 = (1 << 1) - 1
	s.bitfield15 = (s.bitfield15 &^ (mask << 1)) | ((v & mask) << 1)
}

// If_link_status
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_link_status
type If_link_status struct {
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
	storage [120]byte
}

// Ifsr_version returns the Ifsr_version field from the record's packed storage.
func (s *If_link_status) Ifsr_version() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetIfsr_version updates the Ifsr_version field in the record's packed storage.
func (s *If_link_status) SetIfsr_version(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// Ifsr_len returns the Ifsr_len field from the record's packed storage.
func (s *If_link_status) Ifsr_len() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetIfsr_len updates the Ifsr_len field in the record's packed storage.
func (s *If_link_status) SetIfsr_len(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// Ifsr_u returns the Ifsr_u field from the record's packed storage.
func (s *If_link_status) Ifsr_u() [112]byte {
	return *(*[112]byte)(unsafe.Pointer(&s.storage[8]))
}

// SetIfsr_u updates the Ifsr_u field in the record's packed storage.
func (s *If_link_status) SetIfsr_u(v [112]byte) {
	*(*[112]byte)(unsafe.Pointer(&s.storage[8])) = v
}

// If_linkheuristics
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_linkheuristics
type If_linkheuristics struct {
	Iflh_link_heuristics_cnt      U_int64_t
	Iflh_link_heuristics_time     U_int64_t
	Iflh_congested_link_cnt       U_int64_t
	Iflh_congested_link_time      U_int64_t
	Iflh_lqm_good_cnt             U_int64_t
	Iflh_lqm_good_time            U_int64_t
	Iflh_lqm_poor_cnt             U_int64_t
	Iflh_lqm_poor_time            U_int64_t
	Iflh_lqm_min_viable_cnt       U_int64_t
	Iflh_lqm_min_viable_time      U_int64_t
	Iflh_lqm_bad_cnt              U_int64_t
	Iflh_lqm_bad_time             U_int64_t
	Iflh_tcp_linkheur_stealthdrop U_int64_t
	Iflh_tcp_linkheur_noackpri    U_int64_t
	Iflh_tcp_linkheur_comprxmt    U_int64_t
	Iflh_tcp_linkheur_synrxmt     U_int64_t
	Iflh_tcp_linkheur_rxmtfloor   U_int64_t
	Iflh_udp_linkheur_stealthdrop U_int64_t
}

// If_linkparamsreq
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_linkparamsreq
type If_linkparamsreq struct {
	Iflpr_name               [16]int8
	Iflpr_flags              U_int32_t
	Iflpr_output_sched       U_int32_t
	Iflpr_output_tbr_rate    U_int64_t
	Iflpr_output_tbr_percent U_int32_t
	Iflpr_input_tbr_rate     U_int64_t
	Iflpr_output_bw          If_bandwidths
	Iflpr_input_bw           If_bandwidths
	Iflpr_output_lt          If_latencies
	Iflpr_input_lt           If_latencies
	Iflpr_input_netem        If_netem_params
	Iflpr_output_netem       If_netem_params
}

// If_msghdr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_msghdr
type If_msghdr struct {
	Ifm_msglen  uint16
	Ifm_version uint8
	Ifm_type    uint8
	Ifm_addrs   int32
	Ifm_flags   int32
	Ifm_index   uint16
	Ifm_data    If_data
}

// If_msghdr2
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_msghdr2
type If_msghdr2 struct {
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
	storage [160]byte
}

// Ifm_msglen returns the Ifm_msglen field from the record's packed storage.
func (s *If_msghdr2) Ifm_msglen() U_short {
	return U_short(binary.NativeEndian.Uint16(s.storage[0:2]))
}

// SetIfm_msglen updates the Ifm_msglen field in the record's packed storage.
func (s *If_msghdr2) SetIfm_msglen(v U_short) {
	binary.NativeEndian.PutUint16(s.storage[0:2], uint16(v))
}

// Ifm_version returns the Ifm_version field from the record's packed storage.
func (s *If_msghdr2) Ifm_version() U_char {
	return *(*U_char)(unsafe.Pointer(&s.storage[2]))
}

// SetIfm_version updates the Ifm_version field in the record's packed storage.
func (s *If_msghdr2) SetIfm_version(v U_char) {
	*(*U_char)(unsafe.Pointer(&s.storage[2])) = v
}

// Ifm_type returns the Ifm_type field from the record's packed storage.
func (s *If_msghdr2) Ifm_type() U_char {
	return *(*U_char)(unsafe.Pointer(&s.storage[3]))
}

// SetIfm_type updates the Ifm_type field in the record's packed storage.
func (s *If_msghdr2) SetIfm_type(v U_char) {
	*(*U_char)(unsafe.Pointer(&s.storage[3])) = v
}

// Ifm_addrs returns the Ifm_addrs field from the record's packed storage.
func (s *If_msghdr2) Ifm_addrs() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetIfm_addrs updates the Ifm_addrs field in the record's packed storage.
func (s *If_msghdr2) SetIfm_addrs(v int32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// Ifm_flags returns the Ifm_flags field from the record's packed storage.
func (s *If_msghdr2) Ifm_flags() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetIfm_flags updates the Ifm_flags field in the record's packed storage.
func (s *If_msghdr2) SetIfm_flags(v int32) {
	binary.NativeEndian.PutUint32(s.storage[8:12], uint32(v))
}

// Ifm_index returns the Ifm_index field from the record's packed storage.
func (s *If_msghdr2) Ifm_index() U_short {
	return U_short(binary.NativeEndian.Uint16(s.storage[12:16]))
}

// SetIfm_index updates the Ifm_index field in the record's packed storage.
func (s *If_msghdr2) SetIfm_index(v U_short) {
	binary.NativeEndian.PutUint16(s.storage[12:16], uint16(v))
}

// Ifm_snd_len returns the Ifm_snd_len field from the record's packed storage.
func (s *If_msghdr2) Ifm_snd_len() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[16:20]))
}

// SetIfm_snd_len updates the Ifm_snd_len field in the record's packed storage.
func (s *If_msghdr2) SetIfm_snd_len(v int32) {
	binary.NativeEndian.PutUint32(s.storage[16:20], uint32(v))
}

// Ifm_snd_maxlen returns the Ifm_snd_maxlen field from the record's packed storage.
func (s *If_msghdr2) Ifm_snd_maxlen() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[20:24]))
}

// SetIfm_snd_maxlen updates the Ifm_snd_maxlen field in the record's packed storage.
func (s *If_msghdr2) SetIfm_snd_maxlen(v int32) {
	binary.NativeEndian.PutUint32(s.storage[20:24], uint32(v))
}

// Ifm_snd_drops returns the Ifm_snd_drops field from the record's packed storage.
func (s *If_msghdr2) Ifm_snd_drops() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[24:28]))
}

// SetIfm_snd_drops updates the Ifm_snd_drops field in the record's packed storage.
func (s *If_msghdr2) SetIfm_snd_drops(v int32) {
	binary.NativeEndian.PutUint32(s.storage[24:28], uint32(v))
}

// Ifm_timer returns the Ifm_timer field from the record's packed storage.
func (s *If_msghdr2) Ifm_timer() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[28:32]))
}

// SetIfm_timer updates the Ifm_timer field in the record's packed storage.
func (s *If_msghdr2) SetIfm_timer(v int32) {
	binary.NativeEndian.PutUint32(s.storage[28:32], uint32(v))
}

// Ifm_data returns the Ifm_data field from the record's packed storage.
func (s *If_msghdr2) Ifm_data() If_data64 {
	return *(*If_data64)(unsafe.Pointer(&s.storage[32]))
}

// SetIfm_data updates the Ifm_data field in the record's packed storage.
func (s *If_msghdr2) SetIfm_data(v If_data64) {
	*(*If_data64)(unsafe.Pointer(&s.storage[32])) = v
}

// If_nat64req
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_nat64req
type If_nat64req struct {
	Ifnat64_name     [16]int8
	Ifnat64_prefixes [4]Ipv6_prefix
}

// If_netem_params
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_netem_params
type If_netem_params struct {
	Ifnetem_model            If_netem_model_t
	Ifnetem_bandwidth_bps    uint64
	Ifnetem_latency_ms       uint32
	Ifnetem_jitter_ms        uint32
	Ifnetem_corruption_p     uint32
	Ifnetem_duplication_p    uint32
	Ifnetem_loss_p_gr_gl     uint32
	Ifnetem_loss_p_gr_bl     uint32
	Ifnetem_loss_p_bl_br     uint32
	Ifnetem_loss_p_bl_gr     uint32
	Ifnetem_loss_p_br_bl     uint32
	Ifnetem_loss_recovery_ms uint32
	Ifnetem_reordering_p     uint32
	Ifnetem_reordering_ms    uint32
	Ifnetem_output_ival_ms   uint32
}

// If_netidreq
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_netidreq
type If_netidreq struct {
	Ifnetid_name [16]int8
	Ifnetid_len  U_int8_t
	Ifnetid      [32]U_int8_t
}

// If_netif_stats
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_netif_stats
type If_netif_stats struct {
	Ifn_rx_mit_interval          U_int64_t
	Ifn_rx_mit_mode              U_int32_t
	Ifn_rx_mit_packets_avg       U_int32_t
	Ifn_rx_mit_packets_min       U_int32_t
	Ifn_rx_mit_packets_max       U_int32_t
	Ifn_rx_mit_bytes_avg         U_int32_t
	Ifn_rx_mit_bytes_min         U_int32_t
	Ifn_rx_mit_bytes_max         U_int32_t
	Ifn_rx_mit_cfg_idx           U_int32_t
	Ifn_rx_mit_cfg_packets_lowat U_int32_t
	Ifn_rx_mit_cfg_packets_hiwat U_int32_t
	Ifn_rx_mit_cfg_bytes_lowat   U_int32_t
	Ifn_rx_mit_cfg_bytes_hiwat   U_int32_t
	Ifn_rx_mit_cfg_interval      U_int32_t
}

// If_nexusreq
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_nexusreq
type If_nexusreq struct {
	Ifnr_name       [16]int8
	Ifnr_flags      uint64
	Ifnr_netif      [16]uint8
	Ifnr_flowswitch [16]uint8
	Ifnr_reserved   [5]uint64
}

// If_nsreq
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_nsreq
type If_nsreq struct {
	Ifnsr_name   [16]int8
	Ifnsr_family U_int8_t
	Ifnsr_len    U_int8_t
	Ifnsr_flags  U_int16_t
	Ifnsr_data   [20]U_int8_t
}

// If_order
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_order
type If_order struct {
	Ifo_count           U_int32_t
	Ifo_reserved        U_int32_t
	Ifo_ordered_indices Mach_vm_address_t
}

// If_packet_stats
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_packet_stats
type If_packet_stats struct {
	Ifi_tcp_badformat      U_int64_t
	Ifi_tcp_unspecv6       U_int64_t
	Ifi_tcp_synfin         U_int64_t
	Ifi_tcp_badformatipsec U_int64_t
	Ifi_tcp_noconnnolist   U_int64_t
	Ifi_tcp_noconnlist     U_int64_t
	Ifi_tcp_listbadsyn     U_int64_t
	Ifi_tcp_icmp6unreach   U_int64_t
	Ifi_tcp_deprecate6     U_int64_t
	Ifi_tcp_rstinsynrcv    U_int64_t
	Ifi_tcp_ooopacket      U_int64_t
	Ifi_tcp_dospacket      U_int64_t
	Ifi_tcp_cleanup        U_int64_t
	Ifi_tcp_synwindow      U_int64_t
	Reserved               [4]U_int64_t
	Ifi_udp_port_unreach   U_int64_t
	Ifi_udp_faithprefix    U_int64_t
	Ifi_udp_port0          U_int64_t
	Ifi_udp_badlength      U_int64_t
	Ifi_udp_badchksum      U_int64_t
	Ifi_udp_badmcast       U_int64_t
	Ifi_udp_cleanup        U_int64_t
	Ifi_udp_badipsec       U_int64_t
}

// If_protolistreq
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_protolistreq
type If_protolistreq struct {
	Ifpl_name     [16]int8
	Ifpl_count    U_int32_t
	Ifpl_reserved U_int32_t
	Ifpl_list     *U_int32_t
}

// If_qstatsreq
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_qstatsreq
type If_qstatsreq struct {
	Ifqr_name    [16]int8
	Ifqr_grp_idx U_int32_t
	Ifqr_slot    U_int32_t
	Ifqr_buf     unsafe.Pointer
	Ifqr_len     int32
}

// If_rxpoll_stats
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_rxpoll_stats
type If_rxpoll_stats struct {
	Ifi_poll_off_req       U_int32_t
	Ifi_poll_off_err       U_int32_t
	Ifi_poll_on_req        U_int32_t
	Ifi_poll_on_err        U_int32_t
	Ifi_poll_wakeups_avg   U_int32_t
	Ifi_poll_wakeups_lowat U_int32_t
	Ifi_poll_wakeups_hiwat U_int32_t
	Ifi_poll_packets       U_int64_t
	Ifi_poll_packets_avg   U_int32_t
	Ifi_poll_packets_min   U_int32_t
	Ifi_poll_packets_max   U_int32_t
	Ifi_poll_packets_lowat U_int32_t
	Ifi_poll_packets_hiwat U_int32_t
	Ifi_poll_bytes         U_int64_t
	Ifi_poll_bytes_avg     U_int32_t
	Ifi_poll_bytes_min     U_int32_t
	Ifi_poll_bytes_max     U_int32_t
	Ifi_poll_bytes_lowat   U_int32_t
	Ifi_poll_bytes_hiwat   U_int32_t
	Ifi_poll_packets_limit U_int32_t
	Ifi_poll_interval_time U_int64_t
}

// If_tcp_ecn_perf_stat
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_tcp_ecn_perf_stat
type If_tcp_ecn_perf_stat struct {
	Total_txpkts      U_int64_t
	Total_rxmitpkts   U_int64_t
	Total_rxpkts      U_int64_t
	Total_oopkts      U_int64_t
	Total_reorderpkts U_int64_t
	Rtt_avg           U_int64_t
	Rtt_var           U_int64_t
	Sack_episodes     U_int64_t
	Rxmit_drop        U_int64_t
	Rst_drop          U_int64_t
	Oo_percent        U_int64_t
	Reorder_percent   U_int64_t
	Rxmit_percent     U_int64_t
}

// If_tcp_ecn_stat
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_tcp_ecn_stat
type If_tcp_ecn_stat struct {
	Timestamp             U_int64_t
	Ecn_client_setup      U_int64_t
	Ecn_server_setup      U_int64_t
	Ecn_client_success    U_int64_t
	Ecn_server_success    U_int64_t
	Ecn_peer_nosupport    U_int64_t
	Ecn_syn_lost          U_int64_t
	Ecn_synack_lost       U_int64_t
	Ecn_recv_ce           U_int64_t
	Ecn_recv_ece          U_int64_t
	Ecn_conn_recv_ce      U_int64_t
	Ecn_conn_recv_ece     U_int64_t
	Ecn_conn_plnoce       U_int64_t
	Ecn_conn_plce         U_int64_t
	Ecn_conn_noplce       U_int64_t
	Ecn_fallback_synloss  U_int64_t
	Ecn_fallback_reorder  U_int64_t
	Ecn_fallback_ce       U_int64_t
	Ecn_off_conn          U_int64_t
	Ecn_total_conn        U_int64_t
	Ecn_fallback_droprst  U_int64_t
	Ecn_fallback_droprxmt U_int64_t
	Ecn_fallback_synrst   U_int64_t
	Ecn_on                If_tcp_ecn_perf_stat
	Ecn_off               If_tcp_ecn_perf_stat
}

// If_tdmreq
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_tdmreq
type If_tdmreq struct {
	Iftdm_name  [16]int8
	Iftdm_len   U_int32_t
	Iftdm_table unsafe.Pointer
}

// If_throttlereq
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_throttlereq
type If_throttlereq struct {
	Ifthr_name  [16]int8
	Ifthr_level U_int32_t
}

// If_traffic_class
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_traffic_class
type If_traffic_class struct {
	Ifi_ibepackets U_int64_t
	Ifi_ibebytes   U_int64_t
	Ifi_obepackets U_int64_t
	Ifi_obebytes   U_int64_t
	Ifi_ibkpackets U_int64_t
	Ifi_ibkbytes   U_int64_t
	Ifi_obkpackets U_int64_t
	Ifi_obkbytes   U_int64_t
	Ifi_ivipackets U_int64_t
	Ifi_ivibytes   U_int64_t
	Ifi_ovipackets U_int64_t
	Ifi_ovibytes   U_int64_t
	Ifi_ivopackets U_int64_t
	Ifi_ivobytes   U_int64_t
	Ifi_ovopackets U_int64_t
	Ifi_ovobytes   U_int64_t
	Ifi_ipvpackets U_int64_t
	Ifi_ipvbytes   U_int64_t
	Ifi_opvpackets U_int64_t
	Ifi_opvbytes   U_int64_t
}

// If_wifi_status
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_wifi_status
type If_wifi_status struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [112]byte
}

// If_wifi_u returns the If_wifi_u field from the record's packed storage.
func (s *If_wifi_status) If_wifi_u() [112]byte {
	return *(*[112]byte)(unsafe.Pointer(&s.storage[0]))
}

// SetIf_wifi_u updates the If_wifi_u field in the record's packed storage.
func (s *If_wifi_status) SetIf_wifi_u(v [112]byte) {
	*(*[112]byte)(unsafe.Pointer(&s.storage[0])) = v
}

// If_wifi_status_v1
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_wifi_status_v1
type If_wifi_status_v1 struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [112]byte
}

// Valid_bitmask returns the Valid_bitmask field from the record's packed storage.
func (s *If_wifi_status_v1) Valid_bitmask() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetValid_bitmask updates the Valid_bitmask field in the record's packed storage.
func (s *If_wifi_status_v1) SetValid_bitmask(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// Link_quality_metric returns the Link_quality_metric field from the record's packed storage.
func (s *If_wifi_status_v1) Link_quality_metric() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetLink_quality_metric updates the Link_quality_metric field in the record's packed storage.
func (s *If_wifi_status_v1) SetLink_quality_metric(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// Ul_effective_bandwidth returns the Ul_effective_bandwidth field from the record's packed storage.
func (s *If_wifi_status_v1) Ul_effective_bandwidth() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetUl_effective_bandwidth updates the Ul_effective_bandwidth field in the record's packed storage.
func (s *If_wifi_status_v1) SetUl_effective_bandwidth(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[8:12], uint32(v))
}

// Ul_max_bandwidth returns the Ul_max_bandwidth field from the record's packed storage.
func (s *If_wifi_status_v1) Ul_max_bandwidth() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[12:16]))
}

// SetUl_max_bandwidth updates the Ul_max_bandwidth field in the record's packed storage.
func (s *If_wifi_status_v1) SetUl_max_bandwidth(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[12:16], uint32(v))
}

// Ul_min_latency returns the Ul_min_latency field from the record's packed storage.
func (s *If_wifi_status_v1) Ul_min_latency() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[16:20]))
}

// SetUl_min_latency updates the Ul_min_latency field in the record's packed storage.
func (s *If_wifi_status_v1) SetUl_min_latency(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[16:20], uint32(v))
}

// Ul_effective_latency returns the Ul_effective_latency field from the record's packed storage.
func (s *If_wifi_status_v1) Ul_effective_latency() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[20:24]))
}

// SetUl_effective_latency updates the Ul_effective_latency field in the record's packed storage.
func (s *If_wifi_status_v1) SetUl_effective_latency(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[20:24], uint32(v))
}

// Ul_max_latency returns the Ul_max_latency field from the record's packed storage.
func (s *If_wifi_status_v1) Ul_max_latency() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[24:28]))
}

// SetUl_max_latency updates the Ul_max_latency field in the record's packed storage.
func (s *If_wifi_status_v1) SetUl_max_latency(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[24:28], uint32(v))
}

// Ul_retxt_level returns the Ul_retxt_level field from the record's packed storage.
func (s *If_wifi_status_v1) Ul_retxt_level() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[28:32]))
}

// SetUl_retxt_level updates the Ul_retxt_level field in the record's packed storage.
func (s *If_wifi_status_v1) SetUl_retxt_level(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[28:32], uint32(v))
}

// Ul_bytes_lost returns the Ul_bytes_lost field from the record's packed storage.
func (s *If_wifi_status_v1) Ul_bytes_lost() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[32:36]))
}

// SetUl_bytes_lost updates the Ul_bytes_lost field in the record's packed storage.
func (s *If_wifi_status_v1) SetUl_bytes_lost(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[32:36], uint32(v))
}

// Ul_error_rate returns the Ul_error_rate field from the record's packed storage.
func (s *If_wifi_status_v1) Ul_error_rate() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[36:40]))
}

// SetUl_error_rate updates the Ul_error_rate field in the record's packed storage.
func (s *If_wifi_status_v1) SetUl_error_rate(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[36:40], uint32(v))
}

// Dl_effective_bandwidth returns the Dl_effective_bandwidth field from the record's packed storage.
func (s *If_wifi_status_v1) Dl_effective_bandwidth() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[40:44]))
}

// SetDl_effective_bandwidth updates the Dl_effective_bandwidth field in the record's packed storage.
func (s *If_wifi_status_v1) SetDl_effective_bandwidth(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[40:44], uint32(v))
}

// Dl_max_bandwidth returns the Dl_max_bandwidth field from the record's packed storage.
func (s *If_wifi_status_v1) Dl_max_bandwidth() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[44:48]))
}

// SetDl_max_bandwidth updates the Dl_max_bandwidth field in the record's packed storage.
func (s *If_wifi_status_v1) SetDl_max_bandwidth(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[44:48], uint32(v))
}

// Dl_min_latency returns the Dl_min_latency field from the record's packed storage.
func (s *If_wifi_status_v1) Dl_min_latency() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[48:52]))
}

// SetDl_min_latency updates the Dl_min_latency field in the record's packed storage.
func (s *If_wifi_status_v1) SetDl_min_latency(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[48:52], uint32(v))
}

// Dl_effective_latency returns the Dl_effective_latency field from the record's packed storage.
func (s *If_wifi_status_v1) Dl_effective_latency() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[52:56]))
}

// SetDl_effective_latency updates the Dl_effective_latency field in the record's packed storage.
func (s *If_wifi_status_v1) SetDl_effective_latency(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[52:56], uint32(v))
}

// Dl_max_latency returns the Dl_max_latency field from the record's packed storage.
func (s *If_wifi_status_v1) Dl_max_latency() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[56:60]))
}

// SetDl_max_latency updates the Dl_max_latency field in the record's packed storage.
func (s *If_wifi_status_v1) SetDl_max_latency(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[56:60], uint32(v))
}

// Dl_error_rate returns the Dl_error_rate field from the record's packed storage.
func (s *If_wifi_status_v1) Dl_error_rate() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[60:64]))
}

// SetDl_error_rate updates the Dl_error_rate field in the record's packed storage.
func (s *If_wifi_status_v1) SetDl_error_rate(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[60:64], uint32(v))
}

// Config_frequency returns the Config_frequency field from the record's packed storage.
func (s *If_wifi_status_v1) Config_frequency() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[64:68]))
}

// SetConfig_frequency updates the Config_frequency field in the record's packed storage.
func (s *If_wifi_status_v1) SetConfig_frequency(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[64:68], uint32(v))
}

// Config_multicast_rate returns the Config_multicast_rate field from the record's packed storage.
func (s *If_wifi_status_v1) Config_multicast_rate() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[68:72]))
}

// SetConfig_multicast_rate updates the Config_multicast_rate field in the record's packed storage.
func (s *If_wifi_status_v1) SetConfig_multicast_rate(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[68:72], uint32(v))
}

// Scan_count returns the Scan_count field from the record's packed storage.
func (s *If_wifi_status_v1) Scan_count() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[72:76]))
}

// SetScan_count updates the Scan_count field in the record's packed storage.
func (s *If_wifi_status_v1) SetScan_count(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[72:76], uint32(v))
}

// Scan_duration returns the Scan_duration field from the record's packed storage.
func (s *If_wifi_status_v1) Scan_duration() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[76:80]))
}

// SetScan_duration updates the Scan_duration field in the record's packed storage.
func (s *If_wifi_status_v1) SetScan_duration(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[76:80], uint32(v))
}

// Reserved_1 returns the Reserved_1 field from the record's packed storage.
func (s *If_wifi_status_v1) Reserved_1() U_int64_t {
	return U_int64_t(binary.NativeEndian.Uint64(s.storage[80:88]))
}

// SetReserved_1 updates the Reserved_1 field in the record's packed storage.
func (s *If_wifi_status_v1) SetReserved_1(v U_int64_t) {
	binary.NativeEndian.PutUint64(s.storage[80:88], uint64(v))
}

// Reserved_2 returns the Reserved_2 field from the record's packed storage.
func (s *If_wifi_status_v1) Reserved_2() U_int64_t {
	return U_int64_t(binary.NativeEndian.Uint64(s.storage[88:96]))
}

// SetReserved_2 updates the Reserved_2 field in the record's packed storage.
func (s *If_wifi_status_v1) SetReserved_2(v U_int64_t) {
	binary.NativeEndian.PutUint64(s.storage[88:96], uint64(v))
}

// Reserved_3 returns the Reserved_3 field from the record's packed storage.
func (s *If_wifi_status_v1) Reserved_3() U_int64_t {
	return U_int64_t(binary.NativeEndian.Uint64(s.storage[96:104]))
}

// SetReserved_3 updates the Reserved_3 field in the record's packed storage.
func (s *If_wifi_status_v1) SetReserved_3(v U_int64_t) {
	binary.NativeEndian.PutUint64(s.storage[96:104], uint64(v))
}

// Reserved_4 returns the Reserved_4 field from the record's packed storage.
func (s *If_wifi_status_v1) Reserved_4() U_int64_t {
	return U_int64_t(binary.NativeEndian.Uint64(s.storage[104:112]))
}

// SetReserved_4 updates the Reserved_4 field in the record's packed storage.
func (s *If_wifi_status_v1) SetReserved_4(v U_int64_t) {
	binary.NativeEndian.PutUint64(s.storage[104:112], uint64(v))
}

// Ifa_msghdr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifa_msghdr
type Ifa_msghdr struct {
	Ifam_msglen  uint16
	Ifam_version uint8
	Ifam_type    uint8
	Ifam_addrs   int32
	Ifam_flags   int32
	Ifam_index   uint16
	Ifam_metric  int32
}

// Ifaliasreq
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifaliasreq
type Ifaliasreq struct {
	Ifra_name      [16]int8
	Ifra_addr      [16]byte
	Ifra_broadaddr [16]byte
	Ifra_mask      [16]byte
}

// Ifdevmtu
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifdevmtu
type Ifdevmtu struct {
	Ifdm_current int32
	Ifdm_min     int32
	Ifdm_max     int32
}

// Ifdrv
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifdrv
type Ifdrv struct {
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
	storage [40]byte
}

// Ifd_name returns the Ifd_name field from the record's packed storage.
func (s *Ifdrv) Ifd_name() [16]int8 {
	return *(*[16]int8)(unsafe.Pointer(&s.storage[0]))
}

// SetIfd_name updates the Ifd_name field in the record's packed storage.
func (s *Ifdrv) SetIfd_name(v [16]int8) {
	*(*[16]int8)(unsafe.Pointer(&s.storage[0])) = v
}

// Ifd_cmd returns the Ifd_cmd field from the record's packed storage.
func (s *Ifdrv) Ifd_cmd() uint {
	return uint(binary.NativeEndian.Uint64(s.storage[16:24]))
}

// SetIfd_cmd updates the Ifd_cmd field in the record's packed storage.
func (s *Ifdrv) SetIfd_cmd(v uint) {
	binary.NativeEndian.PutUint64(s.storage[16:24], uint64(v))
}

// Ifd_len returns the Ifd_len field from the record's packed storage.
func (s *Ifdrv) Ifd_len() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[24:32]))
}

// SetIfd_len updates the Ifd_len field in the record's packed storage.
func (s *Ifdrv) SetIfd_len(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[24:32], uint64(v))
}

// Ifd_data returns the Ifd_data field from the record's packed storage.
func (s *Ifdrv) Ifd_data() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[32:40]))
}

// SetIfd_data updates the Ifd_data field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *Ifdrv) SetIfd_data(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[32:40], uint64(v))
}

// Iff_filter
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iff_filter
type Iff_filter struct {
	Iff_cookie   unsafe.Pointer    // A kext defined cookie that will be passed to all filter functions.
	Iff_name     *byte             // A filter name used for debugging purposes.
	Iff_protocol Protocol_family_t // The protocol of the packets this filter is interested in. If you specify zero, packets from all protocols will be passed to the filter.
	Iff_input    unsafe.Pointer    // The filter function to handle inbound packets, may be NULL.
	Iff_output   unsafe.Pointer    // The filter function to handle outbound packets, may be NULL.
	Iff_event    unsafe.Pointer    // The filter function to handle interface events, may be null.
	Iff_ioctl    unsafe.Pointer    // The filter function to handle interface ioctls, may be null.
	Iff_detached unsafe.Pointer    // The filter function used to notify the filter that it has been detached.

}

// Ifkpi
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifkpi
type Ifkpi struct {
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

// Ifk_module_id returns the Ifk_module_id field from the record's packed storage.
func (s *Ifkpi) Ifk_module_id() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetIfk_module_id updates the Ifk_module_id field in the record's packed storage.
func (s *Ifkpi) SetIfk_module_id(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// Ifk_type returns the Ifk_type field from the record's packed storage.
func (s *Ifkpi) Ifk_type() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetIfk_type updates the Ifk_type field in the record's packed storage.
func (s *Ifkpi) SetIfk_type(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// Ifk_data returns the Ifk_data field from the record's packed storage.
func (s *Ifkpi) Ifk_data() [2]uint32 {
	return *(*[2]uint32)(unsafe.Pointer(&s.storage[8]))
}

// SetIfk_data updates the Ifk_data field in the record's packed storage.
func (s *Ifkpi) SetIfk_data(v [2]uint32) {
	*(*[2]uint32)(unsafe.Pointer(&s.storage[8])) = v
}

// Ifma_msghdr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifma_msghdr
type Ifma_msghdr struct {
	Ifmam_msglen  uint16
	Ifmam_version uint8
	Ifmam_type    uint8
	Ifmam_addrs   int32
	Ifmam_flags   int32
	Ifmam_index   uint16
}

// Ifma_msghdr2
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifma_msghdr2
type Ifma_msghdr2 struct {
	Ifmam_msglen   U_short
	Ifmam_version  U_char
	Ifmam_type     U_char
	Ifmam_addrs    int32
	Ifmam_flags    int32
	Ifmam_index    U_short
	Ifmam_refcount int32
}

// Ifmedia_description
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifmedia_description
type Ifmedia_description struct {
	Ifmt_word   int32
	Ifmt_string *byte
}

// Ifmibdata
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifmibdata
type Ifmibdata struct {
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
	storage [180]byte
}

// Ifmd_name returns the Ifmd_name field from the record's packed storage.
func (s *Ifmibdata) Ifmd_name() [16]int8 {
	return *(*[16]int8)(unsafe.Pointer(&s.storage[0]))
}

// SetIfmd_name updates the Ifmd_name field in the record's packed storage.
func (s *Ifmibdata) SetIfmd_name(v [16]int8) {
	*(*[16]int8)(unsafe.Pointer(&s.storage[0])) = v
}

// Ifmd_pcount returns the Ifmd_pcount field from the record's packed storage.
func (s *Ifmibdata) Ifmd_pcount() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[16:20]))
}

// SetIfmd_pcount updates the Ifmd_pcount field in the record's packed storage.
func (s *Ifmibdata) SetIfmd_pcount(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[16:20], uint32(v))
}

// Ifmd_flags returns the Ifmd_flags field from the record's packed storage.
func (s *Ifmibdata) Ifmd_flags() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[20:24]))
}

// SetIfmd_flags updates the Ifmd_flags field in the record's packed storage.
func (s *Ifmibdata) SetIfmd_flags(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[20:24], uint32(v))
}

// Ifmd_snd_len returns the Ifmd_snd_len field from the record's packed storage.
func (s *Ifmibdata) Ifmd_snd_len() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[24:28]))
}

// SetIfmd_snd_len updates the Ifmd_snd_len field in the record's packed storage.
func (s *Ifmibdata) SetIfmd_snd_len(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[24:28], uint32(v))
}

// Ifmd_snd_maxlen returns the Ifmd_snd_maxlen field from the record's packed storage.
func (s *Ifmibdata) Ifmd_snd_maxlen() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[28:32]))
}

// SetIfmd_snd_maxlen updates the Ifmd_snd_maxlen field in the record's packed storage.
func (s *Ifmibdata) SetIfmd_snd_maxlen(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[28:32], uint32(v))
}

// Ifmd_snd_drops returns the Ifmd_snd_drops field from the record's packed storage.
func (s *Ifmibdata) Ifmd_snd_drops() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[32:36]))
}

// SetIfmd_snd_drops updates the Ifmd_snd_drops field in the record's packed storage.
func (s *Ifmibdata) SetIfmd_snd_drops(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[32:36], uint32(v))
}

// Ifmd_filler returns the Ifmd_filler field from the record's packed storage.
func (s *Ifmibdata) Ifmd_filler() [4]uint32 {
	return *(*[4]uint32)(unsafe.Pointer(&s.storage[36]))
}

// SetIfmd_filler updates the Ifmd_filler field in the record's packed storage.
func (s *Ifmibdata) SetIfmd_filler(v [4]uint32) {
	*(*[4]uint32)(unsafe.Pointer(&s.storage[36])) = v
}

// Ifmd_data returns the Ifmd_data field from the record's packed storage.
func (s *Ifmibdata) Ifmd_data() If_data64 {
	return *(*If_data64)(unsafe.Pointer(&s.storage[52]))
}

// SetIfmd_data updates the Ifmd_data field in the record's packed storage.
func (s *Ifmibdata) SetIfmd_data(v If_data64) {
	*(*If_data64)(unsafe.Pointer(&s.storage[52])) = v
}

// Ifmibdata_supplemental
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifmibdata_supplemental
type Ifmibdata_supplemental struct {
	Ifmd_traffic_class If_traffic_class
	Ifmd_data_extended If_data_extended
	Ifmd_packet_stats  If_packet_stats
	Ifmd_rxpoll_stats  If_rxpoll_stats
	Ifmd_netif_stats   If_netif_stats
}

// Ifnet_attach_proto_param
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifnet_attach_proto_param
type Ifnet_attach_proto_param struct {
	Demux_array *Ifnet_demux_desc // An array of ifnet_demux_desc structures describing the protocol.
	Demux_count U_int32_t         // The number of entries in the demux_array array.
	Input       unsafe.Pointer    // The function to be called for inbound packets.
	Pre_output  unsafe.Pointer    // The function to be called for outbound packets.
	Event       unsafe.Pointer    // The function to be called for interface events.
	Ioctl       unsafe.Pointer    // The function to be called for ioctls.
	Detached    unsafe.Pointer    // The function to be called for handling the detach.
	Resolve     unsafe.Pointer
	Send_arp    unsafe.Pointer
}

// Ifnet_attach_proto_param_v2
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifnet_attach_proto_param_v2
type Ifnet_attach_proto_param_v2 struct {
	Demux_array *Ifnet_demux_desc
	Demux_count U_int32_t
	Input       unsafe.Pointer
	Pre_output  unsafe.Pointer
	Event       unsafe.Pointer
	Ioctl       unsafe.Pointer
	Detached    unsafe.Pointer
	Resolve     unsafe.Pointer
	Send_arp    unsafe.Pointer
}

// Ifnet_demux_desc
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifnet_demux_desc
type Ifnet_demux_desc struct {
	Type    U_int32_t      // The type of identifier data (i.e. ETHER_DESC_ETYPE2)
	Data    unsafe.Pointer // A pointer to an entry of type (i.e. pointer to 0x0800).
	Datalen U_int32_t      // The number of bytes of data used to describe the packet.

}

// Ifnet_init_params
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifnet_init_params
type Ifnet_init_params struct {
	Uniqueid       unsafe.Pointer // An identifier unique to this instance of the interface.
	Uniqueid_len   U_int32_t      // The length, in bytes, of the uniqueid.
	Name           *byte          // The interface name (i.e. en).
	Unit           U_int32_t      // The interface unit number (en0's unit number is 0).
	Family         Ifnet_family_t // The interface family.
	Type           U_int32_t      // The interface type (see sys/if_types.h). Must be less than 256. For new types, use IFT_OTHER.
	Output         unsafe.Pointer // The output function for the interface. Every packet the stack attempts to send through this interface will go out through this function.
	Demux          unsafe.Pointer // The function used to determine the protocol family of an incoming packet.
	Add_proto      unsafe.Pointer // The function used to attach a protocol to this interface.
	Del_proto      unsafe.Pointer // The function used to remove a protocol from this interface.
	Check_multi    unsafe.Pointer
	Framer         unsafe.Pointer // The function used to frame outbound packets, may be NULL.
	Softc          unsafe.Pointer // Driver specific storage. This value can be retrieved from the ifnet using the ifnet_softc function.
	Ioctl          unsafe.Pointer // The function used to handle ioctls.
	Set_bpf_tap    unsafe.Pointer // The function used to set the bpf_tap function.
	Detach         unsafe.Pointer // The function called to let the driver know the interface has been detached.
	Event          unsafe.Pointer // The function to notify the interface of various interface specific kernel events.
	Broadcast_addr unsafe.Pointer // The link-layer broadcast address for this interface.
	Broadcast_len  U_int32_t      // The length of the link-layer broadcast address.

}

// Ifnet_interface_advisory
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifnet_interface_advisory
type Ifnet_interface_advisory struct {
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
	_       [0]uint64
	storage [96]byte
}

// Version returns the Version field from the record's packed storage.
func (s *Ifnet_interface_advisory) Version() uint8 {
	return uint8(s.storage[0])
}

// SetVersion updates the Version field in the record's packed storage.
func (s *Ifnet_interface_advisory) SetVersion(v uint8) {
	s.storage[0] = uint8(v)
}

// Direction returns the Direction field from the record's packed storage.
func (s *Ifnet_interface_advisory) Direction() uint8 {
	return uint8(s.storage[1])
}

// SetDirection updates the Direction field in the record's packed storage.
func (s *Ifnet_interface_advisory) SetDirection(v uint8) {
	s.storage[1] = uint8(v)
}

// _reserved returns the _reserved field from the record's packed storage.
func (s *Ifnet_interface_advisory) _reserved() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[2:4]))
}

// Set_reserved updates the _reserved field in the record's packed storage.
func (s *Ifnet_interface_advisory) Set_reserved(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[2:4], uint16(v))
}

// Rate_trend_suggestion returns the Rate_trend_suggestion field from the record's packed storage.
func (s *Ifnet_interface_advisory) Rate_trend_suggestion() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetRate_trend_suggestion updates the Rate_trend_suggestion field in the record's packed storage.
func (s *Ifnet_interface_advisory) SetRate_trend_suggestion(v int32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// Timestamp returns the Timestamp field from the record's packed storage.
func (s *Ifnet_interface_advisory) Timestamp() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[8:16]))
}

// SetTimestamp updates the Timestamp field in the record's packed storage.
func (s *Ifnet_interface_advisory) SetTimestamp(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[8:16], uint64(v))
}

// Max_bandwidth returns the Max_bandwidth field from the record's packed storage.
func (s *Ifnet_interface_advisory) Max_bandwidth() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[16:24]))
}

// SetMax_bandwidth updates the Max_bandwidth field in the record's packed storage.
func (s *Ifnet_interface_advisory) SetMax_bandwidth(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[16:24], uint64(v))
}

// Total_byte_count returns the Total_byte_count field from the record's packed storage.
func (s *Ifnet_interface_advisory) Total_byte_count() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[24:32]))
}

// SetTotal_byte_count updates the Total_byte_count field in the record's packed storage.
func (s *Ifnet_interface_advisory) SetTotal_byte_count(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[24:32], uint64(v))
}

// Average_throughput returns the Average_throughput field from the record's packed storage.
func (s *Ifnet_interface_advisory) Average_throughput() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[32:40]))
}

// SetAverage_throughput updates the Average_throughput field in the record's packed storage.
func (s *Ifnet_interface_advisory) SetAverage_throughput(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[32:40], uint64(v))
}

// Flushable_queue_size returns the Flushable_queue_size field from the record's packed storage.
func (s *Ifnet_interface_advisory) Flushable_queue_size() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[40:44]))
}

// SetFlushable_queue_size updates the Flushable_queue_size field in the record's packed storage.
func (s *Ifnet_interface_advisory) SetFlushable_queue_size(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[40:44], uint32(v))
}

// Non_flushable_queue_size returns the Non_flushable_queue_size field from the record's packed storage.
func (s *Ifnet_interface_advisory) Non_flushable_queue_size() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[44:48]))
}

// SetNon_flushable_queue_size updates the Non_flushable_queue_size field in the record's packed storage.
func (s *Ifnet_interface_advisory) SetNon_flushable_queue_size(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[44:48], uint32(v))
}

// Average_delay returns the Average_delay field from the record's packed storage.
func (s *Ifnet_interface_advisory) Average_delay() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[48:52]))
}

// SetAverage_delay updates the Average_delay field in the record's packed storage.
func (s *Ifnet_interface_advisory) SetAverage_delay(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[48:52], uint32(v))
}

// Frequency_band returns the Frequency_band field from the record's packed storage.
func (s *Ifnet_interface_advisory) Frequency_band() uint8 {
	return uint8(s.storage[52])
}

// SetFrequency_band updates the Frequency_band field in the record's packed storage.
func (s *Ifnet_interface_advisory) SetFrequency_band(v uint8) {
	s.storage[52] = uint8(v)
}

// Intermittent_state returns the Intermittent_state field from the record's packed storage.
func (s *Ifnet_interface_advisory) Intermittent_state() uint8 {
	return uint8(s.storage[53])
}

// SetIntermittent_state updates the Intermittent_state field in the record's packed storage.
func (s *Ifnet_interface_advisory) SetIntermittent_state(v uint8) {
	s.storage[53] = uint8(v)
}

// Estimated_intermittent_period returns the Estimated_intermittent_period field from the record's packed storage.
func (s *Ifnet_interface_advisory) Estimated_intermittent_period() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[54:56]))
}

// SetEstimated_intermittent_period updates the Estimated_intermittent_period field in the record's packed storage.
func (s *Ifnet_interface_advisory) SetEstimated_intermittent_period(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[54:56], uint16(v))
}

// Single_outage_period returns the Single_outage_period field from the record's packed storage.
func (s *Ifnet_interface_advisory) Single_outage_period() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[56:58]))
}

// SetSingle_outage_period updates the Single_outage_period field in the record's packed storage.
func (s *Ifnet_interface_advisory) SetSingle_outage_period(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[56:58], uint16(v))
}

// Bt_coex returns the Bt_coex field from the record's packed storage.
func (s *Ifnet_interface_advisory) Bt_coex() uint8 {
	return uint8(s.storage[58])
}

// SetBt_coex updates the Bt_coex field in the record's packed storage.
func (s *Ifnet_interface_advisory) SetBt_coex(v uint8) {
	s.storage[58] = uint8(v)
}

// Quality_score_delay returns the Quality_score_delay field from the record's packed storage.
func (s *Ifnet_interface_advisory) Quality_score_delay() uint8 {
	return uint8(s.storage[59])
}

// SetQuality_score_delay updates the Quality_score_delay field in the record's packed storage.
func (s *Ifnet_interface_advisory) SetQuality_score_delay(v uint8) {
	s.storage[59] = uint8(v)
}

// Quality_score_loss returns the Quality_score_loss field from the record's packed storage.
func (s *Ifnet_interface_advisory) Quality_score_loss() uint8 {
	return uint8(s.storage[60])
}

// SetQuality_score_loss updates the Quality_score_loss field in the record's packed storage.
func (s *Ifnet_interface_advisory) SetQuality_score_loss(v uint8) {
	s.storage[60] = uint8(v)
}

// Quality_score_channel returns the Quality_score_channel field from the record's packed storage.
func (s *Ifnet_interface_advisory) Quality_score_channel() uint8 {
	return uint8(s.storage[61])
}

// SetQuality_score_channel updates the Quality_score_channel field in the record's packed storage.
func (s *Ifnet_interface_advisory) SetQuality_score_channel(v uint8) {
	s.storage[61] = uint8(v)
}

// Ifnet_interface_advisory_capacity
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifnet_interface_advisory_capacity
type Ifnet_interface_advisory_capacity struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [48]byte
}

// Rate_trend_suggestion returns the Rate_trend_suggestion field from the record's packed storage.
func (s *Ifnet_interface_advisory_capacity) Rate_trend_suggestion() [4]byte {
	return *(*[4]byte)(unsafe.Pointer(&s.storage[0]))
}

// SetRate_trend_suggestion updates the Rate_trend_suggestion field in the record's packed storage.
func (s *Ifnet_interface_advisory_capacity) SetRate_trend_suggestion(v [4]byte) {
	*(*[4]byte)(unsafe.Pointer(&s.storage[0])) = v
}

// Timestamp returns the Timestamp field from the record's packed storage.
func (s *Ifnet_interface_advisory_capacity) Timestamp() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[4:12]))
}

// SetTimestamp updates the Timestamp field in the record's packed storage.
func (s *Ifnet_interface_advisory_capacity) SetTimestamp(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[4:12], uint64(v))
}

// Max_bandwidth returns the Max_bandwidth field from the record's packed storage.
func (s *Ifnet_interface_advisory_capacity) Max_bandwidth() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[12:20]))
}

// SetMax_bandwidth updates the Max_bandwidth field in the record's packed storage.
func (s *Ifnet_interface_advisory_capacity) SetMax_bandwidth(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[12:20], uint64(v))
}

// Total_byte_count returns the Total_byte_count field from the record's packed storage.
func (s *Ifnet_interface_advisory_capacity) Total_byte_count() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[20:28]))
}

// SetTotal_byte_count updates the Total_byte_count field in the record's packed storage.
func (s *Ifnet_interface_advisory_capacity) SetTotal_byte_count(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[20:28], uint64(v))
}

// Average_throughput returns the Average_throughput field from the record's packed storage.
func (s *Ifnet_interface_advisory_capacity) Average_throughput() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[28:36]))
}

// SetAverage_throughput updates the Average_throughput field in the record's packed storage.
func (s *Ifnet_interface_advisory_capacity) SetAverage_throughput(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[28:36], uint64(v))
}

// Flushable_queue_size returns the Flushable_queue_size field from the record's packed storage.
func (s *Ifnet_interface_advisory_capacity) Flushable_queue_size() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[36:40]))
}

// SetFlushable_queue_size updates the Flushable_queue_size field in the record's packed storage.
func (s *Ifnet_interface_advisory_capacity) SetFlushable_queue_size(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[36:40], uint32(v))
}

// Non_flushable_queue_size returns the Non_flushable_queue_size field from the record's packed storage.
func (s *Ifnet_interface_advisory_capacity) Non_flushable_queue_size() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[40:44]))
}

// SetNon_flushable_queue_size updates the Non_flushable_queue_size field in the record's packed storage.
func (s *Ifnet_interface_advisory_capacity) SetNon_flushable_queue_size(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[40:44], uint32(v))
}

// Average_delay returns the Average_delay field from the record's packed storage.
func (s *Ifnet_interface_advisory_capacity) Average_delay() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[44:48]))
}

// SetAverage_delay updates the Average_delay field in the record's packed storage.
func (s *Ifnet_interface_advisory_capacity) SetAverage_delay(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[44:48], uint32(v))
}

// Ifnet_interface_advisory_cell_context
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifnet_interface_advisory_cell_context
type Ifnet_interface_advisory_cell_context struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [16]byte
}

// Radio_access_technology returns the Radio_access_technology field from the record's packed storage.
func (s *Ifnet_interface_advisory_cell_context) Radio_access_technology() uint8 {
	return uint8(s.storage[0])
}

// SetRadio_access_technology updates the Radio_access_technology field in the record's packed storage.
func (s *Ifnet_interface_advisory_cell_context) SetRadio_access_technology(v uint8) {
	s.storage[0] = uint8(v)
}

// Reference_signal_level returns the Reference_signal_level field from the record's packed storage.
func (s *Ifnet_interface_advisory_cell_context) Reference_signal_level() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[1:3]))
}

// SetReference_signal_level updates the Reference_signal_level field in the record's packed storage.
func (s *Ifnet_interface_advisory_cell_context) SetReference_signal_level(v int16) {
	binary.NativeEndian.PutUint16(s.storage[1:3], uint16(v))
}

// Signal_level returns the Signal_level field from the record's packed storage.
func (s *Ifnet_interface_advisory_cell_context) Signal_level() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[3:5]))
}

// SetSignal_level updates the Signal_level field in the record's packed storage.
func (s *Ifnet_interface_advisory_cell_context) SetSignal_level(v int16) {
	binary.NativeEndian.PutUint16(s.storage[3:5], uint16(v))
}

// Signal_quality returns the Signal_quality field from the record's packed storage.
func (s *Ifnet_interface_advisory_cell_context) Signal_quality() int8 {
	return int8(s.storage[5])
}

// SetSignal_quality updates the Signal_quality field in the record's packed storage.
func (s *Ifnet_interface_advisory_cell_context) SetSignal_quality(v int8) {
	s.storage[5] = uint8(v)
}

// Uplink_bler returns the Uplink_bler field from the record's packed storage.
func (s *Ifnet_interface_advisory_cell_context) Uplink_bler() uint8 {
	return uint8(s.storage[6])
}

// SetUplink_bler updates the Uplink_bler field in the record's packed storage.
func (s *Ifnet_interface_advisory_cell_context) SetUplink_bler(v uint8) {
	s.storage[6] = uint8(v)
}

// Downlink_bler returns the Downlink_bler field from the record's packed storage.
func (s *Ifnet_interface_advisory_cell_context) Downlink_bler() uint8 {
	return uint8(s.storage[7])
}

// SetDownlink_bler updates the Downlink_bler field in the record's packed storage.
func (s *Ifnet_interface_advisory_cell_context) SetDownlink_bler(v uint8) {
	s.storage[7] = uint8(v)
}

// Bandwidth_limitation_indication returns the Bandwidth_limitation_indication field from the record's packed storage.
func (s *Ifnet_interface_advisory_cell_context) Bandwidth_limitation_indication() uint8 {
	return uint8(s.storage[8])
}

// SetBandwidth_limitation_indication updates the Bandwidth_limitation_indication field in the record's packed storage.
func (s *Ifnet_interface_advisory_cell_context) SetBandwidth_limitation_indication(v uint8) {
	s.storage[8] = uint8(v)
}

// Cdrx_state returns the Cdrx_state field from the record's packed storage.
func (s *Ifnet_interface_advisory_cell_context) Cdrx_state() uint8 {
	return uint8(s.storage[9])
}

// SetCdrx_state updates the Cdrx_state field in the record's packed storage.
func (s *Ifnet_interface_advisory_cell_context) SetCdrx_state(v uint8) {
	s.storage[9] = uint8(v)
}

// Cdrx_cycle returns the Cdrx_cycle field from the record's packed storage.
func (s *Ifnet_interface_advisory_cell_context) Cdrx_cycle() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[10:12]))
}

// SetCdrx_cycle updates the Cdrx_cycle field in the record's packed storage.
func (s *Ifnet_interface_advisory_cell_context) SetCdrx_cycle(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[10:12], uint16(v))
}

// Estimated_outage_period returns the Estimated_outage_period field from the record's packed storage.
func (s *Ifnet_interface_advisory_cell_context) Estimated_outage_period() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[12:14]))
}

// SetEstimated_outage_period updates the Estimated_outage_period field in the record's packed storage.
func (s *Ifnet_interface_advisory_cell_context) SetEstimated_outage_period(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[12:14], uint16(v))
}

// Outage_state returns the Outage_state field from the record's packed storage.
func (s *Ifnet_interface_advisory_cell_context) Outage_state() uint8 {
	return uint8(s.storage[14])
}

// SetOutage_state updates the Outage_state field in the record's packed storage.
func (s *Ifnet_interface_advisory_cell_context) SetOutage_state(v uint8) {
	s.storage[14] = uint8(v)
}

// __pad returns the __pad field from the record's packed storage.
func (s *Ifnet_interface_advisory_cell_context) __pad() uint8 {
	return uint8(s.storage[15])
}

// Set__pad updates the __pad field in the record's packed storage.
func (s *Ifnet_interface_advisory_cell_context) Set__pad(v uint8) {
	s.storage[15] = uint8(v)
}

// Ifnet_interface_advisory_header
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifnet_interface_advisory_header
type Ifnet_interface_advisory_header struct {
	Version           uint8
	Direction         IfInterfaceAdvisoryDirection
	Interface_type    IfInterfaceAdvisoryInterfaceType
	Notification_type [1]byte
}

// Ifnet_interface_advisory_wifi_context
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifnet_interface_advisory_wifi_context
type Ifnet_interface_advisory_wifi_context struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [37]byte
}

// Frequency_band returns the Frequency_band field from the record's packed storage.
func (s *Ifnet_interface_advisory_wifi_context) Frequency_band() IfInterfaceAdvisoryFreqBand {
	return *(*IfInterfaceAdvisoryFreqBand)(unsafe.Pointer(&s.storage[0]))
}

// SetFrequency_band updates the Frequency_band field in the record's packed storage.
func (s *Ifnet_interface_advisory_wifi_context) SetFrequency_band(v IfInterfaceAdvisoryFreqBand) {
	*(*IfInterfaceAdvisoryFreqBand)(unsafe.Pointer(&s.storage[0])) = v
}

// Intermittent_state returns the Intermittent_state field from the record's packed storage.
func (s *Ifnet_interface_advisory_wifi_context) Intermittent_state() uint8 {
	return uint8(s.storage[1])
}

// SetIntermittent_state updates the Intermittent_state field in the record's packed storage.
func (s *Ifnet_interface_advisory_wifi_context) SetIntermittent_state(v uint8) {
	s.storage[1] = uint8(v)
}

// Estimated_intermittent_period returns the Estimated_intermittent_period field from the record's packed storage.
func (s *Ifnet_interface_advisory_wifi_context) Estimated_intermittent_period() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[2:4]))
}

// SetEstimated_intermittent_period updates the Estimated_intermittent_period field in the record's packed storage.
func (s *Ifnet_interface_advisory_wifi_context) SetEstimated_intermittent_period(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[2:4], uint16(v))
}

// Single_outage_period returns the Single_outage_period field from the record's packed storage.
func (s *Ifnet_interface_advisory_wifi_context) Single_outage_period() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[4:6]))
}

// SetSingle_outage_period updates the Single_outage_period field in the record's packed storage.
func (s *Ifnet_interface_advisory_wifi_context) SetSingle_outage_period(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[4:6], uint16(v))
}

// Bt_coex returns the Bt_coex field from the record's packed storage.
func (s *Ifnet_interface_advisory_wifi_context) Bt_coex() uint8 {
	return uint8(s.storage[6])
}

// SetBt_coex updates the Bt_coex field in the record's packed storage.
func (s *Ifnet_interface_advisory_wifi_context) SetBt_coex(v uint8) {
	s.storage[6] = uint8(v)
}

// Quality_score_delay returns the Quality_score_delay field from the record's packed storage.
func (s *Ifnet_interface_advisory_wifi_context) Quality_score_delay() uint8 {
	return uint8(s.storage[7])
}

// SetQuality_score_delay updates the Quality_score_delay field in the record's packed storage.
func (s *Ifnet_interface_advisory_wifi_context) SetQuality_score_delay(v uint8) {
	s.storage[7] = uint8(v)
}

// Quality_score_loss returns the Quality_score_loss field from the record's packed storage.
func (s *Ifnet_interface_advisory_wifi_context) Quality_score_loss() uint8 {
	return uint8(s.storage[8])
}

// SetQuality_score_loss updates the Quality_score_loss field in the record's packed storage.
func (s *Ifnet_interface_advisory_wifi_context) SetQuality_score_loss(v uint8) {
	s.storage[8] = uint8(v)
}

// Quality_score_channel returns the Quality_score_channel field from the record's packed storage.
func (s *Ifnet_interface_advisory_wifi_context) Quality_score_channel() uint8 {
	return uint8(s.storage[9])
}

// SetQuality_score_channel updates the Quality_score_channel field in the record's packed storage.
func (s *Ifnet_interface_advisory_wifi_context) SetQuality_score_channel(v uint8) {
	s.storage[9] = uint8(v)
}

// Radio_coex returns the Radio_coex field from the record's packed storage.
func (s *Ifnet_interface_advisory_wifi_context) Radio_coex() uint8 {
	return uint8(s.storage[10])
}

// SetRadio_coex updates the Radio_coex field in the record's packed storage.
func (s *Ifnet_interface_advisory_wifi_context) SetRadio_coex(v uint8) {
	s.storage[10] = uint8(v)
}

// Wlan_duty_cycle returns the Wlan_duty_cycle field from the record's packed storage.
func (s *Ifnet_interface_advisory_wifi_context) Wlan_duty_cycle() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[11:13]))
}

// SetWlan_duty_cycle updates the Wlan_duty_cycle field in the record's packed storage.
func (s *Ifnet_interface_advisory_wifi_context) SetWlan_duty_cycle(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[11:13], uint16(v))
}

// Wifi_observed_tx_bitrate returns the Wifi_observed_tx_bitrate field from the record's packed storage.
func (s *Ifnet_interface_advisory_wifi_context) Wifi_observed_tx_bitrate() [6]uint32 {
	return *(*[6]uint32)(unsafe.Pointer(&s.storage[13]))
}

// SetWifi_observed_tx_bitrate updates the Wifi_observed_tx_bitrate field in the record's packed storage.
func (s *Ifnet_interface_advisory_wifi_context) SetWifi_observed_tx_bitrate(v [6]uint32) {
	*(*[6]uint32)(unsafe.Pointer(&s.storage[13])) = v
}

// Ifnet_ip_addr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifnet_ip_addr
type Ifnet_ip_addr struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [16]byte
}

// Addr8 returns the Addr8 field from the record's packed storage.
func (s *Ifnet_ip_addr) Addr8() [16]uint8 {
	return *(*[16]uint8)(unsafe.Pointer(&s.storage[0]))
}

// SetAddr8 updates the Addr8 field in the record's packed storage.
func (s *Ifnet_ip_addr) SetAddr8(v [16]uint8) {
	*(*[16]uint8)(unsafe.Pointer(&s.storage[0])) = v
}

// Ifnet_stat_increment_param
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifnet_stat_increment_param
type Ifnet_stat_increment_param struct {
	Packets_in  U_int32_t // The number of packets received.
	Bytes_in    U_int32_t // The number of bytes received.
	Errors_in   U_int32_t // The number of receive errors.
	Packets_out U_int32_t // The number of packets transmitted.
	Bytes_out   U_int32_t // The number of bytes transmitted.
	Errors_out  U_int32_t // The number of transmission errors.
	Collisions  U_int32_t // The number of collisions seen by this interface.
	Dropped     U_int32_t // The number of packets dropped.

}

// Ifnet_stats_param
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifnet_stats_param
type Ifnet_stats_param struct {
	Packets_in     U_int64_t // The number of packets received.
	Bytes_in       U_int64_t // The number of bytes received.
	Multicasts_in  U_int64_t
	Errors_in      U_int64_t // The number of receive errors.
	Packets_out    U_int64_t // The number of packets transmitted.
	Bytes_out      U_int64_t // The number of bytes transmitted.
	Multicasts_out U_int64_t
	Errors_out     U_int64_t // The number of transmission errors.
	Collisions     U_int64_t // The number of collisions seen by this interface.
	Dropped        U_int64_t // The number of packets dropped.
	No_protocol    U_int64_t
}

// Ifnet_stats_per_flow
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifnet_stats_per_flow
type Ifnet_stats_per_flow struct {
	Bk_txpackets           U_int64_t
	Txpackets              U_int64_t
	Rxpackets              U_int64_t
	Txretransmitbytes      U_int32_t
	Rxoutoforderbytes      U_int32_t
	Rxmitpkts              U_int32_t
	Rcvoopack              U_int32_t
	Pawsdrop               U_int32_t
	Sack_recovery_episodes U_int32_t
	Reordered_pkts         U_int32_t
	Dsack_sent             U_int32_t
	Dsack_recvd            U_int32_t
	Srtt                   U_int32_t
	Rttupdated             U_int32_t
	Rttvar                 U_int32_t
	Rttmin                 U_int32_t
	Bw_sndbw_max           U_int32_t
	Bw_rcvbw_max           U_int32_t
	Ecn_recv_ece           U_int32_t
	Ecn_recv_ce            U_int32_t
	Ecn_flags              U_int32_t
	bitfield21             uint16
	_reserved_16           U_int16_t
	_reserved_32           U_int32_t
	Linkheur_noackpri      U_int64_t
	Linkheur_comprxmt      U_int64_t
	Linkheur_synrxmt       U_int64_t
	Linkheur_rxmtfloor     U_int64_t
}

// Ipv4 returns the Ipv4 bitfield.
func (s *Ifnet_stats_per_flow) Ipv4() uint16 {
	return (s.bitfield21 >> 0) & ((1 << 1) - 1)
}

// SetIpv4 updates the Ipv4 bitfield.
func (s *Ifnet_stats_per_flow) SetIpv4(v uint16) {
	const mask uint16 = (1 << 1) - 1
	s.bitfield21 = (s.bitfield21 &^ (mask << 0)) | ((v & mask) << 0)
}

// Local returns the Local bitfield.
func (s *Ifnet_stats_per_flow) Local() uint16 {
	return (s.bitfield21 >> 1) & ((1 << 1) - 1)
}

// SetLocal updates the Local bitfield.
func (s *Ifnet_stats_per_flow) SetLocal(v uint16) {
	const mask uint16 = (1 << 1) - 1
	s.bitfield21 = (s.bitfield21 &^ (mask << 1)) | ((v & mask) << 1)
}

// Connreset returns the Connreset bitfield.
func (s *Ifnet_stats_per_flow) Connreset() uint16 {
	return (s.bitfield21 >> 2) & ((1 << 1) - 1)
}

// SetConnreset updates the Connreset bitfield.
func (s *Ifnet_stats_per_flow) SetConnreset(v uint16) {
	const mask uint16 = (1 << 1) - 1
	s.bitfield21 = (s.bitfield21 &^ (mask << 2)) | ((v & mask) << 2)
}

// Conntimeout returns the Conntimeout bitfield.
func (s *Ifnet_stats_per_flow) Conntimeout() uint16 {
	return (s.bitfield21 >> 3) & ((1 << 1) - 1)
}

// SetConntimeout updates the Conntimeout bitfield.
func (s *Ifnet_stats_per_flow) SetConntimeout(v uint16) {
	const mask uint16 = (1 << 1) - 1
	s.bitfield21 = (s.bitfield21 &^ (mask << 3)) | ((v & mask) << 3)
}

// Rxmit_drop returns the Rxmit_drop bitfield.
func (s *Ifnet_stats_per_flow) Rxmit_drop() uint16 {
	return (s.bitfield21 >> 4) & ((1 << 1) - 1)
}

// SetRxmit_drop updates the Rxmit_drop bitfield.
func (s *Ifnet_stats_per_flow) SetRxmit_drop(v uint16) {
	const mask uint16 = (1 << 1) - 1
	s.bitfield21 = (s.bitfield21 &^ (mask << 4)) | ((v & mask) << 4)
}

// Ecn_fallback_synloss returns the Ecn_fallback_synloss bitfield.
func (s *Ifnet_stats_per_flow) Ecn_fallback_synloss() uint16 {
	return (s.bitfield21 >> 5) & ((1 << 1) - 1)
}

// SetEcn_fallback_synloss updates the Ecn_fallback_synloss bitfield.
func (s *Ifnet_stats_per_flow) SetEcn_fallback_synloss(v uint16) {
	const mask uint16 = (1 << 1) - 1
	s.bitfield21 = (s.bitfield21 &^ (mask << 5)) | ((v & mask) << 5)
}

// Ecn_fallback_droprst returns the Ecn_fallback_droprst bitfield.
func (s *Ifnet_stats_per_flow) Ecn_fallback_droprst() uint16 {
	return (s.bitfield21 >> 6) & ((1 << 1) - 1)
}

// SetEcn_fallback_droprst updates the Ecn_fallback_droprst bitfield.
func (s *Ifnet_stats_per_flow) SetEcn_fallback_droprst(v uint16) {
	const mask uint16 = (1 << 1) - 1
	s.bitfield21 = (s.bitfield21 &^ (mask << 6)) | ((v & mask) << 6)
}

// Ecn_fallback_droprxmt returns the Ecn_fallback_droprxmt bitfield.
func (s *Ifnet_stats_per_flow) Ecn_fallback_droprxmt() uint16 {
	return (s.bitfield21 >> 7) & ((1 << 1) - 1)
}

// SetEcn_fallback_droprxmt updates the Ecn_fallback_droprxmt bitfield.
func (s *Ifnet_stats_per_flow) SetEcn_fallback_droprxmt(v uint16) {
	const mask uint16 = (1 << 1) - 1
	s.bitfield21 = (s.bitfield21 &^ (mask << 7)) | ((v & mask) << 7)
}

// Ecn_fallback_ce returns the Ecn_fallback_ce bitfield.
func (s *Ifnet_stats_per_flow) Ecn_fallback_ce() uint16 {
	return (s.bitfield21 >> 8) & ((1 << 1) - 1)
}

// SetEcn_fallback_ce updates the Ecn_fallback_ce bitfield.
func (s *Ifnet_stats_per_flow) SetEcn_fallback_ce(v uint16) {
	const mask uint16 = (1 << 1) - 1
	s.bitfield21 = (s.bitfield21 &^ (mask << 8)) | ((v & mask) << 8)
}

// Ecn_fallback_reorder returns the Ecn_fallback_reorder bitfield.
func (s *Ifnet_stats_per_flow) Ecn_fallback_reorder() uint16 {
	return (s.bitfield21 >> 9) & ((1 << 1) - 1)
}

// SetEcn_fallback_reorder updates the Ecn_fallback_reorder bitfield.
func (s *Ifnet_stats_per_flow) SetEcn_fallback_reorder(v uint16) {
	const mask uint16 = (1 << 1) - 1
	s.bitfield21 = (s.bitfield21 &^ (mask << 9)) | ((v & mask) << 9)
}

// _reserved_6 returns the _reserved_6 bitfield.
func (s *Ifnet_stats_per_flow) _reserved_6() uint16 {
	return (s.bitfield21 >> 10) & ((1 << 6) - 1)
}

// Set_reserved_6 updates the _reserved_6 bitfield.
func (s *Ifnet_stats_per_flow) Set_reserved_6(v uint16) {
	const mask uint16 = (1 << 6) - 1
	s.bitfield21 = (s.bitfield21 &^ (mask << 10)) | ((v & mask) << 10)
}

// Ifnet_traffic_descriptor_common
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifnet_traffic_descriptor_common
type Ifnet_traffic_descriptor_common struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [8]byte
}

// Itd_type returns the Itd_type field from the record's packed storage.
func (s *Ifnet_traffic_descriptor_common) Itd_type() uint8 {
	return uint8(s.storage[0])
}

// SetItd_type updates the Itd_type field in the record's packed storage.
func (s *Ifnet_traffic_descriptor_common) SetItd_type(v uint8) {
	s.storage[0] = uint8(v)
}

// _reserved returns the _reserved field from the record's packed storage.
func (s *Ifnet_traffic_descriptor_common) _reserved() uint8 {
	return uint8(s.storage[1])
}

// Set_reserved updates the _reserved field in the record's packed storage.
func (s *Ifnet_traffic_descriptor_common) Set_reserved(v uint8) {
	s.storage[1] = uint8(v)
}

// Itd_len returns the Itd_len field from the record's packed storage.
func (s *Ifnet_traffic_descriptor_common) Itd_len() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[2:4]))
}

// SetItd_len updates the Itd_len field in the record's packed storage.
func (s *Ifnet_traffic_descriptor_common) SetItd_len(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[2:4], uint16(v))
}

// Itd_flags returns the Itd_flags field from the record's packed storage.
func (s *Ifnet_traffic_descriptor_common) Itd_flags() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetItd_flags updates the Itd_flags field in the record's packed storage.
func (s *Ifnet_traffic_descriptor_common) SetItd_flags(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// Ifnet_traffic_descriptor_inet
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifnet_traffic_descriptor_inet
type Ifnet_traffic_descriptor_inet struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [48]byte
}

// Inet_common returns the Inet_common field from the record's packed storage.
func (s *Ifnet_traffic_descriptor_inet) Inet_common() Ifnet_traffic_descriptor_common {
	return *(*Ifnet_traffic_descriptor_common)(unsafe.Pointer(&s.storage[0]))
}

// SetInet_common updates the Inet_common field in the record's packed storage.
func (s *Ifnet_traffic_descriptor_inet) SetInet_common(v Ifnet_traffic_descriptor_common) {
	*(*Ifnet_traffic_descriptor_common)(unsafe.Pointer(&s.storage[0])) = v
}

// Inet_mask returns the Inet_mask field from the record's packed storage.
func (s *Ifnet_traffic_descriptor_inet) Inet_mask() uint8 {
	return uint8(s.storage[8])
}

// SetInet_mask updates the Inet_mask field in the record's packed storage.
func (s *Ifnet_traffic_descriptor_inet) SetInet_mask(v uint8) {
	s.storage[8] = uint8(v)
}

// Inet_ipver returns the Inet_ipver field from the record's packed storage.
func (s *Ifnet_traffic_descriptor_inet) Inet_ipver() uint8 {
	return uint8(s.storage[9])
}

// SetInet_ipver updates the Inet_ipver field in the record's packed storage.
func (s *Ifnet_traffic_descriptor_inet) SetInet_ipver(v uint8) {
	s.storage[9] = uint8(v)
}

// Inet_proto returns the Inet_proto field from the record's packed storage.
func (s *Ifnet_traffic_descriptor_inet) Inet_proto() uint8 {
	return uint8(s.storage[10])
}

// SetInet_proto updates the Inet_proto field in the record's packed storage.
func (s *Ifnet_traffic_descriptor_inet) SetInet_proto(v uint8) {
	s.storage[10] = uint8(v)
}

// _reserved returns the _reserved field from the record's packed storage.
func (s *Ifnet_traffic_descriptor_inet) _reserved() uint8 {
	return uint8(s.storage[11])
}

// Set_reserved updates the _reserved field in the record's packed storage.
func (s *Ifnet_traffic_descriptor_inet) Set_reserved(v uint8) {
	s.storage[11] = uint8(v)
}

// Inet_laddr returns the Inet_laddr field from the record's packed storage.
func (s *Ifnet_traffic_descriptor_inet) Inet_laddr() Ifnet_ip_addr {
	return *(*Ifnet_ip_addr)(unsafe.Pointer(&s.storage[12]))
}

// SetInet_laddr updates the Inet_laddr field in the record's packed storage.
func (s *Ifnet_traffic_descriptor_inet) SetInet_laddr(v Ifnet_ip_addr) {
	*(*Ifnet_ip_addr)(unsafe.Pointer(&s.storage[12])) = v
}

// Inet_raddr returns the Inet_raddr field from the record's packed storage.
func (s *Ifnet_traffic_descriptor_inet) Inet_raddr() Ifnet_ip_addr {
	return *(*Ifnet_ip_addr)(unsafe.Pointer(&s.storage[28]))
}

// SetInet_raddr updates the Inet_raddr field in the record's packed storage.
func (s *Ifnet_traffic_descriptor_inet) SetInet_raddr(v Ifnet_ip_addr) {
	*(*Ifnet_ip_addr)(unsafe.Pointer(&s.storage[28])) = v
}

// Inet_lport returns the Inet_lport field from the record's packed storage.
func (s *Ifnet_traffic_descriptor_inet) Inet_lport() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[44:46]))
}

// SetInet_lport updates the Inet_lport field in the record's packed storage.
func (s *Ifnet_traffic_descriptor_inet) SetInet_lport(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[44:46], uint16(v))
}

// Inet_rport returns the Inet_rport field from the record's packed storage.
func (s *Ifnet_traffic_descriptor_inet) Inet_rport() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[46:48]))
}

// SetInet_rport updates the Inet_rport field in the record's packed storage.
func (s *Ifnet_traffic_descriptor_inet) SetInet_rport(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[46:48], uint16(v))
}

// Ifnet_traffic_rule_action
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifnet_traffic_rule_action
type Ifnet_traffic_rule_action struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [4]byte
}

// Ra_type returns the Ra_type field from the record's packed storage.
func (s *Ifnet_traffic_rule_action) Ra_type() uint8 {
	return uint8(s.storage[0])
}

// SetRa_type updates the Ra_type field in the record's packed storage.
func (s *Ifnet_traffic_rule_action) SetRa_type(v uint8) {
	s.storage[0] = uint8(v)
}

// _reserved returns the _reserved field from the record's packed storage.
func (s *Ifnet_traffic_rule_action) _reserved() uint8 {
	return uint8(s.storage[1])
}

// Set_reserved updates the _reserved field in the record's packed storage.
func (s *Ifnet_traffic_rule_action) Set_reserved(v uint8) {
	s.storage[1] = uint8(v)
}

// Ra_len returns the Ra_len field from the record's packed storage.
func (s *Ifnet_traffic_rule_action) Ra_len() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[2:4]))
}

// SetRa_len updates the Ra_len field in the record's packed storage.
func (s *Ifnet_traffic_rule_action) SetRa_len(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[2:4], uint16(v))
}

// Ifnet_traffic_rule_action_steer
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifnet_traffic_rule_action_steer
type Ifnet_traffic_rule_action_steer struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [12]byte
}

// Ras_common returns the Ras_common field from the record's packed storage.
func (s *Ifnet_traffic_rule_action_steer) Ras_common() Ifnet_traffic_rule_action {
	return *(*Ifnet_traffic_rule_action)(unsafe.Pointer(&s.storage[0]))
}

// SetRas_common updates the Ras_common field in the record's packed storage.
func (s *Ifnet_traffic_rule_action_steer) SetRas_common(v Ifnet_traffic_rule_action) {
	*(*Ifnet_traffic_rule_action)(unsafe.Pointer(&s.storage[0])) = v
}

// Ras_qset_id returns the Ras_qset_id field from the record's packed storage.
func (s *Ifnet_traffic_rule_action_steer) Ras_qset_id() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[4:12]))
}

// SetRas_qset_id updates the Ras_qset_id field in the record's packed storage.
func (s *Ifnet_traffic_rule_action_steer) SetRas_qset_id(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[4:12], uint64(v))
}

// Ifqueue
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifqueue
type Ifqueue struct {
	Ifq_head   unsafe.Pointer
	Ifq_tail   unsafe.Pointer
	Ifq_len    int32
	Ifq_maxlen int32
	Ifq_drops  int32
}

// Ifreq
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifreq
type Ifreq struct {
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
	_       [0]uint64
	storage [32]byte
}

// Ifr_name returns the Ifr_name field from the record's packed storage.
func (s *Ifreq) Ifr_name() [16]int8 {
	return *(*[16]int8)(unsafe.Pointer(&s.storage[0]))
}

// SetIfr_name updates the Ifr_name field in the record's packed storage.
func (s *Ifreq) SetIfr_name(v [16]int8) {
	*(*[16]int8)(unsafe.Pointer(&s.storage[0])) = v
}

// Ifr_ifru returns the Ifr_ifru field from the record's packed storage.
func (s *Ifreq) Ifr_ifru() [2]uint64 {
	return *(*[2]uint64)(unsafe.Pointer(&s.storage[16]))
}

// SetIfr_ifru updates the Ifr_ifru field in the record's packed storage.
func (s *Ifreq) SetIfr_ifru(v [2]uint64) {
	*(*[2]uint64)(unsafe.Pointer(&s.storage[16])) = v
}

// Ifs_iso_8802_3
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifs_iso_8802_3
type Ifs_iso_8802_3 struct {
	Dot3StatsAlignmentErrors           U_int32_t
	Dot3StatsFCSErrors                 U_int32_t
	Dot3StatsSingleCollisionFrames     U_int32_t
	Dot3StatsMultipleCollisionFrames   U_int32_t
	Dot3StatsSQETestErrors             U_int32_t
	Dot3StatsDeferredTransmissions     U_int32_t
	Dot3StatsLateCollisions            U_int32_t
	Dot3StatsExcessiveCollisions       U_int32_t
	Dot3StatsInternalMacTransmitErrors U_int32_t
	Dot3StatsCarrierSenseErrors        U_int32_t
	Dot3StatsFrameTooLongs             U_int32_t
	Dot3StatsInternalMacReceiveErrors  U_int32_t
	Dot3StatsEtherChipSet              U_int32_t
	Dot3StatsMissedFrames              U_int32_t
	Dot3StatsCollFrequencies           [16]U_int32_t
	Dot3Compliance                     U_int32_t
}

// Ifstat
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifstat
type Ifstat struct {
	Ifs_name [16]int8
	Ascii    [801]int8
}

// Igmp
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/igmp
type Igmp struct {
	Igmp_type  U_char
	Igmp_code  U_char
	Igmp_cksum U_short
	Igmp_group In_addr
}

// Igmp_grouprec
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/igmp_grouprec
type Igmp_grouprec struct {
	Ig_type    U_char
	Ig_datalen U_char
	Ig_numsrc  U_short
	Ig_group   In_addr
}

// Igmp_report
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/igmp_report
type Igmp_report struct {
	Ir_type    U_char
	Ir_rsv1    U_char
	Ir_cksum   U_short
	Ir_rsv2    U_short
	Ir_numgrps U_short
}

// Igmpstat
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/igmpstat
type Igmpstat struct {
	Igps_rcv_total      U_int
	Igps_rcv_tooshort   U_int
	Igps_rcv_badsum     U_int
	Igps_rcv_queries    U_int
	Igps_rcv_badqueries U_int
	Igps_rcv_reports    U_int
	Igps_rcv_badreports U_int
	Igps_rcv_ourreports U_int
	Igps_snd_reports    U_int
}

// Igmpstat_v3
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/igmpstat_v3
type Igmpstat_v3 struct {
	Igps_version           uint32
	Igps_len               uint32
	Igps_rcv_total         uint64
	Igps_rcv_tooshort      uint64
	Igps_rcv_badttl        uint64
	Igps_rcv_badsum        uint64
	Igps_rcv_v1v2_queries  uint64
	Igps_rcv_v3_queries    uint64
	Igps_rcv_badqueries    uint64
	Igps_rcv_gen_queries   uint64
	Igps_rcv_group_queries uint64
	Igps_rcv_gsr_queries   uint64
	Igps_drop_gsr_queries  uint64
	Igps_rcv_reports       uint64
	Igps_rcv_badreports    uint64
	Igps_rcv_ourreports    uint64
	Igps_rcv_nora          uint64
	Igps_snd_reports       uint64
	__igps_pad             [4]uint64
}

// Igmpv3
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/igmpv3
type Igmpv3 struct {
	Igmp_type   U_char
	Igmp_code   U_char
	Igmp_cksum  U_short
	Igmp_group  In_addr
	Igmp_misc   U_char
	Igmp_qqi    U_char
	Igmp_numsrc U_short
}

// Image_params
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/image_params
type Image_params struct {
	Ip_user_fname                 User_addr_t
	Ip_user_argv                  User_addr_t
	Ip_user_envv                  User_addr_t
	Ip_seg                        int32
	Ip_vp                         unsafe.Pointer
	Ip_vattr                      *Vnode_attr
	Ip_origvattr                  *Vnode_attr
	Ip_origcputype                int32
	Ip_origcpusubtype             int32
	Ip_vdata                      *byte
	Ip_flags                      int32
	Ip_argc                       int32
	Ip_envc                       int32
	Ip_applec                     int32
	Ip_startargv                  *byte
	Ip_endargv                    *byte
	Ip_endenvv                    *byte
	Ip_strings                    *byte
	Ip_strendp                    *byte
	Ip_subsystem_root_path        *byte
	Ip_argspace                   int32
	Ip_strspace                   int32
	Ip_arch_offset                User_size_t
	Ip_arch_size                  User_size_t
	Ip_interp_buffer              [512]int8
	Ip_interp_sugid_fd            int32
	Ip_vfs_context                unsafe.Pointer
	Ip_ndp                        unsafe.Pointer
	Ip_new_thread                 Thread_t
	Ip_execlabelp                 unsafe.Pointer
	Ip_scriptlabelp               unsafe.Pointer
	Ip_scriptvp                   unsafe.Pointer
	Ip_csflags                    uint32
	Ip_mac_return                 int32
	Ip_px_sa                      unsafe.Pointer
	Ip_px_sfa                     unsafe.Pointer
	Ip_px_spa                     unsafe.Pointer
	Ip_free_map                   uint32
	Array                         unsafe.Pointer
	Data                          unsafe.Pointer
	Datalen                       uint64
	Ip_px_persona                 unsafe.Pointer
	Ip_px_pcred_info              unsafe.Pointer
	Ip_cs_error                   unsafe.Pointer
	Ip_inherited_shared_region_id *byte
	Ip_dyld_fsid                  uint64
	Ip_dyld_fsobjid               uint64
	Ip_inherited_jop_pid          uint64
	Ip_flags2                     uint32
	Ip_px_smpx                    [3]uint64
	Ip_simulator_binary           uint32
}

// In6_addrlifetime
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/in6_addrlifetime
type In6_addrlifetime struct {
	Ia6t_expire    int64
	Ia6t_preferred int64
	Ia6t_vltime    U_int32_t
	Ia6t_pltime    U_int32_t
}

// In6_addrpolicy
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/in6_addrpolicy
type In6_addrpolicy struct {
	Addr     Sockaddr_in6
	Addrmask Sockaddr_in6
	Preced   int32
	Label    int32
	Use      U_quad_t
}

// In6_aliasreq
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/in6_aliasreq
type In6_aliasreq struct {
	Ifra_name       [16]int8
	Ifra_addr       Sockaddr_in6
	Ifra_broadaddr  Sockaddr_in6
	Ifra_prefixmask Sockaddr_in6
	Ifra_flags      int32
	Ifra_lifetime   In6_addrlifetime
	Ifra_dstaddr    Sockaddr_in6
}

// In6_ifreq
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/in6_ifreq
type In6_ifreq struct {
	Ifr_name [16]int8
	Ifr_ifru [34]uint64
}

// In6_ifstat
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/in6_ifstat
type In6_ifstat struct {
	Ifs6_in_receive        U_quad_t
	Ifs6_in_hdrerr         U_quad_t
	Ifs6_in_toobig         U_quad_t
	Ifs6_in_noroute        U_quad_t
	Ifs6_in_addrerr        U_quad_t
	Ifs6_in_protounknown   U_quad_t
	Ifs6_in_truncated      U_quad_t
	Ifs6_in_discard        U_quad_t
	Ifs6_in_deliver        U_quad_t
	Ifs6_out_forward       U_quad_t
	Ifs6_out_request       U_quad_t
	Ifs6_out_discard       U_quad_t
	Ifs6_out_fragok        U_quad_t
	Ifs6_out_fragfail      U_quad_t
	Ifs6_out_fragcreat     U_quad_t
	Ifs6_reass_reqd        U_quad_t
	Ifs6_reass_ok          U_quad_t
	Ifs6_atmfrag_rcvd      U_quad_t
	Ifs6_reass_fail        U_quad_t
	Ifs6_in_mcast          U_quad_t
	Ifs6_out_mcast         U_quad_t
	Ifs6_cantfoward_icmp6  U_quad_t
	Ifs6_addr_expiry_cnt   U_quad_t
	Ifs6_pfx_expiry_cnt    U_quad_t
	Ifs6_defrtr_expiry_cnt U_quad_t
}

// In6_pktinfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/in6_pktinfo
type In6_pktinfo struct {
	Ipi6_addr    [4]uint32
	Ipi6_ifindex uint32
}

// In6_prefixreq
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/in6_prefixreq
type In6_prefixreq struct {
	Ipr_name   [16]int8
	Ipr_origin U_char
	Ipr_plen   U_char
	Ipr_vltime U_int32_t
	Ipr_pltime U_int32_t
	Ipr_flags  In6_prflags
	Ipr_prefix Sockaddr_in6
}

// In6_prflags
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/in6_prflags
type In6_prflags struct {
	bitfield0     uint8
	Prf_reserved1 U_char
	Prf_reserved2 U_short
	Prf_reserved3 U_char
	Prf_reserved4 U_short
	Prf_ra        [1]byte
}

// Onlink returns the Onlink bitfield.
func (s *In6_prflags) Onlink() uint8 {
	return (s.bitfield0 >> 0) & ((1 << 1) - 1)
}

// SetOnlink updates the Onlink bitfield.
func (s *In6_prflags) SetOnlink(v uint8) {
	const mask uint8 = (1 << 1) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 0)) | ((v & mask) << 0)
}

// Autonomous returns the Autonomous bitfield.
func (s *In6_prflags) Autonomous() uint8 {
	return (s.bitfield0 >> 1) & ((1 << 1) - 1)
}

// SetAutonomous updates the Autonomous bitfield.
func (s *In6_prflags) SetAutonomous(v uint8) {
	const mask uint8 = (1 << 1) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 1)) | ((v & mask) << 1)
}

// Reserved returns the Reserved bitfield.
func (s *In6_prflags) Reserved() uint8 {
	return (s.bitfield0 >> 2) & ((1 << 6) - 1)
}

// SetReserved updates the Reserved bitfield.
func (s *In6_prflags) SetReserved(v uint8) {
	const mask uint8 = (1 << 6) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 2)) | ((v & mask) << 2)
}

// In6_rrenumreq
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/in6_rrenumreq
type In6_rrenumreq struct {
	Irr_name        [16]int8
	Irr_origin      U_char
	Irr_m_len       U_char
	Irr_m_minlen    U_char
	Irr_m_maxlen    U_char
	Irr_u_uselen    U_char
	Irr_u_keeplen   U_char
	bitfield0       uint8
	Irr_vltime      U_int32_t
	Irr_pltime      U_int32_t
	Irr_flags       In6_prflags
	Irr_matchprefix Sockaddr_in6
	Irr_useprefix   Sockaddr_in6
}

// Onlink returns the Onlink bitfield.
func (s *In6_rrenumreq) Onlink() uint8 {
	return (s.bitfield0 >> 0) & ((1 << 1) - 1)
}

// SetOnlink updates the Onlink bitfield.
func (s *In6_rrenumreq) SetOnlink(v uint8) {
	const mask uint8 = (1 << 1) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 0)) | ((v & mask) << 0)
}

// Autonomous returns the Autonomous bitfield.
func (s *In6_rrenumreq) Autonomous() uint8 {
	return (s.bitfield0 >> 1) & ((1 << 1) - 1)
}

// SetAutonomous updates the Autonomous bitfield.
func (s *In6_rrenumreq) SetAutonomous(v uint8) {
	const mask uint8 = (1 << 1) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 1)) | ((v & mask) << 1)
}

// Reserved returns the Reserved bitfield.
func (s *In6_rrenumreq) Reserved() uint8 {
	return (s.bitfield0 >> 2) & ((1 << 6) - 1)
}

// SetReserved updates the Reserved bitfield.
func (s *In6_rrenumreq) SetReserved(v uint8) {
	const mask uint8 = (1 << 6) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 2)) | ((v & mask) << 2)
}

// In_addr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/in_addr
type In_addr struct {
	S_addr uint32
}

// In_addr_4in6
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/in_addr_4in6
type In_addr_4in6 struct {
	Ia46_pad32 [3]U_int32_t
	Ia46_addr4 In_addr
}

// In_aliasreq
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/in_aliasreq
type In_aliasreq struct {
	Ifra_name      [16]int8
	Ifra_addr      Sockaddr_in
	Ifra_broadaddr Sockaddr_in
	Ifra_mask      Sockaddr_in
}

// In_pktinfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/in_pktinfo
type In_pktinfo struct {
	Ipi_ifindex  uint32
	Ipi_spec_dst In_addr
	Ipi_addr     In_addr
}

// Info_tuple
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/info_tuple
type Info_tuple struct {
	Itpl_proto      U_int8_t
	Itpl_localaddr  [7]uint32
	Itpl_remoteaddr [7]uint32
}

// Inpcb
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/inpcb
type Inpcb struct {
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
	storage [208]byte
}

// Inp_hash returns the Inp_hash field from the record's packed storage.
func (s *Inpcb) Inp_hash() [2]uint32 {
	return *(*[2]uint32)(unsafe.Pointer(&s.storage[0]))
}

// SetInp_hash updates the Inp_hash field in the record's packed storage.
func (s *Inpcb) SetInp_hash(v [2]uint32) {
	*(*[2]uint32)(unsafe.Pointer(&s.storage[0])) = v
}

// Reserved1 returns the Reserved1 field from the record's packed storage.
func (s *Inpcb) Reserved1() In_addr {
	return *(*In_addr)(unsafe.Pointer(&s.storage[8]))
}

// SetReserved1 updates the Reserved1 field in the record's packed storage.
func (s *Inpcb) SetReserved1(v In_addr) {
	*(*In_addr)(unsafe.Pointer(&s.storage[8])) = v
}

// Reserved2 returns the Reserved2 field from the record's packed storage.
func (s *Inpcb) Reserved2() In_addr {
	return *(*In_addr)(unsafe.Pointer(&s.storage[12]))
}

// SetReserved2 updates the Reserved2 field in the record's packed storage.
func (s *Inpcb) SetReserved2(v In_addr) {
	*(*In_addr)(unsafe.Pointer(&s.storage[12])) = v
}

// Inp_fport returns the Inp_fport field from the record's packed storage.
func (s *Inpcb) Inp_fport() U_short {
	return U_short(binary.NativeEndian.Uint16(s.storage[16:18]))
}

// SetInp_fport updates the Inp_fport field in the record's packed storage.
func (s *Inpcb) SetInp_fport(v U_short) {
	binary.NativeEndian.PutUint16(s.storage[16:18], uint16(v))
}

// Inp_lport returns the Inp_lport field from the record's packed storage.
func (s *Inpcb) Inp_lport() U_short {
	return U_short(binary.NativeEndian.Uint16(s.storage[18:20]))
}

// SetInp_lport updates the Inp_lport field in the record's packed storage.
func (s *Inpcb) SetInp_lport(v U_short) {
	binary.NativeEndian.PutUint16(s.storage[18:20], uint16(v))
}

// Inp_list returns the Inp_list field from the record's packed storage.
func (s *Inpcb) Inp_list() [2]uint32 {
	return *(*[2]uint32)(unsafe.Pointer(&s.storage[20]))
}

// SetInp_list updates the Inp_list field in the record's packed storage.
func (s *Inpcb) SetInp_list(v [2]uint32) {
	*(*[2]uint32)(unsafe.Pointer(&s.storage[20])) = v
}

// Inp_ppcb returns the Inp_ppcb field from the record's packed storage.
func (s *Inpcb) Inp_ppcb() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[28:32]))
}

// SetInp_ppcb updates the Inp_ppcb field in the record's packed storage.
func (s *Inpcb) SetInp_ppcb(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[28:32], uint32(v))
}

// Inp_pcbinfo returns the Inp_pcbinfo field from the record's packed storage.
func (s *Inpcb) Inp_pcbinfo() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[32:36]))
}

// SetInp_pcbinfo updates the Inp_pcbinfo field in the record's packed storage.
func (s *Inpcb) SetInp_pcbinfo(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[32:36], uint32(v))
}

// Inp_socket returns the Inp_socket field from the record's packed storage.
func (s *Inpcb) Inp_socket() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[36:40]))
}

// SetInp_socket updates the Inp_socket field in the record's packed storage.
func (s *Inpcb) SetInp_socket(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[36:40], uint32(v))
}

// Nat_owner returns the Nat_owner field from the record's packed storage.
func (s *Inpcb) Nat_owner() U_char {
	return *(*U_char)(unsafe.Pointer(&s.storage[40]))
}

// SetNat_owner updates the Nat_owner field in the record's packed storage.
func (s *Inpcb) SetNat_owner(v U_char) {
	*(*U_char)(unsafe.Pointer(&s.storage[40])) = v
}

// Nat_cookie returns the Nat_cookie field from the record's packed storage.
func (s *Inpcb) Nat_cookie() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[44:48]))
}

// SetNat_cookie updates the Nat_cookie field in the record's packed storage.
func (s *Inpcb) SetNat_cookie(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[44:48], uint32(v))
}

// Inp_portlist returns the Inp_portlist field from the record's packed storage.
func (s *Inpcb) Inp_portlist() [2]uint32 {
	return *(*[2]uint32)(unsafe.Pointer(&s.storage[48]))
}

// SetInp_portlist updates the Inp_portlist field in the record's packed storage.
func (s *Inpcb) SetInp_portlist(v [2]uint32) {
	*(*[2]uint32)(unsafe.Pointer(&s.storage[48])) = v
}

// Inp_phd returns the Inp_phd field from the record's packed storage.
func (s *Inpcb) Inp_phd() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[56:60]))
}

// SetInp_phd updates the Inp_phd field in the record's packed storage.
func (s *Inpcb) SetInp_phd(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[56:60], uint32(v))
}

// Inp_gencnt returns the Inp_gencnt field from the record's packed storage.
func (s *Inpcb) Inp_gencnt() Inp_gen_t {
	return Inp_gen_t(binary.NativeEndian.Uint64(s.storage[60:68]))
}

// SetInp_gencnt updates the Inp_gencnt field in the record's packed storage.
func (s *Inpcb) SetInp_gencnt(v Inp_gen_t) {
	binary.NativeEndian.PutUint64(s.storage[60:68], uint64(v))
}

// Inp_flags returns the Inp_flags field from the record's packed storage.
func (s *Inpcb) Inp_flags() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[68:72]))
}

// SetInp_flags updates the Inp_flags field in the record's packed storage.
func (s *Inpcb) SetInp_flags(v int32) {
	binary.NativeEndian.PutUint32(s.storage[68:72], uint32(v))
}

// Inp_flow returns the Inp_flow field from the record's packed storage.
func (s *Inpcb) Inp_flow() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[72:76]))
}

// SetInp_flow updates the Inp_flow field in the record's packed storage.
func (s *Inpcb) SetInp_flow(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[72:76], uint32(v))
}

// Inp_vflag returns the Inp_vflag field from the record's packed storage.
func (s *Inpcb) Inp_vflag() U_char {
	return *(*U_char)(unsafe.Pointer(&s.storage[76]))
}

// SetInp_vflag updates the Inp_vflag field in the record's packed storage.
func (s *Inpcb) SetInp_vflag(v U_char) {
	*(*U_char)(unsafe.Pointer(&s.storage[76])) = v
}

// Inp_ip_ttl returns the Inp_ip_ttl field from the record's packed storage.
func (s *Inpcb) Inp_ip_ttl() U_char {
	return *(*U_char)(unsafe.Pointer(&s.storage[77]))
}

// SetInp_ip_ttl updates the Inp_ip_ttl field in the record's packed storage.
func (s *Inpcb) SetInp_ip_ttl(v U_char) {
	*(*U_char)(unsafe.Pointer(&s.storage[77])) = v
}

// Inp_ip_p returns the Inp_ip_p field from the record's packed storage.
func (s *Inpcb) Inp_ip_p() U_char {
	return *(*U_char)(unsafe.Pointer(&s.storage[78]))
}

// SetInp_ip_p updates the Inp_ip_p field in the record's packed storage.
func (s *Inpcb) SetInp_ip_p(v U_char) {
	*(*U_char)(unsafe.Pointer(&s.storage[78])) = v
}

// Inp_dependfaddr returns the Inp_dependfaddr field from the record's packed storage.
func (s *Inpcb) Inp_dependfaddr() [4]uint32 {
	return *(*[4]uint32)(unsafe.Pointer(&s.storage[80]))
}

// SetInp_dependfaddr updates the Inp_dependfaddr field in the record's packed storage.
func (s *Inpcb) SetInp_dependfaddr(v [4]uint32) {
	*(*[4]uint32)(unsafe.Pointer(&s.storage[80])) = v
}

// Inp_dependladdr returns the Inp_dependladdr field from the record's packed storage.
func (s *Inpcb) Inp_dependladdr() [4]uint32 {
	return *(*[4]uint32)(unsafe.Pointer(&s.storage[96]))
}

// SetInp_dependladdr updates the Inp_dependladdr field in the record's packed storage.
func (s *Inpcb) SetInp_dependladdr(v [4]uint32) {
	*(*[4]uint32)(unsafe.Pointer(&s.storage[96])) = v
}

// Inp_dependroute returns the Inp_dependroute field from the record's packed storage.
func (s *Inpcb) Inp_dependroute() [32]byte {
	return *(*[32]byte)(unsafe.Pointer(&s.storage[112]))
}

// SetInp_dependroute updates the Inp_dependroute field in the record's packed storage.
func (s *Inpcb) SetInp_dependroute(v [32]byte) {
	*(*[32]byte)(unsafe.Pointer(&s.storage[112])) = v
}

// Inp4_ip_tos returns the Inp4_ip_tos field from the record's packed storage.
func (s *Inpcb) Inp4_ip_tos() U_char {
	return *(*U_char)(unsafe.Pointer(&s.storage[144]))
}

// SetInp4_ip_tos updates the Inp4_ip_tos field in the record's packed storage.
func (s *Inpcb) SetInp4_ip_tos(v U_char) {
	*(*U_char)(unsafe.Pointer(&s.storage[144])) = v
}

// Inp4_options returns the Inp4_options field from the record's packed storage.
func (s *Inpcb) Inp4_options() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[148:152]))
}

// SetInp4_options updates the Inp4_options field in the record's packed storage.
func (s *Inpcb) SetInp4_options(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[148:152], uint32(v))
}

// Inp4_moptions returns the Inp4_moptions field from the record's packed storage.
func (s *Inpcb) Inp4_moptions() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[152:156]))
}

// SetInp4_moptions updates the Inp4_moptions field in the record's packed storage.
func (s *Inpcb) SetInp4_moptions(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[152:156], uint32(v))
}

// Inp6_options returns the Inp6_options field from the record's packed storage.
func (s *Inpcb) Inp6_options() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[156:160]))
}

// SetInp6_options updates the Inp6_options field in the record's packed storage.
func (s *Inpcb) SetInp6_options(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[156:160], uint32(v))
}

// Inp6_hlim returns the Inp6_hlim field from the record's packed storage.
func (s *Inpcb) Inp6_hlim() U_int8_t {
	return *(*U_int8_t)(unsafe.Pointer(&s.storage[160]))
}

// SetInp6_hlim updates the Inp6_hlim field in the record's packed storage.
func (s *Inpcb) SetInp6_hlim(v U_int8_t) {
	*(*U_int8_t)(unsafe.Pointer(&s.storage[160])) = v
}

// Unused_uint8_1 returns the Unused_uint8_1 field from the record's packed storage.
func (s *Inpcb) Unused_uint8_1() U_int8_t {
	return *(*U_int8_t)(unsafe.Pointer(&s.storage[161]))
}

// SetUnused_uint8_1 updates the Unused_uint8_1 field in the record's packed storage.
func (s *Inpcb) SetUnused_uint8_1(v U_int8_t) {
	*(*U_int8_t)(unsafe.Pointer(&s.storage[161])) = v
}

// Unused_uint16_1 returns the Unused_uint16_1 field from the record's packed storage.
func (s *Inpcb) Unused_uint16_1() Ushort {
	return Ushort(binary.NativeEndian.Uint16(s.storage[162:164]))
}

// SetUnused_uint16_1 updates the Unused_uint16_1 field in the record's packed storage.
func (s *Inpcb) SetUnused_uint16_1(v Ushort) {
	binary.NativeEndian.PutUint16(s.storage[162:164], uint16(v))
}

// Inp6_outputopts returns the Inp6_outputopts field from the record's packed storage.
func (s *Inpcb) Inp6_outputopts() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[164:168]))
}

// SetInp6_outputopts updates the Inp6_outputopts field in the record's packed storage.
func (s *Inpcb) SetInp6_outputopts(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[164:168], uint32(v))
}

// Inp6_moptions returns the Inp6_moptions field from the record's packed storage.
func (s *Inpcb) Inp6_moptions() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[168:172]))
}

// SetInp6_moptions updates the Inp6_moptions field in the record's packed storage.
func (s *Inpcb) SetInp6_moptions(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[168:172], uint32(v))
}

// Inp6_icmp6filt returns the Inp6_icmp6filt field from the record's packed storage.
func (s *Inpcb) Inp6_icmp6filt() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[172:176]))
}

// SetInp6_icmp6filt updates the Inp6_icmp6filt field in the record's packed storage.
func (s *Inpcb) SetInp6_icmp6filt(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[172:176], uint32(v))
}

// Inp6_cksum returns the Inp6_cksum field from the record's packed storage.
func (s *Inpcb) Inp6_cksum() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[176:180]))
}

// SetInp6_cksum updates the Inp6_cksum field in the record's packed storage.
func (s *Inpcb) SetInp6_cksum(v int32) {
	binary.NativeEndian.PutUint32(s.storage[176:180], uint32(v))
}

// Inp6_ifindex returns the Inp6_ifindex field from the record's packed storage.
func (s *Inpcb) Inp6_ifindex() U_short {
	return U_short(binary.NativeEndian.Uint16(s.storage[180:182]))
}

// SetInp6_ifindex updates the Inp6_ifindex field in the record's packed storage.
func (s *Inpcb) SetInp6_ifindex(v U_short) {
	binary.NativeEndian.PutUint16(s.storage[180:182], uint16(v))
}

// Inp6_hops returns the Inp6_hops field from the record's packed storage.
func (s *Inpcb) Inp6_hops() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[182:184]))
}

// SetInp6_hops updates the Inp6_hops field in the record's packed storage.
func (s *Inpcb) SetInp6_hops(v int16) {
	binary.NativeEndian.PutUint16(s.storage[182:184], uint16(v))
}

// Hash_element returns the Hash_element field from the record's packed storage.
func (s *Inpcb) Hash_element() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[184:188]))
}

// SetHash_element updates the Hash_element field in the record's packed storage.
func (s *Inpcb) SetHash_element(v int32) {
	binary.NativeEndian.PutUint32(s.storage[184:188], uint32(v))
}

// Inp_saved_ppcb returns the Inp_saved_ppcb field from the record's packed storage.
func (s *Inpcb) Inp_saved_ppcb() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[188:192]))
}

// SetInp_saved_ppcb updates the Inp_saved_ppcb field in the record's packed storage.
func (s *Inpcb) SetInp_saved_ppcb(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[188:192], uint32(v))
}

// Inp_sp returns the Inp_sp field from the record's packed storage.
func (s *Inpcb) Inp_sp() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[192:196]))
}

// SetInp_sp updates the Inp_sp field in the record's packed storage.
func (s *Inpcb) SetInp_sp(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[192:196], uint32(v))
}

// Reserved returns the Reserved field from the record's packed storage.
func (s *Inpcb) Reserved() [3]U_int32_t {
	return *(*[3]U_int32_t)(unsafe.Pointer(&s.storage[196]))
}

// SetReserved updates the Reserved field in the record's packed storage.
func (s *Inpcb) SetReserved(v [3]U_int32_t) {
	*(*[3]U_int32_t)(unsafe.Pointer(&s.storage[196])) = v
}

// Inp_depend4 returns the Inp_depend4 field from the record's packed storage.
func (s *Inpcb) Inp_depend4() [4]byte {
	return *(*[4]byte)(unsafe.Pointer(&s.storage[144]))
}

// SetInp_depend4 updates the Inp_depend4 field in the record's packed storage.
func (s *Inpcb) SetInp_depend4(v [4]byte) {
	*(*[4]byte)(unsafe.Pointer(&s.storage[144])) = v
}

// Inp_depend6 returns the Inp_depend6 field from the record's packed storage.
func (s *Inpcb) Inp_depend6() [4]byte {
	return *(*[4]byte)(unsafe.Pointer(&s.storage[156]))
}

// SetInp_depend6 updates the Inp_depend6 field in the record's packed storage.
func (s *Inpcb) SetInp_depend6(v [4]byte) {
	*(*[4]byte)(unsafe.Pointer(&s.storage[156])) = v
}

// Inpcb64_list_entry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/inpcb64_list_entry
type Inpcb64_list_entry struct {
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

// Le_next returns the Le_next field from the record's packed storage.
func (s *Inpcb64_list_entry) Le_next() U_int64_t {
	return U_int64_t(binary.NativeEndian.Uint64(s.storage[0:8]))
}

// SetLe_next updates the Le_next field in the record's packed storage.
func (s *Inpcb64_list_entry) SetLe_next(v U_int64_t) {
	binary.NativeEndian.PutUint64(s.storage[0:8], uint64(v))
}

// Le_prev returns the Le_prev field from the record's packed storage.
func (s *Inpcb64_list_entry) Le_prev() U_int64_t {
	return U_int64_t(binary.NativeEndian.Uint64(s.storage[8:16]))
}

// SetLe_prev updates the Le_prev field in the record's packed storage.
func (s *Inpcb64_list_entry) SetLe_prev(v U_int64_t) {
	binary.NativeEndian.PutUint64(s.storage[8:16], uint64(v))
}

// Instrs_cycles_snapshot
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/instrs_cycles_snapshot
type Instrs_cycles_snapshot struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [16]byte
}

// Ics_instructions returns the Ics_instructions field from the record's packed storage.
func (s *Instrs_cycles_snapshot) Ics_instructions() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[0:8]))
}

// SetIcs_instructions updates the Ics_instructions field in the record's packed storage.
func (s *Instrs_cycles_snapshot) SetIcs_instructions(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[0:8], uint64(v))
}

// Ics_cycles returns the Ics_cycles field from the record's packed storage.
func (s *Instrs_cycles_snapshot) Ics_cycles() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[8:16]))
}

// SetIcs_cycles updates the Ics_cycles field in the record's packed storage.
func (s *Instrs_cycles_snapshot) SetIcs_cycles(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[8:16], uint64(v))
}

// Instrs_cycles_snapshot_v2
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/instrs_cycles_snapshot_v2
type Instrs_cycles_snapshot_v2 struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [32]byte
}

// Ics_instructions returns the Ics_instructions field from the record's packed storage.
func (s *Instrs_cycles_snapshot_v2) Ics_instructions() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[0:8]))
}

// SetIcs_instructions updates the Ics_instructions field in the record's packed storage.
func (s *Instrs_cycles_snapshot_v2) SetIcs_instructions(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[0:8], uint64(v))
}

// Ics_cycles returns the Ics_cycles field from the record's packed storage.
func (s *Instrs_cycles_snapshot_v2) Ics_cycles() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[8:16]))
}

// SetIcs_cycles updates the Ics_cycles field in the record's packed storage.
func (s *Instrs_cycles_snapshot_v2) SetIcs_cycles(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[8:16], uint64(v))
}

// Ics_p_instructions returns the Ics_p_instructions field from the record's packed storage.
func (s *Instrs_cycles_snapshot_v2) Ics_p_instructions() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[16:24]))
}

// SetIcs_p_instructions updates the Ics_p_instructions field in the record's packed storage.
func (s *Instrs_cycles_snapshot_v2) SetIcs_p_instructions(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[16:24], uint64(v))
}

// Ics_p_cycles returns the Ics_p_cycles field from the record's packed storage.
func (s *Instrs_cycles_snapshot_v2) Ics_p_cycles() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[24:32]))
}

// SetIcs_p_cycles updates the Ics_p_cycles field in the record's packed storage.
func (s *Instrs_cycles_snapshot_v2) SetIcs_p_cycles(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[24:32], uint64(v))
}

// Internal_state
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/internal_state
type Internal_state struct {
	Dummy int32
}

// Io_stat_entry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/io_stat_entry
type Io_stat_entry struct {
	Count uint64
	Size  uint64
}

// Io_stats_snapshot
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/io_stats_snapshot
type Io_stats_snapshot struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [160]byte
}

// Ss_disk_reads_count returns the Ss_disk_reads_count field from the record's packed storage.
func (s *Io_stats_snapshot) Ss_disk_reads_count() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[0:8]))
}

// SetSs_disk_reads_count updates the Ss_disk_reads_count field in the record's packed storage.
func (s *Io_stats_snapshot) SetSs_disk_reads_count(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[0:8], uint64(v))
}

// Ss_disk_reads_size returns the Ss_disk_reads_size field from the record's packed storage.
func (s *Io_stats_snapshot) Ss_disk_reads_size() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[8:16]))
}

// SetSs_disk_reads_size updates the Ss_disk_reads_size field in the record's packed storage.
func (s *Io_stats_snapshot) SetSs_disk_reads_size(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[8:16], uint64(v))
}

// Ss_disk_writes_count returns the Ss_disk_writes_count field from the record's packed storage.
func (s *Io_stats_snapshot) Ss_disk_writes_count() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[16:24]))
}

// SetSs_disk_writes_count updates the Ss_disk_writes_count field in the record's packed storage.
func (s *Io_stats_snapshot) SetSs_disk_writes_count(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[16:24], uint64(v))
}

// Ss_disk_writes_size returns the Ss_disk_writes_size field from the record's packed storage.
func (s *Io_stats_snapshot) Ss_disk_writes_size() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[24:32]))
}

// SetSs_disk_writes_size updates the Ss_disk_writes_size field in the record's packed storage.
func (s *Io_stats_snapshot) SetSs_disk_writes_size(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[24:32], uint64(v))
}

// Ss_io_priority_count returns the Ss_io_priority_count field from the record's packed storage.
func (s *Io_stats_snapshot) Ss_io_priority_count() [4]uint64 {
	return *(*[4]uint64)(unsafe.Pointer(&s.storage[32]))
}

// SetSs_io_priority_count updates the Ss_io_priority_count field in the record's packed storage.
func (s *Io_stats_snapshot) SetSs_io_priority_count(v [4]uint64) {
	*(*[4]uint64)(unsafe.Pointer(&s.storage[32])) = v
}

// Ss_io_priority_size returns the Ss_io_priority_size field from the record's packed storage.
func (s *Io_stats_snapshot) Ss_io_priority_size() [4]uint64 {
	return *(*[4]uint64)(unsafe.Pointer(&s.storage[64]))
}

// SetSs_io_priority_size updates the Ss_io_priority_size field in the record's packed storage.
func (s *Io_stats_snapshot) SetSs_io_priority_size(v [4]uint64) {
	*(*[4]uint64)(unsafe.Pointer(&s.storage[64])) = v
}

// Ss_paging_count returns the Ss_paging_count field from the record's packed storage.
func (s *Io_stats_snapshot) Ss_paging_count() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[96:104]))
}

// SetSs_paging_count updates the Ss_paging_count field in the record's packed storage.
func (s *Io_stats_snapshot) SetSs_paging_count(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[96:104], uint64(v))
}

// Ss_paging_size returns the Ss_paging_size field from the record's packed storage.
func (s *Io_stats_snapshot) Ss_paging_size() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[104:112]))
}

// SetSs_paging_size updates the Ss_paging_size field in the record's packed storage.
func (s *Io_stats_snapshot) SetSs_paging_size(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[104:112], uint64(v))
}

// Ss_non_paging_count returns the Ss_non_paging_count field from the record's packed storage.
func (s *Io_stats_snapshot) Ss_non_paging_count() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[112:120]))
}

// SetSs_non_paging_count updates the Ss_non_paging_count field in the record's packed storage.
func (s *Io_stats_snapshot) SetSs_non_paging_count(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[112:120], uint64(v))
}

// Ss_non_paging_size returns the Ss_non_paging_size field from the record's packed storage.
func (s *Io_stats_snapshot) Ss_non_paging_size() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[120:128]))
}

// SetSs_non_paging_size updates the Ss_non_paging_size field in the record's packed storage.
func (s *Io_stats_snapshot) SetSs_non_paging_size(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[120:128], uint64(v))
}

// Ss_data_count returns the Ss_data_count field from the record's packed storage.
func (s *Io_stats_snapshot) Ss_data_count() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[128:136]))
}

// SetSs_data_count updates the Ss_data_count field in the record's packed storage.
func (s *Io_stats_snapshot) SetSs_data_count(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[128:136], uint64(v))
}

// Ss_data_size returns the Ss_data_size field from the record's packed storage.
func (s *Io_stats_snapshot) Ss_data_size() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[136:144]))
}

// SetSs_data_size updates the Ss_data_size field in the record's packed storage.
func (s *Io_stats_snapshot) SetSs_data_size(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[136:144], uint64(v))
}

// Ss_metadata_count returns the Ss_metadata_count field from the record's packed storage.
func (s *Io_stats_snapshot) Ss_metadata_count() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[144:152]))
}

// SetSs_metadata_count updates the Ss_metadata_count field in the record's packed storage.
func (s *Io_stats_snapshot) SetSs_metadata_count(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[144:152], uint64(v))
}

// Ss_metadata_size returns the Ss_metadata_size field from the record's packed storage.
func (s *Io_stats_snapshot) Ss_metadata_size() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[152:160]))
}

// SetSs_metadata_size updates the Ss_metadata_size field in the record's packed storage.
func (s *Io_stats_snapshot) SetSs_metadata_size(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[152:160], uint64(v))
}

// Iocompressionstats_notification_subsystem
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iocompressionstats_notification_subsystem-oj5
type Iocompressionstats_notification_subsystem struct {
	Server   unsafe.Pointer
	Start    Mach_msg_id_t
	End      Mach_msg_id_t
	Maxsize  uint32
	Reserved Vm_address_t
	Routine  unsafe.Pointer
}

// Iocs_store_buffer_entry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iocs_store_buffer_entry
type Iocs_store_buffer_entry struct {
	Path_name [128]int8
	Iocs      [45]uint64
}

// Iovec
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iovec
type Iovec struct {
	Iov_base unsafe.Pointer
	Iov_len  uintptr
}

// Ip
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ip
type Ip struct {
	bitfield0 uint8
	Ip_tos    U_char
	Ip_len    U_short
	Ip_id     U_short
	Ip_off    U_short
	Ip_ttl    U_char
	Ip_p      U_char
	Ip_sum    U_short
	Ip_src    In_addr
	Ip_dst    In_addr
}

// Ip_hl returns the Ip_hl bitfield.
func (s *Ip) Ip_hl() uint8 {
	return (s.bitfield0 >> 0) & ((1 << 4) - 1)
}

// SetIp_hl updates the Ip_hl bitfield.
func (s *Ip) SetIp_hl(v uint8) {
	const mask uint8 = (1 << 4) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 0)) | ((v & mask) << 0)
}

// Ip_v returns the Ip_v bitfield.
func (s *Ip) Ip_v() uint8 {
	return (s.bitfield0 >> 4) & ((1 << 4) - 1)
}

// SetIp_v updates the Ip_v bitfield.
func (s *Ip) SetIp_v(v uint8) {
	const mask uint8 = (1 << 4) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 4)) | ((v & mask) << 4)
}

// Ip6_dest
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ip6_dest
type Ip6_dest struct {
	Ip6d_nxt U_int8_t
	Ip6d_len U_int8_t
}

// Ip6_ext
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ip6_ext
type Ip6_ext struct {
	Ip6e_nxt U_int8_t
	Ip6e_len U_int8_t
}

// Ip6_frag
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ip6_frag
type Ip6_frag struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [8]byte
}

// Ip6f_nxt returns the Ip6f_nxt field from the record's packed storage.
func (s *Ip6_frag) Ip6f_nxt() U_int8_t {
	return *(*U_int8_t)(unsafe.Pointer(&s.storage[0]))
}

// SetIp6f_nxt updates the Ip6f_nxt field in the record's packed storage.
func (s *Ip6_frag) SetIp6f_nxt(v U_int8_t) {
	*(*U_int8_t)(unsafe.Pointer(&s.storage[0])) = v
}

// Ip6f_reserved returns the Ip6f_reserved field from the record's packed storage.
func (s *Ip6_frag) Ip6f_reserved() U_int8_t {
	return *(*U_int8_t)(unsafe.Pointer(&s.storage[1]))
}

// SetIp6f_reserved updates the Ip6f_reserved field in the record's packed storage.
func (s *Ip6_frag) SetIp6f_reserved(v U_int8_t) {
	*(*U_int8_t)(unsafe.Pointer(&s.storage[1])) = v
}

// Ip6f_offlg returns the Ip6f_offlg field from the record's packed storage.
func (s *Ip6_frag) Ip6f_offlg() U_int16_t {
	return U_int16_t(binary.NativeEndian.Uint16(s.storage[2:4]))
}

// SetIp6f_offlg updates the Ip6f_offlg field in the record's packed storage.
func (s *Ip6_frag) SetIp6f_offlg(v U_int16_t) {
	binary.NativeEndian.PutUint16(s.storage[2:4], uint16(v))
}

// Ip6f_ident returns the Ip6f_ident field from the record's packed storage.
func (s *Ip6_frag) Ip6f_ident() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetIp6f_ident updates the Ip6f_ident field in the record's packed storage.
func (s *Ip6_frag) SetIp6f_ident(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// Ip6_hbh
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ip6_hbh
type Ip6_hbh struct {
	Ip6h_nxt U_int8_t
	Ip6h_len U_int8_t
}

// Ip6_hdr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ip6_hdr
type Ip6_hdr struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [40]byte
}

// Ip6_ctlun returns the Ip6_ctlun field from the record's packed storage.
func (s *Ip6_hdr) Ip6_ctlun() [8]byte {
	return *(*[8]byte)(unsafe.Pointer(&s.storage[0]))
}

// SetIp6_ctlun updates the Ip6_ctlun field in the record's packed storage.
func (s *Ip6_hdr) SetIp6_ctlun(v [8]byte) {
	*(*[8]byte)(unsafe.Pointer(&s.storage[0])) = v
}

// Ip6_src returns the Ip6_src field from the record's packed storage.
func (s *Ip6_hdr) Ip6_src() [4]uint32 {
	return *(*[4]uint32)(unsafe.Pointer(&s.storage[8]))
}

// SetIp6_src updates the Ip6_src field in the record's packed storage.
func (s *Ip6_hdr) SetIp6_src(v [4]uint32) {
	*(*[4]uint32)(unsafe.Pointer(&s.storage[8])) = v
}

// Ip6_dst returns the Ip6_dst field from the record's packed storage.
func (s *Ip6_hdr) Ip6_dst() [4]uint32 {
	return *(*[4]uint32)(unsafe.Pointer(&s.storage[24]))
}

// SetIp6_dst updates the Ip6_dst field in the record's packed storage.
func (s *Ip6_hdr) SetIp6_dst(v [4]uint32) {
	*(*[4]uint32)(unsafe.Pointer(&s.storage[24])) = v
}

// Ip6_mtuinfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ip6_mtuinfo
type Ip6_mtuinfo struct {
	Ip6m_addr Sockaddr_in6
	Ip6m_mtu  uint32
}

// Ip6_opt
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ip6_opt
type Ip6_opt struct {
	Ip6o_type U_int8_t
	Ip6o_len  U_int8_t
}

// Ip6_opt_jumbo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ip6_opt_jumbo
type Ip6_opt_jumbo struct {
	Ip6oj_type      U_int8_t
	Ip6oj_len       U_int8_t
	Ip6oj_jumbo_len [4]U_int8_t
}

// Ip6_opt_nsap
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ip6_opt_nsap
type Ip6_opt_nsap struct {
	Ip6on_type         U_int8_t
	Ip6on_len          U_int8_t
	Ip6on_src_nsap_len U_int8_t
	Ip6on_dst_nsap_len U_int8_t
}

// Ip6_opt_router
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ip6_opt_router
type Ip6_opt_router struct {
	Ip6or_type  U_int8_t
	Ip6or_len   U_int8_t
	Ip6or_value [2]U_int8_t
}

// Ip6_opt_tunnel
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ip6_opt_tunnel
type Ip6_opt_tunnel struct {
	Ip6ot_type        U_int8_t
	Ip6ot_len         U_int8_t
	Ip6ot_encap_limit U_int8_t
}

// Ip6_rthdr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ip6_rthdr
type Ip6_rthdr struct {
	Ip6r_nxt     U_int8_t
	Ip6r_len     U_int8_t
	Ip6r_type    U_int8_t
	Ip6r_segleft U_int8_t
}

// Ip6_rthdr0
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ip6_rthdr0
type Ip6_rthdr0 struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [8]byte
}

// Ip6r0_nxt returns the Ip6r0_nxt field from the record's packed storage.
func (s *Ip6_rthdr0) Ip6r0_nxt() U_int8_t {
	return *(*U_int8_t)(unsafe.Pointer(&s.storage[0]))
}

// SetIp6r0_nxt updates the Ip6r0_nxt field in the record's packed storage.
func (s *Ip6_rthdr0) SetIp6r0_nxt(v U_int8_t) {
	*(*U_int8_t)(unsafe.Pointer(&s.storage[0])) = v
}

// Ip6r0_len returns the Ip6r0_len field from the record's packed storage.
func (s *Ip6_rthdr0) Ip6r0_len() U_int8_t {
	return *(*U_int8_t)(unsafe.Pointer(&s.storage[1]))
}

// SetIp6r0_len updates the Ip6r0_len field in the record's packed storage.
func (s *Ip6_rthdr0) SetIp6r0_len(v U_int8_t) {
	*(*U_int8_t)(unsafe.Pointer(&s.storage[1])) = v
}

// Ip6r0_type returns the Ip6r0_type field from the record's packed storage.
func (s *Ip6_rthdr0) Ip6r0_type() U_int8_t {
	return *(*U_int8_t)(unsafe.Pointer(&s.storage[2]))
}

// SetIp6r0_type updates the Ip6r0_type field in the record's packed storage.
func (s *Ip6_rthdr0) SetIp6r0_type(v U_int8_t) {
	*(*U_int8_t)(unsafe.Pointer(&s.storage[2])) = v
}

// Ip6r0_segleft returns the Ip6r0_segleft field from the record's packed storage.
func (s *Ip6_rthdr0) Ip6r0_segleft() U_int8_t {
	return *(*U_int8_t)(unsafe.Pointer(&s.storage[3]))
}

// SetIp6r0_segleft updates the Ip6r0_segleft field in the record's packed storage.
func (s *Ip6_rthdr0) SetIp6r0_segleft(v U_int8_t) {
	*(*U_int8_t)(unsafe.Pointer(&s.storage[3])) = v
}

// Ip6r0_reserved returns the Ip6r0_reserved field from the record's packed storage.
func (s *Ip6_rthdr0) Ip6r0_reserved() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetIp6r0_reserved updates the Ip6r0_reserved field in the record's packed storage.
func (s *Ip6_rthdr0) SetIp6r0_reserved(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// Ip_linklocal_stat
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ip_linklocal_stat
type Ip_linklocal_stat struct {
	Iplls_in_total   U_int32_t
	Iplls_in_badttl  U_int32_t
	Iplls_out_total  U_int32_t
	Iplls_out_badttl U_int32_t
}

// Ip_mreq
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ip_mreq
type Ip_mreq struct {
	Imr_multiaddr In_addr
	Imr_interface In_addr
}

// Ip_mreq_source
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ip_mreq_source
type Ip_mreq_source struct {
	Imr_multiaddr  In_addr
	Imr_sourceaddr In_addr
	Imr_interface  In_addr
}

// Ip_mreqn
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ip_mreqn
type Ip_mreqn struct {
	Imr_multiaddr In_addr
	Imr_address   In_addr
	Imr_ifindex   int32
}

// Ip_opts
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ip_opts
type Ip_opts struct {
	Ip_dst  In_addr
	Ip_opts [40]int8
}

// Ip_timestamp
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ip_timestamp
type Ip_timestamp struct {
	Ipt_code      U_char
	Ipt_len       U_char
	Ipt_ptr       U_char
	bitfield3     uint8
	Ipt_timestamp [2]uint32
}

// Ipt_flg returns the Ipt_flg bitfield.
func (s *Ip_timestamp) Ipt_flg() uint8 {
	return (s.bitfield3 >> 0) & ((1 << 4) - 1)
}

// SetIpt_flg updates the Ipt_flg bitfield.
func (s *Ip_timestamp) SetIpt_flg(v uint8) {
	const mask uint8 = (1 << 4) - 1
	s.bitfield3 = (s.bitfield3 &^ (mask << 0)) | ((v & mask) << 0)
}

// Ipt_oflw returns the Ipt_oflw bitfield.
func (s *Ip_timestamp) Ipt_oflw() uint8 {
	return (s.bitfield3 >> 4) & ((1 << 4) - 1)
}

// SetIpt_oflw updates the Ipt_oflw bitfield.
func (s *Ip_timestamp) SetIpt_oflw(v uint8) {
	const mask uint8 = (1 << 4) - 1
	s.bitfield3 = (s.bitfield3 &^ (mask << 4)) | ((v & mask) << 4)
}

// Ipc_perm
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ipc_perm
type Ipc_perm struct {
	Uid  uint32
	Gid  uint32
	Cuid uint32
	Cgid uint32
	Mode uint16
	_seq uint16
	_key Key_t
}

// Ipcomp
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ipcomp
type Ipcomp struct {
	Comp_nxt   U_int8_t
	Comp_flags U_int8_t
	Comp_cpi   U_int16_t
}

// Ipovly
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ipovly
type Ipovly struct {
	Ih_x1  [9]U_char
	Ih_pr  U_char
	Ih_len U_short
	Ih_src In_addr
	Ih_dst In_addr
}

// Ipsec_stats_param
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ipsec_stats_param
type Ipsec_stats_param struct {
	Utsp_packets U_int64_t
	Utsp_bytes   U_int64_t
	Utsp_errors  U_int64_t
}

// Ipsec_wake_pkt_event_data
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ipsec_wake_pkt_event_data
type Ipsec_wake_pkt_event_data struct {
	Wake_uuid [37]int8
}

// Ipsec_wake_pkt_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ipsec_wake_pkt_info
type Ipsec_wake_pkt_info struct {
	Wake_pkt     [100]U_int8_t
	Wake_uuid    [37]int8
	Wake_pkt_spi U_int32_t
	Wake_pkt_seq U_int32_t
	Wake_pkt_len U_int16_t
}

// Ipsecstat
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ipsecstat
type Ipsecstat struct {
	In_success     U_quad_t
	In_polvio      U_quad_t
	In_nosa        U_quad_t
	In_inval       U_quad_t
	In_nomem       U_quad_t
	In_badspi      U_quad_t
	In_ahreplay    U_quad_t
	In_espreplay   U_quad_t
	In_ahauthsucc  U_quad_t
	In_ahauthfail  U_quad_t
	In_espauthsucc U_quad_t
	In_espauthfail U_quad_t
	In_esphist     [256]U_quad_t
	In_ahhist      [256]U_quad_t
	In_comphist    [256]U_quad_t
	Out_success    U_quad_t
	Out_polvio     U_quad_t
	Out_nosa       U_quad_t
	Out_inval      U_quad_t
	Out_nomem      U_quad_t
	Out_noroute    U_quad_t
	Out_esphist    [256]U_quad_t
	Out_ahhist     [256]U_quad_t
	Out_comphist   [256]U_quad_t
}

// Ipstat
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ipstat
type Ipstat struct {
	Ips_total              U_int32_t
	Ips_badsum             U_int32_t
	Ips_tooshort           U_int32_t
	Ips_toosmall           U_int32_t
	Ips_badhlen            U_int32_t
	Ips_badlen             U_int32_t
	Ips_fragments          U_int32_t
	Ips_fragdropped        U_int32_t
	Ips_fragtimeout        U_int32_t
	Ips_forward            U_int32_t
	Ips_fastforward        U_int32_t
	Ips_cantforward        U_int32_t
	Ips_redirectsent       U_int32_t
	Ips_noproto            U_int32_t
	Ips_delivered          U_int32_t
	Ips_localout           U_int32_t
	Ips_odropped           U_int32_t
	Ips_reassembled        U_int32_t
	Ips_fragmented         U_int32_t
	Ips_ofragments         U_int32_t
	Ips_cantfrag           U_int32_t
	Ips_badoptions         U_int32_t
	Ips_noroute            U_int32_t
	Ips_badvers            U_int32_t
	Ips_rawout             U_int32_t
	Ips_toolong            U_int32_t
	Ips_notmember          U_int32_t
	Ips_nogif              U_int32_t
	Ips_badaddr            U_int32_t
	Ips_pktdropcntrl       U_int32_t
	Ips_rcv_swcsum         U_int32_t
	Ips_rcv_swcsum_bytes   U_int32_t
	Ips_snd_swcsum         U_int32_t
	Ips_snd_swcsum_bytes   U_int32_t
	Ips_adj                U_int32_t
	Ips_adj_hwcsum_clr     U_int32_t
	Ips_rxc_collisions     U_int32_t
	Ips_rxc_chained        U_int32_t
	Ips_rxc_notchain       U_int32_t
	Ips_rxc_chainsz_gt2    U_int32_t
	Ips_rxc_chainsz_gt4    U_int32_t
	Ips_rxc_notlist        U_int32_t
	Ips_raw_sappend_fail   U_int32_t
	Ips_necp_policy_drop   U_int32_t
	Ips_rcv_if_weak_match  U_int32_t
	Ips_rcv_if_no_match    U_int32_t
	Ips_input_ipf_drop     U_int32_t
	Ips_input_no_proto     U_int32_t
	Ips_src_addr_not_avail U_int32_t
}

// Ipv6_mreq
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ipv6_mreq
type Ipv6_mreq struct {
	Ipv6mr_multiaddr [4]uint32
	Ipv6mr_interface uint32
}

// Ipv6_prefix
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ipv6_prefix
type Ipv6_prefix struct {
	Ipv6_prefix [4]uint32
	Prefix_len  uint32
}

// Itimerval
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/itimerval
type Itimerval struct {
	It_interval Timeval
	It_value    Timeval
}

// Jetsam_coalition_snapshot
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/jetsam_coalition_snapshot
type Jetsam_coalition_snapshot struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [32]byte
}

// Jcs_id returns the Jcs_id field from the record's packed storage.
func (s *Jetsam_coalition_snapshot) Jcs_id() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[0:8]))
}

// SetJcs_id updates the Jcs_id field in the record's packed storage.
func (s *Jetsam_coalition_snapshot) SetJcs_id(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[0:8], uint64(v))
}

// Jcs_flags returns the Jcs_flags field from the record's packed storage.
func (s *Jetsam_coalition_snapshot) Jcs_flags() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[8:16]))
}

// SetJcs_flags updates the Jcs_flags field in the record's packed storage.
func (s *Jetsam_coalition_snapshot) SetJcs_flags(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[8:16], uint64(v))
}

// Jcs_thread_group returns the Jcs_thread_group field from the record's packed storage.
func (s *Jetsam_coalition_snapshot) Jcs_thread_group() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[16:24]))
}

// SetJcs_thread_group updates the Jcs_thread_group field in the record's packed storage.
func (s *Jetsam_coalition_snapshot) SetJcs_thread_group(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[16:24], uint64(v))
}

// Jcs_leader_task_uniqueid returns the Jcs_leader_task_uniqueid field from the record's packed storage.
func (s *Jetsam_coalition_snapshot) Jcs_leader_task_uniqueid() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[24:32]))
}

// SetJcs_leader_task_uniqueid updates the Jcs_leader_task_uniqueid field in the record's packed storage.
func (s *Jetsam_coalition_snapshot) SetJcs_leader_task_uniqueid(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[24:32], uint64(v))
}

// Kauth_cache_sizes
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/kauth_cache_sizes
type Kauth_cache_sizes struct {
	Kcs_group_size U_int32_t
	Kcs_id_size    U_int32_t
}

// Kauth_identity_extlookup
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/kauth_identity_extlookup
type Kauth_identity_extlookup struct {
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
	_       [0]uint64
	storage [304]byte
}

// El_seqno returns the El_seqno field from the record's packed storage.
func (s *Kauth_identity_extlookup) El_seqno() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetEl_seqno updates the El_seqno field in the record's packed storage.
func (s *Kauth_identity_extlookup) SetEl_seqno(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// El_result returns the El_result field from the record's packed storage.
func (s *Kauth_identity_extlookup) El_result() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetEl_result updates the El_result field in the record's packed storage.
func (s *Kauth_identity_extlookup) SetEl_result(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// El_flags returns the El_flags field from the record's packed storage.
func (s *Kauth_identity_extlookup) El_flags() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetEl_flags updates the El_flags field in the record's packed storage.
func (s *Kauth_identity_extlookup) SetEl_flags(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[8:12], uint32(v))
}

// El_info_pid returns the El_info_pid field from the record's packed storage.
func (s *Kauth_identity_extlookup) El_info_pid() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[12:16]))
}

// SetEl_info_pid updates the El_info_pid field in the record's packed storage.
func (s *Kauth_identity_extlookup) SetEl_info_pid(v int32) {
	binary.NativeEndian.PutUint32(s.storage[12:16], uint32(v))
}

// El_extend returns the El_extend field from the record's packed storage.
func (s *Kauth_identity_extlookup) El_extend() U_int64_t {
	return U_int64_t(binary.NativeEndian.Uint64(s.storage[16:24]))
}

// SetEl_extend updates the El_extend field in the record's packed storage.
func (s *Kauth_identity_extlookup) SetEl_extend(v U_int64_t) {
	binary.NativeEndian.PutUint64(s.storage[16:24], uint64(v))
}

// El_info_reserved_1 returns the El_info_reserved_1 field from the record's packed storage.
func (s *Kauth_identity_extlookup) El_info_reserved_1() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[24:28]))
}

// SetEl_info_reserved_1 updates the El_info_reserved_1 field in the record's packed storage.
func (s *Kauth_identity_extlookup) SetEl_info_reserved_1(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[24:28], uint32(v))
}

// El_uid returns the El_uid field from the record's packed storage.
func (s *Kauth_identity_extlookup) El_uid() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[28:32]))
}

// SetEl_uid updates the El_uid field in the record's packed storage.
func (s *Kauth_identity_extlookup) SetEl_uid(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[28:32], uint32(v))
}

// El_uguid returns the El_uguid field from the record's packed storage.
func (s *Kauth_identity_extlookup) El_uguid() [4]uint32 {
	return *(*[4]uint32)(unsafe.Pointer(&s.storage[32]))
}

// SetEl_uguid updates the El_uguid field in the record's packed storage.
func (s *Kauth_identity_extlookup) SetEl_uguid(v [4]uint32) {
	*(*[4]uint32)(unsafe.Pointer(&s.storage[32])) = v
}

// El_uguid_valid returns the El_uguid_valid field from the record's packed storage.
func (s *Kauth_identity_extlookup) El_uguid_valid() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[48:52]))
}

// SetEl_uguid_valid updates the El_uguid_valid field in the record's packed storage.
func (s *Kauth_identity_extlookup) SetEl_uguid_valid(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[48:52], uint32(v))
}

// El_usid returns the El_usid field from the record's packed storage.
func (s *Kauth_identity_extlookup) El_usid() Ntsid_t {
	return *(*Ntsid_t)(unsafe.Pointer(&s.storage[52]))
}

// SetEl_usid updates the El_usid field in the record's packed storage.
func (s *Kauth_identity_extlookup) SetEl_usid(v Ntsid_t) {
	*(*Ntsid_t)(unsafe.Pointer(&s.storage[52])) = v
}

// El_usid_valid returns the El_usid_valid field from the record's packed storage.
func (s *Kauth_identity_extlookup) El_usid_valid() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[124:128]))
}

// SetEl_usid_valid updates the El_usid_valid field in the record's packed storage.
func (s *Kauth_identity_extlookup) SetEl_usid_valid(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[124:128], uint32(v))
}

// El_gid returns the El_gid field from the record's packed storage.
func (s *Kauth_identity_extlookup) El_gid() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[128:132]))
}

// SetEl_gid updates the El_gid field in the record's packed storage.
func (s *Kauth_identity_extlookup) SetEl_gid(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[128:132], uint32(v))
}

// El_gguid returns the El_gguid field from the record's packed storage.
func (s *Kauth_identity_extlookup) El_gguid() [4]uint32 {
	return *(*[4]uint32)(unsafe.Pointer(&s.storage[132]))
}

// SetEl_gguid updates the El_gguid field in the record's packed storage.
func (s *Kauth_identity_extlookup) SetEl_gguid(v [4]uint32) {
	*(*[4]uint32)(unsafe.Pointer(&s.storage[132])) = v
}

// El_gguid_valid returns the El_gguid_valid field from the record's packed storage.
func (s *Kauth_identity_extlookup) El_gguid_valid() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[148:152]))
}

// SetEl_gguid_valid updates the El_gguid_valid field in the record's packed storage.
func (s *Kauth_identity_extlookup) SetEl_gguid_valid(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[148:152], uint32(v))
}

// El_gsid returns the El_gsid field from the record's packed storage.
func (s *Kauth_identity_extlookup) El_gsid() Ntsid_t {
	return *(*Ntsid_t)(unsafe.Pointer(&s.storage[152]))
}

// SetEl_gsid updates the El_gsid field in the record's packed storage.
func (s *Kauth_identity_extlookup) SetEl_gsid(v Ntsid_t) {
	*(*Ntsid_t)(unsafe.Pointer(&s.storage[152])) = v
}

// El_gsid_valid returns the El_gsid_valid field from the record's packed storage.
func (s *Kauth_identity_extlookup) El_gsid_valid() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[224:228]))
}

// SetEl_gsid_valid updates the El_gsid_valid field in the record's packed storage.
func (s *Kauth_identity_extlookup) SetEl_gsid_valid(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[224:228], uint32(v))
}

// El_member_valid returns the El_member_valid field from the record's packed storage.
func (s *Kauth_identity_extlookup) El_member_valid() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[228:232]))
}

// SetEl_member_valid updates the El_member_valid field in the record's packed storage.
func (s *Kauth_identity_extlookup) SetEl_member_valid(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[228:232], uint32(v))
}

// El_sup_grp_cnt returns the El_sup_grp_cnt field from the record's packed storage.
func (s *Kauth_identity_extlookup) El_sup_grp_cnt() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[232:236]))
}

// SetEl_sup_grp_cnt updates the El_sup_grp_cnt field in the record's packed storage.
func (s *Kauth_identity_extlookup) SetEl_sup_grp_cnt(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[232:236], uint32(v))
}

// El_sup_groups returns the El_sup_groups field from the record's packed storage.
func (s *Kauth_identity_extlookup) El_sup_groups() [16]uint32 {
	return *(*[16]uint32)(unsafe.Pointer(&s.storage[236]))
}

// SetEl_sup_groups updates the El_sup_groups field in the record's packed storage.
func (s *Kauth_identity_extlookup) SetEl_sup_groups(v [16]uint32) {
	*(*[16]uint32)(unsafe.Pointer(&s.storage[236])) = v
}

// Kcdata_type_definition
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/kcdata_type_definition
type Kcdata_type_definition struct {
	Kct_type_identifier uint32
	Kct_num_elements    uint32
	Kct_name            [32]int8
}

// Kern_ctl_reg
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/kern_ctl_reg
type Kern_ctl_reg struct {
	Ctl_name       [96]int8       // A Bundle ID string of up to MAX_KCTL_NAME bytes (including the ending zero). This string should not be empty.
	Ctl_id         U_int32_t      // The control ID may be dynamically assigned or it can be a 32-bit creator code assigned by DTS. For a DTS assigned creator code the CTL_FLAG_REG_ID_UNIT flag must be set. For a dynamically assigned control ID, do not set the CTL_FLAG_REG_ID_UNIT flag. The value of the dynamically assigned control ID is set to this field when the registration succeeds.
	Ctl_unit       U_int32_t      // A separate unit number to register multiple units that share the same control ID with DTS assigned creator code when the CTL_FLAG_REG_ID_UNIT flag is set. This field is ignored for a dynamically assigned control ID.
	Ctl_flags      U_int32_t      // CTL_FLAG_PRIVILEGED and/or CTL_FLAG_REG_ID_UNIT.
	Ctl_sendsize   U_int32_t      // Override the default send size. If set to zero, the default send size will be used, and this default value is set to this field to be retrieved by the caller.
	Ctl_recvsize   U_int32_t      // Override the default receive size. If set to zero, the default receive size will be used, and this default value is set to this field to be retrieved by the caller.
	Ctl_connect    unsafe.Pointer // Specify the function to be called whenever a client connects to the kernel control. This field must be specified.
	Ctl_disconnect unsafe.Pointer // Specify a function to be called whenever a client disconnects from the kernel control.
	Ctl_send       unsafe.Pointer // Specify a function to handle data send from the client to the kernel control.
	Ctl_setopt     unsafe.Pointer // Specify a function to handle set socket option operations for the kernel control.
	Ctl_getopt     unsafe.Pointer // Specify a function to handle get socket option operations for the kernel control.

}

// Kern_event_msg
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/kern_event_msg
type Kern_event_msg struct {
	Total_size   U_int32_t    // Total size of the kernel event message including the header.
	Vendor_code  U_int32_t    // The vendor code indicates which vendor generated the kernel event. This gives every vendor a unique set of classes and subclasses to use. Use the SIOCGKEVVENDOR ioctl to look up vendor codes for vendors other than Apple. Apple uses KEV_VENDOR_APPLE.
	Kev_class    U_int32_t    // The class of the kernel event.
	Kev_subclass U_int32_t    // The subclass of the kernel event.
	Id           U_int32_t    // Monotonically increasing value.
	Event_code   U_int32_t    // The event code.
	Event_data   [1]U_int32_t // Any additional data about this event. Format will depend on the vendor_code, kev_class, kev_subclass, and event_code. The length of the event_data can be determined using total_size - KEV_MSG_HEADER_SIZE.

}

// Kernel_triage_info_v1
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/kernel_triage_info_v1
type Kernel_triage_info_v1 struct {
	Triage_string1 [128]int8
	Triage_string2 [128]int8
	Triage_string3 [128]int8
	Triage_string4 [128]int8
	Triage_string5 [128]int8
}

// Kev_d_vectors
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/kev_d_vectors
type Kev_d_vectors struct {
	Data_length U_int32_t // The length of data.
	Data_ptr    unsafe.Pointer
}

// Kev_dl_issues
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/kev_dl_issues
type Kev_dl_issues struct {
	Link_data Net_event_data
	Modid     [20]U_int8_t
	Timestamp U_int64_t
	Info      [12]U_int8_t
}

// Kev_dl_link_quality_metric_data
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/kev_dl_link_quality_metric_data
type Kev_dl_link_quality_metric_data struct {
	Link_data           Net_event_data
	Link_quality_metric int32
}

// Kev_dl_low_power_mode
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/kev_dl_low_power_mode
type Kev_dl_low_power_mode struct {
	Link_data       Net_event_data
	Low_power_event int32
}

// Kev_dl_node_absence
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/kev_dl_node_absence
type Kev_dl_node_absence struct {
	Link_data         Net_event_data
	Sin6_node_address Sockaddr_in6
	Sdl_node_address  Sockaddr_dl
}

// Kev_dl_node_presence
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/kev_dl_node_presence
type Kev_dl_node_presence struct {
	Link_data             Net_event_data
	Sin6_node_address     Sockaddr_in6
	Sdl_node_address      Sockaddr_dl
	Rssi                  int32
	Link_quality_metric   int32
	Node_proximity_metric int32
	Node_service_info     [48]U_int8_t
}

// Kev_dl_proto_data
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/kev_dl_proto_data
type Kev_dl_proto_data struct {
	Link_data             Net_event_data
	Proto_family          U_int32_t
	Proto_remaining_count U_int32_t
}

// Kev_dl_rrc_state
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/kev_dl_rrc_state
type Kev_dl_rrc_state struct {
	Link_data Net_event_data
	Rrc_state U_int32_t
}

// Kev_in6_addrlifetime
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/kev_in6_addrlifetime
type Kev_in6_addrlifetime struct {
	Ia6t_expire    U_int32_t
	Ia6t_preferred U_int32_t
	Ia6t_vltime    U_int32_t
	Ia6t_pltime    U_int32_t
}

// Kev_in6_data
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/kev_in6_data
type Kev_in6_data struct {
	Link_data     Net_event_data
	Ia_addr       Sockaddr_in6
	Ia_net        Sockaddr_in6
	Ia_dstaddr    Sockaddr_in6
	Ia_prefixmask Sockaddr_in6
	Ia_plen       U_int32_t
	Ia6_flags     U_int32_t
	Ia_lifetime   Kev_in6_addrlifetime
	Ia_mac        [6]uint8
}

// Kev_in_arpalive
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/kev_in_arpalive
type Kev_in_arpalive struct {
	Link_data Net_event_data
}

// Kev_in_arpfailure
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/kev_in_arpfailure
type Kev_in_arpfailure struct {
	Link_data Net_event_data
}

// Kev_in_collision
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/kev_in_collision
type Kev_in_collision struct {
	Link_data Net_event_data
	Ia_ipaddr In_addr
	Hw_len    U_char
}

// Kev_in_data
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/kev_in_data
type Kev_in_data struct {
	Link_data       Net_event_data
	Ia_addr         In_addr
	Ia_net          U_int32_t
	Ia_netmask      U_int32_t
	Ia_subnet       U_int32_t
	Ia_subnetmask   U_int32_t
	Ia_netbroadcast In_addr
	Ia_dstaddr      In_addr
}

// Kev_in_portinuse
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/kev_in_portinuse
type Kev_in_portinuse struct {
	Port     U_int16_t
	Req_pid  U_int32_t
	Reserved [2]U_int32_t
}

// Kev_msg
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/kev_msg
type Kev_msg struct {
	Vendor_code  U_int32_t        // The vendor code assigned by kev_vendor_code_find.
	Kev_class    U_int32_t        // The event's class.
	Kev_subclass U_int32_t        // The event's subclass.
	Event_code   U_int32_t        // The event's code.
	Dv           [5]Kev_d_vectors // An array of vectors describing additional data to be appended to the kernel event.

}

// Kev_netevent_apnfallbk_data
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/kev_netevent_apnfallbk_data
type Kev_netevent_apnfallbk_data struct {
	Epid  int32
	Euuid [16]uint8
}

// Kev_netevent_clat46_data
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/kev_netevent_clat46_data
type Kev_netevent_clat46_data struct {
	Clat46_event_code In6_clat46_evhdlr_code_t
	Epid              int32
	Euuid             [16]uint8
}

// Kev_request
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/kev_request
type Kev_request struct {
	Vendor_code  U_int32_t // All kernel events that don't match this vendor code will be ignored. KEV_ANY_VENDOR can be used to receive kernel events with any vendor code.
	Kev_class    U_int32_t // All kernel events that don't match this class will be ignored. KEV_ANY_CLASS can be used to receive kernel events with any class.
	Kev_subclass U_int32_t // All kernel events that don't match this subclass will be ignored. KEV_ANY_SUBCLASS can be used to receive kernel events with any subclass.

}

// Kev_vendor_code
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/kev_vendor_code
type Kev_vendor_code struct {
	Vendor_code   U_int32_t // After making the SIOCGKEVVENDOR ioctl call, this will be filled in with the vendor code if there is one.
	Vendor_string [200]int8 // A bundle style identifier.

}

// Kevent
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/kevent
type Kevent struct {
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
	storage [32]byte
}

// Ident returns the Ident field from the record's packed storage.
func (s *Kevent) Ident() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[0:8]))
}

// SetIdent updates the Ident field in the record's packed storage.
func (s *Kevent) SetIdent(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[0:8], uint64(v))
}

// Filter returns the Filter field from the record's packed storage.
func (s *Kevent) Filter() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[8:10]))
}

// SetFilter updates the Filter field in the record's packed storage.
func (s *Kevent) SetFilter(v int16) {
	binary.NativeEndian.PutUint16(s.storage[8:10], uint16(v))
}

// Flags returns the Flags field from the record's packed storage.
func (s *Kevent) Flags() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[10:12]))
}

// SetFlags updates the Flags field in the record's packed storage.
func (s *Kevent) SetFlags(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[10:12], uint16(v))
}

// Fflags returns the Fflags field from the record's packed storage.
func (s *Kevent) Fflags() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[12:16]))
}

// SetFflags updates the Fflags field in the record's packed storage.
func (s *Kevent) SetFflags(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[12:16], uint32(v))
}

// Data returns the Data field from the record's packed storage.
func (s *Kevent) Data() int {
	return int(binary.NativeEndian.Uint64(s.storage[16:24]))
}

// SetData updates the Data field in the record's packed storage.
func (s *Kevent) SetData(v int) {
	binary.NativeEndian.PutUint64(s.storage[16:24], uint64(v))
}

// Udata returns the Udata field from the record's packed storage.
func (s *Kevent) Udata() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[24:32]))
}

// SetUdata updates the Udata field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *Kevent) SetUdata(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[24:32], uint64(v))
}

// Kevent64_s
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/kevent64_s
type Kevent64_s struct {
	Ident  uint64
	Filter int16
	Flags  uint16
	Fflags uint32
	Data   int64
	Udata  uint64
	Ext    [2]uint64
}

// Ledger_entry_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ledger_entry_info
type Ledger_entry_info struct {
	Lei_balance       int64
	Lei_credit        int64
	Lei_debit         int64
	Lei_limit         uint64
	Lei_refill_period uint64
	Lei_last_refill   uint64
}

// Ledger_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ledger_info
type Ledger_info struct {
	Li_name    [32]int8
	Li_id      int64
	Li_entries int64
}

// Ledger_limit_args
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ledger_limit_args
type Ledger_limit_args struct {
	Lla_name          [32]int8
	Lla_limit         uint64
	Lla_refill_period uint64
}

// Ledger_template_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ledger_template_info
type Ledger_template_info struct {
	Lti_name  [32]int8
	Lti_group [32]int8
	Lti_units [32]int8
}

// Linger
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/linger
type Linger struct {
	L_onoff  int32
	L_linger int32
}

// Linkedit_data_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/linkedit_data_command
type Linkedit_data_command struct {
	Cmd      uint32
	Cmdsize  uint32
	Dataoff  uint32
	Datasize uint32
}

// Linker_option_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/linker_option_command
type Linker_option_command struct {
	Cmd     uint32
	Cmdsize uint32
	Count   uint32
}

// Llc
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/llc
type Llc struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [8]byte
}

// Llc_dsap returns the Llc_dsap field from the record's packed storage.
func (s *Llc) Llc_dsap() U_int8_t {
	return *(*U_int8_t)(unsafe.Pointer(&s.storage[0]))
}

// SetLlc_dsap updates the Llc_dsap field in the record's packed storage.
func (s *Llc) SetLlc_dsap(v U_int8_t) {
	*(*U_int8_t)(unsafe.Pointer(&s.storage[0])) = v
}

// Llc_ssap returns the Llc_ssap field from the record's packed storage.
func (s *Llc) Llc_ssap() U_int8_t {
	return *(*U_int8_t)(unsafe.Pointer(&s.storage[1]))
}

// SetLlc_ssap updates the Llc_ssap field in the record's packed storage.
func (s *Llc) SetLlc_ssap(v U_int8_t) {
	*(*U_int8_t)(unsafe.Pointer(&s.storage[1])) = v
}

// Llc_un returns the Llc_un field from the record's packed storage.
func (s *Llc) Llc_un() [6]byte {
	return *(*[6]byte)(unsafe.Pointer(&s.storage[2]))
}

// SetLlc_un updates the Llc_un field in the record's packed storage.
func (s *Llc) SetLlc_un(v [6]byte) {
	*(*[6]byte)(unsafe.Pointer(&s.storage[2])) = v
}

// Load_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/load_command
type Load_command struct {
	Cmd     uint32
	Cmdsize uint32
}

// Lockd_ans
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/lockd_ans
type Lockd_ans struct {
	La_version int32
	La_errno   int32
	La_xid     U_int64_t
	La_flags   int32
	La_pid     int32
	La_start   int64
	La_len     int64
	La_fh_len  int32
	La_fh      [64]U_int8_t
}

// Lockd_notify
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/lockd_notify
type Lockd_notify struct {
	Ln_version   int32
	Ln_flags     int32
	Ln_pad       int32
	Ln_addrcount int32
	Ln_addr      Pointer
}

// Lockf
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/lockf
type Lockf struct {
	Lf_flags int16
	Lf_type  int16
	Lf_start int64
	Lf_end   int64
	Lf_id    Caddr_t
	Lf_head  *Lockf
	Lf_vnode unsafe.Pointer
	Lf_next  *Lockf
	Lf_blkhd Locklist
	Lf_owner unsafe.Pointer
}

// Locklist
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/locklist
type Locklist struct {
	Tqh_first *Lockf
	Tqh_last  **Lockf
}

// Log2phys
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/log2phys
type Log2phys struct {
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
	storage [20]byte
}

// L2p_flags returns the L2p_flags field from the record's packed storage.
func (s *Log2phys) L2p_flags() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetL2p_flags updates the L2p_flags field in the record's packed storage.
func (s *Log2phys) SetL2p_flags(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// L2p_contigbytes returns the L2p_contigbytes field from the record's packed storage.
func (s *Log2phys) L2p_contigbytes() int64 {
	return int64(binary.NativeEndian.Uint64(s.storage[4:12]))
}

// SetL2p_contigbytes updates the L2p_contigbytes field in the record's packed storage.
func (s *Log2phys) SetL2p_contigbytes(v int64) {
	binary.NativeEndian.PutUint64(s.storage[4:12], uint64(v))
}

// L2p_devoffset returns the L2p_devoffset field from the record's packed storage.
func (s *Log2phys) L2p_devoffset() int64 {
	return int64(binary.NativeEndian.Uint64(s.storage[12:20]))
}

// SetL2p_devoffset updates the L2p_devoffset field in the record's packed storage.
func (s *Log2phys) SetL2p_devoffset(v int64) {
	binary.NativeEndian.PutUint64(s.storage[12:20], uint64(v))
}

// Ltchars
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ltchars
type Ltchars struct {
	T_suspc  int8
	T_dsuspc int8
	T_rprntc int8
	T_flushc int8
	T_werasc int8
	T_lnextc int8
}

// Mach_assert_3x
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/mach_assert_3x
type Mach_assert_3x struct {
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
	storage [36]byte
}

// Hdr returns the Hdr field from the record's packed storage.
func (s *Mach_assert_3x) Hdr() Mach_assert_hdr {
	return *(*Mach_assert_hdr)(unsafe.Pointer(&s.storage[0]))
}

// SetHdr updates the Hdr field in the record's packed storage.
func (s *Mach_assert_3x) SetHdr(v Mach_assert_hdr) {
	*(*Mach_assert_hdr)(unsafe.Pointer(&s.storage[0])) = v
}

// A returns the A field from the record's packed storage.
func (s *Mach_assert_3x) A() *byte {
	return *(**byte)(unsafe.Pointer(&s.storage[12]))
}

// SetA updates the A field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *Mach_assert_3x) SetA(v *byte) {
	*(**byte)(unsafe.Pointer(&s.storage[12])) = v
}

// Op returns the Op field from the record's packed storage.
func (s *Mach_assert_3x) Op() *byte {
	return *(**byte)(unsafe.Pointer(&s.storage[20]))
}

// SetOp updates the Op field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *Mach_assert_3x) SetOp(v *byte) {
	*(**byte)(unsafe.Pointer(&s.storage[20])) = v
}

// B returns the B field from the record's packed storage.
func (s *Mach_assert_3x) B() *byte {
	return *(**byte)(unsafe.Pointer(&s.storage[28]))
}

// SetB updates the B field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *Mach_assert_3x) SetB(v *byte) {
	*(**byte)(unsafe.Pointer(&s.storage[28])) = v
}

// Mach_assert_default
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/mach_assert_default
type Mach_assert_default struct {
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
	storage [20]byte
}

// Hdr returns the Hdr field from the record's packed storage.
func (s *Mach_assert_default) Hdr() Mach_assert_hdr {
	return *(*Mach_assert_hdr)(unsafe.Pointer(&s.storage[0]))
}

// SetHdr updates the Hdr field in the record's packed storage.
func (s *Mach_assert_default) SetHdr(v Mach_assert_hdr) {
	*(*Mach_assert_hdr)(unsafe.Pointer(&s.storage[0])) = v
}

// Expr returns the Expr field from the record's packed storage.
func (s *Mach_assert_default) Expr() *byte {
	return *(**byte)(unsafe.Pointer(&s.storage[12]))
}

// SetExpr updates the Expr field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *Mach_assert_default) SetExpr(v *byte) {
	*(**byte)(unsafe.Pointer(&s.storage[12])) = v
}

// Mach_assert_hdr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/mach_assert_hdr
type Mach_assert_hdr struct {
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
	storage [12]byte
}

// Type returns the Type field from the record's packed storage.
func (s *Mach_assert_hdr) Type() Mach_assert_type_t {
	return *(*Mach_assert_type_t)(unsafe.Pointer(&s.storage[0]))
}

// SetType updates the Type field in the record's packed storage.
func (s *Mach_assert_hdr) SetType(v Mach_assert_type_t) {
	*(*Mach_assert_type_t)(unsafe.Pointer(&s.storage[0])) = v
}

// Lineno returns the Lineno bitfield from the record's packed storage.
func (s *Mach_assert_hdr) Lineno() uint32 {
	return uint32((binary.NativeEndian.Uint32(s.storage[1:5]) >> 0) & 0xffffff)
}

// SetLineno updates the Lineno bitfield in the record's packed storage.
func (s *Mach_assert_hdr) SetLineno(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[1:5], (binary.NativeEndian.Uint32(s.storage[1:5])&^uint32(0xffffff<<0))|uint32((uint32(v)&0xffffff)<<0))
}

// Filename returns the Filename field from the record's packed storage.
func (s *Mach_assert_hdr) Filename() *byte {
	return *(**byte)(unsafe.Pointer(&s.storage[4]))
}

// SetFilename updates the Filename field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *Mach_assert_hdr) SetFilename(v *byte) {
	*(**byte)(unsafe.Pointer(&s.storage[4])) = v
}

// Mach_core_details
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/mach_core_details
type Mach_core_details struct {
	Gzip_offset uint64
	Gzip_length uint64
	Core_name   [16]int8
}

// Mach_core_details_v2
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/mach_core_details_v2
type Mach_core_details_v2 struct {
	Flags     uint64
	Offset    uint64
	Length    uint64
	Core_name [16]int8
}

// Mach_core_fileheader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/mach_core_fileheader
type Mach_core_fileheader struct {
	Signature  uint64
	Log_offset uint64
	Log_length uint64
	Num_files  uint64
	Files      [16]Mach_core_details
}

// Mach_core_fileheader_base
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/mach_core_fileheader_base
type Mach_core_fileheader_base struct {
	Signature uint64
	Version   uint32
}

// Mach_core_fileheader_v2
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/mach_core_fileheader_v2
type Mach_core_fileheader_v2 struct {
	Signature      uint64
	Version        uint32
	Flags          uint64
	Pub_key_offset uint64
	Pub_key_length uint16
	Log_offset     uint64
	Log_length     uint64
	Num_files      uint64
}

// Mach_header
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/mach_header
type Mach_header struct {
	Magic      uint32 // An integer containing a value identifying this file as a 32-bit Mach-O file. Use the constant `MH_MAGIC` if the file is intended for use on a CPU with the same endianness as the computer on which the compiler is running. The constant `MH_CIGAM` can be used when the byte ordering scheme of the target machine is the reverse of the host CPU.
	Cputype    int32
	Cpusubtype int32
	Filetype   uint32
	Ncmds      uint32
	Sizeofcmds uint32
	Flags      uint32
}

// Mach_header_64
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/mach_header_64
type Mach_header_64 struct {
	Magic      uint32 // An integer containing a value identifying this file as a 64-bit Mach-O file. Use the constant `MH_MAGIC_64` if the file is intended for use on a CPU with the same endianness as the computer on which the compiler is running. The constant `MH_CIGAM_64` can be used when the byte ordering scheme of the target machine is the reverse of the host CPU.
	Cputype    int32
	Cpusubtype int32 // An integer specifying the exact model of the CPU. To run on all PowerPC processors supported by the macOS kernel, this should be set to `CPU_SUBTYPE_POWERPC_ALL`.
	Filetype   uint32
	Ncmds      uint32
	Sizeofcmds uint32
	Flags      uint32 // An integer containing a set of bit flags that indicate the state of certain optional features of the Mach-O file format. These are the masks you can use to manipulate this field:
	Reserved   uint32
}

// Macos_panic_header
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/macos_panic_header
type Macos_panic_header struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [56]byte
}

// Mph_magic returns the Mph_magic field from the record's packed storage.
func (s *Macos_panic_header) Mph_magic() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetMph_magic updates the Mph_magic field in the record's packed storage.
func (s *Macos_panic_header) SetMph_magic(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// Mph_crc returns the Mph_crc field from the record's packed storage.
func (s *Macos_panic_header) Mph_crc() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetMph_crc updates the Mph_crc field in the record's packed storage.
func (s *Macos_panic_header) SetMph_crc(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// Mph_version returns the Mph_version field from the record's packed storage.
func (s *Macos_panic_header) Mph_version() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetMph_version updates the Mph_version field in the record's packed storage.
func (s *Macos_panic_header) SetMph_version(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[8:12], uint32(v))
}

// Mph_padding returns the Mph_padding field from the record's packed storage.
func (s *Macos_panic_header) Mph_padding() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[12:16]))
}

// SetMph_padding updates the Mph_padding field in the record's packed storage.
func (s *Macos_panic_header) SetMph_padding(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[12:16], uint32(v))
}

// Mph_panic_flags returns the Mph_panic_flags field from the record's packed storage.
func (s *Macos_panic_header) Mph_panic_flags() Mph_panic_flags_t {
	return Mph_panic_flags_t(binary.NativeEndian.Uint64(s.storage[16:24]))
}

// SetMph_panic_flags updates the Mph_panic_flags field in the record's packed storage.
func (s *Macos_panic_header) SetMph_panic_flags(v Mph_panic_flags_t) {
	binary.NativeEndian.PutUint64(s.storage[16:24], uint64(v))
}

// Mph_panic_log_offset returns the Mph_panic_log_offset field from the record's packed storage.
func (s *Macos_panic_header) Mph_panic_log_offset() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[24:28]))
}

// SetMph_panic_log_offset updates the Mph_panic_log_offset field in the record's packed storage.
func (s *Macos_panic_header) SetMph_panic_log_offset(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[24:28], uint32(v))
}

// Mph_panic_log_len returns the Mph_panic_log_len field from the record's packed storage.
func (s *Macos_panic_header) Mph_panic_log_len() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[28:32]))
}

// SetMph_panic_log_len updates the Mph_panic_log_len field in the record's packed storage.
func (s *Macos_panic_header) SetMph_panic_log_len(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[28:32], uint32(v))
}

// Mph_stackshot_offset returns the Mph_stackshot_offset field from the record's packed storage.
func (s *Macos_panic_header) Mph_stackshot_offset() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[32:36]))
}

// SetMph_stackshot_offset updates the Mph_stackshot_offset field in the record's packed storage.
func (s *Macos_panic_header) SetMph_stackshot_offset(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[32:36], uint32(v))
}

// Mph_stackshot_len returns the Mph_stackshot_len field from the record's packed storage.
func (s *Macos_panic_header) Mph_stackshot_len() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[36:40]))
}

// SetMph_stackshot_len updates the Mph_stackshot_len field in the record's packed storage.
func (s *Macos_panic_header) SetMph_stackshot_len(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[36:40], uint32(v))
}

// Mph_other_log_offset returns the Mph_other_log_offset field from the record's packed storage.
func (s *Macos_panic_header) Mph_other_log_offset() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[40:44]))
}

// SetMph_other_log_offset updates the Mph_other_log_offset field in the record's packed storage.
func (s *Macos_panic_header) SetMph_other_log_offset(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[40:44], uint32(v))
}

// Mph_other_log_len returns the Mph_other_log_len field from the record's packed storage.
func (s *Macos_panic_header) Mph_other_log_len() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[44:48]))
}

// SetMph_other_log_len updates the Mph_other_log_len field in the record's packed storage.
func (s *Macos_panic_header) SetMph_other_log_len(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[44:48], uint32(v))
}

// Mph_roots_installed returns the Mph_roots_installed field from the record's packed storage.
func (s *Macos_panic_header) Mph_roots_installed() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[48:56]))
}

// SetMph_roots_installed updates the Mph_roots_installed field in the record's packed storage.
func (s *Macos_panic_header) SetMph_roots_installed(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[48:56], uint64(v))
}

// Mbstat
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/mbstat
type Mbstat struct {
	M_mbufs        U_int32_t
	M_clusters     U_int32_t
	M_spare        U_int32_t
	M_clfree       U_int32_t
	M_drops        U_int32_t
	M_wait         U_int32_t
	M_drain        U_int32_t
	M_mtypes       [256]U_short
	M_mcfail       U_int32_t
	M_mpfail       U_int32_t
	M_msize        U_int32_t
	M_mclbytes     U_int32_t
	M_minclsize    U_int32_t
	M_mlen         U_int32_t
	M_mhlen        U_int32_t
	M_bigclusters  U_int32_t
	M_bigclfree    U_int32_t
	M_bigmclbytes  U_int32_t
	M_forcedefunct U_int32_t
}

// Mbuf_stat
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/mbuf_stat
type Mbuf_stat struct {
	Mbufs       U_int32_t    // Number of mbufs (free or otherwise).
	Clusters    U_int32_t    // Number of clusters (free or otherwise).
	Clfree      U_int32_t    // Number of free clusters.
	Drops       U_int32_t    // Number of times allocation failed.
	Wait        U_int32_t    // Number of times allocation blocked.
	Drain       U_int32_t    // Number of times protocol drain functions were called.
	Mtypes      [256]U_short // An array of counts of each type of mbuf allocated.
	Mcfail      U_int32_t    // Number of times m_copym failed.
	Mpfail      U_int32_t    // Number of times m_pullup failed.
	Msize       U_int32_t    // Length of an mbuf.
	Mclbytes    U_int32_t    // Length of an mbuf cluster.
	Minclsize   U_int32_t    // Minimum length of data to allocate a cluster. Anything smaller than this should be placed in chained mbufs.
	Mlen        U_int32_t    // Length of data in an mbuf.
	Mhlen       U_int32_t    // Length of data in an mbuf with a packet header.
	Bigclusters U_int32_t    // Number of big clusters.
	Bigclfree   U_int32_t    // Number of unused big clusters.
	Bigmclbytes U_int32_t    // Length of a big mbuf cluster.

}

// Mem_and_io_snapshot
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/mem_and_io_snapshot
type Mem_and_io_snapshot struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [61]byte
}

// Snapshot_magic returns the Snapshot_magic field from the record's packed storage.
func (s *Mem_and_io_snapshot) Snapshot_magic() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetSnapshot_magic updates the Snapshot_magic field in the record's packed storage.
func (s *Mem_and_io_snapshot) SetSnapshot_magic(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// Free_pages returns the Free_pages field from the record's packed storage.
func (s *Mem_and_io_snapshot) Free_pages() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetFree_pages updates the Free_pages field in the record's packed storage.
func (s *Mem_and_io_snapshot) SetFree_pages(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// Active_pages returns the Active_pages field from the record's packed storage.
func (s *Mem_and_io_snapshot) Active_pages() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetActive_pages updates the Active_pages field in the record's packed storage.
func (s *Mem_and_io_snapshot) SetActive_pages(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[8:12], uint32(v))
}

// Inactive_pages returns the Inactive_pages field from the record's packed storage.
func (s *Mem_and_io_snapshot) Inactive_pages() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[12:16]))
}

// SetInactive_pages updates the Inactive_pages field in the record's packed storage.
func (s *Mem_and_io_snapshot) SetInactive_pages(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[12:16], uint32(v))
}

// Purgeable_pages returns the Purgeable_pages field from the record's packed storage.
func (s *Mem_and_io_snapshot) Purgeable_pages() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[16:20]))
}

// SetPurgeable_pages updates the Purgeable_pages field in the record's packed storage.
func (s *Mem_and_io_snapshot) SetPurgeable_pages(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[16:20], uint32(v))
}

// Wired_pages returns the Wired_pages field from the record's packed storage.
func (s *Mem_and_io_snapshot) Wired_pages() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[20:24]))
}

// SetWired_pages updates the Wired_pages field in the record's packed storage.
func (s *Mem_and_io_snapshot) SetWired_pages(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[20:24], uint32(v))
}

// Speculative_pages returns the Speculative_pages field from the record's packed storage.
func (s *Mem_and_io_snapshot) Speculative_pages() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[24:28]))
}

// SetSpeculative_pages updates the Speculative_pages field in the record's packed storage.
func (s *Mem_and_io_snapshot) SetSpeculative_pages(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[24:28], uint32(v))
}

// Throttled_pages returns the Throttled_pages field from the record's packed storage.
func (s *Mem_and_io_snapshot) Throttled_pages() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[28:32]))
}

// SetThrottled_pages updates the Throttled_pages field in the record's packed storage.
func (s *Mem_and_io_snapshot) SetThrottled_pages(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[28:32], uint32(v))
}

// Filebacked_pages returns the Filebacked_pages field from the record's packed storage.
func (s *Mem_and_io_snapshot) Filebacked_pages() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[32:36]))
}

// SetFilebacked_pages updates the Filebacked_pages field in the record's packed storage.
func (s *Mem_and_io_snapshot) SetFilebacked_pages(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[32:36], uint32(v))
}

// Compressions returns the Compressions field from the record's packed storage.
func (s *Mem_and_io_snapshot) Compressions() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[36:40]))
}

// SetCompressions updates the Compressions field in the record's packed storage.
func (s *Mem_and_io_snapshot) SetCompressions(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[36:40], uint32(v))
}

// Decompressions returns the Decompressions field from the record's packed storage.
func (s *Mem_and_io_snapshot) Decompressions() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[40:44]))
}

// SetDecompressions updates the Decompressions field in the record's packed storage.
func (s *Mem_and_io_snapshot) SetDecompressions(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[40:44], uint32(v))
}

// Compressor_size returns the Compressor_size field from the record's packed storage.
func (s *Mem_and_io_snapshot) Compressor_size() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[44:48]))
}

// SetCompressor_size updates the Compressor_size field in the record's packed storage.
func (s *Mem_and_io_snapshot) SetCompressor_size(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[44:48], uint32(v))
}

// Busy_buffer_count returns the Busy_buffer_count field from the record's packed storage.
func (s *Mem_and_io_snapshot) Busy_buffer_count() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[48:52]))
}

// SetBusy_buffer_count updates the Busy_buffer_count field in the record's packed storage.
func (s *Mem_and_io_snapshot) SetBusy_buffer_count(v int32) {
	binary.NativeEndian.PutUint32(s.storage[48:52], uint32(v))
}

// Pages_wanted returns the Pages_wanted field from the record's packed storage.
func (s *Mem_and_io_snapshot) Pages_wanted() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[52:56]))
}

// SetPages_wanted updates the Pages_wanted field in the record's packed storage.
func (s *Mem_and_io_snapshot) SetPages_wanted(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[52:56], uint32(v))
}

// Pages_reclaimed returns the Pages_reclaimed field from the record's packed storage.
func (s *Mem_and_io_snapshot) Pages_reclaimed() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[56:60]))
}

// SetPages_reclaimed updates the Pages_reclaimed field in the record's packed storage.
func (s *Mem_and_io_snapshot) SetPages_reclaimed(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[56:60], uint32(v))
}

// Pages_wanted_reclaimed_valid returns the Pages_wanted_reclaimed_valid field from the record's packed storage.
func (s *Mem_and_io_snapshot) Pages_wanted_reclaimed_valid() uint8 {
	return uint8(s.storage[60])
}

// SetPages_wanted_reclaimed_valid updates the Pages_wanted_reclaimed_valid field in the record's packed storage.
func (s *Mem_and_io_snapshot) SetPages_wanted_reclaimed_valid(v uint8) {
	s.storage[60] = uint8(v)
}

// Memory_error_notification_subsystem
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/memory_error_notification_subsystem-b6h
type Memory_error_notification_subsystem struct {
	Server   unsafe.Pointer
	Start    Mach_msg_id_t
	End      Mach_msg_id_t
	Maxsize  uint32
	Reserved Vm_address_t
	Routine  unsafe.Pointer
}

// Micro_snapshot
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/micro_snapshot
type Micro_snapshot struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [27]byte
}

// Snapshot_magic returns the Snapshot_magic field from the record's packed storage.
func (s *Micro_snapshot) Snapshot_magic() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetSnapshot_magic updates the Snapshot_magic field in the record's packed storage.
func (s *Micro_snapshot) SetSnapshot_magic(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// Ms_cpu returns the Ms_cpu field from the record's packed storage.
func (s *Micro_snapshot) Ms_cpu() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetMs_cpu updates the Ms_cpu field in the record's packed storage.
func (s *Micro_snapshot) SetMs_cpu(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// Ms_time returns the Ms_time field from the record's packed storage.
func (s *Micro_snapshot) Ms_time() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[8:16]))
}

// SetMs_time updates the Ms_time field in the record's packed storage.
func (s *Micro_snapshot) SetMs_time(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[8:16], uint64(v))
}

// Ms_time_microsecs returns the Ms_time_microsecs field from the record's packed storage.
func (s *Micro_snapshot) Ms_time_microsecs() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[16:24]))
}

// SetMs_time_microsecs updates the Ms_time_microsecs field in the record's packed storage.
func (s *Micro_snapshot) SetMs_time_microsecs(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[16:24], uint64(v))
}

// Ms_flags returns the Ms_flags field from the record's packed storage.
func (s *Micro_snapshot) Ms_flags() uint8 {
	return uint8(s.storage[24])
}

// SetMs_flags updates the Ms_flags field in the record's packed storage.
func (s *Micro_snapshot) SetMs_flags(v uint8) {
	s.storage[24] = uint8(v)
}

// Ms_opaque_flags returns the Ms_opaque_flags field from the record's packed storage.
func (s *Micro_snapshot) Ms_opaque_flags() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[25:27]))
}

// SetMs_opaque_flags updates the Ms_opaque_flags field in the record's packed storage.
func (s *Micro_snapshot) SetMs_opaque_flags(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[25:27], uint16(v))
}

// Mld_hdr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/mld_hdr
type Mld_hdr struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [24]byte
}

// Mld_icmp6_hdr returns the Mld_icmp6_hdr field from the record's packed storage.
func (s *Mld_hdr) Mld_icmp6_hdr() Icmp6_hdr {
	return *(*Icmp6_hdr)(unsafe.Pointer(&s.storage[0]))
}

// SetMld_icmp6_hdr updates the Mld_icmp6_hdr field in the record's packed storage.
func (s *Mld_hdr) SetMld_icmp6_hdr(v Icmp6_hdr) {
	*(*Icmp6_hdr)(unsafe.Pointer(&s.storage[0])) = v
}

// Mld_addr returns the Mld_addr field from the record's packed storage.
func (s *Mld_hdr) Mld_addr() [4]uint32 {
	return *(*[4]uint32)(unsafe.Pointer(&s.storage[8]))
}

// SetMld_addr updates the Mld_addr field in the record's packed storage.
func (s *Mld_hdr) SetMld_addr(v [4]uint32) {
	*(*[4]uint32)(unsafe.Pointer(&s.storage[8])) = v
}

// Mptcp_itf_stats
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/mptcp_itf_stats
type Mptcp_itf_stats struct {
	Ifindex            U_short
	Switches           uint16
	bitfield2          uint8
	Mpis_txbytes       uint64
	Mpis_rxbytes       uint64
	Mpis_wifi_txbytes  uint64
	Mpis_wifi_rxbytes  uint64
	Mpis_wired_txbytes uint64
	Mpis_wired_rxbytes uint64
	Mpis_cell_txbytes  uint64
	Mpis_cell_rxbytes  uint64
}

// Is_expensive returns the Is_expensive bitfield.
func (s *Mptcp_itf_stats) Is_expensive() uint8 {
	return (s.bitfield2 >> 0) & ((1 << 1) - 1)
}

// SetIs_expensive updates the Is_expensive bitfield.
func (s *Mptcp_itf_stats) SetIs_expensive(v uint8) {
	const mask uint8 = (1 << 1) - 1
	s.bitfield2 = (s.bitfield2 &^ (mask << 0)) | ((v & mask) << 0)
}

// Msghdr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/msghdr
type Msghdr struct {
	Msg_name       unsafe.Pointer
	Msg_namelen    uint32
	Msg_iov        *Iovec
	Msg_iovlen     int32
	Msg_control    unsafe.Pointer
	Msg_controllen uint32
	Msg_flags      int32
}

// Msginfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/msginfo-lp5
type Msginfo struct {
	Msgmax int32
	Msgmni int32
	Msgmnb int32
	Msgtql int32
	Msgssz int32
	Msgseg int32
}

// Msgmap
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/msgmap
type Msgmap struct {
	Next int16
}

// Msqid_kernel
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/msqid_kernel
type Msqid_kernel struct {
	U     User_msqid_ds
	Label unsafe.Pointer
}

// Mwl_info_hdr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/mwl_info_hdr
type Mwl_info_hdr struct {
	Mwli_version        uint32
	Mwli_page_size      uint16
	Mwli_pointer_format uint16
	Mwli_binds_offset   uint32
	Mwli_binds_count    uint32
	Mwli_chains_offset  uint32
	Mwli_chains_size    uint32
	Mwli_slide          uint64
	Mwli_image_address  uint64
}

// Mwl_region
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/mwl_region
type Mwl_region struct {
	Mwlr_fd          int32
	Mwlr_protections Vm_prot_t
	Mwlr_file_offset uint64
	Mwlr_address     Mach_vm_address_t
	Mwlr_size        Mach_vm_size_t
}

// Mymsg
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/mymsg
type Mymsg struct {
	Mtype int
	Mtext [1]int8
}

// Nd_neighbor_advert
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nd_neighbor_advert
type Nd_neighbor_advert struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [24]byte
}

// Nd_na_hdr returns the Nd_na_hdr field from the record's packed storage.
func (s *Nd_neighbor_advert) Nd_na_hdr() Icmp6_hdr {
	return *(*Icmp6_hdr)(unsafe.Pointer(&s.storage[0]))
}

// SetNd_na_hdr updates the Nd_na_hdr field in the record's packed storage.
func (s *Nd_neighbor_advert) SetNd_na_hdr(v Icmp6_hdr) {
	*(*Icmp6_hdr)(unsafe.Pointer(&s.storage[0])) = v
}

// Nd_na_target returns the Nd_na_target field from the record's packed storage.
func (s *Nd_neighbor_advert) Nd_na_target() [4]uint32 {
	return *(*[4]uint32)(unsafe.Pointer(&s.storage[8]))
}

// SetNd_na_target updates the Nd_na_target field in the record's packed storage.
func (s *Nd_neighbor_advert) SetNd_na_target(v [4]uint32) {
	*(*[4]uint32)(unsafe.Pointer(&s.storage[8])) = v
}

// Nd_neighbor_solicit
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nd_neighbor_solicit
type Nd_neighbor_solicit struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [24]byte
}

// Nd_ns_hdr returns the Nd_ns_hdr field from the record's packed storage.
func (s *Nd_neighbor_solicit) Nd_ns_hdr() Icmp6_hdr {
	return *(*Icmp6_hdr)(unsafe.Pointer(&s.storage[0]))
}

// SetNd_ns_hdr updates the Nd_ns_hdr field in the record's packed storage.
func (s *Nd_neighbor_solicit) SetNd_ns_hdr(v Icmp6_hdr) {
	*(*Icmp6_hdr)(unsafe.Pointer(&s.storage[0])) = v
}

// Nd_ns_target returns the Nd_ns_target field from the record's packed storage.
func (s *Nd_neighbor_solicit) Nd_ns_target() [4]uint32 {
	return *(*[4]uint32)(unsafe.Pointer(&s.storage[8]))
}

// SetNd_ns_target updates the Nd_ns_target field in the record's packed storage.
func (s *Nd_neighbor_solicit) SetNd_ns_target(v [4]uint32) {
	*(*[4]uint32)(unsafe.Pointer(&s.storage[8])) = v
}

// Nd_opt_dnssl
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nd_opt_dnssl
type Nd_opt_dnssl struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [16]byte
}

// Nd_opt_dnssl_type returns the Nd_opt_dnssl_type field from the record's packed storage.
func (s *Nd_opt_dnssl) Nd_opt_dnssl_type() U_int8_t {
	return *(*U_int8_t)(unsafe.Pointer(&s.storage[0]))
}

// SetNd_opt_dnssl_type updates the Nd_opt_dnssl_type field in the record's packed storage.
func (s *Nd_opt_dnssl) SetNd_opt_dnssl_type(v U_int8_t) {
	*(*U_int8_t)(unsafe.Pointer(&s.storage[0])) = v
}

// Nd_opt_dnssl_len returns the Nd_opt_dnssl_len field from the record's packed storage.
func (s *Nd_opt_dnssl) Nd_opt_dnssl_len() U_int8_t {
	return *(*U_int8_t)(unsafe.Pointer(&s.storage[1]))
}

// SetNd_opt_dnssl_len updates the Nd_opt_dnssl_len field in the record's packed storage.
func (s *Nd_opt_dnssl) SetNd_opt_dnssl_len(v U_int8_t) {
	*(*U_int8_t)(unsafe.Pointer(&s.storage[1])) = v
}

// Nd_opt_dnssl_reserved returns the Nd_opt_dnssl_reserved field from the record's packed storage.
func (s *Nd_opt_dnssl) Nd_opt_dnssl_reserved() U_int16_t {
	return U_int16_t(binary.NativeEndian.Uint16(s.storage[2:4]))
}

// SetNd_opt_dnssl_reserved updates the Nd_opt_dnssl_reserved field in the record's packed storage.
func (s *Nd_opt_dnssl) SetNd_opt_dnssl_reserved(v U_int16_t) {
	binary.NativeEndian.PutUint16(s.storage[2:4], uint16(v))
}

// Nd_opt_dnssl_lifetime returns the Nd_opt_dnssl_lifetime field from the record's packed storage.
func (s *Nd_opt_dnssl) Nd_opt_dnssl_lifetime() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetNd_opt_dnssl_lifetime updates the Nd_opt_dnssl_lifetime field in the record's packed storage.
func (s *Nd_opt_dnssl) SetNd_opt_dnssl_lifetime(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// Nd_opt_dnssl_domains returns the Nd_opt_dnssl_domains field from the record's packed storage.
func (s *Nd_opt_dnssl) Nd_opt_dnssl_domains() [8]U_int8_t {
	return *(*[8]U_int8_t)(unsafe.Pointer(&s.storage[8]))
}

// SetNd_opt_dnssl_domains updates the Nd_opt_dnssl_domains field in the record's packed storage.
func (s *Nd_opt_dnssl) SetNd_opt_dnssl_domains(v [8]U_int8_t) {
	*(*[8]U_int8_t)(unsafe.Pointer(&s.storage[8])) = v
}

// Nd_opt_hdr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nd_opt_hdr
type Nd_opt_hdr struct {
	Nd_opt_type U_int8_t
	Nd_opt_len  U_int8_t
}

// Nd_opt_mtu
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nd_opt_mtu
type Nd_opt_mtu struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [8]byte
}

// Nd_opt_mtu_type returns the Nd_opt_mtu_type field from the record's packed storage.
func (s *Nd_opt_mtu) Nd_opt_mtu_type() U_int8_t {
	return *(*U_int8_t)(unsafe.Pointer(&s.storage[0]))
}

// SetNd_opt_mtu_type updates the Nd_opt_mtu_type field in the record's packed storage.
func (s *Nd_opt_mtu) SetNd_opt_mtu_type(v U_int8_t) {
	*(*U_int8_t)(unsafe.Pointer(&s.storage[0])) = v
}

// Nd_opt_mtu_len returns the Nd_opt_mtu_len field from the record's packed storage.
func (s *Nd_opt_mtu) Nd_opt_mtu_len() U_int8_t {
	return *(*U_int8_t)(unsafe.Pointer(&s.storage[1]))
}

// SetNd_opt_mtu_len updates the Nd_opt_mtu_len field in the record's packed storage.
func (s *Nd_opt_mtu) SetNd_opt_mtu_len(v U_int8_t) {
	*(*U_int8_t)(unsafe.Pointer(&s.storage[1])) = v
}

// Nd_opt_mtu_reserved returns the Nd_opt_mtu_reserved field from the record's packed storage.
func (s *Nd_opt_mtu) Nd_opt_mtu_reserved() U_int16_t {
	return U_int16_t(binary.NativeEndian.Uint16(s.storage[2:4]))
}

// SetNd_opt_mtu_reserved updates the Nd_opt_mtu_reserved field in the record's packed storage.
func (s *Nd_opt_mtu) SetNd_opt_mtu_reserved(v U_int16_t) {
	binary.NativeEndian.PutUint16(s.storage[2:4], uint16(v))
}

// Nd_opt_mtu_mtu returns the Nd_opt_mtu_mtu field from the record's packed storage.
func (s *Nd_opt_mtu) Nd_opt_mtu_mtu() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetNd_opt_mtu_mtu updates the Nd_opt_mtu_mtu field in the record's packed storage.
func (s *Nd_opt_mtu) SetNd_opt_mtu_mtu(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// Nd_opt_nonce
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nd_opt_nonce
type Nd_opt_nonce struct {
	Nd_opt_nonce_type U_int8_t
	Nd_opt_nonce_len  U_int8_t
	Nd_opt_nonce      [6]U_int8_t
}

// Nd_opt_pref64
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nd_opt_pref64
type Nd_opt_pref64 struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [16]byte
}

// Nd_opt_pref64_type returns the Nd_opt_pref64_type field from the record's packed storage.
func (s *Nd_opt_pref64) Nd_opt_pref64_type() U_int8_t {
	return *(*U_int8_t)(unsafe.Pointer(&s.storage[0]))
}

// SetNd_opt_pref64_type updates the Nd_opt_pref64_type field in the record's packed storage.
func (s *Nd_opt_pref64) SetNd_opt_pref64_type(v U_int8_t) {
	*(*U_int8_t)(unsafe.Pointer(&s.storage[0])) = v
}

// Nd_opt_pref64_len returns the Nd_opt_pref64_len field from the record's packed storage.
func (s *Nd_opt_pref64) Nd_opt_pref64_len() U_int8_t {
	return *(*U_int8_t)(unsafe.Pointer(&s.storage[1]))
}

// SetNd_opt_pref64_len updates the Nd_opt_pref64_len field in the record's packed storage.
func (s *Nd_opt_pref64) SetNd_opt_pref64_len(v U_int8_t) {
	*(*U_int8_t)(unsafe.Pointer(&s.storage[1])) = v
}

// Nd_opt_pref64_scaled_lifetime_plc returns the Nd_opt_pref64_scaled_lifetime_plc field from the record's packed storage.
func (s *Nd_opt_pref64) Nd_opt_pref64_scaled_lifetime_plc() U_int16_t {
	return U_int16_t(binary.NativeEndian.Uint16(s.storage[2:4]))
}

// SetNd_opt_pref64_scaled_lifetime_plc updates the Nd_opt_pref64_scaled_lifetime_plc field in the record's packed storage.
func (s *Nd_opt_pref64) SetNd_opt_pref64_scaled_lifetime_plc(v U_int16_t) {
	binary.NativeEndian.PutUint16(s.storage[2:4], uint16(v))
}

// Nd_opt_pref64_prefix returns the Nd_opt_pref64_prefix field from the record's packed storage.
func (s *Nd_opt_pref64) Nd_opt_pref64_prefix() [3]U_int32_t {
	return *(*[3]U_int32_t)(unsafe.Pointer(&s.storage[4]))
}

// SetNd_opt_pref64_prefix updates the Nd_opt_pref64_prefix field in the record's packed storage.
func (s *Nd_opt_pref64) SetNd_opt_pref64_prefix(v [3]U_int32_t) {
	*(*[3]U_int32_t)(unsafe.Pointer(&s.storage[4])) = v
}

// Nd_opt_prefix_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nd_opt_prefix_info
type Nd_opt_prefix_info struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [32]byte
}

// Nd_opt_pi_type returns the Nd_opt_pi_type field from the record's packed storage.
func (s *Nd_opt_prefix_info) Nd_opt_pi_type() U_int8_t {
	return *(*U_int8_t)(unsafe.Pointer(&s.storage[0]))
}

// SetNd_opt_pi_type updates the Nd_opt_pi_type field in the record's packed storage.
func (s *Nd_opt_prefix_info) SetNd_opt_pi_type(v U_int8_t) {
	*(*U_int8_t)(unsafe.Pointer(&s.storage[0])) = v
}

// Nd_opt_pi_len returns the Nd_opt_pi_len field from the record's packed storage.
func (s *Nd_opt_prefix_info) Nd_opt_pi_len() U_int8_t {
	return *(*U_int8_t)(unsafe.Pointer(&s.storage[1]))
}

// SetNd_opt_pi_len updates the Nd_opt_pi_len field in the record's packed storage.
func (s *Nd_opt_prefix_info) SetNd_opt_pi_len(v U_int8_t) {
	*(*U_int8_t)(unsafe.Pointer(&s.storage[1])) = v
}

// Nd_opt_pi_prefix_len returns the Nd_opt_pi_prefix_len field from the record's packed storage.
func (s *Nd_opt_prefix_info) Nd_opt_pi_prefix_len() U_int8_t {
	return *(*U_int8_t)(unsafe.Pointer(&s.storage[2]))
}

// SetNd_opt_pi_prefix_len updates the Nd_opt_pi_prefix_len field in the record's packed storage.
func (s *Nd_opt_prefix_info) SetNd_opt_pi_prefix_len(v U_int8_t) {
	*(*U_int8_t)(unsafe.Pointer(&s.storage[2])) = v
}

// Nd_opt_pi_flags_reserved returns the Nd_opt_pi_flags_reserved field from the record's packed storage.
func (s *Nd_opt_prefix_info) Nd_opt_pi_flags_reserved() U_int8_t {
	return *(*U_int8_t)(unsafe.Pointer(&s.storage[3]))
}

// SetNd_opt_pi_flags_reserved updates the Nd_opt_pi_flags_reserved field in the record's packed storage.
func (s *Nd_opt_prefix_info) SetNd_opt_pi_flags_reserved(v U_int8_t) {
	*(*U_int8_t)(unsafe.Pointer(&s.storage[3])) = v
}

// Nd_opt_pi_valid_time returns the Nd_opt_pi_valid_time field from the record's packed storage.
func (s *Nd_opt_prefix_info) Nd_opt_pi_valid_time() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetNd_opt_pi_valid_time updates the Nd_opt_pi_valid_time field in the record's packed storage.
func (s *Nd_opt_prefix_info) SetNd_opt_pi_valid_time(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// Nd_opt_pi_preferred_time returns the Nd_opt_pi_preferred_time field from the record's packed storage.
func (s *Nd_opt_prefix_info) Nd_opt_pi_preferred_time() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetNd_opt_pi_preferred_time updates the Nd_opt_pi_preferred_time field in the record's packed storage.
func (s *Nd_opt_prefix_info) SetNd_opt_pi_preferred_time(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[8:12], uint32(v))
}

// Nd_opt_pi_reserved2 returns the Nd_opt_pi_reserved2 field from the record's packed storage.
func (s *Nd_opt_prefix_info) Nd_opt_pi_reserved2() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[12:16]))
}

// SetNd_opt_pi_reserved2 updates the Nd_opt_pi_reserved2 field in the record's packed storage.
func (s *Nd_opt_prefix_info) SetNd_opt_pi_reserved2(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[12:16], uint32(v))
}

// Nd_opt_pi_prefix returns the Nd_opt_pi_prefix field from the record's packed storage.
func (s *Nd_opt_prefix_info) Nd_opt_pi_prefix() [4]uint32 {
	return *(*[4]uint32)(unsafe.Pointer(&s.storage[16]))
}

// SetNd_opt_pi_prefix updates the Nd_opt_pi_prefix field in the record's packed storage.
func (s *Nd_opt_prefix_info) SetNd_opt_pi_prefix(v [4]uint32) {
	*(*[4]uint32)(unsafe.Pointer(&s.storage[16])) = v
}

// Nd_opt_pvd
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nd_opt_pvd
type Nd_opt_pvd struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [7]byte
}

// Nd_opt_pvd_type returns the Nd_opt_pvd_type field from the record's packed storage.
func (s *Nd_opt_pvd) Nd_opt_pvd_type() U_int8_t {
	return *(*U_int8_t)(unsafe.Pointer(&s.storage[0]))
}

// SetNd_opt_pvd_type updates the Nd_opt_pvd_type field in the record's packed storage.
func (s *Nd_opt_pvd) SetNd_opt_pvd_type(v U_int8_t) {
	*(*U_int8_t)(unsafe.Pointer(&s.storage[0])) = v
}

// Nd_opt_pvd_len returns the Nd_opt_pvd_len field from the record's packed storage.
func (s *Nd_opt_pvd) Nd_opt_pvd_len() U_int8_t {
	return *(*U_int8_t)(unsafe.Pointer(&s.storage[1]))
}

// SetNd_opt_pvd_len updates the Nd_opt_pvd_len field in the record's packed storage.
func (s *Nd_opt_pvd) SetNd_opt_pvd_len(v U_int8_t) {
	*(*U_int8_t)(unsafe.Pointer(&s.storage[1])) = v
}

// Nd_opt_flags_delay returns the Nd_opt_flags_delay field from the record's packed storage.
func (s *Nd_opt_pvd) Nd_opt_flags_delay() [2]U_int8_t {
	return *(*[2]U_int8_t)(unsafe.Pointer(&s.storage[2]))
}

// SetNd_opt_flags_delay updates the Nd_opt_flags_delay field in the record's packed storage.
func (s *Nd_opt_pvd) SetNd_opt_flags_delay(v [2]U_int8_t) {
	*(*[2]U_int8_t)(unsafe.Pointer(&s.storage[2])) = v
}

// Nd_opt_pvd_seq returns the Nd_opt_pvd_seq field from the record's packed storage.
func (s *Nd_opt_pvd) Nd_opt_pvd_seq() U_int16_t {
	return U_int16_t(binary.NativeEndian.Uint16(s.storage[4:6]))
}

// SetNd_opt_pvd_seq updates the Nd_opt_pvd_seq field in the record's packed storage.
func (s *Nd_opt_pvd) SetNd_opt_pvd_seq(v U_int16_t) {
	binary.NativeEndian.PutUint16(s.storage[4:6], uint16(v))
}

// Nd_opt_pvd_id returns the Nd_opt_pvd_id field from the record's packed storage.
func (s *Nd_opt_pvd) Nd_opt_pvd_id() [1]U_int8_t {
	return *(*[1]U_int8_t)(unsafe.Pointer(&s.storage[6]))
}

// SetNd_opt_pvd_id updates the Nd_opt_pvd_id field in the record's packed storage.
func (s *Nd_opt_pvd) SetNd_opt_pvd_id(v [1]U_int8_t) {
	*(*[1]U_int8_t)(unsafe.Pointer(&s.storage[6])) = v
}

// Nd_opt_rd_hdr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nd_opt_rd_hdr
type Nd_opt_rd_hdr struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [8]byte
}

// Nd_opt_rh_type returns the Nd_opt_rh_type field from the record's packed storage.
func (s *Nd_opt_rd_hdr) Nd_opt_rh_type() U_int8_t {
	return *(*U_int8_t)(unsafe.Pointer(&s.storage[0]))
}

// SetNd_opt_rh_type updates the Nd_opt_rh_type field in the record's packed storage.
func (s *Nd_opt_rd_hdr) SetNd_opt_rh_type(v U_int8_t) {
	*(*U_int8_t)(unsafe.Pointer(&s.storage[0])) = v
}

// Nd_opt_rh_len returns the Nd_opt_rh_len field from the record's packed storage.
func (s *Nd_opt_rd_hdr) Nd_opt_rh_len() U_int8_t {
	return *(*U_int8_t)(unsafe.Pointer(&s.storage[1]))
}

// SetNd_opt_rh_len updates the Nd_opt_rh_len field in the record's packed storage.
func (s *Nd_opt_rd_hdr) SetNd_opt_rh_len(v U_int8_t) {
	*(*U_int8_t)(unsafe.Pointer(&s.storage[1])) = v
}

// Nd_opt_rh_reserved1 returns the Nd_opt_rh_reserved1 field from the record's packed storage.
func (s *Nd_opt_rd_hdr) Nd_opt_rh_reserved1() U_int16_t {
	return U_int16_t(binary.NativeEndian.Uint16(s.storage[2:4]))
}

// SetNd_opt_rh_reserved1 updates the Nd_opt_rh_reserved1 field in the record's packed storage.
func (s *Nd_opt_rd_hdr) SetNd_opt_rh_reserved1(v U_int16_t) {
	binary.NativeEndian.PutUint16(s.storage[2:4], uint16(v))
}

// Nd_opt_rh_reserved2 returns the Nd_opt_rh_reserved2 field from the record's packed storage.
func (s *Nd_opt_rd_hdr) Nd_opt_rh_reserved2() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetNd_opt_rh_reserved2 updates the Nd_opt_rh_reserved2 field in the record's packed storage.
func (s *Nd_opt_rd_hdr) SetNd_opt_rh_reserved2(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// Nd_opt_rdnss
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nd_opt_rdnss
type Nd_opt_rdnss struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [24]byte
}

// Nd_opt_rdnss_type returns the Nd_opt_rdnss_type field from the record's packed storage.
func (s *Nd_opt_rdnss) Nd_opt_rdnss_type() U_int8_t {
	return *(*U_int8_t)(unsafe.Pointer(&s.storage[0]))
}

// SetNd_opt_rdnss_type updates the Nd_opt_rdnss_type field in the record's packed storage.
func (s *Nd_opt_rdnss) SetNd_opt_rdnss_type(v U_int8_t) {
	*(*U_int8_t)(unsafe.Pointer(&s.storage[0])) = v
}

// Nd_opt_rdnss_len returns the Nd_opt_rdnss_len field from the record's packed storage.
func (s *Nd_opt_rdnss) Nd_opt_rdnss_len() U_int8_t {
	return *(*U_int8_t)(unsafe.Pointer(&s.storage[1]))
}

// SetNd_opt_rdnss_len updates the Nd_opt_rdnss_len field in the record's packed storage.
func (s *Nd_opt_rdnss) SetNd_opt_rdnss_len(v U_int8_t) {
	*(*U_int8_t)(unsafe.Pointer(&s.storage[1])) = v
}

// Nd_opt_rdnss_reserved returns the Nd_opt_rdnss_reserved field from the record's packed storage.
func (s *Nd_opt_rdnss) Nd_opt_rdnss_reserved() U_int16_t {
	return U_int16_t(binary.NativeEndian.Uint16(s.storage[2:4]))
}

// SetNd_opt_rdnss_reserved updates the Nd_opt_rdnss_reserved field in the record's packed storage.
func (s *Nd_opt_rdnss) SetNd_opt_rdnss_reserved(v U_int16_t) {
	binary.NativeEndian.PutUint16(s.storage[2:4], uint16(v))
}

// Nd_opt_rdnss_lifetime returns the Nd_opt_rdnss_lifetime field from the record's packed storage.
func (s *Nd_opt_rdnss) Nd_opt_rdnss_lifetime() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetNd_opt_rdnss_lifetime updates the Nd_opt_rdnss_lifetime field in the record's packed storage.
func (s *Nd_opt_rdnss) SetNd_opt_rdnss_lifetime(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// Nd_opt_rdnss_addr returns the Nd_opt_rdnss_addr field from the record's packed storage.
func (s *Nd_opt_rdnss) Nd_opt_rdnss_addr() [16]byte {
	return *(*[16]byte)(unsafe.Pointer(&s.storage[8]))
}

// SetNd_opt_rdnss_addr updates the Nd_opt_rdnss_addr field in the record's packed storage.
func (s *Nd_opt_rdnss) SetNd_opt_rdnss_addr(v [16]byte) {
	*(*[16]byte)(unsafe.Pointer(&s.storage[8])) = v
}

// Nd_opt_route_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nd_opt_route_info
type Nd_opt_route_info struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [8]byte
}

// Nd_opt_rti_type returns the Nd_opt_rti_type field from the record's packed storage.
func (s *Nd_opt_route_info) Nd_opt_rti_type() U_int8_t {
	return *(*U_int8_t)(unsafe.Pointer(&s.storage[0]))
}

// SetNd_opt_rti_type updates the Nd_opt_rti_type field in the record's packed storage.
func (s *Nd_opt_route_info) SetNd_opt_rti_type(v U_int8_t) {
	*(*U_int8_t)(unsafe.Pointer(&s.storage[0])) = v
}

// Nd_opt_rti_len returns the Nd_opt_rti_len field from the record's packed storage.
func (s *Nd_opt_route_info) Nd_opt_rti_len() U_int8_t {
	return *(*U_int8_t)(unsafe.Pointer(&s.storage[1]))
}

// SetNd_opt_rti_len updates the Nd_opt_rti_len field in the record's packed storage.
func (s *Nd_opt_route_info) SetNd_opt_rti_len(v U_int8_t) {
	*(*U_int8_t)(unsafe.Pointer(&s.storage[1])) = v
}

// Nd_opt_rti_prefixlen returns the Nd_opt_rti_prefixlen field from the record's packed storage.
func (s *Nd_opt_route_info) Nd_opt_rti_prefixlen() U_int8_t {
	return *(*U_int8_t)(unsafe.Pointer(&s.storage[2]))
}

// SetNd_opt_rti_prefixlen updates the Nd_opt_rti_prefixlen field in the record's packed storage.
func (s *Nd_opt_route_info) SetNd_opt_rti_prefixlen(v U_int8_t) {
	*(*U_int8_t)(unsafe.Pointer(&s.storage[2])) = v
}

// Nd_opt_rti_flags returns the Nd_opt_rti_flags field from the record's packed storage.
func (s *Nd_opt_route_info) Nd_opt_rti_flags() U_int8_t {
	return *(*U_int8_t)(unsafe.Pointer(&s.storage[3]))
}

// SetNd_opt_rti_flags updates the Nd_opt_rti_flags field in the record's packed storage.
func (s *Nd_opt_route_info) SetNd_opt_rti_flags(v U_int8_t) {
	*(*U_int8_t)(unsafe.Pointer(&s.storage[3])) = v
}

// Nd_opt_rti_lifetime returns the Nd_opt_rti_lifetime field from the record's packed storage.
func (s *Nd_opt_route_info) Nd_opt_rti_lifetime() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetNd_opt_rti_lifetime updates the Nd_opt_rti_lifetime field in the record's packed storage.
func (s *Nd_opt_route_info) SetNd_opt_rti_lifetime(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// Nd_redirect
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nd_redirect
type Nd_redirect struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [40]byte
}

// Nd_rd_hdr returns the Nd_rd_hdr field from the record's packed storage.
func (s *Nd_redirect) Nd_rd_hdr() Icmp6_hdr {
	return *(*Icmp6_hdr)(unsafe.Pointer(&s.storage[0]))
}

// SetNd_rd_hdr updates the Nd_rd_hdr field in the record's packed storage.
func (s *Nd_redirect) SetNd_rd_hdr(v Icmp6_hdr) {
	*(*Icmp6_hdr)(unsafe.Pointer(&s.storage[0])) = v
}

// Nd_rd_target returns the Nd_rd_target field from the record's packed storage.
func (s *Nd_redirect) Nd_rd_target() [4]uint32 {
	return *(*[4]uint32)(unsafe.Pointer(&s.storage[8]))
}

// SetNd_rd_target updates the Nd_rd_target field in the record's packed storage.
func (s *Nd_redirect) SetNd_rd_target(v [4]uint32) {
	*(*[4]uint32)(unsafe.Pointer(&s.storage[8])) = v
}

// Nd_rd_dst returns the Nd_rd_dst field from the record's packed storage.
func (s *Nd_redirect) Nd_rd_dst() [4]uint32 {
	return *(*[4]uint32)(unsafe.Pointer(&s.storage[24]))
}

// SetNd_rd_dst updates the Nd_rd_dst field in the record's packed storage.
func (s *Nd_redirect) SetNd_rd_dst(v [4]uint32) {
	*(*[4]uint32)(unsafe.Pointer(&s.storage[24])) = v
}

// Nd_router_advert
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nd_router_advert
type Nd_router_advert struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [16]byte
}

// Nd_ra_hdr returns the Nd_ra_hdr field from the record's packed storage.
func (s *Nd_router_advert) Nd_ra_hdr() Icmp6_hdr {
	return *(*Icmp6_hdr)(unsafe.Pointer(&s.storage[0]))
}

// SetNd_ra_hdr updates the Nd_ra_hdr field in the record's packed storage.
func (s *Nd_router_advert) SetNd_ra_hdr(v Icmp6_hdr) {
	*(*Icmp6_hdr)(unsafe.Pointer(&s.storage[0])) = v
}

// Nd_ra_reachable returns the Nd_ra_reachable field from the record's packed storage.
func (s *Nd_router_advert) Nd_ra_reachable() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetNd_ra_reachable updates the Nd_ra_reachable field in the record's packed storage.
func (s *Nd_router_advert) SetNd_ra_reachable(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[8:12], uint32(v))
}

// Nd_ra_retransmit returns the Nd_ra_retransmit field from the record's packed storage.
func (s *Nd_router_advert) Nd_ra_retransmit() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[12:16]))
}

// SetNd_ra_retransmit updates the Nd_ra_retransmit field in the record's packed storage.
func (s *Nd_router_advert) SetNd_ra_retransmit(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[12:16], uint32(v))
}

// Nd_router_solicit
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nd_router_solicit
type Nd_router_solicit struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [8]byte
}

// Nd_rs_hdr returns the Nd_rs_hdr field from the record's packed storage.
func (s *Nd_router_solicit) Nd_rs_hdr() Icmp6_hdr {
	return *(*Icmp6_hdr)(unsafe.Pointer(&s.storage[0]))
}

// SetNd_rs_hdr updates the Nd_rs_hdr field in the record's packed storage.
func (s *Nd_router_solicit) SetNd_rs_hdr(v Icmp6_hdr) {
	*(*Icmp6_hdr)(unsafe.Pointer(&s.storage[0])) = v
}

// Ndrv_demux_desc
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ndrv_demux_desc
type Ndrv_demux_desc struct {
	Type   U_int16_t
	Length U_int16_t
	Data   [14]uint16
}

// Ndrv_protocol_desc
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ndrv_protocol_desc
type Ndrv_protocol_desc struct {
	Version         U_int32_t
	Protocol_family U_int32_t
	Demux_count     U_int32_t
	Demux_list      *Ndrv_demux_desc
}

// Net_event_data
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/net_event_data
type Net_event_data struct {
	If_family U_int32_t
	If_unit   U_int32_t
	If_name   [16]int8
}

// Netfs_status
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/netfs_status
type Netfs_status struct {
	Ns_status      U_int32_t
	Ns_mountopts   [512]int8
	Ns_waittime    uint32
	Ns_threadcount uint32
}

// Newah
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/newah
type Newah struct {
	Ah_nxt     U_int8_t
	Ah_len     U_int8_t
	Ah_reserve U_int16_t
	Ah_spi     U_int32_t
	Ah_seq     U_int32_t
}

// Newesp
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/newesp
type Newesp struct {
	Esp_spi U_int32_t
	Esp_seq U_int32_t
}

// Nextvend
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nextvend
type Nextvend struct {
	Nv_magic   [4]U_char
	Nv_version U_char
	padding2   [1]byte
	Nv_U       [58]byte
}

// Nfs_args
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nfs_args
type Nfs_args struct {
	Version      int32
	Addr         User32_addr_t
	Addrlen      uint8
	Sotype       int32
	Proto        int32
	Fh           User32_addr_t
	Fhsize       int32
	Flags        int32
	Wsize        int32
	Rsize        int32
	Readdirsize  int32
	Timeo        int32
	Retrans      int32
	Maxgrouplist int32
	Readahead    int32
	Leaseterm    int32
	Deadthresh   int32
	Hostname     User32_addr_t
	Acregmin     int32
	Acregmax     int32
	Acdirmin     int32
	Acdirmax     int32
	Auth         uint32
	Deadtimeout  uint32
}

// Nfs_etype
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nfs_etype
type Nfs_etype struct {
	Count    uint32
	Selected uint32
	Etypes   [3]Nfs_supported_kerberos_etypes
}

// Nfs_exphandle
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nfs_exphandle
type Nfs_exphandle struct {
	Nxh_version  uint32
	Nxh_fsid     uint32
	Nxh_expid    uint32
	Nxh_flags    uint16
	Nxh_reserved uint8
	Nxh_fidlen   uint32
}

// Nfs_export_args
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nfs_export_args
type Nfs_export_args struct {
	Nxa_fsid     uint32
	Nxa_expid    uint32
	Nxa_fspath   User32_addr_t
	Nxa_exppath  User32_addr_t
	Nxa_flags    uint32
	Nxa_netcount uint32
	Nxa_nets     User32_addr_t
}

// Nfs_export_net_args
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nfs_export_net_args
type Nfs_export_net_args struct {
	Nxna_flags uint32
	Nxna_cred  Xucred
	Nxna_addr  [16]uint64
	Nxna_mask  [16]uint64
	Nxna_sec   Nfs_sec
}

// Nfs_export_stat_desc
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nfs_export_stat_desc
type Nfs_export_stat_desc struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [12]byte
}

// Rec_vers returns the Rec_vers field from the record's packed storage.
func (s *Nfs_export_stat_desc) Rec_vers() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetRec_vers updates the Rec_vers field in the record's packed storage.
func (s *Nfs_export_stat_desc) SetRec_vers(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// Rec_count returns the Rec_count field from the record's packed storage.
func (s *Nfs_export_stat_desc) Rec_count() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[4:12]))
}

// SetRec_count updates the Rec_count field in the record's packed storage.
func (s *Nfs_export_stat_desc) SetRec_count(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[4:12], uint64(v))
}

// Nfs_export_stat_rec
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nfs_export_stat_rec
type Nfs_export_stat_rec struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [1049]byte
}

// Path returns the Path field from the record's packed storage.
func (s *Nfs_export_stat_rec) Path() [1025]int8 {
	return *(*[1025]int8)(unsafe.Pointer(&s.storage[0]))
}

// SetPath updates the Path field in the record's packed storage.
func (s *Nfs_export_stat_rec) SetPath(v [1025]int8) {
	*(*[1025]int8)(unsafe.Pointer(&s.storage[0])) = v
}

// Ops returns the Ops field from the record's packed storage.
func (s *Nfs_export_stat_rec) Ops() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[1025:1033]))
}

// SetOps updates the Ops field in the record's packed storage.
func (s *Nfs_export_stat_rec) SetOps(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[1025:1033], uint64(v))
}

// Bytes_read returns the Bytes_read field from the record's packed storage.
func (s *Nfs_export_stat_rec) Bytes_read() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[1033:1041]))
}

// SetBytes_read updates the Bytes_read field in the record's packed storage.
func (s *Nfs_export_stat_rec) SetBytes_read(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[1033:1041], uint64(v))
}

// Bytes_written returns the Bytes_written field from the record's packed storage.
func (s *Nfs_export_stat_rec) Bytes_written() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[1041:1049]))
}

// SetBytes_written updates the Bytes_written field in the record's packed storage.
func (s *Nfs_export_stat_rec) SetBytes_written(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[1041:1049], uint64(v))
}

// Nfs_filehandle
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nfs_filehandle
type Nfs_filehandle struct {
	Nfh_len uint32
	Nfh_xh  Nfs_exphandle
	Nfh_fid [108]uint8
	Nfh_fhp *byte
}

// Nfs_sec
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nfs_sec
type Nfs_sec struct {
	Count   int32
	Flavors [5]uint32
}

// Nfs_testmapid
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nfs_testmapid
type Nfs_testmapid struct {
	Ntm_lookup  uint32
	Ntm_grpflag uint32
	Ntm_id      uint32
	Pad         uint32
	Ntm_guid    [4]uint32
	Ntm_name    [1024]int8
}

// Nfs_user_stat_desc
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nfs_user_stat_desc
type Nfs_user_stat_desc struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [8]byte
}

// Rec_vers returns the Rec_vers field from the record's packed storage.
func (s *Nfs_user_stat_desc) Rec_vers() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetRec_vers updates the Rec_vers field in the record's packed storage.
func (s *Nfs_user_stat_desc) SetRec_vers(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// Rec_count returns the Rec_count field from the record's packed storage.
func (s *Nfs_user_stat_desc) Rec_count() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetRec_count updates the Rec_count field in the record's packed storage.
func (s *Nfs_user_stat_desc) SetRec_count(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// Nfs_user_stat_path_rec
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nfs_user_stat_path_rec
type Nfs_user_stat_path_rec struct {
	Rec_type U_char
	Path     [1025]int8
}

// Nfs_user_stat_user_rec
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nfs_user_stat_user_rec
type Nfs_user_stat_user_rec struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [173]byte
}

// Rec_type returns the Rec_type field from the record's packed storage.
func (s *Nfs_user_stat_user_rec) Rec_type() U_char {
	return *(*U_char)(unsafe.Pointer(&s.storage[0]))
}

// SetRec_type updates the Rec_type field in the record's packed storage.
func (s *Nfs_user_stat_user_rec) SetRec_type(v U_char) {
	*(*U_char)(unsafe.Pointer(&s.storage[0])) = v
}

// Uid returns the Uid field from the record's packed storage.
func (s *Nfs_user_stat_user_rec) Uid() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[1:5]))
}

// SetUid updates the Uid field in the record's packed storage.
func (s *Nfs_user_stat_user_rec) SetUid(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[1:5], uint32(v))
}

// Sock returns the Sock field from the record's packed storage.
func (s *Nfs_user_stat_user_rec) Sock() [16]uint64 {
	return *(*[16]uint64)(unsafe.Pointer(&s.storage[5]))
}

// SetSock updates the Sock field in the record's packed storage.
func (s *Nfs_user_stat_user_rec) SetSock(v [16]uint64) {
	*(*[16]uint64)(unsafe.Pointer(&s.storage[5])) = v
}

// Ops returns the Ops field from the record's packed storage.
func (s *Nfs_user_stat_user_rec) Ops() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[133:141]))
}

// SetOps updates the Ops field in the record's packed storage.
func (s *Nfs_user_stat_user_rec) SetOps(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[133:141], uint64(v))
}

// Bytes_read returns the Bytes_read field from the record's packed storage.
func (s *Nfs_user_stat_user_rec) Bytes_read() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[141:149]))
}

// SetBytes_read updates the Bytes_read field in the record's packed storage.
func (s *Nfs_user_stat_user_rec) SetBytes_read(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[141:149], uint64(v))
}

// Bytes_written returns the Bytes_written field from the record's packed storage.
func (s *Nfs_user_stat_user_rec) Bytes_written() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[149:157]))
}

// SetBytes_written updates the Bytes_written field in the record's packed storage.
func (s *Nfs_user_stat_user_rec) SetBytes_written(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[149:157], uint64(v))
}

// Tm_start returns the Tm_start field from the record's packed storage.
func (s *Nfs_user_stat_user_rec) Tm_start() int64 {
	return int64(binary.NativeEndian.Uint64(s.storage[157:165]))
}

// SetTm_start updates the Tm_start field in the record's packed storage.
func (s *Nfs_user_stat_user_rec) SetTm_start(v int64) {
	binary.NativeEndian.PutUint64(s.storage[157:165], uint64(v))
}

// Tm_last returns the Tm_last field from the record's packed storage.
func (s *Nfs_user_stat_user_rec) Tm_last() int64 {
	return int64(binary.NativeEndian.Uint64(s.storage[165:173]))
}

// SetTm_last updates the Tm_last field in the record's packed storage.
func (s *Nfs_user_stat_user_rec) SetTm_last(v int64) {
	binary.NativeEndian.PutUint64(s.storage[165:173], uint64(v))
}

// Nfsclntstats
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nfsclntstats
type Nfsclntstats struct {
	Attrcache_hits     uint64
	Attrcache_misses   uint64
	Lookupcache_hits   uint64
	Lookupcache_misses uint64
	Direofcache_hits   uint64
	Direofcache_misses uint64
	Accesscache_hits   uint64
	Accesscache_misses uint64
	Biocache_reads     uint64
	Read_bios          uint64
	Read_physios       uint64
	Biocache_writes    uint64
	Write_bios         uint64
	Write_physios      uint64
	Biocache_readlinks uint64
	Readlink_bios      uint64
	Biocache_readdirs  uint64
	Readdir_bios       uint64
	Rpccntv3           [23]uint64
	Nlm_lock           uint64
	Nlm_test           uint64
	Nlm_unlock         uint64
	Opcntv4            [59]uint64
	Cbopcntv4          [15]uint64
	Rpcretries         uint64
	Rpcrequests        uint64
	Rpctimeouts        uint64
	Rpcunexpected      uint64
	Rpcinvalid         uint64
	Pageins            uint64
	Pageouts           uint64
	Errs_common        [30]uint64
	Errs_v4            [78]uint64
	Errs_unknown       uint64
	Nfs_errs           unsafe.Pointer
	Nlmcnt             unsafe.Pointer
}

// Nfsd_args
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nfsd_args
type Nfsd_args struct {
	Sock    int32
	Name    User32_addr_t
	Namelen int32
}

// Nfsrvstats
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nfsrvstats
type Nfsrvstats struct {
	Srvrpccntv3              [23]uint64
	Srvrpc_errs              uint64
	Srv_errs                 uint64
	Srvcache_inproghits      uint64
	Srvcache_idemdonehits    uint64
	Srvcache_nonidemdonehits uint64
	Srvcache_misses          uint64
	Srvvop_writes            uint64
	Errs_common              [30]uint64
	Errs_unknown             uint64
	Nfs_errs                 unsafe.Pointer
}

// Ni_reply_fqdn
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ni_reply_fqdn
type Ni_reply_fqdn struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [8]byte
}

// Ni_fqdn_ttl returns the Ni_fqdn_ttl field from the record's packed storage.
func (s *Ni_reply_fqdn) Ni_fqdn_ttl() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetNi_fqdn_ttl updates the Ni_fqdn_ttl field in the record's packed storage.
func (s *Ni_reply_fqdn) SetNi_fqdn_ttl(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// Ni_fqdn_namelen returns the Ni_fqdn_namelen field from the record's packed storage.
func (s *Ni_reply_fqdn) Ni_fqdn_namelen() U_int8_t {
	return *(*U_int8_t)(unsafe.Pointer(&s.storage[4]))
}

// SetNi_fqdn_namelen updates the Ni_fqdn_namelen field in the record's packed storage.
func (s *Ni_reply_fqdn) SetNi_fqdn_namelen(v U_int8_t) {
	*(*U_int8_t)(unsafe.Pointer(&s.storage[4])) = v
}

// Ni_fqdn_name returns the Ni_fqdn_name field from the record's packed storage.
func (s *Ni_reply_fqdn) Ni_fqdn_name() [3]U_int8_t {
	return *(*[3]U_int8_t)(unsafe.Pointer(&s.storage[5]))
}

// SetNi_fqdn_name updates the Ni_fqdn_name field in the record's packed storage.
func (s *Ni_reply_fqdn) SetNi_fqdn_name(v [3]U_int8_t) {
	*(*[3]U_int8_t)(unsafe.Pointer(&s.storage[5])) = v
}

// Nlist - Describes an entry in the symbol table for 32-bit architectures. Declared in `/usr/include/mach-o/nlist.h`. See also [nlist_64](<https://developer.apple.com/documentation/kernel/nlist_64>).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nlist
type Nlist struct {
	N_name  *byte
	N_type  uint8 // A byte value consisting of data accessed using four bit masks:
	N_other int8
	N_desc  int16          // A 16-bit value providing additional information about the nature of this symbol for non-stab symbols. The reference flags can be accessed using the `REFERENCE_TYPE` mask (0xF) and are defined as follows:
	N_value uint32         // An integer that contains the value of the symbol. The format of this value is different for each type of symbol table entry (as specified by the `n_type` field). For the `N_SECT` symbol type, `n_value` is the address of the symbol. See the description of the `n_type` field for information on other possible values.
	N_un    unsafe.Pointer // A union that holds an index into the string table, `n_strx`. To specify an empty string (`""`), set this value to 0. The `n_name` field is not used in Mach-O files.
	N_sect  uint8          // An integer specifying the number of the section that this symbol can be found in, or `NO_SECT` if the symbol is not to be found in any section of this image. The sections are contiguously numbered across segments, starting from 1, according to the order they appear in the `LC_SEGMENT` load commands.

}

// Nlist_64 - Describes an entry in the symbol table for 64-bit architectures. Declared in `/usr/include/mach-o/nlist.h`.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nlist_64
type Nlist_64 struct {
	N_un    [1]uint32 // A union that holds an index into the string table, `n_strx`. To specify an empty string (`""`), set this value to 0.
	N_type  uint8     // A byte value consisting of data accessed using four bit masks:
	N_sect  uint8     // An integer specifying the number of the section that this symbol can be found in, or `NO_SECT` if the symbol is not to be found in any section of this image. The sections are contiguously numbered across segments, starting from 1, according to the order they appear in the `LC_SEGMENT` load commands.
	N_desc  uint16    // A 16-bit value providing additional information about the nature of this symbol. The reference flags can be accessed using the `REFERENCE_TYPE` mask (0xF) and are defined as follows:
	N_value uint64
}

// Note_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/note_command
type Note_command struct {
	Cmd        uint32
	Cmdsize    uint32
	Data_owner [16]int8
	Offset     uint64
	Size       uint64
}

// Ntptimeval
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ntptimeval
type Ntptimeval struct {
	Time       Timespec
	Maxerror   int
	Esterror   int
	Tai        int
	Time_state int32
}

// Ombstat
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ombstat
type Ombstat struct {
	M_mbufs     U_int32_t
	M_clusters  U_int32_t
	M_spare     U_int32_t
	M_clfree    U_int32_t
	M_drops     U_int32_t
	M_wait      U_int32_t
	M_drain     U_int32_t
	M_mtypes    [256]U_short
	M_mcfail    U_int32_t
	M_mpfail    U_int32_t
	M_msize     U_int32_t
	M_mclbytes  U_int32_t
	M_minclsize U_int32_t
	M_mlen      U_int32_t
	M_mhlen     U_int32_t
}

// Ostat
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ostat
type Ostat struct {
	St_dev       uint16
	St_ino       uint64
	St_mode      uint16
	St_nlink     uint16
	St_uid       uint16
	St_gid       uint16
	St_rdev      uint16
	St_size      int32
	St_atimespec Timespec
	St_mtimespec Timespec
	St_ctimespec Timespec
	St_blksize   int32
	St_blocks    int32
	St_flags     uint32
	St_gen       uint32
}

// Persona_modify_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/persona_modify_info
type Persona_modify_info struct {
	Persona_id uint32
	Unique_pid uint64
}

// Persona_token
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/persona_token
type Persona_token struct {
	Originator Proc_persona_info
	Proximate  Proc_persona_info
}

// Portlabel_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/portlabel_info
type Portlabel_info struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [5]byte
}

// Portlabel_id returns the Portlabel_id field from the record's packed storage.
func (s *Portlabel_info) Portlabel_id() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[0:2]))
}

// SetPortlabel_id updates the Portlabel_id field in the record's packed storage.
func (s *Portlabel_info) SetPortlabel_id(v int16) {
	binary.NativeEndian.PutUint16(s.storage[0:2], uint16(v))
}

// Portlabel_flags returns the Portlabel_flags field from the record's packed storage.
func (s *Portlabel_info) Portlabel_flags() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[2:4]))
}

// SetPortlabel_flags updates the Portlabel_flags field in the record's packed storage.
func (s *Portlabel_info) SetPortlabel_flags(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[2:4], uint16(v))
}

// Portlabel_domain returns the Portlabel_domain field from the record's packed storage.
func (s *Portlabel_info) Portlabel_domain() uint8 {
	return uint8(s.storage[4])
}

// SetPortlabel_domain updates the Portlabel_domain field in the record's packed storage.
func (s *Portlabel_info) SetPortlabel_domain(v uint8) {
	s.storage[4] = uint8(v)
}

// Prebind_cksum_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/prebind_cksum_command
type Prebind_cksum_command struct {
	Cmd     uint32
	Cmdsize uint32
	Cksum   uint32
}

// Prebound_dylib_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/prebound_dylib_command
type Prebound_dylib_command struct {
	Cmd            uint32 // Common to all load command structures. For this structure, set to `LC_PREBOUND_DYLIB`.
	Cmdsize        uint32
	Name           [1]uint32
	Nmodules       uint32
	Linked_modules [1]uint32
}

// Priority_queue_deadline_max
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/priority_queue_deadline_max
type Priority_queue_deadline_max struct {
	Pq_root unsafe.Pointer
}

// Priority_queue_deadline_min
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/priority_queue_deadline_min
type Priority_queue_deadline_min struct {
	Pq_root unsafe.Pointer
}

// Priority_queue_max
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/priority_queue_max
type Priority_queue_max struct {
	Pq_root   unsafe.Pointer
	Pq_cmp_fn Priority_queue_compare_fn_t
}

// Priority_queue_min
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/priority_queue_min
type Priority_queue_min struct {
	Pq_root   unsafe.Pointer
	Pq_cmp_fn Priority_queue_compare_fn_t
}

// Priority_queue_sched_max
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/priority_queue_sched_max
type Priority_queue_sched_max struct {
	Pq_root unsafe.Pointer
}

// Priority_queue_sched_min
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/priority_queue_sched_min
type Priority_queue_sched_min struct {
	Pq_root unsafe.Pointer
}

// Priority_queue_sched_stable_max
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/priority_queue_sched_stable_max
type Priority_queue_sched_stable_max struct {
	Pq_root unsafe.Pointer
}

// Priority_queue_sched_stable_min
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/priority_queue_sched_stable_min
type Priority_queue_sched_stable_min struct {
	Pq_root unsafe.Pointer
}

// Proc_persona_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/proc_persona_info
type Proc_persona_info struct {
	Unique_pid uint64
	Pid        int32
	Flags      uint32
	Pidversion uint32
	Persona_id uint32
	Uid        uint32
	Gid        uint32
	Macho_uuid [16]uint8
}

// Proc_rlimit_control_wakeupmon
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/proc_rlimit_control_wakeupmon
type Proc_rlimit_control_wakeupmon struct {
	Wm_flags uint32
	Wm_rate  int32
}

// Pseminfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/pseminfo
type Pseminfo struct {
	Psem_flags            uint32
	Psem_usecount         uint32
	Psem_mode             uint16
	Psem_uid              uint32
	Psem_gid              uint32
	Psem_name             [32]int8
	Psem_semobject        unsafe.Pointer
	Psem_label            unsafe.Pointer
	Psem_creator_pid      int32
	Psem_creator_uniqueid uint64
}

// Pshminfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/pshminfo
type Pshminfo struct {
	Pshm_flags     uint32
	Pshm_usecount  uint32
	Pshm_length    int64
	Pshm_mode      uint16
	Pshm_uid       uint32
	Pshm_gid       uint32
	Pshm_name      [32]int8
	Pshm_memobject unsafe.Pointer
	Pshm_label     unsafe.Pointer
}

// Radvisory
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/radvisory
type Radvisory struct {
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
	storage [12]byte
}

// Ra_offset returns the Ra_offset field from the record's packed storage.
func (s *Radvisory) Ra_offset() int64 {
	return int64(binary.NativeEndian.Uint64(s.storage[0:8]))
}

// SetRa_offset updates the Ra_offset field in the record's packed storage.
func (s *Radvisory) SetRa_offset(v int64) {
	binary.NativeEndian.PutUint64(s.storage[0:8], uint64(v))
}

// Ra_count returns the Ra_count field from the record's packed storage.
func (s *Radvisory) Ra_count() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetRa_count updates the Ra_count field in the record's packed storage.
func (s *Radvisory) SetRa_count(v int32) {
	binary.NativeEndian.PutUint32(s.storage[8:12], uint32(v))
}

// Receive_sysdiagnose_notification_subsystem
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/receive_sysdiagnose_notification_subsystem
type Receive_sysdiagnose_notification_subsystem struct {
	Server   unsafe.Pointer
	Start    Mach_msg_id_t
	End      Mach_msg_id_t
	Maxsize  uint32
	Reserved Vm_address_t
	Routine  unsafe.Pointer
}

// Receive_vfs_nspace_subsystem
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/receive_vfs_nspace_subsystem-rhm
type Receive_vfs_nspace_subsystem struct {
	Server   unsafe.Pointer
	Start    Mach_msg_id_t
	End      Mach_msg_id_t
	Maxsize  uint32
	Reserved Vm_address_t
	Routine  unsafe.Pointer
}

// Reg_desc
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/reg_desc
type Reg_desc struct {
	Rd_mask   uint32
	Rd_shift  int32
	Rd_name   *byte
	Rd_format *byte
	Rd_values *Reg_values
}

// Reg_values
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/reg_values
type Reg_values struct {
	Rv_value uint32
	Rv_name  *byte
}

// Relocation_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/relocation_info
type Relocation_info struct {
	R_address int32 // In `MH_OBJECT` files, an offset from the start of the section to the item containing the address requiring relocation.
	bitfield1 uint32
}

// R_symbolnum returns the R_symbolnum bitfield.
func (s *Relocation_info) R_symbolnum() uint32 {
	return (s.bitfield1 >> 0) & ((1 << 24) - 1)
}

// SetR_symbolnum updates the R_symbolnum bitfield.
func (s *Relocation_info) SetR_symbolnum(v uint32) {
	const mask uint32 = (1 << 24) - 1
	s.bitfield1 = (s.bitfield1 &^ (mask << 0)) | ((v & mask) << 0)
}

// R_pcrel returns the R_pcrel bitfield.
func (s *Relocation_info) R_pcrel() uint32 {
	return (s.bitfield1 >> 24) & ((1 << 1) - 1)
}

// SetR_pcrel updates the R_pcrel bitfield.
func (s *Relocation_info) SetR_pcrel(v uint32) {
	const mask uint32 = (1 << 1) - 1
	s.bitfield1 = (s.bitfield1 &^ (mask << 24)) | ((v & mask) << 24)
}

// R_length returns the R_length bitfield.
func (s *Relocation_info) R_length() uint32 {
	return (s.bitfield1 >> 25) & ((1 << 2) - 1)
}

// SetR_length updates the R_length bitfield.
func (s *Relocation_info) SetR_length(v uint32) {
	const mask uint32 = (1 << 2) - 1
	s.bitfield1 = (s.bitfield1 &^ (mask << 25)) | ((v & mask) << 25)
}

// R_extern returns the R_extern bitfield.
func (s *Relocation_info) R_extern() uint32 {
	return (s.bitfield1 >> 27) & ((1 << 1) - 1)
}

// SetR_extern updates the R_extern bitfield.
func (s *Relocation_info) SetR_extern(v uint32) {
	const mask uint32 = (1 << 1) - 1
	s.bitfield1 = (s.bitfield1 &^ (mask << 27)) | ((v & mask) << 27)
}

// R_type returns the R_type bitfield.
func (s *Relocation_info) R_type() uint32 {
	return (s.bitfield1 >> 28) & ((1 << 4) - 1)
}

// SetR_type updates the R_type bitfield.
func (s *Relocation_info) SetR_type(v uint32) {
	const mask uint32 = (1 << 4) - 1
	s.bitfield1 = (s.bitfield1 &^ (mask << 28)) | ((v & mask) << 28)
}

// Rip6stat
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/rip6stat
type Rip6stat struct {
	Rip6s_ipackets    U_quad_t
	Rip6s_isum        U_quad_t
	Rip6s_badsum      U_quad_t
	Rip6s_nosock      U_quad_t
	Rip6s_nosockmcast U_quad_t
	Rip6s_fullsock    U_quad_t
	Rip6s_opackets    U_quad_t
}

// Rlimit
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/rlimit
type Rlimit struct {
	Rlim_cur Rlim_t
	Rlim_max Rlim_t
}

// Route_in6_old
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/route_in6_old
type Route_in6_old struct {
	Ro_rt    unsafe.Pointer
	Ro_flags uint32
	Ro_dst   Sockaddr_in6
}

// Route_old
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/route_old
type Route_old struct {
	Ro_rt    unsafe.Pointer
	Ro_flags uint32
	Ro_dst   [16]byte
}

// Routines_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/routines_command
type Routines_command struct {
	Cmd          uint32
	Cmdsize      uint32
	Init_address uint32 // An integer specifying the virtual memory address of the initialization function.
	Init_module  uint32
	Reserved1    uint32
	Reserved2    uint32 // Reserved for future use. Set this field to `0`.
	Reserved3    uint32 // Reserved for future use. Set this field to `0`.
	Reserved4    uint32
	Reserved5    uint32
	Reserved6    uint32
}

// Routines_command_64
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/routines_command_64
type Routines_command_64 struct {
	Cmd          uint32 // Common to all load command structures. For this structure, set to `LC_ROUTINES_64`.
	Cmdsize      uint32
	Init_address uint64
	Init_module  uint64
	Reserved1    uint64
	Reserved2    uint64 // Reserved for future use. Set this field to `0`.
	Reserved3    uint64
	Reserved4    uint64
	Reserved5    uint64 // Reserved for future use. Set this field to `0`.
	Reserved6    uint64
}

// Rpath_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/rpath_command
type Rpath_command struct {
	Cmd     uint32
	Cmdsize uint32
	Path    [1]uint32
}

// Rpc_signature
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/rpc_signature
type Rpc_signature struct {
	Rd  [5]uint64
	Rad unsafe.Pointer
}

// Rr_pco_match
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/rr_pco_match
type Rr_pco_match struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [24]byte
}

// Rpm_code returns the Rpm_code field from the record's packed storage.
func (s *Rr_pco_match) Rpm_code() U_int8_t {
	return *(*U_int8_t)(unsafe.Pointer(&s.storage[0]))
}

// SetRpm_code updates the Rpm_code field in the record's packed storage.
func (s *Rr_pco_match) SetRpm_code(v U_int8_t) {
	*(*U_int8_t)(unsafe.Pointer(&s.storage[0])) = v
}

// Rpm_len returns the Rpm_len field from the record's packed storage.
func (s *Rr_pco_match) Rpm_len() U_int8_t {
	return *(*U_int8_t)(unsafe.Pointer(&s.storage[1]))
}

// SetRpm_len updates the Rpm_len field in the record's packed storage.
func (s *Rr_pco_match) SetRpm_len(v U_int8_t) {
	*(*U_int8_t)(unsafe.Pointer(&s.storage[1])) = v
}

// Rpm_ordinal returns the Rpm_ordinal field from the record's packed storage.
func (s *Rr_pco_match) Rpm_ordinal() U_int8_t {
	return *(*U_int8_t)(unsafe.Pointer(&s.storage[2]))
}

// SetRpm_ordinal updates the Rpm_ordinal field in the record's packed storage.
func (s *Rr_pco_match) SetRpm_ordinal(v U_int8_t) {
	*(*U_int8_t)(unsafe.Pointer(&s.storage[2])) = v
}

// Rpm_matchlen returns the Rpm_matchlen field from the record's packed storage.
func (s *Rr_pco_match) Rpm_matchlen() U_int8_t {
	return *(*U_int8_t)(unsafe.Pointer(&s.storage[3]))
}

// SetRpm_matchlen updates the Rpm_matchlen field in the record's packed storage.
func (s *Rr_pco_match) SetRpm_matchlen(v U_int8_t) {
	*(*U_int8_t)(unsafe.Pointer(&s.storage[3])) = v
}

// Rpm_minlen returns the Rpm_minlen field from the record's packed storage.
func (s *Rr_pco_match) Rpm_minlen() U_int8_t {
	return *(*U_int8_t)(unsafe.Pointer(&s.storage[4]))
}

// SetRpm_minlen updates the Rpm_minlen field in the record's packed storage.
func (s *Rr_pco_match) SetRpm_minlen(v U_int8_t) {
	*(*U_int8_t)(unsafe.Pointer(&s.storage[4])) = v
}

// Rpm_maxlen returns the Rpm_maxlen field from the record's packed storage.
func (s *Rr_pco_match) Rpm_maxlen() U_int8_t {
	return *(*U_int8_t)(unsafe.Pointer(&s.storage[5]))
}

// SetRpm_maxlen updates the Rpm_maxlen field in the record's packed storage.
func (s *Rr_pco_match) SetRpm_maxlen(v U_int8_t) {
	*(*U_int8_t)(unsafe.Pointer(&s.storage[5])) = v
}

// Rpm_reserved returns the Rpm_reserved field from the record's packed storage.
func (s *Rr_pco_match) Rpm_reserved() U_int16_t {
	return U_int16_t(binary.NativeEndian.Uint16(s.storage[6:8]))
}

// SetRpm_reserved updates the Rpm_reserved field in the record's packed storage.
func (s *Rr_pco_match) SetRpm_reserved(v U_int16_t) {
	binary.NativeEndian.PutUint16(s.storage[6:8], uint16(v))
}

// Rpm_prefix returns the Rpm_prefix field from the record's packed storage.
func (s *Rr_pco_match) Rpm_prefix() [4]uint32 {
	return *(*[4]uint32)(unsafe.Pointer(&s.storage[8]))
}

// SetRpm_prefix updates the Rpm_prefix field in the record's packed storage.
func (s *Rr_pco_match) SetRpm_prefix(v [4]uint32) {
	*(*[4]uint32)(unsafe.Pointer(&s.storage[8])) = v
}

// Rr_pco_use
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/rr_pco_use
type Rr_pco_use struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [32]byte
}

// Rpu_uselen returns the Rpu_uselen field from the record's packed storage.
func (s *Rr_pco_use) Rpu_uselen() U_int8_t {
	return *(*U_int8_t)(unsafe.Pointer(&s.storage[0]))
}

// SetRpu_uselen updates the Rpu_uselen field in the record's packed storage.
func (s *Rr_pco_use) SetRpu_uselen(v U_int8_t) {
	*(*U_int8_t)(unsafe.Pointer(&s.storage[0])) = v
}

// Rpu_keeplen returns the Rpu_keeplen field from the record's packed storage.
func (s *Rr_pco_use) Rpu_keeplen() U_int8_t {
	return *(*U_int8_t)(unsafe.Pointer(&s.storage[1]))
}

// SetRpu_keeplen updates the Rpu_keeplen field in the record's packed storage.
func (s *Rr_pco_use) SetRpu_keeplen(v U_int8_t) {
	*(*U_int8_t)(unsafe.Pointer(&s.storage[1])) = v
}

// Rpu_ramask returns the Rpu_ramask field from the record's packed storage.
func (s *Rr_pco_use) Rpu_ramask() U_int8_t {
	return *(*U_int8_t)(unsafe.Pointer(&s.storage[2]))
}

// SetRpu_ramask updates the Rpu_ramask field in the record's packed storage.
func (s *Rr_pco_use) SetRpu_ramask(v U_int8_t) {
	*(*U_int8_t)(unsafe.Pointer(&s.storage[2])) = v
}

// Rpu_raflags returns the Rpu_raflags field from the record's packed storage.
func (s *Rr_pco_use) Rpu_raflags() U_int8_t {
	return *(*U_int8_t)(unsafe.Pointer(&s.storage[3]))
}

// SetRpu_raflags updates the Rpu_raflags field in the record's packed storage.
func (s *Rr_pco_use) SetRpu_raflags(v U_int8_t) {
	*(*U_int8_t)(unsafe.Pointer(&s.storage[3])) = v
}

// Rpu_vltime returns the Rpu_vltime field from the record's packed storage.
func (s *Rr_pco_use) Rpu_vltime() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetRpu_vltime updates the Rpu_vltime field in the record's packed storage.
func (s *Rr_pco_use) SetRpu_vltime(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// Rpu_pltime returns the Rpu_pltime field from the record's packed storage.
func (s *Rr_pco_use) Rpu_pltime() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetRpu_pltime updates the Rpu_pltime field in the record's packed storage.
func (s *Rr_pco_use) SetRpu_pltime(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[8:12], uint32(v))
}

// Rpu_flags returns the Rpu_flags field from the record's packed storage.
func (s *Rr_pco_use) Rpu_flags() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[12:16]))
}

// SetRpu_flags updates the Rpu_flags field in the record's packed storage.
func (s *Rr_pco_use) SetRpu_flags(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[12:16], uint32(v))
}

// Rpu_prefix returns the Rpu_prefix field from the record's packed storage.
func (s *Rr_pco_use) Rpu_prefix() [4]uint32 {
	return *(*[4]uint32)(unsafe.Pointer(&s.storage[16]))
}

// SetRpu_prefix updates the Rpu_prefix field in the record's packed storage.
func (s *Rr_pco_use) SetRpu_prefix(v [4]uint32) {
	*(*[4]uint32)(unsafe.Pointer(&s.storage[16])) = v
}

// Rr_result
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/rr_result
type Rr_result struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [24]byte
}

// Rrr_flags returns the Rrr_flags field from the record's packed storage.
func (s *Rr_result) Rrr_flags() U_int16_t {
	return U_int16_t(binary.NativeEndian.Uint16(s.storage[0:2]))
}

// SetRrr_flags updates the Rrr_flags field in the record's packed storage.
func (s *Rr_result) SetRrr_flags(v U_int16_t) {
	binary.NativeEndian.PutUint16(s.storage[0:2], uint16(v))
}

// Rrr_ordinal returns the Rrr_ordinal field from the record's packed storage.
func (s *Rr_result) Rrr_ordinal() U_int8_t {
	return *(*U_int8_t)(unsafe.Pointer(&s.storage[2]))
}

// SetRrr_ordinal updates the Rrr_ordinal field in the record's packed storage.
func (s *Rr_result) SetRrr_ordinal(v U_int8_t) {
	*(*U_int8_t)(unsafe.Pointer(&s.storage[2])) = v
}

// Rrr_matchedlen returns the Rrr_matchedlen field from the record's packed storage.
func (s *Rr_result) Rrr_matchedlen() U_int8_t {
	return *(*U_int8_t)(unsafe.Pointer(&s.storage[3]))
}

// SetRrr_matchedlen updates the Rrr_matchedlen field in the record's packed storage.
func (s *Rr_result) SetRrr_matchedlen(v U_int8_t) {
	*(*U_int8_t)(unsafe.Pointer(&s.storage[3])) = v
}

// Rrr_ifid returns the Rrr_ifid field from the record's packed storage.
func (s *Rr_result) Rrr_ifid() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetRrr_ifid updates the Rrr_ifid field in the record's packed storage.
func (s *Rr_result) SetRrr_ifid(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// Rrr_prefix returns the Rrr_prefix field from the record's packed storage.
func (s *Rr_result) Rrr_prefix() [4]uint32 {
	return *(*[4]uint32)(unsafe.Pointer(&s.storage[8]))
}

// SetRrr_prefix updates the Rrr_prefix field in the record's packed storage.
func (s *Rr_result) SetRrr_prefix(v [4]uint32) {
	*(*[4]uint32)(unsafe.Pointer(&s.storage[8])) = v
}

// Rslvmulti_req
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/rslvmulti_req
type Rslvmulti_req struct {
	Sa   Pointer
	Llsa *objc.ID
}

// Rt_addrinfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/rt_addrinfo
type Rt_addrinfo struct {
	Rti_addrs int32
	Rti_info  *objc.ID
}

// Rt_addrinfo_ext
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/rt_addrinfo_ext
type Rt_addrinfo_ext struct {
	Rtix_info      Rt_addrinfo
	Rtix_tiny_addr unsafe.Pointer
	Rtix_next_tiny uint8
}

// Rt_metrics
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/rt_metrics
type Rt_metrics struct {
	Rmx_locks    U_int32_t
	Rmx_mtu      U_int32_t
	Rmx_hopcount U_int32_t
	Rmx_expire   int32
	Rmx_recvpipe U_int32_t
	Rmx_sendpipe U_int32_t
	Rmx_ssthresh U_int32_t
	Rmx_rtt      U_int32_t
	Rmx_rttvar   U_int32_t
	Rmx_pksent   U_int32_t
	Rmx_filler   [4]U_int32_t
}

// Rt_msghdr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/rt_msghdr
type Rt_msghdr struct {
	Rtm_msglen  U_short
	Rtm_version U_char
	Rtm_type    U_char
	Rtm_index   U_short
	Rtm_flags   int32
	Rtm_addrs   int32
	Rtm_pid     int32
	Rtm_seq     int32
	Rtm_errno   int32
	Rtm_use     int32
	Rtm_inits   U_int32_t
	Rtm_rmx     Rt_metrics
}

// Rt_msghdr2
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/rt_msghdr2
type Rt_msghdr2 struct {
	Rtm_msglen      U_short
	Rtm_version     U_char
	Rtm_type        U_char
	Rtm_index       U_short
	Rtm_flags       int32
	Rtm_addrs       int32
	Rtm_refcnt      int32
	Rtm_parentflags int32
	Rtm_reserved    int32
	Rtm_use         int32
	Rtm_inits       U_int32_t
	Rtm_rmx         Rt_metrics
}

// Rt_msghdr_common
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/rt_msghdr_common
type Rt_msghdr_common struct {
	Rtm_msglen  U_short
	Rtm_version U_char
	Rtm_type    U_char
	Rtm_index   U_short
	Rtm_flags   int32
	Rtm_addrs   int32
	Rtm_pid     int32
	Rtm_seq     int32
	Rtm_errno   int32
	Rtm_use     int32
}

// Rt_msghdr_ext
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/rt_msghdr_ext
type Rt_msghdr_ext struct {
	Rtm_msglen   U_short
	Rtm_version  U_char
	Rtm_type     U_char
	Rtm_index    U_int32_t
	Rtm_flags    U_int32_t
	Rtm_reserved U_int32_t
	Rtm_addrs    U_int32_t
	Rtm_pid      int32
	Rtm_seq      int32
	Rtm_errno    int32
	Rtm_use      U_int32_t
	Rtm_inits    U_int32_t
	Rtm_rmx      Rt_metrics
	Rtm_ri       Rt_reach_info
}

// Rt_msghdr_prelude
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/rt_msghdr_prelude
type Rt_msghdr_prelude struct {
	Rtm_msglen U_short
}

// Rt_reach_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/rt_reach_info
type Rt_reach_info struct {
	Ri_refcnt     U_int32_t
	Ri_probes     U_int32_t
	Ri_snd_expire U_int64_t
	Ri_rcv_expire U_int64_t
	Ri_rssi       int32
	Ri_lqm        int32
	Ri_npm        int32
}

// Rtstat
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/rtstat
type Rtstat struct {
	Rts_badredirect  int16
	Rts_dynamic      int16
	Rts_newgateway   int16
	Rts_unreach      int16
	Rts_wildcard     int16
	Rts_badrtgwroute int16
}

// Rtstat_64
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/rtstat_64
type Rtstat_64 struct {
	Rts_badredirect  uint64
	Rts_dynamic      uint64
	Rts_newgateway   uint64
	Rts_unreach      uint64
	Rts_wildcard     uint64
	Rts_badrtgwroute uint64
}

// Rusage
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/rusage
type Rusage struct {
	Ru_utime    Timeval
	Ru_stime    Timeval
	Ru_maxrss   int
	Ru_ixrss    int
	Ru_idrss    int
	Ru_isrss    int
	Ru_minflt   int
	Ru_majflt   int
	Ru_nswap    int
	Ru_inblock  int
	Ru_oublock  int
	Ru_msgsnd   int
	Ru_msgrcv   int
	Ru_nsignals int
	Ru_nvcsw    int
	Ru_nivcsw   int
}

// Rusage_info_v0
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/rusage_info_v0
type Rusage_info_v0 struct {
	Ri_uuid               [16]uint8
	Ri_user_time          uint64
	Ri_system_time        uint64
	Ri_pkg_idle_wkups     uint64
	Ri_interrupt_wkups    uint64
	Ri_pageins            uint64
	Ri_wired_size         uint64
	Ri_resident_size      uint64
	Ri_phys_footprint     uint64
	Ri_proc_start_abstime uint64
	Ri_proc_exit_abstime  uint64
}

// Rusage_info_v1
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/rusage_info_v1
type Rusage_info_v1 struct {
	Ri_uuid                  [16]uint8
	Ri_user_time             uint64
	Ri_system_time           uint64
	Ri_pkg_idle_wkups        uint64
	Ri_interrupt_wkups       uint64
	Ri_pageins               uint64
	Ri_wired_size            uint64
	Ri_resident_size         uint64
	Ri_phys_footprint        uint64
	Ri_proc_start_abstime    uint64
	Ri_proc_exit_abstime     uint64
	Ri_child_user_time       uint64
	Ri_child_system_time     uint64
	Ri_child_pkg_idle_wkups  uint64
	Ri_child_interrupt_wkups uint64
	Ri_child_pageins         uint64
	Ri_child_elapsed_abstime uint64
}

// Rusage_info_v2
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/rusage_info_v2
type Rusage_info_v2 struct {
	Ri_uuid                  [16]uint8
	Ri_user_time             uint64
	Ri_system_time           uint64
	Ri_pkg_idle_wkups        uint64
	Ri_interrupt_wkups       uint64
	Ri_pageins               uint64
	Ri_wired_size            uint64
	Ri_resident_size         uint64
	Ri_phys_footprint        uint64
	Ri_proc_start_abstime    uint64
	Ri_proc_exit_abstime     uint64
	Ri_child_user_time       uint64
	Ri_child_system_time     uint64
	Ri_child_pkg_idle_wkups  uint64
	Ri_child_interrupt_wkups uint64
	Ri_child_pageins         uint64
	Ri_child_elapsed_abstime uint64
	Ri_diskio_bytesread      uint64
	Ri_diskio_byteswritten   uint64
}

// Rusage_info_v3
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/rusage_info_v3
type Rusage_info_v3 struct {
	Ri_uuid                          [16]uint8
	Ri_user_time                     uint64
	Ri_system_time                   uint64
	Ri_pkg_idle_wkups                uint64
	Ri_interrupt_wkups               uint64
	Ri_pageins                       uint64
	Ri_wired_size                    uint64
	Ri_resident_size                 uint64
	Ri_phys_footprint                uint64
	Ri_proc_start_abstime            uint64
	Ri_proc_exit_abstime             uint64
	Ri_child_user_time               uint64
	Ri_child_system_time             uint64
	Ri_child_pkg_idle_wkups          uint64
	Ri_child_interrupt_wkups         uint64
	Ri_child_pageins                 uint64
	Ri_child_elapsed_abstime         uint64
	Ri_diskio_bytesread              uint64
	Ri_diskio_byteswritten           uint64
	Ri_cpu_time_qos_default          uint64
	Ri_cpu_time_qos_maintenance      uint64
	Ri_cpu_time_qos_background       uint64
	Ri_cpu_time_qos_utility          uint64
	Ri_cpu_time_qos_legacy           uint64
	Ri_cpu_time_qos_user_initiated   uint64
	Ri_cpu_time_qos_user_interactive uint64
	Ri_billed_system_time            uint64
	Ri_serviced_system_time          uint64
}

// Rusage_info_v4
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/rusage_info_v4
type Rusage_info_v4 struct {
	Ri_uuid                          [16]uint8
	Ri_user_time                     uint64
	Ri_system_time                   uint64
	Ri_pkg_idle_wkups                uint64
	Ri_interrupt_wkups               uint64
	Ri_pageins                       uint64
	Ri_wired_size                    uint64
	Ri_resident_size                 uint64
	Ri_phys_footprint                uint64
	Ri_proc_start_abstime            uint64
	Ri_proc_exit_abstime             uint64
	Ri_child_user_time               uint64
	Ri_child_system_time             uint64
	Ri_child_pkg_idle_wkups          uint64
	Ri_child_interrupt_wkups         uint64
	Ri_child_pageins                 uint64
	Ri_child_elapsed_abstime         uint64
	Ri_diskio_bytesread              uint64
	Ri_diskio_byteswritten           uint64
	Ri_cpu_time_qos_default          uint64
	Ri_cpu_time_qos_maintenance      uint64
	Ri_cpu_time_qos_background       uint64
	Ri_cpu_time_qos_utility          uint64
	Ri_cpu_time_qos_legacy           uint64
	Ri_cpu_time_qos_user_initiated   uint64
	Ri_cpu_time_qos_user_interactive uint64
	Ri_billed_system_time            uint64
	Ri_serviced_system_time          uint64
	Ri_logical_writes                uint64
	Ri_lifetime_max_phys_footprint   uint64
	Ri_instructions                  uint64
	Ri_cycles                        uint64
	Ri_billed_energy                 uint64
	Ri_serviced_energy               uint64
	Ri_interval_max_phys_footprint   uint64
	Ri_runnable_time                 uint64
}

// Rusage_info_v5
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/rusage_info_v5
type Rusage_info_v5 struct {
	Ri_uuid                          [16]uint8
	Ri_user_time                     uint64
	Ri_system_time                   uint64
	Ri_pkg_idle_wkups                uint64
	Ri_interrupt_wkups               uint64
	Ri_pageins                       uint64
	Ri_wired_size                    uint64
	Ri_resident_size                 uint64
	Ri_phys_footprint                uint64
	Ri_proc_start_abstime            uint64
	Ri_proc_exit_abstime             uint64
	Ri_child_user_time               uint64
	Ri_child_system_time             uint64
	Ri_child_pkg_idle_wkups          uint64
	Ri_child_interrupt_wkups         uint64
	Ri_child_pageins                 uint64
	Ri_child_elapsed_abstime         uint64
	Ri_diskio_bytesread              uint64
	Ri_diskio_byteswritten           uint64
	Ri_cpu_time_qos_default          uint64
	Ri_cpu_time_qos_maintenance      uint64
	Ri_cpu_time_qos_background       uint64
	Ri_cpu_time_qos_utility          uint64
	Ri_cpu_time_qos_legacy           uint64
	Ri_cpu_time_qos_user_initiated   uint64
	Ri_cpu_time_qos_user_interactive uint64
	Ri_billed_system_time            uint64
	Ri_serviced_system_time          uint64
	Ri_logical_writes                uint64
	Ri_lifetime_max_phys_footprint   uint64
	Ri_instructions                  uint64
	Ri_cycles                        uint64
	Ri_billed_energy                 uint64
	Ri_serviced_energy               uint64
	Ri_interval_max_phys_footprint   uint64
	Ri_runnable_time                 uint64
	Ri_flags                         uint64
}

// Sadb_address
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sadb_address
type Sadb_address struct {
	Sadb_address_len       U_int16_t
	Sadb_address_exttype   U_int16_t
	Sadb_address_proto     U_int8_t
	Sadb_address_prefixlen U_int8_t
	Sadb_address_reserved  U_int16_t
}

// Sadb_alg
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sadb_alg
type Sadb_alg struct {
	Sadb_alg_id       U_int8_t
	Sadb_alg_ivlen    U_int8_t
	Sadb_alg_minbits  U_int16_t
	Sadb_alg_maxbits  U_int16_t
	Sadb_alg_reserved U_int16_t
}

// Sadb_comb
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sadb_comb
type Sadb_comb struct {
	Sadb_comb_auth             U_int8_t
	Sadb_comb_encrypt          U_int8_t
	Sadb_comb_flags            U_int16_t
	Sadb_comb_auth_minbits     U_int16_t
	Sadb_comb_auth_maxbits     U_int16_t
	Sadb_comb_encrypt_minbits  U_int16_t
	Sadb_comb_encrypt_maxbits  U_int16_t
	Sadb_comb_reserved         U_int32_t
	Sadb_comb_soft_allocations U_int32_t
	Sadb_comb_hard_allocations U_int32_t
	Sadb_comb_soft_bytes       U_int64_t
	Sadb_comb_hard_bytes       U_int64_t
	Sadb_comb_soft_addtime     U_int64_t
	Sadb_comb_hard_addtime     U_int64_t
	Sadb_comb_soft_usetime     U_int64_t
	Sadb_comb_hard_usetime     U_int64_t
}

// Sadb_ext
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sadb_ext
type Sadb_ext struct {
	Sadb_ext_len  U_int16_t
	Sadb_ext_type U_int16_t
}

// Sadb_ident
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sadb_ident
type Sadb_ident struct {
	Sadb_ident_len      U_int16_t
	Sadb_ident_exttype  U_int16_t
	Sadb_ident_type     U_int16_t
	Sadb_ident_reserved U_int16_t
	Sadb_ident_id       U_int64_t
}

// Sadb_key
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sadb_key
type Sadb_key struct {
	Sadb_key_len      U_int16_t
	Sadb_key_exttype  U_int16_t
	Sadb_key_bits     U_int16_t
	Sadb_key_reserved U_int16_t
}

// Sadb_lifetime
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sadb_lifetime
type Sadb_lifetime struct {
	Sadb_lifetime_len         U_int16_t
	Sadb_lifetime_exttype     U_int16_t
	Sadb_lifetime_allocations U_int32_t
	Sadb_lifetime_bytes       U_int64_t
	Sadb_lifetime_addtime     U_int64_t
	Sadb_lifetime_usetime     U_int64_t
}

// Sadb_msg
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sadb_msg
type Sadb_msg struct {
	Sadb_msg_version  U_int8_t
	Sadb_msg_type     U_int8_t
	Sadb_msg_errno    U_int8_t
	Sadb_msg_satype   U_int8_t
	Sadb_msg_len      U_int16_t
	Sadb_msg_reserved U_int16_t
	Sadb_msg_seq      U_int32_t
	Sadb_msg_pid      U_int32_t
}

// Sadb_prop
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sadb_prop
type Sadb_prop struct {
	Sadb_prop_len      U_int16_t
	Sadb_prop_exttype  U_int16_t
	Sadb_prop_replay   U_int8_t
	Sadb_prop_reserved [3]U_int8_t
}

// Sadb_sa
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sadb_sa
type Sadb_sa struct {
	Sadb_sa_len     U_int16_t
	Sadb_sa_exttype U_int16_t
	Sadb_sa_spi     U_int32_t
	Sadb_sa_replay  U_int8_t
	Sadb_sa_state   U_int8_t
	Sadb_sa_auth    U_int8_t
	Sadb_sa_encrypt U_int8_t
	Sadb_sa_flags   U_int32_t
}

// Sadb_sastat
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sadb_sastat
type Sadb_sastat struct {
	Sadb_sastat_len      U_int16_t
	Sadb_sastat_exttype  U_int16_t
	Sadb_sastat_dir      U_int32_t
	Sadb_sastat_reserved U_int32_t
	Sadb_sastat_list_len U_int32_t
}

// Sadb_sens
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sadb_sens
type Sadb_sens struct {
	Sadb_sens_len         U_int16_t
	Sadb_sens_exttype     U_int16_t
	Sadb_sens_dpd         U_int32_t
	Sadb_sens_sens_level  U_int8_t
	Sadb_sens_sens_len    U_int8_t
	Sadb_sens_integ_level U_int8_t
	Sadb_sens_integ_len   U_int8_t
	Sadb_sens_reserved    U_int32_t
}

// Sadb_session_id
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sadb_session_id
type Sadb_session_id struct {
	Sadb_session_id_len     U_int16_t
	Sadb_session_id_exttype U_int16_t
	Sadb_session_id_v       [2]U_int64_t
}

// Sadb_spirange
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sadb_spirange
type Sadb_spirange struct {
	Sadb_spirange_len      U_int16_t
	Sadb_spirange_exttype  U_int16_t
	Sadb_spirange_min      U_int32_t
	Sadb_spirange_max      U_int32_t
	Sadb_spirange_reserved U_int32_t
}

// Sadb_supported
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sadb_supported
type Sadb_supported struct {
	Sadb_supported_len      U_int16_t
	Sadb_supported_exttype  U_int16_t
	Sadb_supported_reserved U_int32_t
}

// Sadb_x_ipsecrequest
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sadb_x_ipsecrequest
type Sadb_x_ipsecrequest struct {
	Sadb_x_ipsecrequest_len   U_int16_t
	Sadb_x_ipsecrequest_proto U_int16_t
	Sadb_x_ipsecrequest_mode  U_int8_t
	Sadb_x_ipsecrequest_level U_int8_t
	Sadb_x_ipsecrequest_reqid U_int16_t
}

// Sadb_x_kmprivate
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sadb_x_kmprivate
type Sadb_x_kmprivate struct {
	Sadb_x_kmprivate_len      U_int16_t
	Sadb_x_kmprivate_exttype  U_int16_t
	Sadb_x_kmprivate_reserved U_int32_t
}

// Sadb_x_policy
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sadb_x_policy
type Sadb_x_policy struct {
	Sadb_x_policy_len       U_int16_t
	Sadb_x_policy_exttype   U_int16_t
	Sadb_x_policy_type      U_int16_t
	Sadb_x_policy_dir       U_int8_t
	Sadb_x_policy_reserved  U_int8_t
	Sadb_x_policy_id        U_int32_t
	Sadb_x_policy_reserved2 U_int32_t
}

// Sadb_x_sa2
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sadb_x_sa2
type Sadb_x_sa2 struct {
	Sadb_x_sa2_len       U_int16_t
	Sadb_x_sa2_exttype   U_int16_t
	Sadb_x_sa2_mode      U_int8_t
	Sadb_x_sa2_reserved1 U_int8_t
	Sadb_x_sa2_reserved2 U_int16_t
	Sadb_x_sa2_sequence  U_int32_t
	Sadb_x_sa2_reqid     U_int32_t
}

// Sastat
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sastat
type Sastat struct {
	Spi     U_int32_t
	Created U_int32_t
	Lft_c   Sadb_lifetime
}

// Sbuf
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sbuf
type Sbuf struct {
	S_buf    *byte
	S_unused unsafe.Pointer
	S_size   int32
	S_len    int32
	S_flags  int32
}

// Scattered_relocation_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/scattered_relocation_info
type Scattered_relocation_info struct {
	bitfield0 uint32
	R_value   int32 // The address of the relocatable expression for the item in the file that needs to be updated if the address is changed. For relocatable expressions with the difference of two section addresses, the address from which to subtract (in mathematical terms, the minuend) is contained in the first relocation entry and the address to subtract (the subtrahend) is contained in the second relocation entry.

}

// R_address returns the R_address bitfield.
func (s *Scattered_relocation_info) R_address() uint32 {
	return (s.bitfield0 >> 0) & ((1 << 24) - 1)
}

// SetR_address updates the R_address bitfield.
func (s *Scattered_relocation_info) SetR_address(v uint32) {
	const mask uint32 = (1 << 24) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 0)) | ((v & mask) << 0)
}

// R_type returns the R_type bitfield.
func (s *Scattered_relocation_info) R_type() uint32 {
	return (s.bitfield0 >> 24) & ((1 << 4) - 1)
}

// SetR_type updates the R_type bitfield.
func (s *Scattered_relocation_info) SetR_type(v uint32) {
	const mask uint32 = (1 << 4) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 24)) | ((v & mask) << 24)
}

// R_length returns the R_length bitfield.
func (s *Scattered_relocation_info) R_length() uint32 {
	return (s.bitfield0 >> 28) & ((1 << 2) - 1)
}

// SetR_length updates the R_length bitfield.
func (s *Scattered_relocation_info) SetR_length(v uint32) {
	const mask uint32 = (1 << 2) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 28)) | ((v & mask) << 28)
}

// R_pcrel returns the R_pcrel bitfield.
func (s *Scattered_relocation_info) R_pcrel() uint32 {
	return (s.bitfield0 >> 30) & ((1 << 1) - 1)
}

// SetR_pcrel updates the R_pcrel bitfield.
func (s *Scattered_relocation_info) SetR_pcrel(v uint32) {
	const mask uint32 = (1 << 1) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 30)) | ((v & mask) << 30)
}

// R_scattered returns the R_scattered bitfield.
func (s *Scattered_relocation_info) R_scattered() uint32 {
	return (s.bitfield0 >> 31) & ((1 << 1) - 1)
}

// SetR_scattered updates the R_scattered bitfield.
func (s *Scattered_relocation_info) SetR_scattered(v uint32) {
	const mask uint32 = (1 << 1) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 31)) | ((v & mask) << 31)
}

// Searchstate
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/searchstate
type Searchstate struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [556]byte
}

// Ss_union_flags returns the Ss_union_flags field from the record's packed storage.
func (s *Searchstate) Ss_union_flags() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetSs_union_flags updates the Ss_union_flags field in the record's packed storage.
func (s *Searchstate) SetSs_union_flags(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// Ss_union_layer returns the Ss_union_layer field from the record's packed storage.
func (s *Searchstate) Ss_union_layer() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetSs_union_layer updates the Ss_union_layer field in the record's packed storage.
func (s *Searchstate) SetSs_union_layer(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// Ss_fsstate returns the Ss_fsstate field from the record's packed storage.
func (s *Searchstate) Ss_fsstate() [548]U_char {
	return *(*[548]U_char)(unsafe.Pointer(&s.storage[8]))
}

// SetSs_fsstate updates the Ss_fsstate field in the record's packed storage.
func (s *Searchstate) SetSs_fsstate(v [548]U_char) {
	*(*[548]U_char)(unsafe.Pointer(&s.storage[8])) = v
}

// Section
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/section
type Section struct {
	Sectname  [16]int8
	Segname   [16]int8
	Addr      uint32
	Size      uint32
	Offset    uint32
	Align     uint32
	Reloff    uint32
	Nreloc    uint32
	Flags     uint32
	Reserved1 uint32
	Reserved2 uint32
}

// Section_64
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/section_64
type Section_64 struct {
	Sectname  [16]int8
	Segname   [16]int8 // A string specifying the name of the segment that should eventually contain this section. For compactness, intermediate object files—files of type `MH_OBJECT`—contain only one segment, in which all sections are placed. The static linker places each section in the named segment when building the final product (any file that is not of type `MH_OBJECT`).
	Addr      uint64
	Size      uint64
	Offset    uint32
	Align     uint32
	Reloff    uint32
	Nreloc    uint32
	Flags     uint32
	Reserved1 uint32
	Reserved2 uint32
	Reserved3 uint32
}

// Segment_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/segment_command
type Segment_command struct {
	Cmd      uint32 // Common to all load command structures. Set to `LC_SEGMENT` for this structure.
	Cmdsize  uint32
	Segname  [16]int8
	Vmaddr   uint32
	Vmsize   uint32 // Indicates the number of bytes of virtual memory occupied by this segment. See also the description of `filesize`, below.
	Fileoff  uint32 // Indicates the offset in this file of the data to be mapped at `vmaddr`.
	Filesize uint32
	Maxprot  Vm_prot_t
	Initprot Vm_prot_t
	Nsects   uint32
	Flags    uint32
}

// Segment_command_64
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/segment_command_64
type Segment_command_64 struct {
	Cmd      uint32
	Cmdsize  uint32
	Segname  [16]int8
	Vmaddr   uint64
	Vmsize   uint64 // Indicates the number of bytes of virtual memory occupied by this segment. See also the description of `filesize`, below.
	Fileoff  uint64
	Filesize uint64 // Indicates the number of bytes occupied by this segment on disk. For segments that require more memory at runtime than they do at build time, `vmsize` can be larger than `filesize`. For example, the `__PAGEZERO` segment generated by the linker for `MH_EXECUTABLE` files has a `vmsize` of 0x1000 but a `filesize` of 0. Because `__PAGEZERO` contains no data, there is no need for it to occupy any space until runtime. Also, the static linker often allocates uninitialized data at the end of the `__DATA` segment; in this case, the `vmsize` is larger than the `filesize`. The loader guarantees that any memory of this sort is initialized with zeros.
	Maxprot  Vm_prot_t
	Initprot Vm_prot_t
	Nsects   uint32 // Indicates the number of section data structures following this load command.
	Flags    uint32
}

// Sem
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sem
type Sem struct {
	Semval  uint16
	Sempid  int32
	Semncnt uint16
	Semzcnt uint16
}

// Sembuf
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sembuf
type Sembuf struct {
	Sem_num uint16
	Sem_op  int16
	Sem_flg int16
}

// Sf_hdtr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sf_hdtr
type Sf_hdtr struct {
	Headers  *Iovec
	Hdr_cnt  int32
	Trailers *Iovec
	Trl_cnt  int32
}

// Sflt_filter
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sflt_filter
type Sflt_filter struct {
	Sf_handle       Sflt_handle    // A value used to find socket filters by applications. An application can use this value to specify that this filter should be attached when using the SO_NKE socket option.
	Sf_flags        int32          // Indicate whether this filter should be attached to all new sockets or just those that request the filter be attached using the SO_NKE socket option. If this filter utilizes the socket filter extension fields, it must also set SFLT_EXTENDED.
	Sf_name         *byte          // A name used for debug purposes.
	Sf_unregistered unsafe.Pointer // Your function for being notified when your filter has been unregistered.
	Sf_attach       unsafe.Pointer // Your function for handling attaches to sockets.
	Sf_detach       unsafe.Pointer // Your function for handling detaches from sockets.
	Sf_notify       unsafe.Pointer // Your function for handling events. May be null.
	Sf_getpeername  unsafe.Pointer
	Sf_getsockname  unsafe.Pointer
	Sf_data_in      unsafe.Pointer // Your function for handling incoming data. May be null.
	Sf_data_out     unsafe.Pointer // Your function for handling outgoing data. May be null.
	Sf_connect_in   unsafe.Pointer // Your function for handling inbound connections. May be null.
	Sf_connect_out  unsafe.Pointer // Your function for handling outbound connections. May be null.
	Sf_bind         unsafe.Pointer // Your function for handling binds. May be null.
	Sf_setoption    unsafe.Pointer // Your function for handling setsockopt. May be null.
	Sf_getoption    unsafe.Pointer // Your function for handling getsockopt. May be null.
	Sf_listen       unsafe.Pointer // Your function for handling listen. May be null.
	Sf_ioctl        unsafe.Pointer // Your function for handling ioctls. May be null.
	Sf_ext_len      uint32
	Sf_ext_accept   unsafe.Pointer
	Sf_ext_rsvd     unsafe.Pointer
	Sf_ext          [7]uint64
}

// Sgttyb
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sgttyb
type Sgttyb struct {
	Sg_ispeed int8
	Sg_ospeed int8
	Sg_erase  int8
	Sg_kill   int8
	Sg_flags  int16
}

// Shared_file_mapping_np
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/shared_file_mapping_np
type Shared_file_mapping_np struct {
	Sfm_address     Mach_vm_address_t
	Sfm_size        Mach_vm_size_t
	Sfm_file_offset Mach_vm_offset_t
	Sfm_max_prot    Vm_prot_t
	Sfm_init_prot   Vm_prot_t
}

// Shared_file_np
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/shared_file_np
type Shared_file_np struct {
	Sf_fd             int32
	Sf_mappings_count uint32
	Sf_slide          uint32
}

// Shared_region_range_np
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/shared_region_range_np
type Shared_region_range_np struct {
	Srr_address Mach_vm_address_t
	Srr_size    Mach_vm_size_t
}

// Sigaction
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sigaction
type Sigaction struct {
	__sigaction_u unsafe.Pointer
	Sa_mask       Sigset_t
	Sa_flags      int32
}

// Sigstack
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sigstack
type Sigstack struct {
	Ss_sp      *byte
	Ss_onstack int32
}

// Sigvec
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sigvec
type Sigvec struct {
	Sv_handler unsafe.Pointer
	Sv_mask    int32
	Sv_flags   int32
}

// Smrq_link
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/smrq_link
type Smrq_link struct {
	Next unsafe.Pointer
	Prev unsafe.Pointer
}

// Smrq_list_head
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/smrq_list_head
type Smrq_list_head struct {
	First unsafe.Pointer
}

// Smrq_slink
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/smrq_slink
type Smrq_slink struct {
	Next unsafe.Pointer
}

// Smrq_slist_head
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/smrq_slist_head
type Smrq_slist_head struct {
	First unsafe.Pointer
}

// Smrq_stailq_head
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/smrq_stailq_head
type Smrq_stailq_head struct {
	First unsafe.Pointer
	Last  unsafe.Pointer
}

// Smrq_tailq_head
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/smrq_tailq_head
type Smrq_tailq_head struct {
	First unsafe.Pointer
	Last  unsafe.Pointer
}

// So_nke
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/so_nke
type So_nke struct {
	Nke_handle uint32
	Nke_where  uint32
	Nke_flags  int32
	Reserved   [4]U_int32_t
}

// So_np_extensions
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/so_np_extensions
type So_np_extensions struct {
	Npx_flags U_int32_t
	Npx_mask  U_int32_t
}

// Sockaddr_ctl
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sockaddr_ctl
type Sockaddr_ctl struct {
	Sc_len      U_char // The length of the structure.
	Sc_family   U_char // AF_SYSTEM.
	Ss_sysaddr  U_int16_t
	Sc_id       U_int32_t
	Sc_unit     U_int32_t
	Sc_reserved [5]U_int32_t
}

// Sockaddr_dl
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sockaddr_dl
type Sockaddr_dl struct {
	Sdl_len    U_char
	Sdl_family U_char
	Sdl_index  U_short
	Sdl_type   U_char
	Sdl_nlen   U_char
	Sdl_alen   U_char
	Sdl_slen   U_char
	Sdl_data   [12]int8
}

// Sockaddr_in
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sockaddr_in
type Sockaddr_in struct {
	Sin_len    uint8
	Sin_family uint8
	Sin_port   uint16
	Sin_addr   In_addr
	Sin_zero   [8]int8
}

// Sockaddr_in6
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sockaddr_in6
type Sockaddr_in6 struct {
	Sin6_len      uint8
	Sin6_family   uint8
	Sin6_port     uint16
	Sin6_flowinfo uint32
	Sin6_addr     [4]uint32
	Sin6_scope_id uint32
}

// Sockaddr_inarp
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sockaddr_inarp
type Sockaddr_inarp struct {
	Sin_len     U_char
	Sin_family  U_char
	Sin_port    U_short
	Sin_addr    In_addr
	Sin_srcaddr In_addr
	Sin_tos     U_short
	Sin_other   U_short
}

// Sockaddr_inifscope
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sockaddr_inifscope
type Sockaddr_inifscope struct {
	Sin_len    uint8
	Sin_family uint8
	Sin_port   uint16
	Sin_addr   In_addr
	Un         [2]uint32
}

// Sockaddr_ndrv
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sockaddr_ndrv
type Sockaddr_ndrv struct {
	Snd_len    uint8
	Snd_family uint8
	Snd_name   [16]uint8
}

// Sockaddr_sys
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sockaddr_sys
type Sockaddr_sys struct {
	Ss_len      U_char
	Ss_family   U_char
	Ss_sysaddr  U_int16_t
	Ss_reserved [7]U_int32_t
}

// Sockaddr_un
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sockaddr_un
type Sockaddr_un struct {
	Sun_len    uint8
	Sun_family uint8
	Sun_path   [104]int8
}

// Sockaddr_vm
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sockaddr_vm
type Sockaddr_vm struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [12]byte
}

// Svm_len returns the Svm_len field from the record's packed storage.
func (s *Sockaddr_vm) Svm_len() uint8 {
	return uint8(s.storage[0])
}

// SetSvm_len updates the Svm_len field in the record's packed storage.
func (s *Sockaddr_vm) SetSvm_len(v uint8) {
	s.storage[0] = uint8(v)
}

// Svm_family returns the Svm_family field from the record's packed storage.
func (s *Sockaddr_vm) Svm_family() uint8 {
	return uint8(s.storage[1])
}

// SetSvm_family updates the Svm_family field in the record's packed storage.
func (s *Sockaddr_vm) SetSvm_family(v uint8) {
	s.storage[1] = uint8(v)
}

// Svm_reserved1 returns the Svm_reserved1 field from the record's packed storage.
func (s *Sockaddr_vm) Svm_reserved1() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[2:4]))
}

// SetSvm_reserved1 updates the Svm_reserved1 field in the record's packed storage.
func (s *Sockaddr_vm) SetSvm_reserved1(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[2:4], uint16(v))
}

// Svm_port returns the Svm_port field from the record's packed storage.
func (s *Sockaddr_vm) Svm_port() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetSvm_port updates the Svm_port field in the record's packed storage.
func (s *Sockaddr_vm) SetSvm_port(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// Svm_cid returns the Svm_cid field from the record's packed storage.
func (s *Sockaddr_vm) Svm_cid() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetSvm_cid updates the Svm_cid field in the record's packed storage.
func (s *Sockaddr_vm) SetSvm_cid(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[8:12], uint32(v))
}

// Sockproto
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sockproto
type Sockproto struct {
	Sp_family   uint16
	Sp_protocol uint16
}

// Source_version_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/source_version_command
type Source_version_command struct {
	Cmd     uint32
	Cmdsize uint32
	Version uint64
}

// Specinfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/specinfo
type Specinfo struct {
	Si_hashchain     *objc.ID
	Si_specnext      unsafe.Pointer
	Si_flags         int
	Si_rdev          int32
	Si_opencount     int32
	Si_size          Daddr_t
	Si_lastr         Daddr64_t
	Si_devsize       U_int64_t
	Si_initted       U_int8_t
	Si_throttleable  U_int8_t
	Si_isssd         U_int16_t
	Si_devbsdunit    U_int32_t
	Si_throttle_mask U_int64_t
	Si_mountingowner Thread_t
}

// Stack_snapshot_frame32
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/stack_snapshot_frame32
type Stack_snapshot_frame32 struct {
	Lr uint32
	Sp uint32
}

// Stack_snapshot_frame64
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/stack_snapshot_frame64
type Stack_snapshot_frame64 struct {
	Lr uint64
	Sp uint64
}

// Stack_snapshot_stacktop
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/stack_snapshot_stacktop
type Stack_snapshot_stacktop struct {
	Sp             uint64
	Stack_contents [8]uint8
}

// Stackshot_cpu_architecture
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/stackshot_cpu_architecture
type Stackshot_cpu_architecture struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [8]byte
}

// Cputype returns the Cputype field from the record's packed storage.
func (s *Stackshot_cpu_architecture) Cputype() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetCputype updates the Cputype field in the record's packed storage.
func (s *Stackshot_cpu_architecture) SetCputype(v int32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// Cpusubtype returns the Cpusubtype field from the record's packed storage.
func (s *Stackshot_cpu_architecture) Cpusubtype() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetCpusubtype updates the Cpusubtype field in the record's packed storage.
func (s *Stackshot_cpu_architecture) SetCpusubtype(v int32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// Stackshot_cpu_times
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/stackshot_cpu_times
type Stackshot_cpu_times struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [16]byte
}

// User_usec returns the User_usec field from the record's packed storage.
func (s *Stackshot_cpu_times) User_usec() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[0:8]))
}

// SetUser_usec updates the User_usec field in the record's packed storage.
func (s *Stackshot_cpu_times) SetUser_usec(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[0:8], uint64(v))
}

// System_usec returns the System_usec field from the record's packed storage.
func (s *Stackshot_cpu_times) System_usec() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[8:16]))
}

// SetSystem_usec updates the System_usec field in the record's packed storage.
func (s *Stackshot_cpu_times) SetSystem_usec(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[8:16], uint64(v))
}

// Stackshot_cpu_times_v2
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/stackshot_cpu_times_v2
type Stackshot_cpu_times_v2 struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [24]byte
}

// User_usec returns the User_usec field from the record's packed storage.
func (s *Stackshot_cpu_times_v2) User_usec() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[0:8]))
}

// SetUser_usec updates the User_usec field in the record's packed storage.
func (s *Stackshot_cpu_times_v2) SetUser_usec(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[0:8], uint64(v))
}

// System_usec returns the System_usec field from the record's packed storage.
func (s *Stackshot_cpu_times_v2) System_usec() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[8:16]))
}

// SetSystem_usec updates the System_usec field in the record's packed storage.
func (s *Stackshot_cpu_times_v2) SetSystem_usec(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[8:16], uint64(v))
}

// Runnable_usec returns the Runnable_usec field from the record's packed storage.
func (s *Stackshot_cpu_times_v2) Runnable_usec() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[16:24]))
}

// SetRunnable_usec updates the Runnable_usec field in the record's packed storage.
func (s *Stackshot_cpu_times_v2) SetRunnable_usec(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[16:24], uint64(v))
}

// Stackshot_duration
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/stackshot_duration
type Stackshot_duration struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [16]byte
}

// Stackshot_duration returns the Stackshot_duration field from the record's packed storage.
func (s *Stackshot_duration) Stackshot_duration() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[0:8]))
}

// SetStackshot_duration updates the Stackshot_duration field in the record's packed storage.
func (s *Stackshot_duration) SetStackshot_duration(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[0:8], uint64(v))
}

// Stackshot_duration_outer returns the Stackshot_duration_outer field from the record's packed storage.
func (s *Stackshot_duration) Stackshot_duration_outer() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[8:16]))
}

// SetStackshot_duration_outer updates the Stackshot_duration_outer field in the record's packed storage.
func (s *Stackshot_duration) SetStackshot_duration_outer(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[8:16], uint64(v))
}

// Stackshot_duration_v2
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/stackshot_duration_v2
type Stackshot_duration_v2 struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [24]byte
}

// Stackshot_duration returns the Stackshot_duration field from the record's packed storage.
func (s *Stackshot_duration_v2) Stackshot_duration() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[0:8]))
}

// SetStackshot_duration updates the Stackshot_duration field in the record's packed storage.
func (s *Stackshot_duration_v2) SetStackshot_duration(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[0:8], uint64(v))
}

// Stackshot_duration_outer returns the Stackshot_duration_outer field from the record's packed storage.
func (s *Stackshot_duration_v2) Stackshot_duration_outer() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[8:16]))
}

// SetStackshot_duration_outer updates the Stackshot_duration_outer field in the record's packed storage.
func (s *Stackshot_duration_v2) SetStackshot_duration_outer(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[8:16], uint64(v))
}

// Stackshot_duration_prior returns the Stackshot_duration_prior field from the record's packed storage.
func (s *Stackshot_duration_v2) Stackshot_duration_prior() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[16:24]))
}

// SetStackshot_duration_prior updates the Stackshot_duration_prior field in the record's packed storage.
func (s *Stackshot_duration_v2) SetStackshot_duration_prior(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[16:24], uint64(v))
}

// Stackshot_fault_stats
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/stackshot_fault_stats
type Stackshot_fault_stats struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [21]byte
}

// Sfs_pages_faulted_in returns the Sfs_pages_faulted_in field from the record's packed storage.
func (s *Stackshot_fault_stats) Sfs_pages_faulted_in() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetSfs_pages_faulted_in updates the Sfs_pages_faulted_in field in the record's packed storage.
func (s *Stackshot_fault_stats) SetSfs_pages_faulted_in(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// Sfs_time_spent_faulting returns the Sfs_time_spent_faulting field from the record's packed storage.
func (s *Stackshot_fault_stats) Sfs_time_spent_faulting() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[4:12]))
}

// SetSfs_time_spent_faulting updates the Sfs_time_spent_faulting field in the record's packed storage.
func (s *Stackshot_fault_stats) SetSfs_time_spent_faulting(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[4:12], uint64(v))
}

// Sfs_system_max_fault_time returns the Sfs_system_max_fault_time field from the record's packed storage.
func (s *Stackshot_fault_stats) Sfs_system_max_fault_time() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[12:20]))
}

// SetSfs_system_max_fault_time updates the Sfs_system_max_fault_time field in the record's packed storage.
func (s *Stackshot_fault_stats) SetSfs_system_max_fault_time(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[12:20], uint64(v))
}

// Sfs_stopped_faulting returns the Sfs_stopped_faulting field from the record's packed storage.
func (s *Stackshot_fault_stats) Sfs_stopped_faulting() uint8 {
	return uint8(s.storage[20])
}

// SetSfs_stopped_faulting updates the Sfs_stopped_faulting field in the record's packed storage.
func (s *Stackshot_fault_stats) SetSfs_stopped_faulting(v uint8) {
	s.storage[20] = uint8(v)
}

// Stackshot_latency_collection
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/stackshot_latency_collection
type Stackshot_latency_collection struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [32]byte
}

// Latency_version returns the Latency_version field from the record's packed storage.
func (s *Stackshot_latency_collection) Latency_version() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[0:8]))
}

// SetLatency_version updates the Latency_version field in the record's packed storage.
func (s *Stackshot_latency_collection) SetLatency_version(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[0:8], uint64(v))
}

// Setup_latency returns the Setup_latency field from the record's packed storage.
func (s *Stackshot_latency_collection) Setup_latency() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[8:16]))
}

// SetSetup_latency updates the Setup_latency field in the record's packed storage.
func (s *Stackshot_latency_collection) SetSetup_latency(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[8:16], uint64(v))
}

// Total_task_iteration_latency returns the Total_task_iteration_latency field from the record's packed storage.
func (s *Stackshot_latency_collection) Total_task_iteration_latency() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[16:24]))
}

// SetTotal_task_iteration_latency updates the Total_task_iteration_latency field in the record's packed storage.
func (s *Stackshot_latency_collection) SetTotal_task_iteration_latency(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[16:24], uint64(v))
}

// Total_terminated_task_iteration_latency returns the Total_terminated_task_iteration_latency field from the record's packed storage.
func (s *Stackshot_latency_collection) Total_terminated_task_iteration_latency() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[24:32]))
}

// SetTotal_terminated_task_iteration_latency updates the Total_terminated_task_iteration_latency field in the record's packed storage.
func (s *Stackshot_latency_collection) SetTotal_terminated_task_iteration_latency(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[24:32], uint64(v))
}

// Stackshot_latency_collection_v2
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/stackshot_latency_collection_v2
type Stackshot_latency_collection_v2 struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [96]byte
}

// Latency_version returns the Latency_version field from the record's packed storage.
func (s *Stackshot_latency_collection_v2) Latency_version() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[0:8]))
}

// SetLatency_version updates the Latency_version field in the record's packed storage.
func (s *Stackshot_latency_collection_v2) SetLatency_version(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[0:8], uint64(v))
}

// Setup_latency_mt returns the Setup_latency_mt field from the record's packed storage.
func (s *Stackshot_latency_collection_v2) Setup_latency_mt() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[8:16]))
}

// SetSetup_latency_mt updates the Setup_latency_mt field in the record's packed storage.
func (s *Stackshot_latency_collection_v2) SetSetup_latency_mt(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[8:16], uint64(v))
}

// Total_task_iteration_latency_mt returns the Total_task_iteration_latency_mt field from the record's packed storage.
func (s *Stackshot_latency_collection_v2) Total_task_iteration_latency_mt() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[16:24]))
}

// SetTotal_task_iteration_latency_mt updates the Total_task_iteration_latency_mt field in the record's packed storage.
func (s *Stackshot_latency_collection_v2) SetTotal_task_iteration_latency_mt(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[16:24], uint64(v))
}

// Total_terminated_task_iteration_latency_mt returns the Total_terminated_task_iteration_latency_mt field from the record's packed storage.
func (s *Stackshot_latency_collection_v2) Total_terminated_task_iteration_latency_mt() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[24:32]))
}

// SetTotal_terminated_task_iteration_latency_mt updates the Total_terminated_task_iteration_latency_mt field in the record's packed storage.
func (s *Stackshot_latency_collection_v2) SetTotal_terminated_task_iteration_latency_mt(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[24:32], uint64(v))
}

// Task_queue_building_latency_mt returns the Task_queue_building_latency_mt field from the record's packed storage.
func (s *Stackshot_latency_collection_v2) Task_queue_building_latency_mt() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[32:40]))
}

// SetTask_queue_building_latency_mt updates the Task_queue_building_latency_mt field in the record's packed storage.
func (s *Stackshot_latency_collection_v2) SetTask_queue_building_latency_mt(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[32:40], uint64(v))
}

// Terminated_task_queue_building_latency_mt returns the Terminated_task_queue_building_latency_mt field from the record's packed storage.
func (s *Stackshot_latency_collection_v2) Terminated_task_queue_building_latency_mt() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[40:48]))
}

// SetTerminated_task_queue_building_latency_mt updates the Terminated_task_queue_building_latency_mt field in the record's packed storage.
func (s *Stackshot_latency_collection_v2) SetTerminated_task_queue_building_latency_mt(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[40:48], uint64(v))
}

// Cpu_wait_latency_mt returns the Cpu_wait_latency_mt field from the record's packed storage.
func (s *Stackshot_latency_collection_v2) Cpu_wait_latency_mt() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[48:56]))
}

// SetCpu_wait_latency_mt updates the Cpu_wait_latency_mt field in the record's packed storage.
func (s *Stackshot_latency_collection_v2) SetCpu_wait_latency_mt(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[48:56], uint64(v))
}

// Main_cpu_number returns the Main_cpu_number field from the record's packed storage.
func (s *Stackshot_latency_collection_v2) Main_cpu_number() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[56:60]))
}

// SetMain_cpu_number updates the Main_cpu_number field in the record's packed storage.
func (s *Stackshot_latency_collection_v2) SetMain_cpu_number(v int32) {
	binary.NativeEndian.PutUint32(s.storage[56:60], uint32(v))
}

// Calling_cpu_number returns the Calling_cpu_number field from the record's packed storage.
func (s *Stackshot_latency_collection_v2) Calling_cpu_number() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[60:64]))
}

// SetCalling_cpu_number updates the Calling_cpu_number field in the record's packed storage.
func (s *Stackshot_latency_collection_v2) SetCalling_cpu_number(v int32) {
	binary.NativeEndian.PutUint32(s.storage[60:64], uint32(v))
}

// Buffer_size returns the Buffer_size field from the record's packed storage.
func (s *Stackshot_latency_collection_v2) Buffer_size() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[64:72]))
}

// SetBuffer_size updates the Buffer_size field in the record's packed storage.
func (s *Stackshot_latency_collection_v2) SetBuffer_size(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[64:72], uint64(v))
}

// Buffer_used returns the Buffer_used field from the record's packed storage.
func (s *Stackshot_latency_collection_v2) Buffer_used() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[72:80]))
}

// SetBuffer_used updates the Buffer_used field in the record's packed storage.
func (s *Stackshot_latency_collection_v2) SetBuffer_used(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[72:80], uint64(v))
}

// Buffer_overhead returns the Buffer_overhead field from the record's packed storage.
func (s *Stackshot_latency_collection_v2) Buffer_overhead() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[80:88]))
}

// SetBuffer_overhead updates the Buffer_overhead field in the record's packed storage.
func (s *Stackshot_latency_collection_v2) SetBuffer_overhead(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[80:88], uint64(v))
}

// Buffer_count returns the Buffer_count field from the record's packed storage.
func (s *Stackshot_latency_collection_v2) Buffer_count() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[88:96]))
}

// SetBuffer_count updates the Buffer_count field in the record's packed storage.
func (s *Stackshot_latency_collection_v2) SetBuffer_count(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[88:96], uint64(v))
}

// Stackshot_latency_cpu
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/stackshot_latency_cpu
type Stackshot_latency_cpu struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [88]byte
}

// Cpu_number returns the Cpu_number field from the record's packed storage.
func (s *Stackshot_latency_cpu) Cpu_number() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetCpu_number updates the Cpu_number field in the record's packed storage.
func (s *Stackshot_latency_cpu) SetCpu_number(v int32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// Cluster_type returns the Cluster_type field from the record's packed storage.
func (s *Stackshot_latency_cpu) Cluster_type() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetCluster_type updates the Cluster_type field in the record's packed storage.
func (s *Stackshot_latency_cpu) SetCluster_type(v int32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// Init_latency_mt returns the Init_latency_mt field from the record's packed storage.
func (s *Stackshot_latency_cpu) Init_latency_mt() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[8:16]))
}

// SetInit_latency_mt updates the Init_latency_mt field in the record's packed storage.
func (s *Stackshot_latency_cpu) SetInit_latency_mt(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[8:16], uint64(v))
}

// Workqueue_latency_mt returns the Workqueue_latency_mt field from the record's packed storage.
func (s *Stackshot_latency_cpu) Workqueue_latency_mt() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[16:24]))
}

// SetWorkqueue_latency_mt updates the Workqueue_latency_mt field in the record's packed storage.
func (s *Stackshot_latency_cpu) SetWorkqueue_latency_mt(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[16:24], uint64(v))
}

// Total_latency_mt returns the Total_latency_mt field from the record's packed storage.
func (s *Stackshot_latency_cpu) Total_latency_mt() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[24:32]))
}

// SetTotal_latency_mt updates the Total_latency_mt field in the record's packed storage.
func (s *Stackshot_latency_cpu) SetTotal_latency_mt(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[24:32], uint64(v))
}

// Total_cycles returns the Total_cycles field from the record's packed storage.
func (s *Stackshot_latency_cpu) Total_cycles() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[32:40]))
}

// SetTotal_cycles updates the Total_cycles field in the record's packed storage.
func (s *Stackshot_latency_cpu) SetTotal_cycles(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[32:40], uint64(v))
}

// Total_instrs returns the Total_instrs field from the record's packed storage.
func (s *Stackshot_latency_cpu) Total_instrs() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[40:48]))
}

// SetTotal_instrs updates the Total_instrs field in the record's packed storage.
func (s *Stackshot_latency_cpu) SetTotal_instrs(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[40:48], uint64(v))
}

// Tasks_processed returns the Tasks_processed field from the record's packed storage.
func (s *Stackshot_latency_cpu) Tasks_processed() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[48:56]))
}

// SetTasks_processed updates the Tasks_processed field in the record's packed storage.
func (s *Stackshot_latency_cpu) SetTasks_processed(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[48:56], uint64(v))
}

// Threads_processed returns the Threads_processed field from the record's packed storage.
func (s *Stackshot_latency_cpu) Threads_processed() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[56:64]))
}

// SetThreads_processed updates the Threads_processed field in the record's packed storage.
func (s *Stackshot_latency_cpu) SetThreads_processed(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[56:64], uint64(v))
}

// Faulting_time_mt returns the Faulting_time_mt field from the record's packed storage.
func (s *Stackshot_latency_cpu) Faulting_time_mt() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[64:72]))
}

// SetFaulting_time_mt updates the Faulting_time_mt field in the record's packed storage.
func (s *Stackshot_latency_cpu) SetFaulting_time_mt(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[64:72], uint64(v))
}

// Total_buf returns the Total_buf field from the record's packed storage.
func (s *Stackshot_latency_cpu) Total_buf() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[72:80]))
}

// SetTotal_buf updates the Total_buf field in the record's packed storage.
func (s *Stackshot_latency_cpu) SetTotal_buf(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[72:80], uint64(v))
}

// Intercluster_buf_used returns the Intercluster_buf_used field from the record's packed storage.
func (s *Stackshot_latency_cpu) Intercluster_buf_used() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[80:88]))
}

// SetIntercluster_buf_used updates the Intercluster_buf_used field in the record's packed storage.
func (s *Stackshot_latency_cpu) SetIntercluster_buf_used(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[80:88], uint64(v))
}

// Stackshot_latency_task
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/stackshot_latency_task
type Stackshot_latency_task struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [80]byte
}

// Task_uniqueid returns the Task_uniqueid field from the record's packed storage.
func (s *Stackshot_latency_task) Task_uniqueid() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[0:8]))
}

// SetTask_uniqueid updates the Task_uniqueid field in the record's packed storage.
func (s *Stackshot_latency_task) SetTask_uniqueid(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[0:8], uint64(v))
}

// Setup_latency returns the Setup_latency field from the record's packed storage.
func (s *Stackshot_latency_task) Setup_latency() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[8:16]))
}

// SetSetup_latency updates the Setup_latency field in the record's packed storage.
func (s *Stackshot_latency_task) SetSetup_latency(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[8:16], uint64(v))
}

// Task_thread_count_loop_latency returns the Task_thread_count_loop_latency field from the record's packed storage.
func (s *Stackshot_latency_task) Task_thread_count_loop_latency() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[16:24]))
}

// SetTask_thread_count_loop_latency updates the Task_thread_count_loop_latency field in the record's packed storage.
func (s *Stackshot_latency_task) SetTask_thread_count_loop_latency(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[16:24], uint64(v))
}

// Task_thread_data_loop_latency returns the Task_thread_data_loop_latency field from the record's packed storage.
func (s *Stackshot_latency_task) Task_thread_data_loop_latency() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[24:32]))
}

// SetTask_thread_data_loop_latency updates the Task_thread_data_loop_latency field in the record's packed storage.
func (s *Stackshot_latency_task) SetTask_thread_data_loop_latency(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[24:32], uint64(v))
}

// Cur_tsnap_latency returns the Cur_tsnap_latency field from the record's packed storage.
func (s *Stackshot_latency_task) Cur_tsnap_latency() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[32:40]))
}

// SetCur_tsnap_latency updates the Cur_tsnap_latency field in the record's packed storage.
func (s *Stackshot_latency_task) SetCur_tsnap_latency(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[32:40], uint64(v))
}

// Pmap_latency returns the Pmap_latency field from the record's packed storage.
func (s *Stackshot_latency_task) Pmap_latency() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[40:48]))
}

// SetPmap_latency updates the Pmap_latency field in the record's packed storage.
func (s *Stackshot_latency_task) SetPmap_latency(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[40:48], uint64(v))
}

// Bsd_proc_ids_latency returns the Bsd_proc_ids_latency field from the record's packed storage.
func (s *Stackshot_latency_task) Bsd_proc_ids_latency() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[48:56]))
}

// SetBsd_proc_ids_latency updates the Bsd_proc_ids_latency field in the record's packed storage.
func (s *Stackshot_latency_task) SetBsd_proc_ids_latency(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[48:56], uint64(v))
}

// Misc_latency returns the Misc_latency field from the record's packed storage.
func (s *Stackshot_latency_task) Misc_latency() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[56:64]))
}

// SetMisc_latency updates the Misc_latency field in the record's packed storage.
func (s *Stackshot_latency_task) SetMisc_latency(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[56:64], uint64(v))
}

// Misc2_latency returns the Misc2_latency field from the record's packed storage.
func (s *Stackshot_latency_task) Misc2_latency() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[64:72]))
}

// SetMisc2_latency updates the Misc2_latency field in the record's packed storage.
func (s *Stackshot_latency_task) SetMisc2_latency(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[64:72], uint64(v))
}

// End_latency returns the End_latency field from the record's packed storage.
func (s *Stackshot_latency_task) End_latency() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[72:80]))
}

// SetEnd_latency updates the End_latency field in the record's packed storage.
func (s *Stackshot_latency_task) SetEnd_latency(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[72:80], uint64(v))
}

// Stackshot_latency_thread
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/stackshot_latency_thread
type Stackshot_latency_thread struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [80]byte
}

// Thread_id returns the Thread_id field from the record's packed storage.
func (s *Stackshot_latency_thread) Thread_id() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[0:8]))
}

// SetThread_id updates the Thread_id field in the record's packed storage.
func (s *Stackshot_latency_thread) SetThread_id(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[0:8], uint64(v))
}

// Cur_thsnap1_latency returns the Cur_thsnap1_latency field from the record's packed storage.
func (s *Stackshot_latency_thread) Cur_thsnap1_latency() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[8:16]))
}

// SetCur_thsnap1_latency updates the Cur_thsnap1_latency field in the record's packed storage.
func (s *Stackshot_latency_thread) SetCur_thsnap1_latency(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[8:16], uint64(v))
}

// Dispatch_serial_latency returns the Dispatch_serial_latency field from the record's packed storage.
func (s *Stackshot_latency_thread) Dispatch_serial_latency() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[16:24]))
}

// SetDispatch_serial_latency updates the Dispatch_serial_latency field in the record's packed storage.
func (s *Stackshot_latency_thread) SetDispatch_serial_latency(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[16:24], uint64(v))
}

// Dispatch_label_latency returns the Dispatch_label_latency field from the record's packed storage.
func (s *Stackshot_latency_thread) Dispatch_label_latency() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[24:32]))
}

// SetDispatch_label_latency updates the Dispatch_label_latency field in the record's packed storage.
func (s *Stackshot_latency_thread) SetDispatch_label_latency(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[24:32], uint64(v))
}

// Cur_thsnap2_latency returns the Cur_thsnap2_latency field from the record's packed storage.
func (s *Stackshot_latency_thread) Cur_thsnap2_latency() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[32:40]))
}

// SetCur_thsnap2_latency updates the Cur_thsnap2_latency field in the record's packed storage.
func (s *Stackshot_latency_thread) SetCur_thsnap2_latency(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[32:40], uint64(v))
}

// Thread_name_latency returns the Thread_name_latency field from the record's packed storage.
func (s *Stackshot_latency_thread) Thread_name_latency() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[40:48]))
}

// SetThread_name_latency updates the Thread_name_latency field in the record's packed storage.
func (s *Stackshot_latency_thread) SetThread_name_latency(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[40:48], uint64(v))
}

// Sur_times_latency returns the Sur_times_latency field from the record's packed storage.
func (s *Stackshot_latency_thread) Sur_times_latency() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[48:56]))
}

// SetSur_times_latency updates the Sur_times_latency field in the record's packed storage.
func (s *Stackshot_latency_thread) SetSur_times_latency(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[48:56], uint64(v))
}

// User_stack_latency returns the User_stack_latency field from the record's packed storage.
func (s *Stackshot_latency_thread) User_stack_latency() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[56:64]))
}

// SetUser_stack_latency updates the User_stack_latency field in the record's packed storage.
func (s *Stackshot_latency_thread) SetUser_stack_latency(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[56:64], uint64(v))
}

// Kernel_stack_latency returns the Kernel_stack_latency field from the record's packed storage.
func (s *Stackshot_latency_thread) Kernel_stack_latency() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[64:72]))
}

// SetKernel_stack_latency updates the Kernel_stack_latency field in the record's packed storage.
func (s *Stackshot_latency_thread) SetKernel_stack_latency(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[64:72], uint64(v))
}

// Misc_latency returns the Misc_latency field from the record's packed storage.
func (s *Stackshot_latency_thread) Misc_latency() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[72:80]))
}

// SetMisc_latency updates the Misc_latency field in the record's packed storage.
func (s *Stackshot_latency_thread) SetMisc_latency(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[72:80], uint64(v))
}

// Stackshot_suspension_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/stackshot_suspension_info
type Stackshot_suspension_info struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [32]byte
}

// Tss_last_start returns the Tss_last_start field from the record's packed storage.
func (s *Stackshot_suspension_info) Tss_last_start() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[0:8]))
}

// SetTss_last_start updates the Tss_last_start field in the record's packed storage.
func (s *Stackshot_suspension_info) SetTss_last_start(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[0:8], uint64(v))
}

// Tss_last_end returns the Tss_last_end field from the record's packed storage.
func (s *Stackshot_suspension_info) Tss_last_end() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[8:16]))
}

// SetTss_last_end updates the Tss_last_end field in the record's packed storage.
func (s *Stackshot_suspension_info) SetTss_last_end(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[8:16], uint64(v))
}

// Tss_count returns the Tss_count field from the record's packed storage.
func (s *Stackshot_suspension_info) Tss_count() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[16:24]))
}

// SetTss_count updates the Tss_count field in the record's packed storage.
func (s *Stackshot_suspension_info) SetTss_count(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[16:24], uint64(v))
}

// Tss_duration returns the Tss_duration field from the record's packed storage.
func (s *Stackshot_suspension_info) Tss_duration() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[24:32]))
}

// SetTss_duration updates the Tss_duration field in the record's packed storage.
func (s *Stackshot_suspension_info) SetTss_duration(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[24:32], uint64(v))
}

// Stackshot_suspension_source
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/stackshot_suspension_source
type Stackshot_suspension_source struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [85]byte
}

// Tss_time returns the Tss_time field from the record's packed storage.
func (s *Stackshot_suspension_source) Tss_time() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[0:8]))
}

// SetTss_time updates the Tss_time field in the record's packed storage.
func (s *Stackshot_suspension_source) SetTss_time(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[0:8], uint64(v))
}

// Tss_tid returns the Tss_tid field from the record's packed storage.
func (s *Stackshot_suspension_source) Tss_tid() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[8:16]))
}

// SetTss_tid updates the Tss_tid field in the record's packed storage.
func (s *Stackshot_suspension_source) SetTss_tid(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[8:16], uint64(v))
}

// Tss_pid returns the Tss_pid field from the record's packed storage.
func (s *Stackshot_suspension_source) Tss_pid() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[16:20]))
}

// SetTss_pid updates the Tss_pid field in the record's packed storage.
func (s *Stackshot_suspension_source) SetTss_pid(v int32) {
	binary.NativeEndian.PutUint32(s.storage[16:20], uint32(v))
}

// Tss_procname returns the Tss_procname field from the record's packed storage.
func (s *Stackshot_suspension_source) Tss_procname() [65]int8 {
	return *(*[65]int8)(unsafe.Pointer(&s.storage[20]))
}

// SetTss_procname updates the Tss_procname field in the record's packed storage.
func (s *Stackshot_suspension_source) SetTss_procname(v [65]int8) {
	*(*[65]int8)(unsafe.Pointer(&s.storage[20])) = v
}

// Stackshot_task_codesigning_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/stackshot_task_codesigning_info
type Stackshot_task_codesigning_info struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [12]byte
}

// Csflags returns the Csflags field from the record's packed storage.
func (s *Stackshot_task_codesigning_info) Csflags() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[0:8]))
}

// SetCsflags updates the Csflags field in the record's packed storage.
func (s *Stackshot_task_codesigning_info) SetCsflags(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[0:8], uint64(v))
}

// Cs_trust_level returns the Cs_trust_level field from the record's packed storage.
func (s *Stackshot_task_codesigning_info) Cs_trust_level() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetCs_trust_level updates the Cs_trust_level field in the record's packed storage.
func (s *Stackshot_task_codesigning_info) SetCs_trust_level(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[8:12], uint32(v))
}

// Stat
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/stat
type Stat struct {
	St_dev           int32
	St_mode          uint16
	St_nlink         uint16
	St_ino           uint64
	St_uid           uint32
	St_gid           uint32
	St_rdev          int32
	St_atimespec     Timespec
	St_mtimespec     Timespec
	St_ctimespec     Timespec
	St_birthtimespec Timespec
	St_size          int64
	St_blocks        int64
	St_blksize       int32
	St_flags         uint32
	St_gen           uint32
	St_lspare        int32
	St_qspare        [2]int64
}

// Sub_client_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sub_client_command
type Sub_client_command struct {
	Cmd     uint32
	Cmdsize uint32
	Client  [1]uint32
}

// Sub_framework_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sub_framework_command
type Sub_framework_command struct {
	Cmd      uint32
	Cmdsize  uint32
	Umbrella [1]uint32
}

// Sub_library_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sub_library_command
type Sub_library_command struct {
	Cmd         uint32
	Cmdsize     uint32
	Sub_library [1]uint32
}

// Sub_umbrella_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sub_umbrella_command
type Sub_umbrella_command struct {
	Cmd          uint32
	Cmdsize      uint32
	Sub_umbrella [1]uint32 // A data structure of type `lc_str`. Specifies the name of the umbrella framework of which this file is a member.

}

// Symseg_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/symseg_command
type Symseg_command struct {
	Cmd     uint32
	Cmdsize uint32
	Offset  uint32
	Size    uint32
}

// Symtab_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/symtab_command
type Symtab_command struct {
	Cmd     uint32
	Cmdsize uint32
	Symoff  uint32
	Nsyms   uint32
	Stroff  uint32
	Strsize uint32
}

// Sysctl_oid
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sysctl_oid
type Sysctl_oid struct {
	Oid_number  int32
	Oid_kind    int32
	Oid_arg1    unsafe.Pointer
	Oid_arg2    int32
	Oid_name    *byte
	Oid_handler unsafe.Pointer
	Oid_fmt     *byte
	Oid_descr   *byte
	Oid_version int32
	Oid_refcnt  int32
	Oid_link    unsafe.Pointer
	Oid_parent  *objc.ID
}

// Sysctl_oid_list
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sysctl_oid_list
type Sysctl_oid_list struct {
	Slh_first *Sysctl_oid
}

// Sysctl_req
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sysctl_req
type Sysctl_req struct {
	P       unsafe.Pointer
	Lock    int32
	Oldptr  User_addr_t
	Oldlen  uintptr
	Oldidx  uintptr
	Oldfunc unsafe.Pointer
	Newptr  User_addr_t
	Newlen  uintptr
	Newidx  uintptr
	Newfunc unsafe.Pointer
}

// Task_access_subsystem
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/task_access_subsystem
type Task_access_subsystem struct {
	Server   unsafe.Pointer
	Start    Mach_msg_id_t
	End      Mach_msg_id_t
	Maxsize  uint32
	Reserved Vm_address_t
	Routine  unsafe.Pointer
}

// Task_delta_snapshot_v2
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/task_delta_snapshot_v2
type Task_delta_snapshot_v2 struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [76]byte
}

// Tds_unique_pid returns the Tds_unique_pid field from the record's packed storage.
func (s *Task_delta_snapshot_v2) Tds_unique_pid() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[0:8]))
}

// SetTds_unique_pid updates the Tds_unique_pid field in the record's packed storage.
func (s *Task_delta_snapshot_v2) SetTds_unique_pid(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[0:8], uint64(v))
}

// Tds_ss_flags returns the Tds_ss_flags field from the record's packed storage.
func (s *Task_delta_snapshot_v2) Tds_ss_flags() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[8:16]))
}

// SetTds_ss_flags updates the Tds_ss_flags field in the record's packed storage.
func (s *Task_delta_snapshot_v2) SetTds_ss_flags(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[8:16], uint64(v))
}

// Tds_user_time_in_terminated_threads returns the Tds_user_time_in_terminated_threads field from the record's packed storage.
func (s *Task_delta_snapshot_v2) Tds_user_time_in_terminated_threads() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[16:24]))
}

// SetTds_user_time_in_terminated_threads updates the Tds_user_time_in_terminated_threads field in the record's packed storage.
func (s *Task_delta_snapshot_v2) SetTds_user_time_in_terminated_threads(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[16:24], uint64(v))
}

// Tds_system_time_in_terminated_threads returns the Tds_system_time_in_terminated_threads field from the record's packed storage.
func (s *Task_delta_snapshot_v2) Tds_system_time_in_terminated_threads() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[24:32]))
}

// SetTds_system_time_in_terminated_threads updates the Tds_system_time_in_terminated_threads field in the record's packed storage.
func (s *Task_delta_snapshot_v2) SetTds_system_time_in_terminated_threads(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[24:32], uint64(v))
}

// Tds_task_size returns the Tds_task_size field from the record's packed storage.
func (s *Task_delta_snapshot_v2) Tds_task_size() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[32:40]))
}

// SetTds_task_size updates the Tds_task_size field in the record's packed storage.
func (s *Task_delta_snapshot_v2) SetTds_task_size(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[32:40], uint64(v))
}

// Tds_max_resident_size returns the Tds_max_resident_size field from the record's packed storage.
func (s *Task_delta_snapshot_v2) Tds_max_resident_size() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[40:48]))
}

// SetTds_max_resident_size updates the Tds_max_resident_size field in the record's packed storage.
func (s *Task_delta_snapshot_v2) SetTds_max_resident_size(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[40:48], uint64(v))
}

// Tds_suspend_count returns the Tds_suspend_count field from the record's packed storage.
func (s *Task_delta_snapshot_v2) Tds_suspend_count() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[48:52]))
}

// SetTds_suspend_count updates the Tds_suspend_count field in the record's packed storage.
func (s *Task_delta_snapshot_v2) SetTds_suspend_count(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[48:52], uint32(v))
}

// Tds_faults returns the Tds_faults field from the record's packed storage.
func (s *Task_delta_snapshot_v2) Tds_faults() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[52:56]))
}

// SetTds_faults updates the Tds_faults field in the record's packed storage.
func (s *Task_delta_snapshot_v2) SetTds_faults(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[52:56], uint32(v))
}

// Tds_pageins returns the Tds_pageins field from the record's packed storage.
func (s *Task_delta_snapshot_v2) Tds_pageins() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[56:60]))
}

// SetTds_pageins updates the Tds_pageins field in the record's packed storage.
func (s *Task_delta_snapshot_v2) SetTds_pageins(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[56:60], uint32(v))
}

// Tds_cow_faults returns the Tds_cow_faults field from the record's packed storage.
func (s *Task_delta_snapshot_v2) Tds_cow_faults() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[60:64]))
}

// SetTds_cow_faults updates the Tds_cow_faults field in the record's packed storage.
func (s *Task_delta_snapshot_v2) SetTds_cow_faults(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[60:64], uint32(v))
}

// Tds_was_throttled returns the Tds_was_throttled field from the record's packed storage.
func (s *Task_delta_snapshot_v2) Tds_was_throttled() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[64:68]))
}

// SetTds_was_throttled updates the Tds_was_throttled field in the record's packed storage.
func (s *Task_delta_snapshot_v2) SetTds_was_throttled(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[64:68], uint32(v))
}

// Tds_did_throttle returns the Tds_did_throttle field from the record's packed storage.
func (s *Task_delta_snapshot_v2) Tds_did_throttle() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[68:72]))
}

// SetTds_did_throttle updates the Tds_did_throttle field in the record's packed storage.
func (s *Task_delta_snapshot_v2) SetTds_did_throttle(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[68:72], uint32(v))
}

// Tds_latency_qos returns the Tds_latency_qos field from the record's packed storage.
func (s *Task_delta_snapshot_v2) Tds_latency_qos() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[72:76]))
}

// SetTds_latency_qos updates the Tds_latency_qos field in the record's packed storage.
func (s *Task_delta_snapshot_v2) SetTds_latency_qos(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[72:76], uint32(v))
}

// Task_snapshot
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/task_snapshot
type Task_snapshot struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [293]byte
}

// Snapshot_magic returns the Snapshot_magic field from the record's packed storage.
func (s *Task_snapshot) Snapshot_magic() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetSnapshot_magic updates the Snapshot_magic field in the record's packed storage.
func (s *Task_snapshot) SetSnapshot_magic(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// Pid returns the Pid field from the record's packed storage.
func (s *Task_snapshot) Pid() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetPid updates the Pid field in the record's packed storage.
func (s *Task_snapshot) SetPid(v int32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// Uniqueid returns the Uniqueid field from the record's packed storage.
func (s *Task_snapshot) Uniqueid() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[8:16]))
}

// SetUniqueid updates the Uniqueid field in the record's packed storage.
func (s *Task_snapshot) SetUniqueid(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[8:16], uint64(v))
}

// User_time_in_terminated_threads returns the User_time_in_terminated_threads field from the record's packed storage.
func (s *Task_snapshot) User_time_in_terminated_threads() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[16:24]))
}

// SetUser_time_in_terminated_threads updates the User_time_in_terminated_threads field in the record's packed storage.
func (s *Task_snapshot) SetUser_time_in_terminated_threads(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[16:24], uint64(v))
}

// System_time_in_terminated_threads returns the System_time_in_terminated_threads field from the record's packed storage.
func (s *Task_snapshot) System_time_in_terminated_threads() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[24:32]))
}

// SetSystem_time_in_terminated_threads updates the System_time_in_terminated_threads field in the record's packed storage.
func (s *Task_snapshot) SetSystem_time_in_terminated_threads(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[24:32], uint64(v))
}

// Shared_cache_identifier returns the Shared_cache_identifier field from the record's packed storage.
func (s *Task_snapshot) Shared_cache_identifier() [16]uint8 {
	return *(*[16]uint8)(unsafe.Pointer(&s.storage[32]))
}

// SetShared_cache_identifier updates the Shared_cache_identifier field in the record's packed storage.
func (s *Task_snapshot) SetShared_cache_identifier(v [16]uint8) {
	*(*[16]uint8)(unsafe.Pointer(&s.storage[32])) = v
}

// Shared_cache_slide returns the Shared_cache_slide field from the record's packed storage.
func (s *Task_snapshot) Shared_cache_slide() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[48:56]))
}

// SetShared_cache_slide updates the Shared_cache_slide field in the record's packed storage.
func (s *Task_snapshot) SetShared_cache_slide(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[48:56], uint64(v))
}

// Nloadinfos returns the Nloadinfos field from the record's packed storage.
func (s *Task_snapshot) Nloadinfos() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[56:60]))
}

// SetNloadinfos updates the Nloadinfos field in the record's packed storage.
func (s *Task_snapshot) SetNloadinfos(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[56:60], uint32(v))
}

// Suspend_count returns the Suspend_count field from the record's packed storage.
func (s *Task_snapshot) Suspend_count() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[60:64]))
}

// SetSuspend_count updates the Suspend_count field in the record's packed storage.
func (s *Task_snapshot) SetSuspend_count(v int32) {
	binary.NativeEndian.PutUint32(s.storage[60:64], uint32(v))
}

// Task_size returns the Task_size field from the record's packed storage.
func (s *Task_snapshot) Task_size() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[64:68]))
}

// SetTask_size updates the Task_size field in the record's packed storage.
func (s *Task_snapshot) SetTask_size(v int32) {
	binary.NativeEndian.PutUint32(s.storage[64:68], uint32(v))
}

// Faults returns the Faults field from the record's packed storage.
func (s *Task_snapshot) Faults() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[68:72]))
}

// SetFaults updates the Faults field in the record's packed storage.
func (s *Task_snapshot) SetFaults(v int32) {
	binary.NativeEndian.PutUint32(s.storage[68:72], uint32(v))
}

// Pageins returns the Pageins field from the record's packed storage.
func (s *Task_snapshot) Pageins() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[72:76]))
}

// SetPageins updates the Pageins field in the record's packed storage.
func (s *Task_snapshot) SetPageins(v int32) {
	binary.NativeEndian.PutUint32(s.storage[72:76], uint32(v))
}

// Cow_faults returns the Cow_faults field from the record's packed storage.
func (s *Task_snapshot) Cow_faults() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[76:80]))
}

// SetCow_faults updates the Cow_faults field in the record's packed storage.
func (s *Task_snapshot) SetCow_faults(v int32) {
	binary.NativeEndian.PutUint32(s.storage[76:80], uint32(v))
}

// Ss_flags returns the Ss_flags field from the record's packed storage.
func (s *Task_snapshot) Ss_flags() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[80:84]))
}

// SetSs_flags updates the Ss_flags field in the record's packed storage.
func (s *Task_snapshot) SetSs_flags(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[80:84], uint32(v))
}

// P_start_sec returns the P_start_sec field from the record's packed storage.
func (s *Task_snapshot) P_start_sec() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[84:92]))
}

// SetP_start_sec updates the P_start_sec field in the record's packed storage.
func (s *Task_snapshot) SetP_start_sec(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[84:92], uint64(v))
}

// P_start_usec returns the P_start_usec field from the record's packed storage.
func (s *Task_snapshot) P_start_usec() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[92:100]))
}

// SetP_start_usec updates the P_start_usec field in the record's packed storage.
func (s *Task_snapshot) SetP_start_usec(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[92:100], uint64(v))
}

// P_comm returns the P_comm field from the record's packed storage.
func (s *Task_snapshot) P_comm() [17]int8 {
	return *(*[17]int8)(unsafe.Pointer(&s.storage[100]))
}

// SetP_comm updates the P_comm field in the record's packed storage.
func (s *Task_snapshot) SetP_comm(v [17]int8) {
	*(*[17]int8)(unsafe.Pointer(&s.storage[100])) = v
}

// Was_throttled returns the Was_throttled field from the record's packed storage.
func (s *Task_snapshot) Was_throttled() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[117:121]))
}

// SetWas_throttled updates the Was_throttled field in the record's packed storage.
func (s *Task_snapshot) SetWas_throttled(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[117:121], uint32(v))
}

// Did_throttle returns the Did_throttle field from the record's packed storage.
func (s *Task_snapshot) Did_throttle() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[121:125]))
}

// SetDid_throttle updates the Did_throttle field in the record's packed storage.
func (s *Task_snapshot) SetDid_throttle(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[121:125], uint32(v))
}

// Latency_qos returns the Latency_qos field from the record's packed storage.
func (s *Task_snapshot) Latency_qos() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[125:129]))
}

// SetLatency_qos updates the Latency_qos field in the record's packed storage.
func (s *Task_snapshot) SetLatency_qos(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[125:129], uint32(v))
}

// Disk_reads_count returns the Disk_reads_count field from the record's packed storage.
func (s *Task_snapshot) Disk_reads_count() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[129:137]))
}

// SetDisk_reads_count updates the Disk_reads_count field in the record's packed storage.
func (s *Task_snapshot) SetDisk_reads_count(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[129:137], uint64(v))
}

// Disk_reads_size returns the Disk_reads_size field from the record's packed storage.
func (s *Task_snapshot) Disk_reads_size() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[137:145]))
}

// SetDisk_reads_size updates the Disk_reads_size field in the record's packed storage.
func (s *Task_snapshot) SetDisk_reads_size(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[137:145], uint64(v))
}

// Disk_writes_count returns the Disk_writes_count field from the record's packed storage.
func (s *Task_snapshot) Disk_writes_count() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[145:153]))
}

// SetDisk_writes_count updates the Disk_writes_count field in the record's packed storage.
func (s *Task_snapshot) SetDisk_writes_count(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[145:153], uint64(v))
}

// Disk_writes_size returns the Disk_writes_size field from the record's packed storage.
func (s *Task_snapshot) Disk_writes_size() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[153:161]))
}

// SetDisk_writes_size updates the Disk_writes_size field in the record's packed storage.
func (s *Task_snapshot) SetDisk_writes_size(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[153:161], uint64(v))
}

// Io_priority_count returns the Io_priority_count field from the record's packed storage.
func (s *Task_snapshot) Io_priority_count() [4]uint64 {
	return *(*[4]uint64)(unsafe.Pointer(&s.storage[161]))
}

// SetIo_priority_count updates the Io_priority_count field in the record's packed storage.
func (s *Task_snapshot) SetIo_priority_count(v [4]uint64) {
	*(*[4]uint64)(unsafe.Pointer(&s.storage[161])) = v
}

// Io_priority_size returns the Io_priority_size field from the record's packed storage.
func (s *Task_snapshot) Io_priority_size() [4]uint64 {
	return *(*[4]uint64)(unsafe.Pointer(&s.storage[193]))
}

// SetIo_priority_size updates the Io_priority_size field in the record's packed storage.
func (s *Task_snapshot) SetIo_priority_size(v [4]uint64) {
	*(*[4]uint64)(unsafe.Pointer(&s.storage[193])) = v
}

// Paging_count returns the Paging_count field from the record's packed storage.
func (s *Task_snapshot) Paging_count() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[225:233]))
}

// SetPaging_count updates the Paging_count field in the record's packed storage.
func (s *Task_snapshot) SetPaging_count(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[225:233], uint64(v))
}

// Paging_size returns the Paging_size field from the record's packed storage.
func (s *Task_snapshot) Paging_size() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[233:241]))
}

// SetPaging_size updates the Paging_size field in the record's packed storage.
func (s *Task_snapshot) SetPaging_size(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[233:241], uint64(v))
}

// Non_paging_count returns the Non_paging_count field from the record's packed storage.
func (s *Task_snapshot) Non_paging_count() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[241:249]))
}

// SetNon_paging_count updates the Non_paging_count field in the record's packed storage.
func (s *Task_snapshot) SetNon_paging_count(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[241:249], uint64(v))
}

// Non_paging_size returns the Non_paging_size field from the record's packed storage.
func (s *Task_snapshot) Non_paging_size() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[249:257]))
}

// SetNon_paging_size updates the Non_paging_size field in the record's packed storage.
func (s *Task_snapshot) SetNon_paging_size(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[249:257], uint64(v))
}

// Data_count returns the Data_count field from the record's packed storage.
func (s *Task_snapshot) Data_count() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[257:265]))
}

// SetData_count updates the Data_count field in the record's packed storage.
func (s *Task_snapshot) SetData_count(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[257:265], uint64(v))
}

// Data_size returns the Data_size field from the record's packed storage.
func (s *Task_snapshot) Data_size() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[265:273]))
}

// SetData_size updates the Data_size field in the record's packed storage.
func (s *Task_snapshot) SetData_size(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[265:273], uint64(v))
}

// Metadata_count returns the Metadata_count field from the record's packed storage.
func (s *Task_snapshot) Metadata_count() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[273:281]))
}

// SetMetadata_count updates the Metadata_count field in the record's packed storage.
func (s *Task_snapshot) SetMetadata_count(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[273:281], uint64(v))
}

// Metadata_size returns the Metadata_size field from the record's packed storage.
func (s *Task_snapshot) Metadata_size() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[281:289]))
}

// SetMetadata_size updates the Metadata_size field in the record's packed storage.
func (s *Task_snapshot) SetMetadata_size(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[281:289], uint64(v))
}

// Donating_pid_count returns the Donating_pid_count field from the record's packed storage.
func (s *Task_snapshot) Donating_pid_count() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[289:293]))
}

// SetDonating_pid_count updates the Donating_pid_count field in the record's packed storage.
func (s *Task_snapshot) SetDonating_pid_count(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[289:293], uint32(v))
}

// Task_snapshot_v2
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/task_snapshot_v2
type Task_snapshot_v2 struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [120]byte
}

// Ts_unique_pid returns the Ts_unique_pid field from the record's packed storage.
func (s *Task_snapshot_v2) Ts_unique_pid() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[0:8]))
}

// SetTs_unique_pid updates the Ts_unique_pid field in the record's packed storage.
func (s *Task_snapshot_v2) SetTs_unique_pid(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[0:8], uint64(v))
}

// Ts_ss_flags returns the Ts_ss_flags field from the record's packed storage.
func (s *Task_snapshot_v2) Ts_ss_flags() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[8:16]))
}

// SetTs_ss_flags updates the Ts_ss_flags field in the record's packed storage.
func (s *Task_snapshot_v2) SetTs_ss_flags(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[8:16], uint64(v))
}

// Ts_user_time_in_terminated_threads returns the Ts_user_time_in_terminated_threads field from the record's packed storage.
func (s *Task_snapshot_v2) Ts_user_time_in_terminated_threads() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[16:24]))
}

// SetTs_user_time_in_terminated_threads updates the Ts_user_time_in_terminated_threads field in the record's packed storage.
func (s *Task_snapshot_v2) SetTs_user_time_in_terminated_threads(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[16:24], uint64(v))
}

// Ts_system_time_in_terminated_threads returns the Ts_system_time_in_terminated_threads field from the record's packed storage.
func (s *Task_snapshot_v2) Ts_system_time_in_terminated_threads() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[24:32]))
}

// SetTs_system_time_in_terminated_threads updates the Ts_system_time_in_terminated_threads field in the record's packed storage.
func (s *Task_snapshot_v2) SetTs_system_time_in_terminated_threads(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[24:32], uint64(v))
}

// Ts_p_start_sec returns the Ts_p_start_sec field from the record's packed storage.
func (s *Task_snapshot_v2) Ts_p_start_sec() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[32:40]))
}

// SetTs_p_start_sec updates the Ts_p_start_sec field in the record's packed storage.
func (s *Task_snapshot_v2) SetTs_p_start_sec(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[32:40], uint64(v))
}

// Ts_task_size returns the Ts_task_size field from the record's packed storage.
func (s *Task_snapshot_v2) Ts_task_size() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[40:48]))
}

// SetTs_task_size updates the Ts_task_size field in the record's packed storage.
func (s *Task_snapshot_v2) SetTs_task_size(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[40:48], uint64(v))
}

// Ts_max_resident_size returns the Ts_max_resident_size field from the record's packed storage.
func (s *Task_snapshot_v2) Ts_max_resident_size() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[48:56]))
}

// SetTs_max_resident_size updates the Ts_max_resident_size field in the record's packed storage.
func (s *Task_snapshot_v2) SetTs_max_resident_size(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[48:56], uint64(v))
}

// Ts_suspend_count returns the Ts_suspend_count field from the record's packed storage.
func (s *Task_snapshot_v2) Ts_suspend_count() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[56:60]))
}

// SetTs_suspend_count updates the Ts_suspend_count field in the record's packed storage.
func (s *Task_snapshot_v2) SetTs_suspend_count(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[56:60], uint32(v))
}

// Ts_faults returns the Ts_faults field from the record's packed storage.
func (s *Task_snapshot_v2) Ts_faults() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[60:64]))
}

// SetTs_faults updates the Ts_faults field in the record's packed storage.
func (s *Task_snapshot_v2) SetTs_faults(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[60:64], uint32(v))
}

// Ts_pageins returns the Ts_pageins field from the record's packed storage.
func (s *Task_snapshot_v2) Ts_pageins() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[64:68]))
}

// SetTs_pageins updates the Ts_pageins field in the record's packed storage.
func (s *Task_snapshot_v2) SetTs_pageins(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[64:68], uint32(v))
}

// Ts_cow_faults returns the Ts_cow_faults field from the record's packed storage.
func (s *Task_snapshot_v2) Ts_cow_faults() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[68:72]))
}

// SetTs_cow_faults updates the Ts_cow_faults field in the record's packed storage.
func (s *Task_snapshot_v2) SetTs_cow_faults(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[68:72], uint32(v))
}

// Ts_was_throttled returns the Ts_was_throttled field from the record's packed storage.
func (s *Task_snapshot_v2) Ts_was_throttled() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[72:76]))
}

// SetTs_was_throttled updates the Ts_was_throttled field in the record's packed storage.
func (s *Task_snapshot_v2) SetTs_was_throttled(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[72:76], uint32(v))
}

// Ts_did_throttle returns the Ts_did_throttle field from the record's packed storage.
func (s *Task_snapshot_v2) Ts_did_throttle() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[76:80]))
}

// SetTs_did_throttle updates the Ts_did_throttle field in the record's packed storage.
func (s *Task_snapshot_v2) SetTs_did_throttle(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[76:80], uint32(v))
}

// Ts_latency_qos returns the Ts_latency_qos field from the record's packed storage.
func (s *Task_snapshot_v2) Ts_latency_qos() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[80:84]))
}

// SetTs_latency_qos updates the Ts_latency_qos field in the record's packed storage.
func (s *Task_snapshot_v2) SetTs_latency_qos(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[80:84], uint32(v))
}

// Ts_pid returns the Ts_pid field from the record's packed storage.
func (s *Task_snapshot_v2) Ts_pid() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[84:88]))
}

// SetTs_pid updates the Ts_pid field in the record's packed storage.
func (s *Task_snapshot_v2) SetTs_pid(v int32) {
	binary.NativeEndian.PutUint32(s.storage[84:88], uint32(v))
}

// Ts_p_comm returns the Ts_p_comm field from the record's packed storage.
func (s *Task_snapshot_v2) Ts_p_comm() [32]int8 {
	return *(*[32]int8)(unsafe.Pointer(&s.storage[88]))
}

// SetTs_p_comm updates the Ts_p_comm field in the record's packed storage.
func (s *Task_snapshot_v2) SetTs_p_comm(v [32]int8) {
	*(*[32]int8)(unsafe.Pointer(&s.storage[88])) = v
}

// Tchars
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/tchars
type Tchars struct {
	T_intrc  int8
	T_quitc  int8
	T_startc int8
	T_stopc  int8
	T_eofc   int8
	T_brkc   int8
}

// Tcp_conn_status
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/tcp_conn_status
type Tcp_conn_status struct {
	bitfield0 uint8
}

// Probe_activated returns the Probe_activated bitfield.
func (s *Tcp_conn_status) Probe_activated() uint8 {
	return (s.bitfield0 >> 0) & ((1 << 1) - 1)
}

// SetProbe_activated updates the Probe_activated bitfield.
func (s *Tcp_conn_status) SetProbe_activated(v uint8) {
	const mask uint8 = (1 << 1) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 0)) | ((v & mask) << 0)
}

// Write_probe_failed returns the Write_probe_failed bitfield.
func (s *Tcp_conn_status) Write_probe_failed() uint8 {
	return (s.bitfield0 >> 1) & ((1 << 1) - 1)
}

// SetWrite_probe_failed updates the Write_probe_failed bitfield.
func (s *Tcp_conn_status) SetWrite_probe_failed(v uint8) {
	const mask uint8 = (1 << 1) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 1)) | ((v & mask) << 1)
}

// Read_probe_failed returns the Read_probe_failed bitfield.
func (s *Tcp_conn_status) Read_probe_failed() uint8 {
	return (s.bitfield0 >> 2) & ((1 << 1) - 1)
}

// SetRead_probe_failed updates the Read_probe_failed bitfield.
func (s *Tcp_conn_status) SetRead_probe_failed(v uint8) {
	const mask uint8 = (1 << 1) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 2)) | ((v & mask) << 2)
}

// Conn_probe_failed returns the Conn_probe_failed bitfield.
func (s *Tcp_conn_status) Conn_probe_failed() uint8 {
	return (s.bitfield0 >> 3) & ((1 << 1) - 1)
}

// SetConn_probe_failed updates the Conn_probe_failed bitfield.
func (s *Tcp_conn_status) SetConn_probe_failed(v uint8) {
	const mask uint8 = (1 << 1) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 3)) | ((v & mask) << 3)
}

// Tcp_connection_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/tcp_connection_info
type Tcp_connection_info struct {
	Tcpi_state               U_int8_t
	Tcpi_snd_wscale          U_int8_t
	Tcpi_rcv_wscale          U_int8_t
	__pad1                   U_int8_t
	Tcpi_options             U_int32_t
	Tcpi_flags               U_int32_t
	Tcpi_rto                 U_int32_t
	Tcpi_maxseg              U_int32_t
	Tcpi_snd_ssthresh        U_int32_t
	Tcpi_snd_cwnd            U_int32_t
	Tcpi_snd_wnd             U_int32_t
	Tcpi_snd_sbbytes         U_int32_t
	Tcpi_rcv_wnd             U_int32_t
	Tcpi_rttcur              U_int32_t
	Tcpi_srtt                U_int32_t
	Tcpi_rttvar              U_int32_t
	bitfield16               uint32
	Tcpi_txpackets           U_int64_t
	Tcpi_txbytes             U_int64_t
	Tcpi_txretransmitbytes   U_int64_t
	Tcpi_rxpackets           U_int64_t
	Tcpi_rxbytes             U_int64_t
	Tcpi_rxoutoforderbytes   U_int64_t
	Tcpi_txretransmitpackets U_int64_t
}

// Tcpi_tfo_cookie_req returns the Tcpi_tfo_cookie_req bitfield.
func (s *Tcp_connection_info) Tcpi_tfo_cookie_req() uint32 {
	return (s.bitfield16 >> 0) & ((1 << 1) - 1)
}

// SetTcpi_tfo_cookie_req updates the Tcpi_tfo_cookie_req bitfield.
func (s *Tcp_connection_info) SetTcpi_tfo_cookie_req(v uint32) {
	const mask uint32 = (1 << 1) - 1
	s.bitfield16 = (s.bitfield16 &^ (mask << 0)) | ((v & mask) << 0)
}

// Tcpi_tfo_cookie_rcv returns the Tcpi_tfo_cookie_rcv bitfield.
func (s *Tcp_connection_info) Tcpi_tfo_cookie_rcv() uint32 {
	return (s.bitfield16 >> 1) & ((1 << 1) - 1)
}

// SetTcpi_tfo_cookie_rcv updates the Tcpi_tfo_cookie_rcv bitfield.
func (s *Tcp_connection_info) SetTcpi_tfo_cookie_rcv(v uint32) {
	const mask uint32 = (1 << 1) - 1
	s.bitfield16 = (s.bitfield16 &^ (mask << 1)) | ((v & mask) << 1)
}

// Tcpi_tfo_syn_loss returns the Tcpi_tfo_syn_loss bitfield.
func (s *Tcp_connection_info) Tcpi_tfo_syn_loss() uint32 {
	return (s.bitfield16 >> 2) & ((1 << 1) - 1)
}

// SetTcpi_tfo_syn_loss updates the Tcpi_tfo_syn_loss bitfield.
func (s *Tcp_connection_info) SetTcpi_tfo_syn_loss(v uint32) {
	const mask uint32 = (1 << 1) - 1
	s.bitfield16 = (s.bitfield16 &^ (mask << 2)) | ((v & mask) << 2)
}

// Tcpi_tfo_syn_data_sent returns the Tcpi_tfo_syn_data_sent bitfield.
func (s *Tcp_connection_info) Tcpi_tfo_syn_data_sent() uint32 {
	return (s.bitfield16 >> 3) & ((1 << 1) - 1)
}

// SetTcpi_tfo_syn_data_sent updates the Tcpi_tfo_syn_data_sent bitfield.
func (s *Tcp_connection_info) SetTcpi_tfo_syn_data_sent(v uint32) {
	const mask uint32 = (1 << 1) - 1
	s.bitfield16 = (s.bitfield16 &^ (mask << 3)) | ((v & mask) << 3)
}

// Tcpi_tfo_syn_data_acked returns the Tcpi_tfo_syn_data_acked bitfield.
func (s *Tcp_connection_info) Tcpi_tfo_syn_data_acked() uint32 {
	return (s.bitfield16 >> 4) & ((1 << 1) - 1)
}

// SetTcpi_tfo_syn_data_acked updates the Tcpi_tfo_syn_data_acked bitfield.
func (s *Tcp_connection_info) SetTcpi_tfo_syn_data_acked(v uint32) {
	const mask uint32 = (1 << 1) - 1
	s.bitfield16 = (s.bitfield16 &^ (mask << 4)) | ((v & mask) << 4)
}

// Tcpi_tfo_syn_data_rcv returns the Tcpi_tfo_syn_data_rcv bitfield.
func (s *Tcp_connection_info) Tcpi_tfo_syn_data_rcv() uint32 {
	return (s.bitfield16 >> 5) & ((1 << 1) - 1)
}

// SetTcpi_tfo_syn_data_rcv updates the Tcpi_tfo_syn_data_rcv bitfield.
func (s *Tcp_connection_info) SetTcpi_tfo_syn_data_rcv(v uint32) {
	const mask uint32 = (1 << 1) - 1
	s.bitfield16 = (s.bitfield16 &^ (mask << 5)) | ((v & mask) << 5)
}

// Tcpi_tfo_cookie_req_rcv returns the Tcpi_tfo_cookie_req_rcv bitfield.
func (s *Tcp_connection_info) Tcpi_tfo_cookie_req_rcv() uint32 {
	return (s.bitfield16 >> 6) & ((1 << 1) - 1)
}

// SetTcpi_tfo_cookie_req_rcv updates the Tcpi_tfo_cookie_req_rcv bitfield.
func (s *Tcp_connection_info) SetTcpi_tfo_cookie_req_rcv(v uint32) {
	const mask uint32 = (1 << 1) - 1
	s.bitfield16 = (s.bitfield16 &^ (mask << 6)) | ((v & mask) << 6)
}

// Tcpi_tfo_cookie_sent returns the Tcpi_tfo_cookie_sent bitfield.
func (s *Tcp_connection_info) Tcpi_tfo_cookie_sent() uint32 {
	return (s.bitfield16 >> 7) & ((1 << 1) - 1)
}

// SetTcpi_tfo_cookie_sent updates the Tcpi_tfo_cookie_sent bitfield.
func (s *Tcp_connection_info) SetTcpi_tfo_cookie_sent(v uint32) {
	const mask uint32 = (1 << 1) - 1
	s.bitfield16 = (s.bitfield16 &^ (mask << 7)) | ((v & mask) << 7)
}

// Tcpi_tfo_cookie_invalid returns the Tcpi_tfo_cookie_invalid bitfield.
func (s *Tcp_connection_info) Tcpi_tfo_cookie_invalid() uint32 {
	return (s.bitfield16 >> 8) & ((1 << 1) - 1)
}

// SetTcpi_tfo_cookie_invalid updates the Tcpi_tfo_cookie_invalid bitfield.
func (s *Tcp_connection_info) SetTcpi_tfo_cookie_invalid(v uint32) {
	const mask uint32 = (1 << 1) - 1
	s.bitfield16 = (s.bitfield16 &^ (mask << 8)) | ((v & mask) << 8)
}

// Tcpi_tfo_cookie_wrong returns the Tcpi_tfo_cookie_wrong bitfield.
func (s *Tcp_connection_info) Tcpi_tfo_cookie_wrong() uint32 {
	return (s.bitfield16 >> 9) & ((1 << 1) - 1)
}

// SetTcpi_tfo_cookie_wrong updates the Tcpi_tfo_cookie_wrong bitfield.
func (s *Tcp_connection_info) SetTcpi_tfo_cookie_wrong(v uint32) {
	const mask uint32 = (1 << 1) - 1
	s.bitfield16 = (s.bitfield16 &^ (mask << 9)) | ((v & mask) << 9)
}

// Tcpi_tfo_no_cookie_rcv returns the Tcpi_tfo_no_cookie_rcv bitfield.
func (s *Tcp_connection_info) Tcpi_tfo_no_cookie_rcv() uint32 {
	return (s.bitfield16 >> 10) & ((1 << 1) - 1)
}

// SetTcpi_tfo_no_cookie_rcv updates the Tcpi_tfo_no_cookie_rcv bitfield.
func (s *Tcp_connection_info) SetTcpi_tfo_no_cookie_rcv(v uint32) {
	const mask uint32 = (1 << 1) - 1
	s.bitfield16 = (s.bitfield16 &^ (mask << 10)) | ((v & mask) << 10)
}

// Tcpi_tfo_heuristics_disable returns the Tcpi_tfo_heuristics_disable bitfield.
func (s *Tcp_connection_info) Tcpi_tfo_heuristics_disable() uint32 {
	return (s.bitfield16 >> 11) & ((1 << 1) - 1)
}

// SetTcpi_tfo_heuristics_disable updates the Tcpi_tfo_heuristics_disable bitfield.
func (s *Tcp_connection_info) SetTcpi_tfo_heuristics_disable(v uint32) {
	const mask uint32 = (1 << 1) - 1
	s.bitfield16 = (s.bitfield16 &^ (mask << 11)) | ((v & mask) << 11)
}

// Tcpi_tfo_send_blackhole returns the Tcpi_tfo_send_blackhole bitfield.
func (s *Tcp_connection_info) Tcpi_tfo_send_blackhole() uint32 {
	return (s.bitfield16 >> 12) & ((1 << 1) - 1)
}

// SetTcpi_tfo_send_blackhole updates the Tcpi_tfo_send_blackhole bitfield.
func (s *Tcp_connection_info) SetTcpi_tfo_send_blackhole(v uint32) {
	const mask uint32 = (1 << 1) - 1
	s.bitfield16 = (s.bitfield16 &^ (mask << 12)) | ((v & mask) << 12)
}

// Tcpi_tfo_recv_blackhole returns the Tcpi_tfo_recv_blackhole bitfield.
func (s *Tcp_connection_info) Tcpi_tfo_recv_blackhole() uint32 {
	return (s.bitfield16 >> 13) & ((1 << 1) - 1)
}

// SetTcpi_tfo_recv_blackhole updates the Tcpi_tfo_recv_blackhole bitfield.
func (s *Tcp_connection_info) SetTcpi_tfo_recv_blackhole(v uint32) {
	const mask uint32 = (1 << 1) - 1
	s.bitfield16 = (s.bitfield16 &^ (mask << 13)) | ((v & mask) << 13)
}

// Tcpi_tfo_onebyte_proxy returns the Tcpi_tfo_onebyte_proxy bitfield.
func (s *Tcp_connection_info) Tcpi_tfo_onebyte_proxy() uint32 {
	return (s.bitfield16 >> 14) & ((1 << 1) - 1)
}

// SetTcpi_tfo_onebyte_proxy updates the Tcpi_tfo_onebyte_proxy bitfield.
func (s *Tcp_connection_info) SetTcpi_tfo_onebyte_proxy(v uint32) {
	const mask uint32 = (1 << 1) - 1
	s.bitfield16 = (s.bitfield16 &^ (mask << 14)) | ((v & mask) << 14)
}

// __pad2 returns the __pad2 bitfield.
func (s *Tcp_connection_info) __pad2() uint32 {
	return (s.bitfield16 >> 15) & ((1 << 17) - 1)
}

// Set__pad2 updates the __pad2 bitfield.
func (s *Tcp_connection_info) Set__pad2(v uint32) {
	const mask uint32 = (1 << 17) - 1
	s.bitfield16 = (s.bitfield16 &^ (mask << 15)) | ((v & mask) << 15)
}

// Tcp_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/tcp_info
type Tcp_info struct {
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
	storage [424]byte
}

// Tcpi_state returns the Tcpi_state field from the record's packed storage.
func (s *Tcp_info) Tcpi_state() U_int8_t {
	return *(*U_int8_t)(unsafe.Pointer(&s.storage[0]))
}

// SetTcpi_state updates the Tcpi_state field in the record's packed storage.
func (s *Tcp_info) SetTcpi_state(v U_int8_t) {
	*(*U_int8_t)(unsafe.Pointer(&s.storage[0])) = v
}

// Tcpi_options returns the Tcpi_options field from the record's packed storage.
func (s *Tcp_info) Tcpi_options() U_int8_t {
	return *(*U_int8_t)(unsafe.Pointer(&s.storage[1]))
}

// SetTcpi_options updates the Tcpi_options field in the record's packed storage.
func (s *Tcp_info) SetTcpi_options(v U_int8_t) {
	*(*U_int8_t)(unsafe.Pointer(&s.storage[1])) = v
}

// Tcpi_snd_wscale returns the Tcpi_snd_wscale field from the record's packed storage.
func (s *Tcp_info) Tcpi_snd_wscale() U_int8_t {
	return *(*U_int8_t)(unsafe.Pointer(&s.storage[2]))
}

// SetTcpi_snd_wscale updates the Tcpi_snd_wscale field in the record's packed storage.
func (s *Tcp_info) SetTcpi_snd_wscale(v U_int8_t) {
	*(*U_int8_t)(unsafe.Pointer(&s.storage[2])) = v
}

// Tcpi_rcv_wscale returns the Tcpi_rcv_wscale field from the record's packed storage.
func (s *Tcp_info) Tcpi_rcv_wscale() U_int8_t {
	return *(*U_int8_t)(unsafe.Pointer(&s.storage[3]))
}

// SetTcpi_rcv_wscale updates the Tcpi_rcv_wscale field in the record's packed storage.
func (s *Tcp_info) SetTcpi_rcv_wscale(v U_int8_t) {
	*(*U_int8_t)(unsafe.Pointer(&s.storage[3])) = v
}

// Tcpi_flags returns the Tcpi_flags field from the record's packed storage.
func (s *Tcp_info) Tcpi_flags() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetTcpi_flags updates the Tcpi_flags field in the record's packed storage.
func (s *Tcp_info) SetTcpi_flags(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// Tcpi_rto returns the Tcpi_rto field from the record's packed storage.
func (s *Tcp_info) Tcpi_rto() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetTcpi_rto updates the Tcpi_rto field in the record's packed storage.
func (s *Tcp_info) SetTcpi_rto(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[8:12], uint32(v))
}

// Tcpi_snd_mss returns the Tcpi_snd_mss field from the record's packed storage.
func (s *Tcp_info) Tcpi_snd_mss() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[12:16]))
}

// SetTcpi_snd_mss updates the Tcpi_snd_mss field in the record's packed storage.
func (s *Tcp_info) SetTcpi_snd_mss(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[12:16], uint32(v))
}

// Tcpi_rcv_mss returns the Tcpi_rcv_mss field from the record's packed storage.
func (s *Tcp_info) Tcpi_rcv_mss() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[16:20]))
}

// SetTcpi_rcv_mss updates the Tcpi_rcv_mss field in the record's packed storage.
func (s *Tcp_info) SetTcpi_rcv_mss(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[16:20], uint32(v))
}

// Tcpi_rttcur returns the Tcpi_rttcur field from the record's packed storage.
func (s *Tcp_info) Tcpi_rttcur() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[20:24]))
}

// SetTcpi_rttcur updates the Tcpi_rttcur field in the record's packed storage.
func (s *Tcp_info) SetTcpi_rttcur(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[20:24], uint32(v))
}

// Tcpi_srtt returns the Tcpi_srtt field from the record's packed storage.
func (s *Tcp_info) Tcpi_srtt() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[24:28]))
}

// SetTcpi_srtt updates the Tcpi_srtt field in the record's packed storage.
func (s *Tcp_info) SetTcpi_srtt(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[24:28], uint32(v))
}

// Tcpi_rttvar returns the Tcpi_rttvar field from the record's packed storage.
func (s *Tcp_info) Tcpi_rttvar() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[28:32]))
}

// SetTcpi_rttvar updates the Tcpi_rttvar field in the record's packed storage.
func (s *Tcp_info) SetTcpi_rttvar(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[28:32], uint32(v))
}

// Tcpi_rttbest returns the Tcpi_rttbest field from the record's packed storage.
func (s *Tcp_info) Tcpi_rttbest() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[32:36]))
}

// SetTcpi_rttbest updates the Tcpi_rttbest field in the record's packed storage.
func (s *Tcp_info) SetTcpi_rttbest(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[32:36], uint32(v))
}

// Tcpi_snd_ssthresh returns the Tcpi_snd_ssthresh field from the record's packed storage.
func (s *Tcp_info) Tcpi_snd_ssthresh() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[36:40]))
}

// SetTcpi_snd_ssthresh updates the Tcpi_snd_ssthresh field in the record's packed storage.
func (s *Tcp_info) SetTcpi_snd_ssthresh(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[36:40], uint32(v))
}

// Tcpi_snd_cwnd returns the Tcpi_snd_cwnd field from the record's packed storage.
func (s *Tcp_info) Tcpi_snd_cwnd() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[40:44]))
}

// SetTcpi_snd_cwnd updates the Tcpi_snd_cwnd field in the record's packed storage.
func (s *Tcp_info) SetTcpi_snd_cwnd(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[40:44], uint32(v))
}

// Tcpi_rcv_space returns the Tcpi_rcv_space field from the record's packed storage.
func (s *Tcp_info) Tcpi_rcv_space() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[44:48]))
}

// SetTcpi_rcv_space updates the Tcpi_rcv_space field in the record's packed storage.
func (s *Tcp_info) SetTcpi_rcv_space(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[44:48], uint32(v))
}

// Tcpi_snd_wnd returns the Tcpi_snd_wnd field from the record's packed storage.
func (s *Tcp_info) Tcpi_snd_wnd() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[48:52]))
}

// SetTcpi_snd_wnd updates the Tcpi_snd_wnd field in the record's packed storage.
func (s *Tcp_info) SetTcpi_snd_wnd(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[48:52], uint32(v))
}

// Tcpi_snd_nxt returns the Tcpi_snd_nxt field from the record's packed storage.
func (s *Tcp_info) Tcpi_snd_nxt() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[52:56]))
}

// SetTcpi_snd_nxt updates the Tcpi_snd_nxt field in the record's packed storage.
func (s *Tcp_info) SetTcpi_snd_nxt(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[52:56], uint32(v))
}

// Tcpi_rcv_nxt returns the Tcpi_rcv_nxt field from the record's packed storage.
func (s *Tcp_info) Tcpi_rcv_nxt() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[56:60]))
}

// SetTcpi_rcv_nxt updates the Tcpi_rcv_nxt field in the record's packed storage.
func (s *Tcp_info) SetTcpi_rcv_nxt(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[56:60], uint32(v))
}

// Tcpi_last_outif returns the Tcpi_last_outif field from the record's packed storage.
func (s *Tcp_info) Tcpi_last_outif() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[60:64]))
}

// SetTcpi_last_outif updates the Tcpi_last_outif field in the record's packed storage.
func (s *Tcp_info) SetTcpi_last_outif(v int32) {
	binary.NativeEndian.PutUint32(s.storage[60:64], uint32(v))
}

// Tcpi_snd_sbbytes returns the Tcpi_snd_sbbytes field from the record's packed storage.
func (s *Tcp_info) Tcpi_snd_sbbytes() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[64:68]))
}

// SetTcpi_snd_sbbytes updates the Tcpi_snd_sbbytes field in the record's packed storage.
func (s *Tcp_info) SetTcpi_snd_sbbytes(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[64:68], uint32(v))
}

// Tcpi_txpackets returns the Tcpi_txpackets field from the record's packed storage.
func (s *Tcp_info) Tcpi_txpackets() U_int64_t {
	return U_int64_t(binary.NativeEndian.Uint64(s.storage[68:76]))
}

// SetTcpi_txpackets updates the Tcpi_txpackets field in the record's packed storage.
func (s *Tcp_info) SetTcpi_txpackets(v U_int64_t) {
	binary.NativeEndian.PutUint64(s.storage[68:76], uint64(v))
}

// Tcpi_txbytes returns the Tcpi_txbytes field from the record's packed storage.
func (s *Tcp_info) Tcpi_txbytes() U_int64_t {
	return U_int64_t(binary.NativeEndian.Uint64(s.storage[76:84]))
}

// SetTcpi_txbytes updates the Tcpi_txbytes field in the record's packed storage.
func (s *Tcp_info) SetTcpi_txbytes(v U_int64_t) {
	binary.NativeEndian.PutUint64(s.storage[76:84], uint64(v))
}

// Tcpi_txretransmitbytes returns the Tcpi_txretransmitbytes field from the record's packed storage.
func (s *Tcp_info) Tcpi_txretransmitbytes() U_int64_t {
	return U_int64_t(binary.NativeEndian.Uint64(s.storage[84:92]))
}

// SetTcpi_txretransmitbytes updates the Tcpi_txretransmitbytes field in the record's packed storage.
func (s *Tcp_info) SetTcpi_txretransmitbytes(v U_int64_t) {
	binary.NativeEndian.PutUint64(s.storage[84:92], uint64(v))
}

// Tcpi_txunacked returns the Tcpi_txunacked field from the record's packed storage.
func (s *Tcp_info) Tcpi_txunacked() U_int64_t {
	return U_int64_t(binary.NativeEndian.Uint64(s.storage[92:100]))
}

// SetTcpi_txunacked updates the Tcpi_txunacked field in the record's packed storage.
func (s *Tcp_info) SetTcpi_txunacked(v U_int64_t) {
	binary.NativeEndian.PutUint64(s.storage[92:100], uint64(v))
}

// Tcpi_rxpackets returns the Tcpi_rxpackets field from the record's packed storage.
func (s *Tcp_info) Tcpi_rxpackets() U_int64_t {
	return U_int64_t(binary.NativeEndian.Uint64(s.storage[100:108]))
}

// SetTcpi_rxpackets updates the Tcpi_rxpackets field in the record's packed storage.
func (s *Tcp_info) SetTcpi_rxpackets(v U_int64_t) {
	binary.NativeEndian.PutUint64(s.storage[100:108], uint64(v))
}

// Tcpi_rxbytes returns the Tcpi_rxbytes field from the record's packed storage.
func (s *Tcp_info) Tcpi_rxbytes() U_int64_t {
	return U_int64_t(binary.NativeEndian.Uint64(s.storage[108:116]))
}

// SetTcpi_rxbytes updates the Tcpi_rxbytes field in the record's packed storage.
func (s *Tcp_info) SetTcpi_rxbytes(v U_int64_t) {
	binary.NativeEndian.PutUint64(s.storage[108:116], uint64(v))
}

// Tcpi_rxduplicatebytes returns the Tcpi_rxduplicatebytes field from the record's packed storage.
func (s *Tcp_info) Tcpi_rxduplicatebytes() U_int64_t {
	return U_int64_t(binary.NativeEndian.Uint64(s.storage[116:124]))
}

// SetTcpi_rxduplicatebytes updates the Tcpi_rxduplicatebytes field in the record's packed storage.
func (s *Tcp_info) SetTcpi_rxduplicatebytes(v U_int64_t) {
	binary.NativeEndian.PutUint64(s.storage[116:124], uint64(v))
}

// Tcpi_rxoutoforderbytes returns the Tcpi_rxoutoforderbytes field from the record's packed storage.
func (s *Tcp_info) Tcpi_rxoutoforderbytes() U_int64_t {
	return U_int64_t(binary.NativeEndian.Uint64(s.storage[124:132]))
}

// SetTcpi_rxoutoforderbytes updates the Tcpi_rxoutoforderbytes field in the record's packed storage.
func (s *Tcp_info) SetTcpi_rxoutoforderbytes(v U_int64_t) {
	binary.NativeEndian.PutUint64(s.storage[124:132], uint64(v))
}

// Tcpi_snd_bw returns the Tcpi_snd_bw field from the record's packed storage.
func (s *Tcp_info) Tcpi_snd_bw() U_int64_t {
	return U_int64_t(binary.NativeEndian.Uint64(s.storage[132:140]))
}

// SetTcpi_snd_bw updates the Tcpi_snd_bw field in the record's packed storage.
func (s *Tcp_info) SetTcpi_snd_bw(v U_int64_t) {
	binary.NativeEndian.PutUint64(s.storage[132:140], uint64(v))
}

// Tcpi_synrexmits returns the Tcpi_synrexmits field from the record's packed storage.
func (s *Tcp_info) Tcpi_synrexmits() U_int8_t {
	return *(*U_int8_t)(unsafe.Pointer(&s.storage[140]))
}

// SetTcpi_synrexmits updates the Tcpi_synrexmits field in the record's packed storage.
func (s *Tcp_info) SetTcpi_synrexmits(v U_int8_t) {
	*(*U_int8_t)(unsafe.Pointer(&s.storage[140])) = v
}

// Tcpi_unused1 returns the Tcpi_unused1 field from the record's packed storage.
func (s *Tcp_info) Tcpi_unused1() U_int8_t {
	return *(*U_int8_t)(unsafe.Pointer(&s.storage[141]))
}

// SetTcpi_unused1 updates the Tcpi_unused1 field in the record's packed storage.
func (s *Tcp_info) SetTcpi_unused1(v U_int8_t) {
	*(*U_int8_t)(unsafe.Pointer(&s.storage[141])) = v
}

// Tcpi_unused2 returns the Tcpi_unused2 field from the record's packed storage.
func (s *Tcp_info) Tcpi_unused2() U_int16_t {
	return U_int16_t(binary.NativeEndian.Uint16(s.storage[142:144]))
}

// SetTcpi_unused2 updates the Tcpi_unused2 field in the record's packed storage.
func (s *Tcp_info) SetTcpi_unused2(v U_int16_t) {
	binary.NativeEndian.PutUint16(s.storage[142:144], uint16(v))
}

// Tcpi_cell_rxpackets returns the Tcpi_cell_rxpackets field from the record's packed storage.
func (s *Tcp_info) Tcpi_cell_rxpackets() U_int64_t {
	return U_int64_t(binary.NativeEndian.Uint64(s.storage[144:152]))
}

// SetTcpi_cell_rxpackets updates the Tcpi_cell_rxpackets field in the record's packed storage.
func (s *Tcp_info) SetTcpi_cell_rxpackets(v U_int64_t) {
	binary.NativeEndian.PutUint64(s.storage[144:152], uint64(v))
}

// Tcpi_cell_rxbytes returns the Tcpi_cell_rxbytes field from the record's packed storage.
func (s *Tcp_info) Tcpi_cell_rxbytes() U_int64_t {
	return U_int64_t(binary.NativeEndian.Uint64(s.storage[152:160]))
}

// SetTcpi_cell_rxbytes updates the Tcpi_cell_rxbytes field in the record's packed storage.
func (s *Tcp_info) SetTcpi_cell_rxbytes(v U_int64_t) {
	binary.NativeEndian.PutUint64(s.storage[152:160], uint64(v))
}

// Tcpi_cell_txpackets returns the Tcpi_cell_txpackets field from the record's packed storage.
func (s *Tcp_info) Tcpi_cell_txpackets() U_int64_t {
	return U_int64_t(binary.NativeEndian.Uint64(s.storage[160:168]))
}

// SetTcpi_cell_txpackets updates the Tcpi_cell_txpackets field in the record's packed storage.
func (s *Tcp_info) SetTcpi_cell_txpackets(v U_int64_t) {
	binary.NativeEndian.PutUint64(s.storage[160:168], uint64(v))
}

// Tcpi_cell_txbytes returns the Tcpi_cell_txbytes field from the record's packed storage.
func (s *Tcp_info) Tcpi_cell_txbytes() U_int64_t {
	return U_int64_t(binary.NativeEndian.Uint64(s.storage[168:176]))
}

// SetTcpi_cell_txbytes updates the Tcpi_cell_txbytes field in the record's packed storage.
func (s *Tcp_info) SetTcpi_cell_txbytes(v U_int64_t) {
	binary.NativeEndian.PutUint64(s.storage[168:176], uint64(v))
}

// Tcpi_wifi_rxpackets returns the Tcpi_wifi_rxpackets field from the record's packed storage.
func (s *Tcp_info) Tcpi_wifi_rxpackets() U_int64_t {
	return U_int64_t(binary.NativeEndian.Uint64(s.storage[176:184]))
}

// SetTcpi_wifi_rxpackets updates the Tcpi_wifi_rxpackets field in the record's packed storage.
func (s *Tcp_info) SetTcpi_wifi_rxpackets(v U_int64_t) {
	binary.NativeEndian.PutUint64(s.storage[176:184], uint64(v))
}

// Tcpi_wifi_rxbytes returns the Tcpi_wifi_rxbytes field from the record's packed storage.
func (s *Tcp_info) Tcpi_wifi_rxbytes() U_int64_t {
	return U_int64_t(binary.NativeEndian.Uint64(s.storage[184:192]))
}

// SetTcpi_wifi_rxbytes updates the Tcpi_wifi_rxbytes field in the record's packed storage.
func (s *Tcp_info) SetTcpi_wifi_rxbytes(v U_int64_t) {
	binary.NativeEndian.PutUint64(s.storage[184:192], uint64(v))
}

// Tcpi_wifi_txpackets returns the Tcpi_wifi_txpackets field from the record's packed storage.
func (s *Tcp_info) Tcpi_wifi_txpackets() U_int64_t {
	return U_int64_t(binary.NativeEndian.Uint64(s.storage[192:200]))
}

// SetTcpi_wifi_txpackets updates the Tcpi_wifi_txpackets field in the record's packed storage.
func (s *Tcp_info) SetTcpi_wifi_txpackets(v U_int64_t) {
	binary.NativeEndian.PutUint64(s.storage[192:200], uint64(v))
}

// Tcpi_wifi_txbytes returns the Tcpi_wifi_txbytes field from the record's packed storage.
func (s *Tcp_info) Tcpi_wifi_txbytes() U_int64_t {
	return U_int64_t(binary.NativeEndian.Uint64(s.storage[200:208]))
}

// SetTcpi_wifi_txbytes updates the Tcpi_wifi_txbytes field in the record's packed storage.
func (s *Tcp_info) SetTcpi_wifi_txbytes(v U_int64_t) {
	binary.NativeEndian.PutUint64(s.storage[200:208], uint64(v))
}

// Tcpi_wired_rxpackets returns the Tcpi_wired_rxpackets field from the record's packed storage.
func (s *Tcp_info) Tcpi_wired_rxpackets() U_int64_t {
	return U_int64_t(binary.NativeEndian.Uint64(s.storage[208:216]))
}

// SetTcpi_wired_rxpackets updates the Tcpi_wired_rxpackets field in the record's packed storage.
func (s *Tcp_info) SetTcpi_wired_rxpackets(v U_int64_t) {
	binary.NativeEndian.PutUint64(s.storage[208:216], uint64(v))
}

// Tcpi_wired_rxbytes returns the Tcpi_wired_rxbytes field from the record's packed storage.
func (s *Tcp_info) Tcpi_wired_rxbytes() U_int64_t {
	return U_int64_t(binary.NativeEndian.Uint64(s.storage[216:224]))
}

// SetTcpi_wired_rxbytes updates the Tcpi_wired_rxbytes field in the record's packed storage.
func (s *Tcp_info) SetTcpi_wired_rxbytes(v U_int64_t) {
	binary.NativeEndian.PutUint64(s.storage[216:224], uint64(v))
}

// Tcpi_wired_txpackets returns the Tcpi_wired_txpackets field from the record's packed storage.
func (s *Tcp_info) Tcpi_wired_txpackets() U_int64_t {
	return U_int64_t(binary.NativeEndian.Uint64(s.storage[224:232]))
}

// SetTcpi_wired_txpackets updates the Tcpi_wired_txpackets field in the record's packed storage.
func (s *Tcp_info) SetTcpi_wired_txpackets(v U_int64_t) {
	binary.NativeEndian.PutUint64(s.storage[224:232], uint64(v))
}

// Tcpi_wired_txbytes returns the Tcpi_wired_txbytes field from the record's packed storage.
func (s *Tcp_info) Tcpi_wired_txbytes() U_int64_t {
	return U_int64_t(binary.NativeEndian.Uint64(s.storage[232:240]))
}

// SetTcpi_wired_txbytes updates the Tcpi_wired_txbytes field in the record's packed storage.
func (s *Tcp_info) SetTcpi_wired_txbytes(v U_int64_t) {
	binary.NativeEndian.PutUint64(s.storage[232:240], uint64(v))
}

// Tcpi_connstatus returns the Tcpi_connstatus field from the record's packed storage.
func (s *Tcp_info) Tcpi_connstatus() Tcp_conn_status {
	return *(*Tcp_conn_status)(unsafe.Pointer(&s.storage[240]))
}

// SetTcpi_connstatus updates the Tcpi_connstatus field in the record's packed storage.
func (s *Tcp_info) SetTcpi_connstatus(v Tcp_conn_status) {
	*(*Tcp_conn_status)(unsafe.Pointer(&s.storage[240])) = v
}

// Tcpi_tfo_cookie_req returns the Tcpi_tfo_cookie_req bitfield from the record's packed storage.
func (s *Tcp_info) Tcpi_tfo_cookie_req() uint32 {
	return uint32((s.storage[244] >> 0) & 0x1)
}

// SetTcpi_tfo_cookie_req updates the Tcpi_tfo_cookie_req bitfield in the record's packed storage.
func (s *Tcp_info) SetTcpi_tfo_cookie_req(v uint32) {
	s.storage[244] = (s.storage[244] &^ uint8(0x1<<0)) | uint8((uint8(v)&0x1)<<0)
}

// Tcpi_tfo_cookie_rcv returns the Tcpi_tfo_cookie_rcv bitfield from the record's packed storage.
func (s *Tcp_info) Tcpi_tfo_cookie_rcv() uint32 {
	return uint32((s.storage[244] >> 1) & 0x1)
}

// SetTcpi_tfo_cookie_rcv updates the Tcpi_tfo_cookie_rcv bitfield in the record's packed storage.
func (s *Tcp_info) SetTcpi_tfo_cookie_rcv(v uint32) {
	s.storage[244] = (s.storage[244] &^ uint8(0x1<<1)) | uint8((uint8(v)&0x1)<<1)
}

// Tcpi_tfo_syn_loss returns the Tcpi_tfo_syn_loss bitfield from the record's packed storage.
func (s *Tcp_info) Tcpi_tfo_syn_loss() uint32 {
	return uint32((s.storage[244] >> 2) & 0x1)
}

// SetTcpi_tfo_syn_loss updates the Tcpi_tfo_syn_loss bitfield in the record's packed storage.
func (s *Tcp_info) SetTcpi_tfo_syn_loss(v uint32) {
	s.storage[244] = (s.storage[244] &^ uint8(0x1<<2)) | uint8((uint8(v)&0x1)<<2)
}

// Tcpi_tfo_syn_data_sent returns the Tcpi_tfo_syn_data_sent bitfield from the record's packed storage.
func (s *Tcp_info) Tcpi_tfo_syn_data_sent() uint32 {
	return uint32((s.storage[244] >> 3) & 0x1)
}

// SetTcpi_tfo_syn_data_sent updates the Tcpi_tfo_syn_data_sent bitfield in the record's packed storage.
func (s *Tcp_info) SetTcpi_tfo_syn_data_sent(v uint32) {
	s.storage[244] = (s.storage[244] &^ uint8(0x1<<3)) | uint8((uint8(v)&0x1)<<3)
}

// Tcpi_tfo_syn_data_acked returns the Tcpi_tfo_syn_data_acked bitfield from the record's packed storage.
func (s *Tcp_info) Tcpi_tfo_syn_data_acked() uint32 {
	return uint32((s.storage[244] >> 4) & 0x1)
}

// SetTcpi_tfo_syn_data_acked updates the Tcpi_tfo_syn_data_acked bitfield in the record's packed storage.
func (s *Tcp_info) SetTcpi_tfo_syn_data_acked(v uint32) {
	s.storage[244] = (s.storage[244] &^ uint8(0x1<<4)) | uint8((uint8(v)&0x1)<<4)
}

// Tcpi_tfo_syn_data_rcv returns the Tcpi_tfo_syn_data_rcv bitfield from the record's packed storage.
func (s *Tcp_info) Tcpi_tfo_syn_data_rcv() uint32 {
	return uint32((s.storage[244] >> 5) & 0x1)
}

// SetTcpi_tfo_syn_data_rcv updates the Tcpi_tfo_syn_data_rcv bitfield in the record's packed storage.
func (s *Tcp_info) SetTcpi_tfo_syn_data_rcv(v uint32) {
	s.storage[244] = (s.storage[244] &^ uint8(0x1<<5)) | uint8((uint8(v)&0x1)<<5)
}

// Tcpi_tfo_cookie_req_rcv returns the Tcpi_tfo_cookie_req_rcv bitfield from the record's packed storage.
func (s *Tcp_info) Tcpi_tfo_cookie_req_rcv() uint32 {
	return uint32((s.storage[244] >> 6) & 0x1)
}

// SetTcpi_tfo_cookie_req_rcv updates the Tcpi_tfo_cookie_req_rcv bitfield in the record's packed storage.
func (s *Tcp_info) SetTcpi_tfo_cookie_req_rcv(v uint32) {
	s.storage[244] = (s.storage[244] &^ uint8(0x1<<6)) | uint8((uint8(v)&0x1)<<6)
}

// Tcpi_tfo_cookie_sent returns the Tcpi_tfo_cookie_sent bitfield from the record's packed storage.
func (s *Tcp_info) Tcpi_tfo_cookie_sent() uint32 {
	return uint32((s.storage[244] >> 7) & 0x1)
}

// SetTcpi_tfo_cookie_sent updates the Tcpi_tfo_cookie_sent bitfield in the record's packed storage.
func (s *Tcp_info) SetTcpi_tfo_cookie_sent(v uint32) {
	s.storage[244] = (s.storage[244] &^ uint8(0x1<<7)) | uint8((uint8(v)&0x1)<<7)
}

// Tcpi_tfo_cookie_invalid returns the Tcpi_tfo_cookie_invalid bitfield from the record's packed storage.
func (s *Tcp_info) Tcpi_tfo_cookie_invalid() uint32 {
	return uint32((s.storage[245] >> 0) & 0x1)
}

// SetTcpi_tfo_cookie_invalid updates the Tcpi_tfo_cookie_invalid bitfield in the record's packed storage.
func (s *Tcp_info) SetTcpi_tfo_cookie_invalid(v uint32) {
	s.storage[245] = (s.storage[245] &^ uint8(0x1<<0)) | uint8((uint8(v)&0x1)<<0)
}

// Tcpi_tfo_cookie_wrong returns the Tcpi_tfo_cookie_wrong bitfield from the record's packed storage.
func (s *Tcp_info) Tcpi_tfo_cookie_wrong() uint32 {
	return uint32((s.storage[245] >> 1) & 0x1)
}

// SetTcpi_tfo_cookie_wrong updates the Tcpi_tfo_cookie_wrong bitfield in the record's packed storage.
func (s *Tcp_info) SetTcpi_tfo_cookie_wrong(v uint32) {
	s.storage[245] = (s.storage[245] &^ uint8(0x1<<1)) | uint8((uint8(v)&0x1)<<1)
}

// Tcpi_tfo_no_cookie_rcv returns the Tcpi_tfo_no_cookie_rcv bitfield from the record's packed storage.
func (s *Tcp_info) Tcpi_tfo_no_cookie_rcv() uint32 {
	return uint32((s.storage[245] >> 2) & 0x1)
}

// SetTcpi_tfo_no_cookie_rcv updates the Tcpi_tfo_no_cookie_rcv bitfield in the record's packed storage.
func (s *Tcp_info) SetTcpi_tfo_no_cookie_rcv(v uint32) {
	s.storage[245] = (s.storage[245] &^ uint8(0x1<<2)) | uint8((uint8(v)&0x1)<<2)
}

// Tcpi_tfo_heuristics_disable returns the Tcpi_tfo_heuristics_disable bitfield from the record's packed storage.
func (s *Tcp_info) Tcpi_tfo_heuristics_disable() uint32 {
	return uint32((s.storage[245] >> 3) & 0x1)
}

// SetTcpi_tfo_heuristics_disable updates the Tcpi_tfo_heuristics_disable bitfield in the record's packed storage.
func (s *Tcp_info) SetTcpi_tfo_heuristics_disable(v uint32) {
	s.storage[245] = (s.storage[245] &^ uint8(0x1<<3)) | uint8((uint8(v)&0x1)<<3)
}

// Tcpi_tfo_send_blackhole returns the Tcpi_tfo_send_blackhole bitfield from the record's packed storage.
func (s *Tcp_info) Tcpi_tfo_send_blackhole() uint32 {
	return uint32((s.storage[245] >> 4) & 0x1)
}

// SetTcpi_tfo_send_blackhole updates the Tcpi_tfo_send_blackhole bitfield in the record's packed storage.
func (s *Tcp_info) SetTcpi_tfo_send_blackhole(v uint32) {
	s.storage[245] = (s.storage[245] &^ uint8(0x1<<4)) | uint8((uint8(v)&0x1)<<4)
}

// Tcpi_tfo_recv_blackhole returns the Tcpi_tfo_recv_blackhole bitfield from the record's packed storage.
func (s *Tcp_info) Tcpi_tfo_recv_blackhole() uint32 {
	return uint32((s.storage[245] >> 5) & 0x1)
}

// SetTcpi_tfo_recv_blackhole updates the Tcpi_tfo_recv_blackhole bitfield in the record's packed storage.
func (s *Tcp_info) SetTcpi_tfo_recv_blackhole(v uint32) {
	s.storage[245] = (s.storage[245] &^ uint8(0x1<<5)) | uint8((uint8(v)&0x1)<<5)
}

// Tcpi_tfo_onebyte_proxy returns the Tcpi_tfo_onebyte_proxy bitfield from the record's packed storage.
func (s *Tcp_info) Tcpi_tfo_onebyte_proxy() uint32 {
	return uint32((s.storage[245] >> 6) & 0x1)
}

// SetTcpi_tfo_onebyte_proxy updates the Tcpi_tfo_onebyte_proxy bitfield in the record's packed storage.
func (s *Tcp_info) SetTcpi_tfo_onebyte_proxy(v uint32) {
	s.storage[245] = (s.storage[245] &^ uint8(0x1<<6)) | uint8((uint8(v)&0x1)<<6)
}

// Tcpi_ecn_client_setup returns the Tcpi_ecn_client_setup bitfield from the record's packed storage.
func (s *Tcp_info) Tcpi_ecn_client_setup() uint32 {
	return uint32((s.storage[245] >> 7) & 0x1)
}

// SetTcpi_ecn_client_setup updates the Tcpi_ecn_client_setup bitfield in the record's packed storage.
func (s *Tcp_info) SetTcpi_ecn_client_setup(v uint32) {
	s.storage[245] = (s.storage[245] &^ uint8(0x1<<7)) | uint8((uint8(v)&0x1)<<7)
}

// Tcpi_ecn_server_setup returns the Tcpi_ecn_server_setup bitfield from the record's packed storage.
func (s *Tcp_info) Tcpi_ecn_server_setup() uint32 {
	return uint32((s.storage[246] >> 0) & 0x1)
}

// SetTcpi_ecn_server_setup updates the Tcpi_ecn_server_setup bitfield in the record's packed storage.
func (s *Tcp_info) SetTcpi_ecn_server_setup(v uint32) {
	s.storage[246] = (s.storage[246] &^ uint8(0x1<<0)) | uint8((uint8(v)&0x1)<<0)
}

// Tcpi_ecn_success returns the Tcpi_ecn_success bitfield from the record's packed storage.
func (s *Tcp_info) Tcpi_ecn_success() uint32 {
	return uint32((s.storage[246] >> 1) & 0x1)
}

// SetTcpi_ecn_success updates the Tcpi_ecn_success bitfield in the record's packed storage.
func (s *Tcp_info) SetTcpi_ecn_success(v uint32) {
	s.storage[246] = (s.storage[246] &^ uint8(0x1<<1)) | uint8((uint8(v)&0x1)<<1)
}

// Tcpi_ecn_lost_syn returns the Tcpi_ecn_lost_syn bitfield from the record's packed storage.
func (s *Tcp_info) Tcpi_ecn_lost_syn() uint32 {
	return uint32((s.storage[246] >> 2) & 0x1)
}

// SetTcpi_ecn_lost_syn updates the Tcpi_ecn_lost_syn bitfield in the record's packed storage.
func (s *Tcp_info) SetTcpi_ecn_lost_syn(v uint32) {
	s.storage[246] = (s.storage[246] &^ uint8(0x1<<2)) | uint8((uint8(v)&0x1)<<2)
}

// Tcpi_ecn_lost_synack returns the Tcpi_ecn_lost_synack bitfield from the record's packed storage.
func (s *Tcp_info) Tcpi_ecn_lost_synack() uint32 {
	return uint32((s.storage[246] >> 3) & 0x1)
}

// SetTcpi_ecn_lost_synack updates the Tcpi_ecn_lost_synack bitfield in the record's packed storage.
func (s *Tcp_info) SetTcpi_ecn_lost_synack(v uint32) {
	s.storage[246] = (s.storage[246] &^ uint8(0x1<<3)) | uint8((uint8(v)&0x1)<<3)
}

// Tcpi_local_peer returns the Tcpi_local_peer bitfield from the record's packed storage.
func (s *Tcp_info) Tcpi_local_peer() uint32 {
	return uint32((s.storage[246] >> 4) & 0x1)
}

// SetTcpi_local_peer updates the Tcpi_local_peer bitfield in the record's packed storage.
func (s *Tcp_info) SetTcpi_local_peer(v uint32) {
	s.storage[246] = (s.storage[246] &^ uint8(0x1<<4)) | uint8((uint8(v)&0x1)<<4)
}

// Tcpi_if_cell returns the Tcpi_if_cell bitfield from the record's packed storage.
func (s *Tcp_info) Tcpi_if_cell() uint32 {
	return uint32((s.storage[246] >> 5) & 0x1)
}

// SetTcpi_if_cell updates the Tcpi_if_cell bitfield in the record's packed storage.
func (s *Tcp_info) SetTcpi_if_cell(v uint32) {
	s.storage[246] = (s.storage[246] &^ uint8(0x1<<5)) | uint8((uint8(v)&0x1)<<5)
}

// Tcpi_if_wifi returns the Tcpi_if_wifi bitfield from the record's packed storage.
func (s *Tcp_info) Tcpi_if_wifi() uint32 {
	return uint32((s.storage[246] >> 6) & 0x1)
}

// SetTcpi_if_wifi updates the Tcpi_if_wifi bitfield in the record's packed storage.
func (s *Tcp_info) SetTcpi_if_wifi(v uint32) {
	s.storage[246] = (s.storage[246] &^ uint8(0x1<<6)) | uint8((uint8(v)&0x1)<<6)
}

// Tcpi_if_wired returns the Tcpi_if_wired bitfield from the record's packed storage.
func (s *Tcp_info) Tcpi_if_wired() uint32 {
	return uint32((s.storage[246] >> 7) & 0x1)
}

// SetTcpi_if_wired updates the Tcpi_if_wired bitfield in the record's packed storage.
func (s *Tcp_info) SetTcpi_if_wired(v uint32) {
	s.storage[246] = (s.storage[246] &^ uint8(0x1<<7)) | uint8((uint8(v)&0x1)<<7)
}

// Tcpi_if_wifi_infra returns the Tcpi_if_wifi_infra bitfield from the record's packed storage.
func (s *Tcp_info) Tcpi_if_wifi_infra() uint32 {
	return uint32((s.storage[247] >> 0) & 0x1)
}

// SetTcpi_if_wifi_infra updates the Tcpi_if_wifi_infra bitfield in the record's packed storage.
func (s *Tcp_info) SetTcpi_if_wifi_infra(v uint32) {
	s.storage[247] = (s.storage[247] &^ uint8(0x1<<0)) | uint8((uint8(v)&0x1)<<0)
}

// Tcpi_if_wifi_awdl returns the Tcpi_if_wifi_awdl bitfield from the record's packed storage.
func (s *Tcp_info) Tcpi_if_wifi_awdl() uint32 {
	return uint32((s.storage[247] >> 1) & 0x1)
}

// SetTcpi_if_wifi_awdl updates the Tcpi_if_wifi_awdl bitfield in the record's packed storage.
func (s *Tcp_info) SetTcpi_if_wifi_awdl(v uint32) {
	s.storage[247] = (s.storage[247] &^ uint8(0x1<<1)) | uint8((uint8(v)&0x1)<<1)
}

// Tcpi_snd_background returns the Tcpi_snd_background bitfield from the record's packed storage.
func (s *Tcp_info) Tcpi_snd_background() uint32 {
	return uint32((s.storage[247] >> 2) & 0x1)
}

// SetTcpi_snd_background updates the Tcpi_snd_background bitfield in the record's packed storage.
func (s *Tcp_info) SetTcpi_snd_background(v uint32) {
	s.storage[247] = (s.storage[247] &^ uint8(0x1<<2)) | uint8((uint8(v)&0x1)<<2)
}

// Tcpi_rcv_background returns the Tcpi_rcv_background bitfield from the record's packed storage.
func (s *Tcp_info) Tcpi_rcv_background() uint32 {
	return uint32((s.storage[247] >> 3) & 0x1)
}

// SetTcpi_rcv_background updates the Tcpi_rcv_background bitfield in the record's packed storage.
func (s *Tcp_info) SetTcpi_rcv_background(v uint32) {
	s.storage[247] = (s.storage[247] &^ uint8(0x1<<3)) | uint8((uint8(v)&0x1)<<3)
}

// Tcpi_l4s_enabled returns the Tcpi_l4s_enabled bitfield from the record's packed storage.
func (s *Tcp_info) Tcpi_l4s_enabled() uint32 {
	return uint32((s.storage[247] >> 4) & 0x1)
}

// SetTcpi_l4s_enabled updates the Tcpi_l4s_enabled bitfield in the record's packed storage.
func (s *Tcp_info) SetTcpi_l4s_enabled(v uint32) {
	s.storage[247] = (s.storage[247] &^ uint8(0x1<<4)) | uint8((uint8(v)&0x1)<<4)
}

// Tcpi_ecn_recv_ce returns the Tcpi_ecn_recv_ce field from the record's packed storage.
func (s *Tcp_info) Tcpi_ecn_recv_ce() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[248:252]))
}

// SetTcpi_ecn_recv_ce updates the Tcpi_ecn_recv_ce field in the record's packed storage.
func (s *Tcp_info) SetTcpi_ecn_recv_ce(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[248:252], uint32(v))
}

// Tcpi_ecn_recv_cwr returns the Tcpi_ecn_recv_cwr field from the record's packed storage.
func (s *Tcp_info) Tcpi_ecn_recv_cwr() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[252:256]))
}

// SetTcpi_ecn_recv_cwr updates the Tcpi_ecn_recv_cwr field in the record's packed storage.
func (s *Tcp_info) SetTcpi_ecn_recv_cwr(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[252:256], uint32(v))
}

// Tcpi_rcvoopack returns the Tcpi_rcvoopack field from the record's packed storage.
func (s *Tcp_info) Tcpi_rcvoopack() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[256:260]))
}

// SetTcpi_rcvoopack updates the Tcpi_rcvoopack field in the record's packed storage.
func (s *Tcp_info) SetTcpi_rcvoopack(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[256:260], uint32(v))
}

// Tcpi_pawsdrop returns the Tcpi_pawsdrop field from the record's packed storage.
func (s *Tcp_info) Tcpi_pawsdrop() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[260:264]))
}

// SetTcpi_pawsdrop updates the Tcpi_pawsdrop field in the record's packed storage.
func (s *Tcp_info) SetTcpi_pawsdrop(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[260:264], uint32(v))
}

// Tcpi_sack_recovery_episode returns the Tcpi_sack_recovery_episode field from the record's packed storage.
func (s *Tcp_info) Tcpi_sack_recovery_episode() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[264:268]))
}

// SetTcpi_sack_recovery_episode updates the Tcpi_sack_recovery_episode field in the record's packed storage.
func (s *Tcp_info) SetTcpi_sack_recovery_episode(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[264:268], uint32(v))
}

// Tcpi_reordered_pkts returns the Tcpi_reordered_pkts field from the record's packed storage.
func (s *Tcp_info) Tcpi_reordered_pkts() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[268:272]))
}

// SetTcpi_reordered_pkts updates the Tcpi_reordered_pkts field in the record's packed storage.
func (s *Tcp_info) SetTcpi_reordered_pkts(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[268:272], uint32(v))
}

// Tcpi_dsack_sent returns the Tcpi_dsack_sent field from the record's packed storage.
func (s *Tcp_info) Tcpi_dsack_sent() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[272:276]))
}

// SetTcpi_dsack_sent updates the Tcpi_dsack_sent field in the record's packed storage.
func (s *Tcp_info) SetTcpi_dsack_sent(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[272:276], uint32(v))
}

// Tcpi_dsack_recvd returns the Tcpi_dsack_recvd field from the record's packed storage.
func (s *Tcp_info) Tcpi_dsack_recvd() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[276:280]))
}

// SetTcpi_dsack_recvd updates the Tcpi_dsack_recvd field in the record's packed storage.
func (s *Tcp_info) SetTcpi_dsack_recvd(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[276:280], uint32(v))
}

// Tcpi_flowhash returns the Tcpi_flowhash field from the record's packed storage.
func (s *Tcp_info) Tcpi_flowhash() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[280:284]))
}

// SetTcpi_flowhash updates the Tcpi_flowhash field in the record's packed storage.
func (s *Tcp_info) SetTcpi_flowhash(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[280:284], uint32(v))
}

// Tcpi_txretransmitpackets returns the Tcpi_txretransmitpackets field from the record's packed storage.
func (s *Tcp_info) Tcpi_txretransmitpackets() U_int64_t {
	return U_int64_t(binary.NativeEndian.Uint64(s.storage[284:292]))
}

// SetTcpi_txretransmitpackets updates the Tcpi_txretransmitpackets field in the record's packed storage.
func (s *Tcp_info) SetTcpi_txretransmitpackets(v U_int64_t) {
	binary.NativeEndian.PutUint64(s.storage[284:292], uint64(v))
}

// Tcpi_rcv_srtt returns the Tcpi_rcv_srtt field from the record's packed storage.
func (s *Tcp_info) Tcpi_rcv_srtt() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[292:296]))
}

// SetTcpi_rcv_srtt updates the Tcpi_rcv_srtt field in the record's packed storage.
func (s *Tcp_info) SetTcpi_rcv_srtt(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[292:296], uint32(v))
}

// Tcpi_client_accecn_state returns the Tcpi_client_accecn_state field from the record's packed storage.
func (s *Tcp_info) Tcpi_client_accecn_state() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[296:300]))
}

// SetTcpi_client_accecn_state updates the Tcpi_client_accecn_state field in the record's packed storage.
func (s *Tcp_info) SetTcpi_client_accecn_state(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[296:300], uint32(v))
}

// Tcpi_server_accecn_state returns the Tcpi_server_accecn_state field from the record's packed storage.
func (s *Tcp_info) Tcpi_server_accecn_state() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[300:304]))
}

// SetTcpi_server_accecn_state updates the Tcpi_server_accecn_state field in the record's packed storage.
func (s *Tcp_info) SetTcpi_server_accecn_state(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[300:304], uint32(v))
}

// Tcpi_ecn_capable_packets_sent returns the Tcpi_ecn_capable_packets_sent field from the record's packed storage.
func (s *Tcp_info) Tcpi_ecn_capable_packets_sent() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[304:312]))
}

// SetTcpi_ecn_capable_packets_sent updates the Tcpi_ecn_capable_packets_sent field in the record's packed storage.
func (s *Tcp_info) SetTcpi_ecn_capable_packets_sent(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[304:312], uint64(v))
}

// Tcpi_ecn_capable_packets_acked returns the Tcpi_ecn_capable_packets_acked field from the record's packed storage.
func (s *Tcp_info) Tcpi_ecn_capable_packets_acked() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[312:320]))
}

// SetTcpi_ecn_capable_packets_acked updates the Tcpi_ecn_capable_packets_acked field in the record's packed storage.
func (s *Tcp_info) SetTcpi_ecn_capable_packets_acked(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[312:320], uint64(v))
}

// Tcpi_ecn_capable_packets_marked returns the Tcpi_ecn_capable_packets_marked field from the record's packed storage.
func (s *Tcp_info) Tcpi_ecn_capable_packets_marked() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[320:328]))
}

// SetTcpi_ecn_capable_packets_marked updates the Tcpi_ecn_capable_packets_marked field in the record's packed storage.
func (s *Tcp_info) SetTcpi_ecn_capable_packets_marked(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[320:328], uint64(v))
}

// Tcpi_ecn_capable_packets_lost returns the Tcpi_ecn_capable_packets_lost field from the record's packed storage.
func (s *Tcp_info) Tcpi_ecn_capable_packets_lost() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[328:336]))
}

// SetTcpi_ecn_capable_packets_lost updates the Tcpi_ecn_capable_packets_lost field in the record's packed storage.
func (s *Tcp_info) SetTcpi_ecn_capable_packets_lost(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[328:336], uint64(v))
}

// Tcpi_received_ce_packets returns the Tcpi_received_ce_packets field from the record's packed storage.
func (s *Tcp_info) Tcpi_received_ce_packets() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[336:344]))
}

// SetTcpi_received_ce_packets updates the Tcpi_received_ce_packets field in the record's packed storage.
func (s *Tcp_info) SetTcpi_received_ce_packets(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[336:344], uint64(v))
}

// Tcpi_received_ect0_bytes returns the Tcpi_received_ect0_bytes field from the record's packed storage.
func (s *Tcp_info) Tcpi_received_ect0_bytes() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[344:352]))
}

// SetTcpi_received_ect0_bytes updates the Tcpi_received_ect0_bytes field in the record's packed storage.
func (s *Tcp_info) SetTcpi_received_ect0_bytes(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[344:352], uint64(v))
}

// Tcpi_received_ect1_bytes returns the Tcpi_received_ect1_bytes field from the record's packed storage.
func (s *Tcp_info) Tcpi_received_ect1_bytes() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[352:360]))
}

// SetTcpi_received_ect1_bytes updates the Tcpi_received_ect1_bytes field in the record's packed storage.
func (s *Tcp_info) SetTcpi_received_ect1_bytes(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[352:360], uint64(v))
}

// Tcpi_received_ce_bytes returns the Tcpi_received_ce_bytes field from the record's packed storage.
func (s *Tcp_info) Tcpi_received_ce_bytes() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[360:368]))
}

// SetTcpi_received_ce_bytes updates the Tcpi_received_ce_bytes field in the record's packed storage.
func (s *Tcp_info) SetTcpi_received_ce_bytes(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[360:368], uint64(v))
}

// Tcpi_delivered_ect0_bytes returns the Tcpi_delivered_ect0_bytes field from the record's packed storage.
func (s *Tcp_info) Tcpi_delivered_ect0_bytes() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[368:376]))
}

// SetTcpi_delivered_ect0_bytes updates the Tcpi_delivered_ect0_bytes field in the record's packed storage.
func (s *Tcp_info) SetTcpi_delivered_ect0_bytes(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[368:376], uint64(v))
}

// Tcpi_delivered_ect1_bytes returns the Tcpi_delivered_ect1_bytes field from the record's packed storage.
func (s *Tcp_info) Tcpi_delivered_ect1_bytes() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[376:384]))
}

// SetTcpi_delivered_ect1_bytes updates the Tcpi_delivered_ect1_bytes field in the record's packed storage.
func (s *Tcp_info) SetTcpi_delivered_ect1_bytes(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[376:384], uint64(v))
}

// Tcpi_delivered_ce_bytes returns the Tcpi_delivered_ce_bytes field from the record's packed storage.
func (s *Tcp_info) Tcpi_delivered_ce_bytes() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[384:392]))
}

// SetTcpi_delivered_ce_bytes updates the Tcpi_delivered_ce_bytes field in the record's packed storage.
func (s *Tcp_info) SetTcpi_delivered_ce_bytes(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[384:392], uint64(v))
}

// Tcpi_flow_control_total_time returns the Tcpi_flow_control_total_time field from the record's packed storage.
func (s *Tcp_info) Tcpi_flow_control_total_time() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[392:400]))
}

// SetTcpi_flow_control_total_time updates the Tcpi_flow_control_total_time field in the record's packed storage.
func (s *Tcp_info) SetTcpi_flow_control_total_time(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[392:400], uint64(v))
}

// Tcpi_rcvwnd_limited_total_time returns the Tcpi_rcvwnd_limited_total_time field from the record's packed storage.
func (s *Tcp_info) Tcpi_rcvwnd_limited_total_time() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[400:408]))
}

// SetTcpi_rcvwnd_limited_total_time updates the Tcpi_rcvwnd_limited_total_time field in the record's packed storage.
func (s *Tcp_info) SetTcpi_rcvwnd_limited_total_time(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[400:408], uint64(v))
}

// Tcpi_pacing_rate returns the Tcpi_pacing_rate field from the record's packed storage.
func (s *Tcp_info) Tcpi_pacing_rate() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[408:416]))
}

// SetTcpi_pacing_rate updates the Tcpi_pacing_rate field in the record's packed storage.
func (s *Tcp_info) SetTcpi_pacing_rate(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[408:416], uint64(v))
}

// Tcpi_max_pacing_rate returns the Tcpi_max_pacing_rate field from the record's packed storage.
func (s *Tcp_info) Tcpi_max_pacing_rate() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[416:424]))
}

// SetTcpi_max_pacing_rate updates the Tcpi_max_pacing_rate field in the record's packed storage.
func (s *Tcp_info) SetTcpi_max_pacing_rate(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[416:424], uint64(v))
}

// Tcp_measure_bw_burst
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/tcp_measure_bw_burst
type Tcp_measure_bw_burst struct {
	Min_burst_size U_int32_t
	Max_burst_size U_int32_t
}

// Tcp_notify_ack_complete
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/tcp_notify_ack_complete
type Tcp_notify_ack_complete struct {
	Notify_pending        U_int32_t
	Notify_complete_count U_int32_t
	Notify_complete_id    [10]Tcp_notify_ack_id_t
}

// Tcpcb
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/tcpcb
type Tcpcb struct {
	T_segq            Tsegqe_head
	T_dupacks         int32
	Unused            U_int32_t
	T_timer           [4]int32
	T_inpcb           U_int32_t
	T_state           int32
	T_flags           U_int
	T_force           int32
	Snd_una           Tcp_seq
	Snd_max           Tcp_seq
	Snd_nxt           Tcp_seq
	Snd_up            Tcp_seq
	Snd_wl1           Tcp_seq
	Snd_wl2           Tcp_seq
	Iss               Tcp_seq
	Irs               Tcp_seq
	Rcv_nxt           Tcp_seq
	Rcv_adv           Tcp_seq
	Rcv_wnd           U_int32_t
	Rcv_up            Tcp_seq
	Snd_wnd           U_int32_t
	Snd_cwnd          U_int32_t
	Snd_ssthresh      U_int32_t
	T_maxopd          U_int
	T_rcvtime         U_int32_t
	T_starttime       U_int32_t
	T_rtttime         int32
	T_rtseq           Tcp_seq
	T_rxtcur          int32
	T_maxseg          U_int
	T_srtt            int32
	T_rttvar          int32
	T_rxtshift        int32
	T_rttmin          U_int
	T_rttupdated      U_int32_t
	Max_sndwnd        U_int32_t
	T_softerror       int32
	T_oobflags        int8
	T_iobc            int8
	Snd_scale         U_char
	Rcv_scale         U_char
	Request_r_scale   U_char
	Requested_s_scale U_char
	Ts_recent         U_int32_t
	Ts_recent_age     U_int32_t
	Last_ack_sent     Tcp_seq
	Cc_send           Tcp_cc
	Cc_recv           Tcp_cc
	Snd_recover       Tcp_seq
	Snd_cwnd_prev     U_int32_t
	Snd_ssthresh_prev U_int32_t
	T_badrxtwin       U_int32_t
}

// Tcphdr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/tcphdr
type Tcphdr struct {
	Th_sport  uint16
	Th_dport  uint16
	Th_seq    Tcp_seq
	Th_ack    Tcp_seq
	bitfield4 uint8
	Th_flags  uint8
	Th_win    uint16
	Th_sum    uint16
	Th_urp    uint16
}

// Th_x2 returns the Th_x2 bitfield.
func (s *Tcphdr) Th_x2() uint8 {
	return (s.bitfield4 >> 0) & ((1 << 4) - 1)
}

// SetTh_x2 updates the Th_x2 bitfield.
func (s *Tcphdr) SetTh_x2(v uint8) {
	const mask uint8 = (1 << 4) - 1
	s.bitfield4 = (s.bitfield4 &^ (mask << 0)) | ((v & mask) << 0)
}

// Th_off returns the Th_off bitfield.
func (s *Tcphdr) Th_off() uint8 {
	return (s.bitfield4 >> 4) & ((1 << 4) - 1)
}

// SetTh_off updates the Th_off bitfield.
func (s *Tcphdr) SetTh_off(v uint8) {
	const mask uint8 = (1 << 4) - 1
	s.bitfield4 = (s.bitfield4 &^ (mask << 4)) | ((v & mask) << 4)
}

// Tcpiphdr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/tcpiphdr
type Tcpiphdr struct {
	Ti_i Ipovly
	Ti_t Tcphdr
}

// Tcpstat
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/tcpstat
type Tcpstat struct {
	Tcps_connattempt                              U_int32_t
	Tcps_accepts                                  U_int32_t
	Tcps_connects                                 U_int32_t
	Tcps_drops                                    U_int32_t
	Tcps_conndrops                                U_int32_t
	Tcps_closed                                   U_int32_t
	Tcps_segstimed                                U_int32_t
	Tcps_rttupdated                               U_int32_t
	Tcps_delack                                   U_int32_t
	Tcps_timeoutdrop                              U_int32_t
	Tcps_rexmttimeo                               U_int32_t
	Tcps_persisttimeo                             U_int32_t
	Tcps_keeptimeo                                U_int32_t
	Tcps_keepprobe                                U_int32_t
	Tcps_keepdrops                                U_int32_t
	Tcps_sndtotal                                 U_int32_t
	Tcps_sndpack                                  U_int32_t
	Tcps_sndbyte                                  U_int32_t
	Tcps_sndrexmitpack                            U_int32_t
	Tcps_sndrexmitbyte                            U_int32_t
	Tcps_sndacks                                  U_int32_t
	Tcps_sndprobe                                 U_int32_t
	Tcps_sndurg                                   U_int32_t
	Tcps_sndwinup                                 U_int32_t
	Tcps_sndctrl                                  U_int32_t
	Tcps_rcvtotal                                 U_int32_t
	Tcps_rcvpack                                  U_int32_t
	Tcps_rcvbyte                                  U_int32_t
	Tcps_rcvbadsum                                U_int32_t
	Tcps_rcvbadoff                                U_int32_t
	Tcps_rcvmemdrop                               U_int32_t
	Tcps_rcvshort                                 U_int32_t
	Tcps_rcvduppack                               U_int32_t
	Tcps_rcvdupbyte                               U_int32_t
	Tcps_rcvpartduppack                           U_int32_t
	Tcps_rcvpartdupbyte                           U_int32_t
	Tcps_rcvoopack                                U_int32_t
	Tcps_rcvoobyte                                U_int32_t
	Tcps_rcvpackafterwin                          U_int32_t
	Tcps_rcvbyteafterwin                          U_int32_t
	Tcps_rcvafterclose                            U_int32_t
	Tcps_rcvwinprobe                              U_int32_t
	Tcps_rcvdupack                                U_int32_t
	Tcps_rcvacktoomuch                            U_int32_t
	Tcps_rcvackpack                               U_int32_t
	Tcps_rcvackbyte                               U_int32_t
	Tcps_rcvwinupd                                U_int32_t
	Tcps_pawsdrop                                 U_int32_t
	Tcps_predack                                  U_int32_t
	Tcps_preddat                                  U_int32_t
	Tcps_cachedrtt                                U_int32_t
	Tcps_cachedrttvar                             U_int32_t
	Tcps_cachedssthresh                           U_int32_t
	Tcps_usedrtt                                  U_int32_t
	Tcps_usedrttvar                               U_int32_t
	Tcps_usedssthresh                             U_int32_t
	Tcps_persistdrop                              U_int32_t
	Tcps_badsyn                                   U_int32_t
	Tcps_mturesent                                U_int32_t
	Tcps_listendrop                               U_int32_t
	Tcps_synchallenge                             U_int32_t
	Tcps_rstchallenge                             U_int32_t
	Tcps_minmssdrops                              U_int32_t
	Tcps_sndrexmitbad                             U_int32_t
	Tcps_badrst                                   U_int32_t
	Tcps_sc_dropped                               U_int32_t
	Tcps_sc_completed                             U_int32_t
	Tcps_sc_aborted                               U_int32_t
	Tcps_sc_sendcookie                            U_int32_t
	Tcps_sc_recvcookie                            U_int32_t
	Tcps_sack_recovery_episode                    U_int32_t
	Tcps_sack_rexmits                             U_int32_t
	Tcps_sack_rexmit_bytes                        U_int32_t
	Tcps_sack_rcv_blocks                          U_int32_t
	Tcps_sack_send_blocks                         U_int32_t
	Tcps_sack_sboverflow                          U_int32_t
	Tcps_rack_recovery_episode                    U_int32_t
	Tcps_rack_reordering_timeout_recovery_episode U_int32_t
	Tcps_rack_rexmits                             U_int32_t
	Tcps_bg_rcvtotal                              U_int32_t
	Tcps_rxtfindrop                               U_int32_t
	Tcps_fcholdpacket                             U_int32_t
	Tcps_limited_txt                              U_int32_t
	Tcps_early_rexmt                              U_int32_t
	Tcps_sack_ackadv                              U_int32_t
	Tcps_rcv_swcsum                               U_int32_t
	Tcps_rcv_swcsum_bytes                         U_int32_t
	Tcps_rcv6_swcsum                              U_int32_t
	Tcps_rcv6_swcsum_bytes                        U_int32_t
	Tcps_snd_swcsum                               U_int32_t
	Tcps_snd_swcsum_bytes                         U_int32_t
	Tcps_snd6_swcsum                              U_int32_t
	Tcps_snd6_swcsum_bytes                        U_int32_t
	Tcps_invalid_mpcap                            U_int32_t
	Tcps_invalid_joins                            U_int32_t
	Tcps_mpcap_fallback                           U_int32_t
	Tcps_join_fallback                            U_int32_t
	Tcps_estab_fallback                           U_int32_t
	Tcps_invalid_opt                              U_int32_t
	Tcps_mp_reducedwin                            U_int32_t
	Tcps_mp_badcsum                               U_int32_t
	Tcps_mp_oodata                                U_int32_t
	Tcps_mp_switches                              U_int32_t
	Tcps_mp_rcvtotal                              U_int32_t
	Tcps_mp_rcvbytes                              U_int32_t
	Tcps_mp_sndpacks                              U_int32_t
	Tcps_mp_sndbytes                              U_int32_t
	Tcps_join_rxmts                               U_int32_t
	Tcps_tailloss_rto                             U_int32_t
	Tcps_reordered_pkts                           U_int32_t
	Tcps_recovered_pkts                           U_int32_t
	Tcps_pto                                      U_int32_t
	Tcps_rto_after_pto                            U_int32_t
	Tcps_tlp_recovery                             U_int32_t
	Tcps_tlp_recoverlastpkt                       U_int32_t
	Tcps_ecn_client_success                       U_int32_t
	Tcps_ecn_recv_ece                             U_int32_t
	Tcps_ecn_sent_ece                             U_int32_t
	Tcps_detect_reordering                        U_int32_t
	Tcps_delay_recovery                           U_int32_t
	Tcps_avoid_rxmt                               U_int32_t
	Tcps_pto_in_recovery                          U_int32_t
	Tcps_pmtudbh_reverted                         U_int32_t
	Tcps_dsack_ackloss                            U_int32_t
	Tcps_dsack_badrexmt                           U_int32_t
	Tcps_dsack_sent                               U_int32_t
	Tcps_dsack_recvd                              U_int32_t
	Tcps_dsack_recvd_old                          U_int32_t
	Tcps_mp_sel_rtt                               U_int32_t
	Tcps_mp_sel_rto                               U_int32_t
	Tcps_mp_num_probes                            U_int32_t
	Tcps_mp_verdowngrade                          U_int32_t
	Tcps_drop_after_sleep                         U_int32_t
	Tcps_probe_if                                 U_int32_t
	Tcps_probe_if_conflict                        U_int32_t
	Tcps_ecn_client_setup                         U_int32_t
	Tcps_ecn_server_setup                         U_int32_t
	Tcps_ecn_server_success                       U_int32_t
	Tcps_ecn_ace_syn_not_ect                      U_int32_t
	Tcps_ecn_ace_syn_ect1                         U_int32_t
	Tcps_ecn_ace_syn_ect0                         U_int32_t
	Tcps_ecn_ace_syn_ce                           U_int32_t
	Tcps_ecn_lost_synack                          U_int32_t
	Tcps_ecn_lost_syn                             U_int32_t
	Tcps_ecn_not_supported                        U_int32_t
	Tcps_ecn_recv_ce                              U_int32_t
	Tcps_ecn_ace_recv_ce                          U_int32_t
	Tcps_ecn_conn_recv_ce                         U_int32_t
	Tcps_ecn_conn_recv_ece                        U_int32_t
	Tcps_ecn_conn_plnoce                          U_int32_t
	Tcps_ecn_conn_pl_ce                           U_int32_t
	Tcps_ecn_conn_nopl_ce                         U_int32_t
	Tcps_ecn_fallback_synloss                     U_int32_t
	Tcps_ecn_fallback_reorder                     U_int32_t
	Tcps_ecn_fallback_ce                          U_int32_t
	Tcps_tfo_syn_data_rcv                         U_int32_t
	Tcps_tfo_cookie_req_rcv                       U_int32_t
	Tcps_tfo_cookie_sent                          U_int32_t
	Tcps_tfo_cookie_invalid                       U_int32_t
	Tcps_tfo_cookie_req                           U_int32_t
	Tcps_tfo_cookie_rcv                           U_int32_t
	Tcps_tfo_syn_data_sent                        U_int32_t
	Tcps_tfo_syn_data_acked                       U_int32_t
	Tcps_tfo_syn_loss                             U_int32_t
	Tcps_tfo_blackhole                            U_int32_t
	Tcps_tfo_cookie_wrong                         U_int32_t
	Tcps_tfo_no_cookie_rcv                        U_int32_t
	Tcps_tfo_heuristics_disable                   U_int32_t
	Tcps_tfo_sndblackhole                         U_int32_t
	Tcps_mss_to_default                           U_int32_t
	Tcps_mss_to_medium                            U_int32_t
	Tcps_mss_to_low                               U_int32_t
	Tcps_ecn_fallback_droprst                     U_int32_t
	Tcps_ecn_fallback_droprxmt                    U_int32_t
	Tcps_ecn_fallback_synrst                      U_int32_t
	Tcps_mptcp_rcvmemdrop                         U_int32_t
	Tcps_mptcp_rcvduppack                         U_int32_t
	Tcps_mptcp_rcvpackafterwin                    U_int32_t
	Tcps_timer_drift_le_1_ms                      U_int32_t
	Tcps_timer_drift_le_10_ms                     U_int32_t
	Tcps_timer_drift_le_20_ms                     U_int32_t
	Tcps_timer_drift_le_50_ms                     U_int32_t
	Tcps_timer_drift_le_100_ms                    U_int32_t
	Tcps_timer_drift_le_200_ms                    U_int32_t
	Tcps_timer_drift_le_500_ms                    U_int32_t
	Tcps_timer_drift_le_1000_ms                   U_int32_t
	Tcps_timer_drift_gt_1000_ms                   U_int32_t
	Tcps_mptcp_handover_attempt                   U_int32_t
	Tcps_mptcp_interactive_attempt                U_int32_t
	Tcps_mptcp_aggregate_attempt                  U_int32_t
	Tcps_mptcp_fp_handover_attempt                U_int32_t
	Tcps_mptcp_fp_interactive_attempt             U_int32_t
	Tcps_mptcp_fp_aggregate_attempt               U_int32_t
	Tcps_mptcp_heuristic_fallback                 U_int32_t
	Tcps_mptcp_fp_heuristic_fallback              U_int32_t
	Tcps_mptcp_handover_success_wifi              U_int32_t
	Tcps_mptcp_handover_success_cell              U_int32_t
	Tcps_mptcp_interactive_success                U_int32_t
	Tcps_mptcp_aggregate_success                  U_int32_t
	Tcps_mptcp_fp_handover_success_wifi           U_int32_t
	Tcps_mptcp_fp_handover_success_cell           U_int32_t
	Tcps_mptcp_fp_interactive_success             U_int32_t
	Tcps_mptcp_fp_aggregate_success               U_int32_t
	Tcps_mptcp_handover_cell_from_wifi            U_int32_t
	Tcps_mptcp_handover_wifi_from_cell            U_int32_t
	Tcps_mptcp_interactive_cell_from_wifi         U_int32_t
	Tcps_mptcp_handover_cell_bytes                U_int64_t
	Tcps_mptcp_interactive_cell_bytes             U_int64_t
	Tcps_mptcp_aggregate_cell_bytes               U_int64_t
	Tcps_mptcp_handover_all_bytes                 U_int64_t
	Tcps_mptcp_interactive_all_bytes              U_int64_t
	Tcps_mptcp_aggregate_all_bytes                U_int64_t
	Tcps_mptcp_back_to_wifi                       U_int32_t
	Tcps_mptcp_wifi_proxy                         U_int32_t
	Tcps_mptcp_cell_proxy                         U_int32_t
	Tcps_ka_offload_drops                         U_int32_t
	Tcps_mptcp_triggered_cell                     U_int32_t
	Tcps_fin_timeout_drops                        U_int32_t
	Tcps_rst_dup_suppressed                       U_int64_t
	Tcps_rst_not_suppressed                       U_int64_t
	Tcps_dsack_disable                            U_int32_t
	Tcps_hc_added                                 U_int32_t
	Tcps_hc_bucketoverflow                        U_int32_t
	Tcps_mp_outofwin                              U_int32_t
	Tcps_mp_sel_peer                              U_int32_t
	Tcps_mp_sel_symtomsd                          U_int32_t
	Tcps_nostretchack                             U_int32_t
	Tcps_pcbcachemiss                             U_int32_t
	Tcps_rescue_rxmt                              U_int32_t
	Tcps_sc_added                                 U_int32_t
	Tcps_sc_badack                                U_int32_t
	Tcps_sc_bucketoverflow                        U_int32_t
	Tcps_sc_cacheoverflow                         U_int32_t
	Tcps_sc_dupsyn                                U_int32_t
	Tcps_sc_reset                                 U_int32_t
	Tcps_sc_retransmitted                         U_int32_t
	Tcps_sc_stale                                 U_int32_t
	Tcps_sc_unreach                               U_int32_t
	Tcps_sc_zonefail                              U_int32_t
	Tcps_unnecessary_rxmt                         U_int32_t
	Tcps_unused_1                                 U_int32_t
	Tcps_unused_2                                 U_int32_t
	Tcps_unused_3                                 U_int32_t
}

// Tcpstat_local
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/tcpstat_local
type Tcpstat_local struct {
	Badformat            U_int64_t
	Unspecv6             U_int64_t
	Synfin               U_int64_t
	Badformatipsec       U_int64_t
	Noconnnolist         U_int64_t
	Noconnlist           U_int64_t
	Listbadsyn           U_int64_t
	Icmp6unreach         U_int64_t
	Deprecate6           U_int64_t
	Ooopacket            U_int64_t
	Rstinsynrcv          U_int64_t
	Dospacket            U_int64_t
	Cleanup              U_int64_t
	Synwindow            U_int64_t
	Linkheur_stealthdrop U_int64_t
	Linkheur_noackpri    U_int64_t
	Linkheur_comprxmt    U_int64_t
	Linkheur_synrxmt     U_int64_t
	Linkheur_rxmtfloor   U_int64_t
}

// Telemetry_notification_subsystem
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/telemetry_notification_subsystem
type Telemetry_notification_subsystem struct {
	Server   unsafe.Pointer
	Start    Mach_msg_id_t
	End      Mach_msg_id_t
	Maxsize  uint32
	Reserved Vm_address_t
	Routine  unsafe.Pointer
}

// Termios
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/termios
type Termios struct {
	C_iflag  Tcflag_t
	C_oflag  Tcflag_t
	C_cflag  Tcflag_t
	C_lflag  Tcflag_t
	C_cc     [20]Cc_t
	C_ispeed Speed_t
	C_ospeed Speed_t
}

// Termios32
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/termios32
type Termios32 struct {
	C_iflag  uint32
	C_oflag  uint32
	C_cflag  uint32
	C_lflag  uint32
	C_cc     [20]Cc_t
	C_ispeed uint32
	C_ospeed uint32
}

// Thread_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/thread_command
type Thread_command struct {
	Cmd     uint32 // Common to all load command structures. For this structure, set to `LC_THREAD` or `LC_UNIXTHREAD`.
	Cmdsize uint32
}

// Thread_crash_exclaves_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/thread_crash_exclaves_info
type Thread_crash_exclaves_info struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [20]byte
}

// Tcei_scid returns the Tcei_scid field from the record's packed storage.
func (s *Thread_crash_exclaves_info) Tcei_scid() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[0:8]))
}

// SetTcei_scid updates the Tcei_scid field in the record's packed storage.
func (s *Thread_crash_exclaves_info) SetTcei_scid(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[0:8], uint64(v))
}

// Tcei_thread_id returns the Tcei_thread_id field from the record's packed storage.
func (s *Thread_crash_exclaves_info) Tcei_thread_id() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[8:16]))
}

// SetTcei_thread_id updates the Tcei_thread_id field in the record's packed storage.
func (s *Thread_crash_exclaves_info) SetTcei_thread_id(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[8:16], uint64(v))
}

// Tcei_flags returns the Tcei_flags field from the record's packed storage.
func (s *Thread_crash_exclaves_info) Tcei_flags() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[16:20]))
}

// SetTcei_flags updates the Tcei_flags field in the record's packed storage.
func (s *Thread_crash_exclaves_info) SetTcei_flags(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[16:20], uint32(v))
}

// Thread_delta_snapshot_v2
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/thread_delta_snapshot_v2
type Thread_delta_snapshot_v2 struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [48]byte
}

// Tds_thread_id returns the Tds_thread_id field from the record's packed storage.
func (s *Thread_delta_snapshot_v2) Tds_thread_id() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[0:8]))
}

// SetTds_thread_id updates the Tds_thread_id field in the record's packed storage.
func (s *Thread_delta_snapshot_v2) SetTds_thread_id(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[0:8], uint64(v))
}

// Tds_voucher_identifier returns the Tds_voucher_identifier field from the record's packed storage.
func (s *Thread_delta_snapshot_v2) Tds_voucher_identifier() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[8:16]))
}

// SetTds_voucher_identifier updates the Tds_voucher_identifier field in the record's packed storage.
func (s *Thread_delta_snapshot_v2) SetTds_voucher_identifier(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[8:16], uint64(v))
}

// Tds_ss_flags returns the Tds_ss_flags field from the record's packed storage.
func (s *Thread_delta_snapshot_v2) Tds_ss_flags() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[16:24]))
}

// SetTds_ss_flags updates the Tds_ss_flags field in the record's packed storage.
func (s *Thread_delta_snapshot_v2) SetTds_ss_flags(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[16:24], uint64(v))
}

// Tds_last_made_runnable_time returns the Tds_last_made_runnable_time field from the record's packed storage.
func (s *Thread_delta_snapshot_v2) Tds_last_made_runnable_time() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[24:32]))
}

// SetTds_last_made_runnable_time updates the Tds_last_made_runnable_time field in the record's packed storage.
func (s *Thread_delta_snapshot_v2) SetTds_last_made_runnable_time(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[24:32], uint64(v))
}

// Tds_state returns the Tds_state field from the record's packed storage.
func (s *Thread_delta_snapshot_v2) Tds_state() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[32:36]))
}

// SetTds_state updates the Tds_state field in the record's packed storage.
func (s *Thread_delta_snapshot_v2) SetTds_state(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[32:36], uint32(v))
}

// Tds_sched_flags returns the Tds_sched_flags field from the record's packed storage.
func (s *Thread_delta_snapshot_v2) Tds_sched_flags() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[36:40]))
}

// SetTds_sched_flags updates the Tds_sched_flags field in the record's packed storage.
func (s *Thread_delta_snapshot_v2) SetTds_sched_flags(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[36:40], uint32(v))
}

// Tds_base_priority returns the Tds_base_priority field from the record's packed storage.
func (s *Thread_delta_snapshot_v2) Tds_base_priority() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[40:42]))
}

// SetTds_base_priority updates the Tds_base_priority field in the record's packed storage.
func (s *Thread_delta_snapshot_v2) SetTds_base_priority(v int16) {
	binary.NativeEndian.PutUint16(s.storage[40:42], uint16(v))
}

// Tds_sched_priority returns the Tds_sched_priority field from the record's packed storage.
func (s *Thread_delta_snapshot_v2) Tds_sched_priority() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[42:44]))
}

// SetTds_sched_priority updates the Tds_sched_priority field in the record's packed storage.
func (s *Thread_delta_snapshot_v2) SetTds_sched_priority(v int16) {
	binary.NativeEndian.PutUint16(s.storage[42:44], uint16(v))
}

// Tds_eqos returns the Tds_eqos field from the record's packed storage.
func (s *Thread_delta_snapshot_v2) Tds_eqos() uint8 {
	return uint8(s.storage[44])
}

// SetTds_eqos updates the Tds_eqos field in the record's packed storage.
func (s *Thread_delta_snapshot_v2) SetTds_eqos(v uint8) {
	s.storage[44] = uint8(v)
}

// Tds_rqos returns the Tds_rqos field from the record's packed storage.
func (s *Thread_delta_snapshot_v2) Tds_rqos() uint8 {
	return uint8(s.storage[45])
}

// SetTds_rqos updates the Tds_rqos field in the record's packed storage.
func (s *Thread_delta_snapshot_v2) SetTds_rqos(v uint8) {
	s.storage[45] = uint8(v)
}

// Tds_rqos_override returns the Tds_rqos_override field from the record's packed storage.
func (s *Thread_delta_snapshot_v2) Tds_rqos_override() uint8 {
	return uint8(s.storage[46])
}

// SetTds_rqos_override updates the Tds_rqos_override field in the record's packed storage.
func (s *Thread_delta_snapshot_v2) SetTds_rqos_override(v uint8) {
	s.storage[46] = uint8(v)
}

// Tds_io_tier returns the Tds_io_tier field from the record's packed storage.
func (s *Thread_delta_snapshot_v2) Tds_io_tier() uint8 {
	return uint8(s.storage[47])
}

// SetTds_io_tier updates the Tds_io_tier field in the record's packed storage.
func (s *Thread_delta_snapshot_v2) SetTds_io_tier(v uint8) {
	s.storage[47] = uint8(v)
}

// Thread_delta_snapshot_v3
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/thread_delta_snapshot_v3
type Thread_delta_snapshot_v3 struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [64]byte
}

// Tds_thread_id returns the Tds_thread_id field from the record's packed storage.
func (s *Thread_delta_snapshot_v3) Tds_thread_id() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[0:8]))
}

// SetTds_thread_id updates the Tds_thread_id field in the record's packed storage.
func (s *Thread_delta_snapshot_v3) SetTds_thread_id(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[0:8], uint64(v))
}

// Tds_voucher_identifier returns the Tds_voucher_identifier field from the record's packed storage.
func (s *Thread_delta_snapshot_v3) Tds_voucher_identifier() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[8:16]))
}

// SetTds_voucher_identifier updates the Tds_voucher_identifier field in the record's packed storage.
func (s *Thread_delta_snapshot_v3) SetTds_voucher_identifier(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[8:16], uint64(v))
}

// Tds_ss_flags returns the Tds_ss_flags field from the record's packed storage.
func (s *Thread_delta_snapshot_v3) Tds_ss_flags() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[16:24]))
}

// SetTds_ss_flags updates the Tds_ss_flags field in the record's packed storage.
func (s *Thread_delta_snapshot_v3) SetTds_ss_flags(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[16:24], uint64(v))
}

// Tds_last_made_runnable_time returns the Tds_last_made_runnable_time field from the record's packed storage.
func (s *Thread_delta_snapshot_v3) Tds_last_made_runnable_time() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[24:32]))
}

// SetTds_last_made_runnable_time updates the Tds_last_made_runnable_time field in the record's packed storage.
func (s *Thread_delta_snapshot_v3) SetTds_last_made_runnable_time(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[24:32], uint64(v))
}

// Tds_state returns the Tds_state field from the record's packed storage.
func (s *Thread_delta_snapshot_v3) Tds_state() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[32:36]))
}

// SetTds_state updates the Tds_state field in the record's packed storage.
func (s *Thread_delta_snapshot_v3) SetTds_state(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[32:36], uint32(v))
}

// Tds_sched_flags returns the Tds_sched_flags field from the record's packed storage.
func (s *Thread_delta_snapshot_v3) Tds_sched_flags() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[36:40]))
}

// SetTds_sched_flags updates the Tds_sched_flags field in the record's packed storage.
func (s *Thread_delta_snapshot_v3) SetTds_sched_flags(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[36:40], uint32(v))
}

// Tds_base_priority returns the Tds_base_priority field from the record's packed storage.
func (s *Thread_delta_snapshot_v3) Tds_base_priority() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[40:42]))
}

// SetTds_base_priority updates the Tds_base_priority field in the record's packed storage.
func (s *Thread_delta_snapshot_v3) SetTds_base_priority(v int16) {
	binary.NativeEndian.PutUint16(s.storage[40:42], uint16(v))
}

// Tds_sched_priority returns the Tds_sched_priority field from the record's packed storage.
func (s *Thread_delta_snapshot_v3) Tds_sched_priority() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[42:44]))
}

// SetTds_sched_priority updates the Tds_sched_priority field in the record's packed storage.
func (s *Thread_delta_snapshot_v3) SetTds_sched_priority(v int16) {
	binary.NativeEndian.PutUint16(s.storage[42:44], uint16(v))
}

// Tds_eqos returns the Tds_eqos field from the record's packed storage.
func (s *Thread_delta_snapshot_v3) Tds_eqos() uint8 {
	return uint8(s.storage[44])
}

// SetTds_eqos updates the Tds_eqos field in the record's packed storage.
func (s *Thread_delta_snapshot_v3) SetTds_eqos(v uint8) {
	s.storage[44] = uint8(v)
}

// Tds_rqos returns the Tds_rqos field from the record's packed storage.
func (s *Thread_delta_snapshot_v3) Tds_rqos() uint8 {
	return uint8(s.storage[45])
}

// SetTds_rqos updates the Tds_rqos field in the record's packed storage.
func (s *Thread_delta_snapshot_v3) SetTds_rqos(v uint8) {
	s.storage[45] = uint8(v)
}

// Tds_rqos_override returns the Tds_rqos_override field from the record's packed storage.
func (s *Thread_delta_snapshot_v3) Tds_rqos_override() uint8 {
	return uint8(s.storage[46])
}

// SetTds_rqos_override updates the Tds_rqos_override field in the record's packed storage.
func (s *Thread_delta_snapshot_v3) SetTds_rqos_override(v uint8) {
	s.storage[46] = uint8(v)
}

// Tds_io_tier returns the Tds_io_tier field from the record's packed storage.
func (s *Thread_delta_snapshot_v3) Tds_io_tier() uint8 {
	return uint8(s.storage[47])
}

// SetTds_io_tier updates the Tds_io_tier field in the record's packed storage.
func (s *Thread_delta_snapshot_v3) SetTds_io_tier(v uint8) {
	s.storage[47] = uint8(v)
}

// Tds_requested_policy returns the Tds_requested_policy field from the record's packed storage.
func (s *Thread_delta_snapshot_v3) Tds_requested_policy() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[48:56]))
}

// SetTds_requested_policy updates the Tds_requested_policy field in the record's packed storage.
func (s *Thread_delta_snapshot_v3) SetTds_requested_policy(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[48:56], uint64(v))
}

// Tds_effective_policy returns the Tds_effective_policy field from the record's packed storage.
func (s *Thread_delta_snapshot_v3) Tds_effective_policy() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[56:64]))
}

// SetTds_effective_policy updates the Tds_effective_policy field in the record's packed storage.
func (s *Thread_delta_snapshot_v3) SetTds_effective_policy(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[56:64], uint64(v))
}

// Thread_exclaves_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/thread_exclaves_info
type Thread_exclaves_info struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [16]byte
}

// Tei_scid returns the Tei_scid field from the record's packed storage.
func (s *Thread_exclaves_info) Tei_scid() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[0:8]))
}

// SetTei_scid updates the Tei_scid field in the record's packed storage.
func (s *Thread_exclaves_info) SetTei_scid(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[0:8], uint64(v))
}

// Tei_thread_offset returns the Tei_thread_offset field from the record's packed storage.
func (s *Thread_exclaves_info) Tei_thread_offset() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetTei_thread_offset updates the Tei_thread_offset field in the record's packed storage.
func (s *Thread_exclaves_info) SetTei_thread_offset(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[8:12], uint32(v))
}

// Tei_flags returns the Tei_flags field from the record's packed storage.
func (s *Thread_exclaves_info) Tei_flags() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[12:16]))
}

// SetTei_flags updates the Tei_flags field in the record's packed storage.
func (s *Thread_exclaves_info) SetTei_flags(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[12:16], uint32(v))
}

// Thread_group_snapshot
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/thread_group_snapshot
type Thread_group_snapshot struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [24]byte
}

// Tgs_id returns the Tgs_id field from the record's packed storage.
func (s *Thread_group_snapshot) Tgs_id() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[0:8]))
}

// SetTgs_id updates the Tgs_id field in the record's packed storage.
func (s *Thread_group_snapshot) SetTgs_id(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[0:8], uint64(v))
}

// Tgs_name returns the Tgs_name field from the record's packed storage.
func (s *Thread_group_snapshot) Tgs_name() [16]int8 {
	return *(*[16]int8)(unsafe.Pointer(&s.storage[8]))
}

// SetTgs_name updates the Tgs_name field in the record's packed storage.
func (s *Thread_group_snapshot) SetTgs_name(v [16]int8) {
	*(*[16]int8)(unsafe.Pointer(&s.storage[8])) = v
}

// Thread_group_snapshot_v2
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/thread_group_snapshot_v2
type Thread_group_snapshot_v2 struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [32]byte
}

// Tgs_id returns the Tgs_id field from the record's packed storage.
func (s *Thread_group_snapshot_v2) Tgs_id() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[0:8]))
}

// SetTgs_id updates the Tgs_id field in the record's packed storage.
func (s *Thread_group_snapshot_v2) SetTgs_id(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[0:8], uint64(v))
}

// Tgs_name returns the Tgs_name field from the record's packed storage.
func (s *Thread_group_snapshot_v2) Tgs_name() [16]int8 {
	return *(*[16]int8)(unsafe.Pointer(&s.storage[8]))
}

// SetTgs_name updates the Tgs_name field in the record's packed storage.
func (s *Thread_group_snapshot_v2) SetTgs_name(v [16]int8) {
	*(*[16]int8)(unsafe.Pointer(&s.storage[8])) = v
}

// Tgs_flags returns the Tgs_flags field from the record's packed storage.
func (s *Thread_group_snapshot_v2) Tgs_flags() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[24:32]))
}

// SetTgs_flags updates the Tgs_flags field in the record's packed storage.
func (s *Thread_group_snapshot_v2) SetTgs_flags(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[24:32], uint64(v))
}

// Thread_group_snapshot_v3
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/thread_group_snapshot_v3
type Thread_group_snapshot_v3 struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [48]byte
}

// Tgs_id returns the Tgs_id field from the record's packed storage.
func (s *Thread_group_snapshot_v3) Tgs_id() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[0:8]))
}

// SetTgs_id updates the Tgs_id field in the record's packed storage.
func (s *Thread_group_snapshot_v3) SetTgs_id(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[0:8], uint64(v))
}

// Tgs_name returns the Tgs_name field from the record's packed storage.
func (s *Thread_group_snapshot_v3) Tgs_name() [16]int8 {
	return *(*[16]int8)(unsafe.Pointer(&s.storage[8]))
}

// SetTgs_name updates the Tgs_name field in the record's packed storage.
func (s *Thread_group_snapshot_v3) SetTgs_name(v [16]int8) {
	*(*[16]int8)(unsafe.Pointer(&s.storage[8])) = v
}

// Tgs_flags returns the Tgs_flags field from the record's packed storage.
func (s *Thread_group_snapshot_v3) Tgs_flags() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[24:32]))
}

// SetTgs_flags updates the Tgs_flags field in the record's packed storage.
func (s *Thread_group_snapshot_v3) SetTgs_flags(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[24:32], uint64(v))
}

// Tgs_name_cont returns the Tgs_name_cont field from the record's packed storage.
func (s *Thread_group_snapshot_v3) Tgs_name_cont() [16]int8 {
	return *(*[16]int8)(unsafe.Pointer(&s.storage[32]))
}

// SetTgs_name_cont updates the Tgs_name_cont field in the record's packed storage.
func (s *Thread_group_snapshot_v3) SetTgs_name_cont(v [16]int8) {
	*(*[16]int8)(unsafe.Pointer(&s.storage[32])) = v
}

// Thread_snapshot
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/thread_snapshot
type Thread_snapshot struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [316]byte
}

// Snapshot_magic returns the Snapshot_magic field from the record's packed storage.
func (s *Thread_snapshot) Snapshot_magic() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetSnapshot_magic updates the Snapshot_magic field in the record's packed storage.
func (s *Thread_snapshot) SetSnapshot_magic(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// Nkern_frames returns the Nkern_frames field from the record's packed storage.
func (s *Thread_snapshot) Nkern_frames() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetNkern_frames updates the Nkern_frames field in the record's packed storage.
func (s *Thread_snapshot) SetNkern_frames(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// Nuser_frames returns the Nuser_frames field from the record's packed storage.
func (s *Thread_snapshot) Nuser_frames() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetNuser_frames updates the Nuser_frames field in the record's packed storage.
func (s *Thread_snapshot) SetNuser_frames(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[8:12], uint32(v))
}

// Wait_event returns the Wait_event field from the record's packed storage.
func (s *Thread_snapshot) Wait_event() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[12:20]))
}

// SetWait_event updates the Wait_event field in the record's packed storage.
func (s *Thread_snapshot) SetWait_event(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[12:20], uint64(v))
}

// Continuation returns the Continuation field from the record's packed storage.
func (s *Thread_snapshot) Continuation() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[20:28]))
}

// SetContinuation updates the Continuation field in the record's packed storage.
func (s *Thread_snapshot) SetContinuation(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[20:28], uint64(v))
}

// Thread_id returns the Thread_id field from the record's packed storage.
func (s *Thread_snapshot) Thread_id() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[28:36]))
}

// SetThread_id updates the Thread_id field in the record's packed storage.
func (s *Thread_snapshot) SetThread_id(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[28:36], uint64(v))
}

// User_time returns the User_time field from the record's packed storage.
func (s *Thread_snapshot) User_time() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[36:44]))
}

// SetUser_time updates the User_time field in the record's packed storage.
func (s *Thread_snapshot) SetUser_time(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[36:44], uint64(v))
}

// System_time returns the System_time field from the record's packed storage.
func (s *Thread_snapshot) System_time() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[44:52]))
}

// SetSystem_time updates the System_time field in the record's packed storage.
func (s *Thread_snapshot) SetSystem_time(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[44:52], uint64(v))
}

// State returns the State field from the record's packed storage.
func (s *Thread_snapshot) State() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[52:56]))
}

// SetState updates the State field in the record's packed storage.
func (s *Thread_snapshot) SetState(v int32) {
	binary.NativeEndian.PutUint32(s.storage[52:56], uint32(v))
}

// Priority returns the Priority field from the record's packed storage.
func (s *Thread_snapshot) Priority() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[56:60]))
}

// SetPriority updates the Priority field in the record's packed storage.
func (s *Thread_snapshot) SetPriority(v int32) {
	binary.NativeEndian.PutUint32(s.storage[56:60], uint32(v))
}

// Sched_pri returns the Sched_pri field from the record's packed storage.
func (s *Thread_snapshot) Sched_pri() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[60:64]))
}

// SetSched_pri updates the Sched_pri field in the record's packed storage.
func (s *Thread_snapshot) SetSched_pri(v int32) {
	binary.NativeEndian.PutUint32(s.storage[60:64], uint32(v))
}

// Sched_flags returns the Sched_flags field from the record's packed storage.
func (s *Thread_snapshot) Sched_flags() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[64:68]))
}

// SetSched_flags updates the Sched_flags field in the record's packed storage.
func (s *Thread_snapshot) SetSched_flags(v int32) {
	binary.NativeEndian.PutUint32(s.storage[64:68], uint32(v))
}

// Ss_flags returns the Ss_flags field from the record's packed storage.
func (s *Thread_snapshot) Ss_flags() int8 {
	return int8(s.storage[68])
}

// SetSs_flags updates the Ss_flags field in the record's packed storage.
func (s *Thread_snapshot) SetSs_flags(v int8) {
	s.storage[68] = uint8(v)
}

// Ts_qos returns the Ts_qos field from the record's packed storage.
func (s *Thread_snapshot) Ts_qos() int8 {
	return int8(s.storage[69])
}

// SetTs_qos updates the Ts_qos field in the record's packed storage.
func (s *Thread_snapshot) SetTs_qos(v int8) {
	s.storage[69] = uint8(v)
}

// Ts_rqos returns the Ts_rqos field from the record's packed storage.
func (s *Thread_snapshot) Ts_rqos() int8 {
	return int8(s.storage[70])
}

// SetTs_rqos updates the Ts_rqos field in the record's packed storage.
func (s *Thread_snapshot) SetTs_rqos(v int8) {
	s.storage[70] = uint8(v)
}

// Ts_rqos_override returns the Ts_rqos_override field from the record's packed storage.
func (s *Thread_snapshot) Ts_rqos_override() int8 {
	return int8(s.storage[71])
}

// SetTs_rqos_override updates the Ts_rqos_override field in the record's packed storage.
func (s *Thread_snapshot) SetTs_rqos_override(v int8) {
	s.storage[71] = uint8(v)
}

// Io_tier returns the Io_tier field from the record's packed storage.
func (s *Thread_snapshot) Io_tier() int8 {
	return int8(s.storage[72])
}

// SetIo_tier updates the Io_tier field in the record's packed storage.
func (s *Thread_snapshot) SetIo_tier(v int8) {
	s.storage[72] = uint8(v)
}

// _reserved returns the _reserved field from the record's packed storage.
func (s *Thread_snapshot) _reserved() [3]int8 {
	return *(*[3]int8)(unsafe.Pointer(&s.storage[73]))
}

// Set_reserved updates the _reserved field in the record's packed storage.
func (s *Thread_snapshot) Set_reserved(v [3]int8) {
	*(*[3]int8)(unsafe.Pointer(&s.storage[73])) = v
}

// Disk_reads_count returns the Disk_reads_count field from the record's packed storage.
func (s *Thread_snapshot) Disk_reads_count() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[76:84]))
}

// SetDisk_reads_count updates the Disk_reads_count field in the record's packed storage.
func (s *Thread_snapshot) SetDisk_reads_count(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[76:84], uint64(v))
}

// Disk_reads_size returns the Disk_reads_size field from the record's packed storage.
func (s *Thread_snapshot) Disk_reads_size() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[84:92]))
}

// SetDisk_reads_size updates the Disk_reads_size field in the record's packed storage.
func (s *Thread_snapshot) SetDisk_reads_size(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[84:92], uint64(v))
}

// Disk_writes_count returns the Disk_writes_count field from the record's packed storage.
func (s *Thread_snapshot) Disk_writes_count() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[92:100]))
}

// SetDisk_writes_count updates the Disk_writes_count field in the record's packed storage.
func (s *Thread_snapshot) SetDisk_writes_count(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[92:100], uint64(v))
}

// Disk_writes_size returns the Disk_writes_size field from the record's packed storage.
func (s *Thread_snapshot) Disk_writes_size() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[100:108]))
}

// SetDisk_writes_size updates the Disk_writes_size field in the record's packed storage.
func (s *Thread_snapshot) SetDisk_writes_size(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[100:108], uint64(v))
}

// Io_priority_count returns the Io_priority_count field from the record's packed storage.
func (s *Thread_snapshot) Io_priority_count() [4]uint64 {
	return *(*[4]uint64)(unsafe.Pointer(&s.storage[108]))
}

// SetIo_priority_count updates the Io_priority_count field in the record's packed storage.
func (s *Thread_snapshot) SetIo_priority_count(v [4]uint64) {
	*(*[4]uint64)(unsafe.Pointer(&s.storage[108])) = v
}

// Io_priority_size returns the Io_priority_size field from the record's packed storage.
func (s *Thread_snapshot) Io_priority_size() [4]uint64 {
	return *(*[4]uint64)(unsafe.Pointer(&s.storage[140]))
}

// SetIo_priority_size updates the Io_priority_size field in the record's packed storage.
func (s *Thread_snapshot) SetIo_priority_size(v [4]uint64) {
	*(*[4]uint64)(unsafe.Pointer(&s.storage[140])) = v
}

// Paging_count returns the Paging_count field from the record's packed storage.
func (s *Thread_snapshot) Paging_count() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[172:180]))
}

// SetPaging_count updates the Paging_count field in the record's packed storage.
func (s *Thread_snapshot) SetPaging_count(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[172:180], uint64(v))
}

// Paging_size returns the Paging_size field from the record's packed storage.
func (s *Thread_snapshot) Paging_size() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[180:188]))
}

// SetPaging_size updates the Paging_size field in the record's packed storage.
func (s *Thread_snapshot) SetPaging_size(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[180:188], uint64(v))
}

// Non_paging_count returns the Non_paging_count field from the record's packed storage.
func (s *Thread_snapshot) Non_paging_count() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[188:196]))
}

// SetNon_paging_count updates the Non_paging_count field in the record's packed storage.
func (s *Thread_snapshot) SetNon_paging_count(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[188:196], uint64(v))
}

// Non_paging_size returns the Non_paging_size field from the record's packed storage.
func (s *Thread_snapshot) Non_paging_size() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[196:204]))
}

// SetNon_paging_size updates the Non_paging_size field in the record's packed storage.
func (s *Thread_snapshot) SetNon_paging_size(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[196:204], uint64(v))
}

// Data_count returns the Data_count field from the record's packed storage.
func (s *Thread_snapshot) Data_count() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[204:212]))
}

// SetData_count updates the Data_count field in the record's packed storage.
func (s *Thread_snapshot) SetData_count(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[204:212], uint64(v))
}

// Data_size returns the Data_size field from the record's packed storage.
func (s *Thread_snapshot) Data_size() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[212:220]))
}

// SetData_size updates the Data_size field in the record's packed storage.
func (s *Thread_snapshot) SetData_size(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[212:220], uint64(v))
}

// Metadata_count returns the Metadata_count field from the record's packed storage.
func (s *Thread_snapshot) Metadata_count() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[220:228]))
}

// SetMetadata_count updates the Metadata_count field in the record's packed storage.
func (s *Thread_snapshot) SetMetadata_count(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[220:228], uint64(v))
}

// Metadata_size returns the Metadata_size field from the record's packed storage.
func (s *Thread_snapshot) Metadata_size() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[228:236]))
}

// SetMetadata_size updates the Metadata_size field in the record's packed storage.
func (s *Thread_snapshot) SetMetadata_size(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[228:236], uint64(v))
}

// Voucher_identifier returns the Voucher_identifier field from the record's packed storage.
func (s *Thread_snapshot) Voucher_identifier() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[236:244]))
}

// SetVoucher_identifier updates the Voucher_identifier field in the record's packed storage.
func (s *Thread_snapshot) SetVoucher_identifier(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[236:244], uint64(v))
}

// Total_syscalls returns the Total_syscalls field from the record's packed storage.
func (s *Thread_snapshot) Total_syscalls() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[244:252]))
}

// SetTotal_syscalls updates the Total_syscalls field in the record's packed storage.
func (s *Thread_snapshot) SetTotal_syscalls(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[244:252], uint64(v))
}

// Pth_name returns the Pth_name field from the record's packed storage.
func (s *Thread_snapshot) Pth_name() [64]int8 {
	return *(*[64]int8)(unsafe.Pointer(&s.storage[252]))
}

// SetPth_name updates the Pth_name field in the record's packed storage.
func (s *Thread_snapshot) SetPth_name(v [64]int8) {
	*(*[64]int8)(unsafe.Pointer(&s.storage[252])) = v
}

// Thread_snapshot_v2
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/thread_snapshot_v2
type Thread_snapshot_v2 struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [104]byte
}

// Ths_thread_id returns the Ths_thread_id field from the record's packed storage.
func (s *Thread_snapshot_v2) Ths_thread_id() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[0:8]))
}

// SetThs_thread_id updates the Ths_thread_id field in the record's packed storage.
func (s *Thread_snapshot_v2) SetThs_thread_id(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[0:8], uint64(v))
}

// Ths_wait_event returns the Ths_wait_event field from the record's packed storage.
func (s *Thread_snapshot_v2) Ths_wait_event() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[8:16]))
}

// SetThs_wait_event updates the Ths_wait_event field in the record's packed storage.
func (s *Thread_snapshot_v2) SetThs_wait_event(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[8:16], uint64(v))
}

// Ths_continuation returns the Ths_continuation field from the record's packed storage.
func (s *Thread_snapshot_v2) Ths_continuation() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[16:24]))
}

// SetThs_continuation updates the Ths_continuation field in the record's packed storage.
func (s *Thread_snapshot_v2) SetThs_continuation(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[16:24], uint64(v))
}

// Ths_total_syscalls returns the Ths_total_syscalls field from the record's packed storage.
func (s *Thread_snapshot_v2) Ths_total_syscalls() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[24:32]))
}

// SetThs_total_syscalls updates the Ths_total_syscalls field in the record's packed storage.
func (s *Thread_snapshot_v2) SetThs_total_syscalls(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[24:32], uint64(v))
}

// Ths_voucher_identifier returns the Ths_voucher_identifier field from the record's packed storage.
func (s *Thread_snapshot_v2) Ths_voucher_identifier() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[32:40]))
}

// SetThs_voucher_identifier updates the Ths_voucher_identifier field in the record's packed storage.
func (s *Thread_snapshot_v2) SetThs_voucher_identifier(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[32:40], uint64(v))
}

// Ths_dqserialnum returns the Ths_dqserialnum field from the record's packed storage.
func (s *Thread_snapshot_v2) Ths_dqserialnum() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[40:48]))
}

// SetThs_dqserialnum updates the Ths_dqserialnum field in the record's packed storage.
func (s *Thread_snapshot_v2) SetThs_dqserialnum(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[40:48], uint64(v))
}

// Ths_user_time returns the Ths_user_time field from the record's packed storage.
func (s *Thread_snapshot_v2) Ths_user_time() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[48:56]))
}

// SetThs_user_time updates the Ths_user_time field in the record's packed storage.
func (s *Thread_snapshot_v2) SetThs_user_time(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[48:56], uint64(v))
}

// Ths_sys_time returns the Ths_sys_time field from the record's packed storage.
func (s *Thread_snapshot_v2) Ths_sys_time() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[56:64]))
}

// SetThs_sys_time updates the Ths_sys_time field in the record's packed storage.
func (s *Thread_snapshot_v2) SetThs_sys_time(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[56:64], uint64(v))
}

// Ths_ss_flags returns the Ths_ss_flags field from the record's packed storage.
func (s *Thread_snapshot_v2) Ths_ss_flags() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[64:72]))
}

// SetThs_ss_flags updates the Ths_ss_flags field in the record's packed storage.
func (s *Thread_snapshot_v2) SetThs_ss_flags(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[64:72], uint64(v))
}

// Ths_last_run_time returns the Ths_last_run_time field from the record's packed storage.
func (s *Thread_snapshot_v2) Ths_last_run_time() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[72:80]))
}

// SetThs_last_run_time updates the Ths_last_run_time field in the record's packed storage.
func (s *Thread_snapshot_v2) SetThs_last_run_time(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[72:80], uint64(v))
}

// Ths_last_made_runnable_time returns the Ths_last_made_runnable_time field from the record's packed storage.
func (s *Thread_snapshot_v2) Ths_last_made_runnable_time() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[80:88]))
}

// SetThs_last_made_runnable_time updates the Ths_last_made_runnable_time field in the record's packed storage.
func (s *Thread_snapshot_v2) SetThs_last_made_runnable_time(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[80:88], uint64(v))
}

// Ths_state returns the Ths_state field from the record's packed storage.
func (s *Thread_snapshot_v2) Ths_state() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[88:92]))
}

// SetThs_state updates the Ths_state field in the record's packed storage.
func (s *Thread_snapshot_v2) SetThs_state(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[88:92], uint32(v))
}

// Ths_sched_flags returns the Ths_sched_flags field from the record's packed storage.
func (s *Thread_snapshot_v2) Ths_sched_flags() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[92:96]))
}

// SetThs_sched_flags updates the Ths_sched_flags field in the record's packed storage.
func (s *Thread_snapshot_v2) SetThs_sched_flags(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[92:96], uint32(v))
}

// Ths_base_priority returns the Ths_base_priority field from the record's packed storage.
func (s *Thread_snapshot_v2) Ths_base_priority() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[96:98]))
}

// SetThs_base_priority updates the Ths_base_priority field in the record's packed storage.
func (s *Thread_snapshot_v2) SetThs_base_priority(v int16) {
	binary.NativeEndian.PutUint16(s.storage[96:98], uint16(v))
}

// Ths_sched_priority returns the Ths_sched_priority field from the record's packed storage.
func (s *Thread_snapshot_v2) Ths_sched_priority() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[98:100]))
}

// SetThs_sched_priority updates the Ths_sched_priority field in the record's packed storage.
func (s *Thread_snapshot_v2) SetThs_sched_priority(v int16) {
	binary.NativeEndian.PutUint16(s.storage[98:100], uint16(v))
}

// Ths_eqos returns the Ths_eqos field from the record's packed storage.
func (s *Thread_snapshot_v2) Ths_eqos() uint8 {
	return uint8(s.storage[100])
}

// SetThs_eqos updates the Ths_eqos field in the record's packed storage.
func (s *Thread_snapshot_v2) SetThs_eqos(v uint8) {
	s.storage[100] = uint8(v)
}

// Ths_rqos returns the Ths_rqos field from the record's packed storage.
func (s *Thread_snapshot_v2) Ths_rqos() uint8 {
	return uint8(s.storage[101])
}

// SetThs_rqos updates the Ths_rqos field in the record's packed storage.
func (s *Thread_snapshot_v2) SetThs_rqos(v uint8) {
	s.storage[101] = uint8(v)
}

// Ths_rqos_override returns the Ths_rqos_override field from the record's packed storage.
func (s *Thread_snapshot_v2) Ths_rqos_override() uint8 {
	return uint8(s.storage[102])
}

// SetThs_rqos_override updates the Ths_rqos_override field in the record's packed storage.
func (s *Thread_snapshot_v2) SetThs_rqos_override(v uint8) {
	s.storage[102] = uint8(v)
}

// Ths_io_tier returns the Ths_io_tier field from the record's packed storage.
func (s *Thread_snapshot_v2) Ths_io_tier() uint8 {
	return uint8(s.storage[103])
}

// SetThs_io_tier updates the Ths_io_tier field in the record's packed storage.
func (s *Thread_snapshot_v2) SetThs_io_tier(v uint8) {
	s.storage[103] = uint8(v)
}

// Thread_snapshot_v3
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/thread_snapshot_v3
type Thread_snapshot_v3 struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [112]byte
}

// Ths_thread_id returns the Ths_thread_id field from the record's packed storage.
func (s *Thread_snapshot_v3) Ths_thread_id() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[0:8]))
}

// SetThs_thread_id updates the Ths_thread_id field in the record's packed storage.
func (s *Thread_snapshot_v3) SetThs_thread_id(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[0:8], uint64(v))
}

// Ths_wait_event returns the Ths_wait_event field from the record's packed storage.
func (s *Thread_snapshot_v3) Ths_wait_event() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[8:16]))
}

// SetThs_wait_event updates the Ths_wait_event field in the record's packed storage.
func (s *Thread_snapshot_v3) SetThs_wait_event(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[8:16], uint64(v))
}

// Ths_continuation returns the Ths_continuation field from the record's packed storage.
func (s *Thread_snapshot_v3) Ths_continuation() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[16:24]))
}

// SetThs_continuation updates the Ths_continuation field in the record's packed storage.
func (s *Thread_snapshot_v3) SetThs_continuation(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[16:24], uint64(v))
}

// Ths_total_syscalls returns the Ths_total_syscalls field from the record's packed storage.
func (s *Thread_snapshot_v3) Ths_total_syscalls() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[24:32]))
}

// SetThs_total_syscalls updates the Ths_total_syscalls field in the record's packed storage.
func (s *Thread_snapshot_v3) SetThs_total_syscalls(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[24:32], uint64(v))
}

// Ths_voucher_identifier returns the Ths_voucher_identifier field from the record's packed storage.
func (s *Thread_snapshot_v3) Ths_voucher_identifier() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[32:40]))
}

// SetThs_voucher_identifier updates the Ths_voucher_identifier field in the record's packed storage.
func (s *Thread_snapshot_v3) SetThs_voucher_identifier(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[32:40], uint64(v))
}

// Ths_dqserialnum returns the Ths_dqserialnum field from the record's packed storage.
func (s *Thread_snapshot_v3) Ths_dqserialnum() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[40:48]))
}

// SetThs_dqserialnum updates the Ths_dqserialnum field in the record's packed storage.
func (s *Thread_snapshot_v3) SetThs_dqserialnum(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[40:48], uint64(v))
}

// Ths_user_time returns the Ths_user_time field from the record's packed storage.
func (s *Thread_snapshot_v3) Ths_user_time() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[48:56]))
}

// SetThs_user_time updates the Ths_user_time field in the record's packed storage.
func (s *Thread_snapshot_v3) SetThs_user_time(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[48:56], uint64(v))
}

// Ths_sys_time returns the Ths_sys_time field from the record's packed storage.
func (s *Thread_snapshot_v3) Ths_sys_time() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[56:64]))
}

// SetThs_sys_time updates the Ths_sys_time field in the record's packed storage.
func (s *Thread_snapshot_v3) SetThs_sys_time(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[56:64], uint64(v))
}

// Ths_ss_flags returns the Ths_ss_flags field from the record's packed storage.
func (s *Thread_snapshot_v3) Ths_ss_flags() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[64:72]))
}

// SetThs_ss_flags updates the Ths_ss_flags field in the record's packed storage.
func (s *Thread_snapshot_v3) SetThs_ss_flags(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[64:72], uint64(v))
}

// Ths_last_run_time returns the Ths_last_run_time field from the record's packed storage.
func (s *Thread_snapshot_v3) Ths_last_run_time() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[72:80]))
}

// SetThs_last_run_time updates the Ths_last_run_time field in the record's packed storage.
func (s *Thread_snapshot_v3) SetThs_last_run_time(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[72:80], uint64(v))
}

// Ths_last_made_runnable_time returns the Ths_last_made_runnable_time field from the record's packed storage.
func (s *Thread_snapshot_v3) Ths_last_made_runnable_time() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[80:88]))
}

// SetThs_last_made_runnable_time updates the Ths_last_made_runnable_time field in the record's packed storage.
func (s *Thread_snapshot_v3) SetThs_last_made_runnable_time(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[80:88], uint64(v))
}

// Ths_state returns the Ths_state field from the record's packed storage.
func (s *Thread_snapshot_v3) Ths_state() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[88:92]))
}

// SetThs_state updates the Ths_state field in the record's packed storage.
func (s *Thread_snapshot_v3) SetThs_state(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[88:92], uint32(v))
}

// Ths_sched_flags returns the Ths_sched_flags field from the record's packed storage.
func (s *Thread_snapshot_v3) Ths_sched_flags() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[92:96]))
}

// SetThs_sched_flags updates the Ths_sched_flags field in the record's packed storage.
func (s *Thread_snapshot_v3) SetThs_sched_flags(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[92:96], uint32(v))
}

// Ths_base_priority returns the Ths_base_priority field from the record's packed storage.
func (s *Thread_snapshot_v3) Ths_base_priority() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[96:98]))
}

// SetThs_base_priority updates the Ths_base_priority field in the record's packed storage.
func (s *Thread_snapshot_v3) SetThs_base_priority(v int16) {
	binary.NativeEndian.PutUint16(s.storage[96:98], uint16(v))
}

// Ths_sched_priority returns the Ths_sched_priority field from the record's packed storage.
func (s *Thread_snapshot_v3) Ths_sched_priority() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[98:100]))
}

// SetThs_sched_priority updates the Ths_sched_priority field in the record's packed storage.
func (s *Thread_snapshot_v3) SetThs_sched_priority(v int16) {
	binary.NativeEndian.PutUint16(s.storage[98:100], uint16(v))
}

// Ths_eqos returns the Ths_eqos field from the record's packed storage.
func (s *Thread_snapshot_v3) Ths_eqos() uint8 {
	return uint8(s.storage[100])
}

// SetThs_eqos updates the Ths_eqos field in the record's packed storage.
func (s *Thread_snapshot_v3) SetThs_eqos(v uint8) {
	s.storage[100] = uint8(v)
}

// Ths_rqos returns the Ths_rqos field from the record's packed storage.
func (s *Thread_snapshot_v3) Ths_rqos() uint8 {
	return uint8(s.storage[101])
}

// SetThs_rqos updates the Ths_rqos field in the record's packed storage.
func (s *Thread_snapshot_v3) SetThs_rqos(v uint8) {
	s.storage[101] = uint8(v)
}

// Ths_rqos_override returns the Ths_rqos_override field from the record's packed storage.
func (s *Thread_snapshot_v3) Ths_rqos_override() uint8 {
	return uint8(s.storage[102])
}

// SetThs_rqos_override updates the Ths_rqos_override field in the record's packed storage.
func (s *Thread_snapshot_v3) SetThs_rqos_override(v uint8) {
	s.storage[102] = uint8(v)
}

// Ths_io_tier returns the Ths_io_tier field from the record's packed storage.
func (s *Thread_snapshot_v3) Ths_io_tier() uint8 {
	return uint8(s.storage[103])
}

// SetThs_io_tier updates the Ths_io_tier field in the record's packed storage.
func (s *Thread_snapshot_v3) SetThs_io_tier(v uint8) {
	s.storage[103] = uint8(v)
}

// Ths_thread_t returns the Ths_thread_t field from the record's packed storage.
func (s *Thread_snapshot_v3) Ths_thread_t() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[104:112]))
}

// SetThs_thread_t updates the Ths_thread_t field in the record's packed storage.
func (s *Thread_snapshot_v3) SetThs_thread_t(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[104:112], uint64(v))
}

// Thread_snapshot_v4
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/thread_snapshot_v4
type Thread_snapshot_v4 struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [128]byte
}

// Ths_thread_id returns the Ths_thread_id field from the record's packed storage.
func (s *Thread_snapshot_v4) Ths_thread_id() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[0:8]))
}

// SetThs_thread_id updates the Ths_thread_id field in the record's packed storage.
func (s *Thread_snapshot_v4) SetThs_thread_id(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[0:8], uint64(v))
}

// Ths_wait_event returns the Ths_wait_event field from the record's packed storage.
func (s *Thread_snapshot_v4) Ths_wait_event() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[8:16]))
}

// SetThs_wait_event updates the Ths_wait_event field in the record's packed storage.
func (s *Thread_snapshot_v4) SetThs_wait_event(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[8:16], uint64(v))
}

// Ths_continuation returns the Ths_continuation field from the record's packed storage.
func (s *Thread_snapshot_v4) Ths_continuation() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[16:24]))
}

// SetThs_continuation updates the Ths_continuation field in the record's packed storage.
func (s *Thread_snapshot_v4) SetThs_continuation(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[16:24], uint64(v))
}

// Ths_total_syscalls returns the Ths_total_syscalls field from the record's packed storage.
func (s *Thread_snapshot_v4) Ths_total_syscalls() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[24:32]))
}

// SetThs_total_syscalls updates the Ths_total_syscalls field in the record's packed storage.
func (s *Thread_snapshot_v4) SetThs_total_syscalls(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[24:32], uint64(v))
}

// Ths_voucher_identifier returns the Ths_voucher_identifier field from the record's packed storage.
func (s *Thread_snapshot_v4) Ths_voucher_identifier() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[32:40]))
}

// SetThs_voucher_identifier updates the Ths_voucher_identifier field in the record's packed storage.
func (s *Thread_snapshot_v4) SetThs_voucher_identifier(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[32:40], uint64(v))
}

// Ths_dqserialnum returns the Ths_dqserialnum field from the record's packed storage.
func (s *Thread_snapshot_v4) Ths_dqserialnum() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[40:48]))
}

// SetThs_dqserialnum updates the Ths_dqserialnum field in the record's packed storage.
func (s *Thread_snapshot_v4) SetThs_dqserialnum(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[40:48], uint64(v))
}

// Ths_user_time returns the Ths_user_time field from the record's packed storage.
func (s *Thread_snapshot_v4) Ths_user_time() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[48:56]))
}

// SetThs_user_time updates the Ths_user_time field in the record's packed storage.
func (s *Thread_snapshot_v4) SetThs_user_time(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[48:56], uint64(v))
}

// Ths_sys_time returns the Ths_sys_time field from the record's packed storage.
func (s *Thread_snapshot_v4) Ths_sys_time() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[56:64]))
}

// SetThs_sys_time updates the Ths_sys_time field in the record's packed storage.
func (s *Thread_snapshot_v4) SetThs_sys_time(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[56:64], uint64(v))
}

// Ths_ss_flags returns the Ths_ss_flags field from the record's packed storage.
func (s *Thread_snapshot_v4) Ths_ss_flags() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[64:72]))
}

// SetThs_ss_flags updates the Ths_ss_flags field in the record's packed storage.
func (s *Thread_snapshot_v4) SetThs_ss_flags(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[64:72], uint64(v))
}

// Ths_last_run_time returns the Ths_last_run_time field from the record's packed storage.
func (s *Thread_snapshot_v4) Ths_last_run_time() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[72:80]))
}

// SetThs_last_run_time updates the Ths_last_run_time field in the record's packed storage.
func (s *Thread_snapshot_v4) SetThs_last_run_time(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[72:80], uint64(v))
}

// Ths_last_made_runnable_time returns the Ths_last_made_runnable_time field from the record's packed storage.
func (s *Thread_snapshot_v4) Ths_last_made_runnable_time() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[80:88]))
}

// SetThs_last_made_runnable_time updates the Ths_last_made_runnable_time field in the record's packed storage.
func (s *Thread_snapshot_v4) SetThs_last_made_runnable_time(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[80:88], uint64(v))
}

// Ths_state returns the Ths_state field from the record's packed storage.
func (s *Thread_snapshot_v4) Ths_state() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[88:92]))
}

// SetThs_state updates the Ths_state field in the record's packed storage.
func (s *Thread_snapshot_v4) SetThs_state(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[88:92], uint32(v))
}

// Ths_sched_flags returns the Ths_sched_flags field from the record's packed storage.
func (s *Thread_snapshot_v4) Ths_sched_flags() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[92:96]))
}

// SetThs_sched_flags updates the Ths_sched_flags field in the record's packed storage.
func (s *Thread_snapshot_v4) SetThs_sched_flags(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[92:96], uint32(v))
}

// Ths_base_priority returns the Ths_base_priority field from the record's packed storage.
func (s *Thread_snapshot_v4) Ths_base_priority() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[96:98]))
}

// SetThs_base_priority updates the Ths_base_priority field in the record's packed storage.
func (s *Thread_snapshot_v4) SetThs_base_priority(v int16) {
	binary.NativeEndian.PutUint16(s.storage[96:98], uint16(v))
}

// Ths_sched_priority returns the Ths_sched_priority field from the record's packed storage.
func (s *Thread_snapshot_v4) Ths_sched_priority() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[98:100]))
}

// SetThs_sched_priority updates the Ths_sched_priority field in the record's packed storage.
func (s *Thread_snapshot_v4) SetThs_sched_priority(v int16) {
	binary.NativeEndian.PutUint16(s.storage[98:100], uint16(v))
}

// Ths_eqos returns the Ths_eqos field from the record's packed storage.
func (s *Thread_snapshot_v4) Ths_eqos() uint8 {
	return uint8(s.storage[100])
}

// SetThs_eqos updates the Ths_eqos field in the record's packed storage.
func (s *Thread_snapshot_v4) SetThs_eqos(v uint8) {
	s.storage[100] = uint8(v)
}

// Ths_rqos returns the Ths_rqos field from the record's packed storage.
func (s *Thread_snapshot_v4) Ths_rqos() uint8 {
	return uint8(s.storage[101])
}

// SetThs_rqos updates the Ths_rqos field in the record's packed storage.
func (s *Thread_snapshot_v4) SetThs_rqos(v uint8) {
	s.storage[101] = uint8(v)
}

// Ths_rqos_override returns the Ths_rqos_override field from the record's packed storage.
func (s *Thread_snapshot_v4) Ths_rqos_override() uint8 {
	return uint8(s.storage[102])
}

// SetThs_rqos_override updates the Ths_rqos_override field in the record's packed storage.
func (s *Thread_snapshot_v4) SetThs_rqos_override(v uint8) {
	s.storage[102] = uint8(v)
}

// Ths_io_tier returns the Ths_io_tier field from the record's packed storage.
func (s *Thread_snapshot_v4) Ths_io_tier() uint8 {
	return uint8(s.storage[103])
}

// SetThs_io_tier updates the Ths_io_tier field in the record's packed storage.
func (s *Thread_snapshot_v4) SetThs_io_tier(v uint8) {
	s.storage[103] = uint8(v)
}

// Ths_thread_t returns the Ths_thread_t field from the record's packed storage.
func (s *Thread_snapshot_v4) Ths_thread_t() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[104:112]))
}

// SetThs_thread_t updates the Ths_thread_t field in the record's packed storage.
func (s *Thread_snapshot_v4) SetThs_thread_t(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[104:112], uint64(v))
}

// Ths_requested_policy returns the Ths_requested_policy field from the record's packed storage.
func (s *Thread_snapshot_v4) Ths_requested_policy() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[112:120]))
}

// SetThs_requested_policy updates the Ths_requested_policy field in the record's packed storage.
func (s *Thread_snapshot_v4) SetThs_requested_policy(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[112:120], uint64(v))
}

// Ths_effective_policy returns the Ths_effective_policy field from the record's packed storage.
func (s *Thread_snapshot_v4) Ths_effective_policy() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[120:128]))
}

// SetThs_effective_policy updates the Ths_effective_policy field in the record's packed storage.
func (s *Thread_snapshot_v4) SetThs_effective_policy(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[120:128], uint64(v))
}

// Thsc_cpi
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/thsc_cpi
type Thsc_cpi struct {
	Tcpi_instructions uint64
	Tcpi_cycles       uint64
}

// Thsc_time_cpi
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/thsc_time_cpi
type Thsc_time_cpi struct {
	Ttci_instructions     uint64
	Ttci_cycles           uint64
	Ttci_user_time_mach   uint64
	Ttci_system_time_mach uint64
}

// Thsc_time_energy_cpi
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/thsc_time_energy_cpi
type Thsc_time_energy_cpi struct {
	Ttec_instructions     uint64
	Ttec_cycles           uint64
	Ttec_user_time_mach   uint64
	Ttec_system_time_mach uint64
	Ttec_energy_nj        uint64
}

// Timebase_freq_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/timebase_freq_t
type Timebase_freq_t struct {
	Timebase_num uint
	Timebase_den uint
}

// Timespec
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/timespec
type Timespec struct {
	Tv_sec  int
	Tv_nsec int
}

// Timeval
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/timeval
type Timeval struct {
	Tv_sec  int
	Tv_usec int32
}

// Timeval32
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/timeval32
type Timeval32 struct {
	Tv_sec  int32
	Tv_usec int32
}

// Timeval64
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/timeval64
type Timeval64 struct {
	Tv_sec  int64
	Tv_usec int64
}

// Timex
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/timex
type Timex struct {
	Modes     uint32
	Offset    int
	Freq      int
	Maxerror  int
	Esterror  int
	Status    int32
	Constant  int
	Precision int
	Tolerance int
	Ppsfreq   int
	Jitter    int
	Shift     int32
	Stabil    int
	Jitcnt    int
	Calcnt    int
	Errcnt    int
	Stbcnt    int
}

// Timezone
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/timezone
type Timezone struct {
	Tz_minuteswest int32
	Tz_dsttime     int32
}

// Tlv_descriptor
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/tlv_descriptor
type Tlv_descriptor struct {
	Thunk  *objc.ID
	Key    uint
	Offset uint
}

// Transitioning_task_snapshot
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/transitioning_task_snapshot
type Transitioning_task_snapshot struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [60]byte
}

// Tts_unique_pid returns the Tts_unique_pid field from the record's packed storage.
func (s *Transitioning_task_snapshot) Tts_unique_pid() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[0:8]))
}

// SetTts_unique_pid updates the Tts_unique_pid field in the record's packed storage.
func (s *Transitioning_task_snapshot) SetTts_unique_pid(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[0:8], uint64(v))
}

// Tts_ss_flags returns the Tts_ss_flags field from the record's packed storage.
func (s *Transitioning_task_snapshot) Tts_ss_flags() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[8:16]))
}

// SetTts_ss_flags updates the Tts_ss_flags field in the record's packed storage.
func (s *Transitioning_task_snapshot) SetTts_ss_flags(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[8:16], uint64(v))
}

// Tts_transition_type returns the Tts_transition_type field from the record's packed storage.
func (s *Transitioning_task_snapshot) Tts_transition_type() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[16:24]))
}

// SetTts_transition_type updates the Tts_transition_type field in the record's packed storage.
func (s *Transitioning_task_snapshot) SetTts_transition_type(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[16:24], uint64(v))
}

// Tts_pid returns the Tts_pid field from the record's packed storage.
func (s *Transitioning_task_snapshot) Tts_pid() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[24:28]))
}

// SetTts_pid updates the Tts_pid field in the record's packed storage.
func (s *Transitioning_task_snapshot) SetTts_pid(v int32) {
	binary.NativeEndian.PutUint32(s.storage[24:28], uint32(v))
}

// Tts_p_comm returns the Tts_p_comm field from the record's packed storage.
func (s *Transitioning_task_snapshot) Tts_p_comm() [32]int8 {
	return *(*[32]int8)(unsafe.Pointer(&s.storage[28]))
}

// SetTts_p_comm updates the Tts_p_comm field in the record's packed storage.
func (s *Transitioning_task_snapshot) SetTts_p_comm(v [32]int8) {
	*(*[32]int8)(unsafe.Pointer(&s.storage[28])) = v
}

// Trust_cache_entry1
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/trust_cache_entry1
type Trust_cache_entry1 struct {
	Cdhash    [20]uint8
	Hash_type uint8
	Flags     uint8
}

// Trust_cache_module1
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/trust_cache_module1
type Trust_cache_module1 struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [24]byte
}

// Version returns the Version field from the record's packed storage.
func (s *Trust_cache_module1) Version() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetVersion updates the Version field in the record's packed storage.
func (s *Trust_cache_module1) SetVersion(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// Uuid returns the Uuid field from the record's packed storage.
func (s *Trust_cache_module1) Uuid() [16]uint8 {
	return *(*[16]uint8)(unsafe.Pointer(&s.storage[4]))
}

// SetUuid updates the Uuid field in the record's packed storage.
func (s *Trust_cache_module1) SetUuid(v [16]uint8) {
	*(*[16]uint8)(unsafe.Pointer(&s.storage[4])) = v
}

// Num_entries returns the Num_entries field from the record's packed storage.
func (s *Trust_cache_module1) Num_entries() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[20:24]))
}

// SetNum_entries updates the Num_entries field in the record's packed storage.
func (s *Trust_cache_module1) SetNum_entries(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[20:24], uint32(v))
}

// Tsegqe_head
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/tsegqe_head
type Tsegqe_head struct {
	Lh_first U_int32_t
}

// Ttysize
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ttysize
type Ttysize struct {
	Ts_lines uint16
	Ts_cols  uint16
	Ts_xxx   uint16
	Ts_yyy   uint16
}

// Twolevel_hint
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/twolevel_hint
type Twolevel_hint struct {
	bitfield0 uint32
}

// Isub_image returns the Isub_image bitfield.
func (s *Twolevel_hint) Isub_image() uint32 {
	return (s.bitfield0 >> 0) & ((1 << 8) - 1)
}

// SetIsub_image updates the Isub_image bitfield.
func (s *Twolevel_hint) SetIsub_image(v uint32) {
	const mask uint32 = (1 << 8) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 0)) | ((v & mask) << 0)
}

// Itoc returns the Itoc bitfield.
func (s *Twolevel_hint) Itoc() uint32 {
	return (s.bitfield0 >> 8) & ((1 << 24) - 1)
}

// SetItoc updates the Itoc bitfield.
func (s *Twolevel_hint) SetItoc(v uint32) {
	const mask uint32 = (1 << 24) - 1
	s.bitfield0 = (s.bitfield0 &^ (mask << 8)) | ((v & mask) << 8)
}

// Twolevel_hints_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/twolevel_hints_command
type Twolevel_hints_command struct {
	Cmd     uint32 // Common to all load command structures. Set to `LC_TWOLEVEL_HINTS` for this structure.
	Cmdsize uint32
	Offset  uint32
	Nhints  uint32
}

// Udphdr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/udphdr
type Udphdr struct {
	Uh_sport U_short
	Uh_dport U_short
	Uh_ulen  U_short
	Uh_sum   U_short
}

// Udpiphdr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/udpiphdr
type Udpiphdr struct {
	Ui_i Ipovly
	Ui_u Udphdr
}

// Udpstat
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/udpstat
type Udpstat struct {
	Udps_ipackets                    U_int32_t
	Udps_hdrops                      U_int32_t
	Udps_badsum                      U_int32_t
	Udps_badlen                      U_int32_t
	Udps_noport                      U_int32_t
	Udps_noportbcast                 U_int32_t
	Udps_fullsock                    U_int32_t
	Udpps_pcbcachemiss               U_int32_t
	Udpps_pcbhashmiss                U_int32_t
	Udps_opackets                    U_int32_t
	Udps_fastout                     U_int32_t
	Udps_nosum                       U_int32_t
	Udps_noportmcast                 U_int32_t
	Udps_filtermcast                 U_int32_t
	Udps_rcv_swcsum                  U_int32_t
	Udps_rcv_swcsum_bytes            U_int32_t
	Udps_rcv6_swcsum                 U_int32_t
	Udps_rcv6_swcsum_bytes           U_int32_t
	Udps_snd_swcsum                  U_int32_t
	Udps_snd_swcsum_bytes            U_int32_t
	Udps_snd6_swcsum                 U_int32_t
	Udps_snd6_swcsum_bytes           U_int32_t
	Udps_port_unreach_dup_suppressed U_int64_t
	Udps_port_unreach_not_suppressed U_int64_t
}

// User32_dyld_uuid_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/user32_dyld_uuid_info
type User32_dyld_uuid_info struct {
	ImageLoadAddress uint32
	ImageUUID        [16]uint8
}

// User32_msqid_ds
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/user32_msqid_ds
type User32_msqid_ds struct {
	Msg_perm   Ipc_perm
	Msg_first  int32
	Msg_last   int32
	Msg_cbytes User32_msglen_t
	Msg_qnum   User32_msgqnum_t
	Msg_qbytes User32_msglen_t
	Msg_lspid  int32
	Msg_lrpid  int32
	Msg_stime  User32_time_t
	Msg_pad1   int32
	Msg_rtime  User32_time_t
	Msg_pad2   int32
	Msg_ctime  User32_time_t
	Msg_pad3   int32
	Msg_pad4   [4]int32
}

// User64_dyld_aot_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/user64_dyld_aot_info
type User64_dyld_aot_info struct {
	X86LoadAddress uint64
	AotLoadAddress uint64
	AotImageSize   uint64
	AotImageKey    [32]uint8
}

// User64_dyld_uuid_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/user64_dyld_uuid_info
type User64_dyld_uuid_info struct {
	ImageLoadAddress uint64
	ImageUUID        [16]uint8
}

// User64_msqid_ds
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/user64_msqid_ds
type User64_msqid_ds struct {
	Msg_perm   Ipc_perm
	Msg_first  int32
	Msg_last   int32
	Msg_cbytes User64_msglen_t
	Msg_qnum   User64_msgqnum_t
	Msg_qbytes User64_msglen_t
	Msg_lspid  int32
	Msg_lrpid  int32
	Msg_stime  User64_time_t
	Msg_pad1   int32
	Msg_rtime  User64_time_t
	Msg_pad2   int32
	Msg_ctime  User64_time_t
	Msg_pad3   int32
	Msg_pad4   [4]int32
}

// User_msqid_ds
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/user_msqid_ds
type User_msqid_ds struct {
	Msg_perm   Ipc_perm
	Msg_first  *Msg
	Msg_last   *Msg
	Msg_cbytes User_msglen_t
	Msg_qnum   User_msgqnum_t
	Msg_qbytes User_msglen_t
	Msg_lspid  int32
	Msg_lrpid  int32
	Msg_stime  User_time_t
	Msg_pad1   int32
	Msg_rtime  User_time_t
	Msg_pad2   int32
	Msg_ctime  User_time_t
	Msg_pad3   int32
	Msg_pad4   [4]int32
}

// User_nfs_args
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/user_nfs_args
type User_nfs_args struct {
	Version      int32
	Addr         User_addr_t
	Addrlen      uint8
	Sotype       int32
	Proto        int32
	Fh           User_addr_t
	Fhsize       int32
	Flags        int32
	Wsize        int32
	Rsize        int32
	Readdirsize  int32
	Timeo        int32
	Retrans      int32
	Maxgrouplist int32
	Readahead    int32
	Leaseterm    int32
	Deadthresh   int32
	Hostname     User_addr_t
	Acregmin     int32
	Acregmax     int32
	Acdirmin     int32
	Acdirmax     int32
	Auth         uint32
	Deadtimeout  uint32
}

// User_nfs_export_args
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/user_nfs_export_args
type User_nfs_export_args struct {
	Nxa_fsid     uint32
	Nxa_expid    uint32
	Nxa_fspath   User_addr_t
	Nxa_exppath  User_addr_t
	Nxa_flags    uint32
	Nxa_netcount uint32
	Nxa_nets     User_addr_t
}

// User_nfsd_args
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/user_nfsd_args
type User_nfsd_args struct {
	Sock    int32
	Name    User_addr_t
	Namelen int32
}

// User_termios
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/user_termios
type User_termios struct {
	C_iflag  User_tcflag_t
	C_oflag  User_tcflag_t
	C_cflag  User_tcflag_t
	C_lflag  User_tcflag_t
	C_cc     [20]Cc_t
	C_ispeed uint64
	C_ospeed uint64
}

// Utun_stats_param
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/utun_stats_param
type Utun_stats_param struct {
	Utsp_packets U_int64_t
	Utsp_bytes   U_int64_t
	Utsp_errors  U_int64_t
}

// Uuid_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/uuid_command
type Uuid_command struct {
	Cmd     uint32
	Cmdsize uint32
	Uuid    [16]uint8
}

// Vend
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vend
type Vend struct {
	V_magic  [4]U_char
	V_flags  U_int32_t
	V_unused [56]U_char
}

// Version_min_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/version_min_command
type Version_min_command struct {
	Cmd     uint32
	Cmdsize uint32
	Version uint32
	Sdk     uint32
}

// Vfs_server
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vfs_server
type Vfs_server struct {
	Vs_minutes     int32
	Vs_server_name [768]U_int8_t
}

// Vfsconf
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vfsconf
type Vfsconf struct {
	Vfc_reserved1 uint32
	Vfc_name      [15]int8
	Vfc_typenum   int32
	Vfc_refcount  int32
	Vfc_flags     int32
	Vfc_reserved2 uint32
	Vfc_reserved3 uint32
}

// Vfsidctl
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vfsidctl
type Vfsidctl struct {
	Vc_vers  int32
	Vc_fsid  Fsid_t
	Vc_ptr   unsafe.Pointer
	Vc_len   uintptr
	Vc_spare [12]U_int32_t
}

// Vfsquery
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vfsquery
type Vfsquery struct {
	Vq_flags U_int32_t
	Vq_spare [31]U_int32_t
}

// Vfsstatfs
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vfsstatfs
type Vfsstatfs struct {
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
	storage [2164]byte
}

// F_bsize returns the F_bsize field from the record's packed storage.
func (s *Vfsstatfs) F_bsize() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetF_bsize updates the F_bsize field in the record's packed storage.
func (s *Vfsstatfs) SetF_bsize(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// F_iosize returns the F_iosize field from the record's packed storage.
func (s *Vfsstatfs) F_iosize() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[4:12]))
}

// SetF_iosize updates the F_iosize field in the record's packed storage.
func (s *Vfsstatfs) SetF_iosize(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[4:12], uint64(v))
}

// F_blocks returns the F_blocks field from the record's packed storage.
func (s *Vfsstatfs) F_blocks() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[12:20]))
}

// SetF_blocks updates the F_blocks field in the record's packed storage.
func (s *Vfsstatfs) SetF_blocks(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[12:20], uint64(v))
}

// F_bfree returns the F_bfree field from the record's packed storage.
func (s *Vfsstatfs) F_bfree() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[20:28]))
}

// SetF_bfree updates the F_bfree field in the record's packed storage.
func (s *Vfsstatfs) SetF_bfree(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[20:28], uint64(v))
}

// F_bavail returns the F_bavail field from the record's packed storage.
func (s *Vfsstatfs) F_bavail() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[28:36]))
}

// SetF_bavail updates the F_bavail field in the record's packed storage.
func (s *Vfsstatfs) SetF_bavail(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[28:36], uint64(v))
}

// F_bused returns the F_bused field from the record's packed storage.
func (s *Vfsstatfs) F_bused() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[36:44]))
}

// SetF_bused updates the F_bused field in the record's packed storage.
func (s *Vfsstatfs) SetF_bused(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[36:44], uint64(v))
}

// F_files returns the F_files field from the record's packed storage.
func (s *Vfsstatfs) F_files() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[44:52]))
}

// SetF_files updates the F_files field in the record's packed storage.
func (s *Vfsstatfs) SetF_files(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[44:52], uint64(v))
}

// F_ffree returns the F_ffree field from the record's packed storage.
func (s *Vfsstatfs) F_ffree() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[52:60]))
}

// SetF_ffree updates the F_ffree field in the record's packed storage.
func (s *Vfsstatfs) SetF_ffree(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[52:60], uint64(v))
}

// F_fsid returns the F_fsid field from the record's packed storage.
func (s *Vfsstatfs) F_fsid() Fsid_t {
	return *(*Fsid_t)(unsafe.Pointer(&s.storage[60]))
}

// SetF_fsid updates the F_fsid field in the record's packed storage.
func (s *Vfsstatfs) SetF_fsid(v Fsid_t) {
	*(*Fsid_t)(unsafe.Pointer(&s.storage[60])) = v
}

// F_owner returns the F_owner field from the record's packed storage.
func (s *Vfsstatfs) F_owner() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[68:72]))
}

// SetF_owner updates the F_owner field in the record's packed storage.
func (s *Vfsstatfs) SetF_owner(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[68:72], uint32(v))
}

// F_flags returns the F_flags field from the record's packed storage.
func (s *Vfsstatfs) F_flags() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[72:80]))
}

// SetF_flags updates the F_flags field in the record's packed storage.
func (s *Vfsstatfs) SetF_flags(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[72:80], uint64(v))
}

// F_fstypename returns the F_fstypename field from the record's packed storage.
func (s *Vfsstatfs) F_fstypename() [16]int8 {
	return *(*[16]int8)(unsafe.Pointer(&s.storage[80]))
}

// SetF_fstypename updates the F_fstypename field in the record's packed storage.
func (s *Vfsstatfs) SetF_fstypename(v [16]int8) {
	*(*[16]int8)(unsafe.Pointer(&s.storage[80])) = v
}

// F_mntonname returns the F_mntonname field from the record's packed storage.
func (s *Vfsstatfs) F_mntonname() [1024]int8 {
	return *(*[1024]int8)(unsafe.Pointer(&s.storage[96]))
}

// SetF_mntonname updates the F_mntonname field in the record's packed storage.
func (s *Vfsstatfs) SetF_mntonname(v [1024]int8) {
	*(*[1024]int8)(unsafe.Pointer(&s.storage[96])) = v
}

// F_mntfromname returns the F_mntfromname field from the record's packed storage.
func (s *Vfsstatfs) F_mntfromname() [1024]int8 {
	return *(*[1024]int8)(unsafe.Pointer(&s.storage[1120]))
}

// SetF_mntfromname updates the F_mntfromname field in the record's packed storage.
func (s *Vfsstatfs) SetF_mntfromname(v [1024]int8) {
	*(*[1024]int8)(unsafe.Pointer(&s.storage[1120])) = v
}

// F_fssubtype returns the F_fssubtype field from the record's packed storage.
func (s *Vfsstatfs) F_fssubtype() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[2144:2148]))
}

// SetF_fssubtype updates the F_fssubtype field in the record's packed storage.
func (s *Vfsstatfs) SetF_fssubtype(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[2144:2148], uint32(v))
}

// F_reserved returns the F_reserved field from the record's packed storage.
func (s *Vfsstatfs) F_reserved() [16]byte {
	return *(*[16]byte)(unsafe.Pointer(&s.storage[2148]))
}

// SetF_reserved updates the F_reserved field in the record's packed storage.
func (s *Vfsstatfs) SetF_reserved(v [16]byte) {
	*(*[16]byte)(unsafe.Pointer(&s.storage[2148])) = v
}

// Vlanreq
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vlanreq
type Vlanreq struct {
	Vlr_parent [16]int8
	Vlr_tag    U_short
}

// Vmspace
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vmspace
type Vmspace struct {
	Dummy  int32
	Dummy2 Caddr_t
	Dummy3 [5]int32
	Dummy4 [3]Caddr_t
}

// Vnode_attr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnode_attr
type Vnode_attr struct {
	Va_supported          uint64
	Va_active             uint64
	Va_vaflags            int32
	Va_rdev               int32
	Va_nlink              uint64
	Va_total_size         uint64
	Va_total_alloc        uint64
	Va_data_size          uint64
	Va_data_alloc         uint64
	Va_iosize             uint32
	Va_uid                uint32
	Va_gid                uint32
	Va_mode               uint16
	Va_flags              uint32
	Va_acl                unsafe.Pointer
	Va_create_time        Timespec
	Va_access_time        Timespec
	Va_modify_time        Timespec
	Va_change_time        Timespec
	Va_backup_time        Timespec
	Va_fileid             uint64
	Va_linkid             uint64
	Va_parentid           uint64
	Va_fsid               uint32
	Va_filerev            uint64
	Va_gen                uint32
	Va_encoding           uint32
	Va_type               Vtype
	Va_name               *byte
	Va_uuuid              [4]uint32
	Va_guuid              [4]uint32
	Va_nchildren          uint64
	Va_dirlinkcount       uint64
	Va_reserved1          unsafe.Pointer
	Va_addedtime          Timespec
	Va_dataprotect_class  uint32
	Va_dataprotect_flags  uint32
	Va_document_id        uint32
	Va_devid              uint32
	Va_objtype            uint32
	Va_objtag             uint32
	Va_user_access        uint32
	Va_finderinfo         [32]uint8
	Va_rsrc_length        uint64
	Va_rsrc_alloc         uint64
	Va_fsid64             Fsid_t
	Va_write_gencount     uint32
	Va_private_size       uint64
	Va_clone_id           uint64
	Va_extflags           uint64
	Va_recursive_gencount uint64
	Va_attribution_tag    uint64
	Va_clone_refcnt       uint32
}

// Winsize
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/winsize
type Winsize struct {
	Ws_row    uint16
	Ws_col    uint16
	Ws_xpixel uint16
	Ws_ypixel uint16
}

// Xdrbuf
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/xdrbuf
type Xdrbuf struct {
	Xb_u        [3]uint64
	Xb_ptr      *byte
	Xb_left     uintptr
	Xb_growsize uintptr
	Xb_type     Xdrbuf_type
	Xb_flags    uint32
}

// Xinpcb
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/xinpcb
type Xinpcb struct {
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
	storage [320]byte
}

// Xi_len returns the Xi_len field from the record's packed storage.
func (s *Xinpcb) Xi_len() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetXi_len updates the Xi_len field in the record's packed storage.
func (s *Xinpcb) SetXi_len(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// Xi_inp returns the Xi_inp field from the record's packed storage.
func (s *Xinpcb) Xi_inp() Inpcb {
	return *(*Inpcb)(unsafe.Pointer(&s.storage[4]))
}

// SetXi_inp updates the Xi_inp field in the record's packed storage.
func (s *Xinpcb) SetXi_inp(v Inpcb) {
	*(*Inpcb)(unsafe.Pointer(&s.storage[4])) = v
}

// Xi_socket returns the Xi_socket field from the record's packed storage.
func (s *Xinpcb) Xi_socket() Xsocket {
	return *(*Xsocket)(unsafe.Pointer(&s.storage[212]))
}

// SetXi_socket updates the Xi_socket field in the record's packed storage.
func (s *Xinpcb) SetXi_socket(v Xsocket) {
	*(*Xsocket)(unsafe.Pointer(&s.storage[212])) = v
}

// Xi_alignment_hack returns the Xi_alignment_hack field from the record's packed storage.
func (s *Xinpcb) Xi_alignment_hack() U_quad_t {
	return U_quad_t(binary.NativeEndian.Uint64(s.storage[312:320]))
}

// SetXi_alignment_hack updates the Xi_alignment_hack field in the record's packed storage.
func (s *Xinpcb) SetXi_alignment_hack(v U_quad_t) {
	binary.NativeEndian.PutUint64(s.storage[312:320], uint64(v))
}

// Xinpcb64
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/xinpcb64
type Xinpcb64 struct {
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
	storage [260]byte
}

// Xi_len returns the Xi_len field from the record's packed storage.
func (s *Xinpcb64) Xi_len() U_int64_t {
	return U_int64_t(binary.NativeEndian.Uint64(s.storage[0:8]))
}

// SetXi_len updates the Xi_len field in the record's packed storage.
func (s *Xinpcb64) SetXi_len(v U_int64_t) {
	binary.NativeEndian.PutUint64(s.storage[0:8], uint64(v))
}

// Xi_inpp returns the Xi_inpp field from the record's packed storage.
func (s *Xinpcb64) Xi_inpp() U_int64_t {
	return U_int64_t(binary.NativeEndian.Uint64(s.storage[8:16]))
}

// SetXi_inpp updates the Xi_inpp field in the record's packed storage.
func (s *Xinpcb64) SetXi_inpp(v U_int64_t) {
	binary.NativeEndian.PutUint64(s.storage[8:16], uint64(v))
}

// Inp_fport returns the Inp_fport field from the record's packed storage.
func (s *Xinpcb64) Inp_fport() U_short {
	return U_short(binary.NativeEndian.Uint16(s.storage[16:18]))
}

// SetInp_fport updates the Inp_fport field in the record's packed storage.
func (s *Xinpcb64) SetInp_fport(v U_short) {
	binary.NativeEndian.PutUint16(s.storage[16:18], uint16(v))
}

// Inp_lport returns the Inp_lport field from the record's packed storage.
func (s *Xinpcb64) Inp_lport() U_short {
	return U_short(binary.NativeEndian.Uint16(s.storage[18:20]))
}

// SetInp_lport updates the Inp_lport field in the record's packed storage.
func (s *Xinpcb64) SetInp_lport(v U_short) {
	binary.NativeEndian.PutUint16(s.storage[18:20], uint16(v))
}

// Inp_list returns the Inp_list field from the record's packed storage.
func (s *Xinpcb64) Inp_list() Inpcb64_list_entry {
	return *(*Inpcb64_list_entry)(unsafe.Pointer(&s.storage[20]))
}

// SetInp_list updates the Inp_list field in the record's packed storage.
func (s *Xinpcb64) SetInp_list(v Inpcb64_list_entry) {
	*(*Inpcb64_list_entry)(unsafe.Pointer(&s.storage[20])) = v
}

// Inp_ppcb returns the Inp_ppcb field from the record's packed storage.
func (s *Xinpcb64) Inp_ppcb() U_int64_t {
	return U_int64_t(binary.NativeEndian.Uint64(s.storage[36:44]))
}

// SetInp_ppcb updates the Inp_ppcb field in the record's packed storage.
func (s *Xinpcb64) SetInp_ppcb(v U_int64_t) {
	binary.NativeEndian.PutUint64(s.storage[36:44], uint64(v))
}

// Inp_pcbinfo returns the Inp_pcbinfo field from the record's packed storage.
func (s *Xinpcb64) Inp_pcbinfo() U_int64_t {
	return U_int64_t(binary.NativeEndian.Uint64(s.storage[44:52]))
}

// SetInp_pcbinfo updates the Inp_pcbinfo field in the record's packed storage.
func (s *Xinpcb64) SetInp_pcbinfo(v U_int64_t) {
	binary.NativeEndian.PutUint64(s.storage[44:52], uint64(v))
}

// Inp_portlist returns the Inp_portlist field from the record's packed storage.
func (s *Xinpcb64) Inp_portlist() Inpcb64_list_entry {
	return *(*Inpcb64_list_entry)(unsafe.Pointer(&s.storage[52]))
}

// SetInp_portlist updates the Inp_portlist field in the record's packed storage.
func (s *Xinpcb64) SetInp_portlist(v Inpcb64_list_entry) {
	*(*Inpcb64_list_entry)(unsafe.Pointer(&s.storage[52])) = v
}

// Inp_phd returns the Inp_phd field from the record's packed storage.
func (s *Xinpcb64) Inp_phd() U_int64_t {
	return U_int64_t(binary.NativeEndian.Uint64(s.storage[68:76]))
}

// SetInp_phd updates the Inp_phd field in the record's packed storage.
func (s *Xinpcb64) SetInp_phd(v U_int64_t) {
	binary.NativeEndian.PutUint64(s.storage[68:76], uint64(v))
}

// Inp_gencnt returns the Inp_gencnt field from the record's packed storage.
func (s *Xinpcb64) Inp_gencnt() Inp_gen_t {
	return Inp_gen_t(binary.NativeEndian.Uint64(s.storage[76:84]))
}

// SetInp_gencnt updates the Inp_gencnt field in the record's packed storage.
func (s *Xinpcb64) SetInp_gencnt(v Inp_gen_t) {
	binary.NativeEndian.PutUint64(s.storage[76:84], uint64(v))
}

// Inp_flags returns the Inp_flags field from the record's packed storage.
func (s *Xinpcb64) Inp_flags() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[84:88]))
}

// SetInp_flags updates the Inp_flags field in the record's packed storage.
func (s *Xinpcb64) SetInp_flags(v int32) {
	binary.NativeEndian.PutUint32(s.storage[84:88], uint32(v))
}

// Inp_flow returns the Inp_flow field from the record's packed storage.
func (s *Xinpcb64) Inp_flow() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[88:92]))
}

// SetInp_flow updates the Inp_flow field in the record's packed storage.
func (s *Xinpcb64) SetInp_flow(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[88:92], uint32(v))
}

// Inp_vflag returns the Inp_vflag field from the record's packed storage.
func (s *Xinpcb64) Inp_vflag() U_char {
	return *(*U_char)(unsafe.Pointer(&s.storage[92]))
}

// SetInp_vflag updates the Inp_vflag field in the record's packed storage.
func (s *Xinpcb64) SetInp_vflag(v U_char) {
	*(*U_char)(unsafe.Pointer(&s.storage[92])) = v
}

// Inp_ip_ttl returns the Inp_ip_ttl field from the record's packed storage.
func (s *Xinpcb64) Inp_ip_ttl() U_char {
	return *(*U_char)(unsafe.Pointer(&s.storage[93]))
}

// SetInp_ip_ttl updates the Inp_ip_ttl field in the record's packed storage.
func (s *Xinpcb64) SetInp_ip_ttl(v U_char) {
	*(*U_char)(unsafe.Pointer(&s.storage[93])) = v
}

// Inp_ip_p returns the Inp_ip_p field from the record's packed storage.
func (s *Xinpcb64) Inp_ip_p() U_char {
	return *(*U_char)(unsafe.Pointer(&s.storage[94]))
}

// SetInp_ip_p updates the Inp_ip_p field in the record's packed storage.
func (s *Xinpcb64) SetInp_ip_p(v U_char) {
	*(*U_char)(unsafe.Pointer(&s.storage[94])) = v
}

// Inp_dependfaddr returns the Inp_dependfaddr field from the record's packed storage.
func (s *Xinpcb64) Inp_dependfaddr() [4]uint32 {
	return *(*[4]uint32)(unsafe.Pointer(&s.storage[96]))
}

// SetInp_dependfaddr updates the Inp_dependfaddr field in the record's packed storage.
func (s *Xinpcb64) SetInp_dependfaddr(v [4]uint32) {
	*(*[4]uint32)(unsafe.Pointer(&s.storage[96])) = v
}

// Inp_dependladdr returns the Inp_dependladdr field from the record's packed storage.
func (s *Xinpcb64) Inp_dependladdr() [4]uint32 {
	return *(*[4]uint32)(unsafe.Pointer(&s.storage[112]))
}

// SetInp_dependladdr updates the Inp_dependladdr field in the record's packed storage.
func (s *Xinpcb64) SetInp_dependladdr(v [4]uint32) {
	*(*[4]uint32)(unsafe.Pointer(&s.storage[112])) = v
}

// Inp4_ip_tos returns the Inp4_ip_tos field from the record's packed storage.
func (s *Xinpcb64) Inp4_ip_tos() U_char {
	return *(*U_char)(unsafe.Pointer(&s.storage[128]))
}

// SetInp4_ip_tos updates the Inp4_ip_tos field in the record's packed storage.
func (s *Xinpcb64) SetInp4_ip_tos(v U_char) {
	*(*U_char)(unsafe.Pointer(&s.storage[128])) = v
}

// Inp6_hlim returns the Inp6_hlim field from the record's packed storage.
func (s *Xinpcb64) Inp6_hlim() U_int8_t {
	return *(*U_int8_t)(unsafe.Pointer(&s.storage[132]))
}

// SetInp6_hlim updates the Inp6_hlim field in the record's packed storage.
func (s *Xinpcb64) SetInp6_hlim(v U_int8_t) {
	*(*U_int8_t)(unsafe.Pointer(&s.storage[132])) = v
}

// Inp6_cksum returns the Inp6_cksum field from the record's packed storage.
func (s *Xinpcb64) Inp6_cksum() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[136:140]))
}

// SetInp6_cksum updates the Inp6_cksum field in the record's packed storage.
func (s *Xinpcb64) SetInp6_cksum(v int32) {
	binary.NativeEndian.PutUint32(s.storage[136:140], uint32(v))
}

// Inp6_ifindex returns the Inp6_ifindex field from the record's packed storage.
func (s *Xinpcb64) Inp6_ifindex() U_short {
	return U_short(binary.NativeEndian.Uint16(s.storage[140:142]))
}

// SetInp6_ifindex updates the Inp6_ifindex field in the record's packed storage.
func (s *Xinpcb64) SetInp6_ifindex(v U_short) {
	binary.NativeEndian.PutUint16(s.storage[140:142], uint16(v))
}

// Inp6_hops returns the Inp6_hops field from the record's packed storage.
func (s *Xinpcb64) Inp6_hops() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[142:144]))
}

// SetInp6_hops updates the Inp6_hops field in the record's packed storage.
func (s *Xinpcb64) SetInp6_hops(v int16) {
	binary.NativeEndian.PutUint16(s.storage[142:144], uint16(v))
}

// Xi_socket returns the Xi_socket field from the record's packed storage.
func (s *Xinpcb64) Xi_socket() Xsocket64 {
	return *(*Xsocket64)(unsafe.Pointer(&s.storage[144]))
}

// SetXi_socket updates the Xi_socket field in the record's packed storage.
func (s *Xinpcb64) SetXi_socket(v Xsocket64) {
	*(*Xsocket64)(unsafe.Pointer(&s.storage[144])) = v
}

// Xi_alignment_hack returns the Xi_alignment_hack field from the record's packed storage.
func (s *Xinpcb64) Xi_alignment_hack() U_quad_t {
	return U_quad_t(binary.NativeEndian.Uint64(s.storage[252:260]))
}

// SetXi_alignment_hack updates the Xi_alignment_hack field in the record's packed storage.
func (s *Xinpcb64) SetXi_alignment_hack(v U_quad_t) {
	binary.NativeEndian.PutUint64(s.storage[252:260], uint64(v))
}

// Xinpgen
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/xinpgen
type Xinpgen struct {
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
	storage [24]byte
}

// Xig_len returns the Xig_len field from the record's packed storage.
func (s *Xinpgen) Xig_len() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetXig_len updates the Xig_len field in the record's packed storage.
func (s *Xinpgen) SetXig_len(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// Xig_count returns the Xig_count field from the record's packed storage.
func (s *Xinpgen) Xig_count() U_int {
	return U_int(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetXig_count updates the Xig_count field in the record's packed storage.
func (s *Xinpgen) SetXig_count(v U_int) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// Xig_gen returns the Xig_gen field from the record's packed storage.
func (s *Xinpgen) Xig_gen() Inp_gen_t {
	return Inp_gen_t(binary.NativeEndian.Uint64(s.storage[8:16]))
}

// SetXig_gen updates the Xig_gen field in the record's packed storage.
func (s *Xinpgen) SetXig_gen(v Inp_gen_t) {
	binary.NativeEndian.PutUint64(s.storage[8:16], uint64(v))
}

// Xig_sogen returns the Xig_sogen field from the record's packed storage.
func (s *Xinpgen) Xig_sogen() So_gen_t {
	return So_gen_t(binary.NativeEndian.Uint64(s.storage[16:24]))
}

// SetXig_sogen updates the Xig_sogen field in the record's packed storage.
func (s *Xinpgen) SetXig_sogen(v So_gen_t) {
	binary.NativeEndian.PutUint64(s.storage[16:24], uint64(v))
}

// Xsockbuf
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/xsockbuf
type Xsockbuf struct {
	Sb_cc    U_int32_t
	Sb_hiwat U_int32_t
	Sb_mbcnt U_int32_t
	Sb_mbmax U_int32_t
	Sb_lowat int32
	Sb_flags int16
	Sb_timeo int16
}

// Xsocket
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/xsocket
type Xsocket struct {
	Xso_len      U_int32_t
	Xso_so       U_int32_t
	So_type      int16
	So_options   int16
	So_linger    int16
	So_state     int16
	So_pcb       U_int32_t
	Xso_protocol int32
	Xso_family   int32
	So_qlen      int16
	So_incqlen   int16
	So_qlimit    int16
	So_timeo     int16
	So_error     U_short
	So_pgid      int32
	So_oobmark   U_int32_t
	So_rcv       Xsockbuf
	So_snd       Xsockbuf
	So_uid       uint32
}

// Xsocket64
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/xsocket64
type Xsocket64 struct {
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
	storage [108]byte
}

// Xso_len returns the Xso_len field from the record's packed storage.
func (s *Xsocket64) Xso_len() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetXso_len updates the Xso_len field in the record's packed storage.
func (s *Xsocket64) SetXso_len(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// Xso_so returns the Xso_so field from the record's packed storage.
func (s *Xsocket64) Xso_so() U_int64_t {
	return U_int64_t(binary.NativeEndian.Uint64(s.storage[4:12]))
}

// SetXso_so updates the Xso_so field in the record's packed storage.
func (s *Xsocket64) SetXso_so(v U_int64_t) {
	binary.NativeEndian.PutUint64(s.storage[4:12], uint64(v))
}

// So_type returns the So_type field from the record's packed storage.
func (s *Xsocket64) So_type() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[12:14]))
}

// SetSo_type updates the So_type field in the record's packed storage.
func (s *Xsocket64) SetSo_type(v int16) {
	binary.NativeEndian.PutUint16(s.storage[12:14], uint16(v))
}

// So_options returns the So_options field from the record's packed storage.
func (s *Xsocket64) So_options() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[14:16]))
}

// SetSo_options updates the So_options field in the record's packed storage.
func (s *Xsocket64) SetSo_options(v int16) {
	binary.NativeEndian.PutUint16(s.storage[14:16], uint16(v))
}

// So_linger returns the So_linger field from the record's packed storage.
func (s *Xsocket64) So_linger() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[16:18]))
}

// SetSo_linger updates the So_linger field in the record's packed storage.
func (s *Xsocket64) SetSo_linger(v int16) {
	binary.NativeEndian.PutUint16(s.storage[16:18], uint16(v))
}

// So_state returns the So_state field from the record's packed storage.
func (s *Xsocket64) So_state() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[18:20]))
}

// SetSo_state updates the So_state field in the record's packed storage.
func (s *Xsocket64) SetSo_state(v int16) {
	binary.NativeEndian.PutUint16(s.storage[18:20], uint16(v))
}

// So_pcb returns the So_pcb field from the record's packed storage.
func (s *Xsocket64) So_pcb() U_int64_t {
	return U_int64_t(binary.NativeEndian.Uint64(s.storage[20:28]))
}

// SetSo_pcb updates the So_pcb field in the record's packed storage.
func (s *Xsocket64) SetSo_pcb(v U_int64_t) {
	binary.NativeEndian.PutUint64(s.storage[20:28], uint64(v))
}

// Xso_protocol returns the Xso_protocol field from the record's packed storage.
func (s *Xsocket64) Xso_protocol() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[28:32]))
}

// SetXso_protocol updates the Xso_protocol field in the record's packed storage.
func (s *Xsocket64) SetXso_protocol(v int32) {
	binary.NativeEndian.PutUint32(s.storage[28:32], uint32(v))
}

// Xso_family returns the Xso_family field from the record's packed storage.
func (s *Xsocket64) Xso_family() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[32:36]))
}

// SetXso_family updates the Xso_family field in the record's packed storage.
func (s *Xsocket64) SetXso_family(v int32) {
	binary.NativeEndian.PutUint32(s.storage[32:36], uint32(v))
}

// So_qlen returns the So_qlen field from the record's packed storage.
func (s *Xsocket64) So_qlen() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[36:38]))
}

// SetSo_qlen updates the So_qlen field in the record's packed storage.
func (s *Xsocket64) SetSo_qlen(v int16) {
	binary.NativeEndian.PutUint16(s.storage[36:38], uint16(v))
}

// So_incqlen returns the So_incqlen field from the record's packed storage.
func (s *Xsocket64) So_incqlen() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[38:40]))
}

// SetSo_incqlen updates the So_incqlen field in the record's packed storage.
func (s *Xsocket64) SetSo_incqlen(v int16) {
	binary.NativeEndian.PutUint16(s.storage[38:40], uint16(v))
}

// So_qlimit returns the So_qlimit field from the record's packed storage.
func (s *Xsocket64) So_qlimit() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[40:42]))
}

// SetSo_qlimit updates the So_qlimit field in the record's packed storage.
func (s *Xsocket64) SetSo_qlimit(v int16) {
	binary.NativeEndian.PutUint16(s.storage[40:42], uint16(v))
}

// So_timeo returns the So_timeo field from the record's packed storage.
func (s *Xsocket64) So_timeo() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[42:44]))
}

// SetSo_timeo updates the So_timeo field in the record's packed storage.
func (s *Xsocket64) SetSo_timeo(v int16) {
	binary.NativeEndian.PutUint16(s.storage[42:44], uint16(v))
}

// So_error returns the So_error field from the record's packed storage.
func (s *Xsocket64) So_error() U_short {
	return U_short(binary.NativeEndian.Uint16(s.storage[44:48]))
}

// SetSo_error updates the So_error field in the record's packed storage.
func (s *Xsocket64) SetSo_error(v U_short) {
	binary.NativeEndian.PutUint16(s.storage[44:48], uint16(v))
}

// So_pgid returns the So_pgid field from the record's packed storage.
func (s *Xsocket64) So_pgid() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[48:52]))
}

// SetSo_pgid updates the So_pgid field in the record's packed storage.
func (s *Xsocket64) SetSo_pgid(v int32) {
	binary.NativeEndian.PutUint32(s.storage[48:52], uint32(v))
}

// So_oobmark returns the So_oobmark field from the record's packed storage.
func (s *Xsocket64) So_oobmark() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[52:56]))
}

// SetSo_oobmark updates the So_oobmark field in the record's packed storage.
func (s *Xsocket64) SetSo_oobmark(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[52:56], uint32(v))
}

// So_rcv returns the So_rcv field from the record's packed storage.
func (s *Xsocket64) So_rcv() Xsockbuf {
	return *(*Xsockbuf)(unsafe.Pointer(&s.storage[56]))
}

// SetSo_rcv updates the So_rcv field in the record's packed storage.
func (s *Xsocket64) SetSo_rcv(v Xsockbuf) {
	*(*Xsockbuf)(unsafe.Pointer(&s.storage[56])) = v
}

// So_snd returns the So_snd field from the record's packed storage.
func (s *Xsocket64) So_snd() Xsockbuf {
	return *(*Xsockbuf)(unsafe.Pointer(&s.storage[80]))
}

// SetSo_snd updates the So_snd field in the record's packed storage.
func (s *Xsocket64) SetSo_snd(v Xsockbuf) {
	*(*Xsockbuf)(unsafe.Pointer(&s.storage[80])) = v
}

// So_uid returns the So_uid field from the record's packed storage.
func (s *Xsocket64) So_uid() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[104:108]))
}

// SetSo_uid updates the So_uid field in the record's packed storage.
func (s *Xsocket64) SetSo_uid(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[104:108], uint32(v))
}

// Xtcpcb
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/xtcpcb
type Xtcpcb struct {
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
	storage [524]byte
}

// Xt_len returns the Xt_len field from the record's packed storage.
func (s *Xtcpcb) Xt_len() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetXt_len updates the Xt_len field in the record's packed storage.
func (s *Xtcpcb) SetXt_len(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// Xt_inp returns the Xt_inp field from the record's packed storage.
func (s *Xtcpcb) Xt_inp() Inpcb {
	return *(*Inpcb)(unsafe.Pointer(&s.storage[4]))
}

// SetXt_inp updates the Xt_inp field in the record's packed storage.
func (s *Xtcpcb) SetXt_inp(v Inpcb) {
	*(*Inpcb)(unsafe.Pointer(&s.storage[4])) = v
}

// Xt_tp returns the Xt_tp field from the record's packed storage.
func (s *Xtcpcb) Xt_tp() Tcpcb {
	return *(*Tcpcb)(unsafe.Pointer(&s.storage[212]))
}

// SetXt_tp updates the Xt_tp field in the record's packed storage.
func (s *Xtcpcb) SetXt_tp(v Tcpcb) {
	*(*Tcpcb)(unsafe.Pointer(&s.storage[212])) = v
}

// Xt_socket returns the Xt_socket field from the record's packed storage.
func (s *Xtcpcb) Xt_socket() Xsocket {
	return *(*Xsocket)(unsafe.Pointer(&s.storage[416]))
}

// SetXt_socket updates the Xt_socket field in the record's packed storage.
func (s *Xtcpcb) SetXt_socket(v Xsocket) {
	*(*Xsocket)(unsafe.Pointer(&s.storage[416])) = v
}

// Xt_alignment_hack returns the Xt_alignment_hack field from the record's packed storage.
func (s *Xtcpcb) Xt_alignment_hack() U_quad_t {
	return U_quad_t(binary.NativeEndian.Uint64(s.storage[516:524]))
}

// SetXt_alignment_hack updates the Xt_alignment_hack field in the record's packed storage.
func (s *Xtcpcb) SetXt_alignment_hack(v U_quad_t) {
	binary.NativeEndian.PutUint64(s.storage[516:524], uint64(v))
}

// Xtcpcb64
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/xtcpcb64
type Xtcpcb64 struct {
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
	storage [472]byte
}

// Xt_len returns the Xt_len field from the record's packed storage.
func (s *Xtcpcb64) Xt_len() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetXt_len updates the Xt_len field in the record's packed storage.
func (s *Xtcpcb64) SetXt_len(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// Xt_inpcb returns the Xt_inpcb field from the record's packed storage.
func (s *Xtcpcb64) Xt_inpcb() Xinpcb64 {
	return *(*Xinpcb64)(unsafe.Pointer(&s.storage[4]))
}

// SetXt_inpcb updates the Xt_inpcb field in the record's packed storage.
func (s *Xtcpcb64) SetXt_inpcb(v Xinpcb64) {
	*(*Xinpcb64)(unsafe.Pointer(&s.storage[4])) = v
}

// T_segq returns the T_segq field from the record's packed storage.
func (s *Xtcpcb64) T_segq() U_int64_t {
	return U_int64_t(binary.NativeEndian.Uint64(s.storage[264:272]))
}

// SetT_segq updates the T_segq field in the record's packed storage.
func (s *Xtcpcb64) SetT_segq(v U_int64_t) {
	binary.NativeEndian.PutUint64(s.storage[264:272], uint64(v))
}

// T_dupacks returns the T_dupacks field from the record's packed storage.
func (s *Xtcpcb64) T_dupacks() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[272:276]))
}

// SetT_dupacks updates the T_dupacks field in the record's packed storage.
func (s *Xtcpcb64) SetT_dupacks(v int32) {
	binary.NativeEndian.PutUint32(s.storage[272:276], uint32(v))
}

// T_timer returns the T_timer field from the record's packed storage.
func (s *Xtcpcb64) T_timer() [4]int32 {
	return *(*[4]int32)(unsafe.Pointer(&s.storage[276]))
}

// SetT_timer updates the T_timer field in the record's packed storage.
func (s *Xtcpcb64) SetT_timer(v [4]int32) {
	*(*[4]int32)(unsafe.Pointer(&s.storage[276])) = v
}

// T_state returns the T_state field from the record's packed storage.
func (s *Xtcpcb64) T_state() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[292:296]))
}

// SetT_state updates the T_state field in the record's packed storage.
func (s *Xtcpcb64) SetT_state(v int32) {
	binary.NativeEndian.PutUint32(s.storage[292:296], uint32(v))
}

// T_flags returns the T_flags field from the record's packed storage.
func (s *Xtcpcb64) T_flags() U_int {
	return U_int(binary.NativeEndian.Uint32(s.storage[296:300]))
}

// SetT_flags updates the T_flags field in the record's packed storage.
func (s *Xtcpcb64) SetT_flags(v U_int) {
	binary.NativeEndian.PutUint32(s.storage[296:300], uint32(v))
}

// T_force returns the T_force field from the record's packed storage.
func (s *Xtcpcb64) T_force() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[300:304]))
}

// SetT_force updates the T_force field in the record's packed storage.
func (s *Xtcpcb64) SetT_force(v int32) {
	binary.NativeEndian.PutUint32(s.storage[300:304], uint32(v))
}

// Snd_una returns the Snd_una field from the record's packed storage.
func (s *Xtcpcb64) Snd_una() Tcp_seq {
	return Tcp_seq(binary.NativeEndian.Uint32(s.storage[304:308]))
}

// SetSnd_una updates the Snd_una field in the record's packed storage.
func (s *Xtcpcb64) SetSnd_una(v Tcp_seq) {
	binary.NativeEndian.PutUint32(s.storage[304:308], uint32(v))
}

// Snd_max returns the Snd_max field from the record's packed storage.
func (s *Xtcpcb64) Snd_max() Tcp_seq {
	return Tcp_seq(binary.NativeEndian.Uint32(s.storage[308:312]))
}

// SetSnd_max updates the Snd_max field in the record's packed storage.
func (s *Xtcpcb64) SetSnd_max(v Tcp_seq) {
	binary.NativeEndian.PutUint32(s.storage[308:312], uint32(v))
}

// Snd_nxt returns the Snd_nxt field from the record's packed storage.
func (s *Xtcpcb64) Snd_nxt() Tcp_seq {
	return Tcp_seq(binary.NativeEndian.Uint32(s.storage[312:316]))
}

// SetSnd_nxt updates the Snd_nxt field in the record's packed storage.
func (s *Xtcpcb64) SetSnd_nxt(v Tcp_seq) {
	binary.NativeEndian.PutUint32(s.storage[312:316], uint32(v))
}

// Snd_up returns the Snd_up field from the record's packed storage.
func (s *Xtcpcb64) Snd_up() Tcp_seq {
	return Tcp_seq(binary.NativeEndian.Uint32(s.storage[316:320]))
}

// SetSnd_up updates the Snd_up field in the record's packed storage.
func (s *Xtcpcb64) SetSnd_up(v Tcp_seq) {
	binary.NativeEndian.PutUint32(s.storage[316:320], uint32(v))
}

// Snd_wl1 returns the Snd_wl1 field from the record's packed storage.
func (s *Xtcpcb64) Snd_wl1() Tcp_seq {
	return Tcp_seq(binary.NativeEndian.Uint32(s.storage[320:324]))
}

// SetSnd_wl1 updates the Snd_wl1 field in the record's packed storage.
func (s *Xtcpcb64) SetSnd_wl1(v Tcp_seq) {
	binary.NativeEndian.PutUint32(s.storage[320:324], uint32(v))
}

// Snd_wl2 returns the Snd_wl2 field from the record's packed storage.
func (s *Xtcpcb64) Snd_wl2() Tcp_seq {
	return Tcp_seq(binary.NativeEndian.Uint32(s.storage[324:328]))
}

// SetSnd_wl2 updates the Snd_wl2 field in the record's packed storage.
func (s *Xtcpcb64) SetSnd_wl2(v Tcp_seq) {
	binary.NativeEndian.PutUint32(s.storage[324:328], uint32(v))
}

// Iss returns the Iss field from the record's packed storage.
func (s *Xtcpcb64) Iss() Tcp_seq {
	return Tcp_seq(binary.NativeEndian.Uint32(s.storage[328:332]))
}

// SetIss updates the Iss field in the record's packed storage.
func (s *Xtcpcb64) SetIss(v Tcp_seq) {
	binary.NativeEndian.PutUint32(s.storage[328:332], uint32(v))
}

// Irs returns the Irs field from the record's packed storage.
func (s *Xtcpcb64) Irs() Tcp_seq {
	return Tcp_seq(binary.NativeEndian.Uint32(s.storage[332:336]))
}

// SetIrs updates the Irs field in the record's packed storage.
func (s *Xtcpcb64) SetIrs(v Tcp_seq) {
	binary.NativeEndian.PutUint32(s.storage[332:336], uint32(v))
}

// Rcv_nxt returns the Rcv_nxt field from the record's packed storage.
func (s *Xtcpcb64) Rcv_nxt() Tcp_seq {
	return Tcp_seq(binary.NativeEndian.Uint32(s.storage[336:340]))
}

// SetRcv_nxt updates the Rcv_nxt field in the record's packed storage.
func (s *Xtcpcb64) SetRcv_nxt(v Tcp_seq) {
	binary.NativeEndian.PutUint32(s.storage[336:340], uint32(v))
}

// Rcv_adv returns the Rcv_adv field from the record's packed storage.
func (s *Xtcpcb64) Rcv_adv() Tcp_seq {
	return Tcp_seq(binary.NativeEndian.Uint32(s.storage[340:344]))
}

// SetRcv_adv updates the Rcv_adv field in the record's packed storage.
func (s *Xtcpcb64) SetRcv_adv(v Tcp_seq) {
	binary.NativeEndian.PutUint32(s.storage[340:344], uint32(v))
}

// Rcv_wnd returns the Rcv_wnd field from the record's packed storage.
func (s *Xtcpcb64) Rcv_wnd() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[344:348]))
}

// SetRcv_wnd updates the Rcv_wnd field in the record's packed storage.
func (s *Xtcpcb64) SetRcv_wnd(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[344:348], uint32(v))
}

// Rcv_up returns the Rcv_up field from the record's packed storage.
func (s *Xtcpcb64) Rcv_up() Tcp_seq {
	return Tcp_seq(binary.NativeEndian.Uint32(s.storage[348:352]))
}

// SetRcv_up updates the Rcv_up field in the record's packed storage.
func (s *Xtcpcb64) SetRcv_up(v Tcp_seq) {
	binary.NativeEndian.PutUint32(s.storage[348:352], uint32(v))
}

// Snd_wnd returns the Snd_wnd field from the record's packed storage.
func (s *Xtcpcb64) Snd_wnd() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[352:356]))
}

// SetSnd_wnd updates the Snd_wnd field in the record's packed storage.
func (s *Xtcpcb64) SetSnd_wnd(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[352:356], uint32(v))
}

// Snd_cwnd returns the Snd_cwnd field from the record's packed storage.
func (s *Xtcpcb64) Snd_cwnd() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[356:360]))
}

// SetSnd_cwnd updates the Snd_cwnd field in the record's packed storage.
func (s *Xtcpcb64) SetSnd_cwnd(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[356:360], uint32(v))
}

// Snd_ssthresh returns the Snd_ssthresh field from the record's packed storage.
func (s *Xtcpcb64) Snd_ssthresh() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[360:364]))
}

// SetSnd_ssthresh updates the Snd_ssthresh field in the record's packed storage.
func (s *Xtcpcb64) SetSnd_ssthresh(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[360:364], uint32(v))
}

// T_maxopd returns the T_maxopd field from the record's packed storage.
func (s *Xtcpcb64) T_maxopd() U_int {
	return U_int(binary.NativeEndian.Uint32(s.storage[364:368]))
}

// SetT_maxopd updates the T_maxopd field in the record's packed storage.
func (s *Xtcpcb64) SetT_maxopd(v U_int) {
	binary.NativeEndian.PutUint32(s.storage[364:368], uint32(v))
}

// T_rcvtime returns the T_rcvtime field from the record's packed storage.
func (s *Xtcpcb64) T_rcvtime() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[368:372]))
}

// SetT_rcvtime updates the T_rcvtime field in the record's packed storage.
func (s *Xtcpcb64) SetT_rcvtime(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[368:372], uint32(v))
}

// T_starttime returns the T_starttime field from the record's packed storage.
func (s *Xtcpcb64) T_starttime() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[372:376]))
}

// SetT_starttime updates the T_starttime field in the record's packed storage.
func (s *Xtcpcb64) SetT_starttime(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[372:376], uint32(v))
}

// T_rtttime returns the T_rtttime field from the record's packed storage.
func (s *Xtcpcb64) T_rtttime() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[376:380]))
}

// SetT_rtttime updates the T_rtttime field in the record's packed storage.
func (s *Xtcpcb64) SetT_rtttime(v int32) {
	binary.NativeEndian.PutUint32(s.storage[376:380], uint32(v))
}

// T_rtseq returns the T_rtseq field from the record's packed storage.
func (s *Xtcpcb64) T_rtseq() Tcp_seq {
	return Tcp_seq(binary.NativeEndian.Uint32(s.storage[380:384]))
}

// SetT_rtseq updates the T_rtseq field in the record's packed storage.
func (s *Xtcpcb64) SetT_rtseq(v Tcp_seq) {
	binary.NativeEndian.PutUint32(s.storage[380:384], uint32(v))
}

// T_rxtcur returns the T_rxtcur field from the record's packed storage.
func (s *Xtcpcb64) T_rxtcur() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[384:388]))
}

// SetT_rxtcur updates the T_rxtcur field in the record's packed storage.
func (s *Xtcpcb64) SetT_rxtcur(v int32) {
	binary.NativeEndian.PutUint32(s.storage[384:388], uint32(v))
}

// T_maxseg returns the T_maxseg field from the record's packed storage.
func (s *Xtcpcb64) T_maxseg() U_int {
	return U_int(binary.NativeEndian.Uint32(s.storage[388:392]))
}

// SetT_maxseg updates the T_maxseg field in the record's packed storage.
func (s *Xtcpcb64) SetT_maxseg(v U_int) {
	binary.NativeEndian.PutUint32(s.storage[388:392], uint32(v))
}

// T_srtt returns the T_srtt field from the record's packed storage.
func (s *Xtcpcb64) T_srtt() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[392:396]))
}

// SetT_srtt updates the T_srtt field in the record's packed storage.
func (s *Xtcpcb64) SetT_srtt(v int32) {
	binary.NativeEndian.PutUint32(s.storage[392:396], uint32(v))
}

// T_rttvar returns the T_rttvar field from the record's packed storage.
func (s *Xtcpcb64) T_rttvar() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[396:400]))
}

// SetT_rttvar updates the T_rttvar field in the record's packed storage.
func (s *Xtcpcb64) SetT_rttvar(v int32) {
	binary.NativeEndian.PutUint32(s.storage[396:400], uint32(v))
}

// T_rxtshift returns the T_rxtshift field from the record's packed storage.
func (s *Xtcpcb64) T_rxtshift() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[400:404]))
}

// SetT_rxtshift updates the T_rxtshift field in the record's packed storage.
func (s *Xtcpcb64) SetT_rxtshift(v int32) {
	binary.NativeEndian.PutUint32(s.storage[400:404], uint32(v))
}

// T_rttmin returns the T_rttmin field from the record's packed storage.
func (s *Xtcpcb64) T_rttmin() U_int {
	return U_int(binary.NativeEndian.Uint32(s.storage[404:408]))
}

// SetT_rttmin updates the T_rttmin field in the record's packed storage.
func (s *Xtcpcb64) SetT_rttmin(v U_int) {
	binary.NativeEndian.PutUint32(s.storage[404:408], uint32(v))
}

// T_rttupdated returns the T_rttupdated field from the record's packed storage.
func (s *Xtcpcb64) T_rttupdated() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[408:412]))
}

// SetT_rttupdated updates the T_rttupdated field in the record's packed storage.
func (s *Xtcpcb64) SetT_rttupdated(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[408:412], uint32(v))
}

// Max_sndwnd returns the Max_sndwnd field from the record's packed storage.
func (s *Xtcpcb64) Max_sndwnd() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[412:416]))
}

// SetMax_sndwnd updates the Max_sndwnd field in the record's packed storage.
func (s *Xtcpcb64) SetMax_sndwnd(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[412:416], uint32(v))
}

// T_softerror returns the T_softerror field from the record's packed storage.
func (s *Xtcpcb64) T_softerror() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[416:420]))
}

// SetT_softerror updates the T_softerror field in the record's packed storage.
func (s *Xtcpcb64) SetT_softerror(v int32) {
	binary.NativeEndian.PutUint32(s.storage[416:420], uint32(v))
}

// T_oobflags returns the T_oobflags field from the record's packed storage.
func (s *Xtcpcb64) T_oobflags() int8 {
	return int8(s.storage[420])
}

// SetT_oobflags updates the T_oobflags field in the record's packed storage.
func (s *Xtcpcb64) SetT_oobflags(v int8) {
	s.storage[420] = uint8(v)
}

// T_iobc returns the T_iobc field from the record's packed storage.
func (s *Xtcpcb64) T_iobc() int8 {
	return int8(s.storage[421])
}

// SetT_iobc updates the T_iobc field in the record's packed storage.
func (s *Xtcpcb64) SetT_iobc(v int8) {
	s.storage[421] = uint8(v)
}

// Snd_scale returns the Snd_scale field from the record's packed storage.
func (s *Xtcpcb64) Snd_scale() U_char {
	return *(*U_char)(unsafe.Pointer(&s.storage[422]))
}

// SetSnd_scale updates the Snd_scale field in the record's packed storage.
func (s *Xtcpcb64) SetSnd_scale(v U_char) {
	*(*U_char)(unsafe.Pointer(&s.storage[422])) = v
}

// Rcv_scale returns the Rcv_scale field from the record's packed storage.
func (s *Xtcpcb64) Rcv_scale() U_char {
	return *(*U_char)(unsafe.Pointer(&s.storage[423]))
}

// SetRcv_scale updates the Rcv_scale field in the record's packed storage.
func (s *Xtcpcb64) SetRcv_scale(v U_char) {
	*(*U_char)(unsafe.Pointer(&s.storage[423])) = v
}

// Request_r_scale returns the Request_r_scale field from the record's packed storage.
func (s *Xtcpcb64) Request_r_scale() U_char {
	return *(*U_char)(unsafe.Pointer(&s.storage[424]))
}

// SetRequest_r_scale updates the Request_r_scale field in the record's packed storage.
func (s *Xtcpcb64) SetRequest_r_scale(v U_char) {
	*(*U_char)(unsafe.Pointer(&s.storage[424])) = v
}

// Requested_s_scale returns the Requested_s_scale field from the record's packed storage.
func (s *Xtcpcb64) Requested_s_scale() U_char {
	return *(*U_char)(unsafe.Pointer(&s.storage[425]))
}

// SetRequested_s_scale updates the Requested_s_scale field in the record's packed storage.
func (s *Xtcpcb64) SetRequested_s_scale(v U_char) {
	*(*U_char)(unsafe.Pointer(&s.storage[425])) = v
}

// Ts_recent returns the Ts_recent field from the record's packed storage.
func (s *Xtcpcb64) Ts_recent() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[428:432]))
}

// SetTs_recent updates the Ts_recent field in the record's packed storage.
func (s *Xtcpcb64) SetTs_recent(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[428:432], uint32(v))
}

// Ts_recent_age returns the Ts_recent_age field from the record's packed storage.
func (s *Xtcpcb64) Ts_recent_age() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[432:436]))
}

// SetTs_recent_age updates the Ts_recent_age field in the record's packed storage.
func (s *Xtcpcb64) SetTs_recent_age(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[432:436], uint32(v))
}

// Last_ack_sent returns the Last_ack_sent field from the record's packed storage.
func (s *Xtcpcb64) Last_ack_sent() Tcp_seq {
	return Tcp_seq(binary.NativeEndian.Uint32(s.storage[436:440]))
}

// SetLast_ack_sent updates the Last_ack_sent field in the record's packed storage.
func (s *Xtcpcb64) SetLast_ack_sent(v Tcp_seq) {
	binary.NativeEndian.PutUint32(s.storage[436:440], uint32(v))
}

// Cc_send returns the Cc_send field from the record's packed storage.
func (s *Xtcpcb64) Cc_send() Tcp_cc {
	return Tcp_cc(binary.NativeEndian.Uint32(s.storage[440:444]))
}

// SetCc_send updates the Cc_send field in the record's packed storage.
func (s *Xtcpcb64) SetCc_send(v Tcp_cc) {
	binary.NativeEndian.PutUint32(s.storage[440:444], uint32(v))
}

// Cc_recv returns the Cc_recv field from the record's packed storage.
func (s *Xtcpcb64) Cc_recv() Tcp_cc {
	return Tcp_cc(binary.NativeEndian.Uint32(s.storage[444:448]))
}

// SetCc_recv updates the Cc_recv field in the record's packed storage.
func (s *Xtcpcb64) SetCc_recv(v Tcp_cc) {
	binary.NativeEndian.PutUint32(s.storage[444:448], uint32(v))
}

// Snd_recover returns the Snd_recover field from the record's packed storage.
func (s *Xtcpcb64) Snd_recover() Tcp_seq {
	return Tcp_seq(binary.NativeEndian.Uint32(s.storage[448:452]))
}

// SetSnd_recover updates the Snd_recover field in the record's packed storage.
func (s *Xtcpcb64) SetSnd_recover(v Tcp_seq) {
	binary.NativeEndian.PutUint32(s.storage[448:452], uint32(v))
}

// Snd_cwnd_prev returns the Snd_cwnd_prev field from the record's packed storage.
func (s *Xtcpcb64) Snd_cwnd_prev() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[452:456]))
}

// SetSnd_cwnd_prev updates the Snd_cwnd_prev field in the record's packed storage.
func (s *Xtcpcb64) SetSnd_cwnd_prev(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[452:456], uint32(v))
}

// Snd_ssthresh_prev returns the Snd_ssthresh_prev field from the record's packed storage.
func (s *Xtcpcb64) Snd_ssthresh_prev() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[456:460]))
}

// SetSnd_ssthresh_prev updates the Snd_ssthresh_prev field in the record's packed storage.
func (s *Xtcpcb64) SetSnd_ssthresh_prev(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[456:460], uint32(v))
}

// T_badrxtwin returns the T_badrxtwin field from the record's packed storage.
func (s *Xtcpcb64) T_badrxtwin() U_int32_t {
	return U_int32_t(binary.NativeEndian.Uint32(s.storage[460:464]))
}

// SetT_badrxtwin updates the T_badrxtwin field in the record's packed storage.
func (s *Xtcpcb64) SetT_badrxtwin(v U_int32_t) {
	binary.NativeEndian.PutUint32(s.storage[460:464], uint32(v))
}

// Xt_alignment_hack returns the Xt_alignment_hack field from the record's packed storage.
func (s *Xtcpcb64) Xt_alignment_hack() U_quad_t {
	return U_quad_t(binary.NativeEndian.Uint64(s.storage[464:472]))
}

// SetXt_alignment_hack updates the Xt_alignment_hack field in the record's packed storage.
func (s *Xtcpcb64) SetXt_alignment_hack(v U_quad_t) {
	binary.NativeEndian.PutUint64(s.storage[464:472], uint64(v))
}

// Xucred
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/xucred
type Xucred struct {
	Cr_version U_int
	Cr_uid     uint32
	Cr_ngroups int16
	Cr_groups  [16]uint32
}

// Xunpgen
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/xunpgen
type Xunpgen struct {
	Xug_len   U_int32_t
	Xug_count U_int
	Xug_gen   Unp_gen_t
	Xug_sogen So_gen_t
}

// Xvsockpcb
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/xvsockpcb
type Xvsockpcb struct {
	Xv_len           U_int32_t
	Xv_vsockpp       U_int64_t
	Xvp_local_cid    U_int32_t
	Xvp_local_port   U_int32_t
	Xvp_remote_cid   U_int32_t
	Xvp_remote_port  U_int32_t
	Xvp_rxcnt        U_int32_t
	Xvp_txcnt        U_int32_t
	Xvp_peer_rxhiwat U_int32_t
	Xvp_peer_rxcnt   U_int32_t
	Xvp_last_pid     int32
	Xvp_gencnt       Vsock_gen_t
	Xv_socket        Xsocket
}

// Xvsockpgen
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/xvsockpgen
type Xvsockpgen struct {
	Xvg_len   U_int32_t
	Xvg_count U_int64_t
	Xvg_gen   Vsock_gen_t
	Xvg_sogen So_gen_t
}
