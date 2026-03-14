// Package ane provides high-level access to the Apple Neural Engine.
//
// Open a [Runtime], compile MIL programs into [Kernel] values, and evaluate
// them on the ANE hardware. IOSurface buffers are managed automatically.
//
//	rt, err := ane.Open()
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer rt.Close()
//
//	k, err := rt.Compile(ane.CompileOptions{
//		MILText:    milText,
//		WeightBlob: blob,
//		ModelType:  ane.ModelTypeMIL,
//	})
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer k.Close()
//
//	k.WriteInputF32(0, input)
//	if err := k.Eval(); err != nil {
//		log.Fatal(err)
//	}
//	k.ReadOutputF32(0, output)
//
// # Compilation
//
// There are two compilation paths. [ModelTypeMIL] compiles MIL text and
// weights in memory:
//
//	MIL text + weights → ANEInMemoryModel → compile → load → Kernel
//
// [ModelTypePackage] loads a pre-compiled .mlmodelc package from disk:
//
//	.mlmodelc path → ANEModel → compile → load → Kernel
//
// After compilation the ANE reports its expected memory layout via model
// attributes. IOSurfaces are created to match these layouts exactly.
// Surface sizes, channels, and spatial dimensions are never specified
// manually — they are parsed from the compiled model.
//
// [Runtime.CompileWithStats] returns a [CompileStats] with wall-clock
// compile and load phase durations.
//
// # Tensor Layout
//
// Data is in channel-first (NCHW) order: data[c*width + x].
// The byte offset in the IOSurface is c*PlaneStride + x*ElemSize.
// The minimum row stride is 64 bytes (the ANE's alignment granularity).
//
// [TensorLayout] describes the memory layout for each input and output.
// AllocSize (Channels * PlaneStride) includes stride padding and is
// always >= the logical data size. Typed I/O methods ([Kernel.WriteInputF32],
// [Kernel.ReadOutputF32], etc.) accept logical element counts and handle
// stride padding internally. Raw I/O ([Kernel.WriteInput],
// [Kernel.ReadOutput]) requires exactly AllocSize bytes.
//
// # Shared Events
//
// [SharedEvent] enables ANE↔GPU/CPU synchronization via IOSurface
// shared events. Use [Kernel.EvalWithSignalEvent],
// [Kernel.EvalBidirectional], and related methods for pipelined
// evaluation across compute domains.
//
// # Telemetry
//
// For hardware performance counters, diagnostics, and runtime
// snapshots, see the [github.com/tmc/apple/x/ane/telemetry] package.
package ane
