// Code generated from Apple documentation for CoreText. DO NOT EDIT.

package coretext

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// CTAdaptiveImageProviding protocol.
//
// See: https://developer.apple.com/documentation/CoreText/CTAdaptiveImageProviding
type CTAdaptiveImageProviding interface {
	objectivec.IObject
}

// CTAdaptiveImageProvidingObject wraps an existing Objective-C object that conforms to the CTAdaptiveImageProviding protocol.
type CTAdaptiveImageProvidingObject struct {
	objectivec.Object
}

func (o CTAdaptiveImageProvidingObject) BaseObject() objectivec.Object {
	return o.Object
}

// CTAdaptiveImageProvidingObjectFromID constructs a [CTAdaptiveImageProvidingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CTAdaptiveImageProvidingObjectFromID(id objc.ID) CTAdaptiveImageProvidingObject {
	return CTAdaptiveImageProvidingObject{
		Object: objectivec.ObjectFromID(id),
	}
}
