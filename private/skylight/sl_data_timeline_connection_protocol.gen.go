// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// SLDataTimelineConnection protocol.
//
// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineConnection
type SLDataTimelineConnection interface {
	objectivec.IObject

	// Close protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineConnection/close
	Close()

	// Connected protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineConnection/connected
	Connected() bool
}

// SLDataTimelineConnectionObject wraps an existing Objective-C object that conforms to the SLDataTimelineConnection protocol.
type SLDataTimelineConnectionObject struct {
	objectivec.Object
}

func (o SLDataTimelineConnectionObject) BaseObject() objectivec.Object {
	return o.Object
}

// SLDataTimelineConnectionObjectFromID constructs a [SLDataTimelineConnectionObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func SLDataTimelineConnectionObjectFromID(id objc.ID) SLDataTimelineConnectionObject {
	return SLDataTimelineConnectionObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineConnection/close
func (o SLDataTimelineConnectionObject) Close() {
	objc.Send[struct{}](o.ID, objc.Sel("close"))
}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineConnection/connected
func (o SLDataTimelineConnectionObject) Connected() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("connected"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineConnection/name
func (o SLDataTimelineConnectionObject) Name() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("name"))
	return objectivec.Object{ID: rv}
}
