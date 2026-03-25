// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"fmt"
	"os"
	"github.com/ebitengine/purego"
)

// registerFunc resolves a framework symbol and registers it as a Go function.
// If the symbol is not found, a warning is printed and the function pointer is left nil.
func registerFunc(fptr any, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: appleneuralengine: symbol %s not found\n", name)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "warning: appleneuralengine: symbol %s registration skipped: %v\n", name, r)
		}
	}()
	purego.RegisterFunc(fptr, sym)
}

// registerSymbol resolves a framework symbol and stores its raw address.
func registerSymbol(dst *uintptr, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: appleneuralengine: symbol %s not found\n", name)
		return
	}
	*dst = sym
}

var _aNEGetValidateNetworkSupportedVersionSymbol uintptr

// ANEGetValidateNetworkSupportedVersion has an opaque C signature in discovered metadata.
// Call ANEGetValidateNetworkSupportedVersionSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/ANEGetValidateNetworkSupportedVersion
func ANEGetValidateNetworkSupportedVersion() {
	panic("appleneuralengine: symbol ANEGetValidateNetworkSupportedVersion has opaque signature; use ANEGetValidateNetworkSupportedVersionSymbol() and a typed manual wrapper")
}

// ANEGetValidateNetworkSupportedVersionSymbol returns the raw symbol address for ANEGetValidateNetworkSupportedVersion.
func ANEGetValidateNetworkSupportedVersionSymbol() uintptr {
	if _aNEGetValidateNetworkSupportedVersionSymbol == 0 {
		return 0
	}
	return _aNEGetValidateNetworkSupportedVersionSymbol
}

var _aNEValidateMILNetworkOnHostSymbol uintptr

// ANEValidateMILNetworkOnHost has an opaque C signature in discovered metadata.
// Call ANEValidateMILNetworkOnHostSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/ANEValidateMILNetworkOnHost
func ANEValidateMILNetworkOnHost() {
	panic("appleneuralengine: symbol ANEValidateMILNetworkOnHost has opaque signature; use ANEValidateMILNetworkOnHostSymbol() and a typed manual wrapper")
}

// ANEValidateMILNetworkOnHostSymbol returns the raw symbol address for ANEValidateMILNetworkOnHost.
func ANEValidateMILNetworkOnHostSymbol() uintptr {
	if _aNEValidateMILNetworkOnHostSymbol == 0 {
		return 0
	}
	return _aNEValidateMILNetworkOnHostSymbol
}

var _aNEValidateMLIRNetworkOnHostSymbol uintptr

// ANEValidateMLIRNetworkOnHost has an opaque C signature in discovered metadata.
// Call ANEValidateMLIRNetworkOnHostSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/ANEValidateMLIRNetworkOnHost
func ANEValidateMLIRNetworkOnHost() {
	panic("appleneuralengine: symbol ANEValidateMLIRNetworkOnHost has opaque signature; use ANEValidateMLIRNetworkOnHostSymbol() and a typed manual wrapper")
}

// ANEValidateMLIRNetworkOnHostSymbol returns the raw symbol address for ANEValidateMLIRNetworkOnHost.
func ANEValidateMLIRNetworkOnHostSymbol() uintptr {
	if _aNEValidateMLIRNetworkOnHostSymbol == 0 {
		return 0
	}
	return _aNEValidateMLIRNetworkOnHostSymbol
}

var _aNEValidateNetworkCreateSymbol uintptr

// ANEValidateNetworkCreate has an opaque C signature in discovered metadata.
// Call ANEValidateNetworkCreateSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/ANEValidateNetworkCreate
func ANEValidateNetworkCreate() {
	panic("appleneuralengine: symbol ANEValidateNetworkCreate has opaque signature; use ANEValidateNetworkCreateSymbol() and a typed manual wrapper")
}

// ANEValidateNetworkCreateSymbol returns the raw symbol address for ANEValidateNetworkCreate.
func ANEValidateNetworkCreateSymbol() uintptr {
	if _aNEValidateNetworkCreateSymbol == 0 {
		return 0
	}
	return _aNEValidateNetworkCreateSymbol
}

var _aNEValidateNetworkCreateGenerateFailedGlobalIdentifierSymbol uintptr

// ANEValidateNetworkCreateGenerateFailedGlobalIdentifier has an opaque C signature in discovered metadata.
// Call ANEValidateNetworkCreateGenerateFailedGlobalIdentifierSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/ANEValidateNetworkCreateGenerateFailedGlobalIdentifier
func ANEValidateNetworkCreateGenerateFailedGlobalIdentifier() {
	panic("appleneuralengine: symbol ANEValidateNetworkCreateGenerateFailedGlobalIdentifier has opaque signature; use ANEValidateNetworkCreateGenerateFailedGlobalIdentifierSymbol() and a typed manual wrapper")
}

// ANEValidateNetworkCreateGenerateFailedGlobalIdentifierSymbol returns the raw symbol address for ANEValidateNetworkCreateGenerateFailedGlobalIdentifier.
func ANEValidateNetworkCreateGenerateFailedGlobalIdentifierSymbol() uintptr {
	if _aNEValidateNetworkCreateGenerateFailedGlobalIdentifierSymbol == 0 {
		return 0
	}
	return _aNEValidateNetworkCreateGenerateFailedGlobalIdentifierSymbol
}

var _aNEValidateNetworkCreateVMHostSymbol uintptr

// ANEValidateNetworkCreateVMHost has an opaque C signature in discovered metadata.
// Call ANEValidateNetworkCreateVMHostSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/ANEValidateNetworkCreateVMHost
func ANEValidateNetworkCreateVMHost() {
	panic("appleneuralengine: symbol ANEValidateNetworkCreateVMHost has opaque signature; use ANEValidateNetworkCreateVMHostSymbol() and a typed manual wrapper")
}

// ANEValidateNetworkCreateVMHostSymbol returns the raw symbol address for ANEValidateNetworkCreateVMHost.
func ANEValidateNetworkCreateVMHostSymbol() uintptr {
	if _aNEValidateNetworkCreateVMHostSymbol == 0 {
		return 0
	}
	return _aNEValidateNetworkCreateVMHostSymbol
}

var _appleNeuralEngineVersionNumberSymbol uintptr

// AppleNeuralEngineVersionNumber has an opaque C signature in discovered metadata.
// Call AppleNeuralEngineVersionNumberSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/AppleNeuralEngineVersionNumber
func AppleNeuralEngineVersionNumber() {
	panic("appleneuralengine: symbol AppleNeuralEngineVersionNumber has opaque signature; use AppleNeuralEngineVersionNumberSymbol() and a typed manual wrapper")
}

// AppleNeuralEngineVersionNumberSymbol returns the raw symbol address for AppleNeuralEngineVersionNumber.
func AppleNeuralEngineVersionNumberSymbol() uintptr {
	if _appleNeuralEngineVersionNumberSymbol == 0 {
		return 0
	}
	return _appleNeuralEngineVersionNumberSymbol
}

var _appleNeuralEngineVersionStringSymbol uintptr

// AppleNeuralEngineVersionString has an opaque C signature in discovered metadata.
// Call AppleNeuralEngineVersionStringSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/AppleNeuralEngineVersionString
func AppleNeuralEngineVersionString() {
	panic("appleneuralengine: symbol AppleNeuralEngineVersionString has opaque signature; use AppleNeuralEngineVersionStringSymbol() and a typed manual wrapper")
}

// AppleNeuralEngineVersionStringSymbol returns the raw symbol address for AppleNeuralEngineVersionString.
func AppleNeuralEngineVersionStringSymbol() uintptr {
	if _appleNeuralEngineVersionStringSymbol == 0 {
		return 0
	}
	return _appleNeuralEngineVersionStringSymbol
}

var __ANEDaemonInterfaceSymbol uintptr

// _ANEDaemonInterface has an opaque C signature in discovered metadata.
// Call _ANEDaemonInterfaceSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDaemonInterface
func _ANEDaemonInterface() {
	panic("appleneuralengine: symbol _ANEDaemonInterface has opaque signature; use _ANEDaemonInterfaceSymbol() and a typed manual wrapper")
}

// _ANEDaemonInterfaceSymbol returns the raw symbol address for _ANEDaemonInterface.
func _ANEDaemonInterfaceSymbol() uintptr {
	if __ANEDaemonInterfaceSymbol == 0 {
		return 0
	}
	return __ANEDaemonInterfaceSymbol
}

var __ANEDaemonInterfacePrivateSymbol uintptr

// _ANEDaemonInterfacePrivate has an opaque C signature in discovered metadata.
// Call _ANEDaemonInterfacePrivateSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDaemonInterfacePrivate
func _ANEDaemonInterfacePrivate() {
	panic("appleneuralengine: symbol _ANEDaemonInterfacePrivate has opaque signature; use _ANEDaemonInterfacePrivateSymbol() and a typed manual wrapper")
}

// _ANEDaemonInterfacePrivateSymbol returns the raw symbol address for _ANEDaemonInterfacePrivate.
func _ANEDaemonInterfacePrivateSymbol() uintptr {
	if __ANEDaemonInterfacePrivateSymbol == 0 {
		return 0
	}
	return __ANEDaemonInterfacePrivateSymbol
}

var _kANEModelKeyAllSegmentsValueSymbol uintptr

// KANEModelKeyAllSegmentsValue has an opaque C signature in discovered metadata.
// Call KANEModelKeyAllSegmentsValueSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/kANEModelKeyAllSegmentsValue
func KANEModelKeyAllSegmentsValue() {
	panic("appleneuralengine: symbol kANEModelKeyAllSegmentsValue has opaque signature; use KANEModelKeyAllSegmentsValueSymbol() and a typed manual wrapper")
}

// KANEModelKeyAllSegmentsValueSymbol returns the raw symbol address for kANEModelKeyAllSegmentsValue.
func KANEModelKeyAllSegmentsValueSymbol() uintptr {
	if _kANEModelKeyAllSegmentsValueSymbol == 0 {
		return 0
	}
	return _kANEModelKeyAllSegmentsValueSymbol
}

var _kANEModelKeyEspressoTranslationOptionsSymbol uintptr

// KANEModelKeyEspressoTranslationOptions has an opaque C signature in discovered metadata.
// Call KANEModelKeyEspressoTranslationOptionsSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/kANEModelKeyEspressoTranslationOptions
func KANEModelKeyEspressoTranslationOptions() {
	panic("appleneuralengine: symbol kANEModelKeyEspressoTranslationOptions has opaque signature; use KANEModelKeyEspressoTranslationOptionsSymbol() and a typed manual wrapper")
}

// KANEModelKeyEspressoTranslationOptionsSymbol returns the raw symbol address for kANEModelKeyEspressoTranslationOptions.
func KANEModelKeyEspressoTranslationOptionsSymbol() uintptr {
	if _kANEModelKeyEspressoTranslationOptionsSymbol == 0 {
		return 0
	}
	return _kANEModelKeyEspressoTranslationOptionsSymbol
}

func init() {
	if frameworkHandle == 0 {
		return
	}
		registerSymbol(&_aNEGetValidateNetworkSupportedVersionSymbol, frameworkHandle, "ANEGetValidateNetworkSupportedVersion")
		registerSymbol(&_aNEValidateMILNetworkOnHostSymbol, frameworkHandle, "ANEValidateMILNetworkOnHost")
		registerSymbol(&_aNEValidateMLIRNetworkOnHostSymbol, frameworkHandle, "ANEValidateMLIRNetworkOnHost")
		registerSymbol(&_aNEValidateNetworkCreateSymbol, frameworkHandle, "ANEValidateNetworkCreate")
		registerSymbol(&_aNEValidateNetworkCreateGenerateFailedGlobalIdentifierSymbol, frameworkHandle, "ANEValidateNetworkCreateGenerateFailedGlobalIdentifier")
		registerSymbol(&_aNEValidateNetworkCreateVMHostSymbol, frameworkHandle, "ANEValidateNetworkCreateVMHost")
		registerSymbol(&_appleNeuralEngineVersionNumberSymbol, frameworkHandle, "AppleNeuralEngineVersionNumber")
		registerSymbol(&_appleNeuralEngineVersionStringSymbol, frameworkHandle, "AppleNeuralEngineVersionString")
		registerSymbol(&__ANEDaemonInterfaceSymbol, frameworkHandle, "_ANEDaemonInterface")
		registerSymbol(&__ANEDaemonInterfacePrivateSymbol, frameworkHandle, "_ANEDaemonInterfacePrivate")
		registerSymbol(&_kANEModelKeyAllSegmentsValueSymbol, frameworkHandle, "kANEModelKeyAllSegmentsValue")
		registerSymbol(&_kANEModelKeyEspressoTranslationOptionsSymbol, frameworkHandle, "kANEModelKeyEspressoTranslationOptions")
	}

