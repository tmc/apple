// Package mach provides a small, hand-written discipline layer over the raw
// Mach bindings in the kernel package: thread ports and real-time thread
// promotion.
//
// The raw calls (thread_policy_set, mach_port_deallocate, ...) are generated
// into the kernel package from Apple's Kernel framework documentation. What
// generation cannot supply lives here: acquiring the right thread port
// without leaking a port right, converting durations to Mach absolute time
// units, and the pin-then-promote calling discipline.
//
// Promotion applies to an OS thread, not a goroutine. Callers must hold the
// thread with [runtime.LockOSThread] before promoting, and the promotion
// lasts until the thread is demoted or exits:
//
//	runtime.LockOSThread()
//	defer runtime.UnlockOSThread()
//	t := mach.ThreadSelf()
//	err := t.SetTimeConstraint(mach.TimeConstraint{
//		Period:      time.Millisecond * 10,
//		Computation: time.Millisecond * 2,
//		Constraint:  time.Millisecond * 4,
//		Preemptible: true,
//	})
package mach
