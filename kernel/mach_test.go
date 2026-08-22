package kernel

import (
	"runtime"
	"testing"
	"unsafe"
)

func TestAbsoluteTimeLayout(t *testing.T) {
	if got := unsafe.Sizeof(AbsoluteTime(0)); got != 8 {
		t.Fatalf("sizeof(AbsoluteTime) = %d, want 8", got)
	}
}

func TestMachPortRoundTrip(t *testing.T) {
	self := Mach_task_self()
	if self == 0 {
		t.Fatal("MachTaskSelf returned 0")
	}

	var port uint32
	if kr := Mach_port_allocate(self, MachPortRightReceive, &port); kr != 0 {
		t.Fatalf("Mach_port_allocate: kr=%d", kr)
	}
	if port == 0 {
		t.Fatal("Mach_port_allocate returned port 0")
	}
	if kr := Mach_port_insert_right(self, port, port, MachMsgTypeMakeSend); kr != 0 {
		t.Fatalf("Mach_port_insert_right: kr=%d", kr)
	}

	// Send a message to ourselves, then receive it back.
	type msg struct {
		hdr MachMsgHdr
		pad [64]byte
	}
	send := msg{hdr: MachMsgHdr{
		Bits:       machMsgBits(MachMsgTypeCopySend, 0),
		Size:       uint32(unsafe.Sizeof(msg{})),
		RemotePort: port,
		ID:         4242,
	}}
	if kr := Mach_msg(unsafe.Pointer(&send.hdr), MachSendMsg|MachSendTimeout,
		uint32(unsafe.Sizeof(msg{})), 0, 0, 1000, 0); kr != 0 {
		t.Fatalf("MachMsg send: kr=%#x", kr)
	}

	// The receive buffer must hold the message plus its trailer.
	var recv struct {
		hdr MachMsgHdr
		pad [256]byte
	}
	if kr := Mach_msg(unsafe.Pointer(&recv.hdr), MachRcvMsg|MachRcvTimeout,
		0, uint32(unsafe.Sizeof(recv)), port, 1000, 0); kr != 0 {
		t.Fatalf("MachMsg recv: kr=%#x", kr)
	}
	if recv.hdr.ID != 4242 {
		t.Errorf("msgh_id = %d, want 4242", recv.hdr.ID)
	}
}

// TestKeventMachPort covers the EVFILT_MACHPORT path end to end. It also
// documents why this cannot be handed to the Go network poller: the wait has
// to happen on a thread this test owns.
func TestKeventMachPort(t *testing.T) {
	self := Mach_task_self()
	var port, pset uint32
	if kr := Mach_port_allocate(self, MachPortRightReceive, &port); kr != 0 {
		t.Fatalf("allocate receive: kr=%d", kr)
	}
	if kr := Mach_port_insert_right(self, port, port, MachMsgTypeMakeSend); kr != 0 {
		t.Fatalf("insert send right: kr=%d", kr)
	}
	if kr := Mach_port_allocate(self, MachPortRightPortSet, &pset); kr != 0 {
		t.Fatalf("allocate port set: kr=%d", kr)
	}
	if kr := Mach_port_insert_member(self, port, pset); kr != 0 {
		t.Fatalf("insert member: kr=%d", kr)
	}

	kq := Kqueue()
	if kq < 0 {
		t.Fatal("Kqueue failed")
	}

	buf := make([]byte, 1024)
	change := Kevent64_s{
		Ident:  uint64(pset),
		Filter: EVFILT_MACHPORT,
		Flags:  EV_ADD | EV_ENABLE,
		Fflags: MachRcvMsgFflag,
		Ext:    [2]uint64{uint64(uintptr(unsafe.Pointer(&buf[0]))), uint64(len(buf))},
	}
	if n := KeventWait(int(kq), []Kevent64_s{change}, nil, 0, &Timespec{}); n < 0 {
		t.Fatalf("Kevent register: n=%d", n)
	}
	runtime.KeepAlive(buf)

	// Nothing queued yet: an immediate wait must time out.
	events := make([]Kevent64_s, 1)
	if n := KeventWait(int(kq), nil, events, 0, &Timespec{}); n != 0 {
		t.Fatalf("expected 0 events before send, got %d (flags=%#x data=%d)",
			n, events[0].Flags, events[0].Data)
	}

	type msg struct {
		hdr MachMsgHdr
		pad [64]byte
	}
	send := msg{hdr: MachMsgHdr{
		Bits:       machMsgBits(MachMsgTypeCopySend, 0),
		Size:       uint32(unsafe.Sizeof(msg{})),
		RemotePort: port,
		ID:         99,
	}}
	if kr := Mach_msg(unsafe.Pointer(&send.hdr), MachSendMsg|MachSendTimeout,
		uint32(unsafe.Sizeof(msg{})), 0, 0, 1000, 0); kr != 0 {
		t.Fatalf("MachMsg send: kr=%#x", kr)
	}

	n := KeventWait(int(kq), nil, events, 0, &Timespec{Tv_sec: 5})
	if n != 1 {
		t.Fatalf("Kevent wait: n=%d, want 1", n)
	}
	if events[0].Flags&EV_ERROR != 0 {
		t.Fatalf("EV_ERROR set, errno=%d", events[0].Data)
	}
	if events[0].Ident != uint64(pset) {
		t.Errorf("ident = %#x, want port set %#x", events[0].Ident, pset)
	}
	runtime.KeepAlive(buf)
}

// machMsgBits builds msgh_bits from the remote and local type names, matching
// the MACH_MSGH_BITS macro.
func machMsgBits(remote, local uint32) uint32 {
	return remote | local<<8
}
