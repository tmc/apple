// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// _DASExtensionRunner protocol.
type DASExtensionRunner interface {
	objectivec.IObject

	// PrepareForActivity protocol.
	PrepareForActivity(activity objectivec.IObject) bool

	// Start protocol.
	Start() byte

	// Stop protocol.
	Stop()
}

// DASExtensionRunnerObject wraps an existing Objective-C object that conforms to the DASExtensionRunner protocol.
type DASExtensionRunnerObject struct {
	objectivec.Object
}

func (o DASExtensionRunnerObject) BaseObject() objectivec.Object {
	return o.Object
}

// DASExtensionRunnerObjectFromID constructs a [DASExtensionRunnerObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func DASExtensionRunnerObjectFromID(id objc.ID) DASExtensionRunnerObject {
	return DASExtensionRunnerObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o DASExtensionRunnerObject) PrepareForActivity(activity objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](o.ID, objc.Sel("prepareForActivity:"), activity)
	return rv
}
func (o DASExtensionRunnerObject) Start() byte {
	rv := objc.SendIfResponds[byte](o.ID, objc.Sel("start"))
	return rv
}
func (o DASExtensionRunnerObject) Stop() {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("stop"))
}
