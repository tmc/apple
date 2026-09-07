// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"fmt"

	"github.com/ebitengine/purego"
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
		return fmt.Sprintf("CoreML: symbol %s unavailable on this system (introduced in macOS %s)", e.symbol, e.introduced)
	}
	return fmt.Sprintf("CoreML: symbol %s unavailable on this system", e.symbol)
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
		return fmt.Errorf("CoreML: symbol %s unavailable because the framework could not be loaded", name)
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
			*errDst = fmt.Errorf("CoreML: register symbol %s: %v", name, r)
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

// SymbolAddress returns the address of name in CoreML, whether or not
// this package generated a binding for it.
//
// What is generated is bounded by what is documented, and for a private
// framework that is whatever a manifest happened to enumerate. The dylib
// usually exports far more. Without this, a symbol nobody wrote down is
// unreachable from a package that has already loaded the image holding it, and
// the generated surface becomes a ceiling instead of a floor.
//
// The lookup is scoped to this framework's handle, not RTLD_DEFAULT, so a
// symbol some other loaded image exports is not reported as this one's.
func SymbolAddress(name string) (uintptr, error) {
	if frameworkHandle == 0 {
		return 0, fmt.Errorf("CoreML: symbol %s unavailable because the framework could not be loaded", name)
	}
	sym, err := purego.Dlsym(frameworkHandle, name)
	if err != nil || sym == 0 {
		return 0, missingSymbolError(name, "", err)
	}
	return sym, nil
}

// BindFunc binds the CoreML symbol name into fptr, which must be a
// pointer to a func variable.
//
// The caller supplies the signature, and nothing checks it. A dylib records no
// argument count or types for a C symbol, so a wrong signature here is not a
// type error: it is the wrong number of machine words moved on a live stack,
// and the failure surfaces somewhere else entirely. Prefer a generated binding,
// whose signature carries recorded evidence, and reach for this only for a
// symbol that has none.
// purego.RegisterFunc panics on a signature it cannot lower; that is recovered
// and returned, because an escape hatch that takes down the process on a
// mistyped experiment is not one anybody can experiment with.
func BindFunc(fptr any, name string) (err error) {
	sym, err := SymbolAddress(name)
	if err != nil {
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("CoreML: bind symbol %s: %v", name, r)
		}
	}()
	purego.RegisterFunc(fptr, sym)
	return nil
}

func init() {
	if frameworkHandle == 0 {
		return
	}
}
