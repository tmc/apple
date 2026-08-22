// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [CBService] class.
var (
	_CBServiceClass     CBServiceClass
	_CBServiceClassOnce sync.Once
)

func getCBServiceClass() CBServiceClass {
	_CBServiceClassOnce.Do(func() {
		_CBServiceClass = CBServiceClass{class: objc.GetClass("CBService")}
	})
	return _CBServiceClass
}

// GetCBServiceClass returns the class object for CBService.
func GetCBServiceClass() CBServiceClass {
	return getCBServiceClass()
}

type CBServiceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CBServiceClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CBServiceClass) Alloc() CBService {
	rv := objc.Send[CBService](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A collection of data and associated behaviors that accomplish a function or
// feature of a device.
//
// # Overview
//
// [CBService] objects represent services of a remote peripheral. Services are
// either primary or secondary and may contain multiple characteristics or
// included services (references to other services).
//
// # Identifying a Service
//
//   - [CBService.Peripheral]: The peripheral to which this service belongs.
//   - [CBService.IsPrimary]: A Boolean value that indicates whether the type of service is primary or secondary.
//
// # Accessing Service Data
//
//   - [CBService.Characteristics]: A list of characteristics discovered in this service.
//   - [CBService.IncludedServices]: A list of included services discovered in this service.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBService
type CBService struct {
	CBAttribute
}

// CBServiceFromID constructs a [CBService] from an objc.ID.
//
// A collection of data and associated behaviors that accomplish a function or
// feature of a device.
func CBServiceFromID(id objc.ID) CBService {
	return CBService{CBAttribute: CBAttributeFromID(id)}
}

// NOTE: CBService adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CBService] class.
//
// # Identifying a Service
//
//   - [ICBService.Peripheral]: The peripheral to which this service belongs.
//   - [ICBService.IsPrimary]: A Boolean value that indicates whether the type of service is primary or secondary.
//
// # Accessing Service Data
//
//   - [ICBService.Characteristics]: A list of characteristics discovered in this service.
//   - [ICBService.IncludedServices]: A list of included services discovered in this service.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBService
type ICBService interface {
	ICBAttribute

	// Topic: Identifying a Service

	// The peripheral to which this service belongs.
	Peripheral() ICBPeripheral
	// A Boolean value that indicates whether the type of service is primary or secondary.
	IsPrimary() bool

	// Topic: Accessing Service Data

	// A list of characteristics discovered in this service.
	Characteristics() []CBCharacteristic
	// A list of included services discovered in this service.
	IncludedServices() []CBService
}

// Init initializes the instance.
func (c CBService) Init() CBService {
	rv := objc.Send[CBService](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CBService) Autorelease() CBService {
	rv := objc.Send[CBService](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCBService creates a new CBService instance.
func NewCBService() CBService {
	class := getCBServiceClass()
	rv := objc.Send[CBService](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The peripheral to which this service belongs.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBService/peripheral
func (c CBService) Peripheral() ICBPeripheral {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("peripheral"))
	return CBPeripheralFromID(objc.ID(rv))
}

// A Boolean value that indicates whether the type of service is primary or
// secondary.
//
// # Discussion
//
// A peripheral’s service is either primary or secondary. A primary service
// describes the primary function of a device. A secondary service describes a
// service that’s relevant only in the context of another service that
// references it. For example, the primary service of a heart rate monitor may
// be to expose heart rate data from the monitor’s heart rate sensor. In
// this example, a secondary service may be to expose the sensor’s battery
// data.
//
// If the value of this property is true, the type of service is primary. If
// the value of this property is false, the type of service is secondary.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBService/isPrimary
func (c CBService) IsPrimary() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isPrimary"))
	return rv
}

// A list of characteristics discovered in this service.
//
// # Discussion
//
// This array contains [CBCharacteristic] objects that represent a service’s
// characteristics. Characteristics provide further details about a
// peripheral’s service. For example, a heart rate service may contain one
// characteristic that describes the intended body location of the device’s
// heart rate sensor, while another characteristic transmits heart rate
// measurement data.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBService/characteristics
func (c CBService) Characteristics() []CBCharacteristic {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("characteristics"))
	return objc.ConvertSlice(rv, func(id objc.ID) CBCharacteristic {
		return CBCharacteristicFromID(id)
	})
}

// A list of included services discovered in this service.
//
// # Discussion
//
// This array contains [CBService] objects that represent the included
// services of a service. A service of a peripheral may contain a reference to
// other services that are available on the peripheral. These other services
// are the included services of the service. You discover included services
// using the [CBPeripheral.DiscoverIncludedServicesForService] method of the
// [CBPeripheral] class.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBService/includedServices
func (c CBService) IncludedServices() []CBService {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("includedServices"))
	return objc.ConvertSlice(rv, func(id objc.ID) CBService {
		return CBServiceFromID(id)
	})
}
