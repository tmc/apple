//go:build darwin

package mil

import (
	"bytes"
	"encoding/binary"
	"math"
	"testing"

	"github.com/tmc/apple/x/coremlcompiler"
)

// metaAt decodes the metadata block at off.
func metaAt(t *testing.T, data []byte, off uint64) (sentinel, dtype uint32, size, dataOff uint64) {
	t.Helper()
	if off%blobAlignment != 0 {
		t.Errorf("metadata offset %d is not %d-byte aligned", off, blobAlignment)
	}
	if int(off)+blobMetaSize > len(data) {
		t.Fatalf("metadata offset %d past end of %d-byte blob", off, len(data))
	}
	m := data[off:]
	return binary.LittleEndian.Uint32(m[0:]),
		binary.LittleEndian.Uint32(m[4:]),
		binary.LittleEndian.Uint64(m[8:]),
		binary.LittleEndian.Uint64(m[16:])
}

func TestBlobWriterSingle(t *testing.T) {
	w := NewBlobWriter()
	idx := w.AddFloat16([]float32{1.0, 2.0, 3.0, 4.0})
	offset := w.Offset(idx)

	data, err := w.Build()
	if err != nil {
		t.Fatal(err)
	}

	if count := binary.LittleEndian.Uint32(data[0:]); count != 1 {
		t.Errorf("blob count = %d, want 1", count)
	}
	if version := binary.LittleEndian.Uint32(data[4:]); version != blobVersion {
		t.Errorf("version = %d, want %d", version, blobVersion)
	}

	// A single blob's metadata directly follows the 64-byte header, which is
	// the offset every BLOBFILE reference in this package's MIL text uses.
	if offset != 64 {
		t.Errorf("metadata offset = %d, want 64", offset)
	}

	sentinel, dtype, size, dataOff := metaAt(t, data, offset)
	if sentinel != blobSentinel {
		t.Errorf("sentinel = 0x%X, want 0x%X", sentinel, blobSentinel)
	}
	if dtype != uint32(BlobFloat16) {
		t.Errorf("dtype = %d, want %d", dtype, BlobFloat16)
	}
	if size != 8 { // 4 fp16 values * 2 bytes
		t.Errorf("data size = %d, want 8", size)
	}
	if dataOff != offset+blobMetaSize {
		t.Errorf("data offset = %d, want %d", dataOff, offset+blobMetaSize)
	}
}

func TestBlobWriterMultiple(t *testing.T) {
	w := NewBlobWriter()
	idx1 := w.AddFloat16([]float32{1.0, 2.0})
	idx2 := w.AddFloat16([]float32{3.0, 4.0, 5.0})

	// Offsets are only final once every blob has been added.
	off1 := w.Offset(idx1)
	off2 := w.Offset(idx2)

	data, err := w.Build()
	if err != nil {
		t.Fatal(err)
	}

	if count := binary.LittleEndian.Uint32(data[0:]); count != 2 {
		t.Errorf("blob count = %d, want 2", count)
	}
	if off2 <= off1 {
		t.Errorf("offset2 %d <= offset1 %d", off2, off1)
	}

	for i, tc := range []struct {
		off  uint64
		size uint64
	}{{off1, 4}, {off2, 6}} {
		_, _, size, dataOff := metaAt(t, data, tc.off)
		if size != tc.size {
			t.Errorf("blob %d: size = %d, want %d", i, size, tc.size)
		}
		if dataOff != tc.off+blobMetaSize {
			t.Errorf("blob %d: data offset = %d, want %d", i, dataOff, tc.off+blobMetaSize)
		}
	}
}

func TestBlobWriterEmpty(t *testing.T) {
	w := NewBlobWriter()
	if _, err := w.Build(); err == nil {
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

	_, _, _, dataOff := metaAt(t, data, offset)
	if got := data[dataOff : int(dataOff)+len(raw)]; !bytes.Equal(got, raw) {
		t.Errorf("data = % x, want % x", got, raw)
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

	_, dtype, _, dataOff := metaAt(t, data, offset)
	if dtype != uint32(BlobFloat32) {
		t.Errorf("dtype = %d, want %d", dtype, BlobFloat32)
	}
	for i, v := range vals {
		got := binary.LittleEndian.Uint32(data[int(dataOff)+i*4:])
		want := math.Float32bits(v)
		if got != want {
			t.Errorf("float32[%d]: got bits 0x%08X, want 0x%08X (value %v)", i, got, want, v)
		}
	}
}

func TestBlobWriterAlignment(t *testing.T) {
	w := NewBlobWriter()
	var idx []int
	for i := range 5 {
		idx = append(idx, w.AddRaw(BlobFloat16, make([]byte, 10+i*7))) // non-aligned sizes
	}

	data, err := w.Build()
	if err != nil {
		t.Fatal(err)
	}

	for i, id := range idx {
		off := w.Offset(id)
		if off%blobAlignment != 0 {
			t.Errorf("blob %d metadata offset %d not %d-byte aligned", i, off, blobAlignment)
		}
		sentinel, _, size, dataOff := metaAt(t, data, off)
		if sentinel != blobSentinel {
			t.Errorf("blob %d: sentinel = 0x%X at offset %d", i, sentinel, off)
		}
		if want := uint64(10 + i*7); size != want {
			t.Errorf("blob %d: size = %d, want %d", i, size, want)
		}
		if dataOff != off+blobMetaSize {
			t.Errorf("blob %d: data offset = %d, want %d", i, dataOff, off+blobMetaSize)
		}
	}
}

// blobSizes are entry payload sizes chosen so that most are not multiples of
// the 64-byte alignment, which is the only case in which a packed layout and
// an aligned one disagree.
var blobSizes = []int{1, 5, 31, 32, 33, 64, 65, 1000}

// TestBlobWriterMultiBlobOffsets pins the multi-blob offsets against a table
// computed by hand from the format rules: entry i's metadata block starts at
// the previous entry's data end rounded up to 64, and its data follows the
// 64-byte metadata block with no padding between the two.
//
// The rounding-up is the load-bearing part. A packed layout, which puts the
// next metadata block directly at the previous data end, agrees with this one
// whenever every payload is a multiple of 64 and disagrees otherwise. Apple's
// own compiler rounds up: of 46 Blob Storage v2 weight files produced by
// coremlcompiler on this machine, all 46 parse under the rounding-up rule and
// only 44 parse under the packed one, and every one of the 4 gaps that tells
// the two apart is rounded up.
func TestBlobWriterMultiBlobOffsets(t *testing.T) {
	want := []uint64{64, 192, 320, 448, 576, 704, 832, 1024}
	const wantSize = 2088 // last entry's data ends at EOF; no tail padding

	w := NewBlobWriter()
	for i, n := range blobSizes {
		if got := w.AddRaw(BlobFloat16, make([]byte, n)); got != i {
			t.Fatalf("AddRaw returned index %d, want %d", got, i)
		}
	}
	data, err := w.Build()
	if err != nil {
		t.Fatal(err)
	}

	if len(data) != wantSize {
		t.Errorf("blob size = %d, want %d", len(data), wantSize)
	}
	for i := range blobSizes {
		if got := w.Offset(i); got != want[i] {
			t.Errorf("Offset(%d) = %d, want %d", i, got, want[i])
		}
	}

	// Walk the file the way a reader does, from the header alone, and check
	// that the walk lands on each offset Offset reported.
	count := binary.LittleEndian.Uint32(data[0:])
	if int(count) != len(blobSizes) {
		t.Fatalf("count = %d, want %d", count, len(blobSizes))
	}
	off := uint64(blobHeaderSize)
	for i := range count {
		if off != w.Offset(int(i)) {
			t.Fatalf("walk reached %d at entry %d, Offset says %d", off, i, w.Offset(int(i)))
		}
		sentinel, _, size, dataOff := metaAt(t, data, off)
		if sentinel != blobSentinel {
			t.Fatalf("entry %d: sentinel = 0x%X at offset %d", i, sentinel, off)
		}
		if size != uint64(blobSizes[i]) {
			t.Errorf("entry %d: size = %d, want %d", i, size, blobSizes[i])
		}
		off = uint64(alignUp(int(dataOff+size), blobAlignment))
	}
}

// TestBlobWriterRoundTrip checks that each entry's payload is recoverable at
// the offset its metadata records, for entries of differing sizes.
func TestBlobWriterRoundTrip(t *testing.T) {
	payload := func(n, seed int) []byte {
		b := make([]byte, n)
		for i := range b {
			b[i] = byte(i*31 + seed*7 + 1)
		}
		return b
	}

	w := NewBlobWriter()
	var want [][]byte
	for i, n := range blobSizes {
		want = append(want, payload(n, i))
		w.AddRaw(BlobFloat16, want[i])
	}
	data, err := w.Build()
	if err != nil {
		t.Fatal(err)
	}

	for i := range want {
		_, _, size, dataOff := metaAt(t, data, w.Offset(i))
		if int(dataOff)+int(size) > len(data) {
			t.Fatalf("entry %d: data [%d,%d) past end of %d-byte blob", i, dataOff, int(dataOff)+int(size), len(data))
		}
		if got := data[dataOff : dataOff+size]; !bytes.Equal(got, want[i]) {
			t.Errorf("entry %d (%d bytes): payload at offset %d does not round-trip", i, len(want[i]), dataOff)
		}
	}
}

// TestBlobWriterMatchesCoreMLCompiler checks this package's encoder against
// x/coremlcompiler's independent implementation of the same format.
func TestBlobWriterMatchesCoreMLCompiler(t *testing.T) {
	sizes := blobSizes

	raw := func(n, seed int) []byte {
		b := make([]byte, n)
		for i := range b {
			b[i] = byte(i*31 + seed)
		}
		return b
	}

	dtypes := []struct {
		mil    BlobDataType
		coreml coremlcompiler.BlobDataType
	}{
		{BlobFloat16, coremlcompiler.BlobDataTypeFloat16},
		{BlobFloat32, coremlcompiler.BlobDataTypeFloat32},
		{BlobUInt8, coremlcompiler.BlobDataTypeUInt8},
		{BlobInt8, coremlcompiler.BlobDataTypeInt8},
	}

	// Every defined data type code must agree.
	for _, d := range dtypes {
		if uint32(d.mil) != uint32(d.coreml) {
			t.Errorf("data type code %v = %d, coremlcompiler = %d", d.mil, d.mil, d.coreml)
		}
	}

	// Single-entry files, one per data type.
	for _, d := range dtypes {
		w := NewBlobWriter()
		idx := w.AddRaw(d.mil, raw(64, 3))
		got, err := w.Build()
		if err != nil {
			t.Fatal(err)
		}
		want, offsets := coremlcompiler.WriteMILBlob([]coremlcompiler.BlobEntry{{DType: d.coreml, Data: raw(64, 3)}})
		if !bytes.Equal(got, want) {
			t.Errorf("dtype %v: %d bytes differ from coremlcompiler", d.mil, len(got))
		}
		if w.Offset(idx) != offsets[0] {
			t.Errorf("dtype %v: offset = %d, coremlcompiler = %d", d.mil, w.Offset(idx), offsets[0])
		}
	}

	// A multi-entry file with sizes that do not divide the 64-byte alignment.
	w := NewBlobWriter()
	var entries []coremlcompiler.BlobEntry
	var idx []int
	for i, n := range sizes {
		idx = append(idx, w.AddRaw(BlobFloat16, raw(n, i)))
		entries = append(entries, coremlcompiler.BlobEntry{DType: coremlcompiler.BlobDataTypeFloat16, Data: raw(n, i)})
	}
	got, err := w.Build()
	if err != nil {
		t.Fatal(err)
	}
	want, offsets := coremlcompiler.WriteMILBlob(entries)
	if !bytes.Equal(got, want) {
		t.Errorf("multi-entry blob: got %d bytes, coremlcompiler %d bytes; contents differ", len(got), len(want))
	}
	for i, id := range idx {
		if w.Offset(id) != offsets[i] {
			t.Errorf("entry %d: offset = %d, coremlcompiler = %d", i, w.Offset(id), offsets[i])
		}
	}
}

// TestSingleBlobWritersAgree checks that the exported single-weight helpers
// all emit the same one-entry blob file.
func TestSingleBlobWritersAgree(t *testing.T) {
	for _, n := range []int{1, 5, 33, 64, 1000} {
		vals := make([]float32, n)
		for i := range vals {
			vals[i] = float32(i%97-48) * 0.03125
		}

		conv, err := BuildWeightBlob(vals, 1, n)
		if err != nil {
			t.Fatal(err)
		}
		flat, err := BuildWeightBlobV1(vals)
		if err != nil {
			t.Fatal(err)
		}
		generic, err := BuildFP16Blob(vals)
		if err != nil {
			t.Fatal(err)
		}

		w := NewBlobWriter()
		w.AddFloat16(vals)
		writer, err := w.Build()
		if err != nil {
			t.Fatal(err)
		}

		for name, got := range map[string][]byte{
			"BuildWeightBlobV1": flat,
			"BuildFP16Blob":     generic,
			"BlobWriter":        writer,
		} {
			if !bytes.Equal(got, conv) {
				t.Errorf("n=%d: %s differs from BuildWeightBlob", n, name)
			}
		}
	}
}
