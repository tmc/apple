// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksBaseSliceOperation] class.
var (
	_NeuralNetworksBaseSliceOperationClass     NeuralNetworksBaseSliceOperationClass
	_NeuralNetworksBaseSliceOperationClassOnce sync.Once
)

func getNeuralNetworksBaseSliceOperationClass() NeuralNetworksBaseSliceOperationClass {
	_NeuralNetworksBaseSliceOperationClassOnce.Do(func() {
		_NeuralNetworksBaseSliceOperationClass = NeuralNetworksBaseSliceOperationClass{class: objc.GetClass("NeuralNetworks.BaseSliceOperation")}
	})
	return _NeuralNetworksBaseSliceOperationClass
}

// GetNeuralNetworksBaseSliceOperationClass returns the class object for NeuralNetworks.BaseSliceOperation.
func GetNeuralNetworksBaseSliceOperationClass() NeuralNetworksBaseSliceOperationClass {
	return getNeuralNetworksBaseSliceOperationClass()
}

type NeuralNetworksBaseSliceOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksBaseSliceOperationClass) Alloc() NeuralNetworksBaseSliceOperation {
	rv := objc.Send[NeuralNetworksBaseSliceOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.BaseSliceOperation
type NeuralNetworksBaseSliceOperation struct {
	objectivec.Object
}

// NeuralNetworksBaseSliceOperationFromID constructs a [NeuralNetworksBaseSliceOperation] from an objc.ID.
func NeuralNetworksBaseSliceOperationFromID(id objc.ID) NeuralNetworksBaseSliceOperation {
	return NeuralNetworksBaseSliceOperation{objectivec.Object{id}}
}
// Ensure NeuralNetworksBaseSliceOperation implements INeuralNetworksBaseSliceOperation.
var _ INeuralNetworksBaseSliceOperation = NeuralNetworksBaseSliceOperation{}





// An interface definition for the [NeuralNetworksBaseSliceOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.BaseSliceOperation
type INeuralNetworksBaseSliceOperation interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksBaseSliceOperation) Init() NeuralNetworksBaseSliceOperation {
	rv := objc.Send[NeuralNetworksBaseSliceOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksBaseSliceOperation) Autorelease() NeuralNetworksBaseSliceOperation {
	rv := objc.Send[NeuralNetworksBaseSliceOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksBaseSliceOperation creates a new NeuralNetworksBaseSliceOperation instance.
func NewNeuralNetworksBaseSliceOperation() NeuralNetworksBaseSliceOperation {
	class := getNeuralNetworksBaseSliceOperationClass()
	rv := objc.Send[NeuralNetworksBaseSliceOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































