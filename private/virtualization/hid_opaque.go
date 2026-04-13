package virtualization

import "unsafe"

// VZOpaqueKeyboardEvents is the opaque buffer type expected by the private
// VZVirtualMachine selector sendKeyboardEvents:keyboardID:.
//
// The runtime encoding for this selector is `void *`, but the pointed-to
// memory is not a raw USB HID keyboard report. Treat this as an opaque private
// ABI contract.
//
// Prefer object-based key injection via [_VZKeyboard.sendKeyEvents:] and
// [_VZKeyEvent] instead of this raw VM selector.
type VZOpaqueKeyboardEvents struct{ ptr unsafe.Pointer }

// UnsafeKeyboardEvents wraps a private keyboard event buffer pointer.
func UnsafeKeyboardEvents(ptr unsafe.Pointer) VZOpaqueKeyboardEvents {
	return VZOpaqueKeyboardEvents{ptr: ptr}
}

// UnsafePointer returns the underlying private ABI pointer.
func (v VZOpaqueKeyboardEvents) UnsafePointer() unsafe.Pointer { return v.ptr }

// VZOpaqueMouseEvents is the opaque buffer type expected by the private
// VZVirtualMachine selector sendMouseEvents:pointingDeviceIndex:.
type VZOpaqueMouseEvents struct{ ptr unsafe.Pointer }

// UnsafeMouseEvents wraps a private mouse event buffer pointer.
func UnsafeMouseEvents(ptr unsafe.Pointer) VZOpaqueMouseEvents { return VZOpaqueMouseEvents{ptr: ptr} }

// UnsafePointer returns the underlying private ABI pointer.
func (v VZOpaqueMouseEvents) UnsafePointer() unsafe.Pointer { return v.ptr }

// VZOpaqueIOHIDEvents is the opaque buffer type expected by the private
// VZVirtualMachine selector sendIOHIDEvents:hidDeviceIndex:.
type VZOpaqueIOHIDEvents struct{ ptr unsafe.Pointer }

// UnsafeIOHIDEvents wraps a private IOHID event buffer pointer.
func UnsafeIOHIDEvents(ptr unsafe.Pointer) VZOpaqueIOHIDEvents { return VZOpaqueIOHIDEvents{ptr: ptr} }

// UnsafePointer returns the underlying private ABI pointer.
func (v VZOpaqueIOHIDEvents) UnsafePointer() unsafe.Pointer { return v.ptr }

// VZOpaqueHIDReports is the opaque buffer type expected by the private
// _processHIDReports:forDevice:deviceType: selector.
type VZOpaqueHIDReports struct{ ptr unsafe.Pointer }

// UnsafeHIDReports wraps a private HID reports buffer pointer.
func UnsafeHIDReports(ptr unsafe.Pointer) VZOpaqueHIDReports { return VZOpaqueHIDReports{ptr: ptr} }

// UnsafePointer returns the underlying private ABI pointer.
func (v VZOpaqueHIDReports) UnsafePointer() unsafe.Pointer { return v.ptr }
