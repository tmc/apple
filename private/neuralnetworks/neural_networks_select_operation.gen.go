// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksSelectOperation] class.
var (
	_NeuralNetworksSelectOperationClass     NeuralNetworksSelectOperationClass
	_NeuralNetworksSelectOperationClassOnce sync.Once
)

func getNeuralNetworksSelectOperationClass() NeuralNetworksSelectOperationClass {
	_NeuralNetworksSelectOperationClassOnce.Do(func() {
		_NeuralNetworksSelectOperationClass = NeuralNetworksSelectOperationClass{class: objc.GetClass("NeuralNetworks.SelectOperation")}
	})
	return _NeuralNetworksSelectOperationClass
}

// GetNeuralNetworksSelectOperationClass returns the class object for NeuralNetworks.SelectOperation.
func GetNeuralNetworksSelectOperationClass() NeuralNetworksSelectOperationClass {
	return getNeuralNetworksSelectOperationClass()
}

type NeuralNetworksSelectOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksSelectOperationClass) Alloc() NeuralNetworksSelectOperation {
	rv := objc.Send[NeuralNetworksSelectOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.SelectOperation
type NeuralNetworksSelectOperation struct {
	objectivec.Object
}

// NeuralNetworksSelectOperationFromID constructs a [NeuralNetworksSelectOperation] from an objc.ID.
func NeuralNetworksSelectOperationFromID(id objc.ID) NeuralNetworksSelectOperation {
	return NeuralNetworksSelectOperation{objectivec.Object{id}}
}
// Ensure NeuralNetworksSelectOperation implements INeuralNetworksSelectOperation.
var _ INeuralNetworksSelectOperation = NeuralNetworksSelectOperation{}





// An interface definition for the [NeuralNetworksSelectOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.SelectOperation
type INeuralNetworksSelectOperation interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksSelectOperation) Init() NeuralNetworksSelectOperation {
	rv := objc.Send[NeuralNetworksSelectOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksSelectOperation) Autorelease() NeuralNetworksSelectOperation {
	rv := objc.Send[NeuralNetworksSelectOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksSelectOperation creates a new NeuralNetworksSelectOperation instance.
func NewNeuralNetworksSelectOperation() NeuralNetworksSelectOperation {
	class := getNeuralNetworksSelectOperationClass()
	rv := objc.Send[NeuralNetworksSelectOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































