// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
)

// The class instance for the [VNDetectedPoint] class.
var (
	_VNDetectedPointClass     VNDetectedPointClass
	_VNDetectedPointClassOnce sync.Once
)

func getVNDetectedPointClass() VNDetectedPointClass {
	_VNDetectedPointClassOnce.Do(func() {
		_VNDetectedPointClass = VNDetectedPointClass{class: objc.GetClass("VNDetectedPoint")}
	})
	return _VNDetectedPointClass
}

// GetVNDetectedPointClass returns the class object for VNDetectedPoint.
func GetVNDetectedPointClass() VNDetectedPointClass {
	return getVNDetectedPointClass()
}

type VNDetectedPointClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNDetectedPointClass) Alloc() VNDetectedPoint {
	rv := objc.Send[VNDetectedPoint](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a normalized point in an image, along with a
// confidence value.
//
// # Inspecting a Point
//
//   - [VNDetectedPoint.Confidence]: A confidence score that indicates the detected point’s accuracy.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectedPoint
type VNDetectedPoint struct {
	VNPoint
}

// VNDetectedPointFromID constructs a [VNDetectedPoint] from an objc.ID.
//
// An object that represents a normalized point in an image, along with a
// confidence value.
func VNDetectedPointFromID(id objc.ID) VNDetectedPoint {
	return VNDetectedPoint{VNPoint: VNPointFromID(id)}
}
// NOTE: VNDetectedPoint adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNDetectedPoint] class.
//
// # Inspecting a Point
//
//   - [IVNDetectedPoint.Confidence]: A confidence score that indicates the detected point’s accuracy.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectedPoint
type IVNDetectedPoint interface {
	IVNPoint

	// Topic: Inspecting a Point

	// A confidence score that indicates the detected point’s accuracy.
	Confidence() VNConfidence
}

// Init initializes the instance.
func (d VNDetectedPoint) Init() VNDetectedPoint {
	rv := objc.Send[VNDetectedPoint](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d VNDetectedPoint) Autorelease() VNDetectedPoint {
	rv := objc.Send[VNDetectedPoint](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNDetectedPoint creates a new VNDetectedPoint instance.
func NewVNDetectedPoint() VNDetectedPoint {
	class := getVNDetectedPointClass()
	rv := objc.Send[VNDetectedPoint](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a point object from the specified Core Graphics point.
//
// location: The Core Graphics point.
//
// See: https://developer.apple.com/documentation/Vision/VNPoint/init(location:)
func NewDetectedPointWithLocation(location corefoundation.CGPoint) VNDetectedPoint {
	instance := getVNDetectedPointClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLocation:"), location)
	return VNDetectedPointFromID(rv)
}

// Creates a point object with the specified coordinates.
//
// x: The x-coordinate value.
//
// y: The y-coordinate value.
//
// See: https://developer.apple.com/documentation/Vision/VNPoint/init(x:y:)
func NewDetectedPointWithXY(x float64, y float64) VNDetectedPoint {
	instance := getVNDetectedPointClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithX:y:"), x, y)
	return VNDetectedPointFromID(rv)
}

// A confidence score that indicates the detected point’s accuracy.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectedPoint/confidence
func (d VNDetectedPoint) Confidence() VNConfidence {
	rv := objc.Send[VNConfidence](d.ID, objc.Sel("confidence"))
	return VNConfidence(rv)
}

