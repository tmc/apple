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
// # Compilation Timing
//
// [Runtime.CompileWithStats] returns a [CompileStats] with wall-clock
// compile and load phase durations. CompileStats has a ReportMetrics
// method compatible with [testing.B].
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
