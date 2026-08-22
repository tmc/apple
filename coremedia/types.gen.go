// Code generated from Apple documentation for CoreMedia. DO NOT EDIT.

package coremedia

import (
	"encoding/binary"
	"unsafe"

	"github.com/tmc/apple/corefoundation"
)

// C struct types

// CMBlockBufferCustomBlockSource - A structure to support custom memory allocation and deallocation for a block used in a block buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferCustomBlockSource
type CMBlockBufferCustomBlockSource struct {
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

// Version returns the Version field from the record's packed storage.
func (s *CMBlockBufferCustomBlockSource) Version() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetVersion updates the Version field in the record's packed storage.
func (s *CMBlockBufferCustomBlockSource) SetVersion(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// AllocateBlock returns a raw C code address from the record's packed storage.
// The address is not callable from Go. To call C through it, use purego to
// bind a Go function value to the address.
func (s *CMBlockBufferCustomBlockSource) AllocateBlock() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[4:12]))
}

// SetAllocateBlock stores a raw C code address in the record's packed storage.
// Use purego.NewCallback to obtain the address for a Go callback. The callback
// and its address must outlive every use C can make through this record.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *CMBlockBufferCustomBlockSource) SetAllocateBlock(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[4:12], uint64(v))
}

// FreeBlock returns a raw C code address from the record's packed storage.
// The address is not callable from Go. To call C through it, use purego to
// bind a Go function value to the address.
func (s *CMBlockBufferCustomBlockSource) FreeBlock() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[12:20]))
}

// SetFreeBlock stores a raw C code address in the record's packed storage.
// Use purego.NewCallback to obtain the address for a Go callback. The callback
// and its address must outlive every use C can make through this record.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *CMBlockBufferCustomBlockSource) SetFreeBlock(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[12:20], uint64(v))
}

// RefCon returns the RefCon field from the record's packed storage.
func (s *CMBlockBufferCustomBlockSource) RefCon() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[20:28]))
}

// SetRefCon updates the RefCon field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *CMBlockBufferCustomBlockSource) SetRefCon(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[20:28], uint64(v))
}

// CMBufferCallbacks - A structure that stores the callbacks that perform buffer operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferCallbacks
type CMBufferCallbacks struct {
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
	storage [68]byte
}

// Version returns the Version field from the record's packed storage.
func (s *CMBufferCallbacks) Version() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetVersion updates the Version field in the record's packed storage.
func (s *CMBufferCallbacks) SetVersion(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// Refcon returns the Refcon field from the record's packed storage.
func (s *CMBufferCallbacks) Refcon() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[4:12]))
}

// SetRefcon updates the Refcon field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *CMBufferCallbacks) SetRefcon(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[4:12], uint64(v))
}

// GetDecodeTimeStamp returns a raw C code address from the record's packed storage.
// The address is not callable from Go. To call C through it, use purego to
// bind a Go function value to the address.
func (s *CMBufferCallbacks) GetDecodeTimeStamp() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[12:20]))
}

// SetGetDecodeTimeStamp stores a raw C code address in the record's packed storage.
// Use purego.NewCallback to obtain the address for a Go callback. The callback
// and its address must outlive every use C can make through this record.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *CMBufferCallbacks) SetGetDecodeTimeStamp(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[12:20], uint64(v))
}

// GetPresentationTimeStamp returns a raw C code address from the record's packed storage.
// The address is not callable from Go. To call C through it, use purego to
// bind a Go function value to the address.
func (s *CMBufferCallbacks) GetPresentationTimeStamp() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[20:28]))
}

// SetGetPresentationTimeStamp stores a raw C code address in the record's packed storage.
// Use purego.NewCallback to obtain the address for a Go callback. The callback
// and its address must outlive every use C can make through this record.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *CMBufferCallbacks) SetGetPresentationTimeStamp(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[20:28], uint64(v))
}

// GetDuration returns a raw C code address from the record's packed storage.
// The address is not callable from Go. To call C through it, use purego to
// bind a Go function value to the address.
func (s *CMBufferCallbacks) GetDuration() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[28:36]))
}

// SetGetDuration stores a raw C code address in the record's packed storage.
// Use purego.NewCallback to obtain the address for a Go callback. The callback
// and its address must outlive every use C can make through this record.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *CMBufferCallbacks) SetGetDuration(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[28:36], uint64(v))
}

// IsDataReady returns a raw C code address from the record's packed storage.
// The address is not callable from Go. To call C through it, use purego to
// bind a Go function value to the address.
func (s *CMBufferCallbacks) IsDataReady() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[36:44]))
}

// SetIsDataReady stores a raw C code address in the record's packed storage.
// Use purego.NewCallback to obtain the address for a Go callback. The callback
// and its address must outlive every use C can make through this record.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *CMBufferCallbacks) SetIsDataReady(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[36:44], uint64(v))
}

// Compare returns a raw C code address from the record's packed storage.
// The address is not callable from Go. To call C through it, use purego to
// bind a Go function value to the address.
func (s *CMBufferCallbacks) Compare() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[44:52]))
}

// SetCompare stores a raw C code address in the record's packed storage.
// Use purego.NewCallback to obtain the address for a Go callback. The callback
// and its address must outlive every use C can make through this record.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *CMBufferCallbacks) SetCompare(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[44:52], uint64(v))
}

// DataBecameReadyNotification returns the DataBecameReadyNotification field from the record's packed storage.
func (s *CMBufferCallbacks) DataBecameReadyNotification() corefoundation.CFStringRef {
	return corefoundation.CFStringRef(binary.NativeEndian.Uint64(s.storage[52:60]))
}

// SetDataBecameReadyNotification updates the DataBecameReadyNotification field in the record's packed storage.
func (s *CMBufferCallbacks) SetDataBecameReadyNotification(v corefoundation.CFStringRef) {
	binary.NativeEndian.PutUint64(s.storage[52:60], uint64(v))
}

// GetSize returns a raw C code address from the record's packed storage.
// The address is not callable from Go. To call C through it, use purego to
// bind a Go function value to the address.
func (s *CMBufferCallbacks) GetSize() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[60:68]))
}

// SetGetSize stores a raw C code address in the record's packed storage.
// Use purego.NewCallback to obtain the address for a Go callback. The callback
// and its address must outlive every use C can make through this record.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *CMBufferCallbacks) SetGetSize(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[60:68], uint64(v))
}

// CMBufferHandlers - A structure that stores the handlers that perform buffer operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferHandlers
type CMBufferHandlers struct {
	Version                     uintptr // The version number.
	GetDecodeTimeStamp          CMBufferGetTimeHandler
	GetPresentationTimeStamp    CMBufferGetTimeHandler
	GetDuration                 CMBufferGetTimeHandler
	IsDataReady                 CMBufferGetBooleanHandler
	Compare                     CMBufferCompareHandler // A handler callback the queue uses to perform an insertion sort of the queue.
	DataBecameReadyNotification corefoundation.CFStringRef
	GetSize                     CMBufferGetSizeHandler
}

// CMSampleTimingInfo - A collection of timing information for a sample in a sample buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleTimingInfo
type CMSampleTimingInfo struct {
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
	storage [72]byte
}

// Duration returns the Duration field from the record's packed storage.
func (s *CMSampleTimingInfo) Duration() CMTime {
	return *(*CMTime)(unsafe.Pointer(&s.storage[0]))
}

// SetDuration updates the Duration field in the record's packed storage.
func (s *CMSampleTimingInfo) SetDuration(v CMTime) {
	*(*CMTime)(unsafe.Pointer(&s.storage[0])) = v
}

// PresentationTimeStamp returns the PresentationTimeStamp field from the record's packed storage.
func (s *CMSampleTimingInfo) PresentationTimeStamp() CMTime {
	return *(*CMTime)(unsafe.Pointer(&s.storage[24]))
}

// SetPresentationTimeStamp updates the PresentationTimeStamp field in the record's packed storage.
func (s *CMSampleTimingInfo) SetPresentationTimeStamp(v CMTime) {
	*(*CMTime)(unsafe.Pointer(&s.storage[24])) = v
}

// DecodeTimeStamp returns the DecodeTimeStamp field from the record's packed storage.
func (s *CMSampleTimingInfo) DecodeTimeStamp() CMTime {
	return *(*CMTime)(unsafe.Pointer(&s.storage[48]))
}

// SetDecodeTimeStamp updates the DecodeTimeStamp field in the record's packed storage.
func (s *CMSampleTimingInfo) SetDecodeTimeStamp(v CMTime) {
	*(*CMTime)(unsafe.Pointer(&s.storage[48])) = v
}

// CMTag - A tag representing additional metadata on tagged media buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTag-c.struct
type CMTag struct {
	Category CMTagCategory // The category assigned to a tag.
	DataType CMTagDataType // The data type for the value stored in the tag.
	Value    CMTagValue    // The value of the tag.

}

// CMTime - A structure that represents time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTime
type CMTime struct {
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

// Value returns the Value field from the record's packed storage.
func (s *CMTime) Value() CMTimeValue {
	return CMTimeValue(binary.NativeEndian.Uint64(s.storage[0:8]))
}

// SetValue updates the Value field in the record's packed storage.
func (s *CMTime) SetValue(v CMTimeValue) {
	binary.NativeEndian.PutUint64(s.storage[0:8], uint64(v))
}

// Timescale returns the Timescale field from the record's packed storage.
func (s *CMTime) Timescale() CMTimeScale {
	return CMTimeScale(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetTimescale updates the Timescale field in the record's packed storage.
func (s *CMTime) SetTimescale(v CMTimeScale) {
	binary.NativeEndian.PutUint32(s.storage[8:12], uint32(v))
}

// Flags returns the Flags field from the record's packed storage.
func (s *CMTime) Flags() CMTimeFlags {
	return *(*CMTimeFlags)(unsafe.Pointer(&s.storage[12]))
}

// SetFlags updates the Flags field in the record's packed storage.
func (s *CMTime) SetFlags(v CMTimeFlags) {
	*(*CMTimeFlags)(unsafe.Pointer(&s.storage[12])) = v
}

// Epoch returns the Epoch field from the record's packed storage.
func (s *CMTime) Epoch() CMTimeEpoch {
	return CMTimeEpoch(binary.NativeEndian.Uint64(s.storage[16:24]))
}

// SetEpoch updates the Epoch field in the record's packed storage.
func (s *CMTime) SetEpoch(v CMTimeEpoch) {
	binary.NativeEndian.PutUint64(s.storage[16:24], uint64(v))
}

// CMTimeMapping - A structure that maps a segment of a source time range to a target time range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeMapping
type CMTimeMapping struct {
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
	storage [96]byte
}

// Source returns the Source field from the record's packed storage.
func (s *CMTimeMapping) Source() CMTimeRange {
	return *(*CMTimeRange)(unsafe.Pointer(&s.storage[0]))
}

// SetSource updates the Source field in the record's packed storage.
func (s *CMTimeMapping) SetSource(v CMTimeRange) {
	*(*CMTimeRange)(unsafe.Pointer(&s.storage[0])) = v
}

// Target returns the Target field from the record's packed storage.
func (s *CMTimeMapping) Target() CMTimeRange {
	return *(*CMTimeRange)(unsafe.Pointer(&s.storage[48]))
}

// SetTarget updates the Target field in the record's packed storage.
func (s *CMTimeMapping) SetTarget(v CMTimeRange) {
	*(*CMTimeRange)(unsafe.Pointer(&s.storage[48])) = v
}

// CMTimeRange - A structure that represents a time range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeRange
type CMTimeRange struct {
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
	storage [48]byte
}

// Start returns the Start field from the record's packed storage.
func (s *CMTimeRange) Start() CMTime {
	return *(*CMTime)(unsafe.Pointer(&s.storage[0]))
}

// SetStart updates the Start field in the record's packed storage.
func (s *CMTimeRange) SetStart(v CMTime) {
	*(*CMTime)(unsafe.Pointer(&s.storage[0])) = v
}

// Duration returns the Duration field from the record's packed storage.
func (s *CMTimeRange) Duration() CMTime {
	return *(*CMTime)(unsafe.Pointer(&s.storage[24]))
}

// SetDuration updates the Duration field in the record's packed storage.
func (s *CMTimeRange) SetDuration(v CMTime) {
	*(*CMTime)(unsafe.Pointer(&s.storage[24])) = v
}

// CMVideoDimensions - A structure that represents video dimensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMVideoDimensions
type CMVideoDimensions struct {
	Width  int32 // The width of the video.
	Height int32 // The height of the video.

}
