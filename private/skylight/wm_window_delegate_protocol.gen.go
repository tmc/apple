// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// WMWindowDelegate protocol.
type WMWindowDelegate interface {
	objectivec.IObject

	// WindowDidUpdateWithChangedProperties protocol.
	WindowDidUpdateWithChangedProperties(window objectivec.IObject, properties uint64)

	// WindowDidUpdateWithChangedServerProperties protocol.
	WindowDidUpdateWithChangedServerProperties(window objectivec.IObject, properties uint64)
}

// WMWindowDelegateObject wraps an existing Objective-C object that conforms to the WMWindowDelegate protocol.
type WMWindowDelegateObject struct {
	objectivec.Object
}

func (o WMWindowDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// WMWindowDelegateObjectFromID constructs a [WMWindowDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func WMWindowDelegateObjectFromID(id objc.ID) WMWindowDelegateObject {
	return WMWindowDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o WMWindowDelegateObject) WindowDidUpdateWithChangedProperties(window objectivec.IObject, properties uint64) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("window:didUpdateWithChangedProperties:"), window, properties)
}
func (o WMWindowDelegateObject) WindowDidUpdateWithChangedServerProperties(window objectivec.IObject, properties uint64) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("window:didUpdateWithChangedServerProperties:"), window, properties)
}
