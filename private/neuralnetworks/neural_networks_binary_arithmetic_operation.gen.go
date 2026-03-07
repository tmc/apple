// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NeuralNetworksBinaryArithmeticOperation] class.
var (
	_NeuralNetworksBinaryArithmeticOperationClass     NeuralNetworksBinaryArithmeticOperationClass
	_NeuralNetworksBinaryArithmeticOperationClassOnce sync.Once
)

func getNeuralNetworksBinaryArithmeticOperationClass() NeuralNetworksBinaryArithmeticOperationClass {
	_NeuralNetworksBinaryArithmeticOperationClassOnce.Do(func() {
		_NeuralNetworksBinaryArithmeticOperationClass = NeuralNetworksBinaryArithmeticOperationClass{class: objc.GetClass("NeuralNetworks.BinaryArithmeticOperation")}
	})
	return _NeuralNetworksBinaryArithmeticOperationClass
}

// GetNeuralNetworksBinaryArithmeticOperationClass returns the class object for NeuralNetworks.BinaryArithmeticOperation.
func GetNeuralNetworksBinaryArithmeticOperationClass() NeuralNetworksBinaryArithmeticOperationClass {
	return getNeuralNetworksBinaryArithmeticOperationClass()
}

type NeuralNetworksBinaryArithmeticOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksBinaryArithmeticOperationClass) Alloc() NeuralNetworksBinaryArithmeticOperation {
	rv := objc.Send[NeuralNetworksBinaryArithmeticOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.BinaryArithmeticOperation
type NeuralNetworksBinaryArithmeticOperation struct {
	NeuralNetworksBinaryElementwiseOperation
}

// NeuralNetworksBinaryArithmeticOperationFromID constructs a [NeuralNetworksBinaryArithmeticOperation] from an objc.ID.
func NeuralNetworksBinaryArithmeticOperationFromID(id objc.ID) NeuralNetworksBinaryArithmeticOperation {
	return NeuralNetworksBinaryArithmeticOperation{NeuralNetworksBinaryElementwiseOperation: NeuralNetworksBinaryElementwiseOperationFromID(id)}
}
// Ensure NeuralNetworksBinaryArithmeticOperation implements INeuralNetworksBinaryArithmeticOperation.
var _ INeuralNetworksBinaryArithmeticOperation = NeuralNetworksBinaryArithmeticOperation{}





// An interface definition for the [NeuralNetworksBinaryArithmeticOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.BinaryArithmeticOperation
type INeuralNetworksBinaryArithmeticOperation interface {
	INeuralNetworksBinaryElementwiseOperation
}





// Init initializes the instance.
func (n NeuralNetworksBinaryArithmeticOperation) Init() NeuralNetworksBinaryArithmeticOperation {
	rv := objc.Send[NeuralNetworksBinaryArithmeticOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksBinaryArithmeticOperation) Autorelease() NeuralNetworksBinaryArithmeticOperation {
	rv := objc.Send[NeuralNetworksBinaryArithmeticOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksBinaryArithmeticOperation creates a new NeuralNetworksBinaryArithmeticOperation instance.
func NewNeuralNetworksBinaryArithmeticOperation() NeuralNetworksBinaryArithmeticOperation {
	class := getNeuralNetworksBinaryArithmeticOperationClass()
	rv := objc.Send[NeuralNetworksBinaryArithmeticOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































