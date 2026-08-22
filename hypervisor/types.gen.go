// Code generated from Apple documentation for Hypervisor. DO NOT EDIT.

package hypervisor

// C struct types

// HVAPICState
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_apic_state
type HVAPICState struct {
	Apic_gpa      uint64
	Apic_controls uint64
	Tsc_deadline  uint64
	Apic_id       uint32
	Ver           uint32
	Tpr           uint32
	Apr           uint32
	Ldr           uint32
	Dfr           uint32
	Svr           uint32
	Isr           [8]uint32
	Tmr           [8]uint32
	Irr           [8]uint32
	Esr           uint32
	Lvt           [7]uint32
	Icr           [2]uint32
	Icr_timer     uint32
	Dcr_timer     uint32
	Ccr_timer     uint32
	Esr_pending   uint32
	Boot_state    HVBootState
	Aeoi          [8]uint32
}

// Hv_apic_state is a type alias for HVAPICState for use in objc.Send[T] calls.
type Hv_apic_state = HVAPICState

// HVAPICStateExt
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_apic_state_ext_t
type HVAPICStateExt struct {
	Version uint32
	State   HVAPICState
}

// Hv_apic_state_ext_t is a type alias for HVAPICStateExt for use in objc.Send[T] calls.
type Hv_apic_state_ext_t = HVAPICStateExt

// HVAtpicState
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_atpic_state
type HVAtpicState struct {
	Ready        bool
	Icw_num      uint8
	Rd_cmd_reg   uint8
	Aeoi         bool
	Poll         bool
	Rotate       bool
	Sfn          bool
	Irq_base     uint8
	Request      uint8
	Service      uint8
	Mask         uint8
	Smm          bool
	Last_request uint8
	Lowprio      uint8
	Intr_raised  bool
	Elc          uint8
}

// Hv_atpic_state is a type alias for HVAtpicState for use in objc.Send[T] calls.
type Hv_atpic_state = HVAtpicState

// HVAtpicStateExt
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_atpic_state_ext_t
type HVAtpicStateExt struct {
	Version uint32
	State   HVAtpicState
}

// Hv_atpic_state_ext_t is a type alias for HVAtpicStateExt for use in objc.Send[T] calls.
type Hv_atpic_state_ext_t = HVAtpicStateExt

// HVIoapicState
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_ioapic_state
type HVIoapicState struct {
	Rtbl     [32]uint64
	Irr      uint32
	Ioa_id   uint32
	Ioregsel uint32
}

// Hv_ioapic_state is a type alias for HVIoapicState for use in objc.Send[T] calls.
type Hv_ioapic_state = HVIoapicState

// HVIoapicStateExt
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_ioapic_state_ext_t
type HVIoapicStateExt struct {
	Version uint32
	State   HVIoapicState
}

// Hv_ioapic_state_ext_t is a type alias for HVIoapicStateExt for use in objc.Send[T] calls.
type Hv_ioapic_state_ext_t = HVIoapicStateExt

// HVVCPUExitException - The structure that describes information about an exit from the virtual CPU (vCPU) to the host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_exit_exception_t
type HVVCPUExitException struct {
	Syndrome         HVExceptionSyndrome // The vCPU exception syndrome causing the exception.
	Virtual_address  HVExceptionAddress  // The vCPU virtual address of the exception.
	Physical_address HVIPA               // The intermediate physical address of the exception in the client.

}

// Hv_vcpu_exit_exception_t is a type alias for HVVCPUExitException for use in objc.Send[T] calls.
type Hv_vcpu_exit_exception_t = HVVCPUExitException

// HVVCPUExit - Information about an exit from the vCPU to the host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_exit_t
type HVVCPUExit struct {
	Reason    HVExitReason        // Information about an exit from the vcpu to the host.
	Exception HVVCPUExitException // Information about an exit exception from the vcpu to the host.

}

// Hv_vcpu_exit_t is a type alias for HVVCPUExit for use in objc.Send[T] calls.
type Hv_vcpu_exit_t = HVVCPUExit

// HVVCPUSMEState
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_sme_state_t
type HVVCPUSMEState struct {
	Streaming_sve_mode_enabled bool
	Za_storage_enabled         bool
}

// Hv_vcpu_sme_state_t is a type alias for HVVCPUSMEState for use in objc.Send[T] calls.
type Hv_vcpu_sme_state_t = HVVCPUSMEState
