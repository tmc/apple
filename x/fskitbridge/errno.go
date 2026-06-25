package fskitbridge

import (
	"errors"
	"io/fs"
	"syscall"
)

// ENOATTR is Darwin's "attribute not found" errno, which the syscall package
// does not name. File systems mapping a backend's missing-attribute error
// report it for extended attribute operations.
const ENOATTR = syscall.Errno(93)

// errnoFor maps an operation error to the errno reported to FSKit.
func errnoFor(err error) syscall.Errno {
	var errno syscall.Errno
	switch {
	case errors.As(err, &errno):
		return errno
	case errors.Is(err, fs.ErrNotExist):
		return syscall.ENOENT
	case errors.Is(err, fs.ErrExist):
		return syscall.EEXIST
	case errors.Is(err, fs.ErrPermission):
		return syscall.EACCES
	case errors.Is(err, fs.ErrInvalid):
		return syscall.EINVAL
	case errors.Is(err, errors.ErrUnsupported):
		return syscall.ENOTSUP
	default:
		return syscall.EIO
	}
}

// xattrErrnoFor is errnoFor for extended attribute operations, where a
// missing attribute reports ENOATTR rather than ENOENT.
func xattrErrnoFor(err error) syscall.Errno {
	var errno syscall.Errno
	if errors.As(err, &errno) {
		return errno
	}
	if errors.Is(err, fs.ErrNotExist) {
		return ENOATTR
	}
	return errnoFor(err)
}
