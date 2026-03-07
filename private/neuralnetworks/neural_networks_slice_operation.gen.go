// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NeuralNetworksSliceOperation] class.
var (
	_NeuralNetworksSliceOperationClass     NeuralNetworksSliceOperationClass
	_NeuralNetworksSliceOperationClassOnce sync.Once
)

func getNeuralNetworksSliceOperationClass() NeuralNetworksSliceOperationClass {
	_NeuralNetworksSliceOperationClassOnce.Do(func() {
		_NeuralNetworksSliceOperationClass = NeuralNetworksSliceOperationClass{class: objc.GetClass("NeuralNetworks.SliceOperation")}
	})
	return _NeuralNetworksSliceOperationClass
}

// GetNeuralNetworksSliceOperationClass returns the class object for NeuralNetworks.SliceOperation.
func GetNeuralNetworksSliceOperationClass() NeuralNetworksSliceOperationClass {
	return getNeuralNetworksSliceOperationClass()
}

type NeuralNetworksSliceOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksSliceOperationClass) Alloc() NeuralNetworksSliceOperation {
	rv := objc.Send[NeuralNetworksSliceOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.SliceOperation
type NeuralNetworksSliceOperation struct {
	NeuralNetworksBaseSliceOperation
}

// NeuralNetworksSliceOperationFromID constructs a [NeuralNetworksSliceOperation] from an objc.ID.
func NeuralNetworksSliceOperationFromID(id objc.ID) NeuralNetworksSliceOperation {
	return NeuralNetworksSliceOperation{NeuralNetworksBaseSliceOperation: NeuralNetworksBaseSliceOperationFromID(id)}
}
// Ensure NeuralNetworksSliceOperation implements INeuralNetworksSliceOperation.
var _ INeuralNetworksSliceOperation = NeuralNetworksSliceOperation{}





// An interface definition for the [NeuralNetworksSliceOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.SliceOperation
type INeuralNetworksSliceOperation interface {
	INeuralNetworksBaseSliceOperation
}





// Init initializes the instance.
func (n NeuralNetworksSliceOperation) Init() NeuralNetworksSliceOperation {
	rv := objc.Send[NeuralNetworksSliceOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksSliceOperation) Autorelease() NeuralNetworksSliceOperation {
	rv := objc.Send[NeuralNetworksSliceOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksSliceOperation creates a new NeuralNetworksSliceOperation instance.
func NewNeuralNetworksSliceOperation() NeuralNetworksSliceOperation {
	class := getNeuralNetworksSliceOperationClass()
	rv := objc.Send[NeuralNetworksSliceOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































