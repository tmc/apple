// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VNContoursObservation] class.
var (
	_VNContoursObservationClass     VNContoursObservationClass
	_VNContoursObservationClassOnce sync.Once
)

func getVNContoursObservationClass() VNContoursObservationClass {
	_VNContoursObservationClassOnce.Do(func() {
		_VNContoursObservationClass = VNContoursObservationClass{class: objc.GetClass("VNContoursObservation")}
	})
	return _VNContoursObservationClass
}

// GetVNContoursObservationClass returns the class object for VNContoursObservation.
func GetVNContoursObservationClass() VNContoursObservationClass {
	return getVNContoursObservationClass()
}

type VNContoursObservationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VNContoursObservationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNContoursObservationClass) Alloc() VNContoursObservation {
	rv := objc.Send[VNContoursObservation](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An object that represents the detected contours in an image.
//
// # Inspecting the Observation
//
//   - [VNContoursObservation.ContourCount]: The total number of detected contours.
//   - [VNContoursObservation.NormalizedPath]: The detected contours as a path object in normalized coordinates.
//   - [VNContoursObservation.TopLevelContours]: An array of contours that don’t have another contour enclosing them.
//   - [VNContoursObservation.TopLevelContourCount]: The total number of detected top-level contours.
//   - [VNContoursObservation.ContourAtIndexError]: Retrieves the contour object at the specified index, irrespective of hierarchy.
//   - [VNContoursObservation.ContourAtIndexPathError]: Retrieves the contour object at the specified index path.
//
// See: https://developer.apple.com/documentation/Vision/VNContoursObservation
type VNContoursObservation struct {
	VNObservation
}

// VNContoursObservationFromID constructs a [VNContoursObservation] from an objc.ID.
//
// An object that represents the detected contours in an image.
func VNContoursObservationFromID(id objc.ID) VNContoursObservation {
	return VNContoursObservation{VNObservation: VNObservationFromID(id)}
}

// NOTE: VNContoursObservation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNContoursObservation] class.
//
// # Inspecting the Observation
//
//   - [IVNContoursObservation.ContourCount]: The total number of detected contours.
//   - [IVNContoursObservation.NormalizedPath]: The detected contours as a path object in normalized coordinates.
//   - [IVNContoursObservation.TopLevelContours]: An array of contours that don’t have another contour enclosing them.
//   - [IVNContoursObservation.TopLevelContourCount]: The total number of detected top-level contours.
//   - [IVNContoursObservation.ContourAtIndexError]: Retrieves the contour object at the specified index, irrespective of hierarchy.
//   - [IVNContoursObservation.ContourAtIndexPathError]: Retrieves the contour object at the specified index path.
//
// See: https://developer.apple.com/documentation/Vision/VNContoursObservation
type IVNContoursObservation interface {
	IVNObservation

	// Topic: Inspecting the Observation

	// The total number of detected contours.
	ContourCount() int
	// The detected contours as a path object in normalized coordinates.
	NormalizedPath() coregraphics.CGPathRef
	// An array of contours that don’t have another contour enclosing them.
	TopLevelContours() []VNContour
	// The total number of detected top-level contours.
	TopLevelContourCount() int
	// Retrieves the contour object at the specified index, irrespective of hierarchy.
	ContourAtIndexError(contourIndex int) (IVNContour, error)
	// Retrieves the contour object at the specified index path.
	ContourAtIndexPathError(indexPath foundation.INSIndexPath) (IVNContour, error)

	// The results of the request to detect contours.
	Results() IVNContoursObservation
	SetResults(value IVNContoursObservation)
}

// Init initializes the instance.
func (c VNContoursObservation) Init() VNContoursObservation {
	rv := objc.Send[VNContoursObservation](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c VNContoursObservation) Autorelease() VNContoursObservation {
	rv := objc.Send[VNContoursObservation](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNContoursObservation creates a new VNContoursObservation instance.
func NewVNContoursObservation() VNContoursObservation {
	class := getVNContoursObservationClass()
	rv := objc.Send[VNContoursObservation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Retrieves the contour object at the specified index, irrespective of
// hierarchy.
//
// contourIndex: The index of the contour to retrieve. Valid values are in the range of 0 to
// [ContourCount] - 1.
//
// # Return Value
//
// The contour object at the specified index.
//
// See: https://developer.apple.com/documentation/Vision/VNContoursObservation/contour(at:)-9on0y
func (c VNContoursObservation) ContourAtIndexError(contourIndex int) (IVNContour, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](c.ID, objc.Sel("contourAtIndex:error:"), contourIndex, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VNContour{}, foundation.NSErrorFrom(errorPtr)
	}
	return VNContourFromID(rv), nil

}

// Retrieves the contour object at the specified index path.
//
// indexPath: The hierarchical index path to the contour.
//
// # Return Value
//
// The contour object at the specified index path.
//
// See: https://developer.apple.com/documentation/Vision/VNContoursObservation/contour(at:)-52odo
func (c VNContoursObservation) ContourAtIndexPathError(indexPath foundation.INSIndexPath) (IVNContour, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](c.ID, objc.Sel("contourAtIndexPath:error:"), indexPath, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VNContour{}, foundation.NSErrorFrom(errorPtr)
	}
	return VNContourFromID(rv), nil

}

// The total number of detected contours.
//
// # Discussion
//
// Use this value to determine the number of indices available for calling
// [ContourAtIndexError].
//
// See: https://developer.apple.com/documentation/Vision/VNContoursObservation/contourCount
func (c VNContoursObservation) ContourCount() int {
	rv := objc.Send[int](c.ID, objc.Sel("contourCount"))
	return rv
}

// The detected contours as a path object in normalized coordinates.
//
// See: https://developer.apple.com/documentation/Vision/VNContoursObservation/normalizedPath
func (c VNContoursObservation) NormalizedPath() coregraphics.CGPathRef {
	rv := objc.Send[coregraphics.CGPathRef](c.ID, objc.Sel("normalizedPath"))
	return coregraphics.CGPathRef(rv)
}

// An array of contours that don’t have another contour enclosing them.
//
// # Discussion
//
// This array constitutes the top of the contour hierarchy. You can iterate
// over each [VNContour] instance to determine its children.
//
// See: https://developer.apple.com/documentation/Vision/VNContoursObservation/topLevelContours
func (c VNContoursObservation) TopLevelContours() []VNContour {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("topLevelContours"))
	return objc.ConvertSlice(rv, func(id objc.ID) VNContour {
		return VNContourFromID(id)
	})
}

// The total number of detected top-level contours.
//
// See: https://developer.apple.com/documentation/Vision/VNContoursObservation/topLevelContourCount
func (c VNContoursObservation) TopLevelContourCount() int {
	rv := objc.Send[int](c.ID, objc.Sel("topLevelContourCount"))
	return rv
}

// The results of the request to detect contours.
//
// See: https://developer.apple.com/documentation/vision/vndetectcontoursrequest/results
func (c VNContoursObservation) Results() IVNContoursObservation {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("results"))
	return VNContoursObservationFromID(objc.ID(rv))
}
func (c VNContoursObservation) SetResults(value IVNContoursObservation) {
	objc.Send[struct{}](c.ID, objc.Sel("setResults:"), value)
}
