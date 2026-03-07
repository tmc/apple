// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NeuralNetworksArgumentReductionOperation] class.
var (
	_NeuralNetworksArgumentReductionOperationClass     NeuralNetworksArgumentReductionOperationClass
	_NeuralNetworksArgumentReductionOperationClassOnce sync.Once
)

func getNeuralNetworksArgumentReductionOperationClass() NeuralNetworksArgumentReductionOperationClass {
	_NeuralNetworksArgumentReductionOperationClassOnce.Do(func() {
		_NeuralNetworksArgumentReductionOperationClass = NeuralNetworksArgumentReductionOperationClass{class: objc.GetClass("NeuralNetworks.ArgumentReductionOperation")}
	})
	return _NeuralNetworksArgumentReductionOperationClass
}

// GetNeuralNetworksArgumentReductionOperationClass returns the class object for NeuralNetworks.ArgumentReductionOperation.
func GetNeuralNetworksArgumentReductionOperationClass() NeuralNetworksArgumentReductionOperationClass {
	return getNeuralNetworksArgumentReductionOperationClass()
}

type NeuralNetworksArgumentReductionOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksArgumentReductionOperationClass) Alloc() NeuralNetworksArgumentReductionOperation {
	rv := objc.Send[NeuralNetworksArgumentReductionOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.ArgumentReductionOperation
type NeuralNetworksArgumentReductionOperation struct {
	NeuralNetworksBaseReductionOperation
}

// NeuralNetworksArgumentReductionOperationFromID constructs a [NeuralNetworksArgumentReductionOperation] from an objc.ID.
func NeuralNetworksArgumentReductionOperationFromID(id objc.ID) NeuralNetworksArgumentReductionOperation {
	return NeuralNetworksArgumentReductionOperation{NeuralNetworksBaseReductionOperation: NeuralNetworksBaseReductionOperationFromID(id)}
}
// Ensure NeuralNetworksArgumentReductionOperation implements INeuralNetworksArgumentReductionOperation.
var _ INeuralNetworksArgumentReductionOperation = NeuralNetworksArgumentReductionOperation{}





// An interface definition for the [NeuralNetworksArgumentReductionOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.ArgumentReductionOperation
type INeuralNetworksArgumentReductionOperation interface {
	INeuralNetworksBaseReductionOperation
}





// Init initializes the instance.
func (n NeuralNetworksArgumentReductionOperation) Init() NeuralNetworksArgumentReductionOperation {
	rv := objc.Send[NeuralNetworksArgumentReductionOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksArgumentReductionOperation) Autorelease() NeuralNetworksArgumentReductionOperation {
	rv := objc.Send[NeuralNetworksArgumentReductionOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksArgumentReductionOperation creates a new NeuralNetworksArgumentReductionOperation instance.
func NewNeuralNetworksArgumentReductionOperation() NeuralNetworksArgumentReductionOperation {
	class := getNeuralNetworksArgumentReductionOperationClass()
	rv := objc.Send[NeuralNetworksArgumentReductionOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































