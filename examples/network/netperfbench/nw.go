//go:build darwin

package main

// Network.framework transport, driven through the Go bindings. Both roles
// run their callbacks on a dispatch queue and hand results to the calling
// goroutine over channels, so the benchmark loop itself stays synchronous
// and matches what the std transport does.

import (
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/network"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// readyTimeout bounds how long a connection or listener may take to reach
// the ready state before the benchmark gives up.
const readyTimeout = 10 * time.Second

// plainTCPParameters returns cleartext TCP parameters with Nagle disabled,
// matching the SetNoDelay the std transport applies.
func plainTCPParameters() network.NWParameters {
	return network.NWParametersCreatePlainTCP(func(options network.NWProtocolOptions) {
		network.NWTCPOptionsSetNoDelay(options, true)
	})
}

func nwErr(e network.NWError) error {
	if e.IsZero() {
		return nil
	}
	return fmt.Errorf("nw error %d: %s", e.Code(), e.Error())
}

type nwServer struct {
	listener network.NWListener
	queue    dispatch.Queue
	port     string
}

func serveNW(port string) (echoServer, error) {
	params := plainTCPParameters()
	defer params.Release()

	listener := network.NWListenerCreateWithPort(port, params)
	if listener.ID == 0 {
		return nil, fmt.Errorf("nw_listener_create_with_port(%s) failed", port)
	}
	queue := dispatch.QueueCreate("netperfbench.server")
	network.NWListenerSetQueue(listener, queue)

	ready := make(chan error, 1)
	network.NWListenerSetStateChangedHandler(listener, func(state network.NWListenerState, e network.NWError) {
		switch state {
		case network.NWListenerStateReady:
			send(ready, nil)
		case network.NWListenerStateFailed, network.NWListenerStateCancelled:
			send(ready, fmt.Errorf("listener state %v: %w", state, nwErr(e)))
		}
	})
	network.NWListenerSetNewConnectionHandler(listener, func(conn network.NWConnection) {
		startEcho(conn, queue)
	})
	network.NWListenerStart(listener)

	select {
	case err := <-ready:
		if err != nil {
			return nil, err
		}
	case <-time.After(readyTimeout):
		return nil, fmt.Errorf("listener not ready after %v", readyTimeout)
	}
	return &nwServer{
		listener: listener,
		queue:    queue,
		port:     fmt.Sprint(network.NWListenerGetPort(listener)),
	}, nil
}

func (s *nwServer) Port() string { return s.port }

func (s *nwServer) Close() { network.NWListenerCancel(s.listener) }

// startEcho drives one accepted connection: every received chunk is sent
// straight back, and the next receive is scheduled from the send
// completion so at most one operation is outstanding per direction.
func startEcho(conn network.NWConnection, queue dispatch.Queue) {
	network.NWConnectionSetQueue(conn, queue)
	context := network.NWContentContextCreate("netperfbench.echo")
	var pump func()
	pump = func() {
		network.NWConnectionReceive(conn, 1, 1<<16, func(content objectivec.Object, _ network.NWContentContext, _ bool, e network.NWError) {
			if !e.IsZero() || content.ID == 0 {
				network.NWConnectionCancel(conn)
				context.Release()
				return
			}
			data := dispatch.DataFromHandle(uintptr(content.ID))
			network.NWConnectionSend(conn, data, context, false, func(sendErr network.NWError) {
				if !sendErr.IsZero() {
					network.NWConnectionCancel(conn)
					context.Release()
					return
				}
				pump()
			})
		})
	}
	network.NWConnectionSetStateChangedHandler(conn, func(state network.NWConnectionState, _ network.NWError) {
		if state == network.NWConnectionStateReady {
			pump()
		}
	})
	network.NWConnectionStart(conn)
}

type nwClient struct {
	conn      network.NWConnection
	queue     dispatch.Queue
	context   network.NWContentContext
	sent      chan error
	got       chan recvResult
	cancelled chan struct{}
}

type recvResult struct {
	n   int
	err error
}

func dialNW(addr string, inflight int) (echoClient, error) {
	host, port, err := net.SplitHostPort(addr)
	if err != nil {
		return nil, err
	}
	params := plainTCPParameters()
	defer params.Release()

	endpoint := network.NWEndpointCreateHost(host, port)
	conn := network.NWConnectionCreate(endpoint, params)
	if conn.ID == 0 {
		return nil, fmt.Errorf("nw_connection_create(%s) failed", addr)
	}
	queue := dispatch.QueueCreate("netperfbench.client")
	network.NWConnectionSetQueue(conn, queue)

	ready := make(chan error, 1)
	cancelled := make(chan struct{})
	var closeOnce sync.Once
	network.NWConnectionSetStateChangedHandler(conn, func(state network.NWConnectionState, e network.NWError) {
		switch state {
		case network.NWConnectionStateReady:
			send(ready, nil)
		case network.NWConnectionStateFailed:
			send(ready, fmt.Errorf("connection state %v: %w", state, nwErr(e)))
		case network.NWConnectionStateCancelled:
			send(ready, fmt.Errorf("connection state %v: %w", state, nwErr(e)))
			closeOnce.Do(func() { close(cancelled) })
		}
	})
	network.NWConnectionStart(conn)

	select {
	case err := <-ready:
		if err != nil {
			return nil, err
		}
	case <-time.After(readyTimeout):
		return nil, fmt.Errorf("connection not ready after %v", readyTimeout)
	}
	return &nwClient{
		conn:      conn,
		queue:     queue,
		context:   network.NWContentContextCreate("netperfbench.send"),
		sent:      make(chan error, max(inflight, 1)),
		got:       make(chan recvResult, 1),
		cancelled: cancelled,
	}, nil
}

// Path reports the path Network.framework actually selected, which is what
// -require-interface and -forbid-loopback assert against.
func (c *nwClient) Path() *pathInfo {
	path := network.NWConnectionCopyCurrentPath(c.conn)
	if path.ID == 0 {
		return nil
	}
	defer path.Release()

	info := &pathInfo{
		Status:       network.NWPathGetStatus(path).String(),
		UsesLoopback: network.NWPathUsesInterfaceType(path, network.NWInterfaceTypeLoopback),
		UsesWifi:     network.NWPathUsesInterfaceType(path, network.NWInterfaceTypeWifi),
		UsesWired:    network.NWPathUsesInterfaceType(path, network.NWInterfaceTypeWired),
	}
	network.NWPathEnumerateInterfaces(path, func(iface network.NWInterface) bool {
		info.Interfaces = append(info.Interfaces, objc.GoString(network.NWInterfaceGetName(iface)))
		return true
	})
	return info
}

// RoundTrip sends n copies of buf before waiting for any of them, then
// reads all n echoes back. With n == 1 the measurement is pure latency;
// with n > 1 the per-operation costs overlap, which is what separates a
// fixed cost per send from a cost per byte.
func (c *nwClient) RoundTrip(buf []byte, n int) error {
	roundTrip := beginSignpost("round-trip")
	defer roundTrip.end()

	data := dispatch.DataCreate(buf)
	defer data.Release()
	for range n {
		send := beginSignpost("send")
		network.NWConnectionSend(c.conn, data, c.context, false, func(e network.NWError) {
			send.end()
			send.event("send-callback")
			send.event("send-signal")
			c.sent <- nwErr(e)
		})
	}

	// The default keeps one receive per echo. -recv-batch is an experimental
	// alternative that moves receive re-arming into Network.framework and
	// crosses back into Go once for the outstanding batch.
	remaining := len(buf) * n
	if *recvBatch && uint64(len(buf))*uint64(n) > uint64(^uint32(0)) {
		return fmt.Errorf("receive length exceeds uint32_t")
	}
	for remaining > 0 {
		want := min(remaining, len(buf))
		if *recvBatch {
			want = remaining
		}
		receive := beginSignpost("receive")
		// The callback runs on a dispatch worker. receive-signal marks the
		// channel handoff there, while receive-wait ends on the benchmark
		// goroutine after the Go scheduler has made it runnable again.
		wait := beginSignpost("receive-wait")
		network.NWConnectionReceive(c.conn, uint32(want), uint32(want), func(content objectivec.Object, _ network.NWContentContext, _ bool, e network.NWError) {
			receive.end()
			receive.event("receive-callback")
			wait.event("receive-signal")
			if err := nwErr(e); err != nil {
				c.got <- recvResult{err: err}
				return
			}
			if content.ID == 0 {
				c.got <- recvResult{err: fmt.Errorf("connection closed by peer")}
				return
			}
			c.got <- recvResult{n: dispatch.DataFromHandle(uintptr(content.ID)).Len()}
		})
		r := <-c.got
		wait.end()
		if r.err != nil {
			return r.err
		}
		if r.n != want {
			return fmt.Errorf("echo returned %d bytes, want %d", r.n, want)
		}
		remaining -= r.n
	}

	for range n {
		wait := beginSignpost("send-wait")
		err := <-c.sent
		wait.end()
		if err != nil {
			return err
		}
	}
	return nil
}

// Close tears the connection down and waits for it to reach the cancelled
// state before releasing. Releasing a connection that has not finished
// cancelling strands its kernel-side buffers, which a matrix runner opening
// many connections notices as both a memory cliff and a slow drift in later
// measurements.
func (c *nwClient) Close() {
	network.NWConnectionForceCancel(c.conn)
	select {
	case <-c.cancelled:
	case <-time.After(readyTimeout):
	}
	c.context.Release()
}

// send delivers on a buffered channel without blocking if the state
// handler fires more than once for the same terminal condition.
func send(ch chan error, err error) {
	select {
	case ch <- err:
	default:
	}
}
