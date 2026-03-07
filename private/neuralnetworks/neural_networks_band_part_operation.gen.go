// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksBandPartOperation] class.
var (
	_NeuralNetworksBandPartOperationClass     NeuralNetworksBandPartOperationClass
	_NeuralNetworksBandPartOperationClassOnce sync.Once
)

func getNeuralNetworksBandPartOperationClass() NeuralNetworksBandPartOperationClass {
	_NeuralNetworksBandPartOperationClassOnce.Do(func() {
		_NeuralNetworksBandPartOperationClass = NeuralNetworksBandPartOperationClass{class: objc.GetClass("NeuralNetworks.BandPartOperation")}
	})
	return _NeuralNetworksBandPartOperationClass
}

// GetNeuralNetworksBandPartOperationClass returns the class object for NeuralNetworks.BandPartOperation.
func GetNeuralNetworksBandPartOperationClass() NeuralNetworksBandPartOperationClass {
	return getNeuralNetworksBandPartOperationClass()
}

type NeuralNetworksBandPartOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksBandPartOperationClass) Alloc() NeuralNetworksBandPartOperation {
	rv := objc.Send[NeuralNetworksBandPartOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.BandPartOperation
type NeuralNetworksBandPartOperation struct {
	objectivec.Object
}

// NeuralNetworksBandPartOperationFromID constructs a [NeuralNetworksBandPartOperation] from an objc.ID.
func NeuralNetworksBandPartOperationFromID(id objc.ID) NeuralNetworksBandPartOperation {
	return NeuralNetworksBandPartOperation{objectivec.Object{id}}
}
// Ensure NeuralNetworksBandPartOperation implements INeuralNetworksBandPartOperation.
var _ INeuralNetworksBandPartOperation = NeuralNetworksBandPartOperation{}





// An interface definition for the [NeuralNetworksBandPartOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.BandPartOperation
type INeuralNetworksBandPartOperation interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksBandPartOperation) Init() NeuralNetworksBandPartOperation {
	rv := objc.Send[NeuralNetworksBandPartOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksBandPartOperation) Autorelease() NeuralNetworksBandPartOperation {
	rv := objc.Send[NeuralNetworksBandPartOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksBandPartOperation creates a new NeuralNetworksBandPartOperation instance.
func NewNeuralNetworksBandPartOperation() NeuralNetworksBandPartOperation {
	class := getNeuralNetworksBandPartOperationClass()
	rv := objc.Send[NeuralNetworksBandPartOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































