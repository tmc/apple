// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NeuralNetworksUnaryElementwiseOperation] class.
var (
	_NeuralNetworksUnaryElementwiseOperationClass     NeuralNetworksUnaryElementwiseOperationClass
	_NeuralNetworksUnaryElementwiseOperationClassOnce sync.Once
)

func getNeuralNetworksUnaryElementwiseOperationClass() NeuralNetworksUnaryElementwiseOperationClass {
	_NeuralNetworksUnaryElementwiseOperationClassOnce.Do(func() {
		_NeuralNetworksUnaryElementwiseOperationClass = NeuralNetworksUnaryElementwiseOperationClass{class: objc.GetClass("NeuralNetworks.UnaryElementwiseOperation")}
	})
	return _NeuralNetworksUnaryElementwiseOperationClass
}

// GetNeuralNetworksUnaryElementwiseOperationClass returns the class object for NeuralNetworks.UnaryElementwiseOperation.
func GetNeuralNetworksUnaryElementwiseOperationClass() NeuralNetworksUnaryElementwiseOperationClass {
	return getNeuralNetworksUnaryElementwiseOperationClass()
}

type NeuralNetworksUnaryElementwiseOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksUnaryElementwiseOperationClass) Alloc() NeuralNetworksUnaryElementwiseOperation {
	rv := objc.Send[NeuralNetworksUnaryElementwiseOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.UnaryElementwiseOperation
type NeuralNetworksUnaryElementwiseOperation struct {
	NeuralNetworksBaseUnaryElementwiseOperation
}

// NeuralNetworksUnaryElementwiseOperationFromID constructs a [NeuralNetworksUnaryElementwiseOperation] from an objc.ID.
func NeuralNetworksUnaryElementwiseOperationFromID(id objc.ID) NeuralNetworksUnaryElementwiseOperation {
	return NeuralNetworksUnaryElementwiseOperation{NeuralNetworksBaseUnaryElementwiseOperation: NeuralNetworksBaseUnaryElementwiseOperationFromID(id)}
}
// Ensure NeuralNetworksUnaryElementwiseOperation implements INeuralNetworksUnaryElementwiseOperation.
var _ INeuralNetworksUnaryElementwiseOperation = NeuralNetworksUnaryElementwiseOperation{}





// An interface definition for the [NeuralNetworksUnaryElementwiseOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.UnaryElementwiseOperation
type INeuralNetworksUnaryElementwiseOperation interface {
	INeuralNetworksBaseUnaryElementwiseOperation
}





// Init initializes the instance.
func (n NeuralNetworksUnaryElementwiseOperation) Init() NeuralNetworksUnaryElementwiseOperation {
	rv := objc.Send[NeuralNetworksUnaryElementwiseOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksUnaryElementwiseOperation) Autorelease() NeuralNetworksUnaryElementwiseOperation {
	rv := objc.Send[NeuralNetworksUnaryElementwiseOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksUnaryElementwiseOperation creates a new NeuralNetworksUnaryElementwiseOperation instance.
func NewNeuralNetworksUnaryElementwiseOperation() NeuralNetworksUnaryElementwiseOperation {
	class := getNeuralNetworksUnaryElementwiseOperationClass()
	rv := objc.Send[NeuralNetworksUnaryElementwiseOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































