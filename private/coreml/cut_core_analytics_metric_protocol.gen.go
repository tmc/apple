// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// CUTCoreAnalyticsMetric protocol.
//
// See: https://developer.apple.com/documentation/CoreML/CUTCoreAnalyticsMetric
type CUTCoreAnalyticsMetric interface {
	objectivec.IObject
}

// CUTCoreAnalyticsMetricObject wraps an existing Objective-C object that conforms to the CUTCoreAnalyticsMetric protocol.
type CUTCoreAnalyticsMetricObject struct {
	objectivec.Object
}
func (o CUTCoreAnalyticsMetricObject) BaseObject() objectivec.Object {
	return o.Object
}

// CUTCoreAnalyticsMetricObjectFromID constructs a [CUTCoreAnalyticsMetricObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CUTCoreAnalyticsMetricObjectFromID(id objc.ID) CUTCoreAnalyticsMetricObject {
	return CUTCoreAnalyticsMetricObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/CoreML/CUTCoreAnalyticsMetric/dictionaryRepresentation
func (o CUTCoreAnalyticsMetricObject) DictionaryRepresentation() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("dictionaryRepresentation"))
	return objectivec.Object{ID: rv}
	}

