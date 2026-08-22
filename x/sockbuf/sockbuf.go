// Package sockbuf reports and sets socket buffer sizes on Darwin, making the
// kernel's silent clamping visible to the caller.
//
// macOS clamps SO_SNDBUF and SO_RCVBUF requests to the kern.ipc.maxsockbuf
// sysctl ceiling without returning an error: a request for 8 MiB may install
// 2 MiB and report success. Code that sizes buffers with setsockopt alone
// cannot tell that it was truncated.
//
// The functions here read the value back after setting it and return what the
// kernel actually installed. A returned size smaller than the requested size
// is not an error; it is the kernel's answer, and comparing the two is the
// only reliable way to detect the clamp.
//
//	actual, err := sockbuf.SetRecvBuffer(fd, 8<<20)
//	if err != nil {
//		return err
//	}
//	if actual < 8<<20 {
//		log.Printf("kernel clamped receive buffer to %d bytes", actual)
//	}
//
// This package uses only public APIs and works on any Darwin version.
package sockbuf

import (
	"fmt"
	"syscall"
)

// MaxSockBuf reports the kern.ipc.maxsockbuf sysctl, the ceiling the kernel
// applies to SO_SNDBUF and SO_RCVBUF requests.
//
// A request larger than this value is silently truncated rather than
// rejected, so callers that need to distinguish "asked for too much" from
// "got what I asked for" should consult it before sizing.
func MaxSockBuf() (int, error) {
	v, err := syscall.SysctlUint32("kern.ipc.maxsockbuf")
	if err != nil {
		return 0, fmt.Errorf("sysctl kern.ipc.maxsockbuf: %w", err)
	}
	return int(v), nil
}

// SetRecvBuffer requests bytes for the socket's receive buffer and reports
// the size the kernel installed, which may be smaller than requested.
func SetRecvBuffer(fd, bytes int) (actual int, err error) {
	return setAndReport(fd, syscall.SO_RCVBUF, bytes)
}

// SetSendBuffer requests bytes for the socket's send buffer and reports the
// size the kernel installed, which may be smaller than requested.
func SetSendBuffer(fd, bytes int) (actual int, err error) {
	return setAndReport(fd, syscall.SO_SNDBUF, bytes)
}

func setAndReport(fd, opt, bytes int) (int, error) {
	if err := syscall.SetsockoptInt(fd, syscall.SOL_SOCKET, opt, bytes); err != nil {
		return 0, fmt.Errorf("setsockopt: %w", err)
	}
	actual, err := syscall.GetsockoptInt(fd, syscall.SOL_SOCKET, opt)
	if err != nil {
		return 0, fmt.Errorf("getsockopt: %w", err)
	}
	return actual, nil
}
