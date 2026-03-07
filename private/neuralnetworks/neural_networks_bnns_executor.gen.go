// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksBNNSExecutor] class.
var (
	_NeuralNetworksBNNSExecutorClass     NeuralNetworksBNNSExecutorClass
	_NeuralNetworksBNNSExecutorClassOnce sync.Once
)

func getNeuralNetworksBNNSExecutorClass() NeuralNetworksBNNSExecutorClass {
	_NeuralNetworksBNNSExecutorClassOnce.Do(func() {
		_NeuralNetworksBNNSExecutorClass = NeuralNetworksBNNSExecutorClass{class: objc.GetClass("NeuralNetworks.BNNSExecutor")}
	})
	return _NeuralNetworksBNNSExecutorClass
}

// GetNeuralNetworksBNNSExecutorClass returns the class object for NeuralNetworks.BNNSExecutor.
func GetNeuralNetworksBNNSExecutorClass() NeuralNetworksBNNSExecutorClass {
	return getNeuralNetworksBNNSExecutorClass()
}

type NeuralNetworksBNNSExecutorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksBNNSExecutorClass) Alloc() NeuralNetworksBNNSExecutor {
	rv := objc.Send[NeuralNetworksBNNSExecutor](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.BNNSExecutor
type NeuralNetworksBNNSExecutor struct {
	objectivec.Object
}

// NeuralNetworksBNNSExecutorFromID constructs a [NeuralNetworksBNNSExecutor] from an objc.ID.
func NeuralNetworksBNNSExecutorFromID(id objc.ID) NeuralNetworksBNNSExecutor {
	return NeuralNetworksBNNSExecutor{objectivec.Object{id}}
}
// Ensure NeuralNetworksBNNSExecutor implements INeuralNetworksBNNSExecutor.
var _ INeuralNetworksBNNSExecutor = NeuralNetworksBNNSExecutor{}





// An interface definition for the [NeuralNetworksBNNSExecutor] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.BNNSExecutor
type INeuralNetworksBNNSExecutor interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksBNNSExecutor) Init() NeuralNetworksBNNSExecutor {
	rv := objc.Send[NeuralNetworksBNNSExecutor](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksBNNSExecutor) Autorelease() NeuralNetworksBNNSExecutor {
	rv := objc.Send[NeuralNetworksBNNSExecutor](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksBNNSExecutor creates a new NeuralNetworksBNNSExecutor instance.
func NewNeuralNetworksBNNSExecutor() NeuralNetworksBNNSExecutor {
	class := getNeuralNetworksBNNSExecutorClass()
	rv := objc.Send[NeuralNetworksBNNSExecutor](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































