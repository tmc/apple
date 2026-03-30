// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// An interface for objects that report progress using a single progress instance.
//
// See: https://developer.apple.com/documentation/Foundation/ProgressReporting
type NSProgressReporting interface {
	objectivec.IObject

	// The progress object returned by the class.
	//
	// See: https://developer.apple.com/documentation/Foundation/ProgressReporting/progress
	Progress() INSProgress
}

// NSProgressReportingObject wraps an existing Objective-C object that conforms to the NSProgressReporting protocol.
type NSProgressReportingObject struct {
	objectivec.Object
}

func (o NSProgressReportingObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSProgressReportingObjectFromID constructs a [NSProgressReportingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSProgressReportingObjectFromID(id objc.ID) NSProgressReportingObject {
	return NSProgressReportingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The progress object returned by the class.
//
// See: https://developer.apple.com/documentation/Foundation/ProgressReporting/progress
func (o NSProgressReportingObject) Progress() INSProgress {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("progress"))
	return NSProgressFromID(rv)
}
