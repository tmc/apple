// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksAnyLayerVariablePathCache] class.
var (
	_NeuralNetworksAnyLayerVariablePathCacheClass     NeuralNetworksAnyLayerVariablePathCacheClass
	_NeuralNetworksAnyLayerVariablePathCacheClassOnce sync.Once
)

func getNeuralNetworksAnyLayerVariablePathCacheClass() NeuralNetworksAnyLayerVariablePathCacheClass {
	_NeuralNetworksAnyLayerVariablePathCacheClassOnce.Do(func() {
		_NeuralNetworksAnyLayerVariablePathCacheClass = NeuralNetworksAnyLayerVariablePathCacheClass{class: objc.GetClass("NeuralNetworks.AnyLayerVariablePathCache")}
	})
	return _NeuralNetworksAnyLayerVariablePathCacheClass
}

// GetNeuralNetworksAnyLayerVariablePathCacheClass returns the class object for NeuralNetworks.AnyLayerVariablePathCache.
func GetNeuralNetworksAnyLayerVariablePathCacheClass() NeuralNetworksAnyLayerVariablePathCacheClass {
	return getNeuralNetworksAnyLayerVariablePathCacheClass()
}

type NeuralNetworksAnyLayerVariablePathCacheClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksAnyLayerVariablePathCacheClass) Alloc() NeuralNetworksAnyLayerVariablePathCache {
	rv := objc.Send[NeuralNetworksAnyLayerVariablePathCache](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.AnyLayerVariablePathCache
type NeuralNetworksAnyLayerVariablePathCache struct {
	objectivec.Object
}

// NeuralNetworksAnyLayerVariablePathCacheFromID constructs a [NeuralNetworksAnyLayerVariablePathCache] from an objc.ID.
func NeuralNetworksAnyLayerVariablePathCacheFromID(id objc.ID) NeuralNetworksAnyLayerVariablePathCache {
	return NeuralNetworksAnyLayerVariablePathCache{objectivec.Object{id}}
}
// Ensure NeuralNetworksAnyLayerVariablePathCache implements INeuralNetworksAnyLayerVariablePathCache.
var _ INeuralNetworksAnyLayerVariablePathCache = NeuralNetworksAnyLayerVariablePathCache{}





// An interface definition for the [NeuralNetworksAnyLayerVariablePathCache] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.AnyLayerVariablePathCache
type INeuralNetworksAnyLayerVariablePathCache interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksAnyLayerVariablePathCache) Init() NeuralNetworksAnyLayerVariablePathCache {
	rv := objc.Send[NeuralNetworksAnyLayerVariablePathCache](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksAnyLayerVariablePathCache) Autorelease() NeuralNetworksAnyLayerVariablePathCache {
	rv := objc.Send[NeuralNetworksAnyLayerVariablePathCache](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksAnyLayerVariablePathCache creates a new NeuralNetworksAnyLayerVariablePathCache instance.
func NewNeuralNetworksAnyLayerVariablePathCache() NeuralNetworksAnyLayerVariablePathCache {
	class := getNeuralNetworksAnyLayerVariablePathCacheClass()
	rv := objc.Send[NeuralNetworksAnyLayerVariablePathCache](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































