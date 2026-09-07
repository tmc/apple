package machdebug

// Regs is arm_thread_state64_t from <mach/arm/_structs.h>.
//
// The layout is ABI-pinned in xnu and absent from Apple's documentation set,
// so it is hand-written here rather than generated. Flags carries
// pointer-authentication state on arm64e; see [Thread.SetRegs].
type Regs struct {
	X     [29]uint64
	FP    uint64
	LR    uint64
	SP    uint64
	PC    uint64
	CPSR  uint32
	Flags uint32
}
