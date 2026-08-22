// Command plaintcp configures cleartext TCP parameters from a block callback.
//
// It demonstrates that the configure block receives the concrete
// [network.NWProtocolOptions] its typedef declares, not an opaque object
// wrapper: the callback both reads the value and passes it straight to
// NWTCPOptionsSetNoDelay, which accepts nothing weaker.
package main

import (
	"fmt"
	"runtime"

	"github.com/tmc/apple/network"
)

func main() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	// The block parameter is typed by NWParametersConfigureProtocolBlock,
	// which declares func(NWProtocolOptions). Naming the concrete type here
	// is what the compiler checks; using the value below is what proves the
	// object arrives live rather than as a zero wrapper.
	called := false
	params := network.NWParametersCreatePlainTCP(func(options network.NWProtocolOptions) {
		called = true
		if options.ID == 0 {
			panic("configure block received a nil nw_protocol_options_t")
		}
		definition := network.NWProtocolOptionsCopyDefinition(options)
		if definition.ID == 0 {
			panic("nw_protocol_options_copy_definition returned nil")
		}
		defer definition.Release()

		// Passing options on unchanged only compiles because the block
		// parameter is NWProtocolOptions; this is the call the generator
		// defect made impossible to write.
		network.NWTCPOptionsSetNoDelay(options, true)
		fmt.Println("configure block ran with a live NWProtocolOptions")
	})
	if !called {
		panic("configure block never ran")
	}
	if params.ID == 0 {
		panic("nw_parameters_create_secure_tcp returned nil")
	}
	defer params.Release()

	// The parameters carry the transport options the block configured.
	stack := network.NWParametersCopyDefaultProtocolStack(params)
	if stack.ID == 0 {
		panic("nw_parameters_copy_default_protocol_stack returned nil")
	}
	defer stack.Release()

	transport := network.NWProtocolStackCopyTransportProtocol(stack)
	if transport.ID == 0 {
		panic("nw_protocol_stack_copy_transport_protocol returned nil")
	}
	defer transport.Release()

	fmt.Println("plain TCP parameters configured")
}
