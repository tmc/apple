package portfwd

import (
	"context"
	"io"
	"net"
	"sync"
	"time"
)

// RelayConfig configures a [RelayTCP] loop.
type RelayConfig struct {
	// Listener accepts host-side connections to forward. RelayTCP closes it on
	// return. The caller chooses the bind address (e.g. ":8080" for all
	// interfaces or "127.0.0.1:8080" for loopback only).
	Listener net.Listener

	// Dial opens a backend connection for one accepted host connection. It is
	// the only transport-specific seam: it may dial a guest TCP address, a
	// vsock port, or anything else yielding a net.Conn. Required.
	Dial func(ctx context.Context) (net.Conn, error)

	// OnConns, if non-nil, is called with +1 when a connection is accepted
	// (before the backend is dialed) and -1 when its relay finishes. The
	// running sum is the active-connection count; the number of +1 calls is the
	// cumulative count of accepted connections, including any whose backend dial
	// failed. It may be called concurrently from multiple goroutines.
	OnConns func(delta int)
}

// RelayTCP accepts connections on cfg.Listener and splices each to a backend
// from cfg.Dial until ctx is canceled or the listener stops accepting. It
// closes the listener before returning and returns ctx.Err() (nil if the
// listener was closed without a context cancellation).
func RelayTCP(ctx context.Context, cfg RelayConfig) error {
	defer cfg.Listener.Close()

	go func() {
		<-ctx.Done()
		cfg.Listener.Close()
	}()

	for {
		conn, err := cfg.Listener.Accept()
		if err != nil {
			if ctx.Err() != nil {
				return ctx.Err()
			}
			// The listener was closed (or hit a terminal error) without a
			// context cancellation; stop the loop without surfacing the
			// post-Close accept error.
			return nil
		}
		go relayConn(ctx, conn, cfg)
	}
}

func relayConn(ctx context.Context, host net.Conn, cfg RelayConfig) {
	defer host.Close()

	// Count the connection from acceptance, before dialing the backend, so a
	// caller tracking cumulative accepts (via the +1 calls) and active conns
	// (via the running sum) sees every accepted connection, including one whose
	// backend dial fails.
	if cfg.OnConns != nil {
		cfg.OnConns(1)
		defer cfg.OnConns(-1)
	}

	backend, err := cfg.Dial(ctx)
	if err != nil {
		return
	}
	defer backend.Close()

	splice(host, backend)
}

// splice copies bytes in both directions between a and b until both halves
// finish. After each direction's copy ends it half-closes the write side of
// the destination so the peer sees EOF: a conn implementing CloseWrite via
// that, anything else (e.g. a vsock conn) via a zero write deadline to unblock
// a blocked read.
func splice(a, b net.Conn) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		io.Copy(b, a)
		halfCloseWrite(b)
	}()
	go func() {
		defer wg.Done()
		io.Copy(a, b)
		halfCloseWrite(a)
	}()
	wg.Wait()
}

func halfCloseWrite(c net.Conn) {
	if cw, ok := c.(interface{ CloseWrite() error }); ok {
		cw.CloseWrite()
		return
	}
	if d, ok := c.(interface{ SetWriteDeadline(time.Time) error }); ok {
		d.SetWriteDeadline(time.Now())
	}
}
