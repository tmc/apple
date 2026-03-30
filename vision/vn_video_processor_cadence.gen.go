// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VNVideoProcessorCadence] class.
var (
	_VNVideoProcessorCadenceClass     VNVideoProcessorCadenceClass
	_VNVideoProcessorCadenceClassOnce sync.Once
)

func getVNVideoProcessorCadenceClass() VNVideoProcessorCadenceClass {
	_VNVideoProcessorCadenceClassOnce.Do(func() {
		_VNVideoProcessorCadenceClass = VNVideoProcessorCadenceClass{class: objc.GetClass("VNVideoProcessorCadence")}
	})
	return _VNVideoProcessorCadenceClass
}

// GetVNVideoProcessorCadenceClass returns the class object for VNVideoProcessorCadence.
func GetVNVideoProcessorCadenceClass() VNVideoProcessorCadenceClass {
	return getVNVideoProcessorCadenceClass()
}

type VNVideoProcessorCadenceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VNVideoProcessorCadenceClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNVideoProcessorCadenceClass) Alloc() VNVideoProcessorCadence {
	rv := objc.Send[VNVideoProcessorCadence](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An object that defines the cadence at which to process video.
//
// See: https://developer.apple.com/documentation/Vision/VNVideoProcessor/Cadence
type VNVideoProcessorCadence struct {
	objectivec.Object
}

// VNVideoProcessorCadenceFromID constructs a [VNVideoProcessorCadence] from an objc.ID.
//
// An object that defines the cadence at which to process video.
func VNVideoProcessorCadenceFromID(id objc.ID) VNVideoProcessorCadence {
	return VNVideoProcessorCadence{objectivec.Object{ID: id}}
}

// NOTE: VNVideoProcessorCadence adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNVideoProcessorCadence] class.
//
// See: https://developer.apple.com/documentation/Vision/VNVideoProcessor/Cadence
type IVNVideoProcessorCadence interface {
	objectivec.IObject

	// The cadence the video processor maintains to process the request.
	Cadence() IVNVideoProcessorCadence
	SetCadence(value IVNVideoProcessorCadence)
}

// Init initializes the instance.
func (v VNVideoProcessorCadence) Init() VNVideoProcessorCadence {
	rv := objc.Send[VNVideoProcessorCadence](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VNVideoProcessorCadence) Autorelease() VNVideoProcessorCadence {
	rv := objc.Send[VNVideoProcessorCadence](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNVideoProcessorCadence creates a new VNVideoProcessorCadence instance.
func NewVNVideoProcessorCadence() VNVideoProcessorCadence {
	class := getVNVideoProcessorCadenceClass()
	rv := objc.Send[VNVideoProcessorCadence](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The cadence the video processor maintains to process the request.
//
// See: https://developer.apple.com/documentation/vision/vnvideoprocessor/requestprocessingoptions/cadence
func (v VNVideoProcessorCadence) Cadence() IVNVideoProcessorCadence {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("cadence"))
	return VNVideoProcessorCadenceFromID(objc.ID(rv))
}
func (v VNVideoProcessorCadence) SetCadence(value IVNVideoProcessorCadence) {
	objc.Send[struct{}](v.ID, objc.Sel("setCadence:"), value)
}
