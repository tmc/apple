// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNReduceFeatureChannelsArgumentMin] class.
var (
	_MPSNNReduceFeatureChannelsArgumentMinClass     MPSNNReduceFeatureChannelsArgumentMinClass
	_MPSNNReduceFeatureChannelsArgumentMinClassOnce sync.Once
)

func getMPSNNReduceFeatureChannelsArgumentMinClass() MPSNNReduceFeatureChannelsArgumentMinClass {
	_MPSNNReduceFeatureChannelsArgumentMinClassOnce.Do(func() {
		_MPSNNReduceFeatureChannelsArgumentMinClass = MPSNNReduceFeatureChannelsArgumentMinClass{class: objc.GetClass("MPSNNReduceFeatureChannelsArgumentMin")}
	})
	return _MPSNNReduceFeatureChannelsArgumentMinClass
}

// GetMPSNNReduceFeatureChannelsArgumentMinClass returns the class object for MPSNNReduceFeatureChannelsArgumentMin.
func GetMPSNNReduceFeatureChannelsArgumentMinClass() MPSNNReduceFeatureChannelsArgumentMinClass {
	return getMPSNNReduceFeatureChannelsArgumentMinClass()
}

type MPSNNReduceFeatureChannelsArgumentMinClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNReduceFeatureChannelsArgumentMinClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReduceFeatureChannelsArgumentMinClass) Alloc() MPSNNReduceFeatureChannelsArgumentMin {
	rv := objc.Send[MPSNNReduceFeatureChannelsArgumentMin](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A reduction filter that returns the index of the location of the minimum
// value for each feature channel in an image.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceFeatureChannelsArgumentMin
type MPSNNReduceFeatureChannelsArgumentMin struct {
	MPSNNReduceUnary
}

// MPSNNReduceFeatureChannelsArgumentMinFromID constructs a [MPSNNReduceFeatureChannelsArgumentMin] from an objc.ID.
//
// A reduction filter that returns the index of the location of the minimum
// value for each feature channel in an image.
func MPSNNReduceFeatureChannelsArgumentMinFromID(id objc.ID) MPSNNReduceFeatureChannelsArgumentMin {
	return MPSNNReduceFeatureChannelsArgumentMin{MPSNNReduceUnary: MPSNNReduceUnaryFromID(id)}
}

// NOTE: MPSNNReduceFeatureChannelsArgumentMin adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNReduceFeatureChannelsArgumentMin] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceFeatureChannelsArgumentMin
type IMPSNNReduceFeatureChannelsArgumentMin interface {
	IMPSNNReduceUnary
}

// Init initializes the instance.
func (r MPSNNReduceFeatureChannelsArgumentMin) Init() MPSNNReduceFeatureChannelsArgumentMin {
	rv := objc.Send[MPSNNReduceFeatureChannelsArgumentMin](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReduceFeatureChannelsArgumentMin) Autorelease() MPSNNReduceFeatureChannelsArgumentMin {
	rv := objc.Send[MPSNNReduceFeatureChannelsArgumentMin](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReduceFeatureChannelsArgumentMin creates a new MPSNNReduceFeatureChannelsArgumentMin instance.
func NewMPSNNReduceFeatureChannelsArgumentMin() MPSNNReduceFeatureChannelsArgumentMin {
	class := getMPSNNReduceFeatureChannelsArgumentMinClass()
	rv := objc.Send[MPSNNReduceFeatureChannelsArgumentMin](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewReduceFeatureChannelsArgumentMinWithCoder(aDecoder foundation.INSCoder) MPSNNReduceFeatureChannelsArgumentMin {
	instance := getMPSNNReduceFeatureChannelsArgumentMinClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSNNReduceFeatureChannelsArgumentMinFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceFeatureChannelsArgumentMin/init(coder:device:)
func NewReduceFeatureChannelsArgumentMinWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSNNReduceFeatureChannelsArgumentMin {
	instance := getMPSNNReduceFeatureChannelsArgumentMinClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSNNReduceFeatureChannelsArgumentMinFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceFeatureChannelsArgumentMin/init(device:)
func NewReduceFeatureChannelsArgumentMinWithDevice(device metal.MTLDevice) MPSNNReduceFeatureChannelsArgumentMin {
	instance := getMPSNNReduceFeatureChannelsArgumentMinClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNReduceFeatureChannelsArgumentMinFromID(rv)
}
