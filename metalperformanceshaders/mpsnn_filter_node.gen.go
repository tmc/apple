// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNFilterNode] class.
var (
	_MPSNNFilterNodeClass     MPSNNFilterNodeClass
	_MPSNNFilterNodeClassOnce sync.Once
)

func getMPSNNFilterNodeClass() MPSNNFilterNodeClass {
	_MPSNNFilterNodeClassOnce.Do(func() {
		_MPSNNFilterNodeClass = MPSNNFilterNodeClass{class: objc.GetClass("MPSNNFilterNode")}
	})
	return _MPSNNFilterNodeClass
}

// GetMPSNNFilterNodeClass returns the class object for MPSNNFilterNode.
func GetMPSNNFilterNodeClass() MPSNNFilterNodeClass {
	return getMPSNNFilterNodeClass()
}

type MPSNNFilterNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNFilterNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNFilterNodeClass) Alloc() MPSNNFilterNode {
	rv := objc.Send[MPSNNFilterNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A placeholder node denoting a neural network filter stage.
//
// # Instance Properties
//
//   - [MPSNNFilterNode.Label]
//   - [MPSNNFilterNode.SetLabel]
//   - [MPSNNFilterNode.PaddingPolicy]
//   - [MPSNNFilterNode.SetPaddingPolicy]
//   - [MPSNNFilterNode.ResultImage]
//   - [MPSNNFilterNode.ResultState]
//   - [MPSNNFilterNode.ResultStates]
//
// # Instance Methods
//
//   - [MPSNNFilterNode.GradientFilterWithSource]
//   - [MPSNNFilterNode.GradientFilterWithSources]
//   - [MPSNNFilterNode.GradientFiltersWithSource]
//   - [MPSNNFilterNode.GradientFiltersWithSources]
//   - [MPSNNFilterNode.TrainingGraphWithSourceGradientNodeHandler]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNFilterNode
type MPSNNFilterNode struct {
	objectivec.Object
}

// MPSNNFilterNodeFromID constructs a [MPSNNFilterNode] from an objc.ID.
//
// A placeholder node denoting a neural network filter stage.
func MPSNNFilterNodeFromID(id objc.ID) MPSNNFilterNode {
	return MPSNNFilterNode{objectivec.Object{ID: id}}
}

// NOTE: MPSNNFilterNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNFilterNode] class.
//
// # Instance Properties
//
//   - [IMPSNNFilterNode.Label]
//   - [IMPSNNFilterNode.SetLabel]
//   - [IMPSNNFilterNode.PaddingPolicy]
//   - [IMPSNNFilterNode.SetPaddingPolicy]
//   - [IMPSNNFilterNode.ResultImage]
//   - [IMPSNNFilterNode.ResultState]
//   - [IMPSNNFilterNode.ResultStates]
//
// # Instance Methods
//
//   - [IMPSNNFilterNode.GradientFilterWithSource]
//   - [IMPSNNFilterNode.GradientFilterWithSources]
//   - [IMPSNNFilterNode.GradientFiltersWithSource]
//   - [IMPSNNFilterNode.GradientFiltersWithSources]
//   - [IMPSNNFilterNode.TrainingGraphWithSourceGradientNodeHandler]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNFilterNode
type IMPSNNFilterNode interface {
	objectivec.IObject

	// Topic: Instance Properties

	Label() string
	SetLabel(value string)
	PaddingPolicy() MPSNNPadding
	SetPaddingPolicy(value MPSNNPadding)
	ResultImage() IMPSNNImageNode
	ResultState() IMPSNNStateNode
	ResultStates() []MPSNNStateNode

	// Topic: Instance Methods

	GradientFilterWithSource(gradientImage IMPSNNImageNode) IMPSNNGradientFilterNode
	GradientFilterWithSources(gradientImages []MPSNNImageNode) IMPSNNGradientFilterNode
	GradientFiltersWithSource(gradientImage IMPSNNImageNode) []MPSNNGradientFilterNode
	GradientFiltersWithSources(gradientImages []MPSNNImageNode) []MPSNNGradientFilterNode
	TrainingGraphWithSourceGradientNodeHandler(gradientImage IMPSNNImageNode, nodeHandler MPSNNFilterNodeMPSNNFilterNodeMPSNNImageNodeMPSNNImageNodeHandler) []MPSNNFilterNode
}

// Init initializes the instance.
func (f MPSNNFilterNode) Init() MPSNNFilterNode {
	rv := objc.Send[MPSNNFilterNode](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f MPSNNFilterNode) Autorelease() MPSNNFilterNode {
	rv := objc.Send[MPSNNFilterNode](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNFilterNode creates a new MPSNNFilterNode instance.
func NewMPSNNFilterNode() MPSNNFilterNode {
	class := getMPSNNFilterNodeClass()
	rv := objc.Send[MPSNNFilterNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNFilterNode/gradientFilter(withSource:)
func (f MPSNNFilterNode) GradientFilterWithSource(gradientImage IMPSNNImageNode) IMPSNNGradientFilterNode {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("gradientFilterWithSource:"), gradientImage)
	return MPSNNGradientFilterNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNFilterNode/gradientFilter(withSources:)
func (f MPSNNFilterNode) GradientFilterWithSources(gradientImages []MPSNNImageNode) IMPSNNGradientFilterNode {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("gradientFilterWithSources:"), objectivec.IObjectSliceToNSArray(gradientImages))
	return MPSNNGradientFilterNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNFilterNode/gradientFilters(withSource:)
func (f MPSNNFilterNode) GradientFiltersWithSource(gradientImage IMPSNNImageNode) []MPSNNGradientFilterNode {
	rv := objc.Send[[]objc.ID](f.ID, objc.Sel("gradientFiltersWithSource:"), gradientImage)
	return objc.ConvertSlice(rv, func(id objc.ID) MPSNNGradientFilterNode {
		return MPSNNGradientFilterNodeFromID(id)
	})
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNFilterNode/gradientFilters(withSources:)
func (f MPSNNFilterNode) GradientFiltersWithSources(gradientImages []MPSNNImageNode) []MPSNNGradientFilterNode {
	rv := objc.Send[[]objc.ID](f.ID, objc.Sel("gradientFiltersWithSources:"), objectivec.IObjectSliceToNSArray(gradientImages))
	return objc.ConvertSlice(rv, func(id objc.ID) MPSNNGradientFilterNode {
		return MPSNNGradientFilterNodeFromID(id)
	})
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNFilterNode/trainingGraph(withSourceGradient:nodeHandler:)
func (f MPSNNFilterNode) TrainingGraphWithSourceGradientNodeHandler(gradientImage IMPSNNImageNode, nodeHandler MPSNNFilterNodeMPSNNFilterNodeMPSNNImageNodeMPSNNImageNodeHandler) []MPSNNFilterNode {
	_block1, _ := NewMPSNNFilterNodeMPSNNFilterNodeMPSNNImageNodeMPSNNImageNodeBlock(nodeHandler)
	rv := objc.Send[[]objc.ID](f.ID, objc.Sel("trainingGraphWithSourceGradient:nodeHandler:"), gradientImage, _block1)
	return objc.ConvertSlice(rv, func(id objc.ID) MPSNNFilterNode {
		return MPSNNFilterNodeFromID(id)
	})
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNFilterNode/label
func (f MPSNNFilterNode) Label() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}
func (f MPSNNFilterNode) SetLabel(value string) {
	objc.Send[struct{}](f.ID, objc.Sel("setLabel:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNFilterNode/paddingPolicy
func (f MPSNNFilterNode) PaddingPolicy() MPSNNPadding {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("paddingPolicy"))
	return MPSNNPaddingObjectFromID(rv)
}
func (f MPSNNFilterNode) SetPaddingPolicy(value MPSNNPadding) {
	objc.Send[struct{}](f.ID, objc.Sel("setPaddingPolicy:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNFilterNode/resultImage
func (f MPSNNFilterNode) ResultImage() IMPSNNImageNode {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("resultImage"))
	return MPSNNImageNodeFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNFilterNode/resultState
func (f MPSNNFilterNode) ResultState() IMPSNNStateNode {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("resultState"))
	return MPSNNStateNodeFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNFilterNode/resultStates
func (f MPSNNFilterNode) ResultStates() []MPSNNStateNode {
	rv := objc.Send[[]objc.ID](f.ID, objc.Sel("resultStates"))
	return objc.ConvertSlice(rv, func(id objc.ID) MPSNNStateNode {
		return MPSNNStateNodeFromID(id)
	})
}
