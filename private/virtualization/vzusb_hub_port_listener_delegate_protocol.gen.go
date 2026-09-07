// Code generated from Apple documentation for virtualization. DO NOT EDIT.

package virtualization

import (
	"fmt"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// _VZUSBHubPortListenerDelegate protocol.
type VZUSBHubPortListenerDelegate interface {
	objectivec.IObject

	// UsbDeviceDidConnectOnPort protocol.
	UsbDeviceDidConnectOnPort(device objectivec.IObject, port uint32)

	// UsbDeviceDidDisconnectOnPort protocol.
	UsbDeviceDidDisconnectOnPort(port uint32)
}

// VZUSBHubPortListenerDelegateObject wraps an existing Objective-C object that conforms to the VZUSBHubPortListenerDelegate protocol.
type VZUSBHubPortListenerDelegateObject struct {
	objectivec.Object
}

func (o VZUSBHubPortListenerDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// VZUSBHubPortListenerDelegateObjectFromID constructs a [VZUSBHubPortListenerDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func VZUSBHubPortListenerDelegateObjectFromID(id objc.ID) VZUSBHubPortListenerDelegateObject {
	return VZUSBHubPortListenerDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o VZUSBHubPortListenerDelegateObject) UsbDeviceDidConnectOnPort(device objectivec.IObject, port uint32) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("usbDevice:didConnectOnPort:"), device, port)
}
func (o VZUSBHubPortListenerDelegateObject) UsbDeviceDidDisconnectOnPort(port uint32) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("usbDeviceDidDisconnectOnPort:"), port)
}

// VZUSBHubPortListenerDelegateConfig holds optional typed callbacks for [_VZUSBHubPortListenerDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
type VZUSBHubPortListenerDelegateConfig struct {

	// Other Methods
	UsbDeviceDidDisconnectOnPort func(port uint32)
}

// NewVZUSBHubPortListenerDelegate creates an Objective-C object implementing the [_VZUSBHubPortListenerDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [VZUSBHubPortListenerDelegateObject] satisfies the [VZUSBHubPortListenerDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
func NewVZUSBHubPortListenerDelegate(config VZUSBHubPortListenerDelegateConfig) VZUSBHubPortListenerDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("Go_VZUSBHubPortListenerDelegate_%d", n)

	var methods []objc.MethodDef

	if config.UsbDeviceDidDisconnectOnPort != nil {
		fn := config.UsbDeviceDidDisconnectOnPort
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("usbDeviceDidDisconnectOnPort:"),
			Fn: func(self objc.ID, _cmd objc.SEL, port uint32) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("_VZUSBHubPortListenerDelegate", "usbDeviceDidDisconnectOnPort:")
					}
				}()
				fn(port)
				_delegateDone = true
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("_VZUSBHubPortListenerDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewVZUSBHubPortListenerDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return VZUSBHubPortListenerDelegateObjectFromID(instance)
}
