// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksMPSHeap] class.
var (
	_NeuralNetworksMPSHeapClass     NeuralNetworksMPSHeapClass
	_NeuralNetworksMPSHeapClassOnce sync.Once
)

func getNeuralNetworksMPSHeapClass() NeuralNetworksMPSHeapClass {
	_NeuralNetworksMPSHeapClassOnce.Do(func() {
		_NeuralNetworksMPSHeapClass = NeuralNetworksMPSHeapClass{class: objc.GetClass("NeuralNetworks.MPSHeap")}
	})
	return _NeuralNetworksMPSHeapClass
}

// GetNeuralNetworksMPSHeapClass returns the class object for NeuralNetworks.MPSHeap.
func GetNeuralNetworksMPSHeapClass() NeuralNetworksMPSHeapClass {
	return getNeuralNetworksMPSHeapClass()
}

type NeuralNetworksMPSHeapClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksMPSHeapClass) Alloc() NeuralNetworksMPSHeap {
	rv := objc.Send[NeuralNetworksMPSHeap](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.MPSHeap
type NeuralNetworksMPSHeap struct {
	objectivec.Object
}

// NeuralNetworksMPSHeapFromID constructs a [NeuralNetworksMPSHeap] from an objc.ID.
func NeuralNetworksMPSHeapFromID(id objc.ID) NeuralNetworksMPSHeap {
	return NeuralNetworksMPSHeap{objectivec.Object{id}}
}
// Ensure NeuralNetworksMPSHeap implements INeuralNetworksMPSHeap.
var _ INeuralNetworksMPSHeap = NeuralNetworksMPSHeap{}





// An interface definition for the [NeuralNetworksMPSHeap] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.MPSHeap
type INeuralNetworksMPSHeap interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksMPSHeap) Init() NeuralNetworksMPSHeap {
	rv := objc.Send[NeuralNetworksMPSHeap](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksMPSHeap) Autorelease() NeuralNetworksMPSHeap {
	rv := objc.Send[NeuralNetworksMPSHeap](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksMPSHeap creates a new NeuralNetworksMPSHeap instance.
func NewNeuralNetworksMPSHeap() NeuralNetworksMPSHeap {
	class := getNeuralNetworksMPSHeapClass()
	rv := objc.Send[NeuralNetworksMPSHeap](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































