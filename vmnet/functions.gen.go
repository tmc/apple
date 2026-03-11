// Code generated from Apple documentation for vmnet. DO NOT EDIT.

package vmnet

import (
	"fmt"
	"os"
	"unsafe"
	"github.com/ebitengine/purego"
)

// registerFunc resolves a framework symbol and registers it as a Go function.
// If the symbol is not found, a warning is printed and the function pointer is left nil.
func registerFunc(fptr any, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: vmnet: symbol %s not found\n", name)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "warning: vmnet: symbol %s registration skipped: %v\n", name, r)
		}
	}()
	purego.RegisterFunc(fptr, sym)
}

// registerSymbol resolves a framework symbol and stores its raw address.
func registerSymbol(dst *uintptr, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: vmnet: symbol %s not found\n", name)
		return
	}
	*dst = sym
}


var _vmnet_copy_shared_interface_list func() unsafe.Pointer

// Vmnet_copy_shared_interface_list.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_copy_shared_interface_list()
func Vmnet_copy_shared_interface_list() unsafe.Pointer {
	if _vmnet_copy_shared_interface_list == nil {
		panic("vmnet: symbol vmnet_copy_shared_interface_list not loaded")
	}
	return _vmnet_copy_shared_interface_list()
}


































func init() {
	if frameworkHandle == 0 {
		return
	}
		registerFunc(&_vmnet_copy_shared_interface_list, frameworkHandle, "vmnet_copy_shared_interface_list")
	}






