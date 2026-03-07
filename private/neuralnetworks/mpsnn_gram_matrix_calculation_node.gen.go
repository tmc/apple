// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNGramMatrixCalculationNode] class.
var (
	_MPSNNGramMatrixCalculationNodeClass     MPSNNGramMatrixCalculationNodeClass
	_MPSNNGramMatrixCalculationNodeClassOnce sync.Once
)

func getMPSNNGramMatrixCalculationNodeClass() MPSNNGramMatrixCalculationNodeClass {
	_MPSNNGramMatrixCalculationNodeClassOnce.Do(func() {
		_MPSNNGramMatrixCalculationNodeClass = MPSNNGramMatrixCalculationNodeClass{class: objc.GetClass("MPSNNGramMatrixCalculationNode")}
	})
	return _MPSNNGramMatrixCalculationNodeClass
}

// GetMPSNNGramMatrixCalculationNodeClass returns the class object for MPSNNGramMatrixCalculationNode.
func GetMPSNNGramMatrixCalculationNodeClass() MPSNNGramMatrixCalculationNodeClass {
	return getMPSNNGramMatrixCalculationNodeClass()
}

type MPSNNGramMatrixCalculationNodeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNGramMatrixCalculationNodeClass) Alloc() MPSNNGramMatrixCalculationNode {
	rv := objc.Send[MPSNNGramMatrixCalculationNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNGramMatrixCalculationNode.Alpha]
//   - [MPSNNGramMatrixCalculationNode.PropertyCallBack]
//   - [MPSNNGramMatrixCalculationNode.SetPropertyCallBack]
//   - [MPSNNGramMatrixCalculationNode.InitWithSource]
//   - [MPSNNGramMatrixCalculationNode.InitWithSourceAlpha]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGramMatrixCalculationNode
type MPSNNGramMatrixCalculationNode struct {
	MPSNNFilterNode
}

// MPSNNGramMatrixCalculationNodeFromID constructs a [MPSNNGramMatrixCalculationNode] from an objc.ID.
func MPSNNGramMatrixCalculationNodeFromID(id objc.ID) MPSNNGramMatrixCalculationNode {
	return MPSNNGramMatrixCalculationNode{MPSNNFilterNode: MPSNNFilterNodeFromID(id)}
}
// Ensure MPSNNGramMatrixCalculationNode implements IMPSNNGramMatrixCalculationNode.
var _ IMPSNNGramMatrixCalculationNode = MPSNNGramMatrixCalculationNode{}





// An interface definition for the [MPSNNGramMatrixCalculationNode] class.
//
// # Methods
//
//   - [IMPSNNGramMatrixCalculationNode.Alpha]
//   - [IMPSNNGramMatrixCalculationNode.PropertyCallBack]
//   - [IMPSNNGramMatrixCalculationNode.SetPropertyCallBack]
//   - [IMPSNNGramMatrixCalculationNode.InitWithSource]
//   - [IMPSNNGramMatrixCalculationNode.InitWithSourceAlpha]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGramMatrixCalculationNode
type IMPSNNGramMatrixCalculationNode interface {
	IMPSNNFilterNode

	// Topic: Methods

	Alpha() float32
	PropertyCallBack() unsafe.Pointer
	SetPropertyCallBack(value unsafe.Pointer)
	InitWithSource(source objectivec.IObject) MPSNNGramMatrixCalculationNode
	InitWithSourceAlpha(source objectivec.IObject, alpha float32) MPSNNGramMatrixCalculationNode
}





// Init initializes the instance.
func (g MPSNNGramMatrixCalculationNode) Init() MPSNNGramMatrixCalculationNode {
	rv := objc.Send[MPSNNGramMatrixCalculationNode](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g MPSNNGramMatrixCalculationNode) Autorelease() MPSNNGramMatrixCalculationNode {
	rv := objc.Send[MPSNNGramMatrixCalculationNode](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNGramMatrixCalculationNode creates a new MPSNNGramMatrixCalculationNode instance.
func NewMPSNNGramMatrixCalculationNode() MPSNNGramMatrixCalculationNode {
	class := getMPSNNGramMatrixCalculationNodeClass()
	rv := objc.Send[MPSNNGramMatrixCalculationNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGramMatrixCalculationNode/initWithSource:
func NewGramMatrixCalculationNodeWithSource(source objectivec.IObject) MPSNNGramMatrixCalculationNode {
	instance := getMPSNNGramMatrixCalculationNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:"), source)
	return MPSNNGramMatrixCalculationNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGramMatrixCalculationNode/initWithSource:alpha:
func NewGramMatrixCalculationNodeWithSourceAlpha(source objectivec.IObject, alpha float32) MPSNNGramMatrixCalculationNode {
	instance := getMPSNNGramMatrixCalculationNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:alpha:"), source, alpha)
	return MPSNNGramMatrixCalculationNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/initWithSourceImages:sourceStates:paddingPolicy:
func NewGramMatrixCalculationNodeWithSourceImagesSourceStatesPaddingPolicy(images objectivec.IObject, states objectivec.IObject, policy objectivec.IObject) MPSNNGramMatrixCalculationNode {
	instance := getMPSNNGramMatrixCalculationNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceImages:sourceStates:paddingPolicy:"), images, states, policy)
	return MPSNNGramMatrixCalculationNodeFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGramMatrixCalculationNode/initWithSource:
func (g MPSNNGramMatrixCalculationNode) InitWithSource(source objectivec.IObject) MPSNNGramMatrixCalculationNode {
	rv := objc.Send[MPSNNGramMatrixCalculationNode](g.ID, objc.Sel("initWithSource:"), source)
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGramMatrixCalculationNode/initWithSource:alpha:
func (g MPSNNGramMatrixCalculationNode) InitWithSourceAlpha(source objectivec.IObject, alpha float32) MPSNNGramMatrixCalculationNode {
	rv := objc.Send[MPSNNGramMatrixCalculationNode](g.ID, objc.Sel("initWithSource:alpha:"), source, alpha)
	return rv
}





//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGramMatrixCalculationNode/nodeWithSource:
func (_MPSNNGramMatrixCalculationNodeClass MPSNNGramMatrixCalculationNodeClass) NodeWithSource(source objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNGramMatrixCalculationNodeClass.class), objc.Sel("nodeWithSource:"), source)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGramMatrixCalculationNode/nodeWithSource:alpha:
func (_MPSNNGramMatrixCalculationNodeClass MPSNNGramMatrixCalculationNodeClass) NodeWithSourceAlpha(source objectivec.IObject, alpha float32) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNGramMatrixCalculationNodeClass.class), objc.Sel("nodeWithSource:alpha:"), source, alpha)
	return objectivec.Object{ID: rv}
}








// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGramMatrixCalculationNode/alpha
func (g MPSNNGramMatrixCalculationNode) Alpha() float32 {
	rv := objc.Send[float32](g.ID, objc.Sel("alpha"))
	return rv
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGramMatrixCalculationNode/propertyCallBack
func (g MPSNNGramMatrixCalculationNode) PropertyCallBack() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("propertyCallBack"))
	return rv
}
func (g MPSNNGramMatrixCalculationNode) SetPropertyCallBack(value unsafe.Pointer) {
	objc.Send[struct{}](g.ID, objc.Sel("setPropertyCallBack:"), value)
}

















