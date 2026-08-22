// Code generated from Apple documentation for IOUSBHost. DO NOT EDIT.

package iousbhost

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOUSBHostIOSource] class.
var (
	_IOUSBHostIOSourceClass     IOUSBHostIOSourceClass
	_IOUSBHostIOSourceClassOnce sync.Once
)

func getIOUSBHostIOSourceClass() IOUSBHostIOSourceClass {
	_IOUSBHostIOSourceClassOnce.Do(func() {
		_IOUSBHostIOSourceClass = IOUSBHostIOSourceClass{class: objc.GetClass("IOUSBHostIOSource")}
	})
	return _IOUSBHostIOSourceClass
}

// GetIOUSBHostIOSourceClass returns the class object for IOUSBHostIOSource.
func GetIOUSBHostIOSourceClass() IOUSBHostIOSourceClass {
	return getIOUSBHostIOSourceClass()
}

type IOUSBHostIOSourceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOUSBHostIOSourceClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOUSBHostIOSourceClass) Alloc() IOUSBHostIOSource {
	rv := objc.Send[IOUSBHostIOSource](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// This class provides basic functionality for deriving pipe and stream
// classes.
//
// # Overview
//
// Don’t create objects of this class or use this class as a subclass.
// Instead, use [IOUSBHostInterface.CopyPipeWithAddressError] and
// [IOUSBHostPipe.CopyStreamWithStreamIDError] when creating an
// [IOUSBHostIOSource].
//
// # Obtaining Device Information
//
//   - [IOUSBHostIOSource.DeviceAddress]: The device’s bus address.
//   - [IOUSBHostIOSource.EndpointAddress]: The pipe or stream’s endpoint address.
//   - [IOUSBHostIOSource.HostInterface]: The interface for the input/output source.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostIOSource
type IOUSBHostIOSource struct {
	objectivec.Object
}

// IOUSBHostIOSourceFromID constructs a [IOUSBHostIOSource] from an objc.ID.
//
// This class provides basic functionality for deriving pipe and stream
// classes.
func IOUSBHostIOSourceFromID(id objc.ID) IOUSBHostIOSource {
	return IOUSBHostIOSource{objectivec.Object{ID: id}}
}

// NOTE: IOUSBHostIOSource adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [IOUSBHostIOSource] class.
//
// # Obtaining Device Information
//
//   - [IIOUSBHostIOSource.DeviceAddress]: The device’s bus address.
//   - [IIOUSBHostIOSource.EndpointAddress]: The pipe or stream’s endpoint address.
//   - [IIOUSBHostIOSource.HostInterface]: The interface for the input/output source.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostIOSource
type IIOUSBHostIOSource interface {
	objectivec.IObject

	// Topic: Obtaining Device Information

	// The device’s bus address.
	DeviceAddress() uint
	// The pipe or stream’s endpoint address.
	EndpointAddress() uint
	// The interface for the input/output source.
	HostInterface() IIOUSBHostInterface
}

// Init initializes the instance.
func (u IOUSBHostIOSource) Init() IOUSBHostIOSource {
	rv := objc.Send[IOUSBHostIOSource](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u IOUSBHostIOSource) Autorelease() IOUSBHostIOSource {
	rv := objc.Send[IOUSBHostIOSource](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOUSBHostIOSource creates a new IOUSBHostIOSource instance.
func NewIOUSBHostIOSource() IOUSBHostIOSource {
	class := getIOUSBHostIOSourceClass()
	rv := objc.Send[IOUSBHostIOSource](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The device’s bus address.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostIOSource/deviceAddress
func (u IOUSBHostIOSource) DeviceAddress() uint {
	rv := objc.Send[uint](u.ID, objc.Sel("deviceAddress"))
	return rv
}

// The pipe or stream’s endpoint address.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostIOSource/endpointAddress
func (u IOUSBHostIOSource) EndpointAddress() uint {
	rv := objc.Send[uint](u.ID, objc.Sel("endpointAddress"))
	return rv
}

// The interface for the input/output source.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostIOSource/hostInterface
func (u IOUSBHostIOSource) HostInterface() IIOUSBHostInterface {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("hostInterface"))
	return IOUSBHostInterfaceFromID(objc.ID(rv))
}
