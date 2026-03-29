// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// CUTMetricLogger protocol.
//
// See: https://developer.apple.com/documentation/CoreML/CUTMetricLogger
type CUTMetricLogger interface {
	objectivec.IObject
}

// CUTMetricLoggerObject wraps an existing Objective-C object that conforms to the CUTMetricLogger protocol.
type CUTMetricLoggerObject struct {
	objectivec.Object
}
func (o CUTMetricLoggerObject) BaseObject() objectivec.Object {
	return o.Object
}

// CUTMetricLoggerObjectFromID constructs a [CUTMetricLoggerObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CUTMetricLoggerObjectFromID(id objc.ID) CUTMetricLoggerObject {
	return CUTMetricLoggerObject{
		Object: objectivec.ObjectFromID(id),
	}
}

//
// See: https://developer.apple.com/documentation/CoreML/CUTMetricLogger/logMetric:
func (o CUTMetricLoggerObject) LogMetric(metric objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("logMetric:"), metric)
	}

