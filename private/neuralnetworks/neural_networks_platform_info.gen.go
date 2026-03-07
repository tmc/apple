// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksPlatformInfo] class.
var (
	_NeuralNetworksPlatformInfoClass     NeuralNetworksPlatformInfoClass
	_NeuralNetworksPlatformInfoClassOnce sync.Once
)

func getNeuralNetworksPlatformInfoClass() NeuralNetworksPlatformInfoClass {
	_NeuralNetworksPlatformInfoClassOnce.Do(func() {
		_NeuralNetworksPlatformInfoClass = NeuralNetworksPlatformInfoClass{class: objc.GetClass("NeuralNetworks.PlatformInfo")}
	})
	return _NeuralNetworksPlatformInfoClass
}

// GetNeuralNetworksPlatformInfoClass returns the class object for NeuralNetworks.PlatformInfo.
func GetNeuralNetworksPlatformInfoClass() NeuralNetworksPlatformInfoClass {
	return getNeuralNetworksPlatformInfoClass()
}

type NeuralNetworksPlatformInfoClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksPlatformInfoClass) Alloc() NeuralNetworksPlatformInfo {
	rv := objc.Send[NeuralNetworksPlatformInfo](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.PlatformInfo
type NeuralNetworksPlatformInfo struct {
	objectivec.Object
}

// NeuralNetworksPlatformInfoFromID constructs a [NeuralNetworksPlatformInfo] from an objc.ID.
func NeuralNetworksPlatformInfoFromID(id objc.ID) NeuralNetworksPlatformInfo {
	return NeuralNetworksPlatformInfo{objectivec.Object{id}}
}
// Ensure NeuralNetworksPlatformInfo implements INeuralNetworksPlatformInfo.
var _ INeuralNetworksPlatformInfo = NeuralNetworksPlatformInfo{}





// An interface definition for the [NeuralNetworksPlatformInfo] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.PlatformInfo
type INeuralNetworksPlatformInfo interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksPlatformInfo) Init() NeuralNetworksPlatformInfo {
	rv := objc.Send[NeuralNetworksPlatformInfo](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksPlatformInfo) Autorelease() NeuralNetworksPlatformInfo {
	rv := objc.Send[NeuralNetworksPlatformInfo](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksPlatformInfo creates a new NeuralNetworksPlatformInfo instance.
func NewNeuralNetworksPlatformInfo() NeuralNetworksPlatformInfo {
	class := getNeuralNetworksPlatformInfoClass()
	rv := objc.Send[NeuralNetworksPlatformInfo](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































