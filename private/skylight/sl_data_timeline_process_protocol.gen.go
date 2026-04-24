// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// SLDataTimelineProcess protocol.
//
// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineProcess
type SLDataTimelineProcess interface {
	objectivec.IObject

	// Pid protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineProcess/pid
	Pid() int
}

// SLDataTimelineProcessObject wraps an existing Objective-C object that conforms to the SLDataTimelineProcess protocol.
type SLDataTimelineProcessObject struct {
	objectivec.Object
}

func (o SLDataTimelineProcessObject) BaseObject() objectivec.Object {
	return o.Object
}

// SLDataTimelineProcessObjectFromID constructs a [SLDataTimelineProcessObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func SLDataTimelineProcessObjectFromID(id objc.ID) SLDataTimelineProcessObject {
	return SLDataTimelineProcessObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineProcess/pid
func (o SLDataTimelineProcessObject) Pid() int {
	rv := objc.Send[int](o.ID, objc.Sel("pid"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineProcess/windowData
func (o SLDataTimelineProcessObject) WindowData() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("windowData"))
	return objectivec.Object{ID: rv}
}
