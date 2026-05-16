// Code generated from Apple documentation. DO NOT EDIT.

package hypervisor

import (
	"unsafe"
)

// See: https://developer.apple.com/documentation/Hypervisor/hv_allocate_flags_t
type HVAllocateFlags = uint64

// HVCapability is the type of system capabilities.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_capability_t
type HVCapability = uint64

// HVExceptionAddress is type of a vCPU exception virtual address.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_exception_address_t
type HVExceptionAddress = uint64

// HVExceptionSyndrome is type of a vCPU exception syndrome.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_exception_syndrome_t
type HVExceptionSyndrome = uint64

// HVGICConfig is an alias for this value type’s equivalent Hypervisor generic interrupt controller (GIC) configuration’s reference type.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_gic_config_t
type HVGICConfig = unsafe.Pointer

// HVGICState is an alias for this value type’s equivalent Hypervisor generic interrupt controller (GIC) state’s reference type.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_gic_state_t
type HVGICState = unsafe.Pointer

// HVGpaddr is the type of a guest physical address (GPA).
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_gpaddr_t
type HVGpaddr = uint64

// HVIonFlags is the bitfield that you use to set the options flags for the I/O notifier.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_ion_flags_t
type HVIonFlags = uint32

// HVIPA is the type of an intermediate physical address, which is a guest physical address space of the VM.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_ipa_t
type HVIPA = uint64

// HVMemoryFlags is the permissions for guest physical memory regions.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_memory_flags_t
type HVMemoryFlags = uint64

// HVMsrFlags is the type representing the native Model-Specific Register (MSR) permissions.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_msr_flags_t
type HVMsrFlags = uint32

// HVShadowFlags is shadow VMCS permissions for the set shadow access function.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_shadow_flags_t
type HVShadowFlags = uint64

// HVUvaddr is the type of a user virtual address.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_uvaddr_t
type HVUvaddr = unsafe.Pointer

// HVVCPUConfig is the type that defines a vCPU configuration.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_config_t
type HVVCPUConfig = unsafe.Pointer

// HVVCPUOptions is options for creating a new vCPU instance.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_options_t
type HVVCPUOptions = uint64

// HVVcpuid is the type that describes a vCPU ID.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vcpuid_t
type HVVcpuid = uint64

// HVVmConfig is the type that defines a virtual-machine configuration.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vm_config_t
type HVVmConfig = unsafe.Pointer

// HVVmOptions is options you use when creating a virtual machine.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vm_options_t
type HVVmOptions = uint64

// HVVmSpace is the type of a guest-address space.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vm_space_t
type HVVmSpace = uint32

// HVVmxMsrInfo is the type that describes Move to Status Register (MSR) information fields.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vmx_msr_info_t
type HVVmxMsrInfo = uint64
