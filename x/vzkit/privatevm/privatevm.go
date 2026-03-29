package privatevm

import (
	"errors"
	"fmt"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	vz "github.com/tmc/apple/virtualization"
	"github.com/tmc/apple/x/vzkit/vm"
)

var (
	ErrVMNotRunning = errors.New("vm not running")
	ErrVMNotStopped = errors.New("vm not stopped")
)

func requireRunning(queue *vm.Queue, machine vz.VZVirtualMachine) error {
	state := vm.State(queue, machine)
	if state != vz.VZVirtualMachineStateRunning {
		return fmt.Errorf("%w: state is %s", ErrVMNotRunning, state)
	}
	return nil
}

// Name returns the private _name property of a virtual machine.
func Name(queue *vm.Queue, machine vz.VZVirtualMachine) string {
	var name string
	queue.Sync(func() {
		rv := objc.Send[objc.ID](machine.ID, objc.Sel("_name"))
		name = foundation.NSStringFromID(rv).String()
	})
	return name
}

// SetName sets the private _name property of a virtual machine.
func SetName(queue *vm.Queue, machine vz.VZVirtualMachine, name string) {
	queue.Sync(func() {
		objc.Send[objc.ID](machine.ID, objc.Sel("_setName:"), foundation.NewNSString().InitWithString(name).ID)
	})
}

// CrashContextMessage returns the private _crashContextMessage property.
func CrashContextMessage(queue *vm.Queue, machine vz.VZVirtualMachine) string {
	var msg string
	queue.Sync(func() {
		rv := objc.Send[objc.ID](machine.ID, objc.Sel("_crashContextMessage"))
		msg = foundation.NSStringFromID(rv).String()
	})
	return msg
}

// SetCrashContextMessage sets the private _crashContextMessage property.
func SetCrashContextMessage(queue *vm.Queue, machine vz.VZVirtualMachine, msg string) {
	queue.Sync(func() {
		objc.Send[objc.ID](machine.ID, objc.Sel("_setCrashContextMessage:"), foundation.NewNSString().InitWithString(msg).ID)
	})
}

// StateDescription returns the private _stateDescription of a running VM.
func StateDescription(queue *vm.Queue, machine vz.VZVirtualMachine) (string, error) {
	if err := requireRunning(queue, machine); err != nil {
		return "", err
	}
	var desc string
	queue.Sync(func() {
		rv := objc.Send[objc.ID](machine.ID, objc.Sel("_stateDescription"))
		desc = foundation.NSStringFromID(rv).String()
	})
	return desc, nil
}

// ServicePID returns the private _serviceProcessIdentifier of a running VM.
func ServicePID(queue *vm.Queue, machine vz.VZVirtualMachine) (int, error) {
	if err := requireRunning(queue, machine); err != nil {
		return 0, err
	}
	var pid int
	queue.Sync(func() {
		pid = objc.Send[int](machine.ID, objc.Sel("_serviceProcessIdentifier"))
	})
	return pid, nil
}

// CanCreateCore returns the private _canCreateCore property of a running VM.
func CanCreateCore(queue *vm.Queue, machine vz.VZVirtualMachine) (bool, error) {
	if err := requireRunning(queue, machine); err != nil {
		return false, err
	}
	var ok bool
	queue.Sync(func() {
		ok = objc.Send[bool](machine.ID, objc.Sel("_canCreateCore"))
	})
	return ok, nil
}

// ShouldSendHIDReports returns the private _shouldSendHIDReports property.
func ShouldSendHIDReports(queue *vm.Queue, machine vz.VZVirtualMachine) (bool, error) {
	if err := requireRunning(queue, machine); err != nil {
		return false, err
	}
	var ok bool
	queue.Sync(func() {
		ok = objc.Send[bool](machine.ID, objc.Sel("_shouldSendHIDReports"))
	})
	return ok, nil
}
