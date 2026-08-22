// Code generated from Apple documentation for ColorSync. DO NOT EDIT.

package colorsync

import (
	"fmt"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
)

type unavailableSymbolError struct {
	symbol     string
	introduced string
	cause      error
}

func (e *unavailableSymbolError) Error() string {
	if e == nil {
		return ""
	}
	if e.introduced != "" {
		return fmt.Sprintf("ColorSync: symbol %s unavailable on this system (introduced in macOS %s)", e.symbol, e.introduced)
	}
	return fmt.Sprintf("ColorSync: symbol %s unavailable on this system", e.symbol)
}

func (e *unavailableSymbolError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.cause
}

func missingSymbolError(name, introduced string, cause error) error {
	return &unavailableSymbolError{
		symbol:     name,
		introduced: introduced,
		cause:      cause,
	}
}

func symbolCallError(name, introduced string, err error) error {
	if err != nil {
		return err
	}
	if frameworkHandle == 0 {
		return fmt.Errorf("ColorSync: symbol %s unavailable because the framework could not be loaded", name)
	}
	return missingSymbolError(name, introduced, nil)
}

// registerFunc resolves a framework symbol and registers it as a Go function.
func registerFunc(fptr any, errDst *error, handle uintptr, name, introduced string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil || sym == 0 {
		*errDst = missingSymbolError(name, introduced, err)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			*errDst = fmt.Errorf("ColorSync: register symbol %s: %v", name, r)
		}
	}()
	purego.RegisterFunc(fptr, sym)
	*errDst = nil
}

// registerSymbol resolves a framework symbol and stores its raw address.
func registerSymbol(dst *uintptr, errDst *error, handle uintptr, name, introduced string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil || sym == 0 {
		*errDst = missingSymbolError(name, introduced, err)
		return
	}
	*dst = sym
	*errDst = nil
}

var _cGDisplayCreateUUIDFromDisplayID func(displayID uint32) corefoundation.CFUUID
var _cGDisplayCreateUUIDFromDisplayIDErr error

func tryCGDisplayCreateUUIDFromDisplayID(displayID uint32) (corefoundation.CFUUID, error) {
	if _cGDisplayCreateUUIDFromDisplayID == nil {
		return *new(corefoundation.CFUUID), symbolCallError("CGDisplayCreateUUIDFromDisplayID", "10.13", _cGDisplayCreateUUIDFromDisplayIDErr)
	}
	return _cGDisplayCreateUUIDFromDisplayID(displayID), nil
}

// CGDisplayCreateUUIDFromDisplayID.
//
// See: https://developer.apple.com/documentation/ColorSync/CGDisplayCreateUUIDFromDisplayID(_:)
func CGDisplayCreateUUIDFromDisplayID(displayID uint32) corefoundation.CFUUID {
	result, callErr := tryCGDisplayCreateUUIDFromDisplayID(displayID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplayGetDisplayIDFromUUID func(uuid corefoundation.CFUUIDRef) uint32
var _cGDisplayGetDisplayIDFromUUIDErr error

func tryCGDisplayGetDisplayIDFromUUID(uuid corefoundation.CFUUIDRef) (uint32, error) {
	if _cGDisplayGetDisplayIDFromUUID == nil {
		return 0, symbolCallError("CGDisplayGetDisplayIDFromUUID", "10.13", _cGDisplayGetDisplayIDFromUUIDErr)
	}
	return _cGDisplayGetDisplayIDFromUUID(uuid), nil
}

// CGDisplayGetDisplayIDFromUUID.
//
// See: https://developer.apple.com/documentation/ColorSync/CGDisplayGetDisplayIDFromUUID(_:)
func CGDisplayGetDisplayIDFromUUID(uuid corefoundation.CFUUIDRef) uint32 {
	result, callErr := tryCGDisplayGetDisplayIDFromUUID(uuid)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _colorSyncAPIVersion func() uint32
var _colorSyncAPIVersionErr error

func tryColorSyncAPIVersion() (uint32, error) {
	if _colorSyncAPIVersion == nil {
		return 0, symbolCallError("ColorSyncAPIVersion", "10.13", _colorSyncAPIVersionErr)
	}
	return _colorSyncAPIVersion(), nil
}

// ColorSyncAPIVersion returns the version of the ColorSync API.
//
// See: https://developer.apple.com/documentation/ColorSync/ColorSyncAPIVersion()
func ColorSyncAPIVersion() uint32 {
	result, callErr := tryColorSyncAPIVersion()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _colorSyncCMMCopyCMMIdentifier func(arg0 unsafe.Pointer) corefoundation.CFStringRef
var _colorSyncCMMCopyCMMIdentifierErr error

func tryColorSyncCMMCopyCMMIdentifier(arg0 unsafe.Pointer) (corefoundation.CFStringRef, error) {
	if _colorSyncCMMCopyCMMIdentifier == nil {
		return *new(corefoundation.CFStringRef), symbolCallError("ColorSyncCMMCopyCMMIdentifier", "10.13", _colorSyncCMMCopyCMMIdentifierErr)
	}
	return _colorSyncCMMCopyCMMIdentifier(arg0), nil
}

// ColorSyncCMMCopyCMMIdentifier copies the identifier of a CMM.
//
// See: https://developer.apple.com/documentation/ColorSync/ColorSyncCMMCopyCMMIdentifier(_:)
func ColorSyncCMMCopyCMMIdentifier(arg0 unsafe.Pointer) corefoundation.CFStringRef {
	result, callErr := tryColorSyncCMMCopyCMMIdentifier(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _colorSyncCMMCopyLocalizedName func(arg0 unsafe.Pointer) corefoundation.CFStringRef
var _colorSyncCMMCopyLocalizedNameErr error

func tryColorSyncCMMCopyLocalizedName(arg0 unsafe.Pointer) (corefoundation.CFStringRef, error) {
	if _colorSyncCMMCopyLocalizedName == nil {
		return *new(corefoundation.CFStringRef), symbolCallError("ColorSyncCMMCopyLocalizedName", "10.13", _colorSyncCMMCopyLocalizedNameErr)
	}
	return _colorSyncCMMCopyLocalizedName(arg0), nil
}

// ColorSyncCMMCopyLocalizedName copies the localized name of a CMM.
//
// See: https://developer.apple.com/documentation/ColorSync/ColorSyncCMMCopyLocalizedName(_:)
func ColorSyncCMMCopyLocalizedName(arg0 unsafe.Pointer) corefoundation.CFStringRef {
	result, callErr := tryColorSyncCMMCopyLocalizedName(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _colorSyncCMMCreate func(cmmBundle corefoundation.CFBundleRef) unsafe.Pointer
var _colorSyncCMMCreateErr error

func tryColorSyncCMMCreate(cmmBundle corefoundation.CFBundleRef) (unsafe.Pointer, error) {
	if _colorSyncCMMCreate == nil {
		return nil, symbolCallError("ColorSyncCMMCreate", "10.13", _colorSyncCMMCreateErr)
	}
	return _colorSyncCMMCreate(cmmBundle), nil
}

// ColorSyncCMMCreate creates a CMM object from a CMM bundle.
//
// See: https://developer.apple.com/documentation/ColorSync/ColorSyncCMMCreate(_:)
func ColorSyncCMMCreate(cmmBundle corefoundation.CFBundleRef) unsafe.Pointer {
	result, callErr := tryColorSyncCMMCreate(cmmBundle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _colorSyncCMMGetBundle func(arg0 unsafe.Pointer) corefoundation.CFBundle
var _colorSyncCMMGetBundleErr error

func tryColorSyncCMMGetBundle(arg0 unsafe.Pointer) (corefoundation.CFBundle, error) {
	if _colorSyncCMMGetBundle == nil {
		return *new(corefoundation.CFBundle), symbolCallError("ColorSyncCMMGetBundle", "10.13", _colorSyncCMMGetBundleErr)
	}
	return _colorSyncCMMGetBundle(arg0), nil
}

// ColorSyncCMMGetBundle returns the bundle associated with a CMM.
//
// See: https://developer.apple.com/documentation/ColorSync/ColorSyncCMMGetBundle(_:)
func ColorSyncCMMGetBundle(arg0 unsafe.Pointer) corefoundation.CFBundle {
	result, callErr := tryColorSyncCMMGetBundle(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _colorSyncCMMGetTypeID func() uint
var _colorSyncCMMGetTypeIDErr error

func tryColorSyncCMMGetTypeID() (uint, error) {
	if _colorSyncCMMGetTypeID == nil {
		return 0, symbolCallError("ColorSyncCMMGetTypeID", "10.13", _colorSyncCMMGetTypeIDErr)
	}
	return _colorSyncCMMGetTypeID(), nil
}

// ColorSyncCMMGetTypeID returns the [CFTypeID] for [ColorSyncCMM]s.
//
// See: https://developer.apple.com/documentation/ColorSync/ColorSyncCMMGetTypeID()
func ColorSyncCMMGetTypeID() uint {
	result, callErr := tryColorSyncCMMGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _colorSyncCreateCodeFragment func(profileSequence corefoundation.CFArrayRef, options corefoundation.CFDictionaryRef) corefoundation.CFTypeRef
var _colorSyncCreateCodeFragmentErr error

func tryColorSyncCreateCodeFragment(profileSequence corefoundation.CFArrayRef, options corefoundation.CFDictionaryRef) (corefoundation.CFTypeRef, error) {
	if _colorSyncCreateCodeFragment == nil {
		return *new(corefoundation.CFTypeRef), symbolCallError("ColorSyncCreateCodeFragment", "10.13", _colorSyncCreateCodeFragmentErr)
	}
	return _colorSyncCreateCodeFragment(profileSequence, options), nil
}

// ColorSyncCreateCodeFragment creates a code fragment from a sequence of profiles.
//
// See: https://developer.apple.com/documentation/ColorSync/ColorSyncCreateCodeFragment(_:_:)
func ColorSyncCreateCodeFragment(profileSequence corefoundation.CFArrayRef, options corefoundation.CFDictionaryRef) corefoundation.CFTypeRef {
	result, callErr := tryColorSyncCreateCodeFragment(profileSequence, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _colorSyncDeviceCopyDeviceInfo func(deviceClass corefoundation.CFStringRef, devID corefoundation.CFUUIDRef) corefoundation.CFDictionaryRef
var _colorSyncDeviceCopyDeviceInfoErr error

func tryColorSyncDeviceCopyDeviceInfo(deviceClass corefoundation.CFStringRef, devID corefoundation.CFUUIDRef) (corefoundation.CFDictionaryRef, error) {
	if _colorSyncDeviceCopyDeviceInfo == nil {
		return *new(corefoundation.CFDictionaryRef), symbolCallError("ColorSyncDeviceCopyDeviceInfo", "10.13", _colorSyncDeviceCopyDeviceInfoErr)
	}
	return _colorSyncDeviceCopyDeviceInfo(deviceClass, devID), nil
}

// ColorSyncDeviceCopyDeviceInfo copies information about a device, resolved for the current host and current user.
//
// See: https://developer.apple.com/documentation/ColorSync/ColorSyncDeviceCopyDeviceInfo(_:_:)
func ColorSyncDeviceCopyDeviceInfo(deviceClass corefoundation.CFStringRef, devID corefoundation.CFUUIDRef) corefoundation.CFDictionaryRef {
	result, callErr := tryColorSyncDeviceCopyDeviceInfo(deviceClass, devID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _colorSyncDeviceSetCustomProfiles func(deviceClass corefoundation.CFStringRef, deviceID corefoundation.CFUUIDRef, profileInfo corefoundation.CFDictionaryRef) bool
var _colorSyncDeviceSetCustomProfilesErr error

func tryColorSyncDeviceSetCustomProfiles(deviceClass corefoundation.CFStringRef, deviceID corefoundation.CFUUIDRef, profileInfo corefoundation.CFDictionaryRef) (bool, error) {
	if _colorSyncDeviceSetCustomProfiles == nil {
		return false, symbolCallError("ColorSyncDeviceSetCustomProfiles", "10.13", _colorSyncDeviceSetCustomProfilesErr)
	}
	return _colorSyncDeviceSetCustomProfiles(deviceClass, deviceID, profileInfo), nil
}

// ColorSyncDeviceSetCustomProfiles sets custom profiles for a device in lieu of its factory profiles.
//
// See: https://developer.apple.com/documentation/ColorSync/ColorSyncDeviceSetCustomProfiles(_:_:_:)
func ColorSyncDeviceSetCustomProfiles(deviceClass corefoundation.CFStringRef, deviceID corefoundation.CFUUIDRef, profileInfo corefoundation.CFDictionaryRef) bool {
	result, callErr := tryColorSyncDeviceSetCustomProfiles(deviceClass, deviceID, profileInfo)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _colorSyncIterateDeviceProfiles func(callBack ColorSyncDeviceProfileIterateCallback, userInfo unsafe.Pointer)
var _colorSyncIterateDeviceProfilesErr error

func tryColorSyncIterateDeviceProfiles(callBack ColorSyncDeviceProfileIterateCallback, userInfo unsafe.Pointer) error {
	if _colorSyncIterateDeviceProfiles == nil {
		return symbolCallError("ColorSyncIterateDeviceProfiles", "10.13", _colorSyncIterateDeviceProfilesErr)
	}
	_colorSyncIterateDeviceProfiles(callBack, userInfo)
	return nil
}

// ColorSyncIterateDeviceProfiles iterates over the profiles registered for all devices, invoking a callback for each.
//
// See: https://developer.apple.com/documentation/ColorSync/ColorSyncIterateDeviceProfiles(_:_:)
func ColorSyncIterateDeviceProfiles(callBack ColorSyncDeviceProfileIterateCallback, userInfo unsafe.Pointer) {
	if callErr := tryColorSyncIterateDeviceProfiles(callBack, userInfo); callErr != nil {
		panic(callErr)
	}
}

var _colorSyncIterateInstalledCMMs func(callBack ColorSyncCMMIterateCallback, userInfo unsafe.Pointer)
var _colorSyncIterateInstalledCMMsErr error

func tryColorSyncIterateInstalledCMMs(callBack ColorSyncCMMIterateCallback, userInfo unsafe.Pointer) error {
	if _colorSyncIterateInstalledCMMs == nil {
		return symbolCallError("ColorSyncIterateInstalledCMMs", "10.13", _colorSyncIterateInstalledCMMsErr)
	}
	_colorSyncIterateInstalledCMMs(callBack, userInfo)
	return nil
}

// ColorSyncIterateInstalledCMMs iterates over the installed CMMs, invoking a callback for each one.
//
// See: https://developer.apple.com/documentation/ColorSync/ColorSyncIterateInstalledCMMs(_:_:)
func ColorSyncIterateInstalledCMMs(callBack ColorSyncCMMIterateCallback, userInfo unsafe.Pointer) {
	if callErr := tryColorSyncIterateInstalledCMMs(callBack, userInfo); callErr != nil {
		panic(callErr)
	}
}

var _colorSyncIterateInstalledProfiles func(callBack ColorSyncProfileIterateCallback, seed *uint32, userInfo unsafe.Pointer, err *corefoundation.CFErrorRef)
var _colorSyncIterateInstalledProfilesErr error

func tryColorSyncIterateInstalledProfiles(callBack ColorSyncProfileIterateCallback, seed *uint32, userInfo unsafe.Pointer, err *corefoundation.CFErrorRef) error {
	if _colorSyncIterateInstalledProfiles == nil {
		return symbolCallError("ColorSyncIterateInstalledProfiles", "10.13", _colorSyncIterateInstalledProfilesErr)
	}
	_colorSyncIterateInstalledProfiles(callBack, seed, userInfo, err)
	return nil
}

// ColorSyncIterateInstalledProfiles iterates over the installed profiles.
//
// See: https://developer.apple.com/documentation/ColorSync/ColorSyncIterateInstalledProfiles(_:_:_:_:)
func ColorSyncIterateInstalledProfiles(callBack ColorSyncProfileIterateCallback, seed *uint32, userInfo unsafe.Pointer, err *corefoundation.CFErrorRef) {
	if callErr := tryColorSyncIterateInstalledProfiles(callBack, seed, userInfo, err); callErr != nil {
		panic(callErr)
	}
}

var _colorSyncIterateInstalledProfilesWithOptions func(callBack ColorSyncProfileIterateCallback, seed *uint32, userInfo unsafe.Pointer, options corefoundation.CFDictionaryRef, err *corefoundation.CFErrorRef)
var _colorSyncIterateInstalledProfilesWithOptionsErr error

func tryColorSyncIterateInstalledProfilesWithOptions(callBack ColorSyncProfileIterateCallback, seed *uint32, userInfo unsafe.Pointer, options corefoundation.CFDictionaryRef, err *corefoundation.CFErrorRef) error {
	if _colorSyncIterateInstalledProfilesWithOptions == nil {
		return symbolCallError("ColorSyncIterateInstalledProfilesWithOptions", "10.13", _colorSyncIterateInstalledProfilesWithOptionsErr)
	}
	_colorSyncIterateInstalledProfilesWithOptions(callBack, seed, userInfo, options, err)
	return nil
}

// ColorSyncIterateInstalledProfilesWithOptions iterates over the installed profiles, using the given options.
//
// See: https://developer.apple.com/documentation/ColorSync/ColorSyncIterateInstalledProfilesWithOptions(_:_:_:_:_:)
func ColorSyncIterateInstalledProfilesWithOptions(callBack ColorSyncProfileIterateCallback, seed *uint32, userInfo unsafe.Pointer, options corefoundation.CFDictionaryRef, err *corefoundation.CFErrorRef) {
	if callErr := tryColorSyncIterateInstalledProfilesWithOptions(callBack, seed, userInfo, options, err); callErr != nil {
		panic(callErr)
	}
}

var _colorSyncProfileContainsTag func(prof coregraphics.ColorSyncProfileRef, signature corefoundation.CFStringRef) bool
var _colorSyncProfileContainsTagErr error

func tryColorSyncProfileContainsTag(prof coregraphics.ColorSyncProfileRef, signature corefoundation.CFStringRef) (bool, error) {
	if _colorSyncProfileContainsTag == nil {
		return false, symbolCallError("ColorSyncProfileContainsTag", "10.13", _colorSyncProfileContainsTagErr)
	}
	return _colorSyncProfileContainsTag(prof, signature), nil
}

// ColorSyncProfileContainsTag returns a Boolean value indicating whether a profile contains a given tag.
//
// See: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileContainsTag(_:_:)
func ColorSyncProfileContainsTag(prof coregraphics.ColorSyncProfileRef, signature corefoundation.CFStringRef) bool {
	result, callErr := tryColorSyncProfileContainsTag(prof, signature)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _colorSyncProfileCopyData func(prof coregraphics.ColorSyncProfileRef, err *corefoundation.CFErrorRef) corefoundation.CFDataRef
var _colorSyncProfileCopyDataErr error

func tryColorSyncProfileCopyData(prof coregraphics.ColorSyncProfileRef, err *corefoundation.CFErrorRef) (corefoundation.CFDataRef, error) {
	if _colorSyncProfileCopyData == nil {
		return *new(corefoundation.CFDataRef), symbolCallError("ColorSyncProfileCopyData", "10.13", _colorSyncProfileCopyDataErr)
	}
	return _colorSyncProfileCopyData(prof, err), nil
}

// ColorSyncProfileCopyData copies the flattened data from a profile.
//
// See: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCopyData(_:_:)
func ColorSyncProfileCopyData(prof coregraphics.ColorSyncProfileRef, err *corefoundation.CFErrorRef) corefoundation.CFDataRef {
	result, callErr := tryColorSyncProfileCopyData(prof, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _colorSyncProfileCopyDescriptionString func(prof coregraphics.ColorSyncProfileRef) corefoundation.CFStringRef
var _colorSyncProfileCopyDescriptionStringErr error

func tryColorSyncProfileCopyDescriptionString(prof coregraphics.ColorSyncProfileRef) (corefoundation.CFStringRef, error) {
	if _colorSyncProfileCopyDescriptionString == nil {
		return *new(corefoundation.CFStringRef), symbolCallError("ColorSyncProfileCopyDescriptionString", "10.13", _colorSyncProfileCopyDescriptionStringErr)
	}
	return _colorSyncProfileCopyDescriptionString(prof), nil
}

// ColorSyncProfileCopyDescriptionString copies the localized description string of a profile.
//
// See: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCopyDescriptionString(_:)
func ColorSyncProfileCopyDescriptionString(prof coregraphics.ColorSyncProfileRef) corefoundation.CFStringRef {
	result, callErr := tryColorSyncProfileCopyDescriptionString(prof)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _colorSyncProfileCopyHeader func(prof coregraphics.ColorSyncProfileRef) corefoundation.CFDataRef
var _colorSyncProfileCopyHeaderErr error

func tryColorSyncProfileCopyHeader(prof coregraphics.ColorSyncProfileRef) (corefoundation.CFDataRef, error) {
	if _colorSyncProfileCopyHeader == nil {
		return *new(corefoundation.CFDataRef), symbolCallError("ColorSyncProfileCopyHeader", "10.13", _colorSyncProfileCopyHeaderErr)
	}
	return _colorSyncProfileCopyHeader(prof), nil
}

// ColorSyncProfileCopyHeader copies the header from a profile.
//
// See: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCopyHeader(_:)
func ColorSyncProfileCopyHeader(prof coregraphics.ColorSyncProfileRef) corefoundation.CFDataRef {
	result, callErr := tryColorSyncProfileCopyHeader(prof)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _colorSyncProfileCopyTag func(prof coregraphics.ColorSyncProfileRef, signature corefoundation.CFStringRef) corefoundation.CFDataRef
var _colorSyncProfileCopyTagErr error

func tryColorSyncProfileCopyTag(prof coregraphics.ColorSyncProfileRef, signature corefoundation.CFStringRef) (corefoundation.CFDataRef, error) {
	if _colorSyncProfileCopyTag == nil {
		return *new(corefoundation.CFDataRef), symbolCallError("ColorSyncProfileCopyTag", "10.13", _colorSyncProfileCopyTagErr)
	}
	return _colorSyncProfileCopyTag(prof, signature), nil
}

// ColorSyncProfileCopyTag copies a tag from a profile.
//
// See: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCopyTag(_:_:)
func ColorSyncProfileCopyTag(prof coregraphics.ColorSyncProfileRef, signature corefoundation.CFStringRef) corefoundation.CFDataRef {
	result, callErr := tryColorSyncProfileCopyTag(prof, signature)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _colorSyncProfileCopyTagSignatures func(prof coregraphics.ColorSyncProfileRef) corefoundation.CFArrayRef
var _colorSyncProfileCopyTagSignaturesErr error

func tryColorSyncProfileCopyTagSignatures(prof coregraphics.ColorSyncProfileRef) (corefoundation.CFArrayRef, error) {
	if _colorSyncProfileCopyTagSignatures == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("ColorSyncProfileCopyTagSignatures", "10.13", _colorSyncProfileCopyTagSignaturesErr)
	}
	return _colorSyncProfileCopyTagSignatures(prof), nil
}

// ColorSyncProfileCopyTagSignatures copies the tag signatures of a profile.
//
// See: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCopyTagSignatures(_:)
func ColorSyncProfileCopyTagSignatures(prof coregraphics.ColorSyncProfileRef) corefoundation.CFArrayRef {
	result, callErr := tryColorSyncProfileCopyTagSignatures(prof)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _colorSyncProfileCreate func(data corefoundation.CFDataRef, err *corefoundation.CFErrorRef) coregraphics.ColorSyncProfileRef
var _colorSyncProfileCreateErr error

func tryColorSyncProfileCreate(data corefoundation.CFDataRef, err *corefoundation.CFErrorRef) (coregraphics.ColorSyncProfileRef, error) {
	if _colorSyncProfileCreate == nil {
		return *new(coregraphics.ColorSyncProfileRef), symbolCallError("ColorSyncProfileCreate", "10.13", _colorSyncProfileCreateErr)
	}
	return _colorSyncProfileCreate(data, err), nil
}

// ColorSyncProfileCreate creates a profile from ICC profile data.
//
// See: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCreate(_:_:)
func ColorSyncProfileCreate(data corefoundation.CFDataRef, err *corefoundation.CFErrorRef) coregraphics.ColorSyncProfileRef {
	result, callErr := tryColorSyncProfileCreate(data, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _colorSyncProfileCreateDeviceProfile func(deviceClass corefoundation.CFStringRef, deviceID corefoundation.CFUUIDRef, profileID corefoundation.CFTypeRef) coregraphics.ColorSyncProfileRef
var _colorSyncProfileCreateDeviceProfileErr error

func tryColorSyncProfileCreateDeviceProfile(deviceClass corefoundation.CFStringRef, deviceID corefoundation.CFUUIDRef, profileID corefoundation.CFTypeRef) (coregraphics.ColorSyncProfileRef, error) {
	if _colorSyncProfileCreateDeviceProfile == nil {
		return *new(coregraphics.ColorSyncProfileRef), symbolCallError("ColorSyncProfileCreateDeviceProfile", "10.13", _colorSyncProfileCreateDeviceProfileErr)
	}
	return _colorSyncProfileCreateDeviceProfile(deviceClass, deviceID, profileID), nil
}

// ColorSyncProfileCreateDeviceProfile creates a profile for a device registered with ColorSync.
//
// See: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCreateDeviceProfile(_:_:_:)
func ColorSyncProfileCreateDeviceProfile(deviceClass corefoundation.CFStringRef, deviceID corefoundation.CFUUIDRef, profileID corefoundation.CFTypeRef) coregraphics.ColorSyncProfileRef {
	result, callErr := tryColorSyncProfileCreateDeviceProfile(deviceClass, deviceID, profileID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _colorSyncProfileCreateDisplayTransferTablesFromVCGT func(profile coregraphics.ColorSyncProfileRef, nSamplesPerChannel *uintptr) corefoundation.CFDataRef
var _colorSyncProfileCreateDisplayTransferTablesFromVCGTErr error

func tryColorSyncProfileCreateDisplayTransferTablesFromVCGT(profile coregraphics.ColorSyncProfileRef, nSamplesPerChannel *uintptr) (corefoundation.CFDataRef, error) {
	if _colorSyncProfileCreateDisplayTransferTablesFromVCGT == nil {
		return *new(corefoundation.CFDataRef), symbolCallError("ColorSyncProfileCreateDisplayTransferTablesFromVCGT", "10.13", _colorSyncProfileCreateDisplayTransferTablesFromVCGTErr)
	}
	return _colorSyncProfileCreateDisplayTransferTablesFromVCGT(profile, nSamplesPerChannel), nil
}

// ColorSyncProfileCreateDisplayTransferTablesFromVCGT creates display transfer tables from the profile’s `vcgt` tag.
//
// See: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCreateDisplayTransferTablesFromVCGT(_:_:)
func ColorSyncProfileCreateDisplayTransferTablesFromVCGT(profile coregraphics.ColorSyncProfileRef, nSamplesPerChannel *uintptr) corefoundation.CFDataRef {
	result, callErr := tryColorSyncProfileCreateDisplayTransferTablesFromVCGT(profile, nSamplesPerChannel)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _colorSyncProfileCreateLink func(profileInfo corefoundation.CFArrayRef, options corefoundation.CFDictionaryRef) coregraphics.ColorSyncProfileRef
var _colorSyncProfileCreateLinkErr error

func tryColorSyncProfileCreateLink(profileInfo corefoundation.CFArrayRef, options corefoundation.CFDictionaryRef) (coregraphics.ColorSyncProfileRef, error) {
	if _colorSyncProfileCreateLink == nil {
		return *new(coregraphics.ColorSyncProfileRef), symbolCallError("ColorSyncProfileCreateLink", "10.13", _colorSyncProfileCreateLinkErr)
	}
	return _colorSyncProfileCreateLink(profileInfo, options), nil
}

// ColorSyncProfileCreateLink creates a device link profile from an array of profiles.
//
// See: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCreateLink(_:_:)
func ColorSyncProfileCreateLink(profileInfo corefoundation.CFArrayRef, options corefoundation.CFDictionaryRef) coregraphics.ColorSyncProfileRef {
	result, callErr := tryColorSyncProfileCreateLink(profileInfo, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _colorSyncProfileCreateMutable func() unsafe.Pointer
var _colorSyncProfileCreateMutableErr error

func tryColorSyncProfileCreateMutable() (unsafe.Pointer, error) {
	if _colorSyncProfileCreateMutable == nil {
		return nil, symbolCallError("ColorSyncProfileCreateMutable", "10.13", _colorSyncProfileCreateMutableErr)
	}
	return _colorSyncProfileCreateMutable(), nil
}

// ColorSyncProfileCreateMutable creates an empty mutable profile.
//
// See: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCreateMutable()
func ColorSyncProfileCreateMutable() unsafe.Pointer {
	result, callErr := tryColorSyncProfileCreateMutable()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _colorSyncProfileCreateMutableCopy func(prof coregraphics.ColorSyncProfileRef) unsafe.Pointer
var _colorSyncProfileCreateMutableCopyErr error

func tryColorSyncProfileCreateMutableCopy(prof coregraphics.ColorSyncProfileRef) (unsafe.Pointer, error) {
	if _colorSyncProfileCreateMutableCopy == nil {
		return nil, symbolCallError("ColorSyncProfileCreateMutableCopy", "10.13", _colorSyncProfileCreateMutableCopyErr)
	}
	return _colorSyncProfileCreateMutableCopy(prof), nil
}

// ColorSyncProfileCreateMutableCopy creates a mutable copy of a profile.
//
// See: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCreateMutableCopy(_:)
func ColorSyncProfileCreateMutableCopy(prof coregraphics.ColorSyncProfileRef) unsafe.Pointer {
	result, callErr := tryColorSyncProfileCreateMutableCopy(prof)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _colorSyncProfileCreateWithDisplayID func(displayID uint32) coregraphics.ColorSyncProfileRef
var _colorSyncProfileCreateWithDisplayIDErr error

func tryColorSyncProfileCreateWithDisplayID(displayID uint32) (coregraphics.ColorSyncProfileRef, error) {
	if _colorSyncProfileCreateWithDisplayID == nil {
		return *new(coregraphics.ColorSyncProfileRef), symbolCallError("ColorSyncProfileCreateWithDisplayID", "10.13", _colorSyncProfileCreateWithDisplayIDErr)
	}
	return _colorSyncProfileCreateWithDisplayID(displayID), nil
}

// ColorSyncProfileCreateWithDisplayID creates a profile for the specified display.
//
// See: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCreateWithDisplayID(_:)
func ColorSyncProfileCreateWithDisplayID(displayID uint32) coregraphics.ColorSyncProfileRef {
	result, callErr := tryColorSyncProfileCreateWithDisplayID(displayID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _colorSyncProfileCreateWithName func(name corefoundation.CFStringRef) coregraphics.ColorSyncProfileRef
var _colorSyncProfileCreateWithNameErr error

func tryColorSyncProfileCreateWithName(name corefoundation.CFStringRef) (coregraphics.ColorSyncProfileRef, error) {
	if _colorSyncProfileCreateWithName == nil {
		return *new(coregraphics.ColorSyncProfileRef), symbolCallError("ColorSyncProfileCreateWithName", "10.13", _colorSyncProfileCreateWithNameErr)
	}
	return _colorSyncProfileCreateWithName(name), nil
}

// ColorSyncProfileCreateWithName creates a profile from a predefined profile name.
//
// See: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCreateWithName(_:)
func ColorSyncProfileCreateWithName(name corefoundation.CFStringRef) coregraphics.ColorSyncProfileRef {
	result, callErr := tryColorSyncProfileCreateWithName(name)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _colorSyncProfileCreateWithURL func(url corefoundation.CFURLRef, err *corefoundation.CFErrorRef) coregraphics.ColorSyncProfileRef
var _colorSyncProfileCreateWithURLErr error

func tryColorSyncProfileCreateWithURL(url corefoundation.CFURLRef, err *corefoundation.CFErrorRef) (coregraphics.ColorSyncProfileRef, error) {
	if _colorSyncProfileCreateWithURL == nil {
		return *new(coregraphics.ColorSyncProfileRef), symbolCallError("ColorSyncProfileCreateWithURL", "10.13", _colorSyncProfileCreateWithURLErr)
	}
	return _colorSyncProfileCreateWithURL(url, err), nil
}

// ColorSyncProfileCreateWithURL creates a profile from ICC profile data at a URL.
//
// See: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCreateWithURL(_:_:)
func ColorSyncProfileCreateWithURL(url corefoundation.CFURLRef, err *corefoundation.CFErrorRef) coregraphics.ColorSyncProfileRef {
	result, callErr := tryColorSyncProfileCreateWithURL(url, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _colorSyncProfileCreateWithURLAndOptions func(url corefoundation.CFURLRef, options corefoundation.CFDictionaryRef, err *corefoundation.CFErrorRef) coregraphics.ColorSyncProfileRef
var _colorSyncProfileCreateWithURLAndOptionsErr error

func tryColorSyncProfileCreateWithURLAndOptions(url corefoundation.CFURLRef, options corefoundation.CFDictionaryRef, err *corefoundation.CFErrorRef) (coregraphics.ColorSyncProfileRef, error) {
	if _colorSyncProfileCreateWithURLAndOptions == nil {
		return *new(coregraphics.ColorSyncProfileRef), symbolCallError("ColorSyncProfileCreateWithURLAndOptions", "10.13", _colorSyncProfileCreateWithURLAndOptionsErr)
	}
	return _colorSyncProfileCreateWithURLAndOptions(url, options, err), nil
}

// ColorSyncProfileCreateWithURLAndOptions creates a profile from ICC profile data at a URL, using the given options.
//
// See: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCreateWithURLAndOptions(_:_:_:)
func ColorSyncProfileCreateWithURLAndOptions(url corefoundation.CFURLRef, options corefoundation.CFDictionaryRef, err *corefoundation.CFErrorRef) coregraphics.ColorSyncProfileRef {
	result, callErr := tryColorSyncProfileCreateWithURLAndOptions(url, options, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _colorSyncProfileEstimateGamma func(prof coregraphics.ColorSyncProfileRef, err *corefoundation.CFErrorRef) float32
var _colorSyncProfileEstimateGammaErr error

func tryColorSyncProfileEstimateGamma(prof coregraphics.ColorSyncProfileRef, err *corefoundation.CFErrorRef) (float32, error) {
	if _colorSyncProfileEstimateGamma == nil {
		return 0.0, symbolCallError("ColorSyncProfileEstimateGamma", "10.13", _colorSyncProfileEstimateGammaErr)
	}
	return _colorSyncProfileEstimateGamma(prof, err), nil
}

// ColorSyncProfileEstimateGamma estimates the gamma of a profile.
//
// See: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileEstimateGamma(_:_:)
func ColorSyncProfileEstimateGamma(prof coregraphics.ColorSyncProfileRef, err *corefoundation.CFErrorRef) float32 {
	result, callErr := tryColorSyncProfileEstimateGamma(prof, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _colorSyncProfileEstimateGammaWithDisplayID func(displayID int32, err *corefoundation.CFErrorRef) float32
var _colorSyncProfileEstimateGammaWithDisplayIDErr error

func tryColorSyncProfileEstimateGammaWithDisplayID(displayID int32, err *corefoundation.CFErrorRef) (float32, error) {
	if _colorSyncProfileEstimateGammaWithDisplayID == nil {
		return 0.0, symbolCallError("ColorSyncProfileEstimateGammaWithDisplayID", "10.13", _colorSyncProfileEstimateGammaWithDisplayIDErr)
	}
	return _colorSyncProfileEstimateGammaWithDisplayID(displayID, err), nil
}

// ColorSyncProfileEstimateGammaWithDisplayID estimates the gamma of the profile for the specified display.
//
// See: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileEstimateGammaWithDisplayID(_:_:)
func ColorSyncProfileEstimateGammaWithDisplayID(displayID int32, err *corefoundation.CFErrorRef) float32 {
	result, callErr := tryColorSyncProfileEstimateGammaWithDisplayID(displayID, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _colorSyncProfileGetDisplayTransferFormulaFromVCGT func(profile coregraphics.ColorSyncProfileRef, redMin *float32, redMax *float32, redGamma *float32, greenMin *float32, greenMax *float32, greenGamma *float32, blueMin *float32, blueMax *float32, blueGamma *float32) bool
var _colorSyncProfileGetDisplayTransferFormulaFromVCGTErr error

func tryColorSyncProfileGetDisplayTransferFormulaFromVCGT(profile coregraphics.ColorSyncProfileRef, redMin []float32, redMax []float32, redGamma []float32, greenMin []float32, greenMax []float32, greenGamma []float32, blueMin []float32, blueMax []float32, blueGamma []float32) (bool, error) {
	if _colorSyncProfileGetDisplayTransferFormulaFromVCGT == nil {
		return false, symbolCallError("ColorSyncProfileGetDisplayTransferFormulaFromVCGT", "10.13", _colorSyncProfileGetDisplayTransferFormulaFromVCGTErr)
	}
	return _colorSyncProfileGetDisplayTransferFormulaFromVCGT(profile, unsafe.SliceData(redMin), unsafe.SliceData(redMax), unsafe.SliceData(redGamma), unsafe.SliceData(greenMin), unsafe.SliceData(greenMax), unsafe.SliceData(greenGamma), unsafe.SliceData(blueMin), unsafe.SliceData(blueMax), unsafe.SliceData(blueGamma)), nil
}

// ColorSyncProfileGetDisplayTransferFormulaFromVCGT converts the profile’s `vcgt` tag to formula components used by [CGSetDisplayTransferByFormula].
//
// See: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileGetDisplayTransferFormulaFromVCGT(_:_:_:_:_:_:_:_:_:_:)
func ColorSyncProfileGetDisplayTransferFormulaFromVCGT(profile coregraphics.ColorSyncProfileRef, redMin []float32, redMax []float32, redGamma []float32, greenMin []float32, greenMax []float32, greenGamma []float32, blueMin []float32, blueMax []float32, blueGamma []float32) bool {
	result, callErr := tryColorSyncProfileGetDisplayTransferFormulaFromVCGT(profile, redMin, redMax, redGamma, greenMin, greenMax, greenGamma, blueMin, blueMax, blueGamma)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _colorSyncProfileGetMD5 func(prof coregraphics.ColorSyncProfileRef) ColorSyncMD5
var _colorSyncProfileGetMD5Err error

func tryColorSyncProfileGetMD5(prof coregraphics.ColorSyncProfileRef) (ColorSyncMD5, error) {
	if _colorSyncProfileGetMD5 == nil {
		return ColorSyncMD5{}, symbolCallError("ColorSyncProfileGetMD5", "10.13", _colorSyncProfileGetMD5Err)
	}
	return _colorSyncProfileGetMD5(prof), nil
}

// ColorSyncProfileGetMD5 returns the MD5 digest for a profile.
//
// See: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileGetMD5(_:)
func ColorSyncProfileGetMD5(prof coregraphics.ColorSyncProfileRef) ColorSyncMD5 {
	result, callErr := tryColorSyncProfileGetMD5(prof)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _colorSyncProfileGetTypeID func() uint
var _colorSyncProfileGetTypeIDErr error

func tryColorSyncProfileGetTypeID() (uint, error) {
	if _colorSyncProfileGetTypeID == nil {
		return 0, symbolCallError("ColorSyncProfileGetTypeID", "10.13", _colorSyncProfileGetTypeIDErr)
	}
	return _colorSyncProfileGetTypeID(), nil
}

// ColorSyncProfileGetTypeID returns the unique identifier for the ColorSync profile opaque type.
//
// See: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileGetTypeID()
func ColorSyncProfileGetTypeID() uint {
	result, callErr := tryColorSyncProfileGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _colorSyncProfileGetURL func(prof coregraphics.ColorSyncProfileRef, err *corefoundation.CFErrorRef) corefoundation.CFURLRef
var _colorSyncProfileGetURLErr error

func tryColorSyncProfileGetURL(prof coregraphics.ColorSyncProfileRef, err *corefoundation.CFErrorRef) (corefoundation.CFURLRef, error) {
	if _colorSyncProfileGetURL == nil {
		return *new(corefoundation.CFURLRef), symbolCallError("ColorSyncProfileGetURL", "10.13", _colorSyncProfileGetURLErr)
	}
	return _colorSyncProfileGetURL(prof, err), nil
}

// ColorSyncProfileGetURL returns the URL of a profile.
//
// See: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileGetURL(_:_:)
func ColorSyncProfileGetURL(prof coregraphics.ColorSyncProfileRef, err *corefoundation.CFErrorRef) corefoundation.CFURLRef {
	result, callErr := tryColorSyncProfileGetURL(prof, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _colorSyncProfileInstall func(profile coregraphics.ColorSyncProfileRef, domain corefoundation.CFStringRef, subpath corefoundation.CFStringRef, err *corefoundation.CFErrorRef) bool
var _colorSyncProfileInstallErr error

func tryColorSyncProfileInstall(profile coregraphics.ColorSyncProfileRef, domain corefoundation.CFStringRef, subpath corefoundation.CFStringRef, err *corefoundation.CFErrorRef) (bool, error) {
	if _colorSyncProfileInstall == nil {
		return false, symbolCallError("ColorSyncProfileInstall", "10.13", _colorSyncProfileInstallErr)
	}
	return _colorSyncProfileInstall(profile, domain, subpath, err), nil
}

// ColorSyncProfileInstall installs a profile in the specified domain.
//
// See: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileInstall(_:_:_:_:)
func ColorSyncProfileInstall(profile coregraphics.ColorSyncProfileRef, domain corefoundation.CFStringRef, subpath corefoundation.CFStringRef, err *corefoundation.CFErrorRef) bool {
	result, callErr := tryColorSyncProfileInstall(profile, domain, subpath, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _colorSyncProfileIsHLGBased func(arg0 coregraphics.ColorSyncProfileRef) bool
var _colorSyncProfileIsHLGBasedErr error

func tryColorSyncProfileIsHLGBased(arg0 coregraphics.ColorSyncProfileRef) (bool, error) {
	if _colorSyncProfileIsHLGBased == nil {
		return false, symbolCallError("ColorSyncProfileIsHLGBased", "10.13", _colorSyncProfileIsHLGBasedErr)
	}
	return _colorSyncProfileIsHLGBased(arg0), nil
}

// ColorSyncProfileIsHLGBased returns a Boolean value indicating whether the profile uses ITU BT.2100 HLG transfer functions.
//
// See: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileIsHLGBased(_:)
func ColorSyncProfileIsHLGBased(arg0 coregraphics.ColorSyncProfileRef) bool {
	result, callErr := tryColorSyncProfileIsHLGBased(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _colorSyncProfileIsMatrixBased func(arg0 coregraphics.ColorSyncProfileRef) bool
var _colorSyncProfileIsMatrixBasedErr error

func tryColorSyncProfileIsMatrixBased(arg0 coregraphics.ColorSyncProfileRef) (bool, error) {
	if _colorSyncProfileIsMatrixBased == nil {
		return false, symbolCallError("ColorSyncProfileIsMatrixBased", "10.13", _colorSyncProfileIsMatrixBasedErr)
	}
	return _colorSyncProfileIsMatrixBased(arg0), nil
}

// ColorSyncProfileIsMatrixBased returns a Boolean value indicating whether the profile is matrix-based.
//
// See: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileIsMatrixBased(_:)
func ColorSyncProfileIsMatrixBased(arg0 coregraphics.ColorSyncProfileRef) bool {
	result, callErr := tryColorSyncProfileIsMatrixBased(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _colorSyncProfileIsPQBased func(arg0 coregraphics.ColorSyncProfileRef) bool
var _colorSyncProfileIsPQBasedErr error

func tryColorSyncProfileIsPQBased(arg0 coregraphics.ColorSyncProfileRef) (bool, error) {
	if _colorSyncProfileIsPQBased == nil {
		return false, symbolCallError("ColorSyncProfileIsPQBased", "10.13", _colorSyncProfileIsPQBasedErr)
	}
	return _colorSyncProfileIsPQBased(arg0), nil
}

// ColorSyncProfileIsPQBased returns a Boolean value indicating whether the profile uses ITU BT.2100 PQ transfer functions.
//
// See: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileIsPQBased(_:)
func ColorSyncProfileIsPQBased(arg0 coregraphics.ColorSyncProfileRef) bool {
	result, callErr := tryColorSyncProfileIsPQBased(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _colorSyncProfileIsWideGamut func(arg0 coregraphics.ColorSyncProfileRef) bool
var _colorSyncProfileIsWideGamutErr error

func tryColorSyncProfileIsWideGamut(arg0 coregraphics.ColorSyncProfileRef) (bool, error) {
	if _colorSyncProfileIsWideGamut == nil {
		return false, symbolCallError("ColorSyncProfileIsWideGamut", "10.13", _colorSyncProfileIsWideGamutErr)
	}
	return _colorSyncProfileIsWideGamut(arg0), nil
}

// ColorSyncProfileIsWideGamut returns a Boolean value indicating whether the display profile describes a wide-gamut color space.
//
// See: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileIsWideGamut(_:)
func ColorSyncProfileIsWideGamut(arg0 coregraphics.ColorSyncProfileRef) bool {
	result, callErr := tryColorSyncProfileIsWideGamut(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _colorSyncProfileRemoveTag func(prof unsafe.Pointer, signature corefoundation.CFStringRef)
var _colorSyncProfileRemoveTagErr error

func tryColorSyncProfileRemoveTag(prof unsafe.Pointer, signature corefoundation.CFStringRef) error {
	if _colorSyncProfileRemoveTag == nil {
		return symbolCallError("ColorSyncProfileRemoveTag", "10.13", _colorSyncProfileRemoveTagErr)
	}
	_colorSyncProfileRemoveTag(prof, signature)
	return nil
}

// ColorSyncProfileRemoveTag removes a tag from a mutable profile.
//
// See: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileRemoveTag(_:_:)
func ColorSyncProfileRemoveTag(prof unsafe.Pointer, signature corefoundation.CFStringRef) {
	if callErr := tryColorSyncProfileRemoveTag(prof, signature); callErr != nil {
		panic(callErr)
	}
}

var _colorSyncProfileSetHeader func(prof unsafe.Pointer, header corefoundation.CFDataRef)
var _colorSyncProfileSetHeaderErr error

func tryColorSyncProfileSetHeader(prof unsafe.Pointer, header corefoundation.CFDataRef) error {
	if _colorSyncProfileSetHeader == nil {
		return symbolCallError("ColorSyncProfileSetHeader", "10.13", _colorSyncProfileSetHeaderErr)
	}
	_colorSyncProfileSetHeader(prof, header)
	return nil
}

// ColorSyncProfileSetHeader sets the header of a mutable profile.
//
// See: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileSetHeader(_:_:)
func ColorSyncProfileSetHeader(prof unsafe.Pointer, header corefoundation.CFDataRef) {
	if callErr := tryColorSyncProfileSetHeader(prof, header); callErr != nil {
		panic(callErr)
	}
}

var _colorSyncProfileSetTag func(prof unsafe.Pointer, signature corefoundation.CFStringRef, data corefoundation.CFDataRef)
var _colorSyncProfileSetTagErr error

func tryColorSyncProfileSetTag(prof unsafe.Pointer, signature corefoundation.CFStringRef, data corefoundation.CFDataRef) error {
	if _colorSyncProfileSetTag == nil {
		return symbolCallError("ColorSyncProfileSetTag", "10.13", _colorSyncProfileSetTagErr)
	}
	_colorSyncProfileSetTag(prof, signature, data)
	return nil
}

// ColorSyncProfileSetTag sets a tag in a mutable profile.
//
// See: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileSetTag(_:_:_:)
func ColorSyncProfileSetTag(prof unsafe.Pointer, signature corefoundation.CFStringRef, data corefoundation.CFDataRef) {
	if callErr := tryColorSyncProfileSetTag(prof, signature, data); callErr != nil {
		panic(callErr)
	}
}

var _colorSyncProfileUninstall func(profile coregraphics.ColorSyncProfileRef, err *corefoundation.CFErrorRef) bool
var _colorSyncProfileUninstallErr error

func tryColorSyncProfileUninstall(profile coregraphics.ColorSyncProfileRef, err *corefoundation.CFErrorRef) (bool, error) {
	if _colorSyncProfileUninstall == nil {
		return false, symbolCallError("ColorSyncProfileUninstall", "10.13", _colorSyncProfileUninstallErr)
	}
	return _colorSyncProfileUninstall(profile, err), nil
}

// ColorSyncProfileUninstall uninstalls a profile.
//
// See: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileUninstall(_:_:)
func ColorSyncProfileUninstall(profile coregraphics.ColorSyncProfileRef, err *corefoundation.CFErrorRef) bool {
	result, callErr := tryColorSyncProfileUninstall(profile, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _colorSyncProfileVerify func(prof coregraphics.ColorSyncProfileRef, errors *corefoundation.CFErrorRef, warnings *corefoundation.CFErrorRef) bool
var _colorSyncProfileVerifyErr error

func tryColorSyncProfileVerify(prof coregraphics.ColorSyncProfileRef, errors *corefoundation.CFErrorRef, warnings *corefoundation.CFErrorRef) (bool, error) {
	if _colorSyncProfileVerify == nil {
		return false, symbolCallError("ColorSyncProfileVerify", "10.13", _colorSyncProfileVerifyErr)
	}
	return _colorSyncProfileVerify(prof, errors, warnings), nil
}

// ColorSyncProfileVerify verifies whether a profile can be used.
//
// See: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileVerify(_:_:_:)
func ColorSyncProfileVerify(prof coregraphics.ColorSyncProfileRef, errors *corefoundation.CFErrorRef, warnings *corefoundation.CFErrorRef) bool {
	result, callErr := tryColorSyncProfileVerify(prof, errors, warnings)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _colorSyncRegisterDevice func(deviceClass corefoundation.CFStringRef, deviceID corefoundation.CFUUIDRef, deviceInfo corefoundation.CFDictionaryRef) bool
var _colorSyncRegisterDeviceErr error

func tryColorSyncRegisterDevice(deviceClass corefoundation.CFStringRef, deviceID corefoundation.CFUUIDRef, deviceInfo corefoundation.CFDictionaryRef) (bool, error) {
	if _colorSyncRegisterDevice == nil {
		return false, symbolCallError("ColorSyncRegisterDevice", "10.13", _colorSyncRegisterDeviceErr)
	}
	return _colorSyncRegisterDevice(deviceClass, deviceID, deviceInfo), nil
}

// ColorSyncRegisterDevice registers a device of the given class with ColorSync.
//
// See: https://developer.apple.com/documentation/ColorSync/ColorSyncRegisterDevice(_:_:_:)
func ColorSyncRegisterDevice(deviceClass corefoundation.CFStringRef, deviceID corefoundation.CFUUIDRef, deviceInfo corefoundation.CFDictionaryRef) bool {
	result, callErr := tryColorSyncRegisterDevice(deviceClass, deviceID, deviceInfo)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _colorSyncTransformConvert func(transform unsafe.Pointer, width uintptr, height uintptr, dst unsafe.Pointer, dstDepth ColorSyncDataDepth, dstLayout uint32, dstBytesPerRow uintptr, src unsafe.Pointer, srcDepth ColorSyncDataDepth, srcLayout uint32, srcBytesPerRow uintptr, options corefoundation.CFDictionaryRef) bool
var _colorSyncTransformConvertErr error

func tryColorSyncTransformConvert(transform unsafe.Pointer, width uintptr, height uintptr, dst unsafe.Pointer, dstDepth ColorSyncDataDepth, dstLayout uint32, dstBytesPerRow uintptr, src unsafe.Pointer, srcDepth ColorSyncDataDepth, srcLayout uint32, srcBytesPerRow uintptr, options corefoundation.CFDictionaryRef) (bool, error) {
	if _colorSyncTransformConvert == nil {
		return false, symbolCallError("ColorSyncTransformConvert", "10.13", _colorSyncTransformConvertErr)
	}
	return _colorSyncTransformConvert(transform, width, height, dst, dstDepth, dstLayout, dstBytesPerRow, src, srcDepth, srcLayout, srcBytesPerRow, options), nil
}

// ColorSyncTransformConvert converts color data from a source layout to a destination layout using a color transform.
//
// See: https://developer.apple.com/documentation/ColorSync/ColorSyncTransformConvert(_:_:_:_:_:_:_:_:_:_:_:_:)
func ColorSyncTransformConvert(transform unsafe.Pointer, width uintptr, height uintptr, dst unsafe.Pointer, dstDepth ColorSyncDataDepth, dstLayout uint32, dstBytesPerRow uintptr, src unsafe.Pointer, srcDepth ColorSyncDataDepth, srcLayout uint32, srcBytesPerRow uintptr, options corefoundation.CFDictionaryRef) bool {
	result, callErr := tryColorSyncTransformConvert(transform, width, height, dst, dstDepth, dstLayout, dstBytesPerRow, src, srcDepth, srcLayout, srcBytesPerRow, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _colorSyncTransformCopyProperty func(transform unsafe.Pointer, key corefoundation.CFTypeRef, options corefoundation.CFDictionaryRef) corefoundation.CFTypeRef
var _colorSyncTransformCopyPropertyErr error

func tryColorSyncTransformCopyProperty(transform unsafe.Pointer, key corefoundation.CFTypeRef, options corefoundation.CFDictionaryRef) (corefoundation.CFTypeRef, error) {
	if _colorSyncTransformCopyProperty == nil {
		return *new(corefoundation.CFTypeRef), symbolCallError("ColorSyncTransformCopyProperty", "10.13", _colorSyncTransformCopyPropertyErr)
	}
	return _colorSyncTransformCopyProperty(transform, key, options), nil
}

// ColorSyncTransformCopyProperty copies a property from a color transform.
//
// See: https://developer.apple.com/documentation/ColorSync/ColorSyncTransformCopyProperty(_:_:_:)
func ColorSyncTransformCopyProperty(transform unsafe.Pointer, key corefoundation.CFTypeRef, options corefoundation.CFDictionaryRef) corefoundation.CFTypeRef {
	result, callErr := tryColorSyncTransformCopyProperty(transform, key, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _colorSyncTransformCreate func(profileSequence corefoundation.CFArrayRef, options corefoundation.CFDictionaryRef) unsafe.Pointer
var _colorSyncTransformCreateErr error

func tryColorSyncTransformCreate(profileSequence corefoundation.CFArrayRef, options corefoundation.CFDictionaryRef) (unsafe.Pointer, error) {
	if _colorSyncTransformCreate == nil {
		return nil, symbolCallError("ColorSyncTransformCreate", "10.13", _colorSyncTransformCreateErr)
	}
	return _colorSyncTransformCreate(profileSequence, options), nil
}

// ColorSyncTransformCreate creates a color transform from a sequence of profiles.
//
// See: https://developer.apple.com/documentation/ColorSync/ColorSyncTransformCreate(_:_:)
func ColorSyncTransformCreate(profileSequence corefoundation.CFArrayRef, options corefoundation.CFDictionaryRef) unsafe.Pointer {
	result, callErr := tryColorSyncTransformCreate(profileSequence, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _colorSyncTransformGetProfileSequence func(transform unsafe.Pointer) corefoundation.CFArrayRef
var _colorSyncTransformGetProfileSequenceErr error

func tryColorSyncTransformGetProfileSequence(transform unsafe.Pointer) (corefoundation.CFArrayRef, error) {
	if _colorSyncTransformGetProfileSequence == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("ColorSyncTransformGetProfileSequence", "10.13", _colorSyncTransformGetProfileSequenceErr)
	}
	return _colorSyncTransformGetProfileSequence(transform), nil
}

// ColorSyncTransformGetProfileSequence returns the profile sequence used to create a color transform.
//
// See: https://developer.apple.com/documentation/ColorSync/ColorSyncTransformGetProfileSequence(_:)
func ColorSyncTransformGetProfileSequence(transform unsafe.Pointer) corefoundation.CFArrayRef {
	result, callErr := tryColorSyncTransformGetProfileSequence(transform)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _colorSyncTransformGetTypeID func() uint
var _colorSyncTransformGetTypeIDErr error

func tryColorSyncTransformGetTypeID() (uint, error) {
	if _colorSyncTransformGetTypeID == nil {
		return 0, symbolCallError("ColorSyncTransformGetTypeID", "10.13", _colorSyncTransformGetTypeIDErr)
	}
	return _colorSyncTransformGetTypeID(), nil
}

// ColorSyncTransformGetTypeID returns the type identifier for the [ColorSyncTransform] opaque type.
//
// See: https://developer.apple.com/documentation/ColorSync/ColorSyncTransformGetTypeID()
func ColorSyncTransformGetTypeID() uint {
	result, callErr := tryColorSyncTransformGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _colorSyncTransformSetProperty func(transform unsafe.Pointer, key corefoundation.CFTypeRef, property corefoundation.CFTypeRef)
var _colorSyncTransformSetPropertyErr error

func tryColorSyncTransformSetProperty(transform unsafe.Pointer, key corefoundation.CFTypeRef, property corefoundation.CFTypeRef) error {
	if _colorSyncTransformSetProperty == nil {
		return symbolCallError("ColorSyncTransformSetProperty", "10.13", _colorSyncTransformSetPropertyErr)
	}
	_colorSyncTransformSetProperty(transform, key, property)
	return nil
}

// ColorSyncTransformSetProperty sets a property on a color transform.
//
// See: https://developer.apple.com/documentation/ColorSync/ColorSyncTransformSetProperty(_:_:_:)
func ColorSyncTransformSetProperty(transform unsafe.Pointer, key corefoundation.CFTypeRef, property corefoundation.CFTypeRef) {
	if callErr := tryColorSyncTransformSetProperty(transform, key, property); callErr != nil {
		panic(callErr)
	}
}

var _colorSyncUnregisterDevice func(deviceClass corefoundation.CFStringRef, deviceID corefoundation.CFUUIDRef) bool
var _colorSyncUnregisterDeviceErr error

func tryColorSyncUnregisterDevice(deviceClass corefoundation.CFStringRef, deviceID corefoundation.CFUUIDRef) (bool, error) {
	if _colorSyncUnregisterDevice == nil {
		return false, symbolCallError("ColorSyncUnregisterDevice", "10.13", _colorSyncUnregisterDeviceErr)
	}
	return _colorSyncUnregisterDevice(deviceClass, deviceID), nil
}

// ColorSyncUnregisterDevice unregisters a device of the given class and identifier.
//
// See: https://developer.apple.com/documentation/ColorSync/ColorSyncUnregisterDevice(_:_:)
func ColorSyncUnregisterDevice(deviceClass corefoundation.CFStringRef, deviceID corefoundation.CFUUIDRef) bool {
	result, callErr := tryColorSyncUnregisterDevice(deviceClass, deviceID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_cGDisplayCreateUUIDFromDisplayID, &_cGDisplayCreateUUIDFromDisplayIDErr, frameworkHandle, "CGDisplayCreateUUIDFromDisplayID", "10.13")
	registerFunc(&_cGDisplayGetDisplayIDFromUUID, &_cGDisplayGetDisplayIDFromUUIDErr, frameworkHandle, "CGDisplayGetDisplayIDFromUUID", "10.13")
	registerFunc(&_colorSyncAPIVersion, &_colorSyncAPIVersionErr, frameworkHandle, "ColorSyncAPIVersion", "10.13")
	registerFunc(&_colorSyncCMMCopyCMMIdentifier, &_colorSyncCMMCopyCMMIdentifierErr, frameworkHandle, "ColorSyncCMMCopyCMMIdentifier", "10.13")
	registerFunc(&_colorSyncCMMCopyLocalizedName, &_colorSyncCMMCopyLocalizedNameErr, frameworkHandle, "ColorSyncCMMCopyLocalizedName", "10.13")
	registerFunc(&_colorSyncCMMCreate, &_colorSyncCMMCreateErr, frameworkHandle, "ColorSyncCMMCreate", "10.13")
	registerFunc(&_colorSyncCMMGetBundle, &_colorSyncCMMGetBundleErr, frameworkHandle, "ColorSyncCMMGetBundle", "10.13")
	registerFunc(&_colorSyncCMMGetTypeID, &_colorSyncCMMGetTypeIDErr, frameworkHandle, "ColorSyncCMMGetTypeID", "10.13")
	registerFunc(&_colorSyncCreateCodeFragment, &_colorSyncCreateCodeFragmentErr, frameworkHandle, "ColorSyncCreateCodeFragment", "10.13")
	registerFunc(&_colorSyncDeviceCopyDeviceInfo, &_colorSyncDeviceCopyDeviceInfoErr, frameworkHandle, "ColorSyncDeviceCopyDeviceInfo", "10.13")
	registerFunc(&_colorSyncDeviceSetCustomProfiles, &_colorSyncDeviceSetCustomProfilesErr, frameworkHandle, "ColorSyncDeviceSetCustomProfiles", "10.13")
	registerFunc(&_colorSyncIterateDeviceProfiles, &_colorSyncIterateDeviceProfilesErr, frameworkHandle, "ColorSyncIterateDeviceProfiles", "10.13")
	registerFunc(&_colorSyncIterateInstalledCMMs, &_colorSyncIterateInstalledCMMsErr, frameworkHandle, "ColorSyncIterateInstalledCMMs", "10.13")
	registerFunc(&_colorSyncIterateInstalledProfiles, &_colorSyncIterateInstalledProfilesErr, frameworkHandle, "ColorSyncIterateInstalledProfiles", "10.13")
	registerFunc(&_colorSyncIterateInstalledProfilesWithOptions, &_colorSyncIterateInstalledProfilesWithOptionsErr, frameworkHandle, "ColorSyncIterateInstalledProfilesWithOptions", "10.13")
	registerFunc(&_colorSyncProfileContainsTag, &_colorSyncProfileContainsTagErr, frameworkHandle, "ColorSyncProfileContainsTag", "10.13")
	registerFunc(&_colorSyncProfileCopyData, &_colorSyncProfileCopyDataErr, frameworkHandle, "ColorSyncProfileCopyData", "10.13")
	registerFunc(&_colorSyncProfileCopyDescriptionString, &_colorSyncProfileCopyDescriptionStringErr, frameworkHandle, "ColorSyncProfileCopyDescriptionString", "10.13")
	registerFunc(&_colorSyncProfileCopyHeader, &_colorSyncProfileCopyHeaderErr, frameworkHandle, "ColorSyncProfileCopyHeader", "10.13")
	registerFunc(&_colorSyncProfileCopyTag, &_colorSyncProfileCopyTagErr, frameworkHandle, "ColorSyncProfileCopyTag", "10.13")
	registerFunc(&_colorSyncProfileCopyTagSignatures, &_colorSyncProfileCopyTagSignaturesErr, frameworkHandle, "ColorSyncProfileCopyTagSignatures", "10.13")
	registerFunc(&_colorSyncProfileCreate, &_colorSyncProfileCreateErr, frameworkHandle, "ColorSyncProfileCreate", "10.13")
	registerFunc(&_colorSyncProfileCreateDeviceProfile, &_colorSyncProfileCreateDeviceProfileErr, frameworkHandle, "ColorSyncProfileCreateDeviceProfile", "10.13")
	registerFunc(&_colorSyncProfileCreateDisplayTransferTablesFromVCGT, &_colorSyncProfileCreateDisplayTransferTablesFromVCGTErr, frameworkHandle, "ColorSyncProfileCreateDisplayTransferTablesFromVCGT", "10.13")
	registerFunc(&_colorSyncProfileCreateLink, &_colorSyncProfileCreateLinkErr, frameworkHandle, "ColorSyncProfileCreateLink", "10.13")
	registerFunc(&_colorSyncProfileCreateMutable, &_colorSyncProfileCreateMutableErr, frameworkHandle, "ColorSyncProfileCreateMutable", "10.13")
	registerFunc(&_colorSyncProfileCreateMutableCopy, &_colorSyncProfileCreateMutableCopyErr, frameworkHandle, "ColorSyncProfileCreateMutableCopy", "10.13")
	registerFunc(&_colorSyncProfileCreateWithDisplayID, &_colorSyncProfileCreateWithDisplayIDErr, frameworkHandle, "ColorSyncProfileCreateWithDisplayID", "10.13")
	registerFunc(&_colorSyncProfileCreateWithName, &_colorSyncProfileCreateWithNameErr, frameworkHandle, "ColorSyncProfileCreateWithName", "10.13")
	registerFunc(&_colorSyncProfileCreateWithURL, &_colorSyncProfileCreateWithURLErr, frameworkHandle, "ColorSyncProfileCreateWithURL", "10.13")
	registerFunc(&_colorSyncProfileCreateWithURLAndOptions, &_colorSyncProfileCreateWithURLAndOptionsErr, frameworkHandle, "ColorSyncProfileCreateWithURLAndOptions", "10.13")
	registerFunc(&_colorSyncProfileEstimateGamma, &_colorSyncProfileEstimateGammaErr, frameworkHandle, "ColorSyncProfileEstimateGamma", "10.13")
	registerFunc(&_colorSyncProfileEstimateGammaWithDisplayID, &_colorSyncProfileEstimateGammaWithDisplayIDErr, frameworkHandle, "ColorSyncProfileEstimateGammaWithDisplayID", "10.13")
	registerFunc(&_colorSyncProfileGetDisplayTransferFormulaFromVCGT, &_colorSyncProfileGetDisplayTransferFormulaFromVCGTErr, frameworkHandle, "ColorSyncProfileGetDisplayTransferFormulaFromVCGT", "10.13")
	registerFunc(&_colorSyncProfileGetMD5, &_colorSyncProfileGetMD5Err, frameworkHandle, "ColorSyncProfileGetMD5", "10.13")
	registerFunc(&_colorSyncProfileGetTypeID, &_colorSyncProfileGetTypeIDErr, frameworkHandle, "ColorSyncProfileGetTypeID", "10.13")
	registerFunc(&_colorSyncProfileGetURL, &_colorSyncProfileGetURLErr, frameworkHandle, "ColorSyncProfileGetURL", "10.13")
	registerFunc(&_colorSyncProfileInstall, &_colorSyncProfileInstallErr, frameworkHandle, "ColorSyncProfileInstall", "10.13")
	registerFunc(&_colorSyncProfileIsHLGBased, &_colorSyncProfileIsHLGBasedErr, frameworkHandle, "ColorSyncProfileIsHLGBased", "10.13")
	registerFunc(&_colorSyncProfileIsMatrixBased, &_colorSyncProfileIsMatrixBasedErr, frameworkHandle, "ColorSyncProfileIsMatrixBased", "10.13")
	registerFunc(&_colorSyncProfileIsPQBased, &_colorSyncProfileIsPQBasedErr, frameworkHandle, "ColorSyncProfileIsPQBased", "10.13")
	registerFunc(&_colorSyncProfileIsWideGamut, &_colorSyncProfileIsWideGamutErr, frameworkHandle, "ColorSyncProfileIsWideGamut", "10.13")
	registerFunc(&_colorSyncProfileRemoveTag, &_colorSyncProfileRemoveTagErr, frameworkHandle, "ColorSyncProfileRemoveTag", "10.13")
	registerFunc(&_colorSyncProfileSetHeader, &_colorSyncProfileSetHeaderErr, frameworkHandle, "ColorSyncProfileSetHeader", "10.13")
	registerFunc(&_colorSyncProfileSetTag, &_colorSyncProfileSetTagErr, frameworkHandle, "ColorSyncProfileSetTag", "10.13")
	registerFunc(&_colorSyncProfileUninstall, &_colorSyncProfileUninstallErr, frameworkHandle, "ColorSyncProfileUninstall", "10.13")
	registerFunc(&_colorSyncProfileVerify, &_colorSyncProfileVerifyErr, frameworkHandle, "ColorSyncProfileVerify", "10.13")
	registerFunc(&_colorSyncRegisterDevice, &_colorSyncRegisterDeviceErr, frameworkHandle, "ColorSyncRegisterDevice", "10.13")
	registerFunc(&_colorSyncTransformConvert, &_colorSyncTransformConvertErr, frameworkHandle, "ColorSyncTransformConvert", "10.13")
	registerFunc(&_colorSyncTransformCopyProperty, &_colorSyncTransformCopyPropertyErr, frameworkHandle, "ColorSyncTransformCopyProperty", "10.13")
	registerFunc(&_colorSyncTransformCreate, &_colorSyncTransformCreateErr, frameworkHandle, "ColorSyncTransformCreate", "10.13")
	registerFunc(&_colorSyncTransformGetProfileSequence, &_colorSyncTransformGetProfileSequenceErr, frameworkHandle, "ColorSyncTransformGetProfileSequence", "10.13")
	registerFunc(&_colorSyncTransformGetTypeID, &_colorSyncTransformGetTypeIDErr, frameworkHandle, "ColorSyncTransformGetTypeID", "10.13")
	registerFunc(&_colorSyncTransformSetProperty, &_colorSyncTransformSetPropertyErr, frameworkHandle, "ColorSyncTransformSetProperty", "10.13")
	registerFunc(&_colorSyncUnregisterDevice, &_colorSyncUnregisterDeviceErr, frameworkHandle, "ColorSyncUnregisterDevice", "10.13")
}
