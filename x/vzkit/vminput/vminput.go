package vminput

import (
	"context"
	"errors"
	"fmt"
	"time"
	"unsafe"

	"github.com/tmc/apple/objectivec"
	pvz "github.com/tmc/apple/private/virtualization"
	vz "github.com/tmc/apple/virtualization"
	"github.com/tmc/apple/x/vzkit/privatevm"
	"github.com/tmc/apple/x/vzkit/vm"
)

var (
	// ErrVMNotRunning reports that the VM is not running.
	ErrVMNotRunning = errors.New("vm not running")
	// ErrNotReady reports that the VM is not accepting HID reports.
	ErrNotReady = errors.New("vm not accepting hid reports")
)

// Sender sends direct input events to a VM.
type Sender struct {
	queue *vm.Queue
	vm    vz.VZVirtualMachine
}

// NewSender creates a direct input sender for a VM.
func NewSender(queue *vm.Queue, vm vz.VZVirtualMachine) *Sender {
	return &Sender{queue: queue, vm: vm}
}

func (s *Sender) sync(fn func()) {
	if s.queue == nil {
		fn()
		return
	}
	s.queue.Sync(fn)
}

func (s *Sender) privateVM() pvz.VZVirtualMachine {
	return pvz.VZVirtualMachineFromID(s.vm.ID)
}

// Ready reports whether the VM is running and ready for HID reports.
func (s *Sender) Ready() bool {
	return s.requireReady() == nil
}

func (s *Sender) requireReady() error {
	if s.queue == nil || s.queue.Queue().Handle() == 0 {
		return fmt.Errorf("missing vm queue")
	}
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

func (s *Sender) send(fn func()) error {
	if err := s.requireReady(); err != nil {
		return err
	}
	s.sync(fn)
	return nil
}

// SendKeyboardEvents sends keyboard event objects to the VM.
func (s *Sender) SendKeyboardEvents(events unsafe.Pointer, keyboardID uint32) error {
	return s.send(func() {
		s.privateVM().SendKeyboardEventsKeyboardID(events, keyboardID)
	})
}

// SendMouseEvents sends mouse event objects to the VM.
func (s *Sender) SendMouseEvents(events unsafe.Pointer, deviceIndex uint32) error {
	return s.send(func() {
		s.privateVM().SendMouseEventsPointingDeviceIndex(events, deviceIndex)
	})
}

// SendPointerNSEvent forwards a single NSEvent to the VM.
func (s *Sender) SendPointerNSEvent(event objectivec.IObject, deviceIndex uint32) error {
	return s.send(func() {
		s.privateVM().SendPointerNSEventPointingDeviceIndex(event, deviceIndex)
	})
}

// SendScrollWheelEvents sends scroll wheel events to the VM.
func (s *Sender) SendScrollWheelEvents(events unsafe.Pointer, deviceIndex uint32) error {
	return s.send(func() {
		s.privateVM().SendScrollWheelEventsPointingDeviceIndex(events, deviceIndex)
	})
}

// SendDigitizerEvents sends digitizer events to the VM.
func (s *Sender) SendDigitizerEvents(events unsafe.Pointer, deviceIndex uint32) error {
	return s.send(func() {
		s.privateVM().SendDigitizerEventsPointingDeviceIndex(events, deviceIndex)
	})
}

// SendMagnifyEvents sends magnify events to the VM.
func (s *Sender) SendMagnifyEvents(events unsafe.Pointer, deviceIndex uint32) error {
	return s.send(func() {
		s.privateVM().SendMagnifyEventsPointingDeviceIndex(events, deviceIndex)
	})
}

// SendRotationEvents sends rotation events to the VM.
func (s *Sender) SendRotationEvents(events unsafe.Pointer, deviceIndex uint32) error {
	return s.send(func() {
		s.privateVM().SendRotationEventsPointingDeviceIndex(events, deviceIndex)
	})
}

// SendSmartMagnifyEvents sends smart magnify events to the VM.
func (s *Sender) SendSmartMagnifyEvents(events unsafe.Pointer, deviceIndex uint32) error {
	return s.send(func() {
		s.privateVM().SendSmartMagnifyEventsPointingDeviceIndex(events, deviceIndex)
	})
}

// SendQuickLookEvents sends Quick Look events to the VM.
func (s *Sender) SendQuickLookEvents(events unsafe.Pointer, deviceIndex uint32) error {
	return s.send(func() {
		s.privateVM().SendQuickLookEventsPointingDeviceIndex(events, deviceIndex)
	})
}

// SendMultiTouchEvents sends multitouch events to the VM.
func (s *Sender) SendMultiTouchEvents(events unsafe.Pointer, deviceIndex uint32) error {
	return s.send(func() {
		s.privateVM().SendMultiTouchEventsMultiTouchDeviceIndex(events, deviceIndex)
	})
}

// SendText sends a UTF-8 string as a sequence of keyboard events.
// This is intentionally small: callers that need keycode-level control can
// build and send their own event arrays.
func (s *Sender) SendText(text string) error {
	if text == "" {
		return nil
	}
	return fmt.Errorf("SendText is not implemented for direct VM injection; build keyboard events explicitly")
}

// WaitReady blocks until the VM is ready or the context is cancelled.
func (s *Sender) WaitReady(ctx context.Context) error {
	for {
		if s.Ready() {
			return nil
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(10 * time.Millisecond):
		}
	}
}
