// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// GTMioTraceDurationProvider protocol.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDurationProvider
type GTMioTraceDurationProvider interface {
	objectivec.IObject

	// KickDurationForEncoder protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDurationProvider/kickDurationForEncoder:
	KickDurationForEncoder(encoder uint32) uint64

	// KickDurationForEncoderDataMaster protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDurationProvider/kickDurationForEncoder:dataMaster:
	KickDurationForEncoderDataMaster(encoder uint32, master uint16) uint64
}

// GTMioTraceDurationProviderObject wraps an existing Objective-C object that conforms to the GTMioTraceDurationProvider protocol.
type GTMioTraceDurationProviderObject struct {
	objectivec.Object
}

func (o GTMioTraceDurationProviderObject) BaseObject() objectivec.Object {
	return o.Object
}

// GTMioTraceDurationProviderObjectFromID constructs a [GTMioTraceDurationProviderObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func GTMioTraceDurationProviderObjectFromID(id objc.ID) GTMioTraceDurationProviderObject {
	return GTMioTraceDurationProviderObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDurationProvider/kickDurationForEncoder:
func (o GTMioTraceDurationProviderObject) KickDurationForEncoder(encoder uint32) uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("kickDurationForEncoder:"), encoder)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDurationProvider/kickDurationForEncoder:dataMaster:
func (o GTMioTraceDurationProviderObject) KickDurationForEncoderDataMaster(encoder uint32, master uint16) uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("kickDurationForEncoder:dataMaster:"), encoder, master)
	return rv
}
