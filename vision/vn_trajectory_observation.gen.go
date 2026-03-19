// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VNTrajectoryObservation] class.
var (
	_VNTrajectoryObservationClass     VNTrajectoryObservationClass
	_VNTrajectoryObservationClassOnce sync.Once
)

func getVNTrajectoryObservationClass() VNTrajectoryObservationClass {
	_VNTrajectoryObservationClassOnce.Do(func() {
		_VNTrajectoryObservationClass = VNTrajectoryObservationClass{class: objc.GetClass("VNTrajectoryObservation")}
	})
	return _VNTrajectoryObservationClass
}

// GetVNTrajectoryObservationClass returns the class object for VNTrajectoryObservation.
func GetVNTrajectoryObservationClass() VNTrajectoryObservationClass {
	return getVNTrajectoryObservationClass()
}

type VNTrajectoryObservationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNTrajectoryObservationClass) Alloc() VNTrajectoryObservation {
	rv := objc.Send[VNTrajectoryObservation](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An observation that describes a detected trajectory.
//
// # Evaluating an Observation
//
//   - [VNTrajectoryObservation.DetectedPoints]: The centroid points of the detected contour along the trajectory.
//   - [VNTrajectoryObservation.ProjectedPoints]: The centroids of the calculated trajectory from the detected points.
//   - [VNTrajectoryObservation.EquationCoefficients]: The coefficients of the parabolic equation.
//   - [VNTrajectoryObservation.MovingAverageRadius]: The moving average radius of the object the request is tracking.
//
// See: https://developer.apple.com/documentation/Vision/VNTrajectoryObservation
type VNTrajectoryObservation struct {
	VNObservation
}

// VNTrajectoryObservationFromID constructs a [VNTrajectoryObservation] from an objc.ID.
//
// An observation that describes a detected trajectory.
func VNTrajectoryObservationFromID(id objc.ID) VNTrajectoryObservation {
	return VNTrajectoryObservation{VNObservation: VNObservationFromID(id)}
}
// NOTE: VNTrajectoryObservation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNTrajectoryObservation] class.
//
// # Evaluating an Observation
//
//   - [IVNTrajectoryObservation.DetectedPoints]: The centroid points of the detected contour along the trajectory.
//   - [IVNTrajectoryObservation.ProjectedPoints]: The centroids of the calculated trajectory from the detected points.
//   - [IVNTrajectoryObservation.EquationCoefficients]: The coefficients of the parabolic equation.
//   - [IVNTrajectoryObservation.MovingAverageRadius]: The moving average radius of the object the request is tracking.
//
// See: https://developer.apple.com/documentation/Vision/VNTrajectoryObservation
type IVNTrajectoryObservation interface {
	IVNObservation

	// Topic: Evaluating an Observation

	// The centroid points of the detected contour along the trajectory.
	DetectedPoints() []VNPoint
	// The centroids of the calculated trajectory from the detected points.
	ProjectedPoints() []VNPoint
	// The coefficients of the parabolic equation.
	EquationCoefficients() objectivec.IObject
	// The moving average radius of the object the request is tracking.
	MovingAverageRadius() float64

	// The array of detected trajectory observations.
	Results() IVNTrajectoryObservation
	SetResults(value IVNTrajectoryObservation)
}

// Init initializes the instance.
func (t VNTrajectoryObservation) Init() VNTrajectoryObservation {
	rv := objc.Send[VNTrajectoryObservation](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t VNTrajectoryObservation) Autorelease() VNTrajectoryObservation {
	rv := objc.Send[VNTrajectoryObservation](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNTrajectoryObservation creates a new VNTrajectoryObservation instance.
func NewVNTrajectoryObservation() VNTrajectoryObservation {
	class := getVNTrajectoryObservationClass()
	rv := objc.Send[VNTrajectoryObservation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The centroid points of the detected contour along the trajectory.
//
// # Discussion
// 
// The detected points may differ slightly from the ideal trajectory because
// they fall within the allowed tolerance. The system limits the maximum
// number of points based on the maximum trajectory length set in the request.
//
// See: https://developer.apple.com/documentation/Vision/VNTrajectoryObservation/detectedPoints
func (t VNTrajectoryObservation) DetectedPoints() []VNPoint {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("detectedPoints"))
	return objc.ConvertSlice(rv, func(id objc.ID) VNPoint {
		return VNPointFromID(id)
	})
}
// The centroids of the calculated trajectory from the detected points.
//
// # Discussion
// 
// The projected points define the ideal trajectory described by the parabolic
// equation. The equation’s coefficients and the projected points of the
// detected trajectory get refined over time. The system limits the maximum
// number of cached points to the maximum points needed to describe the
// trajectory together with the parabolic equation.
//
// See: https://developer.apple.com/documentation/Vision/VNTrajectoryObservation/projectedPoints
func (t VNTrajectoryObservation) ProjectedPoints() []VNPoint {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("projectedPoints"))
	return objc.ConvertSlice(rv, func(id objc.ID) VNPoint {
		return VNPointFromID(id)
	})
}
// The coefficients of the parabolic equation.
//
// # Discussion
// 
// This equation describes the parabola on which the detected contour is
// traveling. The equation and the projected points get refined over time.
//
// See: https://developer.apple.com/documentation/Vision/VNTrajectoryObservation/equationCoefficients
func (t VNTrajectoryObservation) EquationCoefficients() objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("equationCoefficients"))
	return objectivec.Object{ID: rv}
}
// The moving average radius of the object the request is tracking.
//
// See: https://developer.apple.com/documentation/Vision/VNTrajectoryObservation/movingAverageRadius
func (t VNTrajectoryObservation) MovingAverageRadius() float64 {
	rv := objc.Send[float64](t.ID, objc.Sel("movingAverageRadius"))
	return rv
}
// The array of detected trajectory observations.
//
// See: https://developer.apple.com/documentation/vision/vndetecttrajectoriesrequest/results
func (t VNTrajectoryObservation) Results() IVNTrajectoryObservation {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("results"))
	return VNTrajectoryObservationFromID(objc.ID(rv))
}
func (t VNTrajectoryObservation) SetResults(value IVNTrajectoryObservation) {
	objc.Send[struct{}](t.ID, objc.Sel("setResults:"), value)
}

