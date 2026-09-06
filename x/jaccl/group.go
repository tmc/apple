package jaccl

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

// Group is a live RDMA communicator.
type Group struct {
	rank int
	size int

	mu         sync.Mutex
	closing    bool
	poison     error
	closeDone  chan struct{}
	closeErr   error
	operations sync.WaitGroup
	backend    groupBackend
}

type groupBackend interface {
	beginClose()
	close() error
	send(context.Context, int, []byte) error
	recv(context.Context, int, []byte) error
	barrier(context.Context) error
	allGather(context.Context, []byte, []byte) error
	allReduce(context.Context, []byte, []byte, DType, ReduceOp) error
}

// Open validates cfg and opens an RDMA group.
func Open(ctx context.Context, cfg Config) (*Group, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	if _, err := cfg.validate(); err != nil {
		return nil, err
	}
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	return open(ctx, cfg)
}

// OpenEnv reads a standalone JACCL-compatible environment configuration and
// opens the resulting group.
func OpenEnv(ctx context.Context) (*Group, error) {
	cfg, err := ConfigFromEnv()
	if err != nil {
		return nil, err
	}
	return Open(ctx, cfg)
}

// Rank reports the local rank.
func (g *Group) Rank() int {
	if g == nil {
		return -1
	}
	return g.rank
}

// Size reports the number of ranks in the group.
func (g *Group) Size() int {
	if g == nil {
		return 0
	}
	return g.size
}

// Close releases the group's resources. It is safe to call more than once.
func (g *Group) Close() error {
	if g == nil {
		return nil
	}
	g.mu.Lock()
	if g.closing {
		done := g.closeDone
		g.mu.Unlock()
		<-done
		return g.closeErr
	}
	g.closing = true
	g.closeDone = make(chan struct{})
	backend := g.backend
	g.mu.Unlock()

	if backend != nil {
		backend.beginClose()
	}
	g.operations.Wait()

	var err error
	if backend != nil {
		err = backend.close()
	}
	g.mu.Lock()
	g.closeErr = err
	close(g.closeDone)
	g.mu.Unlock()
	return err
}

// Send transfers src to dst.
func (g *Group) Send(ctx context.Context, dst int, src []byte) error {
	if g == nil {
		return ErrClosed
	}
	if dst < 0 || dst >= g.size || dst == g.rank {
		return fmt.Errorf("send destination %d: %w", dst, ErrProtocol)
	}
	return g.call(ctx, "send", func(b groupBackend) error { return b.send(ctx, dst, src) })
}

// Recv transfers bytes from src into dst.
func (g *Group) Recv(ctx context.Context, src int, dst []byte) error {
	if g == nil {
		return ErrClosed
	}
	if src < 0 || src >= g.size || src == g.rank {
		return fmt.Errorf("recv source %d: %w", src, ErrProtocol)
	}
	return g.call(ctx, "recv", func(b groupBackend) error { return b.recv(ctx, src, dst) })
}

// Barrier waits until every rank enters the barrier.
func (g *Group) Barrier(ctx context.Context) error {
	return g.call(ctx, "barrier", func(b groupBackend) error { return b.barrier(ctx) })
}

// AllGather gathers src from every rank into rank-major dst.
func (g *Group) AllGather(ctx context.Context, dst, src []byte) error {
	if g == nil {
		return ErrClosed
	}
	if len(src) > 0 && g.size > int(^uint(0)>>1)/len(src) {
		return fmt.Errorf("all gather length overflows: %w", ErrProtocol)
	}
	if len(dst) != g.size*len(src) {
		return fmt.Errorf("all gather destination length %d, want %d: %w", len(dst), g.size*len(src), ErrProtocol)
	}
	return g.call(ctx, "all gather", func(b groupBackend) error { return b.allGather(ctx, dst, src) })
}

// AllReduce reduces src into dst.
func (g *Group) AllReduce(ctx context.Context, dst, src []byte, dtype DType, op ReduceOp) error {
	if g == nil {
		return ErrClosed
	}
	width, err := dtype.Size()
	if err != nil {
		return err
	}
	if !op.valid() {
		return fmt.Errorf("reduce operation %d: %w", op, ErrProtocol)
	}
	if len(dst) != len(src) || len(src)%width != 0 {
		return fmt.Errorf("all reduce lengths dst=%d src=%d dtype=%d: %w", len(dst), len(src), dtype, ErrProtocol)
	}
	if slicesOverlap(dst, src) && !sameSlice(dst, src) {
		return fmt.Errorf("all reduce source and destination partially overlap: %w", ErrProtocol)
	}
	return g.call(ctx, "all reduce", func(b groupBackend) error { return b.allReduce(ctx, dst, src, dtype, op) })
}

// AllSum reduces src into dst using addition.
func (g *Group) AllSum(ctx context.Context, dst, src []byte, dtype DType) error {
	return g.AllReduce(ctx, dst, src, dtype, Sum)
}

// AllMax reduces src into dst using the greater value.
func (g *Group) AllMax(ctx context.Context, dst, src []byte, dtype DType) error {
	return g.AllReduce(ctx, dst, src, dtype, Max)
}

// AllMin reduces src into dst using the lesser value.
func (g *Group) AllMin(ctx context.Context, dst, src []byte, dtype DType) error {
	return g.AllReduce(ctx, dst, src, dtype, Min)
}

func (g *Group) call(ctx context.Context, operation string, fn func(groupBackend) error) error {
	if ctx == nil {
		ctx = context.Background()
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("%s: %w", operation, err)
	}
	if g == nil {
		return fmt.Errorf("%s: %w", operation, ErrClosed)
	}
	g.mu.Lock()
	if g.closing {
		g.mu.Unlock()
		return fmt.Errorf("%s: %w", operation, ErrClosed)
	}
	if g.poison != nil {
		err := g.poison
		g.mu.Unlock()
		return fmt.Errorf("%s: %w: %w", operation, ErrPoisoned, err)
	}
	backend := g.backend
	if backend == nil {
		g.mu.Unlock()
		return fmt.Errorf("%s: %w", operation, ErrUnsupported)
	}
	g.operations.Add(1)
	g.mu.Unlock()
	defer g.operations.Done()
	err := fn(backend)
	if err == nil || errors.Is(err, ErrClosed) {
		return err
	}
	g.mu.Lock()
	if !g.closing && g.poison == nil {
		g.poison = err
	}
	g.mu.Unlock()
	backend.beginClose()
	return err
}
