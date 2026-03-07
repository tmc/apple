// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NeuralNetworksUnaryArithmeticOperation] class.
var (
	_NeuralNetworksUnaryArithmeticOperationClass     NeuralNetworksUnaryArithmeticOperationClass
	_NeuralNetworksUnaryArithmeticOperationClassOnce sync.Once
)

func getNeuralNetworksUnaryArithmeticOperationClass() NeuralNetworksUnaryArithmeticOperationClass {
	_NeuralNetworksUnaryArithmeticOperationClassOnce.Do(func() {
		_NeuralNetworksUnaryArithmeticOperationClass = NeuralNetworksUnaryArithmeticOperationClass{class: objc.GetClass("NeuralNetworks.UnaryArithmeticOperation")}
	})
	return _NeuralNetworksUnaryArithmeticOperationClass
}

// GetNeuralNetworksUnaryArithmeticOperationClass returns the class object for NeuralNetworks.UnaryArithmeticOperation.
func GetNeuralNetworksUnaryArithmeticOperationClass() NeuralNetworksUnaryArithmeticOperationClass {
	return getNeuralNetworksUnaryArithmeticOperationClass()
}

type NeuralNetworksUnaryArithmeticOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksUnaryArithmeticOperationClass) Alloc() NeuralNetworksUnaryArithmeticOperation {
	rv := objc.Send[NeuralNetworksUnaryArithmeticOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.UnaryArithmeticOperation
type NeuralNetworksUnaryArithmeticOperation struct {
	NeuralNetworksUnaryElementwiseOperation
}

// NeuralNetworksUnaryArithmeticOperationFromID constructs a [NeuralNetworksUnaryArithmeticOperation] from an objc.ID.
func NeuralNetworksUnaryArithmeticOperationFromID(id objc.ID) NeuralNetworksUnaryArithmeticOperation {
	return NeuralNetworksUnaryArithmeticOperation{NeuralNetworksUnaryElementwiseOperation: NeuralNetworksUnaryElementwiseOperationFromID(id)}
}
// Ensure NeuralNetworksUnaryArithmeticOperation implements INeuralNetworksUnaryArithmeticOperation.
var _ INeuralNetworksUnaryArithmeticOperation = NeuralNetworksUnaryArithmeticOperation{}





// An interface definition for the [NeuralNetworksUnaryArithmeticOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.UnaryArithmeticOperation
type INeuralNetworksUnaryArithmeticOperation interface {
	INeuralNetworksUnaryElementwiseOperation
}





// Init initializes the instance.
func (n NeuralNetworksUnaryArithmeticOperation) Init() NeuralNetworksUnaryArithmeticOperation {
	rv := objc.Send[NeuralNetworksUnaryArithmeticOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksUnaryArithmeticOperation) Autorelease() NeuralNetworksUnaryArithmeticOperation {
	rv := objc.Send[NeuralNetworksUnaryArithmeticOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksUnaryArithmeticOperation creates a new NeuralNetworksUnaryArithmeticOperation instance.
func NewNeuralNetworksUnaryArithmeticOperation() NeuralNetworksUnaryArithmeticOperation {
	class := getNeuralNetworksUnaryArithmeticOperationClass()
	rv := objc.Send[NeuralNetworksUnaryArithmeticOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































