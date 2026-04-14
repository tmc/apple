package coremlcompiler

import (
	"encoding/binary"
	"fmt"
)

// MILBlob Storage v2 format for CoreML weight.bin files.
//
// The format consists of a 64-byte storage header followed by
// 64-byte-aligned entries, each with a 64-byte metadata block
// and the raw tensor data.
//
// BLOBFILE offsets in MIL text point to the blob_metadata block
// (not the raw data). The metadata contains the absolute file
// offset of the raw data.

const (
	milBlobAlignment     = 64
	milBlobVersion       = 2
	milBlobSentinel      = 0xDEADBEEF
	milBlobHeaderSize    = 64
	milBlobMetadataSize  = 64
)

// BlobDataType identifies element types in MIL blob storage.
type BlobDataType uint32

const (
	BlobDataTypeFloat16    BlobDataType = 1
	BlobDataTypeFloat32    BlobDataType = 2
	BlobDataTypeUInt8      BlobDataType = 3
	BlobDataTypeInt8       BlobDataType = 4
	BlobDataTypeBFloat16   BlobDataType = 5
	BlobDataTypeInt16      BlobDataType = 6
	BlobDataTypeUInt16     BlobDataType = 7
	BlobDataTypeInt4       BlobDataType = 8
	BlobDataTypeUInt1      BlobDataType = 9
	BlobDataTypeUInt2      BlobDataType = 10
	BlobDataTypeUInt4      BlobDataType = 11
	BlobDataTypeUInt3      BlobDataType = 12
	BlobDataTypeUInt6      BlobDataType = 13
	BlobDataTypeInt32      BlobDataType = 14
	BlobDataTypeUInt32     BlobDataType = 15
	BlobDataTypeFloat8E4M3 BlobDataType = 16
	BlobDataTypeFloat8E5M2 BlobDataType = 17
)

// MILBlobDataTypeName returns the MIL DataType enum for a BlobDataType.
func DataTypeToBlobDataType(dt DataType) (BlobDataType, error) {
	switch dt {
	case DataTypeFloat16:
		return BlobDataTypeFloat16, nil
	case DataTypeFloat32:
		return BlobDataTypeFloat32, nil
	case DataTypeBFloat16:
		return BlobDataTypeBFloat16, nil
	case DataTypeInt8:
		return BlobDataTypeInt8, nil
	case DataTypeInt16:
		return BlobDataTypeInt16, nil
	case DataTypeInt32:
		return BlobDataTypeInt32, nil
	case DataTypeUInt8:
		return BlobDataTypeUInt8, nil
	case DataTypeUInt16:
		return BlobDataTypeUInt16, nil
	case DataTypeUInt32:
		return BlobDataTypeUInt32, nil
	case DataTypeInt4:
		return BlobDataTypeInt4, nil
	default:
		return 0, fmt.Errorf("unsupported data type for MILBlob: %v", dt)
	}
}

// BlobEntry describes a single tensor to write into a MILBlob weight file.
type BlobEntry struct {
	DType BlobDataType
	Data  []byte
}

// WriteMILBlob builds a MIL Blob Storage v2 weight file from the given
// entries. Returns the complete file bytes and the BLOBFILE offsets
// (one per entry) that should be used in MIL text.
func WriteMILBlob(entries []BlobEntry) (data []byte, offsets []uint64) {
	// Storage header.
	header := make([]byte, milBlobHeaderSize)
	binary.LittleEndian.PutUint32(header[0:], uint32(len(entries)))
	binary.LittleEndian.PutUint32(header[4:], milBlobVersion)
	data = header

	offsets = make([]uint64, len(entries))
	for i, entry := range entries {
		// Align to 64 bytes.
		data = alignTo(data, milBlobAlignment)
		metadataOffset := uint64(len(data))
		offsets[i] = metadataOffset

		// Raw data starts immediately after 64-byte metadata.
		dataOffset := metadataOffset + milBlobMetadataSize

		// blob_metadata.
		meta := make([]byte, milBlobMetadataSize)
		binary.LittleEndian.PutUint32(meta[0:], milBlobSentinel)
		binary.LittleEndian.PutUint32(meta[4:], uint32(entry.DType))
		binary.LittleEndian.PutUint64(meta[8:], uint64(len(entry.Data)))
		binary.LittleEndian.PutUint64(meta[16:], dataOffset)
		// remaining fields are zero (reserved + padding_size_in_bits)
		data = append(data, meta...)

		// Raw tensor data.
		data = append(data, entry.Data...)
	}

	return data, offsets
}

func alignTo(data []byte, alignment int) []byte {
	rem := len(data) % alignment
	if rem == 0 {
		return data
	}
	pad := alignment - rem
	return append(data, make([]byte, pad)...)
}
