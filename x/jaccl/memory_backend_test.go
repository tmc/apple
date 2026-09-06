package jaccl

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

// memoryBackend is a reliable test transport. It exercises Group's public
// operations but is not used as evidence for UC flow control; strictUCWire is
// the separate adjudicating control for that purpose.
type memoryBackend struct {
	rank   int
	fabric *memoryFabric
	closed chan struct{}
	once   sync.Once
}

func (b *memoryBackend) beginClose()  { b.once.Do(func() { close(b.closed) }) }
func (b *memoryBackend) close() error { return nil }

func (b *memoryBackend) send(ctx context.Context, dst int, src []byte) error {
	return b.fabric.send(ctx, b.rank, dst, src)
}

func (b *memoryBackend) recv(ctx context.Context, src int, dst []byte) error {
	return b.fabric.recv(ctx, src, b.rank, dst)
}

func (b *memoryBackend) barrier(ctx context.Context) error {
	_, err := b.fabric.collect(ctx, b.rank, memoryBarrier, nil, 0, 0)
	return err
}

func (b *memoryBackend) allGather(ctx context.Context, dst, src []byte) error {
	result, err := b.fabric.collect(ctx, b.rank, memoryGather, src, 0, 0)
	if err != nil {
		return err
	}
	copy(dst, result)
	return nil
}

func (b *memoryBackend) allReduce(ctx context.Context, dst, src []byte, dtype DType, op ReduceOp) error {
	result, err := b.fabric.collect(ctx, b.rank, memoryReduce, src, dtype, op)
	if err != nil {
		return err
	}
	copy(dst, result)
	return nil
}

type memoryOperation uint8

const (
	memoryBarrier memoryOperation = iota + 1
	memoryGather
	memoryReduce
)

type memoryFabric struct {
	size int

	mu       sync.Mutex
	round    *memoryRound
	messages map[[2]int]chan []byte
}

type memoryRound struct {
	kind    memoryOperation
	dtype   DType
	op      ReduceOp
	entered []bool
	inputs  [][]byte
	result  []byte
	err     error
	ready   chan struct{}
	drained chan struct{}
	closed  bool
	waited  int
}

func newMemoryGroups(size int) []*Group {
	fabric := &memoryFabric{size: size, messages: make(map[[2]int]chan []byte)}
	groups := make([]*Group, size)
	for rank := range groups {
		groups[rank] = &Group{rank: rank, size: size, backend: &memoryBackend{
			rank: rank, fabric: fabric, closed: make(chan struct{}),
		}}
	}
	return groups
}

func (f *memoryFabric) collect(ctx context.Context, rank int, kind memoryOperation, input []byte, dtype DType, op ReduceOp) ([]byte, error) {
	f.mu.Lock()
	for f.round != nil && f.round.closed {
		drained := f.round.drained
		f.mu.Unlock()
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-drained:
		}
		f.mu.Lock()
	}
	r := f.round
	if r == nil {
		r = &memoryRound{kind: kind, dtype: dtype, op: op, entered: make([]bool, f.size), inputs: make([][]byte, f.size), ready: make(chan struct{}), drained: make(chan struct{})}
		f.round = r
	}
	if r.kind != kind || r.dtype != dtype || r.op != op || r.entered[rank] {
		f.failRoundLocked(r, fmt.Errorf("memory collective mismatch: %w", ErrProtocol))
	}
	if r.err == nil {
		r.entered[rank] = true
		r.inputs[rank] = append([]byte(nil), input...)
		if allRanksEntered(r.entered) {
			r.result, r.err = memoryCollect(r)
			r.closed = true
			close(r.ready)
		}
	}
	ready := r.ready
	f.mu.Unlock()

	select {
	case <-ctx.Done():
		f.mu.Lock()
		f.failRoundLocked(r, ctx.Err())
		f.mu.Unlock()
		return nil, ctx.Err()
	case <-ready:
	}

	f.mu.Lock()
	defer f.mu.Unlock()
	r.waited++
	if r.waited == f.size {
		f.round = nil
		close(r.drained)
	}
	if r.err != nil {
		return nil, r.err
	}
	return append([]byte(nil), r.result...), nil
}

func (f *memoryFabric) failRoundLocked(r *memoryRound, err error) {
	if r.closed {
		return
	}
	r.err = err
	r.closed = true
	close(r.ready)
	if f.round == r {
		f.round = nil
	}
	close(r.drained)
}

func allRanksEntered(entered []bool) bool {
	for _, entered := range entered {
		if !entered {
			return false
		}
	}
	return true
}

func memoryCollect(r *memoryRound) ([]byte, error) {
	switch r.kind {
	case memoryBarrier:
		return nil, nil
	case memoryGather:
		var result []byte
		for _, input := range r.inputs {
			result = append(result, input...)
		}
		return result, nil
	case memoryReduce:
		result := append([]byte(nil), r.inputs[0]...)
		for _, input := range r.inputs[1:] {
			if err := reduce(result, input, r.dtype, r.op); err != nil {
				return nil, err
			}
		}
		return result, nil
	default:
		return nil, fmt.Errorf("memory collective kind %d: %w", r.kind, ErrProtocol)
	}
}

func (f *memoryFabric) message(src, dst int) chan []byte {
	key := [2]int{src, dst}
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.messages[key] == nil {
		f.messages[key] = make(chan []byte, 1)
	}
	return f.messages[key]
}

func (f *memoryFabric) send(ctx context.Context, src, dst int, data []byte) error {
	select {
	case f.message(src, dst) <- append([]byte(nil), data...):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (f *memoryFabric) recv(ctx context.Context, src, dst int, data []byte) error {
	select {
	case got := <-f.message(src, dst):
		if len(got) != len(data) {
			return fmt.Errorf("memory receive length %d, want %d: %w", len(got), len(data), ErrProtocol)
		}
		copy(data, got)
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func TestGroupCollectivesWithMemoryBackend(t *testing.T) {
	groups := newMemoryGroups(3)
	runGroups(t, groups, func(rank int, g *Group) error {
		if err := g.Barrier(context.Background()); err != nil {
			return err
		}
		src := []byte{byte(rank + 1), byte(rank + 11)}
		gathered := make([]byte, len(groups)*len(src))
		if err := g.AllGather(context.Background(), gathered, src); err != nil {
			return err
		}
		wantGathered := []byte{1, 11, 2, 12, 3, 13}
		if string(gathered) != string(wantGathered) {
			return fmt.Errorf("all gather = %v, want %v", gathered, wantGathered)
		}
		reduced := make([]byte, len(src))
		if err := g.AllSum(context.Background(), reduced, src, Uint8); err != nil {
			return err
		}
		if string(reduced) != string([]byte{6, 36}) {
			return fmt.Errorf("all sum = %v, want [6 36]", reduced)
		}
		if err := g.AllMax(context.Background(), reduced, src, Uint8); err != nil {
			return err
		}
		if string(reduced) != string([]byte{3, 13}) {
			return fmt.Errorf("all max = %v, want [3 13]", reduced)
		}
		if err := g.AllMin(context.Background(), reduced, src, Uint8); err != nil {
			return err
		}
		if string(reduced) != string([]byte{1, 11}) {
			return fmt.Errorf("all min = %v, want [1 11]", reduced)
		}
		return nil
	})
}

func TestGroupSendRecvWithMemoryBackend(t *testing.T) {
	groups := newMemoryGroups(2)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	err := make(chan error, 2)
	go func() { err <- groups[0].Send(ctx, 1, []byte("hello")) }()
	go func() {
		got := make([]byte, 5)
		if e := groups[1].Recv(ctx, 0, got); e != nil {
			err <- e
			return
		}
		if string(got) != "hello" {
			err <- fmt.Errorf("received %q, want hello", got)
			return
		}
		err <- nil
	}()
	for range 2 {
		if err := <-err; err != nil {
			t.Fatal(err)
		}
	}
}

func runGroups(t *testing.T, groups []*Group, fn func(int, *Group) error) {
	t.Helper()
	err := make(chan error, len(groups))
	for rank, group := range groups {
		go func(rank int, group *Group) { err <- fn(rank, group) }(rank, group)
	}
	for range groups {
		if err := <-err; err != nil {
			t.Fatal(err)
		}
	}
}
