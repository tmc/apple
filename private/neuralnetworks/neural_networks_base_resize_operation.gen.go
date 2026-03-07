// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksBaseResizeOperation] class.
var (
	_NeuralNetworksBaseResizeOperationClass     NeuralNetworksBaseResizeOperationClass
	_NeuralNetworksBaseResizeOperationClassOnce sync.Once
)

func getNeuralNetworksBaseResizeOperationClass() NeuralNetworksBaseResizeOperationClass {
	_NeuralNetworksBaseResizeOperationClassOnce.Do(func() {
		_NeuralNetworksBaseResizeOperationClass = NeuralNetworksBaseResizeOperationClass{class: objc.GetClass("NeuralNetworks.BaseResizeOperation")}
	})
	return _NeuralNetworksBaseResizeOperationClass
}

// GetNeuralNetworksBaseResizeOperationClass returns the class object for NeuralNetworks.BaseResizeOperation.
func GetNeuralNetworksBaseResizeOperationClass() NeuralNetworksBaseResizeOperationClass {
	return getNeuralNetworksBaseResizeOperationClass()
}

type NeuralNetworksBaseResizeOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksBaseResizeOperationClass) Alloc() NeuralNetworksBaseResizeOperation {
	rv := objc.Send[NeuralNetworksBaseResizeOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.BaseResizeOperation
type NeuralNetworksBaseResizeOperation struct {
	objectivec.Object
}

// NeuralNetworksBaseResizeOperationFromID constructs a [NeuralNetworksBaseResizeOperation] from an objc.ID.
func NeuralNetworksBaseResizeOperationFromID(id objc.ID) NeuralNetworksBaseResizeOperation {
	return NeuralNetworksBaseResizeOperation{objectivec.Object{id}}
}
// Ensure NeuralNetworksBaseResizeOperation implements INeuralNetworksBaseResizeOperation.
var _ INeuralNetworksBaseResizeOperation = NeuralNetworksBaseResizeOperation{}





// An interface definition for the [NeuralNetworksBaseResizeOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.BaseResizeOperation
type INeuralNetworksBaseResizeOperation interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksBaseResizeOperation) Init() NeuralNetworksBaseResizeOperation {
	rv := objc.Send[NeuralNetworksBaseResizeOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksBaseResizeOperation) Autorelease() NeuralNetworksBaseResizeOperation {
	rv := objc.Send[NeuralNetworksBaseResizeOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksBaseResizeOperation creates a new NeuralNetworksBaseResizeOperation instance.
func NewNeuralNetworksBaseResizeOperation() NeuralNetworksBaseResizeOperation {
	class := getNeuralNetworksBaseResizeOperationClass()
	rv := objc.Send[NeuralNetworksBaseResizeOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































