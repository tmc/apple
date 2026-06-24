package fskitbridge

import (
	"errors"
	"fmt"
	"io/fs"
	"syscall"
	"testing"
)

func TestErrnoFor(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want syscall.Errno
	}{
		{"errno", syscall.ENOSPC, syscall.ENOSPC},
		{"wrapped errno", fmt.Errorf("write: %w", syscall.EDQUOT), syscall.EDQUOT},
		{"not exist", fs.ErrNotExist, syscall.ENOENT},
		{"exist", fs.ErrExist, syscall.EEXIST},
		{"permission", fs.ErrPermission, syscall.EACCES},
		{"invalid", fs.ErrInvalid, syscall.EINVAL},
		{"unsupported", errors.ErrUnsupported, syscall.ENOTSUP},
		{"unknown", errors.New("transport broke"), syscall.EIO},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := errnoFor(tt.err); got != tt.want {
				t.Fatalf("errnoFor(%v) = %v, want %v", tt.err, got, tt.want)
			}
		})
	}
}

func TestXattrErrnoFor(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want syscall.Errno
	}{
		{"not exist", fs.ErrNotExist, ENOATTR},
		{"wrapped not exist", fmt.Errorf("attr: %w", fs.ErrNotExist), ENOATTR},
		{"errno wins", fmt.Errorf("%w: %w", syscall.ENOENT, fs.ErrNotExist), syscall.ENOENT},
		{"unknown", errors.New("transport broke"), syscall.EIO},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := xattrErrnoFor(tt.err); got != tt.want {
				t.Fatalf("xattrErrnoFor(%v) = %v, want %v", tt.err, got, tt.want)
			}
		})
	}
}

func TestNewServerMissingFileSystem(t *testing.T) {
	if _, err := NewServer(ServerConfig{FileSystemName: "F", VolumeName: "V", ItemName: "I"}); err == nil {
		t.Fatal("NewServer() error = nil, want error")
	}
}
