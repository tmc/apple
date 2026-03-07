// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksInputOperation] class.
var (
	_NeuralNetworksInputOperationClass     NeuralNetworksInputOperationClass
	_NeuralNetworksInputOperationClassOnce sync.Once
)

func getNeuralNetworksInputOperationClass() NeuralNetworksInputOperationClass {
	_NeuralNetworksInputOperationClassOnce.Do(func() {
		_NeuralNetworksInputOperationClass = NeuralNetworksInputOperationClass{class: objc.GetClass("NeuralNetworks.InputOperation")}
	})
	return _NeuralNetworksInputOperationClass
}

// GetNeuralNetworksInputOperationClass returns the class object for NeuralNetworks.InputOperation.
func GetNeuralNetworksInputOperationClass() NeuralNetworksInputOperationClass {
	return getNeuralNetworksInputOperationClass()
}

type NeuralNetworksInputOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksInputOperationClass) Alloc() NeuralNetworksInputOperation {
	rv := objc.Send[NeuralNetworksInputOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.InputOperation
type NeuralNetworksInputOperation struct {
	objectivec.Object
}

// NeuralNetworksInputOperationFromID constructs a [NeuralNetworksInputOperation] from an objc.ID.
func NeuralNetworksInputOperationFromID(id objc.ID) NeuralNetworksInputOperation {
	return NeuralNetworksInputOperation{objectivec.Object{id}}
}
// Ensure NeuralNetworksInputOperation implements INeuralNetworksInputOperation.
var _ INeuralNetworksInputOperation = NeuralNetworksInputOperation{}





// An interface definition for the [NeuralNetworksInputOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.InputOperation
type INeuralNetworksInputOperation interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksInputOperation) Init() NeuralNetworksInputOperation {
	rv := objc.Send[NeuralNetworksInputOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksInputOperation) Autorelease() NeuralNetworksInputOperation {
	rv := objc.Send[NeuralNetworksInputOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksInputOperation creates a new NeuralNetworksInputOperation instance.
func NewNeuralNetworksInputOperation() NeuralNetworksInputOperation {
	class := getNeuralNetworksInputOperationClass()
	rv := objc.Send[NeuralNetworksInputOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































