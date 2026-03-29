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

var _ net.Conn = (*VsockConn)(nil)
