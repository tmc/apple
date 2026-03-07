// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NeuralNetworksBinaryLogicalOperation] class.
var (
	_NeuralNetworksBinaryLogicalOperationClass     NeuralNetworksBinaryLogicalOperationClass
	_NeuralNetworksBinaryLogicalOperationClassOnce sync.Once
)

func getNeuralNetworksBinaryLogicalOperationClass() NeuralNetworksBinaryLogicalOperationClass {
	_NeuralNetworksBinaryLogicalOperationClassOnce.Do(func() {
		_NeuralNetworksBinaryLogicalOperationClass = NeuralNetworksBinaryLogicalOperationClass{class: objc.GetClass("NeuralNetworks.BinaryLogicalOperation")}
	})
	return _NeuralNetworksBinaryLogicalOperationClass
}

// GetNeuralNetworksBinaryLogicalOperationClass returns the class object for NeuralNetworks.BinaryLogicalOperation.
func GetNeuralNetworksBinaryLogicalOperationClass() NeuralNetworksBinaryLogicalOperationClass {
	return getNeuralNetworksBinaryLogicalOperationClass()
}

type NeuralNetworksBinaryLogicalOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksBinaryLogicalOperationClass) Alloc() NeuralNetworksBinaryLogicalOperation {
	rv := objc.Send[NeuralNetworksBinaryLogicalOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.BinaryLogicalOperation
type NeuralNetworksBinaryLogicalOperation struct {
	NeuralNetworksBaseBinaryElementwiseOperation
}

// NeuralNetworksBinaryLogicalOperationFromID constructs a [NeuralNetworksBinaryLogicalOperation] from an objc.ID.
func NeuralNetworksBinaryLogicalOperationFromID(id objc.ID) NeuralNetworksBinaryLogicalOperation {
	return NeuralNetworksBinaryLogicalOperation{NeuralNetworksBaseBinaryElementwiseOperation: NeuralNetworksBaseBinaryElementwiseOperationFromID(id)}
}
// Ensure NeuralNetworksBinaryLogicalOperation implements INeuralNetworksBinaryLogicalOperation.
var _ INeuralNetworksBinaryLogicalOperation = NeuralNetworksBinaryLogicalOperation{}





// An interface definition for the [NeuralNetworksBinaryLogicalOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.BinaryLogicalOperation
type INeuralNetworksBinaryLogicalOperation interface {
	INeuralNetworksBaseBinaryElementwiseOperation
}





// Init initializes the instance.
func (n NeuralNetworksBinaryLogicalOperation) Init() NeuralNetworksBinaryLogicalOperation {
	rv := objc.Send[NeuralNetworksBinaryLogicalOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksBinaryLogicalOperation) Autorelease() NeuralNetworksBinaryLogicalOperation {
	rv := objc.Send[NeuralNetworksBinaryLogicalOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksBinaryLogicalOperation creates a new NeuralNetworksBinaryLogicalOperation instance.
func NewNeuralNetworksBinaryLogicalOperation() NeuralNetworksBinaryLogicalOperation {
	class := getNeuralNetworksBinaryLogicalOperationClass()
	rv := objc.Send[NeuralNetworksBinaryLogicalOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































