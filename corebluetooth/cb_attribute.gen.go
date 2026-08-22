// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CBAttribute] class.
var (
	_CBAttributeClass     CBAttributeClass
	_CBAttributeClassOnce sync.Once
)

func getCBAttributeClass() CBAttributeClass {
	_CBAttributeClassOnce.Do(func() {
		_CBAttributeClass = CBAttributeClass{class: objc.GetClass("CBAttribute")}
	})
	return _CBAttributeClass
}

// GetCBAttributeClass returns the class object for CBAttribute.
func GetCBAttributeClass() CBAttributeClass {
	return getCBAttributeClass()
}

type CBAttributeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CBAttributeClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CBAttributeClass) Alloc() CBAttribute {
	rv := objc.Send[CBAttribute](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A representation of common aspects of services offered by a peripheral.
//
// # Overview
//
// Concrete subclasses of [CBAttribute] (and their mutable counterparts)
// represent the services a peripheral offers, the characteristics of those
// services, and the descriptors attached to those characteristics. The
// concrete subclasses are:
//
// - [CBService]
// - [CBCharacteristic]
// - [CBDescriptor]
//
// # Identifying an Attribute
//
//   - [CBAttribute.UUID]: The Bluetooth-specific UUID of the attribute.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBAttribute
type CBAttribute struct {
	objectivec.Object
}

// CBAttributeFromID constructs a [CBAttribute] from an objc.ID.
//
// A representation of common aspects of services offered by a peripheral.
func CBAttributeFromID(id objc.ID) CBAttribute {
	return CBAttribute{objectivec.Object{ID: id}}
}

// NOTE: CBAttribute adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CBAttribute] class.
//
// # Identifying an Attribute
//
//   - [ICBAttribute.UUID]: The Bluetooth-specific UUID of the attribute.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBAttribute
type ICBAttribute interface {
	objectivec.IObject

	// Topic: Identifying an Attribute

	// The Bluetooth-specific UUID of the attribute.
	UUID() ICBUUID
}

// Init initializes the instance.
func (c CBAttribute) Init() CBAttribute {
	rv := objc.Send[CBAttribute](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CBAttribute) Autorelease() CBAttribute {
	rv := objc.Send[CBAttribute](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCBAttribute creates a new CBAttribute instance.
func NewCBAttribute() CBAttribute {
	class := getCBAttributeClass()
	rv := objc.Send[CBAttribute](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The Bluetooth-specific UUID of the attribute.
//
// # Discussion
//
// This property is a 128-bit UUID that identifies the attribute.
//
// For more information, see [Core Bluetooth Programming Guide].
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBAttribute/uuid
//
// [Core Bluetooth Programming Guide]: https://developer.apple.com/library/archive/documentation/NetworkingInternetWeb/Conceptual/CoreBluetooth_concepts/AboutCoreBluetooth/Introduction.html#//apple_ref/doc/uid/TP40013257
func (c CBAttribute) UUID() ICBUUID {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("UUID"))
	return CBUUIDFromID(objc.ID(rv))
}
