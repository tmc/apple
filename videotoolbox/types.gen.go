// Code generated from Apple documentation for VideoToolbox. DO NOT EDIT.

package videotoolbox

import (
	"encoding/binary"
)

// C struct types

// VTDecompressionOutputCallbackRecord
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/VideoToolbox/VTDecompressionOutputCallbackRecord
type VTDecompressionOutputCallbackRecord struct {
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

// DecompressionOutputCallback returns a raw C code address from the record's packed storage.
// The address is not callable from Go. To call C through it, use purego to
// bind a Go function value to the address.
func (s *VTDecompressionOutputCallbackRecord) DecompressionOutputCallback() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[0:8]))
}

// SetDecompressionOutputCallback stores a raw C code address in the record's packed storage.
// Use purego.NewCallback to obtain the address for a Go callback. The callback
// and its address must outlive every use C can make through this record.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *VTDecompressionOutputCallbackRecord) SetDecompressionOutputCallback(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[0:8], uint64(v))
}

// DecompressionOutputRefCon returns the DecompressionOutputRefCon field from the record's packed storage.
func (s *VTDecompressionOutputCallbackRecord) DecompressionOutputRefCon() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[8:16]))
}

// SetDecompressionOutputRefCon updates the DecompressionOutputRefCon field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *VTDecompressionOutputCallbackRecord) SetDecompressionOutputRefCon(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[8:16], uint64(v))
}

// VTInt32Point - A structure that represents a 32-bit integer point value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/VideoToolbox/VTInt32Point
type VTInt32Point struct {
	X int32 // The x-coordinate of the point.
	Y int32 // The y-coordinate of the point.

}

// VTInt32Size - A structure that represents a 32-bit integer size value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/VideoToolbox/VTInt32Size
type VTInt32Size struct {
	Width  int32 // The width of the size.
	Height int32 // The height of the size.

}
