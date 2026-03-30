// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// _VZMacBatterySourceObserver protocol.
//
// See: https://developer.apple.com/documentation/Virtualization/_VZMacBatterySourceObserver
type VZMacBatterySourceObserver interface {
	objectivec.IObject
}

// VZMacBatterySourceObserverObject wraps an existing Objective-C object that conforms to the VZMacBatterySourceObserver protocol.
type VZMacBatterySourceObserverObject struct {
	objectivec.Object
}

func (o VZMacBatterySourceObserverObject) BaseObject() objectivec.Object {
	return o.Object
}

// VZMacBatterySourceObserverObjectFromID constructs a [VZMacBatterySourceObserverObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func VZMacBatterySourceObserverObjectFromID(id objc.ID) VZMacBatterySourceObserverObject {
	return VZMacBatterySourceObserverObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMacBatterySourceObserver/batterySource:didUpdateCharge:
func (o VZMacBatterySourceObserverObject) BatterySourceDidUpdateCharge(source objectivec.IObject, charge float64) {
	objc.Send[struct{}](o.ID, objc.Sel("batterySource:didUpdateCharge:"), source, charge)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMacBatterySourceObserver/batterySource:didUpdateConnectivity:
func (o VZMacBatterySourceObserverObject) BatterySourceDidUpdateConnectivity(source objectivec.IObject, connectivity int64) {
	objc.Send[struct{}](o.ID, objc.Sel("batterySource:didUpdateConnectivity:"), source, connectivity)
}
