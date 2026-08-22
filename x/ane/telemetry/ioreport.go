//go:build darwin

package telemetry

import (
	"fmt"
	"strings"
	"sync"
	"time"
	"unsafe"

	"github.com/ebitengine/purego"

	cf "github.com/tmc/apple/corefoundation"
)

// The DRAM byte counters come from the memory controller's per-agent
// performance counters, the "AMC Stats" / "Perf Counters" channel group
// in the private libIOReport.dylib. x/powersample reads the "Energy
// Model" group the same way; it exposes no general channel read, so the
// symbol binding is repeated here rather than shared.
//
// Measured on macOS 26.6.1 (25G76): the ANE channels enumerate without
// privileges under the names "ANE AF RD", "ANE AF WR", "ANE DCS RD",
// "ANE DCS WR" and the ANEXL equivalents, but an unprivileged process
// cannot sample them. IOReportCreateSubscription fails outright on a
// channel set drawn from "AMC Stats", and on an all-channels
// subscription — which does succeed — no "AMC Stats" channel appears in
// the samples at all. So StartDRAM returns an error on the unentitled
// path here.
//
// TODO: unverified whether root lifts this. It was not testable on the
// development host (no passwordless sudo), and the ANE channels'
// behavior under root is therefore unknown. If it does, StartDRAM's
// error should say so.
//
// Every symbol is resolved at first use and a miss reports itself rather
// than crashing.
var (
	iorOnce sync.Once
	iorErr  error

	iorCopyChannelsInGroup func(group, subgroup cf.CFStringRef, a, b, c uint64) cf.CFMutableDictionaryRef
	iorCreateSubscription  func(a uintptr, desired cf.CFMutableDictionaryRef, subbed *cf.CFMutableDictionaryRef, channelID uint64, b uintptr) uintptr
	iorCreateSamples       func(sub uintptr, subbed cf.CFMutableDictionaryRef, a uintptr) cf.CFDictionaryRef
	iorCreateSamplesDelta  func(prev, cur cf.CFDictionaryRef, a uintptr) cf.CFDictionaryRef
	iorChannelName         func(ch cf.CFDictionaryRef) cf.CFStringRef
	iorChannelUnitLabel    func(ch cf.CFDictionaryRef) cf.CFStringRef
	iorSimpleIntegerValue  func(ch cf.CFDictionaryRef, err uintptr) int64
)

func iorInit() error {
	iorOnce.Do(func() {
		lib, err := purego.Dlopen("/usr/lib/libIOReport.dylib", purego.RTLD_NOW|purego.RTLD_GLOBAL)
		if err != nil {
			iorErr = fmt.Errorf("load libIOReport: %w", err)
			return
		}
		for _, s := range []struct {
			name string
			fn   any
		}{
			{"IOReportCopyChannelsInGroup", &iorCopyChannelsInGroup},
			{"IOReportCreateSubscription", &iorCreateSubscription},
			{"IOReportCreateSamples", &iorCreateSamples},
			{"IOReportCreateSamplesDelta", &iorCreateSamplesDelta},
			{"IOReportChannelGetChannelName", &iorChannelName},
			{"IOReportChannelGetUnitLabel", &iorChannelUnitLabel},
			{"IOReportSimpleGetIntegerValue", &iorSimpleIntegerValue},
		} {
			if _, err := purego.Dlsym(lib, s.name); err != nil {
				iorErr = fmt.Errorf("resolve %s: %w", s.name, err)
				return
			}
			purego.RegisterLibFunc(s.fn, lib, s.name)
		}
	})
	return iorErr
}

// IOReport group and subgroup holding the memory-controller per-agent
// byte counters.
const (
	dramGroup    = "AMC Stats"
	dramSubgroup = "Perf Counters"
)

// DRAMBytes is a count of bytes moved between the ANE and DRAM, as
// measured by the memory controller. The counters are whole-engine and
// system-wide: they include traffic from every process using the ANE,
// not only the caller's.
type DRAMBytes struct {
	Read            uint64 // bytes read on the raw fabric path
	Write           uint64 // bytes written on the raw fabric path
	CompressedRead  uint64 // bytes read via the compression-subsystem path
	CompressedWrite uint64 // bytes written via the compression-subsystem path
}

// Total returns the sum of all four counters.
func (b DRAMBytes) Total() uint64 {
	return b.Read + b.Write + b.CompressedRead + b.CompressedWrite
}

// Available reports whether any bytes were counted.
func (b DRAMBytes) Available() bool { return b.Total() > 0 }

// ReportMetrics emits the counters to a testing.B-compatible reporter.
func (b DRAMBytes) ReportMetrics(r interface {
	ReportMetric(float64, string)
}) {
	if !b.Available() {
		return
	}
	r.ReportMetric(float64(b.Read), "dram_rd_B")
	r.ReportMetric(float64(b.Write), "dram_wr_B")
	r.ReportMetric(float64(b.CompressedRead), "dram_dcs_rd_B")
	r.ReportMetric(float64(b.CompressedWrite), "dram_dcs_wr_B")
}

// A DRAMMeter measures ANE DRAM traffic over a region. It takes one
// IOReport sample when it starts and one when it stops; there is no
// sampling loop.
//
// The zero DRAMMeter is not usable; call [StartDRAM].
type DRAMMeter struct {
	sub    uintptr
	subbed cf.CFMutableDictionaryRef
	first  cf.CFDictionaryRef
	start  time.Time
	done   bool
}

// StartDRAM begins measuring ANE DRAM traffic.
//
// It reports an error when libIOReport, its symbols, or the ANE
// memory-controller channels are unavailable, and on macOS 26.6.1
// (25G76) it always does so for an unprivileged process: the channels
// enumerate but do not subscribe (see [DRAMChannels]). A caller should
// treat the error as "unmeasurable here", not as a failure of the work
// being measured.
func StartDRAM() (*DRAMMeter, error) {
	if err := iorInit(); err != nil {
		return nil, fmt.Errorf("telemetry: dram counters: %w", err)
	}
	group := cfstr(dramGroup)
	defer cf.CFRelease(cf.CFTypeRef(group))
	sub := cfstr(dramSubgroup)
	defer cf.CFRelease(cf.CFTypeRef(sub))
	channels := iorCopyChannelsInGroup(group, sub, 0, 0, 0)
	if channels == 0 {
		return nil, fmt.Errorf("telemetry: no %q/%q channels", dramGroup, dramSubgroup)
	}
	// IOReportCopyChannelsInGroup returns a +1 reference. The subscription
	// takes what it needs from it, so it is released either way; the
	// failure branch below is the one exercised on an unprivileged host.
	defer release(cf.CFDictionaryRef(channels))
	var subbed cf.CFMutableDictionaryRef
	s := iorCreateSubscription(0, channels, &subbed, 0, 0)
	if s == 0 {
		return nil, fmt.Errorf("telemetry: IOReport subscription to %q/%q failed (unprivileged process?)", dramGroup, dramSubgroup)
	}
	first := iorCreateSamples(s, subbed, 0)
	if first == 0 {
		return nil, fmt.Errorf("telemetry: IOReport initial sample failed")
	}
	return &DRAMMeter{sub: s, subbed: subbed, first: first, start: time.Now()}, nil
}

// Stop ends the measurement and returns the bytes moved since
// [StartDRAM], with the elapsed wall-clock time. The meter cannot be
// reused. Stop reports an error when the counters did not move, which on
// an idle engine is expected and on a busy one means the channel names
// changed.
func (m *DRAMMeter) Stop() (DRAMBytes, time.Duration, error) {
	var b DRAMBytes
	if m.done {
		return b, 0, fmt.Errorf("telemetry: dram meter already stopped")
	}
	m.done = true
	elapsed := time.Since(m.start)
	// The subscription and its subscribed-channels dictionary are owned by
	// the meter and die with it. This path is not reachable on an
	// unprivileged host, where StartDRAM never returns a meter, so it is
	// unexercised here.
	defer func() {
		release(cf.CFDictionaryRef(m.subbed))
		if m.sub != 0 {
			cf.CFRelease(cf.CFTypeRef(m.sub))
		}
	}()
	cur := iorCreateSamples(m.sub, m.subbed, 0)
	if cur == 0 {
		release(m.first)
		return b, elapsed, fmt.Errorf("telemetry: IOReport final sample failed")
	}
	delta := iorCreateSamplesDelta(m.first, cur, 0)
	release(m.first)
	release(cur)
	if delta == 0 {
		return b, elapsed, fmt.Errorf("telemetry: IOReport delta failed")
	}
	defer release(delta)
	if err := accumulate(delta, &b); err != nil {
		return b, elapsed, err
	}
	if !b.Available() {
		return b, elapsed, fmt.Errorf("telemetry: no ANE DRAM bytes over %v — engine idle, or the %q channel names changed", elapsed, dramGroup)
	}
	return b, elapsed, nil
}

// accumulate adds the ANE channels of an IOReport delta dictionary into b.
func accumulate(delta cf.CFDictionaryRef, b *DRAMBytes) error {
	key := cfstr("IOReportChannels")
	defer cf.CFRelease(cf.CFTypeRef(key))
	arr := cf.CFArrayRef(uintptr(cf.CFDictionaryGetValue(delta, unsafe.Pointer(uintptr(key)))))
	if arr == 0 {
		return fmt.Errorf("telemetry: IOReport delta has no channels")
	}
	for i := range cf.CFArrayGetCount(arr) {
		ch := cf.CFDictionaryRef(uintptr(cf.CFArrayGetValueAtIndex(arr, i)))
		field := dramField(gostr(iorChannelName(ch)), b)
		if field == nil {
			continue
		}
		// The unit label guards against a rename that reuses an ANE
		// channel name for something that is not a byte count.
		if u := gostr(iorChannelUnitLabel(ch)); u != "B" {
			continue
		}
		if v := iorSimpleIntegerValue(ch, 0); v > 0 {
			*field += uint64(v)
		}
	}
	return nil
}

// dramField maps an IOReport channel name to the DRAMBytes field it
// belongs to, or nil when the channel is not an ANE DRAM byte counter.
//
// It matches "ANE" and "ANEXL" (the extra-large engine tile on some
// parts) and ignores unrelated agents whose names merely contain "ANE".
func dramField(name string, b *DRAMBytes) *uint64 {
	f := strings.Fields(name)
	if len(f) != 3 || (f[0] != "ANE" && f[0] != "ANEXL") {
		return nil
	}
	switch f[1] + " " + f[2] {
	case "AF RD":
		return &b.Read
	case "AF WR":
		return &b.Write
	case "DCS RD":
		return &b.CompressedRead
	case "DCS WR":
		return &b.CompressedWrite
	}
	return nil
}

// DRAMChannels returns the names of the ANE DRAM byte channels this
// machine reports. The channels enumerate even where they cannot be
// sampled, so a non-empty result does not imply [StartDRAM] will
// succeed; it is for diagnosing a version-brittle rename.
func DRAMChannels() ([]string, error) {
	if err := iorInit(); err != nil {
		return nil, err
	}
	group := cfstr(dramGroup)
	defer cf.CFRelease(cf.CFTypeRef(group))
	sub := cfstr(dramSubgroup)
	defer cf.CFRelease(cf.CFTypeRef(sub))
	channels := iorCopyChannelsInGroup(group, sub, 0, 0, 0)
	if channels == 0 {
		return nil, fmt.Errorf("no %q/%q channels", dramGroup, dramSubgroup)
	}
	defer release(cf.CFDictionaryRef(channels))
	key := cfstr("IOReportChannels")
	defer cf.CFRelease(cf.CFTypeRef(key))
	arr := cf.CFArrayRef(uintptr(cf.CFDictionaryGetValue(cf.CFDictionaryRef(channels), unsafe.Pointer(uintptr(key)))))
	if arr == 0 {
		return nil, fmt.Errorf("no channel list")
	}
	var names []string
	var scratch DRAMBytes
	for i := range cf.CFArrayGetCount(arr) {
		ch := cf.CFDictionaryRef(uintptr(cf.CFArrayGetValueAtIndex(arr, i)))
		n := gostr(iorChannelName(ch))
		if dramField(n, &scratch) != nil {
			names = append(names, n)
		}
	}
	return names, nil
}

func cfstr(s string) cf.CFStringRef {
	return cf.CFStringCreateWithCString(0, s, uint32(cf.KCFStringEncodingUTF8))
}

func gostr(s cf.CFStringRef) string {
	if s == 0 {
		return ""
	}
	buf := make([]byte, 256)
	if !cf.CFStringGetCString(s, &buf[0], len(buf), uint32(cf.KCFStringEncodingUTF8)) {
		return ""
	}
	if i := strings.IndexByte(string(buf), 0); i >= 0 {
		return string(buf[:i])
	}
	return string(buf)
}

func release(d cf.CFDictionaryRef) {
	if d != 0 {
		cf.CFRelease(cf.CFTypeRef(d))
	}
}
