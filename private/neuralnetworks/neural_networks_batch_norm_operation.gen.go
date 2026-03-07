// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NeuralNetworksBatchNormOperation] class.
var (
	_NeuralNetworksBatchNormOperationClass     NeuralNetworksBatchNormOperationClass
	_NeuralNetworksBatchNormOperationClassOnce sync.Once
)

func getNeuralNetworksBatchNormOperationClass() NeuralNetworksBatchNormOperationClass {
	_NeuralNetworksBatchNormOperationClassOnce.Do(func() {
		_NeuralNetworksBatchNormOperationClass = NeuralNetworksBatchNormOperationClass{class: objc.GetClass("NeuralNetworks.BatchNormOperation")}
	})
	return _NeuralNetworksBatchNormOperationClass
}

// GetNeuralNetworksBatchNormOperationClass returns the class object for NeuralNetworks.BatchNormOperation.
func GetNeuralNetworksBatchNormOperationClass() NeuralNetworksBatchNormOperationClass {
	return getNeuralNetworksBatchNormOperationClass()
}

type NeuralNetworksBatchNormOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksBatchNormOperationClass) Alloc() NeuralNetworksBatchNormOperation {
	rv := objc.Send[NeuralNetworksBatchNormOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.BatchNormOperation
type NeuralNetworksBatchNormOperation struct {
	NeuralNetworksBaseForwardNormalizeOperation
}

// NeuralNetworksBatchNormOperationFromID constructs a [NeuralNetworksBatchNormOperation] from an objc.ID.
func NeuralNetworksBatchNormOperationFromID(id objc.ID) NeuralNetworksBatchNormOperation {
	return NeuralNetworksBatchNormOperation{NeuralNetworksBaseForwardNormalizeOperation: NeuralNetworksBaseForwardNormalizeOperationFromID(id)}
}
// Ensure NeuralNetworksBatchNormOperation implements INeuralNetworksBatchNormOperation.
var _ INeuralNetworksBatchNormOperation = NeuralNetworksBatchNormOperation{}





// An interface definition for the [NeuralNetworksBatchNormOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.BatchNormOperation
type INeuralNetworksBatchNormOperation interface {
	INeuralNetworksBaseForwardNormalizeOperation
}





// Init initializes the instance.
func (n NeuralNetworksBatchNormOperation) Init() NeuralNetworksBatchNormOperation {
	rv := objc.Send[NeuralNetworksBatchNormOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksBatchNormOperation) Autorelease() NeuralNetworksBatchNormOperation {
	rv := objc.Send[NeuralNetworksBatchNormOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksBatchNormOperation creates a new NeuralNetworksBatchNormOperation instance.
func NewNeuralNetworksBatchNormOperation() NeuralNetworksBatchNormOperation {
	class := getNeuralNetworksBatchNormOperationClass()
	rv := objc.Send[NeuralNetworksBatchNormOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































