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
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacPlatformConfiguration
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
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacPlatformConfiguration
type IVZMacPlatformConfiguration interface {
	IVZPlatformConfiguration

	// Topic: Methods

	_guestEncryptionWrappingKey() *VZWrappingKey
	Set_guestEncryptionWrappingKey(value *VZWrappingKey)
	_hostAttributeShareOptions() uint64
	Set_hostAttributeShareOptions(value uint64)
	_isFairPlayEnabled() bool
	_isFakeEncryptionEnabled() bool
	_isProductionModeEnabled() bool
	_isSIODescramblerEnabled() bool
	_isStrongIdentityEnabled() bool
	_remoteServiceDiscoveryConfiguration() *VZMacRemoteServiceDiscoveryConfiguration
	Set_remoteServiceDiscoveryConfiguration(value *VZMacRemoteServiceDiscoveryConfiguration)
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
func (m VZMacPlatformConfiguration) Init() VZMacPlatformConfiguration {
	rv := objc.Send[VZMacPlatformConfiguration](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m VZMacPlatformConfiguration) Autorelease() VZMacPlatformConfiguration {
	rv := objc.Send[VZMacPlatformConfiguration](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacPlatformConfiguration creates a new VZMacPlatformConfiguration instance.
func NewVZMacPlatformConfiguration() VZMacPlatformConfiguration {
	class := getVZMacPlatformConfigurationClass()
	rv := objc.Send[VZMacPlatformConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacPlatformConfiguration/_isFairPlayEnabled
func (m VZMacPlatformConfiguration) _isFairPlayEnabled() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("_isFairPlayEnabled"))
	return rv
}

// IsFairPlayEnabled is an exported wrapper for the private method _isFairPlayEnabled.
func (m VZMacPlatformConfiguration) IsFairPlayEnabled() (bool, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_isFairPlayEnabled")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_isFairPlayEnabled"}
		return false, err
	}
	return m._isFairPlayEnabled(), nil
}

// CanIsFairPlayEnabled reports whether the receiver responds to the private selector _isFairPlayEnabled.
func (m VZMacPlatformConfiguration) CanIsFairPlayEnabled() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_isFairPlayEnabled"))
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacPlatformConfiguration/_isFakeEncryptionEnabled
func (m VZMacPlatformConfiguration) _isFakeEncryptionEnabled() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("_isFakeEncryptionEnabled"))
	return rv
}

// IsFakeEncryptionEnabled is an exported wrapper for the private method _isFakeEncryptionEnabled.
func (m VZMacPlatformConfiguration) IsFakeEncryptionEnabled() (bool, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_isFakeEncryptionEnabled")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_isFakeEncryptionEnabled"}
		return false, err
	}
	return m._isFakeEncryptionEnabled(), nil
}

// CanIsFakeEncryptionEnabled reports whether the receiver responds to the private selector _isFakeEncryptionEnabled.
func (m VZMacPlatformConfiguration) CanIsFakeEncryptionEnabled() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_isFakeEncryptionEnabled"))
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacPlatformConfiguration/_isProductionModeEnabled
func (m VZMacPlatformConfiguration) _isProductionModeEnabled() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("_isProductionModeEnabled"))
	return rv
}

// IsProductionModeEnabled is an exported wrapper for the private method _isProductionModeEnabled.
func (m VZMacPlatformConfiguration) IsProductionModeEnabled() (bool, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_isProductionModeEnabled")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_isProductionModeEnabled"}
		return false, err
	}
	return m._isProductionModeEnabled(), nil
}

// CanIsProductionModeEnabled reports whether the receiver responds to the private selector _isProductionModeEnabled.
func (m VZMacPlatformConfiguration) CanIsProductionModeEnabled() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_isProductionModeEnabled"))
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacPlatformConfiguration/_isSIODescramblerEnabled
func (m VZMacPlatformConfiguration) _isSIODescramblerEnabled() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("_isSIODescramblerEnabled"))
	return rv
}

// IsSIODescramblerEnabled is an exported wrapper for the private method _isSIODescramblerEnabled.
func (m VZMacPlatformConfiguration) IsSIODescramblerEnabled() (bool, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_isSIODescramblerEnabled")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_isSIODescramblerEnabled"}
		return false, err
	}
	return m._isSIODescramblerEnabled(), nil
}

// CanIsSIODescramblerEnabled reports whether the receiver responds to the private selector _isSIODescramblerEnabled.
func (m VZMacPlatformConfiguration) CanIsSIODescramblerEnabled() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_isSIODescramblerEnabled"))
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacPlatformConfiguration/_isStrongIdentityEnabled
func (m VZMacPlatformConfiguration) _isStrongIdentityEnabled() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("_isStrongIdentityEnabled"))
	return rv
}

// IsStrongIdentityEnabled is an exported wrapper for the private method _isStrongIdentityEnabled.
func (m VZMacPlatformConfiguration) IsStrongIdentityEnabled() (bool, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_isStrongIdentityEnabled")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_isStrongIdentityEnabled"}
		return false, err
	}
	return m._isStrongIdentityEnabled(), nil
}

// CanIsStrongIdentityEnabled reports whether the receiver responds to the private selector _isStrongIdentityEnabled.
func (m VZMacPlatformConfiguration) CanIsStrongIdentityEnabled() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_isStrongIdentityEnabled"))
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacPlatformConfiguration/_setFairPlayEnabled:
func (m VZMacPlatformConfiguration) _setFairPlayEnabled(enabled bool) {
	objc.Send[objc.ID](m.ID, objc.Sel("_setFairPlayEnabled:"), enabled)
}

// SetFairPlayEnabled is an exported wrapper for the private method _setFairPlayEnabled.
func (m VZMacPlatformConfiguration) SetFairPlayEnabled(enabled bool) error {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_setFairPlayEnabled:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setFairPlayEnabled:"}
		return err
	}
	m._setFairPlayEnabled(enabled)
	return nil
}

// CanSetFairPlayEnabled reports whether the receiver responds to the private selector _setFairPlayEnabled:.
func (m VZMacPlatformConfiguration) CanSetFairPlayEnabled() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_setFairPlayEnabled:"))
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacPlatformConfiguration/_setFakeEncryptionEnabled:
func (m VZMacPlatformConfiguration) _setFakeEncryptionEnabled(enabled bool) {
	objc.Send[objc.ID](m.ID, objc.Sel("_setFakeEncryptionEnabled:"), enabled)
}

// SetFakeEncryptionEnabled is an exported wrapper for the private method _setFakeEncryptionEnabled.
func (m VZMacPlatformConfiguration) SetFakeEncryptionEnabled(enabled bool) error {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_setFakeEncryptionEnabled:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setFakeEncryptionEnabled:"}
		return err
	}
	m._setFakeEncryptionEnabled(enabled)
	return nil
}

// CanSetFakeEncryptionEnabled reports whether the receiver responds to the private selector _setFakeEncryptionEnabled:.
func (m VZMacPlatformConfiguration) CanSetFakeEncryptionEnabled() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_setFakeEncryptionEnabled:"))
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacPlatformConfiguration/_setGuestEncryptionWrappingKey:
func (m VZMacPlatformConfiguration) _setGuestEncryptionWrappingKey(key objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("_setGuestEncryptionWrappingKey:"), key)
}

// SetGuestEncryptionWrappingKey is an exported wrapper for the private method _setGuestEncryptionWrappingKey.
func (m VZMacPlatformConfiguration) SetGuestEncryptionWrappingKey(key objectivec.IObject) error {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_setGuestEncryptionWrappingKey:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setGuestEncryptionWrappingKey:"}
		return err
	}
	m._setGuestEncryptionWrappingKey(key)
	return nil
}

// CanSetGuestEncryptionWrappingKey reports whether the receiver responds to the private selector _setGuestEncryptionWrappingKey:.
func (m VZMacPlatformConfiguration) CanSetGuestEncryptionWrappingKey() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_setGuestEncryptionWrappingKey:"))
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacPlatformConfiguration/_setHostAttributeShareOptions:
func (m VZMacPlatformConfiguration) _setHostAttributeShareOptions(options uint64) {
	objc.Send[objc.ID](m.ID, objc.Sel("_setHostAttributeShareOptions:"), options)
}

// SetHostAttributeShareOptions is an exported wrapper for the private method _setHostAttributeShareOptions.
func (m VZMacPlatformConfiguration) SetHostAttributeShareOptions(options uint64) error {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_setHostAttributeShareOptions:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setHostAttributeShareOptions:"}
		return err
	}
	m._setHostAttributeShareOptions(options)
	return nil
}

// CanSetHostAttributeShareOptions reports whether the receiver responds to the private selector _setHostAttributeShareOptions:.
func (m VZMacPlatformConfiguration) CanSetHostAttributeShareOptions() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_setHostAttributeShareOptions:"))
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacPlatformConfiguration/_setProductionModeEnabled:
func (m VZMacPlatformConfiguration) _setProductionModeEnabled(enabled bool) {
	objc.Send[objc.ID](m.ID, objc.Sel("_setProductionModeEnabled:"), enabled)
}

// SetProductionModeEnabled is an exported wrapper for the private method _setProductionModeEnabled.
func (m VZMacPlatformConfiguration) SetProductionModeEnabled(enabled bool) error {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_setProductionModeEnabled:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setProductionModeEnabled:"}
		return err
	}
	m._setProductionModeEnabled(enabled)
	return nil
}

// CanSetProductionModeEnabled reports whether the receiver responds to the private selector _setProductionModeEnabled:.
func (m VZMacPlatformConfiguration) CanSetProductionModeEnabled() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_setProductionModeEnabled:"))
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacPlatformConfiguration/_setRemoteServiceDiscoveryConfiguration:
func (m VZMacPlatformConfiguration) _setRemoteServiceDiscoveryConfiguration(configuration objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("_setRemoteServiceDiscoveryConfiguration:"), configuration)
}

// SetRemoteServiceDiscoveryConfiguration is an exported wrapper for the private method _setRemoteServiceDiscoveryConfiguration.
func (m VZMacPlatformConfiguration) SetRemoteServiceDiscoveryConfiguration(configuration objectivec.IObject) error {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_setRemoteServiceDiscoveryConfiguration:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setRemoteServiceDiscoveryConfiguration:"}
		return err
	}
	m._setRemoteServiceDiscoveryConfiguration(configuration)
	return nil
}

// CanSetRemoteServiceDiscoveryConfiguration reports whether the receiver responds to the private selector _setRemoteServiceDiscoveryConfiguration:.
func (m VZMacPlatformConfiguration) CanSetRemoteServiceDiscoveryConfiguration() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_setRemoteServiceDiscoveryConfiguration:"))
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacPlatformConfiguration/_setSIODescramblerEnabled:
func (m VZMacPlatformConfiguration) _setSIODescramblerEnabled(enabled bool) {
	objc.Send[objc.ID](m.ID, objc.Sel("_setSIODescramblerEnabled:"), enabled)
}

// SetSIODescramblerEnabled is an exported wrapper for the private method _setSIODescramblerEnabled.
func (m VZMacPlatformConfiguration) SetSIODescramblerEnabled(enabled bool) error {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_setSIODescramblerEnabled:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setSIODescramblerEnabled:"}
		return err
	}
	m._setSIODescramblerEnabled(enabled)
	return nil
}

// CanSetSIODescramblerEnabled reports whether the receiver responds to the private selector _setSIODescramblerEnabled:.
func (m VZMacPlatformConfiguration) CanSetSIODescramblerEnabled() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_setSIODescramblerEnabled:"))
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacPlatformConfiguration/_setStrongIdentityEnabled:
func (m VZMacPlatformConfiguration) _setStrongIdentityEnabled(enabled bool) {
	objc.Send[objc.ID](m.ID, objc.Sel("_setStrongIdentityEnabled:"), enabled)
}

// SetStrongIdentityEnabled is an exported wrapper for the private method _setStrongIdentityEnabled.
func (m VZMacPlatformConfiguration) SetStrongIdentityEnabled(enabled bool) error {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_setStrongIdentityEnabled:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setStrongIdentityEnabled:"}
		return err
	}
	m._setStrongIdentityEnabled(enabled)
	return nil
}

// CanSetStrongIdentityEnabled reports whether the receiver responds to the private selector _setStrongIdentityEnabled:.
func (m VZMacPlatformConfiguration) CanSetStrongIdentityEnabled() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_setStrongIdentityEnabled:"))
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacPlatformConfiguration/_fairPlayEnabled
func (m VZMacPlatformConfiguration) _fairPlayEnabled() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("_fairPlayEnabled"))
	return rv
}

// CanFairPlayEnabled reports whether the receiver responds to the private selector _fairPlayEnabled.
func (m VZMacPlatformConfiguration) CanFairPlayEnabled() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_fairPlayEnabled"))
}

// FairPlayEnabled is an exported wrapper for the private property _fairPlayEnabled.
func (m VZMacPlatformConfiguration) FairPlayEnabled() (bool, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_fairPlayEnabled")) {
		return false, &objc.UnrecognizedSelectorError{Selector: "_fairPlayEnabled"}
	}
	return m._fairPlayEnabled(), nil
}
func (m VZMacPlatformConfiguration) Set_fairPlayEnabled(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("set_fairPlayEnabled:"), value)
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacPlatformConfiguration/_fakeEncryptionEnabled
func (m VZMacPlatformConfiguration) _fakeEncryptionEnabled() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("_fakeEncryptionEnabled"))
	return rv
}

// CanFakeEncryptionEnabled reports whether the receiver responds to the private selector _fakeEncryptionEnabled.
func (m VZMacPlatformConfiguration) CanFakeEncryptionEnabled() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_fakeEncryptionEnabled"))
}

// FakeEncryptionEnabled is an exported wrapper for the private property _fakeEncryptionEnabled.
func (m VZMacPlatformConfiguration) FakeEncryptionEnabled() (bool, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_fakeEncryptionEnabled")) {
		return false, &objc.UnrecognizedSelectorError{Selector: "_fakeEncryptionEnabled"}
	}
	return m._fakeEncryptionEnabled(), nil
}
func (m VZMacPlatformConfiguration) Set_fakeEncryptionEnabled(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("set_fakeEncryptionEnabled:"), value)
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacPlatformConfiguration/_guestEncryptionWrappingKey
func (m VZMacPlatformConfiguration) _guestEncryptionWrappingKey() *VZWrappingKey {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("_guestEncryptionWrappingKey"))
	if rv == 0 {
		return nil
	}
	val := VZWrappingKeyFromID(objc.ID(rv))
	return &val
}

// CanGuestEncryptionWrappingKey reports whether the receiver responds to the private selector _guestEncryptionWrappingKey.
func (m VZMacPlatformConfiguration) CanGuestEncryptionWrappingKey() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_guestEncryptionWrappingKey"))
}

// GuestEncryptionWrappingKey is an exported wrapper for the private property _guestEncryptionWrappingKey.
func (m VZMacPlatformConfiguration) GuestEncryptionWrappingKey() (*VZWrappingKey, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_guestEncryptionWrappingKey")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_guestEncryptionWrappingKey"}
	}
	return m._guestEncryptionWrappingKey(), nil
}
func (m VZMacPlatformConfiguration) Set_guestEncryptionWrappingKey(value *VZWrappingKey) {
	if value == nil {
		objc.Send[struct{}](m.ID, objc.Sel("set_guestEncryptionWrappingKey:"), objc.ID(0))
		return
	}
	objc.Send[struct{}](m.ID, objc.Sel("set_guestEncryptionWrappingKey:"), value)
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacPlatformConfiguration/_hostAttributeShareOptions
func (m VZMacPlatformConfiguration) _hostAttributeShareOptions() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("_hostAttributeShareOptions"))
	return rv
}

// CanHostAttributeShareOptions reports whether the receiver responds to the private selector _hostAttributeShareOptions.
func (m VZMacPlatformConfiguration) CanHostAttributeShareOptions() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_hostAttributeShareOptions"))
}

// HostAttributeShareOptions is an exported wrapper for the private property _hostAttributeShareOptions.
func (m VZMacPlatformConfiguration) HostAttributeShareOptions() (uint64, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_hostAttributeShareOptions")) {
		return 0, &objc.UnrecognizedSelectorError{Selector: "_hostAttributeShareOptions"}
	}
	return m._hostAttributeShareOptions(), nil
}
func (m VZMacPlatformConfiguration) Set_hostAttributeShareOptions(value uint64) {
	objc.Send[struct{}](m.ID, objc.Sel("set_hostAttributeShareOptions:"), value)
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacPlatformConfiguration/_productionModeEnabled
func (m VZMacPlatformConfiguration) _productionModeEnabled() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("_productionModeEnabled"))
	return rv
}

// CanProductionModeEnabled reports whether the receiver responds to the private selector _productionModeEnabled.
func (m VZMacPlatformConfiguration) CanProductionModeEnabled() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_productionModeEnabled"))
}

// ProductionModeEnabled is an exported wrapper for the private property _productionModeEnabled.
func (m VZMacPlatformConfiguration) ProductionModeEnabled() (bool, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_productionModeEnabled")) {
		return false, &objc.UnrecognizedSelectorError{Selector: "_productionModeEnabled"}
	}
	return m._productionModeEnabled(), nil
}
func (m VZMacPlatformConfiguration) Set_productionModeEnabled(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("set_productionModeEnabled:"), value)
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacPlatformConfiguration/_remoteServiceDiscoveryConfiguration
func (m VZMacPlatformConfiguration) _remoteServiceDiscoveryConfiguration() *VZMacRemoteServiceDiscoveryConfiguration {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("_remoteServiceDiscoveryConfiguration"))
	if rv == 0 {
		return nil
	}
	val := VZMacRemoteServiceDiscoveryConfigurationFromID(objc.ID(rv))
	return &val
}

// CanRemoteServiceDiscoveryConfiguration reports whether the receiver responds to the private selector _remoteServiceDiscoveryConfiguration.
func (m VZMacPlatformConfiguration) CanRemoteServiceDiscoveryConfiguration() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_remoteServiceDiscoveryConfiguration"))
}

// RemoteServiceDiscoveryConfiguration is an exported wrapper for the private property _remoteServiceDiscoveryConfiguration.
func (m VZMacPlatformConfiguration) RemoteServiceDiscoveryConfiguration() (*VZMacRemoteServiceDiscoveryConfiguration, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_remoteServiceDiscoveryConfiguration")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_remoteServiceDiscoveryConfiguration"}
	}
	return m._remoteServiceDiscoveryConfiguration(), nil
}
func (m VZMacPlatformConfiguration) Set_remoteServiceDiscoveryConfiguration(value *VZMacRemoteServiceDiscoveryConfiguration) {
	if value == nil {
		objc.Send[struct{}](m.ID, objc.Sel("set_remoteServiceDiscoveryConfiguration:"), objc.ID(0))
		return
	}
	objc.Send[struct{}](m.ID, objc.Sel("set_remoteServiceDiscoveryConfiguration:"), value)
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacPlatformConfiguration/_sioDescramblerEnabled
func (m VZMacPlatformConfiguration) _sioDescramblerEnabled() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("_sioDescramblerEnabled"))
	return rv
}

// CanSioDescramblerEnabled reports whether the receiver responds to the private selector _sioDescramblerEnabled.
func (m VZMacPlatformConfiguration) CanSioDescramblerEnabled() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_sioDescramblerEnabled"))
}

// SioDescramblerEnabled is an exported wrapper for the private property _sioDescramblerEnabled.
func (m VZMacPlatformConfiguration) SioDescramblerEnabled() (bool, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_sioDescramblerEnabled")) {
		return false, &objc.UnrecognizedSelectorError{Selector: "_sioDescramblerEnabled"}
	}
	return m._sioDescramblerEnabled(), nil
}
func (m VZMacPlatformConfiguration) Set_sioDescramblerEnabled(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("set_sioDescramblerEnabled:"), value)
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacPlatformConfiguration/_strongIdentityEnabled
func (m VZMacPlatformConfiguration) _strongIdentityEnabled() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("_strongIdentityEnabled"))
	return rv
}

// CanStrongIdentityEnabled reports whether the receiver responds to the private selector _strongIdentityEnabled.
func (m VZMacPlatformConfiguration) CanStrongIdentityEnabled() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_strongIdentityEnabled"))
}

// StrongIdentityEnabled is an exported wrapper for the private property _strongIdentityEnabled.
func (m VZMacPlatformConfiguration) StrongIdentityEnabled() (bool, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_strongIdentityEnabled")) {
		return false, &objc.UnrecognizedSelectorError{Selector: "_strongIdentityEnabled"}
	}
	return m._strongIdentityEnabled(), nil
}
func (m VZMacPlatformConfiguration) Set_strongIdentityEnabled(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("set_strongIdentityEnabled:"), value)
}
