package vzkit

import (
	"fmt"
	"net"
	"os"
	"sync"
	"syscall"
	"time"

	vz "github.com/tmc/apple/virtualization"
)

// VsockConn wraps a VZVirtioSocketConnection as a net.Conn.
type VsockConn struct {
	raw     net.Conn
	vzConn  vz.VZVirtioSocketConnection
	dstPort uint32
	srcPort uint32
}

// NewVsockConn converts a VZVirtioSocketConnection into a net.Conn.
// It extracts the file descriptor and wraps it with net.FileConn.
func NewVsockConn(vzConn vz.VZVirtioSocketConnection) (*VsockConn, error) {
	fd := vzConn.FileDescriptor()
	if fd < 0 {
		return nil, fmt.Errorf("vsock connection closed (fd=%d)", fd)
	}

	// Dup the fd so vzConn retains ownership of the original.
	dupFd, err := syscall.Dup(int(fd))
	if err != nil {
		return nil, fmt.Errorf("dup vsock fd: %w", err)
	}
	file := os.NewFile(uintptr(dupFd), "vsock")
	rawConn, err := net.FileConn(file)
	file.Close() // FileConn dups the fd
	if err != nil {
		return nil, fmt.Errorf("file conn: %w", err)
	}

	return &VsockConn{
		raw:     rawConn,
		vzConn:  vzConn,
		dstPort: vzConn.DestinationPort(),
		srcPort: vzConn.SourcePort(),
	}, nil
}

func (c *VsockConn) Read(b []byte) (int, error)         { return c.raw.Read(b) }
func (c *VsockConn) Write(b []byte) (int, error)        { return c.raw.Write(b) }
func (c *VsockConn) Close() error                       { return c.raw.Close() }
func (c *VsockConn) LocalAddr() net.Addr                { return c.raw.LocalAddr() }
func (c *VsockConn) RemoteAddr() net.Addr               { return c.raw.RemoteAddr() }
func (c *VsockConn) SetDeadline(t time.Time) error      { return c.raw.SetDeadline(t) }
func (c *VsockConn) SetReadDeadline(t time.Time) error  { return c.raw.SetReadDeadline(t) }
func (c *VsockConn) SetWriteDeadline(t time.Time) error { return c.raw.SetWriteDeadline(t) }

// VsockManager manages the VZVirtioSocketDevice for a running VM.
type VsockManager struct {
	device vz.VZVirtioSocketDevice
	mu     sync.Mutex
	// DispatchFunc, if set, dispatches the VZ API call onto the VM's queue.
	// Virtualization framework calls must run on the VM's dispatch queue;
	// calling from an arbitrary goroutine causes SIGTRAP crashes.
	DispatchFunc func(fn func())
}

// NewVsockManager wraps the first VZVirtioSocketDevice from a running VM.
func NewVsockManager(vm vz.VZVirtualMachine) (*VsockManager, error) {
	socketDevices := vm.SocketDevices()
	if len(socketDevices) == 0 {
		return nil, fmt.Errorf("no socket devices configured on VM")
	}

	device := vz.VZVirtioSocketDeviceFromID(socketDevices[0].GetID())
	if device.ID == 0 {
		return nil, fmt.Errorf("failed to get VirtioSocketDevice")
	}

	return &VsockManager{device: device}, nil
}

// Connect establishes a vsock connection to the guest on the given port.
func (m *VsockManager) Connect(port uint32) (net.Conn, error) {
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
			netConn, wrapErr := NewVsockConn(*conn)
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
