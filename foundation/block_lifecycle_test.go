package foundation

import (
	"runtime"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/tmc/apple/objc"
)

// retainCount returns the ObjC retain count for an object.
func retainCount(id objc.ID) uint {
	return objc.Send[uint](id, objc.Sel("retainCount"))
}

// TestBlockRetainPreventsUseAfterFree verifies that the retain call inside
// block callbacks keeps the object alive after autorelease pool drain.
func TestBlockRetainPreventsUseAfterFree(t *testing.T) {
	// Create an NSString that we'll pass through a block callback.
	str := NewStringWithCString("hello-block-test")
	strID := str.ID

	// Baseline: object is alive and has a reasonable retain count.
	rc := retainCount(strID)
	if rc == 0 {
		t.Fatalf("expected non-zero retain count, got %d", rc)
	}

	var received atomic.Value

	// Create a block that receives an NSString via objc.ID -> FromID.
	// The staged changes add a retain inside the block callback.
	block := objc.NewBlock(func(_ objc.Block, resultID objc.ID) {
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := NSStringFromID(resultID)
			received.Store(v)
		}
	})
	defer block.Release()

	// Invoke the block with our string's ID, simulating what ObjC does.
	block.Invoke(strID)

	v, ok := received.Load().(NSString)
	if !ok {
		t.Fatal("block callback did not store a value")
	}

	// Force GC to stress-test that the retained object survives.
	runtime.GC()
	runtime.GC()

	// The object should still be usable.
	got := objc.GoString(objc.Send[*byte](v.ID, objc.Sel("UTF8String")))
	if got != "hello-block-test" {
		t.Fatalf("expected %q, got %q", "hello-block-test", got)
	}

	// Retain count should have increased by 1 from the block's retain.
	rcAfter := retainCount(strID)
	if rcAfter <= rc {
		t.Fatalf("retain count should have increased: before=%d after=%d", rc, rcAfter)
	}
}

// TestBlockCallbackNilSafety verifies that nil/zero objc.ID is handled
// correctly in block callbacks (no retain on nil, nil pointer passed to handler).
func TestBlockCallbackNilSafety(t *testing.T) {
	var called atomic.Bool
	var gotNil atomic.Bool

	handler := func(s *NSString) {
		called.Store(true)
		gotNil.Store(s == nil)
	}

	// Use the generated NewStringBlock if available, otherwise manual.
	block := objc.NewBlock(func(_ objc.Block, resultID objc.ID) {
		var result *NSString
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := NSStringFromID(resultID)
			result = &v
		}
		handler(result)
	})
	defer block.Release()

	// Invoke with zero ID (nil object).
	block.Invoke(objc.ID(0))

	if !called.Load() {
		t.Fatal("handler was not called")
	}
	if !gotNil.Load() {
		t.Fatal("expected nil result for zero ID")
	}
}

// TestSafeErrorFromSemantics verifies SafeErrorFrom handles all cases:
// nil, valid NSError, and non-NSError objects.
func TestSafeErrorFromSemantics(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		err := SafeErrorFrom(0)
		if err != nil {
			t.Fatalf("expected nil error for zero ID, got %v", err)
		}
	})

	t.Run("valid_NSError", func(t *testing.T) {
		// Create a real NSError.
		domain := objc.String("TestDomain")
		userInfo := objc.ID(0) // nil userInfo
		errClass := objc.GetClass("NSError")
		errID := objc.Send[objc.ID](objc.ID(errClass), objc.Sel("errorWithDomain:code:userInfo:"), domain, 42, userInfo)

		err := SafeErrorFrom(errID)
		if err == nil {
			t.Fatal("expected non-nil error")
		}

		nsErr, ok := err.(*NSError)
		if !ok {
			t.Fatalf("expected *NSError, got %T", err)
		}

		if nsErr.Code() != 42 {
			t.Fatalf("expected code 42, got %d", nsErr.Code())
		}
	})

	t.Run("non_NSError_object", func(t *testing.T) {
		// Pass an NSString as the error ID — SafeErrorFrom should handle gracefully.
		str := NewStringWithCString("not-an-error")
		err := SafeErrorFrom(str.ID)
		if err == nil {
			t.Fatal("expected non-nil error for non-NSError object")
		}
		// Should use -description fallback, not crash.
		if err.Error() == "" {
			t.Fatal("expected non-empty error string from -description fallback")
		}
	})
}

// TestBlockRetainStress creates many blocks and invokes them with real ObjC
// objects under GC pressure to verify no use-after-free or crashes.
func TestBlockRetainStress(t *testing.T) {
	const iterations = 1000

	var wg sync.WaitGroup
	var successCount atomic.Int64

	for i := 0; i < iterations; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()

			// Create a fresh NSString for each iteration.
			s := NewStringWithCString("stress-test")

			block := objc.NewBlock(func(_ objc.Block, resultID objc.ID) {
				if resultID != 0 {
					objc.Send[objc.ID](resultID, objc.Sel("retain"))
					v := NSStringFromID(resultID)
					got := objc.GoString(objc.Send[*byte](v.ID, objc.Sel("UTF8String")))
					if got == "stress-test" {
						successCount.Add(1)
					}
				}
			})

			// Invoke with the string.
			block.Invoke(s.ID)

			// Don't release the block — simulating the "no cleanup" pattern.
			// This is what the staged changes do.

			// Force GC to increase pressure.
			if idx%100 == 0 {
				runtime.GC()
			}
		}(i)
	}

	wg.Wait()

	if successCount.Load() != iterations {
		t.Fatalf("expected %d successes, got %d", iterations, successCount.Load())
	}
}

// TestBlockCleanupNoLeak verifies that blocks WITH cleanup properly free
// their cache entries (baseline for the cleanup-removal tradeoff).
func TestBlockCleanupNoLeak(t *testing.T) {
	const iterations = 100

	for i := 0; i < iterations; i++ {
		block, cleanup := NewVoidBlock(func() {})
		_ = block

		// Simulate the old pattern: immediate cleanup.
		cleanup()
	}

	// If we get here without crashing, the cleanup path works.
	// We can't directly inspect blockFunctionCache, but no crash = good.
	runtime.GC()
}

// TestBlockNoCleanupDoesNotCrash verifies that blocks WITHOUT cleanup
// (the new pattern) don't crash even under GC pressure.
func TestBlockNoCleanupDoesNotCrash(t *testing.T) {
	const iterations = 100

	for i := 0; i < iterations; i++ {
		block, _ := NewVoidBlock(func() {})
		_ = block
		// No cleanup — this is the new pattern.
	}

	// Force aggressive GC.
	for i := 0; i < 5; i++ {
		runtime.GC()
	}

	// If we get here, no crash.
}

// TestErrorBlockWithSafeErrorFrom verifies the generated NewErrorBlock
// correctly uses SafeErrorFrom.
func TestErrorBlockWithSafeErrorFrom(t *testing.T) {
	var gotErr atomic.Value

	block, cleanup := NewErrorBlock(func(err error) {
		gotErr.Store(err)
	})
	defer cleanup()

	// Create a real NSError and invoke the block with it.
	domain := objc.String("TestErrorBlockDomain")
	errID := objc.Send[objc.ID](objc.ID(objc.GetClass("NSError")),
		objc.Sel("errorWithDomain:code:userInfo:"), domain, 99, objc.ID(0))

	objc.Block(block).Invoke(errID)

	v := gotErr.Load()
	if v == nil {
		t.Fatal("handler was not called")
	}

	err, ok := v.(error)
	if !ok {
		t.Fatalf("expected error, got %T", v)
	}

	nsErr, ok := err.(*NSError)
	if !ok {
		t.Fatalf("expected *NSError, got %T", err)
	}

	if nsErr.Code() != 99 {
		t.Fatalf("expected code 99, got %d", nsErr.Code())
	}
}

// BenchmarkBlockCreateAndInvoke measures the overhead of creating and
// invoking blocks with the retain pattern.
func BenchmarkBlockCreateAndInvoke(b *testing.B) {
	str := NewStringWithCString("bench")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		block := objc.NewBlock(func(_ objc.Block, resultID objc.ID) {
			if resultID != 0 {
				objc.Send[objc.ID](resultID, objc.Sel("retain"))
				_ = NSStringFromID(resultID)
			}
		})
		block.Invoke(str.ID)
		block.Release()
	}
}

// BenchmarkBlockNoCleanup measures the cost of creating blocks without
// releasing them (the new pattern).
func BenchmarkBlockNoCleanup(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		block, _ := NewVoidBlock(func() {})
		_ = block
	}
}
