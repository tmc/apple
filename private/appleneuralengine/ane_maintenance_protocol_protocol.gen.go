// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// _ANEMaintenanceProtocol protocol.
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEMaintenanceProtocol
type ANEMaintenanceProtocol interface {
	objectivec.IObject
}

// ANEMaintenanceProtocolObject wraps an existing Objective-C object that conforms to the ANEMaintenanceProtocol protocol.
type ANEMaintenanceProtocolObject struct {
	objectivec.Object
}
func (o ANEMaintenanceProtocolObject) BaseObject() objectivec.Object {
	return o.Object
}

// ANEMaintenanceProtocolObjectFromID constructs a [ANEMaintenanceProtocolObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func ANEMaintenanceProtocolObjectFromID(id objc.ID) ANEMaintenanceProtocolObject {
	return ANEMaintenanceProtocolObject{
		Object: objectivec.ObjectFromID(id),
	}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEMaintenanceProtocol/scheduleMaintenanceWithName:directoryPaths:
func (o ANEMaintenanceProtocolObject) ScheduleMaintenanceWithNameDirectoryPaths(name objectivec.IObject, paths objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("scheduleMaintenanceWithName:directoryPaths:"), name, paths)
	}

