// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// GTGPUTimelineProtocol protocol.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTGPUTimelineProtocol
type GTGPUTimelineProtocol interface {
	objectivec.IObject

	// ConsistentStateAchieved protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTGPUTimelineProtocol/consistentStateAchieved
	ConsistentStateAchieved() bool

	// IsMio protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTGPUTimelineProtocol/isMio
	IsMio() bool

	// MetalFXCallDuration protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTGPUTimelineProtocol/metalFXCallDuration:
	MetalFXCallDuration(duration uint64) uint64

	// ProfiledState protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTGPUTimelineProtocol/profiledState
	ProfiledState() uint32

	// Version protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTGPUTimelineProtocol/version
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

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTGPUTimelineProtocol/consistentStateAchieved
func (o GTGPUTimelineProtocolObject) ConsistentStateAchieved() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("consistentStateAchieved"))
	return rv
	}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTGPUTimelineProtocol/isMio
func (o GTGPUTimelineProtocolObject) IsMio() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isMio"))
	return rv
	}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTGPUTimelineProtocol/metalFXCallDuration:
func (o GTGPUTimelineProtocolObject) MetalFXCallDuration(duration uint64) uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("metalFXCallDuration:"), duration)
	return rv
	}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTGPUTimelineProtocol/profiledState
func (o GTGPUTimelineProtocolObject) ProfiledState() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("profiledState"))
	return rv
	}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTGPUTimelineProtocol/version
func (o GTGPUTimelineProtocolObject) Version() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("version"))
	return rv
	}

