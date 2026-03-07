// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"fmt"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf


// The methods you use to respond to changes in the state of the VM.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineDelegate
type VZVirtualMachineDelegate interface {
	objectivec.IObject
}



// VZVirtualMachineDelegateObject wraps an existing Objective-C object that conforms to the VZVirtualMachineDelegate protocol.
type VZVirtualMachineDelegateObject struct {
	objectivec.Object
}
func (o VZVirtualMachineDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}



// VZVirtualMachineDelegateObjectFromID constructs a [VZVirtualMachineDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func VZVirtualMachineDelegateObjectFromID(id objc.ID) VZVirtualMachineDelegateObject {
	return VZVirtualMachineDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}




// Tells the delegate that the guest operating system stopped the VM.
//
// virtualMachine: The VM that called the delegate method.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineDelegate/guestDidStop(_:)

func (o VZVirtualMachineDelegateObject) GuestDidStopVirtualMachine(virtualMachine IVZVirtualMachine) {
	
	objc.Send[struct{}](o.ID, objc.Sel("guestDidStopVirtualMachine:"), virtualMachine)
	}

// Tells the delegate that the VM stopped because of an error.
//
// virtualMachine: The VM that called the delegate method.
//
// error: The error.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineDelegate/virtualMachine(_:didStopWithError:)

func (o VZVirtualMachineDelegateObject) VirtualMachineDidStopWithError(virtualMachine IVZVirtualMachine, error_ foundation.INSError) {
	
	objc.Send[struct{}](o.ID, objc.Sel("virtualMachine:didStopWithError:"), virtualMachine, error_)
	}

// The method the framework calls when an error causes a VM’s network
// attachment to disconnect.
//
// virtualMachine: The VM invoking the delegate method.
//
// networkDevice: The disconnected network device.
//
// error: The error that describes why the virtual machine disconnected the network
// device.
//
// # Discussion
// 
// The system invokes this method when the network interface fails to start,
// which results in the disconnection of the network attachment. This can
// happen in many situations such as initial boot, device reset, reboot, and
// so on. The system may invoke this method several times during a VM’s life
// cycle. After the system calls this method, the [Attachment] property is
// `nil`.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineDelegate/virtualMachine(_:networkDevice:attachmentWasDisconnectedWithError:)

func (o VZVirtualMachineDelegateObject) VirtualMachineNetworkDeviceAttachmentWasDisconnectedWithError(virtualMachine IVZVirtualMachine, networkDevice IVZNetworkDevice, error_ foundation.INSError) {
	
	objc.Send[struct{}](o.ID, objc.Sel("virtualMachine:networkDevice:attachmentWasDisconnectedWithError:"), virtualMachine, networkDevice, error_)
	}





// VZVirtualMachineDelegateConfig holds optional typed callbacks for [VZVirtualMachineDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/virtualization/vzvirtualmachinedelegate
type VZVirtualMachineDelegateConfig struct {

	// Stopping the VM
	// VirtualMachineDidStopWithError — Tells the delegate that the VM stopped because of an error.
	VirtualMachineDidStopWithError func(virtualMachine VZVirtualMachine, error_ foundation.NSError)

	// Responding to network device errors
	// VirtualMachineNetworkDeviceAttachmentWasDisconnectedWithError — The method the framework calls when an error causes a VM’s network attachment to disconnect.
	VirtualMachineNetworkDeviceAttachmentWasDisconnectedWithError func(virtualMachine VZVirtualMachine, networkDevice VZNetworkDevice, error_ foundation.NSError)

	// Other Methods
	// GuestDidStopVirtualMachine — Tells the delegate that the guest operating system stopped the VM.
	GuestDidStopVirtualMachine func(virtualMachine VZVirtualMachine)
}

// NewVZVirtualMachineDelegate creates an Objective-C object implementing the [VZVirtualMachineDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [VZVirtualMachineDelegateObject] satisfies the [VZVirtualMachineDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/virtualization/vzvirtualmachinedelegate
func NewVZVirtualMachineDelegate(config VZVirtualMachineDelegateConfig) VZVirtualMachineDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoVZVirtualMachineDelegate_%d", n)

	var methods []objc.MethodDef

	if config.GuestDidStopVirtualMachine != nil {
		fn := config.GuestDidStopVirtualMachine
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("guestDidStopVirtualMachine:"),
			Fn: func(self objc.ID, _cmd objc.SEL, virtualMachineID objc.ID) {
				virtualMachine := VZVirtualMachineFromID(virtualMachineID)
				fn(virtualMachine)
			},
		})
	}

	if config.VirtualMachineDidStopWithError != nil {
		fn := config.VirtualMachineDidStopWithError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("virtualMachine:didStopWithError:"),
			Fn: func(self objc.ID, _cmd objc.SEL, virtualMachineID objc.ID, error_ID objc.ID) {
				virtualMachine := VZVirtualMachineFromID(virtualMachineID)
				error_ := foundation.NSErrorFromID(error_ID)
				fn(virtualMachine, error_)
			},
		})
	}

	if config.VirtualMachineNetworkDeviceAttachmentWasDisconnectedWithError != nil {
		fn := config.VirtualMachineNetworkDeviceAttachmentWasDisconnectedWithError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("virtualMachine:networkDevice:attachmentWasDisconnectedWithError:"),
			Fn: func(self objc.ID, _cmd objc.SEL, virtualMachineID objc.ID, networkDeviceID objc.ID, error_ID objc.ID) {
				virtualMachine := VZVirtualMachineFromID(virtualMachineID)
				networkDevice := VZNetworkDeviceFromID(networkDeviceID)
				error_ := foundation.NSErrorFromID(error_ID)
				fn(virtualMachine, networkDevice, error_)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("VZVirtualMachineDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewVZVirtualMachineDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return VZVirtualMachineDelegateObjectFromID(instance)
}






