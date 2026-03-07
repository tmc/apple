// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NeuralNetworksConvolutionTranspose2DGradientOperation] class.
var (
	_NeuralNetworksConvolutionTranspose2DGradientOperationClass     NeuralNetworksConvolutionTranspose2DGradientOperationClass
	_NeuralNetworksConvolutionTranspose2DGradientOperationClassOnce sync.Once
)

func getNeuralNetworksConvolutionTranspose2DGradientOperationClass() NeuralNetworksConvolutionTranspose2DGradientOperationClass {
	_NeuralNetworksConvolutionTranspose2DGradientOperationClassOnce.Do(func() {
		_NeuralNetworksConvolutionTranspose2DGradientOperationClass = NeuralNetworksConvolutionTranspose2DGradientOperationClass{class: objc.GetClass("NeuralNetworks.ConvolutionTranspose2DGradientOperation")}
	})
	return _NeuralNetworksConvolutionTranspose2DGradientOperationClass
}

// GetNeuralNetworksConvolutionTranspose2DGradientOperationClass returns the class object for NeuralNetworks.ConvolutionTranspose2DGradientOperation.
func GetNeuralNetworksConvolutionTranspose2DGradientOperationClass() NeuralNetworksConvolutionTranspose2DGradientOperationClass {
	return getNeuralNetworksConvolutionTranspose2DGradientOperationClass()
}

type NeuralNetworksConvolutionTranspose2DGradientOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksConvolutionTranspose2DGradientOperationClass) Alloc() NeuralNetworksConvolutionTranspose2DGradientOperation {
	rv := objc.Send[NeuralNetworksConvolutionTranspose2DGradientOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.ConvolutionTranspose2DGradientOperation
type NeuralNetworksConvolutionTranspose2DGradientOperation struct {
	NeuralNetworksBaseConvolution2DOperation
}

// NeuralNetworksConvolutionTranspose2DGradientOperationFromID constructs a [NeuralNetworksConvolutionTranspose2DGradientOperation] from an objc.ID.
func NeuralNetworksConvolutionTranspose2DGradientOperationFromID(id objc.ID) NeuralNetworksConvolutionTranspose2DGradientOperation {
	return NeuralNetworksConvolutionTranspose2DGradientOperation{NeuralNetworksBaseConvolution2DOperation: NeuralNetworksBaseConvolution2DOperationFromID(id)}
}
// Ensure NeuralNetworksConvolutionTranspose2DGradientOperation implements INeuralNetworksConvolutionTranspose2DGradientOperation.
var _ INeuralNetworksConvolutionTranspose2DGradientOperation = NeuralNetworksConvolutionTranspose2DGradientOperation{}





// An interface definition for the [NeuralNetworksConvolutionTranspose2DGradientOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.ConvolutionTranspose2DGradientOperation
type INeuralNetworksConvolutionTranspose2DGradientOperation interface {
	INeuralNetworksBaseConvolution2DOperation
}





// Init initializes the instance.
func (n NeuralNetworksConvolutionTranspose2DGradientOperation) Init() NeuralNetworksConvolutionTranspose2DGradientOperation {
	rv := objc.Send[NeuralNetworksConvolutionTranspose2DGradientOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksConvolutionTranspose2DGradientOperation) Autorelease() NeuralNetworksConvolutionTranspose2DGradientOperation {
	rv := objc.Send[NeuralNetworksConvolutionTranspose2DGradientOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksConvolutionTranspose2DGradientOperation creates a new NeuralNetworksConvolutionTranspose2DGradientOperation instance.
func NewNeuralNetworksConvolutionTranspose2DGradientOperation() NeuralNetworksConvolutionTranspose2DGradientOperation {
	class := getNeuralNetworksConvolutionTranspose2DGradientOperationClass()
	rv := objc.Send[NeuralNetworksConvolutionTranspose2DGradientOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































