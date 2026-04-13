package usbhid

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	pvz "github.com/tmc/apple/private/virtualization"
	vz "github.com/tmc/apple/virtualization"
	"github.com/tmc/apple/x/vzkit/privatevm"
	"github.com/tmc/apple/x/vzkit/vm"
)

var (
	ErrVMNotRunning = errors.New("vm not running")
	ErrNotReady     = errors.New("vm not accepting HID reports")
	ErrKeyboardID   = errors.New("keyboard id out of range")
	ErrHIDDeviceID  = errors.New("hid device index out of range")
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

func nsArrayFromObjects[T objectivec.IObject](objs []T) foundation.NSArray {
	return foundation.NSArrayFromID(objectivec.IObjectSliceToNSArray(objs))
}

func (s *Sender) keyboard(keyboardID uint32) (pvz.VZKeyboard, error) {
	keyboards := foundation.NSArrayFromID(objc.Send[objc.ID](s.privateVM().ID, objc.Sel("_keyboards")))
	if keyboardID >= uint32(keyboards.Count()) {
		return pvz.VZKeyboard{}, fmt.Errorf("%w: %d", ErrKeyboardID, keyboardID)
	}
	return pvz.VZKeyboardFromID(keyboards.ObjectAtIndex(uint(keyboardID)).GetID()), nil
}

func (s *Sender) hidDevice(hidDeviceIndex uint32) (pvz.VZHIDDevice, error) {
	devices := foundation.NSArrayFromID(objc.Send[objc.ID](s.privateVM().ID, objc.Sel("_hidDevices")))
	if hidDeviceIndex >= uint32(devices.Count()) {
		return pvz.VZHIDDevice{}, fmt.Errorf("%w: %d", ErrHIDDeviceID, hidDeviceIndex)
	}
	return pvz.VZHIDDeviceFromID(devices.ObjectAtIndex(uint(hidDeviceIndex)).GetID()), nil
}

// SendKeyboardEvents sends an opaque private keyboard-event payload to the VM.
//
// Deprecated: Prefer [Sender.SendKeyEvents] or [Sender.SendKeyboardNSEvents].
// This method forwards a private `void *` contract whose pointed-to layout is
// not documented and is not a raw USB HID report buffer.
func (s *Sender) SendKeyboardEvents(events pvz.VZOpaqueKeyboardEvents, keyboardID uint32) error {
	if err := s.requireReady(); err != nil {
		return err
	}
	s.queue.Sync(func() {
		s.privateVM().SendKeyboardEventsKeyboardID(events, keyboardID)
	})
	return nil
}

// SendKeyEvents sends typed private key event objects to the selected keyboard.
//
// This is the recommended keyboard path. It avoids the undocumented
// `sendKeyboardEvents:keyboardID:` buffer contract and instead uses
// _VZKeyboard.sendKeyEvents: with _VZKeyEvent objects.
func (s *Sender) SendKeyEvents(events []pvz.VZKeyEvent, keyboardID uint32) error {
	if err := s.requireReady(); err != nil {
		return err
	}
	keyboard, err := s.keyboard(keyboardID)
	if err != nil {
		return err
	}
	s.queue.Sync(func() {
		keyboard.SendKeyEvents(nsArrayFromObjects(events))
	})
	return nil
}

// SendKeyboardNSEvents converts NSEvent keyboard events into _VZKeyEvent
// objects, then sends them through _VZKeyboard.sendKeyEvents:.
func (s *Sender) SendKeyboardNSEvents(events []objectivec.IObject, keyboardID uint32) error {
	keyEvents := make([]pvz.VZKeyEvent, 0, len(events))
	for _, event := range events {
		keyEvents = append(keyEvents, pvz.NewVZKeyEventWithEvent(event))
	}
	return s.SendKeyEvents(keyEvents, keyboardID)
}

// SendKeyboardNSEvent sends a single NSEvent keyboard event.
func (s *Sender) SendKeyboardNSEvent(event objectivec.IObject, keyboardID uint32) error {
	return s.SendKeyboardNSEvents([]objectivec.IObject{event}, keyboardID)
}

// SendIOHIDEvents sends an opaque private IOHID event payload to the VM.
//
// Deprecated: Prefer [Sender.SendIOHIDEventObjects] when you already have
// _VZIOHIDEvent objects. This raw method still exposes an undocumented
// `void *` contract.
func (s *Sender) SendIOHIDEvents(events pvz.VZOpaqueIOHIDEvents, hidDeviceIndex uint32) error {
	if err := s.requireReady(); err != nil {
		return err
	}
	s.queue.Sync(func() {
		s.privateVM().SendIOHIDEventsHidDeviceIndex(events, hidDeviceIndex)
	})
	return nil
}

// SendIOHIDEventObjects sends typed _VZIOHIDEvent objects to the selected HID device.
//
// This is only safe when the caller already has valid _VZIOHIDEvent objects.
// The private _VZIOHIDEvent initializer takes an IOHIDEventRef-style pointer,
// not an Objective-C event object, so this package does not currently expose a
// higher-level constructor for keyboard IOHID events.
func (s *Sender) SendIOHIDEventObjects(events []pvz.VZIOHIDEvent, hidDeviceIndex uint32) error {
	if err := s.requireReady(); err != nil {
		return err
	}
	device, err := s.hidDevice(hidDeviceIndex)
	if err != nil {
		return err
	}
	s.queue.Sync(func() {
		device.SendIOHIDEvents(nsArrayFromObjects(events))
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
//
// Deprecated: this still forwards an opaque private payload.
func (s *Sender) SendMouseEvents(events pvz.VZOpaqueMouseEvents, deviceIndex uint32) error {
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
