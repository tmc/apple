package jaccl

import (
	"encoding/binary"
	"errors"
	"math"
	"sync"
	"testing"
	"time"
)

func TestRingGatherBlocks(t *testing.T) {
	tests := []struct {
		size  int
		bytes int
		wires int
	}{
		{size: 2, bytes: 1, wires: 4},
		{size: 3, bytes: 10, wires: 3},
		{size: 4, bytes: 9, wires: 2},
	}
	for _, test := range tests {
		t.Run("ring", func(t *testing.T) {
			network := newMemoryRing(test.size)
			sources := make([][]byte, test.size)
			destinations := make([][]byte, test.size)
			for rank := range sources {
				sources[rank] = make([]byte, test.bytes)
				for index := range sources[rank] {
					sources[rank][index] = byte(rank*31 + index)
				}
				destinations[rank] = make([]byte, test.size*test.bytes)
			}
			errs := make(chan error, test.size)
			var group sync.WaitGroup
			for rank := range sources {
				group.Add(1)
				go func() {
					defer group.Done()
					errs <- ringGatherBlocks(rank, test.size, destinations[rank], sources[rank], test.wires, network.stage(rank))
				}()
			}
			done := make(chan struct{})
			go func() {
				group.Wait()
				close(done)
			}()
			select {
			case <-done:
			case <-time.After(time.Second):
				t.Fatal("ring gather did not complete")
			}
			for range sources {
				if err := <-errs; err != nil {
					t.Fatal(err)
				}
			}
			for rank, got := range destinations {
				for source, want := range sources {
					segment := got[source*test.bytes : (source+1)*test.bytes]
					if string(segment) != string(want) {
						t.Fatalf("rank %d source %d = %x, want %x", rank, source, segment, want)
					}
				}
			}
		})
	}
}

func TestReduceRingUsesStandaloneChunkOrder(t *testing.T) {
	tests := []struct {
		name  string
		count int
		wires int
		want  []float32
	}{
		{"one direction", 3, 1, []float32{1, 0, 0}},
		{"two directions", 6, 1, []float32{1, 0, 0, 1, 0, 0}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			blockSize := test.count * 4
			gathered := make([]byte, 3*blockSize)
			for rank, value := range []float32{1e20, -1e20, 1} {
				for element := 0; element < test.count; element++ {
					binary.LittleEndian.PutUint32(gathered[rank*blockSize+element*4:], math.Float32bits(value))
				}
			}
			dst := make([]byte, blockSize)
			if err := reduceRing(dst, gathered, 3, blockSize, Float32, Sum, test.wires); err != nil {
				t.Fatal(err)
			}
			for element, want := range test.want {
				got := math.Float32frombits(binary.LittleEndian.Uint32(dst[element*4:]))
				if got != want {
					t.Fatalf("element %d = %g, want %g", element, got, want)
				}
			}
		})
	}
}

func TestRingReductionPlan(t *testing.T) {
	tests := []struct {
		name                 string
		count, size, wires   int
		bytes                int
		wantChunk, wantSplit int
		wantReverse          bool
	}{
		{"small one direction", 5, 3, 2, 20, 2, 2, true},
		{"small two directions", 6, 3, 1, 24, 2, 1, false},
		{"exactly 64 KiB uses one wire", 16_384, 3, 2, 65_536, 5_462, 2_731, false},
		{"above 64 KiB uses all wires", 16_385, 3, 2, 65_540, 5_462, 2_732, false},
		{"large all wires", 20_000, 3, 2, 80_000, 6_667, 3_334, false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			chunk, split, reverse := ringReductionPlan(test.count, test.size, test.wires, test.bytes)
			if chunk != test.wantChunk || split != test.wantSplit || reverse != test.wantReverse {
				t.Fatalf("plan = chunk=%d split=%d reverse=%t, want chunk=%d split=%d reverse=%t", chunk, split, reverse, test.wantChunk, test.wantSplit, test.wantReverse)
			}
		})
	}
}

func TestMeshUsesRingReduction(t *testing.T) {
	tests := []struct {
		name      string
		size      int
		blockSize int
		dtype     DType
		want      bool
	}{
		{"two ranks stay mesh", 2, 8 * 1024 * 1024, Float32, false},
		{"float32 below threshold", 3, (2*1024*1024 - 1) * 4, Float32, false},
		{"float32 at threshold", 3, 2 * 1024 * 1024 * 4, Float32, true},
		{"float64 at threshold", 3, 1 * 1024 * 1024 * 8, Float64, true},
		{"bfloat16 at threshold", 3, 256 * 1024 * 2, BFloat16, false},
		{"bfloat16 above threshold", 3, (256*1024 + 1) * 2, BFloat16, true},
		{"invalid shape", 3, 3, Float32, false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := meshUsesRingReduction(test.size, test.blockSize, test.dtype); got != test.want {
				t.Fatalf("meshUsesRingReduction(%d, %d, %d) = %t, want %t", test.size, test.blockSize, test.dtype, got, test.want)
			}
		})
	}
}

func TestRingGatherBlocksRejectsInvalidShape(t *testing.T) {
	if err := ringGatherBlocks(0, 2, make([]byte, 3), []byte{1, 2}, 1, func(int, int, []byte, []byte) error { return nil }); !errors.Is(err, ErrProtocol) {
		t.Fatalf("ringGatherBlocks error = %v, want ErrProtocol", err)
	}
}

type memoryRing struct {
	size int

	mu         sync.Mutex
	cond       *sync.Cond
	generation int
	arrived    int
	sends      [][]byte
	receives   [][]byte
}

func newMemoryRing(size int) *memoryRing {
	ring := &memoryRing{size: size, sends: make([][]byte, size), receives: make([][]byte, size)}
	ring.cond = sync.NewCond(&ring.mu)
	return ring
}

func (r *memoryRing) stage(rank int) func(int, int, []byte, []byte) error {
	return func(step, wire int, send, receive []byte) error {
		r.mu.Lock()
		generation := r.generation
		r.sends[rank] = append(r.sends[rank][:0], send...)
		r.receives[rank] = receive
		r.arrived++
		if r.arrived == r.size {
			for destination := 0; destination < r.size; destination++ {
				source := (destination + r.size - 1) % r.size
				copy(r.receives[destination], r.sends[source])
			}
			r.arrived = 0
			r.generation++
			r.cond.Broadcast()
			r.mu.Unlock()
			return nil
		}
		for generation == r.generation {
			r.cond.Wait()
		}
		r.mu.Unlock()
		return nil
	}
}
