// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VNVideoProcessorRequestProcessingOptions] class.
var (
	_VNVideoProcessorRequestProcessingOptionsClass     VNVideoProcessorRequestProcessingOptionsClass
	_VNVideoProcessorRequestProcessingOptionsClassOnce sync.Once
)

func getVNVideoProcessorRequestProcessingOptionsClass() VNVideoProcessorRequestProcessingOptionsClass {
	_VNVideoProcessorRequestProcessingOptionsClassOnce.Do(func() {
		_VNVideoProcessorRequestProcessingOptionsClass = VNVideoProcessorRequestProcessingOptionsClass{class: objc.GetClass("VNVideoProcessorRequestProcessingOptions")}
	})
	return _VNVideoProcessorRequestProcessingOptionsClass
}

// GetVNVideoProcessorRequestProcessingOptionsClass returns the class object for VNVideoProcessorRequestProcessingOptions.
func GetVNVideoProcessorRequestProcessingOptionsClass() VNVideoProcessorRequestProcessingOptionsClass {
	return getVNVideoProcessorRequestProcessingOptionsClass()
}

type VNVideoProcessorRequestProcessingOptionsClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNVideoProcessorRequestProcessingOptionsClass) Alloc() VNVideoProcessorRequestProcessingOptions {
	rv := objc.Send[VNVideoProcessorRequestProcessingOptions](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}







// An object that defines a video processor’s configuration options.
//
// # Configuring Options
//
//   - [VNVideoProcessorRequestProcessingOptions.Cadence]: The cadence the video processor maintains to process the request.
//   - [VNVideoProcessorRequestProcessingOptions.SetCadence]
//
// See: https://developer.apple.com/documentation/Vision/VNVideoProcessor/RequestProcessingOptions
type VNVideoProcessorRequestProcessingOptions struct {
	objectivec.Object
}

// VNVideoProcessorRequestProcessingOptionsFromID constructs a [VNVideoProcessorRequestProcessingOptions] from an objc.ID.
//
// An object that defines a video processor’s configuration options.
func VNVideoProcessorRequestProcessingOptionsFromID(id objc.ID) VNVideoProcessorRequestProcessingOptions {
	return VNVideoProcessorRequestProcessingOptions{objectivec.Object{id}}
}
// NOTE: VNVideoProcessorRequestProcessingOptions adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [VNVideoProcessorRequestProcessingOptions] class.
//
// # Configuring Options
//
//   - [IVNVideoProcessorRequestProcessingOptions.Cadence]: The cadence the video processor maintains to process the request.
//   - [IVNVideoProcessorRequestProcessingOptions.SetCadence]
//
// See: https://developer.apple.com/documentation/Vision/VNVideoProcessor/RequestProcessingOptions
type IVNVideoProcessorRequestProcessingOptions interface {
	objectivec.IObject

	// Topic: Configuring Options

	// The cadence the video processor maintains to process the request.
	Cadence() IVNVideoProcessorCadence
	SetCadence(value IVNVideoProcessorCadence)
}





// Init initializes the instance.
func (v VNVideoProcessorRequestProcessingOptions) Init() VNVideoProcessorRequestProcessingOptions {
	rv := objc.Send[VNVideoProcessorRequestProcessingOptions](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VNVideoProcessorRequestProcessingOptions) Autorelease() VNVideoProcessorRequestProcessingOptions {
	rv := objc.Send[VNVideoProcessorRequestProcessingOptions](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNVideoProcessorRequestProcessingOptions creates a new VNVideoProcessorRequestProcessingOptions instance.
func NewVNVideoProcessorRequestProcessingOptions() VNVideoProcessorRequestProcessingOptions {
	class := getVNVideoProcessorRequestProcessingOptionsClass()
	rv := objc.Send[VNVideoProcessorRequestProcessingOptions](objc.ID(class.class), objc.Sel("new"))
	return rv
}





















// The cadence the video processor maintains to process the request.
//
// # Discussion
// 
// By default, the system processes every frame of video if you don’t
// provide a value for this property.
//
// See: https://developer.apple.com/documentation/Vision/VNVideoProcessor/RequestProcessingOptions/cadence
func (v VNVideoProcessorRequestProcessingOptions) Cadence() IVNVideoProcessorCadence {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("cadence"))
	return VNVideoProcessorCadenceFromID(objc.ID(rv))
}
func (v VNVideoProcessorRequestProcessingOptions) SetCadence(value IVNVideoProcessorCadence) {
	objc.Send[struct{}](v.ID, objc.Sel("setCadence:"), value)
}
























