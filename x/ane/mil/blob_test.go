//go:build darwin

package mil

import (
	"encoding/binary"
	"math"
	"testing"
)

func TestBlobWriterSingle(t *testing.T) {
	w := NewBlobWriter()
	idx := w.AddFloat16([]float32{1.0, 2.0, 3.0, 4.0})

	offset := w.Offset(idx)

	data, err := w.Build()
	if err != nil {
		t.Fatal(err)
	}

	// Check file header.
	count := binary.LittleEndian.Uint64(data[0:])
	if count != 1 {
		t.Errorf("blob count = %d, want 1", count)
	}
	version := binary.LittleEndian.Uint32(data[8:])
	if version != 2 {
		t.Errorf("version = %d, want 2", version)
	}

	// Check chunk metadata.
	magic := binary.LittleEndian.Uint32(data[64:])
	if magic != 0xDEADBEEF {
		t.Errorf("magic = 0x%X, want 0xDEADBEEF", magic)
	}
	dtype := binary.LittleEndian.Uint32(data[68:])
	if dtype != uint32(BlobFloat16) {
		t.Errorf("dtype = %d, want %d", dtype, BlobFloat16)
	}
	size := binary.LittleEndian.Uint64(data[72:])
	if size != 8 { // 4 fp16 values * 2 bytes
		t.Errorf("data size = %d, want 8", size)
	}
	storedOffset := binary.LittleEndian.Uint64(data[80:])
	if storedOffset != offset {
		t.Errorf("stored offset = %d, want %d", storedOffset, offset)
	}

	// Verify data offset is 64-byte aligned.
	if offset%64 != 0 {
		t.Errorf("offset %d is not 64-byte aligned", offset)
	}
}

func TestBlobWriterMultiple(t *testing.T) {
	w := NewBlobWriter()
	idx1 := w.AddFloat16([]float32{1.0, 2.0})
	idx2 := w.AddFloat16([]float32{3.0, 4.0, 5.0})

	// Offsets computed after all blobs added.
	off1 := w.Offset(idx1)
	off2 := w.Offset(idx2)

	data, err := w.Build()
	if err != nil {
		t.Fatal(err)
	}

	count := binary.LittleEndian.Uint64(data[0:])
	if count != 2 {
		t.Errorf("blob count = %d, want 2", count)
	}

	// Both offsets must be 64-byte aligned.
	if off1%64 != 0 {
		t.Errorf("offset1 %d not 64-byte aligned", off1)
	}
	if off2%64 != 0 {
		t.Errorf("offset2 %d not 64-byte aligned", off2)
	}

	// Second blob must come after first.
	if off2 <= off1 {
		t.Errorf("offset2 %d <= offset1 %d", off2, off1)
	}

	// Verify stored offsets match computed offsets.
	stored1 := binary.LittleEndian.Uint64(data[64+16:])  // first chunk meta
	stored2 := binary.LittleEndian.Uint64(data[128+16:]) // second chunk meta
	if stored1 != off1 {
		t.Errorf("stored offset1 = %d, want %d", stored1, off1)
	}
	if stored2 != off2 {
		t.Errorf("stored offset2 = %d, want %d", stored2, off2)
	}
}

func TestBlobWriterEmpty(t *testing.T) {
	w := NewBlobWriter()
	_, err := w.Build()
	if err == nil {
		t.Error("expected error for empty blob writer")
	}
}

func TestBlobWriterRaw(t *testing.T) {
	w := NewBlobWriter()
	raw := []byte{0x01, 0x02, 0x03, 0x04}
	idx := w.AddRaw(BlobUInt8, raw)
	offset := w.Offset(idx)

	data, err := w.Build()
	if err != nil {
		t.Fatal(err)
	}

	// Verify data at offset.
	for i, v := range raw {
		if data[int(offset)+i] != v {
			t.Errorf("data[%d] = 0x%02X, want 0x%02X", int(offset)+i, data[int(offset)+i], v)
		}
	}
}

func TestBlobWriterFloat32(t *testing.T) {
	w := NewBlobWriter()
	vals := []float32{1.0, -2.5, 3.14159, 0.0, math.SmallestNonzeroFloat32, math.MaxFloat32}
	idx := w.AddFloat32(vals)
	offset := w.Offset(idx)

	data, err := w.Build()
	if err != nil {
		t.Fatal(err)
	}

	// Verify IEEE 754 bit patterns are preserved.
	for i, v := range vals {
		got := binary.LittleEndian.Uint32(data[int(offset)+i*4:])
		want := math.Float32bits(v)
		if got != want {
			t.Errorf("float32[%d]: got bits 0x%08X, want 0x%08X (value %v)", i, got, want, v)
		}
	}

	// Verify dtype in metadata.
	dtype := binary.LittleEndian.Uint32(data[68:])
	if dtype != uint32(BlobFloat32) {
		t.Errorf("dtype = %d, want %d", dtype, BlobFloat32)
	}
}

func TestBlobWriterAlignment(t *testing.T) {
	w := NewBlobWriter()
	// Add blobs of varying sizes to test alignment.
	for i := range 5 {
		raw := make([]byte, 10+i*7) // non-aligned sizes
		w.AddRaw(BlobFloat16, raw)
	}

	data, err := w.Build()
	if err != nil {
		t.Fatal(err)
	}

	// Verify all stored offsets are 64-byte aligned.
	for i := range 5 {
		metaOff := 64 + 64*i
		storedOffset := binary.LittleEndian.Uint64(data[metaOff+16:])
		if storedOffset%64 != 0 {
			t.Errorf("blob %d offset %d not 64-byte aligned", i, storedOffset)
		}
	}
}
