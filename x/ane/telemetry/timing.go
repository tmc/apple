//go:build darwin

package telemetry

import (
	"time"

	"github.com/tmc/apple/x/ane"
)

// EvalWithSignalTiming evaluates the kernel with a signal event and returns
// wall-clock timing for the enqueue call.
func EvalWithSignalTiming(k *ane.Kernel, port uint32, val uint64, cfg ane.SharedEventEvalOptions) (EventTiming, error) {
	start := time.Now()
	err := k.EvalWithSignalEvent(port, val, cfg)
	ns := time.Since(start).Nanoseconds()
	return EventTiming{EnqueueNS: ns, TotalNS: ns}, err
}

// KernelTelemetry collects post-evaluation telemetry from the kernel's current state.
// Pass in the EvalStats obtained from a prior EvalWithStats call;
// this function does not re-evaluate the model.
func KernelTelemetry(k *ane.Kernel, stats EvalStats) EvalTelemetry {
	return EvalTelemetry{
		Stats:       stats,
		Diagnostics: ProbeDiagnostics(k),
	}
}

// TelemetryWithSignal evaluates the kernel with a signal event and returns
// an EvalTelemetry value with timing and diagnostics populated.
func TelemetryWithSignal(k *ane.Kernel, port uint32, val uint64, cfg ane.SharedEventEvalOptions) (EvalTelemetry, error) {
	timing, err := EvalWithSignalTiming(k, port, val, cfg)
	if err != nil {
		return EvalTelemetry{}, err
	}
	return EvalTelemetry{
		Diagnostics: ProbeDiagnostics(k),
		Timing:      timing,
	}, nil
}
