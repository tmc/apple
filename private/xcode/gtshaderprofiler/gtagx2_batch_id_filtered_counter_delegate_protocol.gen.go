// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// GTAGX2BatchIdFilteredCounterDelegate protocol.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2BatchIdFilteredCounterDelegate
type GTAGX2BatchIdFilteredCounterDelegate interface {
	objectivec.IObject
}

// GTAGX2BatchIdFilteredCounterDelegateObject wraps an existing Objective-C object that conforms to the GTAGX2BatchIdFilteredCounterDelegate protocol.
type GTAGX2BatchIdFilteredCounterDelegateObject struct {
	objectivec.Object
}
func (o GTAGX2BatchIdFilteredCounterDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// GTAGX2BatchIdFilteredCounterDelegateObjectFromID constructs a [GTAGX2BatchIdFilteredCounterDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func GTAGX2BatchIdFilteredCounterDelegateObjectFromID(id objc.ID) GTAGX2BatchIdFilteredCounterDelegateObject {
	return GTAGX2BatchIdFilteredCounterDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2BatchIdFilteredCounterDelegate/streamDataProcessorBatchIdFilteredCountersUpdated:observerInfo:
func (o GTAGX2BatchIdFilteredCounterDelegateObject) StreamDataProcessorBatchIdFilteredCountersUpdatedObserverInfo(updated objectivec.IObject, info objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("streamDataProcessorBatchIdFilteredCountersUpdated:observerInfo:"), updated, info)
	}

