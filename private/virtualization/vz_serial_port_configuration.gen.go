// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

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
	rv := objc.SendIfResponds[VZSerialPortConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZSerialPortConfiguration._init]
//   - [VZSerialPortConfiguration._serialPort]
//   - [VZSerialPortConfiguration.MakeSerialPortForVirtualMachineSerialPortIndex]
//   - [VZSerialPortConfiguration.DebugDescription]
//   - [VZSerialPortConfiguration.Description]
//   - [VZSerialPortConfiguration.Hash]
//   - [VZSerialPortConfiguration.Superclass]
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
//   - [IVZSerialPortConfiguration._serialPort]
//   - [IVZSerialPortConfiguration.MakeSerialPortForVirtualMachineSerialPortIndex]
//   - [IVZSerialPortConfiguration.DebugDescription]
//   - [IVZSerialPortConfiguration.Description]
//   - [IVZSerialPortConfiguration.Hash]
//   - [IVZSerialPortConfiguration.Superclass]
type IVZSerialPortConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	_init() objectivec.IObject
	_serialPort() unsafe.Pointer
	MakeSerialPortForVirtualMachineSerialPortIndex(machine objectivec.IObject, index uint64) objectivec.IObject
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (v VZSerialPortConfiguration) Init() VZSerialPortConfiguration {
	rv := objc.SendIfResponds[VZSerialPortConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZSerialPortConfiguration) Autorelease() VZSerialPortConfiguration {
	rv := objc.SendIfResponds[VZSerialPortConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZSerialPortConfiguration creates a new VZSerialPortConfiguration instance.
func NewVZSerialPortConfiguration() VZSerialPortConfiguration {
	class := getVZSerialPortConfigurationClass()
	rv := objc.SendIfResponds[VZSerialPortConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VZSerialPortConfiguration) _init() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}
func (v VZSerialPortConfiguration) MakeSerialPortForVirtualMachineSerialPortIndex(machine objectivec.IObject, index uint64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("makeSerialPortForVirtualMachine:serialPortIndex:"), machine, index)
	return objectivec.Object{ID: rv}
}

func (_VZSerialPortConfigurationClass VZSerialPortConfigurationClass) SerialPortType() int64 {
	rv := objc.SendIfResponds[int64](objc.ID(_VZSerialPortConfigurationClass.class), objc.Sel("serialPortType"))
	return rv
}

func (v VZSerialPortConfiguration) _serialPort() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](v.ID, objc.Sel("_serialPort"))
	return rv
}

// CanSerialPort reports whether the receiver responds to the private selector _serialPort.
func (v VZSerialPortConfiguration) CanSerialPort() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_serialPort"))
}

// SerialPort is an exported wrapper for the private property _serialPort.
func (v VZSerialPortConfiguration) SerialPort() (unsafe.Pointer, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_serialPort")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_serialPort"}
	}
	return v._serialPort(), nil
}
func (v VZSerialPortConfiguration) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZSerialPortConfiguration) Description() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZSerialPortConfiguration) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](v.ID, objc.Sel("hash"))
	return rv
}
func (v VZSerialPortConfiguration) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](v.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
