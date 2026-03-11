//go:build darwin

package mil

import (
	"encoding/binary"
	"fmt"
	"math"
)

// BlobDataType identifies the element type in a weight blob entry.
type BlobDataType uint32

const (
	BlobFloat16 BlobDataType = 1
	BlobFloat32 BlobDataType = 2
	BlobUInt8   BlobDataType = 3
	BlobInt8    BlobDataType = 8
)

const (
	blobAlignment  = 64
	blobMagic      = 0xDEADBEEF
	blobVersion    = 2
	blobFileHeader = 64
	blobChunkMeta  = 64
)

type blobEntry struct {
	dtype BlobDataType
	data  []byte
}

// BlobWriter accumulates weight blobs and builds a multi-weight
// MIL Blob Storage v2 binary.
//
// The format consists of a 64-byte file header, followed by 64-byte
// per-blob metadata entries, then 64-byte-aligned data segments.
type BlobWriter struct {
	blobs []blobEntry
}

// NewBlobWriter creates a new BlobWriter.
func NewBlobWriter() *BlobWriter {
	return &BlobWriter{}
}

// AddFloat16 converts float32 data to fp16 and appends it as a blob entry.
// Returns the blob index. Use Offset after all blobs are added to get the data offset.
func (w *BlobWriter) AddFloat16(data []float32) int {
	raw := make([]byte, len(data)*2)
	for i, v := range data {
		binary.LittleEndian.PutUint16(raw[i*2:], float32ToFP16(v))
	}
	return w.AddRaw(BlobFloat16, raw)
}

// AddFloat32 appends float32 data as a blob entry.
// Returns the blob index. Use Offset after all blobs are added to get the data offset.
func (w *BlobWriter) AddFloat32(data []float32) int {
	raw := make([]byte, len(data)*4)
	for i, v := range data {
		binary.LittleEndian.PutUint32(raw[i*4:], math.Float32bits(v))
	}
	return w.AddRaw(BlobFloat32, raw)
}

// AddRaw appends raw byte data as a blob entry.
// Returns the blob index. Use Offset after all blobs are added to get the data offset.
func (w *BlobWriter) AddRaw(dtype BlobDataType, data []byte) int {
	idx := len(w.blobs)
	w.blobs = append(w.blobs, blobEntry{dtype: dtype, data: data})
	return idx
}

// Offset returns the byte offset where blob i's data starts in the built output.
// This must be called after all blobs have been added, as the offset depends
// on the total number of blobs (which determines the metadata section size).
func (w *BlobWriter) Offset(i int) uint64 {
	if i < 0 || i >= len(w.blobs) {
		return 0
	}
	metaSize := alignUp(blobFileHeader+blobChunkMeta*len(w.blobs), blobAlignment)
	offset := metaSize
	for j := range i {
		offset += alignUp(len(w.blobs[j].data), blobAlignment)
	}
	return uint64(offset)
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

	// Compute total size.
	metaSize := alignUp(blobFileHeader+blobChunkMeta*len(w.blobs), blobAlignment)
	totalSize := metaSize
	for _, b := range w.blobs {
		totalSize += alignUp(len(b.data), blobAlignment)
	}

	buf := make([]byte, totalSize)

	// File header.
	binary.LittleEndian.PutUint64(buf[0:], uint64(len(w.blobs))) // count
	binary.LittleEndian.PutUint32(buf[8:], blobVersion)          // version

	// Per-blob metadata and data.
	dataOffset := metaSize
	for i, b := range w.blobs {
		metaOff := blobFileHeader + blobChunkMeta*i

		binary.LittleEndian.PutUint32(buf[metaOff:], blobMagic)              // sentinel
		binary.LittleEndian.PutUint32(buf[metaOff+4:], uint32(b.dtype))      // dtype
		binary.LittleEndian.PutUint64(buf[metaOff+8:], uint64(len(b.data)))  // size
		binary.LittleEndian.PutUint64(buf[metaOff+16:], uint64(dataOffset))  // offset

		copy(buf[dataOffset:], b.data)
		dataOffset += alignUp(len(b.data), blobAlignment)
	}

	return buf, nil
}

func alignUp(n, align int) int {
	return (n + align - 1) &^ (align - 1)
}
