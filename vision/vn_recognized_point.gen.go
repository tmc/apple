// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [VNRecognizedPoint] class.
var (
	_VNRecognizedPointClass     VNRecognizedPointClass
	_VNRecognizedPointClassOnce sync.Once
)

func getVNRecognizedPointClass() VNRecognizedPointClass {
	_VNRecognizedPointClassOnce.Do(func() {
		_VNRecognizedPointClass = VNRecognizedPointClass{class: objc.GetClass("VNRecognizedPoint")}
	})
	return _VNRecognizedPointClass
}

// GetVNRecognizedPointClass returns the class object for VNRecognizedPoint.
func GetVNRecognizedPointClass() VNRecognizedPointClass {
	return getVNRecognizedPointClass()
}

type VNRecognizedPointClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNRecognizedPointClass) Alloc() VNRecognizedPoint {
	rv := objc.Send[VNRecognizedPoint](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a normalized point in an image, along with an
// identifier label and a confidence value.
//
// # Inspecting a Point
//
//   - [VNRecognizedPoint.Identifier]: The point’s identifier label.
//
// See: https://developer.apple.com/documentation/Vision/VNRecognizedPoint
type VNRecognizedPoint struct {
	VNDetectedPoint
}

// VNRecognizedPointFromID constructs a [VNRecognizedPoint] from an objc.ID.
//
// An object that represents a normalized point in an image, along with an
// identifier label and a confidence value.
func VNRecognizedPointFromID(id objc.ID) VNRecognizedPoint {
	return VNRecognizedPoint{VNDetectedPoint: VNDetectedPointFromID(id)}
}
// NOTE: VNRecognizedPoint adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNRecognizedPoint] class.
//
// # Inspecting a Point
//
//   - [IVNRecognizedPoint.Identifier]: The point’s identifier label.
//
// See: https://developer.apple.com/documentation/Vision/VNRecognizedPoint
type IVNRecognizedPoint interface {
	IVNDetectedPoint

	// Topic: Inspecting a Point

	// The point’s identifier label.
	Identifier() VNRecognizedPointKey
}

// Init initializes the instance.
func (r VNRecognizedPoint) Init() VNRecognizedPoint {
	rv := objc.Send[VNRecognizedPoint](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r VNRecognizedPoint) Autorelease() VNRecognizedPoint {
	rv := objc.Send[VNRecognizedPoint](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNRecognizedPoint creates a new VNRecognizedPoint instance.
func NewVNRecognizedPoint() VNRecognizedPoint {
	class := getVNRecognizedPointClass()
	rv := objc.Send[VNRecognizedPoint](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a point object from the specified Core Graphics point.
//
// location: The Core Graphics point.
//
// See: https://developer.apple.com/documentation/Vision/VNPoint/init(location:)
func NewRecognizedPointWithLocation(location corefoundation.CGPoint) VNRecognizedPoint {
	instance := getVNRecognizedPointClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLocation:"), location)
	return VNRecognizedPointFromID(rv)
}

// Creates a point object with the specified coordinates.
//
// x: The x-coordinate value.
//
// y: The y-coordinate value.
//
// See: https://developer.apple.com/documentation/Vision/VNPoint/init(x:y:)
func NewRecognizedPointWithXY(x float64, y float64) VNRecognizedPoint {
	instance := getVNRecognizedPointClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithX:y:"), x, y)
	return VNRecognizedPointFromID(rv)
}

// The point’s identifier label.
//
// See: https://developer.apple.com/documentation/Vision/VNRecognizedPoint/identifier
func (r VNRecognizedPoint) Identifier() VNRecognizedPointKey {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("identifier"))
	return VNRecognizedPointKey(foundation.NSStringFromID(rv).String())
}

