// Code generated from Apple documentation for CoreWLAN. DO NOT EDIT.

package corewlan

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CWWiFiClient] class.
var (
	_CWWiFiClientClass     CWWiFiClientClass
	_CWWiFiClientClassOnce sync.Once
)

func getCWWiFiClientClass() CWWiFiClientClass {
	_CWWiFiClientClassOnce.Do(func() {
		_CWWiFiClientClass = CWWiFiClientClass{class: objc.GetClass("CWWiFiClient")}
	})
	return _CWWiFiClientClass
}

// GetCWWiFiClientClass returns the class object for CWWiFiClient.
func GetCWWiFiClientClass() CWWiFiClientClass {
	return getCWWiFiClientClass()
}

type CWWiFiClientClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CWWiFiClientClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CWWiFiClientClass) Alloc() CWWiFiClient {
	rv := objc.Send[CWWiFiClient](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A wrapper around the entire Wi-Fi subsystem that you use to access
// interfaces and set up event notifications.
//
// # Overview
//
// Wi-Fi client objects are heavy. Therefore, it’s more efficient to use a
// single, long-running client instance, rather than creating several
// short-lived instances. For convenience, you can use the singleton instance
// returned by the [CWWiFiClientClass.SharedWiFiClient] class method.
//
// Instead of instantiating [CWInterface] objects directly, use the ones
// provided by the instance methods of this class. For example, the
// [CWWiFiClient.Interface] method returns the default Wi-Fi interface.
//
// # Setting a Delegate
//
//   - [CWWiFiClient.Delegate]: An object that provides Wi-Fi event handling.
//   - [CWWiFiClient.SetDelegate]
//
// # Getting Interfaces
//
//   - [CWWiFiClient.Interface]: Returns the default Wi-Fi interface.
//   - [CWWiFiClient.InterfaceWithName]: Returns the Wi-Fi interface with the given name.
//   - [CWWiFiClient.Interfaces]: Returns all available Wi-Fi interfaces.
//
// # Monitoring Events
//
//   - [CWWiFiClient.StartMonitoringEventWithTypeError]: Register for specific Wi-Fi event notifications.
//   - [CWWiFiClient.StopMonitoringAllEventsAndReturnError]: Unregister for all Wi-Fi event notifications.
//   - [CWWiFiClient.StopMonitoringEventWithTypeError]: Unregister for specific Wi-Fi event notifications.
//
// # Instance Methods
//
//   - [CWWiFiClient.InterfaceNames]
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWWiFiClient
type CWWiFiClient struct {
	objectivec.Object
}

// CWWiFiClientFromID constructs a [CWWiFiClient] from an objc.ID.
//
// A wrapper around the entire Wi-Fi subsystem that you use to access
// interfaces and set up event notifications.
func CWWiFiClientFromID(id objc.ID) CWWiFiClient {
	return CWWiFiClient{objectivec.Object{ID: id}}
}

// NOTE: CWWiFiClient adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CWWiFiClient] class.
//
// # Setting a Delegate
//
//   - [ICWWiFiClient.Delegate]: An object that provides Wi-Fi event handling.
//   - [ICWWiFiClient.SetDelegate]
//
// # Getting Interfaces
//
//   - [ICWWiFiClient.Interface]: Returns the default Wi-Fi interface.
//   - [ICWWiFiClient.InterfaceWithName]: Returns the Wi-Fi interface with the given name.
//   - [ICWWiFiClient.Interfaces]: Returns all available Wi-Fi interfaces.
//
// # Monitoring Events
//
//   - [ICWWiFiClient.StartMonitoringEventWithTypeError]: Register for specific Wi-Fi event notifications.
//   - [ICWWiFiClient.StopMonitoringAllEventsAndReturnError]: Unregister for all Wi-Fi event notifications.
//   - [ICWWiFiClient.StopMonitoringEventWithTypeError]: Unregister for specific Wi-Fi event notifications.
//
// # Instance Methods
//
//   - [ICWWiFiClient.InterfaceNames]
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWWiFiClient
type ICWWiFiClient interface {
	objectivec.IObject

	// Topic: Setting a Delegate

	// An object that provides Wi-Fi event handling.
	Delegate() objectivec.IObject
	SetDelegate(value objectivec.IObject)

	// Topic: Getting Interfaces

	// Returns the default Wi-Fi interface.
	Interface() ICWInterface
	// Returns the Wi-Fi interface with the given name.
	InterfaceWithName(interfaceName string) ICWInterface
	// Returns all available Wi-Fi interfaces.
	Interfaces() []CWInterface

	// Topic: Monitoring Events

	// Register for specific Wi-Fi event notifications.
	StartMonitoringEventWithTypeError(type_ CWEventType) (bool, error)
	// Unregister for all Wi-Fi event notifications.
	StopMonitoringAllEventsAndReturnError() (bool, error)
	// Unregister for specific Wi-Fi event notifications.
	StopMonitoringEventWithTypeError(type_ CWEventType) (bool, error)

	// Topic: Instance Methods

	InterfaceNames() []string
}

// Init initializes the instance.
func (c CWWiFiClient) Init() CWWiFiClient {
	rv := objc.Send[CWWiFiClient](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CWWiFiClient) Autorelease() CWWiFiClient {
	rv := objc.Send[CWWiFiClient](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCWWiFiClient creates a new CWWiFiClient instance.
func NewCWWiFiClient() CWWiFiClient {
	class := getCWWiFiClientClass()
	rv := objc.Send[CWWiFiClient](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns the default Wi-Fi interface.
//
// # Return Value
//
// The [CWInterface] object that represents the default Wi-Fi interface.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWWiFiClient/interface()
func (c CWWiFiClient) Interface() ICWInterface {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("interface"))
	return CWInterfaceFromID(rv)
}

// Returns the Wi-Fi interface with the given name.
//
// interfaceName: The name of an available Wi-Fi interface. Use the [interfaceNames()] class
// method to obtain a list of valid interface names.
//
// # Return Value
//
// The [CWInterface] object bound to the given interface name, or the default
// interface if no name is specified.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWWiFiClient/interface(withName:)
//
// [interfaceNames()]: https://developer.apple.com/documentation/CoreWLAN/CWWiFiClient/interfaceNames()-swift.type.method
func (c CWWiFiClient) InterfaceWithName(interfaceName string) ICWInterface {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("interfaceWithName:"), objc.String(interfaceName))
	return CWInterfaceFromID(rv)
}

// Returns all available Wi-Fi interfaces.
//
// # Return Value
//
// An array of [CWInterface] objects, representing all of the available Wi-Fi
// interfaces in the system.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWWiFiClient/interfaces()
func (c CWWiFiClient) Interfaces() []CWInterface {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("interfaces"))
	return objc.ConvertSlice(rv, func(id objc.ID) CWInterface {
		return CWInterfaceFromID(id)
	})
}

// Register for specific Wi-Fi event notifications.
//
// type: The type of event notifications to register for. See [CWEventType] for a
// list of possible values.
//
// # Discussion
//
// After registering for notifications, when an event of the given type
// happens, the client sends an appropriate message to its delegate. See the
// [CWEventDelegate] protocol for the complete list of possible messages.
//
// Use the [CWWiFiClient.StopMonitoringEventWithTypeError] method when you
// want to stop receiving notifications of the given event type. Use
// [CWWiFiClient.StopMonitoringAllEventsAndReturnError] to stop receiving all
// notifications from a client.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWWiFiClient/startMonitoringEvent(with:)
//
// [CWEventType]: https://developer.apple.com/documentation/CoreWLAN/CWEventType
func (c CWWiFiClient) StartMonitoringEventWithTypeError(type_ CWEventType) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](c.ID, objc.Sel("startMonitoringEventWithType:error:"), type_, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("startMonitoringEventWithType:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Unregister for all Wi-Fi event notifications.
//
// # Discussion
//
// Use this method when you no longer want to receive any Wi-Fi notifications
// from the client.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWWiFiClient/stopMonitoringAllEvents()
func (c CWWiFiClient) StopMonitoringAllEventsAndReturnError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](c.ID, objc.Sel("stopMonitoringAllEventsAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("stopMonitoringAllEventsAndReturnError: returned NO with nil NSError")
	}
	return rv, nil

}

// Unregister for specific Wi-Fi event notifications.
//
// type: The type of event notifications to unregister for. See [CWEventType] for a
// list of possible values.
//
// # Discussion
//
// Use this method to indicate that the client should no longer send
// notifications for the given event type.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWWiFiClient/stopMonitoringEvent(with:)
//
// [CWEventType]: https://developer.apple.com/documentation/CoreWLAN/CWEventType
func (c CWWiFiClient) StopMonitoringEventWithTypeError(type_ CWEventType) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](c.ID, objc.Sel("stopMonitoringEventWithType:error:"), type_, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("stopMonitoringEventWithType:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreWLAN/CWWiFiClient/interfaceNames()-swift.method
func (c CWWiFiClient) InterfaceNames() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("interfaceNames"))
	return objc.ConvertSliceToStrings(rv)
}

// The shared Wi-Fi client object.
//
// # Discussion
//
// Each process spawns its own shared instance.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWWiFiClient/shared()
func (_CWWiFiClientClass CWWiFiClientClass) SharedWiFiClient() CWWiFiClient {
	rv := objc.Send[objc.ID](objc.ID(_CWWiFiClientClass.class), objc.Sel("sharedWiFiClient"))
	return CWWiFiClientFromID(rv)
}

// An object that provides Wi-Fi event handling.
//
// # Discussion
//
// When a client registers for Wi-Fi events with the
// [CWWiFiClient.StartMonitoringEventWithTypeError] method, the client’s
// delegate receives messages in response to Wi-Fi events. The delegate should
// adopt the [CWEventDelegate] protocol to receive these messages.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWWiFiClient/delegate
func (c CWWiFiClient) Delegate() objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("delegate"))
	return objectivec.Object{ID: rv}
}
func (c CWWiFiClient) SetDelegate(value objectivec.IObject) {
	objc.Send[struct{}](c.ID, objc.Sel("setDelegate:"), value)
}
