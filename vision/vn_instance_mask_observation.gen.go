// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corevideo"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [VNInstanceMaskObservation] class.
var (
	_VNInstanceMaskObservationClass     VNInstanceMaskObservationClass
	_VNInstanceMaskObservationClassOnce sync.Once
)

func getVNInstanceMaskObservationClass() VNInstanceMaskObservationClass {
	_VNInstanceMaskObservationClassOnce.Do(func() {
		_VNInstanceMaskObservationClass = VNInstanceMaskObservationClass{class: objc.GetClass("VNInstanceMaskObservation")}
	})
	return _VNInstanceMaskObservationClass
}

// GetVNInstanceMaskObservationClass returns the class object for VNInstanceMaskObservation.
func GetVNInstanceMaskObservationClass() VNInstanceMaskObservationClass {
	return getVNInstanceMaskObservationClass()
}

type VNInstanceMaskObservationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNInstanceMaskObservationClass) Alloc() VNInstanceMaskObservation {
	rv := objc.Send[VNInstanceMaskObservation](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An observation that contains an instance mask that labels instances in the
// mask.
//
// # Accessing Instances
//
//   - [VNInstanceMaskObservation.AllInstances]: The collection that contains all instances, excluding the background.
//   - [VNInstanceMaskObservation.InstanceMask]: The resulting mask that represents all instances.
//
// # Creating a Mask
//
//   - [VNInstanceMaskObservation.GenerateMaskForInstancesError]: Creates a low-resolution mask from the instances you specify.
//   - [VNInstanceMaskObservation.GenerateMaskedImageOfInstancesFromRequestHandlerCroppedToInstancesExtentError]: Creates a high-resolution image where everything becomes transparent black, except for the instances you specify.
//   - [VNInstanceMaskObservation.GenerateScaledMaskForImageForInstancesFromRequestHandlerError]: Creates a high-resolution mask where everything becomes transparent black, except for the instances you specify.
//
// See: https://developer.apple.com/documentation/Vision/VNInstanceMaskObservation
type VNInstanceMaskObservation struct {
	VNObservation
}

// VNInstanceMaskObservationFromID constructs a [VNInstanceMaskObservation] from an objc.ID.
//
// An observation that contains an instance mask that labels instances in the
// mask.
func VNInstanceMaskObservationFromID(id objc.ID) VNInstanceMaskObservation {
	return VNInstanceMaskObservation{VNObservation: VNObservationFromID(id)}
}
// NOTE: VNInstanceMaskObservation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNInstanceMaskObservation] class.
//
// # Accessing Instances
//
//   - [IVNInstanceMaskObservation.AllInstances]: The collection that contains all instances, excluding the background.
//   - [IVNInstanceMaskObservation.InstanceMask]: The resulting mask that represents all instances.
//
// # Creating a Mask
//
//   - [IVNInstanceMaskObservation.GenerateMaskForInstancesError]: Creates a low-resolution mask from the instances you specify.
//   - [IVNInstanceMaskObservation.GenerateMaskedImageOfInstancesFromRequestHandlerCroppedToInstancesExtentError]: Creates a high-resolution image where everything becomes transparent black, except for the instances you specify.
//   - [IVNInstanceMaskObservation.GenerateScaledMaskForImageForInstancesFromRequestHandlerError]: Creates a high-resolution mask where everything becomes transparent black, except for the instances you specify.
//
// See: https://developer.apple.com/documentation/Vision/VNInstanceMaskObservation
type IVNInstanceMaskObservation interface {
	IVNObservation

	// Topic: Accessing Instances

	// The collection that contains all instances, excluding the background.
	AllInstances() foundation.NSIndexSet
	// The resulting mask that represents all instances.
	InstanceMask() corevideo.CVImageBufferRef

	// Topic: Creating a Mask

	// Creates a low-resolution mask from the instances you specify.
	GenerateMaskForInstancesError(instances foundation.NSIndexSet) (corevideo.CVImageBufferRef, error)
	// Creates a high-resolution image where everything becomes transparent black, except for the instances you specify.
	GenerateMaskedImageOfInstancesFromRequestHandlerCroppedToInstancesExtentError(instances foundation.NSIndexSet, requestHandler IVNImageRequestHandler, cropResult bool) (corevideo.CVImageBufferRef, error)
	// Creates a high-resolution mask where everything becomes transparent black, except for the instances you specify.
	GenerateScaledMaskForImageForInstancesFromRequestHandlerError(instances foundation.NSIndexSet, requestHandler IVNImageRequestHandler) (corevideo.CVImageBufferRef, error)

	// A constant for specifying the first revision of the foreground instance mask request.
	VNGenerateForegroundInstanceMaskRequestRevision1() int
}

// Init initializes the instance.
func (i VNInstanceMaskObservation) Init() VNInstanceMaskObservation {
	rv := objc.Send[VNInstanceMaskObservation](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i VNInstanceMaskObservation) Autorelease() VNInstanceMaskObservation {
	rv := objc.Send[VNInstanceMaskObservation](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNInstanceMaskObservation creates a new VNInstanceMaskObservation instance.
func NewVNInstanceMaskObservation() VNInstanceMaskObservation {
	class := getVNInstanceMaskObservationClass()
	rv := objc.Send[VNInstanceMaskObservation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a low-resolution mask from the instances you specify.
//
// instances: The collection of instances.
//
// # Return Value
// 
// The pixel buffer that contains the image.
//
// See: https://developer.apple.com/documentation/Vision/VNInstanceMaskObservation/generateMask(forInstances:)
func (i VNInstanceMaskObservation) GenerateMaskForInstancesError(instances foundation.NSIndexSet) (corevideo.CVImageBufferRef, error) {
	var errorPtr objc.ID
	rv := objc.Send[corevideo.CVImageBufferRef](i.ID, objc.Sel("generateMaskForInstances:error:"), instances, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0, foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}

// Creates a high-resolution image where everything becomes transparent black,
// except for the instances you specify.
//
// instances: The collection of instances.
//
// requestHandler: The image request callback.
//
// cropResult: A Boolean value that indicates whether to crop the image to the smallest
// rectangle that contains all instances.
//
// # Return Value
// 
// The pixel buffer that contains the image.
//
// See: https://developer.apple.com/documentation/Vision/VNInstanceMaskObservation/generateMaskedImage(ofInstances:from:croppedToInstancesExtent:)
func (i VNInstanceMaskObservation) GenerateMaskedImageOfInstancesFromRequestHandlerCroppedToInstancesExtentError(instances foundation.NSIndexSet, requestHandler IVNImageRequestHandler, cropResult bool) (corevideo.CVImageBufferRef, error) {
	var errorPtr objc.ID
	rv := objc.Send[corevideo.CVImageBufferRef](i.ID, objc.Sel("generateMaskedImageOfInstances:fromRequestHandler:croppedToInstancesExtent:error:"), instances, requestHandler, cropResult, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0, foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}

// Creates a high-resolution mask where everything becomes transparent black,
// except for the instances you specify.
//
// instances: The collection of instances.
//
// requestHandler: The image request callback.
//
// # Return Value
// 
// The pixel buffer that contains the image.
//
// See: https://developer.apple.com/documentation/Vision/VNInstanceMaskObservation/generateScaledMaskForImage(forInstances:from:)
func (i VNInstanceMaskObservation) GenerateScaledMaskForImageForInstancesFromRequestHandlerError(instances foundation.NSIndexSet, requestHandler IVNImageRequestHandler) (corevideo.CVImageBufferRef, error) {
	var errorPtr objc.ID
	rv := objc.Send[corevideo.CVImageBufferRef](i.ID, objc.Sel("generateScaledMaskForImageForInstances:fromRequestHandler:error:"), instances, requestHandler, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0, foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}

// The collection that contains all instances, excluding the background.
//
// See: https://developer.apple.com/documentation/Vision/VNInstanceMaskObservation/allInstances
func (i VNInstanceMaskObservation) AllInstances() foundation.NSIndexSet {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("allInstances"))
	return foundation.NSIndexSetFromID(objc.ID(rv))
}

// The resulting mask that represents all instances.
//
// # Discussion
// 
// A pixel can only correspond to one instance. A `0` represents the
// background, and all other values represent the indices of the instances.
//
// See: https://developer.apple.com/documentation/Vision/VNInstanceMaskObservation/instanceMask
func (i VNInstanceMaskObservation) InstanceMask() corevideo.CVImageBufferRef {
	rv := objc.Send[corevideo.CVImageBufferRef](i.ID, objc.Sel("instanceMask"))
	return corevideo.CVImageBufferRef(rv)
}

// A constant for specifying the first revision of the foreground instance
// mask request.
//
// See: https://developer.apple.com/documentation/vision/vngenerateforegroundinstancemaskrequestrevision1
func (i VNInstanceMaskObservation) VNGenerateForegroundInstanceMaskRequestRevision1() int {
	rv := objc.Send[int](i.ID, objc.Sel("VNGenerateForegroundInstanceMaskRequestRevision1"))
	return rv
}

