package vzkit

import (
	"errors"
	"fmt"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	vz "github.com/tmc/apple/virtualization"
)

var (
	ErrVMNotRunning = errors.New("vm not running")
	ErrVMNotStopped = errors.New("vm not stopped")
)

// requireRunning returns an error if the VM is not in the running state.
func requireRunning(queue *Queue, vm vz.VZVirtualMachine) error {
	state := VMState(queue, vm)
	if state != vz.VZVirtualMachineStateRunning {
		return fmt.Errorf("%w: state is %s", ErrVMNotRunning, state)
	}
	return nil
}

// VMName returns the private _name property of a VZVirtualMachine.
func VMName(queue *Queue, vm vz.VZVirtualMachine) string {
	var name string
	queue.Sync(func() {
		rv := objc.Send[objc.ID](vm.ID, objc.Sel("_name"))
		name = foundation.NSStringFromID(rv).String()
	})
	return name
}

// SetVMName sets the private _name property of a VZVirtualMachine.
func SetVMName(queue *Queue, vm vz.VZVirtualMachine, name string) {
	queue.Sync(func() {
		objc.Send[objc.ID](vm.ID, objc.Sel("_setName:"), foundation.NewNSString().InitWithString(name).ID)
	})
}

// VMCrashContextMessage returns the private _crashContextMessage property.
func VMCrashContextMessage(queue *Queue, vm vz.VZVirtualMachine) string {
	var msg string
	queue.Sync(func() {
		rv := objc.Send[objc.ID](vm.ID, objc.Sel("_crashContextMessage"))
		msg = foundation.NSStringFromID(rv).String()
	})
	return msg
}

// SetVMCrashContextMessage sets the private _crashContextMessage property.
func SetVMCrashContextMessage(queue *Queue, vm vz.VZVirtualMachine, msg string) {
	queue.Sync(func() {
		objc.Send[objc.ID](vm.ID, objc.Sel("_setCrashContextMessage:"), foundation.NewNSString().InitWithString(msg).ID)
	})
}

// VMStateDescription returns the private _stateDescription of a running VM.
func VMStateDescription(queue *Queue, vm vz.VZVirtualMachine) (string, error) {
	if err := requireRunning(queue, vm); err != nil {
		return "", err
	}
	var desc string
	queue.Sync(func() {
		rv := objc.Send[objc.ID](vm.ID, objc.Sel("_stateDescription"))
		desc = foundation.NSStringFromID(rv).String()
	})
	return desc, nil
}

// VMServicePID returns the private _serviceProcessIdentifier of a running VM.
func VMServicePID(queue *Queue, vm vz.VZVirtualMachine) (int, error) {
	if err := requireRunning(queue, vm); err != nil {
		return 0, err
	}
	var pid int
	queue.Sync(func() {
		pid = objc.Send[int](vm.ID, objc.Sel("_serviceProcessIdentifier"))
	})
	return pid, nil
}

// VMCanCreateCore returns the private _canCreateCore property of a running VM.
func VMCanCreateCore(queue *Queue, vm vz.VZVirtualMachine) (bool, error) {
	if err := requireRunning(queue, vm); err != nil {
		return false, err
	}
	var ok bool
	queue.Sync(func() {
		ok = objc.Send[bool](vm.ID, objc.Sel("_canCreateCore"))
	})
	return ok, nil
}

// VMShouldSendHIDReports returns the private _shouldSendHIDReports property of a running VM.
func VMShouldSendHIDReports(queue *Queue, vm vz.VZVirtualMachine) (bool, error) {
	if err := requireRunning(queue, vm); err != nil {
		return false, err
	}
	var ok bool
	queue.Sync(func() {
		ok = objc.Send[bool](vm.ID, objc.Sel("_shouldSendHIDReports"))
	})
	return ok, nil
}
