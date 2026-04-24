// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// CPXLaunchServicesProviding protocol.
//
// See: https://developer.apple.com/documentation/SkyLight/CPXLaunchServicesProviding
type CPXLaunchServicesProviding interface {
	objectivec.IObject
}

// CPXLaunchServicesProvidingObject wraps an existing Objective-C object that conforms to the CPXLaunchServicesProviding protocol.
type CPXLaunchServicesProvidingObject struct {
	objectivec.Object
}

func (o CPXLaunchServicesProvidingObject) BaseObject() objectivec.Object {
	return o.Object
}

// CPXLaunchServicesProvidingObjectFromID constructs a [CPXLaunchServicesProvidingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CPXLaunchServicesProvidingObjectFromID(id objc.ID) CPXLaunchServicesProvidingObject {
	return CPXLaunchServicesProvidingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/SkyLight/CPXLaunchServicesProviding/launchServicesInterface
func (o CPXLaunchServicesProvidingObject) LaunchServicesInterface() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("launchServicesInterface"))
	return objectivec.Object{ID: rv}
}
