// Code generated from Apple documentation for virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VZMacGuestProvisioningOptions] class.
var (
	_VZMacGuestProvisioningOptionsClass     VZMacGuestProvisioningOptionsClass
	_VZMacGuestProvisioningOptionsClassOnce sync.Once
)

func getVZMacGuestProvisioningOptionsClass() VZMacGuestProvisioningOptionsClass {
	_VZMacGuestProvisioningOptionsClassOnce.Do(func() {
		_VZMacGuestProvisioningOptionsClass = VZMacGuestProvisioningOptionsClass{class: objc.GetClass("VZMacGuestProvisioningOptions")}
	})
	return _VZMacGuestProvisioningOptionsClass
}

// GetVZMacGuestProvisioningOptionsClass returns the class object for VZMacGuestProvisioningOptions.
func GetVZMacGuestProvisioningOptionsClass() VZMacGuestProvisioningOptionsClass {
	return getVZMacGuestProvisioningOptionsClass()
}

type VZMacGuestProvisioningOptionsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZMacGuestProvisioningOptionsClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMacGuestProvisioningOptionsClass) Alloc() VZMacGuestProvisioningOptions {
	rv := objc.SendIfResponds[VZMacGuestProvisioningOptions](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZMacGuestProvisioningOptions.EnablesRemoteLogin]
//   - [VZMacGuestProvisioningOptions.SetEnablesRemoteLogin]
//   - [VZMacGuestProvisioningOptions.FullName]
//   - [VZMacGuestProvisioningOptions.SetFullName]
//   - [VZMacGuestProvisioningOptions.LogsInAutomatically]
//   - [VZMacGuestProvisioningOptions.SetLogsInAutomatically]
//   - [VZMacGuestProvisioningOptions.Password]
//   - [VZMacGuestProvisioningOptions.SetPassword]
//   - [VZMacGuestProvisioningOptions.Username]
//   - [VZMacGuestProvisioningOptions.SetUsername]
type VZMacGuestProvisioningOptions struct {
	VZGuestProvisioningOptions
}

// VZMacGuestProvisioningOptionsFromID constructs a [VZMacGuestProvisioningOptions] from an objc.ID.
func VZMacGuestProvisioningOptionsFromID(id objc.ID) VZMacGuestProvisioningOptions {
	return VZMacGuestProvisioningOptions{VZGuestProvisioningOptions: VZGuestProvisioningOptionsFromID(id)}
}

// Ensure VZMacGuestProvisioningOptions implements IVZMacGuestProvisioningOptions.
var _ IVZMacGuestProvisioningOptions = VZMacGuestProvisioningOptions{}

// An interface definition for the [VZMacGuestProvisioningOptions] class.
//
// # Methods
//
//   - [IVZMacGuestProvisioningOptions.EnablesRemoteLogin]
//   - [IVZMacGuestProvisioningOptions.SetEnablesRemoteLogin]
//   - [IVZMacGuestProvisioningOptions.FullName]
//   - [IVZMacGuestProvisioningOptions.SetFullName]
//   - [IVZMacGuestProvisioningOptions.LogsInAutomatically]
//   - [IVZMacGuestProvisioningOptions.SetLogsInAutomatically]
//   - [IVZMacGuestProvisioningOptions.Password]
//   - [IVZMacGuestProvisioningOptions.SetPassword]
//   - [IVZMacGuestProvisioningOptions.Username]
//   - [IVZMacGuestProvisioningOptions.SetUsername]
type IVZMacGuestProvisioningOptions interface {
	IVZGuestProvisioningOptions

	// Topic: Methods

	EnablesRemoteLogin() bool
	SetEnablesRemoteLogin(value bool)
	FullName() string
	SetFullName(value string)
	LogsInAutomatically() bool
	SetLogsInAutomatically(value bool)
	Password() string
	SetPassword(value string)
	Username() string
	SetUsername(value string)
}

// Init initializes the instance.
func (v VZMacGuestProvisioningOptions) Init() VZMacGuestProvisioningOptions {
	rv := objc.SendIfResponds[VZMacGuestProvisioningOptions](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZMacGuestProvisioningOptions) Autorelease() VZMacGuestProvisioningOptions {
	rv := objc.SendIfResponds[VZMacGuestProvisioningOptions](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacGuestProvisioningOptions creates a new VZMacGuestProvisioningOptions instance.
func NewVZMacGuestProvisioningOptions() VZMacGuestProvisioningOptions {
	class := getVZMacGuestProvisioningOptionsClass()
	rv := objc.SendIfResponds[VZMacGuestProvisioningOptions](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VZMacGuestProvisioningOptions) EnablesRemoteLogin() bool {
	rv := objc.SendIfResponds[bool](v.ID, objc.Sel("enablesRemoteLogin"))
	return rv
}
func (v VZMacGuestProvisioningOptions) SetEnablesRemoteLogin(value bool) {
	objc.SendIfResponds[struct{}](v.ID, objc.Sel("setEnablesRemoteLogin:"), value)
}
func (v VZMacGuestProvisioningOptions) FullName() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("fullName"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZMacGuestProvisioningOptions) SetFullName(value string) {
	objc.SendIfResponds[struct{}](v.ID, objc.Sel("setFullName:"), objc.String(value))
}
func (v VZMacGuestProvisioningOptions) LogsInAutomatically() bool {
	rv := objc.SendIfResponds[bool](v.ID, objc.Sel("logsInAutomatically"))
	return rv
}
func (v VZMacGuestProvisioningOptions) SetLogsInAutomatically(value bool) {
	objc.SendIfResponds[struct{}](v.ID, objc.Sel("setLogsInAutomatically:"), value)
}
func (v VZMacGuestProvisioningOptions) Password() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("password"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZMacGuestProvisioningOptions) SetPassword(value string) {
	objc.SendIfResponds[struct{}](v.ID, objc.Sel("setPassword:"), objc.String(value))
}
func (v VZMacGuestProvisioningOptions) Username() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("username"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZMacGuestProvisioningOptions) SetUsername(value string) {
	objc.SendIfResponds[struct{}](v.ID, objc.Sel("setUsername:"), objc.String(value))
}
