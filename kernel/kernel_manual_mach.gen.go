// Code generated from Apple documentation for kernel. DO NOT EDIT.

package kernel

import (
	"runtime"
	"unsafe"

	"github.com/ebitengine/purego"
)

// Mach IPC, descriptors, return codes, and kqueue support.
//
// Layouts extracted directly from Apple SDK headers (SDK build 25F70 / macOS 26.5):
//   - <mach/message.h> (SDK build 25F70)
//   - <mach/kern_return.h> (SDK build 25F70)
//   - <mach/port.h> (SDK build 25F70)
//   - <sys/attr.h> (SDK build 25F70)
//   - <sys/stat.h> (SDK build 25F70)

// MachMsgHdr matches mach_msg_header_t from <mach/message.h> (SDK build 25F70).
// The generator emits mach_msg_header_t as an alias for unsafe.Pointer because Apple
// documents the type without a body.
type MachMsgHdr struct {
	Bits        uint32
	Size        uint32
	RemotePort  uint32
	LocalPort   uint32
	VoucherPort uint32
	ID          int32
}

// Mach_msg_body_t matches mach_msg_body_t from <mach/message.h> (SDK build 25F70).
type Mach_msg_body_t struct {
	Msgh_descriptor_count Mach_msg_size_t
}

// Mach_msg_type_descriptor_t matches mach_msg_type_descriptor_t from <mach/message.h> (SDK build 25F70).
type Mach_msg_type_descriptor_t struct {
	Pad1 Natural_t
	Pad2 Mach_msg_size_t
	Pad3 uint32
}

// Mach_msg_port_descriptor_t matches mach_msg_port_descriptor_t from <mach/message.h> (SDK build 25F70).
type Mach_msg_port_descriptor_t struct {
	Name Mach_port_t
	Pad1 Mach_msg_size_t
	Pad2 uint32
}

// Mach_msg_ool_descriptor_t matches mach_msg_ool_descriptor_t (LP64) from <mach/message.h> (SDK build 25F70).
type Mach_msg_ool_descriptor_t struct {
	Address unsafe.Pointer
	Flags   uint32
	Size    Mach_msg_size_t
}

// Mach_msg_trailer_t matches mach_msg_trailer_t from <mach/message.h> (SDK build 25F70).
type Mach_msg_trailer_t struct {
	Msgh_trailer_type Mach_msg_trailer_type_t
	Msgh_trailer_size Mach_msg_trailer_size_t
}

// Mach_msg_audit_trailer_t matches mach_msg_audit_trailer_t from <mach/message.h> (SDK build 25F70).
type Mach_msg_audit_trailer_t struct {
	Msgh_trailer_type Mach_msg_trailer_type_t
	Msgh_trailer_size Mach_msg_trailer_size_t
	Msgh_seqno        Mach_port_seqno_t
	Msgh_sender       [2]uint32
	Msgh_audit        [8]uint32
}

// Mach port right types, for Mach_port_allocate (from <mach/port.h>, SDK build 25F70).
const (
	MachPortRightSend     Mach_port_right_t = 0
	MachPortRightReceive  Mach_port_right_t = 1
	MachPortRightSendOnce Mach_port_right_t = 2
	MachPortRightPortSet  Mach_port_right_t = 3
	MachPortRightDeadName Mach_port_right_t = 4
)

// Message type names, for Mach_port_insert_right (from <mach/message.h>, SDK build 25F70).
const (
	MachMsgTypeMoveReceive  = 16
	MachMsgTypeMoveSend     = 17
	MachMsgTypeMoveSendOnce = 18
	MachMsgTypeCopySend     = 19
	MachMsgTypeMakeSend     = 20
	MachMsgTypeMakeSendOnce = 21
)

// Options for Mach_msg (from <mach/message.h>, SDK build 25F70).
const (
	MachSendMsg     Mach_msg_option_t = 0x00000001
	MachRcvMsg      Mach_msg_option_t = 0x00000002
	MachSendTimeout Mach_msg_option_t = 0x00000010
	MachRcvTimeout  Mach_msg_option_t = 0x00000100
)

// MachMsgTimeoutNone waits indefinitely.
const MachMsgTimeoutNone = 0

// EVFILT_MACHPORT selects the kevent filter that watches a Mach port or port set (from <sys/event.h>, SDK build 25F70).
const (
	EVFILT_MACHPORT = -8
	EVFILT_USER     = -10
)

// MachRcvMsgFflag asks kevent to receive the message inline (from <sys/event.h>, SDK build 25F70).
const MachRcvMsgFflag = 0x2

// kevent flags (from <sys/event.h>, SDK build 25F70).
const (
	EV_ADD    = 0x0001
	EV_DELETE = 0x0002
	EV_ENABLE = 0x0004
	EV_CLEAR  = 0x0020
	EV_ERROR  = 0x4000
)

// Kernel return code constants (from <mach/kern_return.h>, SDK build 25F70, signed 32-bit integer).
const (
	KERN_SUCCESS                  int32 = 0
	KERN_INVALID_ADDRESS          int32 = 1
	KERN_PROTECTION_FAILURE       int32 = 2
	KERN_NO_SPACE                 int32 = 3
	KERN_INVALID_ARGUMENT         int32 = 4
	KERN_FAILURE                  int32 = 5
	KERN_RESOURCE_SHORTAGE        int32 = 6
	KERN_NOT_RECEIVER             int32 = 7
	KERN_NO_ACCESS                int32 = 8
	KERN_MEMORY_FAILURE           int32 = 9
	KERN_MEMORY_ERROR             int32 = 10
	KERN_ALREADY_IN_SET           int32 = 11
	KERN_NOT_IN_SET               int32 = 12
	KERN_NAME_EXISTS              int32 = 13
	KERN_ABORTED                  int32 = 14
	KERN_INVALID_NAME             int32 = 15
	KERN_INVALID_TASK             int32 = 16
	KERN_INVALID_RIGHT            int32 = 17
	KERN_INVALID_VALUE            int32 = 18
	KERN_URESTRICTED              int32 = 19
	KERN_INVALID_CAPABILITY       int32 = 20
	KERN_RIGHT_EXISTS             int32 = 21
	KERN_INVALID_HOST             int32 = 22
	KERN_MEMORY_PRESENT           int32 = 23
	KERN_MEMORY_DATA_MOVED        int32 = 24
	KERN_MEMORY_RESTART_COPY      int32 = 25
	KERN_INVALID_PROCESSOR_SET    int32 = 26
	KERN_POLICY_LIMIT             int32 = 27
	KERN_INVALID_POLICY           int32 = 28
	KERN_INVALID_OBJECT           int32 = 29
	KERN_ALREADY_WAITING          int32 = 30
	KERN_DEFAULT_SET              int32 = 31
	KERN_EXCEPTION_PROTECTED      int32 = 32
	KERN_INVALID_LEDGER           int32 = 33
	KERN_INVALID_MEMORY_CONTROL   int32 = 34
	KERN_INVALID_SECURITY         int32 = 35
	KERN_NOT_DEPRESSED            int32 = 36
	KERN_TERMINATED               int32 = 37
	KERN_LOCK_SET_DESTROYED       int32 = 38
	KERN_LOCK_UNSTABLE            int32 = 39
	KERN_LOCK_OWNED               int32 = 40
	KERN_LOCK_OWNED_SELF          int32 = 41
	KERN_SEMAPHORE_DESTROYED      int32 = 42
	KERN_RPC_SERVER_TERMINATED    int32 = 43
	KERN_RPC_TERMINATED           int32 = 44
	KERN_RPC_LOCKED               int32 = 45
	KERN_NOT_WAITING              int32 = 46
	KERN_OPERATION_TIMED_OUT      int32 = 47
	KERN_CODESIGN_ERROR           int32 = 48
	KERN_POLICY_STATIC            int32 = 49
	KERN_INSUFFICIENT_BUFFER_SIZE int32 = 50
	KERN_DENIED                   int32 = 51
	KERN_MISSING_HEX              int32 = 52
)

// Mach message return code constants (from <mach/message.h>, SDK build 25F70).
const (
	MACH_MSG_SUCCESS          int32 = 0x00000000
	MACH_MSG_MASK             int32 = 0x00003e00
	MACH_SEND_IN_PROGRESS     int32 = 0x10000001
	MACH_SEND_INVALID_DATA    int32 = 0x10000002
	MACH_SEND_INVALID_DEST    int32 = 0x10000003
	MACH_SEND_TIMED_OUT       int32 = 0x10000004
	MACH_SEND_INTERRUPTED     int32 = 0x10000007
	MACH_SEND_MSG_TOO_SMALL   int32 = 0x10000008
	MACH_SEND_INVALID_REPLY   int32 = 0x10000009
	MACH_SEND_INVALID_RIGHT   int32 = 0x1000000a
	MACH_SEND_INVALID_NOTIFY  int32 = 0x1000000b
	MACH_SEND_INVALID_MEMORY  int32 = 0x1000000c
	MACH_SEND_NO_BUFFER       int32 = 0x1000000d
	MACH_SEND_TOO_LARGE       int32 = 0x1000000e
	MACH_SEND_INVALID_TYPE    int32 = 0x1000000f
	MACH_SEND_INVALID_HEADER  int32 = 0x10000010
	MACH_SEND_INVALID_TRAILER int32 = 0x10000011

	MACH_RCV_IN_PROGRESS       int32 = 0x10004001
	MACH_RCV_INVALID_NAME      int32 = 0x10004002
	MACH_RCV_TIMED_OUT         int32 = 0x10004003
	MACH_RCV_TOO_LARGE         int32 = 0x10004004
	MACH_RCV_INTERRUPTED       int32 = 0x10004005
	MACH_RCV_PORT_CHANGED      int32 = 0x10004006
	MACH_RCV_INVALID_NOTIFY    int32 = 0x10004007
	MACH_RCV_INVALID_DATA      int32 = 0x10004008
	MACH_RCV_PORT_DIED         int32 = 0x10004009
	MACH_RCV_IN_SET            int32 = 0x1000400a
	MACH_RCV_HEADER_ERROR      int32 = 0x1000400b
	MACH_RCV_BODY_ERROR        int32 = 0x1000400c
	MACH_RCV_IN_PROGRESS_TIMED int32 = 0x10004011
)

var _close func(int32) int32

func init() {
	if frameworkHandle == 0 {
		return
	}
	defer func() { recover() }()
	purego.RegisterLibFunc(&_close, frameworkHandle, "close")
}

// CloseFD closes a descriptor returned by Kqueue. It returns 0 on success.
func CloseFD(fd int) int32 {
	if _close == nil {
		return -1
	}
	return _close(int32(fd))
}

// KeventWait registers changes and waits for events on kq. A nil timeout blocks
// indefinitely. It returns the number of events placed in events, or -1.
//
// Callers must check EV_ERROR on each returned event: kevent64 reports a bad
// registration through the event itself, with errno in Data, rather than
// through the return value.
//
// This wraps the generated Kevent64, which takes its buffers as uintptr.
func KeventWait(kq int, changes []Kevent64_s, events []Kevent64_s, flags uint32, timeout *Timespec) int {
	var chPtr, evPtr, toPtr uintptr
	if len(changes) > 0 {
		chPtr = uintptr(unsafe.Pointer(&changes[0]))
	}
	if len(events) > 0 {
		evPtr = uintptr(unsafe.Pointer(&events[0]))
	}
	if timeout != nil {
		toPtr = uintptr(unsafe.Pointer(timeout))
	}
	n := Kevent64(int32(kq), chPtr, int32(len(changes)), evPtr, int32(len(events)), flags, toPtr)
	runtime.KeepAlive(changes)
	runtime.KeepAlive(events)
	runtime.KeepAlive(timeout)
	return int(n)
}
