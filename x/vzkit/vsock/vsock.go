package vsock

import (
	"fmt"
	"net"
	"os"
	"sync"
	"syscall"
	"time"

	vz "github.com/tmc/apple/virtualization"
)

// Conn wraps a Virtio socket connection as a net.Conn.
type Conn struct {
	raw     net.Conn
	vzConn  vz.VZVirtioSocketConnection
	dstPort uint32
	srcPort uint32
}

// NewConn converts a Virtio socket connection into a net.Conn.
func NewConn(vzConn vz.VZVirtioSocketConnection) (*Conn, error) {
	fd := vzConn.FileDescriptor()
	if fd < 0 {
		return nil, fmt.Errorf("vsock connection closed (fd=%d)", fd)
	}
	dupFd, err := syscall.Dup(fd)
	if err != nil {
		return nil, fmt.Errorf("dup vsock fd: %w", err)
	}
	file := os.NewFile(uintptr(dupFd), "vsock")
	rawConn, err := net.FileConn(file)
	file.Close()
	if err != nil {
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
func (c *Conn) Close() error                       { return c.raw.Close() }
func (c *Conn) LocalAddr() net.Addr                { return c.raw.LocalAddr() }
func (c *Conn) RemoteAddr() net.Addr               { return c.raw.RemoteAddr() }
func (c *Conn) SetDeadline(t time.Time) error      { return c.raw.SetDeadline(t) }
func (c *Conn) SetReadDeadline(t time.Time) error  { return c.raw.SetReadDeadline(t) }
func (c *Conn) SetWriteDeadline(t time.Time) error { return c.raw.SetWriteDeadline(t) }

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
