// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// VZGraphicsDisplayObserver protocol.
//
// See: https://developer.apple.com/documentation/Virtualization/VZGraphicsDisplayObserver
type VZGraphicsDisplayObserver interface {
	objectivec.IObject
}

// VZGraphicsDisplayObserverObject wraps an existing Objective-C object that conforms to the VZGraphicsDisplayObserver protocol.
type VZGraphicsDisplayObserverObject struct {
	objectivec.Object
}
func (o VZGraphicsDisplayObserverObject) BaseObject() objectivec.Object {
	return o.Object
}

// VZGraphicsDisplayObserverObjectFromID constructs a [VZGraphicsDisplayObserverObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func VZGraphicsDisplayObserverObjectFromID(id objc.ID) VZGraphicsDisplayObserverObject {
	return VZGraphicsDisplayObserverObject{
		Object: objectivec.ObjectFromID(id),
	}
}

//
// See: https://developer.apple.com/documentation/Virtualization/VZGraphicsDisplayObserver/displayDidBeginReconfiguration:
func (o VZGraphicsDisplayObserverObject) DisplayDidBeginReconfiguration(reconfiguration objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("displayDidBeginReconfiguration:"), reconfiguration)
	}
//
// See: https://developer.apple.com/documentation/Virtualization/VZGraphicsDisplayObserver/displayDidEndReconfiguration:
func (o VZGraphicsDisplayObserverObject) DisplayDidEndReconfiguration(reconfiguration objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("displayDidEndReconfiguration:"), reconfiguration)
	}

