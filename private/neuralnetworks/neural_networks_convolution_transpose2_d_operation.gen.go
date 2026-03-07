// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NeuralNetworksConvolutionTranspose2DOperation] class.
var (
	_NeuralNetworksConvolutionTranspose2DOperationClass     NeuralNetworksConvolutionTranspose2DOperationClass
	_NeuralNetworksConvolutionTranspose2DOperationClassOnce sync.Once
)

func getNeuralNetworksConvolutionTranspose2DOperationClass() NeuralNetworksConvolutionTranspose2DOperationClass {
	_NeuralNetworksConvolutionTranspose2DOperationClassOnce.Do(func() {
		_NeuralNetworksConvolutionTranspose2DOperationClass = NeuralNetworksConvolutionTranspose2DOperationClass{class: objc.GetClass("NeuralNetworks.ConvolutionTranspose2DOperation")}
	})
	return _NeuralNetworksConvolutionTranspose2DOperationClass
}

// GetNeuralNetworksConvolutionTranspose2DOperationClass returns the class object for NeuralNetworks.ConvolutionTranspose2DOperation.
func GetNeuralNetworksConvolutionTranspose2DOperationClass() NeuralNetworksConvolutionTranspose2DOperationClass {
	return getNeuralNetworksConvolutionTranspose2DOperationClass()
}

type NeuralNetworksConvolutionTranspose2DOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksConvolutionTranspose2DOperationClass) Alloc() NeuralNetworksConvolutionTranspose2DOperation {
	rv := objc.Send[NeuralNetworksConvolutionTranspose2DOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.ConvolutionTranspose2DOperation
type NeuralNetworksConvolutionTranspose2DOperation struct {
	NeuralNetworksBaseConvolution2DOperation
}

// NeuralNetworksConvolutionTranspose2DOperationFromID constructs a [NeuralNetworksConvolutionTranspose2DOperation] from an objc.ID.
func NeuralNetworksConvolutionTranspose2DOperationFromID(id objc.ID) NeuralNetworksConvolutionTranspose2DOperation {
	return NeuralNetworksConvolutionTranspose2DOperation{NeuralNetworksBaseConvolution2DOperation: NeuralNetworksBaseConvolution2DOperationFromID(id)}
}
// Ensure NeuralNetworksConvolutionTranspose2DOperation implements INeuralNetworksConvolutionTranspose2DOperation.
var _ INeuralNetworksConvolutionTranspose2DOperation = NeuralNetworksConvolutionTranspose2DOperation{}





// An interface definition for the [NeuralNetworksConvolutionTranspose2DOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.ConvolutionTranspose2DOperation
type INeuralNetworksConvolutionTranspose2DOperation interface {
	INeuralNetworksBaseConvolution2DOperation
}





// Init initializes the instance.
func (n NeuralNetworksConvolutionTranspose2DOperation) Init() NeuralNetworksConvolutionTranspose2DOperation {
	rv := objc.Send[NeuralNetworksConvolutionTranspose2DOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksConvolutionTranspose2DOperation) Autorelease() NeuralNetworksConvolutionTranspose2DOperation {
	rv := objc.Send[NeuralNetworksConvolutionTranspose2DOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksConvolutionTranspose2DOperation creates a new NeuralNetworksConvolutionTranspose2DOperation instance.
func NewNeuralNetworksConvolutionTranspose2DOperation() NeuralNetworksConvolutionTranspose2DOperation {
	class := getNeuralNetworksConvolutionTranspose2DOperationClass()
	rv := objc.Send[NeuralNetworksConvolutionTranspose2DOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































