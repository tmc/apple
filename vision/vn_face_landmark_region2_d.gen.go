// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [VNFaceLandmarkRegion2D] class.
var (
	_VNFaceLandmarkRegion2DClass     VNFaceLandmarkRegion2DClass
	_VNFaceLandmarkRegion2DClassOnce sync.Once
)

func getVNFaceLandmarkRegion2DClass() VNFaceLandmarkRegion2DClass {
	_VNFaceLandmarkRegion2DClassOnce.Do(func() {
		_VNFaceLandmarkRegion2DClass = VNFaceLandmarkRegion2DClass{class: objc.GetClass("VNFaceLandmarkRegion2D")}
	})
	return _VNFaceLandmarkRegion2DClass
}

// GetVNFaceLandmarkRegion2DClass returns the class object for VNFaceLandmarkRegion2D.
func GetVNFaceLandmarkRegion2DClass() VNFaceLandmarkRegion2DClass {
	return getVNFaceLandmarkRegion2DClass()
}

type VNFaceLandmarkRegion2DClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNFaceLandmarkRegion2DClass) Alloc() VNFaceLandmarkRegion2D {
	rv := objc.Send[VNFaceLandmarkRegion2D](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// 2D geometry information for a specific facial feature.
//
// # Overview
// 
// This class represents the set of all facial landmark regions in 2D, exposed
// as properties.
//
// # Describing Region Points
//
//   - [VNFaceLandmarkRegion2D.PointsClassification]: An enumeration that describes how to interpret the points the region provides.
//
// See: https://developer.apple.com/documentation/Vision/VNFaceLandmarkRegion2D
type VNFaceLandmarkRegion2D struct {
	VNFaceLandmarkRegion
}

// VNFaceLandmarkRegion2DFromID constructs a [VNFaceLandmarkRegion2D] from an objc.ID.
//
// 2D geometry information for a specific facial feature.
func VNFaceLandmarkRegion2DFromID(id objc.ID) VNFaceLandmarkRegion2D {
	return VNFaceLandmarkRegion2D{VNFaceLandmarkRegion: VNFaceLandmarkRegionFromID(id)}
}
// NOTE: VNFaceLandmarkRegion2D adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNFaceLandmarkRegion2D] class.
//
// # Describing Region Points
//
//   - [IVNFaceLandmarkRegion2D.PointsClassification]: An enumeration that describes how to interpret the points the region provides.
//
// See: https://developer.apple.com/documentation/Vision/VNFaceLandmarkRegion2D
type IVNFaceLandmarkRegion2D interface {
	IVNFaceLandmarkRegion

	// Topic: Describing Region Points

	// An enumeration that describes how to interpret the points the region provides.
	PointsClassification() VNPointsClassification

	// A buffer in memory containing normalized landmark points.
	NormalizedPoints() unsafe.Pointer
	// An array of precision estimates for each landmark point.
	PrecisionEstimatesPerPoint() []foundation.NSNumber
	// A buffer in memory containing landmark points in the coordinate space of the specified image size.
	PointsInImageOfSize(imageSize corefoundation.CGSize) corefoundation.CGPoint
}

// Init initializes the instance.
func (f VNFaceLandmarkRegion2D) Init() VNFaceLandmarkRegion2D {
	rv := objc.Send[VNFaceLandmarkRegion2D](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f VNFaceLandmarkRegion2D) Autorelease() VNFaceLandmarkRegion2D {
	rv := objc.Send[VNFaceLandmarkRegion2D](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNFaceLandmarkRegion2D creates a new VNFaceLandmarkRegion2D instance.
func NewVNFaceLandmarkRegion2D() VNFaceLandmarkRegion2D {
	class := getVNFaceLandmarkRegion2DClass()
	rv := objc.Send[VNFaceLandmarkRegion2D](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A buffer in memory containing landmark points in the coordinate space of
// the specified image size.
//
// imageSize: The pixel dimensions of the image in which to present landmark points.
//
// # Return Value
// 
// A pointer to a buffer containing a [CGPoint] for each landmark detected in
// the image, expressed in the coordinate space of the specified image size.
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
//
// See: https://developer.apple.com/documentation/Vision/VNFaceLandmarkRegion2D/pointsInImageOfSize:
func (f VNFaceLandmarkRegion2D) PointsInImageOfSize(imageSize corefoundation.CGSize) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](f.ID, objc.Sel("pointsInImageOfSize:"), imageSize)
	return corefoundation.CGPoint(rv)
}

// An enumeration that describes how to interpret the points the region
// provides.
//
// See: https://developer.apple.com/documentation/Vision/VNFaceLandmarkRegion2D/pointsClassification
func (f VNFaceLandmarkRegion2D) PointsClassification() VNPointsClassification {
	rv := objc.Send[VNPointsClassification](f.ID, objc.Sel("pointsClassification"))
	return VNPointsClassification(rv)
}

// A buffer in memory containing normalized landmark points.
//
// # Discussion
// 
// This pointer points to the address of a buffer containing [CGPoint] structs
// representing landmark points.
// 
// The target object owns this buffer, which is guaranteed to exist as long as
// the corresponding [VNFaceLandmarkRegion2D] exists.
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
//
// See: https://developer.apple.com/documentation/Vision/VNFaceLandmarkRegion2D/normalizedPoints-1o38f
func (f VNFaceLandmarkRegion2D) NormalizedPoints() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f.ID, objc.Sel("normalizedPoints"))
	return rv
}

// An array of precision estimates for each landmark point.
//
// # Discussion
// 
// This property is only populated when you configure your
// [VNDetectFaceLandmarksRequest] object with
// [RequestFaceLandmarksConstellation76Points]. For other constellation types,
// this array is set to `nil`.
//
// See: https://developer.apple.com/documentation/Vision/VNFaceLandmarkRegion2D/precisionEstimatesPerPoint-3kx5a
func (f VNFaceLandmarkRegion2D) PrecisionEstimatesPerPoint() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](f.ID, objc.Sel("precisionEstimatesPerPoint"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}

