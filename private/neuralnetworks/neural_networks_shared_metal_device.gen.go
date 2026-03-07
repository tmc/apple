// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksSharedMetalDevice] class.
var (
	_NeuralNetworksSharedMetalDeviceClass     NeuralNetworksSharedMetalDeviceClass
	_NeuralNetworksSharedMetalDeviceClassOnce sync.Once
)

func getNeuralNetworksSharedMetalDeviceClass() NeuralNetworksSharedMetalDeviceClass {
	_NeuralNetworksSharedMetalDeviceClassOnce.Do(func() {
		_NeuralNetworksSharedMetalDeviceClass = NeuralNetworksSharedMetalDeviceClass{class: objc.GetClass("NeuralNetworks.SharedMetalDevice")}
	})
	return _NeuralNetworksSharedMetalDeviceClass
}

// GetNeuralNetworksSharedMetalDeviceClass returns the class object for NeuralNetworks.SharedMetalDevice.
func GetNeuralNetworksSharedMetalDeviceClass() NeuralNetworksSharedMetalDeviceClass {
	return getNeuralNetworksSharedMetalDeviceClass()
}

type NeuralNetworksSharedMetalDeviceClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksSharedMetalDeviceClass) Alloc() NeuralNetworksSharedMetalDevice {
	rv := objc.Send[NeuralNetworksSharedMetalDevice](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.SharedMetalDevice
type NeuralNetworksSharedMetalDevice struct {
	objectivec.Object
}

// NeuralNetworksSharedMetalDeviceFromID constructs a [NeuralNetworksSharedMetalDevice] from an objc.ID.
func NeuralNetworksSharedMetalDeviceFromID(id objc.ID) NeuralNetworksSharedMetalDevice {
	return NeuralNetworksSharedMetalDevice{objectivec.Object{id}}
}
// Ensure NeuralNetworksSharedMetalDevice implements INeuralNetworksSharedMetalDevice.
var _ INeuralNetworksSharedMetalDevice = NeuralNetworksSharedMetalDevice{}





// An interface definition for the [NeuralNetworksSharedMetalDevice] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.SharedMetalDevice
type INeuralNetworksSharedMetalDevice interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksSharedMetalDevice) Init() NeuralNetworksSharedMetalDevice {
	rv := objc.Send[NeuralNetworksSharedMetalDevice](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksSharedMetalDevice) Autorelease() NeuralNetworksSharedMetalDevice {
	rv := objc.Send[NeuralNetworksSharedMetalDevice](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksSharedMetalDevice creates a new NeuralNetworksSharedMetalDevice instance.
func NewNeuralNetworksSharedMetalDevice() NeuralNetworksSharedMetalDevice {
	class := getNeuralNetworksSharedMetalDeviceClass()
	rv := objc.Send[NeuralNetworksSharedMetalDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































