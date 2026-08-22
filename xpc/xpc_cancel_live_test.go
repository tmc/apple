// Copyright 2026 The apple Authors.

//go:build xpclive

// This file is the internal half of the live suite. It is in package xpc, not
// package xpc_test, because the resource under measurement — purego's callback
// table — has no exported accessor: blockCount reads len(xpcBlockKeepalive)/3
// on our side of the FFI boundary. It shares the test binary (and therefore
// TestMain, the LaunchAgent, and the running service) with xpc_live_test.go.
//
// The service name is recomputed rather than shared, since the external test
// package's variable is not visible here. It must stay in step with
// xpc_live_test.go's serviceName; liveServiceName checks that the recomputed
// name actually answers before any verdict is taken, so a drift in the formula
// reports as "no service" rather than as a clean measurement.
package xpc

import (
	"context"
	"fmt"
	"os"
	"sync"
	"testing"
	"time"
)

// blockCountSensitivity proves once per process that blockCount can still see a
// registration, in units of one. Every callback-cost assertion in this file
// expects zero, and a zero from a blind instrument is not a measurement.
//
// It is once per process on purpose. The probe registers a real callback, and
// purego's table holds 2000 with no way to give one back, so a probe run
// per test invocation makes the suite itself die under -count=2000 — which is
// exactly the run used to bracket the leak this file measures. A control that
// destroys the experiment is worse than no control; proving it once is enough,
// because the property proven is a property of the process.
var blockCountSensitivityOnce sync.Once

func blockCountSensitivity(t *testing.T) {
	t.Helper()
	blockCountSensitivityOnce.Do(func() {
		before := blockCount()
		if _, err := newXPCBlock(func(_ uintptr) {}); err != nil {
			t.Fatalf("newXPCBlock: %v", err)
		}
		if spent := blockCount() - before; spent != 1 {
			t.Fatalf("one newXPCBlock registered %d blocks, want 1: blockCount is in the wrong unit, "+
				"so every zero measured in this process is uninterpretable", spent)
		}
		t.Logf("sensitivity: blockCount sees one registration as 1")
	})
}

func liveServiceName() string {
	return fmt.Sprintf("dev.tmc.apple.xpc.live.%d", os.Getpid())
}

func liveDial(t *testing.T) *Session {
	t.Helper()
	s, err := DialMachService(liveServiceName(), SessionOptions{})
	if err != nil {
		t.Fatalf("dial %s: %v", liveServiceName(), err)
	}
	t.Cleanup(func() { s.Cancel() })
	// Prove the peer is really there. Without this, every measurement below
	// could be of a session that answers nothing.
	msg, err := s.CallDictionary(context.Background(), Dictionary{"op": "add", "firstNumber": int64(2), "secondNumber": int64(3)})
	if err != nil {
		t.Fatalf("probe call: %v", err)
	}
	if got := msg.Dictionary()["sumValue"]; got != int64(5) {
		t.Fatalf("probe call sumValue = %#v, want int64(5): the service is not the one this file expects", got)
	}
	return s
}

// TestLiveCancelledCallLeaksReplySlot measures what a cancelled Call costs.
//
// CallDictionary documents that "a cancelled call costs one leaked reply slot".
// That sentence names two distinct resources and this test separates them:
//
//   - the libxpc-side pending reply context, which is observable only
//     indirectly: the peer keeps running and the reply is still delivered to
//     the block after the caller has given up. Arm 2 observes that directly.
//   - the purego callback slot that callAsyncDictionary used to register per
//     message, which is counted exactly by blockCount. Arm 1 counts it, and it
//     is now zero on all three paths: the reply block carries a token and
//     dispatches to one shared trampoline. What the token table costs instead
//     is measured at the end, in liveReplyTokens.
func TestLiveCancelledCallLeaksReplySlot(t *testing.T) {
	s := liveDial(t)

	// Warm-up: pay every one-off registration (symbol resolution, shared
	// applier trampolines) before either measurement.
	for i := 0; i < 2; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		if _, err := s.CallDictionary(ctx, Dictionary{"op": "add", "firstNumber": int64(1), "secondNumber": int64(1)}); err != nil {
			cancel()
			t.Fatalf("warm-up call: %v", err)
		}
		cancel()
	}

	blockCountSensitivity(t)

	const n = 8

	// Arm 1a: n calls that are cancelled mid-flight.
	before := blockCount()
	for i := 0; i < n; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
		_, err := s.CallDictionary(ctx, Dictionary{"op": "slow", "firstNumber": int64(400)})
		cancel()
		if err != context.DeadlineExceeded {
			t.Fatalf("cancelled call %d returned %v, want context.DeadlineExceeded: the call was not in flight, so nothing below measures cancellation", i, err)
		}
	}
	cancelledCost := blockCount() - before

	// Arm 1b: the same number of calls that complete normally, on the same
	// cancellable path. This is the negative control for the attribution: if
	// both arms cost the same, the leak is not caused by cancellation.
	before = blockCount()
	for i := 0; i < n; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		msg, err := s.CallDictionary(ctx, Dictionary{"op": "add", "firstNumber": int64(i), "secondNumber": int64(1)})
		cancel()
		if err != nil {
			t.Fatalf("completed call %d: %v", i, err)
		}
		if got := msg.Dictionary()["sumValue"]; got != int64(i+1) {
			t.Fatalf("completed call %d sumValue = %#v: the work did not happen, so its cost is not a cost", i, got)
		}
	}
	completedCost := blockCount() - before

	// Arm 1c: the blocking path (ctx.Done() == nil), which does not register.
	before = blockCount()
	for i := 0; i < n; i++ {
		if _, err := s.CallDictionary(context.Background(), Dictionary{"op": "add", "firstNumber": int64(i), "secondNumber": int64(1)}); err != nil {
			t.Fatalf("sync call %d: %v", i, err)
		}
	}
	syncCost := blockCount() - before

	t.Logf("callback slots for %d messages: cancelled=%d completed=%d sync=%d", n, cancelledCost, completedCost, syncCost)

	// Every arm must now cost zero callback slots. The reply block carries a
	// token and dispatches to one shared trampoline, so the exhaustible
	// resource is no longer spent per message on any of the three paths.
	if cancelledCost != 0 {
		t.Errorf("cancelled calls cost %d callback slots for %d messages, want 0: purego's table holds 2000 "+
			"and panics from a native thread when full, which aborts the process", cancelledCost, n)
	}
	if completedCost != 0 {
		t.Errorf("completed calls cost %d callback slots for %d messages, want 0", completedCost, n)
	}
	if syncCost != 0 {
		t.Errorf("the blocking path cost %d slots for %d messages, want 0", syncCost, n)
	}

	// What replaced the callback slot is a map entry per call in flight. It is
	// ordinary memory rather than a fixed table, but nothing else reports it,
	// so it is measured here. A cancelled call's entry is freed when the late
	// reply arrives — TestLiveCancelledCallStillDeliversReply shows it does —
	// so the table must drain rather than grow with the cancelled calls above.
	deadline := time.Now().Add(10 * time.Second)
	held := liveReplyTokens()
	for held > 0 && time.Now().Before(deadline) {
		time.Sleep(50 * time.Millisecond)
		held = liveReplyTokens()
	}
	if held != 0 {
		t.Errorf("%d reply tokens are still held after %d cancelled and %d completed calls: the token table replaced "+
			"a bounded leak with an unbounded one", held, n, n)
	} else {
		t.Logf("reply tokens held at rest: 0 (the %d cancelled calls' entries were freed by their late replies)", n)
	}
}

// TestLiveCancelledCallStillDeliversReply observes what the reply block does
// after the caller has abandoned the call. It drives callAsyncDictionary
// directly, which is what CallDictionary does, so the block is the same block.
func TestLiveCancelledCallStillDeliversReply(t *testing.T) {
	s := liveDial(t)

	type outcome struct {
		dict Dictionary
		err  error
	}
	done := make(chan outcome, 1)
	start := time.Now()
	if err := s.callAsyncDictionary(Dictionary{"op": "slow", "firstNumber": int64(400)}, func(d Dictionary, err error) {
		done <- outcome{d, err}
	}); err != nil {
		t.Fatalf("callAsyncDictionary: %v", err)
	}

	// The caller gives up here, exactly as CallDictionary's ctx.Done() arm does.
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()
	select {
	case o := <-done:
		t.Fatalf("reply arrived in %v, before the caller gave up: the call was not in flight (%v, %v)", time.Since(start), o.dict, o.err)
	case <-ctx.Done():
	}
	abandoned := time.Since(start)

	select {
	case o := <-done:
		if o.err != nil {
			t.Fatalf("reply after abandonment carried an error: %v", o.err)
		}
		if got := o.dict["sumValue"]; got != int64(400) {
			t.Fatalf("late reply sumValue = %#v, want int64(400)", got)
		}
		t.Logf("the reply block ran %v after the caller abandoned the call (at %v) and delivered the peer's real answer: the pending reply context stays alive in libxpc",
			time.Since(start)-abandoned, abandoned)
	case <-time.After(5 * time.Second):
		t.Fatalf("the reply block never ran; the pending reply context did not complete")
	}

	// The session is still usable afterwards.
	msg, err := s.CallDictionary(context.Background(), Dictionary{"op": "add", "firstNumber": int64(20), "secondNumber": int64(22)})
	if err != nil {
		t.Fatalf("call after abandonment: %v", err)
	}
	if got := msg.Dictionary()["sumValue"]; got != int64(42) {
		t.Fatalf("after abandonment sumValue = %#v, want int64(42)", got)
	}
}

// TestLiveAsyncCallCostIsConstantInN is the O(1) claim, measured live against a
// real peer rather than argued from the source.
//
// The budget gate in callbackbudget_test.go asserts constancy in call count for
// paths a non-live test can drive; the cancellable-context path is not one of
// them, because it needs a peer to answer. This is that arm. It is the shape of
// the check that matters, not a pinned constant: a fix that merely moved the
// threshold — say, one callback per ten calls — passes a "fewer than 2000"
// assertion and fails this one.
func TestLiveAsyncCallCostIsConstantInN(t *testing.T) {
	s := liveDial(t)

	call := func(i int) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		msg, err := s.CallDictionary(ctx, Dictionary{"op": "add", "firstNumber": int64(i), "secondNumber": int64(1)})
		if err != nil {
			t.Fatalf("call %d: %v", i, err)
		}
		if got := msg.Dictionary()["sumValue"]; got != int64(i+1) {
			t.Fatalf("call %d sumValue = %#v, want %d: the work did not happen, so its cost is not a cost", i, got, i+1)
		}
	}

	// Warm-up outside both measurements.
	call(0)

	blockCountSensitivity(t)

	const small, large = 20, 400 // 20x apart: a per-call cost cannot hide in the ratio

	beforeSmall := blockCount()
	beforeSmallCalls := replyCount()
	for i := 0; i < small; i++ {
		call(i)
	}
	smallCost := blockCount() - beforeSmall
	if made := replyCount() - beforeSmallCalls; made != uint64(small) {
		t.Fatalf("the small arm made %d reply-carrying calls, want %d: the denominator is wrong, so the ratio is not about N", made, small)
	}

	beforeLarge := blockCount()
	beforeLargeCalls := replyCount()
	for i := 0; i < large; i++ {
		call(i)
	}
	largeCost := blockCount() - beforeLarge
	if made := replyCount() - beforeLargeCalls; made != uint64(large) {
		t.Fatalf("the large arm made %d reply-carrying calls, want %d", made, large)
	}

	t.Logf("callback slots: %d calls cost %d, %d calls cost %d", small, smallCost, large, largeCost)
	if smallCost != 0 || largeCost != 0 {
		t.Errorf("cost is not constant in N: %d calls cost %d slots and %d calls cost %d, want 0 and 0",
			small, smallCost, large, largeCost)
	}
}
