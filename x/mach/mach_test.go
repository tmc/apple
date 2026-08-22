package mach

import (
	"runtime"
	"testing"
	"time"
)

func TestThreadSelf(t *testing.T) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	th, err := ThreadSelf()
	if err != nil {
		t.Fatal(err)
	}
	if th == 0 {
		t.Fatal("ThreadSelf returned port 0")
	}
	th2, err := ThreadSelf()
	if err != nil {
		t.Fatal(err)
	}
	if th != th2 {
		t.Fatalf("ThreadSelf not stable on a locked thread: %d != %d", th, th2)
	}
}

func TestSetTimeConstraint(t *testing.T) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	th, err := ThreadSelf()
	if err != nil {
		t.Fatal(err)
	}
	tc := TimeConstraint{
		Period:      10 * time.Millisecond,
		Computation: 2 * time.Millisecond,
		Constraint:  4 * time.Millisecond,
		Preemptible: true,
	}
	if err := th.SetTimeConstraint(tc); err != nil {
		t.Fatal(err)
	}
	// Read the policy back: a 0 return from thread_policy_set is not proof
	// the class changed.
	got, err := th.getTimeConstraint()
	if err != nil {
		t.Fatal(err)
	}
	if got[0] == 0 || got[1] == 0 {
		t.Fatalf("time constraint not in effect after set: %v", got)
	}
	if err := th.SetStandard(); err != nil {
		t.Fatal(err)
	}
}

func TestSetTimeConstraintValidation(t *testing.T) {
	th := Thread(1)
	err := th.SetTimeConstraint(TimeConstraint{
		Period:      time.Millisecond,
		Computation: 2 * time.Millisecond, // computation > period
		Constraint:  time.Millisecond,
	})
	if err == nil {
		t.Fatal("want validation error for computation > constraint")
	}
}
