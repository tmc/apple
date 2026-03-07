// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NeuralNetworksReductionOperation] class.
var (
	_NeuralNetworksReductionOperationClass     NeuralNetworksReductionOperationClass
	_NeuralNetworksReductionOperationClassOnce sync.Once
)

func getNeuralNetworksReductionOperationClass() NeuralNetworksReductionOperationClass {
	_NeuralNetworksReductionOperationClassOnce.Do(func() {
		_NeuralNetworksReductionOperationClass = NeuralNetworksReductionOperationClass{class: objc.GetClass("NeuralNetworks.ReductionOperation")}
	})
	return _NeuralNetworksReductionOperationClass
}

// GetNeuralNetworksReductionOperationClass returns the class object for NeuralNetworks.ReductionOperation.
func GetNeuralNetworksReductionOperationClass() NeuralNetworksReductionOperationClass {
	return getNeuralNetworksReductionOperationClass()
}

type NeuralNetworksReductionOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksReductionOperationClass) Alloc() NeuralNetworksReductionOperation {
	rv := objc.Send[NeuralNetworksReductionOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.ReductionOperation
type NeuralNetworksReductionOperation struct {
	NeuralNetworksBaseReductionOperation
}

// NeuralNetworksReductionOperationFromID constructs a [NeuralNetworksReductionOperation] from an objc.ID.
func NeuralNetworksReductionOperationFromID(id objc.ID) NeuralNetworksReductionOperation {
	return NeuralNetworksReductionOperation{NeuralNetworksBaseReductionOperation: NeuralNetworksBaseReductionOperationFromID(id)}
}
// Ensure NeuralNetworksReductionOperation implements INeuralNetworksReductionOperation.
var _ INeuralNetworksReductionOperation = NeuralNetworksReductionOperation{}





// An interface definition for the [NeuralNetworksReductionOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.ReductionOperation
type INeuralNetworksReductionOperation interface {
	INeuralNetworksBaseReductionOperation
}





// Init initializes the instance.
func (n NeuralNetworksReductionOperation) Init() NeuralNetworksReductionOperation {
	rv := objc.Send[NeuralNetworksReductionOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksReductionOperation) Autorelease() NeuralNetworksReductionOperation {
	rv := objc.Send[NeuralNetworksReductionOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksReductionOperation creates a new NeuralNetworksReductionOperation instance.
func NewNeuralNetworksReductionOperation() NeuralNetworksReductionOperation {
	class := getNeuralNetworksReductionOperationClass()
	rv := objc.Send[NeuralNetworksReductionOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































