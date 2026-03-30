// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"fmt"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// _VZPluginConnectionDelegate protocol.
//
// See: https://developer.apple.com/documentation/Virtualization/_VZPluginConnectionDelegate
type VZPluginConnectionDelegate interface {
	objectivec.IObject

	// InvalidateConnection protocol.
	//
	// See: https://developer.apple.com/documentation/Virtualization/_VZPluginConnectionDelegate/invalidateConnection
	InvalidateConnection()
}

// VZPluginConnectionDelegateObject wraps an existing Objective-C object that conforms to the VZPluginConnectionDelegate protocol.
type VZPluginConnectionDelegateObject struct {
	objectivec.Object
}

func (o VZPluginConnectionDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// VZPluginConnectionDelegateObjectFromID constructs a [VZPluginConnectionDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func VZPluginConnectionDelegateObjectFromID(id objc.ID) VZPluginConnectionDelegateObject {
	return VZPluginConnectionDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/Virtualization/_VZPluginConnectionDelegate/handleConnectionError:
func (o VZPluginConnectionDelegateObject) HandleConnectionError(error_ objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("handleConnectionError:"), error_)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZPluginConnectionDelegate/invalidateConnection
func (o VZPluginConnectionDelegateObject) InvalidateConnection() {
	objc.Send[struct{}](o.ID, objc.Sel("invalidateConnection"))
}

// VZPluginConnectionDelegateConfig holds optional typed callbacks for [_VZPluginConnectionDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/virtualization/_vzpluginconnectiondelegate
type VZPluginConnectionDelegateConfig struct {

	// Other Methods
	InvalidateConnection func()
}

// NewVZPluginConnectionDelegate creates an Objective-C object implementing the [_VZPluginConnectionDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [VZPluginConnectionDelegateObject] satisfies the [VZPluginConnectionDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/virtualization/_vzpluginconnectiondelegate
func NewVZPluginConnectionDelegate(config VZPluginConnectionDelegateConfig) VZPluginConnectionDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("Go_VZPluginConnectionDelegate_%d", n)

	var methods []objc.MethodDef

	if config.InvalidateConnection != nil {
		fn := config.InvalidateConnection
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("invalidateConnection"),
			Fn: func(self objc.ID, _cmd objc.SEL) {
				fn()
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("_VZPluginConnectionDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewVZPluginConnectionDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return VZPluginConnectionDelegateObjectFromID(instance)
}
