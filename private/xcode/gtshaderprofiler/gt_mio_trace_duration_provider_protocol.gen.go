// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// GTMioTraceDurationProvider protocol.
type GTMioTraceDurationProvider interface {
	objectivec.IObject

	// KickDurationForEncoder protocol.
	KickDurationForEncoder(encoder uint32) uint64

	// KickDurationForEncoderDataMaster protocol.
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

func (o GTMioTraceDurationProviderObject) KickDurationForEncoder(encoder uint32) uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("kickDurationForEncoder:"), encoder)
	return rv
}
func (o GTMioTraceDurationProviderObject) KickDurationForEncoderDataMaster(encoder uint32, master uint16) uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("kickDurationForEncoder:dataMaster:"), encoder, master)
	return rv
}
