// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NeuralNetworksBinaryComparisonOperation] class.
var (
	_NeuralNetworksBinaryComparisonOperationClass     NeuralNetworksBinaryComparisonOperationClass
	_NeuralNetworksBinaryComparisonOperationClassOnce sync.Once
)

func getNeuralNetworksBinaryComparisonOperationClass() NeuralNetworksBinaryComparisonOperationClass {
	_NeuralNetworksBinaryComparisonOperationClassOnce.Do(func() {
		_NeuralNetworksBinaryComparisonOperationClass = NeuralNetworksBinaryComparisonOperationClass{class: objc.GetClass("NeuralNetworks.BinaryComparisonOperation")}
	})
	return _NeuralNetworksBinaryComparisonOperationClass
}

// GetNeuralNetworksBinaryComparisonOperationClass returns the class object for NeuralNetworks.BinaryComparisonOperation.
func GetNeuralNetworksBinaryComparisonOperationClass() NeuralNetworksBinaryComparisonOperationClass {
	return getNeuralNetworksBinaryComparisonOperationClass()
}

type NeuralNetworksBinaryComparisonOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksBinaryComparisonOperationClass) Alloc() NeuralNetworksBinaryComparisonOperation {
	rv := objc.Send[NeuralNetworksBinaryComparisonOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.BinaryComparisonOperation
type NeuralNetworksBinaryComparisonOperation struct {
	NeuralNetworksBaseBinaryElementwiseOperation
}

// NeuralNetworksBinaryComparisonOperationFromID constructs a [NeuralNetworksBinaryComparisonOperation] from an objc.ID.
func NeuralNetworksBinaryComparisonOperationFromID(id objc.ID) NeuralNetworksBinaryComparisonOperation {
	return NeuralNetworksBinaryComparisonOperation{NeuralNetworksBaseBinaryElementwiseOperation: NeuralNetworksBaseBinaryElementwiseOperationFromID(id)}
}
// Ensure NeuralNetworksBinaryComparisonOperation implements INeuralNetworksBinaryComparisonOperation.
var _ INeuralNetworksBinaryComparisonOperation = NeuralNetworksBinaryComparisonOperation{}





// An interface definition for the [NeuralNetworksBinaryComparisonOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.BinaryComparisonOperation
type INeuralNetworksBinaryComparisonOperation interface {
	INeuralNetworksBaseBinaryElementwiseOperation
}





// Init initializes the instance.
func (n NeuralNetworksBinaryComparisonOperation) Init() NeuralNetworksBinaryComparisonOperation {
	rv := objc.Send[NeuralNetworksBinaryComparisonOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksBinaryComparisonOperation) Autorelease() NeuralNetworksBinaryComparisonOperation {
	rv := objc.Send[NeuralNetworksBinaryComparisonOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksBinaryComparisonOperation creates a new NeuralNetworksBinaryComparisonOperation instance.
func NewNeuralNetworksBinaryComparisonOperation() NeuralNetworksBinaryComparisonOperation {
	class := getNeuralNetworksBinaryComparisonOperationClass()
	rv := objc.Send[NeuralNetworksBinaryComparisonOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































