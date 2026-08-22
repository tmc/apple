//go:build darwin

package mil

import (
	"encoding/binary"
	"fmt"
	"math"
)

// BlobDataType identifies the element type in a weight blob entry.
// The values are the MIL Blob Storage type codes.
type BlobDataType uint32

const (
	BlobFloat16 BlobDataType = 1
	BlobFloat32 BlobDataType = 2
	BlobUInt8   BlobDataType = 3
	BlobInt8    BlobDataType = 4
)

const (
	blobAlignment  = 64
	blobSentinel   = 0xDEADBEEF
	blobVersion    = 2
	blobHeaderSize = 64
	blobMetaSize   = 64
)

type blobEntry struct {
	dtype BlobDataType
	data  []byte
}

// BlobWriter accumulates weight blobs and builds a MIL Blob Storage v2 binary.
//
// A file begins with a 64-byte header holding the entry count and the format
// version. Each entry follows at a 64-byte-aligned position as a 64-byte
// metadata block immediately followed by its raw data.
type BlobWriter struct {
	blobs []blobEntry
}

// NewBlobWriter returns an empty BlobWriter.
func NewBlobWriter() *BlobWriter {
	return &BlobWriter{}
}

// AddFloat16 converts data to fp16 and appends it as a blob entry.
// It returns the blob index; pass it to Offset once all blobs are added.
func (w *BlobWriter) AddFloat16(data []float32) int {
	return w.AddRaw(BlobFloat16, fp16Bytes(data))
}

// AddFloat32 appends data as a blob entry.
// It returns the blob index; pass it to Offset once all blobs are added.
func (w *BlobWriter) AddFloat32(data []float32) int {
	raw := make([]byte, len(data)*4)
	for i, v := range data {
		binary.LittleEndian.PutUint32(raw[i*4:], math.Float32bits(v))
	}
	return w.AddRaw(BlobFloat32, raw)
}

// AddRaw appends raw byte data as a blob entry.
// It returns the blob index; pass it to Offset once all blobs are added.
func (w *BlobWriter) AddRaw(dtype BlobDataType, data []byte) int {
	idx := len(w.blobs)
	w.blobs = append(w.blobs, blobEntry{dtype: dtype, data: data})
	return idx
}

// Offset returns the byte offset of blob i's metadata block, which is the
// value a BLOBFILE(offset = uint64(...)) reference in MIL text takes. The raw
// data starts 64 bytes later; the metadata records that address.
//
// The offset depends on every preceding blob, so call Offset only after all
// blobs have been added.
func (w *BlobWriter) Offset(i int) uint64 {
	if i < 0 || i >= len(w.blobs) {
		return 0
	}
	off := blobHeaderSize
	for j := range i {
		off = alignUp(off, blobAlignment) + blobMetaSize + len(w.blobs[j].data)
	}
	return uint64(alignUp(off, blobAlignment))
}

// Count returns the number of blobs added.
func (w *BlobWriter) Count() int {
	return len(w.blobs)
}

// Build produces the complete binary blob.
func (w *BlobWriter) Build() ([]byte, error) {
	if len(w.blobs) == 0 {
		return nil, fmt.Errorf("mil: no blobs to build")
	}
	return buildBlob(w.blobs), nil
}

// buildBlob encodes entries as a MIL Blob Storage v2 file.
func buildBlob(blobs []blobEntry) []byte {
	buf := make([]byte, blobHeaderSize)
	binary.LittleEndian.PutUint32(buf[0:], uint32(len(blobs)))
	binary.LittleEndian.PutUint32(buf[4:], blobVersion)

	for _, b := range blobs {
		buf = append(buf, make([]byte, alignUp(len(buf), blobAlignment)-len(buf))...)
		metaOff := len(buf)

		meta := make([]byte, blobMetaSize)
		binary.LittleEndian.PutUint32(meta[0:], blobSentinel)
		binary.LittleEndian.PutUint32(meta[4:], uint32(b.dtype))
		binary.LittleEndian.PutUint64(meta[8:], uint64(len(b.data)))
		binary.LittleEndian.PutUint64(meta[16:], uint64(metaOff+blobMetaSize))
		buf = append(buf, meta...)
		buf = append(buf, b.data...)
	}
	return buf
}

// fp16Blob returns a one-entry blob file holding the given fp16 payload.
// Its BLOBFILE offset is 64.
func fp16Blob(fp16Data []byte) []byte {
	return buildBlob([]blobEntry{{dtype: BlobFloat16, data: fp16Data}})
}

// fp16Bytes converts data to little-endian fp16.
func fp16Bytes(data []float32) []byte {
	raw := make([]byte, len(data)*2)
	for i, v := range data {
		binary.LittleEndian.PutUint16(raw[i*2:], float32ToFP16(v))
	}
	return raw
}

func alignUp(n, align int) int {
	return (n + align - 1) &^ (align - 1)
}
