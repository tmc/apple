// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksBaseLossOperation] class.
var (
	_NeuralNetworksBaseLossOperationClass     NeuralNetworksBaseLossOperationClass
	_NeuralNetworksBaseLossOperationClassOnce sync.Once
)

func getNeuralNetworksBaseLossOperationClass() NeuralNetworksBaseLossOperationClass {
	_NeuralNetworksBaseLossOperationClassOnce.Do(func() {
		_NeuralNetworksBaseLossOperationClass = NeuralNetworksBaseLossOperationClass{class: objc.GetClass("NeuralNetworks.BaseLossOperation")}
	})
	return _NeuralNetworksBaseLossOperationClass
}

// GetNeuralNetworksBaseLossOperationClass returns the class object for NeuralNetworks.BaseLossOperation.
func GetNeuralNetworksBaseLossOperationClass() NeuralNetworksBaseLossOperationClass {
	return getNeuralNetworksBaseLossOperationClass()
}

type NeuralNetworksBaseLossOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksBaseLossOperationClass) Alloc() NeuralNetworksBaseLossOperation {
	rv := objc.Send[NeuralNetworksBaseLossOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.BaseLossOperation
type NeuralNetworksBaseLossOperation struct {
	objectivec.Object
}

// NeuralNetworksBaseLossOperationFromID constructs a [NeuralNetworksBaseLossOperation] from an objc.ID.
func NeuralNetworksBaseLossOperationFromID(id objc.ID) NeuralNetworksBaseLossOperation {
	return NeuralNetworksBaseLossOperation{objectivec.Object{id}}
}
// Ensure NeuralNetworksBaseLossOperation implements INeuralNetworksBaseLossOperation.
var _ INeuralNetworksBaseLossOperation = NeuralNetworksBaseLossOperation{}





// An interface definition for the [NeuralNetworksBaseLossOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.BaseLossOperation
type INeuralNetworksBaseLossOperation interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksBaseLossOperation) Init() NeuralNetworksBaseLossOperation {
	rv := objc.Send[NeuralNetworksBaseLossOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksBaseLossOperation) Autorelease() NeuralNetworksBaseLossOperation {
	rv := objc.Send[NeuralNetworksBaseLossOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksBaseLossOperation creates a new NeuralNetworksBaseLossOperation instance.
func NewNeuralNetworksBaseLossOperation() NeuralNetworksBaseLossOperation {
	class := getNeuralNetworksBaseLossOperationClass()
	rv := objc.Send[NeuralNetworksBaseLossOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































