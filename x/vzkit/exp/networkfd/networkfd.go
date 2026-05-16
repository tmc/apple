package networkfd

import (
	"fmt"
	"net"
	"os"
	"syscall"

	"github.com/tmc/apple/foundation"
	vz "github.com/tmc/apple/virtualization"
)

const DefaultMTU = 1500

// Pair is a connected datagram socket pair for a VZ file-handle network device.
type Pair struct {
	HostFD  int
	GuestFD int
}

// NewSocketPair returns a connected AF_UNIX datagram socket pair.
func NewSocketPair(mtu int) (Pair, error) {
	fds, err := syscall.Socketpair(syscall.AF_UNIX, syscall.SOCK_DGRAM, 0)
	if err != nil {
		return Pair{}, fmt.Errorf("socketpair: %w", err)
	}
	p := Pair{HostFD: fds[0], GuestFD: fds[1]}
	if err := ConfigureDatagramSocketBuffers(p.HostFD, mtu); err != nil {
		p.Close()
		return Pair{}, err
	}
	if err := ConfigureDatagramSocketBuffers(p.GuestFD, mtu); err != nil {
		p.Close()
		return Pair{}, err
	}
	return p, nil
}

// Close closes both descriptors in p.
func (p Pair) Close() {
	if p.HostFD >= 0 {
		_ = syscall.Close(p.HostFD)
	}
	if p.GuestFD >= 0 {
		_ = syscall.Close(p.GuestFD)
	}
}

// ConfigureDatagramSocketBuffers sets socket buffers sized for mtu.
func ConfigureDatagramSocketBuffers(fd int, mtu int) error {
	if mtu <= 0 {
		mtu = DefaultMTU
	}
	snd := mtu * 4
	rcv := mtu * 8
	if snd < 65536 {
		snd = 65536
	}
	if rcv < snd*2 {
		rcv = snd * 2
	}
	if err := syscall.SetsockoptInt(fd, syscall.SOL_SOCKET, syscall.SO_SNDBUF, snd); err != nil {
		return fmt.Errorf("set SO_SNDBUF: %w", err)
	}
	if err := syscall.SetsockoptInt(fd, syscall.SOL_SOCKET, syscall.SO_RCVBUF, rcv); err != nil {
		return fmt.Errorf("set SO_RCVBUF: %w", err)
	}
	return nil
}

// NewFileHandleFromFD wraps fd in an NSFileHandle that closes on dealloc.
func NewFileHandleFromFD(fd int) foundation.NSFileHandle {
	handle := foundation.NewFileHandleWithFileDescriptorCloseOnDealloc(fd, true)
	handle.Retain()
	return handle
}

// NewAttachment creates a VZ file-handle network attachment from guestFD.
func NewAttachment(guestFD int, mtu int) (foundation.NSFileHandle, vz.VZFileHandleNetworkDeviceAttachment, error) {
	handle := NewFileHandleFromFD(guestFD)
	attachment := vz.NewFileHandleNetworkDeviceAttachmentWithFileHandle(handle)
	if attachment.ID == 0 {
		_, _ = handle.CloseAndReturnError()
		return handle, attachment, fmt.Errorf("create filehandle network attachment")
	}
	attachment.Retain()
	if mtu <= 0 {
		mtu = DefaultMTU
	}
	attachment.SetMaximumTransmissionUnit(mtu)
	return handle, attachment, nil
}

// NewDevice creates a Virtio network device for attachment.
func NewDevice(attachment vz.VZFileHandleNetworkDeviceAttachment) (vz.VZVirtioNetworkDeviceConfiguration, error) {
	if attachment.ID == 0 {
		return vz.VZVirtioNetworkDeviceConfiguration{}, fmt.Errorf("network attachment required")
	}
	device := vz.NewVZVirtioNetworkDeviceConfiguration()
	if device.ID == 0 {
		return device, fmt.Errorf("create virtio network device configuration")
	}
	device.SetAttachment(&attachment.VZNetworkDeviceAttachment)
	device.Retain()
	return device, nil
}

// HostConn wraps hostFD as a net.Conn. The returned file is retained so callers
// can close both the connection and file explicitly.
func HostConn(hostFD int) (*os.File, net.Conn, error) {
	file := os.NewFile(uintptr(hostFD), "networkfd-host")
	if file == nil {
		_ = syscall.Close(hostFD)
		return nil, nil, fmt.Errorf("create host file")
	}
	conn, err := net.FileConn(file)
	if err != nil {
		file.Close()
		return nil, nil, fmt.Errorf("wrap host file: %w", err)
	}
	return file, conn, nil
}
