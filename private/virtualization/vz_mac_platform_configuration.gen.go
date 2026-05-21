// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZMacPlatformConfiguration] class.
var (
	_VZMacPlatformConfigurationClass     VZMacPlatformConfigurationClass
	_VZMacPlatformConfigurationClassOnce sync.Once
)

func getVZMacPlatformConfigurationClass() VZMacPlatformConfigurationClass {
	_VZMacPlatformConfigurationClassOnce.Do(func() {
		_VZMacPlatformConfigurationClass = VZMacPlatformConfigurationClass{class: objc.GetClass("VZMacPlatformConfiguration")}
	})
	return _VZMacPlatformConfigurationClass
}

// GetVZMacPlatformConfigurationClass returns the class object for VZMacPlatformConfiguration.
func GetVZMacPlatformConfigurationClass() VZMacPlatformConfigurationClass {
	return getVZMacPlatformConfigurationClass()
}

type VZMacPlatformConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZMacPlatformConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMacPlatformConfigurationClass) Alloc() VZMacPlatformConfiguration {
	rv := objc.Send[VZMacPlatformConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZMacPlatformConfiguration._guestEncryptionWrappingKey]
//   - [VZMacPlatformConfiguration.Set_guestEncryptionWrappingKey]
//   - [VZMacPlatformConfiguration._hostAttributeShareOptions]
//   - [VZMacPlatformConfiguration.Set_hostAttributeShareOptions]
//   - [VZMacPlatformConfiguration._isFairPlayEnabled]
//   - [VZMacPlatformConfiguration._isFakeEncryptionEnabled]
//   - [VZMacPlatformConfiguration._isProductionModeEnabled]
//   - [VZMacPlatformConfiguration._isSIODescramblerEnabled]
//   - [VZMacPlatformConfiguration._isStrongIdentityEnabled]
//   - [VZMacPlatformConfiguration._remoteServiceDiscoveryConfiguration]
//   - [VZMacPlatformConfiguration.Set_remoteServiceDiscoveryConfiguration]
//   - [VZMacPlatformConfiguration._setFairPlayEnabled]
//   - [VZMacPlatformConfiguration._setFakeEncryptionEnabled]
//   - [VZMacPlatformConfiguration._setGuestEncryptionWrappingKey]
//   - [VZMacPlatformConfiguration._setHostAttributeShareOptions]
//   - [VZMacPlatformConfiguration._setProductionModeEnabled]
//   - [VZMacPlatformConfiguration._setRemoteServiceDiscoveryConfiguration]
//   - [VZMacPlatformConfiguration._setSIODescramblerEnabled]
//   - [VZMacPlatformConfiguration._setStrongIdentityEnabled]
//   - [VZMacPlatformConfiguration._fairPlayEnabled]
//   - [VZMacPlatformConfiguration.Set_fairPlayEnabled]
//   - [VZMacPlatformConfiguration._fakeEncryptionEnabled]
//   - [VZMacPlatformConfiguration.Set_fakeEncryptionEnabled]
//   - [VZMacPlatformConfiguration._productionModeEnabled]
//   - [VZMacPlatformConfiguration.Set_productionModeEnabled]
//   - [VZMacPlatformConfiguration._sioDescramblerEnabled]
//   - [VZMacPlatformConfiguration.Set_sioDescramblerEnabled]
//   - [VZMacPlatformConfiguration._strongIdentityEnabled]
//   - [VZMacPlatformConfiguration.Set_strongIdentityEnabled]
type VZMacPlatformConfiguration struct {
	VZPlatformConfiguration
}

// VZMacPlatformConfigurationFromID constructs a [VZMacPlatformConfiguration] from an objc.ID.
func VZMacPlatformConfigurationFromID(id objc.ID) VZMacPlatformConfiguration {
	return VZMacPlatformConfiguration{VZPlatformConfiguration: VZPlatformConfigurationFromID(id)}
}

// Ensure VZMacPlatformConfiguration implements IVZMacPlatformConfiguration.
var _ IVZMacPlatformConfiguration = VZMacPlatformConfiguration{}

// An interface definition for the [VZMacPlatformConfiguration] class.
//
// # Methods
//
//   - [IVZMacPlatformConfiguration._guestEncryptionWrappingKey]
//   - [IVZMacPlatformConfiguration.Set_guestEncryptionWrappingKey]
//   - [IVZMacPlatformConfiguration._hostAttributeShareOptions]
//   - [IVZMacPlatformConfiguration.Set_hostAttributeShareOptions]
//   - [IVZMacPlatformConfiguration._isFairPlayEnabled]
//   - [IVZMacPlatformConfiguration._isFakeEncryptionEnabled]
//   - [IVZMacPlatformConfiguration._isProductionModeEnabled]
//   - [IVZMacPlatformConfiguration._isSIODescramblerEnabled]
//   - [IVZMacPlatformConfiguration._isStrongIdentityEnabled]
//   - [IVZMacPlatformConfiguration._remoteServiceDiscoveryConfiguration]
//   - [IVZMacPlatformConfiguration.Set_remoteServiceDiscoveryConfiguration]
//   - [IVZMacPlatformConfiguration._setFairPlayEnabled]
//   - [IVZMacPlatformConfiguration._setFakeEncryptionEnabled]
//   - [IVZMacPlatformConfiguration._setGuestEncryptionWrappingKey]
//   - [IVZMacPlatformConfiguration._setHostAttributeShareOptions]
//   - [IVZMacPlatformConfiguration._setProductionModeEnabled]
//   - [IVZMacPlatformConfiguration._setRemoteServiceDiscoveryConfiguration]
//   - [IVZMacPlatformConfiguration._setSIODescramblerEnabled]
//   - [IVZMacPlatformConfiguration._setStrongIdentityEnabled]
//   - [IVZMacPlatformConfiguration._fairPlayEnabled]
//   - [IVZMacPlatformConfiguration.Set_fairPlayEnabled]
//   - [IVZMacPlatformConfiguration._fakeEncryptionEnabled]
//   - [IVZMacPlatformConfiguration.Set_fakeEncryptionEnabled]
//   - [IVZMacPlatformConfiguration._productionModeEnabled]
//   - [IVZMacPlatformConfiguration.Set_productionModeEnabled]
//   - [IVZMacPlatformConfiguration._sioDescramblerEnabled]
//   - [IVZMacPlatformConfiguration.Set_sioDescramblerEnabled]
//   - [IVZMacPlatformConfiguration._strongIdentityEnabled]
//   - [IVZMacPlatformConfiguration.Set_strongIdentityEnabled]
type IVZMacPlatformConfiguration interface {
	IVZPlatformConfiguration

	// Topic: Methods

	_guestEncryptionWrappingKey() IVZWrappingKey
	Set_guestEncryptionWrappingKey(value IVZWrappingKey)
	_hostAttributeShareOptions() uint64
	Set_hostAttributeShareOptions(value uint64)
	_isFairPlayEnabled() bool
	_isFakeEncryptionEnabled() bool
	_isProductionModeEnabled() bool
	_isSIODescramblerEnabled() bool
	_isStrongIdentityEnabled() bool
	_remoteServiceDiscoveryConfiguration() IVZMacRemoteServiceDiscoveryConfiguration
	Set_remoteServiceDiscoveryConfiguration(value IVZMacRemoteServiceDiscoveryConfiguration)
	_setFairPlayEnabled(enabled bool)
	_setFakeEncryptionEnabled(enabled bool)
	_setGuestEncryptionWrappingKey(key objectivec.IObject)
	_setHostAttributeShareOptions(options uint64)
	_setProductionModeEnabled(enabled bool)
	_setRemoteServiceDiscoveryConfiguration(configuration objectivec.IObject)
	_setSIODescramblerEnabled(enabled bool)
	_setStrongIdentityEnabled(enabled bool)
	_fairPlayEnabled() bool
	Set_fairPlayEnabled(value bool)
	_fakeEncryptionEnabled() bool
	Set_fakeEncryptionEnabled(value bool)
	_productionModeEnabled() bool
	Set_productionModeEnabled(value bool)
	_sioDescramblerEnabled() bool
	Set_sioDescramblerEnabled(value bool)
	_strongIdentityEnabled() bool
	Set_strongIdentityEnabled(value bool)
}

// Init initializes the instance.
func (v VZMacPlatformConfiguration) Init() VZMacPlatformConfiguration {
	rv := objc.Send[VZMacPlatformConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZMacPlatformConfiguration) Autorelease() VZMacPlatformConfiguration {
	rv := objc.Send[VZMacPlatformConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacPlatformConfiguration creates a new VZMacPlatformConfiguration instance.
func NewVZMacPlatformConfiguration() VZMacPlatformConfiguration {
	class := getVZMacPlatformConfigurationClass()
	rv := objc.Send[VZMacPlatformConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VZMacPlatformConfiguration) _isFairPlayEnabled() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("_isFairPlayEnabled"))
	return rv
}

// IsFairPlayEnabled is an exported wrapper for the private method _isFairPlayEnabled.
func (v VZMacPlatformConfiguration) IsFairPlayEnabled() (bool, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_isFairPlayEnabled")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_isFairPlayEnabled"}
		return false, err
	}
	return v._isFairPlayEnabled(), nil
}

// CanIsFairPlayEnabled reports whether the receiver responds to the private selector _isFairPlayEnabled.
func (v VZMacPlatformConfiguration) CanIsFairPlayEnabled() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_isFairPlayEnabled"))
}
func (v VZMacPlatformConfiguration) _isFakeEncryptionEnabled() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("_isFakeEncryptionEnabled"))
	return rv
}

// IsFakeEncryptionEnabled is an exported wrapper for the private method _isFakeEncryptionEnabled.
func (v VZMacPlatformConfiguration) IsFakeEncryptionEnabled() (bool, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_isFakeEncryptionEnabled")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_isFakeEncryptionEnabled"}
		return false, err
	}
	return v._isFakeEncryptionEnabled(), nil
}

// CanIsFakeEncryptionEnabled reports whether the receiver responds to the private selector _isFakeEncryptionEnabled.
func (v VZMacPlatformConfiguration) CanIsFakeEncryptionEnabled() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_isFakeEncryptionEnabled"))
}
func (v VZMacPlatformConfiguration) _isProductionModeEnabled() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("_isProductionModeEnabled"))
	return rv
}

// IsProductionModeEnabled is an exported wrapper for the private method _isProductionModeEnabled.
func (v VZMacPlatformConfiguration) IsProductionModeEnabled() (bool, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_isProductionModeEnabled")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_isProductionModeEnabled"}
		return false, err
	}
	return v._isProductionModeEnabled(), nil
}

// CanIsProductionModeEnabled reports whether the receiver responds to the private selector _isProductionModeEnabled.
func (v VZMacPlatformConfiguration) CanIsProductionModeEnabled() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_isProductionModeEnabled"))
}
func (v VZMacPlatformConfiguration) _isSIODescramblerEnabled() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("_isSIODescramblerEnabled"))
	return rv
}

// IsSIODescramblerEnabled is an exported wrapper for the private method _isSIODescramblerEnabled.
func (v VZMacPlatformConfiguration) IsSIODescramblerEnabled() (bool, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_isSIODescramblerEnabled")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_isSIODescramblerEnabled"}
		return false, err
	}
	return v._isSIODescramblerEnabled(), nil
}

// CanIsSIODescramblerEnabled reports whether the receiver responds to the private selector _isSIODescramblerEnabled.
func (v VZMacPlatformConfiguration) CanIsSIODescramblerEnabled() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_isSIODescramblerEnabled"))
}
func (v VZMacPlatformConfiguration) _isStrongIdentityEnabled() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("_isStrongIdentityEnabled"))
	return rv
}

// IsStrongIdentityEnabled is an exported wrapper for the private method _isStrongIdentityEnabled.
func (v VZMacPlatformConfiguration) IsStrongIdentityEnabled() (bool, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_isStrongIdentityEnabled")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_isStrongIdentityEnabled"}
		return false, err
	}
	return v._isStrongIdentityEnabled(), nil
}

// CanIsStrongIdentityEnabled reports whether the receiver responds to the private selector _isStrongIdentityEnabled.
func (v VZMacPlatformConfiguration) CanIsStrongIdentityEnabled() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_isStrongIdentityEnabled"))
}
func (v VZMacPlatformConfiguration) _setFairPlayEnabled(enabled bool) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setFairPlayEnabled:"), enabled)
}

// SetFairPlayEnabled is an exported wrapper for the private method _setFairPlayEnabled.
func (v VZMacPlatformConfiguration) SetFairPlayEnabled(enabled bool) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_setFairPlayEnabled:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setFairPlayEnabled:"}
		return err
	}
	v._setFairPlayEnabled(enabled)
	return nil
}

// CanSetFairPlayEnabled reports whether the receiver responds to the private selector _setFairPlayEnabled:.
func (v VZMacPlatformConfiguration) CanSetFairPlayEnabled() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_setFairPlayEnabled:"))
}
func (v VZMacPlatformConfiguration) _setFakeEncryptionEnabled(enabled bool) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setFakeEncryptionEnabled:"), enabled)
}

// SetFakeEncryptionEnabled is an exported wrapper for the private method _setFakeEncryptionEnabled.
func (v VZMacPlatformConfiguration) SetFakeEncryptionEnabled(enabled bool) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_setFakeEncryptionEnabled:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setFakeEncryptionEnabled:"}
		return err
	}
	v._setFakeEncryptionEnabled(enabled)
	return nil
}

// CanSetFakeEncryptionEnabled reports whether the receiver responds to the private selector _setFakeEncryptionEnabled:.
func (v VZMacPlatformConfiguration) CanSetFakeEncryptionEnabled() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_setFakeEncryptionEnabled:"))
}
func (v VZMacPlatformConfiguration) _setGuestEncryptionWrappingKey(key objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setGuestEncryptionWrappingKey:"), key)
}

// SetGuestEncryptionWrappingKey is an exported wrapper for the private method _setGuestEncryptionWrappingKey.
func (v VZMacPlatformConfiguration) SetGuestEncryptionWrappingKey(key objectivec.IObject) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_setGuestEncryptionWrappingKey:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setGuestEncryptionWrappingKey:"}
		return err
	}
	v._setGuestEncryptionWrappingKey(key)
	return nil
}

// CanSetGuestEncryptionWrappingKey reports whether the receiver responds to the private selector _setGuestEncryptionWrappingKey:.
func (v VZMacPlatformConfiguration) CanSetGuestEncryptionWrappingKey() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_setGuestEncryptionWrappingKey:"))
}
func (v VZMacPlatformConfiguration) _setHostAttributeShareOptions(options uint64) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setHostAttributeShareOptions:"), options)
}

// SetHostAttributeShareOptions is an exported wrapper for the private method _setHostAttributeShareOptions.
func (v VZMacPlatformConfiguration) SetHostAttributeShareOptions(options uint64) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_setHostAttributeShareOptions:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setHostAttributeShareOptions:"}
		return err
	}
	v._setHostAttributeShareOptions(options)
	return nil
}

// CanSetHostAttributeShareOptions reports whether the receiver responds to the private selector _setHostAttributeShareOptions:.
func (v VZMacPlatformConfiguration) CanSetHostAttributeShareOptions() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_setHostAttributeShareOptions:"))
}
func (v VZMacPlatformConfiguration) _setProductionModeEnabled(enabled bool) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setProductionModeEnabled:"), enabled)
}

// SetProductionModeEnabled is an exported wrapper for the private method _setProductionModeEnabled.
func (v VZMacPlatformConfiguration) SetProductionModeEnabled(enabled bool) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_setProductionModeEnabled:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setProductionModeEnabled:"}
		return err
	}
	v._setProductionModeEnabled(enabled)
	return nil
}

// CanSetProductionModeEnabled reports whether the receiver responds to the private selector _setProductionModeEnabled:.
func (v VZMacPlatformConfiguration) CanSetProductionModeEnabled() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_setProductionModeEnabled:"))
}
func (v VZMacPlatformConfiguration) _setRemoteServiceDiscoveryConfiguration(configuration objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setRemoteServiceDiscoveryConfiguration:"), configuration)
}

// SetRemoteServiceDiscoveryConfiguration is an exported wrapper for the private method _setRemoteServiceDiscoveryConfiguration.
func (v VZMacPlatformConfiguration) SetRemoteServiceDiscoveryConfiguration(configuration objectivec.IObject) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_setRemoteServiceDiscoveryConfiguration:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setRemoteServiceDiscoveryConfiguration:"}
		return err
	}
	v._setRemoteServiceDiscoveryConfiguration(configuration)
	return nil
}

// CanSetRemoteServiceDiscoveryConfiguration reports whether the receiver responds to the private selector _setRemoteServiceDiscoveryConfiguration:.
func (v VZMacPlatformConfiguration) CanSetRemoteServiceDiscoveryConfiguration() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_setRemoteServiceDiscoveryConfiguration:"))
}
func (v VZMacPlatformConfiguration) _setSIODescramblerEnabled(enabled bool) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setSIODescramblerEnabled:"), enabled)
}

// SetSIODescramblerEnabled is an exported wrapper for the private method _setSIODescramblerEnabled.
func (v VZMacPlatformConfiguration) SetSIODescramblerEnabled(enabled bool) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_setSIODescramblerEnabled:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setSIODescramblerEnabled:"}
		return err
	}
	v._setSIODescramblerEnabled(enabled)
	return nil
}

// CanSetSIODescramblerEnabled reports whether the receiver responds to the private selector _setSIODescramblerEnabled:.
func (v VZMacPlatformConfiguration) CanSetSIODescramblerEnabled() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_setSIODescramblerEnabled:"))
}
func (v VZMacPlatformConfiguration) _setStrongIdentityEnabled(enabled bool) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setStrongIdentityEnabled:"), enabled)
}

// SetStrongIdentityEnabled is an exported wrapper for the private method _setStrongIdentityEnabled.
func (v VZMacPlatformConfiguration) SetStrongIdentityEnabled(enabled bool) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_setStrongIdentityEnabled:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setStrongIdentityEnabled:"}
		return err
	}
	v._setStrongIdentityEnabled(enabled)
	return nil
}

// CanSetStrongIdentityEnabled reports whether the receiver responds to the private selector _setStrongIdentityEnabled:.
func (v VZMacPlatformConfiguration) CanSetStrongIdentityEnabled() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_setStrongIdentityEnabled:"))
}

func (v VZMacPlatformConfiguration) _fairPlayEnabled() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("_fairPlayEnabled"))
	return rv
}

// CanFairPlayEnabled reports whether the receiver responds to the private selector _fairPlayEnabled.
func (v VZMacPlatformConfiguration) CanFairPlayEnabled() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_fairPlayEnabled"))
}

// FairPlayEnabled is an exported wrapper for the private property _fairPlayEnabled.
func (v VZMacPlatformConfiguration) FairPlayEnabled() (bool, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_fairPlayEnabled")) {
		return false, &objc.UnrecognizedSelectorError{Selector: "_fairPlayEnabled"}
	}
	return v._fairPlayEnabled(), nil
}
func (v VZMacPlatformConfiguration) Set_fairPlayEnabled(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("set_fairPlayEnabled:"), value)
}
func (v VZMacPlatformConfiguration) _fakeEncryptionEnabled() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("_fakeEncryptionEnabled"))
	return rv
}

// CanFakeEncryptionEnabled reports whether the receiver responds to the private selector _fakeEncryptionEnabled.
func (v VZMacPlatformConfiguration) CanFakeEncryptionEnabled() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_fakeEncryptionEnabled"))
}

// FakeEncryptionEnabled is an exported wrapper for the private property _fakeEncryptionEnabled.
func (v VZMacPlatformConfiguration) FakeEncryptionEnabled() (bool, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_fakeEncryptionEnabled")) {
		return false, &objc.UnrecognizedSelectorError{Selector: "_fakeEncryptionEnabled"}
	}
	return v._fakeEncryptionEnabled(), nil
}
func (v VZMacPlatformConfiguration) Set_fakeEncryptionEnabled(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("set_fakeEncryptionEnabled:"), value)
}
func (v VZMacPlatformConfiguration) _guestEncryptionWrappingKey() IVZWrappingKey {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_guestEncryptionWrappingKey"))
	return VZWrappingKeyFromID(objc.ID(rv))
}

// CanGuestEncryptionWrappingKey reports whether the receiver responds to the private selector _guestEncryptionWrappingKey.
func (v VZMacPlatformConfiguration) CanGuestEncryptionWrappingKey() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_guestEncryptionWrappingKey"))
}

// GuestEncryptionWrappingKey is an exported wrapper for the private property _guestEncryptionWrappingKey.
func (v VZMacPlatformConfiguration) GuestEncryptionWrappingKey() (IVZWrappingKey, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_guestEncryptionWrappingKey")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_guestEncryptionWrappingKey"}
	}
	return v._guestEncryptionWrappingKey(), nil
}
func (v VZMacPlatformConfiguration) Set_guestEncryptionWrappingKey(value IVZWrappingKey) {
	objc.Send[struct{}](v.ID, objc.Sel("set_guestEncryptionWrappingKey:"), value)
}
func (v VZMacPlatformConfiguration) _hostAttributeShareOptions() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("_hostAttributeShareOptions"))
	return rv
}

// CanHostAttributeShareOptions reports whether the receiver responds to the private selector _hostAttributeShareOptions.
func (v VZMacPlatformConfiguration) CanHostAttributeShareOptions() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_hostAttributeShareOptions"))
}

// HostAttributeShareOptions is an exported wrapper for the private property _hostAttributeShareOptions.
func (v VZMacPlatformConfiguration) HostAttributeShareOptions() (uint64, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_hostAttributeShareOptions")) {
		return 0, &objc.UnrecognizedSelectorError{Selector: "_hostAttributeShareOptions"}
	}
	return v._hostAttributeShareOptions(), nil
}
func (v VZMacPlatformConfiguration) Set_hostAttributeShareOptions(value uint64) {
	objc.Send[struct{}](v.ID, objc.Sel("set_hostAttributeShareOptions:"), value)
}
func (v VZMacPlatformConfiguration) _productionModeEnabled() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("_productionModeEnabled"))
	return rv
}

// CanProductionModeEnabled reports whether the receiver responds to the private selector _productionModeEnabled.
func (v VZMacPlatformConfiguration) CanProductionModeEnabled() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_productionModeEnabled"))
}

// ProductionModeEnabled is an exported wrapper for the private property _productionModeEnabled.
func (v VZMacPlatformConfiguration) ProductionModeEnabled() (bool, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_productionModeEnabled")) {
		return false, &objc.UnrecognizedSelectorError{Selector: "_productionModeEnabled"}
	}
	return v._productionModeEnabled(), nil
}
func (v VZMacPlatformConfiguration) Set_productionModeEnabled(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("set_productionModeEnabled:"), value)
}
func (v VZMacPlatformConfiguration) _remoteServiceDiscoveryConfiguration() IVZMacRemoteServiceDiscoveryConfiguration {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_remoteServiceDiscoveryConfiguration"))
	return VZMacRemoteServiceDiscoveryConfigurationFromID(objc.ID(rv))
}

// CanRemoteServiceDiscoveryConfiguration reports whether the receiver responds to the private selector _remoteServiceDiscoveryConfiguration.
func (v VZMacPlatformConfiguration) CanRemoteServiceDiscoveryConfiguration() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_remoteServiceDiscoveryConfiguration"))
}

// RemoteServiceDiscoveryConfiguration is an exported wrapper for the private property _remoteServiceDiscoveryConfiguration.
func (v VZMacPlatformConfiguration) RemoteServiceDiscoveryConfiguration() (IVZMacRemoteServiceDiscoveryConfiguration, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_remoteServiceDiscoveryConfiguration")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_remoteServiceDiscoveryConfiguration"}
	}
	return v._remoteServiceDiscoveryConfiguration(), nil
}
func (v VZMacPlatformConfiguration) Set_remoteServiceDiscoveryConfiguration(value IVZMacRemoteServiceDiscoveryConfiguration) {
	objc.Send[struct{}](v.ID, objc.Sel("set_remoteServiceDiscoveryConfiguration:"), value)
}
func (v VZMacPlatformConfiguration) _sioDescramblerEnabled() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("_sioDescramblerEnabled"))
	return rv
}

// CanSioDescramblerEnabled reports whether the receiver responds to the private selector _sioDescramblerEnabled.
func (v VZMacPlatformConfiguration) CanSioDescramblerEnabled() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_sioDescramblerEnabled"))
}

// SioDescramblerEnabled is an exported wrapper for the private property _sioDescramblerEnabled.
func (v VZMacPlatformConfiguration) SioDescramblerEnabled() (bool, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_sioDescramblerEnabled")) {
		return false, &objc.UnrecognizedSelectorError{Selector: "_sioDescramblerEnabled"}
	}
	return v._sioDescramblerEnabled(), nil
}
func (v VZMacPlatformConfiguration) Set_sioDescramblerEnabled(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("set_sioDescramblerEnabled:"), value)
}
func (v VZMacPlatformConfiguration) _strongIdentityEnabled() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("_strongIdentityEnabled"))
	return rv
}

// CanStrongIdentityEnabled reports whether the receiver responds to the private selector _strongIdentityEnabled.
func (v VZMacPlatformConfiguration) CanStrongIdentityEnabled() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_strongIdentityEnabled"))
}

// StrongIdentityEnabled is an exported wrapper for the private property _strongIdentityEnabled.
func (v VZMacPlatformConfiguration) StrongIdentityEnabled() (bool, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_strongIdentityEnabled")) {
		return false, &objc.UnrecognizedSelectorError{Selector: "_strongIdentityEnabled"}
	}
	return v._strongIdentityEnabled(), nil
}
func (v VZMacPlatformConfiguration) Set_strongIdentityEnabled(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("set_strongIdentityEnabled:"), value)
}
