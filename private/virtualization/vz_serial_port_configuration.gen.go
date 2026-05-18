// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZSerialPortConfiguration] class.
var (
	_VZSerialPortConfigurationClass     VZSerialPortConfigurationClass
	_VZSerialPortConfigurationClassOnce sync.Once
)

func getVZSerialPortConfigurationClass() VZSerialPortConfigurationClass {
	_VZSerialPortConfigurationClassOnce.Do(func() {
		_VZSerialPortConfigurationClass = VZSerialPortConfigurationClass{class: objc.GetClass("VZSerialPortConfiguration")}
	})
	return _VZSerialPortConfigurationClass
}

// GetVZSerialPortConfigurationClass returns the class object for VZSerialPortConfiguration.
func GetVZSerialPortConfigurationClass() VZSerialPortConfigurationClass {
	return getVZSerialPortConfigurationClass()
}

type VZSerialPortConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZSerialPortConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZSerialPortConfigurationClass) Alloc() VZSerialPortConfiguration {
	rv := objc.Send[VZSerialPortConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZSerialPortConfiguration._init]
//   - [VZSerialPortConfiguration.MakeSerialPortForVirtualMachineSerialPortIndex]
//   - [VZSerialPortConfiguration._serialPort]
//   - [VZSerialPortConfiguration.DebugDescription]
//   - [VZSerialPortConfiguration.Description]
//   - [VZSerialPortConfiguration.Hash]
//   - [VZSerialPortConfiguration.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/VZSerialPortConfiguration
type VZSerialPortConfiguration struct {
	objectivec.Object
}

// VZSerialPortConfigurationFromID constructs a [VZSerialPortConfiguration] from an objc.ID.
func VZSerialPortConfigurationFromID(id objc.ID) VZSerialPortConfiguration {
	return VZSerialPortConfiguration{objectivec.Object{ID: id}}
}

// Ensure VZSerialPortConfiguration implements IVZSerialPortConfiguration.
var _ IVZSerialPortConfiguration = VZSerialPortConfiguration{}

// An interface definition for the [VZSerialPortConfiguration] class.
//
// # Methods
//
//   - [IVZSerialPortConfiguration._init]
//   - [IVZSerialPortConfiguration.MakeSerialPortForVirtualMachineSerialPortIndex]
//   - [IVZSerialPortConfiguration._serialPort]
//   - [IVZSerialPortConfiguration.DebugDescription]
//   - [IVZSerialPortConfiguration.Description]
//   - [IVZSerialPortConfiguration.Hash]
//   - [IVZSerialPortConfiguration.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/VZSerialPortConfiguration
type IVZSerialPortConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	_init() objectivec.IObject
	MakeSerialPortForVirtualMachineSerialPortIndex(machine objectivec.IObject, index uint64) objectivec.IObject
	_serialPort() objectivec.IObject
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (v VZSerialPortConfiguration) Init() VZSerialPortConfiguration {
	rv := objc.Send[VZSerialPortConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZSerialPortConfiguration) Autorelease() VZSerialPortConfiguration {
	rv := objc.Send[VZSerialPortConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZSerialPortConfiguration creates a new VZSerialPortConfiguration instance.
func NewVZSerialPortConfiguration() VZSerialPortConfiguration {
	class := getVZSerialPortConfigurationClass()
	rv := objc.Send[VZSerialPortConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZSerialPortConfiguration/_init
func (v VZSerialPortConfiguration) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/VZSerialPortConfiguration/makeSerialPortForVirtualMachine:serialPortIndex:
func (v VZSerialPortConfiguration) MakeSerialPortForVirtualMachineSerialPortIndex(machine objectivec.IObject, index uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("makeSerialPortForVirtualMachine:serialPortIndex:"), machine, index)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/VZSerialPortConfiguration/serialPortType
func (_VZSerialPortConfigurationClass VZSerialPortConfigurationClass) SerialPortType() int64 {
	rv := objc.Send[int64](objc.ID(_VZSerialPortConfigurationClass.class), objc.Sel("serialPortType"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZSerialPortConfiguration/_serialPort
func (v VZSerialPortConfiguration) _serialPort() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_serialPort"))
	return objectivec.Object{ID: rv}
}

// CanSerialPort reports whether the receiver responds to the private selector _serialPort.
func (v VZSerialPortConfiguration) CanSerialPort() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_serialPort"))
}

// SerialPort is an exported wrapper for the private property _serialPort.
func (v VZSerialPortConfiguration) SerialPort() (objectivec.IObject, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_serialPort")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_serialPort"}
	}
	return v._serialPort(), nil
}

// See: https://developer.apple.com/documentation/Virtualization/VZSerialPortConfiguration/debugDescription
func (v VZSerialPortConfiguration) DebugDescription() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Virtualization/VZSerialPortConfiguration/description
func (v VZSerialPortConfiguration) Description() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Virtualization/VZSerialPortConfiguration/hash
func (v VZSerialPortConfiguration) Hash() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZSerialPortConfiguration/superclass
func (v VZSerialPortConfiguration) Superclass() objc.Class {
	rv := objc.Send[objc.Class](v.ID, objc.Sel("superclass"))
	return rv
}
