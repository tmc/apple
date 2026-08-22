// Code generated from Apple documentation for ParavirtualizedGraphics. DO NOT EDIT.

package paravirtualizedgraphics

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A paravirtualized GPU device object.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDevice
type PGDevice interface {
	objectivec.IObject

	// Reads data from the virtual graphics device’s memory-mapped I/O region.
	//
	// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDevice/mmioRead(atOffset:)
	MmioReadAtOffset(offset uintptr) uint32

	// Writes data to the virtual graphics device’s memory-mapped I/O region.
	//
	// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDevice/mmioWrite(atOffset:value:)
	MmioWriteAtOffsetValue(offset uintptr, value uint32)

	// Notifies the virtual graphics device to start suspending graphics activities.
	//
	// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDevice/willSuspend()
	WillSuspend()

	// Notifies the virtualized graphics device to finish suspending graphics activities.
	//
	// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDevice/finishSuspend()
	FinishSuspend() foundation.NSData

	// Tells a new device object to load a previously saved device’s suspend state.
	//
	// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDevice/willResume(withSuspendState:error:)
	WillResumeWithSuspendStateError(suspendState foundation.NSData) (bool, error)

	// Tells the device object to finish any remaining work to resume processing of a previously saved device’s suspend state.
	//
	// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDevice/didResume()
	DidResume()

	// Create a display from the specified descriptor and uniquifying parameters.
	//
	// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDevice/newDisplay(with:port:serialNum:)
	NewDisplayWithDescriptorPortSerialNum(descriptor IPGDisplayDescriptor, port uint, serialNum uint32) PGDisplay

	// Pause protocol.
	//
	// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDevice/pause()
	Pause()

	// Reset protocol.
	//
	// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDevice/reset()
	Reset()

	// Stop protocol.
	//
	// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDevice/stop()
	Stop()

	// Unpause protocol.
	//
	// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDevice/unpause()
	Unpause()
}

// PGDeviceObject wraps an existing Objective-C object that conforms to the PGDevice protocol.
type PGDeviceObject struct {
	objectivec.Object
}

func (o PGDeviceObject) BaseObject() objectivec.Object {
	return o.Object
}

// PGDeviceObjectFromID constructs a [PGDeviceObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func PGDeviceObjectFromID(id objc.ID) PGDeviceObject {
	return PGDeviceObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Reads data from the virtual graphics device’s memory-mapped I/O region.
//
// offset: The offset into the MMIO bar to write to.
//
// # Return Value
//
// The 32-bit unsigned integer from the MMIO region.
//
// # Discussion
//
// Call this method whenever the guest virtual machine reads from the graphics
// device’s memory-mapped I/O region.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDevice/mmioRead(atOffset:)
func (o PGDeviceObject) MmioReadAtOffset(offset uintptr) uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("mmioReadAtOffset:"), offset)
	return rv
}

// Writes data to the virtual graphics device’s memory-mapped I/O region.
//
// offset: The offset into the MMIO bar to write to.
//
// value: The value to write to memory.
//
// # Discussion
//
// Call this method whenever the guest virtual machine writes to the graphics
// device’s memory-mapped I/O region.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDevice/mmioWrite(atOffset:value:)
func (o PGDeviceObject) MmioWriteAtOffsetValue(offset uintptr, value uint32) {
	objc.Send[struct{}](o.ID, objc.Sel("mmioWriteAtOffset:value:"), offset, value)
}

// Notifies the virtual graphics device to start suspending graphics
// activities.
//
// # Discussion
//
// The virtualized device stops generating interrupts and won’t accept new
// commands from the guest. You must halt any guest CPUs within a short
// interval after you call this method.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDevice/willSuspend()
func (o PGDeviceObject) WillSuspend() {
	objc.Send[struct{}](o.ID, objc.Sel("willSuspend"))
}

// Notifies the virtualized graphics device to finish suspending graphics
// activities.
//
// # Return Value
//
// The suspend state data, or `nil` if an error occurred.
//
// # Discussion
//
// This method may take an arbitrary amount of time as the device needs to
// complete any unfinished GPU work. After this call completes, you can’t
// perform any further operations on this device object and must release it.
//
// Typically, your app serializes the suspend state data to persistant
// storage. Pass the suspend state data to a new device object when you want
// to resume graphics operations.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDevice/finishSuspend()
func (o PGDeviceObject) FinishSuspend() foundation.NSData {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("finishSuspend"))
	return foundation.NSDataFromID(rv)
}

// Tells a new device object to load a previously saved device’s suspend
// state.
//
// suspendState: The suspend data that you previously saved when you suspended an earlier
// device object.
//
// # Return Value
//
// A Boolean value that indicates whether the framework was able to start
// restoring the graphics data.
//
// # Discussion
//
// This method sets up the device to appear in the same state it was in before
// the suspend. The device object doesn’t access guest memory during the
// call to this method.
//
// When resuming from an earlier suspended device, this must be the first
// method that you call on the newly created device object. When you call this
// method, ensure that the guest CPUs aren’t running.
//
// After you call this method, reattach any suspended displays before calling
// [DidResume].
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDevice/willResume(withSuspendState:error:)
func (o PGDeviceObject) WillResumeWithSuspendStateError(suspendState foundation.NSData) (bool, error) {
	rv, err := objc.SendWithError[bool](o.ID, objc.Sel("willResumeWithSuspendState:error:"), suspendState)
	if err != nil {
		return false, err
	}
	return rv, nil
}

// Tells the device object to finish any remaining work to resume processing
// of a previously saved device’s suspend state.
//
// # Discussion
//
// After you call this method, the virtualized device can generate new
// interrupts immediately, even before the call completes. Similarly, guest
// memory must also be accessible before you call this method.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDevice/didResume()
func (o PGDeviceObject) DidResume() {
	objc.Send[struct{}](o.ID, objc.Sel("didResume"))
}

// Create a display from the specified descriptor and uniquifying parameters.
//
// descriptor: The description of the new display.
//
// port: The port number on the accelerator for the device to use for the new
// display. Specify a unique port number for each display.
//
// serialNum: A number that uniquely identifies the display. Ensure that the display
// persists across multiple launches so that the guest compositor can maintain
// a consistent display layout.
//
// # Return Value
//
// A new display object, or `nil` if an error occurred.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDevice/newDisplay(with:port:serialNum:)
func (o PGDeviceObject) NewDisplayWithDescriptorPortSerialNum(descriptor IPGDisplayDescriptor, port uint, serialNum uint32) PGDisplay {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newDisplayWithDescriptor:port:serialNum:"), descriptor, port, serialNum)
	return PGDisplayObjectFromID(rv)
}

// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDevice/pause()
func (o PGDeviceObject) Pause() {
	objc.Send[struct{}](o.ID, objc.Sel("pause"))
}

// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDevice/reset()
func (o PGDeviceObject) Reset() {
	objc.Send[struct{}](o.ID, objc.Sel("reset"))
}

// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDevice/stop()
func (o PGDeviceObject) Stop() {
	objc.Send[struct{}](o.ID, objc.Sel("stop"))
}

// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDevice/unpause()
func (o PGDeviceObject) Unpause() {
	objc.Send[struct{}](o.ID, objc.Sel("unpause"))
}
