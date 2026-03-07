// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksSequentialSampler] class.
var (
	_NeuralNetworksSequentialSamplerClass     NeuralNetworksSequentialSamplerClass
	_NeuralNetworksSequentialSamplerClassOnce sync.Once
)

func getNeuralNetworksSequentialSamplerClass() NeuralNetworksSequentialSamplerClass {
	_NeuralNetworksSequentialSamplerClassOnce.Do(func() {
		_NeuralNetworksSequentialSamplerClass = NeuralNetworksSequentialSamplerClass{class: objc.GetClass("NeuralNetworks.SequentialSampler")}
	})
	return _NeuralNetworksSequentialSamplerClass
}

// GetNeuralNetworksSequentialSamplerClass returns the class object for NeuralNetworks.SequentialSampler.
func GetNeuralNetworksSequentialSamplerClass() NeuralNetworksSequentialSamplerClass {
	return getNeuralNetworksSequentialSamplerClass()
}

type NeuralNetworksSequentialSamplerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksSequentialSamplerClass) Alloc() NeuralNetworksSequentialSampler {
	rv := objc.Send[NeuralNetworksSequentialSampler](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.SequentialSampler
type NeuralNetworksSequentialSampler struct {
	objectivec.Object
}

// NeuralNetworksSequentialSamplerFromID constructs a [NeuralNetworksSequentialSampler] from an objc.ID.
func NeuralNetworksSequentialSamplerFromID(id objc.ID) NeuralNetworksSequentialSampler {
	return NeuralNetworksSequentialSampler{objectivec.Object{id}}
}
// Ensure NeuralNetworksSequentialSampler implements INeuralNetworksSequentialSampler.
var _ INeuralNetworksSequentialSampler = NeuralNetworksSequentialSampler{}





// An interface definition for the [NeuralNetworksSequentialSampler] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.SequentialSampler
type INeuralNetworksSequentialSampler interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksSequentialSampler) Init() NeuralNetworksSequentialSampler {
	rv := objc.Send[NeuralNetworksSequentialSampler](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksSequentialSampler) Autorelease() NeuralNetworksSequentialSampler {
	rv := objc.Send[NeuralNetworksSequentialSampler](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksSequentialSampler creates a new NeuralNetworksSequentialSampler instance.
func NewNeuralNetworksSequentialSampler() NeuralNetworksSequentialSampler {
	class := getNeuralNetworksSequentialSamplerClass()
	rv := objc.Send[NeuralNetworksSequentialSampler](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































