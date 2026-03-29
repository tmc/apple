// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ETTaskState] class.
var (
	_ETTaskStateClass     ETTaskStateClass
	_ETTaskStateClassOnce sync.Once
)

func getETTaskStateClass() ETTaskStateClass {
	_ETTaskStateClassOnce.Do(func() {
		_ETTaskStateClass = ETTaskStateClass{class: objc.GetClass("ETTaskState")}
	})
	return _ETTaskStateClass
}

// GetETTaskStateClass returns the class object for ETTaskState.
func GetETTaskStateClass() ETTaskStateClass {
	return getETTaskStateClass()
}

type ETTaskStateClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec ETTaskStateClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec ETTaskStateClass) Alloc() ETTaskState {
	rv := objc.Send[ETTaskState](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [ETTaskState.Blobs]
//   - [ETTaskState.SetBlobs]
//   - [ETTaskState.NetworkPointer]
//   - [ETTaskState.SetNetworkPointer]
//   - [ETTaskState.InitWithBlobMap]
//   - [ETTaskState.InitWithNetwork]
// See: https://developer.apple.com/documentation/Espresso/ETTaskState
type ETTaskState struct {
	objectivec.Object
}

// ETTaskStateFromID constructs a [ETTaskState] from an objc.ID.
func ETTaskStateFromID(id objc.ID) ETTaskState {
	return ETTaskState{objectivec.Object{ID: id}}
}
// Ensure ETTaskState implements IETTaskState.
var _ IETTaskState = ETTaskState{}

// An interface definition for the [ETTaskState] class.
//
// # Methods
//
//   - [IETTaskState.Blobs]
//   - [IETTaskState.SetBlobs]
//   - [IETTaskState.NetworkPointer]
//   - [IETTaskState.SetNetworkPointer]
//   - [IETTaskState.InitWithBlobMap]
//   - [IETTaskState.InitWithNetwork]
//
// See: https://developer.apple.com/documentation/Espresso/ETTaskState
type IETTaskState interface {
	objectivec.IObject

	// Topic: Methods

	Blobs() objectivec.IObject
	SetBlobs(value objectivec.IObject)
	NetworkPointer() objectivec.IObject
	SetNetworkPointer(value objectivec.IObject)
	InitWithBlobMap(map_ unsafe.Pointer) ETTaskState
	InitWithNetwork(network objectivec.IObject) ETTaskState
}

// Init initializes the instance.
func (e ETTaskState) Init() ETTaskState {
	rv := objc.Send[ETTaskState](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e ETTaskState) Autorelease() ETTaskState {
	rv := objc.Send[ETTaskState](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewETTaskState creates a new ETTaskState instance.
func NewETTaskState() ETTaskState {
	class := getETTaskStateClass()
	rv := objc.Send[ETTaskState](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Espresso/ETTaskState/initWithBlobMap:
func NewETTaskStateWithBlobMap(map_ unsafe.Pointer) ETTaskState {
	instance := getETTaskStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBlobMap:"), map_)
	return ETTaskStateFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Espresso/ETTaskState/initWithNetwork:
func NewETTaskStateWithNetwork(network objectivec.IObject) ETTaskState {
	instance := getETTaskStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithNetwork:"), network)
	return ETTaskStateFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Espresso/ETTaskState/initWithBlobMap:
func (e ETTaskState) InitWithBlobMap(map_ unsafe.Pointer) ETTaskState {
	rv := objc.Send[ETTaskState](e.ID, objc.Sel("initWithBlobMap:"), map_)
	return rv
}
//
// See: https://developer.apple.com/documentation/Espresso/ETTaskState/initWithNetwork:
func (e ETTaskState) InitWithNetwork(network objectivec.IObject) ETTaskState {
	rv := objc.Send[ETTaskState](e.ID, objc.Sel("initWithNetwork:"), network)
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/ETTaskState/blobs
func (e ETTaskState) Blobs() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("blobs"))
	return objectivec.Object{ID: rv}
}
func (e ETTaskState) SetBlobs(value objectivec.IObject) {
	objc.Send[struct{}](e.ID, objc.Sel("setBlobs:"), value)
}
// See: https://developer.apple.com/documentation/Espresso/ETTaskState/networkPointer
func (e ETTaskState) NetworkPointer() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("networkPointer"))
	return objectivec.Object{ID: rv}
}
func (e ETTaskState) SetNetworkPointer(value objectivec.IObject) {
	objc.Send[struct{}](e.ID, objc.Sel("setNetworkPointer:"), value)
}

