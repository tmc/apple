// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
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

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNFilterNodeClass) Alloc() MPSNNFilterNode {
	rv := objc.Send[MPSNNFilterNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNFilterNode.DebugQuickLookObject]
//   - [MPSNNFilterNode.GradientClass]
//   - [MPSNNFilterNode.GradientFilterWithSource]
//   - [MPSNNFilterNode.GradientFilterWithSources]
//   - [MPSNNFilterNode.GradientFiltersWithSource]
//   - [MPSNNFilterNode.GradientFiltersWithSources]
//   - [MPSNNFilterNode.Label]
//   - [MPSNNFilterNode.SetLabel]
//   - [MPSNNFilterNode.NewFilterNode]
//   - [MPSNNFilterNode.PaddingPolicy]
//   - [MPSNNFilterNode.SetPaddingPolicy]
//   - [MPSNNFilterNode.ResultImage]
//   - [MPSNNFilterNode.ResultState]
//   - [MPSNNFilterNode.ResultStates]
//   - [MPSNNFilterNode.ResultStatesNoAllocate]
//   - [MPSNNFilterNode.SourceImages]
//   - [MPSNNFilterNode.SourceStates]
//   - [MPSNNFilterNode.TrainingGraphWithSourceGradientNodeHandler]
//   - [MPSNNFilterNode.InitWithSourceImagesSourceStatesPaddingPolicy]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode
type MPSNNFilterNode struct {
	objectivec.Object
}

// MPSNNFilterNodeFromID constructs a [MPSNNFilterNode] from an objc.ID.
func MPSNNFilterNodeFromID(id objc.ID) MPSNNFilterNode {
	return MPSNNFilterNode{objectivec.Object{id}}
}
// Ensure MPSNNFilterNode implements IMPSNNFilterNode.
var _ IMPSNNFilterNode = MPSNNFilterNode{}





// An interface definition for the [MPSNNFilterNode] class.
//
// # Methods
//
//   - [IMPSNNFilterNode.DebugQuickLookObject]
//   - [IMPSNNFilterNode.GradientClass]
//   - [IMPSNNFilterNode.GradientFilterWithSource]
//   - [IMPSNNFilterNode.GradientFilterWithSources]
//   - [IMPSNNFilterNode.GradientFiltersWithSource]
//   - [IMPSNNFilterNode.GradientFiltersWithSources]
//   - [IMPSNNFilterNode.Label]
//   - [IMPSNNFilterNode.SetLabel]
//   - [IMPSNNFilterNode.NewFilterNode]
//   - [IMPSNNFilterNode.PaddingPolicy]
//   - [IMPSNNFilterNode.SetPaddingPolicy]
//   - [IMPSNNFilterNode.ResultImage]
//   - [IMPSNNFilterNode.ResultState]
//   - [IMPSNNFilterNode.ResultStates]
//   - [IMPSNNFilterNode.ResultStatesNoAllocate]
//   - [IMPSNNFilterNode.SourceImages]
//   - [IMPSNNFilterNode.SourceStates]
//   - [IMPSNNFilterNode.TrainingGraphWithSourceGradientNodeHandler]
//   - [IMPSNNFilterNode.InitWithSourceImagesSourceStatesPaddingPolicy]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode
type IMPSNNFilterNode interface {
	objectivec.IObject

	// Topic: Methods

	DebugQuickLookObject() objectivec.IObject
	GradientClass() objc.Class
	GradientFilterWithSource(source objectivec.IObject) objectivec.IObject
	GradientFilterWithSources(sources objectivec.IObject) objectivec.IObject
	GradientFiltersWithSource(source objectivec.IObject) objectivec.IObject
	GradientFiltersWithSources(sources objectivec.IObject) objectivec.IObject
	Label() string
	SetLabel(value string)
	NewFilterNode() FilterGraphNode
	PaddingPolicy() unsafe.Pointer
	SetPaddingPolicy(value unsafe.Pointer)
	ResultImage() IMPSNNImageNode
	ResultState() IMPSNNStateNode
	ResultStates() foundation.INSArray
	ResultStatesNoAllocate() objectivec.IObject
	SourceImages() objectivec.IObject
	SourceStates() objectivec.IObject
	TrainingGraphWithSourceGradientNodeHandler(gradient objectivec.IObject, handler ErrorHandler) objectivec.IObject
	InitWithSourceImagesSourceStatesPaddingPolicy(images objectivec.IObject, states objectivec.IObject, policy objectivec.IObject) MPSNNFilterNode
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






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/initWithSourceImages:sourceStates:paddingPolicy:
func NewFilterNodeWithSourceImagesSourceStatesPaddingPolicy(images objectivec.IObject, states objectivec.IObject, policy objectivec.IObject) MPSNNFilterNode {
	instance := getMPSNNFilterNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceImages:sourceStates:paddingPolicy:"), images, states, policy)
	return MPSNNFilterNodeFromID(rv)
}







// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/debugQuickLookObject
func (f MPSNNFilterNode) DebugQuickLookObject() objectivec.IObject {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("debugQuickLookObject"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/gradientClass
func (f MPSNNFilterNode) GradientClass() objc.Class {
	rv := objc.Send[objc.Class](f.ID, objc.Sel("gradientClass"))
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/gradientFilterWithSource:
func (f MPSNNFilterNode) GradientFilterWithSource(source objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("gradientFilterWithSource:"), source)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/gradientFilterWithSources:
func (f MPSNNFilterNode) GradientFilterWithSources(sources objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("gradientFilterWithSources:"), sources)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/gradientFiltersWithSource:
func (f MPSNNFilterNode) GradientFiltersWithSource(source objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("gradientFiltersWithSource:"), source)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/gradientFiltersWithSources:
func (f MPSNNFilterNode) GradientFiltersWithSources(sources objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("gradientFiltersWithSources:"), sources)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/newFilterNode
func (f MPSNNFilterNode) NewFilterNode() FilterGraphNode {
	rv := objc.Send[FilterGraphNode](f.ID, objc.Sel("newFilterNode"))
	return FilterGraphNode(rv)
}

// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/resultStatesNoAllocate
func (f MPSNNFilterNode) ResultStatesNoAllocate() objectivec.IObject {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("resultStatesNoAllocate"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/sourceImages
func (f MPSNNFilterNode) SourceImages() objectivec.IObject {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("sourceImages"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/sourceStates
func (f MPSNNFilterNode) SourceStates() objectivec.IObject {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("sourceStates"))
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/trainingGraphWithSourceGradient:nodeHandler:
func (f MPSNNFilterNode) TrainingGraphWithSourceGradientNodeHandler(gradient objectivec.IObject, handler ErrorHandler) objectivec.IObject {
		_block1, _cleanup1 := NewErrorBlock(handler)
	defer _cleanup1()
		rv := objc.Send[objc.ID](f.ID, objc.Sel("trainingGraphWithSourceGradient:nodeHandler:"), gradient, _block1)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/initWithSourceImages:sourceStates:paddingPolicy:
func (f MPSNNFilterNode) InitWithSourceImagesSourceStatesPaddingPolicy(images objectivec.IObject, states objectivec.IObject, policy objectivec.IObject) MPSNNFilterNode {
	rv := objc.Send[MPSNNFilterNode](f.ID, objc.Sel("initWithSourceImages:sourceStates:paddingPolicy:"), images, states, policy)
	return rv
}












// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/label
func (f MPSNNFilterNode) Label() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}
func (f MPSNNFilterNode) SetLabel(value string) {
	objc.Send[struct{}](f.ID, objc.Sel("setLabel:"), objc.String(value))
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/paddingPolicy
func (f MPSNNFilterNode) PaddingPolicy() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f.ID, objc.Sel("paddingPolicy"))
	return rv
}
func (f MPSNNFilterNode) SetPaddingPolicy(value unsafe.Pointer) {
	objc.Send[struct{}](f.ID, objc.Sel("setPaddingPolicy:"), value)
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/resultImage
func (f MPSNNFilterNode) ResultImage() IMPSNNImageNode {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("resultImage"))
	return MPSNNImageNodeFromID(objc.ID(rv))
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/resultState
func (f MPSNNFilterNode) ResultState() IMPSNNStateNode {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("resultState"))
	return MPSNNStateNodeFromID(objc.ID(rv))
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/resultStates
func (f MPSNNFilterNode) ResultStates() foundation.INSArray {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("resultStates"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

















