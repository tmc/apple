//go:build darwin

package powersample

import (
	"fmt"
	"strings"
	"sync"
	"time"
	"unsafe"

	"github.com/ebitengine/purego"

	cf "github.com/tmc/apple/corefoundation"
)

// The IOReport backend reads the SoC's cumulative energy counters (the
// "Energy Model" channel group in libIOReport.dylib) directly — the same
// source powermetrics aggregates — and needs no privileges at all. It is
// private API: every symbol is resolved at first use and a miss reports
// itself rather than crashing.
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

// iorMeter accumulates energy deltas from periodic IOReport samples. The
// counters are cumulative, so sampling exists only to bound counter-wrap
// exposure; the sum of window deltas equals the whole-run delta.
type iorMeter struct {
	sub    uintptr
	subbed cf.CFMutableDictionaryRef
	prev   cf.CFDictionaryRef
	start  time.Time

	quit chan struct{}
	done chan struct{}

	mu      sync.Mutex
	energy  Power // joules
	samples int
}

func startIOReport(interval time.Duration) (*iorMeter, error) {
	if err := iorInit(); err != nil {
		return nil, err
	}
	group := cfstr("Energy Model")
	defer cf.CFRelease(cf.CFTypeRef(group))
	channels := iorCopyChannelsInGroup(group, 0, 0, 0, 0)
	if channels == 0 {
		return nil, fmt.Errorf("no Energy Model channels (not Apple silicon?)")
	}
	var subbed cf.CFMutableDictionaryRef
	sub := iorCreateSubscription(0, channels, &subbed, 0, 0)
	if sub == 0 {
		return nil, fmt.Errorf("IOReport subscription failed")
	}
	first := iorCreateSamples(sub, subbed, 0)
	if first == 0 {
		return nil, fmt.Errorf("IOReport initial sample failed")
	}
	m := &iorMeter{
		sub:    sub,
		subbed: subbed,
		prev:   first,
		start:  time.Now(),
		quit:   make(chan struct{}),
		done:   make(chan struct{}),
	}
	go func() {
		defer close(m.done)
		t := time.NewTicker(interval)
		defer t.Stop()
		for {
			select {
			case <-t.C:
				m.sample()
			case <-m.quit:
				return
			}
		}
	}()
	return m, nil
}

// sample takes a new snapshot and accumulates its delta against the last.
func (m *iorMeter) sample() {
	m.mu.Lock()
	defer m.mu.Unlock()
	cur := iorCreateSamples(m.sub, m.subbed, 0)
	if cur == 0 {
		return
	}
	delta := iorCreateSamplesDelta(m.prev, cur, 0)
	release(m.prev)
	m.prev = cur
	if delta == 0 {
		return
	}
	defer release(delta)
	key := cfstr("IOReportChannels")
	defer cf.CFRelease(cf.CFTypeRef(key))
	arr := cf.CFArrayRef(uintptr(cf.CFDictionaryGetValue(cf.CFDictionaryRef(delta), unsafe.Pointer(uintptr(key)))))
	if arr == 0 {
		return
	}
	for i := range cf.CFArrayGetCount(arr) {
		ch := cf.CFDictionaryRef(uintptr(cf.CFArrayGetValueAtIndex(arr, i)))
		name := gostr(iorChannelName(ch))
		var field *float64
		switch {
		case name == "CPU Energy":
			field = &m.energy.CPU
		case name == "GPU Energy":
			field = &m.energy.GPU
		case strings.HasPrefix(name, "ANE"):
			field = &m.energy.ANE
		default:
			continue
		}
		j, ok := toJoules(iorSimpleIntegerValue(ch, 0), gostr(iorChannelUnitLabel(ch)))
		if !ok {
			continue
		}
		*field += j
		m.energy.Combined += j
	}
	m.samples++
}

func (m *iorMeter) stop() (*Report, error) {
	close(m.quit)
	<-m.done
	m.sample() // final partial window
	m.mu.Lock()
	defer m.mu.Unlock()
	release(m.prev)
	r := &Report{
		Duration: time.Since(m.start),
		Samples:  m.samples,
		Energy:   m.energy,
	}
	if sec := r.Duration.Seconds(); sec > 0 {
		r.Average = Power{
			CPU:      r.Energy.CPU / sec,
			GPU:      r.Energy.GPU / sec,
			ANE:      r.Energy.ANE / sec,
			Combined: r.Energy.Combined / sec,
		}
	}
	if r.Samples == 0 || r.Energy.Combined == 0 {
		return r, fmt.Errorf("powersample: IOReport returned no energy over %v — channel names changed on this macOS version?", r.Duration)
	}
	return r, nil
}

// toJoules converts a counter value with an IOReport unit label to joules.
func toJoules(v int64, unit string) (float64, bool) {
	switch unit {
	case "mJ":
		return float64(v) / 1e3, true
	case "uJ", "µJ":
		return float64(v) / 1e6, true
	case "nJ":
		return float64(v) / 1e9, true
	}
	return 0, false
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
