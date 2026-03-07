// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksRandomSampler] class.
var (
	_NeuralNetworksRandomSamplerClass     NeuralNetworksRandomSamplerClass
	_NeuralNetworksRandomSamplerClassOnce sync.Once
)

func getNeuralNetworksRandomSamplerClass() NeuralNetworksRandomSamplerClass {
	_NeuralNetworksRandomSamplerClassOnce.Do(func() {
		_NeuralNetworksRandomSamplerClass = NeuralNetworksRandomSamplerClass{class: objc.GetClass("NeuralNetworks.RandomSampler")}
	})
	return _NeuralNetworksRandomSamplerClass
}

// GetNeuralNetworksRandomSamplerClass returns the class object for NeuralNetworks.RandomSampler.
func GetNeuralNetworksRandomSamplerClass() NeuralNetworksRandomSamplerClass {
	return getNeuralNetworksRandomSamplerClass()
}

type NeuralNetworksRandomSamplerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksRandomSamplerClass) Alloc() NeuralNetworksRandomSampler {
	rv := objc.Send[NeuralNetworksRandomSampler](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.RandomSampler
type NeuralNetworksRandomSampler struct {
	objectivec.Object
}

// NeuralNetworksRandomSamplerFromID constructs a [NeuralNetworksRandomSampler] from an objc.ID.
func NeuralNetworksRandomSamplerFromID(id objc.ID) NeuralNetworksRandomSampler {
	return NeuralNetworksRandomSampler{objectivec.Object{id}}
}
// Ensure NeuralNetworksRandomSampler implements INeuralNetworksRandomSampler.
var _ INeuralNetworksRandomSampler = NeuralNetworksRandomSampler{}





// An interface definition for the [NeuralNetworksRandomSampler] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.RandomSampler
type INeuralNetworksRandomSampler interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksRandomSampler) Init() NeuralNetworksRandomSampler {
	rv := objc.Send[NeuralNetworksRandomSampler](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksRandomSampler) Autorelease() NeuralNetworksRandomSampler {
	rv := objc.Send[NeuralNetworksRandomSampler](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksRandomSampler creates a new NeuralNetworksRandomSampler instance.
func NewNeuralNetworksRandomSampler() NeuralNetworksRandomSampler {
	class := getNeuralNetworksRandomSamplerClass()
	rv := objc.Send[NeuralNetworksRandomSampler](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































