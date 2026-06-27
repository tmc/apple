package portfwd

import (
	"context"
	"io"
	"net"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// echoServer accepts TCP connections and echoes each back, upper-cased only by
// reversing nothing — it just mirrors bytes, which is enough to prove the relay
// moves data in both directions.
func echoServer(t *testing.T) (addr string, stop func()) {
	t.Helper()
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	var wg sync.WaitGroup
	go func() {
		for {
			c, err := ln.Accept()
			if err != nil {
				return
			}
			wg.Add(1)
			go func() {
				defer wg.Done()
				defer c.Close()
				io.Copy(c, c)
			}()
		}
	}()
	return ln.Addr().String(), func() {
		ln.Close()
		wg.Wait()
	}
}

func TestRelayTCP(t *testing.T) {
	backendAddr, stopBackend := echoServer(t)
	defer stopBackend()

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}

	var active int64
	var maxActive int64
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	relayDone := make(chan error, 1)
	go func() {
		relayDone <- RelayTCP(ctx, RelayConfig{
			Listener: ln,
			Dial: func(ctx context.Context) (net.Conn, error) {
				return net.Dial("tcp", backendAddr)
			},
			OnConns: func(delta int) {
				n := atomic.AddInt64(&active, int64(delta))
				for {
					m := atomic.LoadInt64(&maxActive)
					if n <= m || atomic.CompareAndSwapInt64(&maxActive, m, n) {
						break
					}
				}
			},
		})
	}()

	// Drive a few connections through the relay and check the echo round-trips.
	for i := 0; i < 3; i++ {
		conn, err := net.Dial("tcp", ln.Addr().String())
		if err != nil {
			t.Fatalf("dial relay: %v", err)
		}
		msg := []byte("hello relay")
		if _, err := conn.Write(msg); err != nil {
			t.Fatalf("write: %v", err)
		}
		got := make([]byte, len(msg))
		if _, err := io.ReadFull(conn, got); err != nil {
			t.Fatalf("read: %v", err)
		}
		if string(got) != string(msg) {
			t.Errorf("echo = %q, want %q", got, msg)
		}
		conn.Close()
	}

	// Shutdown via context cancel; RelayTCP must return ctx.Err().
	cancel()
	select {
	case err := <-relayDone:
		if err != context.Canceled {
			t.Errorf("RelayTCP returned %v, want context.Canceled", err)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("RelayTCP did not return after context cancel")
	}

	// Active connections drain asynchronously after the relay returns (it does
	// not block on in-flight conns, matching the donor behavior). Poll briefly.
	deadline := time.Now().Add(2 * time.Second)
	for atomic.LoadInt64(&active) != 0 && time.Now().Before(deadline) {
		time.Sleep(5 * time.Millisecond)
	}
	if n := atomic.LoadInt64(&active); n != 0 {
		t.Errorf("active connections = %d after drain, want 0", n)
	}
	if atomic.LoadInt64(&maxActive) < 1 {
		t.Errorf("OnConns never reported an active connection")
	}
}

func TestRelayTCPDialFailureClosesHostConn(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go RelayTCP(ctx, RelayConfig{
		Listener: ln,
		Dial: func(ctx context.Context) (net.Conn, error) {
			return nil, io.EOF // always fail to dial the backend
		},
	})

	conn, err := net.Dial("tcp", ln.Addr().String())
	if err != nil {
		t.Fatalf("dial relay: %v", err)
	}
	defer conn.Close()

	// With no backend, the relay closes the host connection immediately, so a
	// read returns EOF rather than hanging.
	conn.SetReadDeadline(time.Now().Add(2 * time.Second))
	if _, err := conn.Read(make([]byte, 1)); err == nil {
		t.Error("expected host connection to be closed when dial fails")
	}
}

func TestRelayTCPOnConnsCountsFailedDial(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var opens, closes int64
	go RelayTCP(ctx, RelayConfig{
		Listener: ln,
		Dial: func(ctx context.Context) (net.Conn, error) {
			return nil, io.EOF // backend dial always fails
		},
		OnConns: func(delta int) {
			if delta > 0 {
				atomic.AddInt64(&opens, 1)
			} else {
				atomic.AddInt64(&closes, 1)
			}
		},
	})

	conn, err := net.Dial("tcp", ln.Addr().String())
	if err != nil {
		t.Fatalf("dial relay: %v", err)
	}
	conn.Close()

	// Even though the backend dial fails, the accepted connection must be
	// counted (opened then closed).
	deadline := time.Now().Add(2 * time.Second)
	for atomic.LoadInt64(&closes) == 0 && time.Now().Before(deadline) {
		time.Sleep(5 * time.Millisecond)
	}
	if atomic.LoadInt64(&opens) != 1 {
		t.Errorf("opens = %d, want 1 (accepted conn must count even on dial failure)", opens)
	}
	if atomic.LoadInt64(&closes) != 1 {
		t.Errorf("closes = %d, want 1", closes)
	}
}

func TestRelayTCPListenerCloseStops(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	done := make(chan error, 1)
	go func() {
		done <- RelayTCP(context.Background(), RelayConfig{
			Listener: ln,
			Dial:     func(ctx context.Context) (net.Conn, error) { return nil, io.EOF },
		})
	}()
	// Closing the listener (no context cancellation) stops the loop with nil.
	ln.Close()
	select {
	case err := <-done:
		if err != nil {
			t.Errorf("RelayTCP returned %v, want nil on listener close", err)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("RelayTCP did not return after listener close")
	}
}
