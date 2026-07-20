// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"unsafe"

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

// # Methods
//
//   - [ETTaskState.Blobs]
//   - [ETTaskState.SetBlobs]
//   - [ETTaskState.NetworkPointer]
//   - [ETTaskState.SetNetworkPointer]
//   - [ETTaskState.InitWithBlobMap]
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
type IETTaskState interface {
	objectivec.IObject

	// Topic: Methods

	Blobs() unsafe.Pointer
	SetBlobs(value unsafe.Pointer)
	NetworkPointer() unsafe.Pointer
	SetNetworkPointer(value unsafe.Pointer)
	InitWithBlobMap(map_ unsafe.Pointer) ETTaskState
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

func NewETTaskStateWithBlobMap(map_ unsafe.Pointer) ETTaskState {
	instance := getETTaskStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBlobMap:"), map_)
	return ETTaskStateFromID(rv)
}

func NewETTaskStateWithNetwork(network unsafe.Pointer) ETTaskState {
	instance := getETTaskStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithNetwork:"), network)
	return ETTaskStateFromID(rv)
}

func (e ETTaskState) InitWithBlobMap(map_ unsafe.Pointer) ETTaskState {
	rv := objc.Send[ETTaskState](e.ID, objc.Sel("initWithBlobMap:"), map_)
	return rv
}

func (e ETTaskState) Blobs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e.ID, objc.Sel("blobs"))
	return rv
}
func (e ETTaskState) SetBlobs(value unsafe.Pointer) {
	objc.Send[struct{}](e.ID, objc.Sel("setBlobs:"), value)
}
func (e ETTaskState) NetworkPointer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e.ID, objc.Sel("networkPointer"))
	return rv
}
func (e ETTaskState) SetNetworkPointer(value unsafe.Pointer) {
	objc.Send[struct{}](e.ID, objc.Sel("setNetworkPointer:"), value)
}
