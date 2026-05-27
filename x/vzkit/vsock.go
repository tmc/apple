package vzkit

import (
	"net"

	vz "github.com/tmc/apple/virtualization"
	vsockx "github.com/tmc/apple/x/vzkit/vsock"
)

// VsockConn wraps a VZVirtioSocketConnection as a net.Conn.
type VsockConn = vsockx.Conn

// NewVsockConn converts a VZVirtioSocketConnection into a net.Conn.
func NewVsockConn(vzConn vz.VZVirtioSocketConnection) (*VsockConn, error) {
	return vsockx.NewConn(vzConn)
}

// VsockManager manages the VZVirtioSocketDevice for a running VM.
type VsockManager = vsockx.Manager

// NewVsockManager wraps the first VZVirtioSocketDevice from a running VM.
func NewVsockManager(vm vz.VZVirtualMachine) (*VsockManager, error) {
	return vsockx.NewManager(vm)
}

// VsockListener accepts guest-initiated Virtio socket connections.
type VsockListener = vsockx.Listener

// NewVsockListener installs a Virtio socket listener for port on device.
func NewVsockListener(device vz.VZVirtioSocketDevice, port uint32, dispatch func(func())) (*VsockListener, error) {
	return vsockx.Listen(device, port, dispatch)
}

var _ net.Conn = (*VsockConn)(nil)
