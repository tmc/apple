//go:build darwin

package irecovery

import (
	"context"
	"errors"
	"fmt"
	"time"
)

type finalization struct {
	done chan struct{}
	err  error
}

// FinalizeDFU completes a successful Upload's DFU download and retires the
// connection. It sends the next zero-length block, waits for manifestation, and
// resets USB. Success requires a later WaitOpen to establish the next endpoint;
// it does not prove recovery, restore or boot completion. Calling Control or
// BulkWrite after Upload invalidates the upload recorded for finalization.
//
// Cancellation stops the caller's wait. Native reset cannot be interrupted;
// ownership stays with the operation until reset and cleanup finish. Another
// FinalizeDFU call waits for the same result without repeating writes. Close also
// waits for completion. No method can reuse the retired native handle.
func (c *Conn) FinalizeDFU(ctx context.Context) error {
	if err := c.lock(ctx); err != nil {
		return err
	}
	op := c.finalization
	if op == nil {
		if c.library == nil {
			c.mu.Unlock()
			return fmt.Errorf("recovery connection is closed")
		}
		if c.info.Mode != "dfu" || c.dfuBlocks == 0 {
			c.mu.Unlock()
			return fmt.Errorf("no completed dfu upload to finalize")
		}
		op = &finalization{done: make(chan struct{})}
		worker := &Conn{library: c.library, handle: c.handle, info: c.info}
		blocks := c.dfuBlocks
		c.library = nil
		c.handle = 0
		c.dfuBlocks = 0
		c.finalization = op
		// The worker retains both the library and native handle through blocking reset
		// and cleanup, even when the original caller has already stopped waiting.
		go func() {
			op.err = errors.Join(worker.finishDFU(ctx, blocks), worker.release(true))
			close(op.done)
		}()
	}
	c.mu.Unlock()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-op.done:
		return op.err
	}
}

func (c *Conn) finishDFU(ctx context.Context, blocks uint16) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Minute)
	defer cancel()
	if err := c.controlExact(ctx, 0x21, 1, blocks, nil); err != nil {
		return fmt.Errorf("notify dfu manifestation: %w", err)
	}
	for {
		var response [6]byte
		if err := c.controlExact(ctx, 0xa1, 3, 0, response[:]); err != nil {
			return fmt.Errorf("read manifestation status: %w", err)
		}
		status, state := response[0], response[4]
		if status != 0 {
			return &StatusError{Status: status, State: state}
		}
		if state != 2 && state != 6 && state != 7 && state != 8 {
			return &StatusError{Status: status, State: state}
		}
		delay := time.Duration(uint32(response[1])|uint32(response[2])<<8|uint32(response[3])<<16) * time.Millisecond
		if delay > 0 || state == 6 || state == 7 {
			timer := time.NewTimer(max(delay, time.Millisecond))
			select {
			case <-ctx.Done():
				timer.Stop()
				return ctx.Err()
			case <-timer.C:
			}
		}
		if state == 2 || state == 8 {
			break
		}
	}
	if err := ctx.Err(); err != nil {
		return err
	}
	code := c.library.reset(c.handle)
	// Both removal and the documented rediscovery result retire this handle.
	// Neither proves boot; callers must rediscover and check the target identity.
	if code != 0 && code != -4 && code != -5 {
		return fmt.Errorf("reset recovery device: %d", code)
	}
	return nil
}

// WaitOpen waits for an unambiguous endpoint with the exact ECID and mode
// ("dfu" or "recovery"), then claims its interface. It has a two-minute ceiling
// or the caller's earlier deadline. It retries absence, not permission or claim
// failures. Call only after finalization has finished releasing the old handle.
// The returned connection must be closed by the caller.
func WaitOpen(ctx context.Context, path string, ecid uint64, mode string) (*Conn, error) {
	if ecid == 0 {
		return nil, fmt.Errorf("nonzero ECID is required")
	}
	if mode != "dfu" && mode != "recovery" {
		return nil, fmt.Errorf("invalid recovery mode %q", mode)
	}
	return waitOpen(ctx, func(ctx context.Context) (*Conn, error) { return openDevice(ctx, path, ecid, mode) })
}

func waitOpen(ctx context.Context, open func(context.Context) (*Conn, error)) (*Conn, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Minute)
	defer cancel()
	for {
		if err := ctx.Err(); err != nil {
			return nil, err
		}
		conn, err := open(ctx)
		if err == nil {
			return conn, nil
		}
		if !errors.Is(err, ErrNotFound) {
			return nil, err
		}
		timer := time.NewTimer(100 * time.Millisecond)
		select {
		case <-ctx.Done():
			timer.Stop()
			return nil, ctx.Err()
		case <-timer.C:
		}
	}
}
