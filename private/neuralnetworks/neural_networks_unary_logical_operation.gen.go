// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NeuralNetworksUnaryLogicalOperation] class.
var (
	_NeuralNetworksUnaryLogicalOperationClass     NeuralNetworksUnaryLogicalOperationClass
	_NeuralNetworksUnaryLogicalOperationClassOnce sync.Once
)

func getNeuralNetworksUnaryLogicalOperationClass() NeuralNetworksUnaryLogicalOperationClass {
	_NeuralNetworksUnaryLogicalOperationClassOnce.Do(func() {
		_NeuralNetworksUnaryLogicalOperationClass = NeuralNetworksUnaryLogicalOperationClass{class: objc.GetClass("NeuralNetworks.UnaryLogicalOperation")}
	})
	return _NeuralNetworksUnaryLogicalOperationClass
}

// GetNeuralNetworksUnaryLogicalOperationClass returns the class object for NeuralNetworks.UnaryLogicalOperation.
func GetNeuralNetworksUnaryLogicalOperationClass() NeuralNetworksUnaryLogicalOperationClass {
	return getNeuralNetworksUnaryLogicalOperationClass()
}

type NeuralNetworksUnaryLogicalOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksUnaryLogicalOperationClass) Alloc() NeuralNetworksUnaryLogicalOperation {
	rv := objc.Send[NeuralNetworksUnaryLogicalOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.UnaryLogicalOperation
type NeuralNetworksUnaryLogicalOperation struct {
	NeuralNetworksBaseUnaryElementwiseOperation
}

// NeuralNetworksUnaryLogicalOperationFromID constructs a [NeuralNetworksUnaryLogicalOperation] from an objc.ID.
func NeuralNetworksUnaryLogicalOperationFromID(id objc.ID) NeuralNetworksUnaryLogicalOperation {
	return NeuralNetworksUnaryLogicalOperation{NeuralNetworksBaseUnaryElementwiseOperation: NeuralNetworksBaseUnaryElementwiseOperationFromID(id)}
}
// Ensure NeuralNetworksUnaryLogicalOperation implements INeuralNetworksUnaryLogicalOperation.
var _ INeuralNetworksUnaryLogicalOperation = NeuralNetworksUnaryLogicalOperation{}





// An interface definition for the [NeuralNetworksUnaryLogicalOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.UnaryLogicalOperation
type INeuralNetworksUnaryLogicalOperation interface {
	INeuralNetworksBaseUnaryElementwiseOperation
}





// Init initializes the instance.
func (n NeuralNetworksUnaryLogicalOperation) Init() NeuralNetworksUnaryLogicalOperation {
	rv := objc.Send[NeuralNetworksUnaryLogicalOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksUnaryLogicalOperation) Autorelease() NeuralNetworksUnaryLogicalOperation {
	rv := objc.Send[NeuralNetworksUnaryLogicalOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksUnaryLogicalOperation creates a new NeuralNetworksUnaryLogicalOperation instance.
func NewNeuralNetworksUnaryLogicalOperation() NeuralNetworksUnaryLogicalOperation {
	class := getNeuralNetworksUnaryLogicalOperationClass()
	rv := objc.Send[NeuralNetworksUnaryLogicalOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































