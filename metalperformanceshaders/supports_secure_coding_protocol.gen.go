// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// supportsSecureCoding protocol.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationDataSource/supportsSecureCoding
type supportsSecureCoding interface {
	objectivec.IObject
}

// supportsSecureCodingObject wraps an existing Objective-C object that conforms to the supportsSecureCoding protocol.
type supportsSecureCodingObject struct {
	objectivec.Object
}

func (o supportsSecureCodingObject) BaseObject() objectivec.Object {
	return o.Object
}

// supportsSecureCodingObjectFromID constructs a [supportsSecureCodingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func supportsSecureCodingObjectFromID(id objc.ID) supportsSecureCodingObject {
	return supportsSecureCodingObject{
		Object: objectivec.ObjectFromID(id),
	}
}
