// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [EspressoNetwork] class.
var (
	_EspressoNetworkClass     EspressoNetworkClass
	_EspressoNetworkClassOnce sync.Once
)

func getEspressoNetworkClass() EspressoNetworkClass {
	_EspressoNetworkClassOnce.Do(func() {
		_EspressoNetworkClass = EspressoNetworkClass{class: objc.GetClass("EspressoNetwork")}
	})
	return _EspressoNetworkClass
}

// GetEspressoNetworkClass returns the class object for EspressoNetwork.
func GetEspressoNetworkClass() EspressoNetworkClass {
	return getEspressoNetworkClass()
}

type EspressoNetworkClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoNetworkClass) Alloc() EspressoNetwork {
	rv := objc.Send[EspressoNetwork](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [EspressoNetwork.Ctx]
//   - [EspressoNetwork.Layers_size]
//   - [EspressoNetwork.Net]
//   - [EspressoNetwork.Wipe_layers_blobs]
//   - [EspressoNetwork.InitWithJSFileBinSerializerIdContextComputePath]
//   - [EspressoNetwork.InitWithJSFileContextComputePath]
// See: https://developer.apple.com/documentation/Espresso/EspressoNetwork
type EspressoNetwork struct {
	objectivec.Object
}

// EspressoNetworkFromID constructs a [EspressoNetwork] from an objc.ID.
func EspressoNetworkFromID(id objc.ID) EspressoNetwork {
	return EspressoNetwork{objectivec.Object{ID: id}}
}
// Ensure EspressoNetwork implements IEspressoNetwork.
var _ IEspressoNetwork = EspressoNetwork{}

// An interface definition for the [EspressoNetwork] class.
//
// # Methods
//
//   - [IEspressoNetwork.Ctx]
//   - [IEspressoNetwork.Layers_size]
//   - [IEspressoNetwork.Net]
//   - [IEspressoNetwork.Wipe_layers_blobs]
//   - [IEspressoNetwork.InitWithJSFileBinSerializerIdContextComputePath]
//   - [IEspressoNetwork.InitWithJSFileContextComputePath]
//
// See: https://developer.apple.com/documentation/Espresso/EspressoNetwork
type IEspressoNetwork interface {
	objectivec.IObject

	// Topic: Methods

	Ctx() IEspressoContext
	Layers_size() uint64
	Net() objectivec.IObject
	Wipe_layers_blobs()
	InitWithJSFileBinSerializerIdContextComputePath(jSFile []byte, id []byte, context objectivec.IObject, path int) EspressoNetwork
	InitWithJSFileContextComputePath(jSFile []byte, context objectivec.IObject, path int) EspressoNetwork
}

// Init initializes the instance.
func (e EspressoNetwork) Init() EspressoNetwork {
	rv := objc.Send[EspressoNetwork](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoNetwork) Autorelease() EspressoNetwork {
	rv := objc.Send[EspressoNetwork](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoNetwork creates a new EspressoNetwork instance.
func NewEspressoNetwork() EspressoNetwork {
	class := getEspressoNetworkClass()
	rv := objc.Send[EspressoNetwork](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Espresso/EspressoNetwork/initWithJSFile:binSerializerId:context:computePath:
func NewEspressoNetworkWithJSFileBinSerializerIdContextComputePath(jSFile []byte, id []byte, context objectivec.IObject, path int) EspressoNetwork {
	instance := getEspressoNetworkClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithJSFile:binSerializerId:context:computePath:"), unsafe.Pointer(unsafe.SliceData(jSFile)), unsafe.Pointer(unsafe.SliceData(id)), context, path)
	return EspressoNetworkFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Espresso/EspressoNetwork/initWithJSFile:context:computePath:
func NewEspressoNetworkWithJSFileContextComputePath(jSFile []byte, context objectivec.IObject, path int) EspressoNetwork {
	instance := getEspressoNetworkClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithJSFile:context:computePath:"), unsafe.Pointer(unsafe.SliceData(jSFile)), context, path)
	return EspressoNetworkFromID(rv)
}

// See: https://developer.apple.com/documentation/Espresso/EspressoNetwork/wipe_layers_blobs
func (e EspressoNetwork) Wipe_layers_blobs() {
	objc.Send[objc.ID](e.ID, objc.Sel("wipe_layers_blobs"))
}
//
// See: https://developer.apple.com/documentation/Espresso/EspressoNetwork/initWithJSFile:binSerializerId:context:computePath:
func (e EspressoNetwork) InitWithJSFileBinSerializerIdContextComputePath(jSFile []byte, id []byte, context objectivec.IObject, path int) EspressoNetwork {
	rv := objc.Send[EspressoNetwork](e.ID, objc.Sel("initWithJSFile:binSerializerId:context:computePath:"), unsafe.Pointer(unsafe.SliceData(jSFile)), unsafe.Pointer(unsafe.SliceData(id)), context, path)
	return rv
}
//
// See: https://developer.apple.com/documentation/Espresso/EspressoNetwork/initWithJSFile:context:computePath:
func (e EspressoNetwork) InitWithJSFileContextComputePath(jSFile []byte, context objectivec.IObject, path int) EspressoNetwork {
	rv := objc.Send[EspressoNetwork](e.ID, objc.Sel("initWithJSFile:context:computePath:"), unsafe.Pointer(unsafe.SliceData(jSFile)), context, path)
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoNetwork/ctx
func (e EspressoNetwork) Ctx() IEspressoContext {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("ctx"))
	return EspressoContextFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/Espresso/EspressoNetwork/layers_size
func (e EspressoNetwork) Layers_size() uint64 {
	rv := objc.Send[uint64](e.ID, objc.Sel("layers_size"))
	return rv
}
// See: https://developer.apple.com/documentation/Espresso/EspressoNetwork/net
func (e EspressoNetwork) Net() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("net"))
	return objectivec.Object{ID: rv}
}

