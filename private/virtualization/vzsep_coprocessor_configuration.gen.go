// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZSEPCoprocessorConfiguration] class.
var (
	_VZSEPCoprocessorConfigurationClass     VZSEPCoprocessorConfigurationClass
	_VZSEPCoprocessorConfigurationClassOnce sync.Once
)

func getVZSEPCoprocessorConfigurationClass() VZSEPCoprocessorConfigurationClass {
	_VZSEPCoprocessorConfigurationClassOnce.Do(func() {
		_VZSEPCoprocessorConfigurationClass = VZSEPCoprocessorConfigurationClass{class: objc.GetClass("_VZSEPCoprocessorConfiguration")}
	})
	return _VZSEPCoprocessorConfigurationClass
}

// GetVZSEPCoprocessorConfigurationClass returns the class object for _VZSEPCoprocessorConfiguration.
func GetVZSEPCoprocessorConfigurationClass() VZSEPCoprocessorConfigurationClass {
	return getVZSEPCoprocessorConfigurationClass()
}

type VZSEPCoprocessorConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZSEPCoprocessorConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZSEPCoprocessorConfigurationClass) Alloc() VZSEPCoprocessorConfiguration {
	rv := objc.Send[VZSEPCoprocessorConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZSEPCoprocessorConfiguration.DebugStub]
//   - [VZSEPCoprocessorConfiguration.SetDebugStub]
//   - [VZSEPCoprocessorConfiguration.RomBinaryURL]
//   - [VZSEPCoprocessorConfiguration.SetRomBinaryURL]
//   - [VZSEPCoprocessorConfiguration.Storage]
//   - [VZSEPCoprocessorConfiguration.InitWithStorage]
//   - [VZSEPCoprocessorConfiguration.InitWithStorageURL]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZSEPCoprocessorConfiguration
type VZSEPCoprocessorConfiguration struct {
	VZCoprocessorConfiguration
}

// VZSEPCoprocessorConfigurationFromID constructs a [VZSEPCoprocessorConfiguration] from an objc.ID.
func VZSEPCoprocessorConfigurationFromID(id objc.ID) VZSEPCoprocessorConfiguration {
	return VZSEPCoprocessorConfiguration{VZCoprocessorConfiguration: VZCoprocessorConfigurationFromID(id)}
}

// Ensure VZSEPCoprocessorConfiguration implements IVZSEPCoprocessorConfiguration.
var _ IVZSEPCoprocessorConfiguration = VZSEPCoprocessorConfiguration{}

// An interface definition for the [VZSEPCoprocessorConfiguration] class.
//
// # Methods
//
//   - [IVZSEPCoprocessorConfiguration.DebugStub]
//   - [IVZSEPCoprocessorConfiguration.SetDebugStub]
//   - [IVZSEPCoprocessorConfiguration.RomBinaryURL]
//   - [IVZSEPCoprocessorConfiguration.SetRomBinaryURL]
//   - [IVZSEPCoprocessorConfiguration.Storage]
//   - [IVZSEPCoprocessorConfiguration.InitWithStorage]
//   - [IVZSEPCoprocessorConfiguration.InitWithStorageURL]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZSEPCoprocessorConfiguration
type IVZSEPCoprocessorConfiguration interface {
	IVZCoprocessorConfiguration

	// Topic: Methods

	DebugStub() IVZDebugStubConfiguration
	SetDebugStub(value IVZDebugStubConfiguration)
	RomBinaryURL() foundation.INSURL
	SetRomBinaryURL(value foundation.INSURL)
	Storage() IVZSEPStorage
	InitWithStorage(storage objectivec.IObject) VZSEPCoprocessorConfiguration
	InitWithStorageURL(url foundation.INSURL) VZSEPCoprocessorConfiguration
}

// Init initializes the instance.
func (v VZSEPCoprocessorConfiguration) Init() VZSEPCoprocessorConfiguration {
	rv := objc.Send[VZSEPCoprocessorConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZSEPCoprocessorConfiguration) Autorelease() VZSEPCoprocessorConfiguration {
	rv := objc.Send[VZSEPCoprocessorConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZSEPCoprocessorConfiguration creates a new VZSEPCoprocessorConfiguration instance.
func NewVZSEPCoprocessorConfiguration() VZSEPCoprocessorConfiguration {
	class := getVZSEPCoprocessorConfigurationClass()
	rv := objc.Send[VZSEPCoprocessorConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZSEPCoprocessorConfiguration/initWithStorage:
func NewVZSEPCoprocessorConfigurationWithStorage(storage objectivec.IObject) VZSEPCoprocessorConfiguration {
	instance := getVZSEPCoprocessorConfigurationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithStorage:"), storage)
	return VZSEPCoprocessorConfigurationFromID(rv)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZSEPCoprocessorConfiguration/initWithStorageURL:
func NewVZSEPCoprocessorConfigurationWithStorageURL(url foundation.INSURL) VZSEPCoprocessorConfiguration {
	instance := getVZSEPCoprocessorConfigurationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithStorageURL:"), url)
	return VZSEPCoprocessorConfigurationFromID(rv)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZSEPCoprocessorConfiguration/initWithStorage:
func (v VZSEPCoprocessorConfiguration) InitWithStorage(storage objectivec.IObject) VZSEPCoprocessorConfiguration {
	rv := objc.Send[VZSEPCoprocessorConfiguration](v.ID, objc.Sel("initWithStorage:"), storage)
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZSEPCoprocessorConfiguration/initWithStorageURL:
func (v VZSEPCoprocessorConfiguration) InitWithStorageURL(url foundation.INSURL) VZSEPCoprocessorConfiguration {
	rv := objc.Send[VZSEPCoprocessorConfiguration](v.ID, objc.Sel("initWithStorageURL:"), url)
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZSEPCoprocessorConfiguration/debugStub
func (v VZSEPCoprocessorConfiguration) DebugStub() IVZDebugStubConfiguration {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("debugStub"))
	return VZDebugStubConfigurationFromID(objc.ID(rv))
}
func (v VZSEPCoprocessorConfiguration) SetDebugStub(value IVZDebugStubConfiguration) {
	objc.Send[struct{}](v.ID, objc.Sel("setDebugStub:"), value)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZSEPCoprocessorConfiguration/romBinaryURL
func (v VZSEPCoprocessorConfiguration) RomBinaryURL() foundation.INSURL {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("romBinaryURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (v VZSEPCoprocessorConfiguration) SetRomBinaryURL(value foundation.INSURL) {
	objc.Send[struct{}](v.ID, objc.Sel("setRomBinaryURL:"), value)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZSEPCoprocessorConfiguration/storage
func (v VZSEPCoprocessorConfiguration) Storage() IVZSEPStorage {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("storage"))
	return VZSEPStorageFromID(objc.ID(rv))
}
