// Package usbhid provides safe HID input event helpers for
// Apple Virtualization framework virtual machines.
//
// It wraps the private VZVirtualMachine send* methods with state guards
// that return errors instead of crashing when the VM is not running or
// not yet accepting HID reports.
//
// All methods dispatch on the caller-supplied [vm.Queue] so they are
// safe to call from any goroutine.
package usbhid
