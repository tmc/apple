// Code generated from Apple documentation for SoundAnalysis. DO NOT EDIT.

package soundanalysis

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol that represents sound analysis results.
//
// See: https://developer.apple.com/documentation/SoundAnalysis/SNResult
type SNResult interface {
	objectivec.IObject
}

// SNResultObject wraps an existing Objective-C object that conforms to the SNResult protocol.
type SNResultObject struct {
	objectivec.Object
}

func (o SNResultObject) BaseObject() objectivec.Object {
	return o.Object
}

// SNResultObjectFromID constructs a [SNResultObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func SNResultObjectFromID(id objc.ID) SNResultObject {
	return SNResultObject{
		Object: objectivec.ObjectFromID(id),
	}
}
