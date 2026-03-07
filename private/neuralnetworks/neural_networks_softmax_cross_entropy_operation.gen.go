// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NeuralNetworksSoftmaxCrossEntropyOperation] class.
var (
	_NeuralNetworksSoftmaxCrossEntropyOperationClass     NeuralNetworksSoftmaxCrossEntropyOperationClass
	_NeuralNetworksSoftmaxCrossEntropyOperationClassOnce sync.Once
)

func getNeuralNetworksSoftmaxCrossEntropyOperationClass() NeuralNetworksSoftmaxCrossEntropyOperationClass {
	_NeuralNetworksSoftmaxCrossEntropyOperationClassOnce.Do(func() {
		_NeuralNetworksSoftmaxCrossEntropyOperationClass = NeuralNetworksSoftmaxCrossEntropyOperationClass{class: objc.GetClass("NeuralNetworks.SoftmaxCrossEntropyOperation")}
	})
	return _NeuralNetworksSoftmaxCrossEntropyOperationClass
}

// GetNeuralNetworksSoftmaxCrossEntropyOperationClass returns the class object for NeuralNetworks.SoftmaxCrossEntropyOperation.
func GetNeuralNetworksSoftmaxCrossEntropyOperationClass() NeuralNetworksSoftmaxCrossEntropyOperationClass {
	return getNeuralNetworksSoftmaxCrossEntropyOperationClass()
}

type NeuralNetworksSoftmaxCrossEntropyOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksSoftmaxCrossEntropyOperationClass) Alloc() NeuralNetworksSoftmaxCrossEntropyOperation {
	rv := objc.Send[NeuralNetworksSoftmaxCrossEntropyOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.SoftmaxCrossEntropyOperation
type NeuralNetworksSoftmaxCrossEntropyOperation struct {
	NeuralNetworksLossOperation
}

// NeuralNetworksSoftmaxCrossEntropyOperationFromID constructs a [NeuralNetworksSoftmaxCrossEntropyOperation] from an objc.ID.
func NeuralNetworksSoftmaxCrossEntropyOperationFromID(id objc.ID) NeuralNetworksSoftmaxCrossEntropyOperation {
	return NeuralNetworksSoftmaxCrossEntropyOperation{NeuralNetworksLossOperation: NeuralNetworksLossOperationFromID(id)}
}
// Ensure NeuralNetworksSoftmaxCrossEntropyOperation implements INeuralNetworksSoftmaxCrossEntropyOperation.
var _ INeuralNetworksSoftmaxCrossEntropyOperation = NeuralNetworksSoftmaxCrossEntropyOperation{}





// An interface definition for the [NeuralNetworksSoftmaxCrossEntropyOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.SoftmaxCrossEntropyOperation
type INeuralNetworksSoftmaxCrossEntropyOperation interface {
	INeuralNetworksLossOperation
}





// Init initializes the instance.
func (n NeuralNetworksSoftmaxCrossEntropyOperation) Init() NeuralNetworksSoftmaxCrossEntropyOperation {
	rv := objc.Send[NeuralNetworksSoftmaxCrossEntropyOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksSoftmaxCrossEntropyOperation) Autorelease() NeuralNetworksSoftmaxCrossEntropyOperation {
	rv := objc.Send[NeuralNetworksSoftmaxCrossEntropyOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksSoftmaxCrossEntropyOperation creates a new NeuralNetworksSoftmaxCrossEntropyOperation instance.
func NewNeuralNetworksSoftmaxCrossEntropyOperation() NeuralNetworksSoftmaxCrossEntropyOperation {
	class := getNeuralNetworksSoftmaxCrossEntropyOperationClass()
	rv := objc.Send[NeuralNetworksSoftmaxCrossEntropyOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































