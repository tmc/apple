// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol you implement to observe state changes in graphic displays.
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




// The method the framework calls when the reconfiguration operation has
// begun.
//
// display: The [VZGraphicsDisplay] whose state is changing.
//
// # Discussion
// 
// The framework issued a configuration change, such as a resize, and you can
// expect new frames with a new size or configuration.
// 
// The framework invokes this method on the VM’s queue.
//
// See: https://developer.apple.com/documentation/Virtualization/VZGraphicsDisplayObserver/displayDidBeginReconfiguration(_:)

func (o VZGraphicsDisplayObserverObject) DisplayDidBeginReconfiguration(display IVZGraphicsDisplay) {
	
	objc.Send[struct{}](o.ID, objc.Sel("displayDidBeginReconfiguration:"), display)
	}

// The method the framework calls when the reconfiguration operation ends.
//
// display: The [VZGraphicsDisplay] whose state is changing.
//
// # Discussion
// 
// Frame updates have arrived at the most recently requested display size and
// configuration.
// 
// The framework invokes this method on the VM’s queue.
//
// See: https://developer.apple.com/documentation/Virtualization/VZGraphicsDisplayObserver/displayDidEndReconfiguration(_:)

func (o VZGraphicsDisplayObserverObject) DisplayDidEndReconfiguration(display IVZGraphicsDisplay) {
	
	objc.Send[struct{}](o.ID, objc.Sel("displayDidEndReconfiguration:"), display)
	}







