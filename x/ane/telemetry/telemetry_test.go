//go:build darwin

package telemetry_test

import (
	"errors"
	"testing"

	"github.com/tmc/apple/x/ane"
	"github.com/tmc/apple/x/ane/mil"
	"github.com/tmc/apple/x/ane/telemetry"
)

func openOrSkip(t *testing.T) *ane.Runtime {
	t.Helper()
	rt, err := ane.Open()
	if errors.Is(err, ane.ErrNoANE) {
		t.Skip("no ANE available")
	}
	if err != nil {
		t.Fatal(err)
	}
	return rt
}

func TestEvalWithStats(t *testing.T) {
	rt := openOrSkip(t)
	defer rt.Close()

	const channels = 1
	milText := mil.GenIdentity(channels, 1)
	blob, err := mil.BuildIdentityWeightBlob(channels)
	if err != nil {
		t.Fatal(err)
	}

	k, err := rt.Compile(ane.CompileOptions{
		ModelType:     ane.ModelTypeMIL,
		MILText:       []byte(milText),
		WeightBlob:    blob,
		PerfStatsMask: 0xF,
	})
	if err != nil {
		t.Fatal(err)
	}
	defer k.Close()

	input := []float32{7}
	if err := k.WriteInputF32(0, input); err != nil {
		t.Fatal(err)
	}

	stats, err := telemetry.EvalWithStats(k)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("HWExecutionNS=%d PerfCounterData=%d bytes RawStatsData=%d bytes PerfCounters=%d PerfCountersTruncated=%v PerfStatsEntries=%d",
		stats.HWExecutionNS, len(stats.PerfCounterData), len(stats.RawStatsData), len(stats.PerfCounters), stats.PerfCountersTruncated, stats.PerfStatsEntries)
	for _, pc := range stats.PerfCounters {
		t.Logf("  counter[%d] = %s (value=%d)", pc.Index, pc.Name, pc.Value)
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

func TestDiagnosticsExtended(t *testing.T) {
	rt := openOrSkip(t)
	defer rt.Close()

	const channels = 1
	milText := mil.GenIdentity(channels, 1)
	blob, err := mil.BuildIdentityWeightBlob(channels)
	if err != nil {
		t.Fatal(err)
	}

	k, err := rt.Compile(ane.CompileOptions{
		ModelType:  ane.ModelTypeMIL,
		MILText:    []byte(milText),
		WeightBlob: blob,
	})
	if err != nil {
		t.Fatal(err)
	}
	defer k.Close()

	d := telemetry.ProbeDiagnostics(k)
	t.Logf("ProgramClass=%s (known=%v)", d.ProgramClass, d.ProgramClassKnown)
	t.Logf("ModelQueueDepth=%d (known=%v)", d.ModelQueueDepth, d.ModelQueueDepthKnown)
	t.Logf("AsyncRequestsInFlight=%d (known=%v)", d.AsyncRequestsInFlight, d.AsyncRequestsInFlightKnown)
	t.Logf("ProgramQueueDepth=%d (known=%v)", d.ProgramQueueDepth, d.ProgramQueueDepthKnown)
	t.Logf("ProgramHandle=0x%x (known=%v)", d.ProgramHandle, d.ProgramHandleKnown)
	t.Logf("ModelState=%d (known=%v)", d.ModelState, d.ModelStateKnown)

	if !d.ProgramClassKnown || d.ProgramClass != "ANEInMemoryModel" {
		t.Error("expected ProgramClass=ANEInMemoryModel for MIL model")
	}
}

func TestClientInfo(t *testing.T) {
	rt := openOrSkip(t)
	defer rt.Close()

	ci := telemetry.ProbeClientInfo(rt)
	t.Logf("NumANEs=%d (known=%v) NumCores=%d (known=%v) Arch=%s (known=%v) BoardType=%d (known=%v) CapabilityMask=0x%x (known=%v) DataInterfaceVersion=%d (known=%v) PrecompiledBinarySupported=%v (known=%v)",
		ci.NumANEs, ci.NumANEsKnown,
		ci.NumCores, ci.NumCoresKnown,
		ci.ArchitectureStr, ci.ArchitectureStrKnown,
		ci.BoardType, ci.BoardTypeKnown,
		ci.CapabilityMask, ci.CapabilityMaskKnown,
		ci.DataInterfaceVersion, ci.DataInterfaceVersionKnown,
		ci.PrecompiledBinarySupported, ci.PrecompiledBinarySupportedKnown)

	if !ci.NumCoresKnown {
		t.Log("NumCores not known (VirtualClient may be unavailable)")
	}
}

func TestCacheInfo(t *testing.T) {
	rt := openOrSkip(t)
	defer rt.Close()

	ci := telemetry.ProbeCacheInfo(rt)
	t.Logf("CacheDir=%s (known=%v)", ci.CacheDir, ci.CacheDirKnown)
}

func TestSnapshot(t *testing.T) {
	rt := openOrSkip(t)
	defer rt.Close()

	snap := telemetry.Snapshot(rt)
	if !snap.Device.HasANE {
		t.Error("snapshot device should report HasANE=true")
	}
	t.Logf("Device: Cores=%d Arch=%s Product=%s", snap.Device.NumCores, snap.Device.Architecture, snap.Device.Product)
	t.Logf("Client: Available=%v NumCores=%d (known=%v)", snap.Client.Available(), snap.Client.NumCores, snap.Client.NumCoresKnown)
	t.Logf("Cache: Available=%v Dir=%s", snap.Cache.Available(), snap.Cache.CacheDir)
}

func TestTelemetry(t *testing.T) {
	rt := openOrSkip(t)
	defer rt.Close()

	const channels = 1
	milText := mil.GenIdentity(channels, 1)
	blob, err := mil.BuildIdentityWeightBlob(channels)
	if err != nil {
		t.Fatal(err)
	}

	k, err := rt.Compile(ane.CompileOptions{
		ModelType:     ane.ModelTypeMIL,
		MILText:       []byte(milText),
		WeightBlob:    blob,
		PerfStatsMask: 0xF,
	})
	if err != nil {
		t.Fatal(err)
	}
	defer k.Close()

	if err := k.WriteInputF32(0, []float32{7}); err != nil {
		t.Fatal(err)
	}

	stats, err := telemetry.EvalWithStats(k)
	if err != nil {
		t.Fatal(err)
	}
	tel := telemetry.KernelTelemetry(k, stats)
	t.Logf("Stats.Available=%v Diagnostics.Available=%v Timing.Available=%v",
		tel.Stats.Available(), tel.Diagnostics.Available(), tel.Timing.Available())
	t.Logf("HWExecutionNS=%d ProgramClass=%s", tel.Stats.HWExecutionNS, tel.Diagnostics.ProgramClass)

	if !tel.Diagnostics.Available() {
		t.Error("expected diagnostics to be available")
	}
	if tel.Diagnostics.ProgramClass != "ANEInMemoryModel" {
		t.Errorf("expected ProgramClass=ANEInMemoryModel, got %s", tel.Diagnostics.ProgramClass)
	}
	if tel.Timing.Available() {
		t.Error("Timing should not be available without shared-event eval")
	}
}

func TestAvailableMethods(t *testing.T) {
	if (telemetry.ClientInfo{}).Available() {
		t.Error("zero ClientInfo should not be available")
	}
	if (telemetry.CacheInfo{}).Available() {
		t.Error("zero CacheInfo should not be available")
	}
	if (telemetry.EvalStats{}).Available() {
		t.Error("zero EvalStats should not be available")
	}
	if (telemetry.EventTiming{}).Available() {
		t.Error("zero EventTiming should not be available")
	}
	if (telemetry.Diagnostics{}).Available() {
		t.Error("zero Diagnostics should not be available")
	}
	if (telemetry.RuntimeSnapshot{}).Available() {
		t.Error("zero RuntimeSnapshot should not be available")
	}
	if (telemetry.EvalTelemetry{}).Available() {
		t.Error("zero EvalTelemetry should not be available")
	}

	if !(telemetry.ClientInfo{NumCoresKnown: true}).Available() {
		t.Error("ClientInfo with NumCoresKnown should be available")
	}
	if !(telemetry.CacheInfo{CacheDirKnown: true}).Available() {
		t.Error("CacheInfo with CacheDirKnown should be available")
	}
	if !(telemetry.EvalStats{HWExecutionNS: 1}).Available() {
		t.Error("EvalStats with HWExecutionNS should be available")
	}
	if !(telemetry.EventTiming{TotalNS: 1}).Available() {
		t.Error("EventTiming with TotalNS should be available")
	}
	if !(telemetry.Diagnostics{ProgramClassKnown: true}).Available() {
		t.Error("Diagnostics with ProgramClassKnown should be available")
	}
	if !(telemetry.RuntimeSnapshot{Device: ane.DeviceInfo{HasANE: true}}).Available() {
		t.Error("RuntimeSnapshot with HasANE should be available")
	}
	if !(telemetry.EvalTelemetry{Diagnostics: telemetry.Diagnostics{ProgramClassKnown: true}}).Available() {
		t.Error("EvalTelemetry with Diagnostics should be available")
	}
	if !(telemetry.EvalStats{RawStatsData: []byte{1}}).Available() {
		t.Error("EvalStats with RawStatsData should be available")
	}
	if !(telemetry.EvalStats{PerfStatsEntries: 1}).Available() {
		t.Error("EvalStats with PerfStatsEntries should be available")
	}
}

// metricSink captures ReportMetric calls for testing.
type metricSink struct {
	metrics []metricEntry
}

type metricEntry struct {
	value float64
	name  string
}

func (s *metricSink) ReportMetric(value float64, name string) {
	s.metrics = append(s.metrics, metricEntry{value, name})
}

func (s *metricSink) has(name string) bool {
	for _, m := range s.metrics {
		if m.name == name {
			return true
		}
	}
	return false
}

func (s *metricSink) get(name string) (float64, bool) {
	for _, m := range s.metrics {
		if m.name == name {
			return m.value, true
		}
	}
	return 0, false
}

func TestReportMetricsEvalStats(t *testing.T) {
	var sink metricSink
	stats := telemetry.EvalStats{
		HWExecutionNS:         1234,
		PerfCounterData:       []byte{1, 2, 3},
		RawStatsData:          []byte{4, 5},
		PerfCountersTruncated: true,
		PerfCounters: []telemetry.PerfCounter{
			{Index: 0, Name: "cycles", Value: 100},
			{Index: 1, Name: "", Value: 200},
		},
	}
	stats.ReportMetrics(&sink)

	if v, ok := sink.get("hw-ns/op"); !ok || v != 1234 {
		t.Errorf("hw-ns/op = %v, %v", v, ok)
	}
	if v, ok := sink.get("cycles/op"); !ok || v != 100 {
		t.Errorf("cycles/op = %v, %v", v, ok)
	}
	if v, ok := sink.get("perf-counter-1/op"); !ok || v != 200 {
		t.Errorf("perf-counter-1/op = %v, %v (unnamed fallback)", v, ok)
	}
	if v, ok := sink.get("perf-counter-bytes/op"); !ok || v != 3 {
		t.Errorf("perf-counter-bytes/op = %v, %v", v, ok)
	}
	if v, ok := sink.get("raw-stats-bytes/op"); !ok || v != 2 {
		t.Errorf("raw-stats-bytes/op = %v, %v", v, ok)
	}
	if v, ok := sink.get("perf-counters-truncated/op"); !ok || v != 1 {
		t.Errorf("perf-counters-truncated/op = %v, %v", v, ok)
	}
}

func TestReportMetricsDiagnostics(t *testing.T) {
	var sink metricSink
	d := telemetry.Diagnostics{
		ModelQueueDepth: 127, ModelQueueDepthKnown: true,
		AsyncRequestsInFlight: 3, AsyncRequestsInFlightKnown: true,
		ProgramHandle: 0xff, ProgramHandleKnown: true,
		ModelState: 2, ModelStateKnown: true,
	}
	d.ReportMetrics(&sink)

	if v, ok := sink.get("model-queue-depth"); !ok || v != 127 {
		t.Errorf("model-queue-depth = %v, %v", v, ok)
	}
	if v, ok := sink.get("async-in-flight"); !ok || v != 3 {
		t.Errorf("async-in-flight = %v, %v", v, ok)
	}
	if v, ok := sink.get("program-handle"); !ok || v != 0xff {
		t.Errorf("program-handle = %v, %v", v, ok)
	}
	if v, ok := sink.get("model-state"); !ok || v != 2 {
		t.Errorf("model-state = %v, %v", v, ok)
	}

	var sink2 metricSink
	(telemetry.Diagnostics{}).ReportMetrics(&sink2)
	if len(sink2.metrics) != 0 {
		t.Errorf("zero Diagnostics should emit no metrics, got %d", len(sink2.metrics))
	}
}

func TestReportMetricsEventTiming(t *testing.T) {
	var sink metricSink
	et := telemetry.EventTiming{EnqueueNS: 500, TotalNS: 1000}
	et.ReportMetrics(&sink)

	if v, ok := sink.get("enqueue-ns/op"); !ok || v != 500 {
		t.Errorf("enqueue-ns/op = %v, %v", v, ok)
	}
	if v, ok := sink.get("total-ns/op"); !ok || v != 1000 {
		t.Errorf("total-ns/op = %v, %v", v, ok)
	}

	var sink2 metricSink
	(telemetry.EventTiming{}).ReportMetrics(&sink2)
	if len(sink2.metrics) != 0 {
		t.Errorf("zero EventTiming should emit no metrics, got %d", len(sink2.metrics))
	}
}

func TestReportMetricsRuntimeSnapshot(t *testing.T) {
	var sink metricSink
	snap := telemetry.RuntimeSnapshot{
		Device: ane.DeviceInfo{HasANE: true, NumCores: 16, NumANEs: 1, BoardType: 272, IsVM: false},
		Client: telemetry.ClientInfo{NumCores: 16, NumCoresKnown: true, CapabilityMask: 0xff, CapabilityMaskKnown: true},
	}
	snap.ReportMetrics(&sink)

	if v, ok := sink.get("ane-cores"); !ok || v != 16 {
		t.Errorf("ane-cores = %v, %v", v, ok)
	}
	if v, ok := sink.get("ane-count"); !ok || v != 1 {
		t.Errorf("ane-count = %v, %v", v, ok)
	}
	if sink.has("ane-is-vm") {
		t.Error("ane-is-vm should not be emitted when IsVM=false")
	}
	if v, ok := sink.get("ane-client-cores"); !ok || v != 16 {
		t.Errorf("ane-client-cores = %v, %v", v, ok)
	}
	if v, ok := sink.get("ane-capability-mask"); !ok || v != 0xff {
		t.Errorf("ane-capability-mask = %v, %v", v, ok)
	}
}

func TestReportMetricsEvalTelemetry(t *testing.T) {
	var sink metricSink
	tel := telemetry.EvalTelemetry{
		Stats:       telemetry.EvalStats{HWExecutionNS: 42},
		Diagnostics: telemetry.Diagnostics{ModelQueueDepth: 10, ModelQueueDepthKnown: true},
		Timing:      telemetry.EventTiming{EnqueueNS: 100, TotalNS: 200},
	}
	tel.ReportMetrics(&sink)

	if !sink.has("hw-ns/op") {
		t.Error("EvalTelemetry.ReportMetrics should include EvalStats metrics")
	}
	if !sink.has("model-queue-depth") {
		t.Error("EvalTelemetry.ReportMetrics should include Diagnostics metrics")
	}
	if !sink.has("enqueue-ns/op") {
		t.Error("EvalTelemetry.ReportMetrics should include EventTiming metrics")
	}
}

func TestStringMethods(t *testing.T) {
	if s := (telemetry.Diagnostics{}).String(); s != "Diagnostics{}" {
		t.Errorf("zero Diagnostics.String() = %q", s)
	}
	if s := (telemetry.EventTiming{}).String(); s != "EventTiming{}" {
		t.Errorf("zero EventTiming.String() = %q", s)
	}
	if s := (telemetry.RuntimeSnapshot{}).String(); s != "RuntimeSnapshot{}" {
		t.Errorf("zero RuntimeSnapshot.String() = %q", s)
	}
	if s := (telemetry.EvalTelemetry{}).String(); s != "EvalTelemetry{}" {
		t.Errorf("zero EvalTelemetry.String() = %q", s)
	}

	d := telemetry.Diagnostics{ProgramClass: "ANEInMemoryModel", ProgramClassKnown: true, ModelState: 3, ModelStateKnown: true}
	if s := d.String(); s != "Diagnostics{class=ANEInMemoryModel state=3}" {
		t.Errorf("Diagnostics.String() = %q", s)
	}

	et := telemetry.EventTiming{EnqueueNS: 500, TotalNS: 1000}
	if s := et.String(); s != "EventTiming{enqueue=500ns total=1000ns}" {
		t.Errorf("EventTiming.String() = %q", s)
	}
}

func TestSanitizeMetricName(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"simple", "simple"},
		{"CamelCase", "camelcase"},
		{"snake_case", "snake-case"},
		{"with spaces", "with-spaces"},
		{"ANE.HW.Cycles", "ane-hw-cycles"},
		{"perf_counter__double", "perf-counter-double"},
		{"  leading_trailing  ", "leading-trailing"},
		{"MixOf.Everything_Here 123", "mixof-everything-here-123"},
		{"kANEFPerfCounter", "kanefperfcounter"},
		{"", ""},
		{"---", ""},
		{"a", "a"},
		{"123", "123"},
	}
	for _, tt := range tests {
		got := telemetry.SanitizeMetricName(tt.in)
		if got != tt.want {
			t.Errorf("SanitizeMetricName(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestReportMetricsUnnamedCounterFallback(t *testing.T) {
	var sink metricSink
	stats := telemetry.EvalStats{
		PerfCounters: []telemetry.PerfCounter{
			{Index: 0, Name: "ANE.HW.Cycles", Value: 100},
			{Index: 1, Name: "", Value: 200},
			{Index: 5, Name: "perf_counter__weird", Value: 300},
		},
	}
	stats.ReportMetrics(&sink)

	if v, ok := sink.get("ane-hw-cycles/op"); !ok || v != 100 {
		t.Errorf("expected sanitized named counter: got %v, %v", v, ok)
	}
	if v, ok := sink.get("perf-counter-1/op"); !ok || v != 200 {
		t.Errorf("expected unnamed fallback: got %v, %v", v, ok)
	}
	if v, ok := sink.get("perf-counter-weird/op"); !ok || v != 300 {
		t.Errorf("expected sanitized weird counter: got %v, %v", v, ok)
	}
}

func TestPerfStatsEntries(t *testing.T) {
	if !(telemetry.EvalStats{PerfStatsEntries: 1}).Available() {
		t.Error("EvalStats with PerfStatsEntries should be available")
	}

	var sink metricSink
	stats := telemetry.EvalStats{PerfStatsEntries: 3}
	stats.ReportMetrics(&sink)
	if v, ok := sink.get("perf-stats-entries/op"); !ok || v != 3 {
		t.Errorf("perf-stats-entries/op = %v, %v", v, ok)
	}

	var sink2 metricSink
	(telemetry.EvalStats{}).ReportMetrics(&sink2)
	if sink2.has("perf-stats-entries/op") {
		t.Error("zero PerfStatsEntries should not be reported")
	}
}

func TestRuntimeSnapshotReportMetricsBroadened(t *testing.T) {
	var sink metricSink
	snap := telemetry.RuntimeSnapshot{
		Device: ane.DeviceInfo{HasANE: true, NumCores: 16, NumANEs: 1, BoardType: 272, IsVM: true, InternalBuild: true},
		Client: telemetry.ClientInfo{
			NumCores: 16, NumCoresKnown: true,
			PrecompiledBinarySupported: true, PrecompiledBinarySupportedKnown: true,
		},
		Cache: telemetry.CacheInfo{CacheDir: "/tmp/cache", CacheDirKnown: true},
	}
	snap.ReportMetrics(&sink)

	if v, ok := sink.get("ane-is-vm"); !ok || v != 1 {
		t.Errorf("ane-is-vm = %v, %v", v, ok)
	}
	if v, ok := sink.get("ane-internal-build"); !ok || v != 1 {
		t.Errorf("ane-internal-build = %v, %v", v, ok)
	}
	if v, ok := sink.get("ane-precompiled-binary-supported"); !ok || v != 1 {
		t.Errorf("ane-precompiled-binary-supported = %v, %v", v, ok)
	}
	if v, ok := sink.get("ane-cache-available"); !ok || v != 1 {
		t.Errorf("ane-cache-available = %v, %v", v, ok)
	}

	var sink2 metricSink
	snap2 := telemetry.RuntimeSnapshot{
		Device: ane.DeviceInfo{HasANE: true, NumCores: 8},
	}
	snap2.ReportMetrics(&sink2)
	if sink2.has("ane-is-vm") {
		t.Error("ane-is-vm should not be emitted when IsVM=false")
	}
	if sink2.has("ane-internal-build") {
		t.Error("ane-internal-build should not be emitted when InternalBuild=false")
	}
	if sink2.has("ane-precompiled-binary-supported") {
		t.Error("ane-precompiled-binary-supported should not be emitted when not known")
	}
	if sink2.has("ane-cache-available") {
		t.Error("ane-cache-available should not be emitted when not known")
	}
}

func TestPerfCounterPlaceholderFiltering(t *testing.T) {
	type entry struct {
		name string
	}
	names := []entry{
		{"kANE_AF_TO_L2_DATA"},
		{"kANE_FP16_CYCLES"},
		{"kANE_NE_COMPUTE_CYCLES"},
		{telemetry.PerfCounterPlaceholder()},
		{"garbage_beyond"},
	}

	var counters []telemetry.PerfCounter
	var truncated bool

	const scanCap = 256
	for i := range min(scanCap, len(names)) {
		name := names[i].name
		if name == "" || name == telemetry.PerfCounterPlaceholder() {
			break
		}
		counters = append(counters, telemetry.PerfCounter{Index: i, Name: name})
	}
	if len(counters) == scanCap {
		if scanCap < len(names) {
			name := names[scanCap].name
			if name != "" && name != telemetry.PerfCounterPlaceholder() {
				truncated = true
			}
		}
	}

	if len(counters) != 3 {
		t.Errorf("expected 3 counters, got %d", len(counters))
	}
	if truncated {
		t.Error("should not be truncated when stopped at placeholder")
	}

	want := []string{"kANE_AF_TO_L2_DATA", "kANE_FP16_CYCLES", "kANE_NE_COMPUTE_CYCLES"}
	for i, c := range counters {
		if c.Name != want[i] {
			t.Errorf("counter[%d] = %q, want %q", i, c.Name, want[i])
		}
		if c.Index != i {
			t.Errorf("counter[%d].Index = %d, want %d", i, c.Index, i)
		}
	}
}

func TestPerfCounterPlaceholderNotReported(t *testing.T) {
	var sink metricSink
	stats := telemetry.EvalStats{
		PerfCounters: []telemetry.PerfCounter{
			{Index: 0, Name: "kANE_FP16_CYCLES", Value: 42},
		},
	}
	stats.ReportMetrics(&sink)

	if !sink.has("kane-fp16-cycles/op") {
		t.Error("expected sanitized counter name")
	}
}
