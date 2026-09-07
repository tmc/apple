// Code generated from Apple documentation for virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VZCustomGuestProvisioningOptions] class.
var (
	_VZCustomGuestProvisioningOptionsClass     VZCustomGuestProvisioningOptionsClass
	_VZCustomGuestProvisioningOptionsClassOnce sync.Once
)

func getVZCustomGuestProvisioningOptionsClass() VZCustomGuestProvisioningOptionsClass {
	_VZCustomGuestProvisioningOptionsClassOnce.Do(func() {
		_VZCustomGuestProvisioningOptionsClass = VZCustomGuestProvisioningOptionsClass{class: objc.GetClass("_VZCustomGuestProvisioningOptions")}
	})
	return _VZCustomGuestProvisioningOptionsClass
}

// GetVZCustomGuestProvisioningOptionsClass returns the class object for _VZCustomGuestProvisioningOptions.
func GetVZCustomGuestProvisioningOptionsClass() VZCustomGuestProvisioningOptionsClass {
	return getVZCustomGuestProvisioningOptionsClass()
}

type VZCustomGuestProvisioningOptionsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZCustomGuestProvisioningOptionsClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZCustomGuestProvisioningOptionsClass) Alloc() VZCustomGuestProvisioningOptions {
	rv := objc.SendIfResponds[VZCustomGuestProvisioningOptions](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZCustomGuestProvisioningOptions.ApplicationIdentifier]
//   - [VZCustomGuestProvisioningOptions.SetApplicationIdentifier]
//   - [VZCustomGuestProvisioningOptions.Data]
//   - [VZCustomGuestProvisioningOptions.SetData]
//   - [VZCustomGuestProvisioningOptions.RequiredEntitlement]
//   - [VZCustomGuestProvisioningOptions.SetRequiredEntitlement]
type VZCustomGuestProvisioningOptions struct {
	VZGuestProvisioningOptions
}

// VZCustomGuestProvisioningOptionsFromID constructs a [VZCustomGuestProvisioningOptions] from an objc.ID.
func VZCustomGuestProvisioningOptionsFromID(id objc.ID) VZCustomGuestProvisioningOptions {
	return VZCustomGuestProvisioningOptions{VZGuestProvisioningOptions: VZGuestProvisioningOptionsFromID(id)}
}

// Ensure VZCustomGuestProvisioningOptions implements IVZCustomGuestProvisioningOptions.
var _ IVZCustomGuestProvisioningOptions = VZCustomGuestProvisioningOptions{}

// An interface definition for the [VZCustomGuestProvisioningOptions] class.
//
// # Methods
//
//   - [IVZCustomGuestProvisioningOptions.ApplicationIdentifier]
//   - [IVZCustomGuestProvisioningOptions.SetApplicationIdentifier]
//   - [IVZCustomGuestProvisioningOptions.Data]
//   - [IVZCustomGuestProvisioningOptions.SetData]
//   - [IVZCustomGuestProvisioningOptions.RequiredEntitlement]
//   - [IVZCustomGuestProvisioningOptions.SetRequiredEntitlement]
type IVZCustomGuestProvisioningOptions interface {
	IVZGuestProvisioningOptions

	// Topic: Methods

	ApplicationIdentifier() string
	SetApplicationIdentifier(value string)
	Data() foundation.NSData
	SetData(value foundation.NSData)
	RequiredEntitlement() string
	SetRequiredEntitlement(value string)
}

// Init initializes the instance.
func (v VZCustomGuestProvisioningOptions) Init() VZCustomGuestProvisioningOptions {
	rv := objc.SendIfResponds[VZCustomGuestProvisioningOptions](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZCustomGuestProvisioningOptions) Autorelease() VZCustomGuestProvisioningOptions {
	rv := objc.SendIfResponds[VZCustomGuestProvisioningOptions](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZCustomGuestProvisioningOptions creates a new VZCustomGuestProvisioningOptions instance.
func NewVZCustomGuestProvisioningOptions() VZCustomGuestProvisioningOptions {
	class := getVZCustomGuestProvisioningOptionsClass()
	rv := objc.SendIfResponds[VZCustomGuestProvisioningOptions](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VZCustomGuestProvisioningOptions) ApplicationIdentifier() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("applicationIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZCustomGuestProvisioningOptions) SetApplicationIdentifier(value string) {
	objc.SendIfResponds[struct{}](v.ID, objc.Sel("setApplicationIdentifier:"), objc.String(value))
}
func (v VZCustomGuestProvisioningOptions) Data() foundation.NSData {
	rv := objc.SendIfResponds[foundation.NSData](v.ID, objc.Sel("data"))
	return foundation.NSData(rv)
}
func (v VZCustomGuestProvisioningOptions) SetData(value foundation.NSData) {
	objc.SendIfResponds[struct{}](v.ID, objc.Sel("setData:"), value)
}
func (v VZCustomGuestProvisioningOptions) RequiredEntitlement() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("requiredEntitlement"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZCustomGuestProvisioningOptions) SetRequiredEntitlement(value string) {
	objc.SendIfResponds[struct{}](v.ID, objc.Sel("setRequiredEntitlement:"), objc.String(value))
}
