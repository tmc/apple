// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksRuntimeConfiguration] class.
var (
	_NeuralNetworksRuntimeConfigurationClass     NeuralNetworksRuntimeConfigurationClass
	_NeuralNetworksRuntimeConfigurationClassOnce sync.Once
)

func getNeuralNetworksRuntimeConfigurationClass() NeuralNetworksRuntimeConfigurationClass {
	_NeuralNetworksRuntimeConfigurationClassOnce.Do(func() {
		_NeuralNetworksRuntimeConfigurationClass = NeuralNetworksRuntimeConfigurationClass{class: objc.GetClass("NeuralNetworks.RuntimeConfiguration")}
	})
	return _NeuralNetworksRuntimeConfigurationClass
}

// GetNeuralNetworksRuntimeConfigurationClass returns the class object for NeuralNetworks.RuntimeConfiguration.
func GetNeuralNetworksRuntimeConfigurationClass() NeuralNetworksRuntimeConfigurationClass {
	return getNeuralNetworksRuntimeConfigurationClass()
}

type NeuralNetworksRuntimeConfigurationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksRuntimeConfigurationClass) Alloc() NeuralNetworksRuntimeConfiguration {
	rv := objc.Send[NeuralNetworksRuntimeConfiguration](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.RuntimeConfiguration
type NeuralNetworksRuntimeConfiguration struct {
	objectivec.Object
}

// NeuralNetworksRuntimeConfigurationFromID constructs a [NeuralNetworksRuntimeConfiguration] from an objc.ID.
func NeuralNetworksRuntimeConfigurationFromID(id objc.ID) NeuralNetworksRuntimeConfiguration {
	return NeuralNetworksRuntimeConfiguration{objectivec.Object{id}}
}
// Ensure NeuralNetworksRuntimeConfiguration implements INeuralNetworksRuntimeConfiguration.
var _ INeuralNetworksRuntimeConfiguration = NeuralNetworksRuntimeConfiguration{}





// An interface definition for the [NeuralNetworksRuntimeConfiguration] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.RuntimeConfiguration
type INeuralNetworksRuntimeConfiguration interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksRuntimeConfiguration) Init() NeuralNetworksRuntimeConfiguration {
	rv := objc.Send[NeuralNetworksRuntimeConfiguration](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksRuntimeConfiguration) Autorelease() NeuralNetworksRuntimeConfiguration {
	rv := objc.Send[NeuralNetworksRuntimeConfiguration](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksRuntimeConfiguration creates a new NeuralNetworksRuntimeConfiguration instance.
func NewNeuralNetworksRuntimeConfiguration() NeuralNetworksRuntimeConfiguration {
	class := getNeuralNetworksRuntimeConfigurationClass()
	rv := objc.Send[NeuralNetworksRuntimeConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































