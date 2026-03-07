// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksMPSAllocator] class.
var (
	_NeuralNetworksMPSAllocatorClass     NeuralNetworksMPSAllocatorClass
	_NeuralNetworksMPSAllocatorClassOnce sync.Once
)

func getNeuralNetworksMPSAllocatorClass() NeuralNetworksMPSAllocatorClass {
	_NeuralNetworksMPSAllocatorClassOnce.Do(func() {
		_NeuralNetworksMPSAllocatorClass = NeuralNetworksMPSAllocatorClass{class: objc.GetClass("NeuralNetworks.MPSAllocator")}
	})
	return _NeuralNetworksMPSAllocatorClass
}

// GetNeuralNetworksMPSAllocatorClass returns the class object for NeuralNetworks.MPSAllocator.
func GetNeuralNetworksMPSAllocatorClass() NeuralNetworksMPSAllocatorClass {
	return getNeuralNetworksMPSAllocatorClass()
}

type NeuralNetworksMPSAllocatorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksMPSAllocatorClass) Alloc() NeuralNetworksMPSAllocator {
	rv := objc.Send[NeuralNetworksMPSAllocator](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.MPSAllocator
type NeuralNetworksMPSAllocator struct {
	objectivec.Object
}

// NeuralNetworksMPSAllocatorFromID constructs a [NeuralNetworksMPSAllocator] from an objc.ID.
func NeuralNetworksMPSAllocatorFromID(id objc.ID) NeuralNetworksMPSAllocator {
	return NeuralNetworksMPSAllocator{objectivec.Object{id}}
}
// Ensure NeuralNetworksMPSAllocator implements INeuralNetworksMPSAllocator.
var _ INeuralNetworksMPSAllocator = NeuralNetworksMPSAllocator{}





// An interface definition for the [NeuralNetworksMPSAllocator] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.MPSAllocator
type INeuralNetworksMPSAllocator interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksMPSAllocator) Init() NeuralNetworksMPSAllocator {
	rv := objc.Send[NeuralNetworksMPSAllocator](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksMPSAllocator) Autorelease() NeuralNetworksMPSAllocator {
	rv := objc.Send[NeuralNetworksMPSAllocator](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksMPSAllocator creates a new NeuralNetworksMPSAllocator instance.
func NewNeuralNetworksMPSAllocator() NeuralNetworksMPSAllocator {
	class := getNeuralNetworksMPSAllocatorClass()
	rv := objc.Send[NeuralNetworksMPSAllocator](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































