// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CBManager] class.
var (
	_CBManagerClass     CBManagerClass
	_CBManagerClassOnce sync.Once
)

func getCBManagerClass() CBManagerClass {
	_CBManagerClassOnce.Do(func() {
		_CBManagerClass = CBManagerClass{class: objc.GetClass("CBManager")}
	})
	return _CBManagerClass
}

// GetCBManagerClass returns the class object for CBManager.
func GetCBManagerClass() CBManagerClass {
	return getCBManagerClass()
}

type CBManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CBManagerClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CBManagerClass) Alloc() CBManager {
	rv := objc.Send[CBManager](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// The abstract base class that manages central and peripheral objects.
//
// # Accessing the Manager’s Properties
//
//   - [CBManager.State]: The current state of the manager.
//
// # Deprecated Properties
//
//   - [CBManager.Authorization]: The current authorization status for using Bluetooth.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBManager
type CBManager struct {
	objectivec.Object
}

// CBManagerFromID constructs a [CBManager] from an objc.ID.
//
// The abstract base class that manages central and peripheral objects.
func CBManagerFromID(id objc.ID) CBManager {
	return CBManager{objectivec.Object{ID: id}}
}

// NOTE: CBManager adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CBManager] class.
//
// # Accessing the Manager’s Properties
//
//   - [ICBManager.State]: The current state of the manager.
//
// # Deprecated Properties
//
//   - [ICBManager.Authorization]: The current authorization status for using Bluetooth.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBManager
type ICBManager interface {
	objectivec.IObject

	// Topic: Accessing the Manager’s Properties

	// The current state of the manager.
	State() CBManagerState

	// Topic: Deprecated Properties

	// The current authorization status for using Bluetooth.
	Authorization() CBManagerAuthorization
}

// Init initializes the instance.
func (c CBManager) Init() CBManager {
	rv := objc.Send[CBManager](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CBManager) Autorelease() CBManager {
	rv := objc.Send[CBManager](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCBManager creates a new CBManager instance.
func NewCBManager() CBManager {
	class := getCBManagerClass()
	rv := objc.Send[CBManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The current state of the manager.
//
// # Discussion
//
// This state is initially set to [CBManagerState.unknown]. When the state
// updates, the manager calls its delegate’s [CentralManagerDidUpdateState]
// method.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBManager/state
//
// [CBManagerState.unknown]: https://developer.apple.com/documentation/CoreBluetooth/CBManagerState/unknown
func (c CBManager) State() CBManagerState {
	rv := objc.Send[CBManagerState](c.ID, objc.Sel("state"))
	return CBManagerState(rv)
}

// The current authorization status for using Bluetooth.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBManager/authorization-swift.property
func (c CBManager) Authorization() CBManagerAuthorization {
	rv := objc.Send[CBManagerAuthorization](c.ID, objc.Sel("authorization"))
	return CBManagerAuthorization(rv)
}
