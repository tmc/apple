// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NeuralNetworksSoftmaxOperation] class.
var (
	_NeuralNetworksSoftmaxOperationClass     NeuralNetworksSoftmaxOperationClass
	_NeuralNetworksSoftmaxOperationClassOnce sync.Once
)

func getNeuralNetworksSoftmaxOperationClass() NeuralNetworksSoftmaxOperationClass {
	_NeuralNetworksSoftmaxOperationClassOnce.Do(func() {
		_NeuralNetworksSoftmaxOperationClass = NeuralNetworksSoftmaxOperationClass{class: objc.GetClass("NeuralNetworks.SoftmaxOperation")}
	})
	return _NeuralNetworksSoftmaxOperationClass
}

// GetNeuralNetworksSoftmaxOperationClass returns the class object for NeuralNetworks.SoftmaxOperation.
func GetNeuralNetworksSoftmaxOperationClass() NeuralNetworksSoftmaxOperationClass {
	return getNeuralNetworksSoftmaxOperationClass()
}

type NeuralNetworksSoftmaxOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksSoftmaxOperationClass) Alloc() NeuralNetworksSoftmaxOperation {
	rv := objc.Send[NeuralNetworksSoftmaxOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.SoftmaxOperation
type NeuralNetworksSoftmaxOperation struct {
	NeuralNetworksUnaryElementwiseOperation
}

// NeuralNetworksSoftmaxOperationFromID constructs a [NeuralNetworksSoftmaxOperation] from an objc.ID.
func NeuralNetworksSoftmaxOperationFromID(id objc.ID) NeuralNetworksSoftmaxOperation {
	return NeuralNetworksSoftmaxOperation{NeuralNetworksUnaryElementwiseOperation: NeuralNetworksUnaryElementwiseOperationFromID(id)}
}
// Ensure NeuralNetworksSoftmaxOperation implements INeuralNetworksSoftmaxOperation.
var _ INeuralNetworksSoftmaxOperation = NeuralNetworksSoftmaxOperation{}





// An interface definition for the [NeuralNetworksSoftmaxOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.SoftmaxOperation
type INeuralNetworksSoftmaxOperation interface {
	INeuralNetworksUnaryElementwiseOperation
}





// Init initializes the instance.
func (n NeuralNetworksSoftmaxOperation) Init() NeuralNetworksSoftmaxOperation {
	rv := objc.Send[NeuralNetworksSoftmaxOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksSoftmaxOperation) Autorelease() NeuralNetworksSoftmaxOperation {
	rv := objc.Send[NeuralNetworksSoftmaxOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksSoftmaxOperation creates a new NeuralNetworksSoftmaxOperation instance.
func NewNeuralNetworksSoftmaxOperation() NeuralNetworksSoftmaxOperation {
	class := getNeuralNetworksSoftmaxOperationClass()
	rv := objc.Send[NeuralNetworksSoftmaxOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































