// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NeuralNetworksLogicalReductionOperation] class.
var (
	_NeuralNetworksLogicalReductionOperationClass     NeuralNetworksLogicalReductionOperationClass
	_NeuralNetworksLogicalReductionOperationClassOnce sync.Once
)

func getNeuralNetworksLogicalReductionOperationClass() NeuralNetworksLogicalReductionOperationClass {
	_NeuralNetworksLogicalReductionOperationClassOnce.Do(func() {
		_NeuralNetworksLogicalReductionOperationClass = NeuralNetworksLogicalReductionOperationClass{class: objc.GetClass("NeuralNetworks.LogicalReductionOperation")}
	})
	return _NeuralNetworksLogicalReductionOperationClass
}

// GetNeuralNetworksLogicalReductionOperationClass returns the class object for NeuralNetworks.LogicalReductionOperation.
func GetNeuralNetworksLogicalReductionOperationClass() NeuralNetworksLogicalReductionOperationClass {
	return getNeuralNetworksLogicalReductionOperationClass()
}

type NeuralNetworksLogicalReductionOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksLogicalReductionOperationClass) Alloc() NeuralNetworksLogicalReductionOperation {
	rv := objc.Send[NeuralNetworksLogicalReductionOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.LogicalReductionOperation
type NeuralNetworksLogicalReductionOperation struct {
	NeuralNetworksBaseReductionOperation
}

// NeuralNetworksLogicalReductionOperationFromID constructs a [NeuralNetworksLogicalReductionOperation] from an objc.ID.
func NeuralNetworksLogicalReductionOperationFromID(id objc.ID) NeuralNetworksLogicalReductionOperation {
	return NeuralNetworksLogicalReductionOperation{NeuralNetworksBaseReductionOperation: NeuralNetworksBaseReductionOperationFromID(id)}
}
// Ensure NeuralNetworksLogicalReductionOperation implements INeuralNetworksLogicalReductionOperation.
var _ INeuralNetworksLogicalReductionOperation = NeuralNetworksLogicalReductionOperation{}





// An interface definition for the [NeuralNetworksLogicalReductionOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.LogicalReductionOperation
type INeuralNetworksLogicalReductionOperation interface {
	INeuralNetworksBaseReductionOperation
}





// Init initializes the instance.
func (n NeuralNetworksLogicalReductionOperation) Init() NeuralNetworksLogicalReductionOperation {
	rv := objc.Send[NeuralNetworksLogicalReductionOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksLogicalReductionOperation) Autorelease() NeuralNetworksLogicalReductionOperation {
	rv := objc.Send[NeuralNetworksLogicalReductionOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksLogicalReductionOperation creates a new NeuralNetworksLogicalReductionOperation instance.
func NewNeuralNetworksLogicalReductionOperation() NeuralNetworksLogicalReductionOperation {
	class := getNeuralNetworksLogicalReductionOperationClass()
	rv := objc.Send[NeuralNetworksLogicalReductionOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































