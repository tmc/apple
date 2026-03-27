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
// Returns nil for a zero object ID. The NSError is retained so it
// survives autorelease pool drain when sent across goroutine boundaries.
func NSErrorFrom(errorPtr objc.ID) error {
	if errorPtr == 0 {
		return nil
	}
	objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
	v := NSErrorFromID(errorPtr)
	return &v
}

// ErrorFrom is a compatibility alias for NSErrorFrom.
func ErrorFrom(errorPtr objc.ID) error {
	return NSErrorFrom(errorPtr)
}

// SafeErrorFrom converts an Objective-C object ID to Go error, checking at
// runtime whether the object is actually an NSError. If not (e.g., a private
// framework passes NSString or another type where NSError is declared), it
// falls back to calling -description and wrapping the result as a plain error.
//
// The object is retained so it survives autorelease pool drain. This is
// necessary when the error is sent across goroutine boundaries (e.g.,
// from an async completion handler block to a channel consumer).
func SafeErrorFrom(errorPtr objc.ID) error {
	if errorPtr == 0 {
		return nil
	}
	// Retain so the object survives autorelease pool drain on the
	// calling thread (critical for async completion handler blocks).
	objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
	nsErrorClass := objc.GetClass("NSError")
	if nsErrorClass != 0 && objc.Send[bool](errorPtr, objc.Sel("isKindOfClass:"), nsErrorClass) {
		v := NSErrorFromID(errorPtr)
		return &v
	}
	// Not an NSError — use -description as the error message.
	desc := objc.Send[*byte](errorPtr, objc.Sel("description"))
	objc.Send[objc.ID](errorPtr, objc.Sel("release"))
	if desc != nil {
		return errors.New(objc.GoString(desc))
	}
	return errors.New("unknown Objective-C error")
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

// NSDictionaryToMap converts an NSDictionary to a Go map using converter functions.
func NSDictionaryToMap[K comparable, V any](dict NSDictionary, fromK func(objc.ID) K, fromV func(objc.ID) V) map[K]V {
	if dict.GetID() == 0 {
		return nil
	}
	keys := dict.AllKeys()
	m := make(map[K]V, len(keys))
	for _, k := range keys {
		v := dict.ObjectForKey(k)
		m[fromK(k.GetID())] = fromV(v.GetID())
	}
	return m
}

// NSDictionaryToStringMap converts an NSDictionary with NSString keys to a map[string]V.
func NSDictionaryToStringMap[V any](dict NSDictionary, fromV func(objc.ID) V) map[string]V {
	return NSDictionaryToMap(dict, objc.IDToString, fromV)
}

// NSSetToSlice converts an NSSet to a typed Go slice using a converter function.
func NSSetToSlice[T any](set NSSet, from func(objc.ID) T) []T {
	if set.GetID() == 0 {
		return nil
	}
	objs := set.AllObjects()
	result := make([]T, len(objs))
	for i, o := range objs {
		result[i] = from(o.GetID())
	}
	return result
}

