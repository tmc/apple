// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksBackpropQueue] class.
var (
	_NeuralNetworksBackpropQueueClass     NeuralNetworksBackpropQueueClass
	_NeuralNetworksBackpropQueueClassOnce sync.Once
)

func getNeuralNetworksBackpropQueueClass() NeuralNetworksBackpropQueueClass {
	_NeuralNetworksBackpropQueueClassOnce.Do(func() {
		_NeuralNetworksBackpropQueueClass = NeuralNetworksBackpropQueueClass{class: objc.GetClass("NeuralNetworks.BackpropQueue")}
	})
	return _NeuralNetworksBackpropQueueClass
}

// GetNeuralNetworksBackpropQueueClass returns the class object for NeuralNetworks.BackpropQueue.
func GetNeuralNetworksBackpropQueueClass() NeuralNetworksBackpropQueueClass {
	return getNeuralNetworksBackpropQueueClass()
}

type NeuralNetworksBackpropQueueClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksBackpropQueueClass) Alloc() NeuralNetworksBackpropQueue {
	rv := objc.Send[NeuralNetworksBackpropQueue](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.BackpropQueue
type NeuralNetworksBackpropQueue struct {
	objectivec.Object
}

// NeuralNetworksBackpropQueueFromID constructs a [NeuralNetworksBackpropQueue] from an objc.ID.
func NeuralNetworksBackpropQueueFromID(id objc.ID) NeuralNetworksBackpropQueue {
	return NeuralNetworksBackpropQueue{objectivec.Object{id}}
}
// Ensure NeuralNetworksBackpropQueue implements INeuralNetworksBackpropQueue.
var _ INeuralNetworksBackpropQueue = NeuralNetworksBackpropQueue{}





// An interface definition for the [NeuralNetworksBackpropQueue] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.BackpropQueue
type INeuralNetworksBackpropQueue interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksBackpropQueue) Init() NeuralNetworksBackpropQueue {
	rv := objc.Send[NeuralNetworksBackpropQueue](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksBackpropQueue) Autorelease() NeuralNetworksBackpropQueue {
	rv := objc.Send[NeuralNetworksBackpropQueue](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksBackpropQueue creates a new NeuralNetworksBackpropQueue instance.
func NewNeuralNetworksBackpropQueue() NeuralNetworksBackpropQueue {
	class := getNeuralNetworksBackpropQueueClass()
	rv := objc.Send[NeuralNetworksBackpropQueue](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































