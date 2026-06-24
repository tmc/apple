//go:build darwin

package main

import (
	"errors"
	"fmt"
	"io/fs"
	"syscall"
	"testing"

	"github.com/hugelgupf/p9/linux"
)

// TestBackendErrorLinuxErrno covers the gap the in-memory smoke backend
// masks: the real 9P2000.L backend returns linux.Errno, which is not a
// syscall.Errno and does not match the io/fs sentinels, so without
// translation every such error collapses to EIO at the bridge. backendError
// must surface the Darwin errno through errors.As.
func TestBackendErrorLinuxErrno(t *testing.T) {
	tests := []struct {
		name string
		in   error
		want syscall.Errno
	}{
		{"EACCES", linux.EACCES, syscall.EACCES},
		{"EEXIST", linux.EEXIST, syscall.EEXIST},
		{"ENOSPC", linux.ENOSPC, syscall.ENOSPC},
		{"ENOTEMPTY", linux.ENOTEMPTY, syscall.ENOTEMPTY},
		{"EISDIR", linux.EISDIR, syscall.EISDIR},
		{"ENOTDIR", linux.ENOTDIR, syscall.ENOTDIR},
		{"EROFS", linux.EROFS, syscall.EROFS},
		{"ENODATA maps to Darwin ENOATTR", linux.ENODATA, errnoENOATTR},
		{"ENOTSUP", linux.ENOTSUP, syscall.ENOTSUP},
		{"wrapped EACCES", fmt.Errorf("write %s: %w", "f", linux.EACCES), syscall.EACCES},
		{"unknown linux errno falls to EIO", linux.Errno(0xfff), syscall.EIO},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var errno syscall.Errno
			if !errors.As(backendError(tt.in), &errno) {
				t.Fatalf("backendError(%v): errors.As(&syscall.Errno) = false, want errno %d", tt.in, tt.want)
			}
			if errno != tt.want {
				t.Fatalf("backendError(%v) errno = %d, want %d", tt.in, errno, tt.want)
			}
		})
	}
}

// TestBackendErrorClassic9P covers the classic 9P2000 backend, whose errors
// are plain strings from the server: well-known messages must still map to
// errnos rather than collapsing to EIO.
func TestBackendErrorClassic9P(t *testing.T) {
	tests := []struct {
		name string
		in   error
		want syscall.Errno
	}{
		{"not found", errors.New("file 'x' not found"), syscall.ENOENT},
		{"permission denied", errors.New("open: permission denied"), syscall.EACCES},
		{"already exists", errors.New("create: file already exists"), syscall.EEXIST},
		{"directory not empty", errors.New("remove: directory not empty"), syscall.ENOTEMPTY},
		{"unknown string falls through", errors.New("some transport glitch"), 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var errno syscall.Errno
			ok := errors.As(backendError(tt.in), &errno)
			if tt.want == 0 {
				if ok {
					t.Fatalf("backendError(%v): mapped to errno %d, want no errno (bridge defaults to EIO)", tt.in, errno)
				}
				return
			}
			if !ok || errno != tt.want {
				t.Fatalf("backendError(%v) errno = %d (ok=%v), want %d", tt.in, errno, ok, tt.want)
			}
		})
	}
}

// TestBackendErrorPassthrough verifies errors already carrying the bridge's
// expected types are returned unchanged (the bridge maps them itself).
func TestBackendErrorPassthrough(t *testing.T) {
	if got := backendError(nil); got != nil {
		t.Fatalf("backendError(nil) = %v, want nil", got)
	}
	se := syscall.ENOENT
	if got := backendError(se); !errors.Is(got, syscall.ENOENT) {
		t.Fatalf("backendError(syscall.ENOENT) lost its errno: %v", got)
	}
	if got := backendError(fs.ErrInvalid); !errors.Is(got, fs.ErrInvalid) {
		t.Fatalf("backendError(fs.ErrInvalid) lost its sentinel: %v", got)
	}
}
