// Code generated from Apple documentation for Hypervisor. DO NOT EDIT.

package hypervisor

import (
	"unsafe"
)

// C struct types

// HVAPICState
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_apic_state
type HVAPICState struct {
	Aeoi          uint32
	Apic_controls uint64
	Apic_gpa      uint64
	Apic_id       uint32
	Apr           uint32
	Boot_state    unsafe.Pointer
	Ccr_timer     uint32
	Dcr_timer     uint32
	Dfr           uint32
	Esr           uint32
	Esr_pending   uint32
	Icr           uint32
	Icr_timer     uint32
	Irr           uint32
	Isr           uint32
	Ldr           uint32
	Lvt           uint32
	Svr           uint32
	Tmr           uint32
	Tpr           uint32
	Tsc_deadline  uint64
	Ver           uint32
}

// HVAPICStateExt
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_apic_state_ext_t
type HVAPICStateExt struct {
	State   unsafe.Pointer
	Version uint32
}

// HVAtpicState
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_atpic_state
type HVAtpicState struct {
	Aeoi         bool
	Elc          uint8
	Icw_num      uint8
	Intr_raised  bool
	Irq_base     uint8
	Last_request uint8
	Lowprio      uint8
	Mask         uint8
	Poll         bool
	Rd_cmd_reg   uint8
	Ready        bool
	Request      uint8
	Rotate       bool
	Service      uint8
	Sfn          bool
	Smm          bool
}

// HVAtpicStateExt
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_atpic_state_ext_t
type HVAtpicStateExt struct {
	State   unsafe.Pointer
	Version uint32
}

// HVIoapicState
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_ioapic_state
type HVIoapicState struct {
	Ioa_id   uint32
	Ioregsel uint32
	Irr      uint32
	Rtbl     uint64
}

// HVIoapicStateExt
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_ioapic_state_ext_t
type HVIoapicStateExt struct {
	State   unsafe.Pointer
	Version uint32
}

// HVIonMessage - The structure that describes the Mach message that the Hypervisor sends when an I/O notifier delivers the notifications you request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_ion_message_t
type HVIonMessage struct {
	Addr    uint64         // The address of the I/O write.
	Header  unsafe.Pointer // The Mach message header.
	Size    uint64         // The size of the value written by the notifier.
	Trailer unsafe.Pointer // The Mach message trailer.
	Value   uint64         // An unsigned 64-bit integer that represents the contents of an I/O notifier message.

}

// HVVCPUExitException - The structure that describes information about an exit from the virtual CPU (vCPU) to the host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_exit_exception_t
type HVVCPUExitException struct {
	Syndrome         uint64 // The vCPU exception syndrome causing the exception.
	Virtual_address  uint64 // The vCPU virtual address of the exception.
	Physical_address uint64 // The intermediate physical address of the exception in the client.

}

// HVVCPUExit - Information about an exit from the vCPU to the host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_exit_t
type HVVCPUExit struct {
	Reason    HVExitReason   // Information about an exit from the vcpu to the host.
	Exception unsafe.Pointer // Information about an exit exception from the vcpu to the host.

}

// HVVCPUSMEState
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_sme_state_t
type HVVCPUSMEState struct {
	Streaming_sve_mode_enabled bool
	Za_storage_enabled         bool
}
