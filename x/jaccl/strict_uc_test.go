package jaccl

import (
	"bytes"
	"encoding/binary"
	"errors"
	"math"
	"math/rand"
	"strconv"
	"sync"
	"testing"
)

// strictUCWire drops sends for which the receiver has no posted work request.
// It is deliberately stricter than a buffered Go channel: UC has no RNR retry.
type strictUCWire struct {
	recv    *receiveSlots
	buffers [][]byte
	dropped int
}

func newStrictUCWire(depth int) (*strictUCWire, error) {
	recv, err := newReceiveSlots(1, 1, depth)
	if err != nil {
		return nil, err
	}
	if _, err := recv.postInitial(); err != nil {
		return nil, err
	}
	return &strictUCWire{recv: recv, buffers: make([][]byte, depth)}, nil
}

func (w *strictUCWire) send(src []byte) (uint64, error) {
	id, err := w.recv.takeCredit()
	if err != nil {
		w.dropped++
		return 0, err
	}
	slot, err := w.recv.complete(id)
	if err != nil {
		return 0, err
	}
	w.buffers[slot] = append(w.buffers[slot][:0], src...)
	return id, nil
}

func (w *strictUCWire) consume(id uint64) ([]byte, error) {
	slot := decodeWorkID(id).Slot
	if err := w.recv.consume(slot); err != nil {
		return nil, err
	}
	value := append([]byte(nil), w.buffers[slot]...)
	if _, err := w.recv.repost(slot); err != nil {
		return nil, err
	}
	return value, nil
}

func TestStrictUCShippingDriver(t *testing.T) {
	wire, err := newStrictUCWire(2)
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range [][]byte{[]byte("one"), []byte("two"), []byte("three"), []byte("four")} {
		id, err := wire.send(want)
		if err != nil {
			t.Fatal(err)
		}
		got, err := wire.consume(id)
		if err != nil {
			t.Fatal(err)
		}
		if string(got) != string(want) {
			t.Fatalf("received %q, want %q", got, want)
		}
	}
	if wire.dropped != 0 {
		t.Fatalf("dropped = %d, want 0", wire.dropped)
	}
}

// TestStrictUCRejectsCreditFreeDriver is a non-vacuous control. The frozen
// driver sends beyond the receive depth before consuming a completion. A
// transport that lets it succeed is not modeling UC flow control.
func TestStrictUCRejectsCreditFreeDriver(t *testing.T) {
	wire, err := newStrictUCWire(2)
	if err != nil {
		t.Fatal(err)
	}
	var ids []uint64
	for i := 0; i < 4; i++ {
		id, err := wire.send([]byte{byte(i)})
		if err == nil {
			ids = append(ids, id)
		}
	}
	if wire.dropped == 0 {
		t.Fatal("credit-free driver dropped no sends; strict UC adjudication is vacuous")
	}
	for _, id := range ids {
		if _, err := wire.consume(id); err != nil {
			t.Fatal(err)
		}
	}
}

// TestStrictUCRejectsRepostBeforeConsume is the corresponding slot-ownership
// control. A completed slot is still owned by its consumer.
func TestStrictUCRejectsRepostBeforeConsume(t *testing.T) {
	wire, err := newStrictUCWire(1)
	if err != nil {
		t.Fatal(err)
	}
	id, err := wire.send([]byte("first"))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := wire.recv.repost(decodeWorkID(id).Slot); !errors.Is(err, ErrProtocol) {
		t.Fatalf("premature repost error = %v, want ErrProtocol", err)
	}
}

// TestStrictUCFrozenRepostDriverCorrupts is the mutation control for slot
// ownership. It deliberately bypasses receiveSlots' public rejection, as a
// frozen broken driver would, and proves that reposting overwrites bytes still
// owned by the first completion.
func TestStrictUCFrozenRepostDriverCorrupts(t *testing.T) {
	wire, err := newStrictUCWire(1)
	if err != nil {
		t.Fatal(err)
	}
	first, err := wire.send([]byte("first"))
	if err != nil {
		t.Fatal(err)
	}
	slot := decodeWorkID(first).Slot
	// This assignment is the frozen defect: it claims consumption without the
	// consumer copying the completed bytes.
	wire.recv.slots[slot].state = receiveSlotConsumed
	if _, err := wire.recv.repost(slot); err != nil {
		t.Fatal(err)
	}
	if _, err := wire.send([]byte("second")); err != nil {
		t.Fatal(err)
	}
	got, err := wire.consume(first)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) == "first" {
		t.Fatalf("frozen repost driver preserved %q; strict UC corruption control is vacuous", got)
	}
}

// strictUCFabric is the test-only shipping driver. Unlike the negative
// controls above, it transfers chunks only after strictUCWire has a posted
// receive and consumes each completion before returning the next credit.
type strictUCFabric struct {
	wires map[[2]int]*strictUCWire
	drops int
}

func newStrictUCFabric(size, depth int) (*strictUCFabric, error) {
	fabric := &strictUCFabric{wires: make(map[[2]int]*strictUCWire)}
	for source := 0; source < size; source++ {
		for destination := 0; destination < size; destination++ {
			if source == destination {
				continue
			}
			wire, err := newStrictUCWire(depth)
			if err != nil {
				return nil, err
			}
			fabric.wires[[2]int{source, destination}] = wire
		}
	}
	return fabric, nil
}

func (f *strictUCFabric) transfer(source, destination int, src []byte, chunkSize int) ([]byte, error) {
	wire := f.wires[[2]int{source, destination}]
	if wire == nil || chunkSize <= 0 {
		return nil, ErrProtocol
	}
	var dst []byte
	for len(src) > 0 {
		n := min(chunkSize, len(src))
		id, err := wire.send(src[:n])
		if err != nil {
			f.drops += wire.dropped
			return nil, err
		}
		got, err := wire.consume(id)
		if err != nil {
			return nil, err
		}
		dst = append(dst, got...)
		src = src[n:]
	}
	f.drops += wire.dropped
	return dst, nil
}

func (f *strictUCFabric) allGather(inputs [][]byte, random *rand.Rand) ([][]byte, error) {
	outputs := make([][]byte, len(inputs))
	received := make([][][]byte, len(inputs))
	for destination := range received {
		received[destination] = make([][]byte, len(inputs))
	}
	order := make([]int, len(inputs))
	for rank := range order {
		order[rank] = rank
	}
	for source := range inputs {
		random.Shuffle(len(order), func(i, j int) { order[i], order[j] = order[j], order[i] })
		for _, destination := range order {
			if source == destination {
				continue
			}
			got, err := f.transfer(source, destination, inputs[source], 3)
			if err != nil {
				return nil, err
			}
			if !bytes.Equal(got, inputs[source]) {
				return nil, ErrProtocol
			}
			received[destination][source] = got
		}
	}
	for destination := range outputs {
		for source := range inputs {
			if source == destination {
				outputs[destination] = append(outputs[destination], inputs[source]...)
				continue
			}
			outputs[destination] = append(outputs[destination], received[destination][source]...)
		}
	}
	return outputs, nil
}

func TestStrictUCShippingDriverMultiRank(t *testing.T) {
	for _, size := range []int{2, 3} {
		t.Run(strconv.Itoa(size)+" ranks", func(t *testing.T) {
			fabric, err := newStrictUCFabric(size, 2)
			if err != nil {
				t.Fatal(err)
			}
			inputs := make([][]byte, size)
			var want []byte
			for rank := range inputs {
				inputs[rank] = bytes.Repeat([]byte{byte(rank + 1)}, 11)
				want = append(want, inputs[rank]...)
			}
			got, err := fabric.allGather(inputs, rand.New(rand.NewSource(int64(size))))
			if err != nil {
				t.Fatal(err)
			}
			for rank, output := range got {
				if !bytes.Equal(output, want) {
					t.Fatalf("rank %d gather = %v, want %v", rank, output, want)
				}
			}
			if fabric.drops != 0 {
				t.Fatalf("strict UC dropped %d shipping-driver sends", fabric.drops)
			}
		})
	}
}

func TestStrictUCShippingDriverReductions(t *testing.T) {
	types := []struct {
		name  string
		dtype DType
		input func(int) []byte
	}{
		{"bool", Bool, func(rank int) []byte { return []byte{byte(rank % 2)} }},
		{"int8", Int8, func(rank int) []byte { return []byte{byte(int8(rank - 2))} }},
		{"int16", Int16, func(rank int) []byte { return int16Bytes(int16(rank - 2)) }},
		{"uint8", Uint8, func(rank int) []byte { return []byte{byte(rank + 2), byte(9 - rank)} }},
		{"uint16", UInt16, func(rank int) []byte { return uint16Bytes(uint16(rank + 2)) }},
		{"uint32", UInt32, func(rank int) []byte { return uint32Bytes(uint32(rank + 2)) }},
		{"uint64", UInt64, func(rank int) []byte { return uint64Bytes(uint64(rank + 2)) }},
		{"int32", Int32, func(rank int) []byte {
			data := make([]byte, 8)
			binary.LittleEndian.PutUint32(data, uint32(int32(rank+2)))
			binary.LittleEndian.PutUint32(data[4:], uint32(int32(9-rank)))
			return data
		}},
		{"int64", Int64, func(rank int) []byte { return int64Bytes(int64(rank - 2)) }},
		{"float16", Float16, func(rank int) []byte { return uint16Bytes(float32ToHalf(float32(rank) + 1.5)) }},
		{"bfloat16", BFloat16, func(rank int) []byte { return uint16Bytes(float32ToBFloat16(float32(rank) + 1.5)) }},
		{"float32", Float32, func(rank int) []byte {
			data := make([]byte, 8)
			binary.LittleEndian.PutUint32(data, math.Float32bits(float32(rank)+1.5))
			binary.LittleEndian.PutUint32(data[4:], math.Float32bits(float32(8-rank)+0.5))
			return data
		}},
		{"float64", Float64, func(rank int) []byte { return float64Bytes(float64(rank) + 1.5) }},
		{"complex64", Complex64, func(rank int) []byte { return complex64Bytes(complex(float32(rank)+1.5, float32(8-rank)+0.5)) }},
	}
	for _, typ := range types {
		for _, op := range []ReduceOp{Sum, Min, Max} {
			t.Run(typ.name+"/"+reduceName(op), func(t *testing.T) {
				inputs := make([][]byte, 3)
				for rank := range inputs {
					inputs[rank] = typ.input(rank)
				}
				fabric, err := newStrictUCFabric(len(inputs), 2)
				if err != nil {
					t.Fatal(err)
				}
				gathered, err := fabric.allGather(inputs, rand.New(rand.NewSource(int64(op))))
				if err != nil {
					t.Fatal(err)
				}
				want := append([]byte(nil), inputs[0]...)
				for _, input := range inputs[1:] {
					if err := reduce(want, input, typ.dtype, op); err != nil {
						t.Fatal(err)
					}
				}
				for rank, data := range gathered {
					got := append([]byte(nil), data[:len(inputs[0])]...)
					for peer := 1; peer < len(inputs); peer++ {
						start := peer * len(inputs[0])
						if err := reduce(got, data[start:start+len(inputs[0])], typ.dtype, op); err != nil {
							t.Fatal(err)
						}
					}
					if !bytes.Equal(got, want) {
						t.Fatalf("rank %d reduction = %v, want %v", rank, got, want)
					}
				}
				if fabric.drops != 0 {
					t.Fatalf("strict UC dropped %d shipping-driver sends", fabric.drops)
				}
			})
		}
	}
}

func TestStrictUCRingGather(t *testing.T) {
	const (
		size  = 4
		wires = 3
	)
	fabric, err := newStrictUCRing(size, wires)
	if err != nil {
		t.Fatal(err)
	}
	sources := make([][]byte, size)
	destinations := make([][]byte, size)
	for rank := range sources {
		sources[rank] = bytes.Repeat([]byte{byte(rank + 1)}, 11)
		destinations[rank] = make([]byte, size*len(sources[rank]))
	}
	errs := make(chan error, size)
	for rank := range sources {
		go func(rank int) {
			errs <- ringGatherBlocks(rank, size, destinations[rank], sources[rank], wires, fabric.stage(rank))
		}(rank)
	}
	for range sources {
		if err := <-errs; err != nil {
			t.Fatal(err)
		}
	}
	for rank, output := range destinations {
		for source, want := range sources {
			got := output[source*len(want) : (source+1)*len(want)]
			if !bytes.Equal(got, want) {
				t.Fatalf("rank %d source %d = %x, want %x", rank, source, got, want)
			}
		}
	}
	if fabric.drops != 0 {
		t.Fatalf("strict UC ring dropped %d sends", fabric.drops)
	}
}

type strictUCRing struct {
	size  int
	wires int

	mu         sync.Mutex
	cond       *sync.Cond
	generation int
	arrived    int
	ids        []uint64
	receives   [][]byte
	links      map[[3]int]*strictUCWire
	drops      int
}

func newStrictUCRing(size, wires int) (*strictUCRing, error) {
	ring := &strictUCRing{
		size: size, wires: wires, ids: make([]uint64, size), receives: make([][]byte, size), links: make(map[[3]int]*strictUCWire),
	}
	ring.cond = sync.NewCond(&ring.mu)
	for source := 0; source < size; source++ {
		destination := (source + 1) % size
		for wire := 0; wire < wires; wire++ {
			link, err := newStrictUCWire(1)
			if err != nil {
				return nil, err
			}
			ring.links[[3]int{source, destination, wire}] = link
		}
	}
	return ring, nil
}

func (r *strictUCRing) stage(rank int) func(int, int, []byte, []byte) error {
	return func(step, wire int, send, receive []byte) error {
		destination := (rank + 1) % r.size
		link := r.links[[3]int{rank, destination, wire}]
		if link == nil {
			return ErrProtocol
		}
		id, err := link.send(send)
		if err != nil {
			r.mu.Lock()
			r.drops += link.dropped
			r.mu.Unlock()
			return err
		}
		r.mu.Lock()
		generation := r.generation
		r.ids[rank] = id
		r.receives[rank] = receive
		r.arrived++
		if r.arrived == r.size {
			for destination := 0; destination < r.size; destination++ {
				source := (destination + r.size - 1) % r.size
				inbound := r.links[[3]int{source, destination, wire}]
				data, err := inbound.consume(r.ids[source])
				if err != nil {
					r.mu.Unlock()
					return err
				}
				copy(r.receives[destination], data)
				r.drops += inbound.dropped
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

func reduceName(op ReduceOp) string {
	switch op {
	case Sum:
		return "sum"
	case Min:
		return "min"
	case Max:
		return "max"
	default:
		return "unknown"
	}
}
