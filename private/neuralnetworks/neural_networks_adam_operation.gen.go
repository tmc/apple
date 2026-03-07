// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NeuralNetworksAdamOperation] class.
var (
	_NeuralNetworksAdamOperationClass     NeuralNetworksAdamOperationClass
	_NeuralNetworksAdamOperationClassOnce sync.Once
)

func getNeuralNetworksAdamOperationClass() NeuralNetworksAdamOperationClass {
	_NeuralNetworksAdamOperationClassOnce.Do(func() {
		_NeuralNetworksAdamOperationClass = NeuralNetworksAdamOperationClass{class: objc.GetClass("NeuralNetworks.AdamOperation")}
	})
	return _NeuralNetworksAdamOperationClass
}

// GetNeuralNetworksAdamOperationClass returns the class object for NeuralNetworks.AdamOperation.
func GetNeuralNetworksAdamOperationClass() NeuralNetworksAdamOperationClass {
	return getNeuralNetworksAdamOperationClass()
}

type NeuralNetworksAdamOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksAdamOperationClass) Alloc() NeuralNetworksAdamOperation {
	rv := objc.Send[NeuralNetworksAdamOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.AdamOperation
type NeuralNetworksAdamOperation struct {
	NeuralNetworksBaseOptimizerOperation
}

// NeuralNetworksAdamOperationFromID constructs a [NeuralNetworksAdamOperation] from an objc.ID.
func NeuralNetworksAdamOperationFromID(id objc.ID) NeuralNetworksAdamOperation {
	return NeuralNetworksAdamOperation{NeuralNetworksBaseOptimizerOperation: NeuralNetworksBaseOptimizerOperationFromID(id)}
}
// Ensure NeuralNetworksAdamOperation implements INeuralNetworksAdamOperation.
var _ INeuralNetworksAdamOperation = NeuralNetworksAdamOperation{}





// An interface definition for the [NeuralNetworksAdamOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.AdamOperation
type INeuralNetworksAdamOperation interface {
	INeuralNetworksBaseOptimizerOperation
}





// Init initializes the instance.
func (n NeuralNetworksAdamOperation) Init() NeuralNetworksAdamOperation {
	rv := objc.Send[NeuralNetworksAdamOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksAdamOperation) Autorelease() NeuralNetworksAdamOperation {
	rv := objc.Send[NeuralNetworksAdamOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksAdamOperation creates a new NeuralNetworksAdamOperation instance.
func NewNeuralNetworksAdamOperation() NeuralNetworksAdamOperation {
	class := getNeuralNetworksAdamOperationClass()
	rv := objc.Send[NeuralNetworksAdamOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































