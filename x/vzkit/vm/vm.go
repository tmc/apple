// Package vm provides queue management and lifecycle operations for virtual
// machines.
package vm

import (
	"fmt"

	"github.com/tmc/apple/dispatch"
	vz "github.com/tmc/apple/virtualization"
)

// Queue wraps a dispatch queue for VM operations.
// The Virtualization framework requires VM operations to happen on a specific queue.
type Queue struct {
	q dispatch.Queue
}

// NewQueue creates a serial dispatch queue with the given label.
func NewQueue(label string) *Queue {
	return &Queue{q: dispatch.QueueCreate(label)}
}

// WrapQueue creates a Queue from an existing dispatch.Queue.
func WrapQueue(q dispatch.Queue) *Queue {
	return &Queue{q: q}
}

// Queue returns the underlying dispatch.Queue.
func (q *Queue) Queue() dispatch.Queue {
	return q.q
}

// Sync executes work synchronously on the dispatch queue.
func (q *Queue) Sync(work func()) {
	q.q.Sync(work)
}

// Async executes work asynchronously on the dispatch queue.
func (q *Queue) Async(work func()) {
	q.q.Async(work)
}

// Create creates a VZVirtualMachine bound to the given dispatch queue.
func Create(config vz.VZVirtualMachineConfiguration, queue *Queue) vz.VZVirtualMachine {
	return vz.NewVirtualMachineWithConfigurationQueue(&config, queue.Queue())
}

// Start starts a VZVirtualMachine on its queue and calls completion with the result.
func Start(queue *Queue, machine vz.VZVirtualMachine, completion func(error)) {
	queue.Sync(func() {
		machine.StartWithCompletionHandler(completion)
	})
}

// Stop requests a stop of a VZVirtualMachine on its queue.
func Stop(queue *Queue, machine vz.VZVirtualMachine, completion func(error)) {
	queue.Sync(func() {
		machine.StopWithCompletionHandler(completion)
	})
}

// State returns the current state of a VZVirtualMachine, dispatching to queue.
func State(queue *Queue, machine vz.VZVirtualMachine) vz.VZVirtualMachineState {
	var state vz.VZVirtualMachineState
	queue.Sync(func() {
		state = vz.VZVirtualMachineState(machine.State())
	})
	return state
}

// CanStart reports whether the VM is in a state that allows starting.
func CanStart(queue *Queue, machine vz.VZVirtualMachine) bool {
	var ok bool
	queue.Sync(func() {
		ok = machine.CanStart()
	})
	return ok
}

// CanStop reports whether the VM is in a state that allows stopping.
func CanStop(queue *Queue, machine vz.VZVirtualMachine) bool {
	var ok bool
	queue.Sync(func() {
		ok = machine.CanStop()
	})
	return ok
}

// Pause pauses a running VZVirtualMachine on its queue.
func Pause(queue *Queue, machine vz.VZVirtualMachine, completion func(error)) {
	queue.Sync(func() {
		machine.PauseWithCompletionHandler(completion)
	})
}

// Resume resumes a paused VZVirtualMachine on its queue.
func Resume(queue *Queue, machine vz.VZVirtualMachine, completion func(error)) {
	queue.Sync(func() {
		machine.ResumeWithCompletionHandler(completion)
	})
}

// CanPause reports whether the VM is in a state that allows pausing.
func CanPause(queue *Queue, machine vz.VZVirtualMachine) bool {
	var ok bool
	queue.Sync(func() {
		ok = machine.CanPause()
	})
	return ok
}

// CanResume reports whether the VM is in a state that allows resuming.
func CanResume(queue *Queue, machine vz.VZVirtualMachine) bool {
	var ok bool
	queue.Sync(func() {
		ok = machine.CanResume()
	})
	return ok
}

// Validate validates a VZVirtualMachineConfiguration and returns any error.
func Validate(config vz.VZVirtualMachineConfiguration) error {
	valid, err := config.ValidateWithError()
	if !valid {
		if err != nil {
			return fmt.Errorf("vm config: %w", err)
		}
		return fmt.Errorf("vm config: validation failed")
	}
	return nil
}
