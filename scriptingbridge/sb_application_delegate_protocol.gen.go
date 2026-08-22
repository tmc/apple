// Code generated from Apple documentation for ScriptingBridge. DO NOT EDIT.

package scriptingbridge

import (
	"unsafe"

	"github.com/tmc/apple/coreservices"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// This informal protocol defines a delegation method for handling Apple event errors that are sent from a target application to an [SBApplication](<https://developer.apple.com/documentation/ScriptingBridge/SBApplication>) object.
//
// See: https://developer.apple.com/documentation/ScriptingBridge/SBApplicationDelegate
type SBApplicationDelegate interface {
	objectivec.IObject

	// Sent by an [SBApplication] object when a target application returns an error Apple event.
	//
	// See: https://developer.apple.com/documentation/ScriptingBridge/SBApplicationDelegate/eventDidFail(_:withError:)
	EventDidFailWithError(event *coreservices.AEDesc) (objectivec.IObject, error)
}

// SBApplicationDelegateObject wraps an existing Objective-C object that conforms to the SBApplicationDelegate protocol.
type SBApplicationDelegateObject struct {
	objectivec.Object
}

func (o SBApplicationDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// SBApplicationDelegateObjectFromID constructs a [SBApplicationDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func SBApplicationDelegateObjectFromID(id objc.ID) SBApplicationDelegateObject {
	return SBApplicationDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Sent by an [SBApplication] object when a target application returns an
// error Apple event.
//
// event: A pointer to the Apple event sent to the target application causing the
// error.
//
// error: An object containing information about the error Apple event. Specific
// information may be included in the `useInfo` dictionary of the error
// object. The following table shows the possible keys for this dictionary.
//
// [Table data omitted]
//
// # Return Value
//
// If you return a result, it will become the result of the [sendEvent(_:)]
// that failed. Can be `nil`.
//
// See: https://developer.apple.com/documentation/ScriptingBridge/SBApplicationDelegate/eventDidFail(_:withError:)
//
// [sendEvent(_:)]: https://developer.apple.com/documentation/AppKit/NSApplication/sendEvent(_:)
func (o SBApplicationDelegateObject) EventDidFailWithError(event *coreservices.AEDesc) (objectivec.IObject, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("eventDidFail:withError:"), unsafe.Pointer(event))
	if err != nil {
		return nil, err
	}
	return objectivec.Object{ID: rv}, nil
}
