// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VNContour] class.
var (
	_VNContourClass     VNContourClass
	_VNContourClassOnce sync.Once
)

func getVNContourClass() VNContourClass {
	_VNContourClassOnce.Do(func() {
		_VNContourClass = VNContourClass{class: objc.GetClass("VNContour")}
	})
	return _VNContourClass
}

// GetVNContourClass returns the class object for VNContour.
func GetVNContourClass() VNContourClass {
	return getVNContourClass()
}

type VNContourClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNContourClass) Alloc() VNContour {
	rv := objc.Send[VNContour](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}







// A class that represents a detected contour in an image.
//
// # Inspecting the Contour
//
//   - [VNContour.AspectRatio]: The aspect ratio of the contour.
//   - [VNContour.IndexPath]: The contour object’s index path.
//   - [VNContour.NormalizedPath]: The contour object as a path in normalized coordinates.
//   - [VNContour.PointCount]: The contour’s number of points.
//   - [VNContour.PolygonApproximationWithEpsilonError]: Simplifies the contour to a polygon using a Ramer-Douglas-Peucker algorithm.
//
// # Accessing Child Contours
//
//   - [VNContour.ChildContourCount]: The total number of detected child contours.
//   - [VNContour.ChildContours]: An array of contours that this contour encloses.
//   - [VNContour.ChildContourAtIndexError]: Retrieves the child contour object at the specified index.
//
// See: https://developer.apple.com/documentation/Vision/VNContour
type VNContour struct {
	objectivec.Object
}

// VNContourFromID constructs a [VNContour] from an objc.ID.
//
// A class that represents a detected contour in an image.
func VNContourFromID(id objc.ID) VNContour {
	return VNContour{objectivec.Object{id}}
}
// NOTE: VNContour adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [VNContour] class.
//
// # Inspecting the Contour
//
//   - [IVNContour.AspectRatio]: The aspect ratio of the contour.
//   - [IVNContour.IndexPath]: The contour object’s index path.
//   - [IVNContour.NormalizedPath]: The contour object as a path in normalized coordinates.
//   - [IVNContour.PointCount]: The contour’s number of points.
//   - [IVNContour.PolygonApproximationWithEpsilonError]: Simplifies the contour to a polygon using a Ramer-Douglas-Peucker algorithm.
//
// # Accessing Child Contours
//
//   - [IVNContour.ChildContourCount]: The total number of detected child contours.
//   - [IVNContour.ChildContours]: An array of contours that this contour encloses.
//   - [IVNContour.ChildContourAtIndexError]: Retrieves the child contour object at the specified index.
//
// See: https://developer.apple.com/documentation/Vision/VNContour
type IVNContour interface {
	objectivec.IObject
	VNRequestRevisionProviding

	// Topic: Inspecting the Contour

	// The aspect ratio of the contour.
	AspectRatio() float32
	// The contour object’s index path.
	IndexPath() objc.ID
	// The contour object as a path in normalized coordinates.
	NormalizedPath() coregraphics.CGPathRef
	// The contour’s number of points.
	PointCount() int
	// Simplifies the contour to a polygon using a Ramer-Douglas-Peucker algorithm.
	PolygonApproximationWithEpsilonError(epsilon float32) (IVNContour, error)

	// Topic: Accessing Child Contours

	// The total number of detected child contours.
	ChildContourCount() int
	// An array of contours that this contour encloses.
	ChildContours() []VNContour
	// Retrieves the child contour object at the specified index.
	ChildContourAtIndexError(childContourIndex uint) (IVNContour, error)

	// The contour’s array of points in normalized coordinates.
	NormalizedPoints() objectivec.IObject
	// The total number of detected contours.
	ContourCount() int
	SetContourCount(value int)
	// The total number of detected top-level contours.
	TopLevelContourCount() int
	SetTopLevelContourCount(value int)
	// An array of contours that don’t have another contour enclosing them.
	TopLevelContours() IVNContour
	SetTopLevelContours(value IVNContour)
}





// Init initializes the instance.
func (c VNContour) Init() VNContour {
	rv := objc.Send[VNContour](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c VNContour) Autorelease() VNContour {
	rv := objc.Send[VNContour](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNContour creates a new VNContour instance.
func NewVNContour() VNContour {
	class := getVNContourClass()
	rv := objc.Send[VNContour](objc.ID(class.class), objc.Sel("new"))
	return rv
}










// Simplifies the contour to a polygon using a Ramer-Douglas-Peucker
// algorithm.
//
// epsilon: This parameter defines the distance threshold the algorithm uses. It
// preserves points whose perpendicular distance to the line segment they are
// on is greater than `epsilon`, and removes all others.
//
// # Return Value
// 
// A simplified polygon contour from the points of the original contour.
//
// See: https://developer.apple.com/documentation/Vision/VNContour/polygonApproximation(epsilon:)
func (c VNContour) PolygonApproximationWithEpsilonError(epsilon float32) (IVNContour, error) {
			var errorPtr objc.ID
	rv := objc.Send[objc.ID](c.ID, objc.Sel("polygonApproximationWithEpsilon:error:"), epsilon, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VNContour{}, foundation.NSErrorFrom(errorPtr)
	}
	return VNContourFromID(rv), nil

}

// Retrieves the child contour object at the specified index.
//
// childContourIndex: The child contour index value.
//
// # Return Value
// 
// The child contour object.
//
// See: https://developer.apple.com/documentation/Vision/VNContour/childContour(at:)
func (c VNContour) ChildContourAtIndexError(childContourIndex uint) (IVNContour, error) {
			var errorPtr objc.ID
	rv := objc.Send[objc.ID](c.ID, objc.Sel("childContourAtIndex:error:"), childContourIndex, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VNContour{}, foundation.NSErrorFrom(errorPtr)
	}
	return VNContourFromID(rv), nil

}












// The aspect ratio of the contour.
//
// # Discussion
// 
// The aspect ratio is the original image’s width divided by its height.
//
// See: https://developer.apple.com/documentation/Vision/VNContour/aspectRatio
func (c VNContour) AspectRatio() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("aspectRatio"))
	return rv
}



// The contour object’s index path.
//
// See: https://developer.apple.com/documentation/Vision/VNContour/indexPath
func (c VNContour) IndexPath() objc.ID {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("indexPath"))
	return rv
}



// The contour object as a path in normalized coordinates.
//
// See: https://developer.apple.com/documentation/Vision/VNContour/normalizedPath
func (c VNContour) NormalizedPath() coregraphics.CGPathRef {
	rv := objc.Send[coregraphics.CGPathRef](c.ID, objc.Sel("normalizedPath"))
	return coregraphics.CGPathRef(rv)
}



// The contour’s number of points.
//
// See: https://developer.apple.com/documentation/Vision/VNContour/pointCount
func (c VNContour) PointCount() int {
	rv := objc.Send[int](c.ID, objc.Sel("pointCount"))
	return rv
}



// The total number of detected child contours.
//
// See: https://developer.apple.com/documentation/Vision/VNContour/childContourCount
func (c VNContour) ChildContourCount() int {
	rv := objc.Send[int](c.ID, objc.Sel("childContourCount"))
	return rv
}



// An array of contours that this contour encloses.
//
// See: https://developer.apple.com/documentation/Vision/VNContour/childContours
func (c VNContour) ChildContours() []VNContour {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("childContours"))
	return objc.ConvertSlice(rv, func(id objc.ID) VNContour {
		return VNContourFromID(id)
	})
}



// The contour’s array of points in normalized coordinates.
//
// # Discussion
// 
// This property value provides the address of the buffer that contains the
// array of [CGPoint] values.
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
//
// See: https://developer.apple.com/documentation/Vision/VNContour/normalizedPoints-2orqj
func (c VNContour) NormalizedPoints() objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("normalizedPoints"))
	return objectivec.Object{ID: rv}
}



// The total number of detected contours.
//
// See: https://developer.apple.com/documentation/vision/vncontoursobservation/contourcount
func (c VNContour) ContourCount() int {
	rv := objc.Send[int](c.ID, objc.Sel("contourCount"))
	return rv
}
func (c VNContour) SetContourCount(value int) {
	objc.Send[struct{}](c.ID, objc.Sel("setContourCount:"), value)
}



// The total number of detected top-level contours.
//
// See: https://developer.apple.com/documentation/vision/vncontoursobservation/toplevelcontourcount
func (c VNContour) TopLevelContourCount() int {
	rv := objc.Send[int](c.ID, objc.Sel("topLevelContourCount"))
	return rv
}
func (c VNContour) SetTopLevelContourCount(value int) {
	objc.Send[struct{}](c.ID, objc.Sel("setTopLevelContourCount:"), value)
}



// An array of contours that don’t have another contour enclosing them.
//
// See: https://developer.apple.com/documentation/vision/vncontoursobservation/toplevelcontours
func (c VNContour) TopLevelContours() IVNContour {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("topLevelContours"))
	return VNContourFromID(objc.ID(rv))
}
func (c VNContour) SetTopLevelContours(value IVNContour) {
	objc.Send[struct{}](c.ID, objc.Sel("setTopLevelContours:"), value)
}



// The revision of the [VNRequest] subclass used to generate the implementing
// object.
//
// See: https://developer.apple.com/documentation/Vision/VNRequestRevisionProviding/requestRevision
func (c VNContour) RequestRevision() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("requestRevision"))
	return rv
}















			// Protocol methods for VNRequestRevisionProviding
			













