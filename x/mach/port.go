package mach

import (
	"github.com/tmc/apple/kernel"
)

// Port is a Mach port name in this task's port namespace.
//
// Port rights are reference-counted in the kernel and leak silently: a
// leaked receive right keeps its object alive forever, and a dropped send
// right becomes a dead name whose next use fails far from the bug. A Port
// is a name, not an object — Go's GC cannot manage right lifetime, so every
// acquired right is balanced explicitly with Deallocate or DestroyReceive,
// typically via defer.
type Port uint32

// PortNull is the null port name.
const PortNull Port = 0

// Right identifies a port right kind for ModRefs and Refs.
type Right uint32

// Port right kinds (mach_port_right_t).
const (
	RightSend     Right = 0
	RightReceive  Right = 1
	RightSendOnce Right = 2
)

// Disposition says how a port right crosses in a message
// (mach_msg_type_name_t). MoveSend consumes the sender's right; CopySend
// does not. Getting this backwards leaks a right or produces a dead name —
// see the leak tests.
type Disposition uint32

const (
	MoveReceive  Disposition = 16
	MoveSend     Disposition = 17
	MoveSendOnce Disposition = 18
	CopySend     Disposition = 19
	MakeSend     Disposition = 20
	MakeSendOnce Disposition = 21
)

// TaskSelf returns this task's port.
func TaskSelf() Port {
	return Port(kernel.Mach_task_self())
}

// NewPort allocates a receive right. Balance with DestroyReceive.
func NewPort() (Port, error) {
	var name kernel.Mach_port_name_t
	kr := kernel.Mach_port_allocate(kernel.Mach_task_self(), kernel.Mach_port_right_t(RightReceive), &name)
	if err := kernError("mach_port_allocate", kr); err != nil {
		return 0, err
	}
	return Port(name), nil
}

// MakeSendRight adds a send right for p's receive right under the same
// name. Balance with Deallocate.
func (p Port) MakeSendRight() error {
	kr := kernel.Mach_port_insert_right(kernel.Mach_task_self(), kernel.Mach_port_name_t(p), uint32(p), kernel.Mach_msg_type_name_t(MakeSend))
	return kernError("mach_port_insert_right", kr)
}

// Deallocate releases one send, send-once, or dead-name reference.
func (p Port) Deallocate() error {
	kr := kernel.Mach_port_deallocate(kernel.Mach_task_self(), kernel.Mach_port_name_t(p))
	return kernError("mach_port_deallocate", kr)
}

// ModRefs adjusts the user-reference count for one right kind.
func (p Port) ModRefs(right Right, delta int) error {
	kr := kernel.Mach_port_mod_refs(kernel.Mach_task_self(), kernel.Mach_port_name_t(p), kernel.Mach_port_right_t(right), kernel.Mach_port_delta_t(delta))
	return kernError("mach_port_mod_refs", kr)
}

// DestroyReceive releases p's receive right.
func (p Port) DestroyReceive() error {
	return p.ModRefs(RightReceive, -1)
}

// Refs reports the user-reference count this task holds for one right kind
// on p. A name with no such right reports zero (the kernel returns
// KERN_INVALID_NAME, which is the answer, not an error, for leak checks).
func (p Port) Refs(right Right) (int, error) {
	var refs kernel.Mach_port_urefs_t
	kr := kernel.Mach_port_get_refs(kernel.Mach_task_self(), kernel.Mach_port_name_t(p), kernel.Mach_port_right_t(right), &refs)
	const kernInvalidName = 15
	if kr == kernInvalidName {
		return 0, nil
	}
	if err := kernError("mach_port_get_refs", kr); err != nil {
		return 0, err
	}
	return int(refs), nil
}
