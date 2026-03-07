// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksEspressoV1Executor] class.
var (
	_NeuralNetworksEspressoV1ExecutorClass     NeuralNetworksEspressoV1ExecutorClass
	_NeuralNetworksEspressoV1ExecutorClassOnce sync.Once
)

func getNeuralNetworksEspressoV1ExecutorClass() NeuralNetworksEspressoV1ExecutorClass {
	_NeuralNetworksEspressoV1ExecutorClassOnce.Do(func() {
		_NeuralNetworksEspressoV1ExecutorClass = NeuralNetworksEspressoV1ExecutorClass{class: objc.GetClass("NeuralNetworks.EspressoV1Executor")}
	})
	return _NeuralNetworksEspressoV1ExecutorClass
}

// GetNeuralNetworksEspressoV1ExecutorClass returns the class object for NeuralNetworks.EspressoV1Executor.
func GetNeuralNetworksEspressoV1ExecutorClass() NeuralNetworksEspressoV1ExecutorClass {
	return getNeuralNetworksEspressoV1ExecutorClass()
}

type NeuralNetworksEspressoV1ExecutorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksEspressoV1ExecutorClass) Alloc() NeuralNetworksEspressoV1Executor {
	rv := objc.Send[NeuralNetworksEspressoV1Executor](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.EspressoV1Executor
type NeuralNetworksEspressoV1Executor struct {
	objectivec.Object
}

// NeuralNetworksEspressoV1ExecutorFromID constructs a [NeuralNetworksEspressoV1Executor] from an objc.ID.
func NeuralNetworksEspressoV1ExecutorFromID(id objc.ID) NeuralNetworksEspressoV1Executor {
	return NeuralNetworksEspressoV1Executor{objectivec.Object{id}}
}
// Ensure NeuralNetworksEspressoV1Executor implements INeuralNetworksEspressoV1Executor.
var _ INeuralNetworksEspressoV1Executor = NeuralNetworksEspressoV1Executor{}





// An interface definition for the [NeuralNetworksEspressoV1Executor] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.EspressoV1Executor
type INeuralNetworksEspressoV1Executor interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksEspressoV1Executor) Init() NeuralNetworksEspressoV1Executor {
	rv := objc.Send[NeuralNetworksEspressoV1Executor](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksEspressoV1Executor) Autorelease() NeuralNetworksEspressoV1Executor {
	rv := objc.Send[NeuralNetworksEspressoV1Executor](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksEspressoV1Executor creates a new NeuralNetworksEspressoV1Executor instance.
func NewNeuralNetworksEspressoV1Executor() NeuralNetworksEspressoV1Executor {
	class := getNeuralNetworksEspressoV1ExecutorClass()
	rv := objc.Send[NeuralNetworksEspressoV1Executor](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































