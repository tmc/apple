package vzkit

import (
	"fmt"

	vz "github.com/tmc/apple/virtualization"
)

// CreateVM creates a VZVirtualMachine bound to the given dispatch queue.
func CreateVM(config vz.VZVirtualMachineConfiguration, queue *Queue) vz.VZVirtualMachine {
	return vz.NewVirtualMachineWithConfigurationQueue(&config, queue.Queue())
}

// StartVM starts a VZVirtualMachine on its queue and calls completion with
// the result. The call dispatches synchronously to the queue to issue the
// start, but the completion fires asynchronously.
func StartVM(queue *Queue, vm vz.VZVirtualMachine, completion func(error)) {
	queue.Sync(func() {
		vm.StartWithCompletionHandler(completion)
	})
}

// StopVM requests a stop of a VZVirtualMachine on its queue.
func StopVM(queue *Queue, vm vz.VZVirtualMachine, completion func(error)) {
	queue.Sync(func() {
		vm.StopWithCompletionHandler(completion)
	})
}

// VMState returns the current state of a VZVirtualMachine, dispatching to queue.
func VMState(queue *Queue, vm vz.VZVirtualMachine) vz.VZVirtualMachineState {
	var state vz.VZVirtualMachineState
	queue.Sync(func() {
		state = vz.VZVirtualMachineState(vm.State())
	})
	return state
}

// CanStart reports whether the VM is in a state that allows starting.
func CanStart(queue *Queue, vm vz.VZVirtualMachine) bool {
	var ok bool
	queue.Sync(func() {
		ok = vm.CanStart()
	})
	return ok
}

// CanStop reports whether the VM is in a state that allows stopping.
func CanStop(queue *Queue, vm vz.VZVirtualMachine) bool {
	var ok bool
	queue.Sync(func() {
		ok = vm.CanStop()
	})
	return ok
}

// PauseVM pauses a running VZVirtualMachine on its queue.
func PauseVM(queue *Queue, vm vz.VZVirtualMachine, completion func(error)) {
	queue.Sync(func() {
		vm.PauseWithCompletionHandler(completion)
	})
}

// ResumeVM resumes a paused VZVirtualMachine on its queue.
func ResumeVM(queue *Queue, vm vz.VZVirtualMachine, completion func(error)) {
	queue.Sync(func() {
		vm.ResumeWithCompletionHandler(completion)
	})
}

// CanPause reports whether the VM is in a state that allows pausing.
func CanPause(queue *Queue, vm vz.VZVirtualMachine) bool {
	var ok bool
	queue.Sync(func() {
		ok = vm.CanPause()
	})
	return ok
}

// CanResume reports whether the VM is in a state that allows resuming.
func CanResume(queue *Queue, vm vz.VZVirtualMachine) bool {
	var ok bool
	queue.Sync(func() {
		ok = vm.CanResume()
	})
	return ok
}

// ValidateConfig validates a VZVirtualMachineConfiguration and returns any error.
func ValidateConfig(config vz.VZVirtualMachineConfiguration) error {
	valid, err := config.ValidateWithError()
	if !valid {
		if err != nil {
			return fmt.Errorf("vm config: %w", err)
		}
		return fmt.Errorf("vm config: validation failed")
	}
	return nil
}

// Retain retains an Objective-C object to prevent premature deallocation.
// This is needed for objects passed to the Virtualization framework.
func Retain(id interface{ Retain() }) {
	id.Retain()
}
