// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"fmt"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)
var _ = fmt.Sprintf

// Optional methods that you use to respond when a console port opens or closes in the virtual machine.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioConsoleDeviceDelegate
type VZVirtioConsoleDeviceDelegate interface {
	objectivec.IObject
}

// VZVirtioConsoleDeviceDelegateObject wraps an existing Objective-C object that conforms to the VZVirtioConsoleDeviceDelegate protocol.
type VZVirtioConsoleDeviceDelegateObject struct {
	objectivec.Object
}
func (o VZVirtioConsoleDeviceDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// VZVirtioConsoleDeviceDelegateObjectFromID constructs a [VZVirtioConsoleDeviceDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func VZVirtioConsoleDeviceDelegateObjectFromID(id objc.ID) VZVirtioConsoleDeviceDelegateObject {
	return VZVirtioConsoleDeviceDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Tells the delegate that the framework opened a console port.
//
// consoleDevice: The console port’s console device.
//
// consolePort: The [VZVirtioConsolePort] port that the framework opened.
//
// # Discussion
// 
// Be sure to process or flush any pending data from the [VZVirtioConsolePort]
// attachment before communicating with a new virtual machine process, or
// additional data might remain on the serial port from the previous session.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioConsoleDeviceDelegate/consoleDevice(_:didOpen:)
func (o VZVirtioConsoleDeviceDelegateObject) ConsoleDeviceDidOpenPort(consoleDevice IVZVirtioConsoleDevice, consolePort IVZVirtioConsolePort) {
	objc.Send[struct{}](o.ID, objc.Sel("consoleDevice:didOpenPort:"), consoleDevice, consolePort)
	}
// Tells the delegate that the framework closed a console port.
//
// consoleDevice: The console port’s console device.
//
// consolePort: The [VZVirtioConsolePort] port that the framework closed.
//
// # Discussion
// 
// Be sure to finish processing or flushing any remaining data from the
// [VZVirtioConsolePort] attachment after closing a port, or the additional
// data might remain on the serial port.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioConsoleDeviceDelegate/consoleDevice(_:didClose:)
func (o VZVirtioConsoleDeviceDelegateObject) ConsoleDeviceDidClosePort(consoleDevice IVZVirtioConsoleDevice, consolePort IVZVirtioConsolePort) {
	objc.Send[struct{}](o.ID, objc.Sel("consoleDevice:didClosePort:"), consoleDevice, consolePort)
	}

// VZVirtioConsoleDeviceDelegateConfig holds optional typed callbacks for [VZVirtioConsoleDeviceDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/virtualization/vzvirtioconsoledevicedelegate
type VZVirtioConsoleDeviceDelegateConfig struct {

	// Other Methods
	// ConsoleDeviceDidOpenPort — Tells the delegate that the framework opened a console port.
	ConsoleDeviceDidOpenPort func(consoleDevice VZVirtioConsoleDevice, consolePort VZVirtioConsolePort)
	// ConsoleDeviceDidClosePort — Tells the delegate that the framework closed a console port.
	ConsoleDeviceDidClosePort func(consoleDevice VZVirtioConsoleDevice, consolePort VZVirtioConsolePort)
}

// NewVZVirtioConsoleDeviceDelegate creates an Objective-C object implementing the [VZVirtioConsoleDeviceDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [VZVirtioConsoleDeviceDelegateObject] satisfies the [VZVirtioConsoleDeviceDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/virtualization/vzvirtioconsoledevicedelegate
func NewVZVirtioConsoleDeviceDelegate(config VZVirtioConsoleDeviceDelegateConfig) VZVirtioConsoleDeviceDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoVZVirtioConsoleDeviceDelegate_%d", n)

	var methods []objc.MethodDef

	if config.ConsoleDeviceDidOpenPort != nil {
		fn := config.ConsoleDeviceDidOpenPort
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("consoleDevice:didOpenPort:"),
			Fn: func(self objc.ID, _cmd objc.SEL, consoleDeviceID objc.ID, consolePortID objc.ID) {
				consoleDevice := VZVirtioConsoleDeviceFromID(consoleDeviceID)
				consolePort := VZVirtioConsolePortFromID(consolePortID)
				fn(consoleDevice, consolePort)
			},
		})
	}

	if config.ConsoleDeviceDidClosePort != nil {
		fn := config.ConsoleDeviceDidClosePort
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("consoleDevice:didClosePort:"),
			Fn: func(self objc.ID, _cmd objc.SEL, consoleDeviceID objc.ID, consolePortID objc.ID) {
				consoleDevice := VZVirtioConsoleDeviceFromID(consoleDeviceID)
				consolePort := VZVirtioConsolePortFromID(consolePortID)
				fn(consoleDevice, consolePort)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("VZVirtioConsoleDeviceDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewVZVirtioConsoleDeviceDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return VZVirtioConsoleDeviceDelegateObjectFromID(instance)
}

