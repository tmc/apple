// Code generated from Apple documentation for CoreWLAN. DO NOT EDIT.

package corewlan

import (
	"fmt"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// The interface a Wi-Fi client object uses to notify its delegate about Wi-Fi events.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWEventDelegate
type CWEventDelegate interface {
	objectivec.IObject
}

// CWEventDelegateObject wraps an existing Objective-C object that conforms to the CWEventDelegate protocol.
type CWEventDelegateObject struct {
	objectivec.Object
}

func (o CWEventDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// CWEventDelegateObjectFromID constructs a [CWEventDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CWEventDelegateObjectFromID(id objc.ID) CWEventDelegateObject {
	return CWEventDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Tells the delegate that the current BSSID has changed.
//
// interfaceName: The name of the Wi-Fi interface on which the BSSID changed.
//
// # Discussion
//
// Register for BSSID change notifications by sending the
// [CWWiFiClient.StartMonitoringEventWithTypeError] message to a Wi-Fi client
// object with the [CWEventType.bssidDidChange] event type.
//
// Use the Wi-Fi interface’s [CWInterface.Bssid] method to query the current
// BSSID.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWEventDelegate/bssidDidChangeForWiFiInterface(withName:)
//
// [CWEventType.bssidDidChange]: https://developer.apple.com/documentation/CoreWLAN/CWEventType/bssidDidChange
func (o CWEventDelegateObject) BssidDidChangeForWiFiInterfaceWithName(interfaceName string) {
	objc.Send[struct{}](o.ID, objc.Sel("bssidDidChangeForWiFiInterfaceWithName:"), objc.String(interfaceName))
}

// Tells the delegate that the connection to the Wi-Fi subsystem is
// temporarily interrupted.
//
// # Discussion
//
// All event notifications for which the Wi-Fi client is registered are
// automatically re-registered when the connection resumes.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWEventDelegate/clientConnectionInterrupted()
func (o CWEventDelegateObject) ClientConnectionInterrupted() {
	objc.Send[struct{}](o.ID, objc.Sel("clientConnectionInterrupted"))
}

// Tells the delegate that the connection to the Wi-Fi subsystem is
// permanently invalidated.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWEventDelegate/clientConnectionInvalidated()
func (o CWEventDelegateObject) ClientConnectionInvalidated() {
	objc.Send[struct{}](o.ID, objc.Sel("clientConnectionInvalidated"))
}

// Tells the delegate that the currently adopted country code has changed.
//
// interfaceName: The name of the Wi-Fi interface on which the country code changed.
//
// # Discussion
//
// Register for country code change notifications by sending the
// [CWWiFiClient.StartMonitoringEventWithTypeError] message to a Wi-Fi client
// object with the [CWEventType.countryCodeDidChange] event type.
//
// Use the Wi-Fi interface’s [CWInterface.CountryCode] method to query the
// currently adopted country code.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWEventDelegate/countryCodeDidChangeForWiFiInterface(withName:)
//
// [CWEventType.countryCodeDidChange]: https://developer.apple.com/documentation/CoreWLAN/CWEventType/countryCodeDidChange
func (o CWEventDelegateObject) CountryCodeDidChangeForWiFiInterfaceWithName(interfaceName string) {
	objc.Send[struct{}](o.ID, objc.Sel("countryCodeDidChangeForWiFiInterfaceWithName:"), objc.String(interfaceName))
}

// Tells the delegate that the Wi-Fi link state changed.
//
// interfaceName: The name of the Wi-Fi interface on which the link state changed.
//
// # Discussion
//
// Register for link state change notifications by sending the
// [CWWiFiClient.StartMonitoringEventWithTypeError] message to a Wi-Fi client
// object with the [CWEventType.linkDidChange] event type.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWEventDelegate/linkDidChangeForWiFiInterface(withName:)
//
// [CWEventType.linkDidChange]: https://developer.apple.com/documentation/CoreWLAN/CWEventType/linkDidChange
func (o CWEventDelegateObject) LinkDidChangeForWiFiInterfaceWithName(interfaceName string) {
	objc.Send[struct{}](o.ID, objc.Sel("linkDidChangeForWiFiInterfaceWithName:"), objc.String(interfaceName))
}

// Tells the delegate that the link quality has changed.
//
// interfaceName: The name of the Wi-Fi interface on which the link quality changed.
//
// rssi: The receive signal strength indicator (RSSI) value measured in dBm for the
// currently associated network.
//
// transmitRate: The transmit rate measured in Mbps for the currently associated network.
//
// # Discussion
//
// Register for link quality change notifications by sending the
// [CWWiFiClient.StartMonitoringEventWithTypeError] message to a Wi-Fi client
// object with the [CWEventType.linkQualityDidChange] event type.
//
// Use the Wi-Fi interface’s [CWInterface.RssiValue] and
// [CWInterface.TransmitRate] methods to query the current RSSI and transmit
// rate, respectively.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWEventDelegate/linkQualityDidChangeForWiFiInterface(withName:rssi:transmitRate:)
//
// [CWEventType.linkQualityDidChange]: https://developer.apple.com/documentation/CoreWLAN/CWEventType/linkQualityDidChange
func (o CWEventDelegateObject) LinkQualityDidChangeForWiFiInterfaceWithNameRssiTransmitRate(interfaceName string, rssi int, transmitRate float64) {
	objc.Send[struct{}](o.ID, objc.Sel("linkQualityDidChangeForWiFiInterfaceWithName:rssi:transmitRate:"), objc.String(interfaceName), rssi, transmitRate)
}

// Tells the delegate that the operating mode has changed.
//
// interfaceName: The name of the Wi-Fi interface on which the operating mode changed.
//
// # Discussion
//
// Register for operating mode change notifications by sending the
// [CWWiFiClient.StartMonitoringEventWithTypeError] message to a Wi-Fi client
// object with the [CWEventType.modeDidChange] event type.
//
// Use the Wi-Fi interface’s [CWInterface.InterfaceMode] method to query the
// current operating mode.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWEventDelegate/modeDidChangeForWiFiInterface(withName:)
//
// [CWEventType.modeDidChange]: https://developer.apple.com/documentation/CoreWLAN/CWEventType/modeDidChange
func (o CWEventDelegateObject) ModeDidChangeForWiFiInterfaceWithName(interfaceName string) {
	objc.Send[struct{}](o.ID, objc.Sel("modeDidChangeForWiFiInterfaceWithName:"), objc.String(interfaceName))
}

// Tells the delegate that the Wi-Fi power state changed.
//
// interfaceName: The name of the Wi-Fi interface on which the power changed.
//
// # Discussion
//
// Register for power state change notifications by sending the
// [CWWiFiClient.StartMonitoringEventWithTypeError] message to a Wi-Fi client
// object with the [CWEventType.powerDidChange] event type.
//
// Use the Wi-Fi interface’s [CWInterface.PowerOn] method to query the
// current Wi-Fi power state.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWEventDelegate/powerStateDidChangeForWiFiInterface(withName:)
//
// [CWEventType.powerDidChange]: https://developer.apple.com/documentation/CoreWLAN/CWEventType/powerDidChange
func (o CWEventDelegateObject) PowerStateDidChangeForWiFiInterfaceWithName(interfaceName string) {
	objc.Send[struct{}](o.ID, objc.Sel("powerStateDidChangeForWiFiInterfaceWithName:"), objc.String(interfaceName))
}

// Tells the delegate that the Wi-Fi interface’s scan cache has been updated
// with new results.
//
// interfaceName: The name of the Wi-Fi interface for which the scan cache results have been
// updated.
//
// # Discussion
//
// Register for scan cache update notifications by sending the
// [CWWiFiClient.StartMonitoringEventWithTypeError] message to a Wi-Fi client
// object with the [CWEventType.scanCacheUpdated] event type.
//
// Use the Wi-Fi interface’s [CWInterface.CachedScanResults] method to query
// the scan cache results from the last scan.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWEventDelegate/scanCacheUpdatedForWiFiInterface(withName:)
//
// [CWEventType.scanCacheUpdated]: https://developer.apple.com/documentation/CoreWLAN/CWEventType/scanCacheUpdated
func (o CWEventDelegateObject) ScanCacheUpdatedForWiFiInterfaceWithName(interfaceName string) {
	objc.Send[struct{}](o.ID, objc.Sel("scanCacheUpdatedForWiFiInterfaceWithName:"), objc.String(interfaceName))
}

// Tells the delegate that the current SSID has changed.
//
// interfaceName: The name of the Wi-Fi interface on which the SSID changed.
//
// # Discussion
//
// Register for SSID change notifications by sending the
// [CWWiFiClient.StartMonitoringEventWithTypeError] message to a Wi-Fi client
// object with the [CWEventType.ssidDidChange] event type.
//
// Use the Wi-Fi interface’s [CWInterface.SsidData] or [CWInterface.Ssid]
// method to query the current SSID.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWEventDelegate/ssidDidChangeForWiFiInterface(withName:)
//
// [CWEventType.ssidDidChange]: https://developer.apple.com/documentation/CoreWLAN/CWEventType/ssidDidChange
func (o CWEventDelegateObject) SsidDidChangeForWiFiInterfaceWithName(interfaceName string) {
	objc.Send[struct{}](o.ID, objc.Sel("ssidDidChangeForWiFiInterfaceWithName:"), objc.String(interfaceName))
}

// CWEventDelegateConfig holds optional typed callbacks for [CWEventDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/corewlan/cweventdelegate
type CWEventDelegateConfig struct {

	// Instance Methods
	// ClientConnectionInterrupted — Tells the delegate that the connection to the Wi-Fi subsystem is temporarily interrupted.
	ClientConnectionInterrupted func()
	// ClientConnectionInvalidated — Tells the delegate that the connection to the Wi-Fi subsystem is permanently invalidated.
	ClientConnectionInvalidated func()
}

// NewCWEventDelegate creates an Objective-C object implementing the [CWEventDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [CWEventDelegateObject] satisfies the [CWEventDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/corewlan/cweventdelegate
func NewCWEventDelegate(config CWEventDelegateConfig) CWEventDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoCWEventDelegate_%d", n)

	var methods []objc.MethodDef

	if config.ClientConnectionInterrupted != nil {
		fn := config.ClientConnectionInterrupted
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("clientConnectionInterrupted"),
			Fn: func(self objc.ID, _cmd objc.SEL) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("CWEventDelegate", "clientConnectionInterrupted")
					}
				}()
				fn()
				_delegateDone = true
			},
		})
	}

	if config.ClientConnectionInvalidated != nil {
		fn := config.ClientConnectionInvalidated
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("clientConnectionInvalidated"),
			Fn: func(self objc.ID, _cmd objc.SEL) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("CWEventDelegate", "clientConnectionInvalidated")
					}
				}()
				fn()
				_delegateDone = true
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("CWEventDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewCWEventDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return CWEventDelegateObjectFromID(instance)
}
