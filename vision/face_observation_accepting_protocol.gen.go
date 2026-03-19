// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// An image analysis request that operates on face observations.
//
// See: https://developer.apple.com/documentation/Vision/VNFaceObservationAccepting
type VNFaceObservationAccepting interface {
	objectivec.IObject

	// An array of [VNFaceObservation](<doc://Vision/documentation/Vision/VNFaceObservation>) objects to process as part of the request.
	//
	// See: https://developer.apple.com/documentation/Vision/VNFaceObservationAccepting/inputFaceObservations
	InputFaceObservations() []VNFaceObservation

	// An array of [VNFaceObservation](<doc://Vision/documentation/Vision/VNFaceObservation>) objects to process as part of the request.
	//
	// See: https://developer.apple.com/documentation/Vision/VNFaceObservationAccepting/inputFaceObservations
	SetInputFaceObservations(value []VNFaceObservation)
}

// VNFaceObservationAcceptingObject wraps an existing Objective-C object that conforms to the VNFaceObservationAccepting protocol.
type VNFaceObservationAcceptingObject struct {
	objectivec.Object
}
func (o VNFaceObservationAcceptingObject) BaseObject() objectivec.Object {
	return o.Object
}

// VNFaceObservationAcceptingObjectFromID constructs a [VNFaceObservationAcceptingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func VNFaceObservationAcceptingObjectFromID(id objc.ID) VNFaceObservationAcceptingObject {
	return VNFaceObservationAcceptingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// An array of [VNFaceObservation] objects to process as part of the request.
//
// See: https://developer.apple.com/documentation/Vision/VNFaceObservationAccepting/inputFaceObservations
func (o VNFaceObservationAcceptingObject) InputFaceObservations() []VNFaceObservation {
	
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("inputFaceObservations"))
	return objc.ConvertSlice(rv, func(id objc.ID) VNFaceObservation {
		return VNFaceObservationFromID(id)
	})
	}

func (o VNFaceObservationAcceptingObject) SetInputFaceObservations(value []VNFaceObservation) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputFaceObservations:"), objectivec.IObjectSliceToNSArray(value))
}

