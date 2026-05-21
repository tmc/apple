// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// GTGPUTimelineProtocol protocol.
type GTGPUTimelineProtocol interface {
	objectivec.IObject

	// ConsistentStateAchieved protocol.
	ConsistentStateAchieved() bool

	// IsMio protocol.
	IsMio() bool

	// MetalFXCallDuration protocol.
	MetalFXCallDuration(duration uint64) uint64

	// ProfiledState protocol.
	ProfiledState() uint32

	// Version protocol.
	Version() uint32
}

// GTGPUTimelineProtocolObject wraps an existing Objective-C object that conforms to the GTGPUTimelineProtocol protocol.
type GTGPUTimelineProtocolObject struct {
	objectivec.Object
}

func (o GTGPUTimelineProtocolObject) BaseObject() objectivec.Object {
	return o.Object
}

// GTGPUTimelineProtocolObjectFromID constructs a [GTGPUTimelineProtocolObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func GTGPUTimelineProtocolObjectFromID(id objc.ID) GTGPUTimelineProtocolObject {
	return GTGPUTimelineProtocolObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o GTGPUTimelineProtocolObject) ConsistentStateAchieved() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("consistentStateAchieved"))
	return rv
}
func (o GTGPUTimelineProtocolObject) IsMio() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isMio"))
	return rv
}
func (o GTGPUTimelineProtocolObject) MetalFXCallDuration(duration uint64) uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("metalFXCallDuration:"), duration)
	return rv
}
func (o GTGPUTimelineProtocolObject) ProfiledState() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("profiledState"))
	return rv
}
func (o GTGPUTimelineProtocolObject) Version() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("version"))
	return rv
}
