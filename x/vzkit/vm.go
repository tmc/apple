package vzkit

import (
	vz "github.com/tmc/apple/virtualization"
	vmruntime "github.com/tmc/apple/x/vzkit/vm"
)

// CreateVM creates a VZVirtualMachine bound to the given dispatch queue.
func CreateVM(config vz.VZVirtualMachineConfiguration, queue *Queue) vz.VZVirtualMachine {
	return vmruntime.Create(config, queue)
}

// StartVM starts a VZVirtualMachine on its queue.
func StartVM(queue *Queue, vm vz.VZVirtualMachine, completion func(error)) {
	vmruntime.Start(queue, vm, completion)
}

// StopVM requests a stop of a VZVirtualMachine on its queue.
func StopVM(queue *Queue, vm vz.VZVirtualMachine, completion func(error)) {
	vmruntime.Stop(queue, vm, completion)
}

// VMState returns the current state of a VZVirtualMachine.
func VMState(queue *Queue, vm vz.VZVirtualMachine) vz.VZVirtualMachineState {
	return vmruntime.State(queue, vm)
}

// CanStart reports whether the VM is in a state that allows starting.
func CanStart(queue *Queue, vm vz.VZVirtualMachine) bool { return vmruntime.CanStart(queue, vm) }

// CanStop reports whether the VM is in a state that allows stopping.
func CanStop(queue *Queue, vm vz.VZVirtualMachine) bool { return vmruntime.CanStop(queue, vm) }

// PauseVM pauses a running VZVirtualMachine on its queue.
func PauseVM(queue *Queue, vm vz.VZVirtualMachine, completion func(error)) {
	vmruntime.Pause(queue, vm, completion)
}

// ResumeVM resumes a paused VZVirtualMachine on its queue.
func ResumeVM(queue *Queue, vm vz.VZVirtualMachine, completion func(error)) {
	vmruntime.Resume(queue, vm, completion)
}

// CanPause reports whether the VM is in a state that allows pausing.
func CanPause(queue *Queue, vm vz.VZVirtualMachine) bool { return vmruntime.CanPause(queue, vm) }

// CanResume reports whether the VM is in a state that allows resuming.
func CanResume(queue *Queue, vm vz.VZVirtualMachine) bool { return vmruntime.CanResume(queue, vm) }

// ValidateConfig validates a VZVirtualMachineConfiguration and returns any error.
func ValidateConfig(config vz.VZVirtualMachineConfiguration) error { return vmruntime.Validate(config) }

// Retain retains an Objective-C object to prevent premature deallocation.
func Retain(id interface{ Retain() }) { id.Retain() }
