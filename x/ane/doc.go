// Package ane provides high-level access to the Apple Neural Engine.
//
// Open a [Client], compile MIL programs into [Model] values, and evaluate
// them on the ANE hardware. IOSurface buffers are managed automatically.
//
//	c, err := ane.Open()
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer c.Close()
//
//	m, err := c.Compile(ane.CompileOptions{
//		MILText:    milText,
//		WeightBlob: blob,
//		ModelType:  ane.ModelTypeMIL,
//	})
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer m.Close()
//
//	m.WriteInputF32(0, input)
//	if err := m.Eval(); err != nil {
//		log.Fatal(err)
//	}
//	m.ReadOutputF32(0, output)
//
// # Compilation
//
// There are two compilation paths. [ModelTypeMIL] compiles MIL text and
// weights in memory:
//
//	MIL text + weights → ANEInMemoryModel → compile → load → Model
//
// [ModelTypePackage] loads a pre-compiled .mlmodelc package from disk:
//
//	.mlmodelc path → ANEModel → compile → load → Model
//
// After compilation the ANE reports its expected memory layout via model
// attributes. IOSurfaces are created to match these layouts exactly.
// Surface sizes, channels, and spatial dimensions are never specified
// manually — they are parsed from the compiled model.
//
// [Client.CompileWithStats] returns a [CompileStats] with wall-clock
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
// always >= the logical data size. Typed I/O methods ([Model.WriteInputF32],
// [Model.ReadOutputF32], etc.) accept logical element counts and handle
// stride padding internally. Raw I/O ([Model.WriteInput],
// [Model.ReadOutput]) requires exactly AllocSize bytes.
//
// # Shared Events
//
// [SharedEvent] enables ANE↔GPU/CPU synchronization via IOSurface
// shared events. Use [Model.EvalWithSignalEvent],
// [Model.EvalBidirectional], and related methods for pipelined
// evaluation across compute domains.
//
// Shared events require package-backed models ([ModelTypePackage]).
// In-memory MIL models ([ModelTypeMIL]) do not support shared events;
// attempts to use EvalWithSignalEvent or EvalBidirectional on MIL models
// return an error. To use shared events, compile the model through
// [github.com/tmc/apple/x/coremlcompiler] and load the .mlmodelc
// bundle via ModelTypePackage.
//
// # Telemetry
//
// For hardware performance counters, diagnostics, and runtime
// snapshots, see the [github.com/tmc/apple/x/ane/telemetry] package.
//
// Hardware execution counters (HwExecutionTime, PerfCounters) require
// [CompileOptions].PerfStatsMask to be non-zero. For in-memory MIL
// models, hardware counters may be unavailable; the telemetry package
// returns wall-clock time as a fallback. Package-backed models provide
// full hardware counter access when running on non-VM hosts.
package ane
