// Code generated from Apple documentation for virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZEFISignatureDatabaseConfiguration] class.
var (
	_VZEFISignatureDatabaseConfigurationClass     VZEFISignatureDatabaseConfigurationClass
	_VZEFISignatureDatabaseConfigurationClassOnce sync.Once
)

func getVZEFISignatureDatabaseConfigurationClass() VZEFISignatureDatabaseConfigurationClass {
	_VZEFISignatureDatabaseConfigurationClassOnce.Do(func() {
		_VZEFISignatureDatabaseConfigurationClass = VZEFISignatureDatabaseConfigurationClass{class: objc.GetClass("VZEFISignatureDatabaseConfiguration")}
	})
	return _VZEFISignatureDatabaseConfigurationClass
}

// GetVZEFISignatureDatabaseConfigurationClass returns the class object for VZEFISignatureDatabaseConfiguration.
func GetVZEFISignatureDatabaseConfigurationClass() VZEFISignatureDatabaseConfigurationClass {
	return getVZEFISignatureDatabaseConfigurationClass()
}

type VZEFISignatureDatabaseConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZEFISignatureDatabaseConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZEFISignatureDatabaseConfigurationClass) Alloc() VZEFISignatureDatabaseConfiguration {
	rv := objc.SendIfResponds[VZEFISignatureDatabaseConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZEFISignatureDatabaseConfiguration.DbSignatures]
//   - [VZEFISignatureDatabaseConfiguration.DbxSignatures]
//   - [VZEFISignatureDatabaseConfiguration.KeyExchangeKeys]
//   - [VZEFISignatureDatabaseConfiguration.InitWithKeyExchangeKeysDbSignaturesDbxSignatures]
type VZEFISignatureDatabaseConfiguration struct {
	objectivec.Object
}

// VZEFISignatureDatabaseConfigurationFromID constructs a [VZEFISignatureDatabaseConfiguration] from an objc.ID.
func VZEFISignatureDatabaseConfigurationFromID(id objc.ID) VZEFISignatureDatabaseConfiguration {
	return VZEFISignatureDatabaseConfiguration{objectivec.Object{ID: id}}
}

// Ensure VZEFISignatureDatabaseConfiguration implements IVZEFISignatureDatabaseConfiguration.
var _ IVZEFISignatureDatabaseConfiguration = VZEFISignatureDatabaseConfiguration{}

// An interface definition for the [VZEFISignatureDatabaseConfiguration] class.
//
// # Methods
//
//   - [IVZEFISignatureDatabaseConfiguration.DbSignatures]
//   - [IVZEFISignatureDatabaseConfiguration.DbxSignatures]
//   - [IVZEFISignatureDatabaseConfiguration.KeyExchangeKeys]
//   - [IVZEFISignatureDatabaseConfiguration.InitWithKeyExchangeKeysDbSignaturesDbxSignatures]
type IVZEFISignatureDatabaseConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	DbSignatures() foundation.INSArray
	DbxSignatures() foundation.INSArray
	KeyExchangeKeys() foundation.INSArray
	InitWithKeyExchangeKeysDbSignaturesDbxSignatures(keys objectivec.IObject, signatures objectivec.IObject, signatures2 objectivec.IObject) VZEFISignatureDatabaseConfiguration
}

// Init initializes the instance.
func (v VZEFISignatureDatabaseConfiguration) Init() VZEFISignatureDatabaseConfiguration {
	rv := objc.SendIfResponds[VZEFISignatureDatabaseConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZEFISignatureDatabaseConfiguration) Autorelease() VZEFISignatureDatabaseConfiguration {
	rv := objc.SendIfResponds[VZEFISignatureDatabaseConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZEFISignatureDatabaseConfiguration creates a new VZEFISignatureDatabaseConfiguration instance.
func NewVZEFISignatureDatabaseConfiguration() VZEFISignatureDatabaseConfiguration {
	class := getVZEFISignatureDatabaseConfigurationClass()
	rv := objc.SendIfResponds[VZEFISignatureDatabaseConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewVZEFISignatureDatabaseConfigurationWithKeyExchangeKeysDbSignaturesDbxSignatures(keys objectivec.IObject, signatures objectivec.IObject, signatures2 objectivec.IObject) VZEFISignatureDatabaseConfiguration {
	instance := getVZEFISignatureDatabaseConfigurationClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithKeyExchangeKeys:dbSignatures:dbxSignatures:"), keys, signatures, signatures2)
	return VZEFISignatureDatabaseConfigurationFromID(rv)
}

func (v VZEFISignatureDatabaseConfiguration) InitWithKeyExchangeKeysDbSignaturesDbxSignatures(keys objectivec.IObject, signatures objectivec.IObject, signatures2 objectivec.IObject) VZEFISignatureDatabaseConfiguration {
	rv := objc.SendIfResponds[VZEFISignatureDatabaseConfiguration](v.ID, objc.Sel("initWithKeyExchangeKeys:dbSignatures:dbxSignatures:"), keys, signatures, signatures2)
	return rv
}

func (v VZEFISignatureDatabaseConfiguration) DbSignatures() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("dbSignatures"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (v VZEFISignatureDatabaseConfiguration) DbxSignatures() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("dbxSignatures"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (v VZEFISignatureDatabaseConfiguration) KeyExchangeKeys() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("keyExchangeKeys"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
