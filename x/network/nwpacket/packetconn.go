//go:build darwin

package nwpacket

import (
	"context"
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/tmc/apple/dispatch"
	applenetwork "github.com/tmc/apple/network"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// Config describes a Network.framework UDP packet connection.
type Config struct {
	// InterfaceName is used for link-local zones and optional exact interface
	// selection.
	InterfaceName string

	// LocalAddr is the local UDP address to bind. Port 0 asks Network.framework
	// to choose a port.
	LocalAddr *net.UDPAddr

	RequiredInterfaceType applenetwork.NWInterfaceType
	SetRequiredInterface  bool
	IncludePeerToPeer     bool
	RequireInterface      bool
	ReuseLocalAddress     bool

	// ConnectTimeout bounds outbound connection readiness before retry. It also
	// bounds send completion when the caller has not set a write deadline. Zero
	// means five seconds.
	ConnectTimeout time.Duration

	// ConnectRetries is the number of extra outbound connection attempts after
	// a readiness timeout. Retries cancel and recreate only outbound peers.
	ConnectRetries int

	QueueLabel string
	Tracef     func(format string, args ...any)
}

type nwPacket struct {
	data []byte
	addr *net.UDPAddr
}

type nwPeerConn struct {
	conn     applenetwork.NWConnection
	addr     *net.UDPAddr
	ready    chan error
	outbound bool
	once     sync.Once
}

type nwPacketConn struct {
	config Config

	queue    dispatch.Queue
	listener applenetwork.NWListener
	local    *net.UDPAddr

	packets chan nwPacket
	closed  chan struct{}
	once    sync.Once

	mu                   sync.Mutex
	conns                map[string]*nwPeerConn
	sendSeq              atomic.Uint64
	readDeadline         time.Time
	writeDeadline        time.Time
	readDeadlineChanged  chan struct{}
	writeDeadlineChanged chan struct{}
}

// ListenPacket creates a Network.framework-backed net.PacketConn.
func ListenPacket(config Config) (net.PacketConn, error) {
	return ListenPacketContext(context.Background(), config)
}

// ListenPacketContext creates a Network.framework-backed net.PacketConn.
//
// The context bounds listener startup. After startup succeeds, use deadlines
// and Close to control packet I/O.
func ListenPacketContext(ctx context.Context, config Config) (net.PacketConn, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	if config.LocalAddr == nil {
		return nil, errors.New("missing local address")
	}
	if config.InterfaceName == "" && config.LocalAddr.Zone != "" {
		config.InterfaceName = config.LocalAddr.Zone
	}
	if config.QueueLabel == "" {
		config.QueueLabel = "com.github.tmc.apple.x.network.nwpacket"
	}
	local := copyUDPAddr(config.LocalAddr)
	params, err := newNWParameters(config, local)
	if err != nil {
		return nil, err
	}
	listener := applenetwork.NWListenerCreate(params)
	if listener.ID == 0 {
		return nil, errors.New("nw_listener_create returned nil")
	}

	c := &nwPacketConn{
		config:               config,
		queue:                dispatch.QueueCreate(config.QueueLabel),
		listener:             listener,
		local:                local,
		packets:              make(chan nwPacket, 16384),
		closed:               make(chan struct{}),
		conns:                make(map[string]*nwPeerConn),
		readDeadlineChanged:  make(chan struct{}),
		writeDeadlineChanged: make(chan struct{}),
	}

	ready := make(chan error, 1)
	applenetwork.NWListenerSetQueue(listener, c.queue)
	applenetwork.NWListenerSetStateChangedHandler(listener, func(state applenetwork.NWListenerState, nwErr applenetwork.NWError) {
		c.tracef("listener %s state=%s err=%s", c.local, state, nwErrorString(nwErr))
		switch state {
		case applenetwork.NWListenerStateReady:
			select {
			case ready <- nil:
			default:
			}
		case applenetwork.NWListenerStateFailed, applenetwork.NWListenerStateCancelled:
			err := fmt.Errorf("nw listener %s", state)
			if !nwErr.IsZero() {
				err = fmt.Errorf("%s: %w", err, nwErr)
			}
			select {
			case ready <- err:
			default:
			}
		}
	})
	applenetwork.NWListenerSetNewConnectionHandler(listener, func(obj objectivec.Object) {
		c.accept(applenetwork.NWConnectionFromID(obj.ID))
	})
	applenetwork.NWListenerStart(listener)

	select {
	case err := <-ready:
		if err != nil {
			_ = c.Close()
			return nil, err
		}
	case <-ctx.Done():
		_ = c.Close()
		return nil, ctx.Err()
	case <-time.After(5 * time.Second):
		_ = c.Close()
		return nil, errors.New("timed out waiting for nw listener readiness")
	}
	if got := applenetwork.NWListenerGetPort(listener); got != 0 {
		c.local.Port = int(got)
	}
	if c.local.Port == 0 {
		_ = c.Close()
		return nil, errors.New("nw listener returned port 0")
	}
	return c, nil
}

func newNWParameters(config Config, local *net.UDPAddr) (params applenetwork.NWParameters, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("configure nw parameters: %v", r)
		}
	}()
	params = applenetwork.NWParametersCreate()
	if params.ID == 0 {
		return applenetwork.NWParameters{}, errors.New("nw_parameters_create returned nil")
	}
	stack := applenetwork.NWParametersCopyDefaultProtocolStack(params)
	if stack.ID == 0 {
		return applenetwork.NWParameters{}, errors.New("nw_parameters_copy_default_protocol_stack returned nil")
	}
	udpOptions := applenetwork.NWUDPCreateOptions()
	if udpOptions.ID == 0 {
		return applenetwork.NWParameters{}, errors.New("nw_udp_create_options returned nil")
	}
	applenetwork.NWProtocolStackClearApplicationProtocols(stack)
	applenetwork.NWProtocolStackSetTransportProtocol(stack, udpOptions)
	applenetwork.NWParametersSetIncludePeerToPeer(params, config.IncludePeerToPeer)
	if config.SetRequiredInterface {
		applenetwork.NWParametersSetRequiredInterfaceType(params, config.RequiredInterfaceType)
	}
	applenetwork.NWParametersSetReuseLocalAddress(params, config.ReuseLocalAddress)
	if local != nil {
		applenetwork.NWParametersSetLocalEndpoint(params, nwEndpointForUDPAddr(local))
	}
	if config.RequireInterface {
		if err := requireNWInterface(params, config.InterfaceName); err != nil {
			return applenetwork.NWParameters{}, err
		}
	}
	return params, nil
}

func requireNWInterface(params applenetwork.NWParameters, name string) error {
	privateIface := newPrivateNWInterfaceWithName(name)
	if privateIface.ID == 0 {
		return fmt.Errorf("private NWInterface(%s) returned nil", name)
	}
	ciface := privateGetObject(privateIface, "cInterface")
	if ciface.ID == 0 {
		return fmt.Errorf("private NWInterface(%s).cInterface returned nil", name)
	}
	applenetwork.NWParametersRequireInterface(params, applenetwork.NWInterfaceFromID(ciface.ID))
	return nil
}

func newPrivateNWInterfaceWithName(name string) objectivec.Object {
	instance := objc.Send[objc.ID](objc.ID(objc.GetClass("NWInterface")), objc.Sel("alloc"))
	return objectivec.ObjectFromID(objc.Send[objc.ID](instance, objc.Sel("initWithInterfaceName:"), objectivec.ObjectFromID(objc.String(name))))
}

func privateGetObject(obj objectivec.Object, selector string) objectivec.Object {
	return objectivec.ObjectFromID(objc.Send[objc.ID](obj.ID, objc.Sel(selector)))
}

func (c *nwPacketConn) ReadFrom(b []byte) (int, net.Addr, error) {
	for {
		deadline, changed := c.readDeadlineState()
		var timer *time.Timer
		var timeout <-chan time.Time
		if !deadline.IsZero() {
			wait := time.Until(deadline)
			if wait <= 0 {
				return 0, nil, nwTimeoutError{op: "read"}
			}
			timer = time.NewTimer(wait)
			timeout = timer.C
		}

		select {
		case pkt := <-c.packets:
			if timer != nil {
				timer.Stop()
			}
			n := copy(b, pkt.data)
			return n, copyUDPAddr(pkt.addr), nil
		case <-c.closed:
			if timer != nil {
				timer.Stop()
			}
			return 0, nil, net.ErrClosed
		case <-changed:
			if timer != nil {
				timer.Stop()
			}
			continue
		case <-timeout:
			return 0, nil, nwTimeoutError{op: "read"}
		}
	}
}

func (c *nwPacketConn) WriteTo(b []byte, addr net.Addr) (int, error) {
	select {
	case <-c.closed:
		return 0, net.ErrClosed
	default:
	}
	deadline, _ := c.writeDeadlineState()
	if !deadline.IsZero() && time.Until(deadline) <= 0 {
		return 0, nwTimeoutError{op: "write"}
	}
	retries := c.connectRetries()
	for attempt := 0; ; attempt++ {
		peer, err := c.peerConn(addr)
		if err != nil {
			return 0, err
		}
		if err := peer.waitReady(deadline, c.connectTimeout()); err != nil {
			if attempt < retries && peer.outbound && isTimeoutError(err) {
				c.tracef("outbound %s readiness timeout; retrying connection", peer.addr)
				c.dropPeer(peer)
				continue
			}
			return 0, err
		}
		return c.writeReadyPeer(b, peer, deadline)
	}
}

func (c *nwPacketConn) writeReadyPeer(b []byte, peer *nwPeerConn, deadline time.Time) (int, error) {
	data := dispatch.DataCreate(b)
	context := applenetwork.NWContentContextCreate(fmt.Sprintf("tmc-apple-nwpacket-%d", c.sendSeq.Add(1)))
	done := make(chan error, 1)
	applenetwork.NWConnectionSend(peer.conn, data, context, true, func(nwErr applenetwork.NWError) {
		c.tracef("send %s bytes=%d err=%s", peer.addr, len(b), nwErrorString(nwErr))
		data.Release()
		if context.ID != 0 {
			context.Release()
		}
		if nwErr.IsZero() {
			done <- nil
			return
		}
		done <- nwErr
	})
	wait := c.connectTimeout()
	if !deadline.IsZero() {
		wait = time.Until(deadline)
	}
	timer := time.NewTimer(wait)
	defer timer.Stop()
	select {
	case err := <-done:
		if err != nil {
			return 0, fmt.Errorf("nw send %s: %w", peer.addr, err)
		}
		return len(b), nil
	case <-timer.C:
		return 0, nwTimeoutError{op: "write"}
	case <-c.closed:
		return 0, net.ErrClosed
	}
}

func (c *nwPacketConn) Close() error {
	c.once.Do(func() {
		close(c.closed)
		applenetwork.NWListenerCancel(c.listener)
		c.mu.Lock()
		defer c.mu.Unlock()
		for _, peer := range c.conns {
			applenetwork.NWConnectionCancel(peer.conn)
		}
		c.conns = nil
	})
	return nil
}

func (c *nwPacketConn) LocalAddr() net.Addr {
	c.mu.Lock()
	defer c.mu.Unlock()
	return copyUDPAddr(c.local)
}

func (c *nwPacketConn) SetDeadline(t time.Time) error {
	if err := c.SetReadDeadline(t); err != nil {
		return err
	}
	return c.SetWriteDeadline(t)
}

func (c *nwPacketConn) SetReadDeadline(t time.Time) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.readDeadline = t
	close(c.readDeadlineChanged)
	c.readDeadlineChanged = make(chan struct{})
	return nil
}

func (c *nwPacketConn) SetWriteDeadline(t time.Time) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.writeDeadline = t
	close(c.writeDeadlineChanged)
	c.writeDeadlineChanged = make(chan struct{})
	return nil
}

func (c *nwPacketConn) accept(conn applenetwork.NWConnection) {
	addr := c.connectionEndpoint(conn)
	peer := &nwPeerConn{conn: conn, addr: addr, ready: make(chan error, 1)}
	c.mu.Lock()
	c.conns[addr.String()] = peer
	c.mu.Unlock()

	applenetwork.NWConnectionSetQueue(conn, c.queue)
	applenetwork.NWConnectionSetStateChangedHandler(conn, func(state applenetwork.NWConnectionState, nwErr applenetwork.NWError) {
		c.tracef("accepted %s state=%s err=%s", addr, state, nwErrorString(nwErr))
		switch state {
		case applenetwork.NWConnectionStateReady:
			peer.markReady(nil)
		case applenetwork.NWConnectionStateFailed, applenetwork.NWConnectionStateCancelled:
			err := fmt.Errorf("nw connection %s %s", addr, state)
			if !nwErr.IsZero() {
				err = fmt.Errorf("%s: %w", err, nwErr)
			}
			peer.markReady(err)
		}
	})
	c.receive(conn, addr)
	applenetwork.NWConnectionStart(conn)
}

func (c *nwPacketConn) peerConn(addr net.Addr) (*nwPeerConn, error) {
	udpAddr, err := netAddrToUDP(addr)
	if err != nil {
		return nil, err
	}
	if udpAddr.Zone == "" && udpAddr.IP.To4() == nil && udpAddr.IP.IsLinkLocalUnicast() {
		udpAddr.Zone = c.config.InterfaceName
	}
	key := udpAddr.String()
	c.mu.Lock()
	if peer := c.conns[key]; peer != nil {
		c.mu.Unlock()
		return peer, nil
	}
	c.mu.Unlock()

	params, err := newNWParameters(c.config, c.local)
	if err != nil {
		return nil, err
	}
	endpoint := nwEndpointForUDPAddr(udpAddr)
	conn := applenetwork.NWConnectionCreate(endpoint, params)
	if conn.ID == 0 {
		return nil, fmt.Errorf("nw_connection_create %s returned nil", udpAddr)
	}
	peer := &nwPeerConn{conn: conn, addr: udpAddr, ready: make(chan error, 1), outbound: true}
	applenetwork.NWConnectionSetQueue(conn, c.queue)
	applenetwork.NWConnectionSetStateChangedHandler(conn, func(state applenetwork.NWConnectionState, nwErr applenetwork.NWError) {
		c.tracef("outbound %s state=%s err=%s", udpAddr, state, nwErrorString(nwErr))
		if state == applenetwork.NWConnectionStateReady && c.config.Tracef != nil {
			c.tracef("outbound %s path=%s", udpAddr, c.connectionPath(conn))
		}
		switch state {
		case applenetwork.NWConnectionStateReady:
			peer.markReady(nil)
		case applenetwork.NWConnectionStateFailed, applenetwork.NWConnectionStateCancelled:
			err := fmt.Errorf("nw connection %s %s", udpAddr, state)
			if !nwErr.IsZero() {
				err = fmt.Errorf("%s: %w", err, nwErr)
			}
			peer.markReady(err)
		}
	})
	c.receive(conn, udpAddr)
	applenetwork.NWConnectionStart(conn)

	c.mu.Lock()
	if existing := c.conns[key]; existing != nil {
		c.mu.Unlock()
		applenetwork.NWConnectionCancel(conn)
		return existing, nil
	}
	c.conns[key] = peer
	c.mu.Unlock()
	return peer, nil
}

func (c *nwPacketConn) dropPeer(peer *nwPeerConn) {
	if peer == nil {
		return
	}
	key := peer.addr.String()
	c.mu.Lock()
	if c.conns[key] == peer {
		delete(c.conns, key)
	}
	c.mu.Unlock()
	applenetwork.NWConnectionCancel(peer.conn)
}

func (c *nwPacketConn) connectTimeout() time.Duration {
	if c.config.ConnectTimeout > 0 {
		return c.config.ConnectTimeout
	}
	return 5 * time.Second
}

func (c *nwPacketConn) connectRetries() int {
	if c.config.ConnectRetries <= 0 {
		return 0
	}
	return c.config.ConnectRetries
}

func (p *nwPeerConn) markReady(err error) {
	p.once.Do(func() {
		p.ready <- err
		close(p.ready)
	})
}

func (p *nwPeerConn) waitReady(deadline time.Time, connectTimeout time.Duration) error {
	wait := connectTimeout
	if !deadline.IsZero() {
		until := time.Until(deadline)
		if until <= 0 {
			return nwTimeoutError{op: "write"}
		}
		if until < wait {
			wait = until
		}
	}
	timer := time.NewTimer(wait)
	defer timer.Stop()

	select {
	case err := <-p.ready:
		return err
	case <-timer.C:
		return fmt.Errorf("nw connection %s readiness: %w", p.addr, nwTimeoutError{op: "write"})
	}
}

func isTimeoutError(err error) bool {
	var netErr net.Error
	return errors.As(err, &netErr) && netErr.Timeout()
}

func (c *nwPacketConn) receive(conn applenetwork.NWConnection, addr *net.UDPAddr) {
	select {
	case <-c.closed:
		return
	default:
	}
	applenetwork.NWConnectionReceiveMessage(conn, func(content objectivec.Object, _ objectivec.Object, _ bool, nwErr applenetwork.NWError) {
		if !nwErr.IsZero() {
			c.enqueueErrorPacket(addr, nwErr)
			return
		}
		if content.ID != 0 {
			data := dispatch.DataFromHandle(uintptr(content.ID))
			c.enqueue(nwPacket{data: data.Bytes(), addr: addr})
		}
		c.receive(conn, addr)
	})
}

func (c *nwPacketConn) enqueue(pkt nwPacket) {
	select {
	case <-c.closed:
	case c.packets <- pkt:
	default:
	}
}

func (c *nwPacketConn) enqueueErrorPacket(addr *net.UDPAddr, err error) {
	_ = addr
	_ = err
}

func (c *nwPacketConn) connectionEndpoint(conn applenetwork.NWConnection) *net.UDPAddr {
	endpoint := applenetwork.NWConnectionCopyEndpoint(conn)
	if endpoint.ID == 0 {
		return &net.UDPAddr{}
	}
	addr, err := nwEndpointToUDPAddr(endpoint, c.config.InterfaceName)
	if err != nil {
		return &net.UDPAddr{}
	}
	return addr
}

func (c *nwPacketConn) connectionPath(conn applenetwork.NWConnection) string {
	path := applenetwork.NWConnectionCopyCurrentPath(conn)
	info, err := pathFromNWPath(path)
	if err != nil {
		return "none"
	}
	return info.String()
}

func (c *nwPacketConn) readDeadlineState() (time.Time, <-chan struct{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.readDeadline, c.readDeadlineChanged
}

func (c *nwPacketConn) writeDeadlineState() (time.Time, <-chan struct{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.writeDeadline, c.writeDeadlineChanged
}

func nwEndpointForUDPAddr(addr *net.UDPAddr) applenetwork.NWEndpoint {
	return applenetwork.NWEndpointCreateHost(nwEndpointHost(addr), strconv.Itoa(addr.Port))
}

func nwEndpointToUDPAddr(endpoint applenetwork.NWEndpoint, defaultZone string) (*net.UDPAddr, error) {
	host := objc.GoString(applenetwork.NWEndpointGetHostname(endpoint))
	port := int(applenetwork.NWEndpointGetPort(endpoint))
	host = strings.TrimPrefix(strings.TrimSuffix(host, "]"), "[")
	zone := ""
	if i := strings.LastIndex(host, "%"); i >= 0 {
		zone = host[i+1:]
		host = host[:i]
	}
	ip := net.ParseIP(host)
	if ip == nil {
		return nil, fmt.Errorf("nw endpoint host %q is not an IP address", host)
	}
	if zone == "" && ip.To4() == nil && ip.IsLinkLocalUnicast() {
		zone = defaultZone
	}
	return &net.UDPAddr{IP: ip, Port: port, Zone: zone}, nil
}

func nwEndpointHost(addr *net.UDPAddr) string {
	host := addr.IP.String()
	if addr.Zone != "" {
		host += "%" + addr.Zone
	}
	return host
}

func netAddrToUDP(addr net.Addr) (*net.UDPAddr, error) {
	switch a := addr.(type) {
	case *net.UDPAddr:
		return copyUDPAddr(a), nil
	default:
		udpAddr, err := net.ResolveUDPAddr("udp", addr.String())
		if err != nil {
			return nil, fmt.Errorf("resolve UDP addr %s: %w", addr, err)
		}
		return udpAddr, nil
	}
}

func copyUDPAddr(addr *net.UDPAddr) *net.UDPAddr {
	if addr == nil {
		return nil
	}
	cp := *addr
	cp.IP = append(net.IP(nil), addr.IP...)
	return &cp
}

type nwTimeoutError struct {
	op string
}

func (e nwTimeoutError) Error() string {
	return e.op + " i/o timeout"
}

func (e nwTimeoutError) Timeout() bool {
	return true
}

func (e nwTimeoutError) Temporary() bool {
	return true
}

func nwErrorString(err applenetwork.NWError) string {
	if err.IsZero() {
		return "-"
	}
	return fmt.Sprintf("%s domain=%s code=%d", err.Error(), err.DomainString(), err.Code())
}

func (c *nwPacketConn) tracef(format string, args ...any) {
	if c.config.Tracef == nil {
		return
	}
	c.config.Tracef(format, args...)
}
