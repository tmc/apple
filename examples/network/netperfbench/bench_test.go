//go:build darwin

package main

// go test -bench form of the same measurement, for when the answer wanted is
// "did this change make round trips slower" rather than "how do the three
// implementations compare". Both roles run in one process here; the matrix
// in run.sh keeps them in separate processes.
//
// The comparison worth running is between two trees, not between two lines
// of one run:
//
//	go test -run '^$' -bench BenchmarkRoundTrip -benchmem -benchtime=2000x -count=6 | tee new.txt
//	benchstat old.txt new.txt
//
// A fixed -benchtime=Nx keeps every arm at the same iteration count, so the
// b.N autoscaling is not itself a source of variance between arms.

import (
	"fmt"
	"testing"
)

var benchSizes = []int{64, 4096, 65536}

// benchInflight sweeps how many messages are outstanding at once: 1 is pure
// latency, 32 overlaps the per-operation costs.
var benchInflight = []int{1, 32}

func benchName(size int) string {
	if size >= 1<<10 {
		return fmt.Sprintf("%dKiB", size>>10)
	}
	return fmt.Sprintf("%dB", size)
}

func BenchmarkRoundTrip(b *testing.B) {
	for _, transport := range []string{"std", "nw"} {
		b.Run(transport, func(b *testing.B) {
			for _, inflight := range benchInflight {
				b.Run(fmt.Sprintf("inflight%d", inflight), func(b *testing.B) {
					for _, size := range benchSizes {
						b.Run(benchName(size), func(b *testing.B) {
							benchRoundTrip(b, transport, size, inflight)
						})
					}
				})
			}
		})
	}
}

// BenchmarkRoundTripBatch measures the -recv-batch receive mode, where the
// client asks for a whole outstanding batch in one receive instead of one per
// echo. It is the largest single effect measured on this path — the framework
// stops re-arming a receive per message, and the bindings stop crossing back
// into Go per message — so it is worth a benchmark rather than only a flag.
//
// It is a different workload, not a faster spelling of the same one: a single
// receive for a whole batch gives up per-message framing. Compare it against
// BenchmarkRoundTrip to size the effect, never as a drop-in replacement.
//
// Only nw runs here. The std transport has no receive re-arming to amortise,
// and depth 1 is skipped because a batch of one is the default path.
func BenchmarkRoundTripBatch(b *testing.B) {
	defer func(previous bool) { *recvBatch = previous }(*recvBatch)
	*recvBatch = true

	for _, inflight := range benchInflight {
		if inflight == 1 {
			continue
		}
		b.Run(fmt.Sprintf("inflight%d", inflight), func(b *testing.B) {
			for _, size := range benchSizes {
				b.Run(benchName(size), func(b *testing.B) {
					benchRoundTrip(b, "nw", size, inflight)
				})
			}
		})
	}
}

func benchRoundTrip(b *testing.B, transport string, size, inflight int) {
	b.Helper()

	srv, err := serve(transport, "0")
	if err != nil {
		b.Fatal(err)
	}
	defer srv.Close()

	c, err := dial(transport, "127.0.0.1:"+srv.Port(), inflight)
	if err != nil {
		b.Fatal(err)
	}
	defer c.Close()

	buf := make([]byte, size)
	for i := range buf {
		buf[i] = byte(i*7 + 13)
	}
	// Warm the connection so the first measured round trip is not paying
	// for slow-start and first-touch faults.
	for range 100 {
		if err := c.RoundTrip(buf, inflight); err != nil {
			b.Fatal(err)
		}
	}

	// One iteration moves inflight messages in both directions.
	b.SetBytes(int64(size) * 2 * int64(inflight))
	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		if err := c.RoundTrip(buf, inflight); err != nil {
			b.Fatal(err)
		}
	}
}
