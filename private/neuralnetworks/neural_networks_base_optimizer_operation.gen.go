// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksBaseOptimizerOperation] class.
var (
	_NeuralNetworksBaseOptimizerOperationClass     NeuralNetworksBaseOptimizerOperationClass
	_NeuralNetworksBaseOptimizerOperationClassOnce sync.Once
)

func getNeuralNetworksBaseOptimizerOperationClass() NeuralNetworksBaseOptimizerOperationClass {
	_NeuralNetworksBaseOptimizerOperationClassOnce.Do(func() {
		_NeuralNetworksBaseOptimizerOperationClass = NeuralNetworksBaseOptimizerOperationClass{class: objc.GetClass("NeuralNetworks.BaseOptimizerOperation")}
	})
	return _NeuralNetworksBaseOptimizerOperationClass
}

// GetNeuralNetworksBaseOptimizerOperationClass returns the class object for NeuralNetworks.BaseOptimizerOperation.
func GetNeuralNetworksBaseOptimizerOperationClass() NeuralNetworksBaseOptimizerOperationClass {
	return getNeuralNetworksBaseOptimizerOperationClass()
}

type NeuralNetworksBaseOptimizerOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksBaseOptimizerOperationClass) Alloc() NeuralNetworksBaseOptimizerOperation {
	rv := objc.Send[NeuralNetworksBaseOptimizerOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.BaseOptimizerOperation
type NeuralNetworksBaseOptimizerOperation struct {
	objectivec.Object
}

// NeuralNetworksBaseOptimizerOperationFromID constructs a [NeuralNetworksBaseOptimizerOperation] from an objc.ID.
func NeuralNetworksBaseOptimizerOperationFromID(id objc.ID) NeuralNetworksBaseOptimizerOperation {
	return NeuralNetworksBaseOptimizerOperation{objectivec.Object{id}}
}
// Ensure NeuralNetworksBaseOptimizerOperation implements INeuralNetworksBaseOptimizerOperation.
var _ INeuralNetworksBaseOptimizerOperation = NeuralNetworksBaseOptimizerOperation{}





// An interface definition for the [NeuralNetworksBaseOptimizerOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.BaseOptimizerOperation
type INeuralNetworksBaseOptimizerOperation interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksBaseOptimizerOperation) Init() NeuralNetworksBaseOptimizerOperation {
	rv := objc.Send[NeuralNetworksBaseOptimizerOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksBaseOptimizerOperation) Autorelease() NeuralNetworksBaseOptimizerOperation {
	rv := objc.Send[NeuralNetworksBaseOptimizerOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksBaseOptimizerOperation creates a new NeuralNetworksBaseOptimizerOperation instance.
func NewNeuralNetworksBaseOptimizerOperation() NeuralNetworksBaseOptimizerOperation {
	class := getNeuralNetworksBaseOptimizerOperationClass()
	rv := objc.Send[NeuralNetworksBaseOptimizerOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































