// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksLinearGradientOperation] class.
var (
	_NeuralNetworksLinearGradientOperationClass     NeuralNetworksLinearGradientOperationClass
	_NeuralNetworksLinearGradientOperationClassOnce sync.Once
)

func getNeuralNetworksLinearGradientOperationClass() NeuralNetworksLinearGradientOperationClass {
	_NeuralNetworksLinearGradientOperationClassOnce.Do(func() {
		_NeuralNetworksLinearGradientOperationClass = NeuralNetworksLinearGradientOperationClass{class: objc.GetClass("NeuralNetworks.LinearGradientOperation")}
	})
	return _NeuralNetworksLinearGradientOperationClass
}

// GetNeuralNetworksLinearGradientOperationClass returns the class object for NeuralNetworks.LinearGradientOperation.
func GetNeuralNetworksLinearGradientOperationClass() NeuralNetworksLinearGradientOperationClass {
	return getNeuralNetworksLinearGradientOperationClass()
}

type NeuralNetworksLinearGradientOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksLinearGradientOperationClass) Alloc() NeuralNetworksLinearGradientOperation {
	rv := objc.Send[NeuralNetworksLinearGradientOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.LinearGradientOperation
type NeuralNetworksLinearGradientOperation struct {
	objectivec.Object
}

// NeuralNetworksLinearGradientOperationFromID constructs a [NeuralNetworksLinearGradientOperation] from an objc.ID.
func NeuralNetworksLinearGradientOperationFromID(id objc.ID) NeuralNetworksLinearGradientOperation {
	return NeuralNetworksLinearGradientOperation{objectivec.Object{id}}
}
// Ensure NeuralNetworksLinearGradientOperation implements INeuralNetworksLinearGradientOperation.
var _ INeuralNetworksLinearGradientOperation = NeuralNetworksLinearGradientOperation{}





// An interface definition for the [NeuralNetworksLinearGradientOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.LinearGradientOperation
type INeuralNetworksLinearGradientOperation interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksLinearGradientOperation) Init() NeuralNetworksLinearGradientOperation {
	rv := objc.Send[NeuralNetworksLinearGradientOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksLinearGradientOperation) Autorelease() NeuralNetworksLinearGradientOperation {
	rv := objc.Send[NeuralNetworksLinearGradientOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksLinearGradientOperation creates a new NeuralNetworksLinearGradientOperation instance.
func NewNeuralNetworksLinearGradientOperation() NeuralNetworksLinearGradientOperation {
	class := getNeuralNetworksLinearGradientOperationClass()
	rv := objc.Send[NeuralNetworksLinearGradientOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































