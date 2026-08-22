// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol that defines methods that determine whether and when neural network training parameters are updated.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNTrainableNode
type MPSNNTrainableNode interface {
	objectivec.IObject

	// trainingStyle protocol.
	//
	// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNTrainableNode/trainingStyle
	TrainingStyle() MPSNNTrainingStyle
	SetTrainingStyle(value MPSNNTrainingStyle)
}

// MPSNNTrainableNodeObject wraps an existing Objective-C object that conforms to the MPSNNTrainableNode protocol.
type MPSNNTrainableNodeObject struct {
	objectivec.Object
}

func (o MPSNNTrainableNodeObject) BaseObject() objectivec.Object {
	return o.Object
}

// MPSNNTrainableNodeObjectFromID constructs a [MPSNNTrainableNodeObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MPSNNTrainableNodeObjectFromID(id objc.ID) MPSNNTrainableNodeObject {
	return MPSNNTrainableNodeObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNTrainableNode/trainingStyle
func (o MPSNNTrainableNodeObject) TrainingStyle() MPSNNTrainingStyle {
	rv := objc.Send[MPSNNTrainingStyle](o.ID, objc.Sel("trainingStyle"))
	return MPSNNTrainingStyle(rv)
}

func (o MPSNNTrainableNodeObject) SetTrainingStyle(value MPSNNTrainingStyle) {
	objc.Send[struct{}](o.ID, objc.Sel("setTrainingStyle:"), value)
}
