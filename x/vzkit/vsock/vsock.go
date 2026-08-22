package vsock

import (
	"fmt"
	"net"
	"os"
	"sync"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/tmc/apple/objc"
	vz "github.com/tmc/apple/virtualization"
)

// Conn wraps a Virtio socket connection as a net.Conn.
type Conn struct {
	raw     net.Conn
	vzConn  vz.VZVirtioSocketConnection
	dstPort uint32
	srcPort uint32
	closed  atomic.Bool
}

// NewConn converts a Virtio socket connection into a net.Conn.
func NewConn(vzConn vz.VZVirtioSocketConnection) (*Conn, error) {
	fd := vzConn.FileDescriptor()
	if fd < 0 {
		return nil, fmt.Errorf("vsock connection closed (fd=%d)", fd)
	}
	if vzConn.ID != 0 {
		objc.Send[objc.ID](vzConn.ID, objc.Sel("retain"))
	}
	dupFd, err := syscall.Dup(int(fd))
	if err != nil {
		releaseConn(vzConn)
		return nil, fmt.Errorf("dup vsock fd: %w", err)
	}
	file := os.NewFile(uintptr(dupFd), "vsock")
	rawConn, err := net.FileConn(file)
	file.Close()
	if err != nil {
		releaseConn(vzConn)
		return nil, fmt.Errorf("file conn: %w", err)
	}
	return &Conn{
		raw:     rawConn,
		vzConn:  vzConn,
		dstPort: vzConn.DestinationPort(),
		srcPort: vzConn.SourcePort(),
	}, nil
}

func (c *Conn) Read(b []byte) (int, error)         { return c.raw.Read(b) }
func (c *Conn) Write(b []byte) (int, error)        { return c.raw.Write(b) }
func (c *Conn) LocalAddr() net.Addr                { return c.raw.LocalAddr() }
func (c *Conn) RemoteAddr() net.Addr               { return c.raw.RemoteAddr() }
func (c *Conn) SetDeadline(t time.Time) error      { return c.raw.SetDeadline(t) }
func (c *Conn) SetReadDeadline(t time.Time) error  { return c.raw.SetReadDeadline(t) }
func (c *Conn) SetWriteDeadline(t time.Time) error { return c.raw.SetWriteDeadline(t) }

func (c *Conn) Close() error {
	if c == nil || c.closed.Swap(true) {
		return nil
	}
	err := c.raw.Close()
	if c.vzConn.ID != 0 {
		c.vzConn.Close()
		releaseConn(c.vzConn)
	}
	return err
}

func releaseConn(conn vz.VZVirtioSocketConnection) {
	if conn.ID != 0 {
		objc.Send[objc.ID](conn.ID, objc.Sel("release"))
	}
}

// Listener accepts guest-initiated Virtio socket connections on a host port.
type Listener struct {
	Port uint32

	device   vz.VZVirtioSocketDevice
	listener vz.VZVirtioSocketListener
	delegate objc.ID
	dispatch func(func())
	conns    chan *Conn
	done     chan struct{}
	closed   atomic.Bool
}

var (
	listenerRegistry      = make(map[uintptr]*Listener)
	listenerRegistryMu    sync.RWMutex
	listenerDelegateClass objc.Class
	listenerDelegateOnce  sync.Once
)

const defaultListenBacklog = 16

// Listen installs a Virtio socket listener for port on device.
// If dispatch is non-nil, Virtualization.framework calls are run through it.
func Listen(device vz.VZVirtioSocketDevice, port uint32, dispatch func(func())) (*Listener, error) {
	l := &Listener{
		Port:     port,
		device:   device,
		dispatch: dispatch,
		conns:    make(chan *Conn, defaultListenBacklog),
		done:     make(chan struct{}),
	}

	var setupErr error
	l.run(func() {
		delegateClass := getListenerDelegateClass()
		if delegateClass == 0 {
			setupErr = fmt.Errorf("create vsock listener delegate class")
			return
		}

		delegate := objc.Send[objc.ID](objc.ID(delegateClass), objc.Sel("alloc"))
		delegate = objc.Send[objc.ID](delegate, objc.Sel("init"))
		if delegate == 0 {
			setupErr = fmt.Errorf("create vsock listener delegate")
			return
		}
		l.delegate = delegate

		listenerRegistryMu.Lock()
		listenerRegistry[uintptr(delegate)] = l
		listenerRegistryMu.Unlock()

		vzListener := vz.NewVZVirtioSocketListener()
		if vzListener.ID == 0 {
			listenerRegistryMu.Lock()
			delete(listenerRegistry, uintptr(delegate))
			listenerRegistryMu.Unlock()
			objc.Send[objc.ID](delegate, objc.Sel("release"))
			setupErr = fmt.Errorf("create vsock listener")
			return
		}
		objc.Send[objc.ID](vzListener.ID, objc.Sel("setDelegate:"), delegate)
		l.listener = vzListener
		device.SetSocketListenerForPort(&vzListener, port)
	})
	if setupErr != nil {
		return nil, setupErr
	}
	return l, nil
}

func (l *Listener) run(fn func()) {
	if l.dispatch != nil {
		l.dispatch(fn)
		return
	}
	fn()
}

// Accept waits for and returns the next guest connection.
func (l *Listener) Accept() (*Conn, error) {
	select {
	case conn := <-l.conns:
		return conn, nil
	case <-l.done:
		return nil, net.ErrClosed
	}
}

// Close removes the listener from the socket device.
func (l *Listener) Close() error {
	if l == nil || l.closed.Swap(true) {
		return nil
	}
	close(l.done)
	l.run(func() {
		if l.device.ID != 0 {
			l.device.RemoveSocketListenerForPort(l.Port)
		}

		listenerRegistryMu.Lock()
		delete(listenerRegistry, uintptr(l.delegate))
		listenerRegistryMu.Unlock()

		if l.delegate != 0 {
			objc.Send[objc.ID](l.delegate, objc.Sel("release"))
			l.delegate = 0
		}
	})
	return nil
}

func getListenerDelegateClass() objc.Class {
	listenerDelegateOnce.Do(func() {
		nsObjectClass := objc.GetClass("NSObject")
		protocol := objc.GetProtocol("VZVirtioSocketListenerDelegate")

		var protocols []*objc.Protocol
		if protocol != nil {
			protocols = []*objc.Protocol{protocol}
		}

		var err error
		listenerDelegateClass, err = objc.RegisterClass(
			"VZKitVirtioSocketListenerDelegate",
			nsObjectClass,
			protocols,
			nil,
			[]objc.MethodDef{{
				Cmd: objc.Sel("listener:shouldAcceptNewConnection:fromSocketDevice:"),
				Fn:  listenerShouldAcceptNewConnection,
			}},
		)
		if err != nil {
			listenerDelegateClass = 0
		}
	})
	return listenerDelegateClass
}

func listenerShouldAcceptNewConnection(self objc.ID, _cmd objc.SEL, _listener, connection, _socketDevice objc.ID) bool {
	listenerRegistryMu.RLock()
	l, ok := listenerRegistry[uintptr(self)]
	listenerRegistryMu.RUnlock()
	if !ok || l.closed.Load() {
		return false
	}

	vzConn := vz.VZVirtioSocketConnectionFromID(connection)
	conn, err := NewConn(vzConn)
	if err != nil {
		return false
	}

	select {
	case l.conns <- conn:
		return true
	case <-l.done:
		_ = conn.Close()
		return false
	default:
		_ = conn.Close()
		return false
	}
}

// Manager manages the first Virtio socket device on a running VM.
type Manager struct {
	device vz.VZVirtioSocketDevice
	mu     sync.Mutex

	// DispatchFunc runs Virtualization calls on the VM's queue.
	DispatchFunc func(fn func())
}

// NewManager wraps the first Virtio socket device from a running VM.
func NewManager(machine vz.VZVirtualMachine) (*Manager, error) {
	socketDevices := machine.SocketDevices()
	if len(socketDevices) == 0 {
		return nil, fmt.Errorf("no socket devices configured on VM")
	}
	device := vz.VZVirtioSocketDeviceFromID(socketDevices[0].GetID())
	if device.ID == 0 {
		return nil, fmt.Errorf("failed to get VirtioSocketDevice")
	}
	return &Manager{device: device}, nil
}

// Connect establishes a vsock connection to port.
func (m *Manager) Connect(port uint32) (net.Conn, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	type result struct {
		conn net.Conn
		err  error
	}
	ch := make(chan result, 1)

	connectFn := func() {
		m.device.ConnectToPortCompletionHandler(port, func(conn *vz.VZVirtioSocketConnection, err error) {
			if err != nil {
				ch <- result{err: fmt.Errorf("connect vsock port %d: %w", port, err)}
				return
			}
			if conn == nil {
				ch <- result{err: fmt.Errorf("connect vsock port %d: nil connection", port)}
				return
			}
			netConn, wrapErr := NewConn(*conn)
			ch <- result{conn: netConn, err: wrapErr}
		})
	}

	if m.DispatchFunc != nil {
		m.DispatchFunc(connectFn)
	} else {
		connectFn()
	}
	r := <-ch
	return r.conn, r.err
}

// Listen installs a listener on the managed socket device.
func (m *Manager) Listen(port uint32) (*Listener, error) {
	return Listen(m.device, port, m.DispatchFunc)
}
