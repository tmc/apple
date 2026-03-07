// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksBase2DPoolOperation] class.
var (
	_NeuralNetworksBase2DPoolOperationClass     NeuralNetworksBase2DPoolOperationClass
	_NeuralNetworksBase2DPoolOperationClassOnce sync.Once
)

func getNeuralNetworksBase2DPoolOperationClass() NeuralNetworksBase2DPoolOperationClass {
	_NeuralNetworksBase2DPoolOperationClassOnce.Do(func() {
		_NeuralNetworksBase2DPoolOperationClass = NeuralNetworksBase2DPoolOperationClass{class: objc.GetClass("NeuralNetworks.Base2DPoolOperation")}
	})
	return _NeuralNetworksBase2DPoolOperationClass
}

// GetNeuralNetworksBase2DPoolOperationClass returns the class object for NeuralNetworks.Base2DPoolOperation.
func GetNeuralNetworksBase2DPoolOperationClass() NeuralNetworksBase2DPoolOperationClass {
	return getNeuralNetworksBase2DPoolOperationClass()
}

type NeuralNetworksBase2DPoolOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksBase2DPoolOperationClass) Alloc() NeuralNetworksBase2DPoolOperation {
	rv := objc.Send[NeuralNetworksBase2DPoolOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.Base2DPoolOperation
type NeuralNetworksBase2DPoolOperation struct {
	objectivec.Object
}

// NeuralNetworksBase2DPoolOperationFromID constructs a [NeuralNetworksBase2DPoolOperation] from an objc.ID.
func NeuralNetworksBase2DPoolOperationFromID(id objc.ID) NeuralNetworksBase2DPoolOperation {
	return NeuralNetworksBase2DPoolOperation{objectivec.Object{id}}
}
// Ensure NeuralNetworksBase2DPoolOperation implements INeuralNetworksBase2DPoolOperation.
var _ INeuralNetworksBase2DPoolOperation = NeuralNetworksBase2DPoolOperation{}





// An interface definition for the [NeuralNetworksBase2DPoolOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.Base2DPoolOperation
type INeuralNetworksBase2DPoolOperation interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksBase2DPoolOperation) Init() NeuralNetworksBase2DPoolOperation {
	rv := objc.Send[NeuralNetworksBase2DPoolOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksBase2DPoolOperation) Autorelease() NeuralNetworksBase2DPoolOperation {
	rv := objc.Send[NeuralNetworksBase2DPoolOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksBase2DPoolOperation creates a new NeuralNetworksBase2DPoolOperation instance.
func NewNeuralNetworksBase2DPoolOperation() NeuralNetworksBase2DPoolOperation {
	class := getNeuralNetworksBase2DPoolOperationClass()
	rv := objc.Send[NeuralNetworksBase2DPoolOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































