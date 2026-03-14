//go:build darwin

package telemetry

import (
	"time"

	"github.com/tmc/apple/x/ane"
)

// EvalWithSignalTiming evaluates the model with a signal event and returns
// wall-clock timing for the enqueue call.
func EvalWithSignalTiming(m *ane.Model, port uint32, val uint64, cfg ane.SharedEventEvalOptions) (EventTiming, error) {
	start := time.Now()
	err := m.EvalWithSignalEvent(port, val, cfg)
	ns := time.Since(start).Nanoseconds()
	return EventTiming{EnqueueNS: ns, TotalNS: ns}, err
}

// ModelTelemetry collects post-evaluation telemetry from the model's current state.
// Pass in the EvalStats obtained from a prior EvalWithStats call;
// this function does not re-evaluate the model.
func ModelTelemetry(m *ane.Model, stats EvalStats) EvalTelemetry {
	return EvalTelemetry{
		Stats:       stats,
		Diagnostics: ProbeDiagnostics(m),
	}
}

// TelemetryWithSignal evaluates the model with a signal event and returns
// an EvalTelemetry value with timing and diagnostics populated.
func TelemetryWithSignal(m *ane.Model, port uint32, val uint64, cfg ane.SharedEventEvalOptions) (EvalTelemetry, error) {
	timing, err := EvalWithSignalTiming(m, port, val, cfg)
	if err != nil {
		return EvalTelemetry{}, err
	}
	return EvalTelemetry{
		Diagnostics: ProbeDiagnostics(m),
		Timing:      timing,
	}, nil
}
