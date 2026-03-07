// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VNImageHomographicAlignmentObservation] class.
var (
	_VNImageHomographicAlignmentObservationClass     VNImageHomographicAlignmentObservationClass
	_VNImageHomographicAlignmentObservationClassOnce sync.Once
)

func getVNImageHomographicAlignmentObservationClass() VNImageHomographicAlignmentObservationClass {
	_VNImageHomographicAlignmentObservationClassOnce.Do(func() {
		_VNImageHomographicAlignmentObservationClass = VNImageHomographicAlignmentObservationClass{class: objc.GetClass("VNImageHomographicAlignmentObservation")}
	})
	return _VNImageHomographicAlignmentObservationClass
}

// GetVNImageHomographicAlignmentObservationClass returns the class object for VNImageHomographicAlignmentObservation.
func GetVNImageHomographicAlignmentObservationClass() VNImageHomographicAlignmentObservationClass {
	return getVNImageHomographicAlignmentObservationClass()
}

type VNImageHomographicAlignmentObservationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNImageHomographicAlignmentObservationClass) Alloc() VNImageHomographicAlignmentObservation {
	rv := objc.Send[VNImageHomographicAlignmentObservation](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}







// An object that represents a perspective warp transformation.
//
// # Overview
// 
// This type of observation results from a
// [VNHomographicImageRegistrationRequest], informing the [VNImageHomographicAlignmentObservation.WarpTransform]
// performed to align the input images.
//
// # Accessing the Transform
//
//   - [VNImageHomographicAlignmentObservation.WarpTransform]: The warp transform matrix to morph the floating image into the reference image.
//
// See: https://developer.apple.com/documentation/Vision/VNImageHomographicAlignmentObservation
type VNImageHomographicAlignmentObservation struct {
	VNImageAlignmentObservation
}

// VNImageHomographicAlignmentObservationFromID constructs a [VNImageHomographicAlignmentObservation] from an objc.ID.
//
// An object that represents a perspective warp transformation.
func VNImageHomographicAlignmentObservationFromID(id objc.ID) VNImageHomographicAlignmentObservation {
	return VNImageHomographicAlignmentObservation{VNImageAlignmentObservation: VNImageAlignmentObservationFromID(id)}
}
// NOTE: VNImageHomographicAlignmentObservation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [VNImageHomographicAlignmentObservation] class.
//
// # Accessing the Transform
//
//   - [IVNImageHomographicAlignmentObservation.WarpTransform]: The warp transform matrix to morph the floating image into the reference image.
//
// See: https://developer.apple.com/documentation/Vision/VNImageHomographicAlignmentObservation
type IVNImageHomographicAlignmentObservation interface {
	IVNImageAlignmentObservation

	// Topic: Accessing the Transform

	// The warp transform matrix to morph the floating image into the reference image.
	WarpTransform() objectivec.IObject

	EncodeWithCoder(coder foundation.INSCoder)
}





// Init initializes the instance.
func (i VNImageHomographicAlignmentObservation) Init() VNImageHomographicAlignmentObservation {
	rv := objc.Send[VNImageHomographicAlignmentObservation](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i VNImageHomographicAlignmentObservation) Autorelease() VNImageHomographicAlignmentObservation {
	rv := objc.Send[VNImageHomographicAlignmentObservation](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNImageHomographicAlignmentObservation creates a new VNImageHomographicAlignmentObservation instance.
func NewVNImageHomographicAlignmentObservation() VNImageHomographicAlignmentObservation {
	class := getVNImageHomographicAlignmentObservationClass()
	rv := objc.Send[VNImageHomographicAlignmentObservation](objc.ID(class.class), objc.Sel("new"))
	return rv
}









func (i VNImageHomographicAlignmentObservation) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](i.ID, objc.Sel("encodeWithCoder:"), coder)
}












// The warp transform matrix to morph the floating image into the reference
// image.
//
// See: https://developer.apple.com/documentation/Vision/VNImageHomographicAlignmentObservation/warpTransform
func (i VNImageHomographicAlignmentObservation) WarpTransform() objectivec.IObject {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("warpTransform"))
	return objectivec.Object{ID: rv}
}



























