// Code generated from Apple documentation for coreservices. DO NOT EDIT.

package coreservices

import (
	"encoding/binary"
	"unsafe"

	"github.com/tmc/apple/corefoundation"
)

// C struct types

// AEBuildError - Defines a structure for storing additional error codeinformation for “AEBuild” routines.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/aebuilderror
type AEBuildError struct {
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
	storage [8]byte
}

// FError returns the FError field from the record's packed storage.
func (s *AEBuildError) FError() AEBuildErrorCode {
	return AEBuildErrorCode(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetFError updates the FError field in the record's packed storage.
func (s *AEBuildError) SetFError(v AEBuildErrorCode) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// FErrorPos returns the FErrorPos field from the record's packed storage.
func (s *AEBuildError) FErrorPos() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetFErrorPos updates the FErrorPos field in the record's packed storage.
func (s *AEBuildError) SetFErrorPos(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// AEDesc - Stores data and an accompanying descriptor type to formthe basic building block of all Apple Events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/aedesc
type AEDesc struct {
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
	storage [12]byte
}

// DescriptorType returns the DescriptorType field from the record's packed storage.
func (s *AEDesc) DescriptorType() DescType {
	return DescType(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetDescriptorType updates the DescriptorType field in the record's packed storage.
func (s *AEDesc) SetDescriptorType(v DescType) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// DataHandle returns the DataHandle field from the record's packed storage.
func (s *AEDesc) DataHandle() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[4:12]))
}

// SetDataHandle updates the DataHandle field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *AEDesc) SetDataHandle(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[4:12], uint64(v))
}

// AEKeyDesc - Associates a keyword with a descriptor to form a keyword-specifieddescriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/aekeydesc
type AEKeyDesc struct {
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

// DescKey returns the DescKey field from the record's packed storage.
func (s *AEKeyDesc) DescKey() AEKeyword {
	return AEKeyword(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetDescKey updates the DescKey field in the record's packed storage.
func (s *AEKeyDesc) SetDescKey(v AEKeyword) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// DescContent returns the DescContent field from the record's packed storage.
func (s *AEKeyDesc) DescContent() AEDesc {
	return *(*AEDesc)(unsafe.Pointer(&s.storage[4]))
}

// SetDescContent updates the DescContent field in the record's packed storage.
func (s *AEKeyDesc) SetDescContent(v AEDesc) {
	*(*AEDesc)(unsafe.Pointer(&s.storage[4])) = v
}

// AERemoteProcessResolverContext - Supplied as a parameter when performing asynchronous resolutionof remote processes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/aeremoteprocessresolvercontext
type AERemoteProcessResolverContext struct {
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
	storage [40]byte
}

// Version returns the Version field from the record's packed storage.
func (s *AERemoteProcessResolverContext) Version() int {
	return int(binary.NativeEndian.Uint64(s.storage[0:8]))
}

// SetVersion updates the Version field in the record's packed storage.
func (s *AERemoteProcessResolverContext) SetVersion(v int) {
	binary.NativeEndian.PutUint64(s.storage[0:8], uint64(v))
}

// Info returns the Info field from the record's packed storage.
func (s *AERemoteProcessResolverContext) Info() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[8:16]))
}

// SetInfo updates the Info field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *AERemoteProcessResolverContext) SetInfo(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[8:16], uint64(v))
}

// Retain returns a raw C code address from the record's packed storage.
// The address is not callable from Go. To call C through it, use purego to
// bind a Go function value to the address.
func (s *AERemoteProcessResolverContext) Retain() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[16:24]))
}

// SetRetain stores a raw C code address in the record's packed storage.
// Use purego.NewCallback to obtain the address for a Go callback. The callback
// and its address must outlive every use C can make through this record.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *AERemoteProcessResolverContext) SetRetain(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[16:24], uint64(v))
}

// Release returns a raw C code address from the record's packed storage.
// The address is not callable from Go. To call C through it, use purego to
// bind a Go function value to the address.
func (s *AERemoteProcessResolverContext) Release() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[24:32]))
}

// SetRelease stores a raw C code address in the record's packed storage.
// Use purego.NewCallback to obtain the address for a Go callback. The callback
// and its address must outlive every use C can make through this record.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *AERemoteProcessResolverContext) SetRelease(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[24:32], uint64(v))
}

// CopyDescription returns a raw C code address from the record's packed storage.
// The address is not callable from Go. To call C through it, use purego to
// bind a Go function value to the address.
func (s *AERemoteProcessResolverContext) CopyDescription() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[32:40]))
}

// SetCopyDescription stores a raw C code address in the record's packed storage.
// Use purego.NewCallback to obtain the address for a Go callback. The callback
// and its address must outlive every use C can make through this record.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *AERemoteProcessResolverContext) SetCopyDescription(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[32:40], uint64(v))
}

// CSIdentityClientContext
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/csidentityclientcontext
type CSIdentityClientContext struct {
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
	storage [48]byte
}

// Version returns the Version field from the record's packed storage.
func (s *CSIdentityClientContext) Version() int {
	return int(binary.NativeEndian.Uint64(s.storage[0:8]))
}

// SetVersion updates the Version field in the record's packed storage.
func (s *CSIdentityClientContext) SetVersion(v int) {
	binary.NativeEndian.PutUint64(s.storage[0:8], uint64(v))
}

// Info returns the Info field from the record's packed storage.
func (s *CSIdentityClientContext) Info() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[8:16]))
}

// SetInfo updates the Info field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *CSIdentityClientContext) SetInfo(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[8:16], uint64(v))
}

// Retain returns a raw C code address from the record's packed storage.
// The address is not callable from Go. To call C through it, use purego to
// bind a Go function value to the address.
func (s *CSIdentityClientContext) Retain() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[16:24]))
}

// SetRetain stores a raw C code address in the record's packed storage.
// Use purego.NewCallback to obtain the address for a Go callback. The callback
// and its address must outlive every use C can make through this record.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *CSIdentityClientContext) SetRetain(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[16:24], uint64(v))
}

// Release returns a raw C code address from the record's packed storage.
// The address is not callable from Go. To call C through it, use purego to
// bind a Go function value to the address.
func (s *CSIdentityClientContext) Release() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[24:32]))
}

// SetRelease stores a raw C code address in the record's packed storage.
// Use purego.NewCallback to obtain the address for a Go callback. The callback
// and its address must outlive every use C can make through this record.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *CSIdentityClientContext) SetRelease(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[24:32], uint64(v))
}

// CopyDescription returns a raw C code address from the record's packed storage.
// The address is not callable from Go. To call C through it, use purego to
// bind a Go function value to the address.
func (s *CSIdentityClientContext) CopyDescription() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[32:40]))
}

// SetCopyDescription stores a raw C code address in the record's packed storage.
// Use purego.NewCallback to obtain the address for a Go callback. The callback
// and its address must outlive every use C can make through this record.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *CSIdentityClientContext) SetCopyDescription(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[32:40], uint64(v))
}

// StatusUpdated returns the StatusUpdated field from the record's packed storage.
func (s *CSIdentityClientContext) StatusUpdated() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[40:48]))
}

// SetStatusUpdated updates the StatusUpdated field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *CSIdentityClientContext) SetStatusUpdated(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[40:48], uint64(v))
}

// CSIdentityQueryClientContext
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/csidentityqueryclientcontext
type CSIdentityQueryClientContext struct {
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
	storage [48]byte
}

// Version returns the Version field from the record's packed storage.
func (s *CSIdentityQueryClientContext) Version() int {
	return int(binary.NativeEndian.Uint64(s.storage[0:8]))
}

// SetVersion updates the Version field in the record's packed storage.
func (s *CSIdentityQueryClientContext) SetVersion(v int) {
	binary.NativeEndian.PutUint64(s.storage[0:8], uint64(v))
}

// Info returns the Info field from the record's packed storage.
func (s *CSIdentityQueryClientContext) Info() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[8:16]))
}

// SetInfo updates the Info field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *CSIdentityQueryClientContext) SetInfo(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[8:16], uint64(v))
}

// RetainInfo returns a raw C code address from the record's packed storage.
// The address is not callable from Go. To call C through it, use purego to
// bind a Go function value to the address.
func (s *CSIdentityQueryClientContext) RetainInfo() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[16:24]))
}

// SetRetainInfo stores a raw C code address in the record's packed storage.
// Use purego.NewCallback to obtain the address for a Go callback. The callback
// and its address must outlive every use C can make through this record.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *CSIdentityQueryClientContext) SetRetainInfo(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[16:24], uint64(v))
}

// ReleaseInfo returns a raw C code address from the record's packed storage.
// The address is not callable from Go. To call C through it, use purego to
// bind a Go function value to the address.
func (s *CSIdentityQueryClientContext) ReleaseInfo() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[24:32]))
}

// SetReleaseInfo stores a raw C code address in the record's packed storage.
// Use purego.NewCallback to obtain the address for a Go callback. The callback
// and its address must outlive every use C can make through this record.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *CSIdentityQueryClientContext) SetReleaseInfo(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[24:32], uint64(v))
}

// CopyInfoDescription returns a raw C code address from the record's packed storage.
// The address is not callable from Go. To call C through it, use purego to
// bind a Go function value to the address.
func (s *CSIdentityQueryClientContext) CopyInfoDescription() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[32:40]))
}

// SetCopyInfoDescription stores a raw C code address in the record's packed storage.
// Use purego.NewCallback to obtain the address for a Go callback. The callback
// and its address must outlive every use C can make through this record.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *CSIdentityQueryClientContext) SetCopyInfoDescription(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[32:40], uint64(v))
}

// ReceiveEvent returns the ReceiveEvent field from the record's packed storage.
func (s *CSIdentityQueryClientContext) ReceiveEvent() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[40:48]))
}

// SetReceiveEvent updates the ReceiveEvent field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *CSIdentityQueryClientContext) SetReceiveEvent(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[40:48], uint64(v))
}

// FSEventStreamContext
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/fseventstreamcontext
type FSEventStreamContext struct {
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
	storage [40]byte
}

// Version returns the Version field from the record's packed storage.
func (s *FSEventStreamContext) Version() int {
	return int(binary.NativeEndian.Uint64(s.storage[0:8]))
}

// SetVersion updates the Version field in the record's packed storage.
func (s *FSEventStreamContext) SetVersion(v int) {
	binary.NativeEndian.PutUint64(s.storage[0:8], uint64(v))
}

// Info returns the Info field from the record's packed storage.
func (s *FSEventStreamContext) Info() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[8:16]))
}

// SetInfo updates the Info field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *FSEventStreamContext) SetInfo(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[8:16], uint64(v))
}

// Retain returns a raw C code address from the record's packed storage.
// The address is not callable from Go. To call C through it, use purego to
// bind a Go function value to the address.
func (s *FSEventStreamContext) Retain() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[16:24]))
}

// SetRetain stores a raw C code address in the record's packed storage.
// Use purego.NewCallback to obtain the address for a Go callback. The callback
// and its address must outlive every use C can make through this record.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *FSEventStreamContext) SetRetain(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[16:24], uint64(v))
}

// Release returns a raw C code address from the record's packed storage.
// The address is not callable from Go. To call C through it, use purego to
// bind a Go function value to the address.
func (s *FSEventStreamContext) Release() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[24:32]))
}

// SetRelease stores a raw C code address in the record's packed storage.
// Use purego.NewCallback to obtain the address for a Go callback. The callback
// and its address must outlive every use C can make through this record.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *FSEventStreamContext) SetRelease(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[24:32], uint64(v))
}

// CopyDescription returns a raw C code address from the record's packed storage.
// The address is not callable from Go. To call C through it, use purego to
// bind a Go function value to the address.
func (s *FSEventStreamContext) CopyDescription() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[32:40]))
}

// SetCopyDescription stores a raw C code address in the record's packed storage.
// Use purego.NewCallback to obtain the address for a Go callback. The callback
// and its address must outlive every use C can make through this record.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *FSEventStreamContext) SetCopyDescription(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[32:40], uint64(v))
}

// IntlText - International text consists of an ordered series of bytes, beginning with a 4-byte language code and a 4-byte script code that together determine the format of the bytes that follow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/intltext
type IntlText struct {
	TheScriptCode int16
	TheLangCode   int16
	TheText       [1]int8
}

// LSApplicationParameters - The specification that defines the app, launch flags, and additional parameters that control how an app launches.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/lsapplicationparameters
type LSApplicationParameters struct {
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
	storage [52]byte
}

// Version returns the Version field from the record's packed storage.
func (s *LSApplicationParameters) Version() int {
	return int(binary.NativeEndian.Uint64(s.storage[0:8]))
}

// SetVersion updates the Version field in the record's packed storage.
func (s *LSApplicationParameters) SetVersion(v int) {
	binary.NativeEndian.PutUint64(s.storage[0:8], uint64(v))
}

// Flags returns the Flags field from the record's packed storage.
func (s *LSApplicationParameters) Flags() LSLaunchFlags {
	return *(*LSLaunchFlags)(unsafe.Pointer(&s.storage[8]))
}

// SetFlags updates the Flags field in the record's packed storage.
func (s *LSApplicationParameters) SetFlags(v LSLaunchFlags) {
	*(*LSLaunchFlags)(unsafe.Pointer(&s.storage[8])) = v
}

// Application returns the Application field from the record's packed storage.
func (s *LSApplicationParameters) Application() *FSRef {
	return *(**FSRef)(unsafe.Pointer(&s.storage[12]))
}

// SetApplication updates the Application field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *LSApplicationParameters) SetApplication(v *FSRef) {
	*(**FSRef)(unsafe.Pointer(&s.storage[12])) = v
}

// AsyncLaunchRefCon returns the AsyncLaunchRefCon field from the record's packed storage.
func (s *LSApplicationParameters) AsyncLaunchRefCon() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[20:28]))
}

// SetAsyncLaunchRefCon updates the AsyncLaunchRefCon field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *LSApplicationParameters) SetAsyncLaunchRefCon(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[20:28], uint64(v))
}

// Environment returns the Environment field from the record's packed storage.
func (s *LSApplicationParameters) Environment() corefoundation.CFDictionaryRef {
	return corefoundation.CFDictionaryRef(binary.NativeEndian.Uint64(s.storage[28:36]))
}

// SetEnvironment updates the Environment field in the record's packed storage.
func (s *LSApplicationParameters) SetEnvironment(v corefoundation.CFDictionaryRef) {
	binary.NativeEndian.PutUint64(s.storage[28:36], uint64(v))
}

// Argv returns the Argv field from the record's packed storage.
func (s *LSApplicationParameters) Argv() corefoundation.CFArrayRef {
	return corefoundation.CFArrayRef(binary.NativeEndian.Uint64(s.storage[36:44]))
}

// SetArgv updates the Argv field in the record's packed storage.
func (s *LSApplicationParameters) SetArgv(v corefoundation.CFArrayRef) {
	binary.NativeEndian.PutUint64(s.storage[36:44], uint64(v))
}

// InitialEvent returns the InitialEvent field from the record's packed storage.
func (s *LSApplicationParameters) InitialEvent() *AEDesc {
	return *(**AEDesc)(unsafe.Pointer(&s.storage[44]))
}

// SetInitialEvent updates the InitialEvent field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *LSApplicationParameters) SetInitialEvent(v *AEDesc) {
	*(**AEDesc)(unsafe.Pointer(&s.storage[44])) = v
}

// LSItemInfoRecord - The specification that contains requested information about an item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/lsiteminforecord
type LSItemInfoRecord struct {
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
	storage [20]byte
}

// Flags returns the Flags field from the record's packed storage.
func (s *LSItemInfoRecord) Flags() LSItemInfoFlags {
	return *(*LSItemInfoFlags)(unsafe.Pointer(&s.storage[0]))
}

// SetFlags updates the Flags field in the record's packed storage.
func (s *LSItemInfoRecord) SetFlags(v LSItemInfoFlags) {
	*(*LSItemInfoFlags)(unsafe.Pointer(&s.storage[0])) = v
}

// Filetype returns the Filetype field from the record's packed storage.
func (s *LSItemInfoRecord) Filetype() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetFiletype updates the Filetype field in the record's packed storage.
func (s *LSItemInfoRecord) SetFiletype(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// Creator returns the Creator field from the record's packed storage.
func (s *LSItemInfoRecord) Creator() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetCreator updates the Creator field in the record's packed storage.
func (s *LSItemInfoRecord) SetCreator(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[8:12], uint32(v))
}

// Extension returns the Extension field from the record's packed storage.
func (s *LSItemInfoRecord) Extension() corefoundation.CFStringRef {
	return corefoundation.CFStringRef(binary.NativeEndian.Uint64(s.storage[12:20]))
}

// SetExtension updates the Extension field in the record's packed storage.
func (s *LSItemInfoRecord) SetExtension(v corefoundation.CFStringRef) {
	binary.NativeEndian.PutUint64(s.storage[12:20], uint64(v))
}

// LSLaunchFSRefSpec - The specification that defines, by file-system reference, an app to launch, items to open, or both, along with related information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/lslaunchfsrefspec
type LSLaunchFSRefSpec struct {
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
	storage [44]byte
}

// AppRef returns the AppRef field from the record's packed storage.
func (s *LSLaunchFSRefSpec) AppRef() *FSRef {
	return *(**FSRef)(unsafe.Pointer(&s.storage[0]))
}

// SetAppRef updates the AppRef field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *LSLaunchFSRefSpec) SetAppRef(v *FSRef) {
	*(**FSRef)(unsafe.Pointer(&s.storage[0])) = v
}

// NumDocs returns the NumDocs field from the record's packed storage.
func (s *LSLaunchFSRefSpec) NumDocs() uint {
	return uint(binary.NativeEndian.Uint64(s.storage[8:16]))
}

// SetNumDocs updates the NumDocs field in the record's packed storage.
func (s *LSLaunchFSRefSpec) SetNumDocs(v uint) {
	binary.NativeEndian.PutUint64(s.storage[8:16], uint64(v))
}

// ItemRefs returns the ItemRefs field from the record's packed storage.
func (s *LSLaunchFSRefSpec) ItemRefs() *FSRef {
	return *(**FSRef)(unsafe.Pointer(&s.storage[16]))
}

// SetItemRefs updates the ItemRefs field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *LSLaunchFSRefSpec) SetItemRefs(v *FSRef) {
	*(**FSRef)(unsafe.Pointer(&s.storage[16])) = v
}

// PassThruParams returns the PassThruParams field from the record's packed storage.
func (s *LSLaunchFSRefSpec) PassThruParams() *AEDesc {
	return *(**AEDesc)(unsafe.Pointer(&s.storage[24]))
}

// SetPassThruParams updates the PassThruParams field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *LSLaunchFSRefSpec) SetPassThruParams(v *AEDesc) {
	*(**AEDesc)(unsafe.Pointer(&s.storage[24])) = v
}

// LaunchFlags returns the LaunchFlags field from the record's packed storage.
func (s *LSLaunchFSRefSpec) LaunchFlags() LSLaunchFlags {
	return *(*LSLaunchFlags)(unsafe.Pointer(&s.storage[32]))
}

// SetLaunchFlags updates the LaunchFlags field in the record's packed storage.
func (s *LSLaunchFSRefSpec) SetLaunchFlags(v LSLaunchFlags) {
	*(*LSLaunchFlags)(unsafe.Pointer(&s.storage[32])) = v
}

// AsyncRefCon returns the AsyncRefCon field from the record's packed storage.
func (s *LSLaunchFSRefSpec) AsyncRefCon() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[36:44]))
}

// SetAsyncRefCon updates the AsyncRefCon field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *LSLaunchFSRefSpec) SetAsyncRefCon(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[36:44], uint64(v))
}

// LSLaunchURLSpec - The specification for launching an app, opening items, or both, along with related information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/lslaunchurlspec
type LSLaunchURLSpec struct {
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
	storage [36]byte
}

// AppURL returns the AppURL field from the record's packed storage.
func (s *LSLaunchURLSpec) AppURL() corefoundation.CFURLRef {
	return corefoundation.CFURLRef(binary.NativeEndian.Uint64(s.storage[0:8]))
}

// SetAppURL updates the AppURL field in the record's packed storage.
func (s *LSLaunchURLSpec) SetAppURL(v corefoundation.CFURLRef) {
	binary.NativeEndian.PutUint64(s.storage[0:8], uint64(v))
}

// ItemURLs returns the ItemURLs field from the record's packed storage.
func (s *LSLaunchURLSpec) ItemURLs() corefoundation.CFArrayRef {
	return corefoundation.CFArrayRef(binary.NativeEndian.Uint64(s.storage[8:16]))
}

// SetItemURLs updates the ItemURLs field in the record's packed storage.
func (s *LSLaunchURLSpec) SetItemURLs(v corefoundation.CFArrayRef) {
	binary.NativeEndian.PutUint64(s.storage[8:16], uint64(v))
}

// PassThruParams returns the PassThruParams field from the record's packed storage.
func (s *LSLaunchURLSpec) PassThruParams() *AEDesc {
	return *(**AEDesc)(unsafe.Pointer(&s.storage[16]))
}

// SetPassThruParams updates the PassThruParams field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *LSLaunchURLSpec) SetPassThruParams(v *AEDesc) {
	*(**AEDesc)(unsafe.Pointer(&s.storage[16])) = v
}

// LaunchFlags returns the LaunchFlags field from the record's packed storage.
func (s *LSLaunchURLSpec) LaunchFlags() LSLaunchFlags {
	return *(*LSLaunchFlags)(unsafe.Pointer(&s.storage[24]))
}

// SetLaunchFlags updates the LaunchFlags field in the record's packed storage.
func (s *LSLaunchURLSpec) SetLaunchFlags(v LSLaunchFlags) {
	*(*LSLaunchFlags)(unsafe.Pointer(&s.storage[24])) = v
}

// AsyncRefCon returns the AsyncRefCon field from the record's packed storage.
func (s *LSLaunchURLSpec) AsyncRefCon() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[28:36]))
}

// SetAsyncRefCon updates the AsyncRefCon field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *LSLaunchURLSpec) SetAsyncRefCon(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[28:36], uint64(v))
}

// MDExporterInterfaceStruct
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/mdexporterinterfacestruct
type MDExporterInterfaceStruct struct {
	_reserved          unsafe.Pointer
	QueryInterface     unsafe.Pointer
	AddRef             unsafe.Pointer
	Release            unsafe.Pointer
	ImporterExportData unsafe.Pointer
}

// MDImporterBundleWrapperURLInterfaceStruct
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/mdimporterbundlewrapperurlinterfacestruct
type MDImporterBundleWrapperURLInterfaceStruct struct {
	_reserved                          unsafe.Pointer
	QueryInterface                     unsafe.Pointer
	AddRef                             unsafe.Pointer
	Release                            unsafe.Pointer
	ImporterImportBundleWrapperURLData unsafe.Pointer
}

// MDImporterInterfaceStruct
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/mdimporterinterfacestruct
type MDImporterInterfaceStruct struct {
	_reserved          unsafe.Pointer
	QueryInterface     unsafe.Pointer
	AddRef             unsafe.Pointer
	Release            unsafe.Pointer
	ImporterImportData unsafe.Pointer
}

// MDImporterURLInterfaceStruct
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/mdimporterurlinterfacestruct
type MDImporterURLInterfaceStruct struct {
	_reserved             unsafe.Pointer
	QueryInterface        unsafe.Pointer
	AddRef                unsafe.Pointer
	Release               unsafe.Pointer
	ImporterImportURLData unsafe.Pointer
}

// MDQueryBatchingParams - Structure containing the progress notification batchingparameters of a MDQuery.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/mdquerybatchingparams
type MDQueryBatchingParams struct {
	First_max_num    uintptr // The maximum number of results that can accumulatebefore the first progress notification is sent. This value is usedonly during the initial result-gathering phase of a query.
	First_max_ms     uintptr // The maximum number of milliseconds that can passbefore the first progress notification is sent. This value is advisory,in that the notification will be triggered at some point after `first_max_ms` millisecondshave passed since the query began accumulating results. This valueis used only during the initial result-gathering phase of a query.
	Progress_max_num uintptr // The maximum number of results that can accumulatebefore additional progress notifications are sent. This value isused only during the initial result-gathering phase of a query.
	Progress_max_ms  uintptr // The maximum number of milliseconds that can passbefore additional progress notifications are sent. This value isadvisory, in that the notification will be triggered at some pointafter `progress_max_ms` millisecondshave passed since the query began accumulating results. This valueis used only during the initial result-gathering phase of a query.
	Update_max_num   uintptr // The maximum number of results that can accumulatebefore an update notification is sent. This value is used only duringthe live-update phase of a query.
	Update_max_ms    uintptr // The maximum number of milliseconds that can passbefore an update notification is sent. This value is advisory, inthat the notification will be triggered at some point after `update_max_ms` millisecondshave passed since the query began accumulating results. This valueis used only during the live-update phase of a query.

}

// OffsetArray - Specifies offsets of ranges of text. Not typically used by developers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/offsetarray
type OffsetArray struct {
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

// FNumOfOffsets returns the FNumOfOffsets field from the record's packed storage.
func (s *OffsetArray) FNumOfOffsets() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[0:2]))
}

// SetFNumOfOffsets updates the FNumOfOffsets field in the record's packed storage.
func (s *OffsetArray) SetFNumOfOffsets(v int16) {
	binary.NativeEndian.PutUint16(s.storage[0:2], uint16(v))
}

// FOffset returns the FOffset field from the record's packed storage.
func (s *OffsetArray) FOffset() [1]int32 {
	return *(*[1]int32)(unsafe.Pointer(&s.storage[2]))
}

// SetFOffset updates the FOffset field in the record's packed storage.
func (s *OffsetArray) SetFOffset(v [1]int32) {
	*(*[1]int32)(unsafe.Pointer(&s.storage[2])) = v
}

// TScriptingSizeResource - Defines a data type to store stack and heap information. Not typically used by developers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/tscriptingsizeresource
type TScriptingSizeResource struct {
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
	storage [26]byte
}

// ScriptingSizeFlags returns the ScriptingSizeFlags field from the record's packed storage.
func (s *TScriptingSizeResource) ScriptingSizeFlags() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[0:2]))
}

// SetScriptingSizeFlags updates the ScriptingSizeFlags field in the record's packed storage.
func (s *TScriptingSizeResource) SetScriptingSizeFlags(v int16) {
	binary.NativeEndian.PutUint16(s.storage[0:2], uint16(v))
}

// MinStackSize returns the MinStackSize field from the record's packed storage.
func (s *TScriptingSizeResource) MinStackSize() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[2:6]))
}

// SetMinStackSize updates the MinStackSize field in the record's packed storage.
func (s *TScriptingSizeResource) SetMinStackSize(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[2:6], uint32(v))
}

// PreferredStackSize returns the PreferredStackSize field from the record's packed storage.
func (s *TScriptingSizeResource) PreferredStackSize() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[6:10]))
}

// SetPreferredStackSize updates the PreferredStackSize field in the record's packed storage.
func (s *TScriptingSizeResource) SetPreferredStackSize(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[6:10], uint32(v))
}

// MaxStackSize returns the MaxStackSize field from the record's packed storage.
func (s *TScriptingSizeResource) MaxStackSize() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[10:14]))
}

// SetMaxStackSize updates the MaxStackSize field in the record's packed storage.
func (s *TScriptingSizeResource) SetMaxStackSize(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[10:14], uint32(v))
}

// MinHeapSize returns the MinHeapSize field from the record's packed storage.
func (s *TScriptingSizeResource) MinHeapSize() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[14:18]))
}

// SetMinHeapSize updates the MinHeapSize field in the record's packed storage.
func (s *TScriptingSizeResource) SetMinHeapSize(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[14:18], uint32(v))
}

// PreferredHeapSize returns the PreferredHeapSize field from the record's packed storage.
func (s *TScriptingSizeResource) PreferredHeapSize() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[18:22]))
}

// SetPreferredHeapSize updates the PreferredHeapSize field in the record's packed storage.
func (s *TScriptingSizeResource) SetPreferredHeapSize(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[18:22], uint32(v))
}

// MaxHeapSize returns the MaxHeapSize field from the record's packed storage.
func (s *TScriptingSizeResource) MaxHeapSize() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[22:26]))
}

// SetMaxHeapSize updates the MaxHeapSize field in the record's packed storage.
func (s *TScriptingSizeResource) SetMaxHeapSize(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[22:26], uint32(v))
}

// TextRange - Specifies a range of text. Not typically used by developers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/textrange
type TextRange struct {
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
	storage [10]byte
}

// FStart returns the FStart field from the record's packed storage.
func (s *TextRange) FStart() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetFStart updates the FStart field in the record's packed storage.
func (s *TextRange) SetFStart(v int32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// FEnd returns the FEnd field from the record's packed storage.
func (s *TextRange) FEnd() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetFEnd updates the FEnd field in the record's packed storage.
func (s *TextRange) SetFEnd(v int32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// FHiliteStyle returns the FHiliteStyle field from the record's packed storage.
func (s *TextRange) FHiliteStyle() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[8:10]))
}

// SetFHiliteStyle updates the FHiliteStyle field in the record's packed storage.
func (s *TextRange) SetFHiliteStyle(v int16) {
	binary.NativeEndian.PutUint16(s.storage[8:10], uint16(v))
}

// TextRangeArray - Specifies an array of text ranges. Not typically used by developers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/textrangearray
type TextRangeArray struct {
	FNumOfRanges int16
	FRange       [1]TextRange
}

// UCKeyLayoutFeatureInfo - Specifies the longest possible output string to be produced by the current `'uchr'` resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/uckeylayoutfeatureinfo
type UCKeyLayoutFeatureInfo struct {
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
	storage [8]byte
}

// KeyLayoutFeatureInfoFormat returns the KeyLayoutFeatureInfoFormat field from the record's packed storage.
func (s *UCKeyLayoutFeatureInfo) KeyLayoutFeatureInfoFormat() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[0:2]))
}

// SetKeyLayoutFeatureInfoFormat updates the KeyLayoutFeatureInfoFormat field in the record's packed storage.
func (s *UCKeyLayoutFeatureInfo) SetKeyLayoutFeatureInfoFormat(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[0:2], uint16(v))
}

// Reserved returns the Reserved field from the record's packed storage.
func (s *UCKeyLayoutFeatureInfo) Reserved() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[2:4]))
}

// SetReserved updates the Reserved field in the record's packed storage.
func (s *UCKeyLayoutFeatureInfo) SetReserved(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[2:4], uint16(v))
}

// MaxOutputStringLength returns the MaxOutputStringLength field from the record's packed storage.
func (s *UCKeyLayoutFeatureInfo) MaxOutputStringLength() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetMaxOutputStringLength updates the MaxOutputStringLength field in the record's packed storage.
func (s *UCKeyLayoutFeatureInfo) SetMaxOutputStringLength(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// UCKeyModifiersToTableNum - Maps a modifier key combination to a particular key-code-to-character table number in a `'uchr'` resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/uckeymodifierstotablenum
type UCKeyModifiersToTableNum struct {
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
	storage [10]byte
}

// KeyModifiersToTableNumFormat returns the KeyModifiersToTableNumFormat field from the record's packed storage.
func (s *UCKeyModifiersToTableNum) KeyModifiersToTableNumFormat() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[0:2]))
}

// SetKeyModifiersToTableNumFormat updates the KeyModifiersToTableNumFormat field in the record's packed storage.
func (s *UCKeyModifiersToTableNum) SetKeyModifiersToTableNumFormat(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[0:2], uint16(v))
}

// DefaultTableNum returns the DefaultTableNum field from the record's packed storage.
func (s *UCKeyModifiersToTableNum) DefaultTableNum() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[2:4]))
}

// SetDefaultTableNum updates the DefaultTableNum field in the record's packed storage.
func (s *UCKeyModifiersToTableNum) SetDefaultTableNum(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[2:4], uint16(v))
}

// ModifiersCount returns the ModifiersCount field from the record's packed storage.
func (s *UCKeyModifiersToTableNum) ModifiersCount() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetModifiersCount updates the ModifiersCount field in the record's packed storage.
func (s *UCKeyModifiersToTableNum) SetModifiersCount(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// TableNum returns the TableNum field from the record's packed storage.
func (s *UCKeyModifiersToTableNum) TableNum() [1]uint8 {
	return *(*[1]uint8)(unsafe.Pointer(&s.storage[8]))
}

// SetTableNum updates the TableNum field in the record's packed storage.
func (s *UCKeyModifiersToTableNum) SetTableNum(v [1]uint8) {
	*(*[1]uint8)(unsafe.Pointer(&s.storage[8])) = v
}

// UCKeySequenceDataIndex - Contains offsets to a list of character sequences for a `'uchr'` resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/uckeysequencedataindex
type UCKeySequenceDataIndex struct {
	KeySequenceDataIndexFormat uint16    // An unsigned 16-bit integer identifying the format of the [UCKeySequenceDataIndex] structure. Set to `kUCKeySequenceDataIndexFormat`.
	CharSequenceCount          uint16    // An unsigned 16-bit integer specifying the number of Unicode character sequences that follow the end of the [UCKeySequenceDataIndex] structure.
	CharSequenceOffsets        [1]uint16 // An array of offsets from the beginning of the [UCKeySequenceDataIndex] structure to the Unicode character sequences that follow it. Because a given offset indicates both the beginning of a new character sequence and the end of the sequence that precedes it, the length of each sequence is determined by the difference between the offset to that sequence and the value of the next offset in the array. The array contains one more entry than the number of character sequences; the final entry is the offset to the end of the final character sequence.

}

// UCKeyStateEntryRange - Maps from a dead-key state to either the resultant Unicode character(s) or the new dead key state produced when the current state is terminated by a given character key for a `'uchr'` resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/uckeystateentryrange
type UCKeyStateEntryRange struct {
	CurStateStart   uint16       // An unsigned 16-bit integer specifying the beginning of a given dead-key state range.
	CurStateRange   uint8        // An unsigned 8-bit integer specifying the number of entries in a given dead-key state range.
	DeltaMultiplier uint8        // An unsigned 8-bit integer.
	CharData        UCKeyCharSeq // A value of type [UCKeyCharSeq]. This base character value is used to determine the actual Unicode character(s) produced when a given dead-key state terminates.
	NextState       uint16       // An unsigned 16-bit integer. This base dead-key state value is used to determine the following dead-key state, if any.

}

// UCKeyStateEntryTerminal - Maps from a dead-key state to the Unicode character(s) produced when that state is terminated by a given character key for a `'uchr'` resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/uckeystateentryterminal
type UCKeyStateEntryTerminal struct {
	CurState uint16       // An unsigned 16-bit integer specifying the current dead-key state.
	CharData UCKeyCharSeq // A value of type [UCKeyCharSeq] specifying the Unicode character(s) produced when a given character key is pressed.

}

// UCKeyStateRecord - Determines dead-key state transitions in a `'uchr'` resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/uckeystaterecord
type UCKeyStateRecord struct {
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
	storage [12]byte
}

// StateZeroCharData returns the StateZeroCharData field from the record's packed storage.
func (s *UCKeyStateRecord) StateZeroCharData() UCKeyCharSeq {
	return UCKeyCharSeq(binary.NativeEndian.Uint16(s.storage[0:2]))
}

// SetStateZeroCharData updates the StateZeroCharData field in the record's packed storage.
func (s *UCKeyStateRecord) SetStateZeroCharData(v UCKeyCharSeq) {
	binary.NativeEndian.PutUint16(s.storage[0:2], uint16(v))
}

// StateZeroNextState returns the StateZeroNextState field from the record's packed storage.
func (s *UCKeyStateRecord) StateZeroNextState() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[2:4]))
}

// SetStateZeroNextState updates the StateZeroNextState field in the record's packed storage.
func (s *UCKeyStateRecord) SetStateZeroNextState(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[2:4], uint16(v))
}

// StateEntryCount returns the StateEntryCount field from the record's packed storage.
func (s *UCKeyStateRecord) StateEntryCount() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[4:6]))
}

// SetStateEntryCount updates the StateEntryCount field in the record's packed storage.
func (s *UCKeyStateRecord) SetStateEntryCount(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[4:6], uint16(v))
}

// StateEntryFormat returns the StateEntryFormat field from the record's packed storage.
func (s *UCKeyStateRecord) StateEntryFormat() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[6:8]))
}

// SetStateEntryFormat updates the StateEntryFormat field in the record's packed storage.
func (s *UCKeyStateRecord) SetStateEntryFormat(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[6:8], uint16(v))
}

// StateEntryData returns the StateEntryData field from the record's packed storage.
func (s *UCKeyStateRecord) StateEntryData() [1]uint32 {
	return *(*[1]uint32)(unsafe.Pointer(&s.storage[8]))
}

// SetStateEntryData updates the StateEntryData field in the record's packed storage.
func (s *UCKeyStateRecord) SetStateEntryData(v [1]uint32) {
	*(*[1]uint32)(unsafe.Pointer(&s.storage[8])) = v
}

// UCKeyStateRecordsIndex - Provides a count of, and offsets to, dead-key state records in a `'uchr'` resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/uckeystaterecordsindex
type UCKeyStateRecordsIndex struct {
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
	storage [8]byte
}

// KeyStateRecordsIndexFormat returns the KeyStateRecordsIndexFormat field from the record's packed storage.
func (s *UCKeyStateRecordsIndex) KeyStateRecordsIndexFormat() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[0:2]))
}

// SetKeyStateRecordsIndexFormat updates the KeyStateRecordsIndexFormat field in the record's packed storage.
func (s *UCKeyStateRecordsIndex) SetKeyStateRecordsIndexFormat(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[0:2], uint16(v))
}

// KeyStateRecordCount returns the KeyStateRecordCount field from the record's packed storage.
func (s *UCKeyStateRecordsIndex) KeyStateRecordCount() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[2:4]))
}

// SetKeyStateRecordCount updates the KeyStateRecordCount field in the record's packed storage.
func (s *UCKeyStateRecordsIndex) SetKeyStateRecordCount(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[2:4], uint16(v))
}

// KeyStateRecordOffsets returns the KeyStateRecordOffsets field from the record's packed storage.
func (s *UCKeyStateRecordsIndex) KeyStateRecordOffsets() [1]uint32 {
	return *(*[1]uint32)(unsafe.Pointer(&s.storage[4]))
}

// SetKeyStateRecordOffsets updates the KeyStateRecordOffsets field in the record's packed storage.
func (s *UCKeyStateRecordsIndex) SetKeyStateRecordOffsets(v [1]uint32) {
	*(*[1]uint32)(unsafe.Pointer(&s.storage[4])) = v
}

// UCKeyStateTerminators - Lists the default terminators for each dead-key state handled by a `'uchr'` resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/uckeystateterminators
type UCKeyStateTerminators struct {
	KeyStateTerminatorsFormat uint16          // An unsigned 16-bit integer identifying the format of the [UCKeyStateTerminators] structure. Set to `kUCKeyStateTerminatorsFormat`.
	KeyStateTerminatorCount   uint16          // An unsigned 16-bit integer specifying the number of default dead-key state terminators contained in the `keyStateTerminators[]` array.
	KeyStateTerminators       [1]UCKeyCharSeq // An array of default dead-key state terminators, described as values of type [UCKeyCharSeq](<https://developer.apple.com/documentation/coreservices/uckeycharseq>); the value `keyStateTerminators[0]` is the terminator for state 1, and so on.

}

// UCKeyToCharTableIndex - Provides a count of, and offsets to, key-code-to-character tables in a `'uchr'` resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/uckeytochartableindex
type UCKeyToCharTableIndex struct {
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
	storage [12]byte
}

// KeyToCharTableIndexFormat returns the KeyToCharTableIndexFormat field from the record's packed storage.
func (s *UCKeyToCharTableIndex) KeyToCharTableIndexFormat() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[0:2]))
}

// SetKeyToCharTableIndexFormat updates the KeyToCharTableIndexFormat field in the record's packed storage.
func (s *UCKeyToCharTableIndex) SetKeyToCharTableIndexFormat(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[0:2], uint16(v))
}

// KeyToCharTableSize returns the KeyToCharTableSize field from the record's packed storage.
func (s *UCKeyToCharTableIndex) KeyToCharTableSize() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[2:4]))
}

// SetKeyToCharTableSize updates the KeyToCharTableSize field in the record's packed storage.
func (s *UCKeyToCharTableIndex) SetKeyToCharTableSize(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[2:4], uint16(v))
}

// KeyToCharTableCount returns the KeyToCharTableCount field from the record's packed storage.
func (s *UCKeyToCharTableIndex) KeyToCharTableCount() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetKeyToCharTableCount updates the KeyToCharTableCount field in the record's packed storage.
func (s *UCKeyToCharTableIndex) SetKeyToCharTableCount(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// KeyToCharTableOffsets returns the KeyToCharTableOffsets field from the record's packed storage.
func (s *UCKeyToCharTableIndex) KeyToCharTableOffsets() [1]uint32 {
	return *(*[1]uint32)(unsafe.Pointer(&s.storage[8]))
}

// SetKeyToCharTableOffsets updates the KeyToCharTableOffsets field in the record's packed storage.
func (s *UCKeyToCharTableIndex) SetKeyToCharTableOffsets(v [1]uint32) {
	*(*[1]uint32)(unsafe.Pointer(&s.storage[8])) = v
}

// UCKeyboardLayout - Provides header data for a `'uchr'` resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/uckeyboardlayout
type UCKeyboardLayout struct {
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
	storage [40]byte
}

// KeyLayoutHeaderFormat returns the KeyLayoutHeaderFormat field from the record's packed storage.
func (s *UCKeyboardLayout) KeyLayoutHeaderFormat() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[0:2]))
}

// SetKeyLayoutHeaderFormat updates the KeyLayoutHeaderFormat field in the record's packed storage.
func (s *UCKeyboardLayout) SetKeyLayoutHeaderFormat(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[0:2], uint16(v))
}

// KeyLayoutDataVersion returns the KeyLayoutDataVersion field from the record's packed storage.
func (s *UCKeyboardLayout) KeyLayoutDataVersion() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[2:4]))
}

// SetKeyLayoutDataVersion updates the KeyLayoutDataVersion field in the record's packed storage.
func (s *UCKeyboardLayout) SetKeyLayoutDataVersion(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[2:4], uint16(v))
}

// KeyLayoutFeatureInfoOffset returns the KeyLayoutFeatureInfoOffset field from the record's packed storage.
func (s *UCKeyboardLayout) KeyLayoutFeatureInfoOffset() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetKeyLayoutFeatureInfoOffset updates the KeyLayoutFeatureInfoOffset field in the record's packed storage.
func (s *UCKeyboardLayout) SetKeyLayoutFeatureInfoOffset(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// KeyboardTypeCount returns the KeyboardTypeCount field from the record's packed storage.
func (s *UCKeyboardLayout) KeyboardTypeCount() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetKeyboardTypeCount updates the KeyboardTypeCount field in the record's packed storage.
func (s *UCKeyboardLayout) SetKeyboardTypeCount(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[8:12], uint32(v))
}

// KeyboardTypeList returns the KeyboardTypeList field from the record's packed storage.
func (s *UCKeyboardLayout) KeyboardTypeList() [1]UCKeyboardTypeHeader {
	return *(*[1]UCKeyboardTypeHeader)(unsafe.Pointer(&s.storage[12]))
}

// SetKeyboardTypeList updates the KeyboardTypeList field in the record's packed storage.
func (s *UCKeyboardLayout) SetKeyboardTypeList(v [1]UCKeyboardTypeHeader) {
	*(*[1]UCKeyboardTypeHeader)(unsafe.Pointer(&s.storage[12])) = v
}

// UCKeyboardTypeHeader - Specifies a range of physical keyboard types in a `'uchr'` resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/uckeyboardtypeheader
type UCKeyboardTypeHeader struct {
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
	storage [28]byte
}

// KeyboardTypeFirst returns the KeyboardTypeFirst field from the record's packed storage.
func (s *UCKeyboardTypeHeader) KeyboardTypeFirst() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetKeyboardTypeFirst updates the KeyboardTypeFirst field in the record's packed storage.
func (s *UCKeyboardTypeHeader) SetKeyboardTypeFirst(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// KeyboardTypeLast returns the KeyboardTypeLast field from the record's packed storage.
func (s *UCKeyboardTypeHeader) KeyboardTypeLast() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetKeyboardTypeLast updates the KeyboardTypeLast field in the record's packed storage.
func (s *UCKeyboardTypeHeader) SetKeyboardTypeLast(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// KeyModifiersToTableNumOffset returns the KeyModifiersToTableNumOffset field from the record's packed storage.
func (s *UCKeyboardTypeHeader) KeyModifiersToTableNumOffset() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetKeyModifiersToTableNumOffset updates the KeyModifiersToTableNumOffset field in the record's packed storage.
func (s *UCKeyboardTypeHeader) SetKeyModifiersToTableNumOffset(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[8:12], uint32(v))
}

// KeyToCharTableIndexOffset returns the KeyToCharTableIndexOffset field from the record's packed storage.
func (s *UCKeyboardTypeHeader) KeyToCharTableIndexOffset() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[12:16]))
}

// SetKeyToCharTableIndexOffset updates the KeyToCharTableIndexOffset field in the record's packed storage.
func (s *UCKeyboardTypeHeader) SetKeyToCharTableIndexOffset(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[12:16], uint32(v))
}

// KeyStateRecordsIndexOffset returns the KeyStateRecordsIndexOffset field from the record's packed storage.
func (s *UCKeyboardTypeHeader) KeyStateRecordsIndexOffset() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[16:20]))
}

// SetKeyStateRecordsIndexOffset updates the KeyStateRecordsIndexOffset field in the record's packed storage.
func (s *UCKeyboardTypeHeader) SetKeyStateRecordsIndexOffset(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[16:20], uint32(v))
}

// KeyStateTerminatorsOffset returns the KeyStateTerminatorsOffset field from the record's packed storage.
func (s *UCKeyboardTypeHeader) KeyStateTerminatorsOffset() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[20:24]))
}

// SetKeyStateTerminatorsOffset updates the KeyStateTerminatorsOffset field in the record's packed storage.
func (s *UCKeyboardTypeHeader) SetKeyStateTerminatorsOffset(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[20:24], uint32(v))
}

// KeySequenceDataIndexOffset returns the KeySequenceDataIndexOffset field from the record's packed storage.
func (s *UCKeyboardTypeHeader) KeySequenceDataIndexOffset() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[24:28]))
}

// SetKeySequenceDataIndexOffset updates the KeySequenceDataIndexOffset field in the record's packed storage.
func (s *UCKeyboardTypeHeader) SetKeySequenceDataIndexOffset(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[24:28], uint32(v))
}

// WritingCode
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/writingcode
type WritingCode struct {
	TheScriptCode int16
	TheLangCode   int16
}

// CcntTokenRecord - Stores token information used by the AEResolve functionwhile locating a range of objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/ccnttokenrecord
type CcntTokenRecord struct {
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

// TokenClass returns the TokenClass field from the record's packed storage.
func (s *CcntTokenRecord) TokenClass() DescType {
	return DescType(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetTokenClass updates the TokenClass field in the record's packed storage.
func (s *CcntTokenRecord) SetTokenClass(v DescType) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// Token returns the Token field from the record's packed storage.
func (s *CcntTokenRecord) Token() AEDesc {
	return *(*AEDesc)(unsafe.Pointer(&s.storage[4]))
}

// SetToken updates the Token field in the record's packed storage.
func (s *CcntTokenRecord) SetToken(v AEDesc) {
	*(*AEDesc)(unsafe.Pointer(&s.storage[4])) = v
}
