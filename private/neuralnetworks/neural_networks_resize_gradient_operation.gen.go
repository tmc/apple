// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NeuralNetworksResizeGradientOperation] class.
var (
	_NeuralNetworksResizeGradientOperationClass     NeuralNetworksResizeGradientOperationClass
	_NeuralNetworksResizeGradientOperationClassOnce sync.Once
)

func getNeuralNetworksResizeGradientOperationClass() NeuralNetworksResizeGradientOperationClass {
	_NeuralNetworksResizeGradientOperationClassOnce.Do(func() {
		_NeuralNetworksResizeGradientOperationClass = NeuralNetworksResizeGradientOperationClass{class: objc.GetClass("NeuralNetworks.ResizeGradientOperation")}
	})
	return _NeuralNetworksResizeGradientOperationClass
}

// GetNeuralNetworksResizeGradientOperationClass returns the class object for NeuralNetworks.ResizeGradientOperation.
func GetNeuralNetworksResizeGradientOperationClass() NeuralNetworksResizeGradientOperationClass {
	return getNeuralNetworksResizeGradientOperationClass()
}

type NeuralNetworksResizeGradientOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksResizeGradientOperationClass) Alloc() NeuralNetworksResizeGradientOperation {
	rv := objc.Send[NeuralNetworksResizeGradientOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.ResizeGradientOperation
type NeuralNetworksResizeGradientOperation struct {
	NeuralNetworksBaseResizeOperation
}

// NeuralNetworksResizeGradientOperationFromID constructs a [NeuralNetworksResizeGradientOperation] from an objc.ID.
func NeuralNetworksResizeGradientOperationFromID(id objc.ID) NeuralNetworksResizeGradientOperation {
	return NeuralNetworksResizeGradientOperation{NeuralNetworksBaseResizeOperation: NeuralNetworksBaseResizeOperationFromID(id)}
}
// Ensure NeuralNetworksResizeGradientOperation implements INeuralNetworksResizeGradientOperation.
var _ INeuralNetworksResizeGradientOperation = NeuralNetworksResizeGradientOperation{}





// An interface definition for the [NeuralNetworksResizeGradientOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.ResizeGradientOperation
type INeuralNetworksResizeGradientOperation interface {
	INeuralNetworksBaseResizeOperation
}





// Init initializes the instance.
func (n NeuralNetworksResizeGradientOperation) Init() NeuralNetworksResizeGradientOperation {
	rv := objc.Send[NeuralNetworksResizeGradientOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksResizeGradientOperation) Autorelease() NeuralNetworksResizeGradientOperation {
	rv := objc.Send[NeuralNetworksResizeGradientOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksResizeGradientOperation creates a new NeuralNetworksResizeGradientOperation instance.
func NewNeuralNetworksResizeGradientOperation() NeuralNetworksResizeGradientOperation {
	class := getNeuralNetworksResizeGradientOperationClass()
	rv := objc.Send[NeuralNetworksResizeGradientOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































