// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NeuralNetworksTileOperation] class.
var (
	_NeuralNetworksTileOperationClass     NeuralNetworksTileOperationClass
	_NeuralNetworksTileOperationClassOnce sync.Once
)

func getNeuralNetworksTileOperationClass() NeuralNetworksTileOperationClass {
	_NeuralNetworksTileOperationClassOnce.Do(func() {
		_NeuralNetworksTileOperationClass = NeuralNetworksTileOperationClass{class: objc.GetClass("NeuralNetworks.TileOperation")}
	})
	return _NeuralNetworksTileOperationClass
}

// GetNeuralNetworksTileOperationClass returns the class object for NeuralNetworks.TileOperation.
func GetNeuralNetworksTileOperationClass() NeuralNetworksTileOperationClass {
	return getNeuralNetworksTileOperationClass()
}

type NeuralNetworksTileOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksTileOperationClass) Alloc() NeuralNetworksTileOperation {
	rv := objc.Send[NeuralNetworksTileOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.TileOperation
type NeuralNetworksTileOperation struct {
	NeuralNetworksBaseTileOperation
}

// NeuralNetworksTileOperationFromID constructs a [NeuralNetworksTileOperation] from an objc.ID.
func NeuralNetworksTileOperationFromID(id objc.ID) NeuralNetworksTileOperation {
	return NeuralNetworksTileOperation{NeuralNetworksBaseTileOperation: NeuralNetworksBaseTileOperationFromID(id)}
}
// Ensure NeuralNetworksTileOperation implements INeuralNetworksTileOperation.
var _ INeuralNetworksTileOperation = NeuralNetworksTileOperation{}





// An interface definition for the [NeuralNetworksTileOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.TileOperation
type INeuralNetworksTileOperation interface {
	INeuralNetworksBaseTileOperation
}





// Init initializes the instance.
func (n NeuralNetworksTileOperation) Init() NeuralNetworksTileOperation {
	rv := objc.Send[NeuralNetworksTileOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksTileOperation) Autorelease() NeuralNetworksTileOperation {
	rv := objc.Send[NeuralNetworksTileOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksTileOperation creates a new NeuralNetworksTileOperation instance.
func NewNeuralNetworksTileOperation() NeuralNetworksTileOperation {
	class := getNeuralNetworksTileOperationClass()
	rv := objc.Send[NeuralNetworksTileOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































