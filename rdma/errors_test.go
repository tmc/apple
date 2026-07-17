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
	if got, want := ErrnoText(102), "errno 102 (EOPNOTSUPP)"; got != want {
		t.Fatalf("ErrnoText(EOPNOTSUPP) = %q, want %q", got, want)
	}
	if got, want := ErrnoName(102), "EOPNOTSUPP"; got != want {
		t.Fatalf("ErrnoName(EOPNOTSUPP) = %q, want %q", got, want)
	}
	for _, errno := range []int{int(syscall.ENOTSUP), 102} {
		if !IsUnsupportedErrno(errno) {
			t.Fatalf("IsUnsupportedErrno(%d) = false, want true", errno)
		}
	}
	if IsUnsupportedErrno(22) {
		t.Fatal("IsUnsupportedErrno(EINVAL) = true, want false")
	}
}
