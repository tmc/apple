// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NeuralNetworksConvolution2DGradientOperation] class.
var (
	_NeuralNetworksConvolution2DGradientOperationClass     NeuralNetworksConvolution2DGradientOperationClass
	_NeuralNetworksConvolution2DGradientOperationClassOnce sync.Once
)

func getNeuralNetworksConvolution2DGradientOperationClass() NeuralNetworksConvolution2DGradientOperationClass {
	_NeuralNetworksConvolution2DGradientOperationClassOnce.Do(func() {
		_NeuralNetworksConvolution2DGradientOperationClass = NeuralNetworksConvolution2DGradientOperationClass{class: objc.GetClass("NeuralNetworks.Convolution2DGradientOperation")}
	})
	return _NeuralNetworksConvolution2DGradientOperationClass
}

// GetNeuralNetworksConvolution2DGradientOperationClass returns the class object for NeuralNetworks.Convolution2DGradientOperation.
func GetNeuralNetworksConvolution2DGradientOperationClass() NeuralNetworksConvolution2DGradientOperationClass {
	return getNeuralNetworksConvolution2DGradientOperationClass()
}

type NeuralNetworksConvolution2DGradientOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksConvolution2DGradientOperationClass) Alloc() NeuralNetworksConvolution2DGradientOperation {
	rv := objc.Send[NeuralNetworksConvolution2DGradientOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.Convolution2DGradientOperation
type NeuralNetworksConvolution2DGradientOperation struct {
	NeuralNetworksBaseConvolution2DOperation
}

// NeuralNetworksConvolution2DGradientOperationFromID constructs a [NeuralNetworksConvolution2DGradientOperation] from an objc.ID.
func NeuralNetworksConvolution2DGradientOperationFromID(id objc.ID) NeuralNetworksConvolution2DGradientOperation {
	return NeuralNetworksConvolution2DGradientOperation{NeuralNetworksBaseConvolution2DOperation: NeuralNetworksBaseConvolution2DOperationFromID(id)}
}
// Ensure NeuralNetworksConvolution2DGradientOperation implements INeuralNetworksConvolution2DGradientOperation.
var _ INeuralNetworksConvolution2DGradientOperation = NeuralNetworksConvolution2DGradientOperation{}





// An interface definition for the [NeuralNetworksConvolution2DGradientOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.Convolution2DGradientOperation
type INeuralNetworksConvolution2DGradientOperation interface {
	INeuralNetworksBaseConvolution2DOperation
}





// Init initializes the instance.
func (n NeuralNetworksConvolution2DGradientOperation) Init() NeuralNetworksConvolution2DGradientOperation {
	rv := objc.Send[NeuralNetworksConvolution2DGradientOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksConvolution2DGradientOperation) Autorelease() NeuralNetworksConvolution2DGradientOperation {
	rv := objc.Send[NeuralNetworksConvolution2DGradientOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksConvolution2DGradientOperation creates a new NeuralNetworksConvolution2DGradientOperation instance.
func NewNeuralNetworksConvolution2DGradientOperation() NeuralNetworksConvolution2DGradientOperation {
	class := getNeuralNetworksConvolution2DGradientOperationClass()
	rv := objc.Send[NeuralNetworksConvolution2DGradientOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































