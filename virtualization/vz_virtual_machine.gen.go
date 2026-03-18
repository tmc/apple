// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"context"
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZVirtualMachine] class.
var (
	_VZVirtualMachineClass     VZVirtualMachineClass
	_VZVirtualMachineClassOnce sync.Once
)

func getVZVirtualMachineClass() VZVirtualMachineClass {
	_VZVirtualMachineClassOnce.Do(func() {
		_VZVirtualMachineClass = VZVirtualMachineClass{class: objc.GetClass("VZVirtualMachine")}
	})
	return _VZVirtualMachineClass
}

// GetVZVirtualMachineClass returns the class object for VZVirtualMachine.
func GetVZVirtualMachineClass() VZVirtualMachineClass {
	return getVZVirtualMachineClass()
}

type VZVirtualMachineClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtualMachineClass) Alloc() VZVirtualMachine {
	rv := objc.Send[VZVirtualMachine](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An object that manages the overall state and configuration of your VM.
//
// # Overview
// 
// A [VZVirtualMachine] object emulates a complete hardware machine of the
// same architecture as the underlying Mac computer. Use the VM to execute a
// guest operating system and any other apps you install. The VM manages the
// resources that the guest operating system uses, providing access to some
// hardware resources while emulating others.
// 
// Create and configure a [VZVirtualMachineConfiguration] object with details
// about how you want to configure your VM, and use that object to create the
// [VZVirtualMachine] object. After creating the VM, call the
// [start(completionHandler:)] method (Swift) or the
// [VZVirtualMachine.StartWithCompletionHandler] method (Objective-C) to start the VM and boot
// the guest operating system.
//
// [start(completionHandler:)]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/start(completionHandler:)
//
// # Creating the VM
//
//   - [VZVirtualMachine.InitWithConfiguration]: Creates the VM and configures it with the specified data.
//   - [VZVirtualMachine.InitWithConfigurationQueue]: Creates and configures the VM with the specified data and dispatch queue.
//
// # Responding to a stopped VM
//
//   - [VZVirtualMachine.Delegate]: A custom object you use to determine when the VM stops.
//   - [VZVirtualMachine.SetDelegate]
//
// # Starting and stopping the VM
//
//   - [VZVirtualMachine.StartWithCompletionHandler]: Starts the VM and notifies the specified completion handler if startup was successful.
//   - [VZVirtualMachine.StartWithOptionsCompletionHandler]: Starts the VM with the options and a completion handler you provide.
//   - [VZVirtualMachine.StopWithCompletionHandler]: Stops a VM that’s in either a running or paused state.
//   - [VZVirtualMachine.RequestStopWithError]: Asks the guest operating system to stop running.
//   - [VZVirtualMachine.PauseWithCompletionHandler]: Pauses a running VM and notifies the specified completion handler of the results.
//   - [VZVirtualMachine.ResumeWithCompletionHandler]: Resumes a paused VM and notifies the specified completion handler of the results.
//
// # Configuring VM attributes at runtime
//
//   - [VZVirtualMachine.ConsoleDevices]: The list of configured console devices on the VM.
//   - [VZVirtualMachine.MemoryBalloonDevices]: The array of devices that you use to adjust the amount of memory available to the guest system.
//   - [VZVirtualMachine.NetworkDevices]: The list of configured network devices on the VM.
//   - [VZVirtualMachine.SocketDevices]: The array of socket devices that the VM configures for use ports in the guest VM.
//   - [VZVirtualMachine.DirectorySharingDevices]: The list of configured directory-sharing devices on the VM.
//   - [VZVirtualMachine.UsbControllers]: The list of runtime USB controller objects.
//
// # Getting the state of the VM
//
//   - [VZVirtualMachine.State]: The current execution state of the VM.
//   - [VZVirtualMachine.GraphicsDevices]: The list of configured graphics devices on the virtual machine.
//   - [VZVirtualMachine.CanStart]: A Boolean value that indicates whether you can start the VM.
//   - [VZVirtualMachine.CanPause]: A Boolean value that indicates whether you can pause the VM.
//   - [VZVirtualMachine.CanResume]: A Boolean value that indicates whether you can resume the VM.
//   - [VZVirtualMachine.CanStop]: A Boolean value that indicates whether you can stop the VM.
//   - [VZVirtualMachine.CanRequestStop]: A Boolean value that indicates whether you can ask the guest operating system to stop running.
//   - [VZVirtualMachine.Queue]: The queue associated with this virtual machine.
//
// # Saving and restoring the VM state
//
//   - [VZVirtualMachine.SaveMachineStateToURLCompletionHandler]: Saves the state of a VM.
//   - [VZVirtualMachine.RestoreMachineStateFromURLCompletionHandler]: Restores a VM from a previously saved state.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine
type VZVirtualMachine struct {
	objectivec.Object
}

// VZVirtualMachineFromID constructs a [VZVirtualMachine] from an objc.ID.
//
// An object that manages the overall state and configuration of your VM.
func VZVirtualMachineFromID(id objc.ID) VZVirtualMachine {
	return VZVirtualMachine{objectivec.Object{ID: id}}
}
// NOTE: VZVirtualMachine adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZVirtualMachine] class.
//
// # Creating the VM
//
//   - [IVZVirtualMachine.InitWithConfiguration]: Creates the VM and configures it with the specified data.
//   - [IVZVirtualMachine.InitWithConfigurationQueue]: Creates and configures the VM with the specified data and dispatch queue.
//
// # Responding to a stopped VM
//
//   - [IVZVirtualMachine.Delegate]: A custom object you use to determine when the VM stops.
//   - [IVZVirtualMachine.SetDelegate]
//
// # Starting and stopping the VM
//
//   - [IVZVirtualMachine.StartWithCompletionHandler]: Starts the VM and notifies the specified completion handler if startup was successful.
//   - [IVZVirtualMachine.StartWithOptionsCompletionHandler]: Starts the VM with the options and a completion handler you provide.
//   - [IVZVirtualMachine.StopWithCompletionHandler]: Stops a VM that’s in either a running or paused state.
//   - [IVZVirtualMachine.RequestStopWithError]: Asks the guest operating system to stop running.
//   - [IVZVirtualMachine.PauseWithCompletionHandler]: Pauses a running VM and notifies the specified completion handler of the results.
//   - [IVZVirtualMachine.ResumeWithCompletionHandler]: Resumes a paused VM and notifies the specified completion handler of the results.
//
// # Configuring VM attributes at runtime
//
//   - [IVZVirtualMachine.ConsoleDevices]: The list of configured console devices on the VM.
//   - [IVZVirtualMachine.MemoryBalloonDevices]: The array of devices that you use to adjust the amount of memory available to the guest system.
//   - [IVZVirtualMachine.NetworkDevices]: The list of configured network devices on the VM.
//   - [IVZVirtualMachine.SocketDevices]: The array of socket devices that the VM configures for use ports in the guest VM.
//   - [IVZVirtualMachine.DirectorySharingDevices]: The list of configured directory-sharing devices on the VM.
//   - [IVZVirtualMachine.UsbControllers]: The list of runtime USB controller objects.
//
// # Getting the state of the VM
//
//   - [IVZVirtualMachine.State]: The current execution state of the VM.
//   - [IVZVirtualMachine.GraphicsDevices]: The list of configured graphics devices on the virtual machine.
//   - [IVZVirtualMachine.CanStart]: A Boolean value that indicates whether you can start the VM.
//   - [IVZVirtualMachine.CanPause]: A Boolean value that indicates whether you can pause the VM.
//   - [IVZVirtualMachine.CanResume]: A Boolean value that indicates whether you can resume the VM.
//   - [IVZVirtualMachine.CanStop]: A Boolean value that indicates whether you can stop the VM.
//   - [IVZVirtualMachine.CanRequestStop]: A Boolean value that indicates whether you can ask the guest operating system to stop running.
//   - [IVZVirtualMachine.Queue]: The queue associated with this virtual machine.
//
// # Saving and restoring the VM state
//
//   - [IVZVirtualMachine.SaveMachineStateToURLCompletionHandler]: Saves the state of a VM.
//   - [IVZVirtualMachine.RestoreMachineStateFromURLCompletionHandler]: Restores a VM from a previously saved state.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine
type IVZVirtualMachine interface {
	objectivec.IObject

	// Topic: Creating the VM

	// Creates the VM and configures it with the specified data.
	InitWithConfiguration(configuration IVZVirtualMachineConfiguration) VZVirtualMachine
	// Creates and configures the VM with the specified data and dispatch queue.
	InitWithConfigurationQueue(configuration IVZVirtualMachineConfiguration, queue dispatch.Queue) VZVirtualMachine

	// Topic: Responding to a stopped VM

	// A custom object you use to determine when the VM stops.
	Delegate() VZVirtualMachineDelegate
	SetDelegate(value VZVirtualMachineDelegate)

	// Topic: Starting and stopping the VM

	// Starts the VM and notifies the specified completion handler if startup was successful.
	StartWithCompletionHandler(completionHandler ErrorHandler)
	// Starts the VM with the options and a completion handler you provide.
	StartWithOptionsCompletionHandler(options IVZVirtualMachineStartOptions, completionHandler ErrorHandler)
	// Stops a VM that’s in either a running or paused state.
	StopWithCompletionHandler(completionHandler ErrorHandler)
	// Asks the guest operating system to stop running.
	RequestStopWithError() (bool, error)
	// Pauses a running VM and notifies the specified completion handler of the results.
	PauseWithCompletionHandler(completionHandler ErrorHandler)
	// Resumes a paused VM and notifies the specified completion handler of the results.
	ResumeWithCompletionHandler(completionHandler ErrorHandler)

	// Topic: Configuring VM attributes at runtime

	// The list of configured console devices on the VM.
	ConsoleDevices() []VZConsoleDevice
	// The array of devices that you use to adjust the amount of memory available to the guest system.
	MemoryBalloonDevices() []VZMemoryBalloonDevice
	// The list of configured network devices on the VM.
	NetworkDevices() []VZNetworkDevice
	// The array of socket devices that the VM configures for use ports in the guest VM.
	SocketDevices() []VZSocketDevice
	// The list of configured directory-sharing devices on the VM.
	DirectorySharingDevices() []VZDirectorySharingDevice
	// The list of runtime USB controller objects.
	UsbControllers() []VZUSBController

	// Topic: Getting the state of the VM

	// The current execution state of the VM.
	State() VZVirtualMachineState
	// The list of configured graphics devices on the virtual machine.
	GraphicsDevices() []VZGraphicsDevice
	// A Boolean value that indicates whether you can start the VM.
	CanStart() bool
	// A Boolean value that indicates whether you can pause the VM.
	CanPause() bool
	// A Boolean value that indicates whether you can resume the VM.
	CanResume() bool
	// A Boolean value that indicates whether you can stop the VM.
	CanStop() bool
	// A Boolean value that indicates whether you can ask the guest operating system to stop running.
	CanRequestStop() bool
	// The queue associated with this virtual machine.
	Queue() dispatch.Queue

	// Topic: Saving and restoring the VM state

	// Saves the state of a VM.
	SaveMachineStateToURLCompletionHandler(saveFileURL foundation.INSURL, completionHandler ErrorHandler)
	// Restores a VM from a previously saved state.
	RestoreMachineStateFromURLCompletionHandler(saveFileURL foundation.INSURL, completionHandler ErrorHandler)
}

// Init initializes the instance.
func (v VZVirtualMachine) Init() VZVirtualMachine {
	rv := objc.Send[VZVirtualMachine](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtualMachine) Autorelease() VZVirtualMachine {
	rv := objc.Send[VZVirtualMachine](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtualMachine creates a new VZVirtualMachine instance.
func NewVZVirtualMachine() VZVirtualMachine {
	class := getVZVirtualMachineClass()
	rv := objc.Send[VZVirtualMachine](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates the VM and configures it with the specified data.
//
// configuration: The configuration of the VM. The configuration must be valid, and you can
// verify that it’s valid by calling its [ValidateWithError] method. The VM
// stores a copy of the configuration.
//
// # Return Value
// 
// An initialized VM object.
//
// # Discussion
// 
// This VM uses your app’s main queue for all operations. You must perform
// all VM-related operations on the main queue, and the VM executes all
// callbacks and delegate methods on the main queue.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/init(configuration:)
func NewVirtualMachineWithConfiguration(configuration IVZVirtualMachineConfiguration) VZVirtualMachine {
	instance := getVZVirtualMachineClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithConfiguration:"), configuration)
	return VZVirtualMachineFromID(rv)
}

// Creates and configures the VM with the specified data and dispatch queue.
//
// configuration: The configuration of the VM. The configuration must be valid, and you can
// verify that it’s valid by calling its [ValidateWithError] method. The VM
// stores a copy of the configuration.
//
// queue: The serial dispatch queue for the VM. You must perform all VM-related
// operations on the specified queue, and the VM executes callbacks and
// delegate methods on the queue. If the queue isn’t serial, the behavior
// isn’t defined.
//
// # Return Value
// 
// An initialized VM object.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/init(configuration:queue:)
func NewVirtualMachineWithConfigurationQueue(configuration IVZVirtualMachineConfiguration, queue dispatch.Queue) VZVirtualMachine {
	instance := getVZVirtualMachineClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithConfiguration:queue:"), configuration, uintptr(queue.Handle()))
	return VZVirtualMachineFromID(rv)
}

// Creates the VM and configures it with the specified data.
//
// configuration: The configuration of the VM. The configuration must be valid, and you can
// verify that it’s valid by calling its [ValidateWithError] method. The VM
// stores a copy of the configuration.
//
// # Return Value
// 
// An initialized VM object.
//
// # Discussion
// 
// This VM uses your app’s main queue for all operations. You must perform
// all VM-related operations on the main queue, and the VM executes all
// callbacks and delegate methods on the main queue.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/init(configuration:)
func (v VZVirtualMachine) InitWithConfiguration(configuration IVZVirtualMachineConfiguration) VZVirtualMachine {
	rv := objc.Send[VZVirtualMachine](v.ID, objc.Sel("initWithConfiguration:"), configuration)
	return rv
}

// Creates and configures the VM with the specified data and dispatch queue.
//
// configuration: The configuration of the VM. The configuration must be valid, and you can
// verify that it’s valid by calling its [ValidateWithError] method. The VM
// stores a copy of the configuration.
//
// queue: The serial dispatch queue for the VM. You must perform all VM-related
// operations on the specified queue, and the VM executes callbacks and
// delegate methods on the queue. If the queue isn’t serial, the behavior
// isn’t defined.
//
// # Return Value
// 
// An initialized VM object.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/init(configuration:queue:)
func (v VZVirtualMachine) InitWithConfigurationQueue(configuration IVZVirtualMachineConfiguration, queue dispatch.Queue) VZVirtualMachine {
	rv := objc.Send[VZVirtualMachine](v.ID, objc.Sel("initWithConfiguration:queue:"), configuration, uintptr(queue.Handle()))
	return rv
}

// Starts the VM and notifies the specified completion handler if startup was
// successful.
//
// # Discussion
// 
// Call this method to start a VM that’s in the [VirtualMachineStateStopped]
// or [VirtualMachineStateError] state.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/start()
func (v VZVirtualMachine) StartWithCompletionHandler(completionHandler ErrorHandler) {
_block0, _cleanup0 := NewErrorBlock(completionHandler)
	defer _cleanup0()
	objc.Send[objc.ID](v.ID, objc.Sel("startWithCompletionHandler:"), _block0)
}

// Starts the VM with the options and a completion handler you provide.
//
// options: A [VZVirtualMachineStartOptions] object that describes controlling startup
// behavior of a VM using [VZMacOSBootLoader].
//
// completionHandler: The block to call with the results of the startup attempt. This block has
// no return value and has one [NSError] object as its parameter:
// 
// error: A result type that contains an error object when the VM fails to
// start.
// //
// [NSError]: https://developer.apple.com/documentation/Foundation/NSError
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/start(options:completionHandler:)
func (v VZVirtualMachine) StartWithOptionsCompletionHandler(options IVZVirtualMachineStartOptions, completionHandler ErrorHandler) {
_block1, _cleanup1 := NewErrorBlock(completionHandler)
	defer _cleanup1()
	objc.Send[objc.ID](v.ID, objc.Sel("startWithOptions:completionHandler:"), options, _block1)
}

// Stops a VM that’s in either a running or paused state.
//
// completionHandler: A block called after the VM stopped successfully, or on error. The error
// parameter passed to the block is `nil` if the stop was successful.
//
// # Discussion
// 
// To determine if a VM is in a state that allows you to stop it, check the
// VM’s [CanStop] or [CanRequestStop] properties.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/stop(completionHandler:)
func (v VZVirtualMachine) StopWithCompletionHandler(completionHandler ErrorHandler) {
_block0, _cleanup0 := NewErrorBlock(completionHandler)
	defer _cleanup0()
	objc.Send[objc.ID](v.ID, objc.Sel("stopWithCompletionHandler:"), _block0)
}

// Asks the guest operating system to stop running.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/requestStop()
func (v VZVirtualMachine) RequestStopWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("requestStopWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("requestStopWithError: returned NO with nil NSError")
	}
	return rv, nil

}

// Pauses a running VM and notifies the specified completion handler of the
// results.
//
// # Discussion
// 
// Call this method to pause a VM that’s in the [VirtualMachineStateRunning]
// state. To determine if a VM is in a state that allows you to pause it,
// check the VM’s [CanPause] property.
// 
// If the VM stops before the attempt to pause it finishes, this method calls
// the completion handler with an error.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/pause()
func (v VZVirtualMachine) PauseWithCompletionHandler(completionHandler ErrorHandler) {
_block0, _cleanup0 := NewErrorBlock(completionHandler)
	defer _cleanup0()
	objc.Send[objc.ID](v.ID, objc.Sel("pauseWithCompletionHandler:"), _block0)
}

// Resumes a paused VM and notifies the specified completion handler of the
// results.
//
// # Discussion
// 
// Call this method to resume a VM that’s in the [VirtualMachineStatePaused]
// state. To determine if a VM is in a state that allows you to resume it,
// check the VM’s [CanResume] property.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/resume()
func (v VZVirtualMachine) ResumeWithCompletionHandler(completionHandler ErrorHandler) {
_block0, _cleanup0 := NewErrorBlock(completionHandler)
	defer _cleanup0()
	objc.Send[objc.ID](v.ID, objc.Sel("resumeWithCompletionHandler:"), _block0)
}

// Saves the state of a VM.
//
// saveFileURL: An [NSURL] that indicates the location where the framework writes the saved
// state of the VM.
// //
// [NSURL]: https://developer.apple.com/documentation/Foundation/NSURL
//
// completionHandler: A block the framework calls after successfully saving the VM or upon
// returning an error.
// 
// The error parameter passed to the block is `nil` if the save was
// successful.
//
// # Discussion
// 
// Use this method to save a paused VM to a file. You can use the contents of
// this file later to restore the state of the paused VM.
// 
// This call fails if the VM isn’t in a paused state or if the
// Virtualization framework can’t save the VM. If this method fails, the
// framework returns an error, and the VM state remains unchanged.
// 
// If this method is successful, the framework writes the file, and the VM
// state remains unchanged.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/saveMachineStateTo(url:completionHandler:)
func (v VZVirtualMachine) SaveMachineStateToURLCompletionHandler(saveFileURL foundation.INSURL, completionHandler ErrorHandler) {
_block1, _cleanup1 := NewErrorBlock(completionHandler)
	defer _cleanup1()
	objc.Send[objc.ID](v.ID, objc.Sel("saveMachineStateToURL:completionHandler:"), saveFileURL, _block1)
}

// Restores a VM from a previously saved state.
//
// saveFileURL: An [NSURL] that indicates the location where the framework reads the saved
// state of the VM.
// //
// [NSURL]: https://developer.apple.com/documentation/Foundation/NSURL
//
// completionHandler: A block the framework calls after the VM has been successfully restored or
// upon an error.
//
// # Discussion
// 
// Use this method to restore a stopped VM to a state previously saved to file
// through [SaveMachineStateToURLCompletionHandler].
// 
// The method fails if any of the following conditions are true:
// 
// - The Virtualization framework can’t open or read the file. - The file
// contents are incompatible with the current configuration. - The VM you’re
// trying to restore isn’t in the [VirtualMachineStateStopped] state.
// 
// If this method fails, the framework returns an error, and the VM state
// doesn’t change.
// 
// If this method is successful, the framework restores the VM and places it
// in the paused state.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/restoreMachineStateFrom(url:completionHandler:)
func (v VZVirtualMachine) RestoreMachineStateFromURLCompletionHandler(saveFileURL foundation.INSURL, completionHandler ErrorHandler) {
_block1, _cleanup1 := NewErrorBlock(completionHandler)
	defer _cleanup1()
	objc.Send[objc.ID](v.ID, objc.Sel("restoreMachineStateFromURL:completionHandler:"), saveFileURL, _block1)
}

// A custom object you use to determine when the VM stops.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/delegate
func (v VZVirtualMachine) Delegate() VZVirtualMachineDelegate {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("delegate"))
	return VZVirtualMachineDelegateObjectFromID(rv)
}
func (v VZVirtualMachine) SetDelegate(value VZVirtualMachineDelegate) {
	objc.Send[struct{}](v.ID, objc.Sel("setDelegate:"), value)
}

// The list of configured console devices on the VM.
//
// # Discussion
// 
// Return an empty array if there are no console devices configured.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/consoleDevices
func (v VZVirtualMachine) ConsoleDevices() []VZConsoleDevice {
	rv := objc.Send[[]objc.ID](v.ID, objc.Sel("consoleDevices"))
	return objc.ConvertSlice(rv, func(id objc.ID) VZConsoleDevice {
		return VZConsoleDeviceFromID(id)
	})
}

// The array of devices that you use to adjust the amount of memory available
// to the guest system.
//
// # Discussion
// 
// If you included a [VZVirtioTraditionalMemoryBalloonDeviceConfiguration]
// object in the configuration of your VM, this property contains a
// corresponding [VZVirtioTraditionalMemoryBalloonDevice] object. Use that
// object to adjust the amount of memory available to the guest operating
// system.
// 
// If you didn’t include a memory balloon object in your configuration, this
// property contains an empty array.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/memoryBalloonDevices
func (v VZVirtualMachine) MemoryBalloonDevices() []VZMemoryBalloonDevice {
	rv := objc.Send[[]objc.ID](v.ID, objc.Sel("memoryBalloonDevices"))
	return objc.ConvertSlice(rv, func(id objc.ID) VZMemoryBalloonDevice {
		return VZMemoryBalloonDeviceFromID(id)
	})
}

// The list of configured network devices on the VM.
//
// # Discussion
// 
// Returns an empty array if there are no configured network devices.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/networkDevices
func (v VZVirtualMachine) NetworkDevices() []VZNetworkDevice {
	rv := objc.Send[[]objc.ID](v.ID, objc.Sel("networkDevices"))
	return objc.ConvertSlice(rv, func(id objc.ID) VZNetworkDevice {
		return VZNetworkDeviceFromID(id)
	})
}

// The array of socket devices that the VM configures for use ports in the
// guest VM.
//
// # Discussion
// 
// If you included a [VZVirtioSocketDeviceConfiguration] object in the
// configuration of your VM, this property contains a corresponding
// [VZVirtioSocketDevice] object. Use that object to configure the ports your
// VM makes available to the guest operating system.
// 
// If you didn’t include a socket device in your configuration, this
// property contains an empty array.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/socketDevices
func (v VZVirtualMachine) SocketDevices() []VZSocketDevice {
	rv := objc.Send[[]objc.ID](v.ID, objc.Sel("socketDevices"))
	return objc.ConvertSlice(rv, func(id objc.ID) VZSocketDevice {
		return VZSocketDeviceFromID(id)
	})
}

// The list of configured directory-sharing devices on the VM.
//
// # Discussion
// 
// Returns an empty array if there are no directory sharing devices associated
// with this virtual machine.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/directorySharingDevices
func (v VZVirtualMachine) DirectorySharingDevices() []VZDirectorySharingDevice {
	rv := objc.Send[[]objc.ID](v.ID, objc.Sel("directorySharingDevices"))
	return objc.ConvertSlice(rv, func(id objc.ID) VZDirectorySharingDevice {
		return VZDirectorySharingDeviceFromID(id)
	})
}

// The list of runtime USB controller objects.
//
// # Discussion
// 
// Return an empty array if there isn’t a configured controller USB.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/usbControllers
func (v VZVirtualMachine) UsbControllers() []VZUSBController {
	rv := objc.Send[[]objc.ID](v.ID, objc.Sel("usbControllers"))
	return objc.ConvertSlice(rv, func(id objc.ID) VZUSBController {
		return VZUSBControllerFromID(id)
	})
}

// The current execution state of the VM.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/state-swift.property
func (v VZVirtualMachine) State() VZVirtualMachineState {
	rv := objc.Send[VZVirtualMachineState](v.ID, objc.Sel("state"))
	return VZVirtualMachineState(rv)
}

// The list of configured graphics devices on the virtual machine.
//
// # Discussion
// 
// Returns an empty array if there are no graphics devices configured.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/graphicsDevices
func (v VZVirtualMachine) GraphicsDevices() []VZGraphicsDevice {
	rv := objc.Send[[]objc.ID](v.ID, objc.Sel("graphicsDevices"))
	return objc.ConvertSlice(rv, func(id objc.ID) VZGraphicsDevice {
		return VZGraphicsDeviceFromID(id)
	})
}

// A Boolean value that indicates whether you can start the VM.
//
// # Discussion
// 
// The value of this property is [true] when the VM is in a state that allows
// you to start it. Call the [start(completionHandler:)] method (Swift) or
// [StartWithCompletionHandler] method (Objective-C) to start the VM.
//
// [start(completionHandler:)]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/start(completionHandler:)
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/canStart
func (v VZVirtualMachine) CanStart() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("canStart"))
	return rv
}

// A Boolean value that indicates whether you can pause the VM.
//
// # Discussion
// 
// The value of this property is [true] when the VM is in a state that allows
// you to pause it. Call the [pause(completionHandler:)] method (Swift) or
// [PauseWithCompletionHandler] method (Objective-C) to pause the VM.
//
// [pause(completionHandler:)]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/pause(completionHandler:)
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/canPause
func (v VZVirtualMachine) CanPause() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("canPause"))
	return rv
}

// A Boolean value that indicates whether you can resume the VM.
//
// # Discussion
// 
// The value of this property is [true] when the VM is in a state that allows
// you to resume it. Call the [resume(completionHandler:)] method (Swift) or
// [ResumeWithCompletionHandler] method (Objective-C) to resume the VM
//
// [resume(completionHandler:)]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/resume(completionHandler:)
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/canResume
func (v VZVirtualMachine) CanResume() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("canResume"))
	return rv
}

// A Boolean value that indicates whether you can stop the VM.
//
// # Discussion
// 
// The value of this property is [true] when the VM is in a state that allows
// you to stop it. Call the [StopWithCompletionHandler] to stop the VM.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/canStop
func (v VZVirtualMachine) CanStop() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("canStop"))
	return rv
}

// A Boolean value that indicates whether you can ask the guest operating
// system to stop running.
//
// # Discussion
// 
// The value of this property is [true] when the VM is in a state that allows
// it to stop running. Call the [RequestStopWithError] method to ask the guest
// operating system to stop running.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/canRequestStop
func (v VZVirtualMachine) CanRequestStop() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("canRequestStop"))
	return rv
}

// The queue associated with this virtual machine.
//
// # Discussion
// 
// This property is a reference to the queue the framework used to create the
// virtual machine when initialized using [InitWithConfigurationQueue]. If not
// specified, the default is the main queue.
// 
// The property is accessible from any queue or actor.
// 
// Other properties or function calls on [VZVirtualMachine] must happen on
// this queue. The framework also invokes any completion handlers from
// asynchronous functions on this queue.
// 
// The following example shows use of the `VZVirtualMachine.Queue()` property
// to check to see if it’s possible to start a VM.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/queue
func (v VZVirtualMachine) Queue() dispatch.Queue {
	rv := objc.Send[uintptr](v.ID, objc.Sel("queue"))
	return dispatch.QueueFromHandle(rv)
}

// A Boolean value that indicates whether the system supports virtualization.
//
// # Discussion
// 
// If virtualization is unavailable on the current device, no configuration is
// valid. If you want to know more about why virtualization is unavailable,
// call the [ValidateWithError] method of [VZVirtualMachineConfiguration] and
// examine the returned error object.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/isSupported
func (_VZVirtualMachineClass VZVirtualMachineClass) Supported() bool {
	rv := objc.Send[bool](objc.ID(_VZVirtualMachineClass.class), objc.Sel("isSupported"))
	return rv
}

// Start is a synchronous wrapper around [VZVirtualMachine.StartWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (v VZVirtualMachine) Start(ctx context.Context) error {
	done := make(chan error, 1)
	v.StartWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// StartWithOptions is a synchronous wrapper around [VZVirtualMachine.StartWithOptionsCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (v VZVirtualMachine) StartWithOptions(ctx context.Context, options IVZVirtualMachineStartOptions) error {
	done := make(chan error, 1)
	v.StartWithOptionsCompletionHandler(options, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// Stop is a synchronous wrapper around [VZVirtualMachine.StopWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (v VZVirtualMachine) Stop(ctx context.Context) error {
	done := make(chan error, 1)
	v.StopWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// Pause is a synchronous wrapper around [VZVirtualMachine.PauseWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (v VZVirtualMachine) Pause(ctx context.Context) error {
	done := make(chan error, 1)
	v.PauseWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// Resume is a synchronous wrapper around [VZVirtualMachine.ResumeWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (v VZVirtualMachine) Resume(ctx context.Context) error {
	done := make(chan error, 1)
	v.ResumeWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SaveMachineStateToURL is a synchronous wrapper around [VZVirtualMachine.SaveMachineStateToURLCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (v VZVirtualMachine) SaveMachineStateToURL(ctx context.Context, saveFileURL foundation.INSURL) error {
	done := make(chan error, 1)
	v.SaveMachineStateToURLCompletionHandler(saveFileURL, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// RestoreMachineStateFromURL is a synchronous wrapper around [VZVirtualMachine.RestoreMachineStateFromURLCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (v VZVirtualMachine) RestoreMachineStateFromURL(ctx context.Context, saveFileURL foundation.INSURL) error {
	done := make(chan error, 1)
	v.RestoreMachineStateFromURLCompletionHandler(saveFileURL, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

