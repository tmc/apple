package vzkit

import (
	vz "github.com/tmc/apple/virtualization"
	privatevmx "github.com/tmc/apple/x/vzkit/privatevm"
)

var (
	ErrVMNotRunning = privatevmx.ErrVMNotRunning
	ErrVMNotStopped = privatevmx.ErrVMNotStopped
)

// VMName returns the private _name property of a VZVirtualMachine.
func VMName(queue *Queue, vm vz.VZVirtualMachine) string { return privatevmx.Name(queue, vm) }

// SetVMName sets the private _name property of a VZVirtualMachine.
func SetVMName(queue *Queue, vm vz.VZVirtualMachine, name string) {
	privatevmx.SetName(queue, vm, name)
}

// VMCrashContextMessage returns the private _crashContextMessage property.
func VMCrashContextMessage(queue *Queue, vm vz.VZVirtualMachine) string {
	return privatevmx.CrashContextMessage(queue, vm)
}

// SetVMCrashContextMessage sets the private _crashContextMessage property.
func SetVMCrashContextMessage(queue *Queue, vm vz.VZVirtualMachine, msg string) {
	privatevmx.SetCrashContextMessage(queue, vm, msg)
}

// VMStateDescription returns the private _stateDescription of a running VM.
func VMStateDescription(queue *Queue, vm vz.VZVirtualMachine) (string, error) {
	return privatevmx.StateDescription(queue, vm)
}

// VMServicePID returns the private _serviceProcessIdentifier of a running VM.
func VMServicePID(queue *Queue, vm vz.VZVirtualMachine) (int, error) {
	return privatevmx.ServicePID(queue, vm)
}

// VMCanCreateCore returns the private _canCreateCore property of a running VM.
func VMCanCreateCore(queue *Queue, vm vz.VZVirtualMachine) (bool, error) {
	return privatevmx.CanCreateCore(queue, vm)
}

// VMShouldSendHIDReports returns the private _shouldSendHIDReports property.
func VMShouldSendHIDReports(queue *Queue, vm vz.VZVirtualMachine) (bool, error) {
	return privatevmx.ShouldSendHIDReports(queue, vm)
}
