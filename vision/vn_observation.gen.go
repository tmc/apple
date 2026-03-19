// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VNObservation] class.
var (
	_VNObservationClass     VNObservationClass
	_VNObservationClassOnce sync.Once
)

func getVNObservationClass() VNObservationClass {
	_VNObservationClassOnce.Do(func() {
		_VNObservationClass = VNObservationClass{class: objc.GetClass("VNObservation")}
	})
	return _VNObservationClass
}

// GetVNObservationClass returns the class object for VNObservation.
func GetVNObservationClass() VNObservationClass {
	return getVNObservationClass()
}

type VNObservationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNObservationClass) Alloc() VNObservation {
	rv := objc.Send[VNObservation](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// The abstract superclass for analysis results.
//
// # Overview
// 
// Observations resulting from Vision image analysis requests inherit from
// this abstract base class. Don’t use this abstract superclass directly.
//
// # Tracking Observations
//
//   - [VNObservation.Uuid]: A unique identifier assigned to the Vision observation.
//
// # Evaluating Observations
//
//   - [VNObservation.TimeRange]: The time range of the reported observation.
//   - [VNObservation.Confidence]: The level of confidence in the observation’s accuracy.
//
// See: https://developer.apple.com/documentation/Vision/VNObservation
type VNObservation struct {
	objectivec.Object
}

// VNObservationFromID constructs a [VNObservation] from an objc.ID.
//
// The abstract superclass for analysis results.
func VNObservationFromID(id objc.ID) VNObservation {
	return VNObservation{objectivec.Object{ID: id}}
}
// NOTE: VNObservation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNObservation] class.
//
// # Tracking Observations
//
//   - [IVNObservation.Uuid]: A unique identifier assigned to the Vision observation.
//
// # Evaluating Observations
//
//   - [IVNObservation.TimeRange]: The time range of the reported observation.
//   - [IVNObservation.Confidence]: The level of confidence in the observation’s accuracy.
//
// See: https://developer.apple.com/documentation/Vision/VNObservation
type IVNObservation interface {
	objectivec.IObject
	VNRequestRevisionProviding

	// Topic: Tracking Observations

	// A unique identifier assigned to the Vision observation.
	Uuid() foundation.NSUUID

	// Topic: Evaluating Observations

	// The time range of the reported observation.
	TimeRange() objectivec.IObject
	// The level of confidence in the observation’s accuracy.
	Confidence() VNConfidence

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (o VNObservation) Init() VNObservation {
	rv := objc.Send[VNObservation](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o VNObservation) Autorelease() VNObservation {
	rv := objc.Send[VNObservation](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNObservation creates a new VNObservation instance.
func NewVNObservation() VNObservation {
	class := getVNObservationClass()
	rv := objc.Send[VNObservation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (o VNObservation) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](o.ID, objc.Sel("encodeWithCoder:"), coder)
}

// A unique identifier assigned to the Vision observation.
//
// See: https://developer.apple.com/documentation/Vision/VNObservation/uuid
func (o VNObservation) Uuid() foundation.NSUUID {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("uuid"))
	return foundation.NSUUIDFromID(objc.ID(rv))
}
// The time range of the reported observation.
//
// # Discussion
// 
// When evaluating a sequence of image buffers, use this property to determine
// each observation’s start time and duration. If a request doesn’t
// support time ranges, or the time range is unknown, the value of this
// property is [zero].
//
// [zero]: https://developer.apple.com/documentation/CoreMedia/CMTimeRange/zero
//
// See: https://developer.apple.com/documentation/Vision/VNObservation/timeRange
func (o VNObservation) TimeRange() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("timeRange"))
	return objectivec.Object{ID: rv}
}
// The level of confidence in the observation’s accuracy.
//
// # Discussion
// 
// The Vision framework normalizes this value to `[0.0, 1.0]` under most
// circumstances. A value of `0.0` indicates no confidence. A value of `1.0`
// indicates highest confidence, or the observation doesn’t support or
// assign meaning to confidence.
//
// See: https://developer.apple.com/documentation/Vision/VNObservation/confidence
func (o VNObservation) Confidence() VNConfidence {
	rv := objc.Send[VNConfidence](o.ID, objc.Sel("confidence"))
	return VNConfidence(rv)
}
// The revision of the [VNRequest] subclass used to generate the implementing
// object.
//
// See: https://developer.apple.com/documentation/Vision/VNRequestRevisionProviding/requestRevision
func (o VNObservation) RequestRevision() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("requestRevision"))
	return rv
}

			// Protocol methods for VNRequestRevisionProviding
			

