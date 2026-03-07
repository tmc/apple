// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NeuralNetworksBinaryElementwiseOperation] class.
var (
	_NeuralNetworksBinaryElementwiseOperationClass     NeuralNetworksBinaryElementwiseOperationClass
	_NeuralNetworksBinaryElementwiseOperationClassOnce sync.Once
)

func getNeuralNetworksBinaryElementwiseOperationClass() NeuralNetworksBinaryElementwiseOperationClass {
	_NeuralNetworksBinaryElementwiseOperationClassOnce.Do(func() {
		_NeuralNetworksBinaryElementwiseOperationClass = NeuralNetworksBinaryElementwiseOperationClass{class: objc.GetClass("NeuralNetworks.BinaryElementwiseOperation")}
	})
	return _NeuralNetworksBinaryElementwiseOperationClass
}

// GetNeuralNetworksBinaryElementwiseOperationClass returns the class object for NeuralNetworks.BinaryElementwiseOperation.
func GetNeuralNetworksBinaryElementwiseOperationClass() NeuralNetworksBinaryElementwiseOperationClass {
	return getNeuralNetworksBinaryElementwiseOperationClass()
}

type NeuralNetworksBinaryElementwiseOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksBinaryElementwiseOperationClass) Alloc() NeuralNetworksBinaryElementwiseOperation {
	rv := objc.Send[NeuralNetworksBinaryElementwiseOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.BinaryElementwiseOperation
type NeuralNetworksBinaryElementwiseOperation struct {
	NeuralNetworksBaseBinaryElementwiseOperation
}

// NeuralNetworksBinaryElementwiseOperationFromID constructs a [NeuralNetworksBinaryElementwiseOperation] from an objc.ID.
func NeuralNetworksBinaryElementwiseOperationFromID(id objc.ID) NeuralNetworksBinaryElementwiseOperation {
	return NeuralNetworksBinaryElementwiseOperation{NeuralNetworksBaseBinaryElementwiseOperation: NeuralNetworksBaseBinaryElementwiseOperationFromID(id)}
}
// Ensure NeuralNetworksBinaryElementwiseOperation implements INeuralNetworksBinaryElementwiseOperation.
var _ INeuralNetworksBinaryElementwiseOperation = NeuralNetworksBinaryElementwiseOperation{}





// An interface definition for the [NeuralNetworksBinaryElementwiseOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.BinaryElementwiseOperation
type INeuralNetworksBinaryElementwiseOperation interface {
	INeuralNetworksBaseBinaryElementwiseOperation
}





// Init initializes the instance.
func (n NeuralNetworksBinaryElementwiseOperation) Init() NeuralNetworksBinaryElementwiseOperation {
	rv := objc.Send[NeuralNetworksBinaryElementwiseOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksBinaryElementwiseOperation) Autorelease() NeuralNetworksBinaryElementwiseOperation {
	rv := objc.Send[NeuralNetworksBinaryElementwiseOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksBinaryElementwiseOperation creates a new NeuralNetworksBinaryElementwiseOperation instance.
func NewNeuralNetworksBinaryElementwiseOperation() NeuralNetworksBinaryElementwiseOperation {
	class := getNeuralNetworksBinaryElementwiseOperationClass()
	rv := objc.Send[NeuralNetworksBinaryElementwiseOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































