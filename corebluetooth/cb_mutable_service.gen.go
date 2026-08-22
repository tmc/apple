// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [CBMutableService] class.
var (
	_CBMutableServiceClass     CBMutableServiceClass
	_CBMutableServiceClassOnce sync.Once
)

func getCBMutableServiceClass() CBMutableServiceClass {
	_CBMutableServiceClassOnce.Do(func() {
		_CBMutableServiceClass = CBMutableServiceClass{class: objc.GetClass("CBMutableService")}
	})
	return _CBMutableServiceClass
}

// GetCBMutableServiceClass returns the class object for CBMutableService.
func GetCBMutableServiceClass() CBMutableServiceClass {
	return getCBMutableServiceClass()
}

type CBMutableServiceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CBMutableServiceClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CBMutableServiceClass) Alloc() CBMutableService {
	rv := objc.Send[CBMutableService](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A service with writeable property values.
//
// # Overview
//
// The [CBMutableService] class adds write access to all of the properties in
// the [CBService] class it inherits from. You use this class to create a
// service or an included service on a local peripheral device (represented by
// a [CBPeripheralManager] object). After creating a service, you can add it
// to the peripheral’s local database using the
// [CBPeripheralManager.AddService] method of the [CBPeripheralManager] class.
// After you add a service to the peripheral’s local database, Core
// Bluetooth caches the service and you can no longer make changes to it.
//
// # Creating a Mutable Service
//
//   - [CBMutableService.InitWithTypePrimary]: Creates a newly initialized mutable service specified by UUID and service type.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBMutableService
type CBMutableService struct {
	CBService
}

// CBMutableServiceFromID constructs a [CBMutableService] from an objc.ID.
//
// A service with writeable property values.
func CBMutableServiceFromID(id objc.ID) CBMutableService {
	return CBMutableService{CBService: CBServiceFromID(id)}
}

// NOTE: CBMutableService adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CBMutableService] class.
//
// # Creating a Mutable Service
//
//   - [ICBMutableService.InitWithTypePrimary]: Creates a newly initialized mutable service specified by UUID and service type.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBMutableService
type ICBMutableService interface {
	ICBService

	// Topic: Creating a Mutable Service

	// Creates a newly initialized mutable service specified by UUID and service type.
	InitWithTypePrimary(UUID ICBUUID, isPrimary bool) CBMutableService
}

// Init initializes the instance.
func (c CBMutableService) Init() CBMutableService {
	rv := objc.Send[CBMutableService](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CBMutableService) Autorelease() CBMutableService {
	rv := objc.Send[CBMutableService](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCBMutableService creates a new CBMutableService instance.
func NewCBMutableService() CBMutableService {
	class := getCBMutableServiceClass()
	rv := objc.Send[CBMutableService](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a newly initialized mutable service specified by UUID and service
// type.
//
// UUID: A 128-bit UUID that identifies the service.
//
// isPrimary: A Boolean value that indicates whether the type of service is primary or
// secondary. If the value is true, the type of service is primary. If the
// value is false, the type of service is secondary.
//
// # Return Value
//
// A newly initialized mutable service.
//
// # Discussion
//
// For more information, see [Core Bluetooth Programming Guide].
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBMutableService/init(type:primary:)
//
// [Core Bluetooth Programming Guide]: https://developer.apple.com/library/archive/documentation/NetworkingInternetWeb/Conceptual/CoreBluetooth_concepts/AboutCoreBluetooth/Introduction.html#//apple_ref/doc/uid/TP40013257
func NewCBMutableServiceWithTypePrimary(UUID ICBUUID, isPrimary bool) CBMutableService {
	instance := getCBMutableServiceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithType:primary:"), UUID, isPrimary)
	return CBMutableServiceFromID(rv)
}

// Creates a newly initialized mutable service specified by UUID and service
// type.
//
// UUID: A 128-bit UUID that identifies the service.
//
// isPrimary: A Boolean value that indicates whether the type of service is primary or
// secondary. If the value is true, the type of service is primary. If the
// value is false, the type of service is secondary.
//
// # Return Value
//
// A newly initialized mutable service.
//
// # Discussion
//
// For more information, see [Core Bluetooth Programming Guide].
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBMutableService/init(type:primary:)
//
// [Core Bluetooth Programming Guide]: https://developer.apple.com/library/archive/documentation/NetworkingInternetWeb/Conceptual/CoreBluetooth_concepts/AboutCoreBluetooth/Introduction.html#//apple_ref/doc/uid/TP40013257
func (c CBMutableService) InitWithTypePrimary(UUID ICBUUID, isPrimary bool) CBMutableService {
	rv := objc.Send[CBMutableService](c.ID, objc.Sel("initWithType:primary:"), UUID, isPrimary)
	return rv
}
