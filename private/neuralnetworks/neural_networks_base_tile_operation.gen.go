// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksBaseTileOperation] class.
var (
	_NeuralNetworksBaseTileOperationClass     NeuralNetworksBaseTileOperationClass
	_NeuralNetworksBaseTileOperationClassOnce sync.Once
)

func getNeuralNetworksBaseTileOperationClass() NeuralNetworksBaseTileOperationClass {
	_NeuralNetworksBaseTileOperationClassOnce.Do(func() {
		_NeuralNetworksBaseTileOperationClass = NeuralNetworksBaseTileOperationClass{class: objc.GetClass("NeuralNetworks.BaseTileOperation")}
	})
	return _NeuralNetworksBaseTileOperationClass
}

// GetNeuralNetworksBaseTileOperationClass returns the class object for NeuralNetworks.BaseTileOperation.
func GetNeuralNetworksBaseTileOperationClass() NeuralNetworksBaseTileOperationClass {
	return getNeuralNetworksBaseTileOperationClass()
}

type NeuralNetworksBaseTileOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksBaseTileOperationClass) Alloc() NeuralNetworksBaseTileOperation {
	rv := objc.Send[NeuralNetworksBaseTileOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.BaseTileOperation
type NeuralNetworksBaseTileOperation struct {
	objectivec.Object
}

// NeuralNetworksBaseTileOperationFromID constructs a [NeuralNetworksBaseTileOperation] from an objc.ID.
func NeuralNetworksBaseTileOperationFromID(id objc.ID) NeuralNetworksBaseTileOperation {
	return NeuralNetworksBaseTileOperation{objectivec.Object{id}}
}
// Ensure NeuralNetworksBaseTileOperation implements INeuralNetworksBaseTileOperation.
var _ INeuralNetworksBaseTileOperation = NeuralNetworksBaseTileOperation{}





// An interface definition for the [NeuralNetworksBaseTileOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.BaseTileOperation
type INeuralNetworksBaseTileOperation interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksBaseTileOperation) Init() NeuralNetworksBaseTileOperation {
	rv := objc.Send[NeuralNetworksBaseTileOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksBaseTileOperation) Autorelease() NeuralNetworksBaseTileOperation {
	rv := objc.Send[NeuralNetworksBaseTileOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksBaseTileOperation creates a new NeuralNetworksBaseTileOperation instance.
func NewNeuralNetworksBaseTileOperation() NeuralNetworksBaseTileOperation {
	class := getNeuralNetworksBaseTileOperationClass()
	rv := objc.Send[NeuralNetworksBaseTileOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































