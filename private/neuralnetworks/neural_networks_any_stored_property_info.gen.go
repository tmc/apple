// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksAnyStoredPropertyInfo] class.
var (
	_NeuralNetworksAnyStoredPropertyInfoClass     NeuralNetworksAnyStoredPropertyInfoClass
	_NeuralNetworksAnyStoredPropertyInfoClassOnce sync.Once
)

func getNeuralNetworksAnyStoredPropertyInfoClass() NeuralNetworksAnyStoredPropertyInfoClass {
	_NeuralNetworksAnyStoredPropertyInfoClassOnce.Do(func() {
		_NeuralNetworksAnyStoredPropertyInfoClass = NeuralNetworksAnyStoredPropertyInfoClass{class: objc.GetClass("NeuralNetworks.AnyStoredPropertyInfo")}
	})
	return _NeuralNetworksAnyStoredPropertyInfoClass
}

// GetNeuralNetworksAnyStoredPropertyInfoClass returns the class object for NeuralNetworks.AnyStoredPropertyInfo.
func GetNeuralNetworksAnyStoredPropertyInfoClass() NeuralNetworksAnyStoredPropertyInfoClass {
	return getNeuralNetworksAnyStoredPropertyInfoClass()
}

type NeuralNetworksAnyStoredPropertyInfoClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksAnyStoredPropertyInfoClass) Alloc() NeuralNetworksAnyStoredPropertyInfo {
	rv := objc.Send[NeuralNetworksAnyStoredPropertyInfo](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.AnyStoredPropertyInfo
type NeuralNetworksAnyStoredPropertyInfo struct {
	objectivec.Object
}

// NeuralNetworksAnyStoredPropertyInfoFromID constructs a [NeuralNetworksAnyStoredPropertyInfo] from an objc.ID.
func NeuralNetworksAnyStoredPropertyInfoFromID(id objc.ID) NeuralNetworksAnyStoredPropertyInfo {
	return NeuralNetworksAnyStoredPropertyInfo{objectivec.Object{id}}
}
// Ensure NeuralNetworksAnyStoredPropertyInfo implements INeuralNetworksAnyStoredPropertyInfo.
var _ INeuralNetworksAnyStoredPropertyInfo = NeuralNetworksAnyStoredPropertyInfo{}





// An interface definition for the [NeuralNetworksAnyStoredPropertyInfo] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.AnyStoredPropertyInfo
type INeuralNetworksAnyStoredPropertyInfo interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksAnyStoredPropertyInfo) Init() NeuralNetworksAnyStoredPropertyInfo {
	rv := objc.Send[NeuralNetworksAnyStoredPropertyInfo](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksAnyStoredPropertyInfo) Autorelease() NeuralNetworksAnyStoredPropertyInfo {
	rv := objc.Send[NeuralNetworksAnyStoredPropertyInfo](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksAnyStoredPropertyInfo creates a new NeuralNetworksAnyStoredPropertyInfo instance.
func NewNeuralNetworksAnyStoredPropertyInfo() NeuralNetworksAnyStoredPropertyInfo {
	class := getNeuralNetworksAnyStoredPropertyInfoClass()
	rv := objc.Send[NeuralNetworksAnyStoredPropertyInfo](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































