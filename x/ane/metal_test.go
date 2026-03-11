//go:build darwin

package ane

import (
	"errors"
	"testing"

	"github.com/tmc/apple/x/ane/mil"
)

func TestOpenMetal(t *testing.T) {
	md, err := OpenMetal()
	if err != nil {
		t.Fatal(err)
	}
	defer md.Close()

	if md.Device().GetID() == 0 {
		t.Fatal("Metal device ID is zero")
	}
}

func TestMetalBuffer(t *testing.T) {
	rt := openOrSkip(t)
	defer rt.Close()

	md, err := OpenMetal()
	if err != nil {
		t.Fatal(err)
	}
	defer md.Close()

	const channels = 1
	milText := mil.GenIdentity(channels, 1)
	blob, err := mil.BuildIdentityWeightBlob(channels)
	if err != nil {
		t.Fatal(err)
	}

	k, err := rt.Compile(CompileOptions{
		ModelType:  ModelTypeMIL,
		MILText:    []byte(milText),
		WeightBlob: blob,
	})
	if err != nil {
		t.Fatal(err)
	}
	defer k.Close()

	inBuf, err := k.MetalInputBuffer(md, 0)
	if err != nil {
		t.Fatal(err)
	}
	if inBuf.GetID() == 0 {
		t.Fatal("input MTLBuffer is nil")
	}
	t.Logf("input MTLBuffer length=%d", inBuf.Length())

	outBuf, err := k.MetalOutputBuffer(md, 0)
	if err != nil {
		t.Fatal(err)
	}
	if outBuf.GetID() == 0 {
		t.Fatal("output MTLBuffer is nil")
	}
	t.Logf("output MTLBuffer length=%d", outBuf.Length())
}

func TestMetalSharedEvent(t *testing.T) {
	rt := openOrSkip(t)
	defer rt.Close()

	md, err := OpenMetal()
	if err != nil {
		t.Fatal(err)
	}
	defer md.Close()

	ev, err := NewSharedEvent()
	if err != nil {
		t.Fatal(err)
	}
	defer ev.Close()

	port := ev.Port()
	if port == 0 {
		t.Fatal("shared event port is 0")
	}

	mtlEv, err := md.MetalSharedEvent(ev)
	if err != nil {
		t.Fatal(err)
	}
	if mtlEv.GetID() == 0 {
		t.Fatal("Metal shared event is nil")
	}
	t.Logf("Metal shared event signaled value=%d", mtlEv.SignaledValue())

	// Test NewMetalSharedEvent round-trip.
	mtlEv2, aneEv2, err := md.NewMetalSharedEvent()
	if err != nil {
		t.Fatal(err)
	}
	defer aneEv2.Close()

	if mtlEv2.GetID() == 0 {
		t.Fatal("NewMetalSharedEvent: Metal event is nil")
	}
	if aneEv2.Port() == 0 {
		t.Fatal("NewMetalSharedEvent: ANE event port is 0")
	}
	_ = rt // keep compiler happy
}

func TestEvalAsync(t *testing.T) {
	rt := openOrSkip(t)
	defer rt.Close()

	const channels = 1
	milText := mil.GenIdentity(channels, 1)
	blob, err := mil.BuildIdentityWeightBlob(channels)
	if err != nil {
		t.Fatal(err)
	}

	k, err := rt.Compile(CompileOptions{
		ModelType:  ModelTypeMIL,
		MILText:    []byte(milText),
		WeightBlob: blob,
	})
	if err != nil {
		t.Fatal(err)
	}
	defer k.Close()

	input := []float32{42}
	if err := k.WriteInputF32(0, input); err != nil {
		t.Fatal(err)
	}

	ch := k.EvalAsync()
	if err := <-ch; err != nil {
		t.Fatal(err)
	}

	output := make([]float32, channels)
	if err := k.ReadOutputF32(0, output); err != nil {
		t.Fatal(err)
	}

	diff := input[0] - output[0]
	if diff < -0.1 || diff > 0.1 {
		t.Errorf("output = %v, want %v", output[0], input[0])
	}
}

func TestRequestPool(t *testing.T) {
	rt := openOrSkip(t)
	defer rt.Close()

	const channels = 1
	milText := mil.GenIdentity(channels, 1)
	blob, err := mil.BuildIdentityWeightBlob(channels)
	if err != nil {
		t.Fatal(err)
	}

	k, err := rt.Compile(CompileOptions{
		ModelType:  ModelTypeMIL,
		MILText:    []byte(milText),
		WeightBlob: blob,
	})
	if err != nil {
		t.Fatal(err)
	}
	defer k.Close()

	pool, err := NewRequestPool(k, 4)
	if err != nil {
		t.Fatal(err)
	}
	defer pool.Close()

	for i := 0; i < 8; i++ {
		pr := pool.Acquire()

		// Write input via the kernel's surface (shared memory).
		input := []float32{float32(i + 1)}
		if err := k.WriteInputF32(0, input); err != nil {
			t.Fatal(err)
		}

		if err := pr.Eval(); err != nil {
			t.Fatalf("iter %d: %v", i, err)
		}
		pr.Release()

		output := make([]float32, channels)
		if err := k.ReadOutputF32(0, output); err != nil {
			t.Fatal(err)
		}

		diff := input[0] - output[0]
		if diff < -0.1 || diff > 0.1 {
			t.Errorf("iter %d: output = %v, want %v", i, output[0], input[0])
		}
	}
}

func TestRequestPoolDepthLimit(t *testing.T) {
	rt := openOrSkip(t)
	defer rt.Close()

	const channels = 1
	milText := mil.GenIdentity(channels, 1)
	blob, err := mil.BuildIdentityWeightBlob(channels)
	if err != nil {
		t.Fatal(err)
	}

	k, err := rt.Compile(CompileOptions{
		ModelType:  ModelTypeMIL,
		MILText:    []byte(milText),
		WeightBlob: blob,
	})
	if err != nil {
		t.Fatal(err)
	}
	defer k.Close()

	_, err = NewRequestPool(k, 128)
	if err == nil {
		t.Fatal("expected error for depth > 127")
	}

	_, err = NewRequestPool(k, 0)
	if err == nil {
		t.Fatal("expected error for depth < 1")
	}

	// Depth of 1 should work (reuses kernel's request).
	pool, err := NewRequestPool(k, 1)
	if err != nil {
		t.Fatal(err)
	}
	pool.Close()

	// Max valid depth.
	pool, err = NewRequestPool(k, maxPoolDepth)
	if errors.Is(err, ErrNoANE) {
		t.Skip("no ANE available")
	}
	if err != nil {
		t.Fatal(err)
	}
	pool.Close()
}
