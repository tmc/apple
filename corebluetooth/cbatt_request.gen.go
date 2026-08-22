// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CBATTRequest] class.
var (
	_CBATTRequestClass     CBATTRequestClass
	_CBATTRequestClassOnce sync.Once
)

func getCBATTRequestClass() CBATTRequestClass {
	_CBATTRequestClassOnce.Do(func() {
		_CBATTRequestClass = CBATTRequestClass{class: objc.GetClass("CBATTRequest")}
	})
	return _CBATTRequestClass
}

// GetCBATTRequestClass returns the class object for CBATTRequest.
func GetCBATTRequestClass() CBATTRequestClass {
	return getCBATTRequestClass()
}

type CBATTRequestClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CBATTRequestClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CBATTRequestClass) Alloc() CBATTRequest {
	rv := objc.Send[CBATTRequest](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A request that uses the Attribute Protocol (ATT).
//
// # Overview
//
// The [CBATTRequest] class represents Attribute Protocol (ATT) read and write
// requests from remote central devices (represented by [CBCentral] objects).
// Remote centrals use these ATT requests to read and write characteristic
// values on local peripherals (represented by [CBPeripheralManager] objects).
// Local peripherals, on the other hand, use the properties of [CBATTRequest]
// objects to respond to the read and write requests appropriately, using the
// [CBPeripheralManager.RespondToRequestWithResult] method of the
// [CBPeripheralManager] class.
//
// # Requesting to Read and Write Characteristic Values
//
//   - [CBATTRequest.Central]: The remote central device that originated the request.
//   - [CBATTRequest.Characteristic]: The characteristic to read or write the value of.
//   - [CBATTRequest.Value]: The data that the central reads from or writes to the peripheral.
//   - [CBATTRequest.SetValue]
//   - [CBATTRequest.Offset]: The zero-based index of the first byte for the read or write request.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBATTRequest
type CBATTRequest struct {
	objectivec.Object
}

// CBATTRequestFromID constructs a [CBATTRequest] from an objc.ID.
//
// A request that uses the Attribute Protocol (ATT).
func CBATTRequestFromID(id objc.ID) CBATTRequest {
	return CBATTRequest{objectivec.Object{ID: id}}
}

// NOTE: CBATTRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CBATTRequest] class.
//
// # Requesting to Read and Write Characteristic Values
//
//   - [ICBATTRequest.Central]: The remote central device that originated the request.
//   - [ICBATTRequest.Characteristic]: The characteristic to read or write the value of.
//   - [ICBATTRequest.Value]: The data that the central reads from or writes to the peripheral.
//   - [ICBATTRequest.SetValue]
//   - [ICBATTRequest.Offset]: The zero-based index of the first byte for the read or write request.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBATTRequest
type ICBATTRequest interface {
	objectivec.IObject

	// Topic: Requesting to Read and Write Characteristic Values

	// The remote central device that originated the request.
	Central() ICBCentral
	// The characteristic to read or write the value of.
	Characteristic() ICBCharacteristic
	// The data that the central reads from or writes to the peripheral.
	Value() foundation.NSData
	SetValue(value foundation.NSData)
	// The zero-based index of the first byte for the read or write request.
	Offset() uint
}

// Init initializes the instance.
func (c CBATTRequest) Init() CBATTRequest {
	rv := objc.Send[CBATTRequest](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CBATTRequest) Autorelease() CBATTRequest {
	rv := objc.Send[CBATTRequest](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCBATTRequest creates a new CBATTRequest instance.
func NewCBATTRequest() CBATTRequest {
	class := getCBATTRequestClass()
	rv := objc.Send[CBATTRequest](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The remote central device that originated the request.
//
// # Overview
//
// For more information, see [Core Bluetooth Programming Guide].
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBATTRequest/central
//
// [Core Bluetooth Programming Guide]: https://developer.apple.com/library/archive/documentation/NetworkingInternetWeb/Conceptual/CoreBluetooth_concepts/AboutCoreBluetooth/Introduction.html#//apple_ref/doc/uid/TP40013257
func (c CBATTRequest) Central() ICBCentral {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("central"))
	return CBCentralFromID(objc.ID(rv))
}

// The characteristic to read or write the value of.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBATTRequest/characteristic
func (c CBATTRequest) Characteristic() ICBCharacteristic {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("characteristic"))
	return CBCharacteristicFromID(objc.ID(rv))
}

// The data that the central reads from or writes to the peripheral.
//
// # Discussion
//
// The value of this property depends on whether the request type is read or
// write. For read requests, the property is `nil,` and you should set it
// before responding to the remote central through the
// [CBPeripheralManager.RespondToRequestWithResult] method. For write
// requests, the value is the data to write to the characteristic’s value.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBATTRequest/value
func (c CBATTRequest) Value() foundation.NSData {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("value"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (c CBATTRequest) SetValue(value foundation.NSData) {
	objc.Send[struct{}](c.ID, objc.Sel("setValue:"), value)
}

// The zero-based index of the first byte for the read or write request.
//
// # Discussion
//
// You can use the value of this property to ensure that the ATT request is
// attempting to read or write within the proper bounds of the
// characteristic’s value. For an example of how to take a request’s
// offset property into account when responding to a read or write request,
// see [Responding to Read and Write Requests from a Central].
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBATTRequest/offset
//
// [Responding to Read and Write Requests from a Central]: https://developer.apple.com/library/archive/documentation/NetworkingInternetWeb/Conceptual/CoreBluetooth_concepts/PerformingCommonPeripheralRoleTasks/PerformingCommonPeripheralRoleTasks.html#//apple_ref/doc/uid/TP40013257-CH4-SW6
func (c CBATTRequest) Offset() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("offset"))
	return rv
}
