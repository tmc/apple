// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksGradientAccumulator] class.
var (
	_NeuralNetworksGradientAccumulatorClass     NeuralNetworksGradientAccumulatorClass
	_NeuralNetworksGradientAccumulatorClassOnce sync.Once
)

func getNeuralNetworksGradientAccumulatorClass() NeuralNetworksGradientAccumulatorClass {
	_NeuralNetworksGradientAccumulatorClassOnce.Do(func() {
		_NeuralNetworksGradientAccumulatorClass = NeuralNetworksGradientAccumulatorClass{class: objc.GetClass("NeuralNetworks.GradientAccumulator")}
	})
	return _NeuralNetworksGradientAccumulatorClass
}

// GetNeuralNetworksGradientAccumulatorClass returns the class object for NeuralNetworks.GradientAccumulator.
func GetNeuralNetworksGradientAccumulatorClass() NeuralNetworksGradientAccumulatorClass {
	return getNeuralNetworksGradientAccumulatorClass()
}

type NeuralNetworksGradientAccumulatorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksGradientAccumulatorClass) Alloc() NeuralNetworksGradientAccumulator {
	rv := objc.Send[NeuralNetworksGradientAccumulator](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.GradientAccumulator
type NeuralNetworksGradientAccumulator struct {
	objectivec.Object
}

// NeuralNetworksGradientAccumulatorFromID constructs a [NeuralNetworksGradientAccumulator] from an objc.ID.
func NeuralNetworksGradientAccumulatorFromID(id objc.ID) NeuralNetworksGradientAccumulator {
	return NeuralNetworksGradientAccumulator{objectivec.Object{id}}
}
// Ensure NeuralNetworksGradientAccumulator implements INeuralNetworksGradientAccumulator.
var _ INeuralNetworksGradientAccumulator = NeuralNetworksGradientAccumulator{}





// An interface definition for the [NeuralNetworksGradientAccumulator] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.GradientAccumulator
type INeuralNetworksGradientAccumulator interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksGradientAccumulator) Init() NeuralNetworksGradientAccumulator {
	rv := objc.Send[NeuralNetworksGradientAccumulator](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksGradientAccumulator) Autorelease() NeuralNetworksGradientAccumulator {
	rv := objc.Send[NeuralNetworksGradientAccumulator](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksGradientAccumulator creates a new NeuralNetworksGradientAccumulator instance.
func NewNeuralNetworksGradientAccumulator() NeuralNetworksGradientAccumulator {
	class := getNeuralNetworksGradientAccumulatorClass()
	rv := objc.Send[NeuralNetworksGradientAccumulator](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































