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
	milBlobAlignment    = 64
	milBlobVersion      = 2
	milBlobSentinel     = 0xDEADBEEF
	milBlobHeaderSize   = 64
	milBlobMetadataSize = 64
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
	case DataTypeUInt1:
		return BlobDataTypeUInt1, nil
	case DataTypeUInt2:
		return BlobDataTypeUInt2, nil
	case DataTypeUInt3:
		return BlobDataTypeUInt3, nil
	case DataTypeUInt4:
		return BlobDataTypeUInt4, nil
	case DataTypeUInt6:
		return BlobDataTypeUInt6, nil
	case DataTypeFloat8E4M3FN:
		return BlobDataTypeFloat8E4M3, nil
	case DataTypeFloat8E5M2:
		return BlobDataTypeFloat8E5M2, nil
	default:
		return 0, fmt.Errorf("unsupported data type for MILBlob: %v", dt)
	}
}

// specVersionIOS18 is the model specification version introduced with iOS 18,
// from which integer weights wider than 8 bits may live in the weight file.
const specVersionIOS18 = 9

// ShouldUseWeightFile reports whether a constant of the given element type and
// element count belongs in the weight file rather than inline in MIL text.
// It mirrors coremltools' should_use_weight_file: at least 10 elements and an
// element type the target specification version stores in a blob.
func ShouldUseWeightFile(dt DataType, numElements int, specVersion int32) bool {
	if numElements < 10 {
		return false
	}
	switch dt {
	case DataTypeFloat16, DataTypeFloat32, DataTypeUInt8, DataTypeInt8:
		return true
	case DataTypeUInt16, DataTypeInt16, DataTypeInt32, DataTypeUInt32:
		return specVersion >= specVersionIOS18
	}
	return false
}

// subByteBits reports the bit width of a sub-byte blob element type, or 0 for
// byte-sized types. The runtime derives a sub-byte blob's element count from
// sizeInBytes and padding_size_in_bits, so these widths are part of the format.
func subByteBits(dt BlobDataType) int {
	switch dt {
	case BlobDataTypeUInt1:
		return 1
	case BlobDataTypeUInt2:
		return 2
	case BlobDataTypeUInt3:
		return 3
	case BlobDataTypeInt4, BlobDataTypeUInt4:
		return 4
	case BlobDataTypeUInt6:
		return 6
	}
	return 0
}

// subByteRange reports the inclusive value range of a sub-byte element type.
func subByteRange(dt BlobDataType) (min, max int8) {
	switch dt {
	case BlobDataTypeInt4:
		return -8, 7
	case BlobDataTypeUInt1:
		return 0, 1
	case BlobDataTypeUInt2:
		return 0, 3
	case BlobDataTypeUInt3:
		return 0, 7
	case BlobDataTypeUInt4:
		return 0, 15
	case BlobDataTypeUInt6:
		return 0, 63
	}
	return 0, 0
}

// PackSubByte packs values of a sub-byte blob element type into the byte
// payload MILBlob stores. Element i occupies bits [i*bits, (i+1)*bits) of the
// stream, little-endian within each byte; for widths that do not divide 8 an
// element straddling a byte boundary continues in the low bits of the next
// byte. This mirrors MILBlob's PackSubByteVec.
func PackSubByte(dt BlobDataType, values []int8) ([]byte, error) {
	bits := subByteBits(dt)
	if bits == 0 {
		return nil, fmt.Errorf("not a sub-byte MILBlob data type: %v", dt)
	}
	min, max := subByteRange(dt)
	out := make([]byte, (len(values)*bits+7)/8)
	for i, v := range values {
		if v < min || v > max {
			return nil, fmt.Errorf("value %d is outside allowed subbyte datatype range [%d, %d]", v, min, max)
		}
		startBit := i * bits
		idx, off := startBit/8, startBit%8
		masked := uint8(v) & uint8((1<<bits)-1)
		out[idx] |= masked << off
		if off > 8-bits {
			// The element spills over into the next byte.
			out[idx+1] |= masked >> (8 - off)
		}
	}
	return out, nil
}

// BlobEntry describes a single tensor to write into a MILBlob weight file.
type BlobEntry struct {
	DType BlobDataType
	Data  []byte

	// NumElements is the logical element count. It is required for sub-byte
	// element types, where the trailing partial byte must be reported in the
	// metadata's padding_size_in_bits; it is ignored for byte-sized types.
	NumElements int
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
		// Sub-byte payloads end in a partial byte whose unused high bits the
		// reader must be told about; without it it derives a wrong element
		// count or rejects the blob outright.
		if bits := subByteBits(entry.DType); bits != 0 {
			if rem := (entry.NumElements * bits) % 8; rem != 0 {
				binary.LittleEndian.PutUint64(meta[24:], uint64(8-rem))
			}
		}
		// remaining fields are zero (reserved)
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
