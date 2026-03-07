// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksBaseConvolution2DOperation] class.
var (
	_NeuralNetworksBaseConvolution2DOperationClass     NeuralNetworksBaseConvolution2DOperationClass
	_NeuralNetworksBaseConvolution2DOperationClassOnce sync.Once
)

func getNeuralNetworksBaseConvolution2DOperationClass() NeuralNetworksBaseConvolution2DOperationClass {
	_NeuralNetworksBaseConvolution2DOperationClassOnce.Do(func() {
		_NeuralNetworksBaseConvolution2DOperationClass = NeuralNetworksBaseConvolution2DOperationClass{class: objc.GetClass("NeuralNetworks.BaseConvolution2DOperation")}
	})
	return _NeuralNetworksBaseConvolution2DOperationClass
}

// GetNeuralNetworksBaseConvolution2DOperationClass returns the class object for NeuralNetworks.BaseConvolution2DOperation.
func GetNeuralNetworksBaseConvolution2DOperationClass() NeuralNetworksBaseConvolution2DOperationClass {
	return getNeuralNetworksBaseConvolution2DOperationClass()
}

type NeuralNetworksBaseConvolution2DOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksBaseConvolution2DOperationClass) Alloc() NeuralNetworksBaseConvolution2DOperation {
	rv := objc.Send[NeuralNetworksBaseConvolution2DOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.BaseConvolution2DOperation
type NeuralNetworksBaseConvolution2DOperation struct {
	objectivec.Object
}

// NeuralNetworksBaseConvolution2DOperationFromID constructs a [NeuralNetworksBaseConvolution2DOperation] from an objc.ID.
func NeuralNetworksBaseConvolution2DOperationFromID(id objc.ID) NeuralNetworksBaseConvolution2DOperation {
	return NeuralNetworksBaseConvolution2DOperation{objectivec.Object{id}}
}
// Ensure NeuralNetworksBaseConvolution2DOperation implements INeuralNetworksBaseConvolution2DOperation.
var _ INeuralNetworksBaseConvolution2DOperation = NeuralNetworksBaseConvolution2DOperation{}





// An interface definition for the [NeuralNetworksBaseConvolution2DOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.BaseConvolution2DOperation
type INeuralNetworksBaseConvolution2DOperation interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksBaseConvolution2DOperation) Init() NeuralNetworksBaseConvolution2DOperation {
	rv := objc.Send[NeuralNetworksBaseConvolution2DOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksBaseConvolution2DOperation) Autorelease() NeuralNetworksBaseConvolution2DOperation {
	rv := objc.Send[NeuralNetworksBaseConvolution2DOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksBaseConvolution2DOperation creates a new NeuralNetworksBaseConvolution2DOperation instance.
func NewNeuralNetworksBaseConvolution2DOperation() NeuralNetworksBaseConvolution2DOperation {
	class := getNeuralNetworksBaseConvolution2DOperationClass()
	rv := objc.Send[NeuralNetworksBaseConvolution2DOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































