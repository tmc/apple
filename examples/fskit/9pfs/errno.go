//go:build darwin

package main

import (
	"errors"
	"strings"
	"syscall"

	"github.com/hugelgupf/p9/linux"
)

// The 9P backends report failures in their own vocabularies: the 9P2000.L
// client returns linux.Errno (Linux-numbered, and not a syscall.Errno), and
// the classic 9P2000 client returns plain string errors from the server.
// x/fskitbridge maps an error to an errno only when it is, wraps, or matches
// a syscall.Errno or one of the io/fs sentinels; a linux.Errno or a bare
// string matches none of those and would collapse to EIO, losing EACCES,
// EEXIST, ENOSPC, ENOTEMPTY, and the rest.
//
// backendError translates a backend error into one carrying a Darwin
// syscall.Errno so the bridge reports the right errno. It also bridges the
// Linux/Darwin numeric skew (for example Linux ENODATA 61 is Darwin ENOATTR
// 93, Linux ENOTEMPTY 39 is Darwin 66), which a passthrough of the Linux
// value would get wrong.

// linuxToDarwin maps the linux.Errno values the 9P2000.L client can return to
// their Darwin syscall.Errno equivalents. Values that share a number across
// both systems are listed for clarity.
var linuxToDarwin = map[linux.Errno]syscall.Errno{
	linux.EPERM:        syscall.EPERM,
	linux.ENOENT:       syscall.ENOENT,
	linux.EIO:          syscall.EIO,
	linux.ENXIO:        syscall.ENXIO,
	linux.EBADF:        syscall.EBADF,
	linux.EAGAIN:       syscall.EAGAIN,
	linux.ENOMEM:       syscall.ENOMEM,
	linux.EACCES:       syscall.EACCES,
	linux.EFAULT:       syscall.EFAULT,
	linux.EBUSY:        syscall.EBUSY,
	linux.EEXIST:       syscall.EEXIST,
	linux.EXDEV:        syscall.EXDEV,
	linux.ENODEV:       syscall.ENODEV,
	linux.ENOTDIR:      syscall.ENOTDIR,
	linux.EISDIR:       syscall.EISDIR,
	linux.EINVAL:       syscall.EINVAL,
	linux.EFBIG:        syscall.EFBIG,
	linux.ENOSPC:       syscall.ENOSPC,
	linux.EROFS:        syscall.EROFS,
	linux.EMLINK:       syscall.EMLINK,
	linux.ENAMETOOLONG: syscall.ENAMETOOLONG,
	linux.ENOTEMPTY:    syscall.ENOTEMPTY, // Linux 39, Darwin 66
	linux.ELOOP:        syscall.ELOOP,
	linux.ENODATA:      errnoENOATTR, // Linux ENODATA 61, Darwin ENOATTR 93
	linux.EDQUOT:       syscall.EDQUOT,
	// Linux aliases ENOTSUP and EOPNOTSUPP to the same value (95); Darwin
	// splits them (ENOTSUP 45, EOPNOTSUPP 102). Report the general ENOTSUP.
	linux.ENOTSUP: syscall.ENOTSUP,
}

// errnoENOATTR is Darwin's ENOATTR, which the syscall package does not name.
const errnoENOATTR = syscall.Errno(93)

// stringToErrno maps the substrings of well-known classic 9P (9P2000) server
// error messages to errnos, since that client returns plain string errors.
var stringToErrno = []struct {
	substr string
	errno  syscall.Errno
}{
	{"not found", syscall.ENOENT},
	{"does not exist", syscall.ENOENT},
	{"already exists", syscall.EEXIST},
	{"file exists", syscall.EEXIST},
	{"permission denied", syscall.EACCES},
	{"access denied", syscall.EACCES},
	{"not a directory", syscall.ENOTDIR},
	{"is a directory", syscall.EISDIR},
	{"directory not empty", syscall.ENOTEMPTY},
	{"no space", syscall.ENOSPC},
	{"read-only", syscall.EROFS},
	{"read only", syscall.EROFS},
}

// errnoBackend wraps a backend so every method's error carries a Darwin
// syscall.Errno (see backendError). Wrapping the whole backend once keeps the
// translation in a single place rather than at each of the backend's call
// sites, and guarantees no operation is missed.
type errnoBackend struct{ backend }

func (b errnoBackend) Stat(name string) (nodeInfo, error) {
	info, err := b.backend.Stat(name)
	return info, backendError(err)
}

func (b errnoBackend) ReadDir(name string) ([]nodeInfo, error) {
	nodes, err := b.backend.ReadDir(name)
	return nodes, backendError(err)
}

func (b errnoBackend) ReadFile(name string) ([]byte, error) {
	data, err := b.backend.ReadFile(name)
	return data, backendError(err)
}

func (b errnoBackend) ReadFileAt(name string, offset int64, buf []byte) (int, error) {
	n, err := b.backend.ReadFileAt(name, offset, buf)
	return n, backendError(err)
}

func (b errnoBackend) WriteFile(name string, offset int64, data []byte) (int, error) {
	n, err := b.backend.WriteFile(name, offset, data)
	return n, backendError(err)
}

func (b errnoBackend) Create(name string, mode uint32, directory bool) (nodeInfo, error) {
	info, err := b.backend.Create(name, mode, directory)
	return info, backendError(err)
}

func (b errnoBackend) CreateSymlink(name, target string) (nodeInfo, error) {
	info, err := b.backend.CreateSymlink(name, target)
	return info, backendError(err)
}

func (b errnoBackend) CreateLink(oldName, newName string) (nodeInfo, error) {
	info, err := b.backend.CreateLink(oldName, newName)
	return info, backendError(err)
}

func (b errnoBackend) Readlink(name string) (string, error) {
	target, err := b.backend.Readlink(name)
	return target, backendError(err)
}

func (b errnoBackend) Remove(name string) error {
	return backendError(b.backend.Remove(name))
}

func (b errnoBackend) Rename(oldName, newName string) error {
	return backendError(b.backend.Rename(oldName, newName))
}

func (b errnoBackend) SetAttr(name string, attr setAttr) (nodeInfo, error) {
	info, err := b.backend.SetAttr(name, attr)
	return info, backendError(err)
}

func (b errnoBackend) GetXattr(name, attr string) ([]byte, error) {
	data, err := b.backend.GetXattr(name, attr)
	return data, backendError(err)
}

func (b errnoBackend) SetXattr(name, attr string, data []byte) error {
	return backendError(b.backend.SetXattr(name, attr, data))
}

func (b errnoBackend) ListXattr(name string) ([]string, error) {
	names, err := b.backend.ListXattr(name)
	return names, backendError(err)
}

func (b errnoBackend) RemoveXattr(name, attr string) error {
	return backendError(b.backend.RemoveXattr(name, attr))
}

func (b errnoBackend) Preallocate(name string, offset int64, length uint64) (uint64, error) {
	n, err := b.backend.Preallocate(name, offset, length)
	return n, backendError(err)
}

// backendError annotates err with a Darwin syscall.Errno so x/fskitbridge
// reports the right errno. It returns nil for nil. An error that already
// carries a syscall.Errno or an io/fs sentinel is returned unchanged, since
// the bridge already maps those.
func backendError(err error) error {
	if err == nil {
		return nil
	}
	var already syscall.Errno
	if errors.As(err, &already) {
		return err
	}
	if errno, ok := errnoForBackend(err); ok {
		return &backendErr{err: err, errno: errno}
	}
	return err
}

// errnoForBackend extracts a Darwin errno from a backend error, reporting
// whether one was recognized.
func errnoForBackend(err error) (syscall.Errno, bool) {
	var le linux.Errno
	if errors.As(err, &le) {
		if errno, ok := linuxToDarwin[le]; ok {
			return errno, true
		}
		return syscall.EIO, true
	}
	msg := strings.ToLower(err.Error())
	for _, m := range stringToErrno {
		if strings.Contains(msg, m.substr) {
			return m.errno, true
		}
	}
	return 0, false
}

// backendErr wraps a backend error with the Darwin errno it maps to. It
// unwraps to the original error and reports its errno through errors.As, so
// x/fskitbridge's errnoFor returns the errno.
type backendErr struct {
	err   error
	errno syscall.Errno
}

func (e *backendErr) Error() string { return e.err.Error() }
func (e *backendErr) Unwrap() error { return e.err }

// As lets errors.As(err, *syscall.Errno) recover the mapped errno even though
// the wrapped error is not itself a syscall.Errno.
func (e *backendErr) As(target any) bool {
	if p, ok := target.(*syscall.Errno); ok {
		*p = e.errno
		return true
	}
	return false
}
