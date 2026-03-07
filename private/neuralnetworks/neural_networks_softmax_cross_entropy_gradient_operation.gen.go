// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NeuralNetworksSoftmaxCrossEntropyGradientOperation] class.
var (
	_NeuralNetworksSoftmaxCrossEntropyGradientOperationClass     NeuralNetworksSoftmaxCrossEntropyGradientOperationClass
	_NeuralNetworksSoftmaxCrossEntropyGradientOperationClassOnce sync.Once
)

func getNeuralNetworksSoftmaxCrossEntropyGradientOperationClass() NeuralNetworksSoftmaxCrossEntropyGradientOperationClass {
	_NeuralNetworksSoftmaxCrossEntropyGradientOperationClassOnce.Do(func() {
		_NeuralNetworksSoftmaxCrossEntropyGradientOperationClass = NeuralNetworksSoftmaxCrossEntropyGradientOperationClass{class: objc.GetClass("NeuralNetworks.SoftmaxCrossEntropyGradientOperation")}
	})
	return _NeuralNetworksSoftmaxCrossEntropyGradientOperationClass
}

// GetNeuralNetworksSoftmaxCrossEntropyGradientOperationClass returns the class object for NeuralNetworks.SoftmaxCrossEntropyGradientOperation.
func GetNeuralNetworksSoftmaxCrossEntropyGradientOperationClass() NeuralNetworksSoftmaxCrossEntropyGradientOperationClass {
	return getNeuralNetworksSoftmaxCrossEntropyGradientOperationClass()
}

type NeuralNetworksSoftmaxCrossEntropyGradientOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksSoftmaxCrossEntropyGradientOperationClass) Alloc() NeuralNetworksSoftmaxCrossEntropyGradientOperation {
	rv := objc.Send[NeuralNetworksSoftmaxCrossEntropyGradientOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.SoftmaxCrossEntropyGradientOperation
type NeuralNetworksSoftmaxCrossEntropyGradientOperation struct {
	NeuralNetworksBaseLossOperation
}

// NeuralNetworksSoftmaxCrossEntropyGradientOperationFromID constructs a [NeuralNetworksSoftmaxCrossEntropyGradientOperation] from an objc.ID.
func NeuralNetworksSoftmaxCrossEntropyGradientOperationFromID(id objc.ID) NeuralNetworksSoftmaxCrossEntropyGradientOperation {
	return NeuralNetworksSoftmaxCrossEntropyGradientOperation{NeuralNetworksBaseLossOperation: NeuralNetworksBaseLossOperationFromID(id)}
}
// Ensure NeuralNetworksSoftmaxCrossEntropyGradientOperation implements INeuralNetworksSoftmaxCrossEntropyGradientOperation.
var _ INeuralNetworksSoftmaxCrossEntropyGradientOperation = NeuralNetworksSoftmaxCrossEntropyGradientOperation{}





// An interface definition for the [NeuralNetworksSoftmaxCrossEntropyGradientOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.SoftmaxCrossEntropyGradientOperation
type INeuralNetworksSoftmaxCrossEntropyGradientOperation interface {
	INeuralNetworksBaseLossOperation
}





// Init initializes the instance.
func (n NeuralNetworksSoftmaxCrossEntropyGradientOperation) Init() NeuralNetworksSoftmaxCrossEntropyGradientOperation {
	rv := objc.Send[NeuralNetworksSoftmaxCrossEntropyGradientOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksSoftmaxCrossEntropyGradientOperation) Autorelease() NeuralNetworksSoftmaxCrossEntropyGradientOperation {
	rv := objc.Send[NeuralNetworksSoftmaxCrossEntropyGradientOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksSoftmaxCrossEntropyGradientOperation creates a new NeuralNetworksSoftmaxCrossEntropyGradientOperation instance.
func NewNeuralNetworksSoftmaxCrossEntropyGradientOperation() NeuralNetworksSoftmaxCrossEntropyGradientOperation {
	class := getNeuralNetworksSoftmaxCrossEntropyGradientOperationClass()
	rv := objc.Send[NeuralNetworksSoftmaxCrossEntropyGradientOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































