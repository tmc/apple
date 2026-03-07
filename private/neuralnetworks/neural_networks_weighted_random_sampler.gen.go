// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NeuralNetworksWeightedRandomSampler] class.
var (
	_NeuralNetworksWeightedRandomSamplerClass     NeuralNetworksWeightedRandomSamplerClass
	_NeuralNetworksWeightedRandomSamplerClassOnce sync.Once
)

func getNeuralNetworksWeightedRandomSamplerClass() NeuralNetworksWeightedRandomSamplerClass {
	_NeuralNetworksWeightedRandomSamplerClassOnce.Do(func() {
		_NeuralNetworksWeightedRandomSamplerClass = NeuralNetworksWeightedRandomSamplerClass{class: objc.GetClass("NeuralNetworks.WeightedRandomSampler")}
	})
	return _NeuralNetworksWeightedRandomSamplerClass
}

// GetNeuralNetworksWeightedRandomSamplerClass returns the class object for NeuralNetworks.WeightedRandomSampler.
func GetNeuralNetworksWeightedRandomSamplerClass() NeuralNetworksWeightedRandomSamplerClass {
	return getNeuralNetworksWeightedRandomSamplerClass()
}

type NeuralNetworksWeightedRandomSamplerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksWeightedRandomSamplerClass) Alloc() NeuralNetworksWeightedRandomSampler {
	rv := objc.Send[NeuralNetworksWeightedRandomSampler](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.WeightedRandomSampler
type NeuralNetworksWeightedRandomSampler struct {
	NeuralNetworksRandomSampler
}

// NeuralNetworksWeightedRandomSamplerFromID constructs a [NeuralNetworksWeightedRandomSampler] from an objc.ID.
func NeuralNetworksWeightedRandomSamplerFromID(id objc.ID) NeuralNetworksWeightedRandomSampler {
	return NeuralNetworksWeightedRandomSampler{NeuralNetworksRandomSampler: NeuralNetworksRandomSamplerFromID(id)}
}
// Ensure NeuralNetworksWeightedRandomSampler implements INeuralNetworksWeightedRandomSampler.
var _ INeuralNetworksWeightedRandomSampler = NeuralNetworksWeightedRandomSampler{}





// An interface definition for the [NeuralNetworksWeightedRandomSampler] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.WeightedRandomSampler
type INeuralNetworksWeightedRandomSampler interface {
	INeuralNetworksRandomSampler
}





// Init initializes the instance.
func (n NeuralNetworksWeightedRandomSampler) Init() NeuralNetworksWeightedRandomSampler {
	rv := objc.Send[NeuralNetworksWeightedRandomSampler](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksWeightedRandomSampler) Autorelease() NeuralNetworksWeightedRandomSampler {
	rv := objc.Send[NeuralNetworksWeightedRandomSampler](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksWeightedRandomSampler creates a new NeuralNetworksWeightedRandomSampler instance.
func NewNeuralNetworksWeightedRandomSampler() NeuralNetworksWeightedRandomSampler {
	class := getNeuralNetworksWeightedRandomSamplerClass()
	rv := objc.Send[NeuralNetworksWeightedRandomSampler](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































