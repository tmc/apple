// Code generated from Apple documentation for Network. DO NOT EDIT.

package network

import (
	"fmt"
	"sync"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
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
		return fmt.Sprintf("Network: symbol %s unavailable on this system (introduced in macOS %s)", e.symbol, e.introduced)
	}
	return fmt.Sprintf("Network: symbol %s unavailable on this system", e.symbol)
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
		return fmt.Errorf("Network: symbol %s unavailable because the framework could not be loaded", name)
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
			*errDst = fmt.Errorf("Network: register symbol %s: %v", name, r)
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

type networkAsyncBlockKey struct {
	owner  objc.ID
	setter string
}

var (
	networkAsyncBlockMu sync.Mutex
	networkAsyncBlocks  = make(map[networkAsyncBlockKey]objc.Block)

	_nw_parameters_configure_protocol_default_configurationSymbol uintptr
	_nw_parameters_configure_protocol_default_configurationErr    error
	_nw_parameters_configure_protocol_disableSymbol               uintptr
	_nw_parameters_configure_protocol_disableErr                  error
)

func retainNetworkAsyncBlock(owner objc.ID, setter string, block objc.Block) {
	if owner == 0 || block == 0 {
		return
	}
	key := networkAsyncBlockKey{owner: owner, setter: setter}
	var old objc.Block
	networkAsyncBlockMu.Lock()
	old = networkAsyncBlocks[key]
	networkAsyncBlocks[key] = block
	networkAsyncBlockMu.Unlock()
	if old != 0 {
		old.Release()
	}
}

func clearNetworkAsyncBlock(owner objc.ID, setter string) {
	if owner == 0 {
		return
	}
	key := networkAsyncBlockKey{owner: owner, setter: setter}
	var old objc.Block
	networkAsyncBlockMu.Lock()
	old = networkAsyncBlocks[key]
	delete(networkAsyncBlocks, key)
	networkAsyncBlockMu.Unlock()
	if old != 0 {
		old.Release()
	}
}

func networkProtocolBlockValue(sym uintptr) unsafe.Pointer {
	if sym == 0 {
		return nil
	}
	return *(*unsafe.Pointer)(unsafe.Pointer(sym))
}

var networkCreateSecureTCPForPlain func(configure_tls unsafe.Pointer, configure_tcp unsafe.Pointer) NWParameters
var networkCreateSecureTCPForPlainErr error

func tryNWParametersCreatePlainTCP(configureTCP NWParametersConfigureProtocolBlock) (NWParameters, error) {
	if networkCreateSecureTCPForPlain == nil {
		return *new(NWParameters), symbolCallError("nw_parameters_create_secure_tcp", "10.14", networkCreateSecureTCPForPlainErr)
	}
	if _nw_parameters_configure_protocol_disableSymbol == 0 {
		return *new(NWParameters), symbolCallError("_nw_parameters_configure_protocol_disable", "10.14", _nw_parameters_configure_protocol_disableErr)
	}
	var _block1 unsafe.Pointer
	if configureTCP == nil {
		if _nw_parameters_configure_protocol_default_configurationSymbol == 0 {
			return *new(NWParameters), symbolCallError("_nw_parameters_configure_protocol_default_configuration", "10.14", _nw_parameters_configure_protocol_default_configurationErr)
		}
		_block1 = networkProtocolBlockValue(_nw_parameters_configure_protocol_default_configurationSymbol)
	} else {
		_block1Value := objc.NewBlock(func(_ objc.Block, blockArg0 objc.ID) { configureTCP(objectivec.ObjectFromID(blockArg0)) })
		defer _block1Value.Release()
		_block1 = unsafe.Pointer(_block1Value)
	}
	return networkCreateSecureTCPForPlain(networkProtocolBlockValue(_nw_parameters_configure_protocol_disableSymbol), _block1), nil
}

// TryNWParametersCreatePlainTCP initializes parameters for cleartext TCP.
func TryNWParametersCreatePlainTCP(configureTCP NWParametersConfigureProtocolBlock) (NWParameters, error) {
	return tryNWParametersCreatePlainTCP(configureTCP)
}

// NWParametersCreatePlainTCP initializes parameters for cleartext TCP.
func NWParametersCreatePlainTCP(configureTCP NWParametersConfigureProtocolBlock) NWParameters {
	result, callErr := tryNWParametersCreatePlainTCP(configureTCP)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerSymbol(&_nw_parameters_configure_protocol_default_configurationSymbol, &_nw_parameters_configure_protocol_default_configurationErr, frameworkHandle, "_nw_parameters_configure_protocol_default_configuration", "10.14")
	registerSymbol(&_nw_parameters_configure_protocol_disableSymbol, &_nw_parameters_configure_protocol_disableErr, frameworkHandle, "_nw_parameters_configure_protocol_disable", "10.14")
	registerFunc(&networkCreateSecureTCPForPlain, &networkCreateSecureTCPForPlainErr, frameworkHandle, "nw_parameters_create_secure_tcp", "10.14")
}
