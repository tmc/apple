package usbhid

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/tmc/apple/objectivec"
	pvz "github.com/tmc/apple/private/virtualization"
	vz "github.com/tmc/apple/virtualization"
	"github.com/tmc/apple/x/vzkit/privatevm"
	"github.com/tmc/apple/x/vzkit/vm"
)

var (
	ErrVMNotRunning = errors.New("vm not running")
	ErrNotReady     = errors.New("vm not accepting HID reports")
)

// Sender sends HID events to a virtual machine.
type Sender struct {
	queue *vm.Queue
	vm    vz.VZVirtualMachine
}

// NewSender creates a new HID event sender for a running VM.
func NewSender(queue *vm.Queue, vm vz.VZVirtualMachine) *Sender {
	return &Sender{queue: queue, vm: vm}
}

func (s *Sender) privateVM() pvz.VZVirtualMachine {
	return pvz.VZVirtualMachineFromID(s.vm.ID)
}

// requireReady checks the VM is running and accepting HID reports.
func (s *Sender) requireReady() error {
	state := vm.State(s.queue, s.vm)
	if state != vz.VZVirtualMachineStateRunning {
		return fmt.Errorf("%w: state is %s", ErrVMNotRunning, state)
	}
	ok, err := privatevm.ShouldSendHIDReports(s.queue, s.vm)
	if err != nil {
		return err
	}
	if !ok {
		return ErrNotReady
	}
	return nil
}

// Ready reports whether the VM is running and accepting HID input.
func (s *Sender) Ready() bool {
	return s.requireReady() == nil
}

// SendKeyboardEvents sends keyboard events to the VM.
// The events parameter is an NSArray of keyboard event objects and keyboardID
// identifies which keyboard device to target (typically 0).
func (s *Sender) SendKeyboardEvents(events unsafe.Pointer, keyboardID uint32) error {
	if err := s.requireReady(); err != nil {
		return err
	}
	s.queue.Sync(func() {
		s.privateVM().SendKeyboardEventsKeyboardID(events, keyboardID)
	})
	return nil
}

// SendDigitizerEvents sends digitizer (mouse/touch) events to the VM.
// The events parameter is an NSArray of digitizer event objects and deviceIndex
// identifies which pointing device to target (typically 0).
func (s *Sender) SendDigitizerEvents(events unsafe.Pointer, deviceIndex uint32) error {
	if err := s.requireReady(); err != nil {
		return err
	}
	s.queue.Sync(func() {
		s.privateVM().SendDigitizerEventsPointingDeviceIndex(events, deviceIndex)
	})
	return nil
}

// SendMouseEvents sends mouse events to the VM.
// The events parameter is an NSArray of mouse event objects and deviceIndex
// identifies which pointing device to target (typically 0).
func (s *Sender) SendMouseEvents(events unsafe.Pointer, deviceIndex uint32) error {
	if err := s.requireReady(); err != nil {
		return err
	}
	s.queue.Sync(func() {
		s.privateVM().SendMouseEventsPointingDeviceIndex(events, deviceIndex)
	})
	return nil
}

// SendScrollWheelEvents sends scroll wheel events to the VM.
// The events parameter is an NSArray of scroll wheel event objects and
// deviceIndex identifies which pointing device to target (typically 0).
func (s *Sender) SendScrollWheelEvents(events unsafe.Pointer, deviceIndex uint32) error {
	if err := s.requireReady(); err != nil {
		return err
	}
	s.queue.Sync(func() {
		s.privateVM().SendScrollWheelEventsPointingDeviceIndex(events, deviceIndex)
	})
	return nil
}

// SendPointerNSEvent forwards an AppKit NSEvent to the VM as pointer input.
// The event is sent directly (not wrapped in an array) and deviceIndex
// identifies which pointing device to target (typically 0).
func (s *Sender) SendPointerNSEvent(event objectivec.IObject, deviceIndex uint32) error {
	if err := s.requireReady(); err != nil {
		return err
	}
	s.queue.Sync(func() {
		s.privateVM().SendPointerNSEventPointingDeviceIndex(event, deviceIndex)
	})
	return nil
}
