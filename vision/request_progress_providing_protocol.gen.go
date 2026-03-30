// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol for providing progress information on long-running tasks in Vision.
//
// See: https://developer.apple.com/documentation/Vision/VNRequestProgressProviding
type VNRequestProgressProviding interface {
	objectivec.IObject

	// A block of code executed periodically during a Vision request to report progress on long-running tasks.
	//
	// See: https://developer.apple.com/documentation/Vision/VNRequestProgressProviding/progressHandler
	ProgressHandler() VNRequestProgressHandler

	// A Boolean set to true when a request can’t determine its progress in fractions completed.
	//
	// See: https://developer.apple.com/documentation/Vision/VNRequestProgressProviding/indeterminate
	Indeterminate() bool

	// A block of code executed periodically during a Vision request to report progress on long-running tasks.
	//
	// See: https://developer.apple.com/documentation/Vision/VNRequestProgressProviding/progressHandler
	SetProgressHandler(value VNRequestProgressHandler)
}

// VNRequestProgressProvidingObject wraps an existing Objective-C object that conforms to the VNRequestProgressProviding protocol.
type VNRequestProgressProvidingObject struct {
	objectivec.Object
}

func (o VNRequestProgressProvidingObject) BaseObject() objectivec.Object {
	return o.Object
}

// VNRequestProgressProvidingObjectFromID constructs a [VNRequestProgressProvidingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func VNRequestProgressProvidingObjectFromID(id objc.ID) VNRequestProgressProvidingObject {
	return VNRequestProgressProvidingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// A block of code executed periodically during a Vision request to report
// progress on long-running tasks.
//
// See: https://developer.apple.com/documentation/Vision/VNRequestProgressProviding/progressHandler
func (o VNRequestProgressProvidingObject) ProgressHandler() VNRequestProgressHandler {
	rv := objc.Send[VNRequestProgressHandler](o.ID, objc.Sel("progressHandler"))
	return rv
}

// A Boolean set to true when a request can’t determine its progress in
// fractions completed.
//
// See: https://developer.apple.com/documentation/Vision/VNRequestProgressProviding/indeterminate
func (o VNRequestProgressProvidingObject) Indeterminate() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("indeterminate"))
	return rv
}

// A block of code executed periodically during a Vision request to report
// progress on long-running tasks.
//
// # Discussion
//
// The progress handler is an optional method that allows clients of the
// request to report progress to the user or to display partial results as
// they become available. The Vision framework may call this handler on a
// different dispatch queue from the thread on which you initiated the
// original request, so ensure that your handler can execute asynchronously,
// in a thread-safe manner.
//
// See: https://developer.apple.com/documentation/Vision/VNRequestProgressProviding/progressHandler
func (o VNRequestProgressProvidingObject) SetProgressHandler(value VNRequestProgressHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("setProgressHandler:"), value)
}
