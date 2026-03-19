// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VNFaceLandmarks] class.
var (
	_VNFaceLandmarksClass     VNFaceLandmarksClass
	_VNFaceLandmarksClassOnce sync.Once
)

func getVNFaceLandmarksClass() VNFaceLandmarksClass {
	_VNFaceLandmarksClassOnce.Do(func() {
		_VNFaceLandmarksClass = VNFaceLandmarksClass{class: objc.GetClass("VNFaceLandmarks")}
	})
	return _VNFaceLandmarksClass
}

// GetVNFaceLandmarksClass returns the class object for VNFaceLandmarks.
func GetVNFaceLandmarksClass() VNFaceLandmarksClass {
	return getVNFaceLandmarksClass()
}

type VNFaceLandmarksClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNFaceLandmarksClass) Alloc() VNFaceLandmarks {
	rv := objc.Send[VNFaceLandmarks](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// The abstract superclass for containers of face landmark information.
//
// # Overview
// 
// This class represents the set of all detectable facial landmarks and
// regions, exposed as properties.
//
// # Determining Accuracy
//
//   - [VNFaceLandmarks.Confidence]: A confidence estimate for the detected landmarks.
//
// See: https://developer.apple.com/documentation/Vision/VNFaceLandmarks
type VNFaceLandmarks struct {
	objectivec.Object
}

// VNFaceLandmarksFromID constructs a [VNFaceLandmarks] from an objc.ID.
//
// The abstract superclass for containers of face landmark information.
func VNFaceLandmarksFromID(id objc.ID) VNFaceLandmarks {
	return VNFaceLandmarks{objectivec.Object{ID: id}}
}
// NOTE: VNFaceLandmarks adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNFaceLandmarks] class.
//
// # Determining Accuracy
//
//   - [IVNFaceLandmarks.Confidence]: A confidence estimate for the detected landmarks.
//
// See: https://developer.apple.com/documentation/Vision/VNFaceLandmarks
type IVNFaceLandmarks interface {
	objectivec.IObject
	VNRequestRevisionProviding

	// Topic: Determining Accuracy

	// A confidence estimate for the detected landmarks.
	Confidence() VNConfidence

	// The facial features of the detected face.
	Landmarks() IVNFaceLandmarks2D
	SetLandmarks(value IVNFaceLandmarks2D)
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (f VNFaceLandmarks) Init() VNFaceLandmarks {
	rv := objc.Send[VNFaceLandmarks](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f VNFaceLandmarks) Autorelease() VNFaceLandmarks {
	rv := objc.Send[VNFaceLandmarks](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNFaceLandmarks creates a new VNFaceLandmarks instance.
func NewVNFaceLandmarks() VNFaceLandmarks {
	class := getVNFaceLandmarksClass()
	rv := objc.Send[VNFaceLandmarks](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (f VNFaceLandmarks) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](f.ID, objc.Sel("encodeWithCoder:"), coder)
}

// A confidence estimate for the detected landmarks.
//
// # Discussion
// 
// A value of `0` indicates no confidence. A value of `1` indicates full
// confidence.
//
// See: https://developer.apple.com/documentation/Vision/VNFaceLandmarks/confidence
func (f VNFaceLandmarks) Confidence() VNConfidence {
	rv := objc.Send[VNConfidence](f.ID, objc.Sel("confidence"))
	return VNConfidence(rv)
}
// The facial features of the detected face.
//
// See: https://developer.apple.com/documentation/vision/vnfaceobservation/landmarks
func (f VNFaceLandmarks) Landmarks() IVNFaceLandmarks2D {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("landmarks"))
	return VNFaceLandmarks2DFromID(objc.ID(rv))
}
func (f VNFaceLandmarks) SetLandmarks(value IVNFaceLandmarks2D) {
	objc.Send[struct{}](f.ID, objc.Sel("setLandmarks:"), value)
}
// The revision of the [VNRequest] subclass used to generate the implementing
// object.
//
// See: https://developer.apple.com/documentation/Vision/VNRequestRevisionProviding/requestRevision
func (f VNFaceLandmarks) RequestRevision() uint {
	rv := objc.Send[uint](f.ID, objc.Sel("requestRevision"))
	return rv
}

			// Protocol methods for VNRequestRevisionProviding
			

