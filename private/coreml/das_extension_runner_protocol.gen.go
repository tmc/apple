// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// _DASExtensionRunner protocol.
//
// See: https://developer.apple.com/documentation/CoreML/_DASExtensionRunner
type DASExtensionRunner interface {
	objectivec.IObject

	// Start protocol.
	//
	// See: https://developer.apple.com/documentation/CoreML/_DASExtensionRunner/start
	Start() byte

	// Stop protocol.
	//
	// See: https://developer.apple.com/documentation/CoreML/_DASExtensionRunner/stop
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

//
// See: https://developer.apple.com/documentation/CoreML/_DASExtensionRunner/prepareForActivity:
func (o DASExtensionRunnerObject) PrepareForActivity(activity objectivec.IObject) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("prepareForActivity:"), activity)
	return rv
	}
// See: https://developer.apple.com/documentation/CoreML/_DASExtensionRunner/start
func (o DASExtensionRunnerObject) Start() byte {
	rv := objc.Send[byte](o.ID, objc.Sel("start"))
	return rv
	}
// See: https://developer.apple.com/documentation/CoreML/_DASExtensionRunner/stop
func (o DASExtensionRunnerObject) Stop() {
	objc.Send[struct{}](o.ID, objc.Sel("stop"))
	}

