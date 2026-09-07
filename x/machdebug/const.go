package machdebug

// Mach and AArch64 constants used by this package.
//
// These are ABI-pinned in xnu and absent from Apple's documentation set, so
// they are not generated; each is named with the header it was read from.
// The values were checked against the macOS SDK headers on 2026-09-06.
const (
	// <mach/exception_types.h>
	excBreakpoint     = 6        // EXC_BREAKPOINT
	excMaskBreakpoint = 1 << 6   // EXC_MASK_BREAKPOINT
	exceptionDefault  = 1        // EXCEPTION_DEFAULT: ports only, no state
	machExcCodes      = -1 << 31 // MACH_EXCEPTION_CODES: 64-bit codes, msgh_id 2405

	// mach_exception_raise request and its reply (request + 100), from the
	// MIG-generated mach_exc subsystem.
	msgIDRaise      = 2405
	msgIDRaiseReply = 2505

	// <mach/kern_return.h>
	kernSuccess           = 0
	kernProtectionFailure = 2

	// <mach/message.h>: what a send to a dead task's port reports.
	machSendInvalidDest = 0x10000003

	// <mach/arm/thread_status.h>
	armThreadState64      = 6  // ARM_THREAD_STATE64
	armThreadState64Count = 68 // sizeof(arm_thread_state64_t)/4
	armDebugState64       = 15 // ARM_DEBUG_STATE64
	// ARM_DEBUG_STATE64_COUNT: sizeof(arm_debug_state64_t)/4, the struct
	// being bvr[16] bcr[16] wvr[16] wcr[16] mdscr_el1, all uint64.
	armDebugState64Count = (16*4 + 1) * 2

	// <mach/vm_prot.h>
	vmProtRead    = 0x01 // VM_PROT_READ
	vmProtWrite   = 0x02 // VM_PROT_WRITE
	vmProtExecute = 0x04 // VM_PROT_EXECUTE
	// VM_PROT_COPY forces copy-on-write when raising protection. It is what
	// keeps a text patch private to the target task: shared-cache and
	// file-backed text is mapped read-only and shared by every process that
	// maps it, and without this flag the protect either fails with
	// KERN_PROTECTION_FAILURE or succeeds and rewrites text other processes
	// are executing.
	vmProtCopy = 0x10

	// <mach/vm_attributes.h>
	mattrCache          = 1 // MATTR_CACHE
	mattrValICacheFlush = 8 // MATTR_VAL_ICACHE_FLUSH

	// <mach/vm_region.h>
	vmRegionBasicInfo64      = 9 // VM_REGION_BASIC_INFO_64
	vmRegionBasicInfo64Count = 9 // VM_REGION_BASIC_INFO_COUNT_64, in uint32 words

	// AArch64 A64 encoding of "brk #0", the debugger trap. Confirmed by
	// assembling it: clang -c -arch arm64 emits d4200000.
	brkZero = 0xD4200000

	// MDSCR_EL1.SS (bit 0) and PSTATE.SS (CPSR bit 21) together arm one
	// hardware single step; see D2.12 of the ARMv8-A manual and the
	// mdscr_el1 comment in <mach/arm/_structs.h>.
	mdscrSS = 1 << 0
	cpsrSS  = 1 << 21

	// Debug watchpoint control (DBGWCRn_EL1) fields, ARMv8-A D13.3.11.
	wcrEnable   = 1 << 0 // E
	wcrPACEL0   = 2 << 1 // PAC: privilege access control, EL0 only
	wcrLoad     = 1 << 3 // LSC bit 0: watch loads
	wcrStore    = 1 << 4 // LSC bit 1: watch stores
	wcrBASShift = 5      // BAS: byte address select, one bit per watched byte
)
