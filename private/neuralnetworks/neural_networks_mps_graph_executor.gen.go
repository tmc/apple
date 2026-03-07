// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksMPSGraphExecutor] class.
var (
	_NeuralNetworksMPSGraphExecutorClass     NeuralNetworksMPSGraphExecutorClass
	_NeuralNetworksMPSGraphExecutorClassOnce sync.Once
)

func getNeuralNetworksMPSGraphExecutorClass() NeuralNetworksMPSGraphExecutorClass {
	_NeuralNetworksMPSGraphExecutorClassOnce.Do(func() {
		_NeuralNetworksMPSGraphExecutorClass = NeuralNetworksMPSGraphExecutorClass{class: objc.GetClass("NeuralNetworks.MPSGraphExecutor")}
	})
	return _NeuralNetworksMPSGraphExecutorClass
}

// GetNeuralNetworksMPSGraphExecutorClass returns the class object for NeuralNetworks.MPSGraphExecutor.
func GetNeuralNetworksMPSGraphExecutorClass() NeuralNetworksMPSGraphExecutorClass {
	return getNeuralNetworksMPSGraphExecutorClass()
}

type NeuralNetworksMPSGraphExecutorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksMPSGraphExecutorClass) Alloc() NeuralNetworksMPSGraphExecutor {
	rv := objc.Send[NeuralNetworksMPSGraphExecutor](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.MPSGraphExecutor
type NeuralNetworksMPSGraphExecutor struct {
	objectivec.Object
}

// NeuralNetworksMPSGraphExecutorFromID constructs a [NeuralNetworksMPSGraphExecutor] from an objc.ID.
func NeuralNetworksMPSGraphExecutorFromID(id objc.ID) NeuralNetworksMPSGraphExecutor {
	return NeuralNetworksMPSGraphExecutor{objectivec.Object{id}}
}
// Ensure NeuralNetworksMPSGraphExecutor implements INeuralNetworksMPSGraphExecutor.
var _ INeuralNetworksMPSGraphExecutor = NeuralNetworksMPSGraphExecutor{}





// An interface definition for the [NeuralNetworksMPSGraphExecutor] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.MPSGraphExecutor
type INeuralNetworksMPSGraphExecutor interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksMPSGraphExecutor) Init() NeuralNetworksMPSGraphExecutor {
	rv := objc.Send[NeuralNetworksMPSGraphExecutor](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksMPSGraphExecutor) Autorelease() NeuralNetworksMPSGraphExecutor {
	rv := objc.Send[NeuralNetworksMPSGraphExecutor](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksMPSGraphExecutor creates a new NeuralNetworksMPSGraphExecutor instance.
func NewNeuralNetworksMPSGraphExecutor() NeuralNetworksMPSGraphExecutor {
	class := getNeuralNetworksMPSGraphExecutorClass()
	rv := objc.Send[NeuralNetworksMPSGraphExecutor](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































