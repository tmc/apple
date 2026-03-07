// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"errors"
	"io"
	"unsafe"
	"github.com/tmc/apple/objc"
)

// NSErrorToError wraps an *NSError as a Go error.
// Returns nil if nserr is nil.
func NSErrorToError(nserr *NSError) error {
	if nserr == nil {
		return nil
	}
	return nserr
}

// NSErrorFrom converts an Objective-C NSError object ID to Go error.
// Returns nil for a zero object ID.
func NSErrorFrom(errorPtr objc.ID) error {
	if errorPtr == 0 {
		return nil
	}
	v := NSErrorFromID(errorPtr)
	return &v
}

// ErrorFrom is a compatibility alias for NSErrorFrom.
func ErrorFrom(errorPtr objc.ID) error {
	return NSErrorFrom(errorPtr)
}

// NewDataFromBytes creates an NSData from a Go byte slice.
// The bytes are copied into the Objective-C heap.
func NewDataFromBytes(b []byte) NSData {
	if len(b) == 0 {
		return NewNSData()
	}
	return NewDataWithBytesLength(b)
}

// GoBytes returns the NSData contents as a Go byte slice.
// The returned slice is an independent copy.
func (n NSData) GoBytes() []byte {
	length := n.Length()
	if length == 0 {
		return nil
	}
	return append([]byte(nil), unsafe.Slice((*byte)(n.Bytes()), int(length))...)
}

// NSDataReader wraps an NSData as an io.Reader and io.Seeker.
type NSDataReader struct {
	data NSData
	off  int64
}

// Reader returns an io.ReadSeeker over the NSData contents.
func (n NSData) Reader() *NSDataReader {
	return &NSDataReader{data: n}
}

func (r *NSDataReader) Read(p []byte) (int, error) {
	length := int64(r.data.Length())
	if r.off >= length {
		return 0, io.EOF
	}
	n := copy(p, unsafe.Slice((*byte)(r.data.Bytes()), int(length))[r.off:])
	r.off += int64(n)
	if r.off >= length {
		return n, io.EOF
	}
	return n, nil
}

func (r *NSDataReader) Seek(offset int64, whence int) (int64, error) {
	length := int64(r.data.Length())
	var abs int64
	switch whence {
	case io.SeekStart:
		abs = offset
	case io.SeekCurrent:
		abs = r.off + offset
	case io.SeekEnd:
		abs = length + offset
	default:
		return 0, errors.New("foundation.NSDataReader.Seek: invalid whence")
	}
	if abs < 0 {
		return 0, errors.New("foundation.NSDataReader.Seek: negative position")
	}
	r.off = abs
	return abs, nil
}

// Write appends p to the mutable data, implementing io.Writer.
func (n NSMutableData) Write(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	n.AppendBytesLength(p)
	return len(p), nil
}

