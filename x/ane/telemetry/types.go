package telemetry

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/tmc/apple/x/ane"
)

// perfCounterPlaceholder is the firmware string used for undefined counter
// slots. The ANE driver returns this (note the misspelling) for indices
// beyond the set of real hardware counters. Counter enumeration stops at
// the first occurrence.
const perfCounterPlaceholder = "kANE_UKNOWN"

// PerfCounterPlaceholder returns the firmware placeholder sentinel string.
func PerfCounterPlaceholder() string { return perfCounterPlaceholder }

// PerfCounter represents a named hardware performance counter from an evaluation.
type PerfCounter struct {
	Index int
	Name  string
	Value uint64
}

// EvalStats contains hardware performance statistics from an evaluation.
type EvalStats struct {
	HWExecutionNS         uint64        // hardware execution time in nanoseconds
	PerfCounterData       []byte        // raw bytes from perfStats.PerfCounterData()
	RawStatsData          []byte        // raw bytes from perfStats.PStatsRawData()
	PerfCounters          []PerfCounter // parsed named counters
	PerfCountersTruncated bool          // true if counter list was capped
	PerfStatsEntries      int           // number of entries in the PerfStatsArray (one per procedure)
}

// Available reports whether any hardware stats data was populated.
func (s EvalStats) Available() bool {
	return s.HWExecutionNS > 0 || len(s.PerfCounterData) > 0 || len(s.RawStatsData) > 0 ||
		len(s.PerfCounters) > 0 || s.PerfStatsEntries > 0
}

// ReportMetrics reports evaluation statistics to a testing.B-compatible reporter.
func (s EvalStats) ReportMetrics(b interface{ ReportMetric(float64, string) }) {
	if s.HWExecutionNS > 0 {
		b.ReportMetric(float64(s.HWExecutionNS), "hw-ns/op")
	}
	for _, pc := range s.PerfCounters {
		name := pc.Name
		if name == "" {
			name = fmt.Sprintf("perf-counter-%d", pc.Index)
		} else {
			name = SanitizeMetricName(name)
		}
		b.ReportMetric(float64(pc.Value), name+"/op")
	}
	if len(s.PerfCounterData) > 0 {
		b.ReportMetric(float64(len(s.PerfCounterData)), "perf-counter-bytes/op")
	}
	if len(s.RawStatsData) > 0 {
		b.ReportMetric(float64(len(s.RawStatsData)), "raw-stats-bytes/op")
	}
	if s.PerfCountersTruncated {
		b.ReportMetric(1, "perf-counters-truncated/op")
	}
	if s.PerfStatsEntries > 0 {
		b.ReportMetric(float64(s.PerfStatsEntries), "perf-stats-entries/op")
	}
}

// String returns a compact human-readable summary of the evaluation stats.
func (s EvalStats) String() string {
	if !s.Available() {
		return "EvalStats{}"
	}
	return fmt.Sprintf("EvalStats{hw=%dns counters=%d entries=%d}",
		s.HWExecutionNS, len(s.PerfCounters), s.PerfStatsEntries)
}

// Diagnostics contains best-effort runtime diagnostic information about a kernel.
type Diagnostics struct {
	ModelQueueDepth      int
	ModelQueueDepthKnown bool
	ProgramClass         string
	ProgramClassKnown    bool

	AsyncRequestsInFlight      int64
	AsyncRequestsInFlightKnown bool
	ProgramQueueDepth          int
	ProgramQueueDepthKnown     bool
	ProgramHandle              uint64
	ProgramHandleKnown         bool
	ModelState                 uint64
	ModelStateKnown            bool
}

// Available reports whether any diagnostic field was probed.
func (d Diagnostics) Available() bool {
	return d.ModelQueueDepthKnown || d.ProgramClassKnown || d.AsyncRequestsInFlightKnown ||
		d.ProgramQueueDepthKnown || d.ProgramHandleKnown || d.ModelStateKnown
}

// ReportMetrics reports diagnostic state to a testing.B-compatible reporter.
func (d Diagnostics) ReportMetrics(b interface{ ReportMetric(float64, string) }) {
	if d.ModelQueueDepthKnown {
		b.ReportMetric(float64(d.ModelQueueDepth), "model-queue-depth")
	}
	if d.AsyncRequestsInFlightKnown {
		b.ReportMetric(float64(d.AsyncRequestsInFlight), "async-in-flight")
	}
	if d.ProgramQueueDepthKnown {
		b.ReportMetric(float64(d.ProgramQueueDepth), "program-queue-depth")
	}
	if d.ProgramHandleKnown {
		b.ReportMetric(float64(d.ProgramHandle), "program-handle")
	}
	if d.ModelStateKnown {
		b.ReportMetric(float64(d.ModelState), "model-state")
	}
}

// String returns a compact human-readable summary of the diagnostics.
func (d Diagnostics) String() string {
	if !d.Available() {
		return "Diagnostics{}"
	}
	var parts []string
	if d.ProgramClassKnown {
		parts = append(parts, fmt.Sprintf("class=%s", d.ProgramClass))
	}
	if d.ModelQueueDepthKnown {
		parts = append(parts, fmt.Sprintf("queue=%d", d.ModelQueueDepth))
	}
	if d.AsyncRequestsInFlightKnown {
		parts = append(parts, fmt.Sprintf("inflight=%d", d.AsyncRequestsInFlight))
	}
	if d.ProgramHandleKnown {
		parts = append(parts, fmt.Sprintf("handle=0x%x", d.ProgramHandle))
	}
	if d.ModelStateKnown {
		parts = append(parts, fmt.Sprintf("state=%d", d.ModelState))
	}
	return "Diagnostics{" + strings.Join(parts, " ") + "}"
}

// ClientInfo describes the ANE client connection capabilities.
type ClientInfo struct {
	NumANEs                         uint32
	NumANEsKnown                    bool
	NumCores                        uint32
	NumCoresKnown                   bool
	ArchitectureStr                 string
	ArchitectureStrKnown            bool
	BoardType                       int64
	BoardTypeKnown                  bool
	CapabilityMask                  uint64
	CapabilityMaskKnown             bool
	DataInterfaceVersion            uint64
	DataInterfaceVersionKnown       bool
	PrecompiledBinarySupported      bool
	PrecompiledBinarySupportedKnown bool
}

// Available reports whether any Known field is true.
func (ci ClientInfo) Available() bool {
	return ci.NumANEsKnown || ci.NumCoresKnown || ci.ArchitectureStrKnown ||
		ci.BoardTypeKnown || ci.CapabilityMaskKnown || ci.DataInterfaceVersionKnown ||
		ci.PrecompiledBinarySupportedKnown
}

// CacheInfo describes the ANE model cache.
type CacheInfo struct {
	CacheDir      string
	CacheDirKnown bool
}

// Available reports whether the cache directory was resolved.
func (ci CacheInfo) Available() bool { return ci.CacheDirKnown }

// EventTiming contains timing measurements from a shared event evaluation.
type EventTiming struct {
	EnqueueNS int64 // CPU wall-clock cost to submit the evaluation
	TotalNS   int64 // same as EnqueueNS for synchronous evals
}

// Available reports whether timing data was collected.
func (t EventTiming) Available() bool { return t.TotalNS > 0 }

// ReportMetrics reports event timing to a testing.B-compatible reporter.
func (t EventTiming) ReportMetrics(b interface{ ReportMetric(float64, string) }) {
	if !t.Available() {
		return
	}
	b.ReportMetric(float64(t.EnqueueNS), "enqueue-ns/op")
	b.ReportMetric(float64(t.TotalNS), "total-ns/op")
}

// String returns a compact human-readable summary of the timing.
func (t EventTiming) String() string {
	if !t.Available() {
		return "EventTiming{}"
	}
	return fmt.Sprintf("EventTiming{enqueue=%dns total=%dns}", t.EnqueueNS, t.TotalNS)
}

// ClientSnapshot captures the host and ANE environment at a point in time.
type ClientSnapshot struct {
	Device ane.DeviceInfo
	Client ClientInfo
	Cache  CacheInfo
}

// Available reports whether any component snapshot has data.
func (s ClientSnapshot) Available() bool {
	return s.Device.HasANE || s.Client.Available() || s.Cache.Available()
}

// ReportMetrics reports the runtime environment to a testing.B-compatible reporter.
func (s ClientSnapshot) ReportMetrics(b interface{ ReportMetric(float64, string) }) {
	if s.Device.HasANE {
		b.ReportMetric(float64(s.Device.NumCores), "ane-cores")
		b.ReportMetric(float64(s.Device.NumANEs), "ane-count")
		b.ReportMetric(float64(s.Device.BoardType), "ane-board-type")
		if s.Device.IsVM {
			b.ReportMetric(1, "ane-is-vm")
		}
		if s.Device.InternalBuild {
			b.ReportMetric(1, "ane-internal-build")
		}
	}
	if s.Client.NumCoresKnown {
		b.ReportMetric(float64(s.Client.NumCores), "ane-client-cores")
	}
	if s.Client.CapabilityMaskKnown {
		b.ReportMetric(float64(s.Client.CapabilityMask), "ane-capability-mask")
	}
	if s.Client.DataInterfaceVersionKnown {
		b.ReportMetric(float64(s.Client.DataInterfaceVersion), "ane-data-interface-version")
	}
	if s.Client.PrecompiledBinarySupportedKnown && s.Client.PrecompiledBinarySupported {
		b.ReportMetric(1, "ane-precompiled-binary-supported")
	}
	if s.Cache.CacheDirKnown {
		b.ReportMetric(1, "ane-cache-available")
	}
}

// String returns a compact human-readable summary of the runtime snapshot.
func (s ClientSnapshot) String() string {
	if !s.Available() {
		return "ClientSnapshot{}"
	}
	var parts []string
	if s.Device.HasANE {
		parts = append(parts, fmt.Sprintf("cores=%d", s.Device.NumCores))
		if s.Device.Architecture != "" {
			parts = append(parts, fmt.Sprintf("arch=%s", s.Device.Architecture))
		}
		if s.Device.Product != "" {
			parts = append(parts, fmt.Sprintf("product=%s", s.Device.Product))
		}
		if s.Device.IsVM {
			parts = append(parts, "vm=true")
		}
	}
	if s.Client.Available() {
		if s.Client.CapabilityMaskKnown {
			parts = append(parts, fmt.Sprintf("cap=0x%x", s.Client.CapabilityMask))
		}
	}
	if s.Cache.CacheDirKnown {
		parts = append(parts, fmt.Sprintf("cache=%s", s.Cache.CacheDir))
	}
	return "ClientSnapshot{" + strings.Join(parts, " ") + "}"
}

// EvalTelemetry bundles post-evaluation telemetry from a kernel.
type EvalTelemetry struct {
	Stats       EvalStats
	Diagnostics Diagnostics
	Timing      EventTiming // zero unless set by caller after shared-event eval
}

// Available reports whether any telemetry component has data.
func (t EvalTelemetry) Available() bool {
	return t.Stats.Available() || t.Diagnostics.Available() || t.Timing.Available()
}

// ReportMetrics reports all available telemetry to a testing.B-compatible reporter.
func (t EvalTelemetry) ReportMetrics(b interface{ ReportMetric(float64, string) }) {
	t.Stats.ReportMetrics(b)
	t.Diagnostics.ReportMetrics(b)
	t.Timing.ReportMetrics(b)
}

// String returns a compact human-readable summary of the telemetry bundle.
func (t EvalTelemetry) String() string {
	if !t.Available() {
		return "EvalTelemetry{}"
	}
	var parts []string
	if t.Stats.Available() {
		parts = append(parts, fmt.Sprintf("hw=%dns counters=%d", t.Stats.HWExecutionNS, len(t.Stats.PerfCounters)))
	}
	if t.Diagnostics.Available() {
		parts = append(parts, t.Diagnostics.String())
	}
	if t.Timing.Available() {
		parts = append(parts, t.Timing.String())
	}
	return "EvalTelemetry{" + strings.Join(parts, " ") + "}"
}

// SanitizeMetricName normalizes a raw firmware counter name into a stable,
// parseable metric name suitable for benchmark output. It lowercases the
// input, replaces non-alphanumeric characters with hyphens, collapses
// runs of hyphens, and trims leading/trailing hyphens.
func SanitizeMetricName(name string) string {
	var b strings.Builder
	b.Grow(len(name))
	prevHyphen := true // suppress leading hyphen
	for _, r := range name {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			b.WriteRune(unicode.ToLower(r))
			prevHyphen = false
		} else if !prevHyphen {
			b.WriteByte('-')
			prevHyphen = true
		}
	}
	s := b.String()
	return strings.TrimRight(s, "-")
}
