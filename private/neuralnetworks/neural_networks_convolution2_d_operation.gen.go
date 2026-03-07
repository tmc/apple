// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NeuralNetworksConvolution2DOperation] class.
var (
	_NeuralNetworksConvolution2DOperationClass     NeuralNetworksConvolution2DOperationClass
	_NeuralNetworksConvolution2DOperationClassOnce sync.Once
)

func getNeuralNetworksConvolution2DOperationClass() NeuralNetworksConvolution2DOperationClass {
	_NeuralNetworksConvolution2DOperationClassOnce.Do(func() {
		_NeuralNetworksConvolution2DOperationClass = NeuralNetworksConvolution2DOperationClass{class: objc.GetClass("NeuralNetworks.Convolution2DOperation")}
	})
	return _NeuralNetworksConvolution2DOperationClass
}

// GetNeuralNetworksConvolution2DOperationClass returns the class object for NeuralNetworks.Convolution2DOperation.
func GetNeuralNetworksConvolution2DOperationClass() NeuralNetworksConvolution2DOperationClass {
	return getNeuralNetworksConvolution2DOperationClass()
}

type NeuralNetworksConvolution2DOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksConvolution2DOperationClass) Alloc() NeuralNetworksConvolution2DOperation {
	rv := objc.Send[NeuralNetworksConvolution2DOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.Convolution2DOperation
type NeuralNetworksConvolution2DOperation struct {
	NeuralNetworksBaseConvolution2DOperation
}

// NeuralNetworksConvolution2DOperationFromID constructs a [NeuralNetworksConvolution2DOperation] from an objc.ID.
func NeuralNetworksConvolution2DOperationFromID(id objc.ID) NeuralNetworksConvolution2DOperation {
	return NeuralNetworksConvolution2DOperation{NeuralNetworksBaseConvolution2DOperation: NeuralNetworksBaseConvolution2DOperationFromID(id)}
}
// Ensure NeuralNetworksConvolution2DOperation implements INeuralNetworksConvolution2DOperation.
var _ INeuralNetworksConvolution2DOperation = NeuralNetworksConvolution2DOperation{}





// An interface definition for the [NeuralNetworksConvolution2DOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.Convolution2DOperation
type INeuralNetworksConvolution2DOperation interface {
	INeuralNetworksBaseConvolution2DOperation
}





// Init initializes the instance.
func (n NeuralNetworksConvolution2DOperation) Init() NeuralNetworksConvolution2DOperation {
	rv := objc.Send[NeuralNetworksConvolution2DOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksConvolution2DOperation) Autorelease() NeuralNetworksConvolution2DOperation {
	rv := objc.Send[NeuralNetworksConvolution2DOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksConvolution2DOperation creates a new NeuralNetworksConvolution2DOperation instance.
func NewNeuralNetworksConvolution2DOperation() NeuralNetworksConvolution2DOperation {
	class := getNeuralNetworksConvolution2DOperationClass()
	rv := objc.Send[NeuralNetworksConvolution2DOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































