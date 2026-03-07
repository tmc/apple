// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NeuralNetworksBatchNormGradientOperation] class.
var (
	_NeuralNetworksBatchNormGradientOperationClass     NeuralNetworksBatchNormGradientOperationClass
	_NeuralNetworksBatchNormGradientOperationClassOnce sync.Once
)

func getNeuralNetworksBatchNormGradientOperationClass() NeuralNetworksBatchNormGradientOperationClass {
	_NeuralNetworksBatchNormGradientOperationClassOnce.Do(func() {
		_NeuralNetworksBatchNormGradientOperationClass = NeuralNetworksBatchNormGradientOperationClass{class: objc.GetClass("NeuralNetworks.BatchNormGradientOperation")}
	})
	return _NeuralNetworksBatchNormGradientOperationClass
}

// GetNeuralNetworksBatchNormGradientOperationClass returns the class object for NeuralNetworks.BatchNormGradientOperation.
func GetNeuralNetworksBatchNormGradientOperationClass() NeuralNetworksBatchNormGradientOperationClass {
	return getNeuralNetworksBatchNormGradientOperationClass()
}

type NeuralNetworksBatchNormGradientOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksBatchNormGradientOperationClass) Alloc() NeuralNetworksBatchNormGradientOperation {
	rv := objc.Send[NeuralNetworksBatchNormGradientOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.BatchNormGradientOperation
type NeuralNetworksBatchNormGradientOperation struct {
	NeuralNetworksBaseNormalizeOperation
}

// NeuralNetworksBatchNormGradientOperationFromID constructs a [NeuralNetworksBatchNormGradientOperation] from an objc.ID.
func NeuralNetworksBatchNormGradientOperationFromID(id objc.ID) NeuralNetworksBatchNormGradientOperation {
	return NeuralNetworksBatchNormGradientOperation{NeuralNetworksBaseNormalizeOperation: NeuralNetworksBaseNormalizeOperationFromID(id)}
}
// Ensure NeuralNetworksBatchNormGradientOperation implements INeuralNetworksBatchNormGradientOperation.
var _ INeuralNetworksBatchNormGradientOperation = NeuralNetworksBatchNormGradientOperation{}





// An interface definition for the [NeuralNetworksBatchNormGradientOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.BatchNormGradientOperation
type INeuralNetworksBatchNormGradientOperation interface {
	INeuralNetworksBaseNormalizeOperation
}





// Init initializes the instance.
func (n NeuralNetworksBatchNormGradientOperation) Init() NeuralNetworksBatchNormGradientOperation {
	rv := objc.Send[NeuralNetworksBatchNormGradientOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksBatchNormGradientOperation) Autorelease() NeuralNetworksBatchNormGradientOperation {
	rv := objc.Send[NeuralNetworksBatchNormGradientOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksBatchNormGradientOperation creates a new NeuralNetworksBatchNormGradientOperation instance.
func NewNeuralNetworksBatchNormGradientOperation() NeuralNetworksBatchNormGradientOperation {
	class := getNeuralNetworksBatchNormGradientOperationClass()
	rv := objc.Send[NeuralNetworksBatchNormGradientOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































