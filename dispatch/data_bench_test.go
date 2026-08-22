//go:build darwin

package dispatch

// Cost of getting a Go byte slice into and out of a dispatch data object,
// swept across payload sizes. This is the structural difference between a
// Network.framework send and a plain socket write — nw_connection_send takes
// a dispatch_data_t, so every send pays whatever DataCreate costs — which
// makes it worth measuring on its own rather than inferring it from an
// end-to-end network benchmark.
//
//	go test -run '^$' -bench 'BenchmarkData' -benchmem -benchtime=100x -count=6

import (
	"fmt"
	"testing"
)

// dataBenchSizes spans one page to the region sizes a bulk transfer uses.
var dataBenchSizes = []int{1 << 10, 1 << 16, 1 << 20, 8 << 20}

func sizeName(n int) string {
	switch {
	case n >= 1<<20:
		return fmt.Sprintf("%dMiB", n>>20)
	case n >= 1<<10:
		return fmt.Sprintf("%dKiB", n>>10)
	}
	return fmt.Sprintf("%dB", n)
}

// dataBenchSink defeats dead-code elimination of the read-back benchmarks.
var dataBenchSink byte

// BenchmarkDataCreate measures the copying constructor: libdispatch takes
// its own copy of the buffer, so the Go slice is free immediately after.
func BenchmarkDataCreate(b *testing.B) {
	for _, size := range dataBenchSizes {
		b.Run(sizeName(size), func(b *testing.B) {
			buf := patternBytes(size, 9)
			b.SetBytes(int64(size))
			b.ReportAllocs()
			for b.Loop() {
				d := DataCreate(buf)
				d.Release()
			}
		})
	}
}

// BenchmarkDataCreateNoCopy measures the aliasing constructor, which is the
// same call without the copy. The gap between this and BenchmarkDataCreate
// is what a send pays for copy semantics.
func BenchmarkDataCreateNoCopy(b *testing.B) {
	for _, size := range dataBenchSizes {
		b.Run(sizeName(size), func(b *testing.B) {
			buf := patternBytes(size, 9)
			b.SetBytes(int64(size))
			b.ReportAllocs()
			for b.Loop() {
				d := DataCreateNoCopy(buf)
				d.Release()
			}
		})
	}
}

// BenchmarkDataBytes measures the other direction: a received data object
// flattened into a fresh Go slice, which is what a receive handler does
// when it hands bytes to Go code.
func BenchmarkDataBytes(b *testing.B) {
	for _, size := range dataBenchSizes {
		b.Run(sizeName(size), func(b *testing.B) {
			d := DataCreate(patternBytes(size, 9))
			defer d.Release()
			b.SetBytes(int64(size))
			b.ReportAllocs()
			var sink byte
			for b.Loop() {
				buf := d.Bytes()
				sink ^= buf[len(buf)-1]
			}
			b.StopTimer()
			dataBenchSink ^= sink
		})
	}
}

// BenchmarkDataMap measures reading a data object through a mapped region
// instead of copying it out — the cheap path when the bytes are consumed
// immediately and do not need to outlive the map.
func BenchmarkDataMap(b *testing.B) {
	for _, size := range dataBenchSizes {
		b.Run(sizeName(size), func(b *testing.B) {
			d := DataCreate(patternBytes(size, 9))
			defer d.Release()
			b.SetBytes(int64(size))
			b.ReportAllocs()
			var sink byte
			for b.Loop() {
				m := d.Map()
				buf := m.Bytes()
				sink ^= buf[len(buf)-1]
				m.Release()
			}
			b.StopTimer()
			dataBenchSink ^= sink
		})
	}
}
