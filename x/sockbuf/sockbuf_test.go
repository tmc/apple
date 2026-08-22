package sockbuf

import (
	"syscall"
	"testing"
)

func TestMaxSockBuf(t *testing.T) {
	max, err := MaxSockBuf()
	if err != nil {
		t.Fatalf("MaxSockBuf: %v", err)
	}
	if max <= 0 {
		t.Fatalf("MaxSockBuf = %d, want > 0", max)
	}
	t.Logf("kern.ipc.maxsockbuf = %d", max)
}

func TestSetRecvBufferReportsClamp(t *testing.T) {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, 0)
	if err != nil {
		t.Fatalf("socket: %v", err)
	}
	defer syscall.Close(fd)

	max, err := MaxSockBuf()
	if err != nil {
		t.Fatalf("MaxSockBuf: %v", err)
	}

	// Ask for more than the ceiling: setsockopt fails with ENOBUFS on some
	// releases and silently clamps on others. Either way the caller must end
	// up knowing the installed size, so try the oversized request first and
	// fall back to the ceiling itself.
	actual, err := SetRecvBuffer(fd, max+1<<20)
	if err != nil {
		actual, err = SetRecvBuffer(fd, max)
		if err != nil {
			t.Fatalf("SetRecvBuffer(max): %v", err)
		}
	}
	if actual <= 0 {
		t.Fatalf("SetRecvBuffer reported %d, want > 0", actual)
	}
	if actual > max {
		t.Errorf("installed %d exceeds kern.ipc.maxsockbuf %d", actual, max)
	}
	t.Logf("requested > max, kernel installed %d (ceiling %d)", actual, max)

	// A modest request should be visible via read-back too.
	small, err := SetSendBuffer(fd, 64<<10)
	if err != nil {
		t.Fatalf("SetSendBuffer: %v", err)
	}
	if small < 64<<10 {
		t.Errorf("SetSendBuffer(64KiB) installed %d, want >= 64KiB", small)
	}
}
