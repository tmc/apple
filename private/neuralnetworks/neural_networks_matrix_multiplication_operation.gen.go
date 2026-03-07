// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksMatrixMultiplicationOperation] class.
var (
	_NeuralNetworksMatrixMultiplicationOperationClass     NeuralNetworksMatrixMultiplicationOperationClass
	_NeuralNetworksMatrixMultiplicationOperationClassOnce sync.Once
)

func getNeuralNetworksMatrixMultiplicationOperationClass() NeuralNetworksMatrixMultiplicationOperationClass {
	_NeuralNetworksMatrixMultiplicationOperationClassOnce.Do(func() {
		_NeuralNetworksMatrixMultiplicationOperationClass = NeuralNetworksMatrixMultiplicationOperationClass{class: objc.GetClass("NeuralNetworks.MatrixMultiplicationOperation")}
	})
	return _NeuralNetworksMatrixMultiplicationOperationClass
}

// GetNeuralNetworksMatrixMultiplicationOperationClass returns the class object for NeuralNetworks.MatrixMultiplicationOperation.
func GetNeuralNetworksMatrixMultiplicationOperationClass() NeuralNetworksMatrixMultiplicationOperationClass {
	return getNeuralNetworksMatrixMultiplicationOperationClass()
}

type NeuralNetworksMatrixMultiplicationOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksMatrixMultiplicationOperationClass) Alloc() NeuralNetworksMatrixMultiplicationOperation {
	rv := objc.Send[NeuralNetworksMatrixMultiplicationOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.MatrixMultiplicationOperation
type NeuralNetworksMatrixMultiplicationOperation struct {
	objectivec.Object
}

// NeuralNetworksMatrixMultiplicationOperationFromID constructs a [NeuralNetworksMatrixMultiplicationOperation] from an objc.ID.
func NeuralNetworksMatrixMultiplicationOperationFromID(id objc.ID) NeuralNetworksMatrixMultiplicationOperation {
	return NeuralNetworksMatrixMultiplicationOperation{objectivec.Object{id}}
}
// Ensure NeuralNetworksMatrixMultiplicationOperation implements INeuralNetworksMatrixMultiplicationOperation.
var _ INeuralNetworksMatrixMultiplicationOperation = NeuralNetworksMatrixMultiplicationOperation{}





// An interface definition for the [NeuralNetworksMatrixMultiplicationOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.MatrixMultiplicationOperation
type INeuralNetworksMatrixMultiplicationOperation interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksMatrixMultiplicationOperation) Init() NeuralNetworksMatrixMultiplicationOperation {
	rv := objc.Send[NeuralNetworksMatrixMultiplicationOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksMatrixMultiplicationOperation) Autorelease() NeuralNetworksMatrixMultiplicationOperation {
	rv := objc.Send[NeuralNetworksMatrixMultiplicationOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksMatrixMultiplicationOperation creates a new NeuralNetworksMatrixMultiplicationOperation instance.
func NewNeuralNetworksMatrixMultiplicationOperation() NeuralNetworksMatrixMultiplicationOperation {
	class := getNeuralNetworksMatrixMultiplicationOperationClass()
	rv := objc.Send[NeuralNetworksMatrixMultiplicationOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































