//go:build darwin

package irecovery

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// WaitDisconnected waits for the existing handle to report device removal. It
// does not close the handle or select a replacement device. Call after a command
// that acknowledged a boot transition, then Close before WaitOpen. Read timeouts
// and pipe stalls are retried; neither establishes disconnection. Other errors
// are returned. The wait is bounded
// by two minutes or the caller's earlier deadline.
func (c *Conn) WaitDisconnected(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Minute)
	defer cancel()
	for {
		if err := c.lock(ctx); err != nil {
			return err
		}
		var descriptor [18]byte
		n, err := c.control(ctx, 0x80, 6, 0x0100, 0, descriptor[:])
		c.mu.Unlock()
		if errors.Is(err, ErrNotFound) {
			return nil
		}
		if err != nil {
			var failure *usbFailure
			if !errors.As(err, &failure) || (failure.code != -7 && failure.code != -9) {
				return err
			}
		} else if n != len(descriptor) || descriptor[0] != 18 || descriptor[1] != 1 {
			return fmt.Errorf("invalid USB presence descriptor")
		}
		timer := time.NewTimer(100 * time.Millisecond)
		select {
		case <-ctx.Done():
			timer.Stop()
			return ctx.Err()
		case <-timer.C:
		}
	}
}
