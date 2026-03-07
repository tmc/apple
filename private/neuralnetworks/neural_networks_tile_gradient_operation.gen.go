// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NeuralNetworksTileGradientOperation] class.
var (
	_NeuralNetworksTileGradientOperationClass     NeuralNetworksTileGradientOperationClass
	_NeuralNetworksTileGradientOperationClassOnce sync.Once
)

func getNeuralNetworksTileGradientOperationClass() NeuralNetworksTileGradientOperationClass {
	_NeuralNetworksTileGradientOperationClassOnce.Do(func() {
		_NeuralNetworksTileGradientOperationClass = NeuralNetworksTileGradientOperationClass{class: objc.GetClass("NeuralNetworks.TileGradientOperation")}
	})
	return _NeuralNetworksTileGradientOperationClass
}

// GetNeuralNetworksTileGradientOperationClass returns the class object for NeuralNetworks.TileGradientOperation.
func GetNeuralNetworksTileGradientOperationClass() NeuralNetworksTileGradientOperationClass {
	return getNeuralNetworksTileGradientOperationClass()
}

type NeuralNetworksTileGradientOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksTileGradientOperationClass) Alloc() NeuralNetworksTileGradientOperation {
	rv := objc.Send[NeuralNetworksTileGradientOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.TileGradientOperation
type NeuralNetworksTileGradientOperation struct {
	NeuralNetworksBaseTileOperation
}

// NeuralNetworksTileGradientOperationFromID constructs a [NeuralNetworksTileGradientOperation] from an objc.ID.
func NeuralNetworksTileGradientOperationFromID(id objc.ID) NeuralNetworksTileGradientOperation {
	return NeuralNetworksTileGradientOperation{NeuralNetworksBaseTileOperation: NeuralNetworksBaseTileOperationFromID(id)}
}
// Ensure NeuralNetworksTileGradientOperation implements INeuralNetworksTileGradientOperation.
var _ INeuralNetworksTileGradientOperation = NeuralNetworksTileGradientOperation{}





// An interface definition for the [NeuralNetworksTileGradientOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.TileGradientOperation
type INeuralNetworksTileGradientOperation interface {
	INeuralNetworksBaseTileOperation
}





// Init initializes the instance.
func (n NeuralNetworksTileGradientOperation) Init() NeuralNetworksTileGradientOperation {
	rv := objc.Send[NeuralNetworksTileGradientOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksTileGradientOperation) Autorelease() NeuralNetworksTileGradientOperation {
	rv := objc.Send[NeuralNetworksTileGradientOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksTileGradientOperation creates a new NeuralNetworksTileGradientOperation instance.
func NewNeuralNetworksTileGradientOperation() NeuralNetworksTileGradientOperation {
	class := getNeuralNetworksTileGradientOperationClass()
	rv := objc.Send[NeuralNetworksTileGradientOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































