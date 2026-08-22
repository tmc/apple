//go:build darwin

package dispatch

// Data memory-safety conformance tests.
//
// These prove the ownership contracts of the Data wrappers: DataCreate
// copies (mutating the Go slice afterward is invisible), DataCreateNoCopy
// aliases (mutation IS visible) and survives GC via its Pinner, Apply
// reassembles concatenations exactly, and large payloads round-trip
// byte-for-byte.

import (
	"bytes"
	"crypto/sha256"
	"runtime"
	"testing"
	"time"
)

func patternBytes(n int, seed byte) []byte {
	b := make([]byte, n)
	for i := range b {
		b[i] = byte(i)*7 + seed
	}
	return b
}

func TestDataCreateCopies(t *testing.T) {
	buf := patternBytes(4096, 1)
	want := bytes.Clone(buf)
	d := DataCreate(buf)
	defer d.Release()
	for i := range buf {
		buf[i] = 0xFF // mutate after create: must not be visible
	}
	if got := d.Bytes(); !bytes.Equal(got, want) {
		t.Fatal("DataCreate did not copy: mutation of the source slice leaked into the data object")
	}
}

func TestDataCreateNoCopyAliases(t *testing.T) {
	buf := patternBytes(4096, 2)
	d := DataCreateNoCopy(buf)
	defer d.Release()
	buf[100] = 0xAB // mutate after create: MUST be visible (zero-copy proof)
	m := d.Map()
	defer m.Release()
	if m.Len() != len(buf) {
		t.Fatalf("mapped %d bytes, want %d", m.Len(), len(buf))
	}
	if m.Bytes()[100] != 0xAB {
		t.Fatal("DataCreateNoCopy copied the buffer: mutation not visible through the map")
	}
}

func TestDataCreateNoCopySurvivesGC(t *testing.T) {
	// The only reference to the backing array lives inside libdispatch as a
	// laundered pointer; the Pinner must keep it alive and unmoved.
	d := func() Data {
		buf := patternBytes(1<<16, 3)
		return DataCreateNoCopy(buf)
	}()
	defer d.Release()
	for range 5 {
		runtime.GC()
		ballast := make([][]byte, 64)
		for i := range ballast {
			ballast[i] = make([]byte, 1<<14)
		}
		runtime.KeepAlive(ballast)
	}
	got := d.Bytes()
	want := patternBytes(1<<16, 3)
	if !bytes.Equal(got, want) {
		t.Fatal("no-copy data corrupted after GC cycles: backing array moved or collected")
	}
}

func TestDataLargeRoundTrip(t *testing.T) {
	const size = 8 << 20
	buf := patternBytes(size, 4)
	wantSum := sha256.Sum256(buf)
	d := DataCreate(buf)
	defer d.Release()
	if d.Len() != size {
		t.Fatalf("Len() = %d, want %d", d.Len(), size)
	}
	if gotSum := sha256.Sum256(d.Bytes()); gotSum != wantSum {
		t.Fatal("8MB round-trip corrupted data")
	}
}

func TestDataConcatManyChunksApplyReassembles(t *testing.T) {
	// Build a 100-chunk concatenation, then prove Apply visits regions in
	// order, offsets are exact, and the reassembly is byte-identical.
	const chunks, chunkSize = 100, 1024
	var want []byte
	acc := DataEmpty()
	for i := range chunks {
		chunk := patternBytes(chunkSize, byte(i))
		want = append(want, chunk...)
		cd := DataCreate(chunk)
		next := DataCreateConcat(acc, cd)
		cd.Release()
		acc.Release()
		acc = next
	}
	defer acc.Release()

	if acc.Len() != chunks*chunkSize {
		t.Fatalf("concat Len() = %d, want %d", acc.Len(), chunks*chunkSize)
	}

	var rebuilt []byte
	lastOffset := -1
	ok := acc.Apply(func(region Data, offset int, buf []byte) bool {
		if offset != len(rebuilt) {
			t.Errorf("Apply offset %d, want %d (regions out of order)", offset, len(rebuilt))
		}
		if offset <= lastOffset {
			t.Errorf("Apply offset went backward: %d after %d", offset, lastOffset)
		}
		lastOffset = offset
		rebuilt = append(rebuilt, buf...)
		return true
	})
	if !ok {
		t.Fatal("Apply returned false")
	}
	if !bytes.Equal(rebuilt, want) {
		t.Fatal("Apply reassembly differs from original bytes")
	}

	// Bytes() must flatten to the same content.
	if !bytes.Equal(acc.Bytes(), want) {
		t.Fatal("Bytes() of concatenated data differs from original")
	}
}

func TestDataMapZeroCopyStability(t *testing.T) {
	// A mapped view must stay byte-stable across GC while held.
	buf := patternBytes(1<<20, 5)
	d := DataCreate(buf)
	defer d.Release()
	m := d.Map()
	defer m.Release()
	before := sha256.Sum256(m.Bytes())
	runtime.GC()
	runtime.GC()
	after := sha256.Sum256(m.Bytes())
	if before != after {
		t.Fatal("mapped buffer changed under GC")
	}
}

func TestDataReleaseIdempotentPatterns(t *testing.T) {
	// Zero-value and empty-singleton releases must be safe, repeatedly.
	var zero Data
	for range 3 {
		zero.Release()
		DataEmpty().Release()
	}
	var m DataMap
	m.Release()
	m.Release()

	// Mapping and double-releasing a real map through the pointer receiver
	// must be safe (second call sees mapped==0).
	d := DataCreate([]byte("idempotent"))
	defer d.Release()
	mm := d.Map()
	mm.Release()
	mm.Release()
	if mm.Bytes() != nil {
		t.Fatal("released DataMap still exposes bytes")
	}
}

func TestDataConcurrentReaders(t *testing.T) {
	// dispatch_data_t is immutable; concurrent readers across GCD threads
	// must all see identical bytes.
	buf := patternBytes(1<<18, 6)
	want := sha256.Sum256(buf)
	d := DataCreate(buf)
	defer d.Release()

	q := QueueCreateConcurrent("com.appledocs.dispatch.conformance.data-readers")
	g := GroupCreate()
	const readers = 32
	errs := make(chan string, readers)
	for range readers {
		g.Async(q, func() {
			if got := sha256.Sum256(d.Bytes()); got != want {
				select {
				case errs <- "concurrent reader saw corrupted bytes":
				default:
				}
			}
		})
	}
	if !g.Wait(TimeFromNow(int64(30 * time.Second))) {
		t.Fatal("timeout waiting for concurrent data readers")
	}
	select {
	case msg := <-errs:
		t.Fatal(msg)
	default:
	}
}
