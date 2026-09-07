// Code generated from Apple documentation for virtualization. DO NOT EDIT.

package virtualization

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// _VZRemoteSessionNotifierObserver protocol.
type VZRemoteSessionNotifierObserver interface {
	objectivec.IObject

	// DidDetectRemoteSessionEvent protocol.
	DidDetectRemoteSessionEvent()
}

// VZRemoteSessionNotifierObserverObject wraps an existing Objective-C object that conforms to the VZRemoteSessionNotifierObserver protocol.
type VZRemoteSessionNotifierObserverObject struct {
	objectivec.Object
}

func (o VZRemoteSessionNotifierObserverObject) BaseObject() objectivec.Object {
	return o.Object
}

// VZRemoteSessionNotifierObserverObjectFromID constructs a [VZRemoteSessionNotifierObserverObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func VZRemoteSessionNotifierObserverObjectFromID(id objc.ID) VZRemoteSessionNotifierObserverObject {
	return VZRemoteSessionNotifierObserverObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o VZRemoteSessionNotifierObserverObject) DidDetectRemoteSessionEvent() {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("didDetectRemoteSessionEvent"))
}
