// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksTensorRepresentationIdentityGenerator] class.
var (
	_NeuralNetworksTensorRepresentationIdentityGeneratorClass     NeuralNetworksTensorRepresentationIdentityGeneratorClass
	_NeuralNetworksTensorRepresentationIdentityGeneratorClassOnce sync.Once
)

func getNeuralNetworksTensorRepresentationIdentityGeneratorClass() NeuralNetworksTensorRepresentationIdentityGeneratorClass {
	_NeuralNetworksTensorRepresentationIdentityGeneratorClassOnce.Do(func() {
		_NeuralNetworksTensorRepresentationIdentityGeneratorClass = NeuralNetworksTensorRepresentationIdentityGeneratorClass{class: objc.GetClass("NeuralNetworks.TensorRepresentationIdentityGenerator")}
	})
	return _NeuralNetworksTensorRepresentationIdentityGeneratorClass
}

// GetNeuralNetworksTensorRepresentationIdentityGeneratorClass returns the class object for NeuralNetworks.TensorRepresentationIdentityGenerator.
func GetNeuralNetworksTensorRepresentationIdentityGeneratorClass() NeuralNetworksTensorRepresentationIdentityGeneratorClass {
	return getNeuralNetworksTensorRepresentationIdentityGeneratorClass()
}

type NeuralNetworksTensorRepresentationIdentityGeneratorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksTensorRepresentationIdentityGeneratorClass) Alloc() NeuralNetworksTensorRepresentationIdentityGenerator {
	rv := objc.Send[NeuralNetworksTensorRepresentationIdentityGenerator](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.TensorRepresentationIdentityGenerator
type NeuralNetworksTensorRepresentationIdentityGenerator struct {
	objectivec.Object
}

// NeuralNetworksTensorRepresentationIdentityGeneratorFromID constructs a [NeuralNetworksTensorRepresentationIdentityGenerator] from an objc.ID.
func NeuralNetworksTensorRepresentationIdentityGeneratorFromID(id objc.ID) NeuralNetworksTensorRepresentationIdentityGenerator {
	return NeuralNetworksTensorRepresentationIdentityGenerator{objectivec.Object{id}}
}
// Ensure NeuralNetworksTensorRepresentationIdentityGenerator implements INeuralNetworksTensorRepresentationIdentityGenerator.
var _ INeuralNetworksTensorRepresentationIdentityGenerator = NeuralNetworksTensorRepresentationIdentityGenerator{}





// An interface definition for the [NeuralNetworksTensorRepresentationIdentityGenerator] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.TensorRepresentationIdentityGenerator
type INeuralNetworksTensorRepresentationIdentityGenerator interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksTensorRepresentationIdentityGenerator) Init() NeuralNetworksTensorRepresentationIdentityGenerator {
	rv := objc.Send[NeuralNetworksTensorRepresentationIdentityGenerator](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksTensorRepresentationIdentityGenerator) Autorelease() NeuralNetworksTensorRepresentationIdentityGenerator {
	rv := objc.Send[NeuralNetworksTensorRepresentationIdentityGenerator](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksTensorRepresentationIdentityGenerator creates a new NeuralNetworksTensorRepresentationIdentityGenerator instance.
func NewNeuralNetworksTensorRepresentationIdentityGenerator() NeuralNetworksTensorRepresentationIdentityGenerator {
	class := getNeuralNetworksTensorRepresentationIdentityGeneratorClass()
	rv := objc.Send[NeuralNetworksTensorRepresentationIdentityGenerator](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































