// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksBaseUnaryElementwiseOperation] class.
var (
	_NeuralNetworksBaseUnaryElementwiseOperationClass     NeuralNetworksBaseUnaryElementwiseOperationClass
	_NeuralNetworksBaseUnaryElementwiseOperationClassOnce sync.Once
)

func getNeuralNetworksBaseUnaryElementwiseOperationClass() NeuralNetworksBaseUnaryElementwiseOperationClass {
	_NeuralNetworksBaseUnaryElementwiseOperationClassOnce.Do(func() {
		_NeuralNetworksBaseUnaryElementwiseOperationClass = NeuralNetworksBaseUnaryElementwiseOperationClass{class: objc.GetClass("NeuralNetworks.BaseUnaryElementwiseOperation")}
	})
	return _NeuralNetworksBaseUnaryElementwiseOperationClass
}

// GetNeuralNetworksBaseUnaryElementwiseOperationClass returns the class object for NeuralNetworks.BaseUnaryElementwiseOperation.
func GetNeuralNetworksBaseUnaryElementwiseOperationClass() NeuralNetworksBaseUnaryElementwiseOperationClass {
	return getNeuralNetworksBaseUnaryElementwiseOperationClass()
}

type NeuralNetworksBaseUnaryElementwiseOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksBaseUnaryElementwiseOperationClass) Alloc() NeuralNetworksBaseUnaryElementwiseOperation {
	rv := objc.Send[NeuralNetworksBaseUnaryElementwiseOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.BaseUnaryElementwiseOperation
type NeuralNetworksBaseUnaryElementwiseOperation struct {
	objectivec.Object
}

// NeuralNetworksBaseUnaryElementwiseOperationFromID constructs a [NeuralNetworksBaseUnaryElementwiseOperation] from an objc.ID.
func NeuralNetworksBaseUnaryElementwiseOperationFromID(id objc.ID) NeuralNetworksBaseUnaryElementwiseOperation {
	return NeuralNetworksBaseUnaryElementwiseOperation{objectivec.Object{id}}
}
// Ensure NeuralNetworksBaseUnaryElementwiseOperation implements INeuralNetworksBaseUnaryElementwiseOperation.
var _ INeuralNetworksBaseUnaryElementwiseOperation = NeuralNetworksBaseUnaryElementwiseOperation{}





// An interface definition for the [NeuralNetworksBaseUnaryElementwiseOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.BaseUnaryElementwiseOperation
type INeuralNetworksBaseUnaryElementwiseOperation interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksBaseUnaryElementwiseOperation) Init() NeuralNetworksBaseUnaryElementwiseOperation {
	rv := objc.Send[NeuralNetworksBaseUnaryElementwiseOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksBaseUnaryElementwiseOperation) Autorelease() NeuralNetworksBaseUnaryElementwiseOperation {
	rv := objc.Send[NeuralNetworksBaseUnaryElementwiseOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksBaseUnaryElementwiseOperation creates a new NeuralNetworksBaseUnaryElementwiseOperation instance.
func NewNeuralNetworksBaseUnaryElementwiseOperation() NeuralNetworksBaseUnaryElementwiseOperation {
	class := getNeuralNetworksBaseUnaryElementwiseOperationClass()
	rv := objc.Send[NeuralNetworksBaseUnaryElementwiseOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































