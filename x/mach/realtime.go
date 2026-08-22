package mach

import (
	"fmt"
	"time"
	"unsafe"

	"github.com/tmc/apple/kernel"
)

// Thread policy flavors from <mach/thread_policy.h>.
const (
	threadStandardPolicy       = 1 // THREAD_STANDARD_POLICY
	threadTimeConstraintPolicy = 2 // THREAD_TIME_CONSTRAINT_POLICY
)

// TimeConstraint describes a real-time scheduling contract, as passed to
// THREAD_TIME_CONSTRAINT_POLICY. Period is the interval between deadlines
// (for audio, the buffer duration), Computation the CPU time needed per
// period, and Constraint the window within which that computation must
// complete (Computation <= Constraint <= Period).
type TimeConstraint struct {
	Period      time.Duration
	Computation time.Duration
	Constraint  time.Duration
	Preemptible bool
}

// SetTimeConstraint promotes the thread to the real-time scheduling class.
// The calling goroutine must be locked to the thread t was obtained from.
func (t Thread) SetTimeConstraint(tc TimeConstraint) error {
	if err := load(); err != nil {
		return err
	}
	if tc.Computation > tc.Constraint || tc.Constraint > tc.Period {
		return fmt.Errorf("mach: time constraint wants computation <= constraint <= period, got %v/%v/%v",
			tc.Computation, tc.Constraint, tc.Period)
	}
	// thread_time_constraint_policy_data_t from <mach/thread_policy.h>:
	// uint32_t period, computation, constraint; boolean_t preemptible.
	var policy [4]uint32
	policy[0] = toAbs(tc.Period)
	policy[1] = toAbs(tc.Computation)
	policy[2] = toAbs(tc.Constraint)
	if tc.Preemptible {
		policy[3] = 1
	}
	kr := kernel.Thread_policy_set(uint32(t), threadTimeConstraintPolicy,
		(*int32)(unsafe.Pointer(&policy[0])), uint32(len(policy)))
	return kernError("thread_policy_set(TIME_CONSTRAINT)", kr)
}

// SetStandard demotes the thread back to the standard scheduling class.
func (t Thread) SetStandard() error {
	if err := load(); err != nil {
		return err
	}
	// thread_standard_policy_data_t is a single reserved uint32.
	var policy [1]uint32
	kr := kernel.Thread_policy_set(uint32(t), threadStandardPolicy,
		(*int32)(unsafe.Pointer(&policy[0])), uint32(len(policy)))
	return kernError("thread_policy_set(STANDARD)", kr)
}

// getTimeConstraint reads back the current time-constraint policy, for tests.
func (t Thread) getTimeConstraint() ([4]uint32, error) {
	var policy [4]uint32
	count := uint32(len(policy))
	var def kernel.Boolean_t
	kr := kernel.Thread_policy_get(uint32(t), threadTimeConstraintPolicy,
		(*int32)(unsafe.Pointer(&policy[0])), &count, &def)
	if err := kernError("thread_policy_get(TIME_CONSTRAINT)", kr); err != nil {
		return policy, err
	}
	if def != 0 {
		return policy, fmt.Errorf("mach: thread reports default policy, not the set one")
	}
	return policy, nil
}

// toAbs converts a duration to Mach absolute time units using
// mach_timebase_info (ticks = ns * denom / numer).
func toAbs(d time.Duration) uint32 {
	var tb [2]uint32 // numer, denom
	timebaseInfo(&tb)
	if tb[0] == 0 || tb[1] == 0 {
		return uint32(d.Nanoseconds())
	}
	return uint32(uint64(d.Nanoseconds()) * uint64(tb[1]) / uint64(tb[0]))
}
