// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VNGeometryUtils] class.
var (
	_VNGeometryUtilsClass     VNGeometryUtilsClass
	_VNGeometryUtilsClassOnce sync.Once
)

func getVNGeometryUtilsClass() VNGeometryUtilsClass {
	_VNGeometryUtilsClassOnce.Do(func() {
		_VNGeometryUtilsClass = VNGeometryUtilsClass{class: objc.GetClass("VNGeometryUtils")}
	})
	return _VNGeometryUtilsClass
}

// GetVNGeometryUtilsClass returns the class object for VNGeometryUtils.
func GetVNGeometryUtilsClass() VNGeometryUtilsClass {
	return getVNGeometryUtilsClass()
}

type VNGeometryUtilsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VNGeometryUtilsClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNGeometryUtilsClass) Alloc() VNGeometryUtils {
	rv := objc.Send[VNGeometryUtils](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// Utility methods to determine the geometries of various Vision types.
//
// See: https://developer.apple.com/documentation/Vision/VNGeometryUtils
type VNGeometryUtils struct {
	objectivec.Object
}

// VNGeometryUtilsFromID constructs a [VNGeometryUtils] from an objc.ID.
//
// Utility methods to determine the geometries of various Vision types.
func VNGeometryUtilsFromID(id objc.ID) VNGeometryUtils {
	return VNGeometryUtils{objectivec.Object{ID: id}}
}
// NOTE: VNGeometryUtils adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNGeometryUtils] class.
//
// See: https://developer.apple.com/documentation/Vision/VNGeometryUtils
type IVNGeometryUtils interface {
	objectivec.IObject
}

// Init initializes the instance.
func (g VNGeometryUtils) Init() VNGeometryUtils {
	rv := objc.Send[VNGeometryUtils](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g VNGeometryUtils) Autorelease() VNGeometryUtils {
	rv := objc.Send[VNGeometryUtils](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNGeometryUtils creates a new VNGeometryUtils instance.
func NewVNGeometryUtils() VNGeometryUtils {
	class := getVNGeometryUtilsClass()
	rv := objc.Send[VNGeometryUtils](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Calculates a bounding circle for the specified contour object.
//
// contour: A contour around which to calculate the bounding circle.
//
// # Return Value
// 
// The bounding [VNCircle] object.
//
// See: https://developer.apple.com/documentation/Vision/VNGeometryUtils/boundingCircle(for:)-423ll
func (_VNGeometryUtilsClass VNGeometryUtilsClass) BoundingCircleForContourError(contour IVNContour) (VNCircle, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_VNGeometryUtilsClass.class), objc.Sel("boundingCircleForContour:error:"), contour, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VNCircle{}, foundation.NSErrorFrom(errorPtr)
	}
	return VNCircleFromID(rv), nil

}
// Calculates a bounding circle for the specified array of points.
//
// points: A collection of points around which to calculate the bounding circle.
//
// # Return Value
// 
// The bounding [VNCircle] object.
//
// See: https://developer.apple.com/documentation/Vision/VNGeometryUtils/boundingCircle(for:)-9dggv
func (_VNGeometryUtilsClass VNGeometryUtilsClass) BoundingCircleForPointsError(points []VNPoint) (VNCircle, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_VNGeometryUtilsClass.class), objc.Sel("boundingCircleForPoints:error:"), objectivec.IObjectSliceToNSArray(points), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VNCircle{}, foundation.NSErrorFrom(errorPtr)
	}
	return VNCircleFromID(rv), nil

}
// Calculates a bounding circle for the specified points.
//
// points: A collection of points around which to calculate the bounding circle.
//
// pointCount: The number of points in the `points` argument.
//
// points is a [simd.simd_float2].
//
// # Return Value
// 
// The bounding [VNCircle] object.
//
// See: https://developer.apple.com/documentation/Vision/VNGeometryUtils/boundingCircle(forSIMDPoints:pointCount:)
// points is a [simd.simd_float2].
func (_VNGeometryUtilsClass VNGeometryUtilsClass) BoundingCircleForSIMDPointsPointCountError(points objectivec.IObject, pointCount int) (VNCircle, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_VNGeometryUtilsClass.class), objc.Sel("boundingCircleForSIMDPoints:pointCount:error:"), points, pointCount, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VNCircle{}, foundation.NSErrorFrom(errorPtr)
	}
	return VNCircleFromID(rv), nil

}
// Calculates the area for the specified contour.
//
// area: The output parameter to populate with the calculated contour area.
//
// contour: The contour object for which to calculate the area.
//
// orientedArea: A Boolean value that indicates whether to calculate the signed area
// (positive for counterclockwise-oriented contours and negative for
// clockwise-oriented contours). If you specify [false], the returned area is
// always positive.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
//
// # Discussion
// 
// Attempting to calculate the area for a contour containing random points, or
// with self-crossing edges, produces undefined results.
//
// See: https://developer.apple.com/documentation/Vision/VNGeometryUtils/calculateArea(_:for:orientedArea:)
func (_VNGeometryUtilsClass VNGeometryUtilsClass) CalculateAreaForContourOrientedAreaError(contour VNContour, orientedArea bool) (float64, error) {
	var area float64
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_VNGeometryUtilsClass.class), objc.Sel("calculateArea:forContour:orientedArea:error:"), unsafe.Pointer(&area), contour, orientedArea, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0.0, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return 0.0, errors.New("calculateArea:forContour:orientedArea:error: returned NO with nil NSError")
	}
	return area, nil
}
// Calculates the perimeter of a closed contour.
//
// perimeter: The output parameter to populate with the calculated contour perimeter.
//
// contour: The contour object for which to calculate the perimeter.
//
// See: https://developer.apple.com/documentation/Vision/VNGeometryUtils/calculatePerimeter(_:for:)
func (_VNGeometryUtilsClass VNGeometryUtilsClass) CalculatePerimeterForContourError(contour VNContour) (float64, error) {
	var perimeter float64
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_VNGeometryUtilsClass.class), objc.Sel("calculatePerimeter:forContour:error:"), unsafe.Pointer(&perimeter), contour, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0.0, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return 0.0, errors.New("calculatePerimeter:forContour:error: returned NO with nil NSError")
	}
	return perimeter, nil
}

