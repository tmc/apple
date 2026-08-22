// Code generated from Apple documentation for ModelIO. DO NOT EDIT.

package modelio

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// Adopt this protocol to provide information for use in automatic placement of light probes around a scene.
//
// See: https://developer.apple.com/documentation/ModelIO/MDLLightProbeIrradianceDataSource
type MDLLightProbeIrradianceDataSource interface {
	objectivec.IObject
}

// MDLLightProbeIrradianceDataSourceObject wraps an existing Objective-C object that conforms to the MDLLightProbeIrradianceDataSource protocol.
type MDLLightProbeIrradianceDataSourceObject struct {
	objectivec.Object
}

func (o MDLLightProbeIrradianceDataSourceObject) BaseObject() objectivec.Object {
	return o.Object
}

// MDLLightProbeIrradianceDataSourceObjectFromID constructs a [MDLLightProbeIrradianceDataSourceObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MDLLightProbeIrradianceDataSourceObjectFromID(id objc.ID) MDLLightProbeIrradianceDataSourceObject {
	return MDLLightProbeIrradianceDataSourceObject{
		Object: objectivec.ObjectFromID(id),
	}
}
