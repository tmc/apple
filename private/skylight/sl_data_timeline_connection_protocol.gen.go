// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// SLDataTimelineConnection protocol.
type SLDataTimelineConnection interface {
	objectivec.IObject

	// Close protocol.
	Close()

	// Connected protocol.
	Connected() bool

	// Name protocol.
	Name() objectivec.IObject
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

func (o SLDataTimelineConnectionObject) Close() {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("close"))
}
func (o SLDataTimelineConnectionObject) Connected() bool {
	rv := objc.SendIfResponds[bool](o.ID, objc.Sel("connected"))
	return rv
}
func (o SLDataTimelineConnectionObject) Name() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("name"))
	return objectivec.Object{ID: rv}
}
