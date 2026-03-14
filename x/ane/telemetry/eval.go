//go:build darwin

package telemetry

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/private/appleneuralengine"
	"github.com/tmc/apple/x/ane"
)

// EvalWithStats executes the kernel and returns hardware performance stats.
// If perf stats were not enabled at compile time (PerfStatsMask=0), the
// returned stats will be zero-valued.
func EvalWithStats(m *ane.Model) (EvalStats, error) {
	perfStats := newPerfStats()
	req := appleneuralengine.ANERequestFromID(objc.ID(m.RawRequest()))
	req.SetPerfStats(&perfStats)

	if err := m.Eval(); err != nil {
		return EvalStats{}, err
	}

	var perfStatsEntries int
	func() {
		defer func() { recover() }()
		arr := req.PerfStatsArray()
		if arr != nil {
			perfStatsEntries = int(arr.Count())
			if perfStatsEntries > 0 {
				perfStats = appleneuralengine.ANEPerformanceStatsFromID(arr.ObjectAtIndex(0).GetID())
			}
		}
	}()

	if perfStatsEntries == 0 {
		func() {
			defer func() { recover() }()
			if ps := req.PerfStats(); ps != nil {
				perfStats = *ps
			}
		}()
	}

	var stats EvalStats
	stats.PerfStatsEntries = perfStatsEntries
	func() {
		defer func() { recover() }()
		stats.HWExecutionNS = perfStats.HwExecutionTime()
	}()
	func() {
		defer func() { recover() }()
		if d := perfStats.PerfCounterData(); d != nil {
			data := foundation.NSDataFromID(d.GetID())
			stats.PerfCounterData = data.GoBytes()
		}
	}()
	func() {
		defer func() { recover() }()
		if d := perfStats.PStatsRawData(); d != nil {
			data := foundation.NSDataFromID(d.GetID())
			stats.RawStatsData = data.GoBytes()
		}
	}()
	func() {
		defer func() { recover() }()
		const scanCap = 256
		for i := range scanCap {
			obj := perfStats.StringForPerfCounter(i)
			if obj == nil || obj.GetID() == 0 {
				break
			}
			name := descriptionString(obj.GetID())
			if name == "" || name == perfCounterPlaceholder {
				break
			}
			stats.PerfCounters = append(stats.PerfCounters, PerfCounter{Index: i, Name: name})
		}
		if len(stats.PerfCounters) == scanCap {
			obj := perfStats.StringForPerfCounter(scanCap)
			if obj != nil && obj.GetID() != 0 {
				name := descriptionString(obj.GetID())
				if name != "" && name != perfCounterPlaceholder {
					stats.PerfCountersTruncated = true
				}
			}
		}
	}()
	func() {
		defer func() { recover() }()
		countersObj := perfStats.PerformanceCounters()
		if countersObj == nil || countersObj.GetID() == 0 {
			return
		}
		dict := foundation.NSDictionaryFromID(countersObj.GetID())
		keys := dict.AllKeys()
		for _, key := range keys {
			idx := foundation.NSNumberFromID(key.GetID()).IntValue()
			val := foundation.NSNumberFromID(dict.ObjectForKey(key).GetID()).UnsignedLongLongValue()
			found := false
			for j := range stats.PerfCounters {
				if stats.PerfCounters[j].Index == idx {
					stats.PerfCounters[j].Value = val
					found = true
					break
				}
			}
			if !found {
				stats.PerfCounters = append(stats.PerfCounters, PerfCounter{Index: idx, Value: val})
			}
		}
	}()

	return stats, nil
}

// newPerfStats creates a valid ANEPerformanceStats via the class factory.
func newPerfStats() appleneuralengine.ANEPerformanceStats {
	cls := appleneuralengine.GetANEPerformanceStatsClass()
	obj := cls.StatsWithHardwareExecutionNS(0)
	return appleneuralengine.ANEPerformanceStatsFromID(obj.GetID())
}

// descriptionString calls -description on an ObjC object and returns the Go string.
func descriptionString(id objc.ID) string {
	if id == 0 {
		return ""
	}
	rv := objc.Send[objc.ID](id, objc.Sel("description"))
	if rv == 0 {
		return ""
	}
	return objc.Send[string](rv, objc.Sel("UTF8String"))
}
