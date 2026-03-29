// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VNFaceLandmarkRegion] class.
var (
	_VNFaceLandmarkRegionClass     VNFaceLandmarkRegionClass
	_VNFaceLandmarkRegionClassOnce sync.Once
)

func getVNFaceLandmarkRegionClass() VNFaceLandmarkRegionClass {
	_VNFaceLandmarkRegionClassOnce.Do(func() {
		_VNFaceLandmarkRegionClass = VNFaceLandmarkRegionClass{class: objc.GetClass("VNFaceLandmarkRegion")}
	})
	return _VNFaceLandmarkRegionClass
}

// GetVNFaceLandmarkRegionClass returns the class object for VNFaceLandmarkRegion.
func GetVNFaceLandmarkRegionClass() VNFaceLandmarkRegionClass {
	return getVNFaceLandmarkRegionClass()
}

type VNFaceLandmarkRegionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VNFaceLandmarkRegionClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNFaceLandmarkRegionClass) Alloc() VNFaceLandmarkRegion {
	rv := objc.Send[VNFaceLandmarkRegion](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// The abstract superclass for information about a specific face landmark.
//
// # Instance Properties
//
//   - [VNFaceLandmarkRegion.PointCount]: The number of points in the face region.
//
// See: https://developer.apple.com/documentation/Vision/VNFaceLandmarkRegion
type VNFaceLandmarkRegion struct {
	objectivec.Object
}

// VNFaceLandmarkRegionFromID constructs a [VNFaceLandmarkRegion] from an objc.ID.
//
// The abstract superclass for information about a specific face landmark.
func VNFaceLandmarkRegionFromID(id objc.ID) VNFaceLandmarkRegion {
	return VNFaceLandmarkRegion{objectivec.Object{ID: id}}
}
// NOTE: VNFaceLandmarkRegion adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNFaceLandmarkRegion] class.
//
// # Instance Properties
//
//   - [IVNFaceLandmarkRegion.PointCount]: The number of points in the face region.
//
// See: https://developer.apple.com/documentation/Vision/VNFaceLandmarkRegion
type IVNFaceLandmarkRegion interface {
	objectivec.IObject
	VNRequestRevisionProviding

	// Topic: Instance Properties

	// The number of points in the face region.
	PointCount() uint

	// The facial features of the detected face.
	Landmarks() IVNFaceLandmarks2D
	SetLandmarks(value IVNFaceLandmarks2D)
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (f VNFaceLandmarkRegion) Init() VNFaceLandmarkRegion {
	rv := objc.Send[VNFaceLandmarkRegion](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f VNFaceLandmarkRegion) Autorelease() VNFaceLandmarkRegion {
	rv := objc.Send[VNFaceLandmarkRegion](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNFaceLandmarkRegion creates a new VNFaceLandmarkRegion instance.
func NewVNFaceLandmarkRegion() VNFaceLandmarkRegion {
	class := getVNFaceLandmarkRegionClass()
	rv := objc.Send[VNFaceLandmarkRegion](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (f VNFaceLandmarkRegion) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](f.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The number of points in the face region.
//
// # Discussion
// 
// The value is zero if no points for a region could be found.
//
// See: https://developer.apple.com/documentation/Vision/VNFaceLandmarkRegion/pointCount
func (f VNFaceLandmarkRegion) PointCount() uint {
	rv := objc.Send[uint](f.ID, objc.Sel("pointCount"))
	return rv
}
// The facial features of the detected face.
//
// See: https://developer.apple.com/documentation/vision/vnfaceobservation/landmarks
func (f VNFaceLandmarkRegion) Landmarks() IVNFaceLandmarks2D {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("landmarks"))
	return VNFaceLandmarks2DFromID(objc.ID(rv))
}
func (f VNFaceLandmarkRegion) SetLandmarks(value IVNFaceLandmarks2D) {
	objc.Send[struct{}](f.ID, objc.Sel("setLandmarks:"), value)
}
// The revision of the [VNRequest] subclass used to generate the implementing
// object.
//
// See: https://developer.apple.com/documentation/Vision/VNRequestRevisionProviding/requestRevision
func (f VNFaceLandmarkRegion) RequestRevision() uint {
	rv := objc.Send[uint](f.ID, objc.Sel("requestRevision"))
	return rv
}

			// Protocol methods for VNRequestRevisionProviding
			

