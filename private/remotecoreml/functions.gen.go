// Code generated from Apple documentation for remotecoreml. DO NOT EDIT.

package remotecoreml

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
		fmt.Fprintf(os.Stderr, "warning: remotecoreml: symbol %s not found\n", name)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "warning: remotecoreml: symbol %s registration skipped: %v\n", name, r)
		}
	}()
	purego.RegisterFunc(fptr, sym)
}

// registerSymbol resolves a framework symbol and stores its raw address.
func registerSymbol(dst *uintptr, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: remotecoreml: symbol %s not found\n", name)
		return
	}
	*dst = sym
}


var _remoteCoreMLVersionNumberSymbol uintptr

// RemoteCoreMLVersionNumber has an opaque C signature in discovered metadata.
// Call RemoteCoreMLVersionNumberSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/RemoteCoreML/RemoteCoreMLVersionNumber
func RemoteCoreMLVersionNumber() {
	panic("remotecoreml: symbol RemoteCoreMLVersionNumber has opaque signature; use RemoteCoreMLVersionNumberSymbol() and a typed manual wrapper")
}

// RemoteCoreMLVersionNumberSymbol returns the raw symbol address for RemoteCoreMLVersionNumber.
func RemoteCoreMLVersionNumberSymbol() uintptr {
	if _remoteCoreMLVersionNumberSymbol == 0 {
		panic("remotecoreml: symbol RemoteCoreMLVersionNumber not loaded")
	}
	return _remoteCoreMLVersionNumberSymbol
}


var _remoteCoreMLVersionStringSymbol uintptr

// RemoteCoreMLVersionString has an opaque C signature in discovered metadata.
// Call RemoteCoreMLVersionStringSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/RemoteCoreML/RemoteCoreMLVersionString
func RemoteCoreMLVersionString() {
	panic("remotecoreml: symbol RemoteCoreMLVersionString has opaque signature; use RemoteCoreMLVersionStringSymbol() and a typed manual wrapper")
}

// RemoteCoreMLVersionStringSymbol returns the raw symbol address for RemoteCoreMLVersionString.
func RemoteCoreMLVersionStringSymbol() uintptr {
	if _remoteCoreMLVersionStringSymbol == 0 {
		panic("remotecoreml: symbol RemoteCoreMLVersionString not loaded")
	}
	return _remoteCoreMLVersionStringSymbol
}


var _g_inbound_connectionSymbol uintptr

// G_inbound_connection has an opaque C signature in discovered metadata.
// Call G_inbound_connectionSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/RemoteCoreML/g_inbound_connection
func G_inbound_connection() {
	panic("remotecoreml: symbol g_inbound_connection has opaque signature; use G_inbound_connectionSymbol() and a typed manual wrapper")
}

// G_inbound_connectionSymbol returns the raw symbol address for g_inbound_connection.
func G_inbound_connectionSymbol() uintptr {
	if _g_inbound_connectionSymbol == 0 {
		panic("remotecoreml: symbol g_inbound_connection not loaded")
	}
	return _g_inbound_connectionSymbol
}


var _kMLAckFailNetworkHeaderEncodingSymbol uintptr

// KMLAckFailNetworkHeaderEncoding has an opaque C signature in discovered metadata.
// Call KMLAckFailNetworkHeaderEncodingSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/RemoteCoreML/kMLAckFailNetworkHeaderEncoding
func KMLAckFailNetworkHeaderEncoding() {
	panic("remotecoreml: symbol kMLAckFailNetworkHeaderEncoding has opaque signature; use KMLAckFailNetworkHeaderEncodingSymbol() and a typed manual wrapper")
}

// KMLAckFailNetworkHeaderEncodingSymbol returns the raw symbol address for kMLAckFailNetworkHeaderEncoding.
func KMLAckFailNetworkHeaderEncodingSymbol() uintptr {
	if _kMLAckFailNetworkHeaderEncodingSymbol == 0 {
		panic("remotecoreml: symbol kMLAckFailNetworkHeaderEncoding not loaded")
	}
	return _kMLAckFailNetworkHeaderEncodingSymbol
}


var _kMLAckSuccessNetworkHeaderEncodingSymbol uintptr

// KMLAckSuccessNetworkHeaderEncoding has an opaque C signature in discovered metadata.
// Call KMLAckSuccessNetworkHeaderEncodingSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/RemoteCoreML/kMLAckSuccessNetworkHeaderEncoding
func KMLAckSuccessNetworkHeaderEncoding() {
	panic("remotecoreml: symbol kMLAckSuccessNetworkHeaderEncoding has opaque signature; use KMLAckSuccessNetworkHeaderEncodingSymbol() and a typed manual wrapper")
}

// KMLAckSuccessNetworkHeaderEncodingSymbol returns the raw symbol address for kMLAckSuccessNetworkHeaderEncoding.
func KMLAckSuccessNetworkHeaderEncodingSymbol() uintptr {
	if _kMLAckSuccessNetworkHeaderEncodingSymbol == 0 {
		panic("remotecoreml: symbol kMLAckSuccessNetworkHeaderEncoding not loaded")
	}
	return _kMLAckSuccessNetworkHeaderEncodingSymbol
}


var _kMLCustomNetworkHeaderEncodingSymbol uintptr

// KMLCustomNetworkHeaderEncoding has an opaque C signature in discovered metadata.
// Call KMLCustomNetworkHeaderEncodingSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/RemoteCoreML/kMLCustomNetworkHeaderEncoding
func KMLCustomNetworkHeaderEncoding() {
	panic("remotecoreml: symbol kMLCustomNetworkHeaderEncoding has opaque signature; use KMLCustomNetworkHeaderEncodingSymbol() and a typed manual wrapper")
}

// KMLCustomNetworkHeaderEncodingSymbol returns the raw symbol address for kMLCustomNetworkHeaderEncoding.
func KMLCustomNetworkHeaderEncodingSymbol() uintptr {
	if _kMLCustomNetworkHeaderEncodingSymbol == 0 {
		panic("remotecoreml: symbol kMLCustomNetworkHeaderEncoding not loaded")
	}
	return _kMLCustomNetworkHeaderEncodingSymbol
}


var _kMLErrorNetworkHeaderEncodingSymbol uintptr

// KMLErrorNetworkHeaderEncoding has an opaque C signature in discovered metadata.
// Call KMLErrorNetworkHeaderEncodingSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/RemoteCoreML/kMLErrorNetworkHeaderEncoding
func KMLErrorNetworkHeaderEncoding() {
	panic("remotecoreml: symbol kMLErrorNetworkHeaderEncoding has opaque signature; use KMLErrorNetworkHeaderEncodingSymbol() and a typed manual wrapper")
}

// KMLErrorNetworkHeaderEncodingSymbol returns the raw symbol address for kMLErrorNetworkHeaderEncoding.
func KMLErrorNetworkHeaderEncodingSymbol() uintptr {
	if _kMLErrorNetworkHeaderEncodingSymbol == 0 {
		panic("remotecoreml: symbol kMLErrorNetworkHeaderEncoding not loaded")
	}
	return _kMLErrorNetworkHeaderEncodingSymbol
}


var _kMLIncomingDataNetworkHeaderEncodingSymbol uintptr

// KMLIncomingDataNetworkHeaderEncoding has an opaque C signature in discovered metadata.
// Call KMLIncomingDataNetworkHeaderEncodingSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/RemoteCoreML/kMLIncomingDataNetworkHeaderEncoding
func KMLIncomingDataNetworkHeaderEncoding() {
	panic("remotecoreml: symbol kMLIncomingDataNetworkHeaderEncoding has opaque signature; use KMLIncomingDataNetworkHeaderEncodingSymbol() and a typed manual wrapper")
}

// KMLIncomingDataNetworkHeaderEncodingSymbol returns the raw symbol address for kMLIncomingDataNetworkHeaderEncoding.
func KMLIncomingDataNetworkHeaderEncodingSymbol() uintptr {
	if _kMLIncomingDataNetworkHeaderEncodingSymbol == 0 {
		panic("remotecoreml: symbol kMLIncomingDataNetworkHeaderEncoding not loaded")
	}
	return _kMLIncomingDataNetworkHeaderEncodingSymbol
}


var _kMLLoadNetworkHeaderEncodingSymbol uintptr

// KMLLoadNetworkHeaderEncoding has an opaque C signature in discovered metadata.
// Call KMLLoadNetworkHeaderEncodingSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/RemoteCoreML/kMLLoadNetworkHeaderEncoding
func KMLLoadNetworkHeaderEncoding() {
	panic("remotecoreml: symbol kMLLoadNetworkHeaderEncoding has opaque signature; use KMLLoadNetworkHeaderEncodingSymbol() and a typed manual wrapper")
}

// KMLLoadNetworkHeaderEncodingSymbol returns the raw symbol address for kMLLoadNetworkHeaderEncoding.
func KMLLoadNetworkHeaderEncodingSymbol() uintptr {
	if _kMLLoadNetworkHeaderEncodingSymbol == 0 {
		panic("remotecoreml: symbol kMLLoadNetworkHeaderEncoding not loaded")
	}
	return _kMLLoadNetworkHeaderEncodingSymbol
}


var _kMLNetworkFamilyKeySymbol uintptr

// KMLNetworkFamilyKey has an opaque C signature in discovered metadata.
// Call KMLNetworkFamilyKeySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/RemoteCoreML/kMLNetworkFamilyKey
func KMLNetworkFamilyKey() {
	panic("remotecoreml: symbol kMLNetworkFamilyKey has opaque signature; use KMLNetworkFamilyKeySymbol() and a typed manual wrapper")
}

// KMLNetworkFamilyKeySymbol returns the raw symbol address for kMLNetworkFamilyKey.
func KMLNetworkFamilyKeySymbol() uintptr {
	if _kMLNetworkFamilyKeySymbol == 0 {
		panic("remotecoreml: symbol kMLNetworkFamilyKey not loaded")
	}
	return _kMLNetworkFamilyKeySymbol
}


var _kMLNetworkLocalAddrKeySymbol uintptr

// KMLNetworkLocalAddrKey has an opaque C signature in discovered metadata.
// Call KMLNetworkLocalAddrKeySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/RemoteCoreML/kMLNetworkLocalAddrKey
func KMLNetworkLocalAddrKey() {
	panic("remotecoreml: symbol kMLNetworkLocalAddrKey has opaque signature; use KMLNetworkLocalAddrKeySymbol() and a typed manual wrapper")
}

// KMLNetworkLocalAddrKeySymbol returns the raw symbol address for kMLNetworkLocalAddrKey.
func KMLNetworkLocalAddrKeySymbol() uintptr {
	if _kMLNetworkLocalAddrKeySymbol == 0 {
		panic("remotecoreml: symbol kMLNetworkLocalAddrKey not loaded")
	}
	return _kMLNetworkLocalAddrKeySymbol
}


var _kMLNetworkLocalPortKeySymbol uintptr

// KMLNetworkLocalPortKey has an opaque C signature in discovered metadata.
// Call KMLNetworkLocalPortKeySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/RemoteCoreML/kMLNetworkLocalPortKey
func KMLNetworkLocalPortKey() {
	panic("remotecoreml: symbol kMLNetworkLocalPortKey has opaque signature; use KMLNetworkLocalPortKeySymbol() and a typed manual wrapper")
}

// KMLNetworkLocalPortKeySymbol returns the raw symbol address for kMLNetworkLocalPortKey.
func KMLNetworkLocalPortKeySymbol() uintptr {
	if _kMLNetworkLocalPortKeySymbol == 0 {
		panic("remotecoreml: symbol kMLNetworkLocalPortKey not loaded")
	}
	return _kMLNetworkLocalPortKeySymbol
}


var _kMLNetworkNameIdentifierKeySymbol uintptr

// KMLNetworkNameIdentifierKey has an opaque C signature in discovered metadata.
// Call KMLNetworkNameIdentifierKeySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/RemoteCoreML/kMLNetworkNameIdentifierKey
func KMLNetworkNameIdentifierKey() {
	panic("remotecoreml: symbol kMLNetworkNameIdentifierKey has opaque signature; use KMLNetworkNameIdentifierKeySymbol() and a typed manual wrapper")
}

// KMLNetworkNameIdentifierKeySymbol returns the raw symbol address for kMLNetworkNameIdentifierKey.
func KMLNetworkNameIdentifierKeySymbol() uintptr {
	if _kMLNetworkNameIdentifierKeySymbol == 0 {
		panic("remotecoreml: symbol kMLNetworkNameIdentifierKey not loaded")
	}
	return _kMLNetworkNameIdentifierKeySymbol
}


var _kMLNetworkPortNumberKeySymbol uintptr

// KMLNetworkPortNumberKey has an opaque C signature in discovered metadata.
// Call KMLNetworkPortNumberKeySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/RemoteCoreML/kMLNetworkPortNumberKey
func KMLNetworkPortNumberKey() {
	panic("remotecoreml: symbol kMLNetworkPortNumberKey has opaque signature; use KMLNetworkPortNumberKeySymbol() and a typed manual wrapper")
}

// KMLNetworkPortNumberKeySymbol returns the raw symbol address for kMLNetworkPortNumberKey.
func KMLNetworkPortNumberKeySymbol() uintptr {
	if _kMLNetworkPortNumberKeySymbol == 0 {
		panic("remotecoreml: symbol kMLNetworkPortNumberKey not loaded")
	}
	return _kMLNetworkPortNumberKeySymbol
}


var _kMLNetworkPskKeySymbol uintptr

// KMLNetworkPskKey has an opaque C signature in discovered metadata.
// Call KMLNetworkPskKeySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/RemoteCoreML/kMLNetworkPskKey
func KMLNetworkPskKey() {
	panic("remotecoreml: symbol kMLNetworkPskKey has opaque signature; use KMLNetworkPskKeySymbol() and a typed manual wrapper")
}

// KMLNetworkPskKeySymbol returns the raw symbol address for kMLNetworkPskKey.
func KMLNetworkPskKeySymbol() uintptr {
	if _kMLNetworkPskKeySymbol == 0 {
		panic("remotecoreml: symbol kMLNetworkPskKey not loaded")
	}
	return _kMLNetworkPskKeySymbol
}


var _kMLNetworkUseAWDLKeySymbol uintptr

// KMLNetworkUseAWDLKey has an opaque C signature in discovered metadata.
// Call KMLNetworkUseAWDLKeySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/RemoteCoreML/kMLNetworkUseAWDLKey
func KMLNetworkUseAWDLKey() {
	panic("remotecoreml: symbol kMLNetworkUseAWDLKey has opaque signature; use KMLNetworkUseAWDLKeySymbol() and a typed manual wrapper")
}

// KMLNetworkUseAWDLKeySymbol returns the raw symbol address for kMLNetworkUseAWDLKey.
func KMLNetworkUseAWDLKeySymbol() uintptr {
	if _kMLNetworkUseAWDLKeySymbol == 0 {
		panic("remotecoreml: symbol kMLNetworkUseAWDLKey not loaded")
	}
	return _kMLNetworkUseAWDLKeySymbol
}


var _kMLNetworkUseBonjourKeySymbol uintptr

// KMLNetworkUseBonjourKey has an opaque C signature in discovered metadata.
// Call KMLNetworkUseBonjourKeySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/RemoteCoreML/kMLNetworkUseBonjourKey
func KMLNetworkUseBonjourKey() {
	panic("remotecoreml: symbol kMLNetworkUseBonjourKey has opaque signature; use KMLNetworkUseBonjourKeySymbol() and a typed manual wrapper")
}

// KMLNetworkUseBonjourKeySymbol returns the raw symbol address for kMLNetworkUseBonjourKey.
func KMLNetworkUseBonjourKeySymbol() uintptr {
	if _kMLNetworkUseBonjourKeySymbol == 0 {
		panic("remotecoreml: symbol kMLNetworkUseBonjourKey not loaded")
	}
	return _kMLNetworkUseBonjourKeySymbol
}


var _kMLNetworkUseTLSKeySymbol uintptr

// KMLNetworkUseTLSKey has an opaque C signature in discovered metadata.
// Call KMLNetworkUseTLSKeySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/RemoteCoreML/kMLNetworkUseTLSKey
func KMLNetworkUseTLSKey() {
	panic("remotecoreml: symbol kMLNetworkUseTLSKey has opaque signature; use KMLNetworkUseTLSKeySymbol() and a typed manual wrapper")
}

// KMLNetworkUseTLSKeySymbol returns the raw symbol address for kMLNetworkUseTLSKey.
func KMLNetworkUseTLSKeySymbol() uintptr {
	if _kMLNetworkUseTLSKeySymbol == 0 {
		panic("remotecoreml: symbol kMLNetworkUseTLSKey not loaded")
	}
	return _kMLNetworkUseTLSKeySymbol
}


var _kMLNetworkUseUDPKeySymbol uintptr

// KMLNetworkUseUDPKey has an opaque C signature in discovered metadata.
// Call KMLNetworkUseUDPKeySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/RemoteCoreML/kMLNetworkUseUDPKey
func KMLNetworkUseUDPKey() {
	panic("remotecoreml: symbol kMLNetworkUseUDPKey has opaque signature; use KMLNetworkUseUDPKeySymbol() and a typed manual wrapper")
}

// KMLNetworkUseUDPKeySymbol returns the raw symbol address for kMLNetworkUseUDPKey.
func KMLNetworkUseUDPKeySymbol() uintptr {
	if _kMLNetworkUseUDPKeySymbol == 0 {
		panic("remotecoreml: symbol kMLNetworkUseUDPKey not loaded")
	}
	return _kMLNetworkUseUDPKeySymbol
}


var _kMLPredictNetworkHeaderEncodingSymbol uintptr

// KMLPredictNetworkHeaderEncoding has an opaque C signature in discovered metadata.
// Call KMLPredictNetworkHeaderEncodingSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/RemoteCoreML/kMLPredictNetworkHeaderEncoding
func KMLPredictNetworkHeaderEncoding() {
	panic("remotecoreml: symbol kMLPredictNetworkHeaderEncoding has opaque signature; use KMLPredictNetworkHeaderEncodingSymbol() and a typed manual wrapper")
}

// KMLPredictNetworkHeaderEncodingSymbol returns the raw symbol address for kMLPredictNetworkHeaderEncoding.
func KMLPredictNetworkHeaderEncodingSymbol() uintptr {
	if _kMLPredictNetworkHeaderEncodingSymbol == 0 {
		panic("remotecoreml: symbol kMLPredictNetworkHeaderEncoding not loaded")
	}
	return _kMLPredictNetworkHeaderEncodingSymbol
}


var _kMLTextNetworkHeaderEncodingSymbol uintptr

// KMLTextNetworkHeaderEncoding has an opaque C signature in discovered metadata.
// Call KMLTextNetworkHeaderEncodingSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/RemoteCoreML/kMLTextNetworkHeaderEncoding
func KMLTextNetworkHeaderEncoding() {
	panic("remotecoreml: symbol kMLTextNetworkHeaderEncoding has opaque signature; use KMLTextNetworkHeaderEncodingSymbol() and a typed manual wrapper")
}

// KMLTextNetworkHeaderEncodingSymbol returns the raw symbol address for kMLTextNetworkHeaderEncoding.
func KMLTextNetworkHeaderEncodingSymbol() uintptr {
	if _kMLTextNetworkHeaderEncodingSymbol == 0 {
		panic("remotecoreml: symbol kMLTextNetworkHeaderEncoding not loaded")
	}
	return _kMLTextNetworkHeaderEncodingSymbol
}



func init() {
	if frameworkHandle == 0 {
		return
	}
		registerSymbol(&_remoteCoreMLVersionNumberSymbol, frameworkHandle, "RemoteCoreMLVersionNumber")
		registerSymbol(&_remoteCoreMLVersionStringSymbol, frameworkHandle, "RemoteCoreMLVersionString")
		registerSymbol(&_g_inbound_connectionSymbol, frameworkHandle, "g_inbound_connection")
		registerSymbol(&_kMLAckFailNetworkHeaderEncodingSymbol, frameworkHandle, "kMLAckFailNetworkHeaderEncoding")
		registerSymbol(&_kMLAckSuccessNetworkHeaderEncodingSymbol, frameworkHandle, "kMLAckSuccessNetworkHeaderEncoding")
		registerSymbol(&_kMLCustomNetworkHeaderEncodingSymbol, frameworkHandle, "kMLCustomNetworkHeaderEncoding")
		registerSymbol(&_kMLErrorNetworkHeaderEncodingSymbol, frameworkHandle, "kMLErrorNetworkHeaderEncoding")
		registerSymbol(&_kMLIncomingDataNetworkHeaderEncodingSymbol, frameworkHandle, "kMLIncomingDataNetworkHeaderEncoding")
		registerSymbol(&_kMLLoadNetworkHeaderEncodingSymbol, frameworkHandle, "kMLLoadNetworkHeaderEncoding")
		registerSymbol(&_kMLNetworkFamilyKeySymbol, frameworkHandle, "kMLNetworkFamilyKey")
		registerSymbol(&_kMLNetworkLocalAddrKeySymbol, frameworkHandle, "kMLNetworkLocalAddrKey")
		registerSymbol(&_kMLNetworkLocalPortKeySymbol, frameworkHandle, "kMLNetworkLocalPortKey")
		registerSymbol(&_kMLNetworkNameIdentifierKeySymbol, frameworkHandle, "kMLNetworkNameIdentifierKey")
		registerSymbol(&_kMLNetworkPortNumberKeySymbol, frameworkHandle, "kMLNetworkPortNumberKey")
		registerSymbol(&_kMLNetworkPskKeySymbol, frameworkHandle, "kMLNetworkPskKey")
		registerSymbol(&_kMLNetworkUseAWDLKeySymbol, frameworkHandle, "kMLNetworkUseAWDLKey")
		registerSymbol(&_kMLNetworkUseBonjourKeySymbol, frameworkHandle, "kMLNetworkUseBonjourKey")
		registerSymbol(&_kMLNetworkUseTLSKeySymbol, frameworkHandle, "kMLNetworkUseTLSKey")
		registerSymbol(&_kMLNetworkUseUDPKeySymbol, frameworkHandle, "kMLNetworkUseUDPKey")
		registerSymbol(&_kMLPredictNetworkHeaderEncodingSymbol, frameworkHandle, "kMLPredictNetworkHeaderEncoding")
		registerSymbol(&_kMLTextNetworkHeaderEncodingSymbol, frameworkHandle, "kMLTextNetworkHeaderEncoding")
	}



