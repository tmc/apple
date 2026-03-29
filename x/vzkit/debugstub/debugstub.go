package debugstub

import (
	"fmt"
	"unsafe"

	"github.com/tmc/apple/objectivec"
	pvz "github.com/tmc/apple/private/virtualization"
)

// NewDebugStubConfiguration creates the base debug-stub configuration.
func NewDebugStubConfiguration() pvz.VZDebugStubConfiguration {
	return pvz.NewVZDebugStubConfiguration()
}

// NewGDBDebugStubConfiguration creates a GDB debug-stub configuration.
func NewGDBDebugStubConfiguration(port uint16, listenAll bool) pvz.VZGDBDebugStubConfiguration {
	cfg := pvz.NewVZGDBDebugStubConfigurationWithPort(port)
	cfg.SetListensOnAllNetworkInterfaces(listenAll)
	return cfg
}

// NewForwardingDebugStubConfiguration creates a forwarding debug-stub configuration.
func NewForwardingDebugStubConfiguration(stub unsafe.Pointer) pvz.VZForwardingDebugStubConfiguration {
	cfg := pvz.NewVZForwardingDebugStubConfiguration()
	if raw := cfg.InitWithDebugStub(stub); raw != nil && raw.GetID() != 0 {
		return pvz.VZForwardingDebugStubConfigurationFromID(raw.GetID())
	}
	return cfg
}

// Attach sets a debug stub on a VM configuration.
func Attach(config pvz.VZVirtualMachineConfiguration, stub objectivec.IObject) error {
	if config.ID == 0 {
		return fmt.Errorf("vm configuration is nil")
	}
	if stub == nil || stub.GetID() == 0 {
		return fmt.Errorf("debug stub is nil")
	}
	config.SetDebugStub(stub)
	return nil
}

// Clear removes any debug stub from a VM configuration.
func Clear(config pvz.VZVirtualMachineConfiguration) {
	if config.ID == 0 {
		return
	}
	config.SetDebugStub(nil)
}

// AttachGDB builds a GDB debug stub and attaches it to a VM configuration.
func AttachGDB(config pvz.VZVirtualMachineConfiguration, port uint16, listenAll bool) error {
	return Attach(config, NewGDBDebugStubConfiguration(port, listenAll))
}

// AttachForwarding builds a forwarding debug stub and attaches it to a VM configuration.
func AttachForwarding(config pvz.VZVirtualMachineConfiguration, stub unsafe.Pointer) error {
	return Attach(config, NewForwardingDebugStubConfiguration(stub))
}

// BuildForVM asks a debug-stub configuration to create a VM stub object.
func BuildForVM(stub pvz.VZDebugStubConfiguration, vm objectivec.IObject) objectivec.IObject {
	if stub.ID == 0 {
		return nil
	}
	return stub.MakeDebugStubForVirtualMachine(vm)
}

// BuildForCoprocessor asks a debug-stub configuration to create a coprocessor stub object.
func BuildForCoprocessor(stub pvz.VZDebugStubConfiguration) objectivec.IObject {
	if stub.ID == 0 {
		return nil
	}
	return stub.MakeDebugStubForCoprocessor()
}
