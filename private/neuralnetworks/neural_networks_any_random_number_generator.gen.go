// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksAnyRandomNumberGenerator] class.
var (
	_NeuralNetworksAnyRandomNumberGeneratorClass     NeuralNetworksAnyRandomNumberGeneratorClass
	_NeuralNetworksAnyRandomNumberGeneratorClassOnce sync.Once
)

func getNeuralNetworksAnyRandomNumberGeneratorClass() NeuralNetworksAnyRandomNumberGeneratorClass {
	_NeuralNetworksAnyRandomNumberGeneratorClassOnce.Do(func() {
		_NeuralNetworksAnyRandomNumberGeneratorClass = NeuralNetworksAnyRandomNumberGeneratorClass{class: objc.GetClass("NeuralNetworks.AnyRandomNumberGenerator")}
	})
	return _NeuralNetworksAnyRandomNumberGeneratorClass
}

// GetNeuralNetworksAnyRandomNumberGeneratorClass returns the class object for NeuralNetworks.AnyRandomNumberGenerator.
func GetNeuralNetworksAnyRandomNumberGeneratorClass() NeuralNetworksAnyRandomNumberGeneratorClass {
	return getNeuralNetworksAnyRandomNumberGeneratorClass()
}

type NeuralNetworksAnyRandomNumberGeneratorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksAnyRandomNumberGeneratorClass) Alloc() NeuralNetworksAnyRandomNumberGenerator {
	rv := objc.Send[NeuralNetworksAnyRandomNumberGenerator](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.AnyRandomNumberGenerator
type NeuralNetworksAnyRandomNumberGenerator struct {
	objectivec.Object
}

// NeuralNetworksAnyRandomNumberGeneratorFromID constructs a [NeuralNetworksAnyRandomNumberGenerator] from an objc.ID.
func NeuralNetworksAnyRandomNumberGeneratorFromID(id objc.ID) NeuralNetworksAnyRandomNumberGenerator {
	return NeuralNetworksAnyRandomNumberGenerator{objectivec.Object{id}}
}
// Ensure NeuralNetworksAnyRandomNumberGenerator implements INeuralNetworksAnyRandomNumberGenerator.
var _ INeuralNetworksAnyRandomNumberGenerator = NeuralNetworksAnyRandomNumberGenerator{}





// An interface definition for the [NeuralNetworksAnyRandomNumberGenerator] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.AnyRandomNumberGenerator
type INeuralNetworksAnyRandomNumberGenerator interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksAnyRandomNumberGenerator) Init() NeuralNetworksAnyRandomNumberGenerator {
	rv := objc.Send[NeuralNetworksAnyRandomNumberGenerator](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksAnyRandomNumberGenerator) Autorelease() NeuralNetworksAnyRandomNumberGenerator {
	rv := objc.Send[NeuralNetworksAnyRandomNumberGenerator](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksAnyRandomNumberGenerator creates a new NeuralNetworksAnyRandomNumberGenerator instance.
func NewNeuralNetworksAnyRandomNumberGenerator() NeuralNetworksAnyRandomNumberGenerator {
	class := getNeuralNetworksAnyRandomNumberGeneratorClass()
	rv := objc.Send[NeuralNetworksAnyRandomNumberGenerator](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































