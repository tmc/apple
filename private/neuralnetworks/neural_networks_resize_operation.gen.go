// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NeuralNetworksResizeOperation] class.
var (
	_NeuralNetworksResizeOperationClass     NeuralNetworksResizeOperationClass
	_NeuralNetworksResizeOperationClassOnce sync.Once
)

func getNeuralNetworksResizeOperationClass() NeuralNetworksResizeOperationClass {
	_NeuralNetworksResizeOperationClassOnce.Do(func() {
		_NeuralNetworksResizeOperationClass = NeuralNetworksResizeOperationClass{class: objc.GetClass("NeuralNetworks.ResizeOperation")}
	})
	return _NeuralNetworksResizeOperationClass
}

// GetNeuralNetworksResizeOperationClass returns the class object for NeuralNetworks.ResizeOperation.
func GetNeuralNetworksResizeOperationClass() NeuralNetworksResizeOperationClass {
	return getNeuralNetworksResizeOperationClass()
}

type NeuralNetworksResizeOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksResizeOperationClass) Alloc() NeuralNetworksResizeOperation {
	rv := objc.Send[NeuralNetworksResizeOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.ResizeOperation
type NeuralNetworksResizeOperation struct {
	NeuralNetworksBaseResizeOperation
}

// NeuralNetworksResizeOperationFromID constructs a [NeuralNetworksResizeOperation] from an objc.ID.
func NeuralNetworksResizeOperationFromID(id objc.ID) NeuralNetworksResizeOperation {
	return NeuralNetworksResizeOperation{NeuralNetworksBaseResizeOperation: NeuralNetworksBaseResizeOperationFromID(id)}
}
// Ensure NeuralNetworksResizeOperation implements INeuralNetworksResizeOperation.
var _ INeuralNetworksResizeOperation = NeuralNetworksResizeOperation{}





// An interface definition for the [NeuralNetworksResizeOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.ResizeOperation
type INeuralNetworksResizeOperation interface {
	INeuralNetworksBaseResizeOperation
}





// Init initializes the instance.
func (n NeuralNetworksResizeOperation) Init() NeuralNetworksResizeOperation {
	rv := objc.Send[NeuralNetworksResizeOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksResizeOperation) Autorelease() NeuralNetworksResizeOperation {
	rv := objc.Send[NeuralNetworksResizeOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksResizeOperation creates a new NeuralNetworksResizeOperation instance.
func NewNeuralNetworksResizeOperation() NeuralNetworksResizeOperation {
	class := getNeuralNetworksResizeOperationClass()
	rv := objc.Send[NeuralNetworksResizeOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































