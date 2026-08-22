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

// HVSIMDFPUchar16 is the value that represents an ARM SIMD and FP register.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_simd_fp_uchar16_t
type HVSIMDFPUchar16 = [16]byte

// See: https://developer.apple.com/documentation/Hypervisor/hv_sme_zt0_uchar64_t
type HVSMEZt0Uchar64 = [64]byte

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
type HVVcpuid = uint32

// HVVmConfig is the type that defines a virtual-machine configuration.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vm_config_t
type HVVmConfig = unsafe.Pointer

// HVVmOptions is options you use when creating a virtual machine.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vm_options_t
type HVVmOptions = uint64

// HVVmxMsrInfo is the type that describes Move to Status Register (MSR) information fields.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vmx_msr_info_t
type HVVmxMsrInfo = uint64

// Hv_allocate_flags_t is a C-name alias for HVAllocateFlags.
type Hv_allocate_flags_t = HVAllocateFlags

// Hv_capability_t is a C-name alias for HVCapability.
type Hv_capability_t = HVCapability

// Hv_exception_address_t is a C-name alias for HVExceptionAddress.
type Hv_exception_address_t = HVExceptionAddress

// Hv_exception_syndrome_t is a C-name alias for HVExceptionSyndrome.
type Hv_exception_syndrome_t = HVExceptionSyndrome

// Hv_gic_config_t is a C-name alias for HVGICConfig.
type Hv_gic_config_t = HVGICConfig

// Hv_gic_state_t is a C-name alias for HVGICState.
type Hv_gic_state_t = HVGICState

// Hv_gpaddr_t is a C-name alias for HVGpaddr.
type Hv_gpaddr_t = HVGpaddr

// Hv_ion_flags_t is a C-name alias for HVIonFlags.
type Hv_ion_flags_t = HVIonFlags

// Hv_ipa_t is a C-name alias for HVIPA.
type Hv_ipa_t = HVIPA

// Hv_memory_flags_t is a C-name alias for HVMemoryFlags.
type Hv_memory_flags_t = HVMemoryFlags

// Hv_msr_flags_t is a C-name alias for HVMsrFlags.
type Hv_msr_flags_t = HVMsrFlags

// Hv_return_t is a C-name alias for HVReturn.
type Hv_return_t = HVReturn

// Hv_shadow_flags_t is a C-name alias for HVShadowFlags.
type Hv_shadow_flags_t = HVShadowFlags

// Hv_simd_fp_uchar16_t is a C-name alias for HVSIMDFPUchar16.
type Hv_simd_fp_uchar16_t = HVSIMDFPUchar16

// Hv_sme_zt0_uchar64_t is a C-name alias for HVSMEZt0Uchar64.
type Hv_sme_zt0_uchar64_t = HVSMEZt0Uchar64

// Hv_uvaddr_t is a C-name alias for HVUvaddr.
type Hv_uvaddr_t = HVUvaddr

// Hv_vcpu_config_t is a C-name alias for HVVCPUConfig.
type Hv_vcpu_config_t = HVVCPUConfig

// Hv_vcpu_options_t is a C-name alias for HVVCPUOptions.
type Hv_vcpu_options_t = HVVCPUOptions

// Hv_vcpu_t is a C-name alias for HVVCPU.
type Hv_vcpu_t = HVVCPU

// Hv_vcpuid_t is a C-name alias for HVVcpuid.
type Hv_vcpuid_t = HVVcpuid

// Hv_vm_config_t is a C-name alias for HVVmConfig.
type Hv_vm_config_t = HVVmConfig

// Hv_vm_options_t is a C-name alias for HVVmOptions.
type Hv_vm_options_t = HVVmOptions

// Hv_vm_space_t is a C-name alias for HVVmSpace.
type Hv_vm_space_t = HVVmSpace

// Hv_vmx_msr_info_t is a C-name alias for HVVmxMsrInfo.
type Hv_vmx_msr_info_t = HVVmxMsrInfo
