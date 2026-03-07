// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NeuralNetworksSliceGradientOperation] class.
var (
	_NeuralNetworksSliceGradientOperationClass     NeuralNetworksSliceGradientOperationClass
	_NeuralNetworksSliceGradientOperationClassOnce sync.Once
)

func getNeuralNetworksSliceGradientOperationClass() NeuralNetworksSliceGradientOperationClass {
	_NeuralNetworksSliceGradientOperationClassOnce.Do(func() {
		_NeuralNetworksSliceGradientOperationClass = NeuralNetworksSliceGradientOperationClass{class: objc.GetClass("NeuralNetworks.SliceGradientOperation")}
	})
	return _NeuralNetworksSliceGradientOperationClass
}

// GetNeuralNetworksSliceGradientOperationClass returns the class object for NeuralNetworks.SliceGradientOperation.
func GetNeuralNetworksSliceGradientOperationClass() NeuralNetworksSliceGradientOperationClass {
	return getNeuralNetworksSliceGradientOperationClass()
}

type NeuralNetworksSliceGradientOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksSliceGradientOperationClass) Alloc() NeuralNetworksSliceGradientOperation {
	rv := objc.Send[NeuralNetworksSliceGradientOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.SliceGradientOperation
type NeuralNetworksSliceGradientOperation struct {
	NeuralNetworksBaseSliceOperation
}

// NeuralNetworksSliceGradientOperationFromID constructs a [NeuralNetworksSliceGradientOperation] from an objc.ID.
func NeuralNetworksSliceGradientOperationFromID(id objc.ID) NeuralNetworksSliceGradientOperation {
	return NeuralNetworksSliceGradientOperation{NeuralNetworksBaseSliceOperation: NeuralNetworksBaseSliceOperationFromID(id)}
}
// Ensure NeuralNetworksSliceGradientOperation implements INeuralNetworksSliceGradientOperation.
var _ INeuralNetworksSliceGradientOperation = NeuralNetworksSliceGradientOperation{}





// An interface definition for the [NeuralNetworksSliceGradientOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.SliceGradientOperation
type INeuralNetworksSliceGradientOperation interface {
	INeuralNetworksBaseSliceOperation
}





// Init initializes the instance.
func (n NeuralNetworksSliceGradientOperation) Init() NeuralNetworksSliceGradientOperation {
	rv := objc.Send[NeuralNetworksSliceGradientOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksSliceGradientOperation) Autorelease() NeuralNetworksSliceGradientOperation {
	rv := objc.Send[NeuralNetworksSliceGradientOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksSliceGradientOperation creates a new NeuralNetworksSliceGradientOperation instance.
func NewNeuralNetworksSliceGradientOperation() NeuralNetworksSliceGradientOperation {
	class := getNeuralNetworksSliceGradientOperationClass()
	rv := objc.Send[NeuralNetworksSliceGradientOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































