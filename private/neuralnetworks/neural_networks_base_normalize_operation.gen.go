// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksBaseNormalizeOperation] class.
var (
	_NeuralNetworksBaseNormalizeOperationClass     NeuralNetworksBaseNormalizeOperationClass
	_NeuralNetworksBaseNormalizeOperationClassOnce sync.Once
)

func getNeuralNetworksBaseNormalizeOperationClass() NeuralNetworksBaseNormalizeOperationClass {
	_NeuralNetworksBaseNormalizeOperationClassOnce.Do(func() {
		_NeuralNetworksBaseNormalizeOperationClass = NeuralNetworksBaseNormalizeOperationClass{class: objc.GetClass("NeuralNetworks.BaseNormalizeOperation")}
	})
	return _NeuralNetworksBaseNormalizeOperationClass
}

// GetNeuralNetworksBaseNormalizeOperationClass returns the class object for NeuralNetworks.BaseNormalizeOperation.
func GetNeuralNetworksBaseNormalizeOperationClass() NeuralNetworksBaseNormalizeOperationClass {
	return getNeuralNetworksBaseNormalizeOperationClass()
}

type NeuralNetworksBaseNormalizeOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksBaseNormalizeOperationClass) Alloc() NeuralNetworksBaseNormalizeOperation {
	rv := objc.Send[NeuralNetworksBaseNormalizeOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.BaseNormalizeOperation
type NeuralNetworksBaseNormalizeOperation struct {
	objectivec.Object
}

// NeuralNetworksBaseNormalizeOperationFromID constructs a [NeuralNetworksBaseNormalizeOperation] from an objc.ID.
func NeuralNetworksBaseNormalizeOperationFromID(id objc.ID) NeuralNetworksBaseNormalizeOperation {
	return NeuralNetworksBaseNormalizeOperation{objectivec.Object{id}}
}
// Ensure NeuralNetworksBaseNormalizeOperation implements INeuralNetworksBaseNormalizeOperation.
var _ INeuralNetworksBaseNormalizeOperation = NeuralNetworksBaseNormalizeOperation{}





// An interface definition for the [NeuralNetworksBaseNormalizeOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.BaseNormalizeOperation
type INeuralNetworksBaseNormalizeOperation interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksBaseNormalizeOperation) Init() NeuralNetworksBaseNormalizeOperation {
	rv := objc.Send[NeuralNetworksBaseNormalizeOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksBaseNormalizeOperation) Autorelease() NeuralNetworksBaseNormalizeOperation {
	rv := objc.Send[NeuralNetworksBaseNormalizeOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksBaseNormalizeOperation creates a new NeuralNetworksBaseNormalizeOperation instance.
func NewNeuralNetworksBaseNormalizeOperation() NeuralNetworksBaseNormalizeOperation {
	class := getNeuralNetworksBaseNormalizeOperationClass()
	rv := objc.Send[NeuralNetworksBaseNormalizeOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































