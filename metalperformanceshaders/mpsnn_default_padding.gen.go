// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNDefaultPadding] class.
var (
	_MPSNNDefaultPaddingClass     MPSNNDefaultPaddingClass
	_MPSNNDefaultPaddingClassOnce sync.Once
)

func getMPSNNDefaultPaddingClass() MPSNNDefaultPaddingClass {
	_MPSNNDefaultPaddingClassOnce.Do(func() {
		_MPSNNDefaultPaddingClass = MPSNNDefaultPaddingClass{class: objc.GetClass("MPSNNDefaultPadding")}
	})
	return _MPSNNDefaultPaddingClass
}

// GetMPSNNDefaultPaddingClass returns the class object for MPSNNDefaultPadding.
func GetMPSNNDefaultPaddingClass() MPSNNDefaultPaddingClass {
	return getMPSNNDefaultPaddingClass()
}

type MPSNNDefaultPaddingClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNDefaultPaddingClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNDefaultPaddingClass) Alloc() MPSNNDefaultPadding {
	rv := objc.Send[MPSNNDefaultPadding](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A class that provides predefined padding policies for common tasks.
//
// # Instance Methods
//
//   - [MPSNNDefaultPadding.Label]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNDefaultPadding
type MPSNNDefaultPadding struct {
	objectivec.Object
}

// MPSNNDefaultPaddingFromID constructs a [MPSNNDefaultPadding] from an objc.ID.
//
// A class that provides predefined padding policies for common tasks.
func MPSNNDefaultPaddingFromID(id objc.ID) MPSNNDefaultPadding {
	return MPSNNDefaultPadding{objectivec.Object{ID: id}}
}

// NOTE: MPSNNDefaultPadding adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNDefaultPadding] class.
//
// # Instance Methods
//
//   - [IMPSNNDefaultPadding.Label]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNDefaultPadding
type IMPSNNDefaultPadding interface {
	objectivec.IObject

	// Topic: Instance Methods

	Label() string

	DestinationImageDescriptorForSourceImagesSourceStatesForKernelSuggestedDescriptor(sourceImages []MPSImage, sourceStates []MPSState, kernel IMPSKernel, inDescriptor IMPSImageDescriptor) IMPSImageDescriptor
	Inverse() IMPSNNDefaultPadding
	PaddingMethod() MPSNNPaddingMethod
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (d MPSNNDefaultPadding) Init() MPSNNDefaultPadding {
	rv := objc.Send[MPSNNDefaultPadding](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d MPSNNDefaultPadding) Autorelease() MPSNNDefaultPadding {
	rv := objc.Send[MPSNNDefaultPadding](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNDefaultPadding creates a new MPSNNDefaultPadding instance.
func NewMPSNNDefaultPadding() MPSNNDefaultPadding {
	class := getMPSNNDefaultPaddingClass()
	rv := objc.Send[MPSNNDefaultPadding](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNDefaultPadding/init(method:)
func NewDefaultPaddingWithMethod(method MPSNNPaddingMethod) MPSNNDefaultPadding {
	rv := objc.Send[objc.ID](objc.ID(getMPSNNDefaultPaddingClass().class), objc.Sel("paddingWithMethod:"), method)
	return MPSNNDefaultPaddingFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNDefaultPadding/label()
func (d MPSNNDefaultPadding) Label() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPadding/destinationImageDescriptor(forSourceImages:sourceStates:for:suggestedDescriptor:)
func (d MPSNNDefaultPadding) DestinationImageDescriptorForSourceImagesSourceStatesForKernelSuggestedDescriptor(sourceImages []MPSImage, sourceStates []MPSState, kernel IMPSKernel, inDescriptor IMPSImageDescriptor) IMPSImageDescriptor {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("destinationImageDescriptorForSourceImages:sourceStates:forKernel:suggestedDescriptor:"), objectivec.IObjectSliceToNSArray(sourceImages), objectivec.IObjectSliceToNSArray(sourceStates), kernel, inDescriptor)
	return MPSImageDescriptorFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPadding/inverse()
func (d MPSNNDefaultPadding) Inverse() IMPSNNDefaultPadding {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("inverse"))
	return MPSNNDefaultPaddingFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPadding/paddingMethod()
func (d MPSNNDefaultPadding) PaddingMethod() MPSNNPaddingMethod {
	rv := objc.Send[MPSNNPaddingMethod](d.ID, objc.Sel("paddingMethod"))
	return MPSNNPaddingMethod(rv)
}
func (d MPSNNDefaultPadding) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](d.ID, objc.Sel("encodeWithCoder:"), coder)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNDefaultPadding/forTensorflowAveragePooling()
func (_MPSNNDefaultPaddingClass MPSNNDefaultPaddingClass) PaddingForTensorflowAveragePooling() MPSNNDefaultPadding {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNDefaultPaddingClass.class), objc.Sel("paddingForTensorflowAveragePooling"))
	return MPSNNDefaultPaddingFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNDefaultPadding/forTensorflowAveragePoolingValidOnly()
func (_MPSNNDefaultPaddingClass MPSNNDefaultPaddingClass) PaddingForTensorflowAveragePoolingValidOnly() MPSNNDefaultPadding {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNDefaultPaddingClass.class), objc.Sel("paddingForTensorflowAveragePoolingValidOnly"))
	return MPSNNDefaultPaddingFromID(rv)
}

// Protocol methods for MPSNNPadding
