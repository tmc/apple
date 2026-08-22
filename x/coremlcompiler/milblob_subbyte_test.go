package coremlcompiler

import (
	"encoding/binary"
	"testing"
)

// TestPackSubByte checks the packing order against MILBlob's PackSubByteVec:
// little-endian within a byte, straddling elements continuing in the low bits
// of the next byte.
func TestPackSubByte(t *testing.T) {
	tests := []struct {
		name   string
		dt     BlobDataType
		values []int8
		want   []byte
	}{
		{"uint4 pair", BlobDataTypeUInt4, []int8{1, 2}, []byte{0x21}},
		{"uint4 odd", BlobDataTypeUInt4, []int8{15, 1, 3}, []byte{0x1F, 0x03}},
		{"int4 negative", BlobDataTypeInt4, []int8{-8, -1}, []byte{0xF8}},
		{"uint1", BlobDataTypeUInt1, []int8{1, 0, 1, 1}, []byte{0x0D}},
		{"uint2", BlobDataTypeUInt2, []int8{3, 2, 1, 0}, []byte{0x1B}},
		// 3-bit: 5,7,7 -> byte0 = 101|111<<3|111<<6 = 0xFD, third element's high
		// bit continues in bit 0 of byte1.
		{"uint3 straddle", BlobDataTypeUInt3, []int8{5, 7, 7}, []byte{0xFD, 0x01}},
		{"uint6 straddle", BlobDataTypeUInt6, []int8{63, 63}, []byte{0xFF, 0x0F}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PackSubByte(tt.dt, tt.values)
			if err != nil {
				t.Fatalf("PackSubByte: %v", err)
			}
			if len(got) != len(tt.want) {
				t.Fatalf("got %d bytes %x, want %d bytes %x", len(got), got, len(tt.want), tt.want)
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Fatalf("got %x, want %x", got, tt.want)
				}
			}
		})
	}
}

func TestPackSubByteRange(t *testing.T) {
	if _, err := PackSubByte(BlobDataTypeUInt4, []int8{16}); err == nil {
		t.Fatal("want out-of-range error for 16 as uint4")
	}
	if _, err := PackSubByte(BlobDataTypeFloat16, []int8{1}); err == nil {
		t.Fatal("want error for non-sub-byte type")
	}
}

// TestWriteMILBlobPaddingBits checks blob_metadata.padding_size_in_bits, which
// the runtime uses to recover the element count of a sub-byte blob.
func TestWriteMILBlobPaddingBits(t *testing.T) {
	tests := []struct {
		name        string
		dt          BlobDataType
		numElements int
		want        uint64
	}{
		{"uint3 five elements", BlobDataTypeUInt3, 5, 1},
		{"uint3 eight elements", BlobDataTypeUInt3, 8, 0},
		{"uint4 odd", BlobDataTypeUInt4, 3, 4},
		{"uint4 even", BlobDataTypeUInt4, 4, 0},
		{"uint6 two", BlobDataTypeUInt6, 2, 4},
		{"float16 none", BlobDataTypeFloat16, 4, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bits := subByteBits(tt.dt)
			n := tt.numElements * 16
			if bits != 0 {
				n = (tt.numElements*bits + 7) / 8
			}
			blob, offsets := WriteMILBlob([]BlobEntry{{
				DType:       tt.dt,
				Data:        make([]byte, n),
				NumElements: tt.numElements,
			}})
			meta := blob[offsets[0]:]
			if got := binary.LittleEndian.Uint64(meta[24:]); got != tt.want {
				t.Errorf("padding_size_in_bits = %d, want %d", got, tt.want)
			}
		})
	}
}
