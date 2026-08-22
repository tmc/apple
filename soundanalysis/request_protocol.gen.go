// Code generated from Apple documentation for SoundAnalysis. DO NOT EDIT.

package soundanalysis

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol that represents sound analysis requests.
//
// See: https://developer.apple.com/documentation/SoundAnalysis/SNRequest
type SNRequest interface {
	objectivec.IObject
}

// SNRequestObject wraps an existing Objective-C object that conforms to the SNRequest protocol.
type SNRequestObject struct {
	objectivec.Object
}

func (o SNRequestObject) BaseObject() objectivec.Object {
	return o.Object
}

// SNRequestObjectFromID constructs a [SNRequestObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func SNRequestObjectFromID(id objc.ID) SNRequestObject {
	return SNRequestObject{
		Object: objectivec.ObjectFromID(id),
	}
}
