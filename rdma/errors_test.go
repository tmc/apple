package rdma

import (
	"syscall"
	"testing"
)

func TestErrnoTextDarwinNotSupported(t *testing.T) {
	if got, want := ErrnoText(int(syscall.ENOTSUP)), "errno 45 (ENOTSUP)"; got != want {
		t.Fatalf("ErrnoText(ENOTSUP) = %q, want %q", got, want)
	}
	if got, want := ErrnoName(int(syscall.ENOTSUP)), "ENOTSUP"; got != want {
		t.Fatalf("ErrnoName(ENOTSUP) = %q, want %q", got, want)
	}
}
